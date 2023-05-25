/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0038QosLimitsMaxTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxTres{}

// Dbv0038QosLimitsMaxTres Limits on TRES
type Dbv0038QosLimitsMaxTres struct {
	Minutes *Dbv0038QosLimitsMaxTresMinutes `json:"minutes,omitempty"`
	Per *Dbv0038QosLimitsMaxTresPer `json:"per,omitempty"`
}

// NewDbv0038QosLimitsMaxTres instantiates a new Dbv0038QosLimitsMaxTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxTres() *Dbv0038QosLimitsMaxTres {
	this := Dbv0038QosLimitsMaxTres{}
	return &this
}

// NewDbv0038QosLimitsMaxTresWithDefaults instantiates a new Dbv0038QosLimitsMaxTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxTresWithDefaults() *Dbv0038QosLimitsMaxTres {
	this := Dbv0038QosLimitsMaxTres{}
	return &this
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTres) GetMinutes() Dbv0038QosLimitsMaxTresMinutes {
	if o == nil || IsNil(o.Minutes) {
		var ret Dbv0038QosLimitsMaxTresMinutes
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTres) GetMinutesOk() (*Dbv0038QosLimitsMaxTresMinutes, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTres) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given Dbv0038QosLimitsMaxTresMinutes and assigns it to the Minutes field.
func (o *Dbv0038QosLimitsMaxTres) SetMinutes(v Dbv0038QosLimitsMaxTresMinutes) {
	o.Minutes = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTres) GetPer() Dbv0038QosLimitsMaxTresPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0038QosLimitsMaxTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTres) GetPerOk() (*Dbv0038QosLimitsMaxTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0038QosLimitsMaxTresPer and assigns it to the Per field.
func (o *Dbv0038QosLimitsMaxTres) SetPer(v Dbv0038QosLimitsMaxTresPer) {
	o.Per = &v
}

func (o Dbv0038QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxTres struct {
	value *Dbv0038QosLimitsMaxTres
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxTres) Get() *Dbv0038QosLimitsMaxTres {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxTres) Set(val *Dbv0038QosLimitsMaxTres) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxTres) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxTres(val *Dbv0038QosLimitsMaxTres) *NullableDbv0038QosLimitsMaxTres {
	return &NullableDbv0038QosLimitsMaxTres{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


