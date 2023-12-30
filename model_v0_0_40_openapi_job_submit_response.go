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

// checks if the V0040OpenapiJobSubmitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiJobSubmitResponse{}

// V0040OpenapiJobSubmitResponse struct for V0040OpenapiJobSubmitResponse
type V0040OpenapiJobSubmitResponse struct {
	Result *V0040JobSubmitResponseMsg `json:"result,omitempty"`
	// submited JobId
	JobId *int32 `json:"job_id,omitempty"`
	// submited StepID
	StepId *string `json:"step_id,omitempty"`
	// job submision user message
	JobSubmitUserMsg *string `json:"job_submit_user_msg,omitempty"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

// NewV0040OpenapiJobSubmitResponse instantiates a new V0040OpenapiJobSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiJobSubmitResponse() *V0040OpenapiJobSubmitResponse {
	this := V0040OpenapiJobSubmitResponse{}
	return &this
}

// NewV0040OpenapiJobSubmitResponseWithDefaults instantiates a new V0040OpenapiJobSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiJobSubmitResponseWithDefaults() *V0040OpenapiJobSubmitResponse {
	this := V0040OpenapiJobSubmitResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetResult() V0040JobSubmitResponseMsg {
	if o == nil || IsNil(o.Result) {
		var ret V0040JobSubmitResponseMsg
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetResultOk() (*V0040JobSubmitResponseMsg, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given V0040JobSubmitResponseMsg and assigns it to the Result field.
func (o *V0040OpenapiJobSubmitResponse) SetResult(v V0040JobSubmitResponseMsg) {
	o.Result = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0040OpenapiJobSubmitResponse) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0040OpenapiJobSubmitResponse) SetStepId(v string) {
	o.StepId = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *V0040OpenapiJobSubmitResponse) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiJobSubmitResponse) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiJobSubmitResponse) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiJobSubmitResponse) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobSubmitResponse) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiJobSubmitResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiJobSubmitResponse) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiJobSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiJobSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableV0040OpenapiJobSubmitResponse struct {
	value *V0040OpenapiJobSubmitResponse
	isSet bool
}

func (v NullableV0040OpenapiJobSubmitResponse) Get() *V0040OpenapiJobSubmitResponse {
	return v.value
}

func (v *NullableV0040OpenapiJobSubmitResponse) Set(val *V0040OpenapiJobSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiJobSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiJobSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiJobSubmitResponse(val *V0040OpenapiJobSubmitResponse) *NullableV0040OpenapiJobSubmitResponse {
	return &NullableV0040OpenapiJobSubmitResponse{value: val, isSet: true}
}

func (v NullableV0040OpenapiJobSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiJobSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


