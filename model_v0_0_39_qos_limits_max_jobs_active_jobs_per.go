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

// checks if the V0039QosLimitsMaxJobsActiveJobsPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMaxJobsActiveJobsPer{}

// V0039QosLimitsMaxJobsActiveJobsPer struct for V0039QosLimitsMaxJobsActiveJobsPer
type V0039QosLimitsMaxJobsActiveJobsPer struct {
	Account *V0039Uint32NoVal `json:"account,omitempty"`
	User *V0039Uint32NoVal `json:"user,omitempty"`
}

// NewV0039QosLimitsMaxJobsActiveJobsPer instantiates a new V0039QosLimitsMaxJobsActiveJobsPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMaxJobsActiveJobsPer() *V0039QosLimitsMaxJobsActiveJobsPer {
	this := V0039QosLimitsMaxJobsActiveJobsPer{}
	return &this
}

// NewV0039QosLimitsMaxJobsActiveJobsPerWithDefaults instantiates a new V0039QosLimitsMaxJobsActiveJobsPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMaxJobsActiveJobsPerWithDefaults() *V0039QosLimitsMaxJobsActiveJobsPer {
	this := V0039QosLimitsMaxJobsActiveJobsPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) GetAccount() V0039Uint32NoVal {
	if o == nil || IsNil(o.Account) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) GetAccountOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0039Uint32NoVal and assigns it to the Account field.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) SetAccount(v V0039Uint32NoVal) {
	o.Account = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) GetUser() V0039Uint32NoVal {
	if o == nil || IsNil(o.User) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) GetUserOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given V0039Uint32NoVal and assigns it to the User field.
func (o *V0039QosLimitsMaxJobsActiveJobsPer) SetUser(v V0039Uint32NoVal) {
	o.User = &v
}

func (o V0039QosLimitsMaxJobsActiveJobsPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMaxJobsActiveJobsPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMaxJobsActiveJobsPer struct {
	value *V0039QosLimitsMaxJobsActiveJobsPer
	isSet bool
}

func (v NullableV0039QosLimitsMaxJobsActiveJobsPer) Get() *V0039QosLimitsMaxJobsActiveJobsPer {
	return v.value
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobsPer) Set(val *V0039QosLimitsMaxJobsActiveJobsPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMaxJobsActiveJobsPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobsPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMaxJobsActiveJobsPer(val *V0039QosLimitsMaxJobsActiveJobsPer) *NullableV0039QosLimitsMaxJobsActiveJobsPer {
	return &NullableV0039QosLimitsMaxJobsActiveJobsPer{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMaxJobsActiveJobsPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMaxJobsActiveJobsPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


