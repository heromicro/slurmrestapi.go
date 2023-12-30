/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.40
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040AccountingAllocated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AccountingAllocated{}

// V0040AccountingAllocated struct for V0040AccountingAllocated
type V0040AccountingAllocated struct {
	Seconds *int64 `json:"seconds,omitempty"`
}

// NewV0040AccountingAllocated instantiates a new V0040AccountingAllocated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AccountingAllocated() *V0040AccountingAllocated {
	this := V0040AccountingAllocated{}
	return &this
}

// NewV0040AccountingAllocatedWithDefaults instantiates a new V0040AccountingAllocated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AccountingAllocatedWithDefaults() *V0040AccountingAllocated {
	this := V0040AccountingAllocated{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0040AccountingAllocated) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AccountingAllocated) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0040AccountingAllocated) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0040AccountingAllocated) SetSeconds(v int64) {
	o.Seconds = &v
}

func (o V0040AccountingAllocated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AccountingAllocated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	return toSerialize, nil
}

type NullableV0040AccountingAllocated struct {
	value *V0040AccountingAllocated
	isSet bool
}

func (v NullableV0040AccountingAllocated) Get() *V0040AccountingAllocated {
	return v.value
}

func (v *NullableV0040AccountingAllocated) Set(val *V0040AccountingAllocated) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AccountingAllocated) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AccountingAllocated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AccountingAllocated(val *V0040AccountingAllocated) *NullableV0040AccountingAllocated {
	return &NullableV0040AccountingAllocated{value: val, isSet: true}
}

func (v NullableV0040AccountingAllocated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AccountingAllocated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


