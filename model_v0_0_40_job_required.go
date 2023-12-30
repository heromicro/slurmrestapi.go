/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi/v0.0.38&openapi/dbv0.0.38&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040JobRequired type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobRequired{}

// V0040JobRequired struct for V0040JobRequired
type V0040JobRequired struct {
	CPUs *int32 `json:"CPUs,omitempty"`
	MemoryPerCpu *V0040Uint64NoVal `json:"memory_per_cpu,omitempty"`
	MemoryPerNode *V0040Uint64NoVal `json:"memory_per_node,omitempty"`
}

// NewV0040JobRequired instantiates a new V0040JobRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobRequired() *V0040JobRequired {
	this := V0040JobRequired{}
	return &this
}

// NewV0040JobRequiredWithDefaults instantiates a new V0040JobRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobRequiredWithDefaults() *V0040JobRequired {
	this := V0040JobRequired{}
	return &this
}

// GetCPUs returns the CPUs field value if set, zero value otherwise.
func (o *V0040JobRequired) GetCPUs() int32 {
	if o == nil || IsNil(o.CPUs) {
		var ret int32
		return ret
	}
	return *o.CPUs
}

// GetCPUsOk returns a tuple with the CPUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobRequired) GetCPUsOk() (*int32, bool) {
	if o == nil || IsNil(o.CPUs) {
		return nil, false
	}
	return o.CPUs, true
}

// HasCPUs returns a boolean if a field has been set.
func (o *V0040JobRequired) HasCPUs() bool {
	if o != nil && !IsNil(o.CPUs) {
		return true
	}

	return false
}

// SetCPUs gets a reference to the given int32 and assigns it to the CPUs field.
func (o *V0040JobRequired) SetCPUs(v int32) {
	o.CPUs = &v
}

// GetMemoryPerCpu returns the MemoryPerCpu field value if set, zero value otherwise.
func (o *V0040JobRequired) GetMemoryPerCpu() V0040Uint64NoVal {
	if o == nil || IsNil(o.MemoryPerCpu) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.MemoryPerCpu
}

// GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobRequired) GetMemoryPerCpuOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.MemoryPerCpu) {
		return nil, false
	}
	return o.MemoryPerCpu, true
}

// HasMemoryPerCpu returns a boolean if a field has been set.
func (o *V0040JobRequired) HasMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MemoryPerCpu) {
		return true
	}

	return false
}

// SetMemoryPerCpu gets a reference to the given V0040Uint64NoVal and assigns it to the MemoryPerCpu field.
func (o *V0040JobRequired) SetMemoryPerCpu(v V0040Uint64NoVal) {
	o.MemoryPerCpu = &v
}

// GetMemoryPerNode returns the MemoryPerNode field value if set, zero value otherwise.
func (o *V0040JobRequired) GetMemoryPerNode() V0040Uint64NoVal {
	if o == nil || IsNil(o.MemoryPerNode) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.MemoryPerNode
}

// GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobRequired) GetMemoryPerNodeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.MemoryPerNode) {
		return nil, false
	}
	return o.MemoryPerNode, true
}

// HasMemoryPerNode returns a boolean if a field has been set.
func (o *V0040JobRequired) HasMemoryPerNode() bool {
	if o != nil && !IsNil(o.MemoryPerNode) {
		return true
	}

	return false
}

// SetMemoryPerNode gets a reference to the given V0040Uint64NoVal and assigns it to the MemoryPerNode field.
func (o *V0040JobRequired) SetMemoryPerNode(v V0040Uint64NoVal) {
	o.MemoryPerNode = &v
}

func (o V0040JobRequired) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPUs) {
		toSerialize["CPUs"] = o.CPUs
	}
	if !IsNil(o.MemoryPerCpu) {
		toSerialize["memory_per_cpu"] = o.MemoryPerCpu
	}
	if !IsNil(o.MemoryPerNode) {
		toSerialize["memory_per_node"] = o.MemoryPerNode
	}
	return toSerialize, nil
}

type NullableV0040JobRequired struct {
	value *V0040JobRequired
	isSet bool
}

func (v NullableV0040JobRequired) Get() *V0040JobRequired {
	return v.value
}

func (v *NullableV0040JobRequired) Set(val *V0040JobRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobRequired(val *V0040JobRequired) *NullableV0040JobRequired {
	return &NullableV0040JobRequired{value: val, isSet: true}
}

func (v NullableV0040JobRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


