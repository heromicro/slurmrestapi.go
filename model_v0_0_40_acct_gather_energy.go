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

// checks if the V0040AcctGatherEnergy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AcctGatherEnergy{}

// V0040AcctGatherEnergy struct for V0040AcctGatherEnergy
type V0040AcctGatherEnergy struct {
	AverageWatts *int32 `json:"average_watts,omitempty"`
	BaseConsumedEnergy *int64 `json:"base_consumed_energy,omitempty"`
	ConsumedEnergy *int64 `json:"consumed_energy,omitempty"`
	CurrentWatts *V0040Uint32NoVal `json:"current_watts,omitempty"`
	PreviousConsumedEnergy *int64 `json:"previous_consumed_energy,omitempty"`
	LastCollected *int64 `json:"last_collected,omitempty"`
}

// NewV0040AcctGatherEnergy instantiates a new V0040AcctGatherEnergy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AcctGatherEnergy() *V0040AcctGatherEnergy {
	this := V0040AcctGatherEnergy{}
	return &this
}

// NewV0040AcctGatherEnergyWithDefaults instantiates a new V0040AcctGatherEnergy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AcctGatherEnergyWithDefaults() *V0040AcctGatherEnergy {
	this := V0040AcctGatherEnergy{}
	return &this
}

// GetAverageWatts returns the AverageWatts field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetAverageWatts() int32 {
	if o == nil || IsNil(o.AverageWatts) {
		var ret int32
		return ret
	}
	return *o.AverageWatts
}

// GetAverageWattsOk returns a tuple with the AverageWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetAverageWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageWatts) {
		return nil, false
	}
	return o.AverageWatts, true
}

// HasAverageWatts returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasAverageWatts() bool {
	if o != nil && !IsNil(o.AverageWatts) {
		return true
	}

	return false
}

// SetAverageWatts gets a reference to the given int32 and assigns it to the AverageWatts field.
func (o *V0040AcctGatherEnergy) SetAverageWatts(v int32) {
	o.AverageWatts = &v
}

// GetBaseConsumedEnergy returns the BaseConsumedEnergy field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetBaseConsumedEnergy() int64 {
	if o == nil || IsNil(o.BaseConsumedEnergy) {
		var ret int64
		return ret
	}
	return *o.BaseConsumedEnergy
}

// GetBaseConsumedEnergyOk returns a tuple with the BaseConsumedEnergy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetBaseConsumedEnergyOk() (*int64, bool) {
	if o == nil || IsNil(o.BaseConsumedEnergy) {
		return nil, false
	}
	return o.BaseConsumedEnergy, true
}

// HasBaseConsumedEnergy returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasBaseConsumedEnergy() bool {
	if o != nil && !IsNil(o.BaseConsumedEnergy) {
		return true
	}

	return false
}

// SetBaseConsumedEnergy gets a reference to the given int64 and assigns it to the BaseConsumedEnergy field.
func (o *V0040AcctGatherEnergy) SetBaseConsumedEnergy(v int64) {
	o.BaseConsumedEnergy = &v
}

// GetConsumedEnergy returns the ConsumedEnergy field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetConsumedEnergy() int64 {
	if o == nil || IsNil(o.ConsumedEnergy) {
		var ret int64
		return ret
	}
	return *o.ConsumedEnergy
}

// GetConsumedEnergyOk returns a tuple with the ConsumedEnergy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetConsumedEnergyOk() (*int64, bool) {
	if o == nil || IsNil(o.ConsumedEnergy) {
		return nil, false
	}
	return o.ConsumedEnergy, true
}

// HasConsumedEnergy returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasConsumedEnergy() bool {
	if o != nil && !IsNil(o.ConsumedEnergy) {
		return true
	}

	return false
}

// SetConsumedEnergy gets a reference to the given int64 and assigns it to the ConsumedEnergy field.
func (o *V0040AcctGatherEnergy) SetConsumedEnergy(v int64) {
	o.ConsumedEnergy = &v
}

// GetCurrentWatts returns the CurrentWatts field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetCurrentWatts() V0040Uint32NoVal {
	if o == nil || IsNil(o.CurrentWatts) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.CurrentWatts
}

// GetCurrentWattsOk returns a tuple with the CurrentWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetCurrentWattsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.CurrentWatts) {
		return nil, false
	}
	return o.CurrentWatts, true
}

// HasCurrentWatts returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasCurrentWatts() bool {
	if o != nil && !IsNil(o.CurrentWatts) {
		return true
	}

	return false
}

// SetCurrentWatts gets a reference to the given V0040Uint32NoVal and assigns it to the CurrentWatts field.
func (o *V0040AcctGatherEnergy) SetCurrentWatts(v V0040Uint32NoVal) {
	o.CurrentWatts = &v
}

// GetPreviousConsumedEnergy returns the PreviousConsumedEnergy field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetPreviousConsumedEnergy() int64 {
	if o == nil || IsNil(o.PreviousConsumedEnergy) {
		var ret int64
		return ret
	}
	return *o.PreviousConsumedEnergy
}

// GetPreviousConsumedEnergyOk returns a tuple with the PreviousConsumedEnergy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetPreviousConsumedEnergyOk() (*int64, bool) {
	if o == nil || IsNil(o.PreviousConsumedEnergy) {
		return nil, false
	}
	return o.PreviousConsumedEnergy, true
}

// HasPreviousConsumedEnergy returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasPreviousConsumedEnergy() bool {
	if o != nil && !IsNil(o.PreviousConsumedEnergy) {
		return true
	}

	return false
}

// SetPreviousConsumedEnergy gets a reference to the given int64 and assigns it to the PreviousConsumedEnergy field.
func (o *V0040AcctGatherEnergy) SetPreviousConsumedEnergy(v int64) {
	o.PreviousConsumedEnergy = &v
}

// GetLastCollected returns the LastCollected field value if set, zero value otherwise.
func (o *V0040AcctGatherEnergy) GetLastCollected() int64 {
	if o == nil || IsNil(o.LastCollected) {
		var ret int64
		return ret
	}
	return *o.LastCollected
}

// GetLastCollectedOk returns a tuple with the LastCollected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AcctGatherEnergy) GetLastCollectedOk() (*int64, bool) {
	if o == nil || IsNil(o.LastCollected) {
		return nil, false
	}
	return o.LastCollected, true
}

// HasLastCollected returns a boolean if a field has been set.
func (o *V0040AcctGatherEnergy) HasLastCollected() bool {
	if o != nil && !IsNil(o.LastCollected) {
		return true
	}

	return false
}

// SetLastCollected gets a reference to the given int64 and assigns it to the LastCollected field.
func (o *V0040AcctGatherEnergy) SetLastCollected(v int64) {
	o.LastCollected = &v
}

func (o V0040AcctGatherEnergy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AcctGatherEnergy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageWatts) {
		toSerialize["average_watts"] = o.AverageWatts
	}
	if !IsNil(o.BaseConsumedEnergy) {
		toSerialize["base_consumed_energy"] = o.BaseConsumedEnergy
	}
	if !IsNil(o.ConsumedEnergy) {
		toSerialize["consumed_energy"] = o.ConsumedEnergy
	}
	if !IsNil(o.CurrentWatts) {
		toSerialize["current_watts"] = o.CurrentWatts
	}
	if !IsNil(o.PreviousConsumedEnergy) {
		toSerialize["previous_consumed_energy"] = o.PreviousConsumedEnergy
	}
	if !IsNil(o.LastCollected) {
		toSerialize["last_collected"] = o.LastCollected
	}
	return toSerialize, nil
}

type NullableV0040AcctGatherEnergy struct {
	value *V0040AcctGatherEnergy
	isSet bool
}

func (v NullableV0040AcctGatherEnergy) Get() *V0040AcctGatherEnergy {
	return v.value
}

func (v *NullableV0040AcctGatherEnergy) Set(val *V0040AcctGatherEnergy) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AcctGatherEnergy) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AcctGatherEnergy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AcctGatherEnergy(val *V0040AcctGatherEnergy) *NullableV0040AcctGatherEnergy {
	return &NullableV0040AcctGatherEnergy{value: val, isSet: true}
}

func (v NullableV0040AcctGatherEnergy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AcctGatherEnergy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


