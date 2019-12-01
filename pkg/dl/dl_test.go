// Copyright (c) 2015-2018, NVIDIA CORPORATION. All rights reserved.

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
