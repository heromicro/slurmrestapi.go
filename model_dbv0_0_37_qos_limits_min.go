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

// checks if the Dbv0037QosLimitsMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037QosLimitsMin{}

// Dbv0037QosLimitsMin Min limit settings
type Dbv0037QosLimitsMin struct {
	// Min priority threshold
	PriorityThreshold *int32 `json:"priority_threshold,omitempty"`
	Tres *Dbv0037QosLimitsMinTres `json:"tres,omitempty"`
}

// NewDbv0037QosLimitsMin instantiates a new Dbv0037QosLimitsMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimitsMin() *Dbv0037QosLimitsMin {
	this := Dbv0037QosLimitsMin{}
	return &this
}

// NewDbv0037QosLimitsMinWithDefaults instantiates a new Dbv0037QosLimitsMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsMinWithDefaults() *Dbv0037QosLimitsMin {
	this := Dbv0037QosLimitsMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMin) GetPriorityThreshold() int32 {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret int32
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMin) GetPriorityThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given int32 and assigns it to the PriorityThreshold field.
func (o *Dbv0037QosLimitsMin) SetPriorityThreshold(v int32) {
	o.PriorityThreshold = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMin) GetTres() Dbv0037QosLimitsMinTres {
	if o == nil || IsNil(o.Tres) {
		var ret Dbv0037QosLimitsMinTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMin) GetTresOk() (*Dbv0037QosLimitsMinTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMin) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given Dbv0037QosLimitsMinTres and assigns it to the Tres field.
func (o *Dbv0037QosLimitsMin) SetTres(v Dbv0037QosLimitsMinTres) {
	o.Tres = &v
}

func (o Dbv0037QosLimitsMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037QosLimitsMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0037QosLimitsMin struct {
	value *Dbv0037QosLimitsMin
	isSet bool
}

func (v NullableDbv0037QosLimitsMin) Get() *Dbv0037QosLimitsMin {
	return v.value
}

func (v *NullableDbv0037QosLimitsMin) Set(val *Dbv0037QosLimitsMin) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimitsMin) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimitsMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimitsMin(val *Dbv0037QosLimitsMin) *NullableDbv0037QosLimitsMin {
	return &NullableDbv0037QosLimitsMin{value: val, isSet: true}
}

func (v NullableDbv0037QosLimitsMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimitsMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


