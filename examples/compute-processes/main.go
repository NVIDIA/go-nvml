/**
# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
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
*/

package main

import (
	"fmt"
	"log"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	err := nvml.Init()
	if err != nil {
		log.Fatalf("Unable to initialize NVML: %v", err)
	}
	defer func() {
		err := nvml.Shutdown()
		if err != nil {
			log.Fatalf("Unable to shutdown NVML: %v", err)
		}
	}()

	count, err := nvml.DeviceGetCount()
	if err != nil {
		log.Fatalf("Unable to get device count: %v", err)
	}

	for di := 0; di < count; di++ {
		device, err := nvml.DeviceGetHandleByIndex(di)
		if err != nil {
			log.Fatalf("Unable to get device at index %d: %v", di, err)
		}

		processInfos, err := device.GetComputeRunningProcesses()
		if err != nil {
			log.Fatalf("Unable to get process info for device at index %d: %v", di, err)
		}
		fmt.Printf("Found %d processes on device %d\n", len(processInfos), di)
		for pi, processInfo := range processInfos {
			fmt.Printf("\t[%2d] ProcessInfo: %+v\n", pi, processInfo)
		}
	}
}
