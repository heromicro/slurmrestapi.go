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

// checks if the V0040OpenapiNodesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiNodesResp{}

// V0040OpenapiNodesResp struct for V0040OpenapiNodesResp
type V0040OpenapiNodesResp struct {
	Nodes []V0040Node `json:"nodes"`
	LastUpdate V0040Uint64NoVal `json:"last_update"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiNodesResp V0040OpenapiNodesResp

// NewV0040OpenapiNodesResp instantiates a new V0040OpenapiNodesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiNodesResp(nodes []V0040Node, lastUpdate V0040Uint64NoVal) *V0040OpenapiNodesResp {
	this := V0040OpenapiNodesResp{}
	this.Nodes = nodes
	this.LastUpdate = lastUpdate
	return &this
}

// NewV0040OpenapiNodesRespWithDefaults instantiates a new V0040OpenapiNodesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiNodesRespWithDefaults() *V0040OpenapiNodesResp {
	this := V0040OpenapiNodesResp{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *V0040OpenapiNodesResp) GetNodes() []V0040Node {
	if o == nil {
		var ret []V0040Node
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiNodesResp) GetNodesOk() ([]V0040Node, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *V0040OpenapiNodesResp) SetNodes(v []V0040Node) {
	o.Nodes = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *V0040OpenapiNodesResp) GetLastUpdate() V0040Uint64NoVal {
	if o == nil {
		var ret V0040Uint64NoVal
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiNodesResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *V0040OpenapiNodesResp) SetLastUpdate(v V0040Uint64NoVal) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiNodesResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiNodesResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiNodesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiNodesResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiNodesResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiNodesResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiNodesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiNodesResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiNodesResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiNodesResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiNodesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiNodesResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiNodesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
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

func (o *V0040OpenapiNodesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodes",
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

	varV0040OpenapiNodesResp := _V0040OpenapiNodesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiNodesResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiNodesResp(varV0040OpenapiNodesResp)

	return err
}

type NullableV0040OpenapiNodesResp struct {
	value *V0040OpenapiNodesResp
	isSet bool
}

func (v NullableV0040OpenapiNodesResp) Get() *V0040OpenapiNodesResp {
	return v.value
}

func (v *NullableV0040OpenapiNodesResp) Set(val *V0040OpenapiNodesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiNodesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiNodesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiNodesResp(val *V0040OpenapiNodesResp) *NullableV0040OpenapiNodesResp {
	return &NullableV0040OpenapiNodesResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiNodesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

