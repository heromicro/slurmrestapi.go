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

// checks if the Dbv0038AssociationMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038AssociationMin{}

// Dbv0038AssociationMin Min settings
type Dbv0038AssociationMin struct {
	// Min priority threshold
	PriorityThreshold *int32 `json:"priority_threshold,omitempty"`
}

// NewDbv0038AssociationMin instantiates a new Dbv0038AssociationMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038AssociationMin() *Dbv0038AssociationMin {
	this := Dbv0038AssociationMin{}
	return &this
}

// NewDbv0038AssociationMinWithDefaults instantiates a new Dbv0038AssociationMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AssociationMinWithDefaults() *Dbv0038AssociationMin {
	this := Dbv0038AssociationMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *Dbv0038AssociationMin) GetPriorityThreshold() int32 {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret int32
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationMin) GetPriorityThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *Dbv0038AssociationMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given int32 and assigns it to the PriorityThreshold field.
func (o *Dbv0038AssociationMin) SetPriorityThreshold(v int32) {
	o.PriorityThreshold = &v
}

func (o Dbv0038AssociationMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038AssociationMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	return toSerialize, nil
}

type NullableDbv0038AssociationMin struct {
	value *Dbv0038AssociationMin
	isSet bool
}

func (v NullableDbv0038AssociationMin) Get() *Dbv0038AssociationMin {
	return v.value
}

func (v *NullableDbv0038AssociationMin) Set(val *Dbv0038AssociationMin) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038AssociationMin) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038AssociationMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038AssociationMin(val *Dbv0038AssociationMin) *NullableDbv0038AssociationMin {
	return &NullableDbv0038AssociationMin{value: val, isSet: true}
}

func (v NullableDbv0038AssociationMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038AssociationMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


