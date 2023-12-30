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

// checks if the V0039StepNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepNodes{}

// V0039StepNodes struct for V0039StepNodes
type V0039StepNodes struct {
	List []string `json:"list,omitempty"`
}

// NewV0039StepNodes instantiates a new V0039StepNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepNodes() *V0039StepNodes {
	this := V0039StepNodes{}
	return &this
}

// NewV0039StepNodesWithDefaults instantiates a new V0039StepNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepNodesWithDefaults() *V0039StepNodes {
	this := V0039StepNodes{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V0039StepNodes) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepNodes) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V0039StepNodes) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *V0039StepNodes) SetList(v []string) {
	o.List = v
}

func (o V0039StepNodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepNodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableV0039StepNodes struct {
	value *V0039StepNodes
	isSet bool
}

func (v NullableV0039StepNodes) Get() *V0039StepNodes {
	return v.value
}

func (v *NullableV0039StepNodes) Set(val *V0039StepNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepNodes(val *V0039StepNodes) *NullableV0039StepNodes {
	return &NullableV0039StepNodes{value: val, isSet: true}
}

func (v NullableV0039StepNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


