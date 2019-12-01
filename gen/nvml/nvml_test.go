// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

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
