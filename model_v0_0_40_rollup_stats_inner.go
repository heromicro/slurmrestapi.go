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

// checks if the V0040RollupStatsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040RollupStatsInner{}

// V0040RollupStatsInner recorded rollup statistics
type V0040RollupStatsInner struct {
	// type
	Type *string `json:"type,omitempty"`
	// Last time rollup ran (UNIX timestamp)
	LastRun *int32 `json:"last run,omitempty"`
	// longest rollup time (seconds)
	MaxCycle *int64 `json:"max_cycle,omitempty"`
	// total time spent doing rollups (seconds)
	TotalTime *int64 `json:"total_time,omitempty"`
	// number of rollups since last_run
	TotalCycles *int64 `json:"total_cycles,omitempty"`
	// average time for rollup (seconds)
	MeanCycles *int64 `json:"mean_cycles,omitempty"`
}

// NewV0040RollupStatsInner instantiates a new V0040RollupStatsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040RollupStatsInner() *V0040RollupStatsInner {
	this := V0040RollupStatsInner{}
	return &this
}

// NewV0040RollupStatsInnerWithDefaults instantiates a new V0040RollupStatsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040RollupStatsInnerWithDefaults() *V0040RollupStatsInner {
	this := V0040RollupStatsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V0040RollupStatsInner) SetType(v string) {
	o.Type = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetLastRun() int32 {
	if o == nil || IsNil(o.LastRun) {
		var ret int32
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetLastRunOk() (*int32, bool) {
	if o == nil || IsNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasLastRun() bool {
	if o != nil && !IsNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given int32 and assigns it to the LastRun field.
func (o *V0040RollupStatsInner) SetLastRun(v int32) {
	o.LastRun = &v
}

// GetMaxCycle returns the MaxCycle field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetMaxCycle() int64 {
	if o == nil || IsNil(o.MaxCycle) {
		var ret int64
		return ret
	}
	return *o.MaxCycle
}

// GetMaxCycleOk returns a tuple with the MaxCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetMaxCycleOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxCycle) {
		return nil, false
	}
	return o.MaxCycle, true
}

// HasMaxCycle returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasMaxCycle() bool {
	if o != nil && !IsNil(o.MaxCycle) {
		return true
	}

	return false
}

// SetMaxCycle gets a reference to the given int64 and assigns it to the MaxCycle field.
func (o *V0040RollupStatsInner) SetMaxCycle(v int64) {
	o.MaxCycle = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetTotalTime() int64 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int64
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetTotalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int64 and assigns it to the TotalTime field.
func (o *V0040RollupStatsInner) SetTotalTime(v int64) {
	o.TotalTime = &v
}

// GetTotalCycles returns the TotalCycles field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetTotalCycles() int64 {
	if o == nil || IsNil(o.TotalCycles) {
		var ret int64
		return ret
	}
	return *o.TotalCycles
}

// GetTotalCyclesOk returns a tuple with the TotalCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetTotalCyclesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCycles) {
		return nil, false
	}
	return o.TotalCycles, true
}

// HasTotalCycles returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasTotalCycles() bool {
	if o != nil && !IsNil(o.TotalCycles) {
		return true
	}

	return false
}

// SetTotalCycles gets a reference to the given int64 and assigns it to the TotalCycles field.
func (o *V0040RollupStatsInner) SetTotalCycles(v int64) {
	o.TotalCycles = &v
}

// GetMeanCycles returns the MeanCycles field value if set, zero value otherwise.
func (o *V0040RollupStatsInner) GetMeanCycles() int64 {
	if o == nil || IsNil(o.MeanCycles) {
		var ret int64
		return ret
	}
	return *o.MeanCycles
}

// GetMeanCyclesOk returns a tuple with the MeanCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040RollupStatsInner) GetMeanCyclesOk() (*int64, bool) {
	if o == nil || IsNil(o.MeanCycles) {
		return nil, false
	}
	return o.MeanCycles, true
}

// HasMeanCycles returns a boolean if a field has been set.
func (o *V0040RollupStatsInner) HasMeanCycles() bool {
	if o != nil && !IsNil(o.MeanCycles) {
		return true
	}

	return false
}

// SetMeanCycles gets a reference to the given int64 and assigns it to the MeanCycles field.
func (o *V0040RollupStatsInner) SetMeanCycles(v int64) {
	o.MeanCycles = &v
}

func (o V0040RollupStatsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040RollupStatsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LastRun) {
		toSerialize["last run"] = o.LastRun
	}
	if !IsNil(o.MaxCycle) {
		toSerialize["max_cycle"] = o.MaxCycle
	}
	if !IsNil(o.TotalTime) {
		toSerialize["total_time"] = o.TotalTime
	}
	if !IsNil(o.TotalCycles) {
		toSerialize["total_cycles"] = o.TotalCycles
	}
	if !IsNil(o.MeanCycles) {
		toSerialize["mean_cycles"] = o.MeanCycles
	}
	return toSerialize, nil
}

type NullableV0040RollupStatsInner struct {
	value *V0040RollupStatsInner
	isSet bool
}

func (v NullableV0040RollupStatsInner) Get() *V0040RollupStatsInner {
	return v.value
}

func (v *NullableV0040RollupStatsInner) Set(val *V0040RollupStatsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040RollupStatsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040RollupStatsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040RollupStatsInner(val *V0040RollupStatsInner) *NullableV0040RollupStatsInner {
	return &NullableV0040RollupStatsInner{value: val, isSet: true}
}

func (v NullableV0040RollupStatsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040RollupStatsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


