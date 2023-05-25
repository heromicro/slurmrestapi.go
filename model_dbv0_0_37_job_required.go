/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0037JobRequired type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobRequired{}

// Dbv0037JobRequired Job run requirements
type Dbv0037JobRequired struct {
	// Required number of CPUs
	CPUs *int32 `json:"CPUs,omitempty"`
	// Required amount of memory (MiB)
	Memory *int32 `json:"memory,omitempty"`
}

// NewDbv0037JobRequired instantiates a new Dbv0037JobRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobRequired() *Dbv0037JobRequired {
	this := Dbv0037JobRequired{}
	return &this
}

// NewDbv0037JobRequiredWithDefaults instantiates a new Dbv0037JobRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobRequiredWithDefaults() *Dbv0037JobRequired {
	this := Dbv0037JobRequired{}
	return &this
}

// GetCPUs returns the CPUs field value if set, zero value otherwise.
func (o *Dbv0037JobRequired) GetCPUs() int32 {
	if o == nil || IsNil(o.CPUs) {
		var ret int32
		return ret
	}
	return *o.CPUs
}

// GetCPUsOk returns a tuple with the CPUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobRequired) GetCPUsOk() (*int32, bool) {
	if o == nil || IsNil(o.CPUs) {
		return nil, false
	}
	return o.CPUs, true
}

// HasCPUs returns a boolean if a field has been set.
func (o *Dbv0037JobRequired) HasCPUs() bool {
	if o != nil && !IsNil(o.CPUs) {
		return true
	}

	return false
}

// SetCPUs gets a reference to the given int32 and assigns it to the CPUs field.
func (o *Dbv0037JobRequired) SetCPUs(v int32) {
	o.CPUs = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Dbv0037JobRequired) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobRequired) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Dbv0037JobRequired) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *Dbv0037JobRequired) SetMemory(v int32) {
	o.Memory = &v
}

func (o Dbv0037JobRequired) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPUs) {
		toSerialize["CPUs"] = o.CPUs
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableDbv0037JobRequired struct {
	value *Dbv0037JobRequired
	isSet bool
}

func (v NullableDbv0037JobRequired) Get() *Dbv0037JobRequired {
	return v.value
}

func (v *NullableDbv0037JobRequired) Set(val *Dbv0037JobRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobRequired(val *Dbv0037JobRequired) *NullableDbv0037JobRequired {
	return &NullableDbv0037JobRequired{value: val, isSet: true}
}

func (v NullableDbv0037JobRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


