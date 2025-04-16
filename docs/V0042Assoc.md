# V0042Assoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0042Accounting**](V0042Accounting.md) |  | [optional] 
**Account** | Pointer to **string** | Account name | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Default** | Pointer to [**V0040AssocDefault**](V0040AssocDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Max** | Pointer to [**V0042AssocMax**](V0042AssocMax.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique ID | [optional] 
**IsDefault** | Pointer to **bool** | Is default association for user | [optional] 
**Lineage** | Pointer to **string** | Complete path up the hierarchy to the root association | [optional] 
**Min** | Pointer to [**V0042AssocMin**](V0042AssocMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** | Name of parent account | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**Priority** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 
**User** | **string** | User name | 

## Methods

### NewV0042Assoc

`func NewV0042Assoc(user string, ) *V0042Assoc`

NewV0042Assoc instantiates a new V0042Assoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042AssocWithDefaults

`func NewV0042AssocWithDefaults() *V0042Assoc`

NewV0042AssocWithDefaults instantiates a new V0042Assoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0042Assoc) GetAccounting() []V0042Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0042Assoc) GetAccountingOk() (*[]V0042Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0042Assoc) SetAccounting(v []V0042Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0042Assoc) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetAccount

`func (o *V0042Assoc) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0042Assoc) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0042Assoc) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0042Assoc) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0042Assoc) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0042Assoc) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0042Assoc) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0042Assoc) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComment

`func (o *V0042Assoc) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0042Assoc) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0042Assoc) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0042Assoc) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *V0042Assoc) GetDefault() V0040AssocDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0042Assoc) GetDefaultOk() (*V0040AssocDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0042Assoc) SetDefault(v V0040AssocDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0042Assoc) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0042Assoc) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0042Assoc) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0042Assoc) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0042Assoc) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0042Assoc) GetMax() V0042AssocMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0042Assoc) GetMaxOk() (*V0042AssocMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0042Assoc) SetMax(v V0042AssocMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0042Assoc) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetId

`func (o *V0042Assoc) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0042Assoc) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0042Assoc) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0042Assoc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0042Assoc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0042Assoc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0042Assoc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0042Assoc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLineage

`func (o *V0042Assoc) GetLineage() string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V0042Assoc) GetLineageOk() (*string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V0042Assoc) SetLineage(v string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V0042Assoc) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMin

`func (o *V0042Assoc) GetMin() V0042AssocMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0042Assoc) GetMinOk() (*V0042AssocMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0042Assoc) SetMin(v V0042AssocMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0042Assoc) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0042Assoc) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0042Assoc) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0042Assoc) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0042Assoc) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0042Assoc) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0042Assoc) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0042Assoc) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0042Assoc) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0042Assoc) GetPriority() V0042Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0042Assoc) GetPriorityOk() (*V0042Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0042Assoc) SetPriority(v V0042Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0042Assoc) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0042Assoc) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0042Assoc) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0042Assoc) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0042Assoc) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0042Assoc) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0042Assoc) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0042Assoc) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0042Assoc) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUser

`func (o *V0042Assoc) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0042Assoc) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0042Assoc) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


