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

// checks if the V0039ExtSensorsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039ExtSensorsData{}

// V0039ExtSensorsData struct for V0039ExtSensorsData
type V0039ExtSensorsData struct {
	ConsumedEnergy *V0039Uint64NoVal `json:"consumed_energy,omitempty"`
	Temperature *V0039Uint32NoVal `json:"temperature,omitempty"`
	EnergyUpdateTime *int64 `json:"energy_update_time,omitempty"`
	CurrentWatts *int32 `json:"current_watts,omitempty"`
}

// NewV0039ExtSensorsData instantiates a new V0039ExtSensorsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039ExtSensorsData() *V0039ExtSensorsData {
	this := V0039ExtSensorsData{}
	return &this
}

// NewV0039ExtSensorsDataWithDefaults instantiates a new V0039ExtSensorsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039ExtSensorsDataWithDefaults() *V0039ExtSensorsData {
	this := V0039ExtSensorsData{}
	return &this
}

// GetConsumedEnergy returns the ConsumedEnergy field value if set, zero value otherwise.
func (o *V0039ExtSensorsData) GetConsumedEnergy() V0039Uint64NoVal {
	if o == nil || IsNil(o.ConsumedEnergy) {
		var ret V0039Uint64NoVal
		return ret
	}
	return *o.ConsumedEnergy
}

// GetConsumedEnergyOk returns a tuple with the ConsumedEnergy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ExtSensorsData) GetConsumedEnergyOk() (*V0039Uint64NoVal, bool) {
	if o == nil || IsNil(o.ConsumedEnergy) {
		return nil, false
	}
	return o.ConsumedEnergy, true
}

// HasConsumedEnergy returns a boolean if a field has been set.
func (o *V0039ExtSensorsData) HasConsumedEnergy() bool {
	if o != nil && !IsNil(o.ConsumedEnergy) {
		return true
	}

	return false
}

// SetConsumedEnergy gets a reference to the given V0039Uint64NoVal and assigns it to the ConsumedEnergy field.
func (o *V0039ExtSensorsData) SetConsumedEnergy(v V0039Uint64NoVal) {
	o.ConsumedEnergy = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *V0039ExtSensorsData) GetTemperature() V0039Uint32NoVal {
	if o == nil || IsNil(o.Temperature) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ExtSensorsData) GetTemperatureOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *V0039ExtSensorsData) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given V0039Uint32NoVal and assigns it to the Temperature field.
func (o *V0039ExtSensorsData) SetTemperature(v V0039Uint32NoVal) {
	o.Temperature = &v
}

// GetEnergyUpdateTime returns the EnergyUpdateTime field value if set, zero value otherwise.
func (o *V0039ExtSensorsData) GetEnergyUpdateTime() int64 {
	if o == nil || IsNil(o.EnergyUpdateTime) {
		var ret int64
		return ret
	}
	return *o.EnergyUpdateTime
}

// GetEnergyUpdateTimeOk returns a tuple with the EnergyUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ExtSensorsData) GetEnergyUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EnergyUpdateTime) {
		return nil, false
	}
	return o.EnergyUpdateTime, true
}

// HasEnergyUpdateTime returns a boolean if a field has been set.
func (o *V0039ExtSensorsData) HasEnergyUpdateTime() bool {
	if o != nil && !IsNil(o.EnergyUpdateTime) {
		return true
	}

	return false
}

// SetEnergyUpdateTime gets a reference to the given int64 and assigns it to the EnergyUpdateTime field.
func (o *V0039ExtSensorsData) SetEnergyUpdateTime(v int64) {
	o.EnergyUpdateTime = &v
}

// GetCurrentWatts returns the CurrentWatts field value if set, zero value otherwise.
func (o *V0039ExtSensorsData) GetCurrentWatts() int32 {
	if o == nil || IsNil(o.CurrentWatts) {
		var ret int32
		return ret
	}
	return *o.CurrentWatts
}

// GetCurrentWattsOk returns a tuple with the CurrentWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ExtSensorsData) GetCurrentWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentWatts) {
		return nil, false
	}
	return o.CurrentWatts, true
}

// HasCurrentWatts returns a boolean if a field has been set.
func (o *V0039ExtSensorsData) HasCurrentWatts() bool {
	if o != nil && !IsNil(o.CurrentWatts) {
		return true
	}

	return false
}

// SetCurrentWatts gets a reference to the given int32 and assigns it to the CurrentWatts field.
func (o *V0039ExtSensorsData) SetCurrentWatts(v int32) {
	o.CurrentWatts = &v
}

func (o V0039ExtSensorsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039ExtSensorsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumedEnergy) {
		toSerialize["consumed_energy"] = o.ConsumedEnergy
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !IsNil(o.EnergyUpdateTime) {
		toSerialize["energy_update_time"] = o.EnergyUpdateTime
	}
	if !IsNil(o.CurrentWatts) {
		toSerialize["current_watts"] = o.CurrentWatts
	}
	return toSerialize, nil
}

type NullableV0039ExtSensorsData struct {
	value *V0039ExtSensorsData
	isSet bool
}

func (v NullableV0039ExtSensorsData) Get() *V0039ExtSensorsData {
	return v.value
}

func (v *NullableV0039ExtSensorsData) Set(val *V0039ExtSensorsData) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039ExtSensorsData) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039ExtSensorsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039ExtSensorsData(val *V0039ExtSensorsData) *NullableV0039ExtSensorsData {
	return &NullableV0039ExtSensorsData{value: val, isSet: true}
}

func (v NullableV0039ExtSensorsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039ExtSensorsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


