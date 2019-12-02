// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package nvml

// nvml.UnitGetCount()
func UnitGetCount() (int, Return) {
	var UnitCount uint32
	ret := nvmlUnitGetCount(&UnitCount)
	return int(UnitCount), ret
}

// nvml.UnitGetHandleByIndex()
func UnitGetHandleByIndex(Index int) (Unit, Return) {
	var Unit Unit
	ret := nvmlUnitGetHandleByIndex(uint32(Index), &Unit)
	return Unit, ret
}

// nvml.UnitGetUnitInfo()
func UnitGetUnitInfo(Unit Unit) (UnitInfo, Return) {
	var Info UnitInfo
	ret := nvmlUnitGetUnitInfo(Unit, &Info)
	return Info, ret
}

func (Unit Unit) GetUnitInfo() (UnitInfo, Return) {
	return UnitGetUnitInfo(Unit)
}

// nvml.UnitGetLedState()
func UnitGetLedState(Unit Unit) (LedState, Return) {
	var State LedState
	ret := nvmlUnitGetLedState(Unit, &State)
	return State, ret
}

func (Unit Unit) GetLedState() (LedState, Return) {
	return UnitGetLedState(Unit)
}

// nvml.UnitGetPsuInfo()
func UnitGetPsuInfo(Unit Unit) (PSUInfo, Return) {
	var Psu PSUInfo
	ret := nvmlUnitGetPsuInfo(Unit, &Psu)
	return Psu, ret
}

func (Unit Unit) GetPsuInfo() (PSUInfo, Return) {
	return UnitGetPsuInfo(Unit)
}

// nvml.UnitGetTemperature()
func UnitGetTemperature(Unit Unit, Type int) (uint32, Return) {
	var Temp uint32
	ret := nvmlUnitGetTemperature(Unit, uint32(Type), &Temp)
	return Temp, ret
}

func (Unit Unit) GetTemperature(Type int) (uint32, Return) {
	return UnitGetTemperature(Unit, Type)
}

// nvml.UnitGetFanSpeedInfo()
func UnitGetFanSpeedInfo(Unit Unit) (UnitFanSpeeds, Return) {
	var FanSpeeds UnitFanSpeeds
	ret := nvmlUnitGetFanSpeedInfo(Unit, &FanSpeeds)
	return FanSpeeds, ret
}

func (Unit Unit) GetFanSpeedInfo() (UnitFanSpeeds, Return) {
	return UnitGetFanSpeedInfo(Unit)
}

// nvml.UnitGetDevices()
func UnitGetDevices(Unit Unit) ([]Device, Return) {
	var DeviceCount uint32 = 1 // Will be reduced upon returning
	for {
		Devices := make([]Device, DeviceCount)
		ret := nvmlUnitGetDevices(Unit, &DeviceCount, &Devices[0])
		if ret == SUCCESS {
			return Devices[:DeviceCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		DeviceCount *= 2
	}
}

func (Unit Unit) GetDevices() ([]Device, Return) {
	return UnitGetDevices(Unit)
}

// nvml.UnitSetLedState()
func UnitSetLedState(Unit Unit, Color LedColor) Return {
	return nvmlUnitSetLedState(Unit, Color)
}

func (Unit Unit) SetLedState(Color LedColor) Return {
	return UnitSetLedState(Unit, Color)
}
