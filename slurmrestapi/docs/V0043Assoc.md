# V0043Assoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0043Accounting**](V0043Accounting.md) |  | [optional] 
**Account** | Pointer to **string** | Account name | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Default** | Pointer to [**V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault**](V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Flags on the association | [optional] 
**Max** | Pointer to [**V0043AssocMax**](V0043AssocMax.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique ID (Association ID) | [optional] 
**IsDefault** | Pointer to **bool** | Is default association for user | [optional] 
**Lineage** | Pointer to **string** | Complete path up the hierarchy to the root association | [optional] 
**Min** | Pointer to [**V0043AssocMin**](V0043AssocMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** | Name of parent account | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**Priority** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 
**User** | **string** | User name | 

## Methods

### NewV0043Assoc

`func NewV0043Assoc(user string, ) *V0043Assoc`

NewV0043Assoc instantiates a new V0043Assoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043AssocWithDefaults

`func NewV0043AssocWithDefaults() *V0043Assoc`

NewV0043AssocWithDefaults instantiates a new V0043Assoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0043Assoc) GetAccounting() []V0043Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0043Assoc) GetAccountingOk() (*[]V0043Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0043Assoc) SetAccounting(v []V0043Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0043Assoc) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetAccount

`func (o *V0043Assoc) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0043Assoc) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0043Assoc) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0043Assoc) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0043Assoc) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0043Assoc) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0043Assoc) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0043Assoc) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComment

`func (o *V0043Assoc) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0043Assoc) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0043Assoc) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0043Assoc) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *V0043Assoc) GetDefault() V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0043Assoc) GetDefaultOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0043Assoc) SetDefault(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0043Assoc) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0043Assoc) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0043Assoc) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0043Assoc) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0043Assoc) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0043Assoc) GetMax() V0043AssocMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0043Assoc) GetMaxOk() (*V0043AssocMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0043Assoc) SetMax(v V0043AssocMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0043Assoc) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetId

`func (o *V0043Assoc) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0043Assoc) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0043Assoc) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0043Assoc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0043Assoc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0043Assoc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0043Assoc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0043Assoc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLineage

`func (o *V0043Assoc) GetLineage() string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V0043Assoc) GetLineageOk() (*string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V0043Assoc) SetLineage(v string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V0043Assoc) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMin

`func (o *V0043Assoc) GetMin() V0043AssocMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0043Assoc) GetMinOk() (*V0043AssocMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0043Assoc) SetMin(v V0043AssocMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0043Assoc) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0043Assoc) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0043Assoc) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0043Assoc) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0043Assoc) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0043Assoc) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0043Assoc) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0043Assoc) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0043Assoc) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0043Assoc) GetPriority() V0043Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0043Assoc) GetPriorityOk() (*V0043Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0043Assoc) SetPriority(v V0043Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0043Assoc) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0043Assoc) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0043Assoc) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0043Assoc) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0043Assoc) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0043Assoc) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0043Assoc) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0043Assoc) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0043Assoc) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUser

`func (o *V0043Assoc) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0043Assoc) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0043Assoc) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


