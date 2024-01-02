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

// checks if the V0040QosLimitsMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMaxJobs{}

// V0040QosLimitsMaxJobs struct for V0040QosLimitsMaxJobs
type V0040QosLimitsMaxJobs struct {
	ActiveJobs *V0040QosLimitsMaxJobsActiveJobs `json:"active_jobs,omitempty"`
	Per *V0040QosLimitsMaxJobsActiveJobsPer `json:"per,omitempty"`
}

// NewV0040QosLimitsMaxJobs instantiates a new V0040QosLimitsMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMaxJobs() *V0040QosLimitsMaxJobs {
	this := V0040QosLimitsMaxJobs{}
	return &this
}

// NewV0040QosLimitsMaxJobsWithDefaults instantiates a new V0040QosLimitsMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxJobsWithDefaults() *V0040QosLimitsMaxJobs {
	this := V0040QosLimitsMaxJobs{}
	return &this
}

// GetActiveJobs returns the ActiveJobs field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxJobs) GetActiveJobs() V0040QosLimitsMaxJobsActiveJobs {
	if o == nil || IsNil(o.ActiveJobs) {
		var ret V0040QosLimitsMaxJobsActiveJobs
		return ret
	}
	return *o.ActiveJobs
}

// GetActiveJobsOk returns a tuple with the ActiveJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxJobs) GetActiveJobsOk() (*V0040QosLimitsMaxJobsActiveJobs, bool) {
	if o == nil || IsNil(o.ActiveJobs) {
		return nil, false
	}
	return o.ActiveJobs, true
}

// HasActiveJobs returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxJobs) HasActiveJobs() bool {
	if o != nil && !IsNil(o.ActiveJobs) {
		return true
	}

	return false
}

// SetActiveJobs gets a reference to the given V0040QosLimitsMaxJobsActiveJobs and assigns it to the ActiveJobs field.
func (o *V0040QosLimitsMaxJobs) SetActiveJobs(v V0040QosLimitsMaxJobsActiveJobs) {
	o.ActiveJobs = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxJobs) GetPer() V0040QosLimitsMaxJobsActiveJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040QosLimitsMaxJobsActiveJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxJobs) GetPerOk() (*V0040QosLimitsMaxJobsActiveJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040QosLimitsMaxJobsActiveJobsPer and assigns it to the Per field.
func (o *V0040QosLimitsMaxJobs) SetPer(v V0040QosLimitsMaxJobsActiveJobsPer) {
	o.Per = &v
}

func (o V0040QosLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveJobs) {
		toSerialize["active_jobs"] = o.ActiveJobs
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMaxJobs struct {
	value *V0040QosLimitsMaxJobs
	isSet bool
}

func (v NullableV0040QosLimitsMaxJobs) Get() *V0040QosLimitsMaxJobs {
	return v.value
}

func (v *NullableV0040QosLimitsMaxJobs) Set(val *V0040QosLimitsMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMaxJobs(val *V0040QosLimitsMaxJobs) *NullableV0040QosLimitsMaxJobs {
	return &NullableV0040QosLimitsMaxJobs{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


