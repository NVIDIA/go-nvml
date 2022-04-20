/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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
**/

package cuda

import (
	"fmt"

	"github.com/NVIDIA/go-nvml/pkg/dl"
)

import "C"

const (
	cudaLibraryName      = "libcuda.so.1"
	cudaLibraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

// cuda stores a reference the cuda dynamic library
var cuda *dl.DynamicLibrary

// Load loads the CUDA shared library
func Load() error {
	lib := dl.New(cudaLibraryName, cudaLibraryLoadFlags)
	if lib == nil {
		return fmt.Errorf("error instantiating DynamicLibrary for %v", cudaLibraryName)
	}

	err := lib.Open()
	if err != nil {
		return fmt.Errorf("error opening %v: %v", cudaLibraryName, err)
	}

	cuda = lib
	updateVersionedSymbols()
	return nil
}

// Close closes the shared library
func Close() error {
	if cuda == nil {
		return nil
	}
	err := cuda.Close()
	if err != nil {
		return fmt.Errorf("error closing %v: %v", cudaLibraryName, err)
	}
	cuda = nil
	return nil
}

// udateVersionedSymbols updates the versioned library symbols if required
func updateVersionedSymbols() {
}
