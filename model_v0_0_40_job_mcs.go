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
)

// checks if the V0040JobMcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobMcs{}

// V0040JobMcs struct for V0040JobMcs
type V0040JobMcs struct {
	Label *string `json:"label,omitempty"`
}

// NewV0040JobMcs instantiates a new V0040JobMcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobMcs() *V0040JobMcs {
	this := V0040JobMcs{}
	return &this
}

// NewV0040JobMcsWithDefaults instantiates a new V0040JobMcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobMcsWithDefaults() *V0040JobMcs {
	this := V0040JobMcs{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *V0040JobMcs) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobMcs) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *V0040JobMcs) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *V0040JobMcs) SetLabel(v string) {
	o.Label = &v
}

func (o V0040JobMcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobMcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableV0040JobMcs struct {
	value *V0040JobMcs
	isSet bool
}

func (v NullableV0040JobMcs) Get() *V0040JobMcs {
	return v.value
}

func (v *NullableV0040JobMcs) Set(val *V0040JobMcs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobMcs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobMcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobMcs(val *V0040JobMcs) *NullableV0040JobMcs {
	return &NullableV0040JobMcs{value: val, isSet: true}
}

func (v NullableV0040JobMcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobMcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


