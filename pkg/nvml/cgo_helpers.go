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

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

import "C"

var cgoAllocsUnknown = new(struct{})

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return i
		}
	}
	return len(n)
}

func uint32SliceToIntSlice(s []uint32) []int {
	ret := make([]int, len(s))
	for i := range s {
		ret[i] = int(s[i])
	}
	return ret
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *struct{}) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(h.Data), cgoAllocsUnknown
}

// adjustProcessInfoSlice can be used to adjust a ProcessInfo slice to account for
// differences in the structure across multiple NVML versions. This handles fields that
// were added across versions, for example.
func adjustProcessInfoSlice(in []ProcessInfo) ([]ProcessInfo, error) {
	type v1ProcessInfo struct {
		pid           uint32
		usedGpuMemory uint64
	}

	b := &bytes.Buffer{}
	err := binary.Write(b, binary.LittleEndian, in)
	if err != nil {
		return nil, fmt.Errorf("error creating temporary buffer: %v", err)
	}

	intermediate := make([]v1ProcessInfo, len(in)*2)
	err = binary.Read(b, binary.LittleEndian, intermediate)
	if err != nil {
		return nil, fmt.Errorf("error reading intermediate values: %v", err)
	}

	var out []ProcessInfo
	for i := range in {
		pin := intermediate[i]

		pout := ProcessInfo{
			Pid:           pin.pid,
			UsedGpuMemory: pin.usedGpuMemory,
		}

		out = append(out, pout)
	}

	return out, nil
}
