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
GOLANG_VERSION ?= 1.14.4
C_FOR_GO_TAG ?= 8eeee8c3b71f9c3c90c4a73db54ed08b0bba971d
BUILDIMAGE ?= nvidia/go-nvml-devel:$(GOLANG_VERSION)-$(C_FOR_GO_TAG)

PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg
GEN_BINDINGS_DIR := $(GEN_DIR)/nvml
PKG_BINDINGS_DIR := $(PKG_DIR)/nvml

SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

TARGETS := all test clean bindings test-bindings clean-bindings patch-nvml-h
DOCKER_TARGETS := $(patsubst %, docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

.DEFAULT_GOAL = all

all: bindings
test: test-bindings
clean: clean-bindings

$(PKG_BINDINGS_DIR):
	mkdir -p $(@)

patch-nvml-h: $(PKG_BINDINGS_DIR)/nvml.h
$(PKG_BINDINGS_DIR)/nvml.h: $(GEN_BINDINGS_DIR)/nvml.h | $(PKG_BINDINGS_DIR)
	sed -E 's#(typedef\s+struct)\s+(nvml.*_st\*)\s+(nvml.*_t);#\1\n{\n    struct \2 handle;\n} \3;#g' $(<) > $(@)

bindings: .create-bindings .strip-autogen-comment

.create-bindings: $(PKG_BINDINGS_DIR)/nvml.h $(SOURCES) | $(PKG_BINDINGS_DIR)
	c-for-go -out $(PKG_DIR) $(GEN_BINDINGS_DIR)/nvml.yml
	cp $(GEN_BINDINGS_DIR)/*.go $(PKG_BINDINGS_DIR)
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj

SED_SEARCH_STRING := WARNING: This file has automatically been generated on
SED_REPLACE_STRING := WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
.strip-autogen-comment: | .create-bindings
	grep -l -R "// WARNING: This file has automatically been generated on" pkg \
		| xargs sed -i -E 's/ $(SED_SEARCH_STRING).*$$/ $(SED_REPLACE_STRING)/g'

test-bindings: bindings
	cd $(PKG_BINDINGS_DIR); \
		go test -v .; \
	cd -> /dev/null

clean-bindings:
	rm -rf $(PKG_BINDINGS_DIR)

# Update nvml.h from the specied CUDA_VERSION development image
update-nvml-h:
	if [[ $(CUDA_VERSION) == "" ]]; then echo "define CUDA_VERSION to update"; exit 1; fi
	$(DOCKER) run \
		--rm \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		nvidia/cuda:$(CUDA_VERSION)-devel \
			cp /usr/local/cuda-$(CUDA_VERSION)/targets/x86_64-linux/include/nvml.h $(GEN_BINDINGS_DIR)

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