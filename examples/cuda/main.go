/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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

	"github.com/NVIDIA/go-nvml/pkg/cuda"
)

func main() {
	err := cuda.Load()
	if err != nil {
		log.Fatalf("Error loading library: %v", err)
	}
	defer cuda.Close()
	v, r := cuda.DriverGetVersion()
	if r != cuda.SUCCESS {
		log.Fatalf("Unable to get CUDA version: %v", r)
	}

	fmt.Printf("version=%v\n", v)
}
