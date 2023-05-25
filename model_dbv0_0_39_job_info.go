/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0039JobInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039JobInfo{}

// Dbv0039JobInfo struct for Dbv0039JobInfo
type Dbv0039JobInfo struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []Dbv0039Warning `json:"warnings,omitempty"`
	Jobs []V0039Job `json:"jobs,omitempty"`
}

// NewDbv0039JobInfo instantiates a new Dbv0039JobInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039JobInfo() *Dbv0039JobInfo {
	this := Dbv0039JobInfo{}
	return &this
}

// NewDbv0039JobInfoWithDefaults instantiates a new Dbv0039JobInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039JobInfoWithDefaults() *Dbv0039JobInfo {
	this := Dbv0039JobInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039JobInfo) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039JobInfo) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039JobInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039JobInfo) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039JobInfo) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039JobInfo) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039JobInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039JobInfo) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039JobInfo) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039JobInfo) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039JobInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039JobInfo) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *Dbv0039JobInfo) GetJobs() []V0039Job {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0039Job
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039JobInfo) GetJobsOk() ([]V0039Job, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *Dbv0039JobInfo) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0039Job and assigns it to the Jobs field.
func (o *Dbv0039JobInfo) SetJobs(v []V0039Job) {
	o.Jobs = v
}

func (o Dbv0039JobInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039JobInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	return toSerialize, nil
}

type NullableDbv0039JobInfo struct {
	value *Dbv0039JobInfo
	isSet bool
}

func (v NullableDbv0039JobInfo) Get() *Dbv0039JobInfo {
	return v.value
}

func (v *NullableDbv0039JobInfo) Set(val *Dbv0039JobInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039JobInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039JobInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039JobInfo(val *Dbv0039JobInfo) *NullableDbv0039JobInfo {
	return &NullableDbv0039JobInfo{value: val, isSet: true}
}

func (v NullableDbv0039JobInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039JobInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


