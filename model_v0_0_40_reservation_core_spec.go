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

// checks if the V0040ReservationCoreSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ReservationCoreSpec{}

// V0040ReservationCoreSpec struct for V0040ReservationCoreSpec
type V0040ReservationCoreSpec struct {
	Node *string `json:"node,omitempty"`
	Core *string `json:"core,omitempty"`
}

// NewV0040ReservationCoreSpec instantiates a new V0040ReservationCoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ReservationCoreSpec() *V0040ReservationCoreSpec {
	this := V0040ReservationCoreSpec{}
	return &this
}

// NewV0040ReservationCoreSpecWithDefaults instantiates a new V0040ReservationCoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ReservationCoreSpecWithDefaults() *V0040ReservationCoreSpec {
	this := V0040ReservationCoreSpec{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *V0040ReservationCoreSpec) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationCoreSpec) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *V0040ReservationCoreSpec) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *V0040ReservationCoreSpec) SetNode(v string) {
	o.Node = &v
}

// GetCore returns the Core field value if set, zero value otherwise.
func (o *V0040ReservationCoreSpec) GetCore() string {
	if o == nil || IsNil(o.Core) {
		var ret string
		return ret
	}
	return *o.Core
}

// GetCoreOk returns a tuple with the Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationCoreSpec) GetCoreOk() (*string, bool) {
	if o == nil || IsNil(o.Core) {
		return nil, false
	}
	return o.Core, true
}

// HasCore returns a boolean if a field has been set.
func (o *V0040ReservationCoreSpec) HasCore() bool {
	if o != nil && !IsNil(o.Core) {
		return true
	}

	return false
}

// SetCore gets a reference to the given string and assigns it to the Core field.
func (o *V0040ReservationCoreSpec) SetCore(v string) {
	o.Core = &v
}

func (o V0040ReservationCoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ReservationCoreSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Core) {
		toSerialize["core"] = o.Core
	}
	return toSerialize, nil
}

type NullableV0040ReservationCoreSpec struct {
	value *V0040ReservationCoreSpec
	isSet bool
}

func (v NullableV0040ReservationCoreSpec) Get() *V0040ReservationCoreSpec {
	return v.value
}

func (v *NullableV0040ReservationCoreSpec) Set(val *V0040ReservationCoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ReservationCoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ReservationCoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ReservationCoreSpec(val *V0040ReservationCoreSpec) *NullableV0040ReservationCoreSpec {
	return &NullableV0040ReservationCoreSpec{value: val, isSet: true}
}

func (v NullableV0040ReservationCoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ReservationCoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


