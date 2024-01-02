/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040StepTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepTasks{}

// V0040StepTasks struct for V0040StepTasks
type V0040StepTasks struct {
	Count *int32 `json:"count,omitempty"`
}

// NewV0040StepTasks instantiates a new V0040StepTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepTasks() *V0040StepTasks {
	this := V0040StepTasks{}
	return &this
}

// NewV0040StepTasksWithDefaults instantiates a new V0040StepTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepTasksWithDefaults() *V0040StepTasks {
	this := V0040StepTasks{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040StepTasks) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepTasks) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040StepTasks) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0040StepTasks) SetCount(v int32) {
	o.Count = &v
}

func (o V0040StepTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableV0040StepTasks struct {
	value *V0040StepTasks
	isSet bool
}

func (v NullableV0040StepTasks) Get() *V0040StepTasks {
	return v.value
}

func (v *NullableV0040StepTasks) Set(val *V0040StepTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepTasks(val *V0040StepTasks) *NullableV0040StepTasks {
	return &NullableV0040StepTasks{value: val, isSet: true}
}

func (v NullableV0040StepTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


