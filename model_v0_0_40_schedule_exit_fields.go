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

// checks if the V0040ScheduleExitFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ScheduleExitFields{}

// V0040ScheduleExitFields struct for V0040ScheduleExitFields
type V0040ScheduleExitFields struct {
	EndJobQueue *int32 `json:"end_job_queue,omitempty"`
	DefaultQueueDepth *int32 `json:"default_queue_depth,omitempty"`
	MaxJobStart *int32 `json:"max_job_start,omitempty"`
	MaxRpcCnt *int32 `json:"max_rpc_cnt,omitempty"`
	MaxSchedTime *int32 `json:"max_sched_time,omitempty"`
	Licenses *int32 `json:"licenses,omitempty"`
}

// NewV0040ScheduleExitFields instantiates a new V0040ScheduleExitFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ScheduleExitFields() *V0040ScheduleExitFields {
	this := V0040ScheduleExitFields{}
	return &this
}

// NewV0040ScheduleExitFieldsWithDefaults instantiates a new V0040ScheduleExitFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ScheduleExitFieldsWithDefaults() *V0040ScheduleExitFields {
	this := V0040ScheduleExitFields{}
	return &this
}

// GetEndJobQueue returns the EndJobQueue field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetEndJobQueue() int32 {
	if o == nil || IsNil(o.EndJobQueue) {
		var ret int32
		return ret
	}
	return *o.EndJobQueue
}

// GetEndJobQueueOk returns a tuple with the EndJobQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetEndJobQueueOk() (*int32, bool) {
	if o == nil || IsNil(o.EndJobQueue) {
		return nil, false
	}
	return o.EndJobQueue, true
}

// HasEndJobQueue returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasEndJobQueue() bool {
	if o != nil && !IsNil(o.EndJobQueue) {
		return true
	}

	return false
}

// SetEndJobQueue gets a reference to the given int32 and assigns it to the EndJobQueue field.
func (o *V0040ScheduleExitFields) SetEndJobQueue(v int32) {
	o.EndJobQueue = &v
}

// GetDefaultQueueDepth returns the DefaultQueueDepth field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetDefaultQueueDepth() int32 {
	if o == nil || IsNil(o.DefaultQueueDepth) {
		var ret int32
		return ret
	}
	return *o.DefaultQueueDepth
}

// GetDefaultQueueDepthOk returns a tuple with the DefaultQueueDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetDefaultQueueDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultQueueDepth) {
		return nil, false
	}
	return o.DefaultQueueDepth, true
}

// HasDefaultQueueDepth returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasDefaultQueueDepth() bool {
	if o != nil && !IsNil(o.DefaultQueueDepth) {
		return true
	}

	return false
}

// SetDefaultQueueDepth gets a reference to the given int32 and assigns it to the DefaultQueueDepth field.
func (o *V0040ScheduleExitFields) SetDefaultQueueDepth(v int32) {
	o.DefaultQueueDepth = &v
}

// GetMaxJobStart returns the MaxJobStart field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetMaxJobStart() int32 {
	if o == nil || IsNil(o.MaxJobStart) {
		var ret int32
		return ret
	}
	return *o.MaxJobStart
}

// GetMaxJobStartOk returns a tuple with the MaxJobStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetMaxJobStartOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxJobStart) {
		return nil, false
	}
	return o.MaxJobStart, true
}

// HasMaxJobStart returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasMaxJobStart() bool {
	if o != nil && !IsNil(o.MaxJobStart) {
		return true
	}

	return false
}

// SetMaxJobStart gets a reference to the given int32 and assigns it to the MaxJobStart field.
func (o *V0040ScheduleExitFields) SetMaxJobStart(v int32) {
	o.MaxJobStart = &v
}

// GetMaxRpcCnt returns the MaxRpcCnt field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetMaxRpcCnt() int32 {
	if o == nil || IsNil(o.MaxRpcCnt) {
		var ret int32
		return ret
	}
	return *o.MaxRpcCnt
}

// GetMaxRpcCntOk returns a tuple with the MaxRpcCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetMaxRpcCntOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRpcCnt) {
		return nil, false
	}
	return o.MaxRpcCnt, true
}

// HasMaxRpcCnt returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasMaxRpcCnt() bool {
	if o != nil && !IsNil(o.MaxRpcCnt) {
		return true
	}

	return false
}

// SetMaxRpcCnt gets a reference to the given int32 and assigns it to the MaxRpcCnt field.
func (o *V0040ScheduleExitFields) SetMaxRpcCnt(v int32) {
	o.MaxRpcCnt = &v
}

// GetMaxSchedTime returns the MaxSchedTime field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetMaxSchedTime() int32 {
	if o == nil || IsNil(o.MaxSchedTime) {
		var ret int32
		return ret
	}
	return *o.MaxSchedTime
}

// GetMaxSchedTimeOk returns a tuple with the MaxSchedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetMaxSchedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSchedTime) {
		return nil, false
	}
	return o.MaxSchedTime, true
}

// HasMaxSchedTime returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasMaxSchedTime() bool {
	if o != nil && !IsNil(o.MaxSchedTime) {
		return true
	}

	return false
}

// SetMaxSchedTime gets a reference to the given int32 and assigns it to the MaxSchedTime field.
func (o *V0040ScheduleExitFields) SetMaxSchedTime(v int32) {
	o.MaxSchedTime = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0040ScheduleExitFields) GetLicenses() int32 {
	if o == nil || IsNil(o.Licenses) {
		var ret int32
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ScheduleExitFields) GetLicensesOk() (*int32, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0040ScheduleExitFields) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given int32 and assigns it to the Licenses field.
func (o *V0040ScheduleExitFields) SetLicenses(v int32) {
	o.Licenses = &v
}

func (o V0040ScheduleExitFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ScheduleExitFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndJobQueue) {
		toSerialize["end_job_queue"] = o.EndJobQueue
	}
	if !IsNil(o.DefaultQueueDepth) {
		toSerialize["default_queue_depth"] = o.DefaultQueueDepth
	}
	if !IsNil(o.MaxJobStart) {
		toSerialize["max_job_start"] = o.MaxJobStart
	}
	if !IsNil(o.MaxRpcCnt) {
		toSerialize["max_rpc_cnt"] = o.MaxRpcCnt
	}
	if !IsNil(o.MaxSchedTime) {
		toSerialize["max_sched_time"] = o.MaxSchedTime
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableV0040ScheduleExitFields struct {
	value *V0040ScheduleExitFields
	isSet bool
}

func (v NullableV0040ScheduleExitFields) Get() *V0040ScheduleExitFields {
	return v.value
}

func (v *NullableV0040ScheduleExitFields) Set(val *V0040ScheduleExitFields) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ScheduleExitFields) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ScheduleExitFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ScheduleExitFields(val *V0040ScheduleExitFields) *NullableV0040ScheduleExitFields {
	return &NullableV0040ScheduleExitFields{value: val, isSet: true}
}

func (v NullableV0040ScheduleExitFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ScheduleExitFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

