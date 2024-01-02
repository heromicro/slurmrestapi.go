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

// checks if the Dbv0038Association type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Association{}

// Dbv0038Association Association description
type Dbv0038Association struct {
	// Assigned account
	Account *string `json:"account,omitempty"`
	// Assigned cluster
	Cluster *string `json:"cluster,omitempty"`
	Default *Dbv0038AssociationDefault `json:"default,omitempty"`
	// List of properties of association
	Flags []string `json:"flags,omitempty"`
	Max *Dbv0038AssociationMax `json:"max,omitempty"`
	Min *Dbv0038AssociationMin `json:"min,omitempty"`
	// Parent account name
	ParentAccount *string `json:"parent_account,omitempty"`
	// Assigned partition
	Partition *string `json:"partition,omitempty"`
	// Assigned priority
	Priority *int32 `json:"priority,omitempty"`
	// Raw fairshare shares
	SharesRaw *int32 `json:"shares_raw,omitempty"`
	Usage *Dbv0038AssociationUsage `json:"usage,omitempty"`
	// Assigned user
	User *string `json:"user,omitempty"`
	// Assigned QOS
	QOS []string `json:"QOS,omitempty"`
}

// NewDbv0038Association instantiates a new Dbv0038Association object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Association() *Dbv0038Association {
	this := Dbv0038Association{}
	return &this
}

// NewDbv0038AssociationWithDefaults instantiates a new Dbv0038Association object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AssociationWithDefaults() *Dbv0038Association {
	this := Dbv0038Association{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0038Association) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0038Association) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Dbv0038Association) SetAccount(v string) {
	o.Account = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Dbv0038Association) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Dbv0038Association) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Dbv0038Association) SetCluster(v string) {
	o.Cluster = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Dbv0038Association) GetDefault() Dbv0038AssociationDefault {
	if o == nil || IsNil(o.Default) {
		var ret Dbv0038AssociationDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetDefaultOk() (*Dbv0038AssociationDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Dbv0038Association) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given Dbv0038AssociationDefault and assigns it to the Default field.
func (o *Dbv0038Association) SetDefault(v Dbv0038AssociationDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038Association) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038Association) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038Association) SetFlags(v []string) {
	o.Flags = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0038Association) GetMax() Dbv0038AssociationMax {
	if o == nil || IsNil(o.Max) {
		var ret Dbv0038AssociationMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetMaxOk() (*Dbv0038AssociationMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0038Association) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given Dbv0038AssociationMax and assigns it to the Max field.
func (o *Dbv0038Association) SetMax(v Dbv0038AssociationMax) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Dbv0038Association) GetMin() Dbv0038AssociationMin {
	if o == nil || IsNil(o.Min) {
		var ret Dbv0038AssociationMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetMinOk() (*Dbv0038AssociationMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Dbv0038Association) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given Dbv0038AssociationMin and assigns it to the Min field.
func (o *Dbv0038Association) SetMin(v Dbv0038AssociationMin) {
	o.Min = &v
}

// GetParentAccount returns the ParentAccount field value if set, zero value otherwise.
func (o *Dbv0038Association) GetParentAccount() string {
	if o == nil || IsNil(o.ParentAccount) {
		var ret string
		return ret
	}
	return *o.ParentAccount
}

// GetParentAccountOk returns a tuple with the ParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetParentAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ParentAccount) {
		return nil, false
	}
	return o.ParentAccount, true
}

// HasParentAccount returns a boolean if a field has been set.
func (o *Dbv0038Association) HasParentAccount() bool {
	if o != nil && !IsNil(o.ParentAccount) {
		return true
	}

	return false
}

// SetParentAccount gets a reference to the given string and assigns it to the ParentAccount field.
func (o *Dbv0038Association) SetParentAccount(v string) {
	o.ParentAccount = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *Dbv0038Association) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *Dbv0038Association) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *Dbv0038Association) SetPartition(v string) {
	o.Partition = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Dbv0038Association) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Dbv0038Association) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *Dbv0038Association) SetPriority(v int32) {
	o.Priority = &v
}

// GetSharesRaw returns the SharesRaw field value if set, zero value otherwise.
func (o *Dbv0038Association) GetSharesRaw() int32 {
	if o == nil || IsNil(o.SharesRaw) {
		var ret int32
		return ret
	}
	return *o.SharesRaw
}

// GetSharesRawOk returns a tuple with the SharesRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetSharesRawOk() (*int32, bool) {
	if o == nil || IsNil(o.SharesRaw) {
		return nil, false
	}
	return o.SharesRaw, true
}

// HasSharesRaw returns a boolean if a field has been set.
func (o *Dbv0038Association) HasSharesRaw() bool {
	if o != nil && !IsNil(o.SharesRaw) {
		return true
	}

	return false
}

// SetSharesRaw gets a reference to the given int32 and assigns it to the SharesRaw field.
func (o *Dbv0038Association) SetSharesRaw(v int32) {
	o.SharesRaw = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *Dbv0038Association) GetUsage() Dbv0038AssociationUsage {
	if o == nil || IsNil(o.Usage) {
		var ret Dbv0038AssociationUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetUsageOk() (*Dbv0038AssociationUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *Dbv0038Association) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given Dbv0038AssociationUsage and assigns it to the Usage field.
func (o *Dbv0038Association) SetUsage(v Dbv0038AssociationUsage) {
	o.Usage = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0038Association) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0038Association) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Dbv0038Association) SetUser(v string) {
	o.User = &v
}

// GetQOS returns the QOS field value if set, zero value otherwise.
func (o *Dbv0038Association) GetQOS() []string {
	if o == nil || IsNil(o.QOS) {
		var ret []string
		return ret
	}
	return o.QOS
}

// GetQOSOk returns a tuple with the QOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Association) GetQOSOk() ([]string, bool) {
	if o == nil || IsNil(o.QOS) {
		return nil, false
	}
	return o.QOS, true
}

// HasQOS returns a boolean if a field has been set.
func (o *Dbv0038Association) HasQOS() bool {
	if o != nil && !IsNil(o.QOS) {
		return true
	}

	return false
}

// SetQOS gets a reference to the given []string and assigns it to the QOS field.
func (o *Dbv0038Association) SetQOS(v []string) {
	o.QOS = v
}

func (o Dbv0038Association) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Association) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.ParentAccount) {
		toSerialize["parent_account"] = o.ParentAccount
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.SharesRaw) {
		toSerialize["shares_raw"] = o.SharesRaw
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.QOS) {
		toSerialize["QOS"] = o.QOS
	}
	return toSerialize, nil
}

type NullableDbv0038Association struct {
	value *Dbv0038Association
	isSet bool
}

func (v NullableDbv0038Association) Get() *Dbv0038Association {
	return v.value
}

func (v *NullableDbv0038Association) Set(val *Dbv0038Association) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Association) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Association) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Association(val *Dbv0038Association) *NullableDbv0038Association {
	return &NullableDbv0038Association{value: val, isSet: true}
}

func (v NullableDbv0038Association) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Association) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


