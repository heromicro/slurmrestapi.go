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

// checks if the Dbv0037JobStepTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037JobStepTime{}

// Dbv0037JobStepTime Time properties
type Dbv0037JobStepTime struct {
	// Total time elapsed
	Elapsed *int32 `json:"elapsed,omitempty"`
	// Timestamp of when job ended
	End *int32 `json:"end,omitempty"`
	// Timestamp of when job started
	Start *int32 `json:"start,omitempty"`
	// Timestamp of when job last suspended
	Suspended *int32 `json:"suspended,omitempty"`
	System *Dbv0037JobTimeSystem `json:"system,omitempty"`
	Total *Dbv0037JobTimeTotal `json:"total,omitempty"`
	User *Dbv0037JobTimeUser `json:"user,omitempty"`
}

// NewDbv0037JobStepTime instantiates a new Dbv0037JobStepTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepTime() *Dbv0037JobStepTime {
	this := Dbv0037JobStepTime{}
	return &this
}

// NewDbv0037JobStepTimeWithDefaults instantiates a new Dbv0037JobStepTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepTimeWithDefaults() *Dbv0037JobStepTime {
	this := Dbv0037JobStepTime{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetElapsed() int32 {
	if o == nil || IsNil(o.Elapsed) {
		var ret int32
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetElapsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Elapsed) {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasElapsed() bool {
	if o != nil && !IsNil(o.Elapsed) {
		return true
	}

	return false
}

// SetElapsed gets a reference to the given int32 and assigns it to the Elapsed field.
func (o *Dbv0037JobStepTime) SetElapsed(v int32) {
	o.Elapsed = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *Dbv0037JobStepTime) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *Dbv0037JobStepTime) SetStart(v int32) {
	o.Start = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetSuspended() int32 {
	if o == nil || IsNil(o.Suspended) {
		var ret int32
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetSuspendedOk() (*int32, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given int32 and assigns it to the Suspended field.
func (o *Dbv0037JobStepTime) SetSuspended(v int32) {
	o.Suspended = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetSystem() Dbv0037JobTimeSystem {
	if o == nil || IsNil(o.System) {
		var ret Dbv0037JobTimeSystem
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetSystemOk() (*Dbv0037JobTimeSystem, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given Dbv0037JobTimeSystem and assigns it to the System field.
func (o *Dbv0037JobStepTime) SetSystem(v Dbv0037JobTimeSystem) {
	o.System = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetTotal() Dbv0037JobTimeTotal {
	if o == nil || IsNil(o.Total) {
		var ret Dbv0037JobTimeTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetTotalOk() (*Dbv0037JobTimeTotal, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given Dbv0037JobTimeTotal and assigns it to the Total field.
func (o *Dbv0037JobStepTime) SetTotal(v Dbv0037JobTimeTotal) {
	o.Total = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0037JobStepTime) GetUser() Dbv0037JobTimeUser {
	if o == nil || IsNil(o.User) {
		var ret Dbv0037JobTimeUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTime) GetUserOk() (*Dbv0037JobTimeUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0037JobStepTime) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given Dbv0037JobTimeUser and assigns it to the User field.
func (o *Dbv0037JobStepTime) SetUser(v Dbv0037JobTimeUser) {
	o.User = &v
}

func (o Dbv0037JobStepTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037JobStepTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Elapsed) {
		toSerialize["elapsed"] = o.Elapsed
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableDbv0037JobStepTime struct {
	value *Dbv0037JobStepTime
	isSet bool
}

func (v NullableDbv0037JobStepTime) Get() *Dbv0037JobStepTime {
	return v.value
}

func (v *NullableDbv0037JobStepTime) Set(val *Dbv0037JobStepTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepTime(val *Dbv0037JobStepTime) *NullableDbv0037JobStepTime {
	return &NullableDbv0037JobStepTime{value: val, isSet: true}
}

func (v NullableDbv0037JobStepTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

