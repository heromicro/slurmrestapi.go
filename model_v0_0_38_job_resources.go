/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0038JobResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038JobResources{}

// V0038JobResources struct for V0038JobResources
type V0038JobResources struct {
	// list of assigned job nodes
	Nodes *string `json:"nodes,omitempty"`
	// number of assigned job cpus
	AllocatedCpus *int32 `json:"allocated_cpus,omitempty"`
	// number of assigned job hosts
	AllocatedHosts *int32 `json:"allocated_hosts,omitempty"`
	// array of allocated nodes
	AllocatedNodes []V0038NodeAllocation `json:"allocated_nodes,omitempty"`
}

// NewV0038JobResources instantiates a new V0038JobResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038JobResources() *V0038JobResources {
	this := V0038JobResources{}
	return &this
}

// NewV0038JobResourcesWithDefaults instantiates a new V0038JobResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038JobResourcesWithDefaults() *V0038JobResources {
	this := V0038JobResources{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0038JobResources) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobResources) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0038JobResources) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0038JobResources) SetNodes(v string) {
	o.Nodes = &v
}

// GetAllocatedCpus returns the AllocatedCpus field value if set, zero value otherwise.
func (o *V0038JobResources) GetAllocatedCpus() int32 {
	if o == nil || IsNil(o.AllocatedCpus) {
		var ret int32
		return ret
	}
	return *o.AllocatedCpus
}

// GetAllocatedCpusOk returns a tuple with the AllocatedCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobResources) GetAllocatedCpusOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedCpus) {
		return nil, false
	}
	return o.AllocatedCpus, true
}

// HasAllocatedCpus returns a boolean if a field has been set.
func (o *V0038JobResources) HasAllocatedCpus() bool {
	if o != nil && !IsNil(o.AllocatedCpus) {
		return true
	}

	return false
}

// SetAllocatedCpus gets a reference to the given int32 and assigns it to the AllocatedCpus field.
func (o *V0038JobResources) SetAllocatedCpus(v int32) {
	o.AllocatedCpus = &v
}

// GetAllocatedHosts returns the AllocatedHosts field value if set, zero value otherwise.
func (o *V0038JobResources) GetAllocatedHosts() int32 {
	if o == nil || IsNil(o.AllocatedHosts) {
		var ret int32
		return ret
	}
	return *o.AllocatedHosts
}

// GetAllocatedHostsOk returns a tuple with the AllocatedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobResources) GetAllocatedHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedHosts) {
		return nil, false
	}
	return o.AllocatedHosts, true
}

// HasAllocatedHosts returns a boolean if a field has been set.
func (o *V0038JobResources) HasAllocatedHosts() bool {
	if o != nil && !IsNil(o.AllocatedHosts) {
		return true
	}

	return false
}

// SetAllocatedHosts gets a reference to the given int32 and assigns it to the AllocatedHosts field.
func (o *V0038JobResources) SetAllocatedHosts(v int32) {
	o.AllocatedHosts = &v
}

// GetAllocatedNodes returns the AllocatedNodes field value if set, zero value otherwise.
func (o *V0038JobResources) GetAllocatedNodes() []V0038NodeAllocation {
	if o == nil || IsNil(o.AllocatedNodes) {
		var ret []V0038NodeAllocation
		return ret
	}
	return o.AllocatedNodes
}

// GetAllocatedNodesOk returns a tuple with the AllocatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobResources) GetAllocatedNodesOk() ([]V0038NodeAllocation, bool) {
	if o == nil || IsNil(o.AllocatedNodes) {
		return nil, false
	}
	return o.AllocatedNodes, true
}

// HasAllocatedNodes returns a boolean if a field has been set.
func (o *V0038JobResources) HasAllocatedNodes() bool {
	if o != nil && !IsNil(o.AllocatedNodes) {
		return true
	}

	return false
}

// SetAllocatedNodes gets a reference to the given []V0038NodeAllocation and assigns it to the AllocatedNodes field.
func (o *V0038JobResources) SetAllocatedNodes(v []V0038NodeAllocation) {
	o.AllocatedNodes = v
}

func (o V0038JobResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038JobResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
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

type NullableV0038JobResources struct {
	value *V0038JobResources
	isSet bool
}

func (v NullableV0038JobResources) Get() *V0038JobResources {
	return v.value
}

func (v *NullableV0038JobResources) Set(val *V0038JobResources) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038JobResources) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038JobResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038JobResources(val *V0038JobResources) *NullableV0038JobResources {
	return &NullableV0038JobResources{value: val, isSet: true}
}

func (v NullableV0038JobResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038JobResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


