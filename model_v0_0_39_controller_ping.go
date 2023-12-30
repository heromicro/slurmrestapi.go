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

// checks if the V0039ControllerPing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039ControllerPing{}

// V0039ControllerPing struct for V0039ControllerPing
type V0039ControllerPing struct {
	Hostname *string `json:"hostname,omitempty"`
	Pinged *string `json:"pinged,omitempty"`
	Latency *int64 `json:"latency,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// NewV0039ControllerPing instantiates a new V0039ControllerPing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039ControllerPing() *V0039ControllerPing {
	this := V0039ControllerPing{}
	return &this
}

// NewV0039ControllerPingWithDefaults instantiates a new V0039ControllerPing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039ControllerPingWithDefaults() *V0039ControllerPing {
	this := V0039ControllerPing{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V0039ControllerPing) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ControllerPing) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V0039ControllerPing) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *V0039ControllerPing) SetHostname(v string) {
	o.Hostname = &v
}

// GetPinged returns the Pinged field value if set, zero value otherwise.
func (o *V0039ControllerPing) GetPinged() string {
	if o == nil || IsNil(o.Pinged) {
		var ret string
		return ret
	}
	return *o.Pinged
}

// GetPingedOk returns a tuple with the Pinged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ControllerPing) GetPingedOk() (*string, bool) {
	if o == nil || IsNil(o.Pinged) {
		return nil, false
	}
	return o.Pinged, true
}

// HasPinged returns a boolean if a field has been set.
func (o *V0039ControllerPing) HasPinged() bool {
	if o != nil && !IsNil(o.Pinged) {
		return true
	}

	return false
}

// SetPinged gets a reference to the given string and assigns it to the Pinged field.
func (o *V0039ControllerPing) SetPinged(v string) {
	o.Pinged = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *V0039ControllerPing) GetLatency() int64 {
	if o == nil || IsNil(o.Latency) {
		var ret int64
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ControllerPing) GetLatencyOk() (*int64, bool) {
	if o == nil || IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *V0039ControllerPing) HasLatency() bool {
	if o != nil && !IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int64 and assigns it to the Latency field.
func (o *V0039ControllerPing) SetLatency(v int64) {
	o.Latency = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V0039ControllerPing) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ControllerPing) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V0039ControllerPing) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *V0039ControllerPing) SetMode(v string) {
	o.Mode = &v
}

func (o V0039ControllerPing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039ControllerPing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Pinged) {
		toSerialize["pinged"] = o.Pinged
	}
	if !IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableV0039ControllerPing struct {
	value *V0039ControllerPing
	isSet bool
}

func (v NullableV0039ControllerPing) Get() *V0039ControllerPing {
	return v.value
}

func (v *NullableV0039ControllerPing) Set(val *V0039ControllerPing) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039ControllerPing) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039ControllerPing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039ControllerPing(val *V0039ControllerPing) *NullableV0039ControllerPing {
	return &NullableV0039ControllerPing{value: val, isSet: true}
}

func (v NullableV0039ControllerPing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039ControllerPing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

