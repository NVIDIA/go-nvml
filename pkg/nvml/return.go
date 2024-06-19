// Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nvml

import (
	"fmt"
)

// nvml.ErrorString()
func (l *library) ErrorString(r Return) string {
	return r.String()
}

// String returns the string representation of a ret.error()urn.
func (r Return) String() string {
	return errorStringFunc(r)
}

// error returns the associated golang error.
func (r Return) error() error {
	switch r {
	case nvmlSUCCESS:
		return SUCCESS
	case nvmlERROR_UNINITIALIZED:
		return ERROR_UNINITIALIZED
	case nvmlERROR_INVALID_ARGUMENT:
		return ERROR_INVALID_ARGUMENT
	case nvmlERROR_NOT_SUPPORTED:
		return ERROR_NOT_SUPPORTED
	case nvmlERROR_NO_PERMISSION:
		return ERROR_NO_PERMISSION
	case nvmlERROR_ALREADY_INITIALIZED:
		return ERROR_ALREADY_INITIALIZED
	case nvmlERROR_NOT_FOUND:
		return ERROR_NOT_FOUND
	case nvmlERROR_INSUFFICIENT_SIZE:
		return ERROR_INSUFFICIENT_SIZE
	case nvmlERROR_INSUFFICIENT_POWER:
		return ERROR_INSUFFICIENT_POWER
	case nvmlERROR_DRIVER_NOT_LOADED:
		return ERROR_DRIVER_NOT_LOADED
	case nvmlERROR_TIMEOUT:
		return ERROR_TIMEOUT
	case nvmlERROR_IRQ_ISSUE:
		return ERROR_IRQ_ISSUE
	case nvmlERROR_LIBRARY_NOT_FOUND:
		return ERROR_LIBRARY_NOT_FOUND
	case nvmlERROR_FUNCTION_NOT_FOUND:
		return ERROR_FUNCTION_NOT_FOUND
	case nvmlERROR_CORRUPTED_INFOROM:
		return ERROR_CORRUPTED_INFOROM
	case nvmlERROR_GPU_IS_LOST:
		return ERROR_GPU_IS_LOST
	case nvmlERROR_RESET_REQUIRED:
		return ERROR_RESET_REQUIRED
	case nvmlERROR_OPERATING_SYSTEM:
		return ERROR_OPERATING_SYSTEM
	case nvmlERROR_LIB_RM_VERSION_MISMATCH:
		return ERROR_LIB_RM_VERSION_MISMATCH
	case nvmlERROR_IN_USE:
		return ERROR_IN_USE
	case nvmlERROR_MEMORY:
		return ERROR_MEMORY
	case nvmlERROR_NO_DATA:
		return ERROR_NO_DATA
	case nvmlERROR_VGPU_ECC_NOT_SUPPORTED:
		return ERROR_VGPU_ECC_NOT_SUPPORTED
	case nvmlERROR_INSUFFICIENT_RESOURCES:
		return ERROR_INSUFFICIENT_RESOURCES
	case nvmlERROR_FREQ_NOT_SUPPORTED:
		return ERROR_FREQ_NOT_SUPPORTED
	case nvmlERROR_ARGUMENT_VERSION_MISMATCH:
		return ERROR_ARGUMENT_VERSION_MISMATCH
	case nvmlERROR_DEPRECATED:
		return ERROR_DEPRECATED
	case nvmlERROR_UNKNOWN:
		return ERROR_UNKNOWN
	default:
		return fmt.Errorf("%w: unknown ret.error()urn value: %d", ERROR_UNKNOWN, r)
	}
}

// Assigned to nvml.ErrorString if the system nvml library is in use.
var errorStringFunc = defaultErrorStringFunc

// defaultErrorStringFunc provides a basic nvmlErrorString implementation.
// This allows the nvml.ErrorString function to be used even if the NVML library
// is not loaded.
var defaultErrorStringFunc = func(r Return) string {
	return fmt.Sprintf("%s", r.error())
}
