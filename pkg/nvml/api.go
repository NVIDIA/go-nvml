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
	InitWithFlags(Flags uint32) Return
	Shutdown() Return
	ErrorString(r Return) string
	// System API
	SystemGetDriverVersion() (string, Return)
	SystemGetNVMLVersion() (string, Return)
	SystemGetCudaDriverVersion() (int, Return)
	SystemGetCudaDriverVersion_v2() (int, Return)
	SystemGetProcessName(Pid int) (string, Return)
	SystemGetHicVersion() ([]HwbcEntry, Return)
	SystemGetTopologyGpuSet(CpuNumber int) ([]Device, Return)
	// Device API
	DeviceGetCount() (int, Return)
	DeviceGetHandleByIndex(Index int) (Device, Return)
	DeviceGetHandleBySerial(Serial string) (Device, Return)
	DeviceGetHandleByUUID(Uuid string) (Device, Return)
	DeviceGetHandleByPciBusId(PciBusId string) (Device, Return)
	DeviceGetName(Device Device) (string, Return)
	DeviceGetBrand(Device Device) (BrandType, Return)
	DeviceGetIndex(Device Device) (int, Return)
	DeviceGetSerial(Device Device) (string, Return)
	DeviceGetCpuAffinity(Device Device, NumCPUs int) ([]uint, Return)
	DeviceSetCpuAffinity(Device Device) Return
	DeviceClearCpuAffinity(Device Device) Return
	DeviceGetMemoryAffinity(Device Device, NumNodes int, Scope AffinityScope) ([]uint, Return)
	DeviceGetCpuAffinityWithinScope(Device Device, NumCPUs int, Scope AffinityScope) ([]uint, Return)
	DeviceGetTopologyCommonAncestor(Device1 Device, Device2 Device) (GpuTopologyLevel, Return)
	DeviceGetTopologyNearestGpus(device Device, Level GpuTopologyLevel) ([]Device, Return)
	DeviceGetP2PStatus(Device1 Device, Device2 Device, P2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return)
	DeviceGetUUID(Device Device) (string, Return)
	DeviceGetMinorNumber(Device Device) (int, Return)
	DeviceGetBoardPartNumber(Device Device) (string, Return)
	DeviceGetInforomVersion(Device Device, Object InforomObject) (string, Return)
	DeviceGetInforomImageVersion(Device Device) (string, Return)
	DeviceGetInforomConfigurationChecksum(Device Device) (uint32, Return)
	DeviceValidateInforom(Device Device) Return
	DeviceGetDisplayMode(Device Device) (EnableState, Return)
	DeviceGetDisplayActive(Device Device) (EnableState, Return)
	DeviceGetPersistenceMode(Device Device) (EnableState, Return)
	DeviceGetPciInfo(Device Device) (PciInfo, Return)
	DeviceGetMaxPcieLinkGeneration(Device Device) (int, Return)
	DeviceGetMaxPcieLinkWidth(Device Device) (int, Return)
	DeviceGetCurrPcieLinkGeneration(Device Device) (int, Return)
	DeviceGetCurrPcieLinkWidth(Device Device) (int, Return)
	DeviceGetPcieThroughput(Device Device, Counter PcieUtilCounter) (uint32, Return)
	DeviceGetPcieReplayCounter(Device Device) (int, Return)
	DeviceGetClockInfo(Device Device, _type ClockType) (uint32, Return)
	DeviceGetMaxClockInfo(Device Device, _type ClockType) (uint32, Return)
	DeviceGetApplicationsClock(Device Device, ClockType ClockType) (uint32, Return)
	DeviceGetDefaultApplicationsClock(Device Device, ClockType ClockType) (uint32, Return)
	DeviceResetApplicationsClocks(Device Device) Return
	DeviceGetClock(Device Device, ClockType ClockType, ClockId ClockId) (uint32, Return)
	DeviceGetMaxCustomerBoostClock(Device Device, ClockType ClockType) (uint32, Return)
	DeviceGetSupportedMemoryClocks(Device Device) (int, uint32, Return)
	DeviceGetSupportedGraphicsClocks(Device Device, MemoryClockMHz int) (int, uint32, Return)
	DeviceGetAutoBoostedClocksEnabled(Device Device) (EnableState, EnableState, Return)
	DeviceSetAutoBoostedClocksEnabled(Device Device, Enabled EnableState) Return
	DeviceSetDefaultAutoBoostedClocksEnabled(Device Device, Enabled EnableState, Flags uint32) Return
	DeviceGetFanSpeed(Device Device) (uint32, Return)
	DeviceGetFanSpeed_v2(Device Device, Fan int) (uint32, Return)
	DeviceGetNumFans(Device Device) (int, Return)
	DeviceGetTemperature(Device Device, SensorType TemperatureSensors) (uint32, Return)
	DeviceGetTemperatureThreshold(Device Device, ThresholdType TemperatureThresholds) (uint32, Return)
	DeviceSetTemperatureThreshold(Device Device, ThresholdType TemperatureThresholds, Temp int) Return
	DeviceGetPerformanceState(Device Device) (Pstates, Return)
	DeviceGetCurrentClocksThrottleReasons(Device Device) (uint64, Return)
	DeviceGetSupportedClocksThrottleReasons(Device Device) (uint64, Return)
	DeviceGetPowerState(Device Device) (Pstates, Return)
	DeviceGetPowerManagementMode(Device Device) (EnableState, Return)
	DeviceGetPowerManagementLimit(Device Device) (uint32, Return)
	DeviceGetPowerManagementLimitConstraints(Device Device) (uint32, uint32, Return)
	DeviceGetPowerManagementDefaultLimit(Device Device) (uint32, Return)
	DeviceGetPowerUsage(Device Device) (uint32, Return)
	DeviceGetTotalEnergyConsumption(Device Device) (uint64, Return)
	DeviceGetEnforcedPowerLimit(Device Device) (uint32, Return)
	DeviceGetGpuOperationMode(Device Device) (GpuOperationMode, GpuOperationMode, Return)
	DeviceGetMemoryInfo(Device Device) (Memory, Return)
	DeviceGetMemoryInfo_v2(Device Device) (Memory_v2, Return)
	DeviceGetComputeMode(Device Device) (ComputeMode, Return)
	DeviceGetCudaComputeCapability(Device Device) (int, int, Return)
	DeviceGetEccMode(Device Device) (EnableState, EnableState, Return)
	DeviceGetBoardId(Device Device) (uint32, Return)
	DeviceGetMultiGpuBoard(Device Device) (int, Return)
	DeviceGetTotalEccErrors(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType) (uint64, Return)
	DeviceGetDetailedEccErrors(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType) (EccErrorCounts, Return)
	DeviceGetMemoryErrorCounter(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType, LocationType MemoryLocation) (uint64, Return)
	DeviceGetUtilizationRates(Device Device) (Utilization, Return)
	DeviceGetEncoderUtilization(Device Device) (uint32, uint32, Return)
	DeviceGetEncoderCapacity(Device Device, EncoderQueryType EncoderType) (int, Return)
	DeviceGetEncoderStats(Device Device) (int, uint32, uint32, Return)
	DeviceGetEncoderSessions(Device Device) ([]EncoderSessionInfo, Return)
	DeviceGetDecoderUtilization(Device Device) (uint32, uint32, Return)
	DeviceGetFBCStats(Device Device) (FBCStats, Return)
	DeviceGetFBCSessions(Device Device) ([]FBCSessionInfo, Return)
	DeviceGetDriverModel(Device Device) (DriverModel, DriverModel, Return)
	DeviceGetVbiosVersion(Device Device) (string, Return)
	DeviceGetBridgeChipInfo(Device Device) (BridgeChipHierarchy, Return)
	DeviceOnSameBoard(Device1 Device, Device2 Device) (int, Return)
	DeviceGetAPIRestriction(Device Device, ApiType RestrictedAPI) (EnableState, Return)
	DeviceGetSamples(Device Device, _type SamplingType, LastSeenTimeStamp uint64) (ValueType, []Sample, Return)
	DeviceGetBAR1MemoryInfo(Device Device) (BAR1Memory, Return)
	DeviceGetViolationStatus(Device Device, PerfPolicyType PerfPolicyType) (ViolationTime, Return)
	DeviceGetIrqNum(Device Device) (int, Return)
	DeviceGetNumGpuCores(Device Device) (int, Return)
	DeviceGetPowerSource(Device Device) (PowerSource, Return)
	DeviceGetMemoryBusWidth(Device Device) (uint32, Return)
	DeviceGetPcieLinkMaxSpeed(Device Device) (uint32, Return)
	DeviceGetAdaptiveClockInfoStatus(Device Device) (uint32, Return)
	DeviceGetAccountingMode(Device Device) (EnableState, Return)
	DeviceGetAccountingStats(Device Device, Pid uint32) (AccountingStats, Return)
	DeviceGetAccountingPids(Device Device) ([]int, Return)
	DeviceGetAccountingBufferSize(Device Device) (int, Return)
	DeviceGetRetiredPages(Device Device, Cause PageRetirementCause) ([]uint64, Return)
	DeviceGetRetiredPages_v2(Device Device, Cause PageRetirementCause) ([]uint64, []uint64, Return)
	DeviceGetRetiredPagesPendingStatus(Device Device) (EnableState, Return)
	DeviceSetPersistenceMode(Device Device, Mode EnableState) Return
	DeviceSetComputeMode(Device Device, Mode ComputeMode) Return
	DeviceSetEccMode(Device Device, Ecc EnableState) Return
	DeviceClearEccErrorCounts(Device Device, CounterType EccCounterType) Return
	DeviceSetDriverModel(Device Device, DriverModel DriverModel, Flags uint32) Return
	DeviceSetGpuLockedClocks(Device Device, MinGpuClockMHz uint32, MaxGpuClockMHz uint32) Return
	DeviceResetGpuLockedClocks(Device Device) Return
	DeviceSetMemoryLockedClocks(Device Device, MinMemClockMHz uint32, MaxMemClockMHz uint32) Return
	DeviceResetMemoryLockedClocks(Device Device) Return
	DeviceGetClkMonStatus(Device Device) (ClkMonStatus, Return)
	DeviceSetApplicationsClocks(Device Device, MemClockMHz uint32, GraphicsClockMHz uint32) Return
	DeviceSetPowerManagementLimit(Device Device, Limit uint32) Return
	DeviceSetGpuOperationMode(Device Device, Mode GpuOperationMode) Return
	DeviceSetAPIRestriction(Device Device, ApiType RestrictedAPI, IsRestricted EnableState) Return
	DeviceSetAccountingMode(Device Device, Mode EnableState) Return
	DeviceClearAccountingPids(Device Device) Return
	DeviceGetNvLinkState(Device Device, Link int) (EnableState, Return)
	DeviceGetNvLinkVersion(Device Device, Link int) (uint32, Return)
	DeviceGetNvLinkCapability(Device Device, Link int, Capability NvLinkCapability) (uint32, Return)
	DeviceGetNvLinkRemotePciInfo(Device Device, Link int) (PciInfo, Return)
	DeviceGetNvLinkErrorCounter(Device Device, Link int, Counter NvLinkErrorCounter) (uint64, Return)
	DeviceResetNvLinkErrorCounters(Device Device, Link int) Return
	DeviceSetNvLinkUtilizationControl(Device Device, Link int, Counter int, Control *NvLinkUtilizationControl, Reset bool) Return
	DeviceGetNvLinkUtilizationControl(Device Device, Link int, Counter int) (NvLinkUtilizationControl, Return)
	DeviceGetNvLinkUtilizationCounter(Device Device, Link int, Counter int) (uint64, uint64, Return)
	DeviceFreezeNvLinkUtilizationCounter(Device Device, Link int, Counter int, Freeze EnableState) Return
	DeviceResetNvLinkUtilizationCounter(Device Device, Link int, Counter int) Return
	DeviceGetNvLinkRemoteDeviceType(Device Device, Link int) (IntNvLinkDeviceType, Return)
	DeviceRegisterEvents(Device Device, EventTypes uint64, Set EventSet) Return
	DeviceGetSupportedEventTypes(Device Device) (uint64, Return)
	DeviceModifyDrainState(PciInfo *PciInfo, NewState EnableState) Return
	DeviceQueryDrainState(PciInfo *PciInfo) (EnableState, Return)
	DeviceRemoveGpu(PciInfo *PciInfo) Return
	DeviceRemoveGpu_v2(PciInfo *PciInfo, GpuState DetachGpuState, LinkState PcieLinkState) Return
	DeviceDiscoverGpus() (PciInfo, Return)
	DeviceGetFieldValues(Device Device, Values []FieldValue) Return
	DeviceGetVirtualizationMode(Device Device) (GpuVirtualizationMode, Return)
	DeviceGetHostVgpuMode(Device Device) (HostVgpuMode, Return)
	DeviceSetVirtualizationMode(Device Device, VirtualMode GpuVirtualizationMode) Return
	DeviceGetGridLicensableFeatures(Device Device) (GridLicensableFeatures, Return)
	DeviceGetProcessUtilization(Device Device, LastSeenTimeStamp uint64) ([]ProcessUtilizationSample, Return)
	DeviceGetSupportedVgpus(Device Device) ([]VgpuTypeId, Return)
	DeviceGetCreatableVgpus(Device Device) ([]VgpuTypeId, Return)
	DeviceGetActiveVgpus(Device Device) ([]VgpuInstance, Return)
	DeviceGetVgpuMetadata(Device Device) (VgpuPgpuMetadata, Return)
	DeviceGetPgpuMetadataString(Device Device) (string, Return)
	DeviceGetVgpuUtilization(Device Device, LastSeenTimeStamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return)
	DeviceGetAttributes(Device Device) (DeviceAttributes, Return)
	DeviceGetRemappedRows(Device Device) (int, int, bool, bool, Return)
	DeviceGetRowRemapperHistogram(Device Device) (RowRemapperHistogramValues, Return)
	DeviceGetArchitecture(Device Device) (DeviceArchitecture, Return)
	DeviceGetVgpuProcessUtilization(Device Device, LastSeenTimeStamp uint64) ([]VgpuProcessUtilizationSample, Return)
	GetExcludedDeviceCount() (int, Return)
	GetExcludedDeviceInfoByIndex(Index int) (ExcludedDeviceInfo, Return)
	DeviceSetMigMode(Device Device, Mode int) (Return, Return)
	DeviceGetMigMode(Device Device) (int, int, Return)
	DeviceGetGpuInstanceProfileInfo(Device Device, Profile int) (GpuInstanceProfileInfo, Return)
	DeviceGetGpuInstanceProfileInfoV(Device Device, Profile int) GpuInstanceProfileInfoV
	DeviceGetGpuInstancePossiblePlacements(Device Device, Info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return)
	DeviceGetGpuInstanceRemainingCapacity(Device Device, Info *GpuInstanceProfileInfo) (int, Return)
	DeviceCreateGpuInstance(Device Device, Info *GpuInstanceProfileInfo) (GpuInstance, Return)
	DeviceCreateGpuInstanceWithPlacement(Device Device, Info *GpuInstanceProfileInfo, Placement *GpuInstancePlacement) (GpuInstance, Return)
	GpuInstanceDestroy(GpuInstance GpuInstance) Return
	DeviceGetGpuInstances(Device Device, Info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
	DeviceGetGpuInstanceById(Device Device, Id int) (GpuInstance, Return)
	GpuInstanceGetInfo(GpuInstance GpuInstance) (GpuInstanceInfo, Return)
	GpuInstanceGetComputeInstanceProfileInfo(GpuInstance GpuInstance, Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return)
	GpuInstanceGetComputeInstanceProfileInfoV(GpuInstance GpuInstance, Profile int, EngProfile int) ComputeInstanceProfileInfoV
	GpuInstanceGetComputeInstanceRemainingCapacity(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) (int, Return)
	GpuInstanceCreateComputeInstance(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	ComputeInstanceDestroy(ComputeInstance ComputeInstance) Return
	GpuInstanceGetComputeInstances(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	GpuInstanceGetComputeInstanceById(GpuInstance GpuInstance, Id int) (ComputeInstance, Return)
	ComputeInstanceGetInfo(ComputeInstance ComputeInstance) (ComputeInstanceInfo, Return)
	DeviceIsMigDeviceHandle(Device Device) (bool, Return)
	DeviceGetGpuInstanceId(Device Device) (int, Return)
	DeviceGetComputeInstanceId(Device Device) (int, Return)
	DeviceGetMaxMigDeviceCount(Device Device) (int, Return)
	DeviceGetMigDeviceHandleByIndex(device Device, Index int) (Device, Return)
	DeviceGetDeviceHandleFromMigDeviceHandle(MigDevice Device) (Device, Return)
	DeviceGetBusType(Device Device) (BusType, Return)
	DeviceSetDefaultFanSpeed_v2(Device Device, Fan int) Return
	DeviceGetMinMaxFanSpeed(Device Device) (int, int, Return)
	DeviceGetThermalSettings(Device Device, SensorIndex uint32) (GpuThermalSettings, Return)
	DeviceGetDefaultEccMode(Device Device) (EnableState, Return)
	DeviceGetPcieSpeed(Device Device) (int, Return)
	DeviceGetGspFirmwareVersion(Device Device) (string, Return)
	DeviceGetGspFirmwareMode(Device Device) (bool, bool, Return)
	DeviceGetDynamicPstatesInfo(Device Device) (GpuDynamicPstatesInfo, Return)
	DeviceSetFanSpeed_v2(Device Device, Fan int, Speed int) Return
	DeviceGetGpcClkVfOffset(Device Device) (int, Return)
	DeviceSetGpcClkVfOffset(Device Device, Offset int) Return
	DeviceGetMinMaxClockOfPState(Device Device, _type ClockType, Pstate Pstates) (uint32, uint32, Return)
	DeviceGetSupportedPerformanceStates(Device Device) ([]Pstates, Return)
	DeviceGetTargetFanSpeed(Device Device, Fan int) (int, Return)
	DeviceGetMemClkVfOffset(Device Device) (int, Return)
	DeviceSetMemClkVfOffset(Device Device, Offset int) Return
	DeviceGetGpcClkMinMaxVfOffset(Device Device) (int, int, Return)
	DeviceGetMemClkMinMaxVfOffset(Device Device) (int, int, Return)
	DeviceGetGpuMaxPcieLinkGeneration(Device Device) (int, Return)
	DeviceGetFanControlPolicy_v2(Device Device, Fan int) (FanControlPolicy, Return)
	DeviceSetFanControlPolicy(Device Device, Fan int, Policy FanControlPolicy) Return
	DeviceClearFieldValues(Device Device, Values []FieldValue) Return
	DeviceGetVgpuCapabilities(Device Device, Capability DeviceVgpuCapability) (bool, Return)
	DeviceGetVgpuSchedulerLog(Device Device) (VgpuSchedulerLog, Return)
	DeviceGetVgpuSchedulerState(Device Device) (VgpuSchedulerGetState, Return)
	DeviceSetVgpuSchedulerState(Device Device, PSchedulerState *VgpuSchedulerSetState) Return
	DeviceGetVgpuSchedulerCapabilities(Device Device) (VgpuSchedulerCapabilities, Return)
	GpuInstanceGetComputeInstancePossiblePlacements(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return)
	GpuInstanceCreateComputeInstanceWithPlacement(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo, Placement *ComputeInstancePlacement) (ComputeInstance, Return)
	DeviceGetGpuFabricInfo(Device Device) (GpuFabricInfo, Return)
	DeviceCcuGetStreamState(Device Device) (int, Return)
	DeviceCcuSetStreamState(Device Device, State int) Return
	DeviceSetNvLinkDeviceLowPowerThreshold(Device Device, Info *NvLinkPowerThres) Return
	// Event API
	EventSetCreate() (EventSet, Return)
	EventSetWait(Set EventSet, Timeoutms uint32) (EventData, Return)
	EventSetFree(Set EventSet) Return
	// GPM API
	GpmMetricsGetV(MetricsGet *GpmMetricsGetType) GpmMetricsGetVType
	GpmMetricsGet(MetricsGet *GpmMetricsGetType) Return
	GpmSampleFree(GpmSample GpmSample) Return
	GpmSampleAlloc(GpmSample *GpmSample) Return
	GpmSampleGet(Device Device, GpmSample GpmSample) Return
	GpmQueryDeviceSupportV(Device Device) GpmSupportV
	GpmQueryDeviceSupport(Device Device) (GpmSupport, Return)
	GpmMigSampleGet(Device Device, GpuInstanceId int, GpmSample GpmSample) Return
	// Unit API
	UnitGetCount() (int, Return)
	UnitGetHandleByIndex(Index int) (Unit, Return)
	UnitGetUnitInfo(Unit Unit) (UnitInfo, Return)
	UnitGetLedState(Unit Unit) (LedState, Return)
	UnitGetPsuInfo(Unit Unit) (PSUInfo, Return)
	UnitGetTemperature(Unit Unit, Type int) (uint32, Return)
	UnitGetFanSpeedInfo(Unit Unit) (UnitFanSpeeds, Return)
	UnitGetDevices(Unit Unit) ([]Device, Return)
	UnitSetLedState(Unit Unit, Color LedColor) Return
	// vGPU API
	VgpuTypeGetClass(VgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetName(VgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetGpuInstanceProfileId(VgpuTypeId VgpuTypeId) (uint32, Return)
	VgpuTypeGetDeviceID(VgpuTypeId VgpuTypeId) (uint64, uint64, Return)
	VgpuTypeGetFramebufferSize(VgpuTypeId VgpuTypeId) (uint64, Return)
	VgpuTypeGetNumDisplayHeads(VgpuTypeId VgpuTypeId) (int, Return)
	VgpuTypeGetResolution(VgpuTypeId VgpuTypeId, DisplayIndex int) (uint32, uint32, Return)
	VgpuTypeGetLicense(VgpuTypeId VgpuTypeId) (string, Return)
	VgpuTypeGetFrameRateLimit(VgpuTypeId VgpuTypeId) (uint32, Return)
	VgpuTypeGetMaxInstances(Device Device, VgpuTypeId VgpuTypeId) (int, Return)
	VgpuTypeGetMaxInstancesPerVm(VgpuTypeId VgpuTypeId) (int, Return)
	VgpuInstanceGetVmID(VgpuInstance VgpuInstance) (string, VgpuVmIdType, Return)
	VgpuInstanceGetUUID(VgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetVmDriverVersion(VgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetFbUsage(VgpuInstance VgpuInstance) (uint64, Return)
	VgpuInstanceGetLicenseInfo(VgpuInstance VgpuInstance) (VgpuLicenseInfo, Return)
	VgpuInstanceGetLicenseStatus(VgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceGetType(VgpuInstance VgpuInstance) (VgpuTypeId, Return)
	VgpuInstanceGetFrameRateLimit(VgpuInstance VgpuInstance) (uint32, Return)
	VgpuInstanceGetEccMode(VgpuInstance VgpuInstance) (EnableState, Return)
	VgpuInstanceGetEncoderCapacity(VgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceSetEncoderCapacity(VgpuInstance VgpuInstance, EncoderCapacity int) Return
	VgpuInstanceGetEncoderStats(VgpuInstance VgpuInstance) (int, uint32, uint32, Return)
	VgpuInstanceGetEncoderSessions(VgpuInstance VgpuInstance) (int, EncoderSessionInfo, Return)
	VgpuInstanceGetFBCStats(VgpuInstance VgpuInstance) (FBCStats, Return)
	VgpuInstanceGetFBCSessions(VgpuInstance VgpuInstance) (int, FBCSessionInfo, Return)
	VgpuInstanceGetGpuInstanceId(VgpuInstance VgpuInstance) (int, Return)
	VgpuInstanceGetGpuPciId(VgpuInstance VgpuInstance) (string, Return)
	VgpuInstanceGetMetadata(VgpuInstance VgpuInstance) (VgpuMetadata, Return)
	VgpuInstanceGetAccountingMode(VgpuInstance VgpuInstance) (EnableState, Return)
	VgpuInstanceGetAccountingPids(VgpuInstance VgpuInstance) ([]int, Return)
	VgpuInstanceGetAccountingStats(VgpuInstance VgpuInstance, Pid int) (AccountingStats, Return)
	GetVgpuCompatibility(nvmlVgpuMetadata *nvmlVgpuMetadata, PgpuMetadata *nvmlVgpuPgpuMetadata) (VgpuPgpuCompatibility, Return)
	GetVgpuVersion() (VgpuVersion, VgpuVersion, Return)
	SetVgpuVersion(VgpuVersion *VgpuVersion) Return
	VgpuInstanceClearAccountingPids(VgpuInstance VgpuInstance) Return
	VgpuInstanceGetMdevUUID(VgpuInstance VgpuInstance) (string, Return)
	VgpuTypeGetCapabilities(VgpuTypeId VgpuTypeId, Capability VgpuCapability) (bool, Return)
	GetVgpuDriverCapabilities(Capability VgpuDriverCapability) (bool, Return)
}

type Device interface {
	GetName() (string, Return)
	GetBrand() (BrandType, Return)
	GetIndex() (int, Return)
	GetSerial() (string, Return)
	GetCpuAffinity(NumCPUs int) ([]uint, Return)
	SetCpuAffinity() Return
	ClearCpuAffinity() Return
	GetMemoryAffinity(NumNodes int, Scope AffinityScope) ([]uint, Return)
	GetCpuAffinityWithinScope(NumCPUs int, Scope AffinityScope) ([]uint, Return)
	GetTopologyCommonAncestor(Device2 Device) (GpuTopologyLevel, Return)
	GetTopologyNearestGpus(Level GpuTopologyLevel) ([]Device, Return)
	GetP2PStatus(Device2 Device, P2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return)
	GetUUID() (string, Return)
	GetMinorNumber() (int, Return)
	GetBoardPartNumber() (string, Return)
	GetInforomVersion(Object InforomObject) (string, Return)
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
	GetPcieThroughput(Counter PcieUtilCounter) (uint32, Return)
	GetPcieReplayCounter() (int, Return)
	GetClockInfo(_type ClockType) (uint32, Return)
	GetMaxClockInfo(_type ClockType) (uint32, Return)
	GetApplicationsClock(ClockType ClockType) (uint32, Return)
	GetDefaultApplicationsClock(ClockType ClockType) (uint32, Return)
	ResetApplicationsClocks() Return
	GetClock(ClockType ClockType, ClockId ClockId) (uint32, Return)
	GetMaxCustomerBoostClock(ClockType ClockType) (uint32, Return)
	GetSupportedMemoryClocks() (int, uint32, Return)
	GetSupportedGraphicsClocks(MemoryClockMHz int) (int, uint32, Return)
	GetAutoBoostedClocksEnabled() (EnableState, EnableState, Return)
	SetAutoBoostedClocksEnabled(Enabled EnableState) Return
	SetDefaultAutoBoostedClocksEnabled(Enabled EnableState, Flags uint32) Return
	GetFanSpeed() (uint32, Return)
	GetFanSpeed_v2(Fan int) (uint32, Return)
	GetNumFans() (int, Return)
	GetTemperature(SensorType TemperatureSensors) (uint32, Return)
	GetTemperatureThreshold(ThresholdType TemperatureThresholds) (uint32, Return)
	SetTemperatureThreshold(ThresholdType TemperatureThresholds, Temp int) Return
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
	GetTotalEccErrors(ErrorType MemoryErrorType, CounterType EccCounterType) (uint64, Return)
	GetDetailedEccErrors(ErrorType MemoryErrorType, CounterType EccCounterType) (EccErrorCounts, Return)
	GetMemoryErrorCounter(ErrorType MemoryErrorType, CounterType EccCounterType, LocationType MemoryLocation) (uint64, Return)
	GetUtilizationRates() (Utilization, Return)
	GetEncoderUtilization() (uint32, uint32, Return)
	GetEncoderCapacity(EncoderQueryType EncoderType) (int, Return)
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
	OnSameBoard(Device2 Device) (int, Return)
	GetAPIRestriction(ApiType RestrictedAPI) (EnableState, Return)
	GetSamples(_type SamplingType, LastSeenTimeStamp uint64) (ValueType, []Sample, Return)
	GetBAR1MemoryInfo() (BAR1Memory, Return)
	GetViolationStatus(PerfPolicyType PerfPolicyType) (ViolationTime, Return)
	GetIrqNum() (int, Return)
	GetNumGpuCores() (int, Return)
	GetPowerSource() (PowerSource, Return)
	GetMemoryBusWidth() (uint32, Return)
	GetPcieLinkMaxSpeed() (uint32, Return)
	GetAdaptiveClockInfoStatus() (uint32, Return)
	GetAccountingMode() (EnableState, Return)
	GetAccountingStats(Pid uint32) (AccountingStats, Return)
	GetAccountingPids() ([]int, Return)
	GetAccountingBufferSize() (int, Return)
	GetRetiredPages(Cause PageRetirementCause) ([]uint64, Return)
	GetRetiredPages_v2(Cause PageRetirementCause) ([]uint64, []uint64, Return)
	GetRetiredPagesPendingStatus() (EnableState, Return)
	SetPersistenceMode(Mode EnableState) Return
	SetComputeMode(Mode ComputeMode) Return
	SetEccMode(Ecc EnableState) Return
	ClearEccErrorCounts(CounterType EccCounterType) Return
	SetDriverModel(DriverModel DriverModel, Flags uint32) Return
	SetGpuLockedClocks(MinGpuClockMHz uint32, MaxGpuClockMHz uint32) Return
	ResetGpuLockedClocks() Return
	SetMemoryLockedClocks(NinMemClockMHz uint32, MaxMemClockMHz uint32) Return
	ResetMemoryLockedClocks() Return
	GetClkMonStatus() (ClkMonStatus, Return)
	SetApplicationsClocks(MemClockMHz uint32, GraphicsClockMHz uint32) Return
	SetPowerManagementLimit(Limit uint32) Return
	SetGpuOperationMode(Mode GpuOperationMode) Return
	SetAPIRestriction(ApiType RestrictedAPI, IsRestricted EnableState) Return
	SetAccountingMode(Mode EnableState) Return
	ClearAccountingPids() Return
	GetNvLinkState(Link int) (EnableState, Return)
	GetNvLinkVersion(Link int) (uint32, Return)
	GetNvLinkCapability(Link int, Capability NvLinkCapability) (uint32, Return)
	GetNvLinkRemotePciInfo(Link int) (PciInfo, Return)
	GetNvLinkErrorCounter(Link int, Counter NvLinkErrorCounter) (uint64, Return)
	ResetNvLinkErrorCounters(Link int) Return
	SetNvLinkUtilizationControl(Link int, Counter int, Control *NvLinkUtilizationControl, Reset bool) Return
	GetNvLinkUtilizationControl(Link int, Counter int) (NvLinkUtilizationControl, Return)
	GetNvLinkUtilizationCounter(Link int, Counter int) (uint64, uint64, Return)
	FreezeNvLinkUtilizationCounter(Link int, Counter int, Freeze EnableState) Return
	ResetNvLinkUtilizationCounter(Link int, Counter int) Return
	GetNvLinkRemoteDeviceType(Link int) (IntNvLinkDeviceType, Return)
	RegisterEvents(EventTypes uint64, Set EventSet) Return
	GetSupportedEventTypes() (uint64, Return)
	GetFieldValues(Values []FieldValue) Return
	GetVirtualizationMode() (GpuVirtualizationMode, Return)
	GetHostVgpuMode() (HostVgpuMode, Return)
	SetVirtualizationMode(VirtualMode GpuVirtualizationMode) Return
	GetGridLicensableFeatures() (GridLicensableFeatures, Return)
	GetProcessUtilization(LastSeenTimeStamp uint64) ([]ProcessUtilizationSample, Return)
	GetSupportedVgpus() ([]VgpuTypeId, Return)
	GetCreatableVgpus() ([]VgpuTypeId, Return)
	GetActiveVgpus() ([]VgpuInstance, Return)
	GetVgpuMetadata() (VgpuPgpuMetadata, Return)
	GetPgpuMetadataString() (string, Return)
	GetVgpuUtilization(LastSeenTimeStamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return)
	GetAttributes() (DeviceAttributes, Return)
	GetRemappedRows() (int, int, bool, bool, Return)
	GetRowRemapperHistogram() (RowRemapperHistogramValues, Return)
	GetArchitecture() (DeviceArchitecture, Return)
	GetVgpuProcessUtilization(LastSeenTimeStamp uint64) ([]VgpuProcessUtilizationSample, Return)
	SetMigMode(Mode int) (Return, Return)
	GetMigMode() (int, int, Return)
	GetGpuInstanceProfileInfo(Profile int) (GpuInstanceProfileInfo, Return)
	GetGpuInstanceProfileInfoV(Profile int) GpuInstanceProfileInfoV
	GetGpuInstancePossiblePlacements(Info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return)
	GetGpuInstanceRemainingCapacity(Info *GpuInstanceProfileInfo) (int, Return)
	CreateGpuInstance(Info *GpuInstanceProfileInfo) (GpuInstance, Return)
	CreateGpuInstanceWithPlacement(Info *GpuInstanceProfileInfo, Placement *GpuInstancePlacement) (GpuInstance, Return)
	GetGpuInstances(Info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
	GetGpuInstanceById(Id int) (GpuInstance, Return)
	IsMigDeviceHandle() (bool, Return)
	GetGpuInstanceId() (int, Return)
	GetComputeInstanceId() (int, Return)
	GetMaxMigDeviceCount() (int, Return)
	GetMigDeviceHandleByIndex(Index int) (Device, Return)
	GetDeviceHandleFromMigDeviceHandle() (Device, Return)
	GetBusType() (BusType, Return)
	SetDefaultFanSpeed_v2(Fan int) Return
	GetMinMaxFanSpeed() (int, int, Return)
	GetThermalSettings(SensorIndex uint32) (GpuThermalSettings, Return)
	GetDefaultEccMode() (EnableState, Return)
	GetPcieSpeed() (int, Return)
	GetGspFirmwareVersion() (string, Return)
	GetGspFirmwareMode() (bool, bool, Return)
	GetDynamicPstatesInfo() (GpuDynamicPstatesInfo, Return)
	SetFanSpeed_v2(Fan int, Speed int) Return
	GetGpcClkVfOffset() (int, Return)
	SetGpcClkVfOffset(Offset int) Return
	GetMinMaxClockOfPState(_type ClockType, Pstate Pstates) (uint32, uint32, Return)
	GetSupportedPerformanceStates() ([]Pstates, Return)
	GetTargetFanSpeed(Fan int) (int, Return)
	GetMemClkVfOffset() (int, Return)
	SetMemClkVfOffset(Offset int) Return
	GetGpcClkMinMaxVfOffset() (int, int, Return)
	GetMemClkMinMaxVfOffset() (int, int, Return)
	GetGpuMaxPcieLinkGeneration() (int, Return)
	GetFanControlPolicy_v2(Fan int) (FanControlPolicy, Return)
	SetFanControlPolicy(Fan int, Policy FanControlPolicy) Return
	ClearFieldValues(Values []FieldValue) Return
	GetVgpuCapabilities(Capability DeviceVgpuCapability) (bool, Return)
	GetVgpuSchedulerLog() (VgpuSchedulerLog, Return)
	GetVgpuSchedulerState() (VgpuSchedulerGetState, Return)
	SetVgpuSchedulerState(PSchedulerState *VgpuSchedulerSetState) Return
	GetVgpuSchedulerCapabilities() (VgpuSchedulerCapabilities, Return)
	GetGpuFabricInfo() (GpuFabricInfo, Return)
	CcuGetStreamState() (int, Return)
	CcuSetStreamState(State int) Return
	SetNvLinkDeviceLowPowerThreshold(Info *NvLinkPowerThres) Return
	GpmSampleGet(GpmSample GpmSample) Return
	GpmQueryDeviceSupportV() GpmSupportV
	GpmQueryDeviceSupport() (GpmSupport, Return)
	GpmMigSampleGet(GpuInstanceId int, GpmSample GpmSample) Return
	VgpuTypeGetMaxInstances(VgpuTypeId VgpuTypeId) (int, Return)
}

type GpuInstance interface {
	Destroy() Return
	GetInfo() (GpuInstanceInfo, Return)
	GetComputeInstanceProfileInfo(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return)
	GetComputeInstanceProfileInfoV(Profile int, EngProfile int) ComputeInstanceProfileInfoV
	GetComputeInstanceRemainingCapacity(Info *ComputeInstanceProfileInfo) (int, Return)
	CreateComputeInstance(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	GetComputeInstances(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	GetComputeInstanceById(Id int) (ComputeInstance, Return)
	GetComputeInstancePossiblePlacements(Info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return)
	CreateComputeInstanceWithPlacement(Info *ComputeInstanceProfileInfo, Placement *ComputeInstancePlacement) (ComputeInstance, Return)
}

type EventSet interface {
	Wait(Timeoutms uint32) (EventData, Return)
	Free() Return
}

type Unit interface {
	GetUnitInfo() (UnitInfo, Return)
	GetLedState() (LedState, Return)
	GetPsuInfo() (PSUInfo, Return)
	GetTemperature(Type int) (uint32, Return)
	GetFanSpeedInfo() (UnitFanSpeeds, Return)
	GetDevices() ([]Device, Return)
	SetLedState(Color LedColor) Return
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
	SetEncoderCapacity(EncoderCapacity int) Return
	GetEncoderStats() (int, uint32, uint32, Return)
	GetEncoderSessions() (int, EncoderSessionInfo, Return)
	GetFBCStats() (FBCStats, Return)
	GetFBCSessions() (int, FBCSessionInfo, Return)
	GetGpuInstanceId() (int, Return)
	GetGpuPciId() (string, Return)
	GetMetadata() (VgpuMetadata, Return)
	GetAccountingMode() (EnableState, Return)
	GetAccountingPids() ([]int, Return)
	GetAccountingStats(Pid int) (AccountingStats, Return)
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
	GetResolution(DisplayIndex int) (uint32, uint32, Return)
	GetLicense() (string, Return)
	GetFrameRateLimit() (uint32, Return)
	GetMaxInstances(Device Device) (int, Return)
	GetMaxInstancesPerVm() (int, Return)
	GetCapabilities(Capability VgpuCapability) (bool, Return)
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
