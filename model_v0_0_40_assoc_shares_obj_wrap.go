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

// checks if the V0040AssocSharesObjWrap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocSharesObjWrap{}

// V0040AssocSharesObjWrap struct for V0040AssocSharesObjWrap
type V0040AssocSharesObjWrap struct {
	// assocation id
	Id *int32 `json:"id,omitempty"`
	// cluster name
	Cluster *string `json:"cluster,omitempty"`
	// share name
	Name *string `json:"name,omitempty"`
	// parent name
	Parent *string `json:"parent,omitempty"`
	// partition name
	Partition *string `json:"partition,omitempty"`
	SharesNormalized *V0040Float64NoVal `json:"shares_normalized,omitempty"`
	Shares *V0040Uint32NoVal `json:"shares,omitempty"`
	Tres *V0040AssocSharesObjWrapTres `json:"tres,omitempty"`
	// effective, normalized usage
	EffectiveUsage *float64 `json:"effective_usage,omitempty"`
	UsageNormalized *V0040Float64NoVal `json:"usage_normalized,omitempty"`
	// measure of tresbillableunits usage
	Usage *int64 `json:"usage,omitempty"`
	Fairshare *V0040AssocSharesObjWrapFairshare `json:"fairshare,omitempty"`
	// user or account association
	Type []string `json:"type,omitempty"`
}

// NewV0040AssocSharesObjWrap instantiates a new V0040AssocSharesObjWrap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocSharesObjWrap() *V0040AssocSharesObjWrap {
	this := V0040AssocSharesObjWrap{}
	return &this
}

// NewV0040AssocSharesObjWrapWithDefaults instantiates a new V0040AssocSharesObjWrap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocSharesObjWrapWithDefaults() *V0040AssocSharesObjWrap {
	this := V0040AssocSharesObjWrap{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0040AssocSharesObjWrap) SetId(v int32) {
	o.Id = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *V0040AssocSharesObjWrap) SetCluster(v string) {
	o.Cluster = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040AssocSharesObjWrap) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *V0040AssocSharesObjWrap) SetParent(v string) {
	o.Parent = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0040AssocSharesObjWrap) SetPartition(v string) {
	o.Partition = &v
}

// GetSharesNormalized returns the SharesNormalized field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetSharesNormalized() V0040Float64NoVal {
	if o == nil || IsNil(o.SharesNormalized) {
		var ret V0040Float64NoVal
		return ret
	}
	return *o.SharesNormalized
}

// GetSharesNormalizedOk returns a tuple with the SharesNormalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetSharesNormalizedOk() (*V0040Float64NoVal, bool) {
	if o == nil || IsNil(o.SharesNormalized) {
		return nil, false
	}
	return o.SharesNormalized, true
}

// HasSharesNormalized returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasSharesNormalized() bool {
	if o != nil && !IsNil(o.SharesNormalized) {
		return true
	}

	return false
}

// SetSharesNormalized gets a reference to the given V0040Float64NoVal and assigns it to the SharesNormalized field.
func (o *V0040AssocSharesObjWrap) SetSharesNormalized(v V0040Float64NoVal) {
	o.SharesNormalized = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetShares() V0040Uint32NoVal {
	if o == nil || IsNil(o.Shares) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetSharesOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given V0040Uint32NoVal and assigns it to the Shares field.
func (o *V0040AssocSharesObjWrap) SetShares(v V0040Uint32NoVal) {
	o.Shares = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetTres() V0040AssocSharesObjWrapTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0040AssocSharesObjWrapTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetTresOk() (*V0040AssocSharesObjWrapTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0040AssocSharesObjWrapTres and assigns it to the Tres field.
func (o *V0040AssocSharesObjWrap) SetTres(v V0040AssocSharesObjWrapTres) {
	o.Tres = &v
}

// GetEffectiveUsage returns the EffectiveUsage field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetEffectiveUsage() float64 {
	if o == nil || IsNil(o.EffectiveUsage) {
		var ret float64
		return ret
	}
	return *o.EffectiveUsage
}

// GetEffectiveUsageOk returns a tuple with the EffectiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetEffectiveUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.EffectiveUsage) {
		return nil, false
	}
	return o.EffectiveUsage, true
}

// HasEffectiveUsage returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasEffectiveUsage() bool {
	if o != nil && !IsNil(o.EffectiveUsage) {
		return true
	}

	return false
}

// SetEffectiveUsage gets a reference to the given float64 and assigns it to the EffectiveUsage field.
func (o *V0040AssocSharesObjWrap) SetEffectiveUsage(v float64) {
	o.EffectiveUsage = &v
}

// GetUsageNormalized returns the UsageNormalized field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetUsageNormalized() V0040Float64NoVal {
	if o == nil || IsNil(o.UsageNormalized) {
		var ret V0040Float64NoVal
		return ret
	}
	return *o.UsageNormalized
}

// GetUsageNormalizedOk returns a tuple with the UsageNormalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetUsageNormalizedOk() (*V0040Float64NoVal, bool) {
	if o == nil || IsNil(o.UsageNormalized) {
		return nil, false
	}
	return o.UsageNormalized, true
}

// HasUsageNormalized returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasUsageNormalized() bool {
	if o != nil && !IsNil(o.UsageNormalized) {
		return true
	}

	return false
}

// SetUsageNormalized gets a reference to the given V0040Float64NoVal and assigns it to the UsageNormalized field.
func (o *V0040AssocSharesObjWrap) SetUsageNormalized(v V0040Float64NoVal) {
	o.UsageNormalized = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetUsage() int64 {
	if o == nil || IsNil(o.Usage) {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *V0040AssocSharesObjWrap) SetUsage(v int64) {
	o.Usage = &v
}

// GetFairshare returns the Fairshare field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetFairshare() V0040AssocSharesObjWrapFairshare {
	if o == nil || IsNil(o.Fairshare) {
		var ret V0040AssocSharesObjWrapFairshare
		return ret
	}
	return *o.Fairshare
}

// GetFairshareOk returns a tuple with the Fairshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetFairshareOk() (*V0040AssocSharesObjWrapFairshare, bool) {
	if o == nil || IsNil(o.Fairshare) {
		return nil, false
	}
	return o.Fairshare, true
}

// HasFairshare returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasFairshare() bool {
	if o != nil && !IsNil(o.Fairshare) {
		return true
	}

	return false
}

// SetFairshare gets a reference to the given V0040AssocSharesObjWrapFairshare and assigns it to the Fairshare field.
func (o *V0040AssocSharesObjWrap) SetFairshare(v V0040AssocSharesObjWrapFairshare) {
	o.Fairshare = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrap) GetType() []string {
	if o == nil || IsNil(o.Type) {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrap) GetTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrap) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *V0040AssocSharesObjWrap) SetType(v []string) {
	o.Type = v
}

func (o V0040AssocSharesObjWrap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocSharesObjWrap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.SharesNormalized) {
		toSerialize["shares_normalized"] = o.SharesNormalized
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.EffectiveUsage) {
		toSerialize["effective_usage"] = o.EffectiveUsage
	}
	if !IsNil(o.UsageNormalized) {
		toSerialize["usage_normalized"] = o.UsageNormalized
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Fairshare) {
		toSerialize["fairshare"] = o.Fairshare
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV0040AssocSharesObjWrap struct {
	value *V0040AssocSharesObjWrap
	isSet bool
}

func (v NullableV0040AssocSharesObjWrap) Get() *V0040AssocSharesObjWrap {
	return v.value
}

func (v *NullableV0040AssocSharesObjWrap) Set(val *V0040AssocSharesObjWrap) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocSharesObjWrap) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocSharesObjWrap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocSharesObjWrap(val *V0040AssocSharesObjWrap) *NullableV0040AssocSharesObjWrap {
	return &NullableV0040AssocSharesObjWrap{value: val, isSet: true}
}

func (v NullableV0040AssocSharesObjWrap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocSharesObjWrap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


