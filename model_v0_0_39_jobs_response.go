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

// checks if the V0039JobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobsResponse{}

// V0039JobsResponse struct for V0039JobsResponse
type V0039JobsResponse struct {
	Meta *V0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []V0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []V0039Warning `json:"warnings,omitempty"`
	Jobs []V0039JobInfo `json:"jobs,omitempty"`
}

// NewV0039JobsResponse instantiates a new V0039JobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobsResponse() *V0039JobsResponse {
	this := V0039JobsResponse{}
	return &this
}

// NewV0039JobsResponseWithDefaults instantiates a new V0039JobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobsResponseWithDefaults() *V0039JobsResponse {
	this := V0039JobsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0039JobsResponse) GetMeta() V0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobsResponse) GetMetaOk() (*V0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0039JobsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0039Meta and assigns it to the Meta field.
func (o *V0039JobsResponse) SetMeta(v V0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0039JobsResponse) GetErrors() []V0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobsResponse) GetErrorsOk() ([]V0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0039JobsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0039Error and assigns it to the Errors field.
func (o *V0039JobsResponse) SetErrors(v []V0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0039JobsResponse) GetWarnings() []V0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobsResponse) GetWarningsOk() ([]V0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0039JobsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0039Warning and assigns it to the Warnings field.
func (o *V0039JobsResponse) SetWarnings(v []V0039Warning) {
	o.Warnings = v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0039JobsResponse) GetJobs() []V0039JobInfo {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0039JobInfo
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobsResponse) GetJobsOk() ([]V0039JobInfo, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0039JobsResponse) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0039JobInfo and assigns it to the Jobs field.
func (o *V0039JobsResponse) SetJobs(v []V0039JobInfo) {
	o.Jobs = v
}

func (o V0039JobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableV0039JobsResponse struct {
	value *V0039JobsResponse
	isSet bool
}

func (v NullableV0039JobsResponse) Get() *V0039JobsResponse {
	return v.value
}

func (v *NullableV0039JobsResponse) Set(val *V0039JobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobsResponse(val *V0039JobsResponse) *NullableV0039JobsResponse {
	return &NullableV0039JobsResponse{value: val, isSet: true}
}

func (v NullableV0039JobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

