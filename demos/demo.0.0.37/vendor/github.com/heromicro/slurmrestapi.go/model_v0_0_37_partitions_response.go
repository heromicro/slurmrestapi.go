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

// checks if the V0037PartitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037PartitionsResponse{}

// V0037PartitionsResponse struct for V0037PartitionsResponse
type V0037PartitionsResponse struct {
	// slurm errors
	Errors []V0037Error `json:"errors,omitempty"`
	// partition info
	Partitions []V0037Partition `json:"partitions,omitempty"`
}

// NewV0037PartitionsResponse instantiates a new V0037PartitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037PartitionsResponse() *V0037PartitionsResponse {
	this := V0037PartitionsResponse{}
	return &this
}

// NewV0037PartitionsResponseWithDefaults instantiates a new V0037PartitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037PartitionsResponseWithDefaults() *V0037PartitionsResponse {
	this := V0037PartitionsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0037PartitionsResponse) GetErrors() []V0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037PartitionsResponse) GetErrorsOk() ([]V0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0037PartitionsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0037Error and assigns it to the Errors field.
func (o *V0037PartitionsResponse) SetErrors(v []V0037Error) {
	o.Errors = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *V0037PartitionsResponse) GetPartitions() []V0037Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []V0037Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037PartitionsResponse) GetPartitionsOk() ([]V0037Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *V0037PartitionsResponse) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []V0037Partition and assigns it to the Partitions field.
func (o *V0037PartitionsResponse) SetPartitions(v []V0037Partition) {
	o.Partitions = v
}

func (o V0037PartitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037PartitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return toSerialize, nil
}

type NullableV0037PartitionsResponse struct {
	value *V0037PartitionsResponse
	isSet bool
}

func (v NullableV0037PartitionsResponse) Get() *V0037PartitionsResponse {
	return v.value
}

func (v *NullableV0037PartitionsResponse) Set(val *V0037PartitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037PartitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037PartitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037PartitionsResponse(val *V0037PartitionsResponse) *NullableV0037PartitionsResponse {
	return &NullableV0037PartitionsResponse{value: val, isSet: true}
}

func (v NullableV0037PartitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037PartitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


