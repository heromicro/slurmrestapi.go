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

// checks if the V0039JobRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobRes{}

// V0039JobRes struct for V0039JobRes
type V0039JobRes struct {
	Nodes *string `json:"nodes,omitempty"`
	AllocatedCores *int32 `json:"allocated_cores,omitempty"`
	AllocatedCpus *int32 `json:"allocated_cpus,omitempty"`
	AllocatedHosts *int32 `json:"allocated_hosts,omitempty"`
	// job node resources
	AllocatedNodes []interface{} `json:"allocated_nodes,omitempty"`
}

// NewV0039JobRes instantiates a new V0039JobRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobRes() *V0039JobRes {
	this := V0039JobRes{}
	return &this
}

// NewV0039JobResWithDefaults instantiates a new V0039JobRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobResWithDefaults() *V0039JobRes {
	this := V0039JobRes{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0039JobRes) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobRes) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0039JobRes) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0039JobRes) SetNodes(v string) {
	o.Nodes = &v
}

// GetAllocatedCores returns the AllocatedCores field value if set, zero value otherwise.
func (o *V0039JobRes) GetAllocatedCores() int32 {
	if o == nil || IsNil(o.AllocatedCores) {
		var ret int32
		return ret
	}
	return *o.AllocatedCores
}

// GetAllocatedCoresOk returns a tuple with the AllocatedCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobRes) GetAllocatedCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedCores) {
		return nil, false
	}
	return o.AllocatedCores, true
}

// HasAllocatedCores returns a boolean if a field has been set.
func (o *V0039JobRes) HasAllocatedCores() bool {
	if o != nil && !IsNil(o.AllocatedCores) {
		return true
	}

	return false
}

// SetAllocatedCores gets a reference to the given int32 and assigns it to the AllocatedCores field.
func (o *V0039JobRes) SetAllocatedCores(v int32) {
	o.AllocatedCores = &v
}

// GetAllocatedCpus returns the AllocatedCpus field value if set, zero value otherwise.
func (o *V0039JobRes) GetAllocatedCpus() int32 {
	if o == nil || IsNil(o.AllocatedCpus) {
		var ret int32
		return ret
	}
	return *o.AllocatedCpus
}

// GetAllocatedCpusOk returns a tuple with the AllocatedCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobRes) GetAllocatedCpusOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedCpus) {
		return nil, false
	}
	return o.AllocatedCpus, true
}

// HasAllocatedCpus returns a boolean if a field has been set.
func (o *V0039JobRes) HasAllocatedCpus() bool {
	if o != nil && !IsNil(o.AllocatedCpus) {
		return true
	}

	return false
}

// SetAllocatedCpus gets a reference to the given int32 and assigns it to the AllocatedCpus field.
func (o *V0039JobRes) SetAllocatedCpus(v int32) {
	o.AllocatedCpus = &v
}

// GetAllocatedHosts returns the AllocatedHosts field value if set, zero value otherwise.
func (o *V0039JobRes) GetAllocatedHosts() int32 {
	if o == nil || IsNil(o.AllocatedHosts) {
		var ret int32
		return ret
	}
	return *o.AllocatedHosts
}

// GetAllocatedHostsOk returns a tuple with the AllocatedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobRes) GetAllocatedHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedHosts) {
		return nil, false
	}
	return o.AllocatedHosts, true
}

// HasAllocatedHosts returns a boolean if a field has been set.
func (o *V0039JobRes) HasAllocatedHosts() bool {
	if o != nil && !IsNil(o.AllocatedHosts) {
		return true
	}

	return false
}

// SetAllocatedHosts gets a reference to the given int32 and assigns it to the AllocatedHosts field.
func (o *V0039JobRes) SetAllocatedHosts(v int32) {
	o.AllocatedHosts = &v
}

// GetAllocatedNodes returns the AllocatedNodes field value if set, zero value otherwise.
func (o *V0039JobRes) GetAllocatedNodes() []interface{} {
	if o == nil || IsNil(o.AllocatedNodes) {
		var ret []interface{}
		return ret
	}
	return o.AllocatedNodes
}

// GetAllocatedNodesOk returns a tuple with the AllocatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobRes) GetAllocatedNodesOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.AllocatedNodes) {
		return nil, false
	}
	return o.AllocatedNodes, true
}

// HasAllocatedNodes returns a boolean if a field has been set.
func (o *V0039JobRes) HasAllocatedNodes() bool {
	if o != nil && !IsNil(o.AllocatedNodes) {
		return true
	}

	return false
}

// SetAllocatedNodes gets a reference to the given []interface{} and assigns it to the AllocatedNodes field.
func (o *V0039JobRes) SetAllocatedNodes(v []interface{}) {
	o.AllocatedNodes = v
}

func (o V0039JobRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.AllocatedCores) {
		toSerialize["allocated_cores"] = o.AllocatedCores
	}
	if !IsNil(o.AllocatedCpus) {
		toSerialize["allocated_cpus"] = o.AllocatedCpus
	}
	if !IsNil(o.AllocatedHosts) {
		toSerialize["allocated_hosts"] = o.AllocatedHosts
	}
	if !IsNil(o.AllocatedNodes) {
		toSerialize["allocated_nodes"] = o.AllocatedNodes
	}
	return toSerialize, nil
}

type NullableV0039JobRes struct {
	value *V0039JobRes
	isSet bool
}

func (v NullableV0039JobRes) Get() *V0039JobRes {
	return v.value
}

func (v *NullableV0039JobRes) Set(val *V0039JobRes) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobRes) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobRes(val *V0039JobRes) *NullableV0039JobRes {
	return &NullableV0039JobRes{value: val, isSet: true}
}

func (v NullableV0039JobRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


