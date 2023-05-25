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

// checks if the Dbv0037JobStepStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobStepStep{}

// Dbv0037JobStepStep Step details
type Dbv0037JobStepStep struct {
	// Parent job id
	JobId *int32 `json:"job_id,omitempty"`
	Het *Dbv0037JobStepStepHet `json:"het,omitempty"`
	// Step id
	Id *string `json:"id,omitempty"`
	// Step name
	Name *string `json:"name,omitempty"`
}

// NewDbv0037JobStepStep instantiates a new Dbv0037JobStepStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepStep() *Dbv0037JobStepStep {
	this := Dbv0037JobStepStep{}
	return &this
}

// NewDbv0037JobStepStepWithDefaults instantiates a new Dbv0037JobStepStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepStepWithDefaults() *Dbv0037JobStepStep {
	this := Dbv0037JobStepStep{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *Dbv0037JobStepStep) SetJobId(v int32) {
	o.JobId = &v
}

// GetHet returns the Het field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetHet() Dbv0037JobStepStepHet {
	if o == nil || IsNil(o.Het) {
		var ret Dbv0037JobStepStepHet
		return ret
	}
	return *o.Het
}

// GetHetOk returns a tuple with the Het field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetHetOk() (*Dbv0037JobStepStepHet, bool) {
	if o == nil || IsNil(o.Het) {
		return nil, false
	}
	return o.Het, true
}

// HasHet returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasHet() bool {
	if o != nil && !IsNil(o.Het) {
		return true
	}

	return false
}

// SetHet gets a reference to the given Dbv0037JobStepStepHet and assigns it to the Het field.
func (o *Dbv0037JobStepStep) SetHet(v Dbv0037JobStepStepHet) {
	o.Het = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dbv0037JobStepStep) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037JobStepStep) SetName(v string) {
	o.Name = &v
}

func (o Dbv0037JobStepStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobStepStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.Het) {
		toSerialize["het"] = o.Het
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDbv0037JobStepStep struct {
	value *Dbv0037JobStepStep
	isSet bool
}

func (v NullableDbv0037JobStepStep) Get() *Dbv0037JobStepStep {
	return v.value
}

func (v *NullableDbv0037JobStepStep) Set(val *Dbv0037JobStepStep) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepStep) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepStep(val *Dbv0037JobStepStep) *NullableDbv0037JobStepStep {
	return &NullableDbv0037JobStepStep{value: val, isSet: true}
}

func (v NullableDbv0037JobStepStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

