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
	"bytes"
	"fmt"
)

// checks if the V0040OpenapiSlurmdbdQosRemovedResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiSlurmdbdQosRemovedResp{}

// V0040OpenapiSlurmdbdQosRemovedResp struct for V0040OpenapiSlurmdbdQosRemovedResp
type V0040OpenapiSlurmdbdQosRemovedResp struct {
	RemovedQos []string `json:"removed_qos"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiSlurmdbdQosRemovedResp V0040OpenapiSlurmdbdQosRemovedResp

// NewV0040OpenapiSlurmdbdQosRemovedResp instantiates a new V0040OpenapiSlurmdbdQosRemovedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiSlurmdbdQosRemovedResp(removedQos []string) *V0040OpenapiSlurmdbdQosRemovedResp {
	this := V0040OpenapiSlurmdbdQosRemovedResp{}
	this.RemovedQos = removedQos
	return &this
}

// NewV0040OpenapiSlurmdbdQosRemovedRespWithDefaults instantiates a new V0040OpenapiSlurmdbdQosRemovedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiSlurmdbdQosRemovedRespWithDefaults() *V0040OpenapiSlurmdbdQosRemovedResp {
	this := V0040OpenapiSlurmdbdQosRemovedResp{}
	return &this
}

// GetRemovedQos returns the RemovedQos field value
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetRemovedQos() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedQos
}

// GetRemovedQosOk returns a tuple with the RemovedQos field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetRemovedQosOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedQos, true
}

// SetRemovedQos sets field value
func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetRemovedQos(v []string) {
	o.RemovedQos = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiSlurmdbdQosRemovedResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiSlurmdbdQosRemovedResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removed_qos"] = o.RemovedQos
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

func (o *V0040OpenapiSlurmdbdQosRemovedResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removed_qos",
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

	varV0040OpenapiSlurmdbdQosRemovedResp := _V0040OpenapiSlurmdbdQosRemovedResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiSlurmdbdQosRemovedResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiSlurmdbdQosRemovedResp(varV0040OpenapiSlurmdbdQosRemovedResp)

	return err
}

type NullableV0040OpenapiSlurmdbdQosRemovedResp struct {
	value *V0040OpenapiSlurmdbdQosRemovedResp
	isSet bool
}

func (v NullableV0040OpenapiSlurmdbdQosRemovedResp) Get() *V0040OpenapiSlurmdbdQosRemovedResp {
	return v.value
}

func (v *NullableV0040OpenapiSlurmdbdQosRemovedResp) Set(val *V0040OpenapiSlurmdbdQosRemovedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiSlurmdbdQosRemovedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiSlurmdbdQosRemovedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiSlurmdbdQosRemovedResp(val *V0040OpenapiSlurmdbdQosRemovedResp) *NullableV0040OpenapiSlurmdbdQosRemovedResp {
	return &NullableV0040OpenapiSlurmdbdQosRemovedResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiSlurmdbdQosRemovedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiSlurmdbdQosRemovedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


