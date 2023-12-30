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

// checks if the Dbv0038JobTimeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobTimeUser{}

// Dbv0038JobTimeUser User land time values
type Dbv0038JobTimeUser struct {
	// Total number of CPU-seconds used by the job in user land, in seconds
	Seconds *int32 `json:"seconds,omitempty"`
	// Total number of CPU-seconds used by the job in user land, in microseconds
	Microseconds *int32 `json:"microseconds,omitempty"`
}

// NewDbv0038JobTimeUser instantiates a new Dbv0038JobTimeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobTimeUser() *Dbv0038JobTimeUser {
	this := Dbv0038JobTimeUser{}
	return &this
}

// NewDbv0038JobTimeUserWithDefaults instantiates a new Dbv0038JobTimeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobTimeUserWithDefaults() *Dbv0038JobTimeUser {
	this := Dbv0038JobTimeUser{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *Dbv0038JobTimeUser) GetSeconds() int32 {
	if o == nil || IsNil(o.Seconds) {
		var ret int32
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobTimeUser) GetSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *Dbv0038JobTimeUser) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int32 and assigns it to the Seconds field.
func (o *Dbv0038JobTimeUser) SetSeconds(v int32) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *Dbv0038JobTimeUser) GetMicroseconds() int32 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int32
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobTimeUser) GetMicrosecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *Dbv0038JobTimeUser) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int32 and assigns it to the Microseconds field.
func (o *Dbv0038JobTimeUser) SetMicroseconds(v int32) {
	o.Microseconds = &v
}

func (o Dbv0038JobTimeUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobTimeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableDbv0038JobTimeUser struct {
	value *Dbv0038JobTimeUser
	isSet bool
}

func (v NullableDbv0038JobTimeUser) Get() *Dbv0038JobTimeUser {
	return v.value
}

func (v *NullableDbv0038JobTimeUser) Set(val *Dbv0038JobTimeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobTimeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobTimeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobTimeUser(val *Dbv0038JobTimeUser) *NullableDbv0038JobTimeUser {
	return &NullableDbv0038JobTimeUser{value: val, isSet: true}
}

func (v NullableDbv0038JobTimeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobTimeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


