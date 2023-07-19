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

PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg
GEN_BINDINGS_DIR := $(GEN_DIR)/nvml
PKG_BINDINGS_DIR := $(PKG_DIR)/nvml

ifeq ($(shell uname),Darwin)
	SED := $(DOCKER) run -it --rm -v "$(PWD):$(PWD)" -w "$(PWD)" alpine:latest sed
else
	SED := sed
endif

MODULE := github.com/NVIDIA/go-nvml/pkg

DOCKER ?= docker

GOLANG_VERSION ?= 1.18.10
C_FOR_GO_TAG ?= 8eeee8c3b71f9c3c90c4a73db54ed08b0bba971d

ifeq ($(IMAGE),)
REGISTRY ?= nvidia
IMAGE=$(REGISTRY)/go-nvml
endif
IMAGE_TAG ?= $(GOLANG_VERSION)-$(C_FOR_GO_TAG)
BUILDIMAGE ?= $(IMAGE):$(IMAGE_TAG)-devel

EXAMPLES := $(patsubst ./examples/%/,%,$(sort $(dir $(wildcard ./examples/*/))))
EXAMPLE_TARGETS := $(patsubst %,example-%, $(EXAMPLES))

CMDS := $(patsubst ./cmd/%/,%,$(sort $(dir $(wildcard ./cmd/*/))))
CMD_TARGETS := $(patsubst %,cmd-%, $(CMDS))

CHECK_TARGETS := assert-fmt vet lint

MAKE_TARGETS := binary build all fmt generate test coverage check examples

GENERATE_TARGETS := clean bindings test-bindings clean-bindings patch-nvml-h

TARGETS := $(MAKE_TARGETS) $(EXAMPLE_TARGETS) $(CMD_TARGETS) $(GENERATE_TARGETS) $(CHECK_TARGETS)

DOCKER_TARGETS := $(patsubst %,docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

GOOS := linux

build:
	GOOS=$(GOOS) go build $(MODULE)/...

examples: $(EXAMPLE_TARGETS)
$(EXAMPLE_TARGETS): example-%:
	GOOS=$(GOOS) go build ./examples/$(*)

check: $(CHECK_TARGETS)

# Apply go fmt to the codebase
fmt:
	go list -f '{{.Dir}}' $(MODULE)/... \
		| xargs gofmt -s -l -w

assert-fmt:
	go list -f '{{.Dir}}' $(MODULE)/... \
		| xargs gofmt -s -l > fmt.out
	@if [ -s fmt.out ]; then \
		echo "\nERROR: The following files are not formatted:\n"; \
		cat fmt.out; \
		rm fmt.out; \
		exit 1; \
	else \
		rm fmt.out; \
	fi

generate:
	go generate $(MODULE)/...

lint:
	# We use `go list -f '{{.Dir}}' $(MODULE)/...` to skip the `vendor` folder.
	# One we have fixed the linting issues, we whould add -set_exit_status
	go list -f '{{.Dir}}' $(MODULE)/... | grep -v pkg/nvml | xargs golint

vet:
	go vet $(MODULE)/...

COVERAGE_FILE := coverage.out
test: build
	# go test -v -coverprofile=$(COVERAGE_FILE) $(MODULE)/...
	@echo "TODO: Skipping tests for now"

coverage: test
	cat $(COVERAGE_FILE) | grep -v "_mock.go" > $(COVERAGE_FILE).no-mocks
	go tool cover -func=$(COVERAGE_FILE).no-mocks

# Generate an image for containerized builds
# Note: This image is local only
.PHONY: .build-image .pull-build-image .push-build-image
.build-image: docker/Dockerfile.devel
	if [ "$(SKIP_IMAGE_BUILD)" = "" ]; then \
		$(DOCKER) build \
			--progress=plain \
			--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
			--build-arg C_FOR_GO_TAG="$(C_FOR_GO_TAG)" \
			--tag $(BUILDIMAGE) \
			-f $(^) \
			docker; \
	fi

.pull-build-image:
	$(DOCKER) pull $(BUILDIMAGE)

.push-build-image:
	$(DOCKER) push $(BUILDIMAGE)

# A target for executing make targets in a docker container
$(DOCKER_TARGETS): docker-%: .build-image
	@echo "Running 'make $(*)' in docker container $(BUILDIMAGE)"
	$(DOCKER) run \
		--rm \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			make $(*)

# Start an interactive shell using the development image.
PHONY: .shell
.shell:
	$(DOCKER) run \
		--rm \
		-ti \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE)


SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

.DEFAULT_GOAL = bindings

# In order to build the packages we need to patch the nvml.h file
build: bindings

test: test-bindings
clean: clean-bindings

$(PKG_BINDINGS_DIR):
	mkdir -p $(@)

patch-nvml-h: $(PKG_BINDINGS_DIR)/nvml.h
$(PKG_BINDINGS_DIR)/nvml.h: $(GEN_BINDINGS_DIR)/nvml.h | $(PKG_BINDINGS_DIR)
	cp $(<) $(@)
	$(SED) -i -E 's#(typedef\s+struct)\s+(nvml.*_st\*)\s+(nvml.*_t);#\1\n{\n    struct \2 handle;\n} \3;#g' $(@)
	spatch --in-place --very-quiet --sp-file $(GEN_BINDINGS_DIR)/anonymous_structs.cocci $(@) > /dev/null

bindings: .create-bindings .strip-autogen-comment .strip-nvml-h-linenumber
.create-bindings: $(PKG_BINDINGS_DIR)/nvml.h $(SOURCES) | $(PKG_BINDINGS_DIR)
	cp $(GEN_BINDINGS_DIR)/nvml.yml $(PKG_BINDINGS_DIR)
	c-for-go -out $(PKG_DIR) $(PKG_BINDINGS_DIR)/nvml.yml
	cp $(GEN_BINDINGS_DIR)/*.go $(PKG_BINDINGS_DIR)
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/nvml.yml $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj

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

# Run markdownlint (https://github.com/markdownlint/markdownlint) for README.md
# Note: Tabs are preferred for Golang code blocks
markdownlint: MDL := $(DOCKER) run --rm -v "$(PWD):$(PWD)" -w "$(PWD)" markdownlint/markdownlint:latest
markdownlint:
	@$(MDL) --rules=~no-hard-tabs,~line-length README.md

