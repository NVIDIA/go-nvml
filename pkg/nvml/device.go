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
	"fmt"
	"reflect"
	"unsafe"
)

// nvmlDeviceHandle attempts to convert a device d to an nvmlDevice.
// This is required for functions such as GetTopologyCommonAncestor which
// accept Device arguments that need to be passed to internal nvml* functions
// as nvmlDevice parameters.
func nvmlDeviceHandle(d Device) nvmlDevice {
	var helper func(val reflect.Value) nvmlDevice
	helper = func(val reflect.Value) nvmlDevice {
		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Type() == reflect.TypeOf(nvmlDevice{}) {
			return val.Interface().(nvmlDevice)
		}

		if val.Kind() != reflect.Struct {
			panic(fmt.Errorf("unable to convert non-struct type %v to nvmlDevice", val.Kind()))
		}

		for i := 0; i < val.Type().NumField(); i++ {
			if !val.Type().Field(i).Anonymous {
				continue
			}
			if !val.Field(i).Type().Implements(reflect.TypeOf((*Device)(nil)).Elem()) {
				continue
			}
			return helper(val.Field(i))
		}
		panic(fmt.Errorf("unable to convert %T to nvmlDevice", d))
	}
	return helper(reflect.ValueOf(d))
}

// EccBitType
type EccBitType = MemoryErrorType

// GpuInstanceInfo includes an interface type for Device instead of nvmlDevice
type GpuInstanceInfo struct {
	Device    Device
	Id        uint32
	ProfileId uint32
	Placement GpuInstancePlacement
}

//nolint:unused
func (g GpuInstanceInfo) convert() nvmlGpuInstanceInfo {
	out := nvmlGpuInstanceInfo{
		Device:    g.Device.(nvmlDevice),
		Id:        g.Id,
		ProfileId: g.ProfileId,
		Placement: g.Placement,
	}
	return out
}

func (g nvmlGpuInstanceInfo) convert() GpuInstanceInfo {
	out := GpuInstanceInfo{
		Device:    g.Device,
		Id:        g.Id,
		ProfileId: g.ProfileId,
		Placement: g.Placement,
	}
	return out
}

// ComputeInstanceInfo includes an interface type for Device instead of nvmlDevice
type ComputeInstanceInfo struct {
	Device      Device
	GpuInstance GpuInstance
	Id          uint32
	ProfileId   uint32
	Placement   ComputeInstancePlacement
}

//nolint:unused
func (c ComputeInstanceInfo) convert() nvmlComputeInstanceInfo {
	out := nvmlComputeInstanceInfo{
		Device:      c.Device.(nvmlDevice),
		GpuInstance: c.GpuInstance.(nvmlGpuInstance),
		Id:          c.Id,
		ProfileId:   c.ProfileId,
		Placement:   c.Placement,
	}
	return out
}

func (c nvmlComputeInstanceInfo) convert() ComputeInstanceInfo {
	out := ComputeInstanceInfo{
		Device:      c.Device,
		GpuInstance: c.GpuInstance,
		Id:          c.Id,
		ProfileId:   c.ProfileId,
		Placement:   c.Placement,
	}
	return out
}

// nvml.DeviceGetCount()
func (l *library) DeviceGetCount() (int, error) {
	var deviceCount uint32
	ret := nvmlDeviceGetCount(&deviceCount)
	return int(deviceCount), ret.error()
}

// nvml.DeviceGetHandleByIndex()
func (l *library) DeviceGetHandleByIndex(index int) (Device, error) {
	var device nvmlDevice
	ret := nvmlDeviceGetHandleByIndex(uint32(index), &device)
	return device, ret.error()
}

// nvml.DeviceGetHandleBySerial()
func (l *library) DeviceGetHandleBySerial(serial string) (Device, error) {
	var device nvmlDevice
	ret := nvmlDeviceGetHandleBySerial(serial+string(rune(0)), &device)
	return device, ret.error()
}

// nvml.DeviceGetHandleByUUID()
func (l *library) DeviceGetHandleByUUID(uuid string) (Device, error) {
	var device nvmlDevice
	ret := nvmlDeviceGetHandleByUUID(uuid+string(rune(0)), &device)
	return device, ret.error()
}

// nvml.DeviceGetHandleByPciBusId()
func (l *library) DeviceGetHandleByPciBusId(pciBusId string) (Device, error) {
	var device nvmlDevice
	ret := nvmlDeviceGetHandleByPciBusId(pciBusId+string(rune(0)), &device)
	return device, ret.error()
}

// nvml.DeviceGetName()
func (l *library) DeviceGetName(device Device) (string, error) {
	return device.GetName()
}

func (device nvmlDevice) GetName() (string, error) {
	name := make([]byte, DEVICE_NAME_V2_BUFFER_SIZE)
	ret := nvmlDeviceGetName(device, &name[0], DEVICE_NAME_V2_BUFFER_SIZE)
	return string(name[:clen(name)]), ret.error()
}

// nvml.DeviceGetBrand()
func (l *library) DeviceGetBrand(device Device) (BrandType, error) {
	return device.GetBrand()
}

func (device nvmlDevice) GetBrand() (BrandType, error) {
	var brandType BrandType
	ret := nvmlDeviceGetBrand(device, &brandType)
	return brandType, ret.error()
}

// nvml.DeviceGetIndex()
func (l *library) DeviceGetIndex(device Device) (int, error) {
	return device.GetIndex()
}

func (device nvmlDevice) GetIndex() (int, error) {
	var index uint32
	ret := nvmlDeviceGetIndex(device, &index)
	return int(index), ret.error()
}

// nvml.DeviceGetSerial()
func (l *library) DeviceGetSerial(device Device) (string, error) {
	return device.GetSerial()
}

func (device nvmlDevice) GetSerial() (string, error) {
	serial := make([]byte, DEVICE_SERIAL_BUFFER_SIZE)
	ret := nvmlDeviceGetSerial(device, &serial[0], DEVICE_SERIAL_BUFFER_SIZE)
	return string(serial[:clen(serial)]), ret.error()
}

// nvml.DeviceGetCpuAffinity()
func (l *library) DeviceGetCpuAffinity(device Device, numCPUs int) ([]uint, error) {
	return device.GetCpuAffinity(numCPUs)
}

func (device nvmlDevice) GetCpuAffinity(numCPUs int) ([]uint, error) {
	cpuSetSize := uint32((numCPUs-1)/int(unsafe.Sizeof(uint(0))) + 1)
	cpuSet := make([]uint, cpuSetSize)
	ret := nvmlDeviceGetCpuAffinity(device, cpuSetSize, &cpuSet[0])
	return cpuSet, ret.error()
}

// nvml.DeviceSetCpuAffinity()
func (l *library) DeviceSetCpuAffinity(device Device) error {
	return device.SetCpuAffinity()
}

func (device nvmlDevice) SetCpuAffinity() error {
	return nvmlDeviceSetCpuAffinity(device).error()
}

// nvml.DeviceClearCpuAffinity()
func (l *library) DeviceClearCpuAffinity(device Device) error {
	return device.ClearCpuAffinity()
}

func (device nvmlDevice) ClearCpuAffinity() error {
	return nvmlDeviceClearCpuAffinity(device).error()
}

// nvml.DeviceGetMemoryAffinity()
func (l *library) DeviceGetMemoryAffinity(device Device, numNodes int, scope AffinityScope) ([]uint, error) {
	return device.GetMemoryAffinity(numNodes, scope)
}

func (device nvmlDevice) GetMemoryAffinity(numNodes int, scope AffinityScope) ([]uint, error) {
	nodeSetSize := uint32((numNodes-1)/int(unsafe.Sizeof(uint(0))) + 1)
	nodeSet := make([]uint, nodeSetSize)
	ret := nvmlDeviceGetMemoryAffinity(device, nodeSetSize, &nodeSet[0], scope)
	return nodeSet, ret.error()
}

// nvml.DeviceGetCpuAffinityWithinScope()
func (l *library) DeviceGetCpuAffinityWithinScope(device Device, numCPUs int, scope AffinityScope) ([]uint, error) {
	return device.GetCpuAffinityWithinScope(numCPUs, scope)
}

func (device nvmlDevice) GetCpuAffinityWithinScope(numCPUs int, scope AffinityScope) ([]uint, error) {
	cpuSetSize := uint32((numCPUs-1)/int(unsafe.Sizeof(uint(0))) + 1)
	cpuSet := make([]uint, cpuSetSize)
	ret := nvmlDeviceGetCpuAffinityWithinScope(device, cpuSetSize, &cpuSet[0], scope)
	return cpuSet, ret.error()
}

// nvml.DeviceGetTopologyCommonAncestor()
func (l *library) DeviceGetTopologyCommonAncestor(device1 Device, device2 Device) (GpuTopologyLevel, error) {
	return device1.GetTopologyCommonAncestor(device2)
}

func (device1 nvmlDevice) GetTopologyCommonAncestor(device2 Device) (GpuTopologyLevel, error) {
	var pathInfo GpuTopologyLevel
	ret := nvmlDeviceGetTopologyCommonAncestorStub(device1, nvmlDeviceHandle(device2), &pathInfo)
	return pathInfo, ret.error()
}

// nvmlDeviceGetTopologyCommonAncestorStub allows us to override this for testing.
var nvmlDeviceGetTopologyCommonAncestorStub = nvmlDeviceGetTopologyCommonAncestor

// nvml.DeviceGetTopologyNearestGpus()
func (l *library) DeviceGetTopologyNearestGpus(device Device, level GpuTopologyLevel) ([]Device, error) {
	return device.GetTopologyNearestGpus(level)
}

func (device nvmlDevice) GetTopologyNearestGpus(level GpuTopologyLevel) ([]Device, error) {
	var count uint32
	ret := nvmlDeviceGetTopologyNearestGpus(device, level, &count, nil)
	if ret != nvmlSUCCESS {
		return nil, ret.error()
	}
	if count == 0 {
		return []Device{}, ret.error()
	}
	deviceArray := make([]nvmlDevice, count)
	ret = nvmlDeviceGetTopologyNearestGpus(device, level, &count, &deviceArray[0])
	return convertSlice[nvmlDevice, Device](deviceArray), ret.error()
}

// nvml.DeviceGetP2PStatus()
func (l *library) DeviceGetP2PStatus(device1 Device, device2 Device, p2pIndex GpuP2PCapsIndex) (GpuP2PStatus, error) {
	return device1.GetP2PStatus(device2, p2pIndex)
}

func (device1 nvmlDevice) GetP2PStatus(device2 Device, p2pIndex GpuP2PCapsIndex) (GpuP2PStatus, error) {
	var p2pStatus GpuP2PStatus
	ret := nvmlDeviceGetP2PStatus(device1, nvmlDeviceHandle(device2), p2pIndex, &p2pStatus)
	return p2pStatus, ret.error()
}

// nvml.DeviceGetUUID()
func (l *library) DeviceGetUUID(device Device) (string, error) {
	return device.GetUUID()
}

func (device nvmlDevice) GetUUID() (string, error) {
	uuid := make([]byte, DEVICE_UUID_V2_BUFFER_SIZE)
	ret := nvmlDeviceGetUUID(device, &uuid[0], DEVICE_UUID_V2_BUFFER_SIZE)
	return string(uuid[:clen(uuid)]), ret.error()
}

// nvml.DeviceGetMinorNumber()
func (l *library) DeviceGetMinorNumber(device Device) (int, error) {
	return device.GetMinorNumber()
}

func (device nvmlDevice) GetMinorNumber() (int, error) {
	var minorNumber uint32
	ret := nvmlDeviceGetMinorNumber(device, &minorNumber)
	return int(minorNumber), ret.error()
}

// nvml.DeviceGetBoardPartNumber()
func (l *library) DeviceGetBoardPartNumber(device Device) (string, error) {
	return device.GetBoardPartNumber()
}

func (device nvmlDevice) GetBoardPartNumber() (string, error) {
	partNumber := make([]byte, DEVICE_PART_NUMBER_BUFFER_SIZE)
	ret := nvmlDeviceGetBoardPartNumber(device, &partNumber[0], DEVICE_PART_NUMBER_BUFFER_SIZE)
	return string(partNumber[:clen(partNumber)]), ret.error()
}

// nvml.DeviceGetInforomVersion()
func (l *library) DeviceGetInforomVersion(device Device, object InforomObject) (string, error) {
	return device.GetInforomVersion(object)
}

func (device nvmlDevice) GetInforomVersion(object InforomObject) (string, error) {
	version := make([]byte, DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetInforomVersion(device, object, &version[0], DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	return string(version[:clen(version)]), ret.error()
}

// nvml.DeviceGetInforomImageVersion()
func (l *library) DeviceGetInforomImageVersion(device Device) (string, error) {
	return device.GetInforomImageVersion()
}

func (device nvmlDevice) GetInforomImageVersion() (string, error) {
	version := make([]byte, DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetInforomImageVersion(device, &version[0], DEVICE_INFOROM_VERSION_BUFFER_SIZE)
	return string(version[:clen(version)]), ret.error()
}

// nvml.DeviceGetInforomConfigurationChecksum()
func (l *library) DeviceGetInforomConfigurationChecksum(device Device) (uint32, error) {
	return device.GetInforomConfigurationChecksum()
}

func (device nvmlDevice) GetInforomConfigurationChecksum() (uint32, error) {
	var checksum uint32
	ret := nvmlDeviceGetInforomConfigurationChecksum(device, &checksum)
	return checksum, ret.error()
}

// nvml.DeviceValidateInforom()
func (l *library) DeviceValidateInforom(device Device) error {
	return device.ValidateInforom()
}

func (device nvmlDevice) ValidateInforom() error {
	return nvmlDeviceValidateInforom(device).error()
}

// nvml.DeviceGetDisplayMode()
func (l *library) DeviceGetDisplayMode(device Device) (EnableState, error) {
	return device.GetDisplayMode()
}

func (device nvmlDevice) GetDisplayMode() (EnableState, error) {
	var display EnableState
	ret := nvmlDeviceGetDisplayMode(device, &display)
	return display, ret.error()
}

// nvml.DeviceGetDisplayActive()
func (l *library) DeviceGetDisplayActive(device Device) (EnableState, error) {
	return device.GetDisplayActive()
}

func (device nvmlDevice) GetDisplayActive() (EnableState, error) {
	var isActive EnableState
	ret := nvmlDeviceGetDisplayActive(device, &isActive)
	return isActive, ret.error()
}

// nvml.DeviceGetPersistenceMode()
func (l *library) DeviceGetPersistenceMode(device Device) (EnableState, error) {
	return device.GetPersistenceMode()
}

func (device nvmlDevice) GetPersistenceMode() (EnableState, error) {
	var mode EnableState
	ret := nvmlDeviceGetPersistenceMode(device, &mode)
	return mode, ret.error()
}

// nvml.DeviceGetPciInfo()
func (l *library) DeviceGetPciInfo(device Device) (PciInfo, error) {
	return device.GetPciInfo()
}

func (device nvmlDevice) GetPciInfo() (PciInfo, error) {
	var pci PciInfo
	ret := nvmlDeviceGetPciInfo(device, &pci)
	return pci, ret.error()
}

// nvml.DeviceGetMaxPcieLinkGeneration()
func (l *library) DeviceGetMaxPcieLinkGeneration(device Device) (int, error) {
	return device.GetMaxPcieLinkGeneration()
}

func (device nvmlDevice) GetMaxPcieLinkGeneration() (int, error) {
	var maxLinkGen uint32
	ret := nvmlDeviceGetMaxPcieLinkGeneration(device, &maxLinkGen)
	return int(maxLinkGen), ret.error()
}

// nvml.DeviceGetMaxPcieLinkWidth()
func (l *library) DeviceGetMaxPcieLinkWidth(device Device) (int, error) {
	return device.GetMaxPcieLinkWidth()
}

func (device nvmlDevice) GetMaxPcieLinkWidth() (int, error) {
	var maxLinkWidth uint32
	ret := nvmlDeviceGetMaxPcieLinkWidth(device, &maxLinkWidth)
	return int(maxLinkWidth), ret.error()
}

// nvml.DeviceGetCurrPcieLinkGeneration()
func (l *library) DeviceGetCurrPcieLinkGeneration(device Device) (int, error) {
	return device.GetCurrPcieLinkGeneration()
}

func (device nvmlDevice) GetCurrPcieLinkGeneration() (int, error) {
	var currLinkGen uint32
	ret := nvmlDeviceGetCurrPcieLinkGeneration(device, &currLinkGen)
	return int(currLinkGen), ret.error()
}

// nvml.DeviceGetCurrPcieLinkWidth()
func (l *library) DeviceGetCurrPcieLinkWidth(device Device) (int, error) {
	return device.GetCurrPcieLinkWidth()
}

func (device nvmlDevice) GetCurrPcieLinkWidth() (int, error) {
	var currLinkWidth uint32
	ret := nvmlDeviceGetCurrPcieLinkWidth(device, &currLinkWidth)
	return int(currLinkWidth), ret.error()
}

// nvml.DeviceGetPcieThroughput()
func (l *library) DeviceGetPcieThroughput(device Device, counter PcieUtilCounter) (uint32, error) {
	return device.GetPcieThroughput(counter)
}

func (device nvmlDevice) GetPcieThroughput(counter PcieUtilCounter) (uint32, error) {
	var value uint32
	ret := nvmlDeviceGetPcieThroughput(device, counter, &value)
	return value, ret.error()
}

// nvml.DeviceGetPcieReplayCounter()
func (l *library) DeviceGetPcieReplayCounter(device Device) (int, error) {
	return device.GetPcieReplayCounter()
}

func (device nvmlDevice) GetPcieReplayCounter() (int, error) {
	var value uint32
	ret := nvmlDeviceGetPcieReplayCounter(device, &value)
	return int(value), ret.error()
}

// nvml.nvmlDeviceGetClockInfo()
func (l *library) DeviceGetClockInfo(device Device, clockType ClockType) (uint32, error) {
	return device.GetClockInfo(clockType)
}

func (device nvmlDevice) GetClockInfo(clockType ClockType) (uint32, error) {
	var clock uint32
	ret := nvmlDeviceGetClockInfo(device, clockType, &clock)
	return clock, ret.error()
}

// nvml.DeviceGetMaxClockInfo()
func (l *library) DeviceGetMaxClockInfo(device Device, clockType ClockType) (uint32, error) {
	return device.GetMaxClockInfo(clockType)
}

func (device nvmlDevice) GetMaxClockInfo(clockType ClockType) (uint32, error) {
	var clock uint32
	ret := nvmlDeviceGetMaxClockInfo(device, clockType, &clock)
	return clock, ret.error()
}

// nvml.DeviceGetApplicationsClock()
func (l *library) DeviceGetApplicationsClock(device Device, clockType ClockType) (uint32, error) {
	return device.GetApplicationsClock(clockType)
}

func (device nvmlDevice) GetApplicationsClock(clockType ClockType) (uint32, error) {
	var clockMHz uint32
	ret := nvmlDeviceGetApplicationsClock(device, clockType, &clockMHz)
	return clockMHz, ret.error()
}

// nvml.DeviceGetDefaultApplicationsClock()
func (l *library) DeviceGetDefaultApplicationsClock(device Device, clockType ClockType) (uint32, error) {
	return device.GetDefaultApplicationsClock(clockType)
}

func (device nvmlDevice) GetDefaultApplicationsClock(clockType ClockType) (uint32, error) {
	var clockMHz uint32
	ret := nvmlDeviceGetDefaultApplicationsClock(device, clockType, &clockMHz)
	return clockMHz, ret.error()
}

// nvml.DeviceResetApplicationsClocks()
func (l *library) DeviceResetApplicationsClocks(device Device) error {
	return device.ResetApplicationsClocks()
}

func (device nvmlDevice) ResetApplicationsClocks() error {
	return nvmlDeviceResetApplicationsClocks(device).error()
}

// nvml.DeviceGetClock()
func (l *library) DeviceGetClock(device Device, clockType ClockType, clockId ClockId) (uint32, error) {
	return device.GetClock(clockType, clockId)
}

func (device nvmlDevice) GetClock(clockType ClockType, clockId ClockId) (uint32, error) {
	var clockMHz uint32
	ret := nvmlDeviceGetClock(device, clockType, clockId, &clockMHz)
	return clockMHz, ret.error()
}

// nvml.DeviceGetMaxCustomerBoostClock()
func (l *library) DeviceGetMaxCustomerBoostClock(device Device, clockType ClockType) (uint32, error) {
	return device.GetMaxCustomerBoostClock(clockType)
}

func (device nvmlDevice) GetMaxCustomerBoostClock(clockType ClockType) (uint32, error) {
	var clockMHz uint32
	ret := nvmlDeviceGetMaxCustomerBoostClock(device, clockType, &clockMHz)
	return clockMHz, ret.error()
}

// nvml.DeviceGetSupportedMemoryClocks()
func (l *library) DeviceGetSupportedMemoryClocks(device Device) (int, uint32, error) {
	return device.GetSupportedMemoryClocks()
}

func (device nvmlDevice) GetSupportedMemoryClocks() (int, uint32, error) {
	var count, clocksMHz uint32
	ret := nvmlDeviceGetSupportedMemoryClocks(device, &count, &clocksMHz)
	return int(count), clocksMHz, ret.error()
}

// nvml.DeviceGetSupportedGraphicsClocks()
func (l *library) DeviceGetSupportedGraphicsClocks(device Device, memoryClockMHz int) (int, uint32, error) {
	return device.GetSupportedGraphicsClocks(memoryClockMHz)
}

func (device nvmlDevice) GetSupportedGraphicsClocks(memoryClockMHz int) (int, uint32, error) {
	var count, clocksMHz uint32
	ret := nvmlDeviceGetSupportedGraphicsClocks(device, uint32(memoryClockMHz), &count, &clocksMHz)
	return int(count), clocksMHz, ret.error()
}

// nvml.DeviceGetAutoBoostedClocksEnabled()
func (l *library) DeviceGetAutoBoostedClocksEnabled(device Device) (EnableState, EnableState, error) {
	return device.GetAutoBoostedClocksEnabled()
}

func (device nvmlDevice) GetAutoBoostedClocksEnabled() (EnableState, EnableState, error) {
	var isEnabled, defaultIsEnabled EnableState
	ret := nvmlDeviceGetAutoBoostedClocksEnabled(device, &isEnabled, &defaultIsEnabled)
	return isEnabled, defaultIsEnabled, ret.error()
}

// nvml.DeviceSetAutoBoostedClocksEnabled()
func (l *library) DeviceSetAutoBoostedClocksEnabled(device Device, enabled EnableState) error {
	return device.SetAutoBoostedClocksEnabled(enabled)
}

func (device nvmlDevice) SetAutoBoostedClocksEnabled(enabled EnableState) error {
	return nvmlDeviceSetAutoBoostedClocksEnabled(device, enabled).error()
}

// nvml.DeviceSetDefaultAutoBoostedClocksEnabled()
func (l *library) DeviceSetDefaultAutoBoostedClocksEnabled(device Device, enabled EnableState, flags uint32) error {
	return device.SetDefaultAutoBoostedClocksEnabled(enabled, flags)
}

func (device nvmlDevice) SetDefaultAutoBoostedClocksEnabled(enabled EnableState, flags uint32) error {
	return nvmlDeviceSetDefaultAutoBoostedClocksEnabled(device, enabled, flags).error()
}

// nvml.DeviceGetFanSpeed()
func (l *library) DeviceGetFanSpeed(device Device) (uint32, error) {
	return device.GetFanSpeed()
}

func (device nvmlDevice) GetFanSpeed() (uint32, error) {
	var speed uint32
	ret := nvmlDeviceGetFanSpeed(device, &speed)
	return speed, ret.error()
}

// nvml.DeviceGetFanSpeed_v2()
func (l *library) DeviceGetFanSpeed_v2(device Device, fan int) (uint32, error) {
	return device.GetFanSpeed_v2(fan)
}

func (device nvmlDevice) GetFanSpeed_v2(fan int) (uint32, error) {
	var speed uint32
	ret := nvmlDeviceGetFanSpeed_v2(device, uint32(fan), &speed)
	return speed, ret.error()
}

// nvml.DeviceGetNumFans()
func (l *library) DeviceGetNumFans(device Device) (int, error) {
	return device.GetNumFans()
}

func (device nvmlDevice) GetNumFans() (int, error) {
	var numFans uint32
	ret := nvmlDeviceGetNumFans(device, &numFans)
	return int(numFans), ret.error()
}

// nvml.DeviceGetTemperature()
func (l *library) DeviceGetTemperature(device Device, sensorType TemperatureSensors) (uint32, error) {
	return device.GetTemperature(sensorType)
}

func (device nvmlDevice) GetTemperature(sensorType TemperatureSensors) (uint32, error) {
	var temp uint32
	ret := nvmlDeviceGetTemperature(device, sensorType, &temp)
	return temp, ret.error()
}

// nvml.DeviceGetTemperatureThreshold()
func (l *library) DeviceGetTemperatureThreshold(device Device, thresholdType TemperatureThresholds) (uint32, error) {
	return device.GetTemperatureThreshold(thresholdType)
}

func (device nvmlDevice) GetTemperatureThreshold(thresholdType TemperatureThresholds) (uint32, error) {
	var temp uint32
	ret := nvmlDeviceGetTemperatureThreshold(device, thresholdType, &temp)
	return temp, ret.error()
}

// nvml.DeviceSetTemperatureThreshold()
func (l *library) DeviceSetTemperatureThreshold(device Device, thresholdType TemperatureThresholds, temp int) error {
	return device.SetTemperatureThreshold(thresholdType, temp)
}

func (device nvmlDevice) SetTemperatureThreshold(thresholdType TemperatureThresholds, temp int) error {
	t := int32(temp)
	ret := nvmlDeviceSetTemperatureThreshold(device, thresholdType, &t)
	return ret.error()
}

// nvml.DeviceGetPerformanceState()
func (l *library) DeviceGetPerformanceState(device Device) (Pstates, error) {
	return device.GetPerformanceState()
}

func (device nvmlDevice) GetPerformanceState() (Pstates, error) {
	var pState Pstates
	ret := nvmlDeviceGetPerformanceState(device, &pState)
	return pState, ret.error()
}

// nvml.DeviceGetCurrentClocksThrottleReasons()
func (l *library) DeviceGetCurrentClocksThrottleReasons(device Device) (uint64, error) {
	return device.GetCurrentClocksThrottleReasons()
}

func (device nvmlDevice) GetCurrentClocksThrottleReasons() (uint64, error) {
	var clocksThrottleReasons uint64
	ret := nvmlDeviceGetCurrentClocksThrottleReasons(device, &clocksThrottleReasons)
	return clocksThrottleReasons, ret.error()
}

// nvml.DeviceGetSupportedClocksThrottleReasons()
func (l *library) DeviceGetSupportedClocksThrottleReasons(device Device) (uint64, error) {
	return device.GetSupportedClocksThrottleReasons()
}

func (device nvmlDevice) GetSupportedClocksThrottleReasons() (uint64, error) {
	var supportedClocksThrottleReasons uint64
	ret := nvmlDeviceGetSupportedClocksThrottleReasons(device, &supportedClocksThrottleReasons)
	return supportedClocksThrottleReasons, ret.error()
}

// nvml.DeviceGetPowerState()
func (l *library) DeviceGetPowerState(device Device) (Pstates, error) {
	return device.GetPowerState()
}

func (device nvmlDevice) GetPowerState() (Pstates, error) {
	var pState Pstates
	ret := nvmlDeviceGetPowerState(device, &pState)
	return pState, ret.error()
}

// nvml.DeviceGetPowerManagementMode()
func (l *library) DeviceGetPowerManagementMode(device Device) (EnableState, error) {
	return device.GetPowerManagementMode()
}

func (device nvmlDevice) GetPowerManagementMode() (EnableState, error) {
	var mode EnableState
	ret := nvmlDeviceGetPowerManagementMode(device, &mode)
	return mode, ret.error()
}

// nvml.DeviceGetPowerManagementLimit()
func (l *library) DeviceGetPowerManagementLimit(device Device) (uint32, error) {
	return device.GetPowerManagementLimit()
}

func (device nvmlDevice) GetPowerManagementLimit() (uint32, error) {
	var limit uint32
	ret := nvmlDeviceGetPowerManagementLimit(device, &limit)
	return limit, ret.error()
}

// nvml.DeviceGetPowerManagementLimitConstraints()
func (l *library) DeviceGetPowerManagementLimitConstraints(device Device) (uint32, uint32, error) {
	return device.GetPowerManagementLimitConstraints()
}

func (device nvmlDevice) GetPowerManagementLimitConstraints() (uint32, uint32, error) {
	var minLimit, maxLimit uint32
	ret := nvmlDeviceGetPowerManagementLimitConstraints(device, &minLimit, &maxLimit)
	return minLimit, maxLimit, ret.error()
}

// nvml.DeviceGetPowerManagementDefaultLimit()
func (l *library) DeviceGetPowerManagementDefaultLimit(device Device) (uint32, error) {
	return device.GetPowerManagementDefaultLimit()
}

func (device nvmlDevice) GetPowerManagementDefaultLimit() (uint32, error) {
	var defaultLimit uint32
	ret := nvmlDeviceGetPowerManagementDefaultLimit(device, &defaultLimit)
	return defaultLimit, ret.error()
}

// nvml.DeviceGetPowerUsage()
func (l *library) DeviceGetPowerUsage(device Device) (uint32, error) {
	return device.GetPowerUsage()
}

func (device nvmlDevice) GetPowerUsage() (uint32, error) {
	var power uint32
	ret := nvmlDeviceGetPowerUsage(device, &power)
	return power, ret.error()
}

// nvml.DeviceGetTotalEnergyConsumption()
func (l *library) DeviceGetTotalEnergyConsumption(device Device) (uint64, error) {
	return device.GetTotalEnergyConsumption()
}

func (device nvmlDevice) GetTotalEnergyConsumption() (uint64, error) {
	var energy uint64
	ret := nvmlDeviceGetTotalEnergyConsumption(device, &energy)
	return energy, ret.error()
}

// nvml.DeviceGetEnforcedPowerLimit()
func (l *library) DeviceGetEnforcedPowerLimit(device Device) (uint32, error) {
	return device.GetEnforcedPowerLimit()
}

func (device nvmlDevice) GetEnforcedPowerLimit() (uint32, error) {
	var limit uint32
	ret := nvmlDeviceGetEnforcedPowerLimit(device, &limit)
	return limit, ret.error()
}

// nvml.DeviceGetGpuOperationMode()
func (l *library) DeviceGetGpuOperationMode(device Device) (GpuOperationMode, GpuOperationMode, error) {
	return device.GetGpuOperationMode()
}

func (device nvmlDevice) GetGpuOperationMode() (GpuOperationMode, GpuOperationMode, error) {
	var current, pending GpuOperationMode
	ret := nvmlDeviceGetGpuOperationMode(device, &current, &pending)
	return current, pending, ret.error()
}

// nvml.DeviceGetMemoryInfo()
func (l *library) DeviceGetMemoryInfo(device Device) (Memory, error) {
	return device.GetMemoryInfo()
}

func (device nvmlDevice) GetMemoryInfo() (Memory, error) {
	var memory Memory
	ret := nvmlDeviceGetMemoryInfo(device, &memory)
	return memory, ret.error()
}

// nvml.DeviceGetMemoryInfo_v2()
func (l *library) DeviceGetMemoryInfo_v2(device Device) (Memory_v2, error) {
	return device.GetMemoryInfo_v2()
}

func (device nvmlDevice) GetMemoryInfo_v2() (Memory_v2, error) {
	var memory Memory_v2
	memory.Version = STRUCT_VERSION(memory, 2)
	ret := nvmlDeviceGetMemoryInfo_v2(device, &memory)
	return memory, ret.error()
}

// nvml.DeviceGetComputeMode()
func (l *library) DeviceGetComputeMode(device Device) (ComputeMode, error) {
	return device.GetComputeMode()
}

func (device nvmlDevice) GetComputeMode() (ComputeMode, error) {
	var mode ComputeMode
	ret := nvmlDeviceGetComputeMode(device, &mode)
	return mode, ret.error()
}

// nvml.DeviceGetCudaComputeCapability()
func (l *library) DeviceGetCudaComputeCapability(device Device) (int, int, error) {
	return device.GetCudaComputeCapability()
}

func (device nvmlDevice) GetCudaComputeCapability() (int, int, error) {
	var major, minor int32
	ret := nvmlDeviceGetCudaComputeCapability(device, &major, &minor)
	return int(major), int(minor), ret.error()
}

// nvml.DeviceGetEccMode()
func (l *library) DeviceGetEccMode(device Device) (EnableState, EnableState, error) {
	return device.GetEccMode()
}

func (device nvmlDevice) GetEccMode() (EnableState, EnableState, error) {
	var current, pending EnableState
	ret := nvmlDeviceGetEccMode(device, &current, &pending)
	return current, pending, ret.error()
}

// nvml.DeviceGetBoardId()
func (l *library) DeviceGetBoardId(device Device) (uint32, error) {
	return device.GetBoardId()
}

func (device nvmlDevice) GetBoardId() (uint32, error) {
	var boardId uint32
	ret := nvmlDeviceGetBoardId(device, &boardId)
	return boardId, ret.error()
}

// nvml.DeviceGetMultiGpuBoard()
func (l *library) DeviceGetMultiGpuBoard(device Device) (int, error) {
	return device.GetMultiGpuBoard()
}

func (device nvmlDevice) GetMultiGpuBoard() (int, error) {
	var multiGpuBool uint32
	ret := nvmlDeviceGetMultiGpuBoard(device, &multiGpuBool)
	return int(multiGpuBool), ret.error()
}

// nvml.DeviceGetTotalEccErrors()
func (l *library) DeviceGetTotalEccErrors(device Device, errorType MemoryErrorType, counterType EccCounterType) (uint64, error) {
	return device.GetTotalEccErrors(errorType, counterType)
}

func (device nvmlDevice) GetTotalEccErrors(errorType MemoryErrorType, counterType EccCounterType) (uint64, error) {
	var eccCounts uint64
	ret := nvmlDeviceGetTotalEccErrors(device, errorType, counterType, &eccCounts)
	return eccCounts, ret.error()
}

// nvml.DeviceGetDetailedEccErrors()
func (l *library) DeviceGetDetailedEccErrors(device Device, errorType MemoryErrorType, counterType EccCounterType) (EccErrorCounts, error) {
	return device.GetDetailedEccErrors(errorType, counterType)
}

func (device nvmlDevice) GetDetailedEccErrors(errorType MemoryErrorType, counterType EccCounterType) (EccErrorCounts, error) {
	var eccCounts EccErrorCounts
	ret := nvmlDeviceGetDetailedEccErrors(device, errorType, counterType, &eccCounts)
	return eccCounts, ret.error()
}

// nvml.DeviceGetMemoryErrorCounter()
func (l *library) DeviceGetMemoryErrorCounter(device Device, errorType MemoryErrorType, counterType EccCounterType, locationType MemoryLocation) (uint64, error) {
	return device.GetMemoryErrorCounter(errorType, counterType, locationType)
}

func (device nvmlDevice) GetMemoryErrorCounter(errorType MemoryErrorType, counterType EccCounterType, locationType MemoryLocation) (uint64, error) {
	var count uint64
	ret := nvmlDeviceGetMemoryErrorCounter(device, errorType, counterType, locationType, &count)
	return count, ret.error()
}

// nvml.DeviceGetUtilizationRates()
func (l *library) DeviceGetUtilizationRates(device Device) (Utilization, error) {
	return device.GetUtilizationRates()
}

func (device nvmlDevice) GetUtilizationRates() (Utilization, error) {
	var utilization Utilization
	ret := nvmlDeviceGetUtilizationRates(device, &utilization)
	return utilization, ret.error()
}

// nvml.DeviceGetEncoderUtilization()
func (l *library) DeviceGetEncoderUtilization(device Device) (uint32, uint32, error) {
	return device.GetEncoderUtilization()
}

func (device nvmlDevice) GetEncoderUtilization() (uint32, uint32, error) {
	var utilization, samplingPeriodUs uint32
	ret := nvmlDeviceGetEncoderUtilization(device, &utilization, &samplingPeriodUs)
	return utilization, samplingPeriodUs, ret.error()
}

// nvml.DeviceGetEncoderCapacity()
func (l *library) DeviceGetEncoderCapacity(device Device, encoderQueryType EncoderType) (int, error) {
	return device.GetEncoderCapacity(encoderQueryType)
}

func (device nvmlDevice) GetEncoderCapacity(encoderQueryType EncoderType) (int, error) {
	var encoderCapacity uint32
	ret := nvmlDeviceGetEncoderCapacity(device, encoderQueryType, &encoderCapacity)
	return int(encoderCapacity), ret.error()
}

// nvml.DeviceGetEncoderStats()
func (l *library) DeviceGetEncoderStats(device Device) (int, uint32, uint32, error) {
	return device.GetEncoderStats()
}

func (device nvmlDevice) GetEncoderStats() (int, uint32, uint32, error) {
	var sessionCount, averageFps, averageLatency uint32
	ret := nvmlDeviceGetEncoderStats(device, &sessionCount, &averageFps, &averageLatency)
	return int(sessionCount), averageFps, averageLatency, ret.error()
}

// nvml.DeviceGetEncoderSessions()
func (l *library) DeviceGetEncoderSessions(device Device) ([]EncoderSessionInfo, error) {
	return device.GetEncoderSessions()
}

func (device nvmlDevice) GetEncoderSessions() ([]EncoderSessionInfo, error) {
	var sessionCount uint32 = 1 // Will be reduced upon returning
	for {
		sessionInfos := make([]EncoderSessionInfo, sessionCount)
		ret := nvmlDeviceGetEncoderSessions(device, &sessionCount, &sessionInfos[0])
		if ret == nvmlSUCCESS {
			return sessionInfos[:sessionCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		sessionCount *= 2
	}
}

// nvml.DeviceGetDecoderUtilization()
func (l *library) DeviceGetDecoderUtilization(device Device) (uint32, uint32, error) {
	return device.GetDecoderUtilization()
}

func (device nvmlDevice) GetDecoderUtilization() (uint32, uint32, error) {
	var utilization, samplingPeriodUs uint32
	ret := nvmlDeviceGetDecoderUtilization(device, &utilization, &samplingPeriodUs)
	return utilization, samplingPeriodUs, ret.error()
}

// nvml.DeviceGetFBCStats()
func (l *library) DeviceGetFBCStats(device Device) (FBCStats, error) {
	return device.GetFBCStats()
}

func (device nvmlDevice) GetFBCStats() (FBCStats, error) {
	var fbcStats FBCStats
	ret := nvmlDeviceGetFBCStats(device, &fbcStats)
	return fbcStats, ret.error()
}

// nvml.DeviceGetFBCSessions()
func (l *library) DeviceGetFBCSessions(device Device) ([]FBCSessionInfo, error) {
	return device.GetFBCSessions()
}

func (device nvmlDevice) GetFBCSessions() ([]FBCSessionInfo, error) {
	var sessionCount uint32 = 1 // Will be reduced upon returning
	for {
		sessionInfo := make([]FBCSessionInfo, sessionCount)
		ret := nvmlDeviceGetFBCSessions(device, &sessionCount, &sessionInfo[0])
		if ret == nvmlSUCCESS {
			return sessionInfo[:sessionCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		sessionCount *= 2
	}
}

// nvml.DeviceGetDriverModel()
func (l *library) DeviceGetDriverModel(device Device) (DriverModel, DriverModel, error) {
	return device.GetDriverModel()
}

func (device nvmlDevice) GetDriverModel() (DriverModel, DriverModel, error) {
	var current, pending DriverModel
	ret := nvmlDeviceGetDriverModel(device, &current, &pending)
	return current, pending, ret.error()
}

// nvml.DeviceGetVbiosVersion()
func (l *library) DeviceGetVbiosVersion(device Device) (string, error) {
	return device.GetVbiosVersion()
}

func (device nvmlDevice) GetVbiosVersion() (string, error) {
	version := make([]byte, DEVICE_VBIOS_VERSION_BUFFER_SIZE)
	ret := nvmlDeviceGetVbiosVersion(device, &version[0], DEVICE_VBIOS_VERSION_BUFFER_SIZE)
	return string(version[:clen(version)]), ret.error()
}

// nvml.DeviceGetBridgeChipInfo()
func (l *library) DeviceGetBridgeChipInfo(device Device) (BridgeChipHierarchy, error) {
	return device.GetBridgeChipInfo()
}

func (device nvmlDevice) GetBridgeChipInfo() (BridgeChipHierarchy, error) {
	var bridgeHierarchy BridgeChipHierarchy
	ret := nvmlDeviceGetBridgeChipInfo(device, &bridgeHierarchy)
	return bridgeHierarchy, ret.error()
}

// nvml.DeviceGetComputeRunningProcesses()
func deviceGetComputeRunningProcesses_v1(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v1, infoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v1(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v1Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetComputeRunningProcesses_v2(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v2, infoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v2(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v2Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetComputeRunningProcesses_v3(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo, infoCount)
		ret := nvmlDeviceGetComputeRunningProcesses_v3(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return infos[:infoCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func (l *library) DeviceGetComputeRunningProcesses(device Device) ([]ProcessInfo, error) {
	return device.GetComputeRunningProcesses()
}

func (device nvmlDevice) GetComputeRunningProcesses() ([]ProcessInfo, error) {
	return deviceGetComputeRunningProcesses(device)
}

// nvml.DeviceGetGraphicsRunningProcesses()
func deviceGetGraphicsRunningProcesses_v1(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v1, infoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v1(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v1Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetGraphicsRunningProcesses_v2(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v2, infoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v2(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v2Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetGraphicsRunningProcesses_v3(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo, infoCount)
		ret := nvmlDeviceGetGraphicsRunningProcesses_v3(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return infos[:infoCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func (l *library) DeviceGetGraphicsRunningProcesses(device Device) ([]ProcessInfo, error) {
	return device.GetGraphicsRunningProcesses()
}

func (device nvmlDevice) GetGraphicsRunningProcesses() ([]ProcessInfo, error) {
	return deviceGetGraphicsRunningProcesses(device)
}

// nvml.DeviceGetMPSComputeRunningProcesses()
func deviceGetMPSComputeRunningProcesses_v1(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v1, infoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v1(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v1Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetMPSComputeRunningProcesses_v2(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo_v2, infoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v2(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return ProcessInfo_v2Slice(infos[:infoCount]).ToProcessInfoSlice(), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func deviceGetMPSComputeRunningProcesses_v3(device nvmlDevice) ([]ProcessInfo, error) {
	var infoCount uint32 = 1 // Will be reduced upon returning
	for {
		infos := make([]ProcessInfo, infoCount)
		ret := nvmlDeviceGetMPSComputeRunningProcesses_v3(device, &infoCount, &infos[0])
		if ret == nvmlSUCCESS {
			return infos[:infoCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		infoCount *= 2
	}
}

func (l *library) DeviceGetMPSComputeRunningProcesses(device Device) ([]ProcessInfo, error) {
	return device.GetMPSComputeRunningProcesses()
}

func (device nvmlDevice) GetMPSComputeRunningProcesses() ([]ProcessInfo, error) {
	return deviceGetMPSComputeRunningProcesses(device)
}

// nvml.DeviceOnSameBoard()
func (l *library) DeviceOnSameBoard(device1 Device, device2 Device) (int, error) {
	return device1.OnSameBoard(device2)
}

func (device1 nvmlDevice) OnSameBoard(device2 Device) (int, error) {
	var onSameBoard int32
	ret := nvmlDeviceOnSameBoard(device1, nvmlDeviceHandle(device2), &onSameBoard)
	return int(onSameBoard), ret.error()
}

// nvml.DeviceGetAPIRestriction()
func (l *library) DeviceGetAPIRestriction(device Device, apiType RestrictedAPI) (EnableState, error) {
	return device.GetAPIRestriction(apiType)
}

func (device nvmlDevice) GetAPIRestriction(apiType RestrictedAPI) (EnableState, error) {
	var isRestricted EnableState
	ret := nvmlDeviceGetAPIRestriction(device, apiType, &isRestricted)
	return isRestricted, ret.error()
}

// nvml.DeviceGetSamples()
func (l *library) DeviceGetSamples(device Device, samplingType SamplingType, lastSeenTimestamp uint64) (ValueType, []Sample, error) {
	return device.GetSamples(samplingType, lastSeenTimestamp)
}

func (device nvmlDevice) GetSamples(samplingType SamplingType, lastSeenTimestamp uint64) (ValueType, []Sample, error) {
	var sampleValType ValueType
	var sampleCount uint32
	ret := nvmlDeviceGetSamples(device, samplingType, lastSeenTimestamp, &sampleValType, &sampleCount, nil)
	if ret != nvmlSUCCESS {
		return sampleValType, nil, ret.error()
	}
	if sampleCount == 0 {
		return sampleValType, []Sample{}, ret.error()
	}
	samples := make([]Sample, sampleCount)
	ret = nvmlDeviceGetSamples(device, samplingType, lastSeenTimestamp, &sampleValType, &sampleCount, &samples[0])
	return sampleValType, samples, ret.error()
}

// nvml.DeviceGetBAR1MemoryInfo()
func (l *library) DeviceGetBAR1MemoryInfo(device Device) (BAR1Memory, error) {
	return device.GetBAR1MemoryInfo()
}

func (device nvmlDevice) GetBAR1MemoryInfo() (BAR1Memory, error) {
	var bar1Memory BAR1Memory
	ret := nvmlDeviceGetBAR1MemoryInfo(device, &bar1Memory)
	return bar1Memory, ret.error()
}

// nvml.DeviceGetViolationStatus()
func (l *library) DeviceGetViolationStatus(device Device, perfPolicyType PerfPolicyType) (ViolationTime, error) {
	return device.GetViolationStatus(perfPolicyType)
}

func (device nvmlDevice) GetViolationStatus(perfPolicyType PerfPolicyType) (ViolationTime, error) {
	var violTime ViolationTime
	ret := nvmlDeviceGetViolationStatus(device, perfPolicyType, &violTime)
	return violTime, ret.error()
}

// nvml.DeviceGetIrqNum()
func (l *library) DeviceGetIrqNum(device Device) (int, error) {
	return device.GetIrqNum()
}

func (device nvmlDevice) GetIrqNum() (int, error) {
	var irqNum uint32
	ret := nvmlDeviceGetIrqNum(device, &irqNum)
	return int(irqNum), ret.error()
}

// nvml.DeviceGetNumGpuCores()
func (l *library) DeviceGetNumGpuCores(device Device) (int, error) {
	return device.GetNumGpuCores()
}

func (device nvmlDevice) GetNumGpuCores() (int, error) {
	var numCores uint32
	ret := nvmlDeviceGetNumGpuCores(device, &numCores)
	return int(numCores), ret.error()
}

// nvml.DeviceGetPowerSource()
func (l *library) DeviceGetPowerSource(device Device) (PowerSource, error) {
	return device.GetPowerSource()
}

func (device nvmlDevice) GetPowerSource() (PowerSource, error) {
	var powerSource PowerSource
	ret := nvmlDeviceGetPowerSource(device, &powerSource)
	return powerSource, ret.error()
}

// nvml.DeviceGetMemoryBusWidth()
func (l *library) DeviceGetMemoryBusWidth(device Device) (uint32, error) {
	return device.GetMemoryBusWidth()
}

func (device nvmlDevice) GetMemoryBusWidth() (uint32, error) {
	var busWidth uint32
	ret := nvmlDeviceGetMemoryBusWidth(device, &busWidth)
	return busWidth, ret.error()
}

// nvml.DeviceGetPcieLinkMaxSpeed()
func (l *library) DeviceGetPcieLinkMaxSpeed(device Device) (uint32, error) {
	return device.GetPcieLinkMaxSpeed()
}

func (device nvmlDevice) GetPcieLinkMaxSpeed() (uint32, error) {
	var maxSpeed uint32
	ret := nvmlDeviceGetPcieLinkMaxSpeed(device, &maxSpeed)
	return maxSpeed, ret.error()
}

// nvml.DeviceGetAdaptiveClockInfoStatus()
func (l *library) DeviceGetAdaptiveClockInfoStatus(device Device) (uint32, error) {
	return device.GetAdaptiveClockInfoStatus()
}

func (device nvmlDevice) GetAdaptiveClockInfoStatus() (uint32, error) {
	var adaptiveClockStatus uint32
	ret := nvmlDeviceGetAdaptiveClockInfoStatus(device, &adaptiveClockStatus)
	return adaptiveClockStatus, ret.error()
}

// nvml.DeviceGetAccountingMode()
func (l *library) DeviceGetAccountingMode(device Device) (EnableState, error) {
	return device.GetAccountingMode()
}

func (device nvmlDevice) GetAccountingMode() (EnableState, error) {
	var mode EnableState
	ret := nvmlDeviceGetAccountingMode(device, &mode)
	return mode, ret.error()
}

// nvml.DeviceGetAccountingStats()
func (l *library) DeviceGetAccountingStats(device Device, pid uint32) (AccountingStats, error) {
	return device.GetAccountingStats(pid)
}

func (device nvmlDevice) GetAccountingStats(pid uint32) (AccountingStats, error) {
	var stats AccountingStats
	ret := nvmlDeviceGetAccountingStats(device, pid, &stats)
	return stats, ret.error()
}

// nvml.DeviceGetAccountingPids()
func (l *library) DeviceGetAccountingPids(device Device) ([]int, error) {
	return device.GetAccountingPids()
}

func (device nvmlDevice) GetAccountingPids() ([]int, error) {
	var count uint32 = 1 // Will be reduced upon returning
	for {
		pids := make([]uint32, count)
		ret := nvmlDeviceGetAccountingPids(device, &count, &pids[0])
		if ret == nvmlSUCCESS {
			return uint32SliceToIntSlice(pids[:count]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		count *= 2
	}
}

// nvml.DeviceGetAccountingBufferSize()
func (l *library) DeviceGetAccountingBufferSize(device Device) (int, error) {
	return device.GetAccountingBufferSize()
}

func (device nvmlDevice) GetAccountingBufferSize() (int, error) {
	var bufferSize uint32
	ret := nvmlDeviceGetAccountingBufferSize(device, &bufferSize)
	return int(bufferSize), ret.error()
}

// nvml.DeviceGetRetiredPages()
func (l *library) DeviceGetRetiredPages(device Device, cause PageRetirementCause) ([]uint64, error) {
	return device.GetRetiredPages(cause)
}

func (device nvmlDevice) GetRetiredPages(cause PageRetirementCause) ([]uint64, error) {
	var pageCount uint32 = 1 // Will be reduced upon returning
	for {
		addresses := make([]uint64, pageCount)
		ret := nvmlDeviceGetRetiredPages(device, cause, &pageCount, &addresses[0])
		if ret == nvmlSUCCESS {
			return addresses[:pageCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		pageCount *= 2
	}
}

// nvml.DeviceGetRetiredPages_v2()
func (l *library) DeviceGetRetiredPages_v2(device Device, cause PageRetirementCause) ([]uint64, []uint64, error) {
	return device.GetRetiredPages_v2(cause)
}

func (device nvmlDevice) GetRetiredPages_v2(cause PageRetirementCause) ([]uint64, []uint64, error) {
	var pageCount uint32 = 1 // Will be reduced upon returning
	for {
		addresses := make([]uint64, pageCount)
		timestamps := make([]uint64, pageCount)
		ret := nvmlDeviceGetRetiredPages_v2(device, cause, &pageCount, &addresses[0], &timestamps[0])
		if ret == nvmlSUCCESS {
			return addresses[:pageCount], timestamps[:pageCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, nil, ret.error()
		}
		pageCount *= 2
	}
}

// nvml.DeviceGetRetiredPagesPendingStatus()
func (l *library) DeviceGetRetiredPagesPendingStatus(device Device) (EnableState, error) {
	return device.GetRetiredPagesPendingStatus()
}

func (device nvmlDevice) GetRetiredPagesPendingStatus() (EnableState, error) {
	var isPending EnableState
	ret := nvmlDeviceGetRetiredPagesPendingStatus(device, &isPending)
	return isPending, ret.error()
}

// nvml.DeviceSetPersistenceMode()
func (l *library) DeviceSetPersistenceMode(device Device, mode EnableState) error {
	return device.SetPersistenceMode(mode)
}

func (device nvmlDevice) SetPersistenceMode(mode EnableState) error {
	return nvmlDeviceSetPersistenceMode(device, mode).error()
}

// nvml.DeviceSetComputeMode()
func (l *library) DeviceSetComputeMode(device Device, mode ComputeMode) error {
	return device.SetComputeMode(mode)
}

func (device nvmlDevice) SetComputeMode(mode ComputeMode) error {
	return nvmlDeviceSetComputeMode(device, mode).error()
}

// nvml.DeviceSetEccMode()
func (l *library) DeviceSetEccMode(device Device, ecc EnableState) error {
	return device.SetEccMode(ecc)
}

func (device nvmlDevice) SetEccMode(ecc EnableState) error {
	return nvmlDeviceSetEccMode(device, ecc).error()
}

// nvml.DeviceClearEccErrorCounts()
func (l *library) DeviceClearEccErrorCounts(device Device, counterType EccCounterType) error {
	return device.ClearEccErrorCounts(counterType)
}

func (device nvmlDevice) ClearEccErrorCounts(counterType EccCounterType) error {
	return nvmlDeviceClearEccErrorCounts(device, counterType).error()
}

// nvml.DeviceSetDriverModel()
func (l *library) DeviceSetDriverModel(device Device, driverModel DriverModel, flags uint32) error {
	return device.SetDriverModel(driverModel, flags)
}

func (device nvmlDevice) SetDriverModel(driverModel DriverModel, flags uint32) error {
	return nvmlDeviceSetDriverModel(device, driverModel, flags).error()
}

// nvml.DeviceSetGpuLockedClocks()
func (l *library) DeviceSetGpuLockedClocks(device Device, minGpuClockMHz uint32, maxGpuClockMHz uint32) error {
	return device.SetGpuLockedClocks(minGpuClockMHz, maxGpuClockMHz)
}

func (device nvmlDevice) SetGpuLockedClocks(minGpuClockMHz uint32, maxGpuClockMHz uint32) error {
	return nvmlDeviceSetGpuLockedClocks(device, minGpuClockMHz, maxGpuClockMHz).error()
}

// nvml.DeviceResetGpuLockedClocks()
func (l *library) DeviceResetGpuLockedClocks(device Device) error {
	return device.ResetGpuLockedClocks()
}

func (device nvmlDevice) ResetGpuLockedClocks() error {
	return nvmlDeviceResetGpuLockedClocks(device).error()
}

// nvmlDeviceSetMemoryLockedClocks()
func (l *library) DeviceSetMemoryLockedClocks(device Device, minMemClockMHz uint32, maxMemClockMHz uint32) error {
	return device.SetMemoryLockedClocks(minMemClockMHz, maxMemClockMHz)
}

func (device nvmlDevice) SetMemoryLockedClocks(minMemClockMHz uint32, maxMemClockMHz uint32) error {
	return nvmlDeviceSetMemoryLockedClocks(device, minMemClockMHz, maxMemClockMHz).error()
}

// nvmlDeviceResetMemoryLockedClocks()
func (l *library) DeviceResetMemoryLockedClocks(device Device) error {
	return device.ResetMemoryLockedClocks()
}

func (device nvmlDevice) ResetMemoryLockedClocks() error {
	return nvmlDeviceResetMemoryLockedClocks(device).error()
}

// nvml.DeviceGetClkMonStatus()
func (l *library) DeviceGetClkMonStatus(device Device) (ClkMonStatus, error) {
	return device.GetClkMonStatus()
}

func (device nvmlDevice) GetClkMonStatus() (ClkMonStatus, error) {
	var status ClkMonStatus
	ret := nvmlDeviceGetClkMonStatus(device, &status)
	return status, ret.error()
}

// nvml.DeviceSetApplicationsClocks()
func (l *library) DeviceSetApplicationsClocks(device Device, memClockMHz uint32, graphicsClockMHz uint32) error {
	return device.SetApplicationsClocks(memClockMHz, graphicsClockMHz)
}

func (device nvmlDevice) SetApplicationsClocks(memClockMHz uint32, graphicsClockMHz uint32) error {
	return nvmlDeviceSetApplicationsClocks(device, memClockMHz, graphicsClockMHz).error()
}

// nvml.DeviceSetPowerManagementLimit()
func (l *library) DeviceSetPowerManagementLimit(device Device, limit uint32) error {
	return device.SetPowerManagementLimit(limit)
}

func (device nvmlDevice) SetPowerManagementLimit(limit uint32) error {
	return nvmlDeviceSetPowerManagementLimit(device, limit).error()
}

// nvml.DeviceSetGpuOperationMode()
func (l *library) DeviceSetGpuOperationMode(device Device, mode GpuOperationMode) error {
	return device.SetGpuOperationMode(mode)
}

func (device nvmlDevice) SetGpuOperationMode(mode GpuOperationMode) error {
	return nvmlDeviceSetGpuOperationMode(device, mode).error()
}

// nvml.DeviceSetAPIRestriction()
func (l *library) DeviceSetAPIRestriction(device Device, apiType RestrictedAPI, isRestricted EnableState) error {
	return device.SetAPIRestriction(apiType, isRestricted)
}

func (device nvmlDevice) SetAPIRestriction(apiType RestrictedAPI, isRestricted EnableState) error {
	return nvmlDeviceSetAPIRestriction(device, apiType, isRestricted).error()
}

// nvml.DeviceSetAccountingMode()
func (l *library) DeviceSetAccountingMode(device Device, mode EnableState) error {
	return device.SetAccountingMode(mode)
}

func (device nvmlDevice) SetAccountingMode(mode EnableState) error {
	return nvmlDeviceSetAccountingMode(device, mode).error()
}

// nvml.DeviceClearAccountingPids()
func (l *library) DeviceClearAccountingPids(device Device) error {
	return device.ClearAccountingPids()
}

func (device nvmlDevice) ClearAccountingPids() error {
	return nvmlDeviceClearAccountingPids(device).error()
}

// nvml.DeviceGetNvLinkState()
func (l *library) DeviceGetNvLinkState(device Device, link int) (EnableState, error) {
	return device.GetNvLinkState(link)
}

func (device nvmlDevice) GetNvLinkState(link int) (EnableState, error) {
	var isActive EnableState
	ret := nvmlDeviceGetNvLinkState(device, uint32(link), &isActive)
	return isActive, ret.error()
}

// nvml.DeviceGetNvLinkVersion()
func (l *library) DeviceGetNvLinkVersion(device Device, link int) (uint32, error) {
	return device.GetNvLinkVersion(link)
}

func (device nvmlDevice) GetNvLinkVersion(link int) (uint32, error) {
	var version uint32
	ret := nvmlDeviceGetNvLinkVersion(device, uint32(link), &version)
	return version, ret.error()
}

// nvml.DeviceGetNvLinkCapability()
func (l *library) DeviceGetNvLinkCapability(device Device, link int, capability NvLinkCapability) (uint32, error) {
	return device.GetNvLinkCapability(link, capability)
}

func (device nvmlDevice) GetNvLinkCapability(link int, capability NvLinkCapability) (uint32, error) {
	var capResult uint32
	ret := nvmlDeviceGetNvLinkCapability(device, uint32(link), capability, &capResult)
	return capResult, ret.error()
}

// nvml.DeviceGetNvLinkRemotePciInfo()
func (l *library) DeviceGetNvLinkRemotePciInfo(device Device, link int) (PciInfo, error) {
	return device.GetNvLinkRemotePciInfo(link)
}

func (device nvmlDevice) GetNvLinkRemotePciInfo(link int) (PciInfo, error) {
	var pci PciInfo
	ret := nvmlDeviceGetNvLinkRemotePciInfo(device, uint32(link), &pci)
	return pci, ret.error()
}

// nvml.DeviceGetNvLinkErrorCounter()
func (l *library) DeviceGetNvLinkErrorCounter(device Device, link int, counter NvLinkErrorCounter) (uint64, error) {
	return device.GetNvLinkErrorCounter(link, counter)
}

func (device nvmlDevice) GetNvLinkErrorCounter(link int, counter NvLinkErrorCounter) (uint64, error) {
	var counterValue uint64
	ret := nvmlDeviceGetNvLinkErrorCounter(device, uint32(link), counter, &counterValue)
	return counterValue, ret.error()
}

// nvml.DeviceResetNvLinkErrorCounters()
func (l *library) DeviceResetNvLinkErrorCounters(device Device, link int) error {
	return device.ResetNvLinkErrorCounters(link)
}

func (device nvmlDevice) ResetNvLinkErrorCounters(link int) error {
	return nvmlDeviceResetNvLinkErrorCounters(device, uint32(link)).error()
}

// nvml.DeviceSetNvLinkUtilizationControl()
func (l *library) DeviceSetNvLinkUtilizationControl(device Device, link int, counter int, control *NvLinkUtilizationControl, reset bool) error {
	return device.SetNvLinkUtilizationControl(link, counter, control, reset)
}

func (device nvmlDevice) SetNvLinkUtilizationControl(link int, counter int, control *NvLinkUtilizationControl, reset bool) error {
	resetValue := uint32(0)
	if reset {
		resetValue = 1
	}
	return nvmlDeviceSetNvLinkUtilizationControl(device, uint32(link), uint32(counter), control, resetValue).error()
}

// nvml.DeviceGetNvLinkUtilizationControl()
func (l *library) DeviceGetNvLinkUtilizationControl(device Device, link int, counter int) (NvLinkUtilizationControl, error) {
	return device.GetNvLinkUtilizationControl(link, counter)
}

func (device nvmlDevice) GetNvLinkUtilizationControl(link int, counter int) (NvLinkUtilizationControl, error) {
	var control NvLinkUtilizationControl
	ret := nvmlDeviceGetNvLinkUtilizationControl(device, uint32(link), uint32(counter), &control)
	return control, ret.error()
}

// nvml.DeviceGetNvLinkUtilizationCounter()
func (l *library) DeviceGetNvLinkUtilizationCounter(device Device, link int, counter int) (uint64, uint64, error) {
	return device.GetNvLinkUtilizationCounter(link, counter)
}

func (device nvmlDevice) GetNvLinkUtilizationCounter(link int, counter int) (uint64, uint64, error) {
	var rxCounter, txCounter uint64
	ret := nvmlDeviceGetNvLinkUtilizationCounter(device, uint32(link), uint32(counter), &rxCounter, &txCounter)
	return rxCounter, txCounter, ret.error()
}

// nvml.DeviceFreezeNvLinkUtilizationCounter()
func (l *library) DeviceFreezeNvLinkUtilizationCounter(device Device, link int, counter int, freeze EnableState) error {
	return device.FreezeNvLinkUtilizationCounter(link, counter, freeze)
}

func (device nvmlDevice) FreezeNvLinkUtilizationCounter(link int, counter int, freeze EnableState) error {
	return nvmlDeviceFreezeNvLinkUtilizationCounter(device, uint32(link), uint32(counter), freeze).error()
}

// nvml.DeviceResetNvLinkUtilizationCounter()
func (l *library) DeviceResetNvLinkUtilizationCounter(device Device, link int, counter int) error {
	return device.ResetNvLinkUtilizationCounter(link, counter)
}

func (device nvmlDevice) ResetNvLinkUtilizationCounter(link int, counter int) error {
	return nvmlDeviceResetNvLinkUtilizationCounter(device, uint32(link), uint32(counter)).error()
}

// nvml.DeviceGetNvLinkRemoteDeviceType()
func (l *library) DeviceGetNvLinkRemoteDeviceType(device Device, link int) (IntNvLinkDeviceType, error) {
	return device.GetNvLinkRemoteDeviceType(link)
}

func (device nvmlDevice) GetNvLinkRemoteDeviceType(link int) (IntNvLinkDeviceType, error) {
	var nvLinkDeviceType IntNvLinkDeviceType
	ret := nvmlDeviceGetNvLinkRemoteDeviceType(device, uint32(link), &nvLinkDeviceType)
	return nvLinkDeviceType, ret.error()
}

// nvml.DeviceRegisterEvents()
func (l *library) DeviceRegisterEvents(device Device, eventTypes uint64, set EventSet) error {
	return device.RegisterEvents(eventTypes, set)
}

func (device nvmlDevice) RegisterEvents(eventTypes uint64, set EventSet) error {
	return nvmlDeviceRegisterEvents(device, eventTypes, set.(nvmlEventSet)).error()
}

// nvmlDeviceGetSupportedEventTypes()
func (l *library) DeviceGetSupportedEventTypes(device Device) (uint64, error) {
	return device.GetSupportedEventTypes()
}

func (device nvmlDevice) GetSupportedEventTypes() (uint64, error) {
	var eventTypes uint64
	ret := nvmlDeviceGetSupportedEventTypes(device, &eventTypes)
	return eventTypes, ret.error()
}

// nvml.DeviceModifyDrainState()
func (l *library) DeviceModifyDrainState(pciInfo *PciInfo, newState EnableState) error {
	return nvmlDeviceModifyDrainState(pciInfo, newState).error()
}

// nvml.DeviceQueryDrainState()
func (l *library) DeviceQueryDrainState(pciInfo *PciInfo) (EnableState, error) {
	var currentState EnableState
	ret := nvmlDeviceQueryDrainState(pciInfo, &currentState)
	return currentState, ret.error()
}

// nvml.DeviceRemoveGpu()
func (l *library) DeviceRemoveGpu(pciInfo *PciInfo) error {
	return nvmlDeviceRemoveGpu(pciInfo).error()
}

// nvml.DeviceRemoveGpu_v2()
func (l *library) DeviceRemoveGpu_v2(pciInfo *PciInfo, gpuState DetachGpuState, linkState PcieLinkState) error {
	return nvmlDeviceRemoveGpu_v2(pciInfo, gpuState, linkState).error()
}

// nvml.DeviceDiscoverGpus()
func (l *library) DeviceDiscoverGpus() (PciInfo, error) {
	var pciInfo PciInfo
	ret := nvmlDeviceDiscoverGpus(&pciInfo)
	return pciInfo, ret.error()
}

// nvml.DeviceGetFieldValues()
func (l *library) DeviceGetFieldValues(device Device, values []FieldValue) error {
	return device.GetFieldValues(values)
}

func (device nvmlDevice) GetFieldValues(values []FieldValue) error {
	valuesCount := len(values)
	return nvmlDeviceGetFieldValues(device, int32(valuesCount), &values[0]).error()
}

// nvml.DeviceGetVirtualizationMode()
func (l *library) DeviceGetVirtualizationMode(device Device) (GpuVirtualizationMode, error) {
	return device.GetVirtualizationMode()
}

func (device nvmlDevice) GetVirtualizationMode() (GpuVirtualizationMode, error) {
	var pVirtualMode GpuVirtualizationMode
	ret := nvmlDeviceGetVirtualizationMode(device, &pVirtualMode)
	return pVirtualMode, ret.error()
}

// nvml.DeviceGetHostVgpuMode()
func (l *library) DeviceGetHostVgpuMode(device Device) (HostVgpuMode, error) {
	return device.GetHostVgpuMode()
}

func (device nvmlDevice) GetHostVgpuMode() (HostVgpuMode, error) {
	var pHostVgpuMode HostVgpuMode
	ret := nvmlDeviceGetHostVgpuMode(device, &pHostVgpuMode)
	return pHostVgpuMode, ret.error()
}

// nvml.DeviceSetVirtualizationMode()
func (l *library) DeviceSetVirtualizationMode(device Device, virtualMode GpuVirtualizationMode) error {
	return device.SetVirtualizationMode(virtualMode)
}

func (device nvmlDevice) SetVirtualizationMode(virtualMode GpuVirtualizationMode) error {
	return nvmlDeviceSetVirtualizationMode(device, virtualMode).error()
}

// nvml.DeviceGetGridLicensableFeatures()
func (l *library) DeviceGetGridLicensableFeatures(device Device) (GridLicensableFeatures, error) {
	return device.GetGridLicensableFeatures()
}

func (device nvmlDevice) GetGridLicensableFeatures() (GridLicensableFeatures, error) {
	var pGridLicensableFeatures GridLicensableFeatures
	ret := nvmlDeviceGetGridLicensableFeatures(device, &pGridLicensableFeatures)
	return pGridLicensableFeatures, ret.error()
}

// nvml.DeviceGetProcessUtilization()
func (l *library) DeviceGetProcessUtilization(device Device, lastSeenTimestamp uint64) ([]ProcessUtilizationSample, error) {
	return device.GetProcessUtilization(lastSeenTimestamp)
}

func (device nvmlDevice) GetProcessUtilization(lastSeenTimestamp uint64) ([]ProcessUtilizationSample, error) {
	var processSamplesCount uint32
	ret := nvmlDeviceGetProcessUtilization(device, nil, &processSamplesCount, lastSeenTimestamp)
	if ret != nvmlERROR_INSUFFICIENT_SIZE {
		return nil, ret.error()
	}
	if processSamplesCount == 0 {
		return []ProcessUtilizationSample{}, ret.error()
	}
	utilization := make([]ProcessUtilizationSample, processSamplesCount)
	ret = nvmlDeviceGetProcessUtilization(device, &utilization[0], &processSamplesCount, lastSeenTimestamp)
	return utilization[:processSamplesCount], ret.error()
}

// nvml.DeviceGetSupportedVgpus()
func (l *library) DeviceGetSupportedVgpus(device Device) ([]VgpuTypeId, error) {
	return device.GetSupportedVgpus()
}

func (device nvmlDevice) GetSupportedVgpus() ([]VgpuTypeId, error) {
	var vgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		vgpuTypeIds := make([]nvmlVgpuTypeId, vgpuCount)
		ret := nvmlDeviceGetSupportedVgpus(device, &vgpuCount, &vgpuTypeIds[0])
		if ret == nvmlSUCCESS {
			return convertSlice[nvmlVgpuTypeId, VgpuTypeId](vgpuTypeIds[:vgpuCount]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		vgpuCount *= 2
	}
}

// nvml.DeviceGetCreatableVgpus()
func (l *library) DeviceGetCreatableVgpus(device Device) ([]VgpuTypeId, error) {
	return device.GetCreatableVgpus()
}

func (device nvmlDevice) GetCreatableVgpus() ([]VgpuTypeId, error) {
	var vgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		vgpuTypeIds := make([]nvmlVgpuTypeId, vgpuCount)
		ret := nvmlDeviceGetCreatableVgpus(device, &vgpuCount, &vgpuTypeIds[0])
		if ret == nvmlSUCCESS {
			return convertSlice[nvmlVgpuTypeId, VgpuTypeId](vgpuTypeIds[:vgpuCount]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		vgpuCount *= 2
	}
}

// nvml.DeviceGetActiveVgpus()
func (l *library) DeviceGetActiveVgpus(device Device) ([]VgpuInstance, error) {
	return device.GetActiveVgpus()
}

func (device nvmlDevice) GetActiveVgpus() ([]VgpuInstance, error) {
	var vgpuCount uint32 = 1 // Will be reduced upon returning
	for {
		vgpuInstances := make([]nvmlVgpuInstance, vgpuCount)
		ret := nvmlDeviceGetActiveVgpus(device, &vgpuCount, &vgpuInstances[0])
		if ret == nvmlSUCCESS {
			return convertSlice[nvmlVgpuInstance, VgpuInstance](vgpuInstances[:vgpuCount]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		vgpuCount *= 2
	}
}

// nvml.DeviceGetVgpuMetadata()
func (l *library) DeviceGetVgpuMetadata(device Device) (VgpuPgpuMetadata, error) {
	return device.GetVgpuMetadata()
}

func (device nvmlDevice) GetVgpuMetadata() (VgpuPgpuMetadata, error) {
	var vgpuPgpuMetadata VgpuPgpuMetadata
	opaqueDataSize := unsafe.Sizeof(vgpuPgpuMetadata.nvmlVgpuPgpuMetadata.OpaqueData)
	vgpuPgpuMetadataSize := unsafe.Sizeof(vgpuPgpuMetadata.nvmlVgpuPgpuMetadata) - opaqueDataSize
	for {
		bufferSize := uint32(vgpuPgpuMetadataSize + opaqueDataSize)
		buffer := make([]byte, bufferSize)
		nvmlVgpuPgpuMetadataPtr := (*nvmlVgpuPgpuMetadata)(unsafe.Pointer(&buffer[0]))
		ret := nvmlDeviceGetVgpuMetadata(device, nvmlVgpuPgpuMetadataPtr, &bufferSize)
		if ret == nvmlSUCCESS {
			vgpuPgpuMetadata.nvmlVgpuPgpuMetadata = *nvmlVgpuPgpuMetadataPtr
			vgpuPgpuMetadata.OpaqueData = buffer[vgpuPgpuMetadataSize:bufferSize]
			return vgpuPgpuMetadata, ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return vgpuPgpuMetadata, ret.error()
		}
		opaqueDataSize = 2 * opaqueDataSize
	}
}

// nvml.DeviceGetPgpuMetadataString()
func (l *library) DeviceGetPgpuMetadataString(device Device) (string, error) {
	return device.GetPgpuMetadataString()
}

func (device nvmlDevice) GetPgpuMetadataString() (string, error) {
	var bufferSize uint32 = 1 // Will be reduced upon returning
	for {
		pgpuMetadata := make([]byte, bufferSize)
		ret := nvmlDeviceGetPgpuMetadataString(device, &pgpuMetadata[0], &bufferSize)
		if ret == nvmlSUCCESS {
			return string(pgpuMetadata[:clen(pgpuMetadata)]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return "", ret.error()
		}
		bufferSize *= 2
	}
}

// nvml.DeviceGetVgpuUtilization()
func (l *library) DeviceGetVgpuUtilization(device Device, lastSeenTimestamp uint64) (ValueType, []VgpuInstanceUtilizationSample, error) {
	return device.GetVgpuUtilization(lastSeenTimestamp)
}

func (device nvmlDevice) GetVgpuUtilization(lastSeenTimestamp uint64) (ValueType, []VgpuInstanceUtilizationSample, error) {
	var sampleValType ValueType
	var vgpuInstanceSamplesCount uint32 = 1 // Will be reduced upon returning
	for {
		utilizationSamples := make([]VgpuInstanceUtilizationSample, vgpuInstanceSamplesCount)
		ret := nvmlDeviceGetVgpuUtilization(device, lastSeenTimestamp, &sampleValType, &vgpuInstanceSamplesCount, &utilizationSamples[0])
		if ret == nvmlSUCCESS {
			return sampleValType, utilizationSamples[:vgpuInstanceSamplesCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return sampleValType, nil, ret.error()
		}
		vgpuInstanceSamplesCount *= 2
	}
}

// nvml.DeviceGetAttributes()
func (l *library) DeviceGetAttributes(device Device) (DeviceAttributes, error) {
	return device.GetAttributes()
}

func (device nvmlDevice) GetAttributes() (DeviceAttributes, error) {
	var attributes DeviceAttributes
	ret := nvmlDeviceGetAttributes(device, &attributes)
	return attributes, ret.error()
}

// nvml.DeviceGetRemappedRows()
func (l *library) DeviceGetRemappedRows(device Device) (int, int, bool, bool, error) {
	return device.GetRemappedRows()
}

func (device nvmlDevice) GetRemappedRows() (int, int, bool, bool, error) {
	var corrRows, uncRows, isPending, failureOccured uint32
	ret := nvmlDeviceGetRemappedRows(device, &corrRows, &uncRows, &isPending, &failureOccured)
	return int(corrRows), int(uncRows), (isPending != 0), (failureOccured != 0), ret.error()
}

// nvml.DeviceGetRowRemapperHistogram()
func (l *library) DeviceGetRowRemapperHistogram(device Device) (RowRemapperHistogramValues, error) {
	return device.GetRowRemapperHistogram()
}

func (device nvmlDevice) GetRowRemapperHistogram() (RowRemapperHistogramValues, error) {
	var values RowRemapperHistogramValues
	ret := nvmlDeviceGetRowRemapperHistogram(device, &values)
	return values, ret.error()
}

// nvml.DeviceGetArchitecture()
func (l *library) DeviceGetArchitecture(device Device) (DeviceArchitecture, error) {
	return device.GetArchitecture()
}

func (device nvmlDevice) GetArchitecture() (DeviceArchitecture, error) {
	var arch DeviceArchitecture
	ret := nvmlDeviceGetArchitecture(device, &arch)
	return arch, ret.error()
}

// nvml.DeviceGetVgpuProcessUtilization()
func (l *library) DeviceGetVgpuProcessUtilization(device Device, lastSeenTimestamp uint64) ([]VgpuProcessUtilizationSample, error) {
	return device.GetVgpuProcessUtilization(lastSeenTimestamp)
}

func (device nvmlDevice) GetVgpuProcessUtilization(lastSeenTimestamp uint64) ([]VgpuProcessUtilizationSample, error) {
	var vgpuProcessSamplesCount uint32 = 1 // Will be reduced upon returning
	for {
		utilizationSamples := make([]VgpuProcessUtilizationSample, vgpuProcessSamplesCount)
		ret := nvmlDeviceGetVgpuProcessUtilization(device, lastSeenTimestamp, &vgpuProcessSamplesCount, &utilizationSamples[0])
		if ret == nvmlSUCCESS {
			return utilizationSamples[:vgpuProcessSamplesCount], ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		vgpuProcessSamplesCount *= 2
	}
}

// nvml.GetExcludedDeviceCount()
func (l *library) GetExcludedDeviceCount() (int, error) {
	var deviceCount uint32
	ret := nvmlGetExcludedDeviceCount(&deviceCount)
	return int(deviceCount), ret.error()
}

// nvml.GetExcludedDeviceInfoByIndex()
func (l *library) GetExcludedDeviceInfoByIndex(index int) (ExcludedDeviceInfo, error) {
	var info ExcludedDeviceInfo
	ret := nvmlGetExcludedDeviceInfoByIndex(uint32(index), &info)
	return info, ret.error()
}

// nvml.DeviceSetMigMode()
func (l *library) DeviceSetMigMode(device Device, mode int) (error, error) {
	return device.SetMigMode(mode)
}

func (device nvmlDevice) SetMigMode(mode int) (error, error) {
	var activationStatus Return
	ret := nvmlDeviceSetMigMode(device, uint32(mode), &activationStatus)
	return activationStatus.error(), ret.error()
}

// nvml.DeviceGetMigMode()
func (l *library) DeviceGetMigMode(device Device) (int, int, error) {
	return device.GetMigMode()
}

func (device nvmlDevice) GetMigMode() (int, int, error) {
	var currentMode, pendingMode uint32
	ret := nvmlDeviceGetMigMode(device, &currentMode, &pendingMode)
	return int(currentMode), int(pendingMode), ret.error()
}

// nvml.DeviceGetGpuInstanceProfileInfo()
func (l *library) DeviceGetGpuInstanceProfileInfo(device Device, profile int) (GpuInstanceProfileInfo, error) {
	return device.GetGpuInstanceProfileInfo(profile)
}

func (device nvmlDevice) GetGpuInstanceProfileInfo(profile int) (GpuInstanceProfileInfo, error) {
	var info GpuInstanceProfileInfo
	ret := nvmlDeviceGetGpuInstanceProfileInfo(device, uint32(profile), &info)
	return info, ret.error()
}

// nvml.DeviceGetGpuInstanceProfileInfoV()
type GpuInstanceProfileInfoHandler struct {
	device  nvmlDevice
	profile int
}

func (handler GpuInstanceProfileInfoHandler) V1() (GpuInstanceProfileInfo, error) {
	return DeviceGetGpuInstanceProfileInfo(handler.device, handler.profile)
}

func (handler GpuInstanceProfileInfoHandler) V2() (GpuInstanceProfileInfo_v2, error) {
	var info GpuInstanceProfileInfo_v2
	info.Version = STRUCT_VERSION(info, 2)
	ret := nvmlDeviceGetGpuInstanceProfileInfoV(handler.device, uint32(handler.profile), &info)
	return info, ret.error()
}

func (l *library) DeviceGetGpuInstanceProfileInfoV(device Device, profile int) GpuInstanceProfileInfoHandler {
	return device.GetGpuInstanceProfileInfoV(profile)
}

func (device nvmlDevice) GetGpuInstanceProfileInfoV(profile int) GpuInstanceProfileInfoHandler {
	return GpuInstanceProfileInfoHandler{device, profile}
}

// nvml.DeviceGetGpuInstancePossiblePlacements()
func (l *library) DeviceGetGpuInstancePossiblePlacements(device Device, info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, error) {
	return device.GetGpuInstancePossiblePlacements(info)
}

func (device nvmlDevice) GetGpuInstancePossiblePlacements(info *GpuInstanceProfileInfo) ([]GpuInstancePlacement, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var count uint32
	ret := nvmlDeviceGetGpuInstancePossiblePlacements(device, info.Id, nil, &count)
	if ret != nvmlSUCCESS {
		return nil, ret.error()
	}
	if count == 0 {
		return []GpuInstancePlacement{}, ret.error()
	}
	placements := make([]GpuInstancePlacement, count)
	ret = nvmlDeviceGetGpuInstancePossiblePlacements(device, info.Id, &placements[0], &count)
	return placements[:count], ret.error()
}

// nvml.DeviceGetGpuInstanceRemainingCapacity()
func (l *library) DeviceGetGpuInstanceRemainingCapacity(device Device, info *GpuInstanceProfileInfo) (int, error) {
	return device.GetGpuInstanceRemainingCapacity(info)
}

func (device nvmlDevice) GetGpuInstanceRemainingCapacity(info *GpuInstanceProfileInfo) (int, error) {
	if info == nil {
		return 0, ERROR_INVALID_ARGUMENT
	}
	var count uint32
	ret := nvmlDeviceGetGpuInstanceRemainingCapacity(device, info.Id, &count)
	return int(count), ret.error()
}

// nvml.DeviceCreateGpuInstance()
func (l *library) DeviceCreateGpuInstance(device Device, info *GpuInstanceProfileInfo) (GpuInstance, error) {
	return device.CreateGpuInstance(info)
}

func (device nvmlDevice) CreateGpuInstance(info *GpuInstanceProfileInfo) (GpuInstance, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var gpuInstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstance(device, info.Id, &gpuInstance)
	return gpuInstance, ret.error()
}

// nvml.DeviceCreateGpuInstanceWithPlacement()
func (l *library) DeviceCreateGpuInstanceWithPlacement(device Device, info *GpuInstanceProfileInfo, placement *GpuInstancePlacement) (GpuInstance, error) {
	return device.CreateGpuInstanceWithPlacement(info, placement)
}

func (device nvmlDevice) CreateGpuInstanceWithPlacement(info *GpuInstanceProfileInfo, placement *GpuInstancePlacement) (GpuInstance, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var gpuInstance nvmlGpuInstance
	ret := nvmlDeviceCreateGpuInstanceWithPlacement(device, info.Id, placement, &gpuInstance)
	return gpuInstance, ret.error()
}

// nvml.GpuInstanceDestroy()
func (l *library) GpuInstanceDestroy(gpuInstance GpuInstance) error {
	return gpuInstance.Destroy()
}

func (gpuInstance nvmlGpuInstance) Destroy() error {
	return nvmlGpuInstanceDestroy(gpuInstance).error()
}

// nvml.DeviceGetGpuInstances()
func (l *library) DeviceGetGpuInstances(device Device, info *GpuInstanceProfileInfo) ([]GpuInstance, error) {
	return device.GetGpuInstances(info)
}

func (device nvmlDevice) GetGpuInstances(info *GpuInstanceProfileInfo) ([]GpuInstance, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var count uint32 = info.InstanceCount
	gpuInstances := make([]nvmlGpuInstance, count)
	ret := nvmlDeviceGetGpuInstances(device, info.Id, &gpuInstances[0], &count)
	return convertSlice[nvmlGpuInstance, GpuInstance](gpuInstances[:count]), ret.error()
}

// nvml.DeviceGetGpuInstanceById()
func (l *library) DeviceGetGpuInstanceById(device Device, id int) (GpuInstance, error) {
	return device.GetGpuInstanceById(id)
}

func (device nvmlDevice) GetGpuInstanceById(id int) (GpuInstance, error) {
	var gpuInstance nvmlGpuInstance
	ret := nvmlDeviceGetGpuInstanceById(device, uint32(id), &gpuInstance)
	return gpuInstance, ret.error()
}

// nvml.GpuInstanceGetInfo()
func (l *library) GpuInstanceGetInfo(gpuInstance GpuInstance) (GpuInstanceInfo, error) {
	return gpuInstance.GetInfo()
}

func (gpuInstance nvmlGpuInstance) GetInfo() (GpuInstanceInfo, error) {
	var info nvmlGpuInstanceInfo
	ret := nvmlGpuInstanceGetInfo(gpuInstance, &info)
	return info.convert(), ret.error()
}

// nvml.GpuInstanceGetComputeInstanceProfileInfo()
func (l *library) GpuInstanceGetComputeInstanceProfileInfo(gpuInstance GpuInstance, profile int, engProfile int) (ComputeInstanceProfileInfo, error) {
	return gpuInstance.GetComputeInstanceProfileInfo(profile, engProfile)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstanceProfileInfo(profile int, engProfile int) (ComputeInstanceProfileInfo, error) {
	var info ComputeInstanceProfileInfo
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfo(gpuInstance, uint32(profile), uint32(engProfile), &info)
	return info, ret.error()
}

// nvml.GpuInstanceGetComputeInstanceProfileInfoV()
type ComputeInstanceProfileInfoHandler struct {
	gpuInstance nvmlGpuInstance
	profile     int
	engProfile  int
}

func (handler ComputeInstanceProfileInfoHandler) V1() (ComputeInstanceProfileInfo, error) {
	return GpuInstanceGetComputeInstanceProfileInfo(handler.gpuInstance, handler.profile, handler.engProfile)
}

func (handler ComputeInstanceProfileInfoHandler) V2() (ComputeInstanceProfileInfo_v2, error) {
	var info ComputeInstanceProfileInfo_v2
	info.Version = STRUCT_VERSION(info, 2)
	ret := nvmlGpuInstanceGetComputeInstanceProfileInfoV(handler.gpuInstance, uint32(handler.profile), uint32(handler.engProfile), &info)
	return info, ret.error()
}

func (l *library) GpuInstanceGetComputeInstanceProfileInfoV(gpuInstance GpuInstance, profile int, engProfile int) ComputeInstanceProfileInfoHandler {
	return gpuInstance.GetComputeInstanceProfileInfoV(profile, engProfile)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstanceProfileInfoV(profile int, engProfile int) ComputeInstanceProfileInfoHandler {
	return ComputeInstanceProfileInfoHandler{gpuInstance, profile, engProfile}
}

// nvml.GpuInstanceGetComputeInstanceRemainingCapacity()
func (l *library) GpuInstanceGetComputeInstanceRemainingCapacity(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) (int, error) {
	return gpuInstance.GetComputeInstanceRemainingCapacity(info)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstanceRemainingCapacity(info *ComputeInstanceProfileInfo) (int, error) {
	if info == nil {
		return 0, ERROR_INVALID_ARGUMENT
	}
	var count uint32
	ret := nvmlGpuInstanceGetComputeInstanceRemainingCapacity(gpuInstance, info.Id, &count)
	return int(count), ret.error()
}

// nvml.GpuInstanceCreateComputeInstance()
func (l *library) GpuInstanceCreateComputeInstance(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) (ComputeInstance, error) {
	return gpuInstance.CreateComputeInstance(info)
}

func (gpuInstance nvmlGpuInstance) CreateComputeInstance(info *ComputeInstanceProfileInfo) (ComputeInstance, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var computeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstance(gpuInstance, info.Id, &computeInstance)
	return computeInstance, ret.error()
}

// nvml.ComputeInstanceDestroy()
func (l *library) ComputeInstanceDestroy(computeInstance ComputeInstance) error {
	return computeInstance.Destroy()
}

func (computeInstance nvmlComputeInstance) Destroy() error {
	return nvmlComputeInstanceDestroy(computeInstance).error()
}

// nvml.GpuInstanceGetComputeInstances()
func (l *library) GpuInstanceGetComputeInstances(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) ([]ComputeInstance, error) {
	return gpuInstance.GetComputeInstances(info)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstances(info *ComputeInstanceProfileInfo) ([]ComputeInstance, error) {
	if info == nil {
		return nil, ERROR_INVALID_ARGUMENT
	}
	var count uint32 = info.InstanceCount
	computeInstances := make([]nvmlComputeInstance, count)
	ret := nvmlGpuInstanceGetComputeInstances(gpuInstance, info.Id, &computeInstances[0], &count)
	return convertSlice[nvmlComputeInstance, ComputeInstance](computeInstances[:count]), ret.error()
}

// nvml.GpuInstanceGetComputeInstanceById()
func (l *library) GpuInstanceGetComputeInstanceById(gpuInstance GpuInstance, id int) (ComputeInstance, error) {
	return gpuInstance.GetComputeInstanceById(id)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstanceById(id int) (ComputeInstance, error) {
	var computeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceGetComputeInstanceById(gpuInstance, uint32(id), &computeInstance)
	return computeInstance, ret.error()
}

// nvml.ComputeInstanceGetInfo()
func (l *library) ComputeInstanceGetInfo(computeInstance ComputeInstance) (ComputeInstanceInfo, error) {
	return computeInstance.GetInfo()
}

func (computeInstance nvmlComputeInstance) GetInfo() (ComputeInstanceInfo, error) {
	var info nvmlComputeInstanceInfo
	ret := nvmlComputeInstanceGetInfo(computeInstance, &info)
	return info.convert(), ret.error()
}

// nvml.DeviceIsMigDeviceHandle()
func (l *library) DeviceIsMigDeviceHandle(device Device) (bool, error) {
	return device.IsMigDeviceHandle()
}

func (device nvmlDevice) IsMigDeviceHandle() (bool, error) {
	var isMigDevice uint32
	ret := nvmlDeviceIsMigDeviceHandle(device, &isMigDevice)
	return (isMigDevice != 0), ret.error()
}

// nvml DeviceGetGpuInstanceId()
func (l *library) DeviceGetGpuInstanceId(device Device) (int, error) {
	return device.GetGpuInstanceId()
}

func (device nvmlDevice) GetGpuInstanceId() (int, error) {
	var id uint32
	ret := nvmlDeviceGetGpuInstanceId(device, &id)
	return int(id), ret.error()
}

// nvml.DeviceGetComputeInstanceId()
func (l *library) DeviceGetComputeInstanceId(device Device) (int, error) {
	return device.GetComputeInstanceId()
}

func (device nvmlDevice) GetComputeInstanceId() (int, error) {
	var id uint32
	ret := nvmlDeviceGetComputeInstanceId(device, &id)
	return int(id), ret.error()
}

// nvml.DeviceGetMaxMigDeviceCount()
func (l *library) DeviceGetMaxMigDeviceCount(device Device) (int, error) {
	return device.GetMaxMigDeviceCount()
}

func (device nvmlDevice) GetMaxMigDeviceCount() (int, error) {
	var count uint32
	ret := nvmlDeviceGetMaxMigDeviceCount(device, &count)
	return int(count), ret.error()
}

// nvml.DeviceGetMigDeviceHandleByIndex()
func (l *library) DeviceGetMigDeviceHandleByIndex(device Device, index int) (Device, error) {
	return device.GetMigDeviceHandleByIndex(index)
}

func (device nvmlDevice) GetMigDeviceHandleByIndex(index int) (Device, error) {
	var migDevice nvmlDevice
	ret := nvmlDeviceGetMigDeviceHandleByIndex(device, uint32(index), &migDevice)
	return migDevice, ret.error()
}

// nvml.DeviceGetDeviceHandleFromMigDeviceHandle()
func (l *library) DeviceGetDeviceHandleFromMigDeviceHandle(migdevice Device) (Device, error) {
	return migdevice.GetDeviceHandleFromMigDeviceHandle()
}

func (migDevice nvmlDevice) GetDeviceHandleFromMigDeviceHandle() (Device, error) {
	var device nvmlDevice
	ret := nvmlDeviceGetDeviceHandleFromMigDeviceHandle(migDevice, &device)
	return device, ret.error()
}

// nvml.DeviceGetBusType()
func (l *library) DeviceGetBusType(device Device) (BusType, error) {
	return device.GetBusType()
}

func (device nvmlDevice) GetBusType() (BusType, error) {
	var busType BusType
	ret := nvmlDeviceGetBusType(device, &busType)
	return busType, ret.error()
}

// nvml.DeviceSetDefaultFanSpeed_v2()
func (l *library) DeviceSetDefaultFanSpeed_v2(device Device, fan int) error {
	return device.SetDefaultFanSpeed_v2(fan)
}

func (device nvmlDevice) SetDefaultFanSpeed_v2(fan int) error {
	return nvmlDeviceSetDefaultFanSpeed_v2(device, uint32(fan)).error()
}

// nvml.DeviceGetMinMaxFanSpeed()
func (l *library) DeviceGetMinMaxFanSpeed(device Device) (int, int, error) {
	return device.GetMinMaxFanSpeed()
}

func (device nvmlDevice) GetMinMaxFanSpeed() (int, int, error) {
	var minSpeed, maxSpeed uint32
	ret := nvmlDeviceGetMinMaxFanSpeed(device, &minSpeed, &maxSpeed)
	return int(minSpeed), int(maxSpeed), ret.error()
}

// nvml.DeviceGetThermalSettings()
func (l *library) DeviceGetThermalSettings(device Device, sensorIndex uint32) (GpuThermalSettings, error) {
	return device.GetThermalSettings(sensorIndex)
}

func (device nvmlDevice) GetThermalSettings(sensorIndex uint32) (GpuThermalSettings, error) {
	var pThermalSettings GpuThermalSettings
	ret := nvmlDeviceGetThermalSettings(device, sensorIndex, &pThermalSettings)
	return pThermalSettings, ret.error()
}

// nvml.DeviceGetDefaultEccMode()
func (l *library) DeviceGetDefaultEccMode(device Device) (EnableState, error) {
	return device.GetDefaultEccMode()
}

func (device nvmlDevice) GetDefaultEccMode() (EnableState, error) {
	var defaultMode EnableState
	ret := nvmlDeviceGetDefaultEccMode(device, &defaultMode)
	return defaultMode, ret.error()
}

// nvml.DeviceGetPcieSpeed()
func (l *library) DeviceGetPcieSpeed(device Device) (int, error) {
	return device.GetPcieSpeed()
}

func (device nvmlDevice) GetPcieSpeed() (int, error) {
	var pcieSpeed uint32
	ret := nvmlDeviceGetPcieSpeed(device, &pcieSpeed)
	return int(pcieSpeed), ret.error()
}

// nvml.DeviceGetGspFirmwareVersion()
func (l *library) DeviceGetGspFirmwareVersion(device Device) (string, error) {
	return device.GetGspFirmwareVersion()
}

func (device nvmlDevice) GetGspFirmwareVersion() (string, error) {
	version := make([]byte, GSP_FIRMWARE_VERSION_BUF_SIZE)
	ret := nvmlDeviceGetGspFirmwareVersion(device, &version[0])
	return string(version[:clen(version)]), ret.error()
}

// nvml.DeviceGetGspFirmwareMode()
func (l *library) DeviceGetGspFirmwareMode(device Device) (bool, bool, error) {
	return device.GetGspFirmwareMode()
}

func (device nvmlDevice) GetGspFirmwareMode() (bool, bool, error) {
	var isEnabled, defaultMode uint32
	ret := nvmlDeviceGetGspFirmwareMode(device, &isEnabled, &defaultMode)
	return (isEnabled != 0), (defaultMode != 0), ret.error()
}

// nvml.DeviceGetDynamicPstatesInfo()
func (l *library) DeviceGetDynamicPstatesInfo(device Device) (GpuDynamicPstatesInfo, error) {
	return device.GetDynamicPstatesInfo()
}

func (device nvmlDevice) GetDynamicPstatesInfo() (GpuDynamicPstatesInfo, error) {
	var pDynamicPstatesInfo GpuDynamicPstatesInfo
	ret := nvmlDeviceGetDynamicPstatesInfo(device, &pDynamicPstatesInfo)
	return pDynamicPstatesInfo, ret.error()
}

// nvml.DeviceSetFanSpeed_v2()
func (l *library) DeviceSetFanSpeed_v2(device Device, fan int, speed int) error {
	return device.SetFanSpeed_v2(fan, speed)
}

func (device nvmlDevice) SetFanSpeed_v2(fan int, speed int) error {
	return nvmlDeviceSetFanSpeed_v2(device, uint32(fan), uint32(speed)).error()
}

// nvml.DeviceGetGpcClkVfOffset()
func (l *library) DeviceGetGpcClkVfOffset(device Device) (int, error) {
	return device.GetGpcClkVfOffset()
}

func (device nvmlDevice) GetGpcClkVfOffset() (int, error) {
	var offset int32
	ret := nvmlDeviceGetGpcClkVfOffset(device, &offset)
	return int(offset), ret.error()
}

// nvml.DeviceSetGpcClkVfOffset()
func (l *library) DeviceSetGpcClkVfOffset(device Device, offset int) error {
	return device.SetGpcClkVfOffset(offset)
}

func (device nvmlDevice) SetGpcClkVfOffset(offset int) error {
	return nvmlDeviceSetGpcClkVfOffset(device, int32(offset)).error()
}

// nvml.DeviceGetMinMaxClockOfPState()
func (l *library) DeviceGetMinMaxClockOfPState(device Device, clockType ClockType, pstate Pstates) (uint32, uint32, error) {
	return device.GetMinMaxClockOfPState(clockType, pstate)
}

func (device nvmlDevice) GetMinMaxClockOfPState(clockType ClockType, pstate Pstates) (uint32, uint32, error) {
	var minClockMHz, maxClockMHz uint32
	ret := nvmlDeviceGetMinMaxClockOfPState(device, clockType, pstate, &minClockMHz, &maxClockMHz)
	return minClockMHz, maxClockMHz, ret.error()
}

// nvml.DeviceGetSupportedPerformanceStates()
func (l *library) DeviceGetSupportedPerformanceStates(device Device) ([]Pstates, error) {
	return device.GetSupportedPerformanceStates()
}

func (device nvmlDevice) GetSupportedPerformanceStates() ([]Pstates, error) {
	pstates := make([]Pstates, MAX_GPU_PERF_PSTATES)
	ret := nvmlDeviceGetSupportedPerformanceStates(device, &pstates[0], MAX_GPU_PERF_PSTATES)
	for i := 0; i < MAX_GPU_PERF_PSTATES; i++ {
		if pstates[i] == PSTATE_UNKNOWN {
			return pstates[0:i], ret.error()
		}
	}
	return pstates, ret.error()
}

// nvml.DeviceGetTargetFanSpeed()
func (l *library) DeviceGetTargetFanSpeed(device Device, fan int) (int, error) {
	return device.GetTargetFanSpeed(fan)
}

func (device nvmlDevice) GetTargetFanSpeed(fan int) (int, error) {
	var targetSpeed uint32
	ret := nvmlDeviceGetTargetFanSpeed(device, uint32(fan), &targetSpeed)
	return int(targetSpeed), ret.error()
}

// nvml.DeviceGetMemClkVfOffset()
func (l *library) DeviceGetMemClkVfOffset(device Device) (int, error) {
	return device.GetMemClkVfOffset()
}

func (device nvmlDevice) GetMemClkVfOffset() (int, error) {
	var offset int32
	ret := nvmlDeviceGetMemClkVfOffset(device, &offset)
	return int(offset), ret.error()
}

// nvml.DeviceSetMemClkVfOffset()
func (l *library) DeviceSetMemClkVfOffset(device Device, offset int) error {
	return device.SetMemClkVfOffset(offset)
}

func (device nvmlDevice) SetMemClkVfOffset(offset int) error {
	return nvmlDeviceSetMemClkVfOffset(device, int32(offset)).error()
}

// nvml.DeviceGetGpcClkMinMaxVfOffset()
func (l *library) DeviceGetGpcClkMinMaxVfOffset(device Device) (int, int, error) {
	return device.GetGpcClkMinMaxVfOffset()
}

func (device nvmlDevice) GetGpcClkMinMaxVfOffset() (int, int, error) {
	var minOffset, maxOffset int32
	ret := nvmlDeviceGetGpcClkMinMaxVfOffset(device, &minOffset, &maxOffset)
	return int(minOffset), int(maxOffset), ret.error()
}

// nvml.DeviceGetMemClkMinMaxVfOffset()
func (l *library) DeviceGetMemClkMinMaxVfOffset(device Device) (int, int, error) {
	return device.GetMemClkMinMaxVfOffset()
}

func (device nvmlDevice) GetMemClkMinMaxVfOffset() (int, int, error) {
	var minOffset, maxOffset int32
	ret := nvmlDeviceGetMemClkMinMaxVfOffset(device, &minOffset, &maxOffset)
	return int(minOffset), int(maxOffset), ret.error()
}

// nvml.DeviceGetGpuMaxPcieLinkGeneration()
func (l *library) DeviceGetGpuMaxPcieLinkGeneration(device Device) (int, error) {
	return device.GetGpuMaxPcieLinkGeneration()
}

func (device nvmlDevice) GetGpuMaxPcieLinkGeneration() (int, error) {
	var maxLinkGenDevice uint32
	ret := nvmlDeviceGetGpuMaxPcieLinkGeneration(device, &maxLinkGenDevice)
	return int(maxLinkGenDevice), ret.error()
}

// nvml.DeviceGetFanControlPolicy_v2()
func (l *library) DeviceGetFanControlPolicy_v2(device Device, fan int) (FanControlPolicy, error) {
	return device.GetFanControlPolicy_v2(fan)
}

func (device nvmlDevice) GetFanControlPolicy_v2(fan int) (FanControlPolicy, error) {
	var policy FanControlPolicy
	ret := nvmlDeviceGetFanControlPolicy_v2(device, uint32(fan), &policy)
	return policy, ret.error()
}

// nvml.DeviceSetFanControlPolicy()
func (l *library) DeviceSetFanControlPolicy(device Device, fan int, policy FanControlPolicy) error {
	return device.SetFanControlPolicy(fan, policy)
}

func (device nvmlDevice) SetFanControlPolicy(fan int, policy FanControlPolicy) error {
	return nvmlDeviceSetFanControlPolicy(device, uint32(fan), policy).error()
}

// nvml.DeviceClearFieldValues()
func (l *library) DeviceClearFieldValues(device Device, values []FieldValue) error {
	return device.ClearFieldValues(values)
}

func (device nvmlDevice) ClearFieldValues(values []FieldValue) error {
	valuesCount := len(values)
	return nvmlDeviceClearFieldValues(device, int32(valuesCount), &values[0]).error()
}

// nvml.DeviceGetVgpuCapabilities()
func (l *library) DeviceGetVgpuCapabilities(device Device, capability DeviceVgpuCapability) (bool, error) {
	return device.GetVgpuCapabilities(capability)
}

func (device nvmlDevice) GetVgpuCapabilities(capability DeviceVgpuCapability) (bool, error) {
	var capResult uint32
	ret := nvmlDeviceGetVgpuCapabilities(device, capability, &capResult)
	return (capResult != 0), ret.error()
}

// nvml.DeviceGetVgpuSchedulerLog()
func (l *library) DeviceGetVgpuSchedulerLog(device Device) (VgpuSchedulerLog, error) {
	return device.GetVgpuSchedulerLog()
}

func (device nvmlDevice) GetVgpuSchedulerLog() (VgpuSchedulerLog, error) {
	var pSchedulerLog VgpuSchedulerLog
	ret := nvmlDeviceGetVgpuSchedulerLog(device, &pSchedulerLog)
	return pSchedulerLog, ret.error()
}

// nvml.DeviceGetVgpuSchedulerState()
func (l *library) DeviceGetVgpuSchedulerState(device Device) (VgpuSchedulerGetState, error) {
	return device.GetVgpuSchedulerState()
}

func (device nvmlDevice) GetVgpuSchedulerState() (VgpuSchedulerGetState, error) {
	var pSchedulerState VgpuSchedulerGetState
	ret := nvmlDeviceGetVgpuSchedulerState(device, &pSchedulerState)
	return pSchedulerState, ret.error()
}

// nvml.DeviceSetVgpuSchedulerState()
func (l *library) DeviceSetVgpuSchedulerState(device Device, pSchedulerState *VgpuSchedulerSetState) error {
	return device.SetVgpuSchedulerState(pSchedulerState)
}

func (device nvmlDevice) SetVgpuSchedulerState(pSchedulerState *VgpuSchedulerSetState) error {
	return nvmlDeviceSetVgpuSchedulerState(device, pSchedulerState).error()
}

// nvml.DeviceGetVgpuSchedulerCapabilities()
func (l *library) DeviceGetVgpuSchedulerCapabilities(device Device) (VgpuSchedulerCapabilities, error) {
	return device.GetVgpuSchedulerCapabilities()
}

func (device nvmlDevice) GetVgpuSchedulerCapabilities() (VgpuSchedulerCapabilities, error) {
	var pCapabilities VgpuSchedulerCapabilities
	ret := nvmlDeviceGetVgpuSchedulerCapabilities(device, &pCapabilities)
	return pCapabilities, ret.error()
}

// nvml.GpuInstanceGetComputeInstancePossiblePlacements()
func (l *library) GpuInstanceGetComputeInstancePossiblePlacements(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, error) {
	return gpuInstance.GetComputeInstancePossiblePlacements(info)
}

func (gpuInstance nvmlGpuInstance) GetComputeInstancePossiblePlacements(info *ComputeInstanceProfileInfo) ([]ComputeInstancePlacement, error) {
	var count uint32
	ret := nvmlGpuInstanceGetComputeInstancePossiblePlacements(gpuInstance, info.Id, nil, &count)
	if ret != nvmlSUCCESS {
		return nil, ret.error()
	}
	if count == 0 {
		return []ComputeInstancePlacement{}, ret.error()
	}
	placementArray := make([]ComputeInstancePlacement, count)
	ret = nvmlGpuInstanceGetComputeInstancePossiblePlacements(gpuInstance, info.Id, &placementArray[0], &count)
	return placementArray, ret.error()
}

// nvml.GpuInstanceCreateComputeInstanceWithPlacement()
func (l *library) GpuInstanceCreateComputeInstanceWithPlacement(gpuInstance GpuInstance, info *ComputeInstanceProfileInfo, placement *ComputeInstancePlacement) (ComputeInstance, error) {
	return gpuInstance.CreateComputeInstanceWithPlacement(info, placement)
}

func (gpuInstance nvmlGpuInstance) CreateComputeInstanceWithPlacement(info *ComputeInstanceProfileInfo, placement *ComputeInstancePlacement) (ComputeInstance, error) {
	var computeInstance nvmlComputeInstance
	ret := nvmlGpuInstanceCreateComputeInstanceWithPlacement(gpuInstance, info.Id, placement, &computeInstance)
	return computeInstance, ret.error()
}

// nvml.DeviceGetGpuFabricInfo()
func (l *library) DeviceGetGpuFabricInfo(device Device) (GpuFabricInfo, error) {
	return device.GetGpuFabricInfo()
}

func (device nvmlDevice) GetGpuFabricInfo() (GpuFabricInfo, error) {
	var gpuFabricInfo GpuFabricInfo
	ret := nvmlDeviceGetGpuFabricInfo(device, &gpuFabricInfo)
	return gpuFabricInfo, ret.error()
}

// nvml.DeviceSetNvLinkDeviceLowPowerThreshold()
func (l *library) DeviceSetNvLinkDeviceLowPowerThreshold(device Device, info *NvLinkPowerThres) error {
	return device.SetNvLinkDeviceLowPowerThreshold(info)
}

func (device nvmlDevice) SetNvLinkDeviceLowPowerThreshold(info *NvLinkPowerThres) error {
	return nvmlDeviceSetNvLinkDeviceLowPowerThreshold(device, info).error()
}

// nvml.DeviceGetModuleId()
func (l *library) DeviceGetModuleId(device Device) (int, error) {
	return device.GetModuleId()
}

func (device nvmlDevice) GetModuleId() (int, error) {
	var moduleID uint32
	ret := nvmlDeviceGetModuleId(device, &moduleID)
	return int(moduleID), ret.error()
}

// nvml.DeviceGetCurrentClocksEventReasons()
func (l *library) DeviceGetCurrentClocksEventReasons(device Device) (uint64, error) {
	return device.GetCurrentClocksEventReasons()
}

func (device nvmlDevice) GetCurrentClocksEventReasons() (uint64, error) {
	var clocksEventReasons uint64
	ret := nvmlDeviceGetCurrentClocksEventReasons(device, &clocksEventReasons)
	return clocksEventReasons, ret.error()
}

// nvml.DeviceGetSupportedClocksEventReasons()
func (l *library) DeviceGetSupportedClocksEventReasons(device Device) (uint64, error) {
	return device.GetSupportedClocksEventReasons()
}

func (device nvmlDevice) GetSupportedClocksEventReasons() (uint64, error) {
	var supportedClocksEventReasons uint64
	ret := nvmlDeviceGetSupportedClocksEventReasons(device, &supportedClocksEventReasons)
	return supportedClocksEventReasons, ret.error()
}

// nvml.DeviceGetJpgUtilization()
func (l *library) DeviceGetJpgUtilization(device Device) (uint32, uint32, error) {
	return device.GetJpgUtilization()
}

func (device nvmlDevice) GetJpgUtilization() (uint32, uint32, error) {
	var utilization, samplingPeriodUs uint32
	ret := nvmlDeviceGetJpgUtilization(device, &utilization, &samplingPeriodUs)
	return utilization, samplingPeriodUs, ret.error()
}

// nvml.DeviceGetOfaUtilization()
func (l *library) DeviceGetOfaUtilization(device Device) (uint32, uint32, error) {
	return device.GetOfaUtilization()
}

func (device nvmlDevice) GetOfaUtilization() (uint32, uint32, error) {
	var utilization, samplingPeriodUs uint32
	ret := nvmlDeviceGetOfaUtilization(device, &utilization, &samplingPeriodUs)
	return utilization, samplingPeriodUs, ret.error()
}

// nvml.DeviceGetRunningProcessDetailList()
func (l *library) DeviceGetRunningProcessDetailList(device Device) (ProcessDetailList, error) {
	return device.GetRunningProcessDetailList()
}

func (device nvmlDevice) GetRunningProcessDetailList() (ProcessDetailList, error) {
	var plist ProcessDetailList
	ret := nvmlDeviceGetRunningProcessDetailList(device, &plist)
	return plist, ret.error()
}

// nvml.DeviceGetConfComputeMemSizeInfo()
func (l *library) DeviceGetConfComputeMemSizeInfo(device Device) (ConfComputeMemSizeInfo, error) {
	return device.GetConfComputeMemSizeInfo()
}

func (device nvmlDevice) GetConfComputeMemSizeInfo() (ConfComputeMemSizeInfo, error) {
	var memInfo ConfComputeMemSizeInfo
	ret := nvmlDeviceGetConfComputeMemSizeInfo(device, &memInfo)
	return memInfo, ret.error()
}

// nvml.DeviceGetConfComputeProtectedMemoryUsage()
func (l *library) DeviceGetConfComputeProtectedMemoryUsage(device Device) (Memory, error) {
	return device.GetConfComputeProtectedMemoryUsage()
}

func (device nvmlDevice) GetConfComputeProtectedMemoryUsage() (Memory, error) {
	var memory Memory
	ret := nvmlDeviceGetConfComputeProtectedMemoryUsage(device, &memory)
	return memory, ret.error()
}

// nvml.DeviceGetConfComputeGpuCertificate()
func (l *library) DeviceGetConfComputeGpuCertificate(device Device) (ConfComputeGpuCertificate, error) {
	return device.GetConfComputeGpuCertificate()
}

func (device nvmlDevice) GetConfComputeGpuCertificate() (ConfComputeGpuCertificate, error) {
	var gpuCert ConfComputeGpuCertificate
	ret := nvmlDeviceGetConfComputeGpuCertificate(device, &gpuCert)
	return gpuCert, ret.error()
}

// nvml.DeviceGetConfComputeGpuAttestationReport()
func (l *library) DeviceGetConfComputeGpuAttestationReport(device Device) (ConfComputeGpuAttestationReport, error) {
	return device.GetConfComputeGpuAttestationReport()
}

func (device nvmlDevice) GetConfComputeGpuAttestationReport() (ConfComputeGpuAttestationReport, error) {
	var gpuAtstReport ConfComputeGpuAttestationReport
	ret := nvmlDeviceGetConfComputeGpuAttestationReport(device, &gpuAtstReport)
	return gpuAtstReport, ret.error()
}

// nvml.DeviceSetConfComputeUnprotectedMemSize()
func (l *library) DeviceSetConfComputeUnprotectedMemSize(device Device, sizeKiB uint64) error {
	return device.SetConfComputeUnprotectedMemSize(sizeKiB)
}

func (device nvmlDevice) SetConfComputeUnprotectedMemSize(sizeKiB uint64) error {
	return nvmlDeviceSetConfComputeUnprotectedMemSize(device, sizeKiB).error()
}

// nvml.DeviceSetPowerManagementLimit_v2()
func (l *library) DeviceSetPowerManagementLimit_v2(device Device, powerValue *PowerValue_v2) error {
	return device.SetPowerManagementLimit_v2(powerValue)
}

func (device nvmlDevice) SetPowerManagementLimit_v2(powerValue *PowerValue_v2) error {
	return nvmlDeviceSetPowerManagementLimit_v2(device, powerValue).error()
}

// nvml.DeviceGetC2cModeInfoV()
type C2cModeInfoHandler struct {
	device nvmlDevice
}

func (handler C2cModeInfoHandler) V1() (C2cModeInfo_v1, error) {
	var c2cModeInfo C2cModeInfo_v1
	ret := nvmlDeviceGetC2cModeInfoV(handler.device, &c2cModeInfo)
	return c2cModeInfo, ret.error()
}

func (l *library) DeviceGetC2cModeInfoV(device Device) C2cModeInfoHandler {
	return device.GetC2cModeInfoV()
}

func (device nvmlDevice) GetC2cModeInfoV() C2cModeInfoHandler {
	return C2cModeInfoHandler{device}
}

// nvml.DeviceGetLastBBXFlushTime()
func (l *library) DeviceGetLastBBXFlushTime(device Device) (uint64, uint, error) {
	return device.GetLastBBXFlushTime()
}

func (device nvmlDevice) GetLastBBXFlushTime() (uint64, uint, error) {
	var timestamp uint64
	var durationUs uint
	ret := nvmlDeviceGetLastBBXFlushTime(device, &timestamp, &durationUs)
	return timestamp, durationUs, ret.error()
}

// nvml.DeviceGetNumaNodeId()
func (l *library) DeviceGetNumaNodeId(device Device) (int, error) {
	return device.GetNumaNodeId()
}

func (device nvmlDevice) GetNumaNodeId() (int, error) {
	var node uint32
	ret := nvmlDeviceGetNumaNodeId(device, &node)
	return int(node), ret.error()
}

// nvml.DeviceGetPciInfoExt()
func (l *library) DeviceGetPciInfoExt(device Device) (PciInfoExt, error) {
	return device.GetPciInfoExt()
}

func (device nvmlDevice) GetPciInfoExt() (PciInfoExt, error) {
	var pciInfo PciInfoExt
	ret := nvmlDeviceGetPciInfoExt(device, &pciInfo)
	return pciInfo, ret.error()
}

// nvml.DeviceGetGpuFabricInfoV()
type GpuFabricInfoHandler struct {
	device nvmlDevice
}

func (handler GpuFabricInfoHandler) V1() (GpuFabricInfo, error) {
	return handler.device.GetGpuFabricInfo()
}

func (handler GpuFabricInfoHandler) V2() (GpuFabricInfo_v2, error) {
	var info GpuFabricInfoV
	info.Version = STRUCT_VERSION(info, 2)
	ret := nvmlDeviceGetGpuFabricInfoV(handler.device, &info)
	return GpuFabricInfo_v2(info), ret.error()
}

func (l *library) DeviceGetGpuFabricInfoV(device Device) GpuFabricInfoHandler {
	return device.GetGpuFabricInfoV()
}

func (device nvmlDevice) GetGpuFabricInfoV() GpuFabricInfoHandler {
	return GpuFabricInfoHandler{device}
}

// nvml.DeviceGetProcessesUtilizationInfo()
func (l *library) DeviceGetProcessesUtilizationInfo(device Device) (ProcessesUtilizationInfo, error) {
	return device.GetProcessesUtilizationInfo()
}

func (device nvmlDevice) GetProcessesUtilizationInfo() (ProcessesUtilizationInfo, error) {
	var processesUtilInfo ProcessesUtilizationInfo
	ret := nvmlDeviceGetProcessesUtilizationInfo(device, &processesUtilInfo)
	return processesUtilInfo, ret.error()
}

// nvml.DeviceGetVgpuHeterogeneousMode()
func (l *library) DeviceGetVgpuHeterogeneousMode(device Device) (VgpuHeterogeneousMode, error) {
	return device.GetVgpuHeterogeneousMode()
}

func (device nvmlDevice) GetVgpuHeterogeneousMode() (VgpuHeterogeneousMode, error) {
	var heterogeneousMode VgpuHeterogeneousMode
	ret := nvmlDeviceGetVgpuHeterogeneousMode(device, &heterogeneousMode)
	return heterogeneousMode, ret.error()
}

// nvml.DeviceSetVgpuHeterogeneousMode()
func (l *library) DeviceSetVgpuHeterogeneousMode(device Device, heterogeneousMode VgpuHeterogeneousMode) error {
	return device.SetVgpuHeterogeneousMode(heterogeneousMode)
}

func (device nvmlDevice) SetVgpuHeterogeneousMode(heterogeneousMode VgpuHeterogeneousMode) error {
	ret := nvmlDeviceSetVgpuHeterogeneousMode(device, &heterogeneousMode)
	return ret.error()
}

// nvml.DeviceGetVgpuTypeSupportedPlacements()
func (l *library) DeviceGetVgpuTypeSupportedPlacements(device Device, vgpuTypeId VgpuTypeId) (VgpuPlacementList, error) {
	return device.GetVgpuTypeSupportedPlacements(vgpuTypeId)
}

func (device nvmlDevice) GetVgpuTypeSupportedPlacements(vgpuTypeId VgpuTypeId) (VgpuPlacementList, error) {
	return vgpuTypeId.GetSupportedPlacements(device)
}

func (vgpuTypeId nvmlVgpuTypeId) GetSupportedPlacements(device Device) (VgpuPlacementList, error) {
	var placementList VgpuPlacementList
	ret := nvmlDeviceGetVgpuTypeSupportedPlacements(nvmlDeviceHandle(device), vgpuTypeId, &placementList)
	return placementList, ret.error()
}

// nvml.DeviceGetVgpuTypeCreatablePlacements()
func (l *library) DeviceGetVgpuTypeCreatablePlacements(device Device, vgpuTypeId VgpuTypeId) (VgpuPlacementList, error) {
	return device.GetVgpuTypeCreatablePlacements(vgpuTypeId)
}

func (device nvmlDevice) GetVgpuTypeCreatablePlacements(vgpuTypeId VgpuTypeId) (VgpuPlacementList, error) {
	return vgpuTypeId.GetCreatablePlacements(device)
}

func (vgpuTypeId nvmlVgpuTypeId) GetCreatablePlacements(device Device) (VgpuPlacementList, error) {
	var placementList VgpuPlacementList
	ret := nvmlDeviceGetVgpuTypeCreatablePlacements(nvmlDeviceHandle(device), vgpuTypeId, &placementList)
	return placementList, ret.error()
}

// nvml.DeviceSetVgpuCapabilities()
func (l *library) DeviceSetVgpuCapabilities(device Device, capability DeviceVgpuCapability, state EnableState) error {
	return device.SetVgpuCapabilities(capability, state)
}

func (device nvmlDevice) SetVgpuCapabilities(capability DeviceVgpuCapability, state EnableState) error {
	ret := nvmlDeviceSetVgpuCapabilities(device, capability, state)
	return ret.error()
}

// nvml.DeviceGetVgpuInstancesUtilizationInfo()
func (l *library) DeviceGetVgpuInstancesUtilizationInfo(device Device) (VgpuInstancesUtilizationInfo, error) {
	return device.GetVgpuInstancesUtilizationInfo()
}

func (device nvmlDevice) GetVgpuInstancesUtilizationInfo() (VgpuInstancesUtilizationInfo, error) {
	var vgpuUtilInfo VgpuInstancesUtilizationInfo
	ret := nvmlDeviceGetVgpuInstancesUtilizationInfo(device, &vgpuUtilInfo)
	return vgpuUtilInfo, ret.error()
}

// nvml.DeviceGetVgpuProcessesUtilizationInfo()
func (l *library) DeviceGetVgpuProcessesUtilizationInfo(device Device) (VgpuProcessesUtilizationInfo, error) {
	return device.GetVgpuProcessesUtilizationInfo()
}

func (device nvmlDevice) GetVgpuProcessesUtilizationInfo() (VgpuProcessesUtilizationInfo, error) {
	var vgpuProcUtilInfo VgpuProcessesUtilizationInfo
	ret := nvmlDeviceGetVgpuProcessesUtilizationInfo(device, &vgpuProcUtilInfo)
	return vgpuProcUtilInfo, ret.error()
}

// nvml.DeviceGetSramEccErrorStatus()
func (l *library) DeviceGetSramEccErrorStatus(device Device) (EccSramErrorStatus, error) {
	return device.GetSramEccErrorStatus()
}

func (device nvmlDevice) GetSramEccErrorStatus() (EccSramErrorStatus, error) {
	var status EccSramErrorStatus
	ret := nvmlDeviceGetSramEccErrorStatus(device, &status)
	return status, ret.error()
}
