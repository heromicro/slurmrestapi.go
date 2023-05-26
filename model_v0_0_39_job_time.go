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

// checks if the V0039JobTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobTime{}

// V0039JobTime struct for V0039JobTime
type V0039JobTime struct {
	User *V0039JobTimeUser `json:"user,omitempty"`
}

// NewV0039JobTime instantiates a new V0039JobTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobTime() *V0039JobTime {
	this := V0039JobTime{}
	return &this
}

// NewV0039JobTimeWithDefaults instantiates a new V0039JobTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobTimeWithDefaults() *V0039JobTime {
	this := V0039JobTime{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0039JobTime) GetUser() V0039JobTimeUser {
	if o == nil || IsNil(o.User) {
		var ret V0039JobTimeUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobTime) GetUserOk() (*V0039JobTimeUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0039JobTime) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given V0039JobTimeUser and assigns it to the User field.
func (o *V0039JobTime) SetUser(v V0039JobTimeUser) {
	o.User = &v
}

func (o V0039JobTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0039JobTime struct {
	value *V0039JobTime
	isSet bool
}

func (v NullableV0039JobTime) Get() *V0039JobTime {
	return v.value
}

func (v *NullableV0039JobTime) Set(val *V0039JobTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobTime(val *V0039JobTime) *NullableV0039JobTime {
	return &NullableV0039JobTime{value: val, isSet: true}
}

func (v NullableV0039JobTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

