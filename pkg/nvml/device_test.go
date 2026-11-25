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
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestGetTopologyCommonAncestor(t *testing.T) {
	type wrappedDevice struct {
		Device
	}

	type wrappedWrappedDevice struct {
		wrappedDevice
	}

	testCases := []struct {
		description string
		device      Device
	}{
		{
			description: "nvmlDevice",
			device:      nvmlDevice{},
		},
		{
			description: "pointer to nvmlDevice",
			device:      &nvmlDevice{},
		},
		{
			description: "wrapped device",
			device: wrappedDevice{
				Device: nvmlDevice{},
			},
		},
		{
			description: "pointer to wrapped device",
			device: &wrappedDevice{
				Device: nvmlDevice{},
			},
		},
		{
			description: "nested wrapped device",
			device: wrappedWrappedDevice{
				wrappedDevice: wrappedDevice{
					Device: nvmlDevice{},
				},
			},
		},
		{
			description: "non-device fields included",
			device: struct {
				name string
				Name string
				Device
			}{
				Device: wrappedDevice{
					Device: nvmlDevice{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			defer setNvmlDeviceGetTopologyCommonAncestorStubForTest(SUCCESS)()

			_, ret := nvmlDevice{}.GetTopologyCommonAncestor(tc.device)
			require.Equal(t, SUCCESS, ret)
		})
	}
}

func setNvmlDeviceGetTopologyCommonAncestorStubForTest(ret Return) func() {
	original := nvmlDeviceGetTopologyCommonAncestorStub

	nvmlDeviceGetTopologyCommonAncestorStub = func(Device1, Device2 nvmlDevice, PathInfo *GpuTopologyLevel) Return {
		return ret
	}
	return func() {
		nvmlDeviceGetTopologyCommonAncestorStub = original
	}
}

func TestDeviceGetFieldValues(t *testing.T) {
	original := nvmlDeviceGetFieldValuesStub
	t.Cleanup(func() { nvmlDeviceGetFieldValuesStub = original })

	entryPoints := []struct {
		name string
		call func(Device, []FieldValue) Return
	}{
		{
			name: "device",
			call: func(device Device, values []FieldValue) Return {
				return device.GetFieldValues(values)
			},
		},
		{name: "library", call: newLibrary().DeviceGetFieldValues},
		{name: "package", call: DeviceGetFieldValues},
	}
	testCases := []struct {
		name      string
		ret       Return
		populated int
	}{
		{name: "mixed field results", ret: SUCCESS, populated: 2},
		{name: "partial results on error", ret: ERROR_UNKNOWN, populated: 1},
		{name: "no results on error", ret: ERROR_GPU_IS_LOST},
	}

	for _, entryPoint := range entryPoints {
		for _, tc := range testCases {
			t.Run(entryPoint.name+"/"+tc.name, func(t *testing.T) {
				device := nvmlDevice{}
				storage := []FieldValue{
					{FieldId: FI_DEV_NVLINK_THROUGHPUT_DATA_TX, ScopeId: 0, NvmlReturn: uint32(SUCCESS)},
					{FieldId: FI_DEV_NVLINK_THROUGHPUT_DATA_RX, ScopeId: 1, NvmlReturn: uint32(ERROR_NOT_SUPPORTED)},
					{FieldId: FI_DEV_NVLINK_THROUGHPUT_DATA_TX, ScopeId: 17, Timestamp: 99, Value: [8]byte{42}},
					{FieldId: 123, ScopeId: 456, Timestamp: 789, Value: [8]byte{99}},
				}
				values := storage[:3]
				sentinel := storage[3]
				initial := append([]FieldValue(nil), values...)
				for i := range initial {
					initial[i].NvmlReturn = uint32(ERROR_INVALID_ARGUMENT)
				}
				expected := append([]FieldValue(nil), initial...)
				for i := 0; i < tc.populated; i++ {
					expected[i].Timestamp = int64(1000 + i)
					expected[i].LatencyUsec = int64(10 + i)
					expected[i].ValueType = uint32(VALUE_TYPE_UNSIGNED_LONG_LONG)
					expected[i].NvmlReturn = uint32(SUCCESS)
					expected[i].Value = [8]byte{byte(20 + i)}
				}
				if tc.populated > 1 {
					expected[1].NvmlReturn = uint32(ERROR_NOT_SUPPORTED)
				}

				calls := 0
				nvmlDeviceGetFieldValuesStub = func(gotDevice nvmlDevice, count int32, fields *FieldValue) Return {
					calls++
					require.Equal(t, device, gotDevice)
					require.Equal(t, int32(len(values)), count)
					require.Equal(t, initial, unsafe.Slice(fields, int(count)))

					// Model the driver's access to one entry beyond the requested count.
					buffer := unsafe.Slice(fields, int(count)+1)
					require.Equal(t, FieldValue{NvmlReturn: uint32(ERROR_INVALID_ARGUMENT)}, buffer[count])
					copy(buffer[:tc.populated], expected[:tc.populated])
					buffer[count] = FieldValue{FieldId: 999, Value: [8]byte{255}}
					return tc.ret
				}

				require.Equal(t, tc.ret, entryPoint.call(device, values))
				require.Equal(t, 1, calls)
				require.Equal(t, expected, values)
				require.Equal(t, sentinel, storage[3], "memory beyond the caller's slice must be unchanged")
			})
		}
	}
}

func TestDeviceGetFieldValuesEmpty(t *testing.T) {
	original := nvmlDeviceGetFieldValuesStub
	t.Cleanup(func() { nvmlDeviceGetFieldValuesStub = original })
	nvmlDeviceGetFieldValuesStub = func(nvmlDevice, int32, *FieldValue) Return {
		t.Fatal("NVML must not be called for empty input")
		return ERROR_UNKNOWN
	}

	device := nvmlDevice{}
	require.Equal(t, ERROR_INVALID_ARGUMENT, device.GetFieldValues(nil))
	require.Equal(t, ERROR_INVALID_ARGUMENT, device.GetFieldValues([]FieldValue{}))
}
