/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	resourcev1beta2 "k8s.io/api/resource/v1beta2"
)

// DeviceSubRequestApplyConfiguration represents a declarative configuration of the DeviceSubRequest type for use
// with apply.
type DeviceSubRequestApplyConfiguration struct {
	Name            *string                                 `json:"name,omitempty"`
	DeviceClassName *string                                 `json:"deviceClassName,omitempty"`
	Selectors       []DeviceSelectorApplyConfiguration      `json:"selectors,omitempty"`
	AllocationMode  *resourcev1beta2.DeviceAllocationMode   `json:"allocationMode,omitempty"`
	Count           *int64                                  `json:"count,omitempty"`
	Tolerations     []DeviceTolerationApplyConfiguration    `json:"tolerations,omitempty"`
	Capacity        *CapacityRequirementsApplyConfiguration `json:"capacity,omitempty"`
}

// DeviceSubRequestApplyConfiguration constructs a declarative configuration of the DeviceSubRequest type for use with
// apply.
func DeviceSubRequest() *DeviceSubRequestApplyConfiguration {
	return &DeviceSubRequestApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *DeviceSubRequestApplyConfiguration) WithName(value string) *DeviceSubRequestApplyConfiguration {
	b.Name = &value
	return b
}

// WithDeviceClassName sets the DeviceClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeviceClassName field is set to the value of the last call.
func (b *DeviceSubRequestApplyConfiguration) WithDeviceClassName(value string) *DeviceSubRequestApplyConfiguration {
	b.DeviceClassName = &value
	return b
}

// WithSelectors adds the given value to the Selectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Selectors field.
func (b *DeviceSubRequestApplyConfiguration) WithSelectors(values ...*DeviceSelectorApplyConfiguration) *DeviceSubRequestApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSelectors")
		}
		b.Selectors = append(b.Selectors, *values[i])
	}
	return b
}

// WithAllocationMode sets the AllocationMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocationMode field is set to the value of the last call.
func (b *DeviceSubRequestApplyConfiguration) WithAllocationMode(value resourcev1beta2.DeviceAllocationMode) *DeviceSubRequestApplyConfiguration {
	b.AllocationMode = &value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *DeviceSubRequestApplyConfiguration) WithCount(value int64) *DeviceSubRequestApplyConfiguration {
	b.Count = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *DeviceSubRequestApplyConfiguration) WithTolerations(values ...*DeviceTolerationApplyConfiguration) *DeviceSubRequestApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *DeviceSubRequestApplyConfiguration) WithCapacity(value *CapacityRequirementsApplyConfiguration) *DeviceSubRequestApplyConfiguration {
	b.Capacity = value
	return b
}
