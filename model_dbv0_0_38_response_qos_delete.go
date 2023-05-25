/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0038ResponseQosDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ResponseQosDelete{}

// Dbv0038ResponseQosDelete struct for Dbv0038ResponseQosDelete
type Dbv0038ResponseQosDelete struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
}

// NewDbv0038ResponseQosDelete instantiates a new Dbv0038ResponseQosDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ResponseQosDelete() *Dbv0038ResponseQosDelete {
	this := Dbv0038ResponseQosDelete{}
	return &this
}

// NewDbv0038ResponseQosDeleteWithDefaults instantiates a new Dbv0038ResponseQosDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ResponseQosDeleteWithDefaults() *Dbv0038ResponseQosDelete {
	this := Dbv0038ResponseQosDelete{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038ResponseQosDelete) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseQosDelete) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038ResponseQosDelete) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038ResponseQosDelete) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038ResponseQosDelete) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseQosDelete) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038ResponseQosDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038ResponseQosDelete) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

func (o Dbv0038ResponseQosDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ResponseQosDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0038ResponseQosDelete struct {
	value *Dbv0038ResponseQosDelete
	isSet bool
}

func (v NullableDbv0038ResponseQosDelete) Get() *Dbv0038ResponseQosDelete {
	return v.value
}

func (v *NullableDbv0038ResponseQosDelete) Set(val *Dbv0038ResponseQosDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ResponseQosDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ResponseQosDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ResponseQosDelete(val *Dbv0038ResponseQosDelete) *NullableDbv0038ResponseQosDelete {
	return &NullableDbv0038ResponseQosDelete{value: val, isSet: true}
}

func (v NullableDbv0038ResponseQosDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ResponseQosDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


