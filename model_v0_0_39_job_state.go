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

// checks if the V0039JobState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobState{}

// V0039JobState struct for V0039JobState
type V0039JobState struct {
	Current *string `json:"current,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewV0039JobState instantiates a new V0039JobState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobState() *V0039JobState {
	this := V0039JobState{}
	return &this
}

// NewV0039JobStateWithDefaults instantiates a new V0039JobState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobStateWithDefaults() *V0039JobState {
	this := V0039JobState{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *V0039JobState) GetCurrent() string {
	if o == nil || IsNil(o.Current) {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobState) GetCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *V0039JobState) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *V0039JobState) SetCurrent(v string) {
	o.Current = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V0039JobState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V0039JobState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V0039JobState) SetReason(v string) {
	o.Reason = &v
}

func (o V0039JobState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableV0039JobState struct {
	value *V0039JobState
	isSet bool
}

func (v NullableV0039JobState) Get() *V0039JobState {
	return v.value
}

func (v *NullableV0039JobState) Set(val *V0039JobState) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobState) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobState(val *V0039JobState) *NullableV0039JobState {
	return &NullableV0039JobState{value: val, isSet: true}
}

func (v NullableV0039JobState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

