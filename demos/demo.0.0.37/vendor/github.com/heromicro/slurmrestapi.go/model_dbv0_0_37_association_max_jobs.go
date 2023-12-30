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

// checks if the Dbv0037AssociationMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037AssociationMaxJobs{}

// Dbv0037AssociationMaxJobs Max jobs settings
type Dbv0037AssociationMaxJobs struct {
	// Max TRES for active total jobs
	Active *int32 `json:"active,omitempty"`
	// Max TRES for job accruing priority
	Accruing *int32 `json:"accruing,omitempty"`
	// Max TRES for job total submitted
	Total *int32 `json:"total,omitempty"`
	Per *Dbv0037AssociationMaxJobsPer `json:"per,omitempty"`
}

// NewDbv0037AssociationMaxJobs instantiates a new Dbv0037AssociationMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMaxJobs() *Dbv0037AssociationMaxJobs {
	this := Dbv0037AssociationMaxJobs{}
	return &this
}

// NewDbv0037AssociationMaxJobsWithDefaults instantiates a new Dbv0037AssociationMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMaxJobsWithDefaults() *Dbv0037AssociationMaxJobs {
	this := Dbv0037AssociationMaxJobs{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxJobs) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxJobs) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxJobs) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *Dbv0037AssociationMaxJobs) SetActive(v int32) {
	o.Active = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxJobs) GetAccruing() int32 {
	if o == nil || IsNil(o.Accruing) {
		var ret int32
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxJobs) GetAccruingOk() (*int32, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxJobs) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given int32 and assigns it to the Accruing field.
func (o *Dbv0037AssociationMaxJobs) SetAccruing(v int32) {
	o.Accruing = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxJobs) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxJobs) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxJobs) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *Dbv0037AssociationMaxJobs) SetTotal(v int32) {
	o.Total = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxJobs) GetPer() Dbv0037AssociationMaxJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0037AssociationMaxJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxJobs) GetPerOk() (*Dbv0037AssociationMaxJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0037AssociationMaxJobsPer and assigns it to the Per field.
func (o *Dbv0037AssociationMaxJobs) SetPer(v Dbv0037AssociationMaxJobsPer) {
	o.Per = &v
}

func (o Dbv0037AssociationMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037AssociationMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0037AssociationMaxJobs struct {
	value *Dbv0037AssociationMaxJobs
	isSet bool
}

func (v NullableDbv0037AssociationMaxJobs) Get() *Dbv0037AssociationMaxJobs {
	return v.value
}

func (v *NullableDbv0037AssociationMaxJobs) Set(val *Dbv0037AssociationMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMaxJobs(val *Dbv0037AssociationMaxJobs) *NullableDbv0037AssociationMaxJobs {
	return &NullableDbv0037AssociationMaxJobs{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


