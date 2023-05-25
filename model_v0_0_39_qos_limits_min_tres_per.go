/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039QosLimitsMinTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMinTresPer{}

// V0039QosLimitsMinTresPer struct for V0039QosLimitsMinTresPer
type V0039QosLimitsMinTresPer struct {
	Job []V0039Tres `json:"job,omitempty"`
}

// NewV0039QosLimitsMinTresPer instantiates a new V0039QosLimitsMinTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMinTresPer() *V0039QosLimitsMinTresPer {
	this := V0039QosLimitsMinTresPer{}
	return &this
}

// NewV0039QosLimitsMinTresPerWithDefaults instantiates a new V0039QosLimitsMinTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMinTresPerWithDefaults() *V0039QosLimitsMinTresPer {
	this := V0039QosLimitsMinTresPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039QosLimitsMinTresPer) GetJob() []V0039Tres {
	if o == nil || IsNil(o.Job) {
		var ret []V0039Tres
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMinTresPer) GetJobOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039QosLimitsMinTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0039Tres and assigns it to the Job field.
func (o *V0039QosLimitsMinTresPer) SetJob(v []V0039Tres) {
	o.Job = v
}

func (o V0039QosLimitsMinTresPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMinTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMinTresPer struct {
	value *V0039QosLimitsMinTresPer
	isSet bool
}

func (v NullableV0039QosLimitsMinTresPer) Get() *V0039QosLimitsMinTresPer {
	return v.value
}

func (v *NullableV0039QosLimitsMinTresPer) Set(val *V0039QosLimitsMinTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMinTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMinTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMinTresPer(val *V0039QosLimitsMinTresPer) *NullableV0039QosLimitsMinTresPer {
	return &NullableV0039QosLimitsMinTresPer{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMinTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMinTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


