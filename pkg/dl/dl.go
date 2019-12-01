// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package dl

import (
	"fmt"
	"unsafe"
)

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"

const (
	RTLD_LAZY = C.RTLD_LAZY
	RTLD_NOW = C.RTLD_NOW
	RTLD_GLOBAL = C.RTLD_GLOBAL
	RTLD_LOCAL = C.RTLD_LOCAL
	RTLD_NODELETE = C.RTLD_NODELETE
	RTLD_NOLOAD = C.RTLD_NOLOAD
	RTLD_DEEPBIND = C.RTLD_DEEPBIND
)

type DynamicLibrary struct{
	Name string
	Flags int
	handle unsafe.Pointer
}

func New(name string, flags int) *DynamicLibrary {
	return &DynamicLibrary{
		Name: name,
		Flags: flags,
		handle: nil,
    }
}

func (dl *DynamicLibrary) Open() error {
	handle := C.dlopen(C.CString(dl.Name), C.int(dl.Flags))
	if handle == C.NULL {
		return fmt.Errorf("%s", C.GoString(C.dlerror()))
	}
	dl.handle = handle
	return nil
}

func (dl *DynamicLibrary) Close() error {
	err := C.dlclose(dl.handle)
	if err != 0 {
		return fmt.Errorf("%s", C.GoString(C.dlerror()))
	}
	return nil
}

func (dl *DynamicLibrary) Lookup(symbol string) error {
	C.dlerror() // Clear out any previous errors
	C.dlsym(dl.handle, C.CString(symbol))
	err := C.dlerror()
	if unsafe.Pointer(err) == C.NULL {
		return nil
	}
	return fmt.Errorf("%s", C.GoString(err))
}
