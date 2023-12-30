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

// checks if the Dbv0037JobArrayLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobArrayLimits{}

// Dbv0037JobArrayLimits Limits on array settings
type Dbv0037JobArrayLimits struct {
	Max *Dbv0037JobArrayLimitsMax `json:"max,omitempty"`
}

// NewDbv0037JobArrayLimits instantiates a new Dbv0037JobArrayLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobArrayLimits() *Dbv0037JobArrayLimits {
	this := Dbv0037JobArrayLimits{}
	return &this
}

// NewDbv0037JobArrayLimitsWithDefaults instantiates a new Dbv0037JobArrayLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobArrayLimitsWithDefaults() *Dbv0037JobArrayLimits {
	this := Dbv0037JobArrayLimits{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0037JobArrayLimits) GetMax() Dbv0037JobArrayLimitsMax {
	if o == nil || IsNil(o.Max) {
		var ret Dbv0037JobArrayLimitsMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobArrayLimits) GetMaxOk() (*Dbv0037JobArrayLimitsMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0037JobArrayLimits) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given Dbv0037JobArrayLimitsMax and assigns it to the Max field.
func (o *Dbv0037JobArrayLimits) SetMax(v Dbv0037JobArrayLimitsMax) {
	o.Max = &v
}

func (o Dbv0037JobArrayLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobArrayLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableDbv0037JobArrayLimits struct {
	value *Dbv0037JobArrayLimits
	isSet bool
}

func (v NullableDbv0037JobArrayLimits) Get() *Dbv0037JobArrayLimits {
	return v.value
}

func (v *NullableDbv0037JobArrayLimits) Set(val *Dbv0037JobArrayLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobArrayLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobArrayLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobArrayLimits(val *Dbv0037JobArrayLimits) *NullableDbv0037JobArrayLimits {
	return &NullableDbv0037JobArrayLimits{value: val, isSet: true}
}

func (v NullableDbv0037JobArrayLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobArrayLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


