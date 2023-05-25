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

// checks if the V0039JobExitCodeSignal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobExitCodeSignal{}

// V0039JobExitCodeSignal Job exited due to signal
type V0039JobExitCodeSignal struct {
	// signal numeric ID
	SignalId *int32 `json:"signal_id,omitempty"`
	// signal name
	Name *string `json:"name,omitempty"`
}

// NewV0039JobExitCodeSignal instantiates a new V0039JobExitCodeSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobExitCodeSignal() *V0039JobExitCodeSignal {
	this := V0039JobExitCodeSignal{}
	return &this
}

// NewV0039JobExitCodeSignalWithDefaults instantiates a new V0039JobExitCodeSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobExitCodeSignalWithDefaults() *V0039JobExitCodeSignal {
	this := V0039JobExitCodeSignal{}
	return &this
}

// GetSignalId returns the SignalId field value if set, zero value otherwise.
func (o *V0039JobExitCodeSignal) GetSignalId() int32 {
	if o == nil || IsNil(o.SignalId) {
		var ret int32
		return ret
	}
	return *o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobExitCodeSignal) GetSignalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SignalId) {
		return nil, false
	}
	return o.SignalId, true
}

// HasSignalId returns a boolean if a field has been set.
func (o *V0039JobExitCodeSignal) HasSignalId() bool {
	if o != nil && !IsNil(o.SignalId) {
		return true
	}

	return false
}

// SetSignalId gets a reference to the given int32 and assigns it to the SignalId field.
func (o *V0039JobExitCodeSignal) SetSignalId(v int32) {
	o.SignalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039JobExitCodeSignal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobExitCodeSignal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039JobExitCodeSignal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039JobExitCodeSignal) SetName(v string) {
	o.Name = &v
}

func (o V0039JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobExitCodeSignal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignalId) {
		toSerialize["signal_id"] = o.SignalId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0039JobExitCodeSignal struct {
	value *V0039JobExitCodeSignal
	isSet bool
}

func (v NullableV0039JobExitCodeSignal) Get() *V0039JobExitCodeSignal {
	return v.value
}

func (v *NullableV0039JobExitCodeSignal) Set(val *V0039JobExitCodeSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobExitCodeSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobExitCodeSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobExitCodeSignal(val *V0039JobExitCodeSignal) *NullableV0039JobExitCodeSignal {
	return &NullableV0039JobExitCodeSignal{value: val, isSet: true}
}

func (v NullableV0039JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobExitCodeSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


