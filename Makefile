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
	sed -E 's#(typedef\s+struct)\s+(nvml.*_st\*)\s+(nvml.*_t);#\1\n{\n    struct \2 handle;\n} \3;#g' $(<) > $(@)

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
		| xargs sed -i -E 's#$(SED_SEARCH_STRING).*$$#$(SED_REPLACE_STRING)#g'

.strip-nvml-h-linenumber: SED_SEARCH_STRING := // (.*) nvml/nvml.h:[0-9]+
.strip-nvml-h-linenumber: SED_REPLACE_STRING := // \1 nvml/nvml.h
.strip-nvml-h-linenumber: | .create-bindings
	grep -l -RE "$(SED_SEARCH_STRING)" pkg \
		| xargs sed -i -E 's#$(SED_SEARCH_STRING)$$#$(SED_REPLACE_STRING)#g'


test-bindings: bindings
	cd $(PKG_BINDINGS_DIR); \
		go test -v .; \
	cd -> /dev/null

clean-bindings:
	rm -rf $(PKG_BINDINGS_DIR)

# Update nvml.h from the Anaconda package repository
update-nvml-h: CONDA_PACKAGE_FILE := $(shell wget -qO - https://conda.anaconda.org/nvidia/linux-64/repodata.json | grep -oh '"cuda-nvml-dev-[^"]*"' | tr -d '"' | sort -rV | head -n 1)
update-nvml-h: NVML_VERSION := $(word 4,$(subst -, ,$(CONDA_PACKAGE_FILE)))
update-nvml-h:
	if [ -z "$(NVML_VERSION)" ]; then echo "Failed to get NVML from anaconda.org, please try again."; exit 1; fi
	@echo "NVML version: $(NVML_VERSION)"
	wget -qO - https://anaconda.org/nvidia/cuda-nvml-dev/$(NVML_VERSION)/download/linux-64/$(CONDA_PACKAGE_FILE) | \
		tar xj --strip-components=1 include/nvml.h --directory=$(GEN_BINDINGS_DIR)
	sed -i -E 's#[[:blank:]]+$$##g' $(GEN_BINDINGS_DIR)/nvml.h

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
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		$(BUILDIMAGE) \
			bash
