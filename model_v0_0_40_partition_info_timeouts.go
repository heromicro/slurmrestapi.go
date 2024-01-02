/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040PartitionInfoTimeouts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoTimeouts{}

// V0040PartitionInfoTimeouts struct for V0040PartitionInfoTimeouts
type V0040PartitionInfoTimeouts struct {
	Resume *V0040Uint16NoVal `json:"resume,omitempty"`
	Suspend *V0040Uint16NoVal `json:"suspend,omitempty"`
}

// NewV0040PartitionInfoTimeouts instantiates a new V0040PartitionInfoTimeouts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoTimeouts() *V0040PartitionInfoTimeouts {
	this := V0040PartitionInfoTimeouts{}
	return &this
}

// NewV0040PartitionInfoTimeoutsWithDefaults instantiates a new V0040PartitionInfoTimeouts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoTimeoutsWithDefaults() *V0040PartitionInfoTimeouts {
	this := V0040PartitionInfoTimeouts{}
	return &this
}

// GetResume returns the Resume field value if set, zero value otherwise.
func (o *V0040PartitionInfoTimeouts) GetResume() V0040Uint16NoVal {
	if o == nil || IsNil(o.Resume) {
		var ret V0040Uint16NoVal
		return ret
	}
	return *o.Resume
}

// GetResumeOk returns a tuple with the Resume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoTimeouts) GetResumeOk() (*V0040Uint16NoVal, bool) {
	if o == nil || IsNil(o.Resume) {
		return nil, false
	}
	return o.Resume, true
}

// HasResume returns a boolean if a field has been set.
func (o *V0040PartitionInfoTimeouts) HasResume() bool {
	if o != nil && !IsNil(o.Resume) {
		return true
	}

	return false
}

// SetResume gets a reference to the given V0040Uint16NoVal and assigns it to the Resume field.
func (o *V0040PartitionInfoTimeouts) SetResume(v V0040Uint16NoVal) {
	o.Resume = &v
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *V0040PartitionInfoTimeouts) GetSuspend() V0040Uint16NoVal {
	if o == nil || IsNil(o.Suspend) {
		var ret V0040Uint16NoVal
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoTimeouts) GetSuspendOk() (*V0040Uint16NoVal, bool) {
	if o == nil || IsNil(o.Suspend) {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *V0040PartitionInfoTimeouts) HasSuspend() bool {
	if o != nil && !IsNil(o.Suspend) {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given V0040Uint16NoVal and assigns it to the Suspend field.
func (o *V0040PartitionInfoTimeouts) SetSuspend(v V0040Uint16NoVal) {
	o.Suspend = &v
}

func (o V0040PartitionInfoTimeouts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoTimeouts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resume) {
		toSerialize["resume"] = o.Resume
	}
	if !IsNil(o.Suspend) {
		toSerialize["suspend"] = o.Suspend
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoTimeouts struct {
	value *V0040PartitionInfoTimeouts
	isSet bool
}

func (v NullableV0040PartitionInfoTimeouts) Get() *V0040PartitionInfoTimeouts {
	return v.value
}

func (v *NullableV0040PartitionInfoTimeouts) Set(val *V0040PartitionInfoTimeouts) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoTimeouts) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoTimeouts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoTimeouts(val *V0040PartitionInfoTimeouts) *NullableV0040PartitionInfoTimeouts {
	return &NullableV0040PartitionInfoTimeouts{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoTimeouts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoTimeouts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


