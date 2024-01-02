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

// checks if the V0040QosLimitsMaxActiveJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMaxActiveJobs{}

// V0040QosLimitsMaxActiveJobs struct for V0040QosLimitsMaxActiveJobs
type V0040QosLimitsMaxActiveJobs struct {
	Accruing *V0040Uint32NoVal `json:"accruing,omitempty"`
	Count *V0040Uint32NoVal `json:"count,omitempty"`
}

// NewV0040QosLimitsMaxActiveJobs instantiates a new V0040QosLimitsMaxActiveJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMaxActiveJobs() *V0040QosLimitsMaxActiveJobs {
	this := V0040QosLimitsMaxActiveJobs{}
	return &this
}

// NewV0040QosLimitsMaxActiveJobsWithDefaults instantiates a new V0040QosLimitsMaxActiveJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxActiveJobsWithDefaults() *V0040QosLimitsMaxActiveJobs {
	this := V0040QosLimitsMaxActiveJobs{}
	return &this
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxActiveJobs) GetAccruing() V0040Uint32NoVal {
	if o == nil || IsNil(o.Accruing) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxActiveJobs) GetAccruingOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxActiveJobs) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0040Uint32NoVal and assigns it to the Accruing field.
func (o *V0040QosLimitsMaxActiveJobs) SetAccruing(v V0040Uint32NoVal) {
	o.Accruing = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxActiveJobs) GetCount() V0040Uint32NoVal {
	if o == nil || IsNil(o.Count) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxActiveJobs) GetCountOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxActiveJobs) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given V0040Uint32NoVal and assigns it to the Count field.
func (o *V0040QosLimitsMaxActiveJobs) SetCount(v V0040Uint32NoVal) {
	o.Count = &v
}

func (o V0040QosLimitsMaxActiveJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMaxActiveJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMaxActiveJobs struct {
	value *V0040QosLimitsMaxActiveJobs
	isSet bool
}

func (v NullableV0040QosLimitsMaxActiveJobs) Get() *V0040QosLimitsMaxActiveJobs {
	return v.value
}

func (v *NullableV0040QosLimitsMaxActiveJobs) Set(val *V0040QosLimitsMaxActiveJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMaxActiveJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMaxActiveJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMaxActiveJobs(val *V0040QosLimitsMaxActiveJobs) *NullableV0040QosLimitsMaxActiveJobs {
	return &NullableV0040QosLimitsMaxActiveJobs{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMaxActiveJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMaxActiveJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


