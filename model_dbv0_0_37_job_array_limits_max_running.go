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

// checks if the Dbv0037JobArrayLimitsMaxRunning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobArrayLimitsMaxRunning{}

// Dbv0037JobArrayLimitsMaxRunning Limits on array settings
type Dbv0037JobArrayLimitsMaxRunning struct {
	// Max running tasks in array at any one time
	Tasks *int32 `json:"tasks,omitempty"`
}

// NewDbv0037JobArrayLimitsMaxRunning instantiates a new Dbv0037JobArrayLimitsMaxRunning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobArrayLimitsMaxRunning() *Dbv0037JobArrayLimitsMaxRunning {
	this := Dbv0037JobArrayLimitsMaxRunning{}
	return &this
}

// NewDbv0037JobArrayLimitsMaxRunningWithDefaults instantiates a new Dbv0037JobArrayLimitsMaxRunning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobArrayLimitsMaxRunningWithDefaults() *Dbv0037JobArrayLimitsMaxRunning {
	this := Dbv0037JobArrayLimitsMaxRunning{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *Dbv0037JobArrayLimitsMaxRunning) GetTasks() int32 {
	if o == nil || IsNil(o.Tasks) {
		var ret int32
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobArrayLimitsMaxRunning) GetTasksOk() (*int32, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *Dbv0037JobArrayLimitsMaxRunning) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given int32 and assigns it to the Tasks field.
func (o *Dbv0037JobArrayLimitsMaxRunning) SetTasks(v int32) {
	o.Tasks = &v
}

func (o Dbv0037JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobArrayLimitsMaxRunning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableDbv0037JobArrayLimitsMaxRunning struct {
	value *Dbv0037JobArrayLimitsMaxRunning
	isSet bool
}

func (v NullableDbv0037JobArrayLimitsMaxRunning) Get() *Dbv0037JobArrayLimitsMaxRunning {
	return v.value
}

func (v *NullableDbv0037JobArrayLimitsMaxRunning) Set(val *Dbv0037JobArrayLimitsMaxRunning) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobArrayLimitsMaxRunning) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobArrayLimitsMaxRunning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobArrayLimitsMaxRunning(val *Dbv0037JobArrayLimitsMaxRunning) *NullableDbv0037JobArrayLimitsMaxRunning {
	return &NullableDbv0037JobArrayLimitsMaxRunning{value: val, isSet: true}
}

func (v NullableDbv0037JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobArrayLimitsMaxRunning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

