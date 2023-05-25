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

// checks if the Dbv0039TresUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039TresUpdate{}

// Dbv0039TresUpdate struct for Dbv0039TresUpdate
type Dbv0039TresUpdate struct {
	Tres []V0039Tres `json:"tres,omitempty"`
}

// NewDbv0039TresUpdate instantiates a new Dbv0039TresUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039TresUpdate() *Dbv0039TresUpdate {
	this := Dbv0039TresUpdate{}
	return &this
}

// NewDbv0039TresUpdateWithDefaults instantiates a new Dbv0039TresUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039TresUpdateWithDefaults() *Dbv0039TresUpdate {
	this := Dbv0039TresUpdate{}
	return &this
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0039TresUpdate) GetTres() []V0039Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0039Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039TresUpdate) GetTresOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0039TresUpdate) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0039Tres and assigns it to the Tres field.
func (o *Dbv0039TresUpdate) SetTres(v []V0039Tres) {
	o.Tres = v
}

func (o Dbv0039TresUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039TresUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0039TresUpdate struct {
	value *Dbv0039TresUpdate
	isSet bool
}

func (v NullableDbv0039TresUpdate) Get() *Dbv0039TresUpdate {
	return v.value
}

func (v *NullableDbv0039TresUpdate) Set(val *Dbv0039TresUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039TresUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039TresUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039TresUpdate(val *Dbv0039TresUpdate) *NullableDbv0039TresUpdate {
	return &NullableDbv0039TresUpdate{value: val, isSet: true}
}

func (v NullableDbv0039TresUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039TresUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


