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

// checks if the V0040StepTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepTask{}

// V0040StepTask struct for V0040StepTask
type V0040StepTask struct {
	Distribution *string `json:"distribution,omitempty"`
}

// NewV0040StepTask instantiates a new V0040StepTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepTask() *V0040StepTask {
	this := V0040StepTask{}
	return &this
}

// NewV0040StepTaskWithDefaults instantiates a new V0040StepTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepTaskWithDefaults() *V0040StepTask {
	this := V0040StepTask{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *V0040StepTask) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepTask) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *V0040StepTask) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *V0040StepTask) SetDistribution(v string) {
	o.Distribution = &v
}

func (o V0040StepTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	return toSerialize, nil
}

type NullableV0040StepTask struct {
	value *V0040StepTask
	isSet bool
}

func (v NullableV0040StepTask) Get() *V0040StepTask {
	return v.value
}

func (v *NullableV0040StepTask) Set(val *V0040StepTask) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepTask) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepTask(val *V0040StepTask) *NullableV0040StepTask {
	return &NullableV0040StepTask{value: val, isSet: true}
}

func (v NullableV0040StepTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


