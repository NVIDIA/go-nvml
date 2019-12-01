// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package nvml

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/dl"
)

import "C"

const (
	nvmlLibraryName      = "libnvidia-ml.so.1"
	nvmlLibraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

var nvml *dl.DynamicLibrary

// nvml.Init()
func Init() Return {
	lib := dl.New(nvmlLibraryName, nvmlLibraryLoadFlags)
	if lib == nil {
		panic(fmt.Sprintf("error instantiating DynamicLibrary for %s", nvmlLibraryName))
	}

	err := lib.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", nvmlLibraryName, err))
	}

	nvml = lib
	updateVersionedSymbols()

	return nvmlInit()
}

// nvml.InitWithFlags()
func InitWithFlags(Flags uint32) Return {
	lib := dl.New(nvmlLibraryName, nvmlLibraryLoadFlags)
	if lib == nil {
		panic(fmt.Sprintf("error instantiating DynamicLibrary for %s", nvmlLibraryName))
	}

	err := lib.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", nvmlLibraryName, err))
	}

	nvml = lib

	return nvmlInitWithFlags(Flags)
}

// nvml.Shutdown()
func Shutdown() Return {
	ret := nvmlShutdown()
	if ret != SUCCESS {
		return ret
	}

	err := nvml.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing %s: %v", nvmlLibraryName, err))
	}

	return ret
}

// Default all versioned APIs to v1 (to infer the types)
var nvmlInit                            = nvmlInit_v1
var nvmlDeviceGetPciInfo                = nvmlDeviceGetPciInfo_v1
var nvmlDeviceGetCount                  = nvmlDeviceGetCount_v1
var nvmlDeviceGetHandleByIndex          = nvmlDeviceGetHandleByIndex_v1
var nvmlDeviceGetHandleByPciBusId       = nvmlDeviceGetHandleByPciBusId_v1
var nvmlDeviceGetNvLinkRemotePciInfo    = nvmlDeviceGetNvLinkRemotePciInfo_v1
var nvmlDeviceRemoveGpu                 = nvmlDeviceRemoveGpu_v1
var nvmlDeviceGetGridLicensableFeatures = nvmlDeviceGetGridLicensableFeatures_v1
var nvmlEventSetWait                    = nvmlEventSetWait_v1

// updateVersionedSymbols()
func updateVersionedSymbols() {
	ret := nvml.Lookup("nvmlInit_v2")
    if ret == SUCCESS {
        nvmlInit = nvmlInit_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetPciInfo_v2")
    if ret == SUCCESS {
        nvmlDeviceGetPciInfo = nvmlDeviceGetPciInfo_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetPciInfo_v3")
    if ret == SUCCESS {
        nvmlDeviceGetPciInfo = nvmlDeviceGetPciInfo_v3
    }
	ret = nvml.Lookup("nvmlDeviceGetCount_v2")
    if ret == SUCCESS {
        nvmlDeviceGetCount = nvmlDeviceGetCount_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetHandleByIndex_v2")
    if ret == SUCCESS {
        nvmlDeviceGetHandleByIndex = nvmlDeviceGetHandleByIndex_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetHandleByPciBusId_v2")
    if ret == SUCCESS {
        nvmlDeviceGetHandleByPciBusId = nvmlDeviceGetHandleByPciBusId_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetNvLinkRemotePciInfo_v2")
    if ret == SUCCESS {
        nvmlDeviceGetNvLinkRemotePciInfo = nvmlDeviceGetNvLinkRemotePciInfo_v2
    }
	// Unable to overwrite nvmlDeviceRemoveGpu() because the v2 function takes
	// a different set of parameters than the v1 function.
	//ret = nvml.Lookup("nvmlDeviceRemoveGpu_v2")
    //if ret == SUCCESS {
    //    nvmlDeviceRemoveGpu = nvmlDeviceRemoveGpu_v2
    //}
	ret = nvml.Lookup("nvmlDeviceGetGridLicensableFeatures_v2")
    if ret == SUCCESS {
        nvmlDeviceGetGridLicensableFeatures = nvmlDeviceGetGridLicensableFeatures_v2
    }
	ret = nvml.Lookup("nvmlDeviceGetGridLicensableFeatures_v3")
    if ret == SUCCESS {
        nvmlDeviceGetGridLicensableFeatures = nvmlDeviceGetGridLicensableFeatures_v3
    }
	ret = nvml.Lookup("nvmlEventSetWait_v2")
    if ret == SUCCESS {
        nvmlEventSetWait = nvmlEventSetWait_v2
    }
}
