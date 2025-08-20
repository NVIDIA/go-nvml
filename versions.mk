# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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


GIT_COMMIT ?= $(shell git describe --match="" --dirty --long --always --abbrev=40 2> /dev/null || echo "")
GIT_TAG ?= $(patsubst v%,%,$(shell git describe --tags 2>/dev/null))

MODULE := github.com/NVIDIA/go-nvml
VERSION ?= $(GIT_TAG)

GOLANG_VERSION ?= 1.24.5
C_FOR_GO_TAG ?= 8eeee8c3b71f9c3c90c4a73db54ed08b0bba971d

ifeq ($(IMAGE),)
REGISTRY ?= nvidia
IMAGE=$(REGISTRY)/go-nvml
endif
IMAGE_TAG ?= $(GOLANG_VERSION)-$(C_FOR_GO_TAG)
BUILDIMAGE ?= $(IMAGE):$(IMAGE_TAG)-devel
