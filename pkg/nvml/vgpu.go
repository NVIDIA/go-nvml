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
	"unsafe"
)

// nvml.VgpuMetadata
type VgpuMetadata struct {
	nvmlVgpuMetadata
	OpaqueData []byte
}

// nvml.VgpuPgpuMetadata
type VgpuPgpuMetadata struct {
	nvmlVgpuPgpuMetadata
	OpaqueData []byte
}

// nvml.VgpuTypeGetClass()
func (l *library) VgpuTypeGetClass(vgpuTypeId VgpuTypeId) (string, error) {
	return vgpuTypeId.GetClass()
}

func (vgpuTypeId nvmlVgpuTypeId) GetClass() (string, error) {
	var size uint32 = DEVICE_NAME_BUFFER_SIZE
	vgpuTypeClass := make([]byte, DEVICE_NAME_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetClass(vgpuTypeId, &vgpuTypeClass[0], &size)
	return string(vgpuTypeClass[:clen(vgpuTypeClass)]), ret.error()
}

// nvml.VgpuTypeGetName()
func (l *library) VgpuTypeGetName(vgpuTypeId VgpuTypeId) (string, error) {
	return vgpuTypeId.GetName()
}

func (vgpuTypeId nvmlVgpuTypeId) GetName() (string, error) {
	var size uint32 = DEVICE_NAME_BUFFER_SIZE
	vgpuTypeName := make([]byte, DEVICE_NAME_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetName(vgpuTypeId, &vgpuTypeName[0], &size)
	return string(vgpuTypeName[:clen(vgpuTypeName)]), ret.error()
}

// nvml.VgpuTypeGetGpuInstanceProfileId()
func (l *library) VgpuTypeGetGpuInstanceProfileId(vgpuTypeId VgpuTypeId) (uint32, error) {
	return vgpuTypeId.GetGpuInstanceProfileId()
}

func (vgpuTypeId nvmlVgpuTypeId) GetGpuInstanceProfileId() (uint32, error) {
	var size uint32
	ret := nvmlVgpuTypeGetGpuInstanceProfileId(vgpuTypeId, &size)
	return size, ret.error()
}

// nvml.VgpuTypeGetDeviceID()
func (l *library) VgpuTypeGetDeviceID(vgpuTypeId VgpuTypeId) (uint64, uint64, error) {
	return vgpuTypeId.GetDeviceID()
}

func (vgpuTypeId nvmlVgpuTypeId) GetDeviceID() (uint64, uint64, error) {
	var deviceID, subsystemID uint64
	ret := nvmlVgpuTypeGetDeviceID(vgpuTypeId, &deviceID, &subsystemID)
	return deviceID, subsystemID, ret.error()
}

// nvml.VgpuTypeGetFramebufferSize()
func (l *library) VgpuTypeGetFramebufferSize(vgpuTypeId VgpuTypeId) (uint64, error) {
	return vgpuTypeId.GetFramebufferSize()
}

func (vgpuTypeId nvmlVgpuTypeId) GetFramebufferSize() (uint64, error) {
	var fbSize uint64
	ret := nvmlVgpuTypeGetFramebufferSize(vgpuTypeId, &fbSize)
	return fbSize, ret.error()
}

// nvml.VgpuTypeGetNumDisplayHeads()
func (l *library) VgpuTypeGetNumDisplayHeads(vgpuTypeId VgpuTypeId) (int, error) {
	return vgpuTypeId.GetNumDisplayHeads()
}

func (vgpuTypeId nvmlVgpuTypeId) GetNumDisplayHeads() (int, error) {
	var numDisplayHeads uint32
	ret := nvmlVgpuTypeGetNumDisplayHeads(vgpuTypeId, &numDisplayHeads)
	return int(numDisplayHeads), ret.error()
}

// nvml.VgpuTypeGetResolution()
func (l *library) VgpuTypeGetResolution(vgpuTypeId VgpuTypeId, displayIndex int) (uint32, uint32, error) {
	return vgpuTypeId.GetResolution(displayIndex)
}

func (vgpuTypeId nvmlVgpuTypeId) GetResolution(displayIndex int) (uint32, uint32, error) {
	var xdim, ydim uint32
	ret := nvmlVgpuTypeGetResolution(vgpuTypeId, uint32(displayIndex), &xdim, &ydim)
	return xdim, ydim, ret.error()
}

// nvml.VgpuTypeGetLicense()
func (l *library) VgpuTypeGetLicense(vgpuTypeId VgpuTypeId) (string, error) {
	return vgpuTypeId.GetLicense()
}

func (vgpuTypeId nvmlVgpuTypeId) GetLicense() (string, error) {
	vgpuTypeLicenseString := make([]byte, GRID_LICENSE_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetLicense(vgpuTypeId, &vgpuTypeLicenseString[0], GRID_LICENSE_BUFFER_SIZE)
	return string(vgpuTypeLicenseString[:clen(vgpuTypeLicenseString)]), ret.error()
}

// nvml.VgpuTypeGetFrameRateLimit()
func (l *library) VgpuTypeGetFrameRateLimit(vgpuTypeId VgpuTypeId) (uint32, error) {
	return vgpuTypeId.GetFrameRateLimit()
}

func (vgpuTypeId nvmlVgpuTypeId) GetFrameRateLimit() (uint32, error) {
	var frameRateLimit uint32
	ret := nvmlVgpuTypeGetFrameRateLimit(vgpuTypeId, &frameRateLimit)
	return frameRateLimit, ret.error()
}

// nvml.VgpuTypeGetMaxInstances()
func (l *library) VgpuTypeGetMaxInstances(device Device, vgpuTypeId VgpuTypeId) (int, error) {
	return vgpuTypeId.GetMaxInstances(device)
}

func (device nvmlDevice) VgpuTypeGetMaxInstances(vgpuTypeId VgpuTypeId) (int, error) {
	return vgpuTypeId.GetMaxInstances(device)
}

func (vgpuTypeId nvmlVgpuTypeId) GetMaxInstances(device Device) (int, error) {
	var vgpuInstanceCount uint32
	ret := nvmlVgpuTypeGetMaxInstances(nvmlDeviceHandle(device), vgpuTypeId, &vgpuInstanceCount)
	return int(vgpuInstanceCount), ret.error()
}

// nvml.VgpuTypeGetMaxInstancesPerVm()
func (l *library) VgpuTypeGetMaxInstancesPerVm(vgpuTypeId VgpuTypeId) (int, error) {
	return vgpuTypeId.GetMaxInstancesPerVm()
}

func (vgpuTypeId nvmlVgpuTypeId) GetMaxInstancesPerVm() (int, error) {
	var vgpuInstanceCountPerVm uint32
	ret := nvmlVgpuTypeGetMaxInstancesPerVm(vgpuTypeId, &vgpuInstanceCountPerVm)
	return int(vgpuInstanceCountPerVm), ret.error()
}

// nvml.VgpuInstanceGetVmID()
func (l *library) VgpuInstanceGetVmID(vgpuInstance VgpuInstance) (string, VgpuVmIdType, error) {
	return vgpuInstance.GetVmID()
}

func (vgpuInstance nvmlVgpuInstance) GetVmID() (string, VgpuVmIdType, error) {
	var vmIdType VgpuVmIdType
	vmId := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetVmID(vgpuInstance, &vmId[0], DEVICE_UUID_BUFFER_SIZE, &vmIdType)
	return string(vmId[:clen(vmId)]), vmIdType, ret.error()
}

// nvml.VgpuInstanceGetUUID()
func (l *library) VgpuInstanceGetUUID(vgpuInstance VgpuInstance) (string, error) {
	return vgpuInstance.GetUUID()
}

func (vgpuInstance nvmlVgpuInstance) GetUUID() (string, error) {
	uuid := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetUUID(vgpuInstance, &uuid[0], DEVICE_UUID_BUFFER_SIZE)
	return string(uuid[:clen(uuid)]), ret.error()
}

// nvml.VgpuInstanceGetVmDriverVersion()
func (l *library) VgpuInstanceGetVmDriverVersion(vgpuInstance VgpuInstance) (string, error) {
	return vgpuInstance.GetVmDriverVersion()
}

func (vgpuInstance nvmlVgpuInstance) GetVmDriverVersion() (string, error) {
	version := make([]byte, SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetVmDriverVersion(vgpuInstance, &version[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	return string(version[:clen(version)]), ret.error()
}

// nvml.VgpuInstanceGetFbUsage()
func (l *library) VgpuInstanceGetFbUsage(vgpuInstance VgpuInstance) (uint64, error) {
	return vgpuInstance.GetFbUsage()
}

func (vgpuInstance nvmlVgpuInstance) GetFbUsage() (uint64, error) {
	var fbUsage uint64
	ret := nvmlVgpuInstanceGetFbUsage(vgpuInstance, &fbUsage)
	return fbUsage, ret.error()
}

// nvml.VgpuInstanceGetLicenseInfo()
func (l *library) VgpuInstanceGetLicenseInfo(vgpuInstance VgpuInstance) (VgpuLicenseInfo, error) {
	return vgpuInstance.GetLicenseInfo()
}

func (vgpuInstance nvmlVgpuInstance) GetLicenseInfo() (VgpuLicenseInfo, error) {
	var licenseInfo VgpuLicenseInfo
	ret := nvmlVgpuInstanceGetLicenseInfo(vgpuInstance, &licenseInfo)
	return licenseInfo, ret.error()
}

// nvml.VgpuInstanceGetLicenseStatus()
func (l *library) VgpuInstanceGetLicenseStatus(vgpuInstance VgpuInstance) (int, error) {
	return vgpuInstance.GetLicenseStatus()
}

func (vgpuInstance nvmlVgpuInstance) GetLicenseStatus() (int, error) {
	var licensed uint32
	ret := nvmlVgpuInstanceGetLicenseStatus(vgpuInstance, &licensed)
	return int(licensed), ret.error()
}

// nvml.VgpuInstanceGetType()
func (l *library) VgpuInstanceGetType(vgpuInstance VgpuInstance) (VgpuTypeId, error) {
	return vgpuInstance.GetType()
}

func (vgpuInstance nvmlVgpuInstance) GetType() (VgpuTypeId, error) {
	var vgpuTypeId nvmlVgpuTypeId
	ret := nvmlVgpuInstanceGetType(vgpuInstance, &vgpuTypeId)
	return vgpuTypeId, ret.error()
}

// nvml.VgpuInstanceGetFrameRateLimit()
func (l *library) VgpuInstanceGetFrameRateLimit(vgpuInstance VgpuInstance) (uint32, error) {
	return vgpuInstance.GetFrameRateLimit()
}

func (vgpuInstance nvmlVgpuInstance) GetFrameRateLimit() (uint32, error) {
	var frameRateLimit uint32
	ret := nvmlVgpuInstanceGetFrameRateLimit(vgpuInstance, &frameRateLimit)
	return frameRateLimit, ret.error()
}

// nvml.VgpuInstanceGetEccMode()
func (l *library) VgpuInstanceGetEccMode(vgpuInstance VgpuInstance) (EnableState, error) {
	return vgpuInstance.GetEccMode()
}

func (vgpuInstance nvmlVgpuInstance) GetEccMode() (EnableState, error) {
	var eccMode EnableState
	ret := nvmlVgpuInstanceGetEccMode(vgpuInstance, &eccMode)
	return eccMode, ret.error()
}

// nvml.VgpuInstanceGetEncoderCapacity()
func (l *library) VgpuInstanceGetEncoderCapacity(vgpuInstance VgpuInstance) (int, error) {
	return vgpuInstance.GetEncoderCapacity()
}

func (vgpuInstance nvmlVgpuInstance) GetEncoderCapacity() (int, error) {
	var encoderCapacity uint32
	ret := nvmlVgpuInstanceGetEncoderCapacity(vgpuInstance, &encoderCapacity)
	return int(encoderCapacity), ret.error()
}

// nvml.VgpuInstanceSetEncoderCapacity()
func (l *library) VgpuInstanceSetEncoderCapacity(vgpuInstance VgpuInstance, encoderCapacity int) error {
	return vgpuInstance.SetEncoderCapacity(encoderCapacity)
}

func (vgpuInstance nvmlVgpuInstance) SetEncoderCapacity(encoderCapacity int) error {
	return nvmlVgpuInstanceSetEncoderCapacity(vgpuInstance, uint32(encoderCapacity)).error()
}

// nvml.VgpuInstanceGetEncoderStats()
func (l *library) VgpuInstanceGetEncoderStats(vgpuInstance VgpuInstance) (int, uint32, uint32, error) {
	return vgpuInstance.GetEncoderStats()
}

func (vgpuInstance nvmlVgpuInstance) GetEncoderStats() (int, uint32, uint32, error) {
	var sessionCount, averageFps, averageLatency uint32
	ret := nvmlVgpuInstanceGetEncoderStats(vgpuInstance, &sessionCount, &averageFps, &averageLatency)
	return int(sessionCount), averageFps, averageLatency, ret.error()
}

// nvml.VgpuInstanceGetEncoderSessions()
func (l *library) VgpuInstanceGetEncoderSessions(vgpuInstance VgpuInstance) (int, EncoderSessionInfo, error) {
	return vgpuInstance.GetEncoderSessions()
}

func (vgpuInstance nvmlVgpuInstance) GetEncoderSessions() (int, EncoderSessionInfo, error) {
	var sessionCount uint32
	var sessionInfo EncoderSessionInfo
	ret := nvmlVgpuInstanceGetEncoderSessions(vgpuInstance, &sessionCount, &sessionInfo)
	return int(sessionCount), sessionInfo, ret.error()
}

// nvml.VgpuInstanceGetFBCStats()
func (l *library) VgpuInstanceGetFBCStats(vgpuInstance VgpuInstance) (FBCStats, error) {
	return vgpuInstance.GetFBCStats()
}

func (vgpuInstance nvmlVgpuInstance) GetFBCStats() (FBCStats, error) {
	var fbcStats FBCStats
	ret := nvmlVgpuInstanceGetFBCStats(vgpuInstance, &fbcStats)
	return fbcStats, ret.error()
}

// nvml.VgpuInstanceGetFBCSessions()
func (l *library) VgpuInstanceGetFBCSessions(vgpuInstance VgpuInstance) (int, FBCSessionInfo, error) {
	return vgpuInstance.GetFBCSessions()
}

func (vgpuInstance nvmlVgpuInstance) GetFBCSessions() (int, FBCSessionInfo, error) {
	var sessionCount uint32
	var sessionInfo FBCSessionInfo
	ret := nvmlVgpuInstanceGetFBCSessions(vgpuInstance, &sessionCount, &sessionInfo)
	return int(sessionCount), sessionInfo, ret.error()
}

// nvml.VgpuInstanceGetGpuInstanceId()
func (l *library) VgpuInstanceGetGpuInstanceId(vgpuInstance VgpuInstance) (int, error) {
	return vgpuInstance.GetGpuInstanceId()
}

func (vgpuInstance nvmlVgpuInstance) GetGpuInstanceId() (int, error) {
	var gpuInstanceId uint32
	ret := nvmlVgpuInstanceGetGpuInstanceId(vgpuInstance, &gpuInstanceId)
	return int(gpuInstanceId), ret.error()
}

// nvml.VgpuInstanceGetGpuPciId()
func (l *library) VgpuInstanceGetGpuPciId(vgpuInstance VgpuInstance) (string, error) {
	return vgpuInstance.GetGpuPciId()
}

func (vgpuInstance nvmlVgpuInstance) GetGpuPciId() (string, error) {
	var length uint32 = 1 // Will be reduced upon returning
	for {
		vgpuPciId := make([]byte, length)
		ret := nvmlVgpuInstanceGetGpuPciId(vgpuInstance, &vgpuPciId[0], &length)
		if ret == nvmlSUCCESS {
			return string(vgpuPciId[:clen(vgpuPciId)]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return "", ret.error()
		}
		length *= 2
	}
}

// nvml.VgpuInstanceGetMetadata()
func (l *library) VgpuInstanceGetMetadata(vgpuInstance VgpuInstance) (VgpuMetadata, error) {
	return vgpuInstance.GetMetadata()
}

func (vgpuInstance nvmlVgpuInstance) GetMetadata() (VgpuMetadata, error) {
	var vgpuMetadata VgpuMetadata
	opaqueDataSize := unsafe.Sizeof(vgpuMetadata.nvmlVgpuMetadata.OpaqueData)
	vgpuMetadataSize := unsafe.Sizeof(vgpuMetadata.nvmlVgpuMetadata) - opaqueDataSize
	for {
		bufferSize := uint32(vgpuMetadataSize + opaqueDataSize)
		buffer := make([]byte, bufferSize)
		nvmlVgpuMetadataPtr := (*nvmlVgpuMetadata)(unsafe.Pointer(&buffer[0]))
		ret := nvmlVgpuInstanceGetMetadata(vgpuInstance, nvmlVgpuMetadataPtr, &bufferSize)
		if ret == nvmlSUCCESS {
			vgpuMetadata.nvmlVgpuMetadata = *nvmlVgpuMetadataPtr
			vgpuMetadata.OpaqueData = buffer[vgpuMetadataSize:bufferSize]
			return vgpuMetadata, ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return vgpuMetadata, ret.error()
		}
		opaqueDataSize = 2 * opaqueDataSize
	}
}

// nvml.VgpuInstanceGetAccountingMode()
func (l *library) VgpuInstanceGetAccountingMode(vgpuInstance VgpuInstance) (EnableState, error) {
	return vgpuInstance.GetAccountingMode()
}

func (vgpuInstance nvmlVgpuInstance) GetAccountingMode() (EnableState, error) {
	var mode EnableState
	ret := nvmlVgpuInstanceGetAccountingMode(vgpuInstance, &mode)
	return mode, ret.error()
}

// nvml.VgpuInstanceGetAccountingPids()
func (l *library) VgpuInstanceGetAccountingPids(vgpuInstance VgpuInstance) ([]int, error) {
	return vgpuInstance.GetAccountingPids()
}

func (vgpuInstance nvmlVgpuInstance) GetAccountingPids() ([]int, error) {
	var count uint32 = 1 // Will be reduced upon returning
	for {
		pids := make([]uint32, count)
		ret := nvmlVgpuInstanceGetAccountingPids(vgpuInstance, &count, &pids[0])
		if ret == nvmlSUCCESS {
			return uint32SliceToIntSlice(pids[:count]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		count *= 2
	}
}

// nvml.VgpuInstanceGetAccountingStats()
func (l *library) VgpuInstanceGetAccountingStats(vgpuInstance VgpuInstance, pid int) (AccountingStats, error) {
	return vgpuInstance.GetAccountingStats(pid)
}

func (vgpuInstance nvmlVgpuInstance) GetAccountingStats(pid int) (AccountingStats, error) {
	var stats AccountingStats
	ret := nvmlVgpuInstanceGetAccountingStats(vgpuInstance, uint32(pid), &stats)
	return stats, ret.error()
}

// nvml.GetVgpuCompatibility()
func (l *library) GetVgpuCompatibility(vgpuMetadata *VgpuMetadata, pgpuMetadata *VgpuPgpuMetadata) (VgpuPgpuCompatibility, error) {
	var compatibilityInfo VgpuPgpuCompatibility
	ret := nvmlGetVgpuCompatibility(&vgpuMetadata.nvmlVgpuMetadata, &pgpuMetadata.nvmlVgpuPgpuMetadata, &compatibilityInfo)
	return compatibilityInfo, ret.error()
}

// nvml.GetVgpuVersion()
func (l *library) GetVgpuVersion() (VgpuVersion, VgpuVersion, error) {
	var supported, current VgpuVersion
	ret := nvmlGetVgpuVersion(&supported, &current)
	return supported, current, ret.error()
}

// nvml.SetVgpuVersion()
func (l *library) SetVgpuVersion(vgpuVersion *VgpuVersion) error {
	return nvmlSetVgpuVersion(vgpuVersion).error()
}

// nvml.VgpuInstanceClearAccountingPids()
func (l *library) VgpuInstanceClearAccountingPids(vgpuInstance VgpuInstance) error {
	return vgpuInstance.ClearAccountingPids()
}

func (vgpuInstance nvmlVgpuInstance) ClearAccountingPids() error {
	return nvmlVgpuInstanceClearAccountingPids(vgpuInstance).error()
}

// nvml.VgpuInstanceGetMdevUUID()
func (l *library) VgpuInstanceGetMdevUUID(vgpuInstance VgpuInstance) (string, error) {
	return vgpuInstance.GetMdevUUID()
}

func (vgpuInstance nvmlVgpuInstance) GetMdevUUID() (string, error) {
	mdevUUID := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetMdevUUID(vgpuInstance, &mdevUUID[0], DEVICE_UUID_BUFFER_SIZE)
	return string(mdevUUID[:clen(mdevUUID)]), ret.error()
}

// nvml.VgpuTypeGetCapabilities()
func (l *library) VgpuTypeGetCapabilities(vgpuTypeId VgpuTypeId, capability VgpuCapability) (bool, error) {
	return vgpuTypeId.GetCapabilities(capability)
}

func (vgpuTypeId nvmlVgpuTypeId) GetCapabilities(capability VgpuCapability) (bool, error) {
	var capResult uint32
	ret := nvmlVgpuTypeGetCapabilities(vgpuTypeId, capability, &capResult)
	return (capResult != 0), ret.error()
}

// nvml.GetVgpuDriverCapabilities()
func (l *library) GetVgpuDriverCapabilities(capability VgpuDriverCapability) (bool, error) {
	var capResult uint32
	ret := nvmlGetVgpuDriverCapabilities(capability, &capResult)
	return (capResult != 0), ret.error()
}
