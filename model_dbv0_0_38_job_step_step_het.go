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

// checks if the Dbv0038JobStepStepHet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStepStepHet{}

// Dbv0038JobStepStepHet Heterogeneous job details
type Dbv0038JobStepStepHet struct {
	// Parent HetJob component id
	Component *int32 `json:"component,omitempty"`
}

// NewDbv0038JobStepStepHet instantiates a new Dbv0038JobStepStepHet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStepStepHet() *Dbv0038JobStepStepHet {
	this := Dbv0038JobStepStepHet{}
	return &this
}

// NewDbv0038JobStepStepHetWithDefaults instantiates a new Dbv0038JobStepStepHet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepStepHetWithDefaults() *Dbv0038JobStepStepHet {
	this := Dbv0038JobStepStepHet{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Dbv0038JobStepStepHet) GetComponent() int32 {
	if o == nil || IsNil(o.Component) {
		var ret int32
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepStepHet) GetComponentOk() (*int32, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Dbv0038JobStepStepHet) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given int32 and assigns it to the Component field.
func (o *Dbv0038JobStepStepHet) SetComponent(v int32) {
	o.Component = &v
}

func (o Dbv0038JobStepStepHet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStepStepHet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableDbv0038JobStepStepHet struct {
	value *Dbv0038JobStepStepHet
	isSet bool
}

func (v NullableDbv0038JobStepStepHet) Get() *Dbv0038JobStepStepHet {
	return v.value
}

func (v *NullableDbv0038JobStepStepHet) Set(val *Dbv0038JobStepStepHet) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStepStepHet) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStepStepHet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStepStepHet(val *Dbv0038JobStepStepHet) *NullableDbv0038JobStepStepHet {
	return &NullableDbv0038JobStepStepHet{value: val, isSet: true}
}

func (v NullableDbv0038JobStepStepHet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStepStepHet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


