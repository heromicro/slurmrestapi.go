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

// checks if the V0039JobSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobSubmission{}

// V0039JobSubmission struct for V0039JobSubmission
type V0039JobSubmission struct {
	// Executable script (full contents) to run in batch step for all job components
	Script *string `json:"script,omitempty"`
	Job *V0039JobDescMsg `json:"job,omitempty"`
	Jobs []V0039JobDescMsg `json:"jobs,omitempty"`
}

// NewV0039JobSubmission instantiates a new V0039JobSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobSubmission() *V0039JobSubmission {
	this := V0039JobSubmission{}
	return &this
}

// NewV0039JobSubmissionWithDefaults instantiates a new V0039JobSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobSubmissionWithDefaults() *V0039JobSubmission {
	this := V0039JobSubmission{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *V0039JobSubmission) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobSubmission) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *V0039JobSubmission) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *V0039JobSubmission) SetScript(v string) {
	o.Script = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039JobSubmission) GetJob() V0039JobDescMsg {
	if o == nil || IsNil(o.Job) {
		var ret V0039JobDescMsg
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobSubmission) GetJobOk() (*V0039JobDescMsg, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039JobSubmission) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0039JobDescMsg and assigns it to the Job field.
func (o *V0039JobSubmission) SetJob(v V0039JobDescMsg) {
	o.Job = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0039JobSubmission) GetJobs() []V0039JobDescMsg {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0039JobDescMsg
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobSubmission) GetJobsOk() ([]V0039JobDescMsg, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0039JobSubmission) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0039JobDescMsg and assigns it to the Jobs field.
func (o *V0039JobSubmission) SetJobs(v []V0039JobDescMsg) {
	o.Jobs = v
}

func (o V0039JobSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	return toSerialize, nil
}

type NullableV0039JobSubmission struct {
	value *V0039JobSubmission
	isSet bool
}

func (v NullableV0039JobSubmission) Get() *V0039JobSubmission {
	return v.value
}

func (v *NullableV0039JobSubmission) Set(val *V0039JobSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobSubmission(val *V0039JobSubmission) *NullableV0039JobSubmission {
	return &NullableV0039JobSubmission{value: val, isSet: true}
}

func (v NullableV0039JobSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


