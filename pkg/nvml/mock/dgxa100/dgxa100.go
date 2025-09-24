/*
 * Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
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

package dgxa100

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/gpus"
)

// Backwards compatible type aliases
type Server = shared.Server
type Device = shared.Device
type GpuInstance = shared.GpuInstance
type ComputeInstance = shared.ComputeInstance
type CudaComputeCapability = shared.CudaComputeCapability

func New() *Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpus.A100_SXM4_40GB,
		GPUCount:          8,
		DriverVersion:     "550.54.15",
		NvmlVersion:       "12.550.54.15",
		CudaDriverVersion: 12040,
	})
}

func NewDevice(index int) *Device {
	return shared.NewDeviceFromConfig(gpus.A100_SXM4_40GB, index)
}

// NewServerWithGPU creates a new server with a specific A100 GPU variant
func NewServerWithGPU(gpuConfig shared.Config) *Server {
	return shared.NewServerFromConfig(shared.ServerConfig{
		Config:            gpuConfig,
		GPUCount:          8,
		DriverVersion:     "550.54.15",
		NvmlVersion:       "12.550.54.15",
		CudaDriverVersion: 12040,
	})
}

// NewDeviceWithGPU creates a new device with a specific A100 GPU variant
func NewDeviceWithGPU(gpuConfig shared.Config, index int) *Device {
	return shared.NewDeviceFromConfig(gpuConfig, index)
}

// Legacy globals for backward compatibility - expose the internal data
var (
	MIGProfiles = struct {
		GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
		ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
	}{
		GpuInstanceProfiles:     gpus.A100_SXM4_40GB.MIGProfiles.GpuInstanceProfiles,
		ComputeInstanceProfiles: gpus.A100_SXM4_40GB.MIGProfiles.ComputeInstanceProfiles,
	}

	MIGPlacements = struct {
		GpuInstancePossiblePlacements     map[int][]nvml.GpuInstancePlacement
		ComputeInstancePossiblePlacements map[int]map[int][]nvml.ComputeInstancePlacement
	}{
		GpuInstancePossiblePlacements:     gpus.A100_SXM4_40GB.MIGProfiles.GpuInstancePlacements,
		ComputeInstancePossiblePlacements: gpus.A100_SXM4_40GB.MIGProfiles.ComputeInstancePlacements,
	}
)
