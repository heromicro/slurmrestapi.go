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

// checks if the V0039LicensesInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039LicensesInfo{}

// V0039LicensesInfo struct for V0039LicensesInfo
type V0039LicensesInfo struct {
	Meta *V0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []V0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []V0039Warning `json:"warnings,omitempty"`
	Licenses []V0039License `json:"licenses,omitempty"`
}

// NewV0039LicensesInfo instantiates a new V0039LicensesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039LicensesInfo() *V0039LicensesInfo {
	this := V0039LicensesInfo{}
	return &this
}

// NewV0039LicensesInfoWithDefaults instantiates a new V0039LicensesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039LicensesInfoWithDefaults() *V0039LicensesInfo {
	this := V0039LicensesInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0039LicensesInfo) GetMeta() V0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039LicensesInfo) GetMetaOk() (*V0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0039LicensesInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0039Meta and assigns it to the Meta field.
func (o *V0039LicensesInfo) SetMeta(v V0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0039LicensesInfo) GetErrors() []V0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039LicensesInfo) GetErrorsOk() ([]V0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0039LicensesInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0039Error and assigns it to the Errors field.
func (o *V0039LicensesInfo) SetErrors(v []V0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0039LicensesInfo) GetWarnings() []V0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039LicensesInfo) GetWarningsOk() ([]V0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0039LicensesInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0039Warning and assigns it to the Warnings field.
func (o *V0039LicensesInfo) SetWarnings(v []V0039Warning) {
	o.Warnings = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0039LicensesInfo) GetLicenses() []V0039License {
	if o == nil || IsNil(o.Licenses) {
		var ret []V0039License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039LicensesInfo) GetLicensesOk() ([]V0039License, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0039LicensesInfo) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []V0039License and assigns it to the Licenses field.
func (o *V0039LicensesInfo) SetLicenses(v []V0039License) {
	o.Licenses = v
}

func (o V0039LicensesInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039LicensesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableV0039LicensesInfo struct {
	value *V0039LicensesInfo
	isSet bool
}

func (v NullableV0039LicensesInfo) Get() *V0039LicensesInfo {
	return v.value
}

func (v *NullableV0039LicensesInfo) Set(val *V0039LicensesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039LicensesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039LicensesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039LicensesInfo(val *V0039LicensesInfo) *NullableV0039LicensesInfo {
	return &NullableV0039LicensesInfo{value: val, isSet: true}
}

func (v NullableV0039LicensesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039LicensesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


