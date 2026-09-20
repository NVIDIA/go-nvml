/**
# Copyright 2026 NVIDIA CORPORATION
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

func TestReturnedCounts(t *testing.T) {
	originalSamples := nvmlDeviceGetSamplesStub
	originalNearest := nvmlDeviceGetTopologyNearestGpusStub
	originalTopology := nvmlSystemGetTopologyGpuSetStub
	t.Cleanup(func() {
		nvmlDeviceGetSamplesStub = originalSamples
		nvmlDeviceGetTopologyNearestGpusStub = originalNearest
		nvmlSystemGetTopologyGpuSetStub = originalTopology
	})

	testCases := []struct {
		description                   string
		initialCount, returnedCount   uint32
		firstReturn, secondReturn     Return
		expectedLength, expectedCalls int
	}{
		{"same count", 2, 2, SUCCESS, SUCCESS, 2, 2},
		{"fewer results", 2, 1, SUCCESS, SUCCESS, 1, 2},
		{"no results after query", 2, 0, SUCCESS, SUCCESS, 0, 2},
		{"initially empty", 0, 0, SUCCESS, SUCCESS, 0, 1},
		{"initial query error", 2, 0, ERROR_UNKNOWN, SUCCESS, 0, 1},
		{"larger count on error", 2, 3, SUCCESS, ERROR_INSUFFICIENT_SIZE, 2, 2},
	}
	queries := []struct {
		name string
		run  func() (int, Return)
	}{
		{"samples", func() (int, Return) {
			_, samples, ret := nvmlDevice{}.GetSamples(TOTAL_POWER_SAMPLES, 0)
			return len(samples), ret
		}},
		{"nearest GPUs", func() (int, Return) {
			devices, ret := nvmlDevice{}.GetTopologyNearestGpus(TOPOLOGY_SYSTEM)
			return len(devices), ret
		}},
		{"system GPU set", func() (int, Return) {
			devices, ret := (&library{}).SystemGetTopologyGpuSet(0)
			return len(devices), ret
		}},
	}
	for _, query := range queries {
		t.Run(query.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {
					calls := 0
					fill := func(count *uint32, hasBuffer bool) Return {
						calls++
						if !hasBuffer {
							*count = tc.initialCount
							return tc.firstReturn
						}
						require.Equal(t, tc.initialCount, *count)
						*count = tc.returnedCount
						return tc.secondReturn
					}
					nvmlDeviceGetSamplesStub = func(_ nvmlDevice, _ SamplingType, _ uint64, _ *ValueType, count *uint32, samples *Sample) Return {
						return fill(count, samples != nil)
					}
					nvmlDeviceGetTopologyNearestGpusStub = func(_ nvmlDevice, _ GpuTopologyLevel, count *uint32, devices *nvmlDevice) Return {
						return fill(count, devices != nil)
					}
					nvmlSystemGetTopologyGpuSetStub = func(_ uint32, count *uint32, devices *nvmlDevice) Return {
						return fill(count, devices != nil)
					}
					length, ret := query.run()
					expectedReturn := tc.secondReturn
					if tc.expectedCalls == 1 {
						expectedReturn = tc.firstReturn
					}
					require.Equal(t, expectedReturn, ret)
					require.Equal(t, tc.expectedCalls, calls)
					require.Equal(t, tc.expectedLength, length)
				})
			}
		})
	}
}
