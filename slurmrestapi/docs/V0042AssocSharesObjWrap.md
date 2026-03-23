# V0042AssocSharesObjWrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Association ID | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Name** | Pointer to **string** | Share name | [optional] 
**Parent** | Pointer to **string** | Parent name | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**SharesNormalized** | Pointer to [**V0042Float64NoValStruct**](V0042Float64NoValStruct.md) |  | [optional] 
**Shares** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Tres** | Pointer to [**V0042AssocSharesObjWrapTres**](V0042AssocSharesObjWrapTres.md) |  | [optional] 
**EffectiveUsage** | Pointer to [**V0042Float64NoValStruct**](V0042Float64NoValStruct.md) |  | [optional] 
**UsageNormalized** | Pointer to [**V0042Float64NoValStruct**](V0042Float64NoValStruct.md) |  | [optional] 
**Usage** | Pointer to **int64** | Measure of tresbillableunits usage | [optional] 
**Fairshare** | Pointer to [**V0042AssocSharesObjWrapFairshare**](V0042AssocSharesObjWrapFairshare.md) |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0042AssocSharesObjWrap

`func NewV0042AssocSharesObjWrap() *V0042AssocSharesObjWrap`

NewV0042AssocSharesObjWrap instantiates a new V0042AssocSharesObjWrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042AssocSharesObjWrapWithDefaults

`func NewV0042AssocSharesObjWrapWithDefaults() *V0042AssocSharesObjWrap`

NewV0042AssocSharesObjWrapWithDefaults instantiates a new V0042AssocSharesObjWrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0042AssocSharesObjWrap) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0042AssocSharesObjWrap) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0042AssocSharesObjWrap) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0042AssocSharesObjWrap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCluster

`func (o *V0042AssocSharesObjWrap) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0042AssocSharesObjWrap) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0042AssocSharesObjWrap) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0042AssocSharesObjWrap) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *V0042AssocSharesObjWrap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042AssocSharesObjWrap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042AssocSharesObjWrap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0042AssocSharesObjWrap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *V0042AssocSharesObjWrap) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0042AssocSharesObjWrap) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0042AssocSharesObjWrap) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0042AssocSharesObjWrap) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPartition

`func (o *V0042AssocSharesObjWrap) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0042AssocSharesObjWrap) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0042AssocSharesObjWrap) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0042AssocSharesObjWrap) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSharesNormalized

`func (o *V0042AssocSharesObjWrap) GetSharesNormalized() V0042Float64NoValStruct`

GetSharesNormalized returns the SharesNormalized field if non-nil, zero value otherwise.

### GetSharesNormalizedOk

`func (o *V0042AssocSharesObjWrap) GetSharesNormalizedOk() (*V0042Float64NoValStruct, bool)`

GetSharesNormalizedOk returns a tuple with the SharesNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesNormalized

`func (o *V0042AssocSharesObjWrap) SetSharesNormalized(v V0042Float64NoValStruct)`

SetSharesNormalized sets SharesNormalized field to given value.

### HasSharesNormalized

`func (o *V0042AssocSharesObjWrap) HasSharesNormalized() bool`

HasSharesNormalized returns a boolean if a field has been set.

### GetShares

`func (o *V0042AssocSharesObjWrap) GetShares() V0042Uint32NoValStruct`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0042AssocSharesObjWrap) GetSharesOk() (*V0042Uint32NoValStruct, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0042AssocSharesObjWrap) SetShares(v V0042Uint32NoValStruct)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0042AssocSharesObjWrap) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTres

`func (o *V0042AssocSharesObjWrap) GetTres() V0042AssocSharesObjWrapTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0042AssocSharesObjWrap) GetTresOk() (*V0042AssocSharesObjWrapTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0042AssocSharesObjWrap) SetTres(v V0042AssocSharesObjWrapTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0042AssocSharesObjWrap) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetEffectiveUsage

`func (o *V0042AssocSharesObjWrap) GetEffectiveUsage() V0042Float64NoValStruct`

GetEffectiveUsage returns the EffectiveUsage field if non-nil, zero value otherwise.

### GetEffectiveUsageOk

`func (o *V0042AssocSharesObjWrap) GetEffectiveUsageOk() (*V0042Float64NoValStruct, bool)`

GetEffectiveUsageOk returns a tuple with the EffectiveUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUsage

`func (o *V0042AssocSharesObjWrap) SetEffectiveUsage(v V0042Float64NoValStruct)`

SetEffectiveUsage sets EffectiveUsage field to given value.

### HasEffectiveUsage

`func (o *V0042AssocSharesObjWrap) HasEffectiveUsage() bool`

HasEffectiveUsage returns a boolean if a field has been set.

### GetUsageNormalized

`func (o *V0042AssocSharesObjWrap) GetUsageNormalized() V0042Float64NoValStruct`

GetUsageNormalized returns the UsageNormalized field if non-nil, zero value otherwise.

### GetUsageNormalizedOk

`func (o *V0042AssocSharesObjWrap) GetUsageNormalizedOk() (*V0042Float64NoValStruct, bool)`

GetUsageNormalizedOk returns a tuple with the UsageNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNormalized

`func (o *V0042AssocSharesObjWrap) SetUsageNormalized(v V0042Float64NoValStruct)`

SetUsageNormalized sets UsageNormalized field to given value.

### HasUsageNormalized

`func (o *V0042AssocSharesObjWrap) HasUsageNormalized() bool`

HasUsageNormalized returns a boolean if a field has been set.

### GetUsage

`func (o *V0042AssocSharesObjWrap) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V0042AssocSharesObjWrap) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V0042AssocSharesObjWrap) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V0042AssocSharesObjWrap) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetFairshare

`func (o *V0042AssocSharesObjWrap) GetFairshare() V0042AssocSharesObjWrapFairshare`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0042AssocSharesObjWrap) GetFairshareOk() (*V0042AssocSharesObjWrapFairshare, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0042AssocSharesObjWrap) SetFairshare(v V0042AssocSharesObjWrapFairshare)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0042AssocSharesObjWrap) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.

### GetType

`func (o *V0042AssocSharesObjWrap) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0042AssocSharesObjWrap) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0042AssocSharesObjWrap) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *V0042AssocSharesObjWrap) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


