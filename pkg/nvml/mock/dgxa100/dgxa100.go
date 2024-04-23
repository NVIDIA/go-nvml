/*
 * Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dgxa100

import (
	"fmt"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock"
	"github.com/google/uuid"
)

type Server struct {
	mock.Interface
	Devices [8]nvml.Device
}
type Device struct {
	mock.Device
	UUID               string
	PciBusID           string
	Index              int
	MigMode            int
	GpuInstances       map[*GpuInstance]struct{}
	GpuInstanceCounter uint32
	MemoryInfo         nvml.Memory
}
type GpuInstance struct {
	mock.GpuInstance
	Info                   nvml.GpuInstanceInfo
	ComputeInstances       map[*ComputeInstance]struct{}
	ComputeInstanceCounter uint32
}
type ComputeInstance struct {
	mock.ComputeInstance
	Info nvml.ComputeInstanceInfo
}

var _ nvml.Interface = (*Server)(nil)
var _ nvml.Device = (*Device)(nil)
var _ nvml.GpuInstance = (*GpuInstance)(nil)
var _ nvml.ComputeInstance = (*ComputeInstance)(nil)

func New() nvml.Interface {
	return &Server{
		Devices: [8]nvml.Device{
			NewDevice(0),
			NewDevice(1),
			NewDevice(2),
			NewDevice(3),
			NewDevice(4),
			NewDevice(5),
			NewDevice(6),
			NewDevice(7),
		},
	}
}

func NewDevice(index int) nvml.Device {
	return &Device{
		UUID:               "GPU-" + uuid.New().String(),
		PciBusID:           fmt.Sprintf("0000:%02x:00.0", index),
		Index:              index,
		GpuInstances:       make(map[*GpuInstance]struct{}),
		GpuInstanceCounter: 0,
		MemoryInfo:         nvml.Memory{42949672960, 0, 0},
	}
}

func NewGpuInstance(info nvml.GpuInstanceInfo) nvml.GpuInstance {
	return &GpuInstance{
		Info:                   info,
		ComputeInstances:       make(map[*ComputeInstance]struct{}),
		ComputeInstanceCounter: 0,
	}
}

func NewComputeInstance(info nvml.ComputeInstanceInfo) nvml.ComputeInstance {
	return &ComputeInstance{
		Info: info,
	}
}

func (n *Server) Init() nvml.Return {
	return nvml.SUCCESS
}

func (n *Server) Shutdown() nvml.Return {
	return nvml.SUCCESS
}

func (n *Server) SystemGetNVMLVersion() (string, nvml.Return) {
	return "11.450.51", nvml.SUCCESS
}

func (n *Server) DeviceGetCount() (int, nvml.Return) {
	return len(n.Devices), nvml.SUCCESS
}

func (n *Server) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	if index < 0 || index >= len(n.Devices) {
		return nil, nvml.ERROR_INVALID_ARGUMENT
	}
	return n.Devices[index], nvml.SUCCESS
}

func (n *Server) DeviceGetHandleByUUID(uuid string) (nvml.Device, nvml.Return) {
	for _, d := range n.Devices {
		if uuid == d.(*Device).UUID {
			return d, nvml.SUCCESS
		}
	}
	return nil, nvml.ERROR_INVALID_ARGUMENT
}

func (n *Server) DeviceGetHandleByPciBusId(busID string) (nvml.Device, nvml.Return) {
	for _, d := range n.Devices {
		if busID == d.(*Device).PciBusID {
			return d, nvml.SUCCESS
		}
	}
	return nil, nvml.ERROR_INVALID_ARGUMENT
}

func (d *Device) GetIndex() (int, nvml.Return) {
	return d.Index, nvml.SUCCESS
}

func (d *Device) GetUUID() (string, nvml.Return) {
	return d.UUID, nvml.SUCCESS
}

func (d *Device) GetMemoryInfo() (nvml.Memory, nvml.Return) {
	return d.MemoryInfo, nvml.SUCCESS
}

func (d *Device) GetPciInfo() (nvml.PciInfo, nvml.Return) {
	p := nvml.PciInfo{
		PciDeviceId: 0x20B010DE,
	}
	return p, nvml.SUCCESS
}

func (d *Device) SetMigMode(mode int) (nvml.Return, nvml.Return) {
	d.MigMode = mode
	return nvml.SUCCESS, nvml.SUCCESS
}

func (d *Device) GetMigMode() (int, int, nvml.Return) {
	return d.MigMode, d.MigMode, nvml.SUCCESS
}

func (d *Device) GetGpuInstanceProfileInfo(giProfileId int) (nvml.GpuInstanceProfileInfo, nvml.Return) {
	if giProfileId < 0 || giProfileId >= nvml.GPU_INSTANCE_PROFILE_COUNT {
		return nvml.GpuInstanceProfileInfo{}, nvml.ERROR_INVALID_ARGUMENT
	}

	if _, exists := MIGProfiles.GpuInstanceProfiles[giProfileId]; !exists {
		return nvml.GpuInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	return MIGProfiles.GpuInstanceProfiles[giProfileId], nvml.SUCCESS
}

func (d *Device) CreateGpuInstance(info *nvml.GpuInstanceProfileInfo) (nvml.GpuInstance, nvml.Return) {
	giInfo := nvml.GpuInstanceInfo{
		Device:    d,
		Id:        d.GpuInstanceCounter,
		ProfileId: info.Id,
	}
	d.GpuInstanceCounter++
	gi := NewGpuInstance(giInfo)
	d.GpuInstances[gi.(*GpuInstance)] = struct{}{}
	return gi, nvml.SUCCESS
}

func (d *Device) CreateGpuInstanceWithPlacement(info *nvml.GpuInstanceProfileInfo, placement *nvml.GpuInstancePlacement) (nvml.GpuInstance, nvml.Return) {
	giInfo := nvml.GpuInstanceInfo{
		Device:    d,
		Id:        d.GpuInstanceCounter,
		ProfileId: info.Id,
		Placement: *placement,
	}
	d.GpuInstanceCounter++
	gi := NewGpuInstance(giInfo)
	d.GpuInstances[gi.(*GpuInstance)] = struct{}{}
	return gi, nvml.SUCCESS
}

func (d *Device) GetGpuInstances(info *nvml.GpuInstanceProfileInfo) ([]nvml.GpuInstance, nvml.Return) {
	var gis []nvml.GpuInstance
	for gi := range d.GpuInstances {
		if gi.Info.ProfileId == info.Id {
			gis = append(gis, gi)
		}
	}
	return gis, nvml.SUCCESS
}

func (gi *GpuInstance) GetInfo() (nvml.GpuInstanceInfo, nvml.Return) {
	return gi.Info, nvml.SUCCESS
}

func (gi *GpuInstance) GetComputeInstanceProfileInfo(ciProfileId int, ciEngProfileId int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
	if ciProfileId < 0 || ciProfileId >= nvml.COMPUTE_INSTANCE_PROFILE_COUNT {
		return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_INVALID_ARGUMENT
	}

	if ciEngProfileId != nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED {
		return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	giProfileId := int(gi.Info.ProfileId)

	if _, exists := MIGProfiles.ComputeInstanceProfiles[giProfileId]; !exists {
		return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	if _, exists := MIGProfiles.ComputeInstanceProfiles[giProfileId][ciProfileId]; !exists {
		return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	return MIGProfiles.ComputeInstanceProfiles[giProfileId][ciProfileId], nvml.SUCCESS
}

func (gi *GpuInstance) CreateComputeInstance(info *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
	ciInfo := nvml.ComputeInstanceInfo{
		Device:      gi.Info.Device,
		GpuInstance: gi,
		Id:          gi.ComputeInstanceCounter,
		ProfileId:   info.Id,
	}
	gi.ComputeInstanceCounter++
	ci := NewComputeInstance(ciInfo)
	gi.ComputeInstances[ci.(*ComputeInstance)] = struct{}{}
	return ci, nvml.SUCCESS
}

func (gi *GpuInstance) GetComputeInstances(info *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
	var cis []nvml.ComputeInstance
	for ci := range gi.ComputeInstances {
		if ci.Info.ProfileId == info.Id {
			cis = append(cis, ci)
		}
	}
	return cis, nvml.SUCCESS
}

func (gi *GpuInstance) Destroy() nvml.Return {
	delete(gi.Info.Device.(*Device).GpuInstances, gi)
	return nvml.SUCCESS
}

func (ci *ComputeInstance) GetInfo() (nvml.ComputeInstanceInfo, nvml.Return) {
	return ci.Info, nvml.SUCCESS
}

func (ci *ComputeInstance) Destroy() nvml.Return {
	delete(ci.Info.GpuInstance.(*GpuInstance).ComputeInstances, ci)
	return nvml.SUCCESS
}
