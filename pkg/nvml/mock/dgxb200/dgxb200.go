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

package dgxb200

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/internal/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/internal/shared/gpus"
)

func New() *shared.Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpus.B200_SXM5_180GB,
		GPUCount:          8,
		DriverVersion:     "560.28.03",
		NvmlVersion:       "12.560.28.03",
		CudaDriverVersion: 12060,
	})
}

func NewDevice(index int) *shared.Device {
	return shared.NewDeviceFromConfig(gpus.B200_SXM5_180GB, index)
}

// NewServerWithGPU creates a new server with a specific B200 GPU variant
func NewServerWithGPU(gpuConfig shared.Config) *shared.Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpuConfig,
		GPUCount:          8,
		DriverVersion:     "560.28.03",
		NvmlVersion:       "12.560.28.03",
		CudaDriverVersion: 12060,
	})
}

// NewDeviceWithGPU creates a new device with a specific B200 GPU variant
func NewDeviceWithGPU(gpuConfig shared.Config, index int) *shared.Device {
	return shared.NewDeviceFromConfig(gpuConfig, index)
}

// NewServerWithGPUs creates a new server with heterogeneous GPU configurations
// Example: NewServerWithGPUs(gpus.B200_SXM5_180GB, gpus.B200_SXM5_180GB, gpus.H200_SXM5_141GB)
func NewServerWithGPUs(gpuConfigs ...shared.Config) *shared.Server {
	return shared.NewServerWithGPUs("560.28.03", "12.560.28.03", 12060, gpuConfigs...)
}

// Legacy globals for backward compatibility - expose the internal data
var (
	MIGProfiles = struct {
		GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
		ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
	}{
		GpuInstanceProfiles:     gpus.B200_SXM5_180GB.MIGProfiles.GpuInstanceProfiles,
		ComputeInstanceProfiles: gpus.B200_SXM5_180GB.MIGProfiles.ComputeInstanceProfiles,
	}

	MIGPlacements = struct {
		GpuInstancePossiblePlacements     map[int][]nvml.GpuInstancePlacement
		ComputeInstancePossiblePlacements map[int]map[int][]nvml.ComputeInstancePlacement
	}{
		GpuInstancePossiblePlacements:     gpus.B200_SXM5_180GB.MIGProfiles.GpuInstancePlacements,
		ComputeInstancePossiblePlacements: gpus.B200_SXM5_180GB.MIGProfiles.ComputeInstancePlacements,
	}
)
