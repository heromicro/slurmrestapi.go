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

// checks if the V0037ReservationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037ReservationsResponse{}

// V0037ReservationsResponse struct for V0037ReservationsResponse
type V0037ReservationsResponse struct {
	// slurm errors
	Errors []V0037Error `json:"errors,omitempty"`
	// reservation info
	Reservations []V0037Reservation `json:"reservations,omitempty"`
}

// NewV0037ReservationsResponse instantiates a new V0037ReservationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037ReservationsResponse() *V0037ReservationsResponse {
	this := V0037ReservationsResponse{}
	return &this
}

// NewV0037ReservationsResponseWithDefaults instantiates a new V0037ReservationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037ReservationsResponseWithDefaults() *V0037ReservationsResponse {
	this := V0037ReservationsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0037ReservationsResponse) GetErrors() []V0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037ReservationsResponse) GetErrorsOk() ([]V0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0037ReservationsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0037Error and assigns it to the Errors field.
func (o *V0037ReservationsResponse) SetErrors(v []V0037Error) {
	o.Errors = v
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *V0037ReservationsResponse) GetReservations() []V0037Reservation {
	if o == nil || IsNil(o.Reservations) {
		var ret []V0037Reservation
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037ReservationsResponse) GetReservationsOk() ([]V0037Reservation, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *V0037ReservationsResponse) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []V0037Reservation and assigns it to the Reservations field.
func (o *V0037ReservationsResponse) SetReservations(v []V0037Reservation) {
	o.Reservations = v
}

func (o V0037ReservationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037ReservationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Reservations) {
		toSerialize["reservations"] = o.Reservations
	}
	return toSerialize, nil
}

type NullableV0037ReservationsResponse struct {
	value *V0037ReservationsResponse
	isSet bool
}

func (v NullableV0037ReservationsResponse) Get() *V0037ReservationsResponse {
	return v.value
}

func (v *NullableV0037ReservationsResponse) Set(val *V0037ReservationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037ReservationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037ReservationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037ReservationsResponse(val *V0037ReservationsResponse) *NullableV0037ReservationsResponse {
	return &NullableV0037ReservationsResponse{value: val, isSet: true}
}

func (v NullableV0037ReservationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037ReservationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


