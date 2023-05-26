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

// checks if the Dbv0037DiagStatisticsUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037DiagStatisticsUsersInner{}

// Dbv0037DiagStatisticsUsersInner Statistics by user RPCs
type Dbv0037DiagStatisticsUsersInner struct {
	// User name
	User *string `json:"user,omitempty"`
	// Number of RPCs
	Count *int32 `json:"count,omitempty"`
	Time *Dbv0037DiagStatisticsUsersInnerTime `json:"time,omitempty"`
}

// NewDbv0037DiagStatisticsUsersInner instantiates a new Dbv0037DiagStatisticsUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037DiagStatisticsUsersInner() *Dbv0037DiagStatisticsUsersInner {
	this := Dbv0037DiagStatisticsUsersInner{}
	return &this
}

// NewDbv0037DiagStatisticsUsersInnerWithDefaults instantiates a new Dbv0037DiagStatisticsUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037DiagStatisticsUsersInnerWithDefaults() *Dbv0037DiagStatisticsUsersInner {
	this := Dbv0037DiagStatisticsUsersInner{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsUsersInner) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsUsersInner) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsUsersInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Dbv0037DiagStatisticsUsersInner) SetUser(v string) {
	o.User = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsUsersInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsUsersInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsUsersInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0037DiagStatisticsUsersInner) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsUsersInner) GetTime() Dbv0037DiagStatisticsUsersInnerTime {
	if o == nil || IsNil(o.Time) {
		var ret Dbv0037DiagStatisticsUsersInnerTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsUsersInner) GetTimeOk() (*Dbv0037DiagStatisticsUsersInnerTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsUsersInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Dbv0037DiagStatisticsUsersInnerTime and assigns it to the Time field.
func (o *Dbv0037DiagStatisticsUsersInner) SetTime(v Dbv0037DiagStatisticsUsersInnerTime) {
	o.Time = &v
}

func (o Dbv0037DiagStatisticsUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037DiagStatisticsUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableDbv0037DiagStatisticsUsersInner struct {
	value *Dbv0037DiagStatisticsUsersInner
	isSet bool
}

func (v NullableDbv0037DiagStatisticsUsersInner) Get() *Dbv0037DiagStatisticsUsersInner {
	return v.value
}

func (v *NullableDbv0037DiagStatisticsUsersInner) Set(val *Dbv0037DiagStatisticsUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037DiagStatisticsUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037DiagStatisticsUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037DiagStatisticsUsersInner(val *Dbv0037DiagStatisticsUsersInner) *NullableDbv0037DiagStatisticsUsersInner {
	return &NullableDbv0037DiagStatisticsUsersInner{value: val, isSet: true}
}

func (v NullableDbv0037DiagStatisticsUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037DiagStatisticsUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

