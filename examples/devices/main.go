/*
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
	"errors"
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

	for i := 0; i < count; i++ {
		device, err := nvml.DeviceGetHandleByIndex(i)
		if err != nil {
			log.Fatalf("Unable to get device at index %d: %v", i, err)
		}

		uuid, err := device.GetUUID()
		if err != nil {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, err)
		}
		fmt.Printf("%v\n", uuid)

		current, pending, err := device.GetMigMode()
		if errors.Is(err, nvml.ERROR_NOT_SUPPORTED) {
			fmt.Printf("MIG is not supported for device %v\n", i)

		} else if err != nil {
			log.Fatalf("Error getting MIG mode for device at index %d: %v", i, err)
		}
		fmt.Printf("MIG mode for device %v: current=%v pending=%v\n", i, current, pending)
	}
}
