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

// checks if the Dbv0037JobHet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobHet{}

// Dbv0037JobHet Heterogeneous Job details (optional)
type Dbv0037JobHet struct {
	// Parent HetJob id
	JobId map[string]interface{} `json:"job_id,omitempty"`
	// Offset of this job to parent
	JobOffset map[string]interface{} `json:"job_offset,omitempty"`
}

// NewDbv0037JobHet instantiates a new Dbv0037JobHet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobHet() *Dbv0037JobHet {
	this := Dbv0037JobHet{}
	return &this
}

// NewDbv0037JobHetWithDefaults instantiates a new Dbv0037JobHet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobHetWithDefaults() *Dbv0037JobHet {
	this := Dbv0037JobHet{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Dbv0037JobHet) GetJobId() map[string]interface{} {
	if o == nil || IsNil(o.JobId) {
		var ret map[string]interface{}
		return ret
	}
	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobHet) GetJobIdOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.JobId) {
		return map[string]interface{}{}, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Dbv0037JobHet) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given map[string]interface{} and assigns it to the JobId field.
func (o *Dbv0037JobHet) SetJobId(v map[string]interface{}) {
	o.JobId = v
}

// GetJobOffset returns the JobOffset field value if set, zero value otherwise.
func (o *Dbv0037JobHet) GetJobOffset() map[string]interface{} {
	if o == nil || IsNil(o.JobOffset) {
		var ret map[string]interface{}
		return ret
	}
	return o.JobOffset
}

// GetJobOffsetOk returns a tuple with the JobOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobHet) GetJobOffsetOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.JobOffset) {
		return map[string]interface{}{}, false
	}
	return o.JobOffset, true
}

// HasJobOffset returns a boolean if a field has been set.
func (o *Dbv0037JobHet) HasJobOffset() bool {
	if o != nil && !IsNil(o.JobOffset) {
		return true
	}

	return false
}

// SetJobOffset gets a reference to the given map[string]interface{} and assigns it to the JobOffset field.
func (o *Dbv0037JobHet) SetJobOffset(v map[string]interface{}) {
	o.JobOffset = v
}

func (o Dbv0037JobHet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobHet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobOffset) {
		toSerialize["job_offset"] = o.JobOffset
	}
	return toSerialize, nil
}

type NullableDbv0037JobHet struct {
	value *Dbv0037JobHet
	isSet bool
}

func (v NullableDbv0037JobHet) Get() *Dbv0037JobHet {
	return v.value
}

func (v *NullableDbv0037JobHet) Set(val *Dbv0037JobHet) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobHet) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobHet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobHet(val *Dbv0037JobHet) *NullableDbv0037JobHet {
	return &NullableDbv0037JobHet{value: val, isSet: true}
}

func (v NullableDbv0037JobHet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobHet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


