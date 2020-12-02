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

package dl

import (
	"testing"
)

func TestNew(t *testing.T) {
	dl := New("libc.so", RTLD_LAZY|RTLD_GLOBAL)

	if dl == nil {
		t.Errorf("Error in New: should not return '%v'", dl)
	}
}

func TestOpenSuccess(t *testing.T) {
	dl := New("libdl.so", RTLD_LAZY|RTLD_GLOBAL)

	err := dl.Open()
	defer dl.Close()

	if err != nil {
		t.Errorf("Error opening shared lib: %v", err)
	}
}

func TestOpenFailed(t *testing.T) {
	dl := New("libbogusbadname.so", RTLD_LAZY|RTLD_GLOBAL)

	err := dl.Open()
	if err == nil {
		t.Errorf("Should have errored opening shared lib but did not")
	}
}

func TestClose(t *testing.T) {
	dl := New("libdl.so", RTLD_LAZY|RTLD_GLOBAL)

	dl.Open()
	err := dl.Close()
	if err != nil {
		t.Errorf("Error closing shared lib: %v", err)
	}
}

func TestLookupSuccess(t *testing.T) {
	dl := New("libdl.so", RTLD_LAZY|RTLD_GLOBAL)

	dl.Open()
	defer dl.Close()

	err := dl.Lookup("dlsym")
	if err != nil {
		t.Errorf("Error looking up symbol: %v", err)
	}
}

func TestLookupFailed(t *testing.T) {
	dl := New("libdl.so", RTLD_LAZY|RTLD_GLOBAL)

	dl.Open()
	defer dl.Close()

	err := dl.Lookup("bogus")
	if err == nil {
		t.Errorf("Should have errored loking up symbol but did not")
	}
}
