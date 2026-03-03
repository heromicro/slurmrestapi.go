# SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Message type as integer | 
**MessageType** | **string** | Message type as string | 
**Count** | **int32** | Number of RPCs received | 
**Queued** | **int32** | Number of RPCs queued | 
**Dropped** | **int64** | Number of RPCs dropped | 
**CycleLast** | **int32** | Number of RPCs processed within the last RPC queue cycle | 
**CycleMax** | **int32** | Maximum number of RPCs processed within a RPC queue cycle since start | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime**](SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime.md) |  | 

## Methods

### NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner

`func NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner(typeId int32, messageType string, count int32, queued int32, dropped int64, cycleLast int32, cycleMax int32, totalTime int64, averageTime SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime, ) *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner`

NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner instantiates a new SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerWithDefaults

`func NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerWithDefaults() *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner`

NewSlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerWithDefaults instantiates a new SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetCount(v int32)`

SetCount sets Count field to given value.


### GetQueued

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetDropped

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetDropped(v int64)`

SetDropped sets Dropped field to given value.


### GetCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCycleLast() int32`

GetCycleLast returns the CycleLast field if non-nil, zero value otherwise.

### GetCycleLastOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCycleLastOk() (*int32, bool)`

GetCycleLastOk returns a tuple with the CycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetCycleLast(v int32)`

SetCycleLast sets CycleLast field to given value.


### GetCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCycleMax() int32`

GetCycleMax returns the CycleMax field if non-nil, zero value otherwise.

### GetCycleMaxOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetCycleMaxOk() (*int32, bool)`

GetCycleMaxOk returns a tuple with the CycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetCycleMax(v int32)`

SetCycleMax sets CycleMax field to given value.


### GetTotalTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetAverageTime() SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) GetAverageTimeOk() (*SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner) SetAverageTime(v SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


