//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The KubeVirt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kubevirt.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUFlavor) DeepCopyInto(out *CPUFlavor) {
	*out = *in
	if in.NUMA != nil {
		in, out := &in.NUMA, &out.NUMA
		*out = new(v1.NUMA)
		(*in).DeepCopyInto(*out)
	}
	if in.Realtime != nil {
		in, out := &in.Realtime, &out.Realtime
		*out = new(v1.Realtime)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUFlavor.
func (in *CPUFlavor) DeepCopy() *CPUFlavor {
	if in == nil {
		return nil
	}
	out := new(CPUFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUPreferences) DeepCopyInto(out *CPUPreferences) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUPreferences.
func (in *CPUPreferences) DeepCopy() *CPUPreferences {
	if in == nil {
		return nil
	}
	out := new(CPUPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClockPreferences) DeepCopyInto(out *ClockPreferences) {
	*out = *in
	if in.PreferredClockOffset != nil {
		in, out := &in.PreferredClockOffset, &out.PreferredClockOffset
		*out = new(v1.ClockOffset)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredTimer != nil {
		in, out := &in.PreferredTimer, &out.PreferredTimer
		*out = new(v1.Timer)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClockPreferences.
func (in *ClockPreferences) DeepCopy() *ClockPreferences {
	if in == nil {
		return nil
	}
	out := new(ClockPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePreferences) DeepCopyInto(out *DevicePreferences) {
	*out = *in
	if in.PreferredAutoattachGraphicsDevice != nil {
		in, out := &in.PreferredAutoattachGraphicsDevice, &out.PreferredAutoattachGraphicsDevice
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachMemBalloon != nil {
		in, out := &in.PreferredAutoattachMemBalloon, &out.PreferredAutoattachMemBalloon
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachPodInterface != nil {
		in, out := &in.PreferredAutoattachPodInterface, &out.PreferredAutoattachPodInterface
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachSerialConsole != nil {
		in, out := &in.PreferredAutoattachSerialConsole, &out.PreferredAutoattachSerialConsole
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDisableHotplug != nil {
		in, out := &in.PreferredDisableHotplug, &out.PreferredDisableHotplug
		*out = new(bool)
		**out = **in
	}
	if in.PreferredVirtualGPUOptions != nil {
		in, out := &in.PreferredVirtualGPUOptions, &out.PreferredVirtualGPUOptions
		*out = new(v1.VGPUOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredUseVirtioTransitional != nil {
		in, out := &in.PreferredUseVirtioTransitional, &out.PreferredUseVirtioTransitional
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDiskDedicatedIoThread != nil {
		in, out := &in.PreferredDiskDedicatedIoThread, &out.PreferredDiskDedicatedIoThread
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDiskBlockSize != nil {
		in, out := &in.PreferredDiskBlockSize, &out.PreferredDiskBlockSize
		*out = new(v1.BlockSize)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredRng != nil {
		in, out := &in.PreferredRng, &out.PreferredRng
		*out = new(v1.Rng)
		**out = **in
	}
	if in.PreferredBlockMultiQueue != nil {
		in, out := &in.PreferredBlockMultiQueue, &out.PreferredBlockMultiQueue
		*out = new(bool)
		**out = **in
	}
	if in.PreferredNetworkInterfaceMultiQueue != nil {
		in, out := &in.PreferredNetworkInterfaceMultiQueue, &out.PreferredNetworkInterfaceMultiQueue
		*out = new(bool)
		**out = **in
	}
	if in.PreferredTPM != nil {
		in, out := &in.PreferredTPM, &out.PreferredTPM
		*out = new(v1.TPMDevice)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePreferences.
func (in *DevicePreferences) DeepCopy() *DevicePreferences {
	if in == nil {
		return nil
	}
	out := new(DevicePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturePreferences) DeepCopyInto(out *FeaturePreferences) {
	*out = *in
	if in.PreferredAcpi != nil {
		in, out := &in.PreferredAcpi, &out.PreferredAcpi
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredApic != nil {
		in, out := &in.PreferredApic, &out.PreferredApic
		*out = new(v1.FeatureAPIC)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredHyperv != nil {
		in, out := &in.PreferredHyperv, &out.PreferredHyperv
		*out = new(v1.FeatureHyperv)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredKvm != nil {
		in, out := &in.PreferredKvm, &out.PreferredKvm
		*out = new(v1.FeatureKVM)
		**out = **in
	}
	if in.PreferredPvspinlock != nil {
		in, out := &in.PreferredPvspinlock, &out.PreferredPvspinlock
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredSmm != nil {
		in, out := &in.PreferredSmm, &out.PreferredSmm
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturePreferences.
func (in *FeaturePreferences) DeepCopy() *FeaturePreferences {
	if in == nil {
		return nil
	}
	out := new(FeaturePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwarePreferences) DeepCopyInto(out *FirmwarePreferences) {
	*out = *in
	if in.PreferredUseBios != nil {
		in, out := &in.PreferredUseBios, &out.PreferredUseBios
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseBiosSerial != nil {
		in, out := &in.PreferredUseBiosSerial, &out.PreferredUseBiosSerial
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseEfi != nil {
		in, out := &in.PreferredUseEfi, &out.PreferredUseEfi
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseSecureBoot != nil {
		in, out := &in.PreferredUseSecureBoot, &out.PreferredUseSecureBoot
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwarePreferences.
func (in *FirmwarePreferences) DeepCopy() *FirmwarePreferences {
	if in == nil {
		return nil
	}
	out := new(FirmwarePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePreferences) DeepCopyInto(out *MachinePreferences) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePreferences.
func (in *MachinePreferences) DeepCopy() *MachinePreferences {
	if in == nil {
		return nil
	}
	out := new(MachinePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryFlavor) DeepCopyInto(out *MemoryFlavor) {
	*out = *in
	if in.Guest != nil {
		in, out := &in.Guest, &out.Guest
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Hugepages != nil {
		in, out := &in.Hugepages, &out.Hugepages
		*out = new(v1.Hugepages)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryFlavor.
func (in *MemoryFlavor) DeepCopy() *MemoryFlavor {
	if in == nil {
		return nil
	}
	out := new(MemoryFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterFlavor) DeepCopyInto(out *VirtualMachineClusterFlavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterFlavor.
func (in *VirtualMachineClusterFlavor) DeepCopy() *VirtualMachineClusterFlavor {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterFlavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterFlavorList) DeepCopyInto(out *VirtualMachineClusterFlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineClusterFlavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterFlavorList.
func (in *VirtualMachineClusterFlavorList) DeepCopy() *VirtualMachineClusterFlavorList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterFlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterFlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterPreference) DeepCopyInto(out *VirtualMachineClusterPreference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterPreference.
func (in *VirtualMachineClusterPreference) DeepCopy() *VirtualMachineClusterPreference {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterPreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterPreference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterPreferenceList) DeepCopyInto(out *VirtualMachineClusterPreferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineClusterPreference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterPreferenceList.
func (in *VirtualMachineClusterPreferenceList) DeepCopy() *VirtualMachineClusterPreferenceList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterPreferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterPreferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFlavor) DeepCopyInto(out *VirtualMachineFlavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFlavor.
func (in *VirtualMachineFlavor) DeepCopy() *VirtualMachineFlavor {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineFlavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFlavorList) DeepCopyInto(out *VirtualMachineFlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineFlavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFlavorList.
func (in *VirtualMachineFlavorList) DeepCopy() *VirtualMachineFlavorList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineFlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFlavorSpec) DeepCopyInto(out *VirtualMachineFlavorSpec) {
	*out = *in
	in.CPU.DeepCopyInto(&out.CPU)
	in.Memory.DeepCopyInto(&out.Memory)
	if in.GPUs != nil {
		in, out := &in.GPUs, &out.GPUs
		*out = make([]v1.GPU, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostDevices != nil {
		in, out := &in.HostDevices, &out.HostDevices
		*out = make([]v1.HostDevice, len(*in))
		copy(*out, *in)
	}
	if in.IOThreadsPolicy != nil {
		in, out := &in.IOThreadsPolicy, &out.IOThreadsPolicy
		*out = new(v1.IOThreadsPolicy)
		**out = **in
	}
	if in.LaunchSecurity != nil {
		in, out := &in.LaunchSecurity, &out.LaunchSecurity
		*out = new(v1.LaunchSecurity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFlavorSpec.
func (in *VirtualMachineFlavorSpec) DeepCopy() *VirtualMachineFlavorSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFlavorSpecRevision) DeepCopyInto(out *VirtualMachineFlavorSpecRevision) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFlavorSpecRevision.
func (in *VirtualMachineFlavorSpecRevision) DeepCopy() *VirtualMachineFlavorSpecRevision {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFlavorSpecRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreference) DeepCopyInto(out *VirtualMachinePreference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreference.
func (in *VirtualMachinePreference) DeepCopy() *VirtualMachinePreference {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachinePreference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreferenceList) DeepCopyInto(out *VirtualMachinePreferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachinePreference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreferenceList.
func (in *VirtualMachinePreferenceList) DeepCopy() *VirtualMachinePreferenceList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachinePreferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreferenceSpec) DeepCopyInto(out *VirtualMachinePreferenceSpec) {
	*out = *in
	if in.Clock != nil {
		in, out := &in.Clock, &out.Clock
		*out = new(ClockPreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUPreferences)
		**out = **in
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = new(DevicePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(FeaturePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(FirmwarePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Machine != nil {
		in, out := &in.Machine, &out.Machine
		*out = new(MachinePreferences)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreferenceSpec.
func (in *VirtualMachinePreferenceSpec) DeepCopy() *VirtualMachinePreferenceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreferenceSpecRevision) DeepCopyInto(out *VirtualMachinePreferenceSpecRevision) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreferenceSpecRevision.
func (in *VirtualMachinePreferenceSpecRevision) DeepCopy() *VirtualMachinePreferenceSpecRevision {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreferenceSpecRevision)
	in.DeepCopyInto(out)
	return out
}
