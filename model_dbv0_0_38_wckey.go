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

// checks if the Dbv0038Wckey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Wckey{}

// Dbv0038Wckey struct for Dbv0038Wckey
type Dbv0038Wckey struct {
	// Cluster name
	Cluster *string `json:"cluster,omitempty"`
	// wckey database unique id
	Id *int32 `json:"id,omitempty"`
	// wckey name
	Name *string `json:"name,omitempty"`
	// wckey user
	User *string `json:"user,omitempty"`
	// List of properties of wckey
	Flags []string `json:"flags,omitempty"`
	// List of accounting records
	Accounting []Dbv0038Accounting `json:"accounting,omitempty"`
}

// NewDbv0038Wckey instantiates a new Dbv0038Wckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Wckey() *Dbv0038Wckey {
	this := Dbv0038Wckey{}
	return &this
}

// NewDbv0038WckeyWithDefaults instantiates a new Dbv0038Wckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038WckeyWithDefaults() *Dbv0038Wckey {
	this := Dbv0038Wckey{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Dbv0038Wckey) SetCluster(v string) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Dbv0038Wckey) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038Wckey) SetName(v string) {
	o.Name = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Dbv0038Wckey) SetUser(v string) {
	o.User = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038Wckey) SetFlags(v []string) {
	o.Flags = v
}

// GetAccounting returns the Accounting field value if set, zero value otherwise.
func (o *Dbv0038Wckey) GetAccounting() []Dbv0038Accounting {
	if o == nil || IsNil(o.Accounting) {
		var ret []Dbv0038Accounting
		return ret
	}
	return o.Accounting
}

// GetAccountingOk returns a tuple with the Accounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Wckey) GetAccountingOk() ([]Dbv0038Accounting, bool) {
	if o == nil || IsNil(o.Accounting) {
		return nil, false
	}
	return o.Accounting, true
}

// HasAccounting returns a boolean if a field has been set.
func (o *Dbv0038Wckey) HasAccounting() bool {
	if o != nil && !IsNil(o.Accounting) {
		return true
	}

	return false
}

// SetAccounting gets a reference to the given []Dbv0038Accounting and assigns it to the Accounting field.
func (o *Dbv0038Wckey) SetAccounting(v []Dbv0038Accounting) {
	o.Accounting = v
}

func (o Dbv0038Wckey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Wckey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Accounting) {
		toSerialize["accounting"] = o.Accounting
	}
	return toSerialize, nil
}

type NullableDbv0038Wckey struct {
	value *Dbv0038Wckey
	isSet bool
}

func (v NullableDbv0038Wckey) Get() *Dbv0038Wckey {
	return v.value
}

func (v *NullableDbv0038Wckey) Set(val *Dbv0038Wckey) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Wckey) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Wckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Wckey(val *Dbv0038Wckey) *NullableDbv0038Wckey {
	return &NullableDbv0038Wckey{value: val, isSet: true}
}

func (v NullableDbv0038Wckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Wckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


