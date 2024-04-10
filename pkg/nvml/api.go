/**
# Copyright 2023 NVIDIA CORPORATION
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

// libraryOptions hold the paramaters than can be set by a LibraryOption
type libraryOptions struct {
	path  string
	flags int
}

// LibraryOption represents a functional option to configure the underlying NVML library
type LibraryOption func(*libraryOptions)

// Library defines a set of functions defined on the underlying dynamic library.
type Library interface {
	Lookup(string) error
}

// WithLibraryPath provides an option to set the library name to be used by the NVML library.
func WithLibraryPath(path string) LibraryOption {
	return func(o *libraryOptions) {
		o.path = path
	}
}

// SetLibraryOptions applies the specified options to the NVML library.
// If this is called when a library is already loaded, an error is raised.
func SetLibraryOptions(opts ...LibraryOption) error {
	libnvml.Lock()
	defer libnvml.Unlock()
	if libnvml.refcount != 0 {
		return errLibraryAlreadyLoaded
	}
	libnvml.init(opts...)
	return nil
}

// Interface represents the interface for the top-level NVML library.
type Interface interface {
	// Library API
	GetLibrary() Library
	// Init API
	Init() Return
	InitWithFlags(flags uint32) Return
	Shutdown() Return
	ErrorString(r Return) string
	// System API
	SystemGetDriverVersion() (string, Return)
	SystemGetNVMLVersion() (string, Return)
	SystemGetCudaDriverVersion() (int, Return)
	SystemGetCudaDriverVersion_v2() (int, Return)
	SystemGetProcessName(pid int) (string, Return)
	SystemGetHicVersion() ([]HwbcEntry, Return)
	SystemGetTopologyGpuSet(cpuNumber int) ([]Device, Return)
	// Device API
	DeviceGetCount() (int, Return)
	DeviceGetHandleByIndex(index int) (Device, Return)
	DeviceGetHandleBySerial(serial string) (Device, Return)
	DeviceGetHandleByUUID(uuid string) (Device, Return)
	DeviceGetHandleByPciBusId(pciBusId string) (Device, Return)
	DeviceGetName(device Device) (string, Return)
	DeviceGetBrand(device Device) (BrandType, Return)
	DeviceGetIndex(device Device) (int, Return)
	DeviceGetSerial(device Device) (string, Return)
	DeviceGetCpuAffinity(device Device, numCPUs int) ([]uint, Return)
	DeviceSetCpuAffinity(device Device) Return
	DeviceClearCpuAffinity(device Device) Return
	DeviceGetMemoryAffinity(device Device, numNodes int, scope AffinityScope) ([]uint, Return)
	DeviceGetCpuAffinityWithinScope(device Device, numCPUs int, scope AffinityScope) ([]uint, Return)
	DeviceGetTopologyCommonAncestor(device1 Device, device2 Device) (GpuTopologyLevel, Return)
	DeviceGetTopologyNearestGpus(device Device, level GpuTopologyLevel) ([]Device, Return)
	DeviceGetP2PStatus(device1 Device, device2 Device, p2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return)
	DeviceGetUUID(device Device) (string, Return)
	DeviceGetMinorNumber(device Device) (int, Return)
	DeviceGetBoardPartNumber(device Device) (string, Return)
	DeviceGetInforomVersion(device Device, object InforomObject) (string, Return)
	DeviceGetInforomImageVersion(device Device) (string, Return)
	DeviceGetInforomConfigurationChecksum(device Device) (uint32, Return)
	DeviceValidateInforom(device Device) Return
	DeviceGetDisplayMode(device Device) (EnableState, Return)
	DeviceGetDisplayActive(device Device) (EnableState, Return)
	DeviceGetPersistenceMode(device Device) (EnableState, Return)
	DeviceGetPciInfo(device Device) (PciInfo, Return)
	DeviceGetMaxPcieLinkGeneration(device Device) (int, Return)
	DeviceGetMaxPcieLinkWidth(device Device) (int, Return)
	DeviceGetCurrPcieLinkGeneration(device Device) (int, Return)
	DeviceGetCurrPcieLinkWidth(device Device) (int, Return)
	DeviceGetPcieThroughput(device Device, counter PcieUtilCounter) (uint32, Return)
	DeviceGetPcieReplayCounter(device Device) (int, Return)
	DeviceGetClockInfo(device Device, clockType ClockType) (uint32, Return)
	DeviceGetMaxClockInfo(device Device, clockType ClockType) (uint32, Return)
	DeviceGetApplicationsClock(device Device, clockType ClockType) (uint32, Return)
	DeviceGetDefaultApplicationsClock(device Device, clockType ClockType) (uint32, Return)
	DeviceResetApplicationsClocks(device Device) Return
	DeviceGetClock(device Device, clockType ClockType, clockId ClockId) (uint32, Return)
	DeviceGetMaxCustomerBoostClock(device Device, clockType ClockType) (uint32, Return)
	DeviceGetSupportedMemoryClocks(device Device) (int, uint32, Return)
	DeviceGetSupportedGraphicsClocks(device Device, memoryClockMHz int) (int, uint32, Return)
	DeviceGetAutoBoostedClocksEnabled(device Device) (EnableState, EnableState, Return)
	DeviceSetAutoBoostedClocksEnabled(device Device, enabled EnableState) Return
	DeviceSetDefaultAutoBoostedClocksEnabled(device Device, enabled EnableState, flags uint32) Return
	DeviceGetFanSpeed(device Device) (uint32, Return)
	DeviceGetFanSpeed_v2(device Device, fan int) (uint32, Return)
	DeviceGetNumFans(device Device) (int, Return)
	DeviceGetTemperature(device Device, sensorType TemperatureSensors) (uint32, Return)
	DeviceGetTemperatureThreshold(device Device, thresholdType TemperatureThresholds) (uint32, Return)
	DeviceSetTemperatureThreshold(device Device, thresholdType TemperatureThresholds, temp int) Return
	DeviceGetPerformanceState(device Device) (Pstates, Return)
	DeviceGetCurrentClocksThrottleReasons(device Device) (uint64, Return)
	DeviceGetSupportedClocksThrottleReasons(device Device) (uint64, Return)
	DeviceGetPowerState(device Device) (Pstates, Return)
	DeviceGetPowerManagementMode(device Device) (EnableState, Return)
	DeviceGetPowerManagementLimit(device Device) (uint32, Return)
	DeviceGetPowerManagementLimitConstraints(device Device) (uint32, uint32, Return)
	DeviceGetPowerManagementDefaultLimit(device Device) (uint32, Return)
	DeviceGetPowerUsage(device Device) (uint32, Return)
	DeviceGetTotalEnergyConsumption(device Device) (uint64, Return)
	DeviceGetEnforcedPowerLimit(device Device) (uint32, Return)
	DeviceGetGpuOperationMode(device Device) (GpuOperationMode, GpuOperationMode, Return)
	DeviceGetMemoryInfo(device Device) (Memory, Return)
	DeviceGetMemoryInfo_v2(device Device) (Memory_v2, Return)
	DeviceGetComputeMode(device Device) (ComputeMode, Return)
	DeviceGetCudaComputeCapability(device Device) (int, int, Return)
	DeviceGetEccMode(device Device) (EnableState, EnableState, Return)
	DeviceGetBoardId(device Device) (uint32, Return)
	DeviceGetMultiGpuBoard(device Device) (int, Return)
	DeviceGetTotalEccErrors(device Device, errorType MemoryErrorType, counterType EccCounterType) (uint64, Return)
	DeviceGetDetailedEccErrors(device Device, errorType MemoryErrorType, counterType EccCounterType) (EccErrorCounts, Return)
	DeviceGetMemoryErrorCounter(device Device, errorType MemoryErrorType, counterType EccCounterType, locationType MemoryLocation) (uint64, Return)
	DeviceGetUtilizationRates(device Device) (Utilization, Return)
	DeviceGetEncoderUtilization(device Device) (uint32, uint32, Return)
	DeviceGetEncoderCapacity(device Device, encoderQueryType EncoderType) (int, Return)
	DeviceGetEncoderStats(device Device) (int, uint32, uint32, Return)
	DeviceGetEncoderSessions(device Device) ([]EncoderSessionInfo, Return)
	DeviceGetDecoderUtilization(device Device) (uint32, uint32, Return)
	DeviceGetFBCStats(device Device) (FBCStats, Return)
	DeviceGetFBCSessions(device Device) ([]FBCSessionInfo, Return)
	DeviceGetDriverModel(device Device) (DriverModel, DriverModel, Return)
	DeviceGetVbiosVersion(device Device) (string, Return)
	DeviceGetBridgeChipInfo(device Device) (BridgeChipHierarchy, Return)
	DeviceGetComputeRunningProcesses(device Device) ([]ProcessInfo, Return)
	DeviceGetGraphicsRunningProcesses(device Device) ([]ProcessInfo, Return)
	DeviceGetMPSComputeRunningProcesses(device Device) ([]ProcessInfo, Return)
	DeviceOnSameBoard(device1 Device, device2 Device) (int, Return)
	DeviceGetAPIRestriction(device Device, apiType RestrictedAPI) (EnableState, Return)
	DeviceGetSamples(device Device, samplingType SamplingType, lastSeenTimestamp uint64) (ValueType, []Sample, Return)
	DeviceGetBAR1MemoryInfo(device Device) (BAR1Memory, Return)
	DeviceGetViolationStatus(device Device, perfPolicyType PerfPolicyType) (ViolationTime, Return)
	DeviceGetIrqNum(device Device) (int, Return)
	DeviceGetNumGpuCores(device Device) (int, Return)
	DeviceGetPowerSource(device Device) (PowerSource, Return)
	DeviceGetMemoryBusWidth(device Device) (uint32, Return)
	DeviceGetPcieLinkMaxSpeed(device Device) (uint32, Return)
	DeviceGetAdaptiveClockInfoStatus(device Device) (uint32, Return)
	DeviceGetAccountingMode(device Device) (EnableState, Return)
	DeviceGetAccountingStats(device Device, pid uint32) (AccountingStats, Return)
	DeviceGetAccountingPids(device Device) ([]int, Return)
	DeviceGetAccountingBufferSize(device Device) (int, Return)
	DeviceGetRetiredPages(device Device, cause PageRetirementCause) ([]uint64, Return)
	DeviceGetRetiredPages_v2(device Device, cause PageRetirementCause) ([]uint64, []uint64, Return)
	DeviceGetRetiredPagesPendingStatus(device Device) (EnableState, Return)
	DeviceSetPersistenceMode(device Device, mode EnableState) Return
	DeviceSetComputeMode(device Device, mode ComputeMode) Return
	DeviceSetEccMode(device Device, ecc EnableState) Return
	DeviceClearEccErrorCounts(device Device, counterType EccCounterType) Return
	DeviceSetDriverModel(device Device, driverModel DriverModel, flags uint32) Return
	DeviceSetGpuLockedClocks(device Device, minGpuClockMHz uint32, maxGpuClockMHz uint32) Return
	DeviceResetGpuLockedClocks(device Device) Return
	DeviceSetMemoryLockedClocks(device Device, minMemClockMHz uint32, maxMemClockMHz uint32) Return
	DeviceResetMemoryLockedClocks(device Device) Return
	DeviceGetClkMonStatus(device Device) (ClkMonStatus, Return)
	DeviceSetApplicationsClocks(device Device, memClockMHz uint32, graphicsClockMHz uint32) Return
	DeviceSetPowerManagementLimit(device Device, limit uint32) Return
	DeviceSetGpuOperationMode(device Device, mode GpuOperationMode) Return
	DeviceSetAPIRestriction(device Device, apiType RestrictedAPI, isRestricted EnableState) Return
	DeviceSetAccountingMode(device Device, mode EnableState) Return
	DeviceClearAccountingPids(device Device) Return
	DeviceGetNvLinkState(device Device, link int) (EnableState, Return)
	DeviceGetNvLinkVersion(device Device, link int) (uint32, Return)
	DeviceGetNvLinkCapability(device Device, link int, capability NvLinkCapability) (uint32, Return)
	DeviceGetNvLinkRemotePciInfo(device Device, link int) (PciInfo, Return)
	DeviceGetNvLinkErrorCounter(device Device, link int, counter NvLinkErrorCounter) (uint64, Return)
	DeviceResetNvLinkErrorCounters(device Device, link int) Return
	DeviceSetNvLinkUtilizationControl(device Device, link int, counter int, control *NvLinkUtilizationControl, reset bool) Return
	DeviceGetNvLinkUtilizationControl(device Device, link int, counter int) (NvLinkUtilizationControl, Return)
	DeviceGetNvLinkUtilizationCounter(device Device, link int, counter int) (uint64, uint64, Return)
	DeviceFreezeNvLinkUtilizationCounter(device Device, link int, counter int, freeze EnableState) Return
	DeviceResetNvLinkUtilizationCounter(device Device, link int, counter int) Return
	DeviceGetNvLinkRemoteDeviceType(device Device, link int) (IntNvLinkDeviceType, Return)
	DeviceRegisterEvents(device Device, eventTypes uint64, set EventSet) Return
	DeviceGetSupportedEventTypes(device Device) (uint64, Return)
	DeviceModifyDrainState(pciInfo *PciInfo, newState EnableState) Return
	DeviceQueryDrainState(pciInfo *PciInfo) (EnableState, Return)
	DeviceRemoveGpu(pciInfo *PciInfo) Return
	DeviceRemoveGpu_v2(pciInfo *PciInfo, gpuState DetachGpuState, linkState PcieLinkState) Return
	DeviceDiscoverGpus() (PciInfo, Return)
	DeviceGetFieldValues(device Device, values []FieldValue) Return
	DeviceGetVirtualizationMode(device Device) (GpuVirtualizationMode, Return)
	DeviceGetHostVgpuMode(device Device) (HostVgpuMode, Return)
	DeviceSetVirtualizationMode(device Device, virtualMode GpuVirtualizationMode) Return
	DeviceGetGridLicensableFeatures(device Device) (GridLicensableFeatures, Return)
	DeviceGetProcessUtilization(device Device, lastSeenTimestamp uint64) ([]ProcessUtilizationSample, Return)
	DeviceGetSupportedVgpus(device Device) ([]VgpuTypeId, Return)
	DeviceGetCreatableVgpus(device Device) ([]VgpuTypeId, Return)
	DeviceGetActiveVgpus(device Device) ([]VgpuInstance, Return)
	DeviceGetVgpuMetadata(device Device) (VgpuPgpuMetadata, Return)
	DeviceGetPgpuMetadataString(device Device) (string, Return)
	DeviceGetVgpuUtilization(device Device, lastSeenTimestamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return)
	DeviceGetAttributes(device Device) (DeviceAttributes, Return)
	DeviceGetRemappedRows(device Device) (int, int, bool, bool, Return)
	DeviceGetRowRemapperHistogram(device Device) (RowRemapperHistogramValues, Return)
	DeviceGetArchitecture(device Device) (DeviceArchitecture, Return)
	DeviceGetVgpuProcessUtilization(device Device, lastSeenTimestamp uint64) ([]VgpuProcessUtilizationSample, Return)
	GetExcludedDeviceCount() (int, Return)
	GetExcludedDeviceInfoByIndex(index int) (ExcludedDeviceInfo, Return)
	DeviceSetMigMode(device Device, mode int) (Return, Return)
	DeviceGetMigMode(device Device) (int, int, Return)
	DeviceGetGpuInstanceProfileInfo(device Device, profile int) (GpuInstanceProfileInfo, Return)
	DeviceGetGpuInstanceProfileInfoV(device Device, profile int) GpuInstanceProfileInfoV
	DeviceGetGpuInstancePossiblePlacements(device Device, info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return)
	DeviceGetGpuInstanceRemainingCapacity(device Device, info *GpuInstanceProfileInfo) (int, Return)
	DeviceCreateGpuInstance(device Device, info *GpuInstanceProfileInfo) (GpuInstance, Return)
	DeviceCreateGpuInstanceWithPlacement(device Device, info *GpuInstanceProfileInfo, placement *GpuInstancePlacement) (GpuInstance, Return)
	GpuInstanceDestroy(gpuInstance GpuInstance) Return
	DeviceGetGpuInstances(device Device, info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
	DeviceGetGpuInstanceById(device Device, id int) (GpuInstance, Return)
	GpuInstanceGetInfo(gpuInstance GpuInstance) (GpuInstanceInfo, Return)
	GpuInstanceGetComputeInstanceProfileInfo(gpuInstance GpuInstance, profile int, engProfile int) (ComputeInstanceProfileInfo, Return)
	GpuInstanceGetComputeInstanceProfileInfoV(gpuInstance GpuInstance, profile int, engProfile int) ComputeInstanceProfileInfoV
	GpuInstanceGetComputeInstanceRemainingCapacity(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) (int, Return)
	GpuInstanceCreateComputeInstance(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	ComputeInstanceDestroy(computeInstance ComputeInstance) Return
	GpuInstanceGetComputeInstances(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	GpuInstanceGetComputeInstanceById(gpuInstance GpuInstance, id int) (ComputeInstance, Return)
	ComputeInstanceGetInfo(computeInstance ComputeInstance) (ComputeInstanceInfo, Return)
	DeviceIsMigDeviceHandle(device Device) (bool, Return)
	DeviceGetGpuInstanceId(device Device) (int, Return)
	DeviceGetComputeInstanceId(device Device) (int, Return)
	DeviceGetMaxMigDeviceCount(device Device) (int, Return)
	DeviceGetMigDeviceHandleByIndex(device Device, index int) (Device, Return)
	DeviceGetDeviceHandleFromMigDeviceHandle(migdevice Device) (Device, Return)
	DeviceGetBusType(device Device) (BusType, Return)
	DeviceSetDefaultFanSpeed_v2(device Device, fan int) Return
	DeviceGetMinMaxFanSpeed(device Device) (int, int, Return)
	DeviceGetThermalSettings(device Device, sensorIndex uint32) (GpuThermalSettings, Return)
	DeviceGetDefaultEccMode(device Device) (EnableState, Return)
	DeviceGetPcieSpeed(device Device) (int, Return)
	DeviceGetGspFirmwareVersion(device Device) (string, Return)
	DeviceGetGspFirmwareMode(device Device) (bool, bool, Return)
	DeviceGetDynamicPstatesInfo(device Device) (GpuDynamicPstatesInfo, Return)
	DeviceSetFanSpeed_v2(device Device, fan int, speed int) Return
	DeviceGetGpcClkVfOffset(device Device) (int, Return)
	DeviceSetGpcClkVfOffset(device Device, offset int) Return
	DeviceGetMinMaxClockOfPState(device Device, clockType ClockType, pstate Pstates) (uint32, uint32, Return)
	DeviceGetSupportedPerformanceStates(device Device) ([]Pstates, Return)
	DeviceGetTargetFanSpeed(device Device, fan int) (int, Return)
	DeviceGetMemClkVfOffset(device Device) (int, Return)
	DeviceSetMemClkVfOffset(device Device, offset int) Return
	DeviceGetGpcClkMinMaxVfOffset(device Device) (int, int, Return)
	DeviceGetMemClkMinMaxVfOffset(device Device) (int, int, Return)
	DeviceGetGpuMaxPcieLinkGeneration(device Device) (int, Return)
	DeviceGetFanControlPolicy_v2(device Device, fan int) (FanControlPolicy, Return)
	DeviceSetFanControlPolicy(device Device, fan int, policy FanControlPolicy) Return
	DeviceClearFieldValues(device Device, values []FieldValue) Return
	DeviceGetVgpuCapabilities(device Device, capability DeviceVgpuCapability) (bool, Return)
	DeviceGetVgpuSchedulerLog(device Device) (VgpuSchedulerLog, Return)
	DeviceGetVgpuSchedulerState(device Device) (VgpuSchedulerGetState, Return)
	DeviceSetVgpuSchedulerState(device Device, pSchedulerState *VgpuSchedulerSetState) Return
	DeviceGetVgpuSchedulerCapabilities(device Device) (VgpuSchedulerCapabilities, Return)
	GpuInstanceGetComputeInstancePossiblePlacements(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return)
	GpuInstanceCreateComputeInstanceWithPlacement(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo, placement *ComputeInstancePlacement) (ComputeInstance, Return)
	DeviceGetGpuFabricInfo(device Device) (GpuFabricInfo, Return)
	DeviceCcuGetStreamState(device Device) (int, Return)
	DeviceCcuSetStreamState(device Device, state int) Return
	DeviceSetNvLinkDeviceLowPowerThreshold(device Device, info *NvLinkPowerThres) Return
	// Event API
	EventSetCreate() (EventSet, Return)
	EventSetWait(set EventSet, timeoutms uint32) (EventData, Return)
	EventSetFree(set EventSet) Return
	// GPM API
	GpmMetricsGetV(metricsGet *GpmMetricsGetType) GpmMetricsGetVType
	GpmMetricsGet(metricsGet *GpmMetricsGetType) Return
	GpmSampleFree(gpmSample GpmSample) Return
	GpmSampleAlloc() (GpmSample, Return)
	GpmSampleGet(device Device, gpmSample GpmSample) Return
	GpmQueryDeviceSupportV(device Device) GpmSupportV
	GpmQueryDeviceSupport(device Device) (GpmSupport, Return)
	GpmMigSampleGet(device Device, gpuInstanceId int, gpmSample GpmSample) Return
	// Unit API
	UnitGetCount() (int, Return)
	UnitGetHandleByIndex(index int) (Unit, Return)
	UnitGetUnitInfo(unit Unit) (UnitInfo, Return)
	UnitGetLedState(unit Unit) (LedState, Return)
	UnitGetPsuInfo(unit Unit) (PSUInfo, Return)
	UnitGetTemperature(unit Unit, ttype int) (uint32, Return)
	UnitGetFanSpeedInfo(unit Unit) (UnitFanSpeeds, Return)
	UnitGetDevices(unit Unit) ([]Device, Return)
	UnitSetLedState(unit Unit, color LedColor) Return
	// vGPU API
	VgpuTypeGetClass(vgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetName(vgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetGpuInstanceProfileId(vgpuTypeId VgpuTypeId) (uint32, Return)
	VgpuTypeGetDeviceID(vgpuTypeId VgpuTypeId) (uint64, uint64, Return)
	VgpuTypeGetFramebufferSize(vgpuTypeId VgpuTypeId) (uint64, Return)
	VgpuTypeGetNumDisplayHeads(vgpuTypeId VgpuTypeId) (int, Return)
	VgpuTypeGetResolution(vgpuTypeId VgpuTypeId, displayIndex int) (uint32, uint32, Return)
	VgpuTypeGetLicense(vgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetFrameRateLimit(vgpuTypeId VgpuTypeId) (uint32, Return)
	VgpuTypeGetMaxInstances(device Device, vgpuTypeId VgpuTypeId) (int, Return)
	VgpuTypeGetMaxInstancesPerVm(vgpuTypeId VgpuTypeId) (int, Return)
	VgpuInstanceGetVmID(vgpuInstance VgpuInstance) (string, VgpuVmIdType, Return)
	VgpuInstanceGetUUID(vgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetVmDriverVersion(vgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetFbUsage(vgpuInstance VgpuInstance) (uint64, Return)
	VgpuInstanceGetLicenseInfo(vgpuInstance VgpuInstance) (VgpuLicenseInfo, Return)
	VgpuInstanceGetLicenseStatus(vgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceGetType(vgpuInstance VgpuInstance) (VgpuTypeId, Return)
	VgpuInstanceGetFrameRateLimit(vgpuInstance VgpuInstance) (uint32, Return)
	VgpuInstanceGetEccMode(vgpuInstance VgpuInstance) (EnableState, Return)
	VgpuInstanceGetEncoderCapacity(vgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceSetEncoderCapacity(vgpuInstance VgpuInstance, encoderCapacity int) Return
	VgpuInstanceGetEncoderStats(vgpuInstance VgpuInstance) (int, uint32, uint32, Return)
	VgpuInstanceGetEncoderSessions(vgpuInstance VgpuInstance) (int, EncoderSessionInfo, Return)
	VgpuInstanceGetFBCStats(vgpuInstance VgpuInstance) (FBCStats, Return)
	VgpuInstanceGetFBCSessions(vgpuInstance VgpuInstance) (int, FBCSessionInfo, Return)
	VgpuInstanceGetGpuInstanceId(vgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceGetGpuPciId(vgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetMetadata(vgpuInstance VgpuInstance) (VgpuMetadata, Return)
	VgpuInstanceGetAccountingMode(vgpuInstance VgpuInstance) (EnableState, Return)
	VgpuInstanceGetAccountingPids(vgpuInstance VgpuInstance) ([]int, Return)
	VgpuInstanceGetAccountingStats(vgpuInstance VgpuInstance, pid int) (AccountingStats, Return)
	GetVgpuCompatibility(vgpuMetadata *VgpuMetadata, pgpuMetadata *VgpuPgpuMetadata) (VgpuPgpuCompatibility, Return)
	GetVgpuVersion() (VgpuVersion, VgpuVersion, Return)
	SetVgpuVersion(vgpuVersion *VgpuVersion) Return
	VgpuInstanceClearAccountingPids(vgpuInstance VgpuInstance) Return
	VgpuInstanceGetMdevUUID(vgpuInstance VgpuInstance) (string, Return)
	VgpuTypeGetCapabilities(vgpuTypeId VgpuTypeId, capability VgpuCapability) (bool, Return)
	GetVgpuDriverCapabilities(capability VgpuDriverCapability) (bool, Return)
}

type Device interface {
	GetName() (string, Return)
	GetBrand() (BrandType, Return)
	GetIndex() (int, Return)
	GetSerial() (string, Return)
	GetCpuAffinity(numCPUs int) ([]uint, Return)
	SetCpuAffinity() Return
	ClearCpuAffinity() Return
	GetMemoryAffinity(numNodes int, scope AffinityScope) ([]uint, Return)
	GetCpuAffinityWithinScope(numCPUs int, scope AffinityScope) ([]uint, Return)
	GetTopologyCommonAncestor(device2 Device) (GpuTopologyLevel, Return)
	GetTopologyNearestGpus(level GpuTopologyLevel) ([]Device, Return)
	GetP2PStatus(device2 Device, p2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return)
	GetUUID() (string, Return)
	GetMinorNumber() (int, Return)
	GetBoardPartNumber() (string, Return)
	GetInforomVersion(object InforomObject) (string, Return)
	GetInforomImageVersion() (string, Return)
	GetInforomConfigurationChecksum() (uint32, Return)
	ValidateInforom() Return
	GetDisplayMode() (EnableState, Return)
	GetDisplayActive() (EnableState, Return)
	GetPersistenceMode() (EnableState, Return)
	GetPciInfo() (PciInfo, Return)
	GetMaxPcieLinkGeneration() (int, Return)
	GetMaxPcieLinkWidth() (int, Return)
	GetCurrPcieLinkGeneration() (int, Return)
	GetCurrPcieLinkWidth() (int, Return)
	GetPcieThroughput(counter PcieUtilCounter) (uint32, Return)
	GetPcieReplayCounter() (int, Return)
	GetClockInfo(clockType ClockType) (uint32, Return)
	GetMaxClockInfo(clockType ClockType) (uint32, Return)
	GetApplicationsClock(clockType ClockType) (uint32, Return)
	GetDefaultApplicationsClock(clockType ClockType) (uint32, Return)
	ResetApplicationsClocks() Return
	GetClock(clockType ClockType, clockId ClockId) (uint32, Return)
	GetMaxCustomerBoostClock(clockType ClockType) (uint32, Return)
	GetSupportedMemoryClocks() (int, uint32, Return)
	GetSupportedGraphicsClocks(memoryClockMHz int) (int, uint32, Return)
	GetAutoBoostedClocksEnabled() (EnableState, EnableState, Return)
	SetAutoBoostedClocksEnabled(enabled EnableState) Return
	SetDefaultAutoBoostedClocksEnabled(enabled EnableState, flags uint32) Return
	GetFanSpeed() (uint32, Return)
	GetFanSpeed_v2(fan int) (uint32, Return)
	GetNumFans() (int, Return)
	GetTemperature(sensorType TemperatureSensors) (uint32, Return)
	GetTemperatureThreshold(thresholdType TemperatureThresholds) (uint32, Return)
	SetTemperatureThreshold(thresholdType TemperatureThresholds, temp int) Return
	GetPerformanceState() (Pstates, Return)
	GetCurrentClocksThrottleReasons() (uint64, Return)
	GetSupportedClocksThrottleReasons() (uint64, Return)
	GetPowerState() (Pstates, Return)
	GetPowerManagementMode() (EnableState, Return)
	GetPowerManagementLimit() (uint32, Return)
	GetPowerManagementLimitConstraints() (uint32, uint32, Return)
	GetPowerManagementDefaultLimit() (uint32, Return)
	GetPowerUsage() (uint32, Return)
	GetTotalEnergyConsumption() (uint64, Return)
	GetEnforcedPowerLimit() (uint32, Return)
	GetGpuOperationMode() (GpuOperationMode, GpuOperationMode, Return)
	GetMemoryInfo() (Memory, Return)
	GetMemoryInfo_v2() (Memory_v2, Return)
	GetComputeMode() (ComputeMode, Return)
	GetCudaComputeCapability() (int, int, Return)
	GetEccMode() (EnableState, EnableState, Return)
	GetBoardId() (uint32, Return)
	GetMultiGpuBoard() (int, Return)
	GetTotalEccErrors(errorType MemoryErrorType, counterType EccCounterType) (uint64, Return)
	GetDetailedEccErrors(errorType MemoryErrorType, counterType EccCounterType) (EccErrorCounts, Return)
	GetMemoryErrorCounter(errorType MemoryErrorType, counterType EccCounterType, locationType MemoryLocation) (uint64, Return)
	GetUtilizationRates() (Utilization, Return)
	GetEncoderUtilization() (uint32, uint32, Return)
	GetEncoderCapacity(encoderQueryType EncoderType) (int, Return)
	GetEncoderStats() (int, uint32, uint32, Return)
	GetEncoderSessions() ([]EncoderSessionInfo, Return)
	GetDecoderUtilization() (uint32, uint32, Return)
	GetFBCStats() (FBCStats, Return)
	GetFBCSessions() ([]FBCSessionInfo, Return)
	GetDriverModel() (DriverModel, DriverModel, Return)
	GetVbiosVersion() (string, Return)
	GetBridgeChipInfo() (BridgeChipHierarchy, Return)
	GetComputeRunningProcesses() ([]ProcessInfo, Return)
	GetGraphicsRunningProcesses() ([]ProcessInfo, Return)
	GetMPSComputeRunningProcesses() ([]ProcessInfo, Return)
	OnSameBoard(device2 Device) (int, Return)
	GetAPIRestriction(apiType RestrictedAPI) (EnableState, Return)
	GetSamples(samplingType SamplingType, lastSeenTimestamp uint64) (ValueType, []Sample, Return)
	GetBAR1MemoryInfo() (BAR1Memory, Return)
	GetViolationStatus(perfPolicyType PerfPolicyType) (ViolationTime, Return)
	GetIrqNum() (int, Return)
	GetNumGpuCores() (int, Return)
	GetPowerSource() (PowerSource, Return)
	GetMemoryBusWidth() (uint32, Return)
	GetPcieLinkMaxSpeed() (uint32, Return)
	GetAdaptiveClockInfoStatus() (uint32, Return)
	GetAccountingMode() (EnableState, Return)
	GetAccountingStats(pid uint32) (AccountingStats, Return)
	GetAccountingPids() ([]int, Return)
	GetAccountingBufferSize() (int, Return)
	GetRetiredPages(cause PageRetirementCause) ([]uint64, Return)
	GetRetiredPages_v2(cause PageRetirementCause) ([]uint64, []uint64, Return)
	GetRetiredPagesPendingStatus() (EnableState, Return)
	SetPersistenceMode(mode EnableState) Return
	SetComputeMode(mode ComputeMode) Return
	SetEccMode(ecc EnableState) Return
	ClearEccErrorCounts(counterType EccCounterType) Return
	SetDriverModel(driverModel DriverModel, flags uint32) Return
	SetGpuLockedClocks(minGpuClockMHz uint32, maxGpuClockMHz uint32) Return
	ResetGpuLockedClocks() Return
	SetMemoryLockedClocks(minMemClockMHz uint32, maxMemClockMHz uint32) Return
	ResetMemoryLockedClocks() Return
	GetClkMonStatus() (ClkMonStatus, Return)
	SetApplicationsClocks(memClockMHz uint32, graphicsClockMHz uint32) Return
	SetPowerManagementLimit(limit uint32) Return
	SetGpuOperationMode(mode GpuOperationMode) Return
	SetAPIRestriction(apiType RestrictedAPI, isRestricted EnableState) Return
	SetAccountingMode(mode EnableState) Return
	ClearAccountingPids() Return
	GetNvLinkState(link int) (EnableState, Return)
	GetNvLinkVersion(link int) (uint32, Return)
	GetNvLinkCapability(link int, capability NvLinkCapability) (uint32, Return)
	GetNvLinkRemotePciInfo(link int) (PciInfo, Return)
	GetNvLinkErrorCounter(link int, counter NvLinkErrorCounter) (uint64, Return)
	ResetNvLinkErrorCounters(link int) Return
	SetNvLinkUtilizationControl(link int, counter int, control *NvLinkUtilizationControl, reset bool) Return
	GetNvLinkUtilizationControl(link int, counter int) (NvLinkUtilizationControl, Return)
	GetNvLinkUtilizationCounter(link int, counter int) (uint64, uint64, Return)
	FreezeNvLinkUtilizationCounter(link int, counter int, freeze EnableState) Return
	ResetNvLinkUtilizationCounter(link int, counter int) Return
	GetNvLinkRemoteDeviceType(link int) (IntNvLinkDeviceType, Return)
	RegisterEvents(eventTypes uint64, set EventSet) Return
	GetSupportedEventTypes() (uint64, Return)
	GetFieldValues(values []FieldValue) Return
	GetVirtualizationMode() (GpuVirtualizationMode, Return)
	GetHostVgpuMode() (HostVgpuMode, Return)
	SetVirtualizationMode(virtualMode GpuVirtualizationMode) Return
	GetGridLicensableFeatures() (GridLicensableFeatures, Return)
	GetProcessUtilization(lastSeenTimestamp uint64) ([]ProcessUtilizationSample, Return)
	GetSupportedVgpus() ([]VgpuTypeId, Return)
	GetCreatableVgpus() ([]VgpuTypeId, Return)
	GetActiveVgpus() ([]VgpuInstance, Return)
	GetVgpuMetadata() (VgpuPgpuMetadata, Return)
	GetPgpuMetadataString() (string, Return)
	GetVgpuUtilization(lastSeenTimestamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return)
	GetAttributes() (DeviceAttributes, Return)
	GetRemappedRows() (int, int, bool, bool, Return)
	GetRowRemapperHistogram() (RowRemapperHistogramValues, Return)
	GetArchitecture() (DeviceArchitecture, Return)
	GetVgpuProcessUtilization(lastSeenTimestamp uint64) ([]VgpuProcessUtilizationSample, Return)
	SetMigMode(mode int) (Return, Return)
	GetMigMode() (int, int, Return)
	GetGpuInstanceProfileInfo(profile int) (GpuInstanceProfileInfo, Return)
	GetGpuInstanceProfileInfoV(profile int) GpuInstanceProfileInfoV
	GetGpuInstancePossiblePlacements(info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return)
	GetGpuInstanceRemainingCapacity(info *GpuInstanceProfileInfo) (int, Return)
	CreateGpuInstance(info *GpuInstanceProfileInfo) (GpuInstance, Return)
	CreateGpuInstanceWithPlacement(info *GpuInstanceProfileInfo, placement *GpuInstancePlacement) (GpuInstance, Return)
	GetGpuInstances(info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
	GetGpuInstanceById(id int) (GpuInstance, Return)
	IsMigDeviceHandle() (bool, Return)
	GetGpuInstanceId() (int, Return)
	GetComputeInstanceId() (int, Return)
	GetMaxMigDeviceCount() (int, Return)
	GetMigDeviceHandleByIndex(index int) (Device, Return)
	GetDeviceHandleFromMigDeviceHandle() (Device, Return)
	GetBusType() (BusType, Return)
	SetDefaultFanSpeed_v2(fan int) Return
	GetMinMaxFanSpeed() (int, int, Return)
	GetThermalSettings(sensorIndex uint32) (GpuThermalSettings, Return)
	GetDefaultEccMode() (EnableState, Return)
	GetPcieSpeed() (int, Return)
	GetGspFirmwareVersion() (string, Return)
	GetGspFirmwareMode() (bool, bool, Return)
	GetDynamicPstatesInfo() (GpuDynamicPstatesInfo, Return)
	SetFanSpeed_v2(fan int, speed int) Return
	GetGpcClkVfOffset() (int, Return)
	SetGpcClkVfOffset(offset int) Return
	GetMinMaxClockOfPState(clockType ClockType, pstate Pstates) (uint32, uint32, Return)
	GetSupportedPerformanceStates() ([]Pstates, Return)
	GetTargetFanSpeed(fan int) (int, Return)
	GetMemClkVfOffset() (int, Return)
	SetMemClkVfOffset(offset int) Return
	GetGpcClkMinMaxVfOffset() (int, int, Return)
	GetMemClkMinMaxVfOffset() (int, int, Return)
	GetGpuMaxPcieLinkGeneration() (int, Return)
	GetFanControlPolicy_v2(fan int) (FanControlPolicy, Return)
	SetFanControlPolicy(fan int, policy FanControlPolicy) Return
	ClearFieldValues(values []FieldValue) Return
	GetVgpuCapabilities(capability DeviceVgpuCapability) (bool, Return)
	GetVgpuSchedulerLog() (VgpuSchedulerLog, Return)
	GetVgpuSchedulerState() (VgpuSchedulerGetState, Return)
	SetVgpuSchedulerState(pSchedulerState *VgpuSchedulerSetState) Return
	GetVgpuSchedulerCapabilities() (VgpuSchedulerCapabilities, Return)
	GetGpuFabricInfo() (GpuFabricInfo, Return)
	CcuGetStreamState() (int, Return)
	CcuSetStreamState(state int) Return
	SetNvLinkDeviceLowPowerThreshold(info *NvLinkPowerThres) Return
	GpmSampleGet(gpmSample GpmSample) Return
	GpmQueryDeviceSupportV() GpmSupportV
	GpmQueryDeviceSupport() (GpmSupport, Return)
	GpmMigSampleGet(gpuInstanceId int, gpmSample GpmSample) Return
	VgpuTypeGetMaxInstances(vgpuTypeId VgpuTypeId) (int, Return)
}

type GpuInstance interface {
	Destroy() Return
	GetInfo() (GpuInstanceInfo, Return)
	GetComputeInstanceProfileInfo(profile int, engProfile int) (ComputeInstanceProfileInfo, Return)
	GetComputeInstanceProfileInfoV(profile int, engProfile int) ComputeInstanceProfileInfoV
	GetComputeInstanceRemainingCapacity(info *ComputeInstanceProfileInfo) (int, Return)
	CreateComputeInstance(info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	GetComputeInstances(info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	GetComputeInstanceById(id int) (ComputeInstance, Return)
	GetComputeInstancePossiblePlacements(info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return)
	CreateComputeInstanceWithPlacement(info *ComputeInstanceProfileInfo, placement *ComputeInstancePlacement) (ComputeInstance, Return)
}

type ComputeInstance interface {
	Destroy() Return
	GetInfo() (ComputeInstanceInfo, Return)
}

type EventSet interface {
	Wait(timeoutms uint32) (EventData, Return)
	Free() Return
}

type GpmSample interface {
	Free() Return
	Get(device Device) Return
	MigGet(device Device, gpuInstanceId int) Return
}

type Unit interface {
	GetUnitInfo() (UnitInfo, Return)
	GetLedState() (LedState, Return)
	GetPsuInfo() (PSUInfo, Return)
	GetTemperature(ttype int) (uint32, Return)
	GetFanSpeedInfo() (UnitFanSpeeds, Return)
	GetDevices() ([]Device, Return)
	SetLedState(color LedColor) Return
}

type VgpuInstance interface {
	GetVmID() (string, VgpuVmIdType, Return)
	GetUUID() (string, Return)
	GetVmDriverVersion() (string, Return)
	GetFbUsage() (uint64, Return)
	GetLicenseInfo() (VgpuLicenseInfo, Return)
	GetLicenseStatus() (int, Return)
	GetType() (VgpuTypeId, Return)
	GetFrameRateLimit() (uint32, Return)
	GetEccMode() (EnableState, Return)
	GetEncoderCapacity() (int, Return)
	SetEncoderCapacity(encoderCapacity int) Return
	GetEncoderStats() (int, uint32, uint32, Return)
	GetEncoderSessions() (int, EncoderSessionInfo, Return)
	GetFBCStats() (FBCStats, Return)
	GetFBCSessions() (int, FBCSessionInfo, Return)
	GetGpuInstanceId() (int, Return)
	GetGpuPciId() (string, Return)
	GetMetadata() (VgpuMetadata, Return)
	GetAccountingMode() (EnableState, Return)
	GetAccountingPids() ([]int, Return)
	GetAccountingStats(pid int) (AccountingStats, Return)
	ClearAccountingPids() Return
	GetMdevUUID() (string, Return)
}

type VgpuTypeId interface {
	GetClass() (string, Return)
	GetName() (string, Return)
	GetGpuInstanceProfileId() (uint32, Return)
	GetDeviceID() (uint64, uint64, Return)
	GetFramebufferSize() (uint64, Return)
	GetNumDisplayHeads() (int, Return)
	GetResolution(displayIndex int) (uint32, uint32, Return)
	GetLicense() (string, Return)
	GetFrameRateLimit() (uint32, Return)
	GetMaxInstances(device Device) (int, Return)
	GetMaxInstancesPerVm() (int, Return)
	GetCapabilities(capability VgpuCapability) (bool, Return)
}

// Define package level methods as aliases to Interface methods of libnvml
var (
	GetLibrary                                      = libnvml.GetLibrary
	Init                                            = libnvml.Init
	InitWithFlags                                   = libnvml.InitWithFlags
	Shutdown                                        = libnvml.Shutdown
	ErrorString                                     = libnvml.ErrorString
	SystemGetDriverVersion                          = libnvml.SystemGetDriverVersion
	SystemGetNVMLVersion                            = libnvml.SystemGetNVMLVersion
	SystemGetCudaDriverVersion                      = libnvml.SystemGetCudaDriverVersion
	SystemGetCudaDriverVersion_v2                   = libnvml.SystemGetCudaDriverVersion_v2
	SystemGetProcessName                            = libnvml.SystemGetProcessName
	SystemGetHicVersion                             = libnvml.SystemGetHicVersion
	SystemGetTopologyGpuSet                         = libnvml.SystemGetTopologyGpuSet
	DeviceGetCount                                  = libnvml.DeviceGetCount
	DeviceGetHandleByIndex                          = libnvml.DeviceGetHandleByIndex
	DeviceGetHandleBySerial                         = libnvml.DeviceGetHandleBySerial
	DeviceGetHandleByUUID                           = libnvml.DeviceGetHandleByUUID
	DeviceGetHandleByPciBusId                       = libnvml.DeviceGetHandleByPciBusId
	DeviceGetName                                   = libnvml.DeviceGetName
	DeviceGetBrand                                  = libnvml.DeviceGetBrand
	DeviceGetIndex                                  = libnvml.DeviceGetIndex
	DeviceGetSerial                                 = libnvml.DeviceGetSerial
	DeviceGetCpuAffinity                            = libnvml.DeviceGetCpuAffinity
	DeviceSetCpuAffinity                            = libnvml.DeviceSetCpuAffinity
	DeviceClearCpuAffinity                          = libnvml.DeviceClearCpuAffinity
	DeviceGetMemoryAffinity                         = libnvml.DeviceGetMemoryAffinity
	DeviceGetCpuAffinityWithinScope                 = libnvml.DeviceGetCpuAffinityWithinScope
	DeviceGetTopologyCommonAncestor                 = libnvml.DeviceGetTopologyCommonAncestor
	DeviceGetTopologyNearestGpus                    = libnvml.DeviceGetTopologyNearestGpus
	DeviceGetP2PStatus                              = libnvml.DeviceGetP2PStatus
	DeviceGetUUID                                   = libnvml.DeviceGetUUID
	DeviceGetMinorNumber                            = libnvml.DeviceGetMinorNumber
	DeviceGetBoardPartNumber                        = libnvml.DeviceGetBoardPartNumber
	DeviceGetInforomVersion                         = libnvml.DeviceGetInforomVersion
	DeviceGetInforomImageVersion                    = libnvml.DeviceGetInforomImageVersion
	DeviceGetInforomConfigurationChecksum           = libnvml.DeviceGetInforomConfigurationChecksum
	DeviceValidateInforom                           = libnvml.DeviceValidateInforom
	DeviceGetDisplayMode                            = libnvml.DeviceGetDisplayMode
	DeviceGetDisplayActive                          = libnvml.DeviceGetDisplayActive
	DeviceGetPersistenceMode                        = libnvml.DeviceGetPersistenceMode
	DeviceGetPciInfo                                = libnvml.DeviceGetPciInfo
	DeviceGetMaxPcieLinkGeneration                  = libnvml.DeviceGetMaxPcieLinkGeneration
	DeviceGetMaxPcieLinkWidth                       = libnvml.DeviceGetMaxPcieLinkWidth
	DeviceGetCurrPcieLinkGeneration                 = libnvml.DeviceGetCurrPcieLinkGeneration
	DeviceGetCurrPcieLinkWidth                      = libnvml.DeviceGetCurrPcieLinkWidth
	DeviceGetPcieThroughput                         = libnvml.DeviceGetPcieThroughput
	DeviceGetPcieReplayCounter                      = libnvml.DeviceGetPcieReplayCounter
	DeviceGetClockInfo                              = libnvml.DeviceGetClockInfo
	DeviceGetMaxClockInfo                           = libnvml.DeviceGetMaxClockInfo
	DeviceGetApplicationsClock                      = libnvml.DeviceGetApplicationsClock
	DeviceGetDefaultApplicationsClock               = libnvml.DeviceGetDefaultApplicationsClock
	DeviceResetApplicationsClocks                   = libnvml.DeviceResetApplicationsClocks
	DeviceGetClock                                  = libnvml.DeviceGetClock
	DeviceGetMaxCustomerBoostClock                  = libnvml.DeviceGetMaxCustomerBoostClock
	DeviceGetSupportedMemoryClocks                  = libnvml.DeviceGetSupportedMemoryClocks
	DeviceGetSupportedGraphicsClocks                = libnvml.DeviceGetSupportedGraphicsClocks
	DeviceGetAutoBoostedClocksEnabled               = libnvml.DeviceGetAutoBoostedClocksEnabled
	DeviceSetAutoBoostedClocksEnabled               = libnvml.DeviceSetAutoBoostedClocksEnabled
	DeviceSetDefaultAutoBoostedClocksEnabled        = libnvml.DeviceSetDefaultAutoBoostedClocksEnabled
	DeviceGetFanSpeed                               = libnvml.DeviceGetFanSpeed
	DeviceGetFanSpeed_v2                            = libnvml.DeviceGetFanSpeed_v2
	DeviceGetNumFans                                = libnvml.DeviceGetNumFans
	DeviceGetTemperature                            = libnvml.DeviceGetTemperature
	DeviceGetTemperatureThreshold                   = libnvml.DeviceGetTemperatureThreshold
	DeviceSetTemperatureThreshold                   = libnvml.DeviceSetTemperatureThreshold
	DeviceGetPerformanceState                       = libnvml.DeviceGetPerformanceState
	DeviceGetCurrentClocksThrottleReasons           = libnvml.DeviceGetCurrentClocksThrottleReasons
	DeviceGetSupportedClocksThrottleReasons         = libnvml.DeviceGetSupportedClocksThrottleReasons
	DeviceGetPowerState                             = libnvml.DeviceGetPowerState
	DeviceGetPowerManagementMode                    = libnvml.DeviceGetPowerManagementMode
	DeviceGetPowerManagementLimit                   = libnvml.DeviceGetPowerManagementLimit
	DeviceGetPowerManagementLimitConstraints        = libnvml.DeviceGetPowerManagementLimitConstraints
	DeviceGetPowerManagementDefaultLimit            = libnvml.DeviceGetPowerManagementDefaultLimit
	DeviceGetPowerUsage                             = libnvml.DeviceGetPowerUsage
	DeviceGetTotalEnergyConsumption                 = libnvml.DeviceGetTotalEnergyConsumption
	DeviceGetEnforcedPowerLimit                     = libnvml.DeviceGetEnforcedPowerLimit
	DeviceGetGpuOperationMode                       = libnvml.DeviceGetGpuOperationMode
	DeviceGetMemoryInfo                             = libnvml.DeviceGetMemoryInfo
	DeviceGetMemoryInfo_v2                          = libnvml.DeviceGetMemoryInfo_v2
	DeviceGetComputeMode                            = libnvml.DeviceGetComputeMode
	DeviceGetCudaComputeCapability                  = libnvml.DeviceGetCudaComputeCapability
	DeviceGetEccMode                                = libnvml.DeviceGetEccMode
	DeviceGetBoardId                                = libnvml.DeviceGetBoardId
	DeviceGetMultiGpuBoard                          = libnvml.DeviceGetMultiGpuBoard
	DeviceGetTotalEccErrors                         = libnvml.DeviceGetTotalEccErrors
	DeviceGetDetailedEccErrors                      = libnvml.DeviceGetDetailedEccErrors
	DeviceGetMemoryErrorCounter                     = libnvml.DeviceGetMemoryErrorCounter
	DeviceGetUtilizationRates                       = libnvml.DeviceGetUtilizationRates
	DeviceGetEncoderUtilization                     = libnvml.DeviceGetEncoderUtilization
	DeviceGetEncoderCapacity                        = libnvml.DeviceGetEncoderCapacity
	DeviceGetEncoderStats                           = libnvml.DeviceGetEncoderStats
	DeviceGetEncoderSessions                        = libnvml.DeviceGetEncoderSessions
	DeviceGetDecoderUtilization                     = libnvml.DeviceGetDecoderUtilization
	DeviceGetFBCStats                               = libnvml.DeviceGetFBCStats
	DeviceGetFBCSessions                            = libnvml.DeviceGetFBCSessions
	DeviceGetDriverModel                            = libnvml.DeviceGetDriverModel
	DeviceGetVbiosVersion                           = libnvml.DeviceGetVbiosVersion
	DeviceGetBridgeChipInfo                         = libnvml.DeviceGetBridgeChipInfo
	DeviceGetComputeRunningProcesses                = libnvml.DeviceGetComputeRunningProcesses
	DeviceGetGraphicsRunningProcesses               = libnvml.DeviceGetGraphicsRunningProcesses
	DeviceGetMPSComputeRunningProcesses             = libnvml.DeviceGetMPSComputeRunningProcesses
	DeviceOnSameBoard                               = libnvml.DeviceOnSameBoard
	DeviceGetAPIRestriction                         = libnvml.DeviceGetAPIRestriction
	DeviceGetSamples                                = libnvml.DeviceGetSamples
	DeviceGetBAR1MemoryInfo                         = libnvml.DeviceGetBAR1MemoryInfo
	DeviceGetViolationStatus                        = libnvml.DeviceGetViolationStatus
	DeviceGetIrqNum                                 = libnvml.DeviceGetIrqNum
	DeviceGetNumGpuCores                            = libnvml.DeviceGetNumGpuCores
	DeviceGetPowerSource                            = libnvml.DeviceGetPowerSource
	DeviceGetMemoryBusWidth                         = libnvml.DeviceGetMemoryBusWidth
	DeviceGetPcieLinkMaxSpeed                       = libnvml.DeviceGetPcieLinkMaxSpeed
	DeviceGetAdaptiveClockInfoStatus                = libnvml.DeviceGetAdaptiveClockInfoStatus
	DeviceGetAccountingMode                         = libnvml.DeviceGetAccountingMode
	DeviceGetAccountingStats                        = libnvml.DeviceGetAccountingStats
	DeviceGetAccountingPids                         = libnvml.DeviceGetAccountingPids
	DeviceGetAccountingBufferSize                   = libnvml.DeviceGetAccountingBufferSize
	DeviceGetRetiredPages                           = libnvml.DeviceGetRetiredPages
	DeviceGetRetiredPages_v2                        = libnvml.DeviceGetRetiredPages_v2
	DeviceGetRetiredPagesPendingStatus              = libnvml.DeviceGetRetiredPagesPendingStatus
	DeviceSetPersistenceMode                        = libnvml.DeviceSetPersistenceMode
	DeviceSetComputeMode                            = libnvml.DeviceSetComputeMode
	DeviceSetEccMode                                = libnvml.DeviceSetEccMode
	DeviceClearEccErrorCounts                       = libnvml.DeviceClearEccErrorCounts
	DeviceSetDriverModel                            = libnvml.DeviceSetDriverModel
	DeviceSetGpuLockedClocks                        = libnvml.DeviceSetGpuLockedClocks
	DeviceResetGpuLockedClocks                      = libnvml.DeviceResetGpuLockedClocks
	DeviceSetMemoryLockedClocks                     = libnvml.DeviceSetMemoryLockedClocks
	DeviceResetMemoryLockedClocks                   = libnvml.DeviceResetMemoryLockedClocks
	DeviceGetClkMonStatus                           = libnvml.DeviceGetClkMonStatus
	DeviceSetApplicationsClocks                     = libnvml.DeviceSetApplicationsClocks
	DeviceSetPowerManagementLimit                   = libnvml.DeviceSetPowerManagementLimit
	DeviceSetGpuOperationMode                       = libnvml.DeviceSetGpuOperationMode
	DeviceSetAPIRestriction                         = libnvml.DeviceSetAPIRestriction
	DeviceSetAccountingMode                         = libnvml.DeviceSetAccountingMode
	DeviceClearAccountingPids                       = libnvml.DeviceClearAccountingPids
	DeviceGetNvLinkState                            = libnvml.DeviceGetNvLinkState
	DeviceGetNvLinkVersion                          = libnvml.DeviceGetNvLinkVersion
	DeviceGetNvLinkCapability                       = libnvml.DeviceGetNvLinkCapability
	DeviceGetNvLinkRemotePciInfo                    = libnvml.DeviceGetNvLinkRemotePciInfo
	DeviceGetNvLinkErrorCounter                     = libnvml.DeviceGetNvLinkErrorCounter
	DeviceResetNvLinkErrorCounters                  = libnvml.DeviceResetNvLinkErrorCounters
	DeviceSetNvLinkUtilizationControl               = libnvml.DeviceSetNvLinkUtilizationControl
	DeviceGetNvLinkUtilizationControl               = libnvml.DeviceGetNvLinkUtilizationControl
	DeviceGetNvLinkUtilizationCounter               = libnvml.DeviceGetNvLinkUtilizationCounter
	DeviceFreezeNvLinkUtilizationCounter            = libnvml.DeviceFreezeNvLinkUtilizationCounter
	DeviceResetNvLinkUtilizationCounter             = libnvml.DeviceResetNvLinkUtilizationCounter
	DeviceGetNvLinkRemoteDeviceType                 = libnvml.DeviceGetNvLinkRemoteDeviceType
	DeviceRegisterEvents                            = libnvml.DeviceRegisterEvents
	DeviceGetSupportedEventTypes                    = libnvml.DeviceGetSupportedEventTypes
	DeviceModifyDrainState                          = libnvml.DeviceModifyDrainState
	DeviceQueryDrainState                           = libnvml.DeviceQueryDrainState
	DeviceRemoveGpu                                 = libnvml.DeviceRemoveGpu
	DeviceRemoveGpu_v2                              = libnvml.DeviceRemoveGpu_v2
	DeviceDiscoverGpus                              = libnvml.DeviceDiscoverGpus
	DeviceGetFieldValues                            = libnvml.DeviceGetFieldValues
	DeviceGetVirtualizationMode                     = libnvml.DeviceGetVirtualizationMode
	DeviceGetHostVgpuMode                           = libnvml.DeviceGetHostVgpuMode
	DeviceSetVirtualizationMode                     = libnvml.DeviceSetVirtualizationMode
	DeviceGetGridLicensableFeatures                 = libnvml.DeviceGetGridLicensableFeatures
	DeviceGetProcessUtilization                     = libnvml.DeviceGetProcessUtilization
	DeviceGetSupportedVgpus                         = libnvml.DeviceGetSupportedVgpus
	DeviceGetCreatableVgpus                         = libnvml.DeviceGetCreatableVgpus
	DeviceGetActiveVgpus                            = libnvml.DeviceGetActiveVgpus
	DeviceGetVgpuMetadata                           = libnvml.DeviceGetVgpuMetadata
	DeviceGetPgpuMetadataString                     = libnvml.DeviceGetPgpuMetadataString
	DeviceGetVgpuUtilization                        = libnvml.DeviceGetVgpuUtilization
	DeviceGetAttributes                             = libnvml.DeviceGetAttributes
	DeviceGetRemappedRows                           = libnvml.DeviceGetRemappedRows
	DeviceGetRowRemapperHistogram                   = libnvml.DeviceGetRowRemapperHistogram
	DeviceGetArchitecture                           = libnvml.DeviceGetArchitecture
	DeviceGetVgpuProcessUtilization                 = libnvml.DeviceGetVgpuProcessUtilization
	GetExcludedDeviceCount                          = libnvml.GetExcludedDeviceCount
	GetExcludedDeviceInfoByIndex                    = libnvml.GetExcludedDeviceInfoByIndex
	DeviceSetMigMode                                = libnvml.DeviceSetMigMode
	DeviceGetMigMode                                = libnvml.DeviceGetMigMode
	DeviceGetGpuInstanceProfileInfo                 = libnvml.DeviceGetGpuInstanceProfileInfo
	DeviceGetGpuInstanceProfileInfoV                = libnvml.DeviceGetGpuInstanceProfileInfoV
	DeviceGetGpuInstancePossiblePlacements          = libnvml.DeviceGetGpuInstancePossiblePlacements
	DeviceGetGpuInstanceRemainingCapacity           = libnvml.DeviceGetGpuInstanceRemainingCapacity
	DeviceCreateGpuInstance                         = libnvml.DeviceCreateGpuInstance
	DeviceCreateGpuInstanceWithPlacement            = libnvml.DeviceCreateGpuInstanceWithPlacement
	GpuInstanceDestroy                              = libnvml.GpuInstanceDestroy
	DeviceGetGpuInstances                           = libnvml.DeviceGetGpuInstances
	DeviceGetGpuInstanceById                        = libnvml.DeviceGetGpuInstanceById
	GpuInstanceGetInfo                              = libnvml.GpuInstanceGetInfo
	GpuInstanceGetComputeInstanceProfileInfo        = libnvml.GpuInstanceGetComputeInstanceProfileInfo
	GpuInstanceGetComputeInstanceProfileInfoV       = libnvml.GpuInstanceGetComputeInstanceProfileInfoV
	GpuInstanceGetComputeInstanceRemainingCapacity  = libnvml.GpuInstanceGetComputeInstanceRemainingCapacity
	GpuInstanceCreateComputeInstance                = libnvml.GpuInstanceCreateComputeInstance
	ComputeInstanceDestroy                          = libnvml.ComputeInstanceDestroy
	GpuInstanceGetComputeInstances                  = libnvml.GpuInstanceGetComputeInstances
	GpuInstanceGetComputeInstanceById               = libnvml.GpuInstanceGetComputeInstanceById
	ComputeInstanceGetInfo                          = libnvml.ComputeInstanceGetInfo
	DeviceIsMigDeviceHandle                         = libnvml.DeviceIsMigDeviceHandle
	DeviceGetGpuInstanceId                          = libnvml.DeviceGetGpuInstanceId
	DeviceGetComputeInstanceId                      = libnvml.DeviceGetComputeInstanceId
	DeviceGetMaxMigDeviceCount                      = libnvml.DeviceGetMaxMigDeviceCount
	DeviceGetMigDeviceHandleByIndex                 = libnvml.DeviceGetMigDeviceHandleByIndex
	DeviceGetDeviceHandleFromMigDeviceHandle        = libnvml.DeviceGetDeviceHandleFromMigDeviceHandle
	DeviceGetBusType                                = libnvml.DeviceGetBusType
	DeviceSetDefaultFanSpeed_v2                     = libnvml.DeviceSetDefaultFanSpeed_v2
	DeviceGetMinMaxFanSpeed                         = libnvml.DeviceGetMinMaxFanSpeed
	DeviceGetThermalSettings                        = libnvml.DeviceGetThermalSettings
	DeviceGetDefaultEccMode                         = libnvml.DeviceGetDefaultEccMode
	DeviceGetPcieSpeed                              = libnvml.DeviceGetPcieSpeed
	DeviceGetGspFirmwareVersion                     = libnvml.DeviceGetGspFirmwareVersion
	DeviceGetGspFirmwareMode                        = libnvml.DeviceGetGspFirmwareMode
	DeviceGetDynamicPstatesInfo                     = libnvml.DeviceGetDynamicPstatesInfo
	DeviceSetFanSpeed_v2                            = libnvml.DeviceSetFanSpeed_v2
	DeviceGetGpcClkVfOffset                         = libnvml.DeviceGetGpcClkVfOffset
	DeviceSetGpcClkVfOffset                         = libnvml.DeviceSetGpcClkVfOffset
	DeviceGetMinMaxClockOfPState                    = libnvml.DeviceGetMinMaxClockOfPState
	DeviceGetSupportedPerformanceStates             = libnvml.DeviceGetSupportedPerformanceStates
	DeviceGetTargetFanSpeed                         = libnvml.DeviceGetTargetFanSpeed
	DeviceGetMemClkVfOffset                         = libnvml.DeviceGetMemClkVfOffset
	DeviceSetMemClkVfOffset                         = libnvml.DeviceSetMemClkVfOffset
	DeviceGetGpcClkMinMaxVfOffset                   = libnvml.DeviceGetGpcClkMinMaxVfOffset
	DeviceGetMemClkMinMaxVfOffset                   = libnvml.DeviceGetMemClkMinMaxVfOffset
	DeviceGetGpuMaxPcieLinkGeneration               = libnvml.DeviceGetGpuMaxPcieLinkGeneration
	DeviceGetFanControlPolicy_v2                    = libnvml.DeviceGetFanControlPolicy_v2
	DeviceSetFanControlPolicy                       = libnvml.DeviceSetFanControlPolicy
	DeviceClearFieldValues                          = libnvml.DeviceClearFieldValues
	DeviceGetVgpuCapabilities                       = libnvml.DeviceGetVgpuCapabilities
	DeviceGetVgpuSchedulerLog                       = libnvml.DeviceGetVgpuSchedulerLog
	DeviceGetVgpuSchedulerState                     = libnvml.DeviceGetVgpuSchedulerState
	DeviceSetVgpuSchedulerState                     = libnvml.DeviceSetVgpuSchedulerState
	DeviceGetVgpuSchedulerCapabilities              = libnvml.DeviceGetVgpuSchedulerCapabilities
	GpuInstanceGetComputeInstancePossiblePlacements = libnvml.GpuInstanceGetComputeInstancePossiblePlacements
	GpuInstanceCreateComputeInstanceWithPlacement   = libnvml.GpuInstanceCreateComputeInstanceWithPlacement
	DeviceGetGpuFabricInfo                          = libnvml.DeviceGetGpuFabricInfo
	DeviceCcuGetStreamState                         = libnvml.DeviceCcuGetStreamState
	DeviceCcuSetStreamState                         = libnvml.DeviceCcuSetStreamState
	DeviceSetNvLinkDeviceLowPowerThreshold          = libnvml.DeviceSetNvLinkDeviceLowPowerThreshold
	EventSetCreate                                  = libnvml.EventSetCreate
	EventSetWait                                    = libnvml.EventSetWait
	EventSetFree                                    = libnvml.EventSetFree
	GpmMetricsGetV                                  = libnvml.GpmMetricsGetV
	GpmMetricsGet                                   = libnvml.GpmMetricsGet
	GpmSampleFree                                   = libnvml.GpmSampleFree
	GpmSampleAlloc                                  = libnvml.GpmSampleAlloc
	GpmSampleGet                                    = libnvml.GpmSampleGet
	GpmQueryDeviceSupportV                          = libnvml.GpmQueryDeviceSupportV
	GpmQueryDeviceSupport                           = libnvml.GpmQueryDeviceSupport
	GpmMigSampleGet                                 = libnvml.GpmMigSampleGet
	UnitGetCount                                    = libnvml.UnitGetCount
	UnitGetHandleByIndex                            = libnvml.UnitGetHandleByIndex
	UnitGetUnitInfo                                 = libnvml.UnitGetUnitInfo
	UnitGetLedState                                 = libnvml.UnitGetLedState
	UnitGetPsuInfo                                  = libnvml.UnitGetPsuInfo
	UnitGetTemperature                              = libnvml.UnitGetTemperature
	UnitGetFanSpeedInfo                             = libnvml.UnitGetFanSpeedInfo
	UnitGetDevices                                  = libnvml.UnitGetDevices
	UnitSetLedState                                 = libnvml.UnitSetLedState
	VgpuTypeGetClass                                = libnvml.VgpuTypeGetClass
	VgpuTypeGetName                                 = libnvml.VgpuTypeGetName
	VgpuTypeGetGpuInstanceProfileId                 = libnvml.VgpuTypeGetGpuInstanceProfileId
	VgpuTypeGetDeviceID                             = libnvml.VgpuTypeGetDeviceID
	VgpuTypeGetFramebufferSize                      = libnvml.VgpuTypeGetFramebufferSize
	VgpuTypeGetNumDisplayHeads                      = libnvml.VgpuTypeGetNumDisplayHeads
	VgpuTypeGetResolution                           = libnvml.VgpuTypeGetResolution
	VgpuTypeGetLicense                              = libnvml.VgpuTypeGetLicense
	VgpuTypeGetFrameRateLimit                       = libnvml.VgpuTypeGetFrameRateLimit
	VgpuTypeGetMaxInstances                         = libnvml.VgpuTypeGetMaxInstances
	VgpuTypeGetMaxInstancesPerVm                    = libnvml.VgpuTypeGetMaxInstancesPerVm
	VgpuInstanceGetVmID                             = libnvml.VgpuInstanceGetVmID
	VgpuInstanceGetUUID                             = libnvml.VgpuInstanceGetUUID
	VgpuInstanceGetVmDriverVersion                  = libnvml.VgpuInstanceGetVmDriverVersion
	VgpuInstanceGetFbUsage                          = libnvml.VgpuInstanceGetFbUsage
	VgpuInstanceGetLicenseInfo                      = libnvml.VgpuInstanceGetLicenseInfo
	VgpuInstanceGetLicenseStatus                    = libnvml.VgpuInstanceGetLicenseStatus
	VgpuInstanceGetType                             = libnvml.VgpuInstanceGetType
	VgpuInstanceGetFrameRateLimit                   = libnvml.VgpuInstanceGetFrameRateLimit
	VgpuInstanceGetEccMode                          = libnvml.VgpuInstanceGetEccMode
	VgpuInstanceGetEncoderCapacity                  = libnvml.VgpuInstanceGetEncoderCapacity
	VgpuInstanceSetEncoderCapacity                  = libnvml.VgpuInstanceSetEncoderCapacity
	VgpuInstanceGetEncoderStats                     = libnvml.VgpuInstanceGetEncoderStats
	VgpuInstanceGetEncoderSessions                  = libnvml.VgpuInstanceGetEncoderSessions
	VgpuInstanceGetFBCStats                         = libnvml.VgpuInstanceGetFBCStats
	VgpuInstanceGetFBCSessions                      = libnvml.VgpuInstanceGetFBCSessions
	VgpuInstanceGetGpuInstanceId                    = libnvml.VgpuInstanceGetGpuInstanceId
	VgpuInstanceGetGpuPciId                         = libnvml.VgpuInstanceGetGpuPciId
	VgpuInstanceGetMetadata                         = libnvml.VgpuInstanceGetMetadata
	VgpuInstanceGetAccountingMode                   = libnvml.VgpuInstanceGetAccountingMode
	VgpuInstanceGetAccountingPids                   = libnvml.VgpuInstanceGetAccountingPids
	VgpuInstanceGetAccountingStats                  = libnvml.VgpuInstanceGetAccountingStats
	GetVgpuCompatibility                            = libnvml.GetVgpuCompatibility
	GetVgpuVersion                                  = libnvml.GetVgpuVersion
	SetVgpuVersion                                  = libnvml.SetVgpuVersion
	VgpuInstanceClearAccountingPids                 = libnvml.VgpuInstanceClearAccountingPids
	VgpuInstanceGetMdevUUID                         = libnvml.VgpuInstanceGetMdevUUID
	VgpuTypeGetCapabilities                         = libnvml.VgpuTypeGetCapabilities
	GetVgpuDriverCapabilities                       = libnvml.GetVgpuDriverCapabilities
)
