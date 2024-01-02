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

// checks if the V0040OpenapiSharesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiSharesResp{}

// V0040OpenapiSharesResp struct for V0040OpenapiSharesResp
type V0040OpenapiSharesResp struct {
	Shares V0040SharesRespMsg `json:"shares"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiSharesResp V0040OpenapiSharesResp

// NewV0040OpenapiSharesResp instantiates a new V0040OpenapiSharesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiSharesResp(shares V0040SharesRespMsg) *V0040OpenapiSharesResp {
	this := V0040OpenapiSharesResp{}
	this.Shares = shares
	return &this
}

// NewV0040OpenapiSharesRespWithDefaults instantiates a new V0040OpenapiSharesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiSharesRespWithDefaults() *V0040OpenapiSharesResp {
	this := V0040OpenapiSharesResp{}
	return &this
}

// GetShares returns the Shares field value
func (o *V0040OpenapiSharesResp) GetShares() V0040SharesRespMsg {
	if o == nil {
		var ret V0040SharesRespMsg
		return ret
	}

	return o.Shares
}

// GetSharesOk returns a tuple with the Shares field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSharesResp) GetSharesOk() (*V0040SharesRespMsg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shares, true
}

// SetShares sets field value
func (o *V0040OpenapiSharesResp) SetShares(v V0040SharesRespMsg) {
	o.Shares = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiSharesResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSharesResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiSharesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiSharesResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiSharesResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSharesResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiSharesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiSharesResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiSharesResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSharesResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiSharesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiSharesResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiSharesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiSharesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shares"] = o.Shares
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

func (o *V0040OpenapiSharesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shares",
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

	varV0040OpenapiSharesResp := _V0040OpenapiSharesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiSharesResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiSharesResp(varV0040OpenapiSharesResp)

	return err
}

type NullableV0040OpenapiSharesResp struct {
	value *V0040OpenapiSharesResp
	isSet bool
}

func (v NullableV0040OpenapiSharesResp) Get() *V0040OpenapiSharesResp {
	return v.value
}

func (v *NullableV0040OpenapiSharesResp) Set(val *V0040OpenapiSharesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiSharesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiSharesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiSharesResp(val *V0040OpenapiSharesResp) *NullableV0040OpenapiSharesResp {
	return &NullableV0040OpenapiSharesResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiSharesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiSharesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

