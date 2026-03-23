# V0042PartitionInfoDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryPerCpu** | Pointer to **int64** | DefMemPerCPU or DefMemPerNode | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**Time** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Job** | Pointer to **string** | JobDefaults | [optional] 

## Methods

### NewV0042PartitionInfoDefaults

`func NewV0042PartitionInfoDefaults() *V0042PartitionInfoDefaults`

NewV0042PartitionInfoDefaults instantiates a new V0042PartitionInfoDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042PartitionInfoDefaultsWithDefaults

`func NewV0042PartitionInfoDefaultsWithDefaults() *V0042PartitionInfoDefaults`

NewV0042PartitionInfoDefaultsWithDefaults instantiates a new V0042PartitionInfoDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0042PartitionInfoDefaults) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) GetPartitionMemoryPerCpu() V0042Uint64NoValStruct`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0042PartitionInfoDefaults) GetPartitionMemoryPerCpuOk() (*V0042Uint64NoValStruct, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) SetPartitionMemoryPerCpu(v V0042Uint64NoValStruct)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0042PartitionInfoDefaults) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0042PartitionInfoDefaults) GetPartitionMemoryPerNode() V0042Uint64NoValStruct`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0042PartitionInfoDefaults) GetPartitionMemoryPerNodeOk() (*V0042Uint64NoValStruct, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0042PartitionInfoDefaults) SetPartitionMemoryPerNode(v V0042Uint64NoValStruct)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0042PartitionInfoDefaults) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetTime

`func (o *V0042PartitionInfoDefaults) GetTime() V0042Uint32NoValStruct`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042PartitionInfoDefaults) GetTimeOk() (*V0042Uint32NoValStruct, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042PartitionInfoDefaults) SetTime(v V0042Uint32NoValStruct)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042PartitionInfoDefaults) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetJob

`func (o *V0042PartitionInfoDefaults) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0042PartitionInfoDefaults) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0042PartitionInfoDefaults) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0042PartitionInfoDefaults) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


