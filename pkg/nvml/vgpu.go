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
func (l *library) VgpuTypeGetClass(VgpuTypeId VgpuTypeId) (string, Return) {
	var Size uint32 = DEVICE_NAME_BUFFER_SIZE
	VgpuTypeClass := make([]byte, DEVICE_NAME_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetClass(VgpuTypeId, &VgpuTypeClass[0], &Size)
	return string(VgpuTypeClass[:clen(VgpuTypeClass)]), ret
}

func (VgpuTypeId VgpuTypeId) GetClass() (string, Return) {
	return VgpuTypeGetClass(VgpuTypeId)
}

// nvml.VgpuTypeGetName()
func (l *library) VgpuTypeGetName(VgpuTypeId VgpuTypeId) (string, Return) {
	var Size uint32 = DEVICE_NAME_BUFFER_SIZE
	VgpuTypeName := make([]byte, DEVICE_NAME_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetName(VgpuTypeId, &VgpuTypeName[0], &Size)
	return string(VgpuTypeName[:clen(VgpuTypeName)]), ret
}

func (VgpuTypeId VgpuTypeId) GetName() (string, Return) {
	return VgpuTypeGetName(VgpuTypeId)
}

// nvml.VgpuTypeGetGpuInstanceProfileId()
func (l *library) VgpuTypeGetGpuInstanceProfileId(VgpuTypeId VgpuTypeId) (uint32, Return) {
	var Size uint32
	ret := nvmlVgpuTypeGetGpuInstanceProfileId(VgpuTypeId, &Size)
	return Size, ret
}

func (VgpuTypeId VgpuTypeId) GetGpuInstanceProfileId() (uint32, Return) {
	return VgpuTypeGetGpuInstanceProfileId(VgpuTypeId)
}

// nvml.VgpuTypeGetDeviceID()
func (l *library) VgpuTypeGetDeviceID(VgpuTypeId VgpuTypeId) (uint64, uint64, Return) {
	var DeviceID, SubsystemID uint64
	ret := nvmlVgpuTypeGetDeviceID(VgpuTypeId, &DeviceID, &SubsystemID)
	return DeviceID, SubsystemID, ret
}

func (VgpuTypeId VgpuTypeId) GetDeviceID() (uint64, uint64, Return) {
	return VgpuTypeGetDeviceID(VgpuTypeId)
}

// nvml.VgpuTypeGetFramebufferSize()
func (l *library) VgpuTypeGetFramebufferSize(VgpuTypeId VgpuTypeId) (uint64, Return) {
	var FbSize uint64
	ret := nvmlVgpuTypeGetFramebufferSize(VgpuTypeId, &FbSize)
	return FbSize, ret
}

func (VgpuTypeId VgpuTypeId) GetFramebufferSize() (uint64, Return) {
	return VgpuTypeGetFramebufferSize(VgpuTypeId)
}

// nvml.VgpuTypeGetNumDisplayHeads()
func (l *library) VgpuTypeGetNumDisplayHeads(VgpuTypeId VgpuTypeId) (int, Return) {
	var NumDisplayHeads uint32
	ret := nvmlVgpuTypeGetNumDisplayHeads(VgpuTypeId, &NumDisplayHeads)
	return int(NumDisplayHeads), ret
}

func (VgpuTypeId VgpuTypeId) GetNumDisplayHeads() (int, Return) {
	return VgpuTypeGetNumDisplayHeads(VgpuTypeId)
}

// nvml.VgpuTypeGetResolution()
func (l *library) VgpuTypeGetResolution(VgpuTypeId VgpuTypeId, DisplayIndex int) (uint32, uint32, Return) {
	var Xdim, Ydim uint32
	ret := nvmlVgpuTypeGetResolution(VgpuTypeId, uint32(DisplayIndex), &Xdim, &Ydim)
	return Xdim, Ydim, ret
}

func (VgpuTypeId VgpuTypeId) GetResolution(DisplayIndex int) (uint32, uint32, Return) {
	return VgpuTypeGetResolution(VgpuTypeId, DisplayIndex)
}

// nvml.VgpuTypeGetLicense()
func (l *library) VgpuTypeGetLicense(VgpuTypeId VgpuTypeId) (string, Return) {
	VgpuTypeLicenseString := make([]byte, GRID_LICENSE_BUFFER_SIZE)
	ret := nvmlVgpuTypeGetLicense(VgpuTypeId, &VgpuTypeLicenseString[0], GRID_LICENSE_BUFFER_SIZE)
	return string(VgpuTypeLicenseString[:clen(VgpuTypeLicenseString)]), ret
}

func (VgpuTypeId VgpuTypeId) GetLicense() (string, Return) {
	return VgpuTypeGetLicense(VgpuTypeId)
}

// nvml.VgpuTypeGetFrameRateLimit()
func (l *library) VgpuTypeGetFrameRateLimit(VgpuTypeId VgpuTypeId) (uint32, Return) {
	var FrameRateLimit uint32
	ret := nvmlVgpuTypeGetFrameRateLimit(VgpuTypeId, &FrameRateLimit)
	return FrameRateLimit, ret
}

func (VgpuTypeId VgpuTypeId) GetFrameRateLimit() (uint32, Return) {
	return VgpuTypeGetFrameRateLimit(VgpuTypeId)
}

// nvml.VgpuTypeGetMaxInstances()
func (l *library) VgpuTypeGetMaxInstances(Device Device, VgpuTypeId VgpuTypeId) (int, Return) {
	return Device.VgpuTypeGetMaxInstances(VgpuTypeId)
}

func (Device nvmlDevice) VgpuTypeGetMaxInstances(VgpuTypeId VgpuTypeId) (int, Return) {
	var VgpuInstanceCount uint32
	ret := nvmlVgpuTypeGetMaxInstances(Device, VgpuTypeId, &VgpuInstanceCount)
	return int(VgpuInstanceCount), ret
}

func (VgpuTypeId VgpuTypeId) GetMaxInstances(Device Device) (int, Return) {
	return VgpuTypeGetMaxInstances(Device, VgpuTypeId)
}

// nvml.VgpuTypeGetMaxInstancesPerVm()
func (l *library) VgpuTypeGetMaxInstancesPerVm(VgpuTypeId VgpuTypeId) (int, Return) {
	var VgpuInstanceCountPerVm uint32
	ret := nvmlVgpuTypeGetMaxInstancesPerVm(VgpuTypeId, &VgpuInstanceCountPerVm)
	return int(VgpuInstanceCountPerVm), ret
}

func (VgpuTypeId VgpuTypeId) GetMaxInstancesPerVm() (int, Return) {
	return VgpuTypeGetMaxInstancesPerVm(VgpuTypeId)
}

// nvml.VgpuInstanceGetVmID()
func (l *library) VgpuInstanceGetVmID(VgpuInstance VgpuInstance) (string, VgpuVmIdType, Return) {
	return VgpuInstance.GetVmID()
}

func (VgpuInstance nvmlVgpuInstance) GetVmID() (string, VgpuVmIdType, Return) {
	var VmIdType VgpuVmIdType
	VmId := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetVmID(VgpuInstance, &VmId[0], DEVICE_UUID_BUFFER_SIZE, &VmIdType)
	return string(VmId[:clen(VmId)]), VmIdType, ret
}

// nvml.VgpuInstanceGetUUID()
func (l *library) VgpuInstanceGetUUID(VgpuInstance VgpuInstance) (string, Return) {
	return VgpuInstance.GetUUID()
}

func (VgpuInstance nvmlVgpuInstance) GetUUID() (string, Return) {
	Uuid := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetUUID(VgpuInstance, &Uuid[0], DEVICE_UUID_BUFFER_SIZE)
	return string(Uuid[:clen(Uuid)]), ret
}

// nvml.VgpuInstanceGetVmDriverVersion()
func (l *library) VgpuInstanceGetVmDriverVersion(VgpuInstance VgpuInstance) (string, Return) {
	return VgpuInstance.GetVmDriverVersion()
}

func (VgpuInstance nvmlVgpuInstance) GetVmDriverVersion() (string, Return) {
	Version := make([]byte, SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetVmDriverVersion(VgpuInstance, &Version[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.VgpuInstanceGetFbUsage()
func (l *library) VgpuInstanceGetFbUsage(VgpuInstance VgpuInstance) (uint64, Return) {
	return VgpuInstance.GetFbUsage()
}

func (VgpuInstance nvmlVgpuInstance) GetFbUsage() (uint64, Return) {
	var FbUsage uint64
	ret := nvmlVgpuInstanceGetFbUsage(VgpuInstance, &FbUsage)
	return FbUsage, ret
}

// nvml.VgpuInstanceGetLicenseInfo()
func (l *library) VgpuInstanceGetLicenseInfo(VgpuInstance VgpuInstance) (VgpuLicenseInfo, Return) {
	return VgpuInstance.GetLicenseInfo()
}

func (VgpuInstance nvmlVgpuInstance) GetLicenseInfo() (VgpuLicenseInfo, Return) {
	var LicenseInfo VgpuLicenseInfo
	ret := nvmlVgpuInstanceGetLicenseInfo(VgpuInstance, &LicenseInfo)
	return LicenseInfo, ret
}

// nvml.VgpuInstanceGetLicenseStatus()
func (l *library) VgpuInstanceGetLicenseStatus(VgpuInstance VgpuInstance) (int, Return) {
	return VgpuInstance.GetLicenseStatus()
}

func (VgpuInstance nvmlVgpuInstance) GetLicenseStatus() (int, Return) {
	var Licensed uint32
	ret := nvmlVgpuInstanceGetLicenseStatus(VgpuInstance, &Licensed)
	return int(Licensed), ret
}

// nvml.VgpuInstanceGetType()
func (l *library) VgpuInstanceGetType(VgpuInstance VgpuInstance) (VgpuTypeId, Return) {
	return VgpuInstance.GetType()
}

func (VgpuInstance nvmlVgpuInstance) GetType() (VgpuTypeId, Return) {
	var VgpuTypeId VgpuTypeId
	ret := nvmlVgpuInstanceGetType(VgpuInstance, &VgpuTypeId)
	return VgpuTypeId, ret
}

// nvml.VgpuInstanceGetFrameRateLimit()
func (l *library) VgpuInstanceGetFrameRateLimit(VgpuInstance VgpuInstance) (uint32, Return) {
	return VgpuInstance.GetFrameRateLimit()
}

func (VgpuInstance nvmlVgpuInstance) GetFrameRateLimit() (uint32, Return) {
	var FrameRateLimit uint32
	ret := nvmlVgpuInstanceGetFrameRateLimit(VgpuInstance, &FrameRateLimit)
	return FrameRateLimit, ret
}

// nvml.VgpuInstanceGetEccMode()
func (l *library) VgpuInstanceGetEccMode(VgpuInstance VgpuInstance) (EnableState, Return) {
	return VgpuInstance.GetEccMode()
}

func (VgpuInstance nvmlVgpuInstance) GetEccMode() (EnableState, Return) {
	var EccMode EnableState
	ret := nvmlVgpuInstanceGetEccMode(VgpuInstance, &EccMode)
	return EccMode, ret
}

// nvml.VgpuInstanceGetEncoderCapacity()
func (l *library) VgpuInstanceGetEncoderCapacity(VgpuInstance VgpuInstance) (int, Return) {
	return VgpuInstance.GetEncoderCapacity()
}

func (VgpuInstance nvmlVgpuInstance) GetEncoderCapacity() (int, Return) {
	var EncoderCapacity uint32
	ret := nvmlVgpuInstanceGetEncoderCapacity(VgpuInstance, &EncoderCapacity)
	return int(EncoderCapacity), ret
}

// nvml.VgpuInstanceSetEncoderCapacity()
func (l *library) VgpuInstanceSetEncoderCapacity(VgpuInstance VgpuInstance, EncoderCapacity int) Return {
	return VgpuInstance.SetEncoderCapacity(EncoderCapacity)
}

func (VgpuInstance nvmlVgpuInstance) SetEncoderCapacity(EncoderCapacity int) Return {
	return nvmlVgpuInstanceSetEncoderCapacity(VgpuInstance, uint32(EncoderCapacity))
}

// nvml.VgpuInstanceGetEncoderStats()
func (l *library) VgpuInstanceGetEncoderStats(VgpuInstance VgpuInstance) (int, uint32, uint32, Return) {
	return VgpuInstance.GetEncoderStats()
}

func (VgpuInstance nvmlVgpuInstance) GetEncoderStats() (int, uint32, uint32, Return) {
	var SessionCount, AverageFps, AverageLatency uint32
	ret := nvmlVgpuInstanceGetEncoderStats(VgpuInstance, &SessionCount, &AverageFps, &AverageLatency)
	return int(SessionCount), AverageFps, AverageLatency, ret
}

// nvml.VgpuInstanceGetEncoderSessions()
func (l *library) VgpuInstanceGetEncoderSessions(VgpuInstance VgpuInstance) (int, EncoderSessionInfo, Return) {
	return VgpuInstance.GetEncoderSessions()
}

func (VgpuInstance nvmlVgpuInstance) GetEncoderSessions() (int, EncoderSessionInfo, Return) {
	var SessionCount uint32
	var SessionInfo EncoderSessionInfo
	ret := nvmlVgpuInstanceGetEncoderSessions(VgpuInstance, &SessionCount, &SessionInfo)
	return int(SessionCount), SessionInfo, ret
}

// nvml.VgpuInstanceGetFBCStats()
func (l *library) VgpuInstanceGetFBCStats(VgpuInstance VgpuInstance) (FBCStats, Return) {
	return VgpuInstance.GetFBCStats()
}

func (VgpuInstance nvmlVgpuInstance) GetFBCStats() (FBCStats, Return) {
	var FbcStats FBCStats
	ret := nvmlVgpuInstanceGetFBCStats(VgpuInstance, &FbcStats)
	return FbcStats, ret
}

// nvml.VgpuInstanceGetFBCSessions()
func (l *library) VgpuInstanceGetFBCSessions(VgpuInstance VgpuInstance) (int, FBCSessionInfo, Return) {
	return VgpuInstance.GetFBCSessions()
}

func (VgpuInstance nvmlVgpuInstance) GetFBCSessions() (int, FBCSessionInfo, Return) {
	var SessionCount uint32
	var SessionInfo FBCSessionInfo
	ret := nvmlVgpuInstanceGetFBCSessions(VgpuInstance, &SessionCount, &SessionInfo)
	return int(SessionCount), SessionInfo, ret
}

// nvml.VgpuInstanceGetGpuInstanceId()
func (l *library) VgpuInstanceGetGpuInstanceId(VgpuInstance VgpuInstance) (int, Return) {
	return VgpuInstance.GetGpuInstanceId()
}

func (VgpuInstance nvmlVgpuInstance) GetGpuInstanceId() (int, Return) {
	var gpuInstanceId uint32
	ret := nvmlVgpuInstanceGetGpuInstanceId(VgpuInstance, &gpuInstanceId)
	return int(gpuInstanceId), ret
}

// nvml.VgpuInstanceGetGpuPciId()
func (l *library) VgpuInstanceGetGpuPciId(VgpuInstance VgpuInstance) (string, Return) {
	return VgpuInstance.GetGpuPciId()
}

func (VgpuInstance nvmlVgpuInstance) GetGpuPciId() (string, Return) {
	var Length uint32 = 1 // Will be reduced upon returning
	for {
		VgpuPciId := make([]byte, Length)
		ret := nvmlVgpuInstanceGetGpuPciId(VgpuInstance, &VgpuPciId[0], &Length)
		if ret == SUCCESS {
			return string(VgpuPciId[:clen(VgpuPciId)]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return "", ret
		}
		Length *= 2
	}
}

// nvml.VgpuInstanceGetMetadata()
func (l *library) VgpuInstanceGetMetadata(VgpuInstance VgpuInstance) (VgpuMetadata, Return) {
	return VgpuInstance.GetMetadata()
}

func (VgpuInstance nvmlVgpuInstance) GetMetadata() (VgpuMetadata, Return) {
	var VgpuMetadata VgpuMetadata
	OpaqueDataSize := unsafe.Sizeof(VgpuMetadata.nvmlVgpuMetadata.OpaqueData)
	VgpuMetadataSize := unsafe.Sizeof(VgpuMetadata.nvmlVgpuMetadata) - OpaqueDataSize
	for {
		BufferSize := uint32(VgpuMetadataSize + OpaqueDataSize)
		Buffer := make([]byte, BufferSize)
		nvmlVgpuMetadataPtr := (*nvmlVgpuMetadata)(unsafe.Pointer(&Buffer[0]))
		ret := nvmlVgpuInstanceGetMetadata(VgpuInstance, nvmlVgpuMetadataPtr, &BufferSize)
		if ret == SUCCESS {
			VgpuMetadata.nvmlVgpuMetadata = *nvmlVgpuMetadataPtr
			VgpuMetadata.OpaqueData = Buffer[VgpuMetadataSize:BufferSize]
			return VgpuMetadata, ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return VgpuMetadata, ret
		}
		OpaqueDataSize = 2 * OpaqueDataSize
	}
}

// nvml.VgpuInstanceGetAccountingMode()
func (l *library) VgpuInstanceGetAccountingMode(VgpuInstance VgpuInstance) (EnableState, Return) {
	return VgpuInstance.GetAccountingMode()
}

func (VgpuInstance nvmlVgpuInstance) GetAccountingMode() (EnableState, Return) {
	var Mode EnableState
	ret := nvmlVgpuInstanceGetAccountingMode(VgpuInstance, &Mode)
	return Mode, ret
}

// nvml.VgpuInstanceGetAccountingPids()
func (l *library) VgpuInstanceGetAccountingPids(VgpuInstance VgpuInstance) ([]int, Return) {
	return VgpuInstance.GetAccountingPids()
}

func (VgpuInstance nvmlVgpuInstance) GetAccountingPids() ([]int, Return) {
	var Count uint32 = 1 // Will be reduced upon returning
	for {
		Pids := make([]uint32, Count)
		ret := nvmlVgpuInstanceGetAccountingPids(VgpuInstance, &Count, &Pids[0])
		if ret == SUCCESS {
			return uint32SliceToIntSlice(Pids[:Count]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		Count *= 2
	}
}

// nvml.VgpuInstanceGetAccountingStats()
func (l *library) VgpuInstanceGetAccountingStats(VgpuInstance VgpuInstance, Pid int) (AccountingStats, Return) {
	return VgpuInstance.GetAccountingStats(Pid)
}

func (VgpuInstance nvmlVgpuInstance) GetAccountingStats(Pid int) (AccountingStats, Return) {
	var Stats AccountingStats
	ret := nvmlVgpuInstanceGetAccountingStats(VgpuInstance, uint32(Pid), &Stats)
	return Stats, ret
}

// nvml.GetVgpuCompatibility()
func (l *library) GetVgpuCompatibility(nvmlVgpuMetadata *nvmlVgpuMetadata, PgpuMetadata *nvmlVgpuPgpuMetadata) (VgpuPgpuCompatibility, Return) {
	var CompatibilityInfo VgpuPgpuCompatibility
	ret := nvmlGetVgpuCompatibility(nvmlVgpuMetadata, PgpuMetadata, &CompatibilityInfo)
	return CompatibilityInfo, ret
}

// nvml.GetVgpuVersion()
func (l *library) GetVgpuVersion() (VgpuVersion, VgpuVersion, Return) {
	var Supported, Current VgpuVersion
	ret := nvmlGetVgpuVersion(&Supported, &Current)
	return Supported, Current, ret
}

// nvml.SetVgpuVersion()
func (l *library) SetVgpuVersion(VgpuVersion *VgpuVersion) Return {
	return nvmlSetVgpuVersion(VgpuVersion)
}

// nvml.VgpuInstanceClearAccountingPids()
func (l *library) VgpuInstanceClearAccountingPids(VgpuInstance VgpuInstance) Return {
	return VgpuInstance.ClearAccountingPids()
}

func (VgpuInstance nvmlVgpuInstance) ClearAccountingPids() Return {
	return nvmlVgpuInstanceClearAccountingPids(VgpuInstance)
}

// nvml.VgpuInstanceGetMdevUUID()
func (l *library) VgpuInstanceGetMdevUUID(VgpuInstance VgpuInstance) (string, Return) {
	return VgpuInstance.GetMdevUUID()
}

func (VgpuInstance nvmlVgpuInstance) GetMdevUUID() (string, Return) {
	MdevUuid := make([]byte, DEVICE_UUID_BUFFER_SIZE)
	ret := nvmlVgpuInstanceGetMdevUUID(VgpuInstance, &MdevUuid[0], DEVICE_UUID_BUFFER_SIZE)
	return string(MdevUuid[:clen(MdevUuid)]), ret
}

// nvml.VgpuTypeGetCapabilities()
func (l *library) VgpuTypeGetCapabilities(VgpuTypeId VgpuTypeId, Capability VgpuCapability) (bool, Return) {
	var CapResult uint32
	ret := nvmlVgpuTypeGetCapabilities(VgpuTypeId, Capability, &CapResult)
	return (CapResult != 0), ret
}

func (VgpuTypeId VgpuTypeId) GetCapabilities(Capability VgpuCapability) (bool, Return) {
	return VgpuTypeGetCapabilities(VgpuTypeId, Capability)
}

// nvml.GetVgpuDriverCapabilities()
func (l *library) GetVgpuDriverCapabilities(Capability VgpuDriverCapability) (bool, Return) {
	var CapResult uint32
	ret := nvmlGetVgpuDriverCapabilities(Capability, &CapResult)
	return (CapResult != 0), ret
}
