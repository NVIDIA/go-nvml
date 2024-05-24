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

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Fatalf("failed to init NVML: %v", ret)
	}
	defer func() {
		_ = nvml.Shutdown()
	}()

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Fatalf("failed to get device count: %v", ret)
	}

	for i := 0; i < count; i++ {
		if err := collectGPMMetrics(i); err != nil {
			log.Printf("failed to get metrics for device %d: %v\n", i, err)
		}
	}
}

// collectGPMMetrics gets GPM metrics for a specified device.
func collectGPMMetrics(i int) error {
	device, ret := nvml.DeviceGetHandleByIndex(i)
	if ret != nvml.SUCCESS {
		return fmt.Errorf("could not get devices handle: %w", ret)
	}

	gpuQuerySupport, ret := device.GpmQueryDeviceSupport()
	if ret != nvml.SUCCESS {
		return fmt.Errorf("could not query GPM support: %w", ret)
	}

	if gpuQuerySupport.IsSupportedDevice == 0 {
		return fmt.Errorf("GPM queries are not supported")
	}

	sample1, ret := nvml.GpmSampleAlloc()
	if ret != nvml.SUCCESS {
		return fmt.Errorf("could not allocate sample: %w", ret)
	}
	defer func() {
		_ = sample1.Free()
	}()
	sample2, ret := nvml.GpmSampleAlloc()
	if ret != nvml.SUCCESS {
		return fmt.Errorf("could not allocate sample: %w", ret)
	}
	defer func() {
		_ = sample2.Free()
	}()

	if ret := device.GpmSampleGet(sample1); ret != nvml.SUCCESS {
		return fmt.Errorf("could not get sample: %w", ret)
	}
	// add a delay between samples.
	time.Sleep(1 * time.Second)
	if ret := device.GpmSampleGet(sample2); ret != nvml.SUCCESS {
		return fmt.Errorf("could not get sample: %w", ret)
	}

	gpmMetric := nvml.GpmMetricsGetType{
		NumMetrics: 1,
		Sample1:    sample1,
		Sample2:    sample2,
		Metrics: [98]nvml.GpmMetric{
			{
				MetricId: uint32(nvml.GPM_METRIC_GRAPHICS_UTIL),
			},
		},
	}

	ret = nvml.GpmMetricsGet(&gpmMetric)
	if ret != nvml.SUCCESS {
		return fmt.Errorf("failed to get gpm metric: %w", ret)
	}

	for i := 0; i < int(gpmMetric.NumMetrics); i++ {
		if gpmMetric.Metrics[i].MetricId > 0 {
			fmt.Printf("%v: %v\n", gpmMetric.Metrics[i].MetricId, gpmMetric.Metrics[i].Value)
		}
	}

	return nil
}
