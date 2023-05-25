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

// checks if the V0039ReservationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039ReservationInfo{}

// V0039ReservationInfo struct for V0039ReservationInfo
type V0039ReservationInfo struct {
	Accounts *string `json:"accounts,omitempty"`
	BurstBuffer *string `json:"burst_buffer,omitempty"`
	CoreCount *int32 `json:"core_count,omitempty"`
	CoreSpecializations []V0039ReservationCoreSpec `json:"core_specializations,omitempty"`
	EndTime *int64 `json:"end_time,omitempty"`
	Features *string `json:"features,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Groups *string `json:"groups,omitempty"`
	Licenses *string `json:"licenses,omitempty"`
	MaxStartDelay *int32 `json:"max_start_delay,omitempty"`
	Name *string `json:"name,omitempty"`
	NodeCount *int32 `json:"node_count,omitempty"`
	NodeList *string `json:"node_list,omitempty"`
	Partition *string `json:"partition,omitempty"`
	PurgeCompleted *V0039ReservationInfoPurgeCompleted `json:"purge_completed,omitempty"`
	StartTime *int64 `json:"start_time,omitempty"`
	Watts *V0039Uint32NoVal `json:"watts,omitempty"`
	Tres *string `json:"tres,omitempty"`
	Users *string `json:"users,omitempty"`
}

// NewV0039ReservationInfo instantiates a new V0039ReservationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039ReservationInfo() *V0039ReservationInfo {
	this := V0039ReservationInfo{}
	return &this
}

// NewV0039ReservationInfoWithDefaults instantiates a new V0039ReservationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039ReservationInfoWithDefaults() *V0039ReservationInfo {
	this := V0039ReservationInfo{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetAccounts() string {
	if o == nil || IsNil(o.Accounts) {
		var ret string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given string and assigns it to the Accounts field.
func (o *V0039ReservationInfo) SetAccounts(v string) {
	o.Accounts = &v
}

// GetBurstBuffer returns the BurstBuffer field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetBurstBuffer() string {
	if o == nil || IsNil(o.BurstBuffer) {
		var ret string
		return ret
	}
	return *o.BurstBuffer
}

// GetBurstBufferOk returns a tuple with the BurstBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetBurstBufferOk() (*string, bool) {
	if o == nil || IsNil(o.BurstBuffer) {
		return nil, false
	}
	return o.BurstBuffer, true
}

// HasBurstBuffer returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasBurstBuffer() bool {
	if o != nil && !IsNil(o.BurstBuffer) {
		return true
	}

	return false
}

// SetBurstBuffer gets a reference to the given string and assigns it to the BurstBuffer field.
func (o *V0039ReservationInfo) SetBurstBuffer(v string) {
	o.BurstBuffer = &v
}

// GetCoreCount returns the CoreCount field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetCoreCount() int32 {
	if o == nil || IsNil(o.CoreCount) {
		var ret int32
		return ret
	}
	return *o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetCoreCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CoreCount) {
		return nil, false
	}
	return o.CoreCount, true
}

// HasCoreCount returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasCoreCount() bool {
	if o != nil && !IsNil(o.CoreCount) {
		return true
	}

	return false
}

// SetCoreCount gets a reference to the given int32 and assigns it to the CoreCount field.
func (o *V0039ReservationInfo) SetCoreCount(v int32) {
	o.CoreCount = &v
}

// GetCoreSpecializations returns the CoreSpecializations field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetCoreSpecializations() []V0039ReservationCoreSpec {
	if o == nil || IsNil(o.CoreSpecializations) {
		var ret []V0039ReservationCoreSpec
		return ret
	}
	return o.CoreSpecializations
}

// GetCoreSpecializationsOk returns a tuple with the CoreSpecializations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetCoreSpecializationsOk() ([]V0039ReservationCoreSpec, bool) {
	if o == nil || IsNil(o.CoreSpecializations) {
		return nil, false
	}
	return o.CoreSpecializations, true
}

// HasCoreSpecializations returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasCoreSpecializations() bool {
	if o != nil && !IsNil(o.CoreSpecializations) {
		return true
	}

	return false
}

// SetCoreSpecializations gets a reference to the given []V0039ReservationCoreSpec and assigns it to the CoreSpecializations field.
func (o *V0039ReservationInfo) SetCoreSpecializations(v []V0039ReservationCoreSpec) {
	o.CoreSpecializations = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *V0039ReservationInfo) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *V0039ReservationInfo) SetFeatures(v string) {
	o.Features = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039ReservationInfo) SetFlags(v []string) {
	o.Flags = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetGroups() string {
	if o == nil || IsNil(o.Groups) {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *V0039ReservationInfo) SetGroups(v string) {
	o.Groups = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetLicenses() string {
	if o == nil || IsNil(o.Licenses) {
		var ret string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetLicensesOk() (*string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given string and assigns it to the Licenses field.
func (o *V0039ReservationInfo) SetLicenses(v string) {
	o.Licenses = &v
}

// GetMaxStartDelay returns the MaxStartDelay field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetMaxStartDelay() int32 {
	if o == nil || IsNil(o.MaxStartDelay) {
		var ret int32
		return ret
	}
	return *o.MaxStartDelay
}

// GetMaxStartDelayOk returns a tuple with the MaxStartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetMaxStartDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxStartDelay) {
		return nil, false
	}
	return o.MaxStartDelay, true
}

// HasMaxStartDelay returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasMaxStartDelay() bool {
	if o != nil && !IsNil(o.MaxStartDelay) {
		return true
	}

	return false
}

// SetMaxStartDelay gets a reference to the given int32 and assigns it to the MaxStartDelay field.
func (o *V0039ReservationInfo) SetMaxStartDelay(v int32) {
	o.MaxStartDelay = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039ReservationInfo) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *V0039ReservationInfo) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetNodeList returns the NodeList field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetNodeList() string {
	if o == nil || IsNil(o.NodeList) {
		var ret string
		return ret
	}
	return *o.NodeList
}

// GetNodeListOk returns a tuple with the NodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetNodeListOk() (*string, bool) {
	if o == nil || IsNil(o.NodeList) {
		return nil, false
	}
	return o.NodeList, true
}

// HasNodeList returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasNodeList() bool {
	if o != nil && !IsNil(o.NodeList) {
		return true
	}

	return false
}

// SetNodeList gets a reference to the given string and assigns it to the NodeList field.
func (o *V0039ReservationInfo) SetNodeList(v string) {
	o.NodeList = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0039ReservationInfo) SetPartition(v string) {
	o.Partition = &v
}

// GetPurgeCompleted returns the PurgeCompleted field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetPurgeCompleted() V0039ReservationInfoPurgeCompleted {
	if o == nil || IsNil(o.PurgeCompleted) {
		var ret V0039ReservationInfoPurgeCompleted
		return ret
	}
	return *o.PurgeCompleted
}

// GetPurgeCompletedOk returns a tuple with the PurgeCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetPurgeCompletedOk() (*V0039ReservationInfoPurgeCompleted, bool) {
	if o == nil || IsNil(o.PurgeCompleted) {
		return nil, false
	}
	return o.PurgeCompleted, true
}

// HasPurgeCompleted returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasPurgeCompleted() bool {
	if o != nil && !IsNil(o.PurgeCompleted) {
		return true
	}

	return false
}

// SetPurgeCompleted gets a reference to the given V0039ReservationInfoPurgeCompleted and assigns it to the PurgeCompleted field.
func (o *V0039ReservationInfo) SetPurgeCompleted(v V0039ReservationInfoPurgeCompleted) {
	o.PurgeCompleted = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *V0039ReservationInfo) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetWatts returns the Watts field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetWatts() V0039Uint32NoVal {
	if o == nil || IsNil(o.Watts) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Watts
}

// GetWattsOk returns a tuple with the Watts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetWattsOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Watts) {
		return nil, false
	}
	return o.Watts, true
}

// HasWatts returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasWatts() bool {
	if o != nil && !IsNil(o.Watts) {
		return true
	}

	return false
}

// SetWatts gets a reference to the given V0039Uint32NoVal and assigns it to the Watts field.
func (o *V0039ReservationInfo) SetWatts(v V0039Uint32NoVal) {
	o.Watts = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetTres() string {
	if o == nil || IsNil(o.Tres) {
		var ret string
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetTresOk() (*string, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given string and assigns it to the Tres field.
func (o *V0039ReservationInfo) SetTres(v string) {
	o.Tres = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V0039ReservationInfo) GetUsers() string {
	if o == nil || IsNil(o.Users) {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ReservationInfo) GetUsersOk() (*string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V0039ReservationInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *V0039ReservationInfo) SetUsers(v string) {
	o.Users = &v
}

func (o V0039ReservationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039ReservationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.BurstBuffer) {
		toSerialize["burst_buffer"] = o.BurstBuffer
	}
	if !IsNil(o.CoreCount) {
		toSerialize["core_count"] = o.CoreCount
	}
	if !IsNil(o.CoreSpecializations) {
		toSerialize["core_specializations"] = o.CoreSpecializations
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.MaxStartDelay) {
		toSerialize["max_start_delay"] = o.MaxStartDelay
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeCount) {
		toSerialize["node_count"] = o.NodeCount
	}
	if !IsNil(o.NodeList) {
		toSerialize["node_list"] = o.NodeList
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.PurgeCompleted) {
		toSerialize["purge_completed"] = o.PurgeCompleted
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Watts) {
		toSerialize["watts"] = o.Watts
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableV0039ReservationInfo struct {
	value *V0039ReservationInfo
	isSet bool
}

func (v NullableV0039ReservationInfo) Get() *V0039ReservationInfo {
	return v.value
}

func (v *NullableV0039ReservationInfo) Set(val *V0039ReservationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039ReservationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039ReservationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039ReservationInfo(val *V0039ReservationInfo) *NullableV0039ReservationInfo {
	return &NullableV0039ReservationInfo{value: val, isSet: true}
}

func (v NullableV0039ReservationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039ReservationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


