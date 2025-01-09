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

// nvml.SystemGetDriverVersion()
func (l *library) SystemGetDriverVersion() (string, error) {
	Version := make([]byte, SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	ret := nvmlSystemGetDriverVersion(&Version[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret.error()
}

// nvml.SystemGetNVMLVersion()
func (l *library) SystemGetNVMLVersion() (string, error) {
	Version := make([]byte, SYSTEM_NVML_VERSION_BUFFER_SIZE)
	ret := nvmlSystemGetNVMLVersion(&Version[0], SYSTEM_NVML_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret.error()
}

// nvml.SystemGetCudaDriverVersion()
func (l *library) SystemGetCudaDriverVersion() (int, error) {
	var CudaDriverVersion int32
	ret := nvmlSystemGetCudaDriverVersion(&CudaDriverVersion)
	return int(CudaDriverVersion), ret.error()
}

// nvml.SystemGetCudaDriverVersion_v2()
func (l *library) SystemGetCudaDriverVersion_v2() (int, error) {
	var CudaDriverVersion int32
	ret := nvmlSystemGetCudaDriverVersion_v2(&CudaDriverVersion)
	return int(CudaDriverVersion), ret.error()
}

// nvml.SystemGetProcessName()
func (l *library) SystemGetProcessName(pid int) (string, error) {
	name := make([]byte, SYSTEM_PROCESS_NAME_BUFFER_SIZE)
	ret := nvmlSystemGetProcessName(uint32(pid), &name[0], SYSTEM_PROCESS_NAME_BUFFER_SIZE)
	return string(name[:clen(name)]), ret.error()
}

// nvml.SystemGetHicVersion()
func (l *library) SystemGetHicVersion() ([]HwbcEntry, error) {
	var hwbcCount uint32 = 1 // Will be reduced upon returning
	for {
		hwbcEntries := make([]HwbcEntry, hwbcCount)
		ret := nvmlSystemGetHicVersion(&hwbcCount, &hwbcEntries[0])
		if ret == nvmlSUCCESS {
			return hwbcEntries[:hwbcCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		hwbcCount *= 2
	}
}

// nvml.SystemGetTopologyGpuSet()
func (l *library) SystemGetTopologyGpuSet(cpuNumber int) ([]Device, error) {
	var count uint32
	ret := nvmlSystemGetTopologyGpuSet(uint32(cpuNumber), &count, nil)
	if ret != nvmlSUCCESS {
		return nil, ret.error()
	}
	if count == 0 {
		return []Device{}, ret.error()
	}
	deviceArray := make([]nvmlDevice, count)
	ret = nvmlSystemGetTopologyGpuSet(uint32(cpuNumber), &count, &deviceArray[0])
	return convertSlice[nvmlDevice, Device](deviceArray), ret.error()
}

// nvml.SystemGetConfComputeCapabilities()
func (l *library) SystemGetConfComputeCapabilities() (ConfComputeSystemCaps, error) {
	var capabilities ConfComputeSystemCaps
	ret := nvmlSystemGetConfComputeCapabilities(&capabilities)
	return capabilities, ret.error()
}

// nvml.SystemGetConfComputeState()
func SystemGetConfComputeState() (ConfComputeSystemState, error) {
	var state ConfComputeSystemState
	ret := nvmlSystemGetConfComputeState(&state)
	return state, ret.error()
}

// nvml.SystemGetConfComputeGpusReadyState()
func SystemGetConfComputeGpusReadyState() (uint32, error) {
	var isAcceptingWork uint32
	ret := nvmlSystemGetConfComputeGpusReadyState(&isAcceptingWork)
	return isAcceptingWork, ret.error()
}

// nvml.SystemSetConfComputeGpusReadyState()
func SystemSetConfComputeGpusReadyState(isAcceptingWork uint32) error {
	return nvmlSystemSetConfComputeGpusReadyState(isAcceptingWork).error()
}

// nvml.SystemSetNvlinkBwMode()
func SystemSetNvlinkBwMode(nvlinkBwMode uint32) error {
	return nvmlSystemSetNvlinkBwMode(nvlinkBwMode).error()
}

// nvml.SystemGetNvlinkBwMode()
func SystemGetNvlinkBwMode() (uint32, error) {
	var nvlinkBwMode uint32
	ret := nvmlSystemGetNvlinkBwMode(&nvlinkBwMode)
	return nvlinkBwMode, ret.error()
}

// nvml.SystemGetConfComputeKeyRotationThresholdInfo()
func (l *library) SystemGetConfComputeKeyRotationThresholdInfo() (ConfComputeGetKeyRotationThresholdInfo, error) {
	var keyRotationThresholdInfo ConfComputeGetKeyRotationThresholdInfo
	keyRotationThresholdInfo.Version = STRUCT_VERSION(keyRotationThresholdInfo, 1)
	ret := nvmlSystemGetConfComputeKeyRotationThresholdInfo(&keyRotationThresholdInfo)
	return keyRotationThresholdInfo, ret.error()
}

// nvml.SystemGetConfComputeSettings()
func (l *library) SystemGetConfComputeSettings() (SystemConfComputeSettings, error) {
	var settings SystemConfComputeSettings
	settings.Version = STRUCT_VERSION(settings, 1)
	ret := nvmlSystemGetConfComputeSettings(&settings)
	return settings, ret.error()
}

// nvml.SystemSetConfComputeKeyRotationThresholdInfo()
func (l *library) SystemSetConfComputeKeyRotationThresholdInfo(keyRotationThresholdInfo ConfComputeSetKeyRotationThresholdInfo) error {
	return nvmlSystemSetConfComputeKeyRotationThresholdInfo(&keyRotationThresholdInfo).error()
}
