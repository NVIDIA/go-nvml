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
	"testing"
)

func TestInit(t *testing.T) {
	ret := Init()
	if ret != SUCCESS {
		t.Errorf("Init: %v", ret)
	} else {
		t.Logf("Init: %v", ret)
	}

	ret = Shutdown()
	if ret != SUCCESS {
		t.Errorf("Shutdown: %v", ret)
	} else {
		t.Logf("Shutdown: %v", ret)
	}
}

func TestSystem(t *testing.T) {
	Init()
	defer Shutdown()

	driverVersion, ret := SystemGetDriverVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetDriverVersion: %v", ret)
	} else {
		t.Logf("SystemGetDriverVersion: %v", ret)
		t.Logf("  version: %v", driverVersion)
	}

	nvmlVersion, ret := SystemGetNVMLVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetNVMLVersion: %v", ret)
	} else {
		t.Logf("SystemGetNVMLVersion: %v", ret)
		t.Logf("  version: %v", nvmlVersion)
	}

	cudaDriverVersion, ret := SystemGetCudaDriverVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetCudaDriverVersion: %v", ret)
	} else {
		t.Logf("SystemGetCudaDriverVersion: %v", ret)
		t.Logf("  version: %v", cudaDriverVersion)
	}

	cudaDriverVersionV2, ret := SystemGetCudaDriverVersion_v2()
	if ret != SUCCESS {
		t.Errorf("SystemGetCudaDriverVersion_v2: %v", ret)
	} else {
		t.Logf("SystemGetCudaDriverVersion_v2: %v", ret)
		t.Logf("  version: %v", cudaDriverVersionV2)
	}

	processName, ret := SystemGetProcessName(1)
	if ret != SUCCESS {
		t.Errorf("SystemGetProcessName: %v", ret)
	} else {
		t.Logf("SystemGetProcessName: %v", ret)
		t.Logf("  name: %v", processName)
	}

	hwbcEntries, ret := SystemGetHicVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetHicVersion: %v", ret)
	} else {
		t.Logf("SystemGetHicVersion: %v", ret)
		t.Logf("  count: %v", len(hwbcEntries))
		for i, entry := range hwbcEntries {
			t.Logf("  device[%v]: %v", i, entry)
		}
	}

	deviceArray, ret := SystemGetTopologyGpuSet(0)
	if ret != SUCCESS {
		t.Errorf("SystemGetTopologyGpuSet: %v", ret)
	} else {
		t.Logf("SystemGetTopologyGpuSet: %v", ret)
		t.Logf("  count: %v", len(deviceArray))
		for i, device := range deviceArray {
			t.Logf("  device[%v]: %v", i, device)
		}
	}
}

func TestUnit(t *testing.T) {
	Init()
	defer Shutdown()

	unitCount, ret := UnitGetCount()
	if ret != SUCCESS {
		t.Errorf("UnitGetCount: %v", ret)
	} else {
		t.Logf("UnitGetCount: %v", ret)
		t.Logf("  count: %v", unitCount)
	}

	if unitCount == 0 {
		t.Skip("Skipping test with no Units.")
	}

	unit, ret := UnitGetHandleByIndex(0)
	if ret != SUCCESS {
		t.Errorf("UnitGetHandleByIndex: %v", ret)
	} else {
		t.Logf("UnitGetHandleByIndex: %v", ret)
		t.Logf("  unit: %v", unit)
	}

	info, ret := UnitGetUnitInfo(unit)
	if ret != SUCCESS {
		t.Errorf("UnitGetUnitInfo: %v", ret)
	} else {
		t.Logf("UnitGetUnitInfo: %v", ret)
		t.Logf("  info: %v", info)
	}

	info, ret = unit.GetUnitInfo()
	if ret != SUCCESS {
		t.Errorf("Unit.GetUnitInfo: %v", ret)
	} else {
		t.Logf("Unit.GetUnitInfo: %v", ret)
		t.Logf("  info: %v", info)
	}

	state, ret := UnitGetLedState(unit)
	if ret != SUCCESS {
		t.Errorf("UnitGetLedState: %v", ret)
	} else {
		t.Logf("UnitGetLedState: %v", ret)
		t.Logf("  state: %v", state)
	}

	state, ret = unit.GetLedState()
	if ret != SUCCESS {
		t.Errorf("Unit.GetLedState: %v", ret)
	} else {
		t.Logf("Unit.GetLedState: %v", ret)
		t.Logf("  state: %v", state)
	}

	psu, ret := UnitGetPsuInfo(unit)
	if ret != SUCCESS {
		t.Errorf("UnitGetPsuInfo: %v", ret)
	} else {
		t.Logf("UnitGetPsuInfo: %v", ret)
		t.Logf("  psu: %v", psu)
	}

	psu, ret = unit.GetPsuInfo()
	if ret != SUCCESS {
		t.Errorf("Unit.GetPsuInfo: %v", ret)
	} else {
		t.Logf("Unit.GetPsuInfo: %v", ret)
		t.Logf("  psu: %v", psu)
	}

	temp, ret := UnitGetTemperature(unit, 0)
	if ret != SUCCESS {
		t.Errorf("UnitGetTemperature: %v", ret)
	} else {
		t.Logf("UnitGetTemperature: %v", ret)
		t.Logf("  temp: %v", temp)
	}

	temp, ret = unit.GetTemperature(0)
	if ret != SUCCESS {
		t.Errorf("Unit.GetTemperature: %v", ret)
	} else {
		t.Logf("Unit.GetTemperature: %v", ret)
		t.Logf("  temp: %v", temp)
	}

	speed, ret := UnitGetFanSpeedInfo(unit)
	if ret != SUCCESS {
		t.Errorf("UnitGetFanSpeedInfo: %v", ret)
	} else {
		t.Logf("UnitGetFanSpeedInfo: %v", ret)
		t.Logf("  speed: %v", speed)
	}

	speed, ret = unit.GetFanSpeedInfo()
	if ret != SUCCESS {
		t.Errorf("Unit.GetFanSpeedInfo: %v", ret)
	} else {
		t.Logf("Unit.GetFanSpeedInfo: %v", ret)
		t.Logf("  speed: %v", speed)
	}

	devices, ret := UnitGetDevices(unit)
	if ret != SUCCESS {
		t.Errorf("UnitGetDevices: %v", ret)
	} else {
		t.Logf("UnitGetDevices: %v", ret)
		t.Logf("  count: %v", len(devices))
		for i, device := range devices {
			t.Logf("  device[%v]: %v", i, device)
		}
	}

	devices, ret = unit.GetDevices()
	if ret != SUCCESS {
		t.Errorf("Unit.GetDevices: %v", ret)
	} else {
		t.Logf("Unit.GetDevices: %v", ret)
		t.Logf("  count: %v", len(devices))
		for i, device := range devices {
			t.Logf("  device[%v]: %v", i, device)
		}
	}

	ret = UnitSetLedState(unit, 0)
	if ret != SUCCESS {
		t.Errorf("UnitSetLedState: %v", ret)
	} else {
		t.Logf("UnitSetLedState: %v", ret)
	}

	ret = unit.SetLedState(0)
	if ret != SUCCESS {
		t.Errorf("Unit.SetLedState: %v", ret)
	} else {
		t.Logf("Unit.SetLedState: %v", ret)
	}
}

func TestEventSet(t *testing.T) {
	Init()
	defer Shutdown()

	set, ret := EventSetCreate()
	if ret != SUCCESS {
		t.Errorf("EventSetCreate: %v", ret)
	} else {
		t.Logf("EventSetCreate: %v", ret)
		t.Logf("  set: %v", set)
	}

	data, ret := EventSetWait(set, 0)
	if ret != ERROR_TIMEOUT {
		t.Errorf("EventSetWait: %v", ret)
	} else {
		t.Logf("EventSetWait: %v", ret)
		t.Logf("  data: %v", data)
	}

	data, ret = set.Wait(0)
	if ret != ERROR_TIMEOUT {
		t.Errorf("EventSet.Wait: %v", ret)
	} else {
		t.Logf("EventSet.Wait: %v", ret)
		t.Logf("  data: %v", data)
	}

	ret = EventSetFree(set)
	if ret != SUCCESS {
		t.Errorf("EventSetFree: %v", ret)
	} else {
		t.Logf("EventSetFree: %v", ret)
	}

	set, _ = EventSetCreate()
	ret = set.Free()
	if ret != SUCCESS {
		t.Errorf("EventSet.Free: %v", ret)
	} else {
		t.Logf("EventSet.Free: %v", ret)
	}
}
