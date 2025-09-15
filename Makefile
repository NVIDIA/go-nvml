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

include $(CURDIR)/versions.mk

PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg
GEN_BINDINGS_DIR := $(GEN_DIR)/nvml
PKG_BINDINGS_DIR := $(PKG_DIR)/nvml

DOCKER ?= docker
ifeq ($(shell uname),Darwin)
	JQ ?= $(DOCKER) run -i --rm -v "$(PWD):$(PWD)" -w "$(PWD)" backplane/jq:latest
	SED ?= $(DOCKER) run -i --rm -v "$(PWD):$(PWD)" -w "$(PWD)" alpine:latest sed
else
	JQ ?= jq
	SED ?= sed
endif

EXAMPLES := $(patsubst ./examples/%/,%,$(sort $(dir $(wildcard ./examples/*/))))
EXAMPLE_TARGETS := $(patsubst %,example-%, $(EXAMPLES))

CMDS := $(patsubst ./cmd/%/,%,$(sort $(dir $(wildcard ./cmd/*/))))
CMD_TARGETS := $(patsubst %,cmd-%, $(CMDS))

CHECK_TARGETS := validate-modules golangci-lint

MAKE_TARGETS := binary build all fmt generate test coverage check examples update-nvml-h

GENERATE_TARGETS := clean bindings test-bindings clean-bindings patch-nvml-h

TARGETS := $(MAKE_TARGETS) $(EXAMPLE_TARGETS) $(CMD_TARGETS) $(GENERATE_TARGETS) $(CHECK_TARGETS)

DOCKER_TARGETS := $(patsubst %,docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

build:
	go build $(MODULE)/pkg/...

examples: $(EXAMPLE_TARGETS)
$(EXAMPLE_TARGETS): example-%:
	go build ./examples/$(*)

check: $(CHECK_TARGETS)

# Apply go fmt to the codebase
fmt:
	go list -f '{{.Dir}}' $(MODULE)/pkg/... $(MODULE)/gen/... \
		| xargs gofmt -s -l -w

golangci-lint:
	golangci-lint run ./pkg/... ./examples/...

generate:
	go generate $(MODULE)/...

COVERAGE_FILE := coverage.out
test: build
	go test -v -coverprofile=$(COVERAGE_FILE) $(MODULE)/pkg/...

coverage: test
	cat $(COVERAGE_FILE) | grep -v "_mock.go" > $(COVERAGE_FILE).no-mocks
	go tool cover -func=$(COVERAGE_FILE).no-mocks

validate-modules:
	@echo "- Verifying that the dependencies have expected content..."
	go mod verify
	@echo "- Checking for any unused/missing packages in go.mod..."
	go mod tidy
	git diff --exit-code HEAD -- go.mod go.sum

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
		-e GOCACHE=$(PWD)/.cache/go \
		-e GOPATH=$(PWD)/.cache/gopath \
		-e GOLANGCI_LINT_CACHE=/tmp/.cache \
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
		-e GOCACHE=$(PWD)/.cache/go \
		-e GOPATH=$(PWD)/.cache/gopath \
		-e GOLANGCI_LINT_CACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE)


SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

.DEFAULT_GOAL = bindings
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
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/nvml.yml $(PKG_BINDINGS_DIR)/cgo_helpers.go $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj
	go run -tags generateapi \
		$(GEN_BINDINGS_DIR)/generateapi.go \
		--sourceDir $(PKG_BINDINGS_DIR) \
		--output $(PKG_BINDINGS_DIR)/zz_generated.api.go
	go run -tags generatecgoapi \
		$(GEN_BINDINGS_DIR)/generatecgoapi.go \
		--sourceDir $(PKG_BINDINGS_DIR) \
		--output $(PKG_BINDINGS_DIR)/zz_generated.cgoapi.go
	make fmt

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
	rm -f $(PKG_BINDINGS_DIR)/cgo_helpers.go
	rm -f $(PKG_BINDINGS_DIR)/cgo_helpers.h
	rm -f $(PKG_BINDINGS_DIR)/const.go
	rm -f $(PKG_BINDINGS_DIR)/doc.go
	rm -f $(PKG_BINDINGS_DIR)/nvml.go
	rm -f $(PKG_BINDINGS_DIR)/nvml.h
	rm -f $(PKG_BINDINGS_DIR)/types_gen.go
	rm -f $(PKG_BINDINGS_DIR)/zz_generated.api.go

# Update nvml.h from the NVIDIA CUDA redistributable JSON
update-nvml-h: CUDA_VERSION := 13.0.0
update-nvml-h: CUDA_REDIST_BASE_URL := https://developer.download.nvidia.com/compute/cuda/redist
update-nvml-h: CUDA_REDIST_JSON_URL := $(CUDA_REDIST_BASE_URL)/redistrib_$(CUDA_VERSION).json
update-nvml-h: NVML_DEV_HEADERS_INFO := $(shell \
		wget -qO - $(CUDA_REDIST_JSON_URL) | \
			$(JQ) '.cuda_nvml_dev.version + "@" + .cuda_nvml_dev."linux-x86_64".relative_path' | \
			tr -d '"' \
	)
update-nvml-h: NVML_DEV_HEADERS_COUNT := $(words $(NVML_DEV_HEADERS_INFO))
update-nvml-h: .list-nvml-packages
update-nvml-h:
	@read -p "Pick an NVML header to update ([1]-$(NVML_DEV_HEADERS_COUNT)): " idx; \
	if [ -z $${idx} ]; then idx=1; fi; \
	if ! [ $${idx} -ge 1 ] || ! [ $${idx} -le $(NVML_DEV_HEADERS_COUNT) ]; then echo "Invalid number: \"$${idx}\""; exit 1; fi; \
	NVML_DEV_HEADERS_INFO="$$(echo "$(NVML_DEV_HEADERS_INFO)" | cut -d ' ' -f$${idx})"; \
	NVML_VERSION="$$(echo "$${NVML_DEV_HEADERS_INFO}" | cut -d '@' -f1)"; \
	NVML_TAR_PATH="$$(echo "$${NVML_DEV_HEADERS_INFO}" | cut -d '@' -f2)"; \
	NVML_TAR_URL="$(CUDA_REDIST_BASE_URL)/$${NVML_TAR_PATH}"; \
	echo; \
	echo "NVML version: $${NVML_VERSION}"; \
	echo "Tar file: $${NVML_TAR_URL}"; \
	echo; \
	echo "Downloading and extracting nvml.h to $${NVML_VERSION} ..."; \
	TMP_DIR=$$(mktemp -d); \
	wget -O $${TMP_DIR}/nvml_dev.tar.xz "$${NVML_TAR_URL}" && \
	tar -xf $${TMP_DIR}/nvml_dev.tar.xz -C $${TMP_DIR} --strip-components=1 && \
	cp $${TMP_DIR}/include/nvml.h "$(GEN_BINDINGS_DIR)/nvml.h" && \
	rm -rf $${TMP_DIR} && \
	$(SED) -i -E 's#[[:blank:]]+$$##g' "$(GEN_BINDINGS_DIR)/nvml.h" && \
	$(SED) -i "1i /*** From $${NVML_TAR_URL} ***/" "$(GEN_BINDINGS_DIR)/nvml.h" && \
	$(SED) -i "1i /*** NVML VERSION: $${NVML_VERSION} ***/" "$(GEN_BINDINGS_DIR)/nvml.h" && \
	echo "Successfully updated nvml.h to $${NVML_VERSION}."

.list-nvml-packages:
	@if [ $(NVML_DEV_HEADERS_COUNT) -eq 0 ]; then \
		echo "Failed to get NVML from CUDA redistributable JSON, please try again."; \
		exit 1; \
	fi
	@echo "Found $(NVML_DEV_HEADERS_COUNT) NVML headers:"; echo
	@printf "%3s  %-8s  %-19s\n" "No." "Version" "Source"
	@idx=0; \
	for info in $(NVML_DEV_HEADERS_INFO); do \
		idx=$$((idx + 1)); \
		NVML_VERSION="$$(echo "$${info}" | cut -d '@' -f1)"; \
		NVML_TAR_PATH="$$(echo "$${info}" | cut -d '@' -f2)"; \
		printf "%3s  %-8s  %-19s\n" "$${idx}" "$${NVML_VERSION}" "CUDA Redist"; \
	done; \
	echo

# Run markdownlint (https://github.com/markdownlint/markdownlint) for README.md
# Note: Tabs are preferred for Golang code blocks
markdownlint: MDL := $(DOCKER) run --rm -v "$(PWD):$(PWD)" -w "$(PWD)" markdownlint/markdownlint:latest
markdownlint:
	@$(MDL) --rules=~no-hard-tabs,~line-length README.md
