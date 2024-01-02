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

// checks if the V0039StepStatisticsEnergy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepStatisticsEnergy{}

// V0039StepStatisticsEnergy struct for V0039StepStatisticsEnergy
type V0039StepStatisticsEnergy struct {
	Consumed *V0039Uint64NoVal `json:"consumed,omitempty"`
}

// NewV0039StepStatisticsEnergy instantiates a new V0039StepStatisticsEnergy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepStatisticsEnergy() *V0039StepStatisticsEnergy {
	this := V0039StepStatisticsEnergy{}
	return &this
}

// NewV0039StepStatisticsEnergyWithDefaults instantiates a new V0039StepStatisticsEnergy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepStatisticsEnergyWithDefaults() *V0039StepStatisticsEnergy {
	this := V0039StepStatisticsEnergy{}
	return &this
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *V0039StepStatisticsEnergy) GetConsumed() V0039Uint64NoVal {
	if o == nil || IsNil(o.Consumed) {
		var ret V0039Uint64NoVal
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepStatisticsEnergy) GetConsumedOk() (*V0039Uint64NoVal, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *V0039StepStatisticsEnergy) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given V0039Uint64NoVal and assigns it to the Consumed field.
func (o *V0039StepStatisticsEnergy) SetConsumed(v V0039Uint64NoVal) {
	o.Consumed = &v
}

func (o V0039StepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepStatisticsEnergy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	return toSerialize, nil
}

type NullableV0039StepStatisticsEnergy struct {
	value *V0039StepStatisticsEnergy
	isSet bool
}

func (v NullableV0039StepStatisticsEnergy) Get() *V0039StepStatisticsEnergy {
	return v.value
}

func (v *NullableV0039StepStatisticsEnergy) Set(val *V0039StepStatisticsEnergy) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepStatisticsEnergy) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepStatisticsEnergy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepStatisticsEnergy(val *V0039StepStatisticsEnergy) *NullableV0039StepStatisticsEnergy {
	return &NullableV0039StepStatisticsEnergy{value: val, isSet: true}
}

func (v NullableV0039StepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepStatisticsEnergy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


