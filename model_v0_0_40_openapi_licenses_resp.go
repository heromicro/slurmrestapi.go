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
	"bytes"
	"fmt"
)

// checks if the V0040OpenapiLicensesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiLicensesResp{}

// V0040OpenapiLicensesResp struct for V0040OpenapiLicensesResp
type V0040OpenapiLicensesResp struct {
	Licenses []V0040License `json:"licenses"`
	LastUpdate V0040Uint64NoVal `json:"last_update"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiLicensesResp V0040OpenapiLicensesResp

// NewV0040OpenapiLicensesResp instantiates a new V0040OpenapiLicensesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiLicensesResp(licenses []V0040License, lastUpdate V0040Uint64NoVal) *V0040OpenapiLicensesResp {
	this := V0040OpenapiLicensesResp{}
	this.Licenses = licenses
	this.LastUpdate = lastUpdate
	return &this
}

// NewV0040OpenapiLicensesRespWithDefaults instantiates a new V0040OpenapiLicensesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiLicensesRespWithDefaults() *V0040OpenapiLicensesResp {
	this := V0040OpenapiLicensesResp{}
	return &this
}

// GetLicenses returns the Licenses field value
func (o *V0040OpenapiLicensesResp) GetLicenses() []V0040License {
	if o == nil {
		var ret []V0040License
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiLicensesResp) GetLicensesOk() ([]V0040License, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *V0040OpenapiLicensesResp) SetLicenses(v []V0040License) {
	o.Licenses = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *V0040OpenapiLicensesResp) GetLastUpdate() V0040Uint64NoVal {
	if o == nil {
		var ret V0040Uint64NoVal
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiLicensesResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *V0040OpenapiLicensesResp) SetLastUpdate(v V0040Uint64NoVal) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiLicensesResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiLicensesResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiLicensesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiLicensesResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiLicensesResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiLicensesResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiLicensesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiLicensesResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiLicensesResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiLicensesResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiLicensesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiLicensesResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiLicensesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiLicensesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["licenses"] = o.Licenses
	toSerialize["last_update"] = o.LastUpdate
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0040OpenapiLicensesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"licenses",
		"last_update",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0040OpenapiLicensesResp := _V0040OpenapiLicensesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiLicensesResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiLicensesResp(varV0040OpenapiLicensesResp)

	return err
}

type NullableV0040OpenapiLicensesResp struct {
	value *V0040OpenapiLicensesResp
	isSet bool
}

func (v NullableV0040OpenapiLicensesResp) Get() *V0040OpenapiLicensesResp {
	return v.value
}

func (v *NullableV0040OpenapiLicensesResp) Set(val *V0040OpenapiLicensesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiLicensesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiLicensesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiLicensesResp(val *V0040OpenapiLicensesResp) *NullableV0040OpenapiLicensesResp {
	return &NullableV0040OpenapiLicensesResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiLicensesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiLicensesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


