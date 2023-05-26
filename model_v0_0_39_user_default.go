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

// checks if the V0039UserDefault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039UserDefault{}

// V0039UserDefault struct for V0039UserDefault
type V0039UserDefault struct {
	Wckey *string `json:"wckey,omitempty"`
}

// NewV0039UserDefault instantiates a new V0039UserDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039UserDefault() *V0039UserDefault {
	this := V0039UserDefault{}
	return &this
}

// NewV0039UserDefaultWithDefaults instantiates a new V0039UserDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039UserDefaultWithDefaults() *V0039UserDefault {
	this := V0039UserDefault{}
	return &this
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *V0039UserDefault) GetWckey() string {
	if o == nil || IsNil(o.Wckey) {
		var ret string
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UserDefault) GetWckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *V0039UserDefault) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given string and assigns it to the Wckey field.
func (o *V0039UserDefault) SetWckey(v string) {
	o.Wckey = &v
}

func (o V0039UserDefault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039UserDefault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	return toSerialize, nil
}

type NullableV0039UserDefault struct {
	value *V0039UserDefault
	isSet bool
}

func (v NullableV0039UserDefault) Get() *V0039UserDefault {
	return v.value
}

func (v *NullableV0039UserDefault) Set(val *V0039UserDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039UserDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039UserDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039UserDefault(val *V0039UserDefault) *NullableV0039UserDefault {
	return &NullableV0039UserDefault{value: val, isSet: true}
}

func (v NullableV0039UserDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039UserDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

