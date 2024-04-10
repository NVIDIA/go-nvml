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

// nvml.UnitGetCount()
func (l *library) UnitGetCount() (int, Return) {
	var UnitCount uint32
	ret := nvmlUnitGetCount(&UnitCount)
	return int(UnitCount), ret
}

// nvml.UnitGetHandleByIndex()
func (l *library) UnitGetHandleByIndex(Index int) (Unit, Return) {
	var Unit nvmlUnit
	ret := nvmlUnitGetHandleByIndex(uint32(Index), &Unit)
	return Unit, ret
}

// nvml.UnitGetUnitInfo()
func (l *library) UnitGetUnitInfo(Unit Unit) (UnitInfo, Return) {
	return Unit.GetUnitInfo()
}

func (Unit nvmlUnit) GetUnitInfo() (UnitInfo, Return) {
	var Info UnitInfo
	ret := nvmlUnitGetUnitInfo(Unit, &Info)
	return Info, ret
}

// nvml.UnitGetLedState()
func (l *library) UnitGetLedState(Unit Unit) (LedState, Return) {
	return Unit.GetLedState()
}

func (Unit nvmlUnit) GetLedState() (LedState, Return) {
	var State LedState
	ret := nvmlUnitGetLedState(Unit, &State)
	return State, ret
}

// nvml.UnitGetPsuInfo()
func (l *library) UnitGetPsuInfo(Unit Unit) (PSUInfo, Return) {
	return Unit.GetPsuInfo()
}

func (Unit nvmlUnit) GetPsuInfo() (PSUInfo, Return) {
	var Psu PSUInfo
	ret := nvmlUnitGetPsuInfo(Unit, &Psu)
	return Psu, ret
}

// nvml.UnitGetTemperature()
func (l *library) UnitGetTemperature(Unit Unit, Type int) (uint32, Return) {
	return Unit.GetTemperature(Type)
}

func (Unit nvmlUnit) GetTemperature(Type int) (uint32, Return) {
	var Temp uint32
	ret := nvmlUnitGetTemperature(Unit, uint32(Type), &Temp)
	return Temp, ret
}

// nvml.UnitGetFanSpeedInfo()
func (l *library) UnitGetFanSpeedInfo(Unit Unit) (UnitFanSpeeds, Return) {
	return Unit.GetFanSpeedInfo()
}

func (Unit nvmlUnit) GetFanSpeedInfo() (UnitFanSpeeds, Return) {
	var FanSpeeds UnitFanSpeeds
	ret := nvmlUnitGetFanSpeedInfo(Unit, &FanSpeeds)
	return FanSpeeds, ret
}

// nvml.UnitGetDevices()
func (l *library) UnitGetDevices(Unit Unit) ([]Device, Return) {
	return Unit.GetDevices()
}

func (Unit nvmlUnit) GetDevices() ([]Device, Return) {
	var DeviceCount uint32 = 1 // Will be reduced upon returning
	for {
		Devices := make([]nvmlDevice, DeviceCount)
		ret := nvmlUnitGetDevices(Unit, &DeviceCount, &Devices[0])
		if ret == SUCCESS {
			return convertSlice[nvmlDevice, Device](Devices[:DeviceCount]), ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		DeviceCount *= 2
	}
}

// nvml.UnitSetLedState()
func (l *library) UnitSetLedState(Unit Unit, Color LedColor) Return {
	return Unit.SetLedState(Color)
}

func (Unit nvmlUnit) SetLedState(Color LedColor) Return {
	return nvmlUnitSetLedState(Unit, Color)
}
