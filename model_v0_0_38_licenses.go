/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi/v0.0.38&openapi/dbv0.0.38&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0038Licenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038Licenses{}

// V0038Licenses struct for V0038Licenses
type V0038Licenses struct {
	// slurm errors
	Errors []V0038Error `json:"errors,omitempty"`
	// licenses info
	Licenses []V0038License `json:"licenses,omitempty"`
}

// NewV0038Licenses instantiates a new V0038Licenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038Licenses() *V0038Licenses {
	this := V0038Licenses{}
	return &this
}

// NewV0038LicensesWithDefaults instantiates a new V0038Licenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038LicensesWithDefaults() *V0038Licenses {
	this := V0038Licenses{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0038Licenses) GetErrors() []V0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Licenses) GetErrorsOk() ([]V0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0038Licenses) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0038Error and assigns it to the Errors field.
func (o *V0038Licenses) SetErrors(v []V0038Error) {
	o.Errors = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0038Licenses) GetLicenses() []V0038License {
	if o == nil || IsNil(o.Licenses) {
		var ret []V0038License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Licenses) GetLicensesOk() ([]V0038License, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0038Licenses) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []V0038License and assigns it to the Licenses field.
func (o *V0038Licenses) SetLicenses(v []V0038License) {
	o.Licenses = v
}

func (o V0038Licenses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038Licenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableV0038Licenses struct {
	value *V0038Licenses
	isSet bool
}

func (v NullableV0038Licenses) Get() *V0038Licenses {
	return v.value
}

func (v *NullableV0038Licenses) Set(val *V0038Licenses) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038Licenses) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038Licenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038Licenses(val *V0038Licenses) *NullableV0038Licenses {
	return &NullableV0038Licenses{value: val, isSet: true}
}

func (v NullableV0038Licenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038Licenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


