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

// checks if the V0039JobUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobUpdateResponse{}

// V0039JobUpdateResponse struct for V0039JobUpdateResponse
type V0039JobUpdateResponse struct {
	Meta *V0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []V0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []V0039Warning `json:"warnings,omitempty"`
	// Result per ArrayJob
	Results []V0039JobArrayResponseMsgInner `json:"results,omitempty"`
}

// NewV0039JobUpdateResponse instantiates a new V0039JobUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobUpdateResponse() *V0039JobUpdateResponse {
	this := V0039JobUpdateResponse{}
	return &this
}

// NewV0039JobUpdateResponseWithDefaults instantiates a new V0039JobUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobUpdateResponseWithDefaults() *V0039JobUpdateResponse {
	this := V0039JobUpdateResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0039JobUpdateResponse) GetMeta() V0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobUpdateResponse) GetMetaOk() (*V0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0039JobUpdateResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0039Meta and assigns it to the Meta field.
func (o *V0039JobUpdateResponse) SetMeta(v V0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0039JobUpdateResponse) GetErrors() []V0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobUpdateResponse) GetErrorsOk() ([]V0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0039JobUpdateResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0039Error and assigns it to the Errors field.
func (o *V0039JobUpdateResponse) SetErrors(v []V0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0039JobUpdateResponse) GetWarnings() []V0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobUpdateResponse) GetWarningsOk() ([]V0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0039JobUpdateResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0039Warning and assigns it to the Warnings field.
func (o *V0039JobUpdateResponse) SetWarnings(v []V0039Warning) {
	o.Warnings = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V0039JobUpdateResponse) GetResults() []V0039JobArrayResponseMsgInner {
	if o == nil || IsNil(o.Results) {
		var ret []V0039JobArrayResponseMsgInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobUpdateResponse) GetResultsOk() ([]V0039JobArrayResponseMsgInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V0039JobUpdateResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []V0039JobArrayResponseMsgInner and assigns it to the Results field.
func (o *V0039JobUpdateResponse) SetResults(v []V0039JobArrayResponseMsgInner) {
	o.Results = v
}

func (o V0039JobUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobUpdateResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableV0039JobUpdateResponse struct {
	value *V0039JobUpdateResponse
	isSet bool
}

func (v NullableV0039JobUpdateResponse) Get() *V0039JobUpdateResponse {
	return v.value
}

func (v *NullableV0039JobUpdateResponse) Set(val *V0039JobUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobUpdateResponse(val *V0039JobUpdateResponse) *NullableV0039JobUpdateResponse {
	return &NullableV0039JobUpdateResponse{value: val, isSet: true}
}

func (v NullableV0039JobUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


