/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0038NodeAllocationSockets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038NodeAllocationSockets{}

// V0038NodeAllocationSockets assignment status of each socket by numeric socket id
type V0038NodeAllocationSockets struct {
	// assignment status of each core by core id in each socket
	Cores map[string]interface{} `json:"cores,omitempty"`
}

// NewV0038NodeAllocationSockets instantiates a new V0038NodeAllocationSockets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038NodeAllocationSockets() *V0038NodeAllocationSockets {
	this := V0038NodeAllocationSockets{}
	return &this
}

// NewV0038NodeAllocationSocketsWithDefaults instantiates a new V0038NodeAllocationSockets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038NodeAllocationSocketsWithDefaults() *V0038NodeAllocationSockets {
	this := V0038NodeAllocationSockets{}
	return &this
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *V0038NodeAllocationSockets) GetCores() map[string]interface{} {
	if o == nil || IsNil(o.Cores) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038NodeAllocationSockets) GetCoresOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Cores) {
		return map[string]interface{}{}, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *V0038NodeAllocationSockets) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given map[string]interface{} and assigns it to the Cores field.
func (o *V0038NodeAllocationSockets) SetCores(v map[string]interface{}) {
	o.Cores = v
}

func (o V0038NodeAllocationSockets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038NodeAllocationSockets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	return toSerialize, nil
}

type NullableV0038NodeAllocationSockets struct {
	value *V0038NodeAllocationSockets
	isSet bool
}

func (v NullableV0038NodeAllocationSockets) Get() *V0038NodeAllocationSockets {
	return v.value
}

func (v *NullableV0038NodeAllocationSockets) Set(val *V0038NodeAllocationSockets) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038NodeAllocationSockets) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038NodeAllocationSockets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038NodeAllocationSockets(val *V0038NodeAllocationSockets) *NullableV0038NodeAllocationSockets {
	return &NullableV0038NodeAllocationSockets{value: val, isSet: true}
}

func (v NullableV0038NodeAllocationSockets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038NodeAllocationSockets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


