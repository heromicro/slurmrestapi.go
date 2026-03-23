# V0043PartitionInfoMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** | Raw value for MaxMemPerCPU or MaxMemPerNode | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**Nodes** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**Shares** | Pointer to **int32** | OverSubscribe - Controls the ability of the partition to execute more than one job at a time on each resource | [optional] 
**Oversubscribe** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe**](V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0043Uint16NoValStruct**](V0043Uint16NoValStruct.md) |  | [optional] 

## Methods

### NewV0043PartitionInfoMaximums

`func NewV0043PartitionInfoMaximums() *V0043PartitionInfoMaximums`

NewV0043PartitionInfoMaximums instantiates a new V0043PartitionInfoMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043PartitionInfoMaximumsWithDefaults

`func NewV0043PartitionInfoMaximumsWithDefaults() *V0043PartitionInfoMaximums`

NewV0043PartitionInfoMaximumsWithDefaults instantiates a new V0043PartitionInfoMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0043PartitionInfoMaximums) GetCpusPerNode() V0043Uint32NoValStruct`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0043PartitionInfoMaximums) GetCpusPerNodeOk() (*V0043Uint32NoValStruct, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0043PartitionInfoMaximums) SetCpusPerNode(v V0043Uint32NoValStruct)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0043PartitionInfoMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0043PartitionInfoMaximums) GetCpusPerSocket() V0043Uint32NoValStruct`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0043PartitionInfoMaximums) GetCpusPerSocketOk() (*V0043Uint32NoValStruct, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0043PartitionInfoMaximums) SetCpusPerSocket(v V0043Uint32NoValStruct)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0043PartitionInfoMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0043PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) GetPartitionMemoryPerCpu() V0043Uint64NoValStruct`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0043PartitionInfoMaximums) GetPartitionMemoryPerCpuOk() (*V0043Uint64NoValStruct, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) SetPartitionMemoryPerCpu(v V0043Uint64NoValStruct)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0043PartitionInfoMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0043PartitionInfoMaximums) GetPartitionMemoryPerNode() V0043Uint64NoValStruct`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0043PartitionInfoMaximums) GetPartitionMemoryPerNodeOk() (*V0043Uint64NoValStruct, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0043PartitionInfoMaximums) SetPartitionMemoryPerNode(v V0043Uint64NoValStruct)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0043PartitionInfoMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0043PartitionInfoMaximums) GetNodes() V0043Uint32NoValStruct`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0043PartitionInfoMaximums) GetNodesOk() (*V0043Uint32NoValStruct, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0043PartitionInfoMaximums) SetNodes(v V0043Uint32NoValStruct)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0043PartitionInfoMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0043PartitionInfoMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0043PartitionInfoMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0043PartitionInfoMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0043PartitionInfoMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0043PartitionInfoMaximums) GetOversubscribe() V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0043PartitionInfoMaximums) GetOversubscribeOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0043PartitionInfoMaximums) SetOversubscribe(v V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0043PartitionInfoMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0043PartitionInfoMaximums) GetTime() V0043Uint32NoValStruct`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0043PartitionInfoMaximums) GetTimeOk() (*V0043Uint32NoValStruct, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0043PartitionInfoMaximums) SetTime(v V0043Uint32NoValStruct)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0043PartitionInfoMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0043PartitionInfoMaximums) GetOverTimeLimit() V0043Uint16NoValStruct`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0043PartitionInfoMaximums) GetOverTimeLimitOk() (*V0043Uint16NoValStruct, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0043PartitionInfoMaximums) SetOverTimeLimit(v V0043Uint16NoValStruct)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0043PartitionInfoMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


