// Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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

// nvml.GpmMetricsGet()
type GpmMetricsGetVType struct {
	metricsGet *GpmMetricsGetType
}

func (l *library) GpmMetricsGetV(MetricsGet *GpmMetricsGetType) GpmMetricsGetVType {
	return GpmMetricsGetVType{MetricsGet}
}
func (MetricsGetV GpmMetricsGetVType) V1() Return {
	MetricsGetV.metricsGet.Version = 1
	return nvmlGpmMetricsGet(MetricsGetV.metricsGet)
}

func (l *library) GpmMetricsGet(MetricsGet *GpmMetricsGetType) Return {
	MetricsGet.Version = GPM_METRICS_GET_VERSION
	return nvmlGpmMetricsGet(MetricsGet)
}

// nvml.GpmSampleFree()
func (l *library) GpmSampleFree(GpmSample GpmSample) Return {
	return nvmlGpmSampleFree(GpmSample)
}

func (GpmSample GpmSample) Free() Return {
	return GpmSampleFree(GpmSample)
}

// nvml.GpmSampleAlloc()
func (l *library) GpmSampleAlloc() (GpmSample, Return) {
	var GpmSample GpmSample
	ret := nvmlGpmSampleAlloc(&GpmSample)
	return GpmSample, ret
}

// nvml.GpmSampleGet()
func (l *library) GpmSampleGet(Device Device, GpmSample GpmSample) Return {
	return GpmSample.Get(Device)
}

func (Device nvmlDevice) GpmSampleGet(GpmSample GpmSample) Return {
	return GpmSample.Get(Device)
}

func (GpmSample GpmSample) Get(Device Device) Return {
	return nvmlGpmSampleGet(Device.(nvmlDevice), GpmSample)
}

// nvml.GpmQueryDeviceSupport()
type GpmSupportV struct {
	device nvmlDevice
}

func (l *library) GpmQueryDeviceSupportV(Device Device) GpmSupportV {
	return Device.GpmQueryDeviceSupportV()
}

func (Device nvmlDevice) GpmQueryDeviceSupportV() GpmSupportV {
	return GpmSupportV{Device}
}

func (GpmSupportV GpmSupportV) V1() (GpmSupport, Return) {
	var GpmSupport GpmSupport
	GpmSupport.Version = 1
	ret := nvmlGpmQueryDeviceSupport(GpmSupportV.device, &GpmSupport)
	return GpmSupport, ret
}

func (l *library) GpmQueryDeviceSupport(Device Device) (GpmSupport, Return) {
	return Device.GpmQueryDeviceSupport()
}

func (Device nvmlDevice) GpmQueryDeviceSupport() (GpmSupport, Return) {
	var GpmSupport GpmSupport
	GpmSupport.Version = GPM_SUPPORT_VERSION
	ret := nvmlGpmQueryDeviceSupport(Device, &GpmSupport)
	return GpmSupport, ret
}

// nvml.GpmMigSampleGet()
func (l *library) GpmMigSampleGet(Device Device, GpuInstanceId int, GpmSample GpmSample) Return {
	return GpmSample.MigGet(Device, GpuInstanceId)
}

func (Device nvmlDevice) GpmMigSampleGet(GpuInstanceId int, GpmSample GpmSample) Return {
	return GpmSample.MigGet(Device, GpuInstanceId)
}

func (GpmSample GpmSample) MigGet(Device Device, GpuInstanceId int) Return {
	return nvmlGpmMigSampleGet(Device.(nvmlDevice), uint32(GpuInstanceId), GpmSample)
}
