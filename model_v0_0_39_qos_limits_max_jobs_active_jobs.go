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

// checks if the V0039QosLimitsMaxJobsActiveJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMaxJobsActiveJobs{}

// V0039QosLimitsMaxJobsActiveJobs struct for V0039QosLimitsMaxJobsActiveJobs
type V0039QosLimitsMaxJobsActiveJobs struct {
	Per *V0039QosLimitsMaxJobsActiveJobsPer `json:"per,omitempty"`
}

// NewV0039QosLimitsMaxJobsActiveJobs instantiates a new V0039QosLimitsMaxJobsActiveJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMaxJobsActiveJobs() *V0039QosLimitsMaxJobsActiveJobs {
	this := V0039QosLimitsMaxJobsActiveJobs{}
	return &this
}

// NewV0039QosLimitsMaxJobsActiveJobsWithDefaults instantiates a new V0039QosLimitsMaxJobsActiveJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMaxJobsActiveJobsWithDefaults() *V0039QosLimitsMaxJobsActiveJobs {
	this := V0039QosLimitsMaxJobsActiveJobs{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxJobsActiveJobs) GetPer() V0039QosLimitsMaxJobsActiveJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0039QosLimitsMaxJobsActiveJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxJobsActiveJobs) GetPerOk() (*V0039QosLimitsMaxJobsActiveJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxJobsActiveJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0039QosLimitsMaxJobsActiveJobsPer and assigns it to the Per field.
func (o *V0039QosLimitsMaxJobsActiveJobs) SetPer(v V0039QosLimitsMaxJobsActiveJobsPer) {
	o.Per = &v
}

func (o V0039QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMaxJobsActiveJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMaxJobsActiveJobs struct {
	value *V0039QosLimitsMaxJobsActiveJobs
	isSet bool
}

func (v NullableV0039QosLimitsMaxJobsActiveJobs) Get() *V0039QosLimitsMaxJobsActiveJobs {
	return v.value
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobs) Set(val *V0039QosLimitsMaxJobsActiveJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMaxJobsActiveJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMaxJobsActiveJobs(val *V0039QosLimitsMaxJobsActiveJobs) *NullableV0039QosLimitsMaxJobsActiveJobs {
	return &NullableV0039QosLimitsMaxJobsActiveJobs{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


