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

// checks if the V0040QosLimitsMaxJobsActiveJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMaxJobsActiveJobs{}

// V0040QosLimitsMaxJobsActiveJobs struct for V0040QosLimitsMaxJobsActiveJobs
type V0040QosLimitsMaxJobsActiveJobs struct {
	Per *V0040QosLimitsMaxJobsActiveJobsPer `json:"per,omitempty"`
}

// NewV0040QosLimitsMaxJobsActiveJobs instantiates a new V0040QosLimitsMaxJobsActiveJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMaxJobsActiveJobs() *V0040QosLimitsMaxJobsActiveJobs {
	this := V0040QosLimitsMaxJobsActiveJobs{}
	return &this
}

// NewV0040QosLimitsMaxJobsActiveJobsWithDefaults instantiates a new V0040QosLimitsMaxJobsActiveJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxJobsActiveJobsWithDefaults() *V0040QosLimitsMaxJobsActiveJobs {
	this := V0040QosLimitsMaxJobsActiveJobs{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxJobsActiveJobs) GetPer() V0040QosLimitsMaxJobsActiveJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040QosLimitsMaxJobsActiveJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxJobsActiveJobs) GetPerOk() (*V0040QosLimitsMaxJobsActiveJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxJobsActiveJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040QosLimitsMaxJobsActiveJobsPer and assigns it to the Per field.
func (o *V0040QosLimitsMaxJobsActiveJobs) SetPer(v V0040QosLimitsMaxJobsActiveJobsPer) {
	o.Per = &v
}

func (o V0040QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMaxJobsActiveJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMaxJobsActiveJobs struct {
	value *V0040QosLimitsMaxJobsActiveJobs
	isSet bool
}

func (v NullableV0040QosLimitsMaxJobsActiveJobs) Get() *V0040QosLimitsMaxJobsActiveJobs {
	return v.value
}

func (v *NullableV0040QosLimitsMaxJobsActiveJobs) Set(val *V0040QosLimitsMaxJobsActiveJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMaxJobsActiveJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMaxJobsActiveJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMaxJobsActiveJobs(val *V0040QosLimitsMaxJobsActiveJobs) *NullableV0040QosLimitsMaxJobsActiveJobs {
	return &NullableV0040QosLimitsMaxJobsActiveJobs{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMaxJobsActiveJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


