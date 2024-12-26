# V0041OpenapiPartitionRespPartitionsInnerMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerNode**](V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerNode.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerSocket**](V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerSocket.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** | MaxMemPerCPU or MaxMemPerNode | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerCpu**](V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerCpu.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerNode**](V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerNode.md) |  | [optional] 
**Nodes** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsNodes**](V0041OpenapiPartitionRespPartitionsInnerMaximumsNodes.md) |  | [optional] 
**Shares** | Pointer to **int32** | OverSubscribe | [optional] 
**Oversubscribe** | Pointer to [**V0040PartitionInfoMaximumsOversubscribe**](V0040PartitionInfoMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsTime**](V0041OpenapiPartitionRespPartitionsInnerMaximumsTime.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsOverTimeLimit**](V0041OpenapiPartitionRespPartitionsInnerMaximumsOverTimeLimit.md) |  | [optional] 

## Methods

### NewV0041OpenapiPartitionRespPartitionsInnerMaximums

`func NewV0041OpenapiPartitionRespPartitionsInnerMaximums() *V0041OpenapiPartitionRespPartitionsInnerMaximums`

NewV0041OpenapiPartitionRespPartitionsInnerMaximums instantiates a new V0041OpenapiPartitionRespPartitionsInnerMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults

`func NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerMaximums`

NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerNode() V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerNode`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerNodeOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerNode, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetCpusPerNode(v V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerNode)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerSocket() V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerSocket`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerSocketOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerSocket, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetCpusPerSocket(v V0041OpenapiPartitionRespPartitionsInnerMaximumsCpusPerSocket)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerCpu() V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerCpu`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerCpuOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerCpu, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetPartitionMemoryPerCpu(v V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerCpu)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerNode() V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerNode`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerNodeOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerNode, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetPartitionMemoryPerNode(v V0041OpenapiPartitionRespPartitionsInnerMaximumsPartitionMemoryPerNode)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetNodes() V0041OpenapiPartitionRespPartitionsInnerMaximumsNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetNodesOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetNodes(v V0041OpenapiPartitionRespPartitionsInnerMaximumsNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOversubscribe() V0040PartitionInfoMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOversubscribeOk() (*V0040PartitionInfoMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetOversubscribe(v V0040PartitionInfoMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetTime() V0041OpenapiPartitionRespPartitionsInnerMaximumsTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetTimeOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetTime(v V0041OpenapiPartitionRespPartitionsInnerMaximumsTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOverTimeLimit() V0041OpenapiPartitionRespPartitionsInnerMaximumsOverTimeLimit`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOverTimeLimitOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsOverTimeLimit, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetOverTimeLimit(v V0041OpenapiPartitionRespPartitionsInnerMaximumsOverTimeLimit)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


