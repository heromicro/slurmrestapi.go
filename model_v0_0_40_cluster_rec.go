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

// checks if the V0040ClusterRec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ClusterRec{}

// V0040ClusterRec struct for V0040ClusterRec
type V0040ClusterRec struct {
	Controller *V0040ClusterRecController `json:"controller,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Name *string `json:"name,omitempty"`
	Nodes *string `json:"nodes,omitempty"`
	SelectPlugin *string `json:"select_plugin,omitempty"`
	Associations *V0040ClusterRecAssociations `json:"associations,omitempty"`
	RpcVersion *int32 `json:"rpc_version,omitempty"`
	Tres []V0040Tres `json:"tres,omitempty"`
}

// NewV0040ClusterRec instantiates a new V0040ClusterRec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ClusterRec() *V0040ClusterRec {
	this := V0040ClusterRec{}
	return &this
}

// NewV0040ClusterRecWithDefaults instantiates a new V0040ClusterRec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ClusterRecWithDefaults() *V0040ClusterRec {
	this := V0040ClusterRec{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetController() V0040ClusterRecController {
	if o == nil || IsNil(o.Controller) {
		var ret V0040ClusterRecController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetControllerOk() (*V0040ClusterRecController, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given V0040ClusterRecController and assigns it to the Controller field.
func (o *V0040ClusterRec) SetController(v V0040ClusterRecController) {
	o.Controller = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0040ClusterRec) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040ClusterRec) SetName(v string) {
	o.Name = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0040ClusterRec) SetNodes(v string) {
	o.Nodes = &v
}

// GetSelectPlugin returns the SelectPlugin field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetSelectPlugin() string {
	if o == nil || IsNil(o.SelectPlugin) {
		var ret string
		return ret
	}
	return *o.SelectPlugin
}

// GetSelectPluginOk returns a tuple with the SelectPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetSelectPluginOk() (*string, bool) {
	if o == nil || IsNil(o.SelectPlugin) {
		return nil, false
	}
	return o.SelectPlugin, true
}

// HasSelectPlugin returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasSelectPlugin() bool {
	if o != nil && !IsNil(o.SelectPlugin) {
		return true
	}

	return false
}

// SetSelectPlugin gets a reference to the given string and assigns it to the SelectPlugin field.
func (o *V0040ClusterRec) SetSelectPlugin(v string) {
	o.SelectPlugin = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetAssociations() V0040ClusterRecAssociations {
	if o == nil || IsNil(o.Associations) {
		var ret V0040ClusterRecAssociations
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetAssociationsOk() (*V0040ClusterRecAssociations, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given V0040ClusterRecAssociations and assigns it to the Associations field.
func (o *V0040ClusterRec) SetAssociations(v V0040ClusterRecAssociations) {
	o.Associations = &v
}

// GetRpcVersion returns the RpcVersion field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetRpcVersion() int32 {
	if o == nil || IsNil(o.RpcVersion) {
		var ret int32
		return ret
	}
	return *o.RpcVersion
}

// GetRpcVersionOk returns a tuple with the RpcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetRpcVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RpcVersion) {
		return nil, false
	}
	return o.RpcVersion, true
}

// HasRpcVersion returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasRpcVersion() bool {
	if o != nil && !IsNil(o.RpcVersion) {
		return true
	}

	return false
}

// SetRpcVersion gets a reference to the given int32 and assigns it to the RpcVersion field.
func (o *V0040ClusterRec) SetRpcVersion(v int32) {
	o.RpcVersion = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040ClusterRec) GetTres() []V0040Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0040Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRec) GetTresOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040ClusterRec) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0040Tres and assigns it to the Tres field.
func (o *V0040ClusterRec) SetTres(v []V0040Tres) {
	o.Tres = v
}

func (o V0040ClusterRec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ClusterRec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.SelectPlugin) {
		toSerialize["select_plugin"] = o.SelectPlugin
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.RpcVersion) {
		toSerialize["rpc_version"] = o.RpcVersion
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableV0040ClusterRec struct {
	value *V0040ClusterRec
	isSet bool
}

func (v NullableV0040ClusterRec) Get() *V0040ClusterRec {
	return v.value
}

func (v *NullableV0040ClusterRec) Set(val *V0040ClusterRec) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ClusterRec) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ClusterRec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ClusterRec(val *V0040ClusterRec) *NullableV0040ClusterRec {
	return &NullableV0040ClusterRec{value: val, isSet: true}
}

func (v NullableV0040ClusterRec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ClusterRec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


