/*
 * Copyright (c) 2025, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dgxh100

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/internal/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/internal/shared/gpus"
)

func New() *shared.Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpus.H100_SXM5_80GB,
		GPUCount:          8,
		DriverVersion:     "550.54.15",
		NvmlVersion:       "12.550.54.15",
		CudaDriverVersion: 12040,
	})
}

func NewDevice(index int) *shared.Device {
	return shared.NewDeviceFromConfig(gpus.H100_SXM5_80GB, index)
}

// NewServerWithGPU creates a new server with a specific H100 GPU variant
func NewServerWithGPU(gpuConfig shared.Config) *shared.Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpuConfig,
		GPUCount:          8,
		DriverVersion:     "550.54.15",
		NvmlVersion:       "12.550.54.15",
		CudaDriverVersion: 12040,
	})
}

// NewDeviceWithGPU creates a new device with a specific H100 GPU variant
func NewDeviceWithGPU(gpuConfig shared.Config, index int) *shared.Device {
	return shared.NewDeviceFromConfig(gpuConfig, index)
}

// NewServerWithGPUs creates a new server with heterogeneous GPU configurations
// Example: NewServerWithGPUs(gpus.H100_SXM5_80GB, gpus.H200_SXM5_141GB)
func NewServerWithGPUs(gpuConfigs ...shared.Config) *shared.Server {
	return shared.NewServerWithGPUs("550.54.15", "12.550.54.15", 12040, gpuConfigs...)
}

// Legacy globals for backward compatibility - expose the internal data
var (
	MIGProfiles = struct {
		GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
		ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
	}{
		GpuInstanceProfiles:     gpus.H100_SXM5_80GB.MIGProfiles.GpuInstanceProfiles,
		ComputeInstanceProfiles: gpus.H100_SXM5_80GB.MIGProfiles.ComputeInstanceProfiles,
	}

	MIGPlacements = struct {
		GpuInstancePossiblePlacements     map[int][]nvml.GpuInstancePlacement
		ComputeInstancePossiblePlacements map[int]map[int][]nvml.ComputeInstancePlacement
	}{
		GpuInstancePossiblePlacements:     gpus.H100_SXM5_80GB.MIGProfiles.GpuInstancePlacements,
		ComputeInstancePossiblePlacements: gpus.H100_SXM5_80GB.MIGProfiles.ComputeInstancePlacements,
	}
)
