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

// checks if the V0040OpenapiAssocsRemovedResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiAssocsRemovedResp{}

// V0040OpenapiAssocsRemovedResp struct for V0040OpenapiAssocsRemovedResp
type V0040OpenapiAssocsRemovedResp struct {
	RemovedAssociations []string `json:"removed_associations"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiAssocsRemovedResp V0040OpenapiAssocsRemovedResp

// NewV0040OpenapiAssocsRemovedResp instantiates a new V0040OpenapiAssocsRemovedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiAssocsRemovedResp(removedAssociations []string) *V0040OpenapiAssocsRemovedResp {
	this := V0040OpenapiAssocsRemovedResp{}
	this.RemovedAssociations = removedAssociations
	return &this
}

// NewV0040OpenapiAssocsRemovedRespWithDefaults instantiates a new V0040OpenapiAssocsRemovedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiAssocsRemovedRespWithDefaults() *V0040OpenapiAssocsRemovedResp {
	this := V0040OpenapiAssocsRemovedResp{}
	return &this
}

// GetRemovedAssociations returns the RemovedAssociations field value
func (o *V0040OpenapiAssocsRemovedResp) GetRemovedAssociations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedAssociations
}

// GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAssocsRemovedResp) GetRemovedAssociationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedAssociations, true
}

// SetRemovedAssociations sets field value
func (o *V0040OpenapiAssocsRemovedResp) SetRemovedAssociations(v []string) {
	o.RemovedAssociations = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiAssocsRemovedResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAssocsRemovedResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiAssocsRemovedResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiAssocsRemovedResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiAssocsRemovedResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAssocsRemovedResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiAssocsRemovedResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiAssocsRemovedResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiAssocsRemovedResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAssocsRemovedResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiAssocsRemovedResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiAssocsRemovedResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiAssocsRemovedResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiAssocsRemovedResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removed_associations"] = o.RemovedAssociations
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

func (o *V0040OpenapiAssocsRemovedResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removed_associations",
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

	varV0040OpenapiAssocsRemovedResp := _V0040OpenapiAssocsRemovedResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiAssocsRemovedResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiAssocsRemovedResp(varV0040OpenapiAssocsRemovedResp)

	return err
}

type NullableV0040OpenapiAssocsRemovedResp struct {
	value *V0040OpenapiAssocsRemovedResp
	isSet bool
}

func (v NullableV0040OpenapiAssocsRemovedResp) Get() *V0040OpenapiAssocsRemovedResp {
	return v.value
}

func (v *NullableV0040OpenapiAssocsRemovedResp) Set(val *V0040OpenapiAssocsRemovedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiAssocsRemovedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiAssocsRemovedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiAssocsRemovedResp(val *V0040OpenapiAssocsRemovedResp) *NullableV0040OpenapiAssocsRemovedResp {
	return &NullableV0040OpenapiAssocsRemovedResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiAssocsRemovedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiAssocsRemovedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

