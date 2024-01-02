/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040ProcessExitCodeVerbose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ProcessExitCodeVerbose{}

// V0040ProcessExitCodeVerbose struct for V0040ProcessExitCodeVerbose
type V0040ProcessExitCodeVerbose struct {
	// Status given by return code
	Status []string `json:"status,omitempty"`
	ReturnCode *V0040Uint32NoVal `json:"return_code,omitempty"`
	Signal *V0040ProcessExitCodeVerboseSignal `json:"signal,omitempty"`
}

// NewV0040ProcessExitCodeVerbose instantiates a new V0040ProcessExitCodeVerbose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ProcessExitCodeVerbose() *V0040ProcessExitCodeVerbose {
	this := V0040ProcessExitCodeVerbose{}
	return &this
}

// NewV0040ProcessExitCodeVerboseWithDefaults instantiates a new V0040ProcessExitCodeVerbose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ProcessExitCodeVerboseWithDefaults() *V0040ProcessExitCodeVerbose {
	this := V0040ProcessExitCodeVerbose{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0040ProcessExitCodeVerbose) GetStatus() []string {
	if o == nil || IsNil(o.Status) {
		var ret []string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ProcessExitCodeVerbose) GetStatusOk() ([]string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0040ProcessExitCodeVerbose) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []string and assigns it to the Status field.
func (o *V0040ProcessExitCodeVerbose) SetStatus(v []string) {
	o.Status = v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *V0040ProcessExitCodeVerbose) GetReturnCode() V0040Uint32NoVal {
	if o == nil || IsNil(o.ReturnCode) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ProcessExitCodeVerbose) GetReturnCodeOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *V0040ProcessExitCodeVerbose) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given V0040Uint32NoVal and assigns it to the ReturnCode field.
func (o *V0040ProcessExitCodeVerbose) SetReturnCode(v V0040Uint32NoVal) {
	o.ReturnCode = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *V0040ProcessExitCodeVerbose) GetSignal() V0040ProcessExitCodeVerboseSignal {
	if o == nil || IsNil(o.Signal) {
		var ret V0040ProcessExitCodeVerboseSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ProcessExitCodeVerbose) GetSignalOk() (*V0040ProcessExitCodeVerboseSignal, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *V0040ProcessExitCodeVerbose) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given V0040ProcessExitCodeVerboseSignal and assigns it to the Signal field.
func (o *V0040ProcessExitCodeVerbose) SetSignal(v V0040ProcessExitCodeVerboseSignal) {
	o.Signal = &v
}

func (o V0040ProcessExitCodeVerbose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ProcessExitCodeVerbose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	return toSerialize, nil
}

type NullableV0040ProcessExitCodeVerbose struct {
	value *V0040ProcessExitCodeVerbose
	isSet bool
}

func (v NullableV0040ProcessExitCodeVerbose) Get() *V0040ProcessExitCodeVerbose {
	return v.value
}

func (v *NullableV0040ProcessExitCodeVerbose) Set(val *V0040ProcessExitCodeVerbose) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ProcessExitCodeVerbose) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ProcessExitCodeVerbose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ProcessExitCodeVerbose(val *V0040ProcessExitCodeVerbose) *NullableV0040ProcessExitCodeVerbose {
	return &NullableV0040ProcessExitCodeVerbose{value: val, isSet: true}
}

func (v NullableV0040ProcessExitCodeVerbose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ProcessExitCodeVerbose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


