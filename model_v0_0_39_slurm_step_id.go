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

// checks if the V0039SlurmStepId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039SlurmStepId{}

// V0039SlurmStepId step details
type V0039SlurmStepId struct {
	// JobID
	JobId *int32 `json:"job_id,omitempty"`
	// HetStep
	StepHetComponent *int32 `json:"step_het_component,omitempty"`
	StepId *string `json:"step_id,omitempty"`
}

// NewV0039SlurmStepId instantiates a new V0039SlurmStepId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039SlurmStepId() *V0039SlurmStepId {
	this := V0039SlurmStepId{}
	return &this
}

// NewV0039SlurmStepIdWithDefaults instantiates a new V0039SlurmStepId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039SlurmStepIdWithDefaults() *V0039SlurmStepId {
	this := V0039SlurmStepId{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0039SlurmStepId) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039SlurmStepId) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0039SlurmStepId) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0039SlurmStepId) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepHetComponent returns the StepHetComponent field value if set, zero value otherwise.
func (o *V0039SlurmStepId) GetStepHetComponent() int32 {
	if o == nil || IsNil(o.StepHetComponent) {
		var ret int32
		return ret
	}
	return *o.StepHetComponent
}

// GetStepHetComponentOk returns a tuple with the StepHetComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039SlurmStepId) GetStepHetComponentOk() (*int32, bool) {
	if o == nil || IsNil(o.StepHetComponent) {
		return nil, false
	}
	return o.StepHetComponent, true
}

// HasStepHetComponent returns a boolean if a field has been set.
func (o *V0039SlurmStepId) HasStepHetComponent() bool {
	if o != nil && !IsNil(o.StepHetComponent) {
		return true
	}

	return false
}

// SetStepHetComponent gets a reference to the given int32 and assigns it to the StepHetComponent field.
func (o *V0039SlurmStepId) SetStepHetComponent(v int32) {
	o.StepHetComponent = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0039SlurmStepId) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039SlurmStepId) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0039SlurmStepId) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0039SlurmStepId) SetStepId(v string) {
	o.StepId = &v
}

func (o V0039SlurmStepId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039SlurmStepId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepHetComponent) {
		toSerialize["step_het_component"] = o.StepHetComponent
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	return toSerialize, nil
}

type NullableV0039SlurmStepId struct {
	value *V0039SlurmStepId
	isSet bool
}

func (v NullableV0039SlurmStepId) Get() *V0039SlurmStepId {
	return v.value
}

func (v *NullableV0039SlurmStepId) Set(val *V0039SlurmStepId) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039SlurmStepId) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039SlurmStepId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039SlurmStepId(val *V0039SlurmStepId) *NullableV0039SlurmStepId {
	return &NullableV0039SlurmStepId{value: val, isSet: true}
}

func (v NullableV0039SlurmStepId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039SlurmStepId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

