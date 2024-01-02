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

// checks if the Dbv0038TresListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038TresListInner{}

// Dbv0038TresListInner struct for Dbv0038TresListInner
type Dbv0038TresListInner struct {
	// TRES type
	Type *string `json:"type,omitempty"`
	// TRES name (optional)
	Name *string `json:"name,omitempty"`
	// database id
	Id *int32 `json:"id,omitempty"`
	// count of TRES
	Count *int32 `json:"count,omitempty"`
}

// NewDbv0038TresListInner instantiates a new Dbv0038TresListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038TresListInner() *Dbv0038TresListInner {
	this := Dbv0038TresListInner{}
	return &this
}

// NewDbv0038TresListInnerWithDefaults instantiates a new Dbv0038TresListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038TresListInnerWithDefaults() *Dbv0038TresListInner {
	this := Dbv0038TresListInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Dbv0038TresListInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038TresListInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Dbv0038TresListInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Dbv0038TresListInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038TresListInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038TresListInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038TresListInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038TresListInner) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0038TresListInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038TresListInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0038TresListInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Dbv0038TresListInner) SetId(v int32) {
	o.Id = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0038TresListInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038TresListInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0038TresListInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0038TresListInner) SetCount(v int32) {
	o.Count = &v
}

func (o Dbv0038TresListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038TresListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableDbv0038TresListInner struct {
	value *Dbv0038TresListInner
	isSet bool
}

func (v NullableDbv0038TresListInner) Get() *Dbv0038TresListInner {
	return v.value
}

func (v *NullableDbv0038TresListInner) Set(val *Dbv0038TresListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038TresListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038TresListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038TresListInner(val *Dbv0038TresListInner) *NullableDbv0038TresListInner {
	return &NullableDbv0038TresListInner{value: val, isSet: true}
}

func (v NullableDbv0038TresListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038TresListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

