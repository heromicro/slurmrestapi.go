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

// checks if the Dbv0038JobWckey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobWckey{}

// Dbv0038JobWckey Job assigned wckey details
type Dbv0038JobWckey struct {
	// Job assigned wckey
	Wckey *string `json:"wckey,omitempty"`
	// wckey flags
	Flags []string `json:"flags,omitempty"`
}

// NewDbv0038JobWckey instantiates a new Dbv0038JobWckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobWckey() *Dbv0038JobWckey {
	this := Dbv0038JobWckey{}
	return &this
}

// NewDbv0038JobWckeyWithDefaults instantiates a new Dbv0038JobWckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobWckeyWithDefaults() *Dbv0038JobWckey {
	this := Dbv0038JobWckey{}
	return &this
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *Dbv0038JobWckey) GetWckey() string {
	if o == nil || IsNil(o.Wckey) {
		var ret string
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobWckey) GetWckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *Dbv0038JobWckey) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given string and assigns it to the Wckey field.
func (o *Dbv0038JobWckey) SetWckey(v string) {
	o.Wckey = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038JobWckey) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobWckey) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038JobWckey) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038JobWckey) SetFlags(v []string) {
	o.Flags = v
}

func (o Dbv0038JobWckey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobWckey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableDbv0038JobWckey struct {
	value *Dbv0038JobWckey
	isSet bool
}

func (v NullableDbv0038JobWckey) Get() *Dbv0038JobWckey {
	return v.value
}

func (v *NullableDbv0038JobWckey) Set(val *Dbv0038JobWckey) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobWckey) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobWckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobWckey(val *Dbv0038JobWckey) *NullableDbv0038JobWckey {
	return &NullableDbv0038JobWckey{value: val, isSet: true}
}

func (v NullableDbv0038JobWckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobWckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


