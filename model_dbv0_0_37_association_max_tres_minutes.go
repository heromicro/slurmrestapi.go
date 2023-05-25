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

// checks if the Dbv0037AssociationMaxTresMinutes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037AssociationMaxTresMinutes{}

// Dbv0037AssociationMaxTresMinutes Max TRES minutes settings
type Dbv0037AssociationMaxTresMinutes struct {
	Per *Dbv0037AssociationMaxTresMinutesPer `json:"per,omitempty"`
	// TRES list of attributes
	Total []Dbv0037TresListInner `json:"total,omitempty"`
}

// NewDbv0037AssociationMaxTresMinutes instantiates a new Dbv0037AssociationMaxTresMinutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMaxTresMinutes() *Dbv0037AssociationMaxTresMinutes {
	this := Dbv0037AssociationMaxTresMinutes{}
	return &this
}

// NewDbv0037AssociationMaxTresMinutesWithDefaults instantiates a new Dbv0037AssociationMaxTresMinutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMaxTresMinutesWithDefaults() *Dbv0037AssociationMaxTresMinutes {
	this := Dbv0037AssociationMaxTresMinutes{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxTresMinutes) GetPer() Dbv0037AssociationMaxTresMinutesPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0037AssociationMaxTresMinutesPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxTresMinutes) GetPerOk() (*Dbv0037AssociationMaxTresMinutesPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxTresMinutes) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0037AssociationMaxTresMinutesPer and assigns it to the Per field.
func (o *Dbv0037AssociationMaxTresMinutes) SetPer(v Dbv0037AssociationMaxTresMinutesPer) {
	o.Per = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxTresMinutes) GetTotal() []Dbv0037TresListInner {
	if o == nil || IsNil(o.Total) {
		var ret []Dbv0037TresListInner
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxTresMinutes) GetTotalOk() ([]Dbv0037TresListInner, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxTresMinutes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []Dbv0037TresListInner and assigns it to the Total field.
func (o *Dbv0037AssociationMaxTresMinutes) SetTotal(v []Dbv0037TresListInner) {
	o.Total = v
}

func (o Dbv0037AssociationMaxTresMinutes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037AssociationMaxTresMinutes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableDbv0037AssociationMaxTresMinutes struct {
	value *Dbv0037AssociationMaxTresMinutes
	isSet bool
}

func (v NullableDbv0037AssociationMaxTresMinutes) Get() *Dbv0037AssociationMaxTresMinutes {
	return v.value
}

func (v *NullableDbv0037AssociationMaxTresMinutes) Set(val *Dbv0037AssociationMaxTresMinutes) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMaxTresMinutes) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMaxTresMinutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMaxTresMinutes(val *Dbv0037AssociationMaxTresMinutes) *NullableDbv0037AssociationMaxTresMinutes {
	return &NullableDbv0037AssociationMaxTresMinutes{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMaxTresMinutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMaxTresMinutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

