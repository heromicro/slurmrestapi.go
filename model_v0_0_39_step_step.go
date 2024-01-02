/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039StepStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepStep{}

// V0039StepStep struct for V0039StepStep
type V0039StepStep struct {
	Id *V0039SlurmStepId `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV0039StepStep instantiates a new V0039StepStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepStep() *V0039StepStep {
	this := V0039StepStep{}
	return &this
}

// NewV0039StepStepWithDefaults instantiates a new V0039StepStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepStepWithDefaults() *V0039StepStep {
	this := V0039StepStep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0039StepStep) GetId() V0039SlurmStepId {
	if o == nil || IsNil(o.Id) {
		var ret V0039SlurmStepId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepStep) GetIdOk() (*V0039SlurmStepId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0039StepStep) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given V0039SlurmStepId and assigns it to the Id field.
func (o *V0039StepStep) SetId(v V0039SlurmStepId) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039StepStep) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepStep) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039StepStep) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039StepStep) SetName(v string) {
	o.Name = &v
}

func (o V0039StepStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0039StepStep struct {
	value *V0039StepStep
	isSet bool
}

func (v NullableV0039StepStep) Get() *V0039StepStep {
	return v.value
}

func (v *NullableV0039StepStep) Set(val *V0039StepStep) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepStep) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepStep(val *V0039StepStep) *NullableV0039StepStep {
	return &NullableV0039StepStep{value: val, isSet: true}
}

func (v NullableV0039StepStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


