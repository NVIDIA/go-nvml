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

// EccBitType
type EccBitType = MemoryErrorType

// nvml.DeviceGetCount()
func (l *library) DeviceGetCount() (int, Return) {
	var DeviceCount uint32
	ret := nvmlDeviceGetCount(&DeviceCount)
	return int(DeviceCount), ret
}

// nvml.DeviceGetHandleByIndex()
func (l *library) DeviceGetHandleByIndex(Index int) (Device, Return) {
	var Device nvmlDevice
	ret := nvmlDeviceGetHandleByIndex(uint32(Index), &Device)
	return Device, ret
}

// nvml.DeviceGetHandleBySerial()
func (l *library) DeviceGetHandleBySerial(Serial string) (Device, Return) {
	var Device nvmlDevice
	ret := nvmlDeviceGetHandleBySerial(Serial+string(rune(0)), &Device)
	return Device, ret
}

// nvml.DeviceGetHandleByUUID()
func (l *library) DeviceGetHandleByUUID(Uuid string) (Device, Return) {
	var Device nvmlDevice
	ret := nvmlDeviceGetHandleByUUID(Uuid+string(rune(0)), &Device)
	return Device, ret
}

// nvml.DeviceGetHandleByPciBusId()
func (l *library) DeviceGetHandleByPciBusId(PciBusId string) (Device, Return) {
	var Device nvmlDevice
	ret := nvmlDeviceGetHandleByPciBusId(PciBusId+string(rune(0)), &Device)
	return Device, ret
}

// nvml.DeviceGetName()
func (l *library) DeviceGetName(Device Device) (string, Return) {
	return Device.GetName()
}

func (Device nvmlDevice) GetName() (string, Return) {
	Name := make([]byte, DEVICE_NAME_V2_BUFFER_SIZE)
	ret := nvmlDeviceGetName(Device, &Name[0], DEVICE_NAME_V2_BUFFER_SIZE)
	return string(Name[:clen(Name)]), ret
}

// nvml.DeviceGetBrand()
func (l *library) DeviceGetBrand(Device Device) (BrandType, Return) {
	return Device.GetBrand()
}

func (Device nvmlDevice) GetBrand() (BrandType, Return) {
	var _type BrandType
	ret := nvmlDeviceGetBrand(Device, &_type)
	return _type, ret
}

// nvml.DeviceGetIndex()
func (l *library) DeviceGetIndex(Device Device) (int, Return) {
	return Device.GetIndex()
}

func (Device nvmlDevice) GetIndex() (int, Return) {
	var Index uint32
	ret := nvmlDeviceGetIndex(Device, &Index)
	return int(Index), ret
}

// nvml.DeviceGetSerial()
func (l *library) DeviceGetSerial(Device Device) (string, Return) {
	return Device.GetSerial()
}

func (Device nvmlDevice) GetSerial() (string, Return) {
	Serial := make([]byte, DEVICE_SERIAL_BUFFER_SIZE)
	ret := nvmlDeviceGetSerial(Device, &Serial[0], DEVICE_SERIAL_BUFFER_SIZE)
	return string(Serial[:clen(Serial)]), ret
}

// nvml.DeviceGetCpuAffinity()
func (l *library) DeviceGetCpuAffinity(Device Device, NumCPUs int) ([]uint, Return) {
	return Device.GetCpuAffinity(NumCPUs)
}

func (Device nvmlDevice) GetCpuAffinity(NumCPUs int) ([]uint, Return) {
	CpuSetSize := uint32((NumCPUs-1)/int(unsafe.Sizeof(uint(0))) + 1)
	CpuSet := make([]uint, CpuSetSize)
	ret := nvmlDeviceGetCpuAffinity(Device, CpuSetSize, &CpuSet[0])
	return CpuSet, ret
}

// nvml.DeviceSetCpuAffinity()
func (l *library) DeviceSetCpuAffinity(Device Device) Return {
	return Device.SetCpuAffinity()
}

func (Device nvmlDevice) SetCpuAffinity() Return {
	return nvmlDeviceSetCpuAffinity(Device)
}

// nvml.DeviceClearCpuAffinity()
func (l *library) DeviceClearCpuAffinity(Device Device) Return {
	return Device.ClearCpuAffinity()
}

func (Device nvmlDevice) ClearCpuAffinity() Return {
	return nvmlDeviceClearCpuAffinity(Device)
}

// nvml.DeviceGetMemoryAffinity()
func (l *library) DeviceGetMemoryAffinity(Device Device, NumNodes int, Scope AffinityScope) ([]uint, Return) {
	return Device.GetMemoryAffinity(NumNodes, Scope)
}

func (Device nvmlDevice) GetMemoryAffinity(NumNodes int, Scope AffinityScope) ([]uint, Return) {
	NodeSetSize := uint32((NumNodes-1)/int(unsafe.Sizeof(uint(0))) + 1)
	NodeSet := make([]uint, NodeSetSize)
	ret := nvmlDeviceGetMemoryAffinity(Device, NodeSetSize, &NodeSet[0], Scope)
	return NodeSet, ret
}

// nvml.DeviceGetCpuAffinityWithinScope()
func (l *library) DeviceGetCpuAffinityWithinScope(Device Device, NumCPUs int, Scope AffinityScope) ([]uint, Return) {
	return Device.GetCpuAffinityWithinScope(NumCPUs, Scope)
}

func (Device nvmlDevice) GetCpuAffinityWithinScope(NumCPUs int, Scope AffinityScope) ([]uint, Return) {
	CpuSetSize := uint32((NumCPUs-1)/int(unsafe.Sizeof(uint(0))) + 1)
	CpuSet := make([]uint, CpuSetSize)
	ret := nvmlDeviceGetCpuAffinityWithinScope(Device, CpuSetSize, &CpuSet[0], Scope)
	return CpuSet, ret
}

// nvml.DeviceGetTopologyCommonAncestor()
func (l *library) DeviceGetTopologyCommonAncestor(Device1 Device, Device2 Device) (GpuTopologyLevel, Return) {
	return Device1.GetTopologyCommonAncestor(Device2)
}

func (Device1 nvmlDevice) GetTopologyCommonAncestor(Device2 Device) (GpuTopologyLevel, Return) {
	var PathInfo GpuTopologyLevel
	ret := nvmlDeviceGetTopologyCommonAncestor(Device1, Device2.(nvmlDevice), &PathInfo)
	return PathInfo, ret
}

// nvml.DeviceGetTopologyNearestGpus()
func (l *library) DeviceGetTopologyNearestGpus(Device Device, Level GpuTopologyLevel) ([]Device, Return) {
	return Device.GetTopologyNearestGpus(Level)
}

func (device nvmlDevice) GetTopologyNearestGpus(Level GpuTopologyLevel) ([]Device, Return) {
	var Count uint32
	ret := nvmlDeviceGetTopologyNearestGpus(device, Level, &Count, nil)
	if ret != SUCCESS {
		return nil, ret
	}
	if Count == 0 {
		return []Device{}, ret
	}
	DeviceArray := make([]nvmlDevice, Count)
	ret = nvmlDeviceGetTopologyNearestGpus(device, Level, &Count, &DeviceArray[0])
	return convertSlice[nvmlDevice, Device](DeviceArray), ret
}

// nvml.DeviceGetP2PStatus()
func (l *library) DeviceGetP2PStatus(Device1 Device, Device2 Device, P2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return) {
	return Device1.GetP2PStatus(Device2, P2pIndex)
}

func (Device1 nvmlDevice) GetP2PStatus(Device2 Device, P2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return) {
	var P2pStatus GpuP2PStatus
	ret := nvmlDeviceGetP2PStatus(Device1, Device2.(nvmlDevice), P2pIndex, &P2pStatus)
	return P2pStatus, ret
}

// nvml.DeviceGetUUID()
func (l *library) DeviceGetUUID(Device Device) (string, Return) {
	return Device.GetUUID()
}

func (Device nvmlDevice) GetUUID() (string, Return) {
	Uuid := make([]byte, DEVICE_UUID_V2_BUFFER_SIZE)
	ret := nvmlDeviceGetUUID(Device, &Uuid[0], DEVICE_UUID_V2_BUFFER_SIZE)
	return string(Uuid[:clen(Uuid)]), ret
}

// nvml.DeviceGetMinorNumber()
func (l *library) DeviceGetMinorNumber(Device Device) (int, Return) {
	return Device.GetMinorNumber()
}

func (Device nvmlDevice) GetMinorNumber() (int, Return) {
	var MinorNumber uint32
	ret := nvmlDeviceGetMinorNumber(Device, &MinorNumber)
	return int(MinorNumber), ret
}

// nvml.DeviceGetBoardPartNumber()
func (l *library) DeviceGetBoardPartNumber(Device Device) (string, Return) {
	return Device.GetBoardPartNumber()
}

func (Device nvmlDevice) GetBoardPartNumber() (string, Return) {
	PartNumber := make([]byte, DEVICE_PART_NUMBER_BUFFER_SIZE)
	ret := nvmlDeviceGetBoardPartNumber(Device, &PartNumber[0], DEVICE_PART_NUMBER_BUFFER_SIZE)
	return string(PartNumber[:clen(PartNumber)]), ret
}

// nvml.DeviceGetInforomVersion()
func (l *library) DeviceGetInforomVersion(Device Device, Object InforomObject) (string, Return) {
	return Device.GetInforomVersion(Object)
}

func (Device nvmlDevice) GetInforomVersion(Object InforomObject) (string, Return) {
	Version := make([]byte, DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetInforomVersion(Device, Object, &Version[0], DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.DeviceGetInforomImageVersion()
func (l *library) DeviceGetInforomImageVersion(Device Device) (string, Return) {
	return Device.GetInforomImageVersion()
}

func (Device nvmlDevice) GetInforomImageVersion() (string, Return) {
	Version := make([]byte, DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetInforomImageVersion(Device, &Version[0], DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.DeviceGetInforomConfigurationChecksum()
func (l *library) DeviceGetInforomConfigurationChecksum(Device Device) (uint32, Return) {
	return Device.GetInforomConfigurationChecksum()
}

func (Device nvmlDevice) GetInforomConfigurationChecksum() (uint32, Return) {
	var Checksum uint32
	ret := nvmlDeviceGetInforomConfigurationChecksum(Device, &Checksum)
	return Checksum, ret
}

// nvml.DeviceValidateInforom()
func (l *library) DeviceValidateInforom(Device Device) Return {
	return Device.ValidateInforom()
}

func (Device nvmlDevice) ValidateInforom() Return {
	return nvmlDeviceValidateInforom(Device)
}

// nvml.DeviceGetDisplayMode()
func (l *library) DeviceGetDisplayMode(Device Device) (EnableState, Return) {
	return Device.GetDisplayMode()
}

func (Device nvmlDevice) GetDisplayMode() (EnableState, Return) {
	var Display EnableState
	ret := nvmlDeviceGetDisplayMode(Device, &Display)
	return Display, ret
}

// nvml.DeviceGetDisplayActive()
func (l *library) DeviceGetDisplayActive(Device Device) (EnableState, Return) {
	return Device.GetDisplayActive()
}

func (Device nvmlDevice) GetDisplayActive() (EnableState, Return) {
	var IsActive EnableState
	ret := nvmlDeviceGetDisplayActive(Device, &IsActive)
	return IsActive, ret
}

// nvml.DeviceGetPersistenceMode()
func (l *library) DeviceGetPersistenceMode(Device Device) (EnableState, Return) {
	return Device.GetPersistenceMode()
}

func (Device nvmlDevice) GetPersistenceMode() (EnableState, Return) {
	var Mode EnableState
	ret := nvmlDeviceGetPersistenceMode(Device, &Mode)
	return Mode, ret
}

// nvml.DeviceGetPciInfo()
func (l *library) DeviceGetPciInfo(Device Device) (PciInfo, Return) {
	return Device.GetPciInfo()
}

func (Device nvmlDevice) GetPciInfo() (PciInfo, Return) {
	var Pci PciInfo
	ret := nvmlDeviceGetPciInfo(Device, &Pci)
	return Pci, ret
}

// nvml.DeviceGetMaxPcieLinkGeneration()
func (l *library) DeviceGetMaxPcieLinkGeneration(Device Device) (int, Return) {
	return Device.GetMaxPcieLinkGeneration()
}

func (Device nvmlDevice) GetMaxPcieLinkGeneration() (int, Return) {
	var MaxLinkGen uint32
	ret := nvmlDeviceGetMaxPcieLinkGeneration(Device, &MaxLinkGen)
	return int(MaxLinkGen), ret
}

// nvml.DeviceGetMaxPcieLinkWidth()
func (l *library) DeviceGetMaxPcieLinkWidth(Device Device) (int, Return) {
	return Device.GetMaxPcieLinkWidth()
}

func (Device nvmlDevice) GetMaxPcieLinkWidth() (int, Return) {
	var MaxLinkWidth uint32
	ret := nvmlDeviceGetMaxPcieLinkWidth(Device, &MaxLinkWidth)
	return int(MaxLinkWidth), ret
}

// nvml.DeviceGetCurrPcieLinkGeneration()
func (l *library) DeviceGetCurrPcieLinkGeneration(Device Device) (int, Return) {
	return Device.GetCurrPcieLinkGeneration()
}

func (Device nvmlDevice) GetCurrPcieLinkGeneration() (int, Return) {
	var CurrLinkGen uint32
	ret := nvmlDeviceGetCurrPcieLinkGeneration(Device, &CurrLinkGen)
	return int(CurrLinkGen), ret
}

// nvml.DeviceGetCurrPcieLinkWidth()
func (l *library) DeviceGetCurrPcieLinkWidth(Device Device) (int, Return) {
	return Device.GetCurrPcieLinkWidth()
}

func (Device nvmlDevice) GetCurrPcieLinkWidth() (int, Return) {
	var CurrLinkWidth uint32
	ret := nvmlDeviceGetCurrPcieLinkWidth(Device, &CurrLinkWidth)
	return int(CurrLinkWidth), ret
}

// nvml.DeviceGetPcieThroughput()
func (l *library) DeviceGetPcieThroughput(Device Device, Counter PcieUtilCounter) (uint32, Return) {
	return Device.GetPcieThroughput(Counter)
}

func (Device nvmlDevice) GetPcieThroughput(Counter PcieUtilCounter) (uint32, Return) {
	var Value uint32
	ret := nvmlDeviceGetPcieThroughput(Device, Counter, &Value)
	return Value, ret
}

// nvml.DeviceGetPcieReplayCounter()
func (l *library) DeviceGetPcieReplayCounter(Device Device) (int, Return) {
	return Device.GetPcieReplayCounter()
}

func (Device nvmlDevice) GetPcieReplayCounter() (int, Return) {
	var Value uint32
	ret := nvmlDeviceGetPcieReplayCounter(Device, &Value)
	return int(Value), ret
}

// nvml.nvmlDeviceGetClockInfo()
func (l *library) DeviceGetClockInfo(Device Device, _type ClockType) (uint32, Return) {
	return Device.GetClockInfo(_type)
}

func (Device nvmlDevice) GetClockInfo(_type ClockType) (uint32, Return) {
	var Clock uint32
	ret := nvmlDeviceGetClockInfo(Device, _type, &Clock)
	return Clock, ret
}

// nvml.DeviceGetMaxClockInfo()
func (l *library) DeviceGetMaxClockInfo(Device Device, _type ClockType) (uint32, Return) {
	return Device.GetMaxClockInfo(_type)
}

func (Device nvmlDevice) GetMaxClockInfo(_type ClockType) (uint32, Return) {
	var Clock uint32
	ret := nvmlDeviceGetMaxClockInfo(Device, _type, &Clock)
	return Clock, ret
}

// nvml.DeviceGetApplicationsClock()
func (l *library) DeviceGetApplicationsClock(Device Device, ClockType ClockType) (uint32, Return) {
	return Device.GetApplicationsClock(ClockType)
}

func (Device nvmlDevice) GetApplicationsClock(ClockType ClockType) (uint32, Return) {
	var ClockMHz uint32
	ret := nvmlDeviceGetApplicationsClock(Device, ClockType, &ClockMHz)
	return ClockMHz, ret
}

// nvml.DeviceGetDefaultApplicationsClock()
func (l *library) DeviceGetDefaultApplicationsClock(Device Device, ClockType ClockType) (uint32, Return) {
	return Device.GetDefaultApplicationsClock(ClockType)
}

func (Device nvmlDevice) GetDefaultApplicationsClock(ClockType ClockType) (uint32, Return) {
	var ClockMHz uint32
	ret := nvmlDeviceGetDefaultApplicationsClock(Device, ClockType, &ClockMHz)
	return ClockMHz, ret
}

// nvml.DeviceResetApplicationsClocks()
func (l *library) DeviceResetApplicationsClocks(Device Device) Return {
	return Device.ResetApplicationsClocks()
}

func (Device nvmlDevice) ResetApplicationsClocks() Return {
	return nvmlDeviceResetApplicationsClocks(Device)
}

// nvml.DeviceGetClock()
func (l *library) DeviceGetClock(Device Device, ClockType ClockType, ClockId ClockId) (uint32, Return) {
	return Device.GetClock(ClockType, ClockId)
}

func (Device nvmlDevice) GetClock(ClockType ClockType, ClockId ClockId) (uint32, Return) {
	var ClockMHz uint32
	ret := nvmlDeviceGetClock(Device, ClockType, ClockId, &ClockMHz)
	return ClockMHz, ret
}

// nvml.DeviceGetMaxCustomerBoostClock()
func (l *library) DeviceGetMaxCustomerBoostClock(Device Device, ClockType ClockType) (uint32, Return) {
	return Device.GetMaxCustomerBoostClock(ClockType)
}

func (Device nvmlDevice) GetMaxCustomerBoostClock(ClockType ClockType) (uint32, Return) {
	var ClockMHz uint32
	ret := nvmlDeviceGetMaxCustomerBoostClock(Device, ClockType, &ClockMHz)
	return ClockMHz, ret
}

// nvml.DeviceGetSupportedMemoryClocks()
func (l *library) DeviceGetSupportedMemoryClocks(Device Device) (int, uint32, Return) {
	return Device.GetSupportedMemoryClocks()
}

func (Device nvmlDevice) GetSupportedMemoryClocks() (int, uint32, Return) {
	var Count, ClocksMHz uint32
	ret := nvmlDeviceGetSupportedMemoryClocks(Device, &Count, &ClocksMHz)
	return int(Count), ClocksMHz, ret
}

// nvml.DeviceGetSupportedGraphicsClocks()
func (l *library) DeviceGetSupportedGraphicsClocks(Device Device, MemoryClockMHz int) (int, uint32, Return) {
	return Device.GetSupportedGraphicsClocks(MemoryClockMHz)
}

func (Device nvmlDevice) GetSupportedGraphicsClocks(MemoryClockMHz int) (int, uint32, Return) {
	var Count, ClocksMHz uint32
	ret := nvmlDeviceGetSupportedGraphicsClocks(Device, uint32(MemoryClockMHz), &Count, &ClocksMHz)
	return int(Count), ClocksMHz, ret
}

// nvml.DeviceGetAutoBoostedClocksEnabled()
func (l *library) DeviceGetAutoBoostedClocksEnabled(Device Device) (EnableState, EnableState, Return) {
	return Device.GetAutoBoostedClocksEnabled()
}

func (Device nvmlDevice) GetAutoBoostedClocksEnabled() (EnableState, EnableState, Return) {
	var IsEnabled, DefaultIsEnabled EnableState
	ret := nvmlDeviceGetAutoBoostedClocksEnabled(Device, &IsEnabled, &DefaultIsEnabled)
	return IsEnabled, DefaultIsEnabled, ret
}

// nvml.DeviceSetAutoBoostedClocksEnabled()
func (l *library) DeviceSetAutoBoostedClocksEnabled(Device Device, Enabled EnableState) Return {
	return Device.SetAutoBoostedClocksEnabled(Enabled)
}

func (Device nvmlDevice) SetAutoBoostedClocksEnabled(Enabled EnableState) Return {
	return nvmlDeviceSetAutoBoostedClocksEnabled(Device, Enabled)
}

// nvml.DeviceSetDefaultAutoBoostedClocksEnabled()
func (l *library) DeviceSetDefaultAutoBoostedClocksEnabled(Device Device, Enabled EnableState, Flags uint32) Return {
	return Device.SetDefaultAutoBoostedClocksEnabled(Enabled, Flags)
}

func (Device nvmlDevice) SetDefaultAutoBoostedClocksEnabled(Enabled EnableState, Flags uint32) Return {
	return nvmlDeviceSetDefaultAutoBoostedClocksEnabled(Device, Enabled, Flags)
}

// nvml.DeviceGetFanSpeed()
func (l *library) DeviceGetFanSpeed(Device Device) (uint32, Return) {
	return Device.GetFanSpeed()
}

func (Device nvmlDevice) GetFanSpeed() (uint32, Return) {
	var Speed uint32
	ret := nvmlDeviceGetFanSpeed(Device, &Speed)
	return Speed, ret
}

// nvml.DeviceGetFanSpeed_v2()
func (l *library) DeviceGetFanSpeed_v2(Device Device, Fan int) (uint32, Return) {
	return Device.GetFanSpeed_v2(Fan)
}

func (Device nvmlDevice) GetFanSpeed_v2(Fan int) (uint32, Return) {
	var Speed uint32
	ret := nvmlDeviceGetFanSpeed_v2(Device, uint32(Fan), &Speed)
	return Speed, ret
}

// nvml.DeviceGetNumFans()
func (l *library) DeviceGetNumFans(Device Device) (int, Return) {
	return Device.GetNumFans()
}

func (Device nvmlDevice) GetNumFans() (int, Return) {
	var NumFans uint32
	ret := nvmlDeviceGetNumFans(Device, &NumFans)
	return int(NumFans), ret
}

// nvml.DeviceGetTemperature()
func (l *library) DeviceGetTemperature(Device Device, SensorType TemperatureSensors) (uint32, Return) {
	return Device.GetTemperature(SensorType)
}

func (Device nvmlDevice) GetTemperature(SensorType TemperatureSensors) (uint32, Return) {
	var Temp uint32
	ret := nvmlDeviceGetTemperature(Device, SensorType, &Temp)
	return Temp, ret
}

// nvml.DeviceGetTemperatureThreshold()
func (l *library) DeviceGetTemperatureThreshold(Device Device, ThresholdType TemperatureThresholds) (uint32, Return) {
	return Device.GetTemperatureThreshold(ThresholdType)
}

func (Device nvmlDevice) GetTemperatureThreshold(ThresholdType TemperatureThresholds) (uint32, Return) {
	var Temp uint32
	ret := nvmlDeviceGetTemperatureThreshold(Device, ThresholdType, &Temp)
	return Temp, ret
}

// nvml.DeviceSetTemperatureThreshold()
func (l *library) DeviceSetTemperatureThreshold(Device Device, ThresholdType TemperatureThresholds, Temp int) Return {
	return Device.SetTemperatureThreshold(ThresholdType, Temp)
}

func (Device nvmlDevice) SetTemperatureThreshold(ThresholdType TemperatureThresholds, Temp int) Return {
	t := int32(Temp)
	ret := nvmlDeviceSetTemperatureThreshold(Device, ThresholdType, &t)
	return ret
}

// nvml.DeviceGetPerformanceState()
func (l *library) DeviceGetPerformanceState(Device Device) (Pstates, Return) {
	return Device.GetPerformanceState()
}

func (Device nvmlDevice) GetPerformanceState() (Pstates, Return) {
	var PState Pstates
	ret := nvmlDeviceGetPerformanceState(Device, &PState)
	return PState, ret
}

// nvml.DeviceGetCurrentClocksThrottleReasons()
func (l *library) DeviceGetCurrentClocksThrottleReasons(Device Device) (uint64, Return) {
	return Device.GetCurrentClocksThrottleReasons()
}

func (Device nvmlDevice) GetCurrentClocksThrottleReasons() (uint64, Return) {
	var ClocksThrottleReasons uint64
	ret := nvmlDeviceGetCurrentClocksThrottleReasons(Device, &ClocksThrottleReasons)
	return ClocksThrottleReasons, ret
}

// nvml.DeviceGetSupportedClocksThrottleReasons()
func (l *library) DeviceGetSupportedClocksThrottleReasons(Device Device) (uint64, Return) {
	return Device.GetSupportedClocksThrottleReasons()
}

func (Device nvmlDevice) GetSupportedClocksThrottleReasons() (uint64, Return) {
	var SupportedClocksThrottleReasons uint64
	ret := nvmlDeviceGetSupportedClocksThrottleReasons(Device, &SupportedClocksThrottleReasons)
	return SupportedClocksThrottleReasons, ret
}

// nvml.DeviceGetPowerState()
func (l *library) DeviceGetPowerState(Device Device) (Pstates, Return) {
	return Device.GetPowerState()
}

func (Device nvmlDevice) GetPowerState() (Pstates, Return) {
	var PState Pstates
	ret := nvmlDeviceGetPowerState(Device, &PState)
	return PState, ret
}

// nvml.DeviceGetPowerManagementMode()
func (l *library) DeviceGetPowerManagementMode(Device Device) (EnableState, Return) {
	return Device.GetPowerManagementMode()
}

func (Device nvmlDevice) GetPowerManagementMode() (EnableState, Return) {
	var Mode EnableState
	ret := nvmlDeviceGetPowerManagementMode(Device, &Mode)
	return Mode, ret
}

// nvml.DeviceGetPowerManagementLimit()
func (l *library) DeviceGetPowerManagementLimit(Device Device) (uint32, Return) {
	return Device.GetPowerManagementLimit()
}

func (Device nvmlDevice) GetPowerManagementLimit() (uint32, Return) {
	var Limit uint32
	ret := nvmlDeviceGetPowerManagementLimit(Device, &Limit)
	return Limit, ret
}

// nvml.DeviceGetPowerManagementLimitConstraints()
func (l *library) DeviceGetPowerManagementLimitConstraints(Device Device) (uint32, uint32, Return) {
	return Device.GetPowerManagementLimitConstraints()
}

func (Device nvmlDevice) GetPowerManagementLimitConstraints() (uint32, uint32, Return) {
	var MinLimit, MaxLimit uint32
	ret := nvmlDeviceGetPowerManagementLimitConstraints(Device, &MinLimit, &MaxLimit)
	return MinLimit, MaxLimit, ret
}

// nvml.DeviceGetPowerManagementDefaultLimit()
func (l *library) DeviceGetPowerManagementDefaultLimit(Device Device) (uint32, Return) {
	return Device.GetPowerManagementDefaultLimit()
}

func (Device nvmlDevice) GetPowerManagementDefaultLimit() (uint32, Return) {
	var DefaultLimit uint32
	ret := nvmlDeviceGetPowerManagementDefaultLimit(Device, &DefaultLimit)
	return DefaultLimit, ret
}

// nvml.DeviceGetPowerUsage()
func (l *library) DeviceGetPowerUsage(Device Device) (uint32, Return) {
	return Device.GetPowerUsage()
}

func (Device nvmlDevice) GetPowerUsage() (uint32, Return) {
	var Power uint32
	ret := nvmlDeviceGetPowerUsage(Device, &Power)
	return Power, ret
}

// nvml.DeviceGetTotalEnergyConsumption()
func (l *library) DeviceGetTotalEnergyConsumption(Device Device) (uint64, Return) {
	return Device.GetTotalEnergyConsumption()
}

func (Device nvmlDevice) GetTotalEnergyConsumption() (uint64, Return) {
	var Energy uint64
	ret := nvmlDeviceGetTotalEnergyConsumption(Device, &Energy)
	return Energy, ret
}

// nvml.DeviceGetEnforcedPowerLimit()
func (l *library) DeviceGetEnforcedPowerLimit(Device Device) (uint32, Return) {
	return Device.GetEnforcedPowerLimit()
}

func (Device nvmlDevice) GetEnforcedPowerLimit() (uint32, Return) {
	var Limit uint32
	ret := nvmlDeviceGetEnforcedPowerLimit(Device, &Limit)
	return Limit, ret
}

// nvml.DeviceGetGpuOperationMode()
func (l *library) DeviceGetGpuOperationMode(Device Device) (GpuOperationMode, GpuOperationMode, Return) {
	return Device.GetGpuOperationMode()
}

func (Device nvmlDevice) GetGpuOperationMode() (GpuOperationMode, GpuOperationMode, Return) {
	var Current, Pending GpuOperationMode
	ret := nvmlDeviceGetGpuOperationMode(Device, &Current, &Pending)
	return Current, Pending, ret
}

// nvml.DeviceGetMemoryInfo()
func (l *library) DeviceGetMemoryInfo(Device Device) (Memory, Return) {
	return Device.GetMemoryInfo()
}

func (Device nvmlDevice) GetMemoryInfo() (Memory, Return) {
	var Memory Memory
	ret := nvmlDeviceGetMemoryInfo(Device, &Memory)
	return Memory, ret
}

// nvml.DeviceGetMemoryInfo_v2()
func (l *library) DeviceGetMemoryInfo_v2(Device Device) (Memory_v2, Return) {
	return Device.GetMemoryInfo_v2()
}

func (Device nvmlDevice) GetMemoryInfo_v2() (Memory_v2, Return) {
	var Memory Memory_v2
	Memory.Version = STRUCT_VERSION(Memory, 2)
	ret := nvmlDeviceGetMemoryInfo_v2(Device, &Memory)
	return Memory, ret
}

// nvml.DeviceGetComputeMode()
func (l *library) DeviceGetComputeMode(Device Device) (ComputeMode, Return) {
	return Device.GetComputeMode()
}

func (Device nvmlDevice) GetComputeMode() (ComputeMode, Return) {
	var Mode ComputeMode
	ret := nvmlDeviceGetComputeMode(Device, &Mode)
	return Mode, ret
}

// nvml.DeviceGetCudaComputeCapability()
func (l *library) DeviceGetCudaComputeCapability(Device Device) (int, int, Return) {
	return Device.GetCudaComputeCapability()
}

func (Device nvmlDevice) GetCudaComputeCapability() (int, int, Return) {
	var Major, Minor int32
	ret := nvmlDeviceGetCudaComputeCapability(Device, &Major, &Minor)
	return int(Major), int(Minor), ret
}

// nvml.DeviceGetEccMode()
func (l *library) DeviceGetEccMode(Device Device) (EnableState, EnableState, Return) {
	return Device.GetEccMode()
}

func (Device nvmlDevice) GetEccMode() (EnableState, EnableState, Return) {
	var Current, Pending EnableState
	ret := nvmlDeviceGetEccMode(Device, &Current, &Pending)
	return Current, Pending, ret
}

// nvml.DeviceGetBoardId()
func (l *library) DeviceGetBoardId(Device Device) (uint32, Return) {
	return Device.GetBoardId()
}

func (Device nvmlDevice) GetBoardId() (uint32, Return) {
	var BoardId uint32
	ret := nvmlDeviceGetBoardId(Device, &BoardId)
	return BoardId, ret
}

// nvml.DeviceGetMultiGpuBoard()
func (l *library) DeviceGetMultiGpuBoard(Device Device) (int, Return) {
	return Device.GetMultiGpuBoard()
}

func (Device nvmlDevice) GetMultiGpuBoard() (int, Return) {
	var MultiGpuBool uint32
	ret := nvmlDeviceGetMultiGpuBoard(Device, &MultiGpuBool)
	return int(MultiGpuBool), ret
}

// nvml.DeviceGetTotalEccErrors()
func (l *library) DeviceGetTotalEccErrors(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType) (uint64, Return) {
	return Device.GetTotalEccErrors(ErrorType, CounterType)
}

func (Device nvmlDevice) GetTotalEccErrors(ErrorType MemoryErrorType, CounterType EccCounterType) (uint64, Return) {
	var EccCounts uint64
	ret := nvmlDeviceGetTotalEccErrors(Device, ErrorType, CounterType, &EccCounts)
	return EccCounts, ret
}

// nvml.DeviceGetDetailedEccErrors()
func (l *library) DeviceGetDetailedEccErrors(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType) (EccErrorCounts, Return) {
	return Device.GetDetailedEccErrors(ErrorType, CounterType)
}

func (Device nvmlDevice) GetDetailedEccErrors(ErrorType MemoryErrorType, CounterType EccCounterType) (EccErrorCounts, Return) {
	var EccCounts EccErrorCounts
	ret := nvmlDeviceGetDetailedEccErrors(Device, ErrorType, CounterType, &EccCounts)
	return EccCounts, ret
}

// nvml.DeviceGetMemoryErrorCounter()
func (l *library) DeviceGetMemoryErrorCounter(Device Device, ErrorType MemoryErrorType, CounterType EccCounterType, LocationType MemoryLocation) (uint64, Return) {
	return Device.GetMemoryErrorCounter(ErrorType, CounterType, LocationType)
}

func (Device nvmlDevice) GetMemoryErrorCounter(ErrorType MemoryErrorType, CounterType EccCounterType, LocationType MemoryLocation) (uint64, Return) {
	var Count uint64
	ret := nvmlDeviceGetMemoryErrorCounter(Device, ErrorType, CounterType, LocationType, &Count)
	return Count, ret
}

// nvml.DeviceGetUtilizationRates()
func (l *library) DeviceGetUtilizationRates(Device Device) (Utilization, Return) {
	return Device.GetUtilizationRates()
}

func (Device nvmlDevice) GetUtilizationRates() (Utilization, Return) {
	var Utilization Utilization
	ret := nvmlDeviceGetUtilizationRates(Device, &Utilization)
	return Utilization, ret
}

// nvml.DeviceGetEncoderUtilization()
func (l *library) DeviceGetEncoderUtilization(Device Device) (uint32, uint32, Return) {
	return Device.GetEncoderUtilization()
}

func (Device nvmlDevice) GetEncoderUtilization() (uint32, uint32, Return) {
	var Utilization, SamplingPeriodUs uint32
	ret := nvmlDeviceGetEncoderUtilization(Device, &Utilization, &SamplingPeriodUs)
	return Utilization, SamplingPeriodUs, ret
}

// nvml.DeviceGetEncoderCapacity()
func (l *library) DeviceGetEncoderCapacity(Device Device, EncoderQueryType EncoderType) (int, Return) {
	return Device.GetEncoderCapacity(EncoderQueryType)
}

func (Device nvmlDevice) GetEncoderCapacity(EncoderQueryType EncoderType) (int, Return) {
	var EncoderCapacity uint32
	ret := nvmlDeviceGetEncoderCapacity(Device, EncoderQueryType, &EncoderCapacity)
	return int(EncoderCapacity), ret
}

// nvml.DeviceGetEncoderStats()
func (l *library) DeviceGetEncoderStats(Device Device) (int, uint32, uint32, Return) {
	return Device.GetEncoderStats()
}

func (Device nvmlDevice) GetEncoderStats() (int, uint32, uint32, Return) {
	var SessionCount, AverageFps, AverageLatency uint32
	ret := nvmlDeviceGetEncoderStats(Device, &SessionCount, &AverageFps, &AverageLatency)
	return int(SessionCount), AverageFps, AverageLatency, ret
}

// nvml.DeviceGetEncoderSessions()
func (l *library) DeviceGetEncoderSessions(Device Device) ([]EncoderSessionInfo, Return) {
	return Device.GetEncoderSessions()
}

func (Device nvmlDevice) GetEncoderSessions() ([]EncoderSessionInfo, Return) {
	var SessionCount uint32 = 1 // Will be reduced upon returning
	for {
		SessionInfos := make([]EncoderSessionInfo, SessionCount)
		ret := nvmlDeviceGetEncoderSessions(Device, &SessionCount, &SessionInfos[0])
		if ret == SUCCESS {
			return SessionInfos[:SessionCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		SessionCount *= 2
	}
}

// nvml.DeviceGetDecoderUtilization()
func (l *library) DeviceGetDecoderUtilization(Device Device) (uint32, uint32, Return) {
	return Device.GetDecoderUtilization()
}

func (Device nvmlDevice) GetDecoderUtilization() (uint32, uint32, Return) {
	var Utilization, SamplingPeriodUs uint32
	ret := nvmlDeviceGetDecoderUtilization(Device, &Utilization, &SamplingPeriodUs)
	return Utilization, SamplingPeriodUs, ret
}

// nvml.DeviceGetFBCStats()
func (l *library) DeviceGetFBCStats(Device Device) (FBCStats, Return) {
	return Device.GetFBCStats()
}

func (Device nvmlDevice) GetFBCStats() (FBCStats, Return) {
	var FbcStats FBCStats
	ret := nvmlDeviceGetFBCStats(Device, &FbcStats)
	return FbcStats, ret
}

// nvml.DeviceGetFBCSessions()
func (l *library) DeviceGetFBCSessions(Device Device) ([]FBCSessionInfo, Return) {
	return Device.GetFBCSessions()
}

func (Device nvmlDevice) GetFBCSessions() ([]FBCSessionInfo, Return) {
	var SessionCount uint32 = 1 // Will be reduced upon returning
	for {
		SessionInfo := make([]FBCSessionInfo, SessionCount)
		ret := nvmlDeviceGetFBCSessions(Device, &SessionCount, &SessionInfo[0])
		if ret == SUCCESS {
			return SessionInfo[:SessionCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		SessionCount *= 2
	}
}

// nvml.DeviceGetDriverModel()
func (l *library) DeviceGetDriverModel(Device Device) (DriverModel, DriverModel, Return) {
	return Device.GetDriverModel()
}

func (Device nvmlDevice) GetDriverModel() (DriverModel, DriverModel, Return) {
	var Current, Pending DriverModel
	ret := nvmlDeviceGetDriverModel(Device, &Current, &Pending)
	return Current, Pending, ret
}

// nvml.DeviceGetVbiosVersion()
func (l *library) DeviceGetVbiosVersion(Device Device) (string, Return) {
	return Device.GetVbiosVersion()
}

func (Device nvmlDevice) GetVbiosVersion() (string, Return) {
	Version := make([]byte, DEVICE_VBIOS_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetVbiosVersion(Device, &Version[0], DEVICE_VBIOS_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.DeviceGetBridgeChipInfo()
func (l *library) DeviceGetBridgeChipInfo(Device Device) (BridgeChipHierarchy, Return) {
	return Device.GetBridgeChipInfo()
}

func (Device nvmlDevice) GetBridgeChipInfo() (BridgeChipHierarchy, Return) {
	var BridgeHierarchy BridgeChipHierarchy
	ret := nvmlDeviceGetBridgeChipInfo(Device, &BridgeHierarchy)
	return BridgeHierarchy, ret
}

// nvml.DeviceGetComputeRunningProcesses()
func deviceGetComputeRunningProcesses_v1(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v1, InfoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v1(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v1Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetComputeRunningProcesses_v2(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v2, InfoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v2(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v2Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetComputeRunningProcesses_v3(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo, InfoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v3(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return Infos[:InfoCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func (l *library) DeviceGetComputeRunningProcesses(Device Device) ([]ProcessInfo, Return) {
	return Device.GetComputeRunningProcesses()
}

func (Device nvmlDevice) GetComputeRunningProcesses() ([]ProcessInfo, Return) {
	return deviceGetComputeRunningProcesses(Device)
}

// nvml.DeviceGetGraphicsRunningProcesses()
func deviceGetGraphicsRunningProcesses_v1(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v1, InfoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v1(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v1Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetGraphicsRunningProcesses_v2(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v2, InfoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v2(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v2Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetGraphicsRunningProcesses_v3(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo, InfoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v3(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return Infos[:InfoCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func (l *library) DeviceGetGraphicsRunningProcesses(Device Device) ([]ProcessInfo, Return) {
	return Device.GetGraphicsRunningProcesses()
}

func (Device nvmlDevice) GetGraphicsRunningProcesses() ([]ProcessInfo, Return) {
	return deviceGetGraphicsRunningProcesses(Device)
}

// nvml.DeviceGetMPSComputeRunningProcesses()
func deviceGetMPSComputeRunningProcesses_v1(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v1, InfoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v1(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v1Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetMPSComputeRunningProcesses_v2(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo_v2, InfoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v2(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return ProcessInfo_v2Slice(Infos[:InfoCount]).ToProcessInfoSlice(), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func deviceGetMPSComputeRunningProcesses_v3(Device nvmlDevice) ([]ProcessInfo, Return) {
	var InfoCount uint32 = 1 // Will be reduced upon returning
	for {
		Infos := make([]ProcessInfo, InfoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v3(Device, &InfoCount, &Infos[0])
		if ret == SUCCESS {
			return Infos[:InfoCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}
}

func (l *library) DeviceGetMPSComputeRunningProcesses(Device Device) ([]ProcessInfo, Return) {
	return Device.GetMPSComputeRunningProcesses()
}

func (Device nvmlDevice) GetMPSComputeRunningProcesses() ([]ProcessInfo, Return) {
	return deviceGetMPSComputeRunningProcesses(Device)
}

// nvml.DeviceOnSameBoard()
func (l *library) DeviceOnSameBoard(Device1 Device, Device2 Device) (int, Return) {
	return Device1.OnSameBoard(Device2)
}

func (Device1 nvmlDevice) OnSameBoard(Device2 Device) (int, Return) {
	var OnSameBoard int32
	ret := nvmlDeviceOnSameBoard(Device1, Device2.(nvmlDevice), &OnSameBoard)
	return int(OnSameBoard), ret
}

// nvml.DeviceGetAPIRestriction()
func (l *library) DeviceGetAPIRestriction(Device Device, ApiType RestrictedAPI) (EnableState, Return) {
	return Device.GetAPIRestriction(ApiType)
}

func (Device nvmlDevice) GetAPIRestriction(ApiType RestrictedAPI) (EnableState, Return) {
	var IsRestricted EnableState
	ret := nvmlDeviceGetAPIRestriction(Device, ApiType, &IsRestricted)
	return IsRestricted, ret
}

// nvml.DeviceGetSamples()
func (l *library) DeviceGetSamples(Device Device, _type SamplingType, LastSeenTimeStamp uint64) (ValueType, []Sample, Return) {
	return Device.GetSamples(_type, LastSeenTimeStamp)
}

func (Device nvmlDevice) GetSamples(_type SamplingType, LastSeenTimeStamp uint64) (ValueType, []Sample, Return) {
	var SampleValType ValueType
	var SampleCount uint32
	ret := nvmlDeviceGetSamples(Device, _type, LastSeenTimeStamp, &SampleValType, &SampleCount, nil)
	if ret != SUCCESS {
		return SampleValType, nil, ret
	}
	if SampleCount == 0 {
		return SampleValType, []Sample{}, ret
	}
	Samples := make([]Sample, SampleCount)
	ret = nvmlDeviceGetSamples(Device, _type, LastSeenTimeStamp, &SampleValType, &SampleCount, &Samples[0])
	return SampleValType, Samples, ret
}

// nvml.DeviceGetBAR1MemoryInfo()
func (l *library) DeviceGetBAR1MemoryInfo(Device Device) (BAR1Memory, Return) {
	return Device.GetBAR1MemoryInfo()
}

func (Device nvmlDevice) GetBAR1MemoryInfo() (BAR1Memory, Return) {
	var Bar1Memory BAR1Memory
	ret := nvmlDeviceGetBAR1MemoryInfo(Device, &Bar1Memory)
	return Bar1Memory, ret
}

// nvml.DeviceGetViolationStatus()
func (l *library) DeviceGetViolationStatus(Device Device, PerfPolicyType PerfPolicyType) (ViolationTime, Return) {
	return Device.GetViolationStatus(PerfPolicyType)
}

func (Device nvmlDevice) GetViolationStatus(PerfPolicyType PerfPolicyType) (ViolationTime, Return) {
	var ViolTime ViolationTime
	ret := nvmlDeviceGetViolationStatus(Device, PerfPolicyType, &ViolTime)
	return ViolTime, ret
}

// nvml.DeviceGetIrqNum()
func (l *library) DeviceGetIrqNum(Device Device) (int, Return) {
	return Device.GetIrqNum()
}

func (Device nvmlDevice) GetIrqNum() (int, Return) {
	var IrqNum uint32
	ret := nvmlDeviceGetIrqNum(Device, &IrqNum)
	return int(IrqNum), ret
}

// nvml.DeviceGetNumGpuCores()
func (l *library) DeviceGetNumGpuCores(Device Device) (int, Return) {
	return Device.GetNumGpuCores()
}

func (Device nvmlDevice) GetNumGpuCores() (int, Return) {
	var NumCores uint32
	ret := nvmlDeviceGetNumGpuCores(Device, &NumCores)
	return int(NumCores), ret
}

// nvml.DeviceGetPowerSource()
func (l *library) DeviceGetPowerSource(Device Device) (PowerSource, Return) {
	return Device.GetPowerSource()
}

func (Device nvmlDevice) GetPowerSource() (PowerSource, Return) {
	var PowerSource PowerSource
	ret := nvmlDeviceGetPowerSource(Device, &PowerSource)
	return PowerSource, ret
}

// nvml.DeviceGetMemoryBusWidth()
func (l *library) DeviceGetMemoryBusWidth(Device Device) (uint32, Return) {
	return Device.GetMemoryBusWidth()
}

func (Device nvmlDevice) GetMemoryBusWidth() (uint32, Return) {
	var BusWidth uint32
	ret := nvmlDeviceGetMemoryBusWidth(Device, &BusWidth)
	return BusWidth, ret
}

// nvml.DeviceGetPcieLinkMaxSpeed()
func (l *library) DeviceGetPcieLinkMaxSpeed(Device Device) (uint32, Return) {
	return Device.GetPcieLinkMaxSpeed()
}

func (Device nvmlDevice) GetPcieLinkMaxSpeed() (uint32, Return) {
	var MaxSpeed uint32
	ret := nvmlDeviceGetPcieLinkMaxSpeed(Device, &MaxSpeed)
	return MaxSpeed, ret
}

// nvml.DeviceGetAdaptiveClockInfoStatus()
func (l *library) DeviceGetAdaptiveClockInfoStatus(Device Device) (uint32, Return) {
	return Device.GetAdaptiveClockInfoStatus()
}

func (Device nvmlDevice) GetAdaptiveClockInfoStatus() (uint32, Return) {
	var AdaptiveClockStatus uint32
	ret := nvmlDeviceGetAdaptiveClockInfoStatus(Device, &AdaptiveClockStatus)
	return AdaptiveClockStatus, ret
}

// nvml.DeviceGetAccountingMode()
func (l *library) DeviceGetAccountingMode(Device Device) (EnableState, Return) {
	return Device.GetAccountingMode()
}

func (Device nvmlDevice) GetAccountingMode() (EnableState, Return) {
	var Mode EnableState
	ret := nvmlDeviceGetAccountingMode(Device, &Mode)
	return Mode, ret
}

// nvml.DeviceGetAccountingStats()
func (l *library) DeviceGetAccountingStats(Device Device, Pid uint32) (AccountingStats, Return) {
	return Device.GetAccountingStats(Pid)
}

func (Device nvmlDevice) GetAccountingStats(Pid uint32) (AccountingStats, Return) {
	var Stats AccountingStats
	ret := nvmlDeviceGetAccountingStats(Device, Pid, &Stats)
	return Stats, ret
}

// nvml.DeviceGetAccountingPids()
func (l *library) DeviceGetAccountingPids(Device Device) ([]int, Return) {
	return Device.GetAccountingPids()
}

func (Device nvmlDevice) GetAccountingPids() ([]int, Return) {
	var Count uint32 = 1 // Will be reduced upon returning
	for {
		Pids := make([]uint32, Count)
		ret := nvmlDeviceGetAccountingPids(Device, &Count, &Pids[0])
		if ret == SUCCESS {
			return uint32SliceToIntSlice(Pids[:Count]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		Count *= 2
	}
}

// nvml.DeviceGetAccountingBufferSize()
func (l *library) DeviceGetAccountingBufferSize(Device Device) (int, Return) {
	return Device.GetAccountingBufferSize()
}

func (Device nvmlDevice) GetAccountingBufferSize() (int, Return) {
	var BufferSize uint32
	ret := nvmlDeviceGetAccountingBufferSize(Device, &BufferSize)
	return int(BufferSize), ret
}

// nvml.DeviceGetRetiredPages()
func (l *library) DeviceGetRetiredPages(Device Device, Cause PageRetirementCause) ([]uint64, Return) {
	return Device.GetRetiredPages(Cause)
}

func (Device nvmlDevice) GetRetiredPages(Cause PageRetirementCause) ([]uint64, Return) {
	var PageCount uint32 = 1 // Will be reduced upon returning
	for {
		Addresses := make([]uint64, PageCount)
		ret := nvmlDeviceGetRetiredPages(Device, Cause, &PageCount, &Addresses[0])
		if ret == SUCCESS {
			return Addresses[:PageCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		PageCount *= 2
	}
}

// nvml.DeviceGetRetiredPages_v2()
func (l *library) DeviceGetRetiredPages_v2(Device Device, Cause PageRetirementCause) ([]uint64, []uint64, Return) {
	return Device.GetRetiredPages_v2(Cause)
}

func (Device nvmlDevice) GetRetiredPages_v2(Cause PageRetirementCause) ([]uint64, []uint64, Return) {
	var PageCount uint32 = 1 // Will be reduced upon returning
	for {
		Addresses := make([]uint64, PageCount)
		Timestamps := make([]uint64, PageCount)
		ret := nvmlDeviceGetRetiredPages_v2(Device, Cause, &PageCount, &Addresses[0], &Timestamps[0])
		if ret == SUCCESS {
			return Addresses[:PageCount], Timestamps[:PageCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, nil, ret
		}
		PageCount *= 2
	}
}

// nvml.DeviceGetRetiredPagesPendingStatus()
func (l *library) DeviceGetRetiredPagesPendingStatus(Device Device) (EnableState, Return) {
	return Device.GetRetiredPagesPendingStatus()
}

func (Device nvmlDevice) GetRetiredPagesPendingStatus() (EnableState, Return) {
	var IsPending EnableState
	ret := nvmlDeviceGetRetiredPagesPendingStatus(Device, &IsPending)
	return IsPending, ret
}

// nvml.DeviceSetPersistenceMode()
func (l *library) DeviceSetPersistenceMode(Device Device, Mode EnableState) Return {
	return Device.SetPersistenceMode(Mode)
}

func (Device nvmlDevice) SetPersistenceMode(Mode EnableState) Return {
	return nvmlDeviceSetPersistenceMode(Device, Mode)
}

// nvml.DeviceSetComputeMode()
func (l *library) DeviceSetComputeMode(Device Device, Mode ComputeMode) Return {
	return Device.SetComputeMode(Mode)
}

func (Device nvmlDevice) SetComputeMode(Mode ComputeMode) Return {
	return nvmlDeviceSetComputeMode(Device, Mode)
}

// nvml.DeviceSetEccMode()
func (l *library) DeviceSetEccMode(Device Device, Ecc EnableState) Return {
	return Device.SetEccMode(Ecc)
}

func (Device nvmlDevice) SetEccMode(Ecc EnableState) Return {
	return nvmlDeviceSetEccMode(Device, Ecc)
}

// nvml.DeviceClearEccErrorCounts()
func (l *library) DeviceClearEccErrorCounts(Device Device, CounterType EccCounterType) Return {
	return Device.ClearEccErrorCounts(CounterType)
}

func (Device nvmlDevice) ClearEccErrorCounts(CounterType EccCounterType) Return {
	return nvmlDeviceClearEccErrorCounts(Device, CounterType)
}

// nvml.DeviceSetDriverModel()
func (l *library) DeviceSetDriverModel(Device Device, DriverModel DriverModel, Flags uint32) Return {
	return Device.SetDriverModel(DriverModel, Flags)
}

func (Device nvmlDevice) SetDriverModel(DriverModel DriverModel, Flags uint32) Return {
	return nvmlDeviceSetDriverModel(Device, DriverModel, Flags)
}

// nvml.DeviceSetGpuLockedClocks()
func (l *library) DeviceSetGpuLockedClocks(Device Device, MinGpuClockMHz uint32, MaxGpuClockMHz uint32) Return {
	return Device.SetGpuLockedClocks(MinGpuClockMHz, MaxGpuClockMHz)
}

func (Device nvmlDevice) SetGpuLockedClocks(MinGpuClockMHz uint32, MaxGpuClockMHz uint32) Return {
	return nvmlDeviceSetGpuLockedClocks(Device, MinGpuClockMHz, MaxGpuClockMHz)
}

// nvml.DeviceResetGpuLockedClocks()
func (l *library) DeviceResetGpuLockedClocks(Device Device) Return {
	return Device.ResetGpuLockedClocks()
}

func (Device nvmlDevice) ResetGpuLockedClocks() Return {
	return nvmlDeviceResetGpuLockedClocks(Device)
}

// nvmlDeviceSetMemoryLockedClocks()
func (l *library) DeviceSetMemoryLockedClocks(Device Device, MinMemClockMHz uint32, MaxMemClockMHz uint32) Return {
	return Device.SetMemoryLockedClocks(MinMemClockMHz, MaxMemClockMHz)
}

func (Device nvmlDevice) SetMemoryLockedClocks(NinMemClockMHz uint32, MaxMemClockMHz uint32) Return {
	return nvmlDeviceSetMemoryLockedClocks(Device, NinMemClockMHz, MaxMemClockMHz)
}

// nvmlDeviceResetMemoryLockedClocks()
func (l *library) DeviceResetMemoryLockedClocks(Device Device) Return {
	return Device.ResetMemoryLockedClocks()
}

func (Device nvmlDevice) ResetMemoryLockedClocks() Return {
	return nvmlDeviceResetMemoryLockedClocks(Device)
}

// nvml.DeviceGetClkMonStatus()
func (l *library) DeviceGetClkMonStatus(Device Device) (ClkMonStatus, Return) {
	return Device.GetClkMonStatus()
}

func (Device nvmlDevice) GetClkMonStatus() (ClkMonStatus, Return) {
	var Status ClkMonStatus
	ret := nvmlDeviceGetClkMonStatus(Device, &Status)
	return Status, ret
}

// nvml.DeviceSetApplicationsClocks()
func (l *library) DeviceSetApplicationsClocks(Device Device, MemClockMHz uint32, GraphicsClockMHz uint32) Return {
	return Device.SetApplicationsClocks(MemClockMHz, GraphicsClockMHz)
}

func (Device nvmlDevice) SetApplicationsClocks(MemClockMHz uint32, GraphicsClockMHz uint32) Return {
	return nvmlDeviceSetApplicationsClocks(Device, MemClockMHz, GraphicsClockMHz)
}

// nvml.DeviceSetPowerManagementLimit()
func (l *library) DeviceSetPowerManagementLimit(Device Device, Limit uint32) Return {
	return Device.SetPowerManagementLimit(Limit)
}

func (Device nvmlDevice) SetPowerManagementLimit(Limit uint32) Return {
	return nvmlDeviceSetPowerManagementLimit(Device, Limit)
}

// nvml.DeviceSetGpuOperationMode()
func (l *library) DeviceSetGpuOperationMode(Device Device, Mode GpuOperationMode) Return {
	return Device.SetGpuOperationMode(Mode)
}

func (Device nvmlDevice) SetGpuOperationMode(Mode GpuOperationMode) Return {
	return nvmlDeviceSetGpuOperationMode(Device, Mode)
}

// nvml.DeviceSetAPIRestriction()
func (l *library) DeviceSetAPIRestriction(Device Device, ApiType RestrictedAPI, IsRestricted EnableState) Return {
	return Device.SetAPIRestriction(ApiType, IsRestricted)
}

func (Device nvmlDevice) SetAPIRestriction(ApiType RestrictedAPI, IsRestricted EnableState) Return {
	return nvmlDeviceSetAPIRestriction(Device, ApiType, IsRestricted)
}

// nvml.DeviceSetAccountingMode()
func (l *library) DeviceSetAccountingMode(Device Device, Mode EnableState) Return {
	return Device.SetAccountingMode(Mode)
}

func (Device nvmlDevice) SetAccountingMode(Mode EnableState) Return {
	return nvmlDeviceSetAccountingMode(Device, Mode)
}

// nvml.DeviceClearAccountingPids()
func (l *library) DeviceClearAccountingPids(Device Device) Return {
	return Device.ClearAccountingPids()
}

func (Device nvmlDevice) ClearAccountingPids() Return {
	return nvmlDeviceClearAccountingPids(Device)
}

// nvml.DeviceGetNvLinkState()
func (l *library) DeviceGetNvLinkState(Device Device, Link int) (EnableState, Return) {
	return Device.GetNvLinkState(Link)
}

func (Device nvmlDevice) GetNvLinkState(Link int) (EnableState, Return) {
	var IsActive EnableState
	ret := nvmlDeviceGetNvLinkState(Device, uint32(Link), &IsActive)
	return IsActive, ret
}

// nvml.DeviceGetNvLinkVersion()
func (l *library) DeviceGetNvLinkVersion(Device Device, Link int) (uint32, Return) {
	return Device.GetNvLinkVersion(Link)
}

func (Device nvmlDevice) GetNvLinkVersion(Link int) (uint32, Return) {
	var Version uint32
	ret := nvmlDeviceGetNvLinkVersion(Device, uint32(Link), &Version)
	return Version, ret
}

// nvml.DeviceGetNvLinkCapability()
func (l *library) DeviceGetNvLinkCapability(Device Device, Link int, Capability NvLinkCapability) (uint32, Return) {
	return Device.GetNvLinkCapability(Link, Capability)
}

func (Device nvmlDevice) GetNvLinkCapability(Link int, Capability NvLinkCapability) (uint32, Return) {
	var CapResult uint32
	ret := nvmlDeviceGetNvLinkCapability(Device, uint32(Link), Capability, &CapResult)
	return CapResult, ret
}

// nvml.DeviceGetNvLinkRemotePciInfo()
func (l *library) DeviceGetNvLinkRemotePciInfo(Device Device, Link int) (PciInfo, Return) {
	return Device.GetNvLinkRemotePciInfo(Link)
}

func (Device nvmlDevice) GetNvLinkRemotePciInfo(Link int) (PciInfo, Return) {
	var Pci PciInfo
	ret := nvmlDeviceGetNvLinkRemotePciInfo(Device, uint32(Link), &Pci)
	return Pci, ret
}

// nvml.DeviceGetNvLinkErrorCounter()
func (l *library) DeviceGetNvLinkErrorCounter(Device Device, Link int, Counter NvLinkErrorCounter) (uint64, Return) {
	return Device.GetNvLinkErrorCounter(Link, Counter)
}

func (Device nvmlDevice) GetNvLinkErrorCounter(Link int, Counter NvLinkErrorCounter) (uint64, Return) {
	var CounterValue uint64
	ret := nvmlDeviceGetNvLinkErrorCounter(Device, uint32(Link), Counter, &CounterValue)
	return CounterValue, ret
}

// nvml.DeviceResetNvLinkErrorCounters()
func (l *library) DeviceResetNvLinkErrorCounters(Device Device, Link int) Return {
	return Device.ResetNvLinkErrorCounters(Link)
}

func (Device nvmlDevice) ResetNvLinkErrorCounters(Link int) Return {
	return nvmlDeviceResetNvLinkErrorCounters(Device, uint32(Link))
}

// nvml.DeviceSetNvLinkUtilizationControl()
func (l *library) DeviceSetNvLinkUtilizationControl(Device Device, Link int, Counter int, Control *NvLinkUtilizationControl, Reset bool) Return {
	return Device.SetNvLinkUtilizationControl(Link, Counter, Control, Reset)
}

func (Device nvmlDevice) SetNvLinkUtilizationControl(Link int, Counter int, Control *NvLinkUtilizationControl, Reset bool) Return {
	reset := uint32(0)
	if Reset {
		reset = 1
	}
	return nvmlDeviceSetNvLinkUtilizationControl(Device, uint32(Link), uint32(Counter), Control, reset)
}

// nvml.DeviceGetNvLinkUtilizationControl()
func (l *library) DeviceGetNvLinkUtilizationControl(Device Device, Link int, Counter int) (NvLinkUtilizationControl, Return) {
	return Device.GetNvLinkUtilizationControl(Link, Counter)
}

func (Device nvmlDevice) GetNvLinkUtilizationControl(Link int, Counter int) (NvLinkUtilizationControl, Return) {
	var Control NvLinkUtilizationControl
	ret := nvmlDeviceGetNvLinkUtilizationControl(Device, uint32(Link), uint32(Counter), &Control)
	return Control, ret
}

// nvml.DeviceGetNvLinkUtilizationCounter()
func (l *library) DeviceGetNvLinkUtilizationCounter(Device Device, Link int, Counter int) (uint64, uint64, Return) {
	return Device.GetNvLinkUtilizationCounter(Link, Counter)
}

func (Device nvmlDevice) GetNvLinkUtilizationCounter(Link int, Counter int) (uint64, uint64, Return) {
	var Rxcounter, Txcounter uint64
	ret := nvmlDeviceGetNvLinkUtilizationCounter(Device, uint32(Link), uint32(Counter), &Rxcounter, &Txcounter)
	return Rxcounter, Txcounter, ret
}

// nvml.DeviceFreezeNvLinkUtilizationCounter()
func (l *library) DeviceFreezeNvLinkUtilizationCounter(Device Device, Link int, Counter int, Freeze EnableState) Return {
	return Device.FreezeNvLinkUtilizationCounter(Link, Counter, Freeze)
}

func (Device nvmlDevice) FreezeNvLinkUtilizationCounter(Link int, Counter int, Freeze EnableState) Return {
	return nvmlDeviceFreezeNvLinkUtilizationCounter(Device, uint32(Link), uint32(Counter), Freeze)
}

// nvml.DeviceResetNvLinkUtilizationCounter()
func (l *library) DeviceResetNvLinkUtilizationCounter(Device Device, Link int, Counter int) Return {
	return Device.ResetNvLinkUtilizationCounter(Link, Counter)
}

func (Device nvmlDevice) ResetNvLinkUtilizationCounter(Link int, Counter int) Return {
	return nvmlDeviceResetNvLinkUtilizationCounter(Device, uint32(Link), uint32(Counter))
}

// nvml.DeviceGetNvLinkRemoteDeviceType()
func (l *library) DeviceGetNvLinkRemoteDeviceType(Device Device, Link int) (IntNvLinkDeviceType, Return) {
	return Device.GetNvLinkRemoteDeviceType(Link)
}

func (Device nvmlDevice) GetNvLinkRemoteDeviceType(Link int) (IntNvLinkDeviceType, Return) {
	var NvLinkDeviceType IntNvLinkDeviceType
	ret := nvmlDeviceGetNvLinkRemoteDeviceType(Device, uint32(Link), &NvLinkDeviceType)
	return NvLinkDeviceType, ret
}

// nvml.DeviceRegisterEvents()
func (l *library) DeviceRegisterEvents(Device Device, EventTypes uint64, Set EventSet) Return {
	return Device.RegisterEvents(EventTypes, Set)
}

func (Device nvmlDevice) RegisterEvents(EventTypes uint64, Set EventSet) Return {
	return nvmlDeviceRegisterEvents(Device, EventTypes, Set.(nvmlEventSet))
}

// nvmlDeviceGetSupportedEventTypes()
func (l *library) DeviceGetSupportedEventTypes(Device Device) (uint64, Return) {
	return Device.GetSupportedEventTypes()
}

func (Device nvmlDevice) GetSupportedEventTypes() (uint64, Return) {
	var EventTypes uint64
	ret := nvmlDeviceGetSupportedEventTypes(Device, &EventTypes)
	return EventTypes, ret
}

// nvml.DeviceModifyDrainState()
func (l *library) DeviceModifyDrainState(PciInfo *PciInfo, NewState EnableState) Return {
	return nvmlDeviceModifyDrainState(PciInfo, NewState)
}

// nvml.DeviceQueryDrainState()
func (l *library) DeviceQueryDrainState(PciInfo *PciInfo) (EnableState, Return) {
	var CurrentState EnableState
	ret := nvmlDeviceQueryDrainState(PciInfo, &CurrentState)
	return CurrentState, ret
}

// nvml.DeviceRemoveGpu()
func (l *library) DeviceRemoveGpu(PciInfo *PciInfo) Return {
	return nvmlDeviceRemoveGpu(PciInfo)
}

// nvml.DeviceRemoveGpu_v2()
func (l *library) DeviceRemoveGpu_v2(PciInfo *PciInfo, GpuState DetachGpuState, LinkState PcieLinkState) Return {
	return nvmlDeviceRemoveGpu_v2(PciInfo, GpuState, LinkState)
}

// nvml.DeviceDiscoverGpus()
func (l *library) DeviceDiscoverGpus() (PciInfo, Return) {
	var PciInfo PciInfo
	ret := nvmlDeviceDiscoverGpus(&PciInfo)
	return PciInfo, ret
}

// nvml.DeviceGetFieldValues()
func (l *library) DeviceGetFieldValues(Device Device, Values []FieldValue) Return {
	return Device.GetFieldValues(Values)
}

func (Device nvmlDevice) GetFieldValues(Values []FieldValue) Return {
	ValuesCount := len(Values)
	return nvmlDeviceGetFieldValues(Device, int32(ValuesCount), &Values[0])
}

// nvml.DeviceGetVirtualizationMode()
func (l *library) DeviceGetVirtualizationMode(Device Device) (GpuVirtualizationMode, Return) {
	return Device.GetVirtualizationMode()
}

func (Device nvmlDevice) GetVirtualizationMode() (GpuVirtualizationMode, Return) {
	var PVirtualMode GpuVirtualizationMode
	ret := nvmlDeviceGetVirtualizationMode(Device, &PVirtualMode)
	return PVirtualMode, ret
}

// nvml.DeviceGetHostVgpuMode()
func (l *library) DeviceGetHostVgpuMode(Device Device) (HostVgpuMode, Return) {
	return Device.GetHostVgpuMode()
}

func (Device nvmlDevice) GetHostVgpuMode() (HostVgpuMode, Return) {
	var PHostVgpuMode HostVgpuMode
	ret := nvmlDeviceGetHostVgpuMode(Device, &PHostVgpuMode)
	return PHostVgpuMode, ret
}

// nvml.DeviceSetVirtualizationMode()
func (l *library) DeviceSetVirtualizationMode(Device Device, VirtualMode GpuVirtualizationMode) Return {
	return Device.SetVirtualizationMode(VirtualMode)
}

func (Device nvmlDevice) SetVirtualizationMode(VirtualMode GpuVirtualizationMode) Return {
	return nvmlDeviceSetVirtualizationMode(Device, VirtualMode)
}

// nvml.DeviceGetGridLicensableFeatures()
func (l *library) DeviceGetGridLicensableFeatures(Device Device) (GridLicensableFeatures, Return) {
	return Device.GetGridLicensableFeatures()
}

func (Device nvmlDevice) GetGridLicensableFeatures() (GridLicensableFeatures, Return) {
	var PGridLicensableFeatures GridLicensableFeatures
	ret := nvmlDeviceGetGridLicensableFeatures(Device, &PGridLicensableFeatures)
	return PGridLicensableFeatures, ret
}

// nvml.DeviceGetProcessUtilization()
func (l *library) DeviceGetProcessUtilization(Device Device, LastSeenTimeStamp uint64) ([]ProcessUtilizationSample, Return) {
	return Device.GetProcessUtilization(LastSeenTimeStamp)
}

func (Device nvmlDevice) GetProcessUtilization(LastSeenTimeStamp uint64) ([]ProcessUtilizationSample, Return) {
	var ProcessSamplesCount uint32
	ret := nvmlDeviceGetProcessUtilization(Device, nil, &ProcessSamplesCount, LastSeenTimeStamp)
	if ret != ERROR_INSUFFICIENT_SIZE {
		return nil, ret
	}
	if ProcessSamplesCount == 0 {
		return []ProcessUtilizationSample{}, ret
	}
	Utilization := make([]ProcessUtilizationSample, ProcessSamplesCount)
	ret = nvmlDeviceGetProcessUtilization(Device, &Utilization[0], &ProcessSamplesCount, LastSeenTimeStamp)
	return Utilization[:ProcessSamplesCount], ret
}

// nvml.DeviceGetSupportedVgpus()
func (l *library) DeviceGetSupportedVgpus(Device Device) ([]VgpuTypeId, Return) {
	return Device.GetSupportedVgpus()
}

func (Device nvmlDevice) GetSupportedVgpus() ([]VgpuTypeId, Return) {
	var VgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		VgpuTypeIds := make([]nvmlVgpuTypeId, VgpuCount)
		ret := nvmlDeviceGetSupportedVgpus(Device, &VgpuCount, &VgpuTypeIds[0])
		if ret == SUCCESS {
			return convertSlice[nvmlVgpuTypeId, VgpuTypeId](VgpuTypeIds[:VgpuCount]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		VgpuCount *= 2
	}
}

// nvml.DeviceGetCreatableVgpus()
func (l *library) DeviceGetCreatableVgpus(Device Device) ([]VgpuTypeId, Return) {
	return Device.GetCreatableVgpus()
}

func (Device nvmlDevice) GetCreatableVgpus() ([]VgpuTypeId, Return) {
	var VgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		VgpuTypeIds := make([]nvmlVgpuTypeId, VgpuCount)
		ret := nvmlDeviceGetCreatableVgpus(Device, &VgpuCount, &VgpuTypeIds[0])
		if ret == SUCCESS {
			return convertSlice[nvmlVgpuTypeId, VgpuTypeId](VgpuTypeIds[:VgpuCount]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		VgpuCount *= 2
	}
}

// nvml.DeviceGetActiveVgpus()
func (l *library) DeviceGetActiveVgpus(Device Device) ([]VgpuInstance, Return) {
	return Device.GetActiveVgpus()
}

func (Device nvmlDevice) GetActiveVgpus() ([]VgpuInstance, Return) {
	var VgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		VgpuInstances := make([]nvmlVgpuInstance, VgpuCount)
		ret := nvmlDeviceGetActiveVgpus(Device, &VgpuCount, &VgpuInstances[0])
		if ret == SUCCESS {
			return convertSlice[nvmlVgpuInstance, VgpuInstance](VgpuInstances[:VgpuCount]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		VgpuCount *= 2
	}
}

// nvml.DeviceGetVgpuMetadata()
func (l *library) DeviceGetVgpuMetadata(Device Device) (VgpuPgpuMetadata, Return) {
	return Device.GetVgpuMetadata()
}

func (Device nvmlDevice) GetVgpuMetadata() (VgpuPgpuMetadata, Return) {
	var VgpuPgpuMetadata VgpuPgpuMetadata
	OpaqueDataSize := unsafe.Sizeof(VgpuPgpuMetadata.nvmlVgpuPgpuMetadata.OpaqueData)
	VgpuPgpuMetadataSize := unsafe.Sizeof(VgpuPgpuMetadata.nvmlVgpuPgpuMetadata) - OpaqueDataSize
	for {
		BufferSize := uint32(VgpuPgpuMetadataSize + OpaqueDataSize)
		Buffer := make([]byte, BufferSize)
		nvmlVgpuPgpuMetadataPtr := (*nvmlVgpuPgpuMetadata)(unsafe.Pointer(&Buffer[0]))
		ret := nvmlDeviceGetVgpuMetadata(Device, nvmlVgpuPgpuMetadataPtr, &BufferSize)
		if ret == SUCCESS {
			VgpuPgpuMetadata.nvmlVgpuPgpuMetadata = *nvmlVgpuPgpuMetadataPtr
			VgpuPgpuMetadata.OpaqueData = Buffer[VgpuPgpuMetadataSize:BufferSize]
			return VgpuPgpuMetadata, ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return VgpuPgpuMetadata, ret
		}
		OpaqueDataSize = 2 * OpaqueDataSize
	}
}

// nvml.DeviceGetPgpuMetadataString()
func (l *library) DeviceGetPgpuMetadataString(Device Device) (string, Return) {
	return Device.GetPgpuMetadataString()
}

func (Device nvmlDevice) GetPgpuMetadataString() (string, Return) {
	var BufferSize uint32 = 1 // Will be reduced upon returning
	for {
		PgpuMetadata := make([]byte, BufferSize)
		ret := nvmlDeviceGetPgpuMetadataString(Device, &PgpuMetadata[0], &BufferSize)
		if ret == SUCCESS {
			return string(PgpuMetadata[:clen(PgpuMetadata)]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return "", ret
		}
		BufferSize *= 2
	}
}

// nvml.DeviceGetVgpuUtilization()
func (l *library) DeviceGetVgpuUtilization(Device Device, LastSeenTimeStamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return) {
	return Device.GetVgpuUtilization(LastSeenTimeStamp)
}

func (Device nvmlDevice) GetVgpuUtilization(LastSeenTimeStamp uint64) (ValueType, []VgpuInstanceUtilizationSample, Return) {
	var SampleValType ValueType
	var VgpuInstanceSamplesCount uint32 = 1 // Will be reduced upon returning
	for {
		UtilizationSamples := make([]VgpuInstanceUtilizationSample, VgpuInstanceSamplesCount)
		ret := nvmlDeviceGetVgpuUtilization(Device, LastSeenTimeStamp, &SampleValType, &VgpuInstanceSamplesCount, &UtilizationSamples[0])
		if ret == SUCCESS {
			return SampleValType, UtilizationSamples[:VgpuInstanceSamplesCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return SampleValType, nil, ret
		}
		VgpuInstanceSamplesCount *= 2
	}
}

// nvml.DeviceGetAttributes()
func (l *library) DeviceGetAttributes(Device Device) (DeviceAttributes, Return) {
	return Device.GetAttributes()
}

func (Device nvmlDevice) GetAttributes() (DeviceAttributes, Return) {
	var Attributes DeviceAttributes
	ret := nvmlDeviceGetAttributes(Device, &Attributes)
	return Attributes, ret
}

// nvml.DeviceGetRemappedRows()
func (l *library) DeviceGetRemappedRows(Device Device) (int, int, bool, bool, Return) {
	return Device.GetRemappedRows()
}

func (Device nvmlDevice) GetRemappedRows() (int, int, bool, bool, Return) {
	var CorrRows, UncRows, IsPending, FailureOccured uint32
	ret := nvmlDeviceGetRemappedRows(Device, &CorrRows, &UncRows, &IsPending, &FailureOccured)
	return int(CorrRows), int(UncRows), (IsPending != 0), (FailureOccured != 0), ret
}

// nvml.DeviceGetRowRemapperHistogram()
func (l *library) DeviceGetRowRemapperHistogram(Device Device) (RowRemapperHistogramValues, Return) {
	return Device.GetRowRemapperHistogram()
}

func (Device nvmlDevice) GetRowRemapperHistogram() (RowRemapperHistogramValues, Return) {
	var Values RowRemapperHistogramValues
	ret := nvmlDeviceGetRowRemapperHistogram(Device, &Values)
	return Values, ret
}

// nvml.DeviceGetArchitecture()
func (l *library) DeviceGetArchitecture(Device Device) (DeviceArchitecture, Return) {
	return Device.GetArchitecture()
}

func (Device nvmlDevice) GetArchitecture() (DeviceArchitecture, Return) {
	var Arch DeviceArchitecture
	ret := nvmlDeviceGetArchitecture(Device, &Arch)
	return Arch, ret
}

// nvml.DeviceGetVgpuProcessUtilization()
func (l *library) DeviceGetVgpuProcessUtilization(Device Device, LastSeenTimeStamp uint64) ([]VgpuProcessUtilizationSample, Return) {
	return Device.GetVgpuProcessUtilization(LastSeenTimeStamp)
}

func (Device nvmlDevice) GetVgpuProcessUtilization(LastSeenTimeStamp uint64) ([]VgpuProcessUtilizationSample, Return) {
	var VgpuProcessSamplesCount uint32 = 1 // Will be reduced upon returning
	for {
		UtilizationSamples := make([]VgpuProcessUtilizationSample, VgpuProcessSamplesCount)
		ret := nvmlDeviceGetVgpuProcessUtilization(Device, LastSeenTimeStamp, &VgpuProcessSamplesCount, &UtilizationSamples[0])
		if ret == SUCCESS {
			return UtilizationSamples[:VgpuProcessSamplesCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		VgpuProcessSamplesCount *= 2
	}
}

// nvml.GetExcludedDeviceCount()
func (l *library) GetExcludedDeviceCount() (int, Return) {
	var DeviceCount uint32
	ret := nvmlGetExcludedDeviceCount(&DeviceCount)
	return int(DeviceCount), ret
}

// nvml.GetExcludedDeviceInfoByIndex()
func (l *library) GetExcludedDeviceInfoByIndex(Index int) (ExcludedDeviceInfo, Return) {
	var Info ExcludedDeviceInfo
	ret := nvmlGetExcludedDeviceInfoByIndex(uint32(Index), &Info)
	return Info, ret
}

// nvml.DeviceSetMigMode()
func (l *library) DeviceSetMigMode(Device Device, Mode int) (Return, Return) {
	return Device.SetMigMode(Mode)
}

func (Device nvmlDevice) SetMigMode(Mode int) (Return, Return) {
	var ActivationStatus Return
	ret := nvmlDeviceSetMigMode(Device, uint32(Mode), &ActivationStatus)
	return ActivationStatus, ret
}

// nvml.DeviceGetMigMode()
func (l *library) DeviceGetMigMode(Device Device) (int, int, Return) {
	return Device.GetMigMode()
}

func (Device nvmlDevice) GetMigMode() (int, int, Return) {
	var CurrentMode, PendingMode uint32
	ret := nvmlDeviceGetMigMode(Device, &CurrentMode, &PendingMode)
	return int(CurrentMode), int(PendingMode), ret
}

// nvml.DeviceGetGpuInstanceProfileInfo()
func (l *library) DeviceGetGpuInstanceProfileInfo(Device Device, Profile int) (GpuInstanceProfileInfo, Return) {
	return Device.GetGpuInstanceProfileInfo(Profile)
}

func (Device nvmlDevice) GetGpuInstanceProfileInfo(Profile int) (GpuInstanceProfileInfo, Return) {
	var Info GpuInstanceProfileInfo
	ret := nvmlDeviceGetGpuInstanceProfileInfo(Device, uint32(Profile), &Info)
	return Info, ret
}

// nvml.DeviceGetGpuInstanceProfileInfoV()
type GpuInstanceProfileInfoV struct {
	device  nvmlDevice
	profile int
}

func (InfoV GpuInstanceProfileInfoV) V1() (GpuInstanceProfileInfo, Return) {
	return DeviceGetGpuInstanceProfileInfo(InfoV.device, InfoV.profile)
}

func (InfoV GpuInstanceProfileInfoV) V2() (GpuInstanceProfileInfo_v2, Return) {
	var Info GpuInstanceProfileInfo_v2
	Info.Version = STRUCT_VERSION(Info, 2)
	ret := nvmlDeviceGetGpuInstanceProfileInfoV(InfoV.device, uint32(InfoV.profile), &Info)
	return Info, ret
}

func (l *library) DeviceGetGpuInstanceProfileInfoV(Device Device, Profile int) GpuInstanceProfileInfoV {
	return Device.GetGpuInstanceProfileInfoV(Profile)
}

func (Device nvmlDevice) GetGpuInstanceProfileInfoV(Profile int) GpuInstanceProfileInfoV {
	return GpuInstanceProfileInfoV{Device, Profile}
}

// nvml.DeviceGetGpuInstancePossiblePlacements()
func (l *library) DeviceGetGpuInstancePossiblePlacements(Device Device, Info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return) {
	return Device.GetGpuInstancePossiblePlacements(Info)
}

func (Device nvmlDevice) GetGpuInstancePossiblePlacements(Info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var Count uint32
	ret := nvmlDeviceGetGpuInstancePossiblePlacements(Device, Info.Id, nil, &Count)
	if ret != SUCCESS {
		return nil, ret
	}
	if Count == 0 {
		return []GpuInstancePlacement{}, ret
	}
	Placements := make([]GpuInstancePlacement, Count)
	ret = nvmlDeviceGetGpuInstancePossiblePlacements(Device, Info.Id, &Placements[0], &Count)
	return Placements[:Count], ret
}

// nvml.DeviceGetGpuInstanceRemainingCapacity()
func (l *library) DeviceGetGpuInstanceRemainingCapacity(Device Device, Info *GpuInstanceProfileInfo) (int, Return) {
	return Device.GetGpuInstanceRemainingCapacity(Info)
}

func (Device nvmlDevice) GetGpuInstanceRemainingCapacity(Info *GpuInstanceProfileInfo) (int, Return) {
	if Info == nil {
		return 0, ERROR_INVALID_ARGUMENT
	}
	var Count uint32
	ret := nvmlDeviceGetGpuInstanceRemainingCapacity(Device, Info.Id, &Count)
	return int(Count), ret
}

// nvml.DeviceCreateGpuInstance()
func (l *library) DeviceCreateGpuInstance(Device Device, Info *GpuInstanceProfileInfo) (GpuInstance, Return) {
	return Device.CreateGpuInstance(Info)
}

func (Device nvmlDevice) CreateGpuInstance(Info *GpuInstanceProfileInfo) (GpuInstance, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var GpuInstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstance(Device, Info.Id, &GpuInstance)
	return GpuInstance, ret
}

// nvml.DeviceCreateGpuInstanceWithPlacement()
func (l *library) DeviceCreateGpuInstanceWithPlacement(Device Device, Info *GpuInstanceProfileInfo, Placement *GpuInstancePlacement) (GpuInstance, Return) {
	return Device.CreateGpuInstanceWithPlacement(Info, Placement)
}

func (Device nvmlDevice) CreateGpuInstanceWithPlacement(Info *GpuInstanceProfileInfo, Placement *GpuInstancePlacement) (GpuInstance, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var GpuInstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstanceWithPlacement(Device, Info.Id, Placement, &GpuInstance)
	return GpuInstance, ret
}

// nvml.GpuInstanceDestroy()
func (l *library) GpuInstanceDestroy(GpuInstance GpuInstance) Return {
	return GpuInstance.Destroy()
}

func (GpuInstance nvmlGpuInstance) Destroy() Return {
	return nvmlGpuInstanceDestroy(GpuInstance)
}

// nvml.DeviceGetGpuInstances()
func (l *library) DeviceGetGpuInstances(Device Device, Info *GpuInstanceProfileInfo) ([]GpuInstance, Return) {
	return Device.GetGpuInstances(Info)
}

func (Device nvmlDevice) GetGpuInstances(Info *GpuInstanceProfileInfo) ([]GpuInstance, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var Count uint32 = Info.InstanceCount
	GpuInstances := make([]nvmlGpuInstance, Count)
	ret := nvmlDeviceGetGpuInstances(Device, Info.Id, &GpuInstances[0], &Count)
	return convertSlice[nvmlGpuInstance, GpuInstance](GpuInstances[:Count]), ret
}

// nvml.DeviceGetGpuInstanceById()
func (l *library) DeviceGetGpuInstanceById(Device Device, Id int) (GpuInstance, Return) {
	return Device.GetGpuInstanceById(Id)
}

func (Device nvmlDevice) GetGpuInstanceById(Id int) (GpuInstance, Return) {
	var GpuInstance nvmlGpuInstance
	ret := nvmlDeviceGetGpuInstanceById(Device, uint32(Id), &GpuInstance)
	return GpuInstance, ret
}

// nvml.GpuInstanceGetInfo()
func (l *library) GpuInstanceGetInfo(GpuInstance GpuInstance) (GpuInstanceInfo, Return) {
	return GpuInstance.GetInfo()
}

func (GpuInstance nvmlGpuInstance) GetInfo() (GpuInstanceInfo, Return) {
	var Info GpuInstanceInfo
	ret := nvmlGpuInstanceGetInfo(GpuInstance, &Info)
	return Info, ret
}

// nvml.GpuInstanceGetComputeInstanceProfileInfo()
func (l *library) GpuInstanceGetComputeInstanceProfileInfo(GpuInstance GpuInstance, Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return) {
	return GpuInstance.GetComputeInstanceProfileInfo(Profile, EngProfile)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstanceProfileInfo(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return) {
	var Info ComputeInstanceProfileInfo
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfo(GpuInstance, uint32(Profile), uint32(EngProfile), &Info)
	return Info, ret
}

// nvml.GpuInstanceGetComputeInstanceProfileInfoV()
type ComputeInstanceProfileInfoV struct {
	gpuInstance nvmlGpuInstance
	profile     int
	engProfile  int
}

func (InfoV ComputeInstanceProfileInfoV) V1() (ComputeInstanceProfileInfo, Return) {
	return GpuInstanceGetComputeInstanceProfileInfo(InfoV.gpuInstance, InfoV.profile, InfoV.engProfile)
}

func (InfoV ComputeInstanceProfileInfoV) V2() (ComputeInstanceProfileInfo_v2, Return) {
	var Info ComputeInstanceProfileInfo_v2
	Info.Version = STRUCT_VERSION(Info, 2)
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfoV(InfoV.gpuInstance, uint32(InfoV.profile), uint32(InfoV.engProfile), &Info)
	return Info, ret
}

func (l *library) GpuInstanceGetComputeInstanceProfileInfoV(GpuInstance GpuInstance, Profile int, EngProfile int) ComputeInstanceProfileInfoV {
	return GpuInstance.GetComputeInstanceProfileInfoV(Profile, EngProfile)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstanceProfileInfoV(Profile int, EngProfile int) ComputeInstanceProfileInfoV {
	return ComputeInstanceProfileInfoV{GpuInstance, Profile, EngProfile}
}

// nvml.GpuInstanceGetComputeInstanceRemainingCapacity()
func (l *library) GpuInstanceGetComputeInstanceRemainingCapacity(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) (int, Return) {
	return GpuInstance.GetComputeInstanceRemainingCapacity(Info)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstanceRemainingCapacity(Info *ComputeInstanceProfileInfo) (int, Return) {
	if Info == nil {
		return 0, ERROR_INVALID_ARGUMENT
	}
	var Count uint32
	ret := nvmlGpuInstanceGetComputeInstanceRemainingCapacity(GpuInstance, Info.Id, &Count)
	return int(Count), ret
}

// nvml.GpuInstanceCreateComputeInstance()
func (l *library) GpuInstanceCreateComputeInstance(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) (ComputeInstance, Return) {
	return GpuInstance.CreateComputeInstance(Info)
}

func (GpuInstance nvmlGpuInstance) CreateComputeInstance(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var ComputeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstance(GpuInstance, Info.Id, &ComputeInstance)
	return ComputeInstance, ret
}

// nvml.ComputeInstanceDestroy()
func (l *library) ComputeInstanceDestroy(ComputeInstance ComputeInstance) Return {
	return ComputeInstance.Destroy()
}

func (ComputeInstance nvmlComputeInstance) Destroy() Return {
	return nvmlComputeInstanceDestroy(ComputeInstance)
}

// nvml.GpuInstanceGetComputeInstances()
func (l *library) GpuInstanceGetComputeInstances(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return) {
	return GpuInstance.GetComputeInstances(Info)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstances(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return) {
	if Info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var Count uint32 = Info.InstanceCount
	ComputeInstances := make([]nvmlComputeInstance, Count)
	ret := nvmlGpuInstanceGetComputeInstances(GpuInstance, Info.Id, &ComputeInstances[0], &Count)
	return convertSlice[nvmlComputeInstance, ComputeInstance](ComputeInstances[:Count]), ret
}

// nvml.GpuInstanceGetComputeInstanceById()
func (l *library) GpuInstanceGetComputeInstanceById(GpuInstance GpuInstance, Id int) (ComputeInstance, Return) {
	return GpuInstance.GetComputeInstanceById(Id)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstanceById(Id int) (ComputeInstance, Return) {
	var ComputeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceGetComputeInstanceById(GpuInstance, uint32(Id), &ComputeInstance)
	return ComputeInstance, ret
}

// nvml.ComputeInstanceGetInfo()
func (l *library) ComputeInstanceGetInfo(ComputeInstance ComputeInstance) (ComputeInstanceInfo, Return) {
	return ComputeInstance.GetInfo()
}

func (ComputeInstance nvmlComputeInstance) GetInfo() (ComputeInstanceInfo, Return) {
	var Info ComputeInstanceInfo
	ret := nvmlComputeInstanceGetInfo(ComputeInstance, &Info)
	return Info, ret
}

// nvml.DeviceIsMigDeviceHandle()
func (l *library) DeviceIsMigDeviceHandle(Device Device) (bool, Return) {
	return Device.IsMigDeviceHandle()
}

func (Device nvmlDevice) IsMigDeviceHandle() (bool, Return) {
	var IsMigDevice uint32
	ret := nvmlDeviceIsMigDeviceHandle(Device, &IsMigDevice)
	return (IsMigDevice != 0), ret
}

// nvml DeviceGetGpuInstanceId()
func (l *library) DeviceGetGpuInstanceId(Device Device) (int, Return) {
	return Device.GetGpuInstanceId()
}

func (Device nvmlDevice) GetGpuInstanceId() (int, Return) {
	var Id uint32
	ret := nvmlDeviceGetGpuInstanceId(Device, &Id)
	return int(Id), ret
}

// nvml.DeviceGetComputeInstanceId()
func (l *library) DeviceGetComputeInstanceId(Device Device) (int, Return) {
	return Device.GetComputeInstanceId()
}

func (Device nvmlDevice) GetComputeInstanceId() (int, Return) {
	var Id uint32
	ret := nvmlDeviceGetComputeInstanceId(Device, &Id)
	return int(Id), ret
}

// nvml.DeviceGetMaxMigDeviceCount()
func (l *library) DeviceGetMaxMigDeviceCount(Device Device) (int, Return) {
	return Device.GetMaxMigDeviceCount()
}

func (Device nvmlDevice) GetMaxMigDeviceCount() (int, Return) {
	var Count uint32
	ret := nvmlDeviceGetMaxMigDeviceCount(Device, &Count)
	return int(Count), ret
}

// nvml.DeviceGetMigDeviceHandleByIndex()
func (l *library) DeviceGetMigDeviceHandleByIndex(device Device, Index int) (Device, Return) {
	return device.GetMigDeviceHandleByIndex(Index)
}

func (Device nvmlDevice) GetMigDeviceHandleByIndex(Index int) (Device, Return) {
	var MigDevice nvmlDevice
	ret := nvmlDeviceGetMigDeviceHandleByIndex(Device, uint32(Index), &MigDevice)
	return MigDevice, ret
}

// nvml.DeviceGetDeviceHandleFromMigDeviceHandle()
func (l *library) DeviceGetDeviceHandleFromMigDeviceHandle(MigDevice Device) (Device, Return) {
	return MigDevice.GetDeviceHandleFromMigDeviceHandle()
}

func (MigDevice nvmlDevice) GetDeviceHandleFromMigDeviceHandle() (Device, Return) {
	var Device nvmlDevice
	ret := nvmlDeviceGetDeviceHandleFromMigDeviceHandle(MigDevice, &Device)
	return Device, ret
}

// nvml.DeviceGetBusType()
func (l *library) DeviceGetBusType(Device Device) (BusType, Return) {
	return Device.GetBusType()
}

func (Device nvmlDevice) GetBusType() (BusType, Return) {
	var Type BusType
	ret := nvmlDeviceGetBusType(Device, &Type)
	return Type, ret
}

// nvml.DeviceSetDefaultFanSpeed_v2()
func (l *library) DeviceSetDefaultFanSpeed_v2(Device Device, Fan int) Return {
	return Device.SetDefaultFanSpeed_v2(Fan)
}

func (Device nvmlDevice) SetDefaultFanSpeed_v2(Fan int) Return {
	return nvmlDeviceSetDefaultFanSpeed_v2(Device, uint32(Fan))
}

// nvml.DeviceGetMinMaxFanSpeed()
func (l *library) DeviceGetMinMaxFanSpeed(Device Device) (int, int, Return) {
	return Device.GetMinMaxFanSpeed()
}

func (Device nvmlDevice) GetMinMaxFanSpeed() (int, int, Return) {
	var MinSpeed, MaxSpeed uint32
	ret := nvmlDeviceGetMinMaxFanSpeed(Device, &MinSpeed, &MaxSpeed)
	return int(MinSpeed), int(MaxSpeed), ret
}

// nvml.DeviceGetThermalSettings()
func (l *library) DeviceGetThermalSettings(Device Device, SensorIndex uint32) (GpuThermalSettings, Return) {
	return Device.GetThermalSettings(SensorIndex)
}

func (Device nvmlDevice) GetThermalSettings(SensorIndex uint32) (GpuThermalSettings, Return) {
	var PThermalSettings GpuThermalSettings
	ret := nvmlDeviceGetThermalSettings(Device, SensorIndex, &PThermalSettings)
	return PThermalSettings, ret
}

// nvml.DeviceGetDefaultEccMode()
func (l *library) DeviceGetDefaultEccMode(Device Device) (EnableState, Return) {
	return Device.GetDefaultEccMode()
}

func (Device nvmlDevice) GetDefaultEccMode() (EnableState, Return) {
	var DefaultMode EnableState
	ret := nvmlDeviceGetDefaultEccMode(Device, &DefaultMode)
	return DefaultMode, ret
}

// nvml.DeviceGetPcieSpeed()
func (l *library) DeviceGetPcieSpeed(Device Device) (int, Return) {
	return Device.GetPcieSpeed()
}

func (Device nvmlDevice) GetPcieSpeed() (int, Return) {
	var PcieSpeed uint32
	ret := nvmlDeviceGetPcieSpeed(Device, &PcieSpeed)
	return int(PcieSpeed), ret
}

// nvml.DeviceGetGspFirmwareVersion()
func (l *library) DeviceGetGspFirmwareVersion(Device Device) (string, Return) {
	return Device.GetGspFirmwareVersion()
}

func (Device nvmlDevice) GetGspFirmwareVersion() (string, Return) {
	Version := make([]byte, GSP_FIRMWARE_VERSION_BUF_SIZE)
	ret := nvmlDeviceGetGspFirmwareVersion(Device, &Version[0])
	return string(Version[:clen(Version)]), ret
}

// nvml.DeviceGetGspFirmwareMode()
func (l *library) DeviceGetGspFirmwareMode(Device Device) (bool, bool, Return) {
	return Device.GetGspFirmwareMode()
}

func (Device nvmlDevice) GetGspFirmwareMode() (bool, bool, Return) {
	var IsEnabled, DefaultMode uint32
	ret := nvmlDeviceGetGspFirmwareMode(Device, &IsEnabled, &DefaultMode)
	return (IsEnabled != 0), (DefaultMode != 0), ret
}

// nvml.DeviceGetDynamicPstatesInfo()
func (l *library) DeviceGetDynamicPstatesInfo(Device Device) (GpuDynamicPstatesInfo, Return) {
	return Device.GetDynamicPstatesInfo()
}

func (Device nvmlDevice) GetDynamicPstatesInfo() (GpuDynamicPstatesInfo, Return) {
	var PDynamicPstatesInfo GpuDynamicPstatesInfo
	ret := nvmlDeviceGetDynamicPstatesInfo(Device, &PDynamicPstatesInfo)
	return PDynamicPstatesInfo, ret
}

// nvml.DeviceSetFanSpeed_v2()
func (l *library) DeviceSetFanSpeed_v2(Device Device, Fan int, Speed int) Return {
	return Device.SetFanSpeed_v2(Fan, Speed)
}

func (Device nvmlDevice) SetFanSpeed_v2(Fan int, Speed int) Return {
	return nvmlDeviceSetFanSpeed_v2(Device, uint32(Fan), uint32(Speed))
}

// nvml.DeviceGetGpcClkVfOffset()
func (l *library) DeviceGetGpcClkVfOffset(Device Device) (int, Return) {
	return Device.GetGpcClkVfOffset()
}

func (Device nvmlDevice) GetGpcClkVfOffset() (int, Return) {
	var Offset int32
	ret := nvmlDeviceGetGpcClkVfOffset(Device, &Offset)
	return int(Offset), ret
}

// nvml.DeviceSetGpcClkVfOffset()
func (l *library) DeviceSetGpcClkVfOffset(Device Device, Offset int) Return {
	return Device.SetGpcClkVfOffset(Offset)
}

func (Device nvmlDevice) SetGpcClkVfOffset(Offset int) Return {
	return nvmlDeviceSetGpcClkVfOffset(Device, int32(Offset))
}

// nvml.DeviceGetMinMaxClockOfPState()
func (l *library) DeviceGetMinMaxClockOfPState(Device Device, _type ClockType, Pstate Pstates) (uint32, uint32, Return) {
	return Device.GetMinMaxClockOfPState(_type, Pstate)
}

func (Device nvmlDevice) GetMinMaxClockOfPState(_type ClockType, Pstate Pstates) (uint32, uint32, Return) {
	var MinClockMHz, MaxClockMHz uint32
	ret := nvmlDeviceGetMinMaxClockOfPState(Device, _type, Pstate, &MinClockMHz, &MaxClockMHz)
	return MinClockMHz, MaxClockMHz, ret
}

// nvml.DeviceGetSupportedPerformanceStates()
func (l *library) DeviceGetSupportedPerformanceStates(Device Device) ([]Pstates, Return) {
	return Device.GetSupportedPerformanceStates()
}

func (Device nvmlDevice) GetSupportedPerformanceStates() ([]Pstates, Return) {
	Pstates := make([]Pstates, MAX_GPU_PERF_PSTATES)
	ret := nvmlDeviceGetSupportedPerformanceStates(Device, &Pstates[0], MAX_GPU_PERF_PSTATES)
	for i := 0; i < MAX_GPU_PERF_PSTATES; i++ {
		if Pstates[i] == PSTATE_UNKNOWN {
			return Pstates[0:i], ret
		}
	}
	return Pstates, ret
}

// nvml.DeviceGetTargetFanSpeed()
func (l *library) DeviceGetTargetFanSpeed(Device Device, Fan int) (int, Return) {
	return Device.GetTargetFanSpeed(Fan)
}

func (Device nvmlDevice) GetTargetFanSpeed(Fan int) (int, Return) {
	var TargetSpeed uint32
	ret := nvmlDeviceGetTargetFanSpeed(Device, uint32(Fan), &TargetSpeed)
	return int(TargetSpeed), ret
}

// nvml.DeviceGetMemClkVfOffset()
func (l *library) DeviceGetMemClkVfOffset(Device Device) (int, Return) {
	return Device.GetMemClkVfOffset()
}

func (Device nvmlDevice) GetMemClkVfOffset() (int, Return) {
	var Offset int32
	ret := nvmlDeviceGetMemClkVfOffset(Device, &Offset)
	return int(Offset), ret
}

// nvml.DeviceSetMemClkVfOffset()
func (l *library) DeviceSetMemClkVfOffset(Device Device, Offset int) Return {
	return Device.SetMemClkVfOffset(Offset)
}

func (Device nvmlDevice) SetMemClkVfOffset(Offset int) Return {
	return nvmlDeviceSetMemClkVfOffset(Device, int32(Offset))
}

// nvml.DeviceGetGpcClkMinMaxVfOffset()
func (l *library) DeviceGetGpcClkMinMaxVfOffset(Device Device) (int, int, Return) {
	return Device.GetGpcClkMinMaxVfOffset()
}

func (Device nvmlDevice) GetGpcClkMinMaxVfOffset() (int, int, Return) {
	var MinOffset, MaxOffset int32
	ret := nvmlDeviceGetGpcClkMinMaxVfOffset(Device, &MinOffset, &MaxOffset)
	return int(MinOffset), int(MaxOffset), ret
}

// nvml.DeviceGetMemClkMinMaxVfOffset()
func (l *library) DeviceGetMemClkMinMaxVfOffset(Device Device) (int, int, Return) {
	return Device.GetMemClkMinMaxVfOffset()
}

func (Device nvmlDevice) GetMemClkMinMaxVfOffset() (int, int, Return) {
	var MinOffset, MaxOffset int32
	ret := nvmlDeviceGetMemClkMinMaxVfOffset(Device, &MinOffset, &MaxOffset)
	return int(MinOffset), int(MaxOffset), ret
}

// nvml.DeviceGetGpuMaxPcieLinkGeneration()
func (l *library) DeviceGetGpuMaxPcieLinkGeneration(Device Device) (int, Return) {
	return Device.GetGpuMaxPcieLinkGeneration()
}

func (Device nvmlDevice) GetGpuMaxPcieLinkGeneration() (int, Return) {
	var MaxLinkGenDevice uint32
	ret := nvmlDeviceGetGpuMaxPcieLinkGeneration(Device, &MaxLinkGenDevice)
	return int(MaxLinkGenDevice), ret
}

// nvml.DeviceGetFanControlPolicy_v2()
func (l *library) DeviceGetFanControlPolicy_v2(Device Device, Fan int) (FanControlPolicy, Return) {
	return Device.GetFanControlPolicy_v2(Fan)
}

func (Device nvmlDevice) GetFanControlPolicy_v2(Fan int) (FanControlPolicy, Return) {
	var Policy FanControlPolicy
	ret := nvmlDeviceGetFanControlPolicy_v2(Device, uint32(Fan), &Policy)
	return Policy, ret
}

// nvml.DeviceSetFanControlPolicy()
func (l *library) DeviceSetFanControlPolicy(Device Device, Fan int, Policy FanControlPolicy) Return {
	return Device.SetFanControlPolicy(Fan, Policy)
}

func (Device nvmlDevice) SetFanControlPolicy(Fan int, Policy FanControlPolicy) Return {
	return nvmlDeviceSetFanControlPolicy(Device, uint32(Fan), Policy)
}

// nvml.DeviceClearFieldValues()
func (l *library) DeviceClearFieldValues(Device Device, Values []FieldValue) Return {
	return Device.ClearFieldValues(Values)
}

func (Device nvmlDevice) ClearFieldValues(Values []FieldValue) Return {
	ValuesCount := len(Values)
	return nvmlDeviceClearFieldValues(Device, int32(ValuesCount), &Values[0])
}

// nvml.DeviceGetVgpuCapabilities()
func (l *library) DeviceGetVgpuCapabilities(Device Device, Capability DeviceVgpuCapability) (bool, Return) {
	return Device.GetVgpuCapabilities(Capability)
}

func (Device nvmlDevice) GetVgpuCapabilities(Capability DeviceVgpuCapability) (bool, Return) {
	var CapResult uint32
	ret := nvmlDeviceGetVgpuCapabilities(Device, Capability, &CapResult)
	return (CapResult != 0), ret
}

// nvml.DeviceGetVgpuSchedulerLog()
func (l *library) DeviceGetVgpuSchedulerLog(Device Device) (VgpuSchedulerLog, Return) {
	return Device.GetVgpuSchedulerLog()
}

func (Device nvmlDevice) GetVgpuSchedulerLog() (VgpuSchedulerLog, Return) {
	var PSchedulerLog VgpuSchedulerLog
	ret := nvmlDeviceGetVgpuSchedulerLog(Device, &PSchedulerLog)
	return PSchedulerLog, ret
}

// nvml.DeviceGetVgpuSchedulerState()
func (l *library) DeviceGetVgpuSchedulerState(Device Device) (VgpuSchedulerGetState, Return) {
	return Device.GetVgpuSchedulerState()
}

func (Device nvmlDevice) GetVgpuSchedulerState() (VgpuSchedulerGetState, Return) {
	var PSchedulerState VgpuSchedulerGetState
	ret := nvmlDeviceGetVgpuSchedulerState(Device, &PSchedulerState)
	return PSchedulerState, ret
}

// nvml.DeviceSetVgpuSchedulerState()
func (l *library) DeviceSetVgpuSchedulerState(Device Device, PSchedulerState *VgpuSchedulerSetState) Return {
	return Device.SetVgpuSchedulerState(PSchedulerState)
}

func (Device nvmlDevice) SetVgpuSchedulerState(PSchedulerState *VgpuSchedulerSetState) Return {
	return nvmlDeviceSetVgpuSchedulerState(Device, PSchedulerState)
}

// nvml.DeviceGetVgpuSchedulerCapabilities()
func (l *library) DeviceGetVgpuSchedulerCapabilities(Device Device) (VgpuSchedulerCapabilities, Return) {
	return Device.GetVgpuSchedulerCapabilities()
}

func (Device nvmlDevice) GetVgpuSchedulerCapabilities() (VgpuSchedulerCapabilities, Return) {
	var PCapabilities VgpuSchedulerCapabilities
	ret := nvmlDeviceGetVgpuSchedulerCapabilities(Device, &PCapabilities)
	return PCapabilities, ret
}

// nvml.GpuInstanceGetComputeInstancePossiblePlacements()
func (l *library) GpuInstanceGetComputeInstancePossiblePlacements(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return) {
	return GpuInstance.GetComputeInstancePossiblePlacements(Info)
}

func (GpuInstance nvmlGpuInstance) GetComputeInstancePossiblePlacements(Info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, Return) {
	var Count uint32
	ret := nvmlGpuInstanceGetComputeInstancePossiblePlacements(GpuInstance, Info.Id, nil, &Count)
	if ret != SUCCESS {
		return nil, ret
	}
	if Count == 0 {
		return []ComputeInstancePlacement{}, ret
	}
	PlacementArray := make([]ComputeInstancePlacement, Count)
	ret = nvmlGpuInstanceGetComputeInstancePossiblePlacements(GpuInstance, Info.Id, &PlacementArray[0], &Count)
	return PlacementArray, ret
}

// nvml.GpuInstanceCreateComputeInstanceWithPlacement()
func (l *library) GpuInstanceCreateComputeInstanceWithPlacement(GpuInstance GpuInstance, Info *ComputeInstanceProfileInfo, Placement *ComputeInstancePlacement) (ComputeInstance, Return) {
	return GpuInstance.CreateComputeInstanceWithPlacement(Info, Placement)
}

func (GpuInstance nvmlGpuInstance) CreateComputeInstanceWithPlacement(Info *ComputeInstanceProfileInfo, Placement *ComputeInstancePlacement) (ComputeInstance, Return) {
	var ComputeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstanceWithPlacement(GpuInstance, Info.Id, Placement, &ComputeInstance)
	return ComputeInstance, ret
}

// nvml.DeviceGetGpuFabricInfo()
func (l *library) DeviceGetGpuFabricInfo(Device Device) (GpuFabricInfo, Return) {
	return Device.GetGpuFabricInfo()
}

func (Device nvmlDevice) GetGpuFabricInfo() (GpuFabricInfo, Return) {
	var GpuFabricInfo GpuFabricInfo
	ret := nvmlDeviceGetGpuFabricInfo(Device, &GpuFabricInfo)
	return GpuFabricInfo, ret
}

// nvml.DeviceCcuGetStreamState()
func (l *library) DeviceCcuGetStreamState(Device Device) (int, Return) {
	return Device.CcuGetStreamState()
}

func (Device nvmlDevice) CcuGetStreamState() (int, Return) {
	var State uint32
	ret := nvmlDeviceCcuGetStreamState(Device, &State)
	return int(State), ret
}

// nvml.DeviceCcuSetStreamState()
func (l *library) DeviceCcuSetStreamState(Device Device, State int) Return {
	return Device.CcuSetStreamState(State)
}

func (Device nvmlDevice) CcuSetStreamState(State int) Return {
	return nvmlDeviceCcuSetStreamState(Device, uint32(State))
}

// nvml.DeviceSetNvLinkDeviceLowPowerThreshold()
func (l *library) DeviceSetNvLinkDeviceLowPowerThreshold(Device Device, Info *NvLinkPowerThres) Return {
	return Device.SetNvLinkDeviceLowPowerThreshold(Info)
}

func (Device nvmlDevice) SetNvLinkDeviceLowPowerThreshold(Info *NvLinkPowerThres) Return {
	return nvmlDeviceSetNvLinkDeviceLowPowerThreshold(Device, Info)
}
