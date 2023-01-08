# Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

DOCKER ?= docker
GOLANG_VERSION ?= 1.15
C_FOR_GO_TAG ?= 8eeee8c3b71f9c3c90c4a73db54ed08b0bba971d
BUILDIMAGE ?= nvidia/go-nvml-devel:$(GOLANG_VERSION)-$(C_FOR_GO_TAG)

PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg
GEN_BINDINGS_DIR := $(GEN_DIR)/nvml
PKG_BINDINGS_DIR := $(PKG_DIR)/nvml

SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

EXAMPLES := $(patsubst ./examples/%/,%,$(sort $(dir $(wildcard ./examples/*/))))
EXAMPLE_TARGETS := $(patsubst %,example-%, $(EXAMPLES))

ifeq ($(shell uname),Darwin)
	SED := $(DOCKER) run -it --rm -v "$(PWD):$(PWD)" -w "$(PWD)" alpine:latest sed
else
	SED := sed
endif

TARGETS := all test clean bindings test-bindings clean-bindings patch-nvml-h examples $(EXAMPLE_TARGETS)
DOCKER_TARGETS := $(patsubst %, docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

.DEFAULT_GOAL = all

all: bindings
test: test-bindings
clean: clean-bindings

GOOS := linux

examples: $(EXAMPLE_TARGETS)
$(EXAMPLE_TARGETS): example-%: bindings
	GOOS=$(GOOS) go build ./examples/$(*)

$(PKG_BINDINGS_DIR):
	mkdir -p $(@)

patch-nvml-h: $(PKG_BINDINGS_DIR)/nvml.h
$(PKG_BINDINGS_DIR)/nvml.h: $(GEN_BINDINGS_DIR)/nvml.h | $(PKG_BINDINGS_DIR)
	$(SED) -E 's#(typedef\s+struct)\s+(nvml.*_st\*)\s+(nvml.*_t);#\1\n{\n    struct \2 handle;\n} \3;#g' $(<) > $(@)

bindings: .create-bindings .strip-autogen-comment .strip-nvml-h-linenumber

.create-bindings: $(PKG_BINDINGS_DIR)/nvml.h $(SOURCES) | $(PKG_BINDINGS_DIR)
	c-for-go -out $(PKG_DIR) $(GEN_BINDINGS_DIR)/nvml.yml
	cp $(GEN_BINDINGS_DIR)/*.go $(PKG_BINDINGS_DIR)
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj

.strip-autogen-comment: SED_SEARCH_STRING := // WARNING: This file has automatically been generated on
.strip-autogen-comment: SED_REPLACE_STRING := // WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
.strip-autogen-comment: | .create-bindings
	grep -l -R "$(SED_SEARCH_STRING)" pkg \
		| xargs $(SED) -i -E 's#$(SED_SEARCH_STRING).*$$#$(SED_REPLACE_STRING)#g'

.strip-nvml-h-linenumber: SED_SEARCH_STRING := // (.*) nvml/nvml.h:[0-9]+
.strip-nvml-h-linenumber: SED_REPLACE_STRING := // \1 nvml/nvml.h
.strip-nvml-h-linenumber: | .create-bindings
	grep -l -RE "$(SED_SEARCH_STRING)" pkg \
		| xargs $(SED) -i -E 's#$(SED_SEARCH_STRING)$$#$(SED_REPLACE_STRING)#g'

test-bindings: bindings
	cd $(PKG_BINDINGS_DIR); \
		go test -v .; \
	cd -> /dev/null

clean-bindings:
	rm -rf $(PKG_BINDINGS_DIR)
	git checkout $(PKG_BINDINGS_DIR)
	rm -rf $(PKG_BINDINGS_DIR)/nvml.h

# Update nvml.h from the Anaconda package repository
update-nvml-h: JQ ?= $(DOCKER) run -i --rm -v "$(PWD):$(PWD)" -w "$(PWD)" stedolan/jq:latest
update-nvml-h: NVML_DEV_PACKAGES_INFO := $(shell \
		wget -qO - https://api.anaconda.org/package/nvidia/cuda-nvml-dev/files | \
			$(JQ) '.[] | select(.attrs.subdir=="linux-64") | .version + "@" + .upload_time[:19] + "@" + .full_name' | \
			tr -d '"' | tr ' ' '-' | sort -rV \
	)
update-nvml-h: NVML_DEV_PACKAGES_COUNT := $(words $(NVML_DEV_PACKAGES_INFO))
update-nvml-h: .list-nvml-packages
update-nvml-h:
	@read -p "Pick an NVML package to update ([1]-$(NVML_DEV_PACKAGES_COUNT)): " idx; \
	if [ -z $${idx} ]; then idx=1; fi; \
	if ! [ $${idx} -ge 1 ] || ! [ $${idx} -le $(NVML_DEV_PACKAGES_COUNT) ]; then echo "Invalid number: \"$${idx}\""; exit 1; fi; \
	NVML_DEV_PACKAGE_INFO="$$(echo "$(NVML_DEV_PACKAGES_INFO)" | cut -d ' ' -f$${idx})"; \
	NVML_VERSION="$$(echo "$${NVML_DEV_PACKAGE_INFO}" | cut -d '@' -f1)"; \
	NVML_DEV_PACKAGE="$$(echo "$${NVML_DEV_PACKAGE_INFO}" | cut -d '@' -f3)"; \
	NVML_DEV_PACKAGE_URL="https://api.anaconda.org/download/$${NVML_DEV_PACKAGE}"; \
	echo; \
	echo "NVML version: $${NVML_VERSION}"; \
	echo "Package: $${NVML_DEV_PACKAGE}"; \
	echo; \
	echo "Updating nvml.h to $${NVML_VERSION} from $${NVML_DEV_PACKAGE_URL} ..."; \
	wget -qO - "$${NVML_DEV_PACKAGE_URL}" | \
	tar -xj --directory="$(GEN_BINDINGS_DIR)" \
		--strip-components=1 include/nvml.h && \
	$(SED) -i -E 's#[[:blank:]]+$$##g' "$(GEN_BINDINGS_DIR)/nvml.h" && \
	$(SED) -i "1i /*** From $${NVML_DEV_PACKAGE_URL} ***/" "$(GEN_BINDINGS_DIR)/nvml.h" && \
	$(SED) -i "1i /*** NVML VERSION: $${NVML_VERSION} ***/" "$(GEN_BINDINGS_DIR)/nvml.h" && \
	echo "Successfully updated nvml.h to $${NVML_VERSION}."

.list-nvml-packages:
	@if [ $(NVML_DEV_PACKAGES_COUNT) -eq 0 ]; then \
		echo "Failed to get NVML from anaconda.org, please try again."; \
		exit 1; \
	fi
	@echo "Found $(NVML_DEV_PACKAGES_COUNT) NVML packages:"; echo
	@printf "%3s  %-8s  %-19s  %-s\n" "No." "Version" "Upload Time" "Package"
	@idx=0; \
	for info in $(NVML_DEV_PACKAGES_INFO); do \
		idx=$$((idx + 1)); \
		NVML_VERSION="$$(echo "$${info}" | cut -d '@' -f1)"; \
		UPLOAD_TIME="$$(echo "$${info}" | cut -d '@' -f2)"; \
		PACKAGE="$$(echo "$${info}" | cut -d '@' -f3)"; \
		printf "%3s  %-8s  %-19s  %-s\n" "$${idx}" "$${NVML_VERSION}" "$${UPLOAD_TIME}" "$${PACKAGE}"; \
	done; \
	echo

# Generate an image for containerized builds
# Note: This image is local only
.build-image: Dockerfile
	$(DOCKER) build \
		--progress=plain \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--build-arg C_FOR_GO_TAG="$(C_FOR_GO_TAG)" \
		--tag $(BUILDIMAGE) \
		.

# A target for executing make targets in a docker container
$(DOCKER_TARGETS): docker-%: .build-image
	@echo "Running 'make $(*)' in docker container $(BUILDIMAGE)"
	$(DOCKER) run \
		--rm \
		-e JQ=jq \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			make $(*)

# A make target to set up an interactive docker environment as the current user.
# This is useful for debugging issues in the make process in the container.
.docker-shell: .build-image
	$(DOCKER) run \
		-ti \
		--rm \
		-e JQ=jq \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			bash

# A make target to set up an interactive docker environment as root.
# This is useful for debugging issues in dependencies or the container build process
.docker-root-shell: .build-image
	$(DOCKER) run \
		-ti \
		--rm \
		-e JQ=jq \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		$(BUILDIMAGE) \
			bash

# Run markdownlint (https://github.com/markdownlint/markdownlint) for README.md
# Note: Tabs are preferred for Golang code blocks
markdownlint: MDL := $(DOCKER) run --rm -v "$(PWD):$(PWD)" -w "$(PWD)" markdownlint/markdownlint:latest
markdownlint:
	@$(MDL) --rules=~no-hard-tabs,~line-length README.md
