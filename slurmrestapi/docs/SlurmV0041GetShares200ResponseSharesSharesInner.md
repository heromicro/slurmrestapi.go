# SlurmV0041GetShares200ResponseSharesSharesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Association ID | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Name** | Pointer to **string** | Share name | [optional] 
**Parent** | Pointer to **string** | Parent name | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**SharesNormalized** | Pointer to [**SlurmV0041GetShares200ResponseSharesSharesInnerSharesNormalized**](SlurmV0041GetShares200ResponseSharesSharesInnerSharesNormalized.md) |  | [optional] 
**Shares** | Pointer to [**SlurmV0041GetShares200ResponseSharesSharesInnerShares**](SlurmV0041GetShares200ResponseSharesSharesInnerShares.md) |  | [optional] 
**Tres** | Pointer to [**SlurmV0041GetShares200ResponseSharesSharesInnerTres**](SlurmV0041GetShares200ResponseSharesSharesInnerTres.md) |  | [optional] 
**EffectiveUsage** | Pointer to **float64** | Effective, normalized usage | [optional] 
**UsageNormalized** | Pointer to [**SlurmV0041GetShares200ResponseSharesSharesInnerUsageNormalized**](SlurmV0041GetShares200ResponseSharesSharesInnerUsageNormalized.md) |  | [optional] 
**Usage** | Pointer to **int64** | Measure of tresbillableunits usage | [optional] 
**Fairshare** | Pointer to [**SlurmV0041GetShares200ResponseSharesSharesInnerFairshare**](SlurmV0041GetShares200ResponseSharesSharesInnerFairshare.md) |  | [optional] 
**Type** | Pointer to **[]string** | User or account association | [optional] 

## Methods

### NewSlurmV0041GetShares200ResponseSharesSharesInner

`func NewSlurmV0041GetShares200ResponseSharesSharesInner() *SlurmV0041GetShares200ResponseSharesSharesInner`

NewSlurmV0041GetShares200ResponseSharesSharesInner instantiates a new SlurmV0041GetShares200ResponseSharesSharesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseSharesSharesInnerWithDefaults

`func NewSlurmV0041GetShares200ResponseSharesSharesInnerWithDefaults() *SlurmV0041GetShares200ResponseSharesSharesInner`

NewSlurmV0041GetShares200ResponseSharesSharesInnerWithDefaults instantiates a new SlurmV0041GetShares200ResponseSharesSharesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCluster

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPartition

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSharesNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetSharesNormalized() SlurmV0041GetShares200ResponseSharesSharesInnerSharesNormalized`

GetSharesNormalized returns the SharesNormalized field if non-nil, zero value otherwise.

### GetSharesNormalizedOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetSharesNormalizedOk() (*SlurmV0041GetShares200ResponseSharesSharesInnerSharesNormalized, bool)`

GetSharesNormalizedOk returns a tuple with the SharesNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetSharesNormalized(v SlurmV0041GetShares200ResponseSharesSharesInnerSharesNormalized)`

SetSharesNormalized sets SharesNormalized field to given value.

### HasSharesNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasSharesNormalized() bool`

HasSharesNormalized returns a boolean if a field has been set.

### GetShares

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetShares() SlurmV0041GetShares200ResponseSharesSharesInnerShares`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetSharesOk() (*SlurmV0041GetShares200ResponseSharesSharesInnerShares, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetShares(v SlurmV0041GetShares200ResponseSharesSharesInnerShares)`

SetShares sets Shares field to given value.

### HasShares

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTres

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetTres() SlurmV0041GetShares200ResponseSharesSharesInnerTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetTresOk() (*SlurmV0041GetShares200ResponseSharesSharesInnerTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetTres(v SlurmV0041GetShares200ResponseSharesSharesInnerTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetEffectiveUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetEffectiveUsage() float64`

GetEffectiveUsage returns the EffectiveUsage field if non-nil, zero value otherwise.

### GetEffectiveUsageOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetEffectiveUsageOk() (*float64, bool)`

GetEffectiveUsageOk returns a tuple with the EffectiveUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetEffectiveUsage(v float64)`

SetEffectiveUsage sets EffectiveUsage field to given value.

### HasEffectiveUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasEffectiveUsage() bool`

HasEffectiveUsage returns a boolean if a field has been set.

### GetUsageNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetUsageNormalized() SlurmV0041GetShares200ResponseSharesSharesInnerUsageNormalized`

GetUsageNormalized returns the UsageNormalized field if non-nil, zero value otherwise.

### GetUsageNormalizedOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetUsageNormalizedOk() (*SlurmV0041GetShares200ResponseSharesSharesInnerUsageNormalized, bool)`

GetUsageNormalizedOk returns a tuple with the UsageNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetUsageNormalized(v SlurmV0041GetShares200ResponseSharesSharesInnerUsageNormalized)`

SetUsageNormalized sets UsageNormalized field to given value.

### HasUsageNormalized

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasUsageNormalized() bool`

HasUsageNormalized returns a boolean if a field has been set.

### GetUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetFairshare

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetFairshare() SlurmV0041GetShares200ResponseSharesSharesInnerFairshare`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetFairshareOk() (*SlurmV0041GetShares200ResponseSharesSharesInnerFairshare, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetFairshare(v SlurmV0041GetShares200ResponseSharesSharesInnerFairshare)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.

### GetType

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *SlurmV0041GetShares200ResponseSharesSharesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


