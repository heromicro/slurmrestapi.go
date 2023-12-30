/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.40
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040StepStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepStatistics{}

// V0040StepStatistics struct for V0040StepStatistics
type V0040StepStatistics struct {
	CPU *V0040StepStatisticsCPU `json:"CPU,omitempty"`
	Energy *V0040StepStatisticsEnergy `json:"energy,omitempty"`
}

// NewV0040StepStatistics instantiates a new V0040StepStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepStatistics() *V0040StepStatistics {
	this := V0040StepStatistics{}
	return &this
}

// NewV0040StepStatisticsWithDefaults instantiates a new V0040StepStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepStatisticsWithDefaults() *V0040StepStatistics {
	this := V0040StepStatistics{}
	return &this
}

// GetCPU returns the CPU field value if set, zero value otherwise.
func (o *V0040StepStatistics) GetCPU() V0040StepStatisticsCPU {
	if o == nil || IsNil(o.CPU) {
		var ret V0040StepStatisticsCPU
		return ret
	}
	return *o.CPU
}

// GetCPUOk returns a tuple with the CPU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepStatistics) GetCPUOk() (*V0040StepStatisticsCPU, bool) {
	if o == nil || IsNil(o.CPU) {
		return nil, false
	}
	return o.CPU, true
}

// HasCPU returns a boolean if a field has been set.
func (o *V0040StepStatistics) HasCPU() bool {
	if o != nil && !IsNil(o.CPU) {
		return true
	}

	return false
}

// SetCPU gets a reference to the given V0040StepStatisticsCPU and assigns it to the CPU field.
func (o *V0040StepStatistics) SetCPU(v V0040StepStatisticsCPU) {
	o.CPU = &v
}

// GetEnergy returns the Energy field value if set, zero value otherwise.
func (o *V0040StepStatistics) GetEnergy() V0040StepStatisticsEnergy {
	if o == nil || IsNil(o.Energy) {
		var ret V0040StepStatisticsEnergy
		return ret
	}
	return *o.Energy
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepStatistics) GetEnergyOk() (*V0040StepStatisticsEnergy, bool) {
	if o == nil || IsNil(o.Energy) {
		return nil, false
	}
	return o.Energy, true
}

// HasEnergy returns a boolean if a field has been set.
func (o *V0040StepStatistics) HasEnergy() bool {
	if o != nil && !IsNil(o.Energy) {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given V0040StepStatisticsEnergy and assigns it to the Energy field.
func (o *V0040StepStatistics) SetEnergy(v V0040StepStatisticsEnergy) {
	o.Energy = &v
}

func (o V0040StepStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPU) {
		toSerialize["CPU"] = o.CPU
	}
	if !IsNil(o.Energy) {
		toSerialize["energy"] = o.Energy
	}
	return toSerialize, nil
}

type NullableV0040StepStatistics struct {
	value *V0040StepStatistics
	isSet bool
}

func (v NullableV0040StepStatistics) Get() *V0040StepStatistics {
	return v.value
}

func (v *NullableV0040StepStatistics) Set(val *V0040StepStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepStatistics(val *V0040StepStatistics) *NullableV0040StepStatistics {
	return &NullableV0040StepStatistics{value: val, isSet: true}
}

func (v NullableV0040StepStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


