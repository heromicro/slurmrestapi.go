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
)

// checks if the V0040SharesUint64Tres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040SharesUint64Tres{}

// V0040SharesUint64Tres struct for V0040SharesUint64Tres
type V0040SharesUint64Tres struct {
	// TRES name
	Name *string `json:"name,omitempty"`
	Value *V0040Uint64NoVal `json:"value,omitempty"`
}

// NewV0040SharesUint64Tres instantiates a new V0040SharesUint64Tres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040SharesUint64Tres() *V0040SharesUint64Tres {
	this := V0040SharesUint64Tres{}
	return &this
}

// NewV0040SharesUint64TresWithDefaults instantiates a new V0040SharesUint64Tres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040SharesUint64TresWithDefaults() *V0040SharesUint64Tres {
	this := V0040SharesUint64Tres{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040SharesUint64Tres) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040SharesUint64Tres) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040SharesUint64Tres) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040SharesUint64Tres) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V0040SharesUint64Tres) GetValue() V0040Uint64NoVal {
	if o == nil || IsNil(o.Value) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040SharesUint64Tres) GetValueOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V0040SharesUint64Tres) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given V0040Uint64NoVal and assigns it to the Value field.
func (o *V0040SharesUint64Tres) SetValue(v V0040Uint64NoVal) {
	o.Value = &v
}

func (o V0040SharesUint64Tres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040SharesUint64Tres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableV0040SharesUint64Tres struct {
	value *V0040SharesUint64Tres
	isSet bool
}

func (v NullableV0040SharesUint64Tres) Get() *V0040SharesUint64Tres {
	return v.value
}

func (v *NullableV0040SharesUint64Tres) Set(val *V0040SharesUint64Tres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040SharesUint64Tres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040SharesUint64Tres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040SharesUint64Tres(val *V0040SharesUint64Tres) *NullableV0040SharesUint64Tres {
	return &NullableV0040SharesUint64Tres{value: val, isSet: true}
}

func (v NullableV0040SharesUint64Tres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040SharesUint64Tres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


