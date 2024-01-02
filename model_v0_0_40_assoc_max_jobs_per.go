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

// checks if the V0040AssocMaxJobsPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMaxJobsPer{}

// V0040AssocMaxJobsPer struct for V0040AssocMaxJobsPer
type V0040AssocMaxJobsPer struct {
	Count *V0040Uint32NoVal `json:"count,omitempty"`
	Accruing *V0040Uint32NoVal `json:"accruing,omitempty"`
	Submitted *V0040Uint32NoVal `json:"submitted,omitempty"`
	WallClock *V0040Uint32NoVal `json:"wall_clock,omitempty"`
}

// NewV0040AssocMaxJobsPer instantiates a new V0040AssocMaxJobsPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMaxJobsPer() *V0040AssocMaxJobsPer {
	this := V0040AssocMaxJobsPer{}
	return &this
}

// NewV0040AssocMaxJobsPerWithDefaults instantiates a new V0040AssocMaxJobsPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxJobsPerWithDefaults() *V0040AssocMaxJobsPer {
	this := V0040AssocMaxJobsPer{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040AssocMaxJobsPer) GetCount() V0040Uint32NoVal {
	if o == nil || IsNil(o.Count) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobsPer) GetCountOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040AssocMaxJobsPer) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given V0040Uint32NoVal and assigns it to the Count field.
func (o *V0040AssocMaxJobsPer) SetCount(v V0040Uint32NoVal) {
	o.Count = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0040AssocMaxJobsPer) GetAccruing() V0040Uint32NoVal {
	if o == nil || IsNil(o.Accruing) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobsPer) GetAccruingOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0040AssocMaxJobsPer) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0040Uint32NoVal and assigns it to the Accruing field.
func (o *V0040AssocMaxJobsPer) SetAccruing(v V0040Uint32NoVal) {
	o.Accruing = &v
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *V0040AssocMaxJobsPer) GetSubmitted() V0040Uint32NoVal {
	if o == nil || IsNil(o.Submitted) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobsPer) GetSubmittedOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *V0040AssocMaxJobsPer) HasSubmitted() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given V0040Uint32NoVal and assigns it to the Submitted field.
func (o *V0040AssocMaxJobsPer) SetSubmitted(v V0040Uint32NoVal) {
	o.Submitted = &v
}

// GetWallClock returns the WallClock field value if set, zero value otherwise.
func (o *V0040AssocMaxJobsPer) GetWallClock() V0040Uint32NoVal {
	if o == nil || IsNil(o.WallClock) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.WallClock
}

// GetWallClockOk returns a tuple with the WallClock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobsPer) GetWallClockOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.WallClock) {
		return nil, false
	}
	return o.WallClock, true
}

// HasWallClock returns a boolean if a field has been set.
func (o *V0040AssocMaxJobsPer) HasWallClock() bool {
	if o != nil && !IsNil(o.WallClock) {
		return true
	}

	return false
}

// SetWallClock gets a reference to the given V0040Uint32NoVal and assigns it to the WallClock field.
func (o *V0040AssocMaxJobsPer) SetWallClock(v V0040Uint32NoVal) {
	o.WallClock = &v
}

func (o V0040AssocMaxJobsPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMaxJobsPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Submitted) {
		toSerialize["submitted"] = o.Submitted
	}
	if !IsNil(o.WallClock) {
		toSerialize["wall_clock"] = o.WallClock
	}
	return toSerialize, nil
}

type NullableV0040AssocMaxJobsPer struct {
	value *V0040AssocMaxJobsPer
	isSet bool
}

func (v NullableV0040AssocMaxJobsPer) Get() *V0040AssocMaxJobsPer {
	return v.value
}

func (v *NullableV0040AssocMaxJobsPer) Set(val *V0040AssocMaxJobsPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMaxJobsPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMaxJobsPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMaxJobsPer(val *V0040AssocMaxJobsPer) *NullableV0040AssocMaxJobsPer {
	return &NullableV0040AssocMaxJobsPer{value: val, isSet: true}
}

func (v NullableV0040AssocMaxJobsPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMaxJobsPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


