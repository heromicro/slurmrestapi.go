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

// checks if the V0039JobTimeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobTimeUser{}

// V0039JobTimeUser struct for V0039JobTimeUser
type V0039JobTimeUser struct {
	Microseconds *int64 `json:"microseconds,omitempty"`
}

// NewV0039JobTimeUser instantiates a new V0039JobTimeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobTimeUser() *V0039JobTimeUser {
	this := V0039JobTimeUser{}
	return &this
}

// NewV0039JobTimeUserWithDefaults instantiates a new V0039JobTimeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobTimeUserWithDefaults() *V0039JobTimeUser {
	this := V0039JobTimeUser{}
	return &this
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0039JobTimeUser) GetMicroseconds() int64 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int64
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobTimeUser) GetMicrosecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0039JobTimeUser) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int64 and assigns it to the Microseconds field.
func (o *V0039JobTimeUser) SetMicroseconds(v int64) {
	o.Microseconds = &v
}

func (o V0039JobTimeUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobTimeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0039JobTimeUser struct {
	value *V0039JobTimeUser
	isSet bool
}

func (v NullableV0039JobTimeUser) Get() *V0039JobTimeUser {
	return v.value
}

func (v *NullableV0039JobTimeUser) Set(val *V0039JobTimeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobTimeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobTimeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobTimeUser(val *V0039JobTimeUser) *NullableV0039JobTimeUser {
	return &NullableV0039JobTimeUser{value: val, isSet: true}
}

func (v NullableV0039JobTimeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobTimeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


