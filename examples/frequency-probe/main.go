package main

import (
	"fmt"
	"log"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	// Initialize NVML
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Fatalf("Failed to initialize NVML: %v", nvml.ErrorString(ret))
	}
	defer nvml.Shutdown()

	// Get the first device
	device, ret := nvml.DeviceGetHandleByIndex(0)
	if ret != nvml.SUCCESS {
		log.Fatalf("Failed to get device handle: %v", nvml.ErrorString(ret))
	}

	// Get device name for reference
	name, ret := device.GetName()
	if ret != nvml.SUCCESS {
		log.Printf("Warning: Failed to get device name: %v", nvml.ErrorString(ret))
		name = "Unknown GPU"
	}

	fmt.Printf("GPU: %s\n", name)
	fmt.Println("Supported Clock Frequencies:")
	fmt.Println("===========================")

	// Get supported memory clocks
	memCount, memClocks, ret := device.GetSupportedMemoryClocks()
	if ret != nvml.SUCCESS {
		log.Fatalf("Failed to get supported memory clocks: %v", nvml.ErrorString(ret))
	}

	fmt.Printf("Found %d supported memory clock speeds\n", memCount)

	// Iterate over each memory clock
	for _, memClock := range memClocks {
		fmt.Printf("\nMemory Clock: %d MHz\n", memClock)
		fmt.Println("Graphics Clocks (MHz):")

		// Get supported graphics clocks for this memory clock
		graphicsCount, graphicsClocks, ret := device.GetSupportedGraphicsClocks(int(memClock))
		if ret != nvml.SUCCESS {
			log.Printf("Warning: Failed to get graphics clocks for memory clock %d MHz: %v", memClock, nvml.ErrorString(ret))
			continue
		}

		// Print each graphics clock
		for _, graphicsClock := range graphicsClocks {
			fmt.Printf("  %d MHz\n", graphicsClock)
		}
		fmt.Printf("Total graphics clocks for this memory speed: %d\n", graphicsCount)
	}

	// Get current clocks
	currentGraphicsClock, ret := device.GetClockInfo(nvml.CLOCK_GRAPHICS)
	if ret == nvml.SUCCESS {
		fmt.Printf("\nCurrent Graphics Clock: %d MHz\n", currentGraphicsClock)
	} else {
		log.Printf("Warning: Failed to get current graphics clock: %v", nvml.ErrorString(ret))
	}

	currentMemClock, ret := device.GetClockInfo(nvml.CLOCK_MEM)
	if ret == nvml.SUCCESS {
		fmt.Printf("Current Memory Clock: %d MHz\n", currentMemClock)
	} else {
		log.Printf("Warning: Failed to get current memory clock: %v", nvml.ErrorString(ret))
	}
}
