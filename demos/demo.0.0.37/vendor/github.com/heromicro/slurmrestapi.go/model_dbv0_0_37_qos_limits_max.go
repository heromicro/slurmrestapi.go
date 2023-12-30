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

// checks if the Dbv0037QosLimitsMax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037QosLimitsMax{}

// Dbv0037QosLimitsMax Limits on max settings
type Dbv0037QosLimitsMax struct {
	WallClock *Dbv0037QosLimitsMaxWallClock `json:"wall_clock,omitempty"`
	Jobs *Dbv0037QosLimitsMaxJobs `json:"jobs,omitempty"`
	Accruing *Dbv0037QosLimitsMaxAccruing `json:"accruing,omitempty"`
	Tres *Dbv0037QosLimitsMaxTres `json:"tres,omitempty"`
}

// NewDbv0037QosLimitsMax instantiates a new Dbv0037QosLimitsMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimitsMax() *Dbv0037QosLimitsMax {
	this := Dbv0037QosLimitsMax{}
	return &this
}

// NewDbv0037QosLimitsMaxWithDefaults instantiates a new Dbv0037QosLimitsMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsMaxWithDefaults() *Dbv0037QosLimitsMax {
	this := Dbv0037QosLimitsMax{}
	return &this
}

// GetWallClock returns the WallClock field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMax) GetWallClock() Dbv0037QosLimitsMaxWallClock {
	if o == nil || IsNil(o.WallClock) {
		var ret Dbv0037QosLimitsMaxWallClock
		return ret
	}
	return *o.WallClock
}

// GetWallClockOk returns a tuple with the WallClock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMax) GetWallClockOk() (*Dbv0037QosLimitsMaxWallClock, bool) {
	if o == nil || IsNil(o.WallClock) {
		return nil, false
	}
	return o.WallClock, true
}

// HasWallClock returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMax) HasWallClock() bool {
	if o != nil && !IsNil(o.WallClock) {
		return true
	}

	return false
}

// SetWallClock gets a reference to the given Dbv0037QosLimitsMaxWallClock and assigns it to the WallClock field.
func (o *Dbv0037QosLimitsMax) SetWallClock(v Dbv0037QosLimitsMaxWallClock) {
	o.WallClock = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMax) GetJobs() Dbv0037QosLimitsMaxJobs {
	if o == nil || IsNil(o.Jobs) {
		var ret Dbv0037QosLimitsMaxJobs
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMax) GetJobsOk() (*Dbv0037QosLimitsMaxJobs, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMax) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given Dbv0037QosLimitsMaxJobs and assigns it to the Jobs field.
func (o *Dbv0037QosLimitsMax) SetJobs(v Dbv0037QosLimitsMaxJobs) {
	o.Jobs = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMax) GetAccruing() Dbv0037QosLimitsMaxAccruing {
	if o == nil || IsNil(o.Accruing) {
		var ret Dbv0037QosLimitsMaxAccruing
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMax) GetAccruingOk() (*Dbv0037QosLimitsMaxAccruing, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMax) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given Dbv0037QosLimitsMaxAccruing and assigns it to the Accruing field.
func (o *Dbv0037QosLimitsMax) SetAccruing(v Dbv0037QosLimitsMaxAccruing) {
	o.Accruing = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMax) GetTres() Dbv0037QosLimitsMaxTres {
	if o == nil || IsNil(o.Tres) {
		var ret Dbv0037QosLimitsMaxTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMax) GetTresOk() (*Dbv0037QosLimitsMaxTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMax) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given Dbv0037QosLimitsMaxTres and assigns it to the Tres field.
func (o *Dbv0037QosLimitsMax) SetTres(v Dbv0037QosLimitsMaxTres) {
	o.Tres = &v
}

func (o Dbv0037QosLimitsMax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037QosLimitsMax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WallClock) {
		toSerialize["wall_clock"] = o.WallClock
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0037QosLimitsMax struct {
	value *Dbv0037QosLimitsMax
	isSet bool
}

func (v NullableDbv0037QosLimitsMax) Get() *Dbv0037QosLimitsMax {
	return v.value
}

func (v *NullableDbv0037QosLimitsMax) Set(val *Dbv0037QosLimitsMax) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimitsMax) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimitsMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimitsMax(val *Dbv0037QosLimitsMax) *NullableDbv0037QosLimitsMax {
	return &NullableDbv0037QosLimitsMax{value: val, isSet: true}
}

func (v NullableDbv0037QosLimitsMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimitsMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


