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

// checks if the V0040JobSubmitResponseMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobSubmitResponseMsg{}

// V0040JobSubmitResponseMsg struct for V0040JobSubmitResponseMsg
type V0040JobSubmitResponseMsg struct {
	JobId *int32 `json:"job_id,omitempty"`
	StepId *string `json:"step_id,omitempty"`
	ErrorCode *int32 `json:"error_code,omitempty"`
	Error *string `json:"error,omitempty"`
	JobSubmitUserMsg *string `json:"job_submit_user_msg,omitempty"`
}

// NewV0040JobSubmitResponseMsg instantiates a new V0040JobSubmitResponseMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobSubmitResponseMsg() *V0040JobSubmitResponseMsg {
	this := V0040JobSubmitResponseMsg{}
	return &this
}

// NewV0040JobSubmitResponseMsgWithDefaults instantiates a new V0040JobSubmitResponseMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobSubmitResponseMsgWithDefaults() *V0040JobSubmitResponseMsg {
	this := V0040JobSubmitResponseMsg{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0040JobSubmitResponseMsg) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitResponseMsg) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0040JobSubmitResponseMsg) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0040JobSubmitResponseMsg) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0040JobSubmitResponseMsg) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitResponseMsg) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0040JobSubmitResponseMsg) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0040JobSubmitResponseMsg) SetStepId(v string) {
	o.StepId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *V0040JobSubmitResponseMsg) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitResponseMsg) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *V0040JobSubmitResponseMsg) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *V0040JobSubmitResponseMsg) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V0040JobSubmitResponseMsg) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitResponseMsg) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V0040JobSubmitResponseMsg) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V0040JobSubmitResponseMsg) SetError(v string) {
	o.Error = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *V0040JobSubmitResponseMsg) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitResponseMsg) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *V0040JobSubmitResponseMsg) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *V0040JobSubmitResponseMsg) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

func (o V0040JobSubmitResponseMsg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobSubmitResponseMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
	}
	return toSerialize, nil
}

type NullableV0040JobSubmitResponseMsg struct {
	value *V0040JobSubmitResponseMsg
	isSet bool
}

func (v NullableV0040JobSubmitResponseMsg) Get() *V0040JobSubmitResponseMsg {
	return v.value
}

func (v *NullableV0040JobSubmitResponseMsg) Set(val *V0040JobSubmitResponseMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobSubmitResponseMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobSubmitResponseMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobSubmitResponseMsg(val *V0040JobSubmitResponseMsg) *NullableV0040JobSubmitResponseMsg {
	return &NullableV0040JobSubmitResponseMsg{value: val, isSet: true}
}

func (v NullableV0040JobSubmitResponseMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobSubmitResponseMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


