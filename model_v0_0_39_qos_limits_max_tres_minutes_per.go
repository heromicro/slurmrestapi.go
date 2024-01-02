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

// checks if the V0039QosLimitsMaxTresMinutesPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMaxTresMinutesPer{}

// V0039QosLimitsMaxTresMinutesPer struct for V0039QosLimitsMaxTresMinutesPer
type V0039QosLimitsMaxTresMinutesPer struct {
	Qos []V0039Tres `json:"qos,omitempty"`
	Job []V0039Tres `json:"job,omitempty"`
	Account []V0039Tres `json:"account,omitempty"`
	User []V0039Tres `json:"user,omitempty"`
}

// NewV0039QosLimitsMaxTresMinutesPer instantiates a new V0039QosLimitsMaxTresMinutesPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMaxTresMinutesPer() *V0039QosLimitsMaxTresMinutesPer {
	this := V0039QosLimitsMaxTresMinutesPer{}
	return &this
}

// NewV0039QosLimitsMaxTresMinutesPerWithDefaults instantiates a new V0039QosLimitsMaxTresMinutesPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMaxTresMinutesPerWithDefaults() *V0039QosLimitsMaxTresMinutesPer {
	this := V0039QosLimitsMaxTresMinutesPer{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxTresMinutesPer) GetQos() []V0039Tres {
	if o == nil || IsNil(o.Qos) {
		var ret []V0039Tres
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) GetQosOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []V0039Tres and assigns it to the Qos field.
func (o *V0039QosLimitsMaxTresMinutesPer) SetQos(v []V0039Tres) {
	o.Qos = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxTresMinutesPer) GetJob() []V0039Tres {
	if o == nil || IsNil(o.Job) {
		var ret []V0039Tres
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) GetJobOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0039Tres and assigns it to the Job field.
func (o *V0039QosLimitsMaxTresMinutesPer) SetJob(v []V0039Tres) {
	o.Job = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxTresMinutesPer) GetAccount() []V0039Tres {
	if o == nil || IsNil(o.Account) {
		var ret []V0039Tres
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) GetAccountOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []V0039Tres and assigns it to the Account field.
func (o *V0039QosLimitsMaxTresMinutesPer) SetAccount(v []V0039Tres) {
	o.Account = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxTresMinutesPer) GetUser() []V0039Tres {
	if o == nil || IsNil(o.User) {
		var ret []V0039Tres
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) GetUserOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxTresMinutesPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []V0039Tres and assigns it to the User field.
func (o *V0039QosLimitsMaxTresMinutesPer) SetUser(v []V0039Tres) {
	o.User = v
}

func (o V0039QosLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMaxTresMinutesPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMaxTresMinutesPer struct {
	value *V0039QosLimitsMaxTresMinutesPer
	isSet bool
}

func (v NullableV0039QosLimitsMaxTresMinutesPer) Get() *V0039QosLimitsMaxTresMinutesPer {
	return v.value
}

func (v *NullableV0039QosLimitsMaxTresMinutesPer) Set(val *V0039QosLimitsMaxTresMinutesPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMaxTresMinutesPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMaxTresMinutesPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMaxTresMinutesPer(val *V0039QosLimitsMaxTresMinutesPer) *NullableV0039QosLimitsMaxTresMinutesPer {
	return &NullableV0039QosLimitsMaxTresMinutesPer{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMaxTresMinutesPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


