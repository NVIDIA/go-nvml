/**
# Copyright 2025 NVIDIA CORPORATION
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

// Generated Code; DO NOT EDIT.

package nvml

// CgoAPI is a global variable providing direct calls to the cgo API rather than the standard wrappers
var CgoAPI = cgoapi{}

type cgoapi struct{}

func (cgoapi) ComputeInstanceDestroy(computeinstance ComputeInstance) Return {
	ret := nvmlComputeInstanceDestroy(nvmlComputeInstanceHandle(computeinstance))
	return ret
}

func (cgoapi) ComputeInstanceGetInfo(computeinstance ComputeInstance, info *ComputeInstanceInfo) Return {
	var nvmlInfo nvmlComputeInstanceInfo
	ret := nvmlComputeInstanceGetInfo(nvmlComputeInstanceHandle(computeinstance), &nvmlInfo)
	return ret
}

func (cgoapi) ComputeInstanceGetInfo_v1(computeinstance ComputeInstance, info *ComputeInstanceInfo) Return {
	var nvmlInfo nvmlComputeInstanceInfo
	ret := nvmlComputeInstanceGetInfo_v1(nvmlComputeInstanceHandle(computeinstance), &nvmlInfo)
	return ret
}

func (cgoapi) ComputeInstanceGetInfo_v2(computeinstance ComputeInstance, info *ComputeInstanceInfo) Return {
	var nvmlInfo nvmlComputeInstanceInfo
	ret := nvmlComputeInstanceGetInfo_v2(nvmlComputeInstanceHandle(computeinstance), &nvmlInfo)
	return ret
}

func (cgoapi) DeviceClearAccountingPids(device Device) Return {
	ret := nvmlDeviceClearAccountingPids(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceClearCpuAffinity(device Device) Return {
	ret := nvmlDeviceClearCpuAffinity(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceClearEccErrorCounts(device Device, countertype EccCounterType) Return {
	ret := nvmlDeviceClearEccErrorCounts(nvmlDeviceHandle(device), countertype)
	return ret
}

func (cgoapi) DeviceClearFieldValues(device Device, valuescount int32, values *FieldValue) Return {
	ret := nvmlDeviceClearFieldValues(nvmlDeviceHandle(device), valuescount, values)
	return ret
}

func (cgoapi) DeviceCreateGpuInstance(device Device, profileid uint32, gpuinstance *GpuInstance) Return {
	var nvmlGpuinstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstance(nvmlDeviceHandle(device), profileid, &nvmlGpuinstance)
	*gpuinstance = GpuInstance(nvmlGpuinstance)
	return ret
}

func (cgoapi) DeviceCreateGpuInstanceWithPlacement(device Device, profileid uint32, placement *GpuInstancePlacement, gpuinstance *GpuInstance) Return {
	var nvmlGpuinstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstanceWithPlacement(nvmlDeviceHandle(device), profileid, placement, &nvmlGpuinstance)
	*gpuinstance = GpuInstance(nvmlGpuinstance)
	return ret
}

func (cgoapi) DeviceDiscoverGpus(pciinfo *PciInfo) Return {
	ret := nvmlDeviceDiscoverGpus(pciinfo)
	return ret
}

func (cgoapi) DeviceFreezeNvLinkUtilizationCounter(device Device, link uint32, counter uint32, freeze EnableState) Return {
	ret := nvmlDeviceFreezeNvLinkUtilizationCounter(nvmlDeviceHandle(device), link, counter, freeze)
	return ret
}

func (cgoapi) DeviceGetAPIRestriction(device Device, apitype RestrictedAPI, isrestricted *EnableState) Return {
	ret := nvmlDeviceGetAPIRestriction(nvmlDeviceHandle(device), apitype, isrestricted)
	return ret
}

func (cgoapi) DeviceGetAccountingBufferSize(device Device, buffersize *uint32) Return {
	ret := nvmlDeviceGetAccountingBufferSize(nvmlDeviceHandle(device), buffersize)
	return ret
}

func (cgoapi) DeviceGetAccountingMode(device Device, mode *EnableState) Return {
	ret := nvmlDeviceGetAccountingMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceGetAccountingPids(device Device, count *uint32, pids *uint32) Return {
	ret := nvmlDeviceGetAccountingPids(nvmlDeviceHandle(device), count, pids)
	return ret
}

func (cgoapi) DeviceGetAccountingStats(device Device, pid uint32, stats *AccountingStats) Return {
	ret := nvmlDeviceGetAccountingStats(nvmlDeviceHandle(device), pid, stats)
	return ret
}

func (cgoapi) DeviceGetActiveVgpus(device Device, vgpucount *uint32, vgpuinstances *VgpuInstance) Return {
	var nvmlVgpuinstances nvmlVgpuInstance
	ret := nvmlDeviceGetActiveVgpus(nvmlDeviceHandle(device), vgpucount, &nvmlVgpuinstances)
	*vgpuinstances = VgpuInstance(nvmlVgpuinstances)
	return ret
}

func (cgoapi) DeviceGetAdaptiveClockInfoStatus(device Device, adaptiveclockstatus *uint32) Return {
	ret := nvmlDeviceGetAdaptiveClockInfoStatus(nvmlDeviceHandle(device), adaptiveclockstatus)
	return ret
}

func (cgoapi) DeviceGetApplicationsClock(device Device, clocktype ClockType, clockmhz *uint32) Return {
	ret := nvmlDeviceGetApplicationsClock(nvmlDeviceHandle(device), clocktype, clockmhz)
	return ret
}

func (cgoapi) DeviceGetArchitecture(device Device, arch *DeviceArchitecture) Return {
	ret := nvmlDeviceGetArchitecture(nvmlDeviceHandle(device), arch)
	return ret
}

func (cgoapi) DeviceGetAttributes(device Device, attributes *DeviceAttributes) Return {
	ret := nvmlDeviceGetAttributes(nvmlDeviceHandle(device), attributes)
	return ret
}

func (cgoapi) DeviceGetAttributes_v1(device Device, attributes *DeviceAttributes) Return {
	ret := nvmlDeviceGetAttributes_v1(nvmlDeviceHandle(device), attributes)
	return ret
}

func (cgoapi) DeviceGetAttributes_v2(device Device, attributes *DeviceAttributes) Return {
	ret := nvmlDeviceGetAttributes_v2(nvmlDeviceHandle(device), attributes)
	return ret
}

func (cgoapi) DeviceGetAutoBoostedClocksEnabled(device Device, isenabled *EnableState, defaultisenabled *EnableState) Return {
	ret := nvmlDeviceGetAutoBoostedClocksEnabled(nvmlDeviceHandle(device), isenabled, defaultisenabled)
	return ret
}

func (cgoapi) DeviceGetBAR1MemoryInfo(device Device, bar1memory *BAR1Memory) Return {
	ret := nvmlDeviceGetBAR1MemoryInfo(nvmlDeviceHandle(device), bar1memory)
	return ret
}

func (cgoapi) DeviceGetBoardId(device Device, boardid *uint32) Return {
	ret := nvmlDeviceGetBoardId(nvmlDeviceHandle(device), boardid)
	return ret
}

func (cgoapi) DeviceGetBoardPartNumber(device Device, partnumber *byte, length uint32) Return {
	ret := nvmlDeviceGetBoardPartNumber(nvmlDeviceHandle(device), partnumber, length)
	return ret
}

func (cgoapi) DeviceGetBrand(device Device, _type *BrandType) Return {
	ret := nvmlDeviceGetBrand(nvmlDeviceHandle(device), _type)
	return ret
}

func (cgoapi) DeviceGetBridgeChipInfo(device Device, bridgehierarchy *BridgeChipHierarchy) Return {
	ret := nvmlDeviceGetBridgeChipInfo(nvmlDeviceHandle(device), bridgehierarchy)
	return ret
}

func (cgoapi) DeviceGetBusType(device Device, _type *BusType) Return {
	ret := nvmlDeviceGetBusType(nvmlDeviceHandle(device), _type)
	return ret
}

func (cgoapi) DeviceGetC2cModeInfoV(device Device, c2cmodeinfo *C2cModeInfo_v1) Return {
	ret := nvmlDeviceGetC2cModeInfoV(nvmlDeviceHandle(device), c2cmodeinfo)
	return ret
}

func (cgoapi) DeviceGetCapabilities(device Device, caps *DeviceCapabilities) Return {
	ret := nvmlDeviceGetCapabilities(nvmlDeviceHandle(device), caps)
	return ret
}

func (cgoapi) DeviceGetClkMonStatus(device Device, status *ClkMonStatus) Return {
	ret := nvmlDeviceGetClkMonStatus(nvmlDeviceHandle(device), status)
	return ret
}

func (cgoapi) DeviceGetClock(device Device, clocktype ClockType, clockid ClockId, clockmhz *uint32) Return {
	ret := nvmlDeviceGetClock(nvmlDeviceHandle(device), clocktype, clockid, clockmhz)
	return ret
}

func (cgoapi) DeviceGetClockInfo(device Device, _type ClockType, clock *uint32) Return {
	ret := nvmlDeviceGetClockInfo(nvmlDeviceHandle(device), _type, clock)
	return ret
}

func (cgoapi) DeviceGetClockOffsets(device Device, info *ClockOffset) Return {
	ret := nvmlDeviceGetClockOffsets(nvmlDeviceHandle(device), info)
	return ret
}

func (cgoapi) DeviceGetComputeInstanceId(device Device, id *uint32) Return {
	ret := nvmlDeviceGetComputeInstanceId(nvmlDeviceHandle(device), id)
	return ret
}

func (cgoapi) DeviceGetComputeMode(device Device, mode *ComputeMode) Return {
	ret := nvmlDeviceGetComputeMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceGetComputeRunningProcesses(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetComputeRunningProcesses(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetComputeRunningProcesses_v1(device Device, infocount *uint32, infos *ProcessInfo_v1) Return {
	ret := nvmlDeviceGetComputeRunningProcesses_v1(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetComputeRunningProcesses_v2(device Device, infocount *uint32, infos *ProcessInfo_v2) Return {
	ret := nvmlDeviceGetComputeRunningProcesses_v2(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetComputeRunningProcesses_v3(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetComputeRunningProcesses_v3(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetConfComputeGpuAttestationReport(device Device, gpuatstreport *ConfComputeGpuAttestationReport) Return {
	ret := nvmlDeviceGetConfComputeGpuAttestationReport(nvmlDeviceHandle(device), gpuatstreport)
	return ret
}

func (cgoapi) DeviceGetConfComputeGpuCertificate(device Device, gpucert *ConfComputeGpuCertificate) Return {
	ret := nvmlDeviceGetConfComputeGpuCertificate(nvmlDeviceHandle(device), gpucert)
	return ret
}

func (cgoapi) DeviceGetConfComputeMemSizeInfo(device Device, meminfo *ConfComputeMemSizeInfo) Return {
	ret := nvmlDeviceGetConfComputeMemSizeInfo(nvmlDeviceHandle(device), meminfo)
	return ret
}

func (cgoapi) DeviceGetConfComputeProtectedMemoryUsage(device Device, memory *Memory) Return {
	ret := nvmlDeviceGetConfComputeProtectedMemoryUsage(nvmlDeviceHandle(device), memory)
	return ret
}

func (cgoapi) DeviceGetCoolerInfo(device Device, coolerinfo *CoolerInfo) Return {
	ret := nvmlDeviceGetCoolerInfo(nvmlDeviceHandle(device), coolerinfo)
	return ret
}

func (cgoapi) DeviceGetCount(devicecount *uint32) Return {
	ret := nvmlDeviceGetCount(devicecount)
	return ret
}

func (cgoapi) DeviceGetCount_v1(devicecount *uint32) Return {
	ret := nvmlDeviceGetCount_v1(devicecount)
	return ret
}

func (cgoapi) DeviceGetCount_v2(devicecount *uint32) Return {
	ret := nvmlDeviceGetCount_v2(devicecount)
	return ret
}

func (cgoapi) DeviceGetCpuAffinity(device Device, cpusetsize uint32, cpuset *uint) Return {
	ret := nvmlDeviceGetCpuAffinity(nvmlDeviceHandle(device), cpusetsize, cpuset)
	return ret
}

func (cgoapi) DeviceGetCpuAffinityWithinScope(device Device, cpusetsize uint32, cpuset *uint, scope AffinityScope) Return {
	ret := nvmlDeviceGetCpuAffinityWithinScope(nvmlDeviceHandle(device), cpusetsize, cpuset, scope)
	return ret
}

func (cgoapi) DeviceGetCreatableVgpus(device Device, vgpucount *uint32, vgputypeids *VgpuTypeId) Return {
	var nvmlVgputypeids nvmlVgpuTypeId
	ret := nvmlDeviceGetCreatableVgpus(nvmlDeviceHandle(device), vgpucount, &nvmlVgputypeids)
	*vgputypeids = VgpuTypeId(nvmlVgputypeids)
	return ret
}

func (cgoapi) DeviceGetCudaComputeCapability(device Device, major *int32, minor *int32) Return {
	ret := nvmlDeviceGetCudaComputeCapability(nvmlDeviceHandle(device), major, minor)
	return ret
}

func (cgoapi) DeviceGetCurrPcieLinkGeneration(device Device, currlinkgen *uint32) Return {
	ret := nvmlDeviceGetCurrPcieLinkGeneration(nvmlDeviceHandle(device), currlinkgen)
	return ret
}

func (cgoapi) DeviceGetCurrPcieLinkWidth(device Device, currlinkwidth *uint32) Return {
	ret := nvmlDeviceGetCurrPcieLinkWidth(nvmlDeviceHandle(device), currlinkwidth)
	return ret
}

func (cgoapi) DeviceGetCurrentClockFreqs(device Device, currentclockfreqs *DeviceCurrentClockFreqs) Return {
	ret := nvmlDeviceGetCurrentClockFreqs(nvmlDeviceHandle(device), currentclockfreqs)
	return ret
}

func (cgoapi) DeviceGetCurrentClocksEventReasons(device Device, clockseventreasons *uint64) Return {
	ret := nvmlDeviceGetCurrentClocksEventReasons(nvmlDeviceHandle(device), clockseventreasons)
	return ret
}

func (cgoapi) DeviceGetCurrentClocksThrottleReasons(device Device, clocksthrottlereasons *uint64) Return {
	ret := nvmlDeviceGetCurrentClocksThrottleReasons(nvmlDeviceHandle(device), clocksthrottlereasons)
	return ret
}

func (cgoapi) DeviceGetDecoderUtilization(device Device, utilization *uint32, samplingperiodus *uint32) Return {
	ret := nvmlDeviceGetDecoderUtilization(nvmlDeviceHandle(device), utilization, samplingperiodus)
	return ret
}

func (cgoapi) DeviceGetDefaultApplicationsClock(device Device, clocktype ClockType, clockmhz *uint32) Return {
	ret := nvmlDeviceGetDefaultApplicationsClock(nvmlDeviceHandle(device), clocktype, clockmhz)
	return ret
}

func (cgoapi) DeviceGetDefaultEccMode(device Device, defaultmode *EnableState) Return {
	ret := nvmlDeviceGetDefaultEccMode(nvmlDeviceHandle(device), defaultmode)
	return ret
}

func (cgoapi) DeviceGetDetailedEccErrors(device Device, errortype MemoryErrorType, countertype EccCounterType, ecccounts *EccErrorCounts) Return {
	ret := nvmlDeviceGetDetailedEccErrors(nvmlDeviceHandle(device), errortype, countertype, ecccounts)
	return ret
}

func (cgoapi) DeviceGetDeviceHandleFromMigDeviceHandle(migdevice Device, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetDeviceHandleFromMigDeviceHandle(nvmlDeviceHandle(migdevice), &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetDisplayActive(device Device, isactive *EnableState) Return {
	ret := nvmlDeviceGetDisplayActive(nvmlDeviceHandle(device), isactive)
	return ret
}

func (cgoapi) DeviceGetDisplayMode(device Device, display *EnableState) Return {
	ret := nvmlDeviceGetDisplayMode(nvmlDeviceHandle(device), display)
	return ret
}

func (cgoapi) DeviceGetDramEncryptionMode(device Device, current *DramEncryptionInfo, pending *DramEncryptionInfo) Return {
	ret := nvmlDeviceGetDramEncryptionMode(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetDriverModel(device Device, current *DriverModel, pending *DriverModel) Return {
	ret := nvmlDeviceGetDriverModel(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetDriverModel_v1(device Device, current *DriverModel, pending *DriverModel) Return {
	ret := nvmlDeviceGetDriverModel_v1(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetDriverModel_v2(device Device, current *DriverModel, pending *DriverModel) Return {
	ret := nvmlDeviceGetDriverModel_v2(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetDynamicPstatesInfo(device Device, pdynamicpstatesinfo *GpuDynamicPstatesInfo) Return {
	ret := nvmlDeviceGetDynamicPstatesInfo(nvmlDeviceHandle(device), pdynamicpstatesinfo)
	return ret
}

func (cgoapi) DeviceGetEccMode(device Device, current *EnableState, pending *EnableState) Return {
	ret := nvmlDeviceGetEccMode(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetEncoderCapacity(device Device, encoderquerytype EncoderType, encodercapacity *uint32) Return {
	ret := nvmlDeviceGetEncoderCapacity(nvmlDeviceHandle(device), encoderquerytype, encodercapacity)
	return ret
}

func (cgoapi) DeviceGetEncoderSessions(device Device, sessioncount *uint32, sessioninfos *EncoderSessionInfo) Return {
	ret := nvmlDeviceGetEncoderSessions(nvmlDeviceHandle(device), sessioncount, sessioninfos)
	return ret
}

func (cgoapi) DeviceGetEncoderStats(device Device, sessioncount *uint32, averagefps *uint32, averagelatency *uint32) Return {
	ret := nvmlDeviceGetEncoderStats(nvmlDeviceHandle(device), sessioncount, averagefps, averagelatency)
	return ret
}

func (cgoapi) DeviceGetEncoderUtilization(device Device, utilization *uint32, samplingperiodus *uint32) Return {
	ret := nvmlDeviceGetEncoderUtilization(nvmlDeviceHandle(device), utilization, samplingperiodus)
	return ret
}

func (cgoapi) DeviceGetEnforcedPowerLimit(device Device, limit *uint32) Return {
	ret := nvmlDeviceGetEnforcedPowerLimit(nvmlDeviceHandle(device), limit)
	return ret
}

func (cgoapi) DeviceGetFBCSessions(device Device, sessioncount *uint32, sessioninfo *FBCSessionInfo) Return {
	ret := nvmlDeviceGetFBCSessions(nvmlDeviceHandle(device), sessioncount, sessioninfo)
	return ret
}

func (cgoapi) DeviceGetFBCStats(device Device, fbcstats *FBCStats) Return {
	ret := nvmlDeviceGetFBCStats(nvmlDeviceHandle(device), fbcstats)
	return ret
}

func (cgoapi) DeviceGetFanControlPolicy_v2(device Device, fan uint32, policy *FanControlPolicy) Return {
	ret := nvmlDeviceGetFanControlPolicy_v2(nvmlDeviceHandle(device), fan, policy)
	return ret
}

func (cgoapi) DeviceGetFanSpeed(device Device, speed *uint32) Return {
	ret := nvmlDeviceGetFanSpeed(nvmlDeviceHandle(device), speed)
	return ret
}

func (cgoapi) DeviceGetFanSpeedRPM(device Device, fanspeed *FanSpeedInfo) Return {
	ret := nvmlDeviceGetFanSpeedRPM(nvmlDeviceHandle(device), fanspeed)
	return ret
}

func (cgoapi) DeviceGetFanSpeed_v2(device Device, fan uint32, speed *uint32) Return {
	ret := nvmlDeviceGetFanSpeed_v2(nvmlDeviceHandle(device), fan, speed)
	return ret
}

func (cgoapi) DeviceGetFieldValues(device Device, valuescount int32, values *FieldValue) Return {
	ret := nvmlDeviceGetFieldValues(nvmlDeviceHandle(device), valuescount, values)
	return ret
}

func (cgoapi) DeviceGetGpcClkMinMaxVfOffset(device Device, minoffset *int32, maxoffset *int32) Return {
	ret := nvmlDeviceGetGpcClkMinMaxVfOffset(nvmlDeviceHandle(device), minoffset, maxoffset)
	return ret
}

func (cgoapi) DeviceGetGpcClkVfOffset(device Device, offset *int32) Return {
	ret := nvmlDeviceGetGpcClkVfOffset(nvmlDeviceHandle(device), offset)
	return ret
}

func (cgoapi) DeviceGetGpuFabricInfo(device Device, gpufabricinfo *GpuFabricInfo) Return {
	ret := nvmlDeviceGetGpuFabricInfo(nvmlDeviceHandle(device), gpufabricinfo)
	return ret
}

func (cgoapi) DeviceGetGpuFabricInfoV(device Device, gpufabricinfo *GpuFabricInfoV) Return {
	ret := nvmlDeviceGetGpuFabricInfoV(nvmlDeviceHandle(device), gpufabricinfo)
	return ret
}

func (cgoapi) DeviceGetGpuInstanceById(device Device, id uint32, gpuinstance *GpuInstance) Return {
	var nvmlGpuinstance nvmlGpuInstance
	ret := nvmlDeviceGetGpuInstanceById(nvmlDeviceHandle(device), id, &nvmlGpuinstance)
	*gpuinstance = GpuInstance(nvmlGpuinstance)
	return ret
}

func (cgoapi) DeviceGetGpuInstanceId(device Device, id *uint32) Return {
	ret := nvmlDeviceGetGpuInstanceId(nvmlDeviceHandle(device), id)
	return ret
}

func (cgoapi) DeviceGetGpuInstancePossiblePlacements(device Device, profileid uint32, placements *GpuInstancePlacement, count *uint32) Return {
	ret := nvmlDeviceGetGpuInstancePossiblePlacements(nvmlDeviceHandle(device), profileid, placements, count)
	return ret
}

func (cgoapi) DeviceGetGpuInstancePossiblePlacements_v1(device Device, profileid uint32, placements *GpuInstancePlacement, count *uint32) Return {
	ret := nvmlDeviceGetGpuInstancePossiblePlacements_v1(nvmlDeviceHandle(device), profileid, placements, count)
	return ret
}

func (cgoapi) DeviceGetGpuInstancePossiblePlacements_v2(device Device, profileid uint32, placements *GpuInstancePlacement, count *uint32) Return {
	ret := nvmlDeviceGetGpuInstancePossiblePlacements_v2(nvmlDeviceHandle(device), profileid, placements, count)
	return ret
}

func (cgoapi) DeviceGetGpuInstanceProfileInfo(device Device, profile uint32, info *GpuInstanceProfileInfo) Return {
	ret := nvmlDeviceGetGpuInstanceProfileInfo(nvmlDeviceHandle(device), profile, info)
	return ret
}

func (cgoapi) DeviceGetGpuInstanceProfileInfoV(device Device, profile uint32, info *GpuInstanceProfileInfo_v2) Return {
	ret := nvmlDeviceGetGpuInstanceProfileInfoV(nvmlDeviceHandle(device), profile, info)
	return ret
}

func (cgoapi) DeviceGetGpuInstanceRemainingCapacity(device Device, profileid uint32, count *uint32) Return {
	ret := nvmlDeviceGetGpuInstanceRemainingCapacity(nvmlDeviceHandle(device), profileid, count)
	return ret
}

func (cgoapi) DeviceGetGpuInstances(device Device, profileid uint32, gpuinstances *GpuInstance, count *uint32) Return {
	var nvmlGpuinstances nvmlGpuInstance
	ret := nvmlDeviceGetGpuInstances(nvmlDeviceHandle(device), profileid, &nvmlGpuinstances, count)
	*gpuinstances = GpuInstance(nvmlGpuinstances)
	return ret
}

func (cgoapi) DeviceGetGpuMaxPcieLinkGeneration(device Device, maxlinkgendevice *uint32) Return {
	ret := nvmlDeviceGetGpuMaxPcieLinkGeneration(nvmlDeviceHandle(device), maxlinkgendevice)
	return ret
}

func (cgoapi) DeviceGetGpuOperationMode(device Device, current *GpuOperationMode, pending *GpuOperationMode) Return {
	ret := nvmlDeviceGetGpuOperationMode(nvmlDeviceHandle(device), current, pending)
	return ret
}

func (cgoapi) DeviceGetGraphicsRunningProcesses(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetGraphicsRunningProcesses(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetGraphicsRunningProcesses_v1(device Device, infocount *uint32, infos *ProcessInfo_v1) Return {
	ret := nvmlDeviceGetGraphicsRunningProcesses_v1(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetGraphicsRunningProcesses_v2(device Device, infocount *uint32, infos *ProcessInfo_v2) Return {
	ret := nvmlDeviceGetGraphicsRunningProcesses_v2(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetGraphicsRunningProcesses_v3(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetGraphicsRunningProcesses_v3(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetGridLicensableFeatures(device Device, pgridlicensablefeatures *GridLicensableFeatures) Return {
	ret := nvmlDeviceGetGridLicensableFeatures(nvmlDeviceHandle(device), pgridlicensablefeatures)
	return ret
}

func (cgoapi) DeviceGetGridLicensableFeatures_v1(device Device, pgridlicensablefeatures *GridLicensableFeatures) Return {
	ret := nvmlDeviceGetGridLicensableFeatures_v1(nvmlDeviceHandle(device), pgridlicensablefeatures)
	return ret
}

func (cgoapi) DeviceGetGridLicensableFeatures_v2(device Device, pgridlicensablefeatures *GridLicensableFeatures) Return {
	ret := nvmlDeviceGetGridLicensableFeatures_v2(nvmlDeviceHandle(device), pgridlicensablefeatures)
	return ret
}

func (cgoapi) DeviceGetGridLicensableFeatures_v3(device Device, pgridlicensablefeatures *GridLicensableFeatures) Return {
	ret := nvmlDeviceGetGridLicensableFeatures_v3(nvmlDeviceHandle(device), pgridlicensablefeatures)
	return ret
}

func (cgoapi) DeviceGetGridLicensableFeatures_v4(device Device, pgridlicensablefeatures *GridLicensableFeatures) Return {
	ret := nvmlDeviceGetGridLicensableFeatures_v4(nvmlDeviceHandle(device), pgridlicensablefeatures)
	return ret
}

func (cgoapi) DeviceGetGspFirmwareMode(device Device, isenabled *uint32, defaultmode *uint32) Return {
	ret := nvmlDeviceGetGspFirmwareMode(nvmlDeviceHandle(device), isenabled, defaultmode)
	return ret
}

func (cgoapi) DeviceGetGspFirmwareVersion(device Device, version *byte) Return {
	ret := nvmlDeviceGetGspFirmwareVersion(nvmlDeviceHandle(device), version)
	return ret
}

func (cgoapi) DeviceGetHandleByIndex(index uint32, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByIndex(index, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByIndex_v1(index uint32, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByIndex_v1(index, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByIndex_v2(index uint32, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByIndex_v2(index, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByPciBusId(pcibusid string, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByPciBusId(pcibusid, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByPciBusId_v1(pcibusid string, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByPciBusId_v1(pcibusid, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByPciBusId_v2(pcibusid string, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByPciBusId_v2(pcibusid, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleBySerial(serial string, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleBySerial(serial, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByUUID(uuid string, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByUUID(uuid, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHandleByUUIDV(uuid *UUID, device *Device) Return {
	var nvmlDevice nvmlDevice
	ret := nvmlDeviceGetHandleByUUIDV(uuid, &nvmlDevice)
	*device = Device(nvmlDevice)
	return ret
}

func (cgoapi) DeviceGetHostVgpuMode(device Device, phostvgpumode *HostVgpuMode) Return {
	ret := nvmlDeviceGetHostVgpuMode(nvmlDeviceHandle(device), phostvgpumode)
	return ret
}

func (cgoapi) DeviceGetIndex(device Device, index *uint32) Return {
	ret := nvmlDeviceGetIndex(nvmlDeviceHandle(device), index)
	return ret
}

func (cgoapi) DeviceGetInforomConfigurationChecksum(device Device, checksum *uint32) Return {
	ret := nvmlDeviceGetInforomConfigurationChecksum(nvmlDeviceHandle(device), checksum)
	return ret
}

func (cgoapi) DeviceGetInforomImageVersion(device Device, version *byte, length uint32) Return {
	ret := nvmlDeviceGetInforomImageVersion(nvmlDeviceHandle(device), version, length)
	return ret
}

func (cgoapi) DeviceGetInforomVersion(device Device, object InforomObject, version *byte, length uint32) Return {
	ret := nvmlDeviceGetInforomVersion(nvmlDeviceHandle(device), object, version, length)
	return ret
}

func (cgoapi) DeviceGetIrqNum(device Device, irqnum *uint32) Return {
	ret := nvmlDeviceGetIrqNum(nvmlDeviceHandle(device), irqnum)
	return ret
}

func (cgoapi) DeviceGetJpgUtilization(device Device, utilization *uint32, samplingperiodus *uint32) Return {
	ret := nvmlDeviceGetJpgUtilization(nvmlDeviceHandle(device), utilization, samplingperiodus)
	return ret
}

func (cgoapi) DeviceGetLastBBXFlushTime(device Device, timestamp *uint64, durationus *uint) Return {
	ret := nvmlDeviceGetLastBBXFlushTime(nvmlDeviceHandle(device), timestamp, durationus)
	return ret
}

func (cgoapi) DeviceGetMPSComputeRunningProcesses(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetMPSComputeRunningProcesses(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetMPSComputeRunningProcesses_v1(device Device, infocount *uint32, infos *ProcessInfo_v1) Return {
	ret := nvmlDeviceGetMPSComputeRunningProcesses_v1(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetMPSComputeRunningProcesses_v2(device Device, infocount *uint32, infos *ProcessInfo_v2) Return {
	ret := nvmlDeviceGetMPSComputeRunningProcesses_v2(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetMPSComputeRunningProcesses_v3(device Device, infocount *uint32, infos *ProcessInfo) Return {
	ret := nvmlDeviceGetMPSComputeRunningProcesses_v3(nvmlDeviceHandle(device), infocount, infos)
	return ret
}

func (cgoapi) DeviceGetMarginTemperature(device Device, margintempinfo *MarginTemperature) Return {
	ret := nvmlDeviceGetMarginTemperature(nvmlDeviceHandle(device), margintempinfo)
	return ret
}

func (cgoapi) DeviceGetMaxClockInfo(device Device, _type ClockType, clock *uint32) Return {
	ret := nvmlDeviceGetMaxClockInfo(nvmlDeviceHandle(device), _type, clock)
	return ret
}

func (cgoapi) DeviceGetMaxCustomerBoostClock(device Device, clocktype ClockType, clockmhz *uint32) Return {
	ret := nvmlDeviceGetMaxCustomerBoostClock(nvmlDeviceHandle(device), clocktype, clockmhz)
	return ret
}

func (cgoapi) DeviceGetMaxMigDeviceCount(device Device, count *uint32) Return {
	ret := nvmlDeviceGetMaxMigDeviceCount(nvmlDeviceHandle(device), count)
	return ret
}

func (cgoapi) DeviceGetMaxPcieLinkGeneration(device Device, maxlinkgen *uint32) Return {
	ret := nvmlDeviceGetMaxPcieLinkGeneration(nvmlDeviceHandle(device), maxlinkgen)
	return ret
}

func (cgoapi) DeviceGetMaxPcieLinkWidth(device Device, maxlinkwidth *uint32) Return {
	ret := nvmlDeviceGetMaxPcieLinkWidth(nvmlDeviceHandle(device), maxlinkwidth)
	return ret
}

func (cgoapi) DeviceGetMemClkMinMaxVfOffset(device Device, minoffset *int32, maxoffset *int32) Return {
	ret := nvmlDeviceGetMemClkMinMaxVfOffset(nvmlDeviceHandle(device), minoffset, maxoffset)
	return ret
}

func (cgoapi) DeviceGetMemClkVfOffset(device Device, offset *int32) Return {
	ret := nvmlDeviceGetMemClkVfOffset(nvmlDeviceHandle(device), offset)
	return ret
}

func (cgoapi) DeviceGetMemoryAffinity(device Device, nodesetsize uint32, nodeset *uint, scope AffinityScope) Return {
	ret := nvmlDeviceGetMemoryAffinity(nvmlDeviceHandle(device), nodesetsize, nodeset, scope)
	return ret
}

func (cgoapi) DeviceGetMemoryBusWidth(device Device, buswidth *uint32) Return {
	ret := nvmlDeviceGetMemoryBusWidth(nvmlDeviceHandle(device), buswidth)
	return ret
}

func (cgoapi) DeviceGetMemoryErrorCounter(device Device, errortype MemoryErrorType, countertype EccCounterType, locationtype MemoryLocation, count *uint64) Return {
	ret := nvmlDeviceGetMemoryErrorCounter(nvmlDeviceHandle(device), errortype, countertype, locationtype, count)
	return ret
}

func (cgoapi) DeviceGetMemoryInfo(device Device, memory *Memory) Return {
	ret := nvmlDeviceGetMemoryInfo(nvmlDeviceHandle(device), memory)
	return ret
}

func (cgoapi) DeviceGetMemoryInfo_v2(device Device, memory *Memory_v2) Return {
	ret := nvmlDeviceGetMemoryInfo_v2(nvmlDeviceHandle(device), memory)
	return ret
}

func (cgoapi) DeviceGetMigDeviceHandleByIndex(device Device, index uint32, migdevice *Device) Return {
	var nvmlMigdevice nvmlDevice
	ret := nvmlDeviceGetMigDeviceHandleByIndex(nvmlDeviceHandle(device), index, &nvmlMigdevice)
	*migdevice = Device(nvmlMigdevice)
	return ret
}

func (cgoapi) DeviceGetMigMode(device Device, currentmode *uint32, pendingmode *uint32) Return {
	ret := nvmlDeviceGetMigMode(nvmlDeviceHandle(device), currentmode, pendingmode)
	return ret
}

func (cgoapi) DeviceGetMinMaxClockOfPState(device Device, _type ClockType, pstate Pstates, minclockmhz *uint32, maxclockmhz *uint32) Return {
	ret := nvmlDeviceGetMinMaxClockOfPState(nvmlDeviceHandle(device), _type, pstate, minclockmhz, maxclockmhz)
	return ret
}

func (cgoapi) DeviceGetMinMaxFanSpeed(device Device, minspeed *uint32, maxspeed *uint32) Return {
	ret := nvmlDeviceGetMinMaxFanSpeed(nvmlDeviceHandle(device), minspeed, maxspeed)
	return ret
}

func (cgoapi) DeviceGetMinorNumber(device Device, minornumber *uint32) Return {
	ret := nvmlDeviceGetMinorNumber(nvmlDeviceHandle(device), minornumber)
	return ret
}

func (cgoapi) DeviceGetModuleId(device Device, moduleid *uint32) Return {
	ret := nvmlDeviceGetModuleId(nvmlDeviceHandle(device), moduleid)
	return ret
}

func (cgoapi) DeviceGetMultiGpuBoard(device Device, multigpubool *uint32) Return {
	ret := nvmlDeviceGetMultiGpuBoard(nvmlDeviceHandle(device), multigpubool)
	return ret
}

func (cgoapi) DeviceGetName(device Device, name *byte, length uint32) Return {
	ret := nvmlDeviceGetName(nvmlDeviceHandle(device), name, length)
	return ret
}

func (cgoapi) DeviceGetNumFans(device Device, numfans *uint32) Return {
	ret := nvmlDeviceGetNumFans(nvmlDeviceHandle(device), numfans)
	return ret
}

func (cgoapi) DeviceGetNumGpuCores(device Device, numcores *uint32) Return {
	ret := nvmlDeviceGetNumGpuCores(nvmlDeviceHandle(device), numcores)
	return ret
}

func (cgoapi) DeviceGetNumaNodeId(device Device, node *uint32) Return {
	ret := nvmlDeviceGetNumaNodeId(nvmlDeviceHandle(device), node)
	return ret
}

func (cgoapi) DeviceGetNvLinkCapability(device Device, link uint32, capability NvLinkCapability, capresult *uint32) Return {
	ret := nvmlDeviceGetNvLinkCapability(nvmlDeviceHandle(device), link, capability, capresult)
	return ret
}

func (cgoapi) DeviceGetNvLinkErrorCounter(device Device, link uint32, counter NvLinkErrorCounter, countervalue *uint64) Return {
	ret := nvmlDeviceGetNvLinkErrorCounter(nvmlDeviceHandle(device), link, counter, countervalue)
	return ret
}

func (cgoapi) DeviceGetNvLinkRemoteDeviceType(device Device, link uint32, pnvlinkdevicetype *IntNvLinkDeviceType) Return {
	ret := nvmlDeviceGetNvLinkRemoteDeviceType(nvmlDeviceHandle(device), link, pnvlinkdevicetype)
	return ret
}

func (cgoapi) DeviceGetNvLinkRemotePciInfo(device Device, link uint32, pci *PciInfo) Return {
	ret := nvmlDeviceGetNvLinkRemotePciInfo(nvmlDeviceHandle(device), link, pci)
	return ret
}

func (cgoapi) DeviceGetNvLinkRemotePciInfo_v1(device Device, link uint32, pci *PciInfo) Return {
	ret := nvmlDeviceGetNvLinkRemotePciInfo_v1(nvmlDeviceHandle(device), link, pci)
	return ret
}

func (cgoapi) DeviceGetNvLinkRemotePciInfo_v2(device Device, link uint32, pci *PciInfo) Return {
	ret := nvmlDeviceGetNvLinkRemotePciInfo_v2(nvmlDeviceHandle(device), link, pci)
	return ret
}

func (cgoapi) DeviceGetNvLinkState(device Device, link uint32, isactive *EnableState) Return {
	ret := nvmlDeviceGetNvLinkState(nvmlDeviceHandle(device), link, isactive)
	return ret
}

func (cgoapi) DeviceGetNvLinkUtilizationControl(device Device, link uint32, counter uint32, control *NvLinkUtilizationControl) Return {
	ret := nvmlDeviceGetNvLinkUtilizationControl(nvmlDeviceHandle(device), link, counter, control)
	return ret
}

func (cgoapi) DeviceGetNvLinkUtilizationCounter(device Device, link uint32, counter uint32, rxcounter *uint64, txcounter *uint64) Return {
	ret := nvmlDeviceGetNvLinkUtilizationCounter(nvmlDeviceHandle(device), link, counter, rxcounter, txcounter)
	return ret
}

func (cgoapi) DeviceGetNvLinkVersion(device Device, link uint32, version *uint32) Return {
	ret := nvmlDeviceGetNvLinkVersion(nvmlDeviceHandle(device), link, version)
	return ret
}

func (cgoapi) DeviceGetNvlinkBwMode(device Device, getbwmode *NvlinkGetBwMode) Return {
	ret := nvmlDeviceGetNvlinkBwMode(nvmlDeviceHandle(device), getbwmode)
	return ret
}

func (cgoapi) DeviceGetNvlinkSupportedBwModes(device Device, supportedbwmode *NvlinkSupportedBwModes) Return {
	ret := nvmlDeviceGetNvlinkSupportedBwModes(nvmlDeviceHandle(device), supportedbwmode)
	return ret
}

func (cgoapi) DeviceGetOfaUtilization(device Device, utilization *uint32, samplingperiodus *uint32) Return {
	ret := nvmlDeviceGetOfaUtilization(nvmlDeviceHandle(device), utilization, samplingperiodus)
	return ret
}

func (cgoapi) DeviceGetP2PStatus(device1 Device, device2 Device, p2pindex GpuP2PCapsIndex, p2pstatus *GpuP2PStatus) Return {
	ret := nvmlDeviceGetP2PStatus(nvmlDeviceHandle(device1), nvmlDeviceHandle(device2), p2pindex, p2pstatus)
	return ret
}

func (cgoapi) DeviceGetPciInfo(device Device, pci *PciInfo) Return {
	ret := nvmlDeviceGetPciInfo(nvmlDeviceHandle(device), pci)
	return ret
}

func (cgoapi) DeviceGetPciInfoExt(device Device, pci *PciInfoExt) Return {
	ret := nvmlDeviceGetPciInfoExt(nvmlDeviceHandle(device), pci)
	return ret
}

func (cgoapi) DeviceGetPciInfo_v1(device Device, pci *PciInfo) Return {
	ret := nvmlDeviceGetPciInfo_v1(nvmlDeviceHandle(device), pci)
	return ret
}

func (cgoapi) DeviceGetPciInfo_v2(device Device, pci *PciInfo) Return {
	ret := nvmlDeviceGetPciInfo_v2(nvmlDeviceHandle(device), pci)
	return ret
}

func (cgoapi) DeviceGetPciInfo_v3(device Device, pci *PciInfo) Return {
	ret := nvmlDeviceGetPciInfo_v3(nvmlDeviceHandle(device), pci)
	return ret
}

func (cgoapi) DeviceGetPcieLinkMaxSpeed(device Device, maxspeed *uint32) Return {
	ret := nvmlDeviceGetPcieLinkMaxSpeed(nvmlDeviceHandle(device), maxspeed)
	return ret
}

func (cgoapi) DeviceGetPcieReplayCounter(device Device, value *uint32) Return {
	ret := nvmlDeviceGetPcieReplayCounter(nvmlDeviceHandle(device), value)
	return ret
}

func (cgoapi) DeviceGetPcieSpeed(device Device, pciespeed *uint32) Return {
	ret := nvmlDeviceGetPcieSpeed(nvmlDeviceHandle(device), pciespeed)
	return ret
}

func (cgoapi) DeviceGetPcieThroughput(device Device, counter PcieUtilCounter, value *uint32) Return {
	ret := nvmlDeviceGetPcieThroughput(nvmlDeviceHandle(device), counter, value)
	return ret
}

func (cgoapi) DeviceGetPerformanceModes(device Device, perfmodes *DevicePerfModes) Return {
	ret := nvmlDeviceGetPerformanceModes(nvmlDeviceHandle(device), perfmodes)
	return ret
}

func (cgoapi) DeviceGetPerformanceState(device Device, pstate *Pstates) Return {
	ret := nvmlDeviceGetPerformanceState(nvmlDeviceHandle(device), pstate)
	return ret
}

func (cgoapi) DeviceGetPersistenceMode(device Device, mode *EnableState) Return {
	ret := nvmlDeviceGetPersistenceMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceGetPgpuMetadataString(device Device, pgpumetadata *byte, buffersize *uint32) Return {
	ret := nvmlDeviceGetPgpuMetadataString(nvmlDeviceHandle(device), pgpumetadata, buffersize)
	return ret
}

func (cgoapi) DeviceGetPlatformInfo(device Device, platforminfo *PlatformInfo) Return {
	ret := nvmlDeviceGetPlatformInfo(nvmlDeviceHandle(device), platforminfo)
	return ret
}

func (cgoapi) DeviceGetPowerManagementDefaultLimit(device Device, defaultlimit *uint32) Return {
	ret := nvmlDeviceGetPowerManagementDefaultLimit(nvmlDeviceHandle(device), defaultlimit)
	return ret
}

func (cgoapi) DeviceGetPowerManagementLimit(device Device, limit *uint32) Return {
	ret := nvmlDeviceGetPowerManagementLimit(nvmlDeviceHandle(device), limit)
	return ret
}

func (cgoapi) DeviceGetPowerManagementLimitConstraints(device Device, minlimit *uint32, maxlimit *uint32) Return {
	ret := nvmlDeviceGetPowerManagementLimitConstraints(nvmlDeviceHandle(device), minlimit, maxlimit)
	return ret
}

func (cgoapi) DeviceGetPowerManagementMode(device Device, mode *EnableState) Return {
	ret := nvmlDeviceGetPowerManagementMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceGetPowerSource(device Device, powersource *PowerSource) Return {
	ret := nvmlDeviceGetPowerSource(nvmlDeviceHandle(device), powersource)
	return ret
}

func (cgoapi) DeviceGetPowerState(device Device, pstate *Pstates) Return {
	ret := nvmlDeviceGetPowerState(nvmlDeviceHandle(device), pstate)
	return ret
}

func (cgoapi) DeviceGetPowerUsage(device Device, power *uint32) Return {
	ret := nvmlDeviceGetPowerUsage(nvmlDeviceHandle(device), power)
	return ret
}

func (cgoapi) DeviceGetProcessUtilization(device Device, utilization *ProcessUtilizationSample, processsamplescount *uint32, lastseentimestamp uint64) Return {
	ret := nvmlDeviceGetProcessUtilization(nvmlDeviceHandle(device), utilization, processsamplescount, lastseentimestamp)
	return ret
}

func (cgoapi) DeviceGetProcessesUtilizationInfo(device Device, procesesutilinfo *ProcessesUtilizationInfo) Return {
	ret := nvmlDeviceGetProcessesUtilizationInfo(nvmlDeviceHandle(device), procesesutilinfo)
	return ret
}

func (cgoapi) DeviceGetRemappedRows(device Device, corrrows *uint32, uncrows *uint32, ispending *uint32, failureoccurred *uint32) Return {
	ret := nvmlDeviceGetRemappedRows(nvmlDeviceHandle(device), corrrows, uncrows, ispending, failureoccurred)
	return ret
}

func (cgoapi) DeviceGetRetiredPages(device Device, cause PageRetirementCause, pagecount *uint32, addresses *uint64) Return {
	ret := nvmlDeviceGetRetiredPages(nvmlDeviceHandle(device), cause, pagecount, addresses)
	return ret
}

func (cgoapi) DeviceGetRetiredPagesPendingStatus(device Device, ispending *EnableState) Return {
	ret := nvmlDeviceGetRetiredPagesPendingStatus(nvmlDeviceHandle(device), ispending)
	return ret
}

func (cgoapi) DeviceGetRetiredPages_v2(device Device, cause PageRetirementCause, pagecount *uint32, addresses *uint64, timestamps *uint64) Return {
	ret := nvmlDeviceGetRetiredPages_v2(nvmlDeviceHandle(device), cause, pagecount, addresses, timestamps)
	return ret
}

func (cgoapi) DeviceGetRowRemapperHistogram(device Device, values *RowRemapperHistogramValues) Return {
	ret := nvmlDeviceGetRowRemapperHistogram(nvmlDeviceHandle(device), values)
	return ret
}

func (cgoapi) DeviceGetRunningProcessDetailList(device Device, plist *ProcessDetailList) Return {
	ret := nvmlDeviceGetRunningProcessDetailList(nvmlDeviceHandle(device), plist)
	return ret
}

func (cgoapi) DeviceGetSamples(device Device, _type SamplingType, lastseentimestamp uint64, samplevaltype *ValueType, samplecount *uint32, samples *Sample) Return {
	ret := nvmlDeviceGetSamples(nvmlDeviceHandle(device), _type, lastseentimestamp, samplevaltype, samplecount, samples)
	return ret
}

func (cgoapi) DeviceGetSerial(device Device, serial *byte, length uint32) Return {
	ret := nvmlDeviceGetSerial(nvmlDeviceHandle(device), serial, length)
	return ret
}

func (cgoapi) DeviceGetSramEccErrorStatus(device Device, status *EccSramErrorStatus) Return {
	ret := nvmlDeviceGetSramEccErrorStatus(nvmlDeviceHandle(device), status)
	return ret
}

func (cgoapi) DeviceGetSupportedClocksEventReasons(device Device, supportedclockseventreasons *uint64) Return {
	ret := nvmlDeviceGetSupportedClocksEventReasons(nvmlDeviceHandle(device), supportedclockseventreasons)
	return ret
}

func (cgoapi) DeviceGetSupportedClocksThrottleReasons(device Device, supportedclocksthrottlereasons *uint64) Return {
	ret := nvmlDeviceGetSupportedClocksThrottleReasons(nvmlDeviceHandle(device), supportedclocksthrottlereasons)
	return ret
}

func (cgoapi) DeviceGetSupportedEventTypes(device Device, eventtypes *uint64) Return {
	ret := nvmlDeviceGetSupportedEventTypes(nvmlDeviceHandle(device), eventtypes)
	return ret
}

func (cgoapi) DeviceGetSupportedGraphicsClocks(device Device, memoryclockmhz uint32, count *uint32, clocksmhz *uint32) Return {
	ret := nvmlDeviceGetSupportedGraphicsClocks(nvmlDeviceHandle(device), memoryclockmhz, count, clocksmhz)
	return ret
}

func (cgoapi) DeviceGetSupportedMemoryClocks(device Device, count *uint32, clocksmhz *uint32) Return {
	ret := nvmlDeviceGetSupportedMemoryClocks(nvmlDeviceHandle(device), count, clocksmhz)
	return ret
}

func (cgoapi) DeviceGetSupportedPerformanceStates(device Device, pstates *Pstates, size uint32) Return {
	ret := nvmlDeviceGetSupportedPerformanceStates(nvmlDeviceHandle(device), pstates, size)
	return ret
}

func (cgoapi) DeviceGetSupportedVgpus(device Device, vgpucount *uint32, vgputypeids *VgpuTypeId) Return {
	var nvmlVgputypeids nvmlVgpuTypeId
	ret := nvmlDeviceGetSupportedVgpus(nvmlDeviceHandle(device), vgpucount, &nvmlVgputypeids)
	*vgputypeids = VgpuTypeId(nvmlVgputypeids)
	return ret
}

func (cgoapi) DeviceGetTargetFanSpeed(device Device, fan uint32, targetspeed *uint32) Return {
	ret := nvmlDeviceGetTargetFanSpeed(nvmlDeviceHandle(device), fan, targetspeed)
	return ret
}

func (cgoapi) DeviceGetTemperature(device Device, sensortype TemperatureSensors, temp *uint32) Return {
	ret := nvmlDeviceGetTemperature(nvmlDeviceHandle(device), sensortype, temp)
	return ret
}

func (cgoapi) DeviceGetTemperatureThreshold(device Device, thresholdtype TemperatureThresholds, temp *uint32) Return {
	ret := nvmlDeviceGetTemperatureThreshold(nvmlDeviceHandle(device), thresholdtype, temp)
	return ret
}

func (cgoapi) DeviceGetTemperatureV(device Device, temperature *Temperature) Return {
	ret := nvmlDeviceGetTemperatureV(nvmlDeviceHandle(device), temperature)
	return ret
}

func (cgoapi) DeviceGetThermalSettings(device Device, sensorindex uint32, pthermalsettings *GpuThermalSettings) Return {
	ret := nvmlDeviceGetThermalSettings(nvmlDeviceHandle(device), sensorindex, pthermalsettings)
	return ret
}

func (cgoapi) DeviceGetTopologyCommonAncestor(device1 Device, device2 Device, pathinfo *GpuTopologyLevel) Return {
	ret := nvmlDeviceGetTopologyCommonAncestor(nvmlDeviceHandle(device1), nvmlDeviceHandle(device2), pathinfo)
	return ret
}

func (cgoapi) DeviceGetTopologyNearestGpus(device Device, level GpuTopologyLevel, count *uint32, devicearray *Device) Return {
	var nvmlDevicearray nvmlDevice
	ret := nvmlDeviceGetTopologyNearestGpus(nvmlDeviceHandle(device), level, count, &nvmlDevicearray)
	*devicearray = Device(nvmlDevicearray)
	return ret
}

func (cgoapi) DeviceGetTotalEccErrors(device Device, errortype MemoryErrorType, countertype EccCounterType, ecccounts *uint64) Return {
	ret := nvmlDeviceGetTotalEccErrors(nvmlDeviceHandle(device), errortype, countertype, ecccounts)
	return ret
}

func (cgoapi) DeviceGetTotalEnergyConsumption(device Device, energy *uint64) Return {
	ret := nvmlDeviceGetTotalEnergyConsumption(nvmlDeviceHandle(device), energy)
	return ret
}

func (cgoapi) DeviceGetUUID(device Device, uuid *byte, length uint32) Return {
	ret := nvmlDeviceGetUUID(nvmlDeviceHandle(device), uuid, length)
	return ret
}

func (cgoapi) DeviceGetUtilizationRates(device Device, utilization *Utilization) Return {
	ret := nvmlDeviceGetUtilizationRates(nvmlDeviceHandle(device), utilization)
	return ret
}

func (cgoapi) DeviceGetVbiosVersion(device Device, version *byte, length uint32) Return {
	ret := nvmlDeviceGetVbiosVersion(nvmlDeviceHandle(device), version, length)
	return ret
}

func (cgoapi) DeviceGetVgpuCapabilities(device Device, capability DeviceVgpuCapability, capresult *uint32) Return {
	ret := nvmlDeviceGetVgpuCapabilities(nvmlDeviceHandle(device), capability, capresult)
	return ret
}

func (cgoapi) DeviceGetVgpuHeterogeneousMode(device Device, pheterogeneousmode *VgpuHeterogeneousMode) Return {
	ret := nvmlDeviceGetVgpuHeterogeneousMode(nvmlDeviceHandle(device), pheterogeneousmode)
	return ret
}

func (cgoapi) DeviceGetVgpuInstancesUtilizationInfo(device Device, vgpuutilinfo *VgpuInstancesUtilizationInfo) Return {
	ret := nvmlDeviceGetVgpuInstancesUtilizationInfo(nvmlDeviceHandle(device), vgpuutilinfo)
	return ret
}

func (cgoapi) DeviceGetVgpuMetadata(device Device, pgpumetadata *VgpuPgpuMetadata, buffersize *uint32) Return {
	ret := nvmlDeviceGetVgpuMetadata(nvmlDeviceHandle(device), &pgpumetadata.nvmlVgpuPgpuMetadata, buffersize)
	return ret
}

func (cgoapi) DeviceGetVgpuProcessUtilization(device Device, lastseentimestamp uint64, vgpuprocesssamplescount *uint32, utilizationsamples *VgpuProcessUtilizationSample) Return {
	ret := nvmlDeviceGetVgpuProcessUtilization(nvmlDeviceHandle(device), lastseentimestamp, vgpuprocesssamplescount, utilizationsamples)
	return ret
}

func (cgoapi) DeviceGetVgpuProcessesUtilizationInfo(device Device, vgpuprocutilinfo *VgpuProcessesUtilizationInfo) Return {
	ret := nvmlDeviceGetVgpuProcessesUtilizationInfo(nvmlDeviceHandle(device), vgpuprocutilinfo)
	return ret
}

func (cgoapi) DeviceGetVgpuSchedulerCapabilities(device Device, pcapabilities *VgpuSchedulerCapabilities) Return {
	ret := nvmlDeviceGetVgpuSchedulerCapabilities(nvmlDeviceHandle(device), pcapabilities)
	return ret
}

func (cgoapi) DeviceGetVgpuSchedulerLog(device Device, pschedulerlog *VgpuSchedulerLog) Return {
	ret := nvmlDeviceGetVgpuSchedulerLog(nvmlDeviceHandle(device), pschedulerlog)
	return ret
}

func (cgoapi) DeviceGetVgpuSchedulerState(device Device, pschedulerstate *VgpuSchedulerGetState) Return {
	ret := nvmlDeviceGetVgpuSchedulerState(nvmlDeviceHandle(device), pschedulerstate)
	return ret
}

func (cgoapi) DeviceGetVgpuTypeCreatablePlacements(device Device, vgputypeid VgpuTypeId, pplacementlist *VgpuPlacementList) Return {
	ret := nvmlDeviceGetVgpuTypeCreatablePlacements(nvmlDeviceHandle(device), nvmlVgpuTypeIdHandle(vgputypeid), pplacementlist)
	return ret
}

func (cgoapi) DeviceGetVgpuTypeSupportedPlacements(device Device, vgputypeid VgpuTypeId, pplacementlist *VgpuPlacementList) Return {
	ret := nvmlDeviceGetVgpuTypeSupportedPlacements(nvmlDeviceHandle(device), nvmlVgpuTypeIdHandle(vgputypeid), pplacementlist)
	return ret
}

func (cgoapi) DeviceGetVgpuUtilization(device Device, lastseentimestamp uint64, samplevaltype *ValueType, vgpuinstancesamplescount *uint32, utilizationsamples *VgpuInstanceUtilizationSample) Return {
	ret := nvmlDeviceGetVgpuUtilization(nvmlDeviceHandle(device), lastseentimestamp, samplevaltype, vgpuinstancesamplescount, utilizationsamples)
	return ret
}

func (cgoapi) DeviceGetViolationStatus(device Device, perfpolicytype PerfPolicyType, violtime *ViolationTime) Return {
	ret := nvmlDeviceGetViolationStatus(nvmlDeviceHandle(device), perfpolicytype, violtime)
	return ret
}

func (cgoapi) DeviceGetVirtualizationMode(device Device, pvirtualmode *GpuVirtualizationMode) Return {
	ret := nvmlDeviceGetVirtualizationMode(nvmlDeviceHandle(device), pvirtualmode)
	return ret
}

func (cgoapi) DeviceIsMigDeviceHandle(device Device, ismigdevice *uint32) Return {
	ret := nvmlDeviceIsMigDeviceHandle(nvmlDeviceHandle(device), ismigdevice)
	return ret
}

func (cgoapi) DeviceModifyDrainState(pciinfo *PciInfo, newstate EnableState) Return {
	ret := nvmlDeviceModifyDrainState(pciinfo, newstate)
	return ret
}

func (cgoapi) DeviceOnSameBoard(device1 Device, device2 Device, onsameboard *int32) Return {
	ret := nvmlDeviceOnSameBoard(nvmlDeviceHandle(device1), nvmlDeviceHandle(device2), onsameboard)
	return ret
}

func (cgoapi) DevicePowerSmoothingActivatePresetProfile(device Device, profile *PowerSmoothingProfile) Return {
	ret := nvmlDevicePowerSmoothingActivatePresetProfile(nvmlDeviceHandle(device), profile)
	return ret
}

func (cgoapi) DevicePowerSmoothingSetState(device Device, state *PowerSmoothingState) Return {
	ret := nvmlDevicePowerSmoothingSetState(nvmlDeviceHandle(device), state)
	return ret
}

func (cgoapi) DevicePowerSmoothingUpdatePresetProfileParam(device Device, profile *PowerSmoothingProfile) Return {
	ret := nvmlDevicePowerSmoothingUpdatePresetProfileParam(nvmlDeviceHandle(device), profile)
	return ret
}

func (cgoapi) DeviceQueryDrainState(pciinfo *PciInfo, currentstate *EnableState) Return {
	ret := nvmlDeviceQueryDrainState(pciinfo, currentstate)
	return ret
}

func (cgoapi) DeviceRegisterEvents(device Device, eventtypes uint64, set EventSet) Return {
	ret := nvmlDeviceRegisterEvents(nvmlDeviceHandle(device), eventtypes, nvmlEventSetHandle(set))
	return ret
}

func (cgoapi) DeviceRemoveGpu(pciinfo *PciInfo) Return {
	ret := nvmlDeviceRemoveGpu(pciinfo)
	return ret
}

func (cgoapi) DeviceRemoveGpu_v1(pciinfo *PciInfo) Return {
	ret := nvmlDeviceRemoveGpu_v1(pciinfo)
	return ret
}

func (cgoapi) DeviceRemoveGpu_v2(pciinfo *PciInfo, gpustate DetachGpuState, linkstate PcieLinkState) Return {
	ret := nvmlDeviceRemoveGpu_v2(pciinfo, gpustate, linkstate)
	return ret
}

func (cgoapi) DeviceResetApplicationsClocks(device Device) Return {
	ret := nvmlDeviceResetApplicationsClocks(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceResetGpuLockedClocks(device Device) Return {
	ret := nvmlDeviceResetGpuLockedClocks(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceResetMemoryLockedClocks(device Device) Return {
	ret := nvmlDeviceResetMemoryLockedClocks(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceResetNvLinkErrorCounters(device Device, link uint32) Return {
	ret := nvmlDeviceResetNvLinkErrorCounters(nvmlDeviceHandle(device), link)
	return ret
}

func (cgoapi) DeviceResetNvLinkUtilizationCounter(device Device, link uint32, counter uint32) Return {
	ret := nvmlDeviceResetNvLinkUtilizationCounter(nvmlDeviceHandle(device), link, counter)
	return ret
}

func (cgoapi) DeviceSetAPIRestriction(device Device, apitype RestrictedAPI, isrestricted EnableState) Return {
	ret := nvmlDeviceSetAPIRestriction(nvmlDeviceHandle(device), apitype, isrestricted)
	return ret
}

func (cgoapi) DeviceSetAccountingMode(device Device, mode EnableState) Return {
	ret := nvmlDeviceSetAccountingMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceSetApplicationsClocks(device Device, memclockmhz uint32, graphicsclockmhz uint32) Return {
	ret := nvmlDeviceSetApplicationsClocks(nvmlDeviceHandle(device), memclockmhz, graphicsclockmhz)
	return ret
}

func (cgoapi) DeviceSetAutoBoostedClocksEnabled(device Device, enabled EnableState) Return {
	ret := nvmlDeviceSetAutoBoostedClocksEnabled(nvmlDeviceHandle(device), enabled)
	return ret
}

func (cgoapi) DeviceSetClockOffsets(device Device, info *ClockOffset) Return {
	ret := nvmlDeviceSetClockOffsets(nvmlDeviceHandle(device), info)
	return ret
}

func (cgoapi) DeviceSetComputeMode(device Device, mode ComputeMode) Return {
	ret := nvmlDeviceSetComputeMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceSetConfComputeUnprotectedMemSize(device Device, sizekib uint64) Return {
	ret := nvmlDeviceSetConfComputeUnprotectedMemSize(nvmlDeviceHandle(device), sizekib)
	return ret
}

func (cgoapi) DeviceSetCpuAffinity(device Device) Return {
	ret := nvmlDeviceSetCpuAffinity(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceSetDefaultAutoBoostedClocksEnabled(device Device, enabled EnableState, flags uint32) Return {
	ret := nvmlDeviceSetDefaultAutoBoostedClocksEnabled(nvmlDeviceHandle(device), enabled, flags)
	return ret
}

func (cgoapi) DeviceSetDefaultFanSpeed_v2(device Device, fan uint32) Return {
	ret := nvmlDeviceSetDefaultFanSpeed_v2(nvmlDeviceHandle(device), fan)
	return ret
}

func (cgoapi) DeviceSetDramEncryptionMode(device Device, dramencryption *DramEncryptionInfo) Return {
	ret := nvmlDeviceSetDramEncryptionMode(nvmlDeviceHandle(device), dramencryption)
	return ret
}

func (cgoapi) DeviceSetDriverModel(device Device, drivermodel DriverModel, flags uint32) Return {
	ret := nvmlDeviceSetDriverModel(nvmlDeviceHandle(device), drivermodel, flags)
	return ret
}

func (cgoapi) DeviceSetEccMode(device Device, ecc EnableState) Return {
	ret := nvmlDeviceSetEccMode(nvmlDeviceHandle(device), ecc)
	return ret
}

func (cgoapi) DeviceSetFanControlPolicy(device Device, fan uint32, policy FanControlPolicy) Return {
	ret := nvmlDeviceSetFanControlPolicy(nvmlDeviceHandle(device), fan, policy)
	return ret
}

func (cgoapi) DeviceSetFanSpeed_v2(device Device, fan uint32, speed uint32) Return {
	ret := nvmlDeviceSetFanSpeed_v2(nvmlDeviceHandle(device), fan, speed)
	return ret
}

func (cgoapi) DeviceSetGpcClkVfOffset(device Device, offset int32) Return {
	ret := nvmlDeviceSetGpcClkVfOffset(nvmlDeviceHandle(device), offset)
	return ret
}

func (cgoapi) DeviceSetGpuLockedClocks(device Device, mingpuclockmhz uint32, maxgpuclockmhz uint32) Return {
	ret := nvmlDeviceSetGpuLockedClocks(nvmlDeviceHandle(device), mingpuclockmhz, maxgpuclockmhz)
	return ret
}

func (cgoapi) DeviceSetGpuOperationMode(device Device, mode GpuOperationMode) Return {
	ret := nvmlDeviceSetGpuOperationMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceSetMemClkVfOffset(device Device, offset int32) Return {
	ret := nvmlDeviceSetMemClkVfOffset(nvmlDeviceHandle(device), offset)
	return ret
}

func (cgoapi) DeviceSetMemoryLockedClocks(device Device, minmemclockmhz uint32, maxmemclockmhz uint32) Return {
	ret := nvmlDeviceSetMemoryLockedClocks(nvmlDeviceHandle(device), minmemclockmhz, maxmemclockmhz)
	return ret
}

func (cgoapi) DeviceSetMigMode(device Device, mode uint32, activationstatus *Return) Return {
	ret := nvmlDeviceSetMigMode(nvmlDeviceHandle(device), mode, activationstatus)
	return ret
}

func (cgoapi) DeviceSetNvLinkDeviceLowPowerThreshold(device Device, info *NvLinkPowerThres) Return {
	ret := nvmlDeviceSetNvLinkDeviceLowPowerThreshold(nvmlDeviceHandle(device), info)
	return ret
}

func (cgoapi) DeviceSetNvLinkUtilizationControl(device Device, link uint32, counter uint32, control *NvLinkUtilizationControl, reset uint32) Return {
	ret := nvmlDeviceSetNvLinkUtilizationControl(nvmlDeviceHandle(device), link, counter, control, reset)
	return ret
}

func (cgoapi) DeviceSetNvlinkBwMode(device Device, setbwmode *NvlinkSetBwMode) Return {
	ret := nvmlDeviceSetNvlinkBwMode(nvmlDeviceHandle(device), setbwmode)
	return ret
}

func (cgoapi) DeviceSetPersistenceMode(device Device, mode EnableState) Return {
	ret := nvmlDeviceSetPersistenceMode(nvmlDeviceHandle(device), mode)
	return ret
}

func (cgoapi) DeviceSetPowerManagementLimit(device Device, limit uint32) Return {
	ret := nvmlDeviceSetPowerManagementLimit(nvmlDeviceHandle(device), limit)
	return ret
}

func (cgoapi) DeviceSetPowerManagementLimit_v2(device Device, powervalue *PowerValue_v2) Return {
	ret := nvmlDeviceSetPowerManagementLimit_v2(nvmlDeviceHandle(device), powervalue)
	return ret
}

func (cgoapi) DeviceSetTemperatureThreshold(device Device, thresholdtype TemperatureThresholds, temp *int32) Return {
	ret := nvmlDeviceSetTemperatureThreshold(nvmlDeviceHandle(device), thresholdtype, temp)
	return ret
}

func (cgoapi) DeviceSetVgpuCapabilities(device Device, capability DeviceVgpuCapability, state EnableState) Return {
	ret := nvmlDeviceSetVgpuCapabilities(nvmlDeviceHandle(device), capability, state)
	return ret
}

func (cgoapi) DeviceSetVgpuHeterogeneousMode(device Device, pheterogeneousmode *VgpuHeterogeneousMode) Return {
	ret := nvmlDeviceSetVgpuHeterogeneousMode(nvmlDeviceHandle(device), pheterogeneousmode)
	return ret
}

func (cgoapi) DeviceSetVgpuSchedulerState(device Device, pschedulerstate *VgpuSchedulerSetState) Return {
	ret := nvmlDeviceSetVgpuSchedulerState(nvmlDeviceHandle(device), pschedulerstate)
	return ret
}

func (cgoapi) DeviceSetVirtualizationMode(device Device, virtualmode GpuVirtualizationMode) Return {
	ret := nvmlDeviceSetVirtualizationMode(nvmlDeviceHandle(device), virtualmode)
	return ret
}

func (cgoapi) DeviceValidateInforom(device Device) Return {
	ret := nvmlDeviceValidateInforom(nvmlDeviceHandle(device))
	return ret
}

func (cgoapi) DeviceWorkloadPowerProfileClearRequestedProfiles(device Device, requestedprofiles *WorkloadPowerProfileRequestedProfiles) Return {
	ret := nvmlDeviceWorkloadPowerProfileClearRequestedProfiles(nvmlDeviceHandle(device), requestedprofiles)
	return ret
}

func (cgoapi) DeviceWorkloadPowerProfileGetCurrentProfiles(device Device, currentprofiles *WorkloadPowerProfileCurrentProfiles) Return {
	ret := nvmlDeviceWorkloadPowerProfileGetCurrentProfiles(nvmlDeviceHandle(device), currentprofiles)
	return ret
}

func (cgoapi) DeviceWorkloadPowerProfileGetProfilesInfo(device Device, profilesinfo *WorkloadPowerProfileProfilesInfo) Return {
	ret := nvmlDeviceWorkloadPowerProfileGetProfilesInfo(nvmlDeviceHandle(device), profilesinfo)
	return ret
}

func (cgoapi) DeviceWorkloadPowerProfileSetRequestedProfiles(device Device, requestedprofiles *WorkloadPowerProfileRequestedProfiles) Return {
	ret := nvmlDeviceWorkloadPowerProfileSetRequestedProfiles(nvmlDeviceHandle(device), requestedprofiles)
	return ret
}

func (cgoapi) ErrorString(result Return) string {
	ret := nvmlErrorString(result)
	return ret
}

func (cgoapi) EventSetCreate(set *EventSet) Return {
	var nvmlSet nvmlEventSet
	ret := nvmlEventSetCreate(&nvmlSet)
	*set = EventSet(nvmlSet)
	return ret
}

func (cgoapi) EventSetFree(set EventSet) Return {
	ret := nvmlEventSetFree(nvmlEventSetHandle(set))
	return ret
}

func (cgoapi) EventSetWait(set EventSet, data *EventData, timeoutms uint32) Return {
	var nvmlData nvmlEventData
	ret := nvmlEventSetWait(nvmlEventSetHandle(set), &nvmlData, timeoutms)
	return ret
}

func (cgoapi) EventSetWait_v1(set EventSet, data *EventData, timeoutms uint32) Return {
	var nvmlData nvmlEventData
	ret := nvmlEventSetWait_v1(nvmlEventSetHandle(set), &nvmlData, timeoutms)
	return ret
}

func (cgoapi) EventSetWait_v2(set EventSet, data *EventData, timeoutms uint32) Return {
	var nvmlData nvmlEventData
	ret := nvmlEventSetWait_v2(nvmlEventSetHandle(set), &nvmlData, timeoutms)
	return ret
}

func (cgoapi) GetBlacklistDeviceCount(devicecount *uint32) Return {
	ret := nvmlGetBlacklistDeviceCount(devicecount)
	return ret
}

func (cgoapi) GetBlacklistDeviceInfoByIndex(index uint32, info *ExcludedDeviceInfo) Return {
	ret := nvmlGetBlacklistDeviceInfoByIndex(index, info)
	return ret
}

func (cgoapi) GetExcludedDeviceCount(devicecount *uint32) Return {
	ret := nvmlGetExcludedDeviceCount(devicecount)
	return ret
}

func (cgoapi) GetExcludedDeviceInfoByIndex(index uint32, info *ExcludedDeviceInfo) Return {
	ret := nvmlGetExcludedDeviceInfoByIndex(index, info)
	return ret
}

func (cgoapi) GetVgpuCompatibility(vgpumetadata *VgpuMetadata, pgpumetadata *VgpuPgpuMetadata, compatibilityinfo *VgpuPgpuCompatibility) Return {
	ret := nvmlGetVgpuCompatibility(&vgpumetadata.nvmlVgpuMetadata, &pgpumetadata.nvmlVgpuPgpuMetadata, compatibilityinfo)
	return ret
}

func (cgoapi) GetVgpuDriverCapabilities(capability VgpuDriverCapability, capresult *uint32) Return {
	ret := nvmlGetVgpuDriverCapabilities(capability, capresult)
	return ret
}

func (cgoapi) GetVgpuVersion(supported *VgpuVersion, current *VgpuVersion) Return {
	ret := nvmlGetVgpuVersion(supported, current)
	return ret
}

func (cgoapi) GpmMetricsGet(metricsget *GpmMetricsGetType) Return {
	var nvmlMetricsget nvmlGpmMetricsGetType
	ret := nvmlGpmMetricsGet(&nvmlMetricsget)
	return ret
}

func (cgoapi) GpmMigSampleGet(device Device, gpuinstanceid uint32, gpmsample GpmSample) Return {
	ret := nvmlGpmMigSampleGet(nvmlDeviceHandle(device), gpuinstanceid, nvmlGpmSampleHandle(gpmsample))
	return ret
}

func (cgoapi) GpmQueryDeviceSupport(device Device, gpmsupport *GpmSupport) Return {
	ret := nvmlGpmQueryDeviceSupport(nvmlDeviceHandle(device), gpmsupport)
	return ret
}

func (cgoapi) GpmQueryIfStreamingEnabled(device Device, state *uint32) Return {
	ret := nvmlGpmQueryIfStreamingEnabled(nvmlDeviceHandle(device), state)
	return ret
}

func (cgoapi) GpmSampleAlloc(gpmsample *GpmSample) Return {
	var nvmlGpmsample nvmlGpmSample
	ret := nvmlGpmSampleAlloc(&nvmlGpmsample)
	*gpmsample = GpmSample(nvmlGpmsample)
	return ret
}

func (cgoapi) GpmSampleFree(gpmsample GpmSample) Return {
	ret := nvmlGpmSampleFree(nvmlGpmSampleHandle(gpmsample))
	return ret
}

func (cgoapi) GpmSampleGet(device Device, gpmsample GpmSample) Return {
	ret := nvmlGpmSampleGet(nvmlDeviceHandle(device), nvmlGpmSampleHandle(gpmsample))
	return ret
}

func (cgoapi) GpmSetStreamingEnabled(device Device, state uint32) Return {
	ret := nvmlGpmSetStreamingEnabled(nvmlDeviceHandle(device), state)
	return ret
}

func (cgoapi) GpuInstanceCreateComputeInstance(gpuinstance GpuInstance, profileid uint32, computeinstance *ComputeInstance) Return {
	var nvmlComputeinstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstance(nvmlGpuInstanceHandle(gpuinstance), profileid, &nvmlComputeinstance)
	*computeinstance = ComputeInstance(nvmlComputeinstance)
	return ret
}

func (cgoapi) GpuInstanceCreateComputeInstanceWithPlacement(gpuinstance GpuInstance, profileid uint32, placement *ComputeInstancePlacement, computeinstance *ComputeInstance) Return {
	var nvmlComputeinstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstanceWithPlacement(nvmlGpuInstanceHandle(gpuinstance), profileid, placement, &nvmlComputeinstance)
	*computeinstance = ComputeInstance(nvmlComputeinstance)
	return ret
}

func (cgoapi) GpuInstanceDestroy(gpuinstance GpuInstance) Return {
	ret := nvmlGpuInstanceDestroy(nvmlGpuInstanceHandle(gpuinstance))
	return ret
}

func (cgoapi) GpuInstanceGetActiveVgpus(gpuinstance GpuInstance, pvgpuinstanceinfo *ActiveVgpuInstanceInfo) Return {
	ret := nvmlGpuInstanceGetActiveVgpus(nvmlGpuInstanceHandle(gpuinstance), pvgpuinstanceinfo)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstanceById(gpuinstance GpuInstance, id uint32, computeinstance *ComputeInstance) Return {
	var nvmlComputeinstance nvmlComputeInstance
	ret := nvmlGpuInstanceGetComputeInstanceById(nvmlGpuInstanceHandle(gpuinstance), id, &nvmlComputeinstance)
	*computeinstance = ComputeInstance(nvmlComputeinstance)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstancePossiblePlacements(gpuinstance GpuInstance, profileid uint32, placements *ComputeInstancePlacement, count *uint32) Return {
	ret := nvmlGpuInstanceGetComputeInstancePossiblePlacements(nvmlGpuInstanceHandle(gpuinstance), profileid, placements, count)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstanceProfileInfo(gpuinstance GpuInstance, profile uint32, engprofile uint32, info *ComputeInstanceProfileInfo) Return {
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfo(nvmlGpuInstanceHandle(gpuinstance), profile, engprofile, info)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstanceProfileInfoV(gpuinstance GpuInstance, profile uint32, engprofile uint32, info *ComputeInstanceProfileInfo_v2) Return {
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfoV(nvmlGpuInstanceHandle(gpuinstance), profile, engprofile, info)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstanceRemainingCapacity(gpuinstance GpuInstance, profileid uint32, count *uint32) Return {
	ret := nvmlGpuInstanceGetComputeInstanceRemainingCapacity(nvmlGpuInstanceHandle(gpuinstance), profileid, count)
	return ret
}

func (cgoapi) GpuInstanceGetComputeInstances(gpuinstance GpuInstance, profileid uint32, computeinstances *ComputeInstance, count *uint32) Return {
	var nvmlComputeinstances nvmlComputeInstance
	ret := nvmlGpuInstanceGetComputeInstances(nvmlGpuInstanceHandle(gpuinstance), profileid, &nvmlComputeinstances, count)
	*computeinstances = ComputeInstance(nvmlComputeinstances)
	return ret
}

func (cgoapi) GpuInstanceGetCreatableVgpus(gpuinstance GpuInstance, pvgpus *VgpuTypeIdInfo) Return {
	ret := nvmlGpuInstanceGetCreatableVgpus(nvmlGpuInstanceHandle(gpuinstance), pvgpus)
	return ret
}

func (cgoapi) GpuInstanceGetInfo(gpuinstance GpuInstance, info *GpuInstanceInfo) Return {
	var nvmlInfo nvmlGpuInstanceInfo
	ret := nvmlGpuInstanceGetInfo(nvmlGpuInstanceHandle(gpuinstance), &nvmlInfo)
	return ret
}

func (cgoapi) GpuInstanceGetVgpuHeterogeneousMode(gpuinstance GpuInstance, pheterogeneousmode *VgpuHeterogeneousMode) Return {
	ret := nvmlGpuInstanceGetVgpuHeterogeneousMode(nvmlGpuInstanceHandle(gpuinstance), pheterogeneousmode)
	return ret
}

func (cgoapi) GpuInstanceGetVgpuSchedulerLog(gpuinstance GpuInstance, pschedulerloginfo *VgpuSchedulerLogInfo) Return {
	ret := nvmlGpuInstanceGetVgpuSchedulerLog(nvmlGpuInstanceHandle(gpuinstance), pschedulerloginfo)
	return ret
}

func (cgoapi) GpuInstanceGetVgpuSchedulerState(gpuinstance GpuInstance, pschedulerstateinfo *VgpuSchedulerStateInfo) Return {
	ret := nvmlGpuInstanceGetVgpuSchedulerState(nvmlGpuInstanceHandle(gpuinstance), pschedulerstateinfo)
	return ret
}

func (cgoapi) GpuInstanceGetVgpuTypeCreatablePlacements(gpuinstance GpuInstance, pcreatableplacementinfo *VgpuCreatablePlacementInfo) Return {
	ret := nvmlGpuInstanceGetVgpuTypeCreatablePlacements(nvmlGpuInstanceHandle(gpuinstance), pcreatableplacementinfo)
	return ret
}

func (cgoapi) GpuInstanceSetVgpuHeterogeneousMode(gpuinstance GpuInstance, pheterogeneousmode *VgpuHeterogeneousMode) Return {
	ret := nvmlGpuInstanceSetVgpuHeterogeneousMode(nvmlGpuInstanceHandle(gpuinstance), pheterogeneousmode)
	return ret
}

func (cgoapi) GpuInstanceSetVgpuSchedulerState(gpuinstance GpuInstance, pscheduler *VgpuSchedulerState) Return {
	ret := nvmlGpuInstanceSetVgpuSchedulerState(nvmlGpuInstanceHandle(gpuinstance), pscheduler)
	return ret
}

func (cgoapi) Init() Return {
	ret := nvmlInit()
	return ret
}

func (cgoapi) InitWithFlags(flags uint32) Return {
	ret := nvmlInitWithFlags(flags)
	return ret
}

func (cgoapi) Init_v1() Return {
	ret := nvmlInit_v1()
	return ret
}

func (cgoapi) Init_v2() Return {
	ret := nvmlInit_v2()
	return ret
}

func (cgoapi) SetVgpuVersion(vgpuversion *VgpuVersion) Return {
	ret := nvmlSetVgpuVersion(vgpuversion)
	return ret
}

func (cgoapi) Shutdown() Return {
	ret := nvmlShutdown()
	return ret
}

func (cgoapi) SystemEventSetCreate(request *SystemEventSetCreateRequest) Return {
	ret := nvmlSystemEventSetCreate(request)
	return ret
}

func (cgoapi) SystemEventSetFree(request *SystemEventSetFreeRequest) Return {
	ret := nvmlSystemEventSetFree(request)
	return ret
}

func (cgoapi) SystemEventSetWait(request *SystemEventSetWaitRequest) Return {
	ret := nvmlSystemEventSetWait(request)
	return ret
}

func (cgoapi) SystemGetConfComputeCapabilities(capabilities *ConfComputeSystemCaps) Return {
	ret := nvmlSystemGetConfComputeCapabilities(capabilities)
	return ret
}

func (cgoapi) SystemGetConfComputeGpusReadyState(isacceptingwork *uint32) Return {
	ret := nvmlSystemGetConfComputeGpusReadyState(isacceptingwork)
	return ret
}

func (cgoapi) SystemGetConfComputeKeyRotationThresholdInfo(pkeyrotationthrinfo *ConfComputeGetKeyRotationThresholdInfo) Return {
	ret := nvmlSystemGetConfComputeKeyRotationThresholdInfo(pkeyrotationthrinfo)
	return ret
}

func (cgoapi) SystemGetConfComputeSettings(settings *SystemConfComputeSettings) Return {
	ret := nvmlSystemGetConfComputeSettings(settings)
	return ret
}

func (cgoapi) SystemGetConfComputeState(state *ConfComputeSystemState) Return {
	ret := nvmlSystemGetConfComputeState(state)
	return ret
}

func (cgoapi) SystemGetCudaDriverVersion(cudadriverversion *int32) Return {
	ret := nvmlSystemGetCudaDriverVersion(cudadriverversion)
	return ret
}

func (cgoapi) SystemGetCudaDriverVersion_v2(cudadriverversion *int32) Return {
	ret := nvmlSystemGetCudaDriverVersion_v2(cudadriverversion)
	return ret
}

func (cgoapi) SystemGetDriverBranch(branchinfo *SystemDriverBranchInfo, length uint32) Return {
	ret := nvmlSystemGetDriverBranch(branchinfo, length)
	return ret
}

func (cgoapi) SystemGetDriverVersion(version *byte, length uint32) Return {
	ret := nvmlSystemGetDriverVersion(version, length)
	return ret
}

func (cgoapi) SystemGetHicVersion(hwbccount *uint32, hwbcentries *HwbcEntry) Return {
	ret := nvmlSystemGetHicVersion(hwbccount, hwbcentries)
	return ret
}

func (cgoapi) SystemGetNVMLVersion(version *byte, length uint32) Return {
	ret := nvmlSystemGetNVMLVersion(version, length)
	return ret
}

func (cgoapi) SystemGetNvlinkBwMode(nvlinkbwmode *uint32) Return {
	ret := nvmlSystemGetNvlinkBwMode(nvlinkbwmode)
	return ret
}

func (cgoapi) SystemGetProcessName(pid uint32, name *byte, length uint32) Return {
	ret := nvmlSystemGetProcessName(pid, name, length)
	return ret
}

func (cgoapi) SystemGetTopologyGpuSet(cpunumber uint32, count *uint32, devicearray *Device) Return {
	var nvmlDevicearray nvmlDevice
	ret := nvmlSystemGetTopologyGpuSet(cpunumber, count, &nvmlDevicearray)
	*devicearray = Device(nvmlDevicearray)
	return ret
}

func (cgoapi) SystemRegisterEvents(request *SystemRegisterEventRequest) Return {
	ret := nvmlSystemRegisterEvents(request)
	return ret
}

func (cgoapi) SystemSetConfComputeGpusReadyState(isacceptingwork uint32) Return {
	ret := nvmlSystemSetConfComputeGpusReadyState(isacceptingwork)
	return ret
}

func (cgoapi) SystemSetConfComputeKeyRotationThresholdInfo(pkeyrotationthrinfo *ConfComputeSetKeyRotationThresholdInfo) Return {
	ret := nvmlSystemSetConfComputeKeyRotationThresholdInfo(pkeyrotationthrinfo)
	return ret
}

func (cgoapi) SystemSetNvlinkBwMode(nvlinkbwmode uint32) Return {
	ret := nvmlSystemSetNvlinkBwMode(nvlinkbwmode)
	return ret
}

func (cgoapi) UnitGetCount(unitcount *uint32) Return {
	ret := nvmlUnitGetCount(unitcount)
	return ret
}

func (cgoapi) UnitGetDevices(unit Unit, devicecount *uint32, devices *Device) Return {
	var nvmlDevices nvmlDevice
	ret := nvmlUnitGetDevices(nvmlUnitHandle(unit), devicecount, &nvmlDevices)
	*devices = Device(nvmlDevices)
	return ret
}

func (cgoapi) UnitGetFanSpeedInfo(unit Unit, fanspeeds *UnitFanSpeeds) Return {
	ret := nvmlUnitGetFanSpeedInfo(nvmlUnitHandle(unit), fanspeeds)
	return ret
}

func (cgoapi) UnitGetHandleByIndex(index uint32, unit *Unit) Return {
	var nvmlUnit nvmlUnit
	ret := nvmlUnitGetHandleByIndex(index, &nvmlUnit)
	*unit = Unit(nvmlUnit)
	return ret
}

func (cgoapi) UnitGetLedState(unit Unit, state *LedState) Return {
	ret := nvmlUnitGetLedState(nvmlUnitHandle(unit), state)
	return ret
}

func (cgoapi) UnitGetPsuInfo(unit Unit, psu *PSUInfo) Return {
	ret := nvmlUnitGetPsuInfo(nvmlUnitHandle(unit), psu)
	return ret
}

func (cgoapi) UnitGetTemperature(unit Unit, _type uint32, temp *uint32) Return {
	ret := nvmlUnitGetTemperature(nvmlUnitHandle(unit), _type, temp)
	return ret
}

func (cgoapi) UnitGetUnitInfo(unit Unit, info *UnitInfo) Return {
	ret := nvmlUnitGetUnitInfo(nvmlUnitHandle(unit), info)
	return ret
}

func (cgoapi) UnitSetLedState(unit Unit, color LedColor) Return {
	ret := nvmlUnitSetLedState(nvmlUnitHandle(unit), color)
	return ret
}

func (cgoapi) VgpuInstanceClearAccountingPids(vgpuinstance VgpuInstance) Return {
	ret := nvmlVgpuInstanceClearAccountingPids(nvmlVgpuInstanceHandle(vgpuinstance))
	return ret
}

func (cgoapi) VgpuInstanceGetAccountingMode(vgpuinstance VgpuInstance, mode *EnableState) Return {
	ret := nvmlVgpuInstanceGetAccountingMode(nvmlVgpuInstanceHandle(vgpuinstance), mode)
	return ret
}

func (cgoapi) VgpuInstanceGetAccountingPids(vgpuinstance VgpuInstance, count *uint32, pids *uint32) Return {
	ret := nvmlVgpuInstanceGetAccountingPids(nvmlVgpuInstanceHandle(vgpuinstance), count, pids)
	return ret
}

func (cgoapi) VgpuInstanceGetAccountingStats(vgpuinstance VgpuInstance, pid uint32, stats *AccountingStats) Return {
	ret := nvmlVgpuInstanceGetAccountingStats(nvmlVgpuInstanceHandle(vgpuinstance), pid, stats)
	return ret
}

func (cgoapi) VgpuInstanceGetEccMode(vgpuinstance VgpuInstance, eccmode *EnableState) Return {
	ret := nvmlVgpuInstanceGetEccMode(nvmlVgpuInstanceHandle(vgpuinstance), eccmode)
	return ret
}

func (cgoapi) VgpuInstanceGetEncoderCapacity(vgpuinstance VgpuInstance, encodercapacity *uint32) Return {
	ret := nvmlVgpuInstanceGetEncoderCapacity(nvmlVgpuInstanceHandle(vgpuinstance), encodercapacity)
	return ret
}

func (cgoapi) VgpuInstanceGetEncoderSessions(vgpuinstance VgpuInstance, sessioncount *uint32, sessioninfo *EncoderSessionInfo) Return {
	ret := nvmlVgpuInstanceGetEncoderSessions(nvmlVgpuInstanceHandle(vgpuinstance), sessioncount, sessioninfo)
	return ret
}

func (cgoapi) VgpuInstanceGetEncoderStats(vgpuinstance VgpuInstance, sessioncount *uint32, averagefps *uint32, averagelatency *uint32) Return {
	ret := nvmlVgpuInstanceGetEncoderStats(nvmlVgpuInstanceHandle(vgpuinstance), sessioncount, averagefps, averagelatency)
	return ret
}

func (cgoapi) VgpuInstanceGetFBCSessions(vgpuinstance VgpuInstance, sessioncount *uint32, sessioninfo *FBCSessionInfo) Return {
	ret := nvmlVgpuInstanceGetFBCSessions(nvmlVgpuInstanceHandle(vgpuinstance), sessioncount, sessioninfo)
	return ret
}

func (cgoapi) VgpuInstanceGetFBCStats(vgpuinstance VgpuInstance, fbcstats *FBCStats) Return {
	ret := nvmlVgpuInstanceGetFBCStats(nvmlVgpuInstanceHandle(vgpuinstance), fbcstats)
	return ret
}

func (cgoapi) VgpuInstanceGetFbUsage(vgpuinstance VgpuInstance, fbusage *uint64) Return {
	ret := nvmlVgpuInstanceGetFbUsage(nvmlVgpuInstanceHandle(vgpuinstance), fbusage)
	return ret
}

func (cgoapi) VgpuInstanceGetFrameRateLimit(vgpuinstance VgpuInstance, frameratelimit *uint32) Return {
	ret := nvmlVgpuInstanceGetFrameRateLimit(nvmlVgpuInstanceHandle(vgpuinstance), frameratelimit)
	return ret
}

func (cgoapi) VgpuInstanceGetGpuInstanceId(vgpuinstance VgpuInstance, gpuinstanceid *uint32) Return {
	ret := nvmlVgpuInstanceGetGpuInstanceId(nvmlVgpuInstanceHandle(vgpuinstance), gpuinstanceid)
	return ret
}

func (cgoapi) VgpuInstanceGetGpuPciId(vgpuinstance VgpuInstance, vgpupciid *byte, length *uint32) Return {
	ret := nvmlVgpuInstanceGetGpuPciId(nvmlVgpuInstanceHandle(vgpuinstance), vgpupciid, length)
	return ret
}

func (cgoapi) VgpuInstanceGetLicenseInfo(vgpuinstance VgpuInstance, licenseinfo *VgpuLicenseInfo) Return {
	ret := nvmlVgpuInstanceGetLicenseInfo(nvmlVgpuInstanceHandle(vgpuinstance), licenseinfo)
	return ret
}

func (cgoapi) VgpuInstanceGetLicenseInfo_v1(vgpuinstance VgpuInstance, licenseinfo *VgpuLicenseInfo) Return {
	ret := nvmlVgpuInstanceGetLicenseInfo_v1(nvmlVgpuInstanceHandle(vgpuinstance), licenseinfo)
	return ret
}

func (cgoapi) VgpuInstanceGetLicenseInfo_v2(vgpuinstance VgpuInstance, licenseinfo *VgpuLicenseInfo) Return {
	ret := nvmlVgpuInstanceGetLicenseInfo_v2(nvmlVgpuInstanceHandle(vgpuinstance), licenseinfo)
	return ret
}

func (cgoapi) VgpuInstanceGetLicenseStatus(vgpuinstance VgpuInstance, licensed *uint32) Return {
	ret := nvmlVgpuInstanceGetLicenseStatus(nvmlVgpuInstanceHandle(vgpuinstance), licensed)
	return ret
}

func (cgoapi) VgpuInstanceGetMdevUUID(vgpuinstance VgpuInstance, mdevuuid *byte, size uint32) Return {
	ret := nvmlVgpuInstanceGetMdevUUID(nvmlVgpuInstanceHandle(vgpuinstance), mdevuuid, size)
	return ret
}

func (cgoapi) VgpuInstanceGetMetadata(vgpuinstance VgpuInstance, vgpumetadata *VgpuMetadata, buffersize *uint32) Return {
	ret := nvmlVgpuInstanceGetMetadata(nvmlVgpuInstanceHandle(vgpuinstance), &vgpumetadata.nvmlVgpuMetadata, buffersize)
	return ret
}

func (cgoapi) VgpuInstanceGetPlacementId(vgpuinstance VgpuInstance, pplacement *VgpuPlacementId) Return {
	ret := nvmlVgpuInstanceGetPlacementId(nvmlVgpuInstanceHandle(vgpuinstance), pplacement)
	return ret
}

func (cgoapi) VgpuInstanceGetRuntimeStateSize(vgpuinstance VgpuInstance, pstate *VgpuRuntimeState) Return {
	ret := nvmlVgpuInstanceGetRuntimeStateSize(nvmlVgpuInstanceHandle(vgpuinstance), pstate)
	return ret
}

func (cgoapi) VgpuInstanceGetType(vgpuinstance VgpuInstance, vgputypeid *VgpuTypeId) Return {
	var nvmlVgputypeid nvmlVgpuTypeId
	ret := nvmlVgpuInstanceGetType(nvmlVgpuInstanceHandle(vgpuinstance), &nvmlVgputypeid)
	*vgputypeid = VgpuTypeId(nvmlVgputypeid)
	return ret
}

func (cgoapi) VgpuInstanceGetUUID(vgpuinstance VgpuInstance, uuid *byte, size uint32) Return {
	ret := nvmlVgpuInstanceGetUUID(nvmlVgpuInstanceHandle(vgpuinstance), uuid, size)
	return ret
}

func (cgoapi) VgpuInstanceGetVmDriverVersion(vgpuinstance VgpuInstance, version *byte, length uint32) Return {
	ret := nvmlVgpuInstanceGetVmDriverVersion(nvmlVgpuInstanceHandle(vgpuinstance), version, length)
	return ret
}

func (cgoapi) VgpuInstanceGetVmID(vgpuinstance VgpuInstance, vmid *byte, size uint32, vmidtype *VgpuVmIdType) Return {
	ret := nvmlVgpuInstanceGetVmID(nvmlVgpuInstanceHandle(vgpuinstance), vmid, size, vmidtype)
	return ret
}

func (cgoapi) VgpuInstanceSetEncoderCapacity(vgpuinstance VgpuInstance, encodercapacity uint32) Return {
	ret := nvmlVgpuInstanceSetEncoderCapacity(nvmlVgpuInstanceHandle(vgpuinstance), encodercapacity)
	return ret
}

func (cgoapi) VgpuTypeGetBAR1Info(vgputypeid VgpuTypeId, bar1info *VgpuTypeBar1Info) Return {
	ret := nvmlVgpuTypeGetBAR1Info(nvmlVgpuTypeIdHandle(vgputypeid), bar1info)
	return ret
}

func (cgoapi) VgpuTypeGetCapabilities(vgputypeid VgpuTypeId, capability VgpuCapability, capresult *uint32) Return {
	ret := nvmlVgpuTypeGetCapabilities(nvmlVgpuTypeIdHandle(vgputypeid), capability, capresult)
	return ret
}

func (cgoapi) VgpuTypeGetClass(vgputypeid VgpuTypeId, vgputypeclass *byte, size *uint32) Return {
	ret := nvmlVgpuTypeGetClass(nvmlVgpuTypeIdHandle(vgputypeid), vgputypeclass, size)
	return ret
}

func (cgoapi) VgpuTypeGetDeviceID(vgputypeid VgpuTypeId, deviceid *uint64, subsystemid *uint64) Return {
	ret := nvmlVgpuTypeGetDeviceID(nvmlVgpuTypeIdHandle(vgputypeid), deviceid, subsystemid)
	return ret
}

func (cgoapi) VgpuTypeGetFbReservation(vgputypeid VgpuTypeId, fbreservation *uint64) Return {
	ret := nvmlVgpuTypeGetFbReservation(nvmlVgpuTypeIdHandle(vgputypeid), fbreservation)
	return ret
}

func (cgoapi) VgpuTypeGetFrameRateLimit(vgputypeid VgpuTypeId, frameratelimit *uint32) Return {
	ret := nvmlVgpuTypeGetFrameRateLimit(nvmlVgpuTypeIdHandle(vgputypeid), frameratelimit)
	return ret
}

func (cgoapi) VgpuTypeGetFramebufferSize(vgputypeid VgpuTypeId, fbsize *uint64) Return {
	ret := nvmlVgpuTypeGetFramebufferSize(nvmlVgpuTypeIdHandle(vgputypeid), fbsize)
	return ret
}

func (cgoapi) VgpuTypeGetGpuInstanceProfileId(vgputypeid VgpuTypeId, gpuinstanceprofileid *uint32) Return {
	ret := nvmlVgpuTypeGetGpuInstanceProfileId(nvmlVgpuTypeIdHandle(vgputypeid), gpuinstanceprofileid)
	return ret
}

func (cgoapi) VgpuTypeGetGspHeapSize(vgputypeid VgpuTypeId, gspheapsize *uint64) Return {
	ret := nvmlVgpuTypeGetGspHeapSize(nvmlVgpuTypeIdHandle(vgputypeid), gspheapsize)
	return ret
}

func (cgoapi) VgpuTypeGetLicense(vgputypeid VgpuTypeId, vgputypelicensestring *byte, size uint32) Return {
	ret := nvmlVgpuTypeGetLicense(nvmlVgpuTypeIdHandle(vgputypeid), vgputypelicensestring, size)
	return ret
}

func (cgoapi) VgpuTypeGetMaxInstances(device Device, vgputypeid VgpuTypeId, vgpuinstancecount *uint32) Return {
	ret := nvmlVgpuTypeGetMaxInstances(nvmlDeviceHandle(device), nvmlVgpuTypeIdHandle(vgputypeid), vgpuinstancecount)
	return ret
}

func (cgoapi) VgpuTypeGetMaxInstancesPerGpuInstance(pmaxinstance *VgpuTypeMaxInstance) Return {
	ret := nvmlVgpuTypeGetMaxInstancesPerGpuInstance(pmaxinstance)
	return ret
}

func (cgoapi) VgpuTypeGetMaxInstancesPerVm(vgputypeid VgpuTypeId, vgpuinstancecountpervm *uint32) Return {
	ret := nvmlVgpuTypeGetMaxInstancesPerVm(nvmlVgpuTypeIdHandle(vgputypeid), vgpuinstancecountpervm)
	return ret
}

func (cgoapi) VgpuTypeGetName(vgputypeid VgpuTypeId, vgputypename *byte, size *uint32) Return {
	ret := nvmlVgpuTypeGetName(nvmlVgpuTypeIdHandle(vgputypeid), vgputypename, size)
	return ret
}

func (cgoapi) VgpuTypeGetNumDisplayHeads(vgputypeid VgpuTypeId, numdisplayheads *uint32) Return {
	ret := nvmlVgpuTypeGetNumDisplayHeads(nvmlVgpuTypeIdHandle(vgputypeid), numdisplayheads)
	return ret
}

func (cgoapi) VgpuTypeGetResolution(vgputypeid VgpuTypeId, displayindex uint32, xdim *uint32, ydim *uint32) Return {
	ret := nvmlVgpuTypeGetResolution(nvmlVgpuTypeIdHandle(vgputypeid), displayindex, xdim, ydim)
	return ret
}
