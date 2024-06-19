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
func (l *library) UnitGetCount() (int, error) {
	var UnitCount uint32
	ret := nvmlUnitGetCount(&UnitCount)
	return int(UnitCount), ret.error()
}

// nvml.UnitGetHandleByIndex()
func (l *library) UnitGetHandleByIndex(index int) (Unit, error) {
	var unit nvmlUnit
	ret := nvmlUnitGetHandleByIndex(uint32(index), &unit)
	return unit, ret.error()
}

// nvml.UnitGetUnitInfo()
func (l *library) UnitGetUnitInfo(unit Unit) (UnitInfo, error) {
	return unit.GetUnitInfo()
}

func (unit nvmlUnit) GetUnitInfo() (UnitInfo, error) {
	var info UnitInfo
	ret := nvmlUnitGetUnitInfo(unit, &info)
	return info, ret.error()
}

// nvml.UnitGetLedState()
func (l *library) UnitGetLedState(unit Unit) (LedState, error) {
	return unit.GetLedState()
}

func (unit nvmlUnit) GetLedState() (LedState, error) {
	var state LedState
	ret := nvmlUnitGetLedState(unit, &state)
	return state, ret.error()
}

// nvml.UnitGetPsuInfo()
func (l *library) UnitGetPsuInfo(unit Unit) (PSUInfo, error) {
	return unit.GetPsuInfo()
}

func (unit nvmlUnit) GetPsuInfo() (PSUInfo, error) {
	var psu PSUInfo
	ret := nvmlUnitGetPsuInfo(unit, &psu)
	return psu, ret.error()
}

// nvml.UnitGetTemperature()
func (l *library) UnitGetTemperature(unit Unit, ttype int) (uint32, error) {
	return unit.GetTemperature(ttype)
}

func (unit nvmlUnit) GetTemperature(ttype int) (uint32, error) {
	var temp uint32
	ret := nvmlUnitGetTemperature(unit, uint32(ttype), &temp)
	return temp, ret.error()
}

// nvml.UnitGetFanSpeedInfo()
func (l *library) UnitGetFanSpeedInfo(unit Unit) (UnitFanSpeeds, error) {
	return unit.GetFanSpeedInfo()
}

func (unit nvmlUnit) GetFanSpeedInfo() (UnitFanSpeeds, error) {
	var fanSpeeds UnitFanSpeeds
	ret := nvmlUnitGetFanSpeedInfo(unit, &fanSpeeds)
	return fanSpeeds, ret.error()
}

// nvml.UnitGetDevices()
func (l *library) UnitGetDevices(unit Unit) ([]Device, error) {
	return unit.GetDevices()
}

func (unit nvmlUnit) GetDevices() ([]Device, error) {
	var deviceCount uint32 = 1 // Will be reduced upon returning
	for {
		devices := make([]nvmlDevice, deviceCount)
		ret := nvmlUnitGetDevices(unit, &deviceCount, &devices[0])
		if ret == nvmlSUCCESS {
			return convertSlice[nvmlDevice, Device](devices[:deviceCount]), ret.error()
		}
		if ret != nvmlERROR_INSUFFICIENT_SIZE {
			return nil, ret.error()
		}
		deviceCount *= 2
	}
}

// nvml.UnitSetLedState()
func (l *library) UnitSetLedState(unit Unit, color LedColor) error {
	return unit.SetLedState(color)
}

func (unit nvmlUnit) SetLedState(color LedColor) error {
	return nvmlUnitSetLedState(unit, color).error()
}
