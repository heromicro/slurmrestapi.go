/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0039Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039Meta{}

// Dbv0039Meta struct for Dbv0039Meta
type Dbv0039Meta struct {
	Plugin *V0039MetaPlugin `json:"plugin,omitempty"`
	Slurm *V0039MetaSlurm `json:"Slurm,omitempty"`
}

// NewDbv0039Meta instantiates a new Dbv0039Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039Meta() *Dbv0039Meta {
	this := Dbv0039Meta{}
	return &this
}

// NewDbv0039MetaWithDefaults instantiates a new Dbv0039Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039MetaWithDefaults() *Dbv0039Meta {
	this := Dbv0039Meta{}
	return &this
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *Dbv0039Meta) GetPlugin() V0039MetaPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret V0039MetaPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Meta) GetPluginOk() (*V0039MetaPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *Dbv0039Meta) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given V0039MetaPlugin and assigns it to the Plugin field.
func (o *Dbv0039Meta) SetPlugin(v V0039MetaPlugin) {
	o.Plugin = &v
}

// GetSlurm returns the Slurm field value if set, zero value otherwise.
func (o *Dbv0039Meta) GetSlurm() V0039MetaSlurm {
	if o == nil || IsNil(o.Slurm) {
		var ret V0039MetaSlurm
		return ret
	}
	return *o.Slurm
}

// GetSlurmOk returns a tuple with the Slurm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Meta) GetSlurmOk() (*V0039MetaSlurm, bool) {
	if o == nil || IsNil(o.Slurm) {
		return nil, false
	}
	return o.Slurm, true
}

// HasSlurm returns a boolean if a field has been set.
func (o *Dbv0039Meta) HasSlurm() bool {
	if o != nil && !IsNil(o.Slurm) {
		return true
	}

	return false
}

// SetSlurm gets a reference to the given V0039MetaSlurm and assigns it to the Slurm field.
func (o *Dbv0039Meta) SetSlurm(v V0039MetaSlurm) {
	o.Slurm = &v
}

func (o Dbv0039Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Slurm) {
		toSerialize["Slurm"] = o.Slurm
	}
	return toSerialize, nil
}

type NullableDbv0039Meta struct {
	value *Dbv0039Meta
	isSet bool
}

func (v NullableDbv0039Meta) Get() *Dbv0039Meta {
	return v.value
}

func (v *NullableDbv0039Meta) Set(val *Dbv0039Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039Meta(val *Dbv0039Meta) *NullableDbv0039Meta {
	return &NullableDbv0039Meta{value: val, isSet: true}
}

func (v NullableDbv0039Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

