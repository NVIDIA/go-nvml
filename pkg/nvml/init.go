// Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
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

import "C"

// nvml.Init()
func (l *library) Init() error {
	if err := l.load(); err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}
	return nvmlInit().error()
}

// nvml.InitWithFlags()
func (l *library) InitWithFlags(flags uint32) error {
	if err := l.load(); err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}
	return nvmlInitWithFlags(flags).error()
}

// nvml.Shutdown()
func (l *library) Shutdown() error {
	ret := nvmlShutdown()
	if ret != nvmlSUCCESS {
		return ret.error()
	}

	err := l.close()
	if err != nil {
		return ERROR_UNKNOWN
	}

	return ret.error()
}
