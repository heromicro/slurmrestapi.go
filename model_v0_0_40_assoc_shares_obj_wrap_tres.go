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

// checks if the V0040AssocSharesObjWrapTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocSharesObjWrapTres{}

// V0040AssocSharesObjWrapTres struct for V0040AssocSharesObjWrapTres
type V0040AssocSharesObjWrapTres struct {
	RunSeconds []V0040SharesUint64Tres `json:"run_seconds,omitempty"`
	GroupMinutes []V0040SharesUint64Tres `json:"group_minutes,omitempty"`
	Usage []V0040SharesFloat128Tres `json:"usage,omitempty"`
}

// NewV0040AssocSharesObjWrapTres instantiates a new V0040AssocSharesObjWrapTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocSharesObjWrapTres() *V0040AssocSharesObjWrapTres {
	this := V0040AssocSharesObjWrapTres{}
	return &this
}

// NewV0040AssocSharesObjWrapTresWithDefaults instantiates a new V0040AssocSharesObjWrapTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocSharesObjWrapTresWithDefaults() *V0040AssocSharesObjWrapTres {
	this := V0040AssocSharesObjWrapTres{}
	return &this
}

// GetRunSeconds returns the RunSeconds field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrapTres) GetRunSeconds() []V0040SharesUint64Tres {
	if o == nil || IsNil(o.RunSeconds) {
		var ret []V0040SharesUint64Tres
		return ret
	}
	return o.RunSeconds
}

// GetRunSecondsOk returns a tuple with the RunSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrapTres) GetRunSecondsOk() ([]V0040SharesUint64Tres, bool) {
	if o == nil || IsNil(o.RunSeconds) {
		return nil, false
	}
	return o.RunSeconds, true
}

// HasRunSeconds returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrapTres) HasRunSeconds() bool {
	if o != nil && !IsNil(o.RunSeconds) {
		return true
	}

	return false
}

// SetRunSeconds gets a reference to the given []V0040SharesUint64Tres and assigns it to the RunSeconds field.
func (o *V0040AssocSharesObjWrapTres) SetRunSeconds(v []V0040SharesUint64Tres) {
	o.RunSeconds = v
}

// GetGroupMinutes returns the GroupMinutes field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrapTres) GetGroupMinutes() []V0040SharesUint64Tres {
	if o == nil || IsNil(o.GroupMinutes) {
		var ret []V0040SharesUint64Tres
		return ret
	}
	return o.GroupMinutes
}

// GetGroupMinutesOk returns a tuple with the GroupMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrapTres) GetGroupMinutesOk() ([]V0040SharesUint64Tres, bool) {
	if o == nil || IsNil(o.GroupMinutes) {
		return nil, false
	}
	return o.GroupMinutes, true
}

// HasGroupMinutes returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrapTres) HasGroupMinutes() bool {
	if o != nil && !IsNil(o.GroupMinutes) {
		return true
	}

	return false
}

// SetGroupMinutes gets a reference to the given []V0040SharesUint64Tres and assigns it to the GroupMinutes field.
func (o *V0040AssocSharesObjWrapTres) SetGroupMinutes(v []V0040SharesUint64Tres) {
	o.GroupMinutes = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrapTres) GetUsage() []V0040SharesFloat128Tres {
	if o == nil || IsNil(o.Usage) {
		var ret []V0040SharesFloat128Tres
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrapTres) GetUsageOk() ([]V0040SharesFloat128Tres, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrapTres) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []V0040SharesFloat128Tres and assigns it to the Usage field.
func (o *V0040AssocSharesObjWrapTres) SetUsage(v []V0040SharesFloat128Tres) {
	o.Usage = v
}

func (o V0040AssocSharesObjWrapTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocSharesObjWrapTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RunSeconds) {
		toSerialize["run_seconds"] = o.RunSeconds
	}
	if !IsNil(o.GroupMinutes) {
		toSerialize["group_minutes"] = o.GroupMinutes
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableV0040AssocSharesObjWrapTres struct {
	value *V0040AssocSharesObjWrapTres
	isSet bool
}

func (v NullableV0040AssocSharesObjWrapTres) Get() *V0040AssocSharesObjWrapTres {
	return v.value
}

func (v *NullableV0040AssocSharesObjWrapTres) Set(val *V0040AssocSharesObjWrapTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocSharesObjWrapTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocSharesObjWrapTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocSharesObjWrapTres(val *V0040AssocSharesObjWrapTres) *NullableV0040AssocSharesObjWrapTres {
	return &NullableV0040AssocSharesObjWrapTres{value: val, isSet: true}
}

func (v NullableV0040AssocSharesObjWrapTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocSharesObjWrapTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


