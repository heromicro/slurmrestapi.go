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

// checks if the Dbv0037ResponseClusterDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037ResponseClusterDelete{}

// Dbv0037ResponseClusterDelete struct for Dbv0037ResponseClusterDelete
type Dbv0037ResponseClusterDelete struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
}

// NewDbv0037ResponseClusterDelete instantiates a new Dbv0037ResponseClusterDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ResponseClusterDelete() *Dbv0037ResponseClusterDelete {
	this := Dbv0037ResponseClusterDelete{}
	return &this
}

// NewDbv0037ResponseClusterDeleteWithDefaults instantiates a new Dbv0037ResponseClusterDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ResponseClusterDeleteWithDefaults() *Dbv0037ResponseClusterDelete {
	this := Dbv0037ResponseClusterDelete{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037ResponseClusterDelete) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ResponseClusterDelete) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037ResponseClusterDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037ResponseClusterDelete) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

func (o Dbv0037ResponseClusterDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037ResponseClusterDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0037ResponseClusterDelete struct {
	value *Dbv0037ResponseClusterDelete
	isSet bool
}

func (v NullableDbv0037ResponseClusterDelete) Get() *Dbv0037ResponseClusterDelete {
	return v.value
}

func (v *NullableDbv0037ResponseClusterDelete) Set(val *Dbv0037ResponseClusterDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ResponseClusterDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ResponseClusterDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ResponseClusterDelete(val *Dbv0037ResponseClusterDelete) *NullableDbv0037ResponseClusterDelete {
	return &NullableDbv0037ResponseClusterDelete{value: val, isSet: true}
}

func (v NullableDbv0037ResponseClusterDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ResponseClusterDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

