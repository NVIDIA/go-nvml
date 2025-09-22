# NVML Mock Framework

This package provides mock implementations of NVIDIA's NVML (NVIDIA Management Library) for testing and development purposes. The framework uses a shared factory system to define GPU configurations that can be easily extended and customized.

## Architecture

```
pkg/nvml/mock/
├── shared/
│   ├── shared.go                 # Core shared factory and types
│   └── gpus/                     # GPU configuration definitions
│       ├── a100.go              # A100 GPU variants
│       └── a30.go               # A30 GPU variants
└── dgxa100/                      # DGX A100 implementation
    ├── dgxa100.go               # Server and device implementation
    ├── gpus.go                  # Legacy A100 configurations and MIG profiles
    └── dgxa100_test.go          # Comprehensive tests
```

## Core Concepts

### Shared Factory (`shared.Config`)
Define the characteristics of individual GPU models including:
- Device properties (name, architecture, brand, PCI device ID)
- Compute capabilities (CUDA version, compute capability)
- Memory configuration
- MIG (Multi-Instance GPU) profiles and placements

### Server Configuration (`shared.ServerConfig`)
Define complete system configurations including:
- GPU configuration and count
- Driver, NVML, and CUDA versions

### MIG Profile Configuration (`shared.MIGProfileConfig`)
Define Multi-Instance GPU capabilities including:
- GPU instance profiles (slice configurations)
- Compute instance profiles
- Placement constraints and possibilities

## Usage Examples

### Basic Usage

```go
import (
    "github.com/NVIDIA/go-nvml/pkg/nvml/mock/dgxa100"
    "github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/gpus"
)

// Create default A100 system
serverA100 := dgxa100.New()   // A100-SXM4-40GB (8 GPUs)

// Create specific A100 variants
serverA100_80GB := dgxa100.NewServerWithGPU(gpus.A100_SXM4_80GB)
serverA100_PCIE := dgxa100.NewServerWithGPU(gpus.A100_PCIE_40GB)
```

### Device Creation

```go
// Create device with default configuration
device := dgxa100.NewDevice(0)

// Create device with specific GPU variant
deviceA100_80GB := dgxa100.NewDeviceWithGPU(gpus.A100_SXM4_80GB, 0)
deviceA100_PCIE := dgxa100.NewDeviceWithGPU(gpus.A100_PCIE_40GB, 1)
```

### Accessing GPU Configurations

```go
// Available GPU configurations
// A100 Family
gpus.A100_SXM4_40GB     // A100 SXM4 40GB
gpus.A100_SXM4_80GB     // A100 SXM4 80GB
gpus.A100_PCIE_40GB     // A100 PCIe 40GB
gpus.A100_PCIE_80GB     // A100 PCIe 80GB

// A30 Family
gpus.A30_PCIE_24GB      // A30 PCIe 24GB

// Inspect configurations
fmt.Printf("GPU: %s\n", gpus.A100_SXM4_80GB.Name)
fmt.Printf("Memory: %d MB\n", gpus.A100_SXM4_80GB.MemoryMB)
fmt.Printf("Architecture: %v\n", gpus.A100_SXM4_80GB.Architecture)
fmt.Printf("PCI Device ID: 0x%X\n", gpus.A100_SXM4_80GB.PciDeviceId)
```

## Available GPU Models

### A100 Family (Ampere Architecture, 108 SMs)

- **A100 SXM4 40GB** (`gpus.A100_SXM4_40GB`)
  - Form factor: SXM4
  - Memory: 40GB HBM2
  - PCI Device ID: 0x20B010DE
  - CUDA Capability: 8.0
  - SMs per slice: 14 (1-slice), 28 (2-slice), 42 (3-slice), 56 (4-slice), 98 (7-slice)
  - MIG P2P: Not supported (`IsP2pSupported: 0`)

- **A100 SXM4 80GB** (`gpus.A100_SXM4_80GB`)
  - Form factor: SXM4
  - Memory: 80GB HBM2e
  - PCI Device ID: 0x20B210DE
  - CUDA Capability: 8.0

- **A100 PCIe 40GB** (`gpus.A100_PCIE_40GB`)
  - Form factor: PCIe
  - Memory: 40GB HBM2
  - PCI Device ID: 0x20F110DE
  - CUDA Capability: 8.0

- **A100 PCIe 80GB** (`gpus.A100_PCIE_80GB`)
  - Form factor: PCIe
  - Memory: 80GB HBM2e
  - PCI Device ID: 0x20B510DE
  - CUDA Capability: 8.0

### A30 Family (Ampere Architecture, 56 SMs)

- **A30 PCIe 24GB** (`gpus.A30_PCIE_24GB`)
  - Form factor: PCIe
  - Memory: 24GB HBM2
  - PCI Device ID: 0x20B710DE
  - CUDA Capability: 8.0
  - SMs per slice: 14 (1-slice), 28 (2-slice), 56 (4-slice)
  - MIG P2P: Not supported (`IsP2pSupported: 0`)
  - MIG slices: 1, 2, 4 (no 3-slice or 7-slice support)

## Available Server Models

### DGX A100 Family

- **DGX A100 40GB** (default)
  - 8x A100 SXM4 40GB GPUs
  - Driver: 550.54.15
  - NVML: 12.550.54.15
  - CUDA: 12040

## MIG (Multi-Instance GPU) Support

All GPU configurations include comprehensive MIG profile definitions:

- **A100**: No P2P support in MIG (`IsP2pSupported: 0`)
  - Memory profiles differ between 40GB and 80GB variants
  - Supports standard NVIDIA MIG slice configurations (1, 2, 3, 4, 7 slices)
- **A30**: No P2P support in MIG (`IsP2pSupported: 0`)
  - Supports limited MIG slice configurations (1, 2, 4 slices only)
  - 56 SMs total with 14 SMs per slice

### MIG Operations

```go
// Create server with MIG support
server := dgxa100.New()
device, _ := server.DeviceGetHandleByIndex(0)

// Enable MIG mode
device.SetMigMode(1)

// Get available GPU instance profiles
profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)

// Create GPU instance
gi, ret := device.CreateGpuInstance(&profileInfo)

// Create compute instance within GPU instance
ciProfileInfo, ret := gi.GetComputeInstanceProfileInfo(
    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
    nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED
)
ci, ret := gi.CreateComputeInstance(&ciProfileInfo)
```

## Testing

The framework includes comprehensive tests covering:
- Server creation and device enumeration
- Device properties and capabilities
- MIG mode operations and lifecycle
- GPU and compute instance management
- Memory and PCI information
- Multi-device scenarios

```bash
# Run all mock tests
go test ./pkg/nvml/mock/...

# Run A100 specific tests
go test -v ./pkg/nvml/mock/dgxa100/

# Run specific test
go test -v ./pkg/nvml/mock/dgxa100/ -run TestMIGProfilesExist
```

## Extending the Framework

### Adding GPU Variants

Add new configurations to the appropriate file in `shared/gpus/`:
```go
var A100_PCIE_24GB = shared.Config{
    Name:         "NVIDIA A100-PCIE-24GB",
    Architecture: nvml.DEVICE_ARCH_AMPERE,
    Brand:        nvml.BRAND_NVIDIA,
    MemoryMB:     24576, // 24GB
    CudaMajor:    8,
    CudaMinor:    0,
    PciDeviceId:  0x20F010DE,
    MIGProfiles:  a100_24gb_MIGProfiles,
}
```

### Adding GPU Generations

1. **Create new package** (e.g., `dgxa100/`)
2. **Define GPU configurations** in `shared/gpus/a100.go`
3. **Define MIG profiles** with appropriate memory and SM allocations
4. **Implement server and device factory functions**
5. **Add comprehensive tests**

Example structure for A100 generation:
```go
// In shared/gpus/a100.go
var A100_SXM4_80GB = shared.Config{
    Name:         "NVIDIA A100 SXM4 80GB",
    Architecture: nvml.DEVICE_ARCH_AMPERE,
    Brand:        nvml.BRAND_NVIDIA,
    MemoryMB:     81920,
    CudaMajor:    8,
    CudaMinor:    0,
    PciDeviceId:  0x20B210DE,
    MIGProfiles:  a100MIGProfiles,
}

// In dgxa100/dgxa100.go
func New() *Server {
    return shared.NewServerFromConfig(shared.ServerConfig{
        Config:            gpus.A100_SXM4_80GB,
        GPUCount:          4,
        DriverVersion:     "550.54.15",
        NvmlVersion:       "12.550.54.15",
        CudaDriverVersion: 12040,
    })
}
```

## Backward Compatibility

The framework maintains full backward compatibility:
- All existing `dgxa100.New()` calls continue to work unchanged
- Legacy global variables (`MIGProfiles`, `MIGPlacements`) are preserved
- Device names maintain "Mock" prefix for test compatibility
- All existing tests pass without modification
- A100 configurations now reference `shared/gpus` package

## Performance Considerations

- Configurations are defined as static variables (no runtime overhead)
- Device creation uses shared factory (fast)
- MIG profiles are shared between devices of the same type
- Mock functions use direct field access (minimal latency)

## Implementation Notes

- **Thread Safety**: Device implementations include proper mutex usage
- **Memory Management**: No memory leaks in device/instance lifecycle
- **Error Handling**: Proper NVML return codes for all operations
- **Standards Compliance**: Follows official NVML API patterns and behaviors
- **Separation of Concerns**: GPU configs in `shared/gpus`, server logic in package-specific files
