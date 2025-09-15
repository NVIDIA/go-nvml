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

package nvml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCgoAPIDeviceHandleConversion verifies that we can convert an nvmlDevice to a Device interface and use it via CgoAPI.
func TestCgoAPIDeviceHandleConversion(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	// Test that we can call CgoAPI methods with Device interface
	// The CgoAPI should automatically convert nvmlDevice to Device interface

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	// Test getting a device handle by index
	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	if ret == SUCCESS {
		// If we got a device, test calling methods on it
		var name [DEVICE_NAME_BUFFER_SIZE]byte
		ret = CgoAPI.DeviceGetName(device, &name[0], DEVICE_NAME_BUFFER_SIZE)
		require.Equal(t, SUCCESS, ret, "DeviceGetName should succeed")

		// Test getting device memory info
		var memory Memory
		ret = CgoAPI.DeviceGetMemoryInfo(device, &memory)
		require.Equal(t, SUCCESS, ret, "DeviceGetMemoryInfo should succeed")

		// Test getting compute mode
		var mode ComputeMode
		ret = CgoAPI.DeviceGetComputeMode(device, &mode)
		require.Equal(t, SUCCESS, ret, "DeviceGetComputeMode should succeed")
	} else {
		t.Logf("No devices available for testing, skipping device-specific tests")
	}
}

// TestCgoAPIDeviceMethods exercises various device-related CgoAPI methods.
func TestCgoAPIDeviceMethods(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	// Test system-level device count
	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")
	t.Logf("Device count: %d", count)

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	// Test getting device handle by index
	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test various device information methods
	var name [DEVICE_NAME_BUFFER_SIZE]byte
	ret = CgoAPI.DeviceGetName(device, &name[0], DEVICE_NAME_BUFFER_SIZE)
	require.Equal(t, SUCCESS, ret, "DeviceGetName should succeed")
	t.Logf("Device name: %s", string(name[:]))

	var uuid [DEVICE_UUID_BUFFER_SIZE]byte
	ret = CgoAPI.DeviceGetUUID(device, &uuid[0], DEVICE_UUID_BUFFER_SIZE)
	require.Equal(t, SUCCESS, ret, "DeviceGetUUID should succeed")
	t.Logf("Device UUID: %s", string(uuid[:]))

	var brand BrandType
	ret = CgoAPI.DeviceGetBrand(device, &brand)
	require.Equal(t, SUCCESS, ret, "DeviceGetBrand should succeed")
	t.Logf("Device brand: %v", brand)

	var busType BusType
	ret = CgoAPI.DeviceGetBusType(device, &busType)
	require.Equal(t, SUCCESS, ret, "DeviceGetBusType should succeed")
	t.Logf("Device bus type: %v", busType)
}

// TestCgoAPIDeviceMemoryInfo tests device memory information methods via CgoAPI.
func TestCgoAPIDeviceMemoryInfo(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test memory info
	var memInfo Memory
	ret = CgoAPI.DeviceGetMemoryInfo(device, &memInfo)
	require.Equal(t, SUCCESS, ret, "DeviceGetMemoryInfo should succeed")
	t.Logf("Memory info: Total=%d, Used=%d, Free=%d",
		memInfo.Total, memInfo.Used, memInfo.Free)

	// Test BAR1 memory info
	var bar1Info BAR1Memory
	ret = CgoAPI.DeviceGetBAR1MemoryInfo(device, &bar1Info)
	require.Equal(t, SUCCESS, ret, "DeviceGetBAR1MemoryInfo should succeed")
	t.Logf("BAR1 memory: Total=%d, Used=%d, Free=%d",
		bar1Info.Bar1Total, bar1Info.Bar1Used, bar1Info.Bar1Free)
}

// TestCgoAPIDeviceUtilization tests device utilization methods via CgoAPI.
func TestCgoAPIDeviceUtilization(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test GPU utilization
	var utilization Utilization
	ret = CgoAPI.DeviceGetUtilizationRates(device, &utilization)
	require.Equal(t, SUCCESS, ret, "DeviceGetUtilizationRates should succeed")
	t.Logf("GPU utilization: %d%%, sampling period: %dμs", utilization.Gpu, utilization.Memory)

	// Test encoder utilization
	var encoderUtil uint32
	var encoderSamplingPeriodUs uint32
	ret = CgoAPI.DeviceGetEncoderUtilization(device, &encoderUtil, &encoderSamplingPeriodUs)
	require.Equal(t, SUCCESS, ret, "DeviceGetEncoderUtilization should succeed")
	t.Logf("Encoder utilization: %d%%, sampling period: %dμs", encoderUtil, encoderSamplingPeriodUs)

	// Test decoder utilization
	var decoderUtil uint32
	var decoderSamplingPeriodUs uint32
	ret = CgoAPI.DeviceGetDecoderUtilization(device, &decoderUtil, &decoderSamplingPeriodUs)
	require.Equal(t, SUCCESS, ret, "DeviceGetDecoderUtilization should succeed")
	t.Logf("Decoder utilization: %d%%, sampling period: %dμs", decoderUtil, decoderSamplingPeriodUs)
}

// TestCgoAPIDeviceTemperature tests device temperature methods via CgoAPI.
func TestCgoAPIDeviceTemperature(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test GPU temperature
	var temp uint32
	ret = CgoAPI.DeviceGetTemperature(device, TEMPERATURE_GPU, &temp)
	require.Equal(t, SUCCESS, ret, "DeviceGetTemperature should succeed")
	t.Logf("GPU temperature: %d°C", temp)

	// Test memory temperature (using GPU temperature as fallback since TEMPERATURE_MEMORY doesn't exist)
	var memTemp uint32
	ret = CgoAPI.DeviceGetTemperature(device, TEMPERATURE_GPU, &memTemp)
	require.Equal(t, SUCCESS, ret, "DeviceGetTemperature should succeed")
	t.Logf("GPU temperature: %d°C", memTemp)
}

// TestCgoAPIDevicePower tests device power methods via CgoAPI.
func TestCgoAPIDevicePower(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test power management mode
	var mode EnableState
	ret = CgoAPI.DeviceGetPowerManagementMode(device, &mode)
	require.Equal(t, SUCCESS, ret, "DeviceGetPowerManagementMode should succeed")
	t.Logf("Power management mode: %v", mode)

	// Test power usage
	var power uint32
	ret = CgoAPI.DeviceGetPowerUsage(device, &power)
	require.Equal(t, SUCCESS, ret, "DeviceGetPowerUsage should succeed")
	t.Logf("Power usage: %d mW", power)

	// Test power limit
	var limit uint32
	ret = CgoAPI.DeviceGetEnforcedPowerLimit(device, &limit)
	require.Equal(t, SUCCESS, ret, "DeviceGetEnforcedPowerLimit should succeed")
	t.Logf("Enforced power limit: %d mW", limit)
}

// TestCgoAPISystemMethods tests system-level CgoAPI methods.
func TestCgoAPISystemMethods(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	// Test driver version
	var driverVersion [SYSTEM_DRIVER_VERSION_BUFFER_SIZE]byte
	ret = CgoAPI.SystemGetDriverVersion(&driverVersion[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	require.Equal(t, SUCCESS, ret, "SystemGetDriverVersion should succeed")
	t.Logf("Driver version: %s", string(driverVersion[:]))

	// Test NVML version
	var nvmlVersion [SYSTEM_NVML_VERSION_BUFFER_SIZE]byte
	ret = CgoAPI.SystemGetNVMLVersion(&nvmlVersion[0], SYSTEM_NVML_VERSION_BUFFER_SIZE)
	require.Equal(t, SUCCESS, ret, "SystemGetNVMLVersion should succeed")
	t.Logf("NVML version: %s", string(nvmlVersion[:]))

	// Test CUDA driver version
	var cudaVersion int32
	ret = CgoAPI.SystemGetCudaDriverVersion(&cudaVersion)
	require.Equal(t, SUCCESS, ret, "SystemGetCudaDriverVersion should succeed")
	t.Logf("CUDA driver version: %d", cudaVersion)
}

// TestCgoAPIDeviceTopology tests device topology methods via CgoAPI.
func TestCgoAPIDeviceTopology(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count < 2 {
		t.Skip("Need at least 2 devices for topology tests")
	}

	var device1, device2 Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device1)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex(0) should succeed")

	ret = CgoAPI.DeviceGetHandleByIndex(1, &device2)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex(1) should succeed")

	// Test topology common ancestor
	var pathInfo GpuTopologyLevel
	ret = CgoAPI.DeviceGetTopologyCommonAncestor(device1, device2, &pathInfo)
	require.Equal(t, SUCCESS, ret, "DeviceGetTopologyCommonAncestor should succeed")
	t.Logf("Topology common ancestor: %v", pathInfo)

	// Test P2P status
	var p2pStatus GpuP2PStatus
	ret = CgoAPI.DeviceGetP2PStatus(device1, device2, P2P_CAPS_INDEX_READ, &p2pStatus)
	require.Equal(t, SUCCESS, ret, "DeviceGetP2PStatus should succeed")
	t.Logf("P2P read status: %v", p2pStatus)

	// Test if devices are on same board
	var onSameBoard int32
	ret = CgoAPI.DeviceOnSameBoard(device1, device2, &onSameBoard)
	require.Equal(t, SUCCESS, ret, "DeviceOnSameBoard should succeed")
	t.Logf("Devices on same board: %v", onSameBoard == 1)
}

// TestCgoAPIDeviceClockInfo tests device clock information methods via CgoAPI.
func TestCgoAPIDeviceClockInfo(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Test current clock frequencies
	var clockFreqs DeviceCurrentClockFreqs
	clockFreqs.Version = STRUCT_VERSION(clockFreqs, 1)
	ret = CgoAPI.DeviceGetCurrentClockFreqs(device, &clockFreqs)
	require.Equal(t, SUCCESS, ret, "DeviceGetCurrentClockFreqs should succeed")
	t.Logf("Current clock freqs: Version=%d, Str=%s",
		clockFreqs.Version, string(convertSlice[int8, uint8](clockFreqs.Str[:])))

	// Test applications clock
	var appClock uint32
	ret = CgoAPI.DeviceGetApplicationsClock(device, CLOCK_SM, &appClock)
	require.Equal(t, SUCCESS, ret, "DeviceGetApplicationsClock should succeed")
	t.Logf("Applications SM clock: %d MHz", appClock)
}

// TestCgoAPI_DeviceHandleInterface verifies that a Device handle obtained via CgoAPI
// can be used to call Device interface methods directly.
func TestCgoAPI_DeviceHandleInterface(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")
	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Now call Device interface methods directly
	name, ret := device.GetName()
	require.Equal(t, SUCCESS, ret, "Device.GetName should succeed")
	t.Logf("Device.GetName: %s", name)

	uuid, ret := device.GetUUID()
	require.Equal(t, SUCCESS, ret, "Device.GetUUID should succeed")
	t.Logf("Device.GetUUID: %s", uuid)

	brand, ret := device.GetBrand()
	require.Equal(t, SUCCESS, ret, "Device.GetBrand should succeed")
	t.Logf("Device.GetBrand: %v", brand)

	busType, ret := device.GetBusType()
	require.Equal(t, SUCCESS, ret, "Device.GetBusType should succeed")
	t.Logf("Device.GetBusType: %v", busType)
}

// requireConfComputeSupport checks if confidential compute is supported and skips the test if not.
func requireConfComputeSupport(t *testing.T, device Device) {
	var attestationReport ConfComputeGpuAttestationReport
	// Initialize nonce to all zeros
	for i := range attestationReport.Nonce {
		attestationReport.Nonce[i] = 0
	}

	ret := CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport)
	if ret != SUCCESS {
		t.Skipf("Confidential compute not supported on this system (returned: %v)", ret)
	}
}

// TestCgoAPIDeviceConfComputeAttestation tests confidential compute GPU attestation report methods via CgoAPI.
func TestCgoAPIDeviceConfComputeAttestation(t *testing.T) {
	requireLibNvidiaML(t)

	// Initialize NVML
	ret := Init()
	require.Equal(t, SUCCESS, ret, "Init should succeed")
	defer Shutdown()

	var count uint32
	ret = CgoAPI.DeviceGetCount(&count)
	require.Equal(t, SUCCESS, ret, "DeviceGetCount should succeed")

	if count == 0 {
		t.Skip("No devices available for testing")
	}

	var device Device
	ret = CgoAPI.DeviceGetHandleByIndex(0, &device)
	require.Equal(t, SUCCESS, ret, "DeviceGetHandleByIndex should succeed")

	// Check if confidential compute is supported
	requireConfComputeSupport(t, device)

	// Test with zero nonce
	var attestationReport ConfComputeGpuAttestationReport
	// Initialize nonce to all zeros
	for i := range attestationReport.Nonce {
		attestationReport.Nonce[i] = 0
	}

	ret = CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport)
	require.Equal(t, SUCCESS, ret, "DeviceGetConfComputeGpuAttestationReport should succeed")
	t.Logf("Attestation report: IsCecPresent=%d, AttestationSize=%d, CecSize=%d",
		attestationReport.IsCecAttestationReportPresent,
		attestationReport.AttestationReportSize,
		attestationReport.CecAttestationReportSize)

	// Test with pattern nonce
	var attestationReport2 ConfComputeGpuAttestationReport
	attestationReport2.Nonce = createPatternNonce(1, 1) // Creates [1, 2, 3, ..., 32]

	ret = CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport2)
	require.Equal(t, SUCCESS, ret, "DeviceGetConfComputeGpuAttestationReport with pattern nonce should succeed")
	t.Logf("Attestation report with pattern nonce: IsCecPresent=%d, AttestationSize=%d, CecSize=%d",
		attestationReport2.IsCecAttestationReportPresent,
		attestationReport2.AttestationReportSize,
		attestationReport2.CecAttestationReportSize)

	// Test with all-ones nonce
	var attestationReport3 ConfComputeGpuAttestationReport
	// Set nonce to all ones
	for i := range attestationReport3.Nonce {
		attestationReport3.Nonce[i] = 0xFF
	}

	ret = CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport3)
	require.Equal(t, SUCCESS, ret, "DeviceGetConfComputeGpuAttestationReport with all-ones nonce should succeed")
	t.Logf("Attestation report with all-ones nonce: IsCecPresent=%d, AttestationSize=%d, CecSize=%d",
		attestationReport3.IsCecAttestationReportPresent,
		attestationReport3.AttestationReportSize,
		attestationReport3.CecAttestationReportSize)

	// Test nonce field access and modification using helper function
	testNonce := createNonceFromString("test-nonce-32-bytes-long-string")

	var attestationReport4 ConfComputeGpuAttestationReport
	attestationReport4.Nonce = testNonce

	// Verify nonce was set correctly
	for i := 0; i < 32; i++ {
		require.Equal(t, testNonce[i], attestationReport4.Nonce[i], "Nonce field should be set correctly")
	}

	ret = CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport4)
	require.Equal(t, SUCCESS, ret, "DeviceGetConfComputeGpuAttestationReport with string nonce should succeed")
	t.Logf("Attestation report with string nonce: IsCecPresent=%d, AttestationSize=%d, CecSize=%d",
		attestationReport4.IsCecAttestationReportPresent,
		attestationReport4.AttestationReportSize,
		attestationReport4.CecAttestationReportSize)

	// Test with pseudo-random nonce
	var attestationReport5 ConfComputeGpuAttestationReport
	attestationReport5.Nonce = createRandomNonce(42) // Use seed 42 for reproducible results

	ret = CgoAPI.DeviceGetConfComputeGpuAttestationReport(device, &attestationReport5)
	require.Equal(t, SUCCESS, ret, "DeviceGetConfComputeGpuAttestationReport with random nonce should succeed")
	t.Logf("Attestation report with random nonce: IsCecPresent=%d, AttestationSize=%d, CecSize=%d",
		attestationReport5.IsCecAttestationReportPresent,
		attestationReport5.AttestationReportSize,
		attestationReport5.CecAttestationReportSize)
}

// createNonceFromString creates a 32-byte nonce from a string, padding with zeros if needed.
func createNonceFromString(s string) [32]uint8 {
	var nonce [32]uint8
	copy(nonce[:], []byte(s))
	return nonce
}

// createPatternNonce creates a 32-byte nonce with a repeating pattern.
func createPatternNonce(start, step uint8) [32]uint8 {
	var nonce [32]uint8
	for i := range nonce {
		nonce[i] = start + uint8(i)*step
	}
	return nonce
}

// createRandomNonce creates a 32-byte nonce with pseudo-random values.
func createRandomNonce(seed int64) [32]uint8 {
	var nonce [32]uint8
	// Simple pseudo-random generation for testing
	for i := range nonce {
		nonce[i] = uint8((seed + int64(i)*7) % 256)
	}
	return nonce
}
