/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.40
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0040JobTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobTime{}

// V0040JobTime struct for V0040JobTime
type V0040JobTime struct {
	Elapsed *int32 `json:"elapsed,omitempty"`
	Eligible *int64 `json:"eligible,omitempty"`
	End *int64 `json:"end,omitempty"`
	Start *int64 `json:"start,omitempty"`
	Submission *int64 `json:"submission,omitempty"`
	Suspended *int32 `json:"suspended,omitempty"`
	System *V0040JobTimeSystem `json:"system,omitempty"`
	Limit *V0040Uint32NoVal `json:"limit,omitempty"`
	Total *V0040JobTimeSystem `json:"total,omitempty"`
	User *V0040JobTimeSystem `json:"user,omitempty"`
}

// NewV0040JobTime instantiates a new V0040JobTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobTime() *V0040JobTime {
	this := V0040JobTime{}
	return &this
}

// NewV0040JobTimeWithDefaults instantiates a new V0040JobTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobTimeWithDefaults() *V0040JobTime {
	this := V0040JobTime{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *V0040JobTime) GetElapsed() int32 {
	if o == nil || IsNil(o.Elapsed) {
		var ret int32
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetElapsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Elapsed) {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *V0040JobTime) HasElapsed() bool {
	if o != nil && !IsNil(o.Elapsed) {
		return true
	}

	return false
}

// SetElapsed gets a reference to the given int32 and assigns it to the Elapsed field.
func (o *V0040JobTime) SetElapsed(v int32) {
	o.Elapsed = &v
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *V0040JobTime) GetEligible() int64 {
	if o == nil || IsNil(o.Eligible) {
		var ret int64
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetEligibleOk() (*int64, bool) {
	if o == nil || IsNil(o.Eligible) {
		return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *V0040JobTime) HasEligible() bool {
	if o != nil && !IsNil(o.Eligible) {
		return true
	}

	return false
}

// SetEligible gets a reference to the given int64 and assigns it to the Eligible field.
func (o *V0040JobTime) SetEligible(v int64) {
	o.Eligible = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V0040JobTime) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V0040JobTime) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *V0040JobTime) SetEnd(v int64) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *V0040JobTime) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *V0040JobTime) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *V0040JobTime) SetStart(v int64) {
	o.Start = &v
}

// GetSubmission returns the Submission field value if set, zero value otherwise.
func (o *V0040JobTime) GetSubmission() int64 {
	if o == nil || IsNil(o.Submission) {
		var ret int64
		return ret
	}
	return *o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetSubmissionOk() (*int64, bool) {
	if o == nil || IsNil(o.Submission) {
		return nil, false
	}
	return o.Submission, true
}

// HasSubmission returns a boolean if a field has been set.
func (o *V0040JobTime) HasSubmission() bool {
	if o != nil && !IsNil(o.Submission) {
		return true
	}

	return false
}

// SetSubmission gets a reference to the given int64 and assigns it to the Submission field.
func (o *V0040JobTime) SetSubmission(v int64) {
	o.Submission = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *V0040JobTime) GetSuspended() int32 {
	if o == nil || IsNil(o.Suspended) {
		var ret int32
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetSuspendedOk() (*int32, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *V0040JobTime) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given int32 and assigns it to the Suspended field.
func (o *V0040JobTime) SetSuspended(v int32) {
	o.Suspended = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *V0040JobTime) GetSystem() V0040JobTimeSystem {
	if o == nil || IsNil(o.System) {
		var ret V0040JobTimeSystem
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetSystemOk() (*V0040JobTimeSystem, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *V0040JobTime) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given V0040JobTimeSystem and assigns it to the System field.
func (o *V0040JobTime) SetSystem(v V0040JobTimeSystem) {
	o.System = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V0040JobTime) GetLimit() V0040Uint32NoVal {
	if o == nil || IsNil(o.Limit) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetLimitOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V0040JobTime) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given V0040Uint32NoVal and assigns it to the Limit field.
func (o *V0040JobTime) SetLimit(v V0040Uint32NoVal) {
	o.Limit = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040JobTime) GetTotal() V0040JobTimeSystem {
	if o == nil || IsNil(o.Total) {
		var ret V0040JobTimeSystem
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetTotalOk() (*V0040JobTimeSystem, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040JobTime) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given V0040JobTimeSystem and assigns it to the Total field.
func (o *V0040JobTime) SetTotal(v V0040JobTimeSystem) {
	o.Total = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0040JobTime) GetUser() V0040JobTimeSystem {
	if o == nil || IsNil(o.User) {
		var ret V0040JobTimeSystem
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTime) GetUserOk() (*V0040JobTimeSystem, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0040JobTime) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given V0040JobTimeSystem and assigns it to the User field.
func (o *V0040JobTime) SetUser(v V0040JobTimeSystem) {
	o.User = &v
}

func (o V0040JobTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Elapsed) {
		toSerialize["elapsed"] = o.Elapsed
	}
	if !IsNil(o.Eligible) {
		toSerialize["eligible"] = o.Eligible
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Submission) {
		toSerialize["submission"] = o.Submission
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0040JobTime struct {
	value *V0040JobTime
	isSet bool
}

func (v NullableV0040JobTime) Get() *V0040JobTime {
	return v.value
}

func (v *NullableV0040JobTime) Set(val *V0040JobTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobTime(val *V0040JobTime) *NullableV0040JobTime {
	return &NullableV0040JobTime{value: val, isSet: true}
}

func (v NullableV0040JobTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


