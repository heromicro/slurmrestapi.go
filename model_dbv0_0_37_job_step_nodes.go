/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0037JobStepNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobStepNodes{}

// Dbv0037JobStepNodes Node details
type Dbv0037JobStepNodes struct {
	// Total number of nodes in step
	Count *int32 `json:"count,omitempty"`
	// Nodes in step
	Range *string `json:"range,omitempty"`
}

// NewDbv0037JobStepNodes instantiates a new Dbv0037JobStepNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepNodes() *Dbv0037JobStepNodes {
	this := Dbv0037JobStepNodes{}
	return &this
}

// NewDbv0037JobStepNodesWithDefaults instantiates a new Dbv0037JobStepNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepNodesWithDefaults() *Dbv0037JobStepNodes {
	this := Dbv0037JobStepNodes{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0037JobStepNodes) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepNodes) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0037JobStepNodes) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0037JobStepNodes) SetCount(v int32) {
	o.Count = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Dbv0037JobStepNodes) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepNodes) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Dbv0037JobStepNodes) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *Dbv0037JobStepNodes) SetRange(v string) {
	o.Range = &v
}

func (o Dbv0037JobStepNodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobStepNodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	return toSerialize, nil
}

type NullableDbv0037JobStepNodes struct {
	value *Dbv0037JobStepNodes
	isSet bool
}

func (v NullableDbv0037JobStepNodes) Get() *Dbv0037JobStepNodes {
	return v.value
}

func (v *NullableDbv0037JobStepNodes) Set(val *Dbv0037JobStepNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepNodes(val *Dbv0037JobStepNodes) *NullableDbv0037JobStepNodes {
	return &NullableDbv0037JobStepNodes{value: val, isSet: true}
}

func (v NullableDbv0037JobStepNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

