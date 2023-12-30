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

// checks if the Dbv0038QosLimitsMaxWallClockPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxWallClockPer{}

// Dbv0038QosLimitsMaxWallClockPer Limit on wallclock per settings
type Dbv0038QosLimitsMaxWallClockPer struct {
	// Max wallclock per QOS
	Qos *int32 `json:"qos,omitempty"`
	// Max wallclock per job
	Job *int32 `json:"job,omitempty"`
}

// NewDbv0038QosLimitsMaxWallClockPer instantiates a new Dbv0038QosLimitsMaxWallClockPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxWallClockPer() *Dbv0038QosLimitsMaxWallClockPer {
	this := Dbv0038QosLimitsMaxWallClockPer{}
	return &this
}

// NewDbv0038QosLimitsMaxWallClockPerWithDefaults instantiates a new Dbv0038QosLimitsMaxWallClockPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxWallClockPerWithDefaults() *Dbv0038QosLimitsMaxWallClockPer {
	this := Dbv0038QosLimitsMaxWallClockPer{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxWallClockPer) GetQos() int32 {
	if o == nil || IsNil(o.Qos) {
		var ret int32
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxWallClockPer) GetQosOk() (*int32, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxWallClockPer) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given int32 and assigns it to the Qos field.
func (o *Dbv0038QosLimitsMaxWallClockPer) SetQos(v int32) {
	o.Qos = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxWallClockPer) GetJob() int32 {
	if o == nil || IsNil(o.Job) {
		var ret int32
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxWallClockPer) GetJobOk() (*int32, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxWallClockPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given int32 and assigns it to the Job field.
func (o *Dbv0038QosLimitsMaxWallClockPer) SetJob(v int32) {
	o.Job = &v
}

func (o Dbv0038QosLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxWallClockPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxWallClockPer struct {
	value *Dbv0038QosLimitsMaxWallClockPer
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxWallClockPer) Get() *Dbv0038QosLimitsMaxWallClockPer {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxWallClockPer) Set(val *Dbv0038QosLimitsMaxWallClockPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxWallClockPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxWallClockPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxWallClockPer(val *Dbv0038QosLimitsMaxWallClockPer) *NullableDbv0038QosLimitsMaxWallClockPer {
	return &NullableDbv0038QosLimitsMaxWallClockPer{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxWallClockPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


