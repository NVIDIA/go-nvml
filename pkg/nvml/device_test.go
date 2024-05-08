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

func TestGetTopologyCommonAncestor(t *testing.T) {
	type wrappedDevice struct {
		Device
	}

	d1 := wrappedDevice{
		Device: nvmlDevice{},
	}

	d2 := wrappedDevice{
		Device: nvmlDevice{},
	}

	defer setNvmlDeviceGetTopologyCommonAncestorStubForTest(SUCCESS)()

	_, ret := d1.GetTopologyCommonAncestor(d2)
	require.Equal(t, SUCCESS, ret)
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
