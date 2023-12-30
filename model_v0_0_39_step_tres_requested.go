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

// checks if the V0039StepTresRequested type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepTresRequested{}

// V0039StepTresRequested struct for V0039StepTresRequested
type V0039StepTresRequested struct {
	Max []V0039Tres `json:"max,omitempty"`
	Min []V0039Tres `json:"min,omitempty"`
	Average []V0039Tres `json:"average,omitempty"`
	Total []V0039Tres `json:"total,omitempty"`
}

// NewV0039StepTresRequested instantiates a new V0039StepTresRequested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepTresRequested() *V0039StepTresRequested {
	this := V0039StepTresRequested{}
	return &this
}

// NewV0039StepTresRequestedWithDefaults instantiates a new V0039StepTresRequested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepTresRequestedWithDefaults() *V0039StepTresRequested {
	this := V0039StepTresRequested{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V0039StepTresRequested) GetMax() []V0039Tres {
	if o == nil || IsNil(o.Max) {
		var ret []V0039Tres
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTresRequested) GetMaxOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V0039StepTresRequested) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given []V0039Tres and assigns it to the Max field.
func (o *V0039StepTresRequested) SetMax(v []V0039Tres) {
	o.Max = v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0039StepTresRequested) GetMin() []V0039Tres {
	if o == nil || IsNil(o.Min) {
		var ret []V0039Tres
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTresRequested) GetMinOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0039StepTresRequested) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given []V0039Tres and assigns it to the Min field.
func (o *V0039StepTresRequested) SetMin(v []V0039Tres) {
	o.Min = v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *V0039StepTresRequested) GetAverage() []V0039Tres {
	if o == nil || IsNil(o.Average) {
		var ret []V0039Tres
		return ret
	}
	return o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTresRequested) GetAverageOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *V0039StepTresRequested) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given []V0039Tres and assigns it to the Average field.
func (o *V0039StepTresRequested) SetAverage(v []V0039Tres) {
	o.Average = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0039StepTresRequested) GetTotal() []V0039Tres {
	if o == nil || IsNil(o.Total) {
		var ret []V0039Tres
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTresRequested) GetTotalOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0039StepTresRequested) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []V0039Tres and assigns it to the Total field.
func (o *V0039StepTresRequested) SetTotal(v []V0039Tres) {
	o.Total = v
}

func (o V0039StepTresRequested) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepTresRequested) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0039StepTresRequested struct {
	value *V0039StepTresRequested
	isSet bool
}

func (v NullableV0039StepTresRequested) Get() *V0039StepTresRequested {
	return v.value
}

func (v *NullableV0039StepTresRequested) Set(val *V0039StepTresRequested) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepTresRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepTresRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepTresRequested(val *V0039StepTresRequested) *NullableV0039StepTresRequested {
	return &NullableV0039StepTresRequested{value: val, isSet: true}
}

func (v NullableV0039StepTresRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepTresRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


