/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0037UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037UserInfo{}

// Dbv0037UserInfo struct for Dbv0037UserInfo
type Dbv0037UserInfo struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
	// Array of users
	Users []Dbv0037User `json:"users,omitempty"`
}

// NewDbv0037UserInfo instantiates a new Dbv0037UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037UserInfo() *Dbv0037UserInfo {
	this := Dbv0037UserInfo{}
	return &this
}

// NewDbv0037UserInfoWithDefaults instantiates a new Dbv0037UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037UserInfoWithDefaults() *Dbv0037UserInfo {
	this := Dbv0037UserInfo{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037UserInfo) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037UserInfo) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037UserInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037UserInfo) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0037UserInfo) GetUsers() []Dbv0037User {
	if o == nil || IsNil(o.Users) {
		var ret []Dbv0037User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037UserInfo) GetUsersOk() ([]Dbv0037User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0037UserInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Dbv0037User and assigns it to the Users field.
func (o *Dbv0037UserInfo) SetUsers(v []Dbv0037User) {
	o.Users = v
}

func (o Dbv0037UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableDbv0037UserInfo struct {
	value *Dbv0037UserInfo
	isSet bool
}

func (v NullableDbv0037UserInfo) Get() *Dbv0037UserInfo {
	return v.value
}

func (v *NullableDbv0037UserInfo) Set(val *Dbv0037UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037UserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037UserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037UserInfo(val *Dbv0037UserInfo) *NullableDbv0037UserInfo {
	return &NullableDbv0037UserInfo{value: val, isSet: true}
}

func (v NullableDbv0037UserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037UserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


