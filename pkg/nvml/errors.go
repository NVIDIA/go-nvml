/**
# Copyright 2024 NVIDIA CORPORATION
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

package nvml

import "errors"

var (
	SUCCESS                         = error(nil)
	ERROR_UNINITIALIZED             = errors.New("ERROR_UNINITIALIZED")
	ERROR_INVALID_ARGUMENT          = errors.New("ERROR_INVALID_ARGUMENT")
	ERROR_NOT_SUPPORTED             = errors.New("ERROR_NOT_SUPPORTED")
	ERROR_NO_PERMISSION             = errors.New("ERROR_NO_PERMISSION")
	ERROR_ALREADY_INITIALIZED       = errors.New("ERROR_ALREADY_INITIALIZED")
	ERROR_NOT_FOUND                 = errors.New("ERROR_NOT_FOUND")
	ERROR_INSUFFICIENT_SIZE         = errors.New("ERROR_INSUFFICIENT_SIZE")
	ERROR_INSUFFICIENT_POWER        = errors.New("ERROR_INSUFFICIENT_POWER")
	ERROR_DRIVER_NOT_LOADED         = errors.New("ERROR_DRIVER_NOT_LOADED")
	ERROR_TIMEOUT                   = errors.New("ERROR_TIMEOUT")
	ERROR_IRQ_ISSUE                 = errors.New("ERROR_IRQ_ISSUE")
	ERROR_LIBRARY_NOT_FOUND         = errors.New("ERROR_LIBRARY_NOT_FOUND")
	ERROR_FUNCTION_NOT_FOUND        = errors.New("ERROR_FUNCTION_NOT_FOUND")
	ERROR_CORRUPTED_INFOROM         = errors.New("ERROR_CORRUPTED_INFOROM")
	ERROR_GPU_IS_LOST               = errors.New("ERROR_GPU_IS_LOST")
	ERROR_RESET_REQUIRED            = errors.New("ERROR_RESET_REQUIRED")
	ERROR_OPERATING_SYSTEM          = errors.New("ERROR_OPERATING_SYSTEM")
	ERROR_LIB_RM_VERSION_MISMATCH   = errors.New("ERROR_LIB_RM_VERSION_MISMATCH")
	ERROR_IN_USE                    = errors.New("ERROR_IN_USE")
	ERROR_MEMORY                    = errors.New("ERROR_MEMORY")
	ERROR_NO_DATA                   = errors.New("ERROR_NO_DATA")
	ERROR_VGPU_ECC_NOT_SUPPORTED    = errors.New("ERROR_VGPU_ECC_NOT_SUPPORTED")
	ERROR_INSUFFICIENT_RESOURCES    = errors.New("ERROR_INSUFFICIENT_RESOURCES")
	ERROR_FREQ_NOT_SUPPORTED        = errors.New("ERROR_FREQ_NOT_SUPPORTED")
	ERROR_ARGUMENT_VERSION_MISMATCH = errors.New("ERROR_ARGUMENT_VERSION_MISMATCH")
	ERROR_DEPRECATED                = errors.New("ERROR_DEPRECATED")
	ERROR_NOT_READY                 = errors.New("ERROR_NOT_READY")
	ERROR_GPU_NOT_FOUND             = errors.New("ERROR_GPU_NOT_FOUND")
	ERROR_INVALID_STATE             = errors.New("ERROR_INVALID_STATE")
	ERROR_UNKNOWN                   = errors.New("ERROR_UNKNOWN")
)
