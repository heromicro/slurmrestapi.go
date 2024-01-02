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

// checks if the V0040StepStatisticsCPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepStatisticsCPU{}

// V0040StepStatisticsCPU struct for V0040StepStatisticsCPU
type V0040StepStatisticsCPU struct {
	ActualFrequency *int64 `json:"actual_frequency,omitempty"`
}

// NewV0040StepStatisticsCPU instantiates a new V0040StepStatisticsCPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepStatisticsCPU() *V0040StepStatisticsCPU {
	this := V0040StepStatisticsCPU{}
	return &this
}

// NewV0040StepStatisticsCPUWithDefaults instantiates a new V0040StepStatisticsCPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepStatisticsCPUWithDefaults() *V0040StepStatisticsCPU {
	this := V0040StepStatisticsCPU{}
	return &this
}

// GetActualFrequency returns the ActualFrequency field value if set, zero value otherwise.
func (o *V0040StepStatisticsCPU) GetActualFrequency() int64 {
	if o == nil || IsNil(o.ActualFrequency) {
		var ret int64
		return ret
	}
	return *o.ActualFrequency
}

// GetActualFrequencyOk returns a tuple with the ActualFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepStatisticsCPU) GetActualFrequencyOk() (*int64, bool) {
	if o == nil || IsNil(o.ActualFrequency) {
		return nil, false
	}
	return o.ActualFrequency, true
}

// HasActualFrequency returns a boolean if a field has been set.
func (o *V0040StepStatisticsCPU) HasActualFrequency() bool {
	if o != nil && !IsNil(o.ActualFrequency) {
		return true
	}

	return false
}

// SetActualFrequency gets a reference to the given int64 and assigns it to the ActualFrequency field.
func (o *V0040StepStatisticsCPU) SetActualFrequency(v int64) {
	o.ActualFrequency = &v
}

func (o V0040StepStatisticsCPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepStatisticsCPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualFrequency) {
		toSerialize["actual_frequency"] = o.ActualFrequency
	}
	return toSerialize, nil
}

type NullableV0040StepStatisticsCPU struct {
	value *V0040StepStatisticsCPU
	isSet bool
}

func (v NullableV0040StepStatisticsCPU) Get() *V0040StepStatisticsCPU {
	return v.value
}

func (v *NullableV0040StepStatisticsCPU) Set(val *V0040StepStatisticsCPU) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepStatisticsCPU) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepStatisticsCPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepStatisticsCPU(val *V0040StepStatisticsCPU) *NullableV0040StepStatisticsCPU {
	return &NullableV0040StepStatisticsCPU{value: val, isSet: true}
}

func (v NullableV0040StepStatisticsCPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepStatisticsCPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


