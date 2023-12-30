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

// checks if the V0040AssocMaxPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMaxPer{}

// V0040AssocMaxPer struct for V0040AssocMaxPer
type V0040AssocMaxPer struct {
	Account *V0040AssocMaxPerAccount `json:"account,omitempty"`
}

// NewV0040AssocMaxPer instantiates a new V0040AssocMaxPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMaxPer() *V0040AssocMaxPer {
	this := V0040AssocMaxPer{}
	return &this
}

// NewV0040AssocMaxPerWithDefaults instantiates a new V0040AssocMaxPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxPerWithDefaults() *V0040AssocMaxPer {
	this := V0040AssocMaxPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0040AssocMaxPer) GetAccount() V0040AssocMaxPerAccount {
	if o == nil || IsNil(o.Account) {
		var ret V0040AssocMaxPerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxPer) GetAccountOk() (*V0040AssocMaxPerAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0040AssocMaxPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0040AssocMaxPerAccount and assigns it to the Account field.
func (o *V0040AssocMaxPer) SetAccount(v V0040AssocMaxPerAccount) {
	o.Account = &v
}

func (o V0040AssocMaxPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMaxPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableV0040AssocMaxPer struct {
	value *V0040AssocMaxPer
	isSet bool
}

func (v NullableV0040AssocMaxPer) Get() *V0040AssocMaxPer {
	return v.value
}

func (v *NullableV0040AssocMaxPer) Set(val *V0040AssocMaxPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMaxPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMaxPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMaxPer(val *V0040AssocMaxPer) *NullableV0040AssocMaxPer {
	return &NullableV0040AssocMaxPer{value: val, isSet: true}
}

func (v NullableV0040AssocMaxPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMaxPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


