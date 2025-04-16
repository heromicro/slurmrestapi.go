# V0042PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**V0040PartitionInfoNodes**](V0040PartitionInfoNodes.md) |  | [optional] 
**Accounts** | Pointer to [**V0040PartitionInfoAccounts**](V0040PartitionInfoAccounts.md) |  | [optional] 
**Groups** | Pointer to [**V0040PartitionInfoGroups**](V0040PartitionInfoGroups.md) |  | [optional] 
**Qos** | Pointer to [**V0040PartitionInfoQos**](V0040PartitionInfoQos.md) |  | [optional] 
**Alternate** | Pointer to **string** | Alternate | [optional] 
**Tres** | Pointer to [**V0040PartitionInfoTres**](V0040PartitionInfoTres.md) |  | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**SelectType** | Pointer to **[]string** |  | [optional] 
**Cpus** | Pointer to [**V0040PartitionInfoCpus**](V0040PartitionInfoCpus.md) |  | [optional] 
**Defaults** | Pointer to [**V0042PartitionInfoDefaults**](V0042PartitionInfoDefaults.md) |  | [optional] 
**GraceTime** | Pointer to **int32** | GraceTime | [optional] 
**Maximums** | Pointer to [**V0042PartitionInfoMaximums**](V0042PartitionInfoMaximums.md) |  | [optional] 
**Minimums** | Pointer to [**V0040PartitionInfoMinimums**](V0040PartitionInfoMinimums.md) |  | [optional] 
**Name** | Pointer to **string** | PartitionName | [optional] 
**NodeSets** | Pointer to **string** | NodeSets | [optional] 
**Priority** | Pointer to [**V0040PartitionInfoPriority**](V0040PartitionInfoPriority.md) |  | [optional] 
**Timeouts** | Pointer to [**V0042PartitionInfoTimeouts**](V0042PartitionInfoTimeouts.md) |  | [optional] 
**Partition** | Pointer to [**V0042PartitionInfoPartition**](V0042PartitionInfoPartition.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0042PartitionInfo

`func NewV0042PartitionInfo() *V0042PartitionInfo`

NewV0042PartitionInfo instantiates a new V0042PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042PartitionInfoWithDefaults

`func NewV0042PartitionInfoWithDefaults() *V0042PartitionInfo`

NewV0042PartitionInfoWithDefaults instantiates a new V0042PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0042PartitionInfo) GetNodes() V0040PartitionInfoNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042PartitionInfo) GetNodesOk() (*V0040PartitionInfoNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042PartitionInfo) SetNodes(v V0040PartitionInfoNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042PartitionInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAccounts

`func (o *V0042PartitionInfo) GetAccounts() V0040PartitionInfoAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0042PartitionInfo) GetAccountsOk() (*V0040PartitionInfoAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0042PartitionInfo) SetAccounts(v V0040PartitionInfoAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0042PartitionInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *V0042PartitionInfo) GetGroups() V0040PartitionInfoGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0042PartitionInfo) GetGroupsOk() (*V0040PartitionInfoGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0042PartitionInfo) SetGroups(v V0040PartitionInfoGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0042PartitionInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetQos

`func (o *V0042PartitionInfo) GetQos() V0040PartitionInfoQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0042PartitionInfo) GetQosOk() (*V0040PartitionInfoQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0042PartitionInfo) SetQos(v V0040PartitionInfoQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0042PartitionInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetAlternate

`func (o *V0042PartitionInfo) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *V0042PartitionInfo) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *V0042PartitionInfo) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *V0042PartitionInfo) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetTres

`func (o *V0042PartitionInfo) GetTres() V0040PartitionInfoTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0042PartitionInfo) GetTresOk() (*V0040PartitionInfoTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0042PartitionInfo) SetTres(v V0040PartitionInfoTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0042PartitionInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetCluster

`func (o *V0042PartitionInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0042PartitionInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0042PartitionInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0042PartitionInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSelectType

`func (o *V0042PartitionInfo) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0042PartitionInfo) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0042PartitionInfo) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0042PartitionInfo) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetCpus

`func (o *V0042PartitionInfo) GetCpus() V0040PartitionInfoCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0042PartitionInfo) GetCpusOk() (*V0040PartitionInfoCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0042PartitionInfo) SetCpus(v V0040PartitionInfoCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0042PartitionInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDefaults

`func (o *V0042PartitionInfo) GetDefaults() V0042PartitionInfoDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *V0042PartitionInfo) GetDefaultsOk() (*V0042PartitionInfoDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *V0042PartitionInfo) SetDefaults(v V0042PartitionInfoDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *V0042PartitionInfo) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGraceTime

`func (o *V0042PartitionInfo) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0042PartitionInfo) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0042PartitionInfo) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0042PartitionInfo) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMaximums

`func (o *V0042PartitionInfo) GetMaximums() V0042PartitionInfoMaximums`

GetMaximums returns the Maximums field if non-nil, zero value otherwise.

### GetMaximumsOk

`func (o *V0042PartitionInfo) GetMaximumsOk() (*V0042PartitionInfoMaximums, bool)`

GetMaximumsOk returns a tuple with the Maximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximums

`func (o *V0042PartitionInfo) SetMaximums(v V0042PartitionInfoMaximums)`

SetMaximums sets Maximums field to given value.

### HasMaximums

`func (o *V0042PartitionInfo) HasMaximums() bool`

HasMaximums returns a boolean if a field has been set.

### GetMinimums

`func (o *V0042PartitionInfo) GetMinimums() V0040PartitionInfoMinimums`

GetMinimums returns the Minimums field if non-nil, zero value otherwise.

### GetMinimumsOk

`func (o *V0042PartitionInfo) GetMinimumsOk() (*V0040PartitionInfoMinimums, bool)`

GetMinimumsOk returns a tuple with the Minimums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimums

`func (o *V0042PartitionInfo) SetMinimums(v V0040PartitionInfoMinimums)`

SetMinimums sets Minimums field to given value.

### HasMinimums

`func (o *V0042PartitionInfo) HasMinimums() bool`

HasMinimums returns a boolean if a field has been set.

### GetName

`func (o *V0042PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042PartitionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0042PartitionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSets

`func (o *V0042PartitionInfo) GetNodeSets() string`

GetNodeSets returns the NodeSets field if non-nil, zero value otherwise.

### GetNodeSetsOk

`func (o *V0042PartitionInfo) GetNodeSetsOk() (*string, bool)`

GetNodeSetsOk returns a tuple with the NodeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSets

`func (o *V0042PartitionInfo) SetNodeSets(v string)`

SetNodeSets sets NodeSets field to given value.

### HasNodeSets

`func (o *V0042PartitionInfo) HasNodeSets() bool`

HasNodeSets returns a boolean if a field has been set.

### GetPriority

`func (o *V0042PartitionInfo) GetPriority() V0040PartitionInfoPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0042PartitionInfo) GetPriorityOk() (*V0040PartitionInfoPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0042PartitionInfo) SetPriority(v V0040PartitionInfoPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0042PartitionInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeouts

`func (o *V0042PartitionInfo) GetTimeouts() V0042PartitionInfoTimeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *V0042PartitionInfo) GetTimeoutsOk() (*V0042PartitionInfoTimeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *V0042PartitionInfo) SetTimeouts(v V0042PartitionInfoTimeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *V0042PartitionInfo) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetPartition

`func (o *V0042PartitionInfo) GetPartition() V0042PartitionInfoPartition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0042PartitionInfo) GetPartitionOk() (*V0042PartitionInfoPartition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0042PartitionInfo) SetPartition(v V0042PartitionInfoPartition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0042PartitionInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0042PartitionInfo) GetSuspendTime() V0042Uint32NoValStruct`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0042PartitionInfo) GetSuspendTimeOk() (*V0042Uint32NoValStruct, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0042PartitionInfo) SetSuspendTime(v V0042Uint32NoValStruct)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0042PartitionInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


