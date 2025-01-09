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

// Generated Code; DO NOT EDIT.

package nvml

// The variables below represent package level methods from the library type.
var (
	ComputeInstanceDestroy                          = libnvml.ComputeInstanceDestroy
	ComputeInstanceGetInfo                          = libnvml.ComputeInstanceGetInfo
	DeviceClearAccountingPids                       = libnvml.DeviceClearAccountingPids
	DeviceClearCpuAffinity                          = libnvml.DeviceClearCpuAffinity
	DeviceClearEccErrorCounts                       = libnvml.DeviceClearEccErrorCounts
	DeviceClearFieldValues                          = libnvml.DeviceClearFieldValues
	DeviceCreateGpuInstance                         = libnvml.DeviceCreateGpuInstance
	DeviceCreateGpuInstanceWithPlacement            = libnvml.DeviceCreateGpuInstanceWithPlacement
	DeviceDiscoverGpus                              = libnvml.DeviceDiscoverGpus
	DeviceFreezeNvLinkUtilizationCounter            = libnvml.DeviceFreezeNvLinkUtilizationCounter
	DeviceGetAPIRestriction                         = libnvml.DeviceGetAPIRestriction
	DeviceGetAccountingBufferSize                   = libnvml.DeviceGetAccountingBufferSize
	DeviceGetAccountingMode                         = libnvml.DeviceGetAccountingMode
	DeviceGetAccountingPids                         = libnvml.DeviceGetAccountingPids
	DeviceGetAccountingStats                        = libnvml.DeviceGetAccountingStats
	DeviceGetActiveVgpus                            = libnvml.DeviceGetActiveVgpus
	DeviceGetAdaptiveClockInfoStatus                = libnvml.DeviceGetAdaptiveClockInfoStatus
	DeviceGetApplicationsClock                      = libnvml.DeviceGetApplicationsClock
	DeviceGetArchitecture                           = libnvml.DeviceGetArchitecture
	DeviceGetAttributes                             = libnvml.DeviceGetAttributes
	DeviceGetAutoBoostedClocksEnabled               = libnvml.DeviceGetAutoBoostedClocksEnabled
	DeviceGetBAR1MemoryInfo                         = libnvml.DeviceGetBAR1MemoryInfo
	DeviceGetBoardId                                = libnvml.DeviceGetBoardId
	DeviceGetBoardPartNumber                        = libnvml.DeviceGetBoardPartNumber
	DeviceGetBrand                                  = libnvml.DeviceGetBrand
	DeviceGetBridgeChipInfo                         = libnvml.DeviceGetBridgeChipInfo
	DeviceGetBusType                                = libnvml.DeviceGetBusType
	DeviceGetC2cModeInfoV                           = libnvml.DeviceGetC2cModeInfoV
	DeviceGetClkMonStatus                           = libnvml.DeviceGetClkMonStatus
	DeviceGetClock                                  = libnvml.DeviceGetClock
	DeviceGetClockInfo                              = libnvml.DeviceGetClockInfo
	DeviceGetComputeInstanceId                      = libnvml.DeviceGetComputeInstanceId
	DeviceGetComputeMode                            = libnvml.DeviceGetComputeMode
	DeviceGetComputeRunningProcesses                = libnvml.DeviceGetComputeRunningProcesses
	DeviceGetConfComputeGpuAttestationReport        = libnvml.DeviceGetConfComputeGpuAttestationReport
	DeviceGetConfComputeGpuCertificate              = libnvml.DeviceGetConfComputeGpuCertificate
	DeviceGetConfComputeMemSizeInfo                 = libnvml.DeviceGetConfComputeMemSizeInfo
	DeviceGetConfComputeProtectedMemoryUsage        = libnvml.DeviceGetConfComputeProtectedMemoryUsage
	DeviceGetCount                                  = libnvml.DeviceGetCount
	DeviceGetCpuAffinity                            = libnvml.DeviceGetCpuAffinity
	DeviceGetCpuAffinityWithinScope                 = libnvml.DeviceGetCpuAffinityWithinScope
	DeviceGetCreatableVgpus                         = libnvml.DeviceGetCreatableVgpus
	DeviceGetCudaComputeCapability                  = libnvml.DeviceGetCudaComputeCapability
	DeviceGetCurrPcieLinkGeneration                 = libnvml.DeviceGetCurrPcieLinkGeneration
	DeviceGetCurrPcieLinkWidth                      = libnvml.DeviceGetCurrPcieLinkWidth
	DeviceGetCurrentClocksEventReasons              = libnvml.DeviceGetCurrentClocksEventReasons
	DeviceGetCurrentClocksThrottleReasons           = libnvml.DeviceGetCurrentClocksThrottleReasons
	DeviceGetDecoderUtilization                     = libnvml.DeviceGetDecoderUtilization
	DeviceGetDefaultApplicationsClock               = libnvml.DeviceGetDefaultApplicationsClock
	DeviceGetDefaultEccMode                         = libnvml.DeviceGetDefaultEccMode
	DeviceGetDetailedEccErrors                      = libnvml.DeviceGetDetailedEccErrors
	DeviceGetDeviceHandleFromMigDeviceHandle        = libnvml.DeviceGetDeviceHandleFromMigDeviceHandle
	DeviceGetDisplayActive                          = libnvml.DeviceGetDisplayActive
	DeviceGetDisplayMode                            = libnvml.DeviceGetDisplayMode
	DeviceGetDriverModel                            = libnvml.DeviceGetDriverModel
	DeviceGetDynamicPstatesInfo                     = libnvml.DeviceGetDynamicPstatesInfo
	DeviceGetEccMode                                = libnvml.DeviceGetEccMode
	DeviceGetEncoderCapacity                        = libnvml.DeviceGetEncoderCapacity
	DeviceGetEncoderSessions                        = libnvml.DeviceGetEncoderSessions
	DeviceGetEncoderStats                           = libnvml.DeviceGetEncoderStats
	DeviceGetEncoderUtilization                     = libnvml.DeviceGetEncoderUtilization
	DeviceGetEnforcedPowerLimit                     = libnvml.DeviceGetEnforcedPowerLimit
	DeviceGetFBCSessions                            = libnvml.DeviceGetFBCSessions
	DeviceGetFBCStats                               = libnvml.DeviceGetFBCStats
	DeviceGetFanControlPolicy_v2                    = libnvml.DeviceGetFanControlPolicy_v2
	DeviceGetFanSpeed                               = libnvml.DeviceGetFanSpeed
	DeviceGetFanSpeed_v2                            = libnvml.DeviceGetFanSpeed_v2
	DeviceGetFieldValues                            = libnvml.DeviceGetFieldValues
	DeviceGetGpcClkMinMaxVfOffset                   = libnvml.DeviceGetGpcClkMinMaxVfOffset
	DeviceGetGpcClkVfOffset                         = libnvml.DeviceGetGpcClkVfOffset
	DeviceGetGpuFabricInfo                          = libnvml.DeviceGetGpuFabricInfo
	DeviceGetGpuFabricInfoV                         = libnvml.DeviceGetGpuFabricInfoV
	DeviceGetGpuInstanceById                        = libnvml.DeviceGetGpuInstanceById
	DeviceGetGpuInstanceId                          = libnvml.DeviceGetGpuInstanceId
	DeviceGetGpuInstancePossiblePlacements          = libnvml.DeviceGetGpuInstancePossiblePlacements
	DeviceGetGpuInstanceProfileInfo                 = libnvml.DeviceGetGpuInstanceProfileInfo
	DeviceGetGpuInstanceProfileInfoV                = libnvml.DeviceGetGpuInstanceProfileInfoV
	DeviceGetGpuInstanceRemainingCapacity           = libnvml.DeviceGetGpuInstanceRemainingCapacity
	DeviceGetGpuInstances                           = libnvml.DeviceGetGpuInstances
	DeviceGetGpuMaxPcieLinkGeneration               = libnvml.DeviceGetGpuMaxPcieLinkGeneration
	DeviceGetGpuOperationMode                       = libnvml.DeviceGetGpuOperationMode
	DeviceGetGraphicsRunningProcesses               = libnvml.DeviceGetGraphicsRunningProcesses
	DeviceGetGridLicensableFeatures                 = libnvml.DeviceGetGridLicensableFeatures
	DeviceGetGspFirmwareMode                        = libnvml.DeviceGetGspFirmwareMode
	DeviceGetGspFirmwareVersion                     = libnvml.DeviceGetGspFirmwareVersion
	DeviceGetHandleByIndex                          = libnvml.DeviceGetHandleByIndex
	DeviceGetHandleByPciBusId                       = libnvml.DeviceGetHandleByPciBusId
	DeviceGetHandleBySerial                         = libnvml.DeviceGetHandleBySerial
	DeviceGetHandleByUUID                           = libnvml.DeviceGetHandleByUUID
	DeviceGetHostVgpuMode                           = libnvml.DeviceGetHostVgpuMode
	DeviceGetIndex                                  = libnvml.DeviceGetIndex
	DeviceGetInforomConfigurationChecksum           = libnvml.DeviceGetInforomConfigurationChecksum
	DeviceGetInforomImageVersion                    = libnvml.DeviceGetInforomImageVersion
	DeviceGetInforomVersion                         = libnvml.DeviceGetInforomVersion
	DeviceGetIrqNum                                 = libnvml.DeviceGetIrqNum
	DeviceGetJpgUtilization                         = libnvml.DeviceGetJpgUtilization
	DeviceGetLastBBXFlushTime                       = libnvml.DeviceGetLastBBXFlushTime
	DeviceGetMPSComputeRunningProcesses             = libnvml.DeviceGetMPSComputeRunningProcesses
	DeviceGetMaxClockInfo                           = libnvml.DeviceGetMaxClockInfo
	DeviceGetMaxCustomerBoostClock                  = libnvml.DeviceGetMaxCustomerBoostClock
	DeviceGetMaxMigDeviceCount                      = libnvml.DeviceGetMaxMigDeviceCount
	DeviceGetMaxPcieLinkGeneration                  = libnvml.DeviceGetMaxPcieLinkGeneration
	DeviceGetMaxPcieLinkWidth                       = libnvml.DeviceGetMaxPcieLinkWidth
	DeviceGetMemClkMinMaxVfOffset                   = libnvml.DeviceGetMemClkMinMaxVfOffset
	DeviceGetMemClkVfOffset                         = libnvml.DeviceGetMemClkVfOffset
	DeviceGetMemoryAffinity                         = libnvml.DeviceGetMemoryAffinity
	DeviceGetMemoryBusWidth                         = libnvml.DeviceGetMemoryBusWidth
	DeviceGetMemoryErrorCounter                     = libnvml.DeviceGetMemoryErrorCounter
	DeviceGetMemoryInfo                             = libnvml.DeviceGetMemoryInfo
	DeviceGetMemoryInfo_v2                          = libnvml.DeviceGetMemoryInfo_v2
	DeviceGetMigDeviceHandleByIndex                 = libnvml.DeviceGetMigDeviceHandleByIndex
	DeviceGetMigMode                                = libnvml.DeviceGetMigMode
	DeviceGetMinMaxClockOfPState                    = libnvml.DeviceGetMinMaxClockOfPState
	DeviceGetMinMaxFanSpeed                         = libnvml.DeviceGetMinMaxFanSpeed
	DeviceGetMinorNumber                            = libnvml.DeviceGetMinorNumber
	DeviceGetModuleId                               = libnvml.DeviceGetModuleId
	DeviceGetMultiGpuBoard                          = libnvml.DeviceGetMultiGpuBoard
	DeviceGetName                                   = libnvml.DeviceGetName
	DeviceGetNumFans                                = libnvml.DeviceGetNumFans
	DeviceGetNumGpuCores                            = libnvml.DeviceGetNumGpuCores
	DeviceGetNumaNodeId                             = libnvml.DeviceGetNumaNodeId
	DeviceGetNvLinkCapability                       = libnvml.DeviceGetNvLinkCapability
	DeviceGetNvLinkErrorCounter                     = libnvml.DeviceGetNvLinkErrorCounter
	DeviceGetNvLinkRemoteDeviceType                 = libnvml.DeviceGetNvLinkRemoteDeviceType
	DeviceGetNvLinkRemotePciInfo                    = libnvml.DeviceGetNvLinkRemotePciInfo
	DeviceGetNvLinkState                            = libnvml.DeviceGetNvLinkState
	DeviceGetNvLinkUtilizationControl               = libnvml.DeviceGetNvLinkUtilizationControl
	DeviceGetNvLinkUtilizationCounter               = libnvml.DeviceGetNvLinkUtilizationCounter
	DeviceGetNvLinkVersion                          = libnvml.DeviceGetNvLinkVersion
	DeviceGetOfaUtilization                         = libnvml.DeviceGetOfaUtilization
	DeviceGetP2PStatus                              = libnvml.DeviceGetP2PStatus
	DeviceGetPciInfo                                = libnvml.DeviceGetPciInfo
	DeviceGetPciInfoExt                             = libnvml.DeviceGetPciInfoExt
	DeviceGetPcieLinkMaxSpeed                       = libnvml.DeviceGetPcieLinkMaxSpeed
	DeviceGetPcieReplayCounter                      = libnvml.DeviceGetPcieReplayCounter
	DeviceGetPcieSpeed                              = libnvml.DeviceGetPcieSpeed
	DeviceGetPcieThroughput                         = libnvml.DeviceGetPcieThroughput
	DeviceGetPerformanceState                       = libnvml.DeviceGetPerformanceState
	DeviceGetPersistenceMode                        = libnvml.DeviceGetPersistenceMode
	DeviceGetPgpuMetadataString                     = libnvml.DeviceGetPgpuMetadataString
	DeviceGetPowerManagementDefaultLimit            = libnvml.DeviceGetPowerManagementDefaultLimit
	DeviceGetPowerManagementLimit                   = libnvml.DeviceGetPowerManagementLimit
	DeviceGetPowerManagementLimitConstraints        = libnvml.DeviceGetPowerManagementLimitConstraints
	DeviceGetPowerManagementMode                    = libnvml.DeviceGetPowerManagementMode
	DeviceGetPowerSource                            = libnvml.DeviceGetPowerSource
	DeviceGetPowerState                             = libnvml.DeviceGetPowerState
	DeviceGetPowerUsage                             = libnvml.DeviceGetPowerUsage
	DeviceGetProcessUtilization                     = libnvml.DeviceGetProcessUtilization
	DeviceGetProcessesUtilizationInfo               = libnvml.DeviceGetProcessesUtilizationInfo
	DeviceGetRemappedRows                           = libnvml.DeviceGetRemappedRows
	DeviceGetRetiredPages                           = libnvml.DeviceGetRetiredPages
	DeviceGetRetiredPagesPendingStatus              = libnvml.DeviceGetRetiredPagesPendingStatus
	DeviceGetRetiredPages_v2                        = libnvml.DeviceGetRetiredPages_v2
	DeviceGetRowRemapperHistogram                   = libnvml.DeviceGetRowRemapperHistogram
	DeviceGetRunningProcessDetailList               = libnvml.DeviceGetRunningProcessDetailList
	DeviceGetSamples                                = libnvml.DeviceGetSamples
	DeviceGetSerial                                 = libnvml.DeviceGetSerial
	DeviceGetSramEccErrorStatus                     = libnvml.DeviceGetSramEccErrorStatus
	DeviceGetSupportedClocksEventReasons            = libnvml.DeviceGetSupportedClocksEventReasons
	DeviceGetSupportedClocksThrottleReasons         = libnvml.DeviceGetSupportedClocksThrottleReasons
	DeviceGetSupportedEventTypes                    = libnvml.DeviceGetSupportedEventTypes
	DeviceGetSupportedGraphicsClocks                = libnvml.DeviceGetSupportedGraphicsClocks
	DeviceGetSupportedMemoryClocks                  = libnvml.DeviceGetSupportedMemoryClocks
	DeviceGetSupportedPerformanceStates             = libnvml.DeviceGetSupportedPerformanceStates
	DeviceGetSupportedVgpus                         = libnvml.DeviceGetSupportedVgpus
	DeviceGetTargetFanSpeed                         = libnvml.DeviceGetTargetFanSpeed
	DeviceGetTemperature                            = libnvml.DeviceGetTemperature
	DeviceGetTemperatureThreshold                   = libnvml.DeviceGetTemperatureThreshold
	DeviceGetThermalSettings                        = libnvml.DeviceGetThermalSettings
	DeviceGetTopologyCommonAncestor                 = libnvml.DeviceGetTopologyCommonAncestor
	DeviceGetTopologyNearestGpus                    = libnvml.DeviceGetTopologyNearestGpus
	DeviceGetTotalEccErrors                         = libnvml.DeviceGetTotalEccErrors
	DeviceGetTotalEnergyConsumption                 = libnvml.DeviceGetTotalEnergyConsumption
	DeviceGetUUID                                   = libnvml.DeviceGetUUID
	DeviceGetUtilizationRates                       = libnvml.DeviceGetUtilizationRates
	DeviceGetVbiosVersion                           = libnvml.DeviceGetVbiosVersion
	DeviceGetVgpuCapabilities                       = libnvml.DeviceGetVgpuCapabilities
	DeviceGetVgpuHeterogeneousMode                  = libnvml.DeviceGetVgpuHeterogeneousMode
	DeviceGetVgpuInstancesUtilizationInfo           = libnvml.DeviceGetVgpuInstancesUtilizationInfo
	DeviceGetVgpuMetadata                           = libnvml.DeviceGetVgpuMetadata
	DeviceGetVgpuProcessUtilization                 = libnvml.DeviceGetVgpuProcessUtilization
	DeviceGetVgpuProcessesUtilizationInfo           = libnvml.DeviceGetVgpuProcessesUtilizationInfo
	DeviceGetVgpuSchedulerCapabilities              = libnvml.DeviceGetVgpuSchedulerCapabilities
	DeviceGetVgpuSchedulerLog                       = libnvml.DeviceGetVgpuSchedulerLog
	DeviceGetVgpuSchedulerState                     = libnvml.DeviceGetVgpuSchedulerState
	DeviceGetVgpuTypeCreatablePlacements            = libnvml.DeviceGetVgpuTypeCreatablePlacements
	DeviceGetVgpuTypeSupportedPlacements            = libnvml.DeviceGetVgpuTypeSupportedPlacements
	DeviceGetVgpuUtilization                        = libnvml.DeviceGetVgpuUtilization
	DeviceGetViolationStatus                        = libnvml.DeviceGetViolationStatus
	DeviceGetVirtualizationMode                     = libnvml.DeviceGetVirtualizationMode
	DeviceIsMigDeviceHandle                         = libnvml.DeviceIsMigDeviceHandle
	DeviceModifyDrainState                          = libnvml.DeviceModifyDrainState
	DeviceOnSameBoard                               = libnvml.DeviceOnSameBoard
	DeviceQueryDrainState                           = libnvml.DeviceQueryDrainState
	DeviceRegisterEvents                            = libnvml.DeviceRegisterEvents
	DeviceRemoveGpu                                 = libnvml.DeviceRemoveGpu
	DeviceRemoveGpu_v2                              = libnvml.DeviceRemoveGpu_v2
	DeviceResetApplicationsClocks                   = libnvml.DeviceResetApplicationsClocks
	DeviceResetGpuLockedClocks                      = libnvml.DeviceResetGpuLockedClocks
	DeviceResetMemoryLockedClocks                   = libnvml.DeviceResetMemoryLockedClocks
	DeviceResetNvLinkErrorCounters                  = libnvml.DeviceResetNvLinkErrorCounters
	DeviceResetNvLinkUtilizationCounter             = libnvml.DeviceResetNvLinkUtilizationCounter
	DeviceSetAPIRestriction                         = libnvml.DeviceSetAPIRestriction
	DeviceSetAccountingMode                         = libnvml.DeviceSetAccountingMode
	DeviceSetApplicationsClocks                     = libnvml.DeviceSetApplicationsClocks
	DeviceSetAutoBoostedClocksEnabled               = libnvml.DeviceSetAutoBoostedClocksEnabled
	DeviceSetComputeMode                            = libnvml.DeviceSetComputeMode
	DeviceSetConfComputeUnprotectedMemSize          = libnvml.DeviceSetConfComputeUnprotectedMemSize
	DeviceSetCpuAffinity                            = libnvml.DeviceSetCpuAffinity
	DeviceSetDefaultAutoBoostedClocksEnabled        = libnvml.DeviceSetDefaultAutoBoostedClocksEnabled
	DeviceSetDefaultFanSpeed_v2                     = libnvml.DeviceSetDefaultFanSpeed_v2
	DeviceSetDriverModel                            = libnvml.DeviceSetDriverModel
	DeviceSetEccMode                                = libnvml.DeviceSetEccMode
	DeviceSetFanControlPolicy                       = libnvml.DeviceSetFanControlPolicy
	DeviceSetFanSpeed_v2                            = libnvml.DeviceSetFanSpeed_v2
	DeviceSetGpcClkVfOffset                         = libnvml.DeviceSetGpcClkVfOffset
	DeviceSetGpuLockedClocks                        = libnvml.DeviceSetGpuLockedClocks
	DeviceSetGpuOperationMode                       = libnvml.DeviceSetGpuOperationMode
	DeviceSetMemClkVfOffset                         = libnvml.DeviceSetMemClkVfOffset
	DeviceSetMemoryLockedClocks                     = libnvml.DeviceSetMemoryLockedClocks
	DeviceSetMigMode                                = libnvml.DeviceSetMigMode
	DeviceSetNvLinkDeviceLowPowerThreshold          = libnvml.DeviceSetNvLinkDeviceLowPowerThreshold
	DeviceSetNvLinkUtilizationControl               = libnvml.DeviceSetNvLinkUtilizationControl
	DeviceSetPersistenceMode                        = libnvml.DeviceSetPersistenceMode
	DeviceSetPowerManagementLimit                   = libnvml.DeviceSetPowerManagementLimit
	DeviceSetPowerManagementLimit_v2                = libnvml.DeviceSetPowerManagementLimit_v2
	DeviceSetTemperatureThreshold                   = libnvml.DeviceSetTemperatureThreshold
	DeviceSetVgpuCapabilities                       = libnvml.DeviceSetVgpuCapabilities
	DeviceSetVgpuHeterogeneousMode                  = libnvml.DeviceSetVgpuHeterogeneousMode
	DeviceSetVgpuSchedulerState                     = libnvml.DeviceSetVgpuSchedulerState
	DeviceSetVirtualizationMode                     = libnvml.DeviceSetVirtualizationMode
	DeviceValidateInforom                           = libnvml.DeviceValidateInforom
	ErrorString                                     = libnvml.ErrorString
	EventSetCreate                                  = libnvml.EventSetCreate
	EventSetFree                                    = libnvml.EventSetFree
	EventSetWait                                    = libnvml.EventSetWait
	Extensions                                      = libnvml.Extensions
	GetExcludedDeviceCount                          = libnvml.GetExcludedDeviceCount
	GetExcludedDeviceInfoByIndex                    = libnvml.GetExcludedDeviceInfoByIndex
	GetVgpuCompatibility                            = libnvml.GetVgpuCompatibility
	GetVgpuDriverCapabilities                       = libnvml.GetVgpuDriverCapabilities
	GetVgpuVersion                                  = libnvml.GetVgpuVersion
	GpmMetricsGet                                   = libnvml.GpmMetricsGet
	GpmMetricsGetV                                  = libnvml.GpmMetricsGetV
	GpmMigSampleGet                                 = libnvml.GpmMigSampleGet
	GpmQueryDeviceSupport                           = libnvml.GpmQueryDeviceSupport
	GpmQueryDeviceSupportV                          = libnvml.GpmQueryDeviceSupportV
	GpmQueryIfStreamingEnabled                      = libnvml.GpmQueryIfStreamingEnabled
	GpmSampleAlloc                                  = libnvml.GpmSampleAlloc
	GpmSampleFree                                   = libnvml.GpmSampleFree
	GpmSampleGet                                    = libnvml.GpmSampleGet
	GpmSetStreamingEnabled                          = libnvml.GpmSetStreamingEnabled
	GpuInstanceCreateComputeInstance                = libnvml.GpuInstanceCreateComputeInstance
	GpuInstanceCreateComputeInstanceWithPlacement   = libnvml.GpuInstanceCreateComputeInstanceWithPlacement
	GpuInstanceDestroy                              = libnvml.GpuInstanceDestroy
	GpuInstanceGetComputeInstanceById               = libnvml.GpuInstanceGetComputeInstanceById
	GpuInstanceGetComputeInstancePossiblePlacements = libnvml.GpuInstanceGetComputeInstancePossiblePlacements
	GpuInstanceGetComputeInstanceProfileInfo        = libnvml.GpuInstanceGetComputeInstanceProfileInfo
	GpuInstanceGetComputeInstanceProfileInfoV       = libnvml.GpuInstanceGetComputeInstanceProfileInfoV
	GpuInstanceGetComputeInstanceRemainingCapacity  = libnvml.GpuInstanceGetComputeInstanceRemainingCapacity
	GpuInstanceGetComputeInstances                  = libnvml.GpuInstanceGetComputeInstances
	GpuInstanceGetInfo                              = libnvml.GpuInstanceGetInfo
	Init                                            = libnvml.Init
	InitWithFlags                                   = libnvml.InitWithFlags
	SetVgpuVersion                                  = libnvml.SetVgpuVersion
	Shutdown                                        = libnvml.Shutdown
	SystemGetConfComputeCapabilities                = libnvml.SystemGetConfComputeCapabilities
	SystemGetConfComputeKeyRotationThresholdInfo    = libnvml.SystemGetConfComputeKeyRotationThresholdInfo
	SystemGetConfComputeSettings                    = libnvml.SystemGetConfComputeSettings
	SystemGetCudaDriverVersion                      = libnvml.SystemGetCudaDriverVersion
	SystemGetCudaDriverVersion_v2                   = libnvml.SystemGetCudaDriverVersion_v2
	SystemGetDriverVersion                          = libnvml.SystemGetDriverVersion
	SystemGetHicVersion                             = libnvml.SystemGetHicVersion
	SystemGetNVMLVersion                            = libnvml.SystemGetNVMLVersion
	SystemGetProcessName                            = libnvml.SystemGetProcessName
	SystemGetTopologyGpuSet                         = libnvml.SystemGetTopologyGpuSet
	SystemSetConfComputeKeyRotationThresholdInfo    = libnvml.SystemSetConfComputeKeyRotationThresholdInfo
	UnitGetCount                                    = libnvml.UnitGetCount
	UnitGetDevices                                  = libnvml.UnitGetDevices
	UnitGetFanSpeedInfo                             = libnvml.UnitGetFanSpeedInfo
	UnitGetHandleByIndex                            = libnvml.UnitGetHandleByIndex
	UnitGetLedState                                 = libnvml.UnitGetLedState
	UnitGetPsuInfo                                  = libnvml.UnitGetPsuInfo
	UnitGetTemperature                              = libnvml.UnitGetTemperature
	UnitGetUnitInfo                                 = libnvml.UnitGetUnitInfo
	UnitSetLedState                                 = libnvml.UnitSetLedState
	VgpuInstanceClearAccountingPids                 = libnvml.VgpuInstanceClearAccountingPids
	VgpuInstanceGetAccountingMode                   = libnvml.VgpuInstanceGetAccountingMode
	VgpuInstanceGetAccountingPids                   = libnvml.VgpuInstanceGetAccountingPids
	VgpuInstanceGetAccountingStats                  = libnvml.VgpuInstanceGetAccountingStats
	VgpuInstanceGetEccMode                          = libnvml.VgpuInstanceGetEccMode
	VgpuInstanceGetEncoderCapacity                  = libnvml.VgpuInstanceGetEncoderCapacity
	VgpuInstanceGetEncoderSessions                  = libnvml.VgpuInstanceGetEncoderSessions
	VgpuInstanceGetEncoderStats                     = libnvml.VgpuInstanceGetEncoderStats
	VgpuInstanceGetFBCSessions                      = libnvml.VgpuInstanceGetFBCSessions
	VgpuInstanceGetFBCStats                         = libnvml.VgpuInstanceGetFBCStats
	VgpuInstanceGetFbUsage                          = libnvml.VgpuInstanceGetFbUsage
	VgpuInstanceGetFrameRateLimit                   = libnvml.VgpuInstanceGetFrameRateLimit
	VgpuInstanceGetGpuInstanceId                    = libnvml.VgpuInstanceGetGpuInstanceId
	VgpuInstanceGetGpuPciId                         = libnvml.VgpuInstanceGetGpuPciId
	VgpuInstanceGetLicenseInfo                      = libnvml.VgpuInstanceGetLicenseInfo
	VgpuInstanceGetLicenseStatus                    = libnvml.VgpuInstanceGetLicenseStatus
	VgpuInstanceGetMdevUUID                         = libnvml.VgpuInstanceGetMdevUUID
	VgpuInstanceGetMetadata                         = libnvml.VgpuInstanceGetMetadata
	VgpuInstanceGetType                             = libnvml.VgpuInstanceGetType
	VgpuInstanceGetUUID                             = libnvml.VgpuInstanceGetUUID
	VgpuInstanceGetVmDriverVersion                  = libnvml.VgpuInstanceGetVmDriverVersion
	VgpuInstanceGetVmID                             = libnvml.VgpuInstanceGetVmID
	VgpuInstanceSetEncoderCapacity                  = libnvml.VgpuInstanceSetEncoderCapacity
	VgpuTypeGetCapabilities                         = libnvml.VgpuTypeGetCapabilities
	VgpuTypeGetClass                                = libnvml.VgpuTypeGetClass
	VgpuTypeGetDeviceID                             = libnvml.VgpuTypeGetDeviceID
	VgpuTypeGetFrameRateLimit                       = libnvml.VgpuTypeGetFrameRateLimit
	VgpuTypeGetFramebufferSize                      = libnvml.VgpuTypeGetFramebufferSize
	VgpuTypeGetGpuInstanceProfileId                 = libnvml.VgpuTypeGetGpuInstanceProfileId
	VgpuTypeGetLicense                              = libnvml.VgpuTypeGetLicense
	VgpuTypeGetMaxInstances                         = libnvml.VgpuTypeGetMaxInstances
	VgpuTypeGetMaxInstancesPerVm                    = libnvml.VgpuTypeGetMaxInstancesPerVm
	VgpuTypeGetName                                 = libnvml.VgpuTypeGetName
	VgpuTypeGetNumDisplayHeads                      = libnvml.VgpuTypeGetNumDisplayHeads
	VgpuTypeGetResolution                           = libnvml.VgpuTypeGetResolution
)

// Interface represents the interface for the library type.
//
//go:generate moq -rm -out mock/interface.go -pkg mock . Interface:Interface
type Interface interface {
	ComputeInstanceDestroy(ComputeInstance) error
	ComputeInstanceGetInfo(ComputeInstance) (ComputeInstanceInfo, error)
	DeviceClearAccountingPids(Device) error
	DeviceClearCpuAffinity(Device) error
	DeviceClearEccErrorCounts(Device, EccCounterType) error
	DeviceClearFieldValues(Device, []FieldValue) error
	DeviceCreateGpuInstance(Device, *GpuInstanceProfileInfo) (GpuInstance, error)
	DeviceCreateGpuInstanceWithPlacement(Device, *GpuInstanceProfileInfo, *GpuInstancePlacement) (GpuInstance, error)
	DeviceDiscoverGpus() (PciInfo, error)
	DeviceFreezeNvLinkUtilizationCounter(Device, int, int, EnableState) error
	DeviceGetAPIRestriction(Device, RestrictedAPI) (EnableState, error)
	DeviceGetAccountingBufferSize(Device) (int, error)
	DeviceGetAccountingMode(Device) (EnableState, error)
	DeviceGetAccountingPids(Device) ([]int, error)
	DeviceGetAccountingStats(Device, uint32) (AccountingStats, error)
	DeviceGetActiveVgpus(Device) ([]VgpuInstance, error)
	DeviceGetAdaptiveClockInfoStatus(Device) (uint32, error)
	DeviceGetApplicationsClock(Device, ClockType) (uint32, error)
	DeviceGetArchitecture(Device) (DeviceArchitecture, error)
	DeviceGetAttributes(Device) (DeviceAttributes, error)
	DeviceGetAutoBoostedClocksEnabled(Device) (EnableState, EnableState, error)
	DeviceGetBAR1MemoryInfo(Device) (BAR1Memory, error)
	DeviceGetBoardId(Device) (uint32, error)
	DeviceGetBoardPartNumber(Device) (string, error)
	DeviceGetBrand(Device) (BrandType, error)
	DeviceGetBridgeChipInfo(Device) (BridgeChipHierarchy, error)
	DeviceGetBusType(Device) (BusType, error)
	DeviceGetC2cModeInfoV(Device) C2cModeInfoHandler
	DeviceGetClkMonStatus(Device) (ClkMonStatus, error)
	DeviceGetClock(Device, ClockType, ClockId) (uint32, error)
	DeviceGetClockInfo(Device, ClockType) (uint32, error)
	DeviceGetComputeInstanceId(Device) (int, error)
	DeviceGetComputeMode(Device) (ComputeMode, error)
	DeviceGetComputeRunningProcesses(Device) ([]ProcessInfo, error)
	DeviceGetConfComputeGpuAttestationReport(Device) (ConfComputeGpuAttestationReport, error)
	DeviceGetConfComputeGpuCertificate(Device) (ConfComputeGpuCertificate, error)
	DeviceGetConfComputeMemSizeInfo(Device) (ConfComputeMemSizeInfo, error)
	DeviceGetConfComputeProtectedMemoryUsage(Device) (Memory, error)
	DeviceGetCount() (int, error)
	DeviceGetCpuAffinity(Device, int) ([]uint, error)
	DeviceGetCpuAffinityWithinScope(Device, int, AffinityScope) ([]uint, error)
	DeviceGetCreatableVgpus(Device) ([]VgpuTypeId, error)
	DeviceGetCudaComputeCapability(Device) (int, int, error)
	DeviceGetCurrPcieLinkGeneration(Device) (int, error)
	DeviceGetCurrPcieLinkWidth(Device) (int, error)
	DeviceGetCurrentClocksEventReasons(Device) (uint64, error)
	DeviceGetCurrentClocksThrottleReasons(Device) (uint64, error)
	DeviceGetDecoderUtilization(Device) (uint32, uint32, error)
	DeviceGetDefaultApplicationsClock(Device, ClockType) (uint32, error)
	DeviceGetDefaultEccMode(Device) (EnableState, error)
	DeviceGetDetailedEccErrors(Device, MemoryErrorType, EccCounterType) (EccErrorCounts, error)
	DeviceGetDeviceHandleFromMigDeviceHandle(Device) (Device, error)
	DeviceGetDisplayActive(Device) (EnableState, error)
	DeviceGetDisplayMode(Device) (EnableState, error)
	DeviceGetDriverModel(Device) (DriverModel, DriverModel, error)
	DeviceGetDynamicPstatesInfo(Device) (GpuDynamicPstatesInfo, error)
	DeviceGetEccMode(Device) (EnableState, EnableState, error)
	DeviceGetEncoderCapacity(Device, EncoderType) (int, error)
	DeviceGetEncoderSessions(Device) ([]EncoderSessionInfo, error)
	DeviceGetEncoderStats(Device) (int, uint32, uint32, error)
	DeviceGetEncoderUtilization(Device) (uint32, uint32, error)
	DeviceGetEnforcedPowerLimit(Device) (uint32, error)
	DeviceGetFBCSessions(Device) ([]FBCSessionInfo, error)
	DeviceGetFBCStats(Device) (FBCStats, error)
	DeviceGetFanControlPolicy_v2(Device, int) (FanControlPolicy, error)
	DeviceGetFanSpeed(Device) (uint32, error)
	DeviceGetFanSpeed_v2(Device, int) (uint32, error)
	DeviceGetFieldValues(Device, []FieldValue) error
	DeviceGetGpcClkMinMaxVfOffset(Device) (int, int, error)
	DeviceGetGpcClkVfOffset(Device) (int, error)
	DeviceGetGpuFabricInfo(Device) (GpuFabricInfo, error)
	DeviceGetGpuFabricInfoV(Device) GpuFabricInfoHandler
	DeviceGetGpuInstanceById(Device, int) (GpuInstance, error)
	DeviceGetGpuInstanceId(Device) (int, error)
	DeviceGetGpuInstancePossiblePlacements(Device, *GpuInstanceProfileInfo) ([]GpuInstancePlacement, error)
	DeviceGetGpuInstanceProfileInfo(Device, int) (GpuInstanceProfileInfo, error)
	DeviceGetGpuInstanceProfileInfoV(Device, int) GpuInstanceProfileInfoHandler
	DeviceGetGpuInstanceRemainingCapacity(Device, *GpuInstanceProfileInfo) (int, error)
	DeviceGetGpuInstances(Device, *GpuInstanceProfileInfo) ([]GpuInstance, error)
	DeviceGetGpuMaxPcieLinkGeneration(Device) (int, error)
	DeviceGetGpuOperationMode(Device) (GpuOperationMode, GpuOperationMode, error)
	DeviceGetGraphicsRunningProcesses(Device) ([]ProcessInfo, error)
	DeviceGetGridLicensableFeatures(Device) (GridLicensableFeatures, error)
	DeviceGetGspFirmwareMode(Device) (bool, bool, error)
	DeviceGetGspFirmwareVersion(Device) (string, error)
	DeviceGetHandleByIndex(int) (Device, error)
	DeviceGetHandleByPciBusId(string) (Device, error)
	DeviceGetHandleBySerial(string) (Device, error)
	DeviceGetHandleByUUID(string) (Device, error)
	DeviceGetHostVgpuMode(Device) (HostVgpuMode, error)
	DeviceGetIndex(Device) (int, error)
	DeviceGetInforomConfigurationChecksum(Device) (uint32, error)
	DeviceGetInforomImageVersion(Device) (string, error)
	DeviceGetInforomVersion(Device, InforomObject) (string, error)
	DeviceGetIrqNum(Device) (int, error)
	DeviceGetJpgUtilization(Device) (uint32, uint32, error)
	DeviceGetLastBBXFlushTime(Device) (uint64, uint, error)
	DeviceGetMPSComputeRunningProcesses(Device) ([]ProcessInfo, error)
	DeviceGetMaxClockInfo(Device, ClockType) (uint32, error)
	DeviceGetMaxCustomerBoostClock(Device, ClockType) (uint32, error)
	DeviceGetMaxMigDeviceCount(Device) (int, error)
	DeviceGetMaxPcieLinkGeneration(Device) (int, error)
	DeviceGetMaxPcieLinkWidth(Device) (int, error)
	DeviceGetMemClkMinMaxVfOffset(Device) (int, int, error)
	DeviceGetMemClkVfOffset(Device) (int, error)
	DeviceGetMemoryAffinity(Device, int, AffinityScope) ([]uint, error)
	DeviceGetMemoryBusWidth(Device) (uint32, error)
	DeviceGetMemoryErrorCounter(Device, MemoryErrorType, EccCounterType, MemoryLocation) (uint64, error)
	DeviceGetMemoryInfo(Device) (Memory, error)
	DeviceGetMemoryInfo_v2(Device) (Memory_v2, error)
	DeviceGetMigDeviceHandleByIndex(Device, int) (Device, error)
	DeviceGetMigMode(Device) (int, int, error)
	DeviceGetMinMaxClockOfPState(Device, ClockType, Pstates) (uint32, uint32, error)
	DeviceGetMinMaxFanSpeed(Device) (int, int, error)
	DeviceGetMinorNumber(Device) (int, error)
	DeviceGetModuleId(Device) (int, error)
	DeviceGetMultiGpuBoard(Device) (int, error)
	DeviceGetName(Device) (string, error)
	DeviceGetNumFans(Device) (int, error)
	DeviceGetNumGpuCores(Device) (int, error)
	DeviceGetNumaNodeId(Device) (int, error)
	DeviceGetNvLinkCapability(Device, int, NvLinkCapability) (uint32, error)
	DeviceGetNvLinkErrorCounter(Device, int, NvLinkErrorCounter) (uint64, error)
	DeviceGetNvLinkRemoteDeviceType(Device, int) (IntNvLinkDeviceType, error)
	DeviceGetNvLinkRemotePciInfo(Device, int) (PciInfo, error)
	DeviceGetNvLinkState(Device, int) (EnableState, error)
	DeviceGetNvLinkUtilizationControl(Device, int, int) (NvLinkUtilizationControl, error)
	DeviceGetNvLinkUtilizationCounter(Device, int, int) (uint64, uint64, error)
	DeviceGetNvLinkVersion(Device, int) (uint32, error)
	DeviceGetOfaUtilization(Device) (uint32, uint32, error)
	DeviceGetP2PStatus(Device, Device, GpuP2PCapsIndex) (GpuP2PStatus, error)
	DeviceGetPciInfo(Device) (PciInfo, error)
	DeviceGetPciInfoExt(Device) (PciInfoExt, error)
	DeviceGetPcieLinkMaxSpeed(Device) (uint32, error)
	DeviceGetPcieReplayCounter(Device) (int, error)
	DeviceGetPcieSpeed(Device) (int, error)
	DeviceGetPcieThroughput(Device, PcieUtilCounter) (uint32, error)
	DeviceGetPerformanceState(Device) (Pstates, error)
	DeviceGetPersistenceMode(Device) (EnableState, error)
	DeviceGetPgpuMetadataString(Device) (string, error)
	DeviceGetPowerManagementDefaultLimit(Device) (uint32, error)
	DeviceGetPowerManagementLimit(Device) (uint32, error)
	DeviceGetPowerManagementLimitConstraints(Device) (uint32, uint32, error)
	DeviceGetPowerManagementMode(Device) (EnableState, error)
	DeviceGetPowerSource(Device) (PowerSource, error)
	DeviceGetPowerState(Device) (Pstates, error)
	DeviceGetPowerUsage(Device) (uint32, error)
	DeviceGetProcessUtilization(Device, uint64) ([]ProcessUtilizationSample, error)
	DeviceGetProcessesUtilizationInfo(Device) (ProcessesUtilizationInfo, error)
	DeviceGetRemappedRows(Device) (int, int, bool, bool, error)
	DeviceGetRetiredPages(Device, PageRetirementCause) ([]uint64, error)
	DeviceGetRetiredPagesPendingStatus(Device) (EnableState, error)
	DeviceGetRetiredPages_v2(Device, PageRetirementCause) ([]uint64, []uint64, error)
	DeviceGetRowRemapperHistogram(Device) (RowRemapperHistogramValues, error)
	DeviceGetRunningProcessDetailList(Device) (ProcessDetailList, error)
	DeviceGetSamples(Device, SamplingType, uint64) (ValueType, []Sample, error)
	DeviceGetSerial(Device) (string, error)
	DeviceGetSramEccErrorStatus(Device) (EccSramErrorStatus, error)
	DeviceGetSupportedClocksEventReasons(Device) (uint64, error)
	DeviceGetSupportedClocksThrottleReasons(Device) (uint64, error)
	DeviceGetSupportedEventTypes(Device) (uint64, error)
	DeviceGetSupportedGraphicsClocks(Device, int) (int, uint32, error)
	DeviceGetSupportedMemoryClocks(Device) (int, uint32, error)
	DeviceGetSupportedPerformanceStates(Device) ([]Pstates, error)
	DeviceGetSupportedVgpus(Device) ([]VgpuTypeId, error)
	DeviceGetTargetFanSpeed(Device, int) (int, error)
	DeviceGetTemperature(Device, TemperatureSensors) (uint32, error)
	DeviceGetTemperatureThreshold(Device, TemperatureThresholds) (uint32, error)
	DeviceGetThermalSettings(Device, uint32) (GpuThermalSettings, error)
	DeviceGetTopologyCommonAncestor(Device, Device) (GpuTopologyLevel, error)
	DeviceGetTopologyNearestGpus(Device, GpuTopologyLevel) ([]Device, error)
	DeviceGetTotalEccErrors(Device, MemoryErrorType, EccCounterType) (uint64, error)
	DeviceGetTotalEnergyConsumption(Device) (uint64, error)
	DeviceGetUUID(Device) (string, error)
	DeviceGetUtilizationRates(Device) (Utilization, error)
	DeviceGetVbiosVersion(Device) (string, error)
	DeviceGetVgpuCapabilities(Device, DeviceVgpuCapability) (bool, error)
	DeviceGetVgpuHeterogeneousMode(Device) (VgpuHeterogeneousMode, error)
	DeviceGetVgpuInstancesUtilizationInfo(Device) (VgpuInstancesUtilizationInfo, error)
	DeviceGetVgpuMetadata(Device) (VgpuPgpuMetadata, error)
	DeviceGetVgpuProcessUtilization(Device, uint64) ([]VgpuProcessUtilizationSample, error)
	DeviceGetVgpuProcessesUtilizationInfo(Device) (VgpuProcessesUtilizationInfo, error)
	DeviceGetVgpuSchedulerCapabilities(Device) (VgpuSchedulerCapabilities, error)
	DeviceGetVgpuSchedulerLog(Device) (VgpuSchedulerLog, error)
	DeviceGetVgpuSchedulerState(Device) (VgpuSchedulerGetState, error)
	DeviceGetVgpuTypeCreatablePlacements(Device, VgpuTypeId) (VgpuPlacementList, error)
	DeviceGetVgpuTypeSupportedPlacements(Device, VgpuTypeId) (VgpuPlacementList, error)
	DeviceGetVgpuUtilization(Device, uint64) (ValueType, []VgpuInstanceUtilizationSample, error)
	DeviceGetViolationStatus(Device, PerfPolicyType) (ViolationTime, error)
	DeviceGetVirtualizationMode(Device) (GpuVirtualizationMode, error)
	DeviceIsMigDeviceHandle(Device) (bool, error)
	DeviceModifyDrainState(*PciInfo, EnableState) error
	DeviceOnSameBoard(Device, Device) (int, error)
	DeviceQueryDrainState(*PciInfo) (EnableState, error)
	DeviceRegisterEvents(Device, uint64, EventSet) error
	DeviceRemoveGpu(*PciInfo) error
	DeviceRemoveGpu_v2(*PciInfo, DetachGpuState, PcieLinkState) error
	DeviceResetApplicationsClocks(Device) error
	DeviceResetGpuLockedClocks(Device) error
	DeviceResetMemoryLockedClocks(Device) error
	DeviceResetNvLinkErrorCounters(Device, int) error
	DeviceResetNvLinkUtilizationCounter(Device, int, int) error
	DeviceSetAPIRestriction(Device, RestrictedAPI, EnableState) error
	DeviceSetAccountingMode(Device, EnableState) error
	DeviceSetApplicationsClocks(Device, uint32, uint32) error
	DeviceSetAutoBoostedClocksEnabled(Device, EnableState) error
	DeviceSetComputeMode(Device, ComputeMode) error
	DeviceSetConfComputeUnprotectedMemSize(Device, uint64) error
	DeviceSetCpuAffinity(Device) error
	DeviceSetDefaultAutoBoostedClocksEnabled(Device, EnableState, uint32) error
	DeviceSetDefaultFanSpeed_v2(Device, int) error
	DeviceSetDriverModel(Device, DriverModel, uint32) error
	DeviceSetEccMode(Device, EnableState) error
	DeviceSetFanControlPolicy(Device, int, FanControlPolicy) error
	DeviceSetFanSpeed_v2(Device, int, int) error
	DeviceSetGpcClkVfOffset(Device, int) error
	DeviceSetGpuLockedClocks(Device, uint32, uint32) error
	DeviceSetGpuOperationMode(Device, GpuOperationMode) error
	DeviceSetMemClkVfOffset(Device, int) error
	DeviceSetMemoryLockedClocks(Device, uint32, uint32) error
	DeviceSetMigMode(Device, int) (error, error)
	DeviceSetNvLinkDeviceLowPowerThreshold(Device, *NvLinkPowerThres) error
	DeviceSetNvLinkUtilizationControl(Device, int, int, *NvLinkUtilizationControl, bool) error
	DeviceSetPersistenceMode(Device, EnableState) error
	DeviceSetPowerManagementLimit(Device, uint32) error
	DeviceSetPowerManagementLimit_v2(Device, *PowerValue_v2) error
	DeviceSetTemperatureThreshold(Device, TemperatureThresholds, int) error
	DeviceSetVgpuCapabilities(Device, DeviceVgpuCapability, EnableState) error
	DeviceSetVgpuHeterogeneousMode(Device, VgpuHeterogeneousMode) error
	DeviceSetVgpuSchedulerState(Device, *VgpuSchedulerSetState) error
	DeviceSetVirtualizationMode(Device, GpuVirtualizationMode) error
	DeviceValidateInforom(Device) error
	ErrorString(Return) string
	EventSetCreate() (EventSet, error)
	EventSetFree(EventSet) error
	EventSetWait(EventSet, uint32) (EventData, error)
	Extensions() ExtendedInterface
	GetExcludedDeviceCount() (int, error)
	GetExcludedDeviceInfoByIndex(int) (ExcludedDeviceInfo, error)
	GetVgpuCompatibility(*VgpuMetadata, *VgpuPgpuMetadata) (VgpuPgpuCompatibility, error)
	GetVgpuDriverCapabilities(VgpuDriverCapability) (bool, error)
	GetVgpuVersion() (VgpuVersion, VgpuVersion, error)
	GpmMetricsGet(*GpmMetricsGetType) error
	GpmMetricsGetV(*GpmMetricsGetType) GpmMetricsGetVType
	GpmMigSampleGet(Device, int, GpmSample) error
	GpmQueryDeviceSupport(Device) (GpmSupport, error)
	GpmQueryDeviceSupportV(Device) GpmSupportV
	GpmQueryIfStreamingEnabled(Device) (uint32, error)
	GpmSampleAlloc() (GpmSample, error)
	GpmSampleFree(GpmSample) error
	GpmSampleGet(Device, GpmSample) error
	GpmSetStreamingEnabled(Device, uint32) error
	GpuInstanceCreateComputeInstance(GpuInstance, *ComputeInstanceProfileInfo) (ComputeInstance, error)
	GpuInstanceCreateComputeInstanceWithPlacement(GpuInstance, *ComputeInstanceProfileInfo, *ComputeInstancePlacement) (ComputeInstance, error)
	GpuInstanceDestroy(GpuInstance) error
	GpuInstanceGetComputeInstanceById(GpuInstance, int) (ComputeInstance, error)
	GpuInstanceGetComputeInstancePossiblePlacements(GpuInstance, *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, error)
	GpuInstanceGetComputeInstanceProfileInfo(GpuInstance, int, int) (ComputeInstanceProfileInfo, error)
	GpuInstanceGetComputeInstanceProfileInfoV(GpuInstance, int, int) ComputeInstanceProfileInfoHandler
	GpuInstanceGetComputeInstanceRemainingCapacity(GpuInstance, *ComputeInstanceProfileInfo) (int, error)
	GpuInstanceGetComputeInstances(GpuInstance, *ComputeInstanceProfileInfo) ([]ComputeInstance, error)
	GpuInstanceGetInfo(GpuInstance) (GpuInstanceInfo, error)
	Init() error
	InitWithFlags(uint32) error
	SetVgpuVersion(*VgpuVersion) error
	Shutdown() error
	SystemGetConfComputeCapabilities() (ConfComputeSystemCaps, error)
	SystemGetConfComputeKeyRotationThresholdInfo() (ConfComputeGetKeyRotationThresholdInfo, error)
	SystemGetConfComputeSettings() (SystemConfComputeSettings, error)
	SystemGetCudaDriverVersion() (int, error)
	SystemGetCudaDriverVersion_v2() (int, error)
	SystemGetDriverVersion() (string, error)
	SystemGetHicVersion() ([]HwbcEntry, error)
	SystemGetNVMLVersion() (string, error)
	SystemGetProcessName(int) (string, error)
	SystemGetTopologyGpuSet(int) ([]Device, error)
	SystemSetConfComputeKeyRotationThresholdInfo(ConfComputeSetKeyRotationThresholdInfo) error
	UnitGetCount() (int, error)
	UnitGetDevices(Unit) ([]Device, error)
	UnitGetFanSpeedInfo(Unit) (UnitFanSpeeds, error)
	UnitGetHandleByIndex(int) (Unit, error)
	UnitGetLedState(Unit) (LedState, error)
	UnitGetPsuInfo(Unit) (PSUInfo, error)
	UnitGetTemperature(Unit, int) (uint32, error)
	UnitGetUnitInfo(Unit) (UnitInfo, error)
	UnitSetLedState(Unit, LedColor) error
	VgpuInstanceClearAccountingPids(VgpuInstance) error
	VgpuInstanceGetAccountingMode(VgpuInstance) (EnableState, error)
	VgpuInstanceGetAccountingPids(VgpuInstance) ([]int, error)
	VgpuInstanceGetAccountingStats(VgpuInstance, int) (AccountingStats, error)
	VgpuInstanceGetEccMode(VgpuInstance) (EnableState, error)
	VgpuInstanceGetEncoderCapacity(VgpuInstance) (int, error)
	VgpuInstanceGetEncoderSessions(VgpuInstance) (int, EncoderSessionInfo, error)
	VgpuInstanceGetEncoderStats(VgpuInstance) (int, uint32, uint32, error)
	VgpuInstanceGetFBCSessions(VgpuInstance) (int, FBCSessionInfo, error)
	VgpuInstanceGetFBCStats(VgpuInstance) (FBCStats, error)
	VgpuInstanceGetFbUsage(VgpuInstance) (uint64, error)
	VgpuInstanceGetFrameRateLimit(VgpuInstance) (uint32, error)
	VgpuInstanceGetGpuInstanceId(VgpuInstance) (int, error)
	VgpuInstanceGetGpuPciId(VgpuInstance) (string, error)
	VgpuInstanceGetLicenseInfo(VgpuInstance) (VgpuLicenseInfo, error)
	VgpuInstanceGetLicenseStatus(VgpuInstance) (int, error)
	VgpuInstanceGetMdevUUID(VgpuInstance) (string, error)
	VgpuInstanceGetMetadata(VgpuInstance) (VgpuMetadata, error)
	VgpuInstanceGetType(VgpuInstance) (VgpuTypeId, error)
	VgpuInstanceGetUUID(VgpuInstance) (string, error)
	VgpuInstanceGetVmDriverVersion(VgpuInstance) (string, error)
	VgpuInstanceGetVmID(VgpuInstance) (string, VgpuVmIdType, error)
	VgpuInstanceSetEncoderCapacity(VgpuInstance, int) error
	VgpuTypeGetCapabilities(VgpuTypeId, VgpuCapability) (bool, error)
	VgpuTypeGetClass(VgpuTypeId) (string, error)
	VgpuTypeGetDeviceID(VgpuTypeId) (uint64, uint64, error)
	VgpuTypeGetFrameRateLimit(VgpuTypeId) (uint32, error)
	VgpuTypeGetFramebufferSize(VgpuTypeId) (uint64, error)
	VgpuTypeGetGpuInstanceProfileId(VgpuTypeId) (uint32, error)
	VgpuTypeGetLicense(VgpuTypeId) (string, error)
	VgpuTypeGetMaxInstances(Device, VgpuTypeId) (int, error)
	VgpuTypeGetMaxInstancesPerVm(VgpuTypeId) (int, error)
	VgpuTypeGetName(VgpuTypeId) (string, error)
	VgpuTypeGetNumDisplayHeads(VgpuTypeId) (int, error)
	VgpuTypeGetResolution(VgpuTypeId, int) (uint32, uint32, error)
}

// Device represents the interface for the nvmlDevice type.
//
//go:generate moq -rm -out mock/device.go -pkg mock . Device:Device
type Device interface {
	ClearAccountingPids() error
	ClearCpuAffinity() error
	ClearEccErrorCounts(EccCounterType) error
	ClearFieldValues([]FieldValue) error
	CreateGpuInstance(*GpuInstanceProfileInfo) (GpuInstance, error)
	CreateGpuInstanceWithPlacement(*GpuInstanceProfileInfo, *GpuInstancePlacement) (GpuInstance, error)
	FreezeNvLinkUtilizationCounter(int, int, EnableState) error
	GetAPIRestriction(RestrictedAPI) (EnableState, error)
	GetAccountingBufferSize() (int, error)
	GetAccountingMode() (EnableState, error)
	GetAccountingPids() ([]int, error)
	GetAccountingStats(uint32) (AccountingStats, error)
	GetActiveVgpus() ([]VgpuInstance, error)
	GetAdaptiveClockInfoStatus() (uint32, error)
	GetApplicationsClock(ClockType) (uint32, error)
	GetArchitecture() (DeviceArchitecture, error)
	GetAttributes() (DeviceAttributes, error)
	GetAutoBoostedClocksEnabled() (EnableState, EnableState, error)
	GetBAR1MemoryInfo() (BAR1Memory, error)
	GetBoardId() (uint32, error)
	GetBoardPartNumber() (string, error)
	GetBrand() (BrandType, error)
	GetBridgeChipInfo() (BridgeChipHierarchy, error)
	GetBusType() (BusType, error)
	GetC2cModeInfoV() C2cModeInfoHandler
	GetClkMonStatus() (ClkMonStatus, error)
	GetClock(ClockType, ClockId) (uint32, error)
	GetClockInfo(ClockType) (uint32, error)
	GetComputeInstanceId() (int, error)
	GetComputeMode() (ComputeMode, error)
	GetComputeRunningProcesses() ([]ProcessInfo, error)
	GetConfComputeGpuAttestationReport() (ConfComputeGpuAttestationReport, error)
	GetConfComputeGpuCertificate() (ConfComputeGpuCertificate, error)
	GetConfComputeMemSizeInfo() (ConfComputeMemSizeInfo, error)
	GetConfComputeProtectedMemoryUsage() (Memory, error)
	GetCpuAffinity(int) ([]uint, error)
	GetCpuAffinityWithinScope(int, AffinityScope) ([]uint, error)
	GetCreatableVgpus() ([]VgpuTypeId, error)
	GetCudaComputeCapability() (int, int, error)
	GetCurrPcieLinkGeneration() (int, error)
	GetCurrPcieLinkWidth() (int, error)
	GetCurrentClocksEventReasons() (uint64, error)
	GetCurrentClocksThrottleReasons() (uint64, error)
	GetDecoderUtilization() (uint32, uint32, error)
	GetDefaultApplicationsClock(ClockType) (uint32, error)
	GetDefaultEccMode() (EnableState, error)
	GetDetailedEccErrors(MemoryErrorType, EccCounterType) (EccErrorCounts, error)
	GetDeviceHandleFromMigDeviceHandle() (Device, error)
	GetDisplayActive() (EnableState, error)
	GetDisplayMode() (EnableState, error)
	GetDriverModel() (DriverModel, DriverModel, error)
	GetDynamicPstatesInfo() (GpuDynamicPstatesInfo, error)
	GetEccMode() (EnableState, EnableState, error)
	GetEncoderCapacity(EncoderType) (int, error)
	GetEncoderSessions() ([]EncoderSessionInfo, error)
	GetEncoderStats() (int, uint32, uint32, error)
	GetEncoderUtilization() (uint32, uint32, error)
	GetEnforcedPowerLimit() (uint32, error)
	GetFBCSessions() ([]FBCSessionInfo, error)
	GetFBCStats() (FBCStats, error)
	GetFanControlPolicy_v2(int) (FanControlPolicy, error)
	GetFanSpeed() (uint32, error)
	GetFanSpeed_v2(int) (uint32, error)
	GetFieldValues([]FieldValue) error
	GetGpcClkMinMaxVfOffset() (int, int, error)
	GetGpcClkVfOffset() (int, error)
	GetGpuFabricInfo() (GpuFabricInfo, error)
	GetGpuFabricInfoV() GpuFabricInfoHandler
	GetGpuInstanceById(int) (GpuInstance, error)
	GetGpuInstanceId() (int, error)
	GetGpuInstancePossiblePlacements(*GpuInstanceProfileInfo) ([]GpuInstancePlacement, error)
	GetGpuInstanceProfileInfo(int) (GpuInstanceProfileInfo, error)
	GetGpuInstanceProfileInfoV(int) GpuInstanceProfileInfoHandler
	GetGpuInstanceRemainingCapacity(*GpuInstanceProfileInfo) (int, error)
	GetGpuInstances(*GpuInstanceProfileInfo) ([]GpuInstance, error)
	GetGpuMaxPcieLinkGeneration() (int, error)
	GetGpuOperationMode() (GpuOperationMode, GpuOperationMode, error)
	GetGraphicsRunningProcesses() ([]ProcessInfo, error)
	GetGridLicensableFeatures() (GridLicensableFeatures, error)
	GetGspFirmwareMode() (bool, bool, error)
	GetGspFirmwareVersion() (string, error)
	GetHostVgpuMode() (HostVgpuMode, error)
	GetIndex() (int, error)
	GetInforomConfigurationChecksum() (uint32, error)
	GetInforomImageVersion() (string, error)
	GetInforomVersion(InforomObject) (string, error)
	GetIrqNum() (int, error)
	GetJpgUtilization() (uint32, uint32, error)
	GetLastBBXFlushTime() (uint64, uint, error)
	GetMPSComputeRunningProcesses() ([]ProcessInfo, error)
	GetMaxClockInfo(ClockType) (uint32, error)
	GetMaxCustomerBoostClock(ClockType) (uint32, error)
	GetMaxMigDeviceCount() (int, error)
	GetMaxPcieLinkGeneration() (int, error)
	GetMaxPcieLinkWidth() (int, error)
	GetMemClkMinMaxVfOffset() (int, int, error)
	GetMemClkVfOffset() (int, error)
	GetMemoryAffinity(int, AffinityScope) ([]uint, error)
	GetMemoryBusWidth() (uint32, error)
	GetMemoryErrorCounter(MemoryErrorType, EccCounterType, MemoryLocation) (uint64, error)
	GetMemoryInfo() (Memory, error)
	GetMemoryInfo_v2() (Memory_v2, error)
	GetMigDeviceHandleByIndex(int) (Device, error)
	GetMigMode() (int, int, error)
	GetMinMaxClockOfPState(ClockType, Pstates) (uint32, uint32, error)
	GetMinMaxFanSpeed() (int, int, error)
	GetMinorNumber() (int, error)
	GetModuleId() (int, error)
	GetMultiGpuBoard() (int, error)
	GetName() (string, error)
	GetNumFans() (int, error)
	GetNumGpuCores() (int, error)
	GetNumaNodeId() (int, error)
	GetNvLinkCapability(int, NvLinkCapability) (uint32, error)
	GetNvLinkErrorCounter(int, NvLinkErrorCounter) (uint64, error)
	GetNvLinkRemoteDeviceType(int) (IntNvLinkDeviceType, error)
	GetNvLinkRemotePciInfo(int) (PciInfo, error)
	GetNvLinkState(int) (EnableState, error)
	GetNvLinkUtilizationControl(int, int) (NvLinkUtilizationControl, error)
	GetNvLinkUtilizationCounter(int, int) (uint64, uint64, error)
	GetNvLinkVersion(int) (uint32, error)
	GetOfaUtilization() (uint32, uint32, error)
	GetP2PStatus(Device, GpuP2PCapsIndex) (GpuP2PStatus, error)
	GetPciInfo() (PciInfo, error)
	GetPciInfoExt() (PciInfoExt, error)
	GetPcieLinkMaxSpeed() (uint32, error)
	GetPcieReplayCounter() (int, error)
	GetPcieSpeed() (int, error)
	GetPcieThroughput(PcieUtilCounter) (uint32, error)
	GetPerformanceState() (Pstates, error)
	GetPersistenceMode() (EnableState, error)
	GetPgpuMetadataString() (string, error)
	GetPowerManagementDefaultLimit() (uint32, error)
	GetPowerManagementLimit() (uint32, error)
	GetPowerManagementLimitConstraints() (uint32, uint32, error)
	GetPowerManagementMode() (EnableState, error)
	GetPowerSource() (PowerSource, error)
	GetPowerState() (Pstates, error)
	GetPowerUsage() (uint32, error)
	GetProcessUtilization(uint64) ([]ProcessUtilizationSample, error)
	GetProcessesUtilizationInfo() (ProcessesUtilizationInfo, error)
	GetRemappedRows() (int, int, bool, bool, error)
	GetRetiredPages(PageRetirementCause) ([]uint64, error)
	GetRetiredPagesPendingStatus() (EnableState, error)
	GetRetiredPages_v2(PageRetirementCause) ([]uint64, []uint64, error)
	GetRowRemapperHistogram() (RowRemapperHistogramValues, error)
	GetRunningProcessDetailList() (ProcessDetailList, error)
	GetSamples(SamplingType, uint64) (ValueType, []Sample, error)
	GetSerial() (string, error)
	GetSramEccErrorStatus() (EccSramErrorStatus, error)
	GetSupportedClocksEventReasons() (uint64, error)
	GetSupportedClocksThrottleReasons() (uint64, error)
	GetSupportedEventTypes() (uint64, error)
	GetSupportedGraphicsClocks(int) (int, uint32, error)
	GetSupportedMemoryClocks() (int, uint32, error)
	GetSupportedPerformanceStates() ([]Pstates, error)
	GetSupportedVgpus() ([]VgpuTypeId, error)
	GetTargetFanSpeed(int) (int, error)
	GetTemperature(TemperatureSensors) (uint32, error)
	GetTemperatureThreshold(TemperatureThresholds) (uint32, error)
	GetThermalSettings(uint32) (GpuThermalSettings, error)
	GetTopologyCommonAncestor(Device) (GpuTopologyLevel, error)
	GetTopologyNearestGpus(GpuTopologyLevel) ([]Device, error)
	GetTotalEccErrors(MemoryErrorType, EccCounterType) (uint64, error)
	GetTotalEnergyConsumption() (uint64, error)
	GetUUID() (string, error)
	GetUtilizationRates() (Utilization, error)
	GetVbiosVersion() (string, error)
	GetVgpuCapabilities(DeviceVgpuCapability) (bool, error)
	GetVgpuHeterogeneousMode() (VgpuHeterogeneousMode, error)
	GetVgpuInstancesUtilizationInfo() (VgpuInstancesUtilizationInfo, error)
	GetVgpuMetadata() (VgpuPgpuMetadata, error)
	GetVgpuProcessUtilization(uint64) ([]VgpuProcessUtilizationSample, error)
	GetVgpuProcessesUtilizationInfo() (VgpuProcessesUtilizationInfo, error)
	GetVgpuSchedulerCapabilities() (VgpuSchedulerCapabilities, error)
	GetVgpuSchedulerLog() (VgpuSchedulerLog, error)
	GetVgpuSchedulerState() (VgpuSchedulerGetState, error)
	GetVgpuTypeCreatablePlacements(VgpuTypeId) (VgpuPlacementList, error)
	GetVgpuTypeSupportedPlacements(VgpuTypeId) (VgpuPlacementList, error)
	GetVgpuUtilization(uint64) (ValueType, []VgpuInstanceUtilizationSample, error)
	GetViolationStatus(PerfPolicyType) (ViolationTime, error)
	GetVirtualizationMode() (GpuVirtualizationMode, error)
	GpmMigSampleGet(int, GpmSample) error
	GpmQueryDeviceSupport() (GpmSupport, error)
	GpmQueryDeviceSupportV() GpmSupportV
	GpmQueryIfStreamingEnabled() (uint32, error)
	GpmSampleGet(GpmSample) error
	GpmSetStreamingEnabled(uint32) error
	IsMigDeviceHandle() (bool, error)
	OnSameBoard(Device) (int, error)
	RegisterEvents(uint64, EventSet) error
	ResetApplicationsClocks() error
	ResetGpuLockedClocks() error
	ResetMemoryLockedClocks() error
	ResetNvLinkErrorCounters(int) error
	ResetNvLinkUtilizationCounter(int, int) error
	SetAPIRestriction(RestrictedAPI, EnableState) error
	SetAccountingMode(EnableState) error
	SetApplicationsClocks(uint32, uint32) error
	SetAutoBoostedClocksEnabled(EnableState) error
	SetComputeMode(ComputeMode) error
	SetConfComputeUnprotectedMemSize(uint64) error
	SetCpuAffinity() error
	SetDefaultAutoBoostedClocksEnabled(EnableState, uint32) error
	SetDefaultFanSpeed_v2(int) error
	SetDriverModel(DriverModel, uint32) error
	SetEccMode(EnableState) error
	SetFanControlPolicy(int, FanControlPolicy) error
	SetFanSpeed_v2(int, int) error
	SetGpcClkVfOffset(int) error
	SetGpuLockedClocks(uint32, uint32) error
	SetGpuOperationMode(GpuOperationMode) error
	SetMemClkVfOffset(int) error
	SetMemoryLockedClocks(uint32, uint32) error
	SetMigMode(int) (error, error)
	SetNvLinkDeviceLowPowerThreshold(*NvLinkPowerThres) error
	SetNvLinkUtilizationControl(int, int, *NvLinkUtilizationControl, bool) error
	SetPersistenceMode(EnableState) error
	SetPowerManagementLimit(uint32) error
	SetPowerManagementLimit_v2(*PowerValue_v2) error
	SetTemperatureThreshold(TemperatureThresholds, int) error
	SetVgpuCapabilities(DeviceVgpuCapability, EnableState) error
	SetVgpuHeterogeneousMode(VgpuHeterogeneousMode) error
	SetVgpuSchedulerState(*VgpuSchedulerSetState) error
	SetVirtualizationMode(GpuVirtualizationMode) error
	ValidateInforom() error
	VgpuTypeGetMaxInstances(VgpuTypeId) (int, error)
}

// GpuInstance represents the interface for the nvmlGpuInstance type.
//
//go:generate moq -rm -out mock/gpuinstance.go -pkg mock . GpuInstance:GpuInstance
type GpuInstance interface {
	CreateComputeInstance(*ComputeInstanceProfileInfo) (ComputeInstance, error)
	CreateComputeInstanceWithPlacement(*ComputeInstanceProfileInfo, *ComputeInstancePlacement) (ComputeInstance, error)
	Destroy() error
	GetComputeInstanceById(int) (ComputeInstance, error)
	GetComputeInstancePossiblePlacements(*ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, error)
	GetComputeInstanceProfileInfo(int, int) (ComputeInstanceProfileInfo, error)
	GetComputeInstanceProfileInfoV(int, int) ComputeInstanceProfileInfoHandler
	GetComputeInstanceRemainingCapacity(*ComputeInstanceProfileInfo) (int, error)
	GetComputeInstances(*ComputeInstanceProfileInfo) ([]ComputeInstance, error)
	GetInfo() (GpuInstanceInfo, error)
}

// ComputeInstance represents the interface for the nvmlComputeInstance type.
//
//go:generate moq -rm -out mock/computeinstance.go -pkg mock . ComputeInstance:ComputeInstance
type ComputeInstance interface {
	Destroy() error
	GetInfo() (ComputeInstanceInfo, error)
}

// EventSet represents the interface for the nvmlEventSet type.
//
//go:generate moq -rm -out mock/eventset.go -pkg mock . EventSet:EventSet
type EventSet interface {
	Free() error
	Wait(uint32) (EventData, error)
}

// GpmSample represents the interface for the nvmlGpmSample type.
//
//go:generate moq -rm -out mock/gpmsample.go -pkg mock . GpmSample:GpmSample
type GpmSample interface {
	Free() error
	Get(Device) error
	MigGet(Device, int) error
}

// Unit represents the interface for the nvmlUnit type.
//
//go:generate moq -rm -out mock/unit.go -pkg mock . Unit:Unit
type Unit interface {
	GetDevices() ([]Device, error)
	GetFanSpeedInfo() (UnitFanSpeeds, error)
	GetLedState() (LedState, error)
	GetPsuInfo() (PSUInfo, error)
	GetTemperature(int) (uint32, error)
	GetUnitInfo() (UnitInfo, error)
	SetLedState(LedColor) error
}

// VgpuInstance represents the interface for the nvmlVgpuInstance type.
//
//go:generate moq -rm -out mock/vgpuinstance.go -pkg mock . VgpuInstance:VgpuInstance
type VgpuInstance interface {
	ClearAccountingPids() error
	GetAccountingMode() (EnableState, error)
	GetAccountingPids() ([]int, error)
	GetAccountingStats(int) (AccountingStats, error)
	GetEccMode() (EnableState, error)
	GetEncoderCapacity() (int, error)
	GetEncoderSessions() (int, EncoderSessionInfo, error)
	GetEncoderStats() (int, uint32, uint32, error)
	GetFBCSessions() (int, FBCSessionInfo, error)
	GetFBCStats() (FBCStats, error)
	GetFbUsage() (uint64, error)
	GetFrameRateLimit() (uint32, error)
	GetGpuInstanceId() (int, error)
	GetGpuPciId() (string, error)
	GetLicenseInfo() (VgpuLicenseInfo, error)
	GetLicenseStatus() (int, error)
	GetMdevUUID() (string, error)
	GetMetadata() (VgpuMetadata, error)
	GetType() (VgpuTypeId, error)
	GetUUID() (string, error)
	GetVmDriverVersion() (string, error)
	GetVmID() (string, VgpuVmIdType, error)
	SetEncoderCapacity(int) error
}

// VgpuTypeId represents the interface for the nvmlVgpuTypeId type.
//
//go:generate moq -rm -out mock/vgputypeid.go -pkg mock . VgpuTypeId:VgpuTypeId
type VgpuTypeId interface {
	GetCapabilities(VgpuCapability) (bool, error)
	GetClass() (string, error)
	GetCreatablePlacements(Device) (VgpuPlacementList, error)
	GetDeviceID() (uint64, uint64, error)
	GetFrameRateLimit() (uint32, error)
	GetFramebufferSize() (uint64, error)
	GetGpuInstanceProfileId() (uint32, error)
	GetLicense() (string, error)
	GetMaxInstances(Device) (int, error)
	GetMaxInstancesPerVm() (int, error)
	GetName() (string, error)
	GetNumDisplayHeads() (int, error)
	GetResolution(int) (uint32, uint32, error)
	GetSupportedPlacements(Device) (VgpuPlacementList, error)
}
