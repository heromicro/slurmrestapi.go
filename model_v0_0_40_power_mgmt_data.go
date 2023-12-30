/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.40
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040PowerMgmtData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PowerMgmtData{}

// V0040PowerMgmtData struct for V0040PowerMgmtData
type V0040PowerMgmtData struct {
	MaximumWatts *V0040Uint32NoVal `json:"maximum_watts,omitempty"`
	CurrentWatts *int32 `json:"current_watts,omitempty"`
	TotalEnergy *int64 `json:"total_energy,omitempty"`
	NewMaximumWatts *int32 `json:"new_maximum_watts,omitempty"`
	PeakWatts *int32 `json:"peak_watts,omitempty"`
	LowestWatts *int32 `json:"lowest_watts,omitempty"`
	NewJobTime *V0040Uint64NoVal `json:"new_job_time,omitempty"`
	State *int32 `json:"state,omitempty"`
	TimeStartDay *int64 `json:"time_start_day,omitempty"`
}

// NewV0040PowerMgmtData instantiates a new V0040PowerMgmtData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PowerMgmtData() *V0040PowerMgmtData {
	this := V0040PowerMgmtData{}
	return &this
}

// NewV0040PowerMgmtDataWithDefaults instantiates a new V0040PowerMgmtData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PowerMgmtDataWithDefaults() *V0040PowerMgmtData {
	this := V0040PowerMgmtData{}
	return &this
}

// GetMaximumWatts returns the MaximumWatts field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetMaximumWatts() V0040Uint32NoVal {
	if o == nil || IsNil(o.MaximumWatts) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.MaximumWatts
}

// GetMaximumWattsOk returns a tuple with the MaximumWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetMaximumWattsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.MaximumWatts) {
		return nil, false
	}
	return o.MaximumWatts, true
}

// HasMaximumWatts returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasMaximumWatts() bool {
	if o != nil && !IsNil(o.MaximumWatts) {
		return true
	}

	return false
}

// SetMaximumWatts gets a reference to the given V0040Uint32NoVal and assigns it to the MaximumWatts field.
func (o *V0040PowerMgmtData) SetMaximumWatts(v V0040Uint32NoVal) {
	o.MaximumWatts = &v
}

// GetCurrentWatts returns the CurrentWatts field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetCurrentWatts() int32 {
	if o == nil || IsNil(o.CurrentWatts) {
		var ret int32
		return ret
	}
	return *o.CurrentWatts
}

// GetCurrentWattsOk returns a tuple with the CurrentWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetCurrentWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentWatts) {
		return nil, false
	}
	return o.CurrentWatts, true
}

// HasCurrentWatts returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasCurrentWatts() bool {
	if o != nil && !IsNil(o.CurrentWatts) {
		return true
	}

	return false
}

// SetCurrentWatts gets a reference to the given int32 and assigns it to the CurrentWatts field.
func (o *V0040PowerMgmtData) SetCurrentWatts(v int32) {
	o.CurrentWatts = &v
}

// GetTotalEnergy returns the TotalEnergy field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetTotalEnergy() int64 {
	if o == nil || IsNil(o.TotalEnergy) {
		var ret int64
		return ret
	}
	return *o.TotalEnergy
}

// GetTotalEnergyOk returns a tuple with the TotalEnergy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetTotalEnergyOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalEnergy) {
		return nil, false
	}
	return o.TotalEnergy, true
}

// HasTotalEnergy returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasTotalEnergy() bool {
	if o != nil && !IsNil(o.TotalEnergy) {
		return true
	}

	return false
}

// SetTotalEnergy gets a reference to the given int64 and assigns it to the TotalEnergy field.
func (o *V0040PowerMgmtData) SetTotalEnergy(v int64) {
	o.TotalEnergy = &v
}

// GetNewMaximumWatts returns the NewMaximumWatts field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetNewMaximumWatts() int32 {
	if o == nil || IsNil(o.NewMaximumWatts) {
		var ret int32
		return ret
	}
	return *o.NewMaximumWatts
}

// GetNewMaximumWattsOk returns a tuple with the NewMaximumWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetNewMaximumWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.NewMaximumWatts) {
		return nil, false
	}
	return o.NewMaximumWatts, true
}

// HasNewMaximumWatts returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasNewMaximumWatts() bool {
	if o != nil && !IsNil(o.NewMaximumWatts) {
		return true
	}

	return false
}

// SetNewMaximumWatts gets a reference to the given int32 and assigns it to the NewMaximumWatts field.
func (o *V0040PowerMgmtData) SetNewMaximumWatts(v int32) {
	o.NewMaximumWatts = &v
}

// GetPeakWatts returns the PeakWatts field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetPeakWatts() int32 {
	if o == nil || IsNil(o.PeakWatts) {
		var ret int32
		return ret
	}
	return *o.PeakWatts
}

// GetPeakWattsOk returns a tuple with the PeakWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetPeakWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.PeakWatts) {
		return nil, false
	}
	return o.PeakWatts, true
}

// HasPeakWatts returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasPeakWatts() bool {
	if o != nil && !IsNil(o.PeakWatts) {
		return true
	}

	return false
}

// SetPeakWatts gets a reference to the given int32 and assigns it to the PeakWatts field.
func (o *V0040PowerMgmtData) SetPeakWatts(v int32) {
	o.PeakWatts = &v
}

// GetLowestWatts returns the LowestWatts field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetLowestWatts() int32 {
	if o == nil || IsNil(o.LowestWatts) {
		var ret int32
		return ret
	}
	return *o.LowestWatts
}

// GetLowestWattsOk returns a tuple with the LowestWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetLowestWattsOk() (*int32, bool) {
	if o == nil || IsNil(o.LowestWatts) {
		return nil, false
	}
	return o.LowestWatts, true
}

// HasLowestWatts returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasLowestWatts() bool {
	if o != nil && !IsNil(o.LowestWatts) {
		return true
	}

	return false
}

// SetLowestWatts gets a reference to the given int32 and assigns it to the LowestWatts field.
func (o *V0040PowerMgmtData) SetLowestWatts(v int32) {
	o.LowestWatts = &v
}

// GetNewJobTime returns the NewJobTime field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetNewJobTime() V0040Uint64NoVal {
	if o == nil || IsNil(o.NewJobTime) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.NewJobTime
}

// GetNewJobTimeOk returns a tuple with the NewJobTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetNewJobTimeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.NewJobTime) {
		return nil, false
	}
	return o.NewJobTime, true
}

// HasNewJobTime returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasNewJobTime() bool {
	if o != nil && !IsNil(o.NewJobTime) {
		return true
	}

	return false
}

// SetNewJobTime gets a reference to the given V0040Uint64NoVal and assigns it to the NewJobTime field.
func (o *V0040PowerMgmtData) SetNewJobTime(v V0040Uint64NoVal) {
	o.NewJobTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetState() int32 {
	if o == nil || IsNil(o.State) {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetStateOk() (*int32, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *V0040PowerMgmtData) SetState(v int32) {
	o.State = &v
}

// GetTimeStartDay returns the TimeStartDay field value if set, zero value otherwise.
func (o *V0040PowerMgmtData) GetTimeStartDay() int64 {
	if o == nil || IsNil(o.TimeStartDay) {
		var ret int64
		return ret
	}
	return *o.TimeStartDay
}

// GetTimeStartDayOk returns a tuple with the TimeStartDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PowerMgmtData) GetTimeStartDayOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeStartDay) {
		return nil, false
	}
	return o.TimeStartDay, true
}

// HasTimeStartDay returns a boolean if a field has been set.
func (o *V0040PowerMgmtData) HasTimeStartDay() bool {
	if o != nil && !IsNil(o.TimeStartDay) {
		return true
	}

	return false
}

// SetTimeStartDay gets a reference to the given int64 and assigns it to the TimeStartDay field.
func (o *V0040PowerMgmtData) SetTimeStartDay(v int64) {
	o.TimeStartDay = &v
}

func (o V0040PowerMgmtData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PowerMgmtData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaximumWatts) {
		toSerialize["maximum_watts"] = o.MaximumWatts
	}
	if !IsNil(o.CurrentWatts) {
		toSerialize["current_watts"] = o.CurrentWatts
	}
	if !IsNil(o.TotalEnergy) {
		toSerialize["total_energy"] = o.TotalEnergy
	}
	if !IsNil(o.NewMaximumWatts) {
		toSerialize["new_maximum_watts"] = o.NewMaximumWatts
	}
	if !IsNil(o.PeakWatts) {
		toSerialize["peak_watts"] = o.PeakWatts
	}
	if !IsNil(o.LowestWatts) {
		toSerialize["lowest_watts"] = o.LowestWatts
	}
	if !IsNil(o.NewJobTime) {
		toSerialize["new_job_time"] = o.NewJobTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TimeStartDay) {
		toSerialize["time_start_day"] = o.TimeStartDay
	}
	return toSerialize, nil
}

type NullableV0040PowerMgmtData struct {
	value *V0040PowerMgmtData
	isSet bool
}

func (v NullableV0040PowerMgmtData) Get() *V0040PowerMgmtData {
	return v.value
}

func (v *NullableV0040PowerMgmtData) Set(val *V0040PowerMgmtData) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PowerMgmtData) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PowerMgmtData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PowerMgmtData(val *V0040PowerMgmtData) *NullableV0040PowerMgmtData {
	return &NullableV0040PowerMgmtData{value: val, isSet: true}
}

func (v NullableV0040PowerMgmtData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PowerMgmtData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


