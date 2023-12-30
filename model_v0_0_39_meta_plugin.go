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

// checks if the V0039MetaPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039MetaPlugin{}

// V0039MetaPlugin struct for V0039MetaPlugin
type V0039MetaPlugin struct {
	// 
	Type *string `json:"type,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
}

// NewV0039MetaPlugin instantiates a new V0039MetaPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039MetaPlugin() *V0039MetaPlugin {
	this := V0039MetaPlugin{}
	return &this
}

// NewV0039MetaPluginWithDefaults instantiates a new V0039MetaPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039MetaPluginWithDefaults() *V0039MetaPlugin {
	this := V0039MetaPlugin{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0039MetaPlugin) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039MetaPlugin) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0039MetaPlugin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V0039MetaPlugin) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039MetaPlugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039MetaPlugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039MetaPlugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039MetaPlugin) SetName(v string) {
	o.Name = &v
}

func (o V0039MetaPlugin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039MetaPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0039MetaPlugin struct {
	value *V0039MetaPlugin
	isSet bool
}

func (v NullableV0039MetaPlugin) Get() *V0039MetaPlugin {
	return v.value
}

func (v *NullableV0039MetaPlugin) Set(val *V0039MetaPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039MetaPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039MetaPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039MetaPlugin(val *V0039MetaPlugin) *NullableV0039MetaPlugin {
	return &NullableV0039MetaPlugin{value: val, isSet: true}
}

func (v NullableV0039MetaPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039MetaPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


