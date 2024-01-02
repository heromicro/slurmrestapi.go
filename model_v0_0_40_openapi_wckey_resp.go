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

// checks if the V0040OpenapiWckeyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiWckeyResp{}

// V0040OpenapiWckeyResp struct for V0040OpenapiWckeyResp
type V0040OpenapiWckeyResp struct {
	Wckeys []V0040Wckey `json:"wckeys"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiWckeyResp V0040OpenapiWckeyResp

// NewV0040OpenapiWckeyResp instantiates a new V0040OpenapiWckeyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiWckeyResp(wckeys []V0040Wckey) *V0040OpenapiWckeyResp {
	this := V0040OpenapiWckeyResp{}
	this.Wckeys = wckeys
	return &this
}

// NewV0040OpenapiWckeyRespWithDefaults instantiates a new V0040OpenapiWckeyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiWckeyRespWithDefaults() *V0040OpenapiWckeyResp {
	this := V0040OpenapiWckeyResp{}
	return &this
}

// GetWckeys returns the Wckeys field value
func (o *V0040OpenapiWckeyResp) GetWckeys() []V0040Wckey {
	if o == nil {
		var ret []V0040Wckey
		return ret
	}

	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWckeyResp) GetWckeysOk() ([]V0040Wckey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wckeys, true
}

// SetWckeys sets field value
func (o *V0040OpenapiWckeyResp) SetWckeys(v []V0040Wckey) {
	o.Wckeys = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiWckeyResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWckeyResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiWckeyResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiWckeyResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiWckeyResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWckeyResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiWckeyResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiWckeyResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiWckeyResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWckeyResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiWckeyResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiWckeyResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiWckeyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiWckeyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wckeys"] = o.Wckeys
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

func (o *V0040OpenapiWckeyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wckeys",
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

	varV0040OpenapiWckeyResp := _V0040OpenapiWckeyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiWckeyResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiWckeyResp(varV0040OpenapiWckeyResp)

	return err
}

type NullableV0040OpenapiWckeyResp struct {
	value *V0040OpenapiWckeyResp
	isSet bool
}

func (v NullableV0040OpenapiWckeyResp) Get() *V0040OpenapiWckeyResp {
	return v.value
}

func (v *NullableV0040OpenapiWckeyResp) Set(val *V0040OpenapiWckeyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiWckeyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiWckeyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiWckeyResp(val *V0040OpenapiWckeyResp) *NullableV0040OpenapiWckeyResp {
	return &NullableV0040OpenapiWckeyResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiWckeyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiWckeyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

