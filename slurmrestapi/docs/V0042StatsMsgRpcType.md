# V0042StatsMsgRpcType

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
**AverageTime** | [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | 

## Methods

### NewV0042StatsMsgRpcType

`func NewV0042StatsMsgRpcType(typeId int32, messageType string, count int32, queued int32, dropped int64, cycleLast int32, cycleMax int32, totalTime int64, averageTime V0042Uint64NoValStruct, ) *V0042StatsMsgRpcType`

NewV0042StatsMsgRpcType instantiates a new V0042StatsMsgRpcType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsMsgRpcTypeWithDefaults

`func NewV0042StatsMsgRpcTypeWithDefaults() *V0042StatsMsgRpcType`

NewV0042StatsMsgRpcTypeWithDefaults instantiates a new V0042StatsMsgRpcType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *V0042StatsMsgRpcType) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *V0042StatsMsgRpcType) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *V0042StatsMsgRpcType) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *V0042StatsMsgRpcType) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V0042StatsMsgRpcType) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V0042StatsMsgRpcType) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *V0042StatsMsgRpcType) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042StatsMsgRpcType) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042StatsMsgRpcType) SetCount(v int32)`

SetCount sets Count field to given value.


### GetQueued

`func (o *V0042StatsMsgRpcType) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *V0042StatsMsgRpcType) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *V0042StatsMsgRpcType) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetDropped

`func (o *V0042StatsMsgRpcType) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *V0042StatsMsgRpcType) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *V0042StatsMsgRpcType) SetDropped(v int64)`

SetDropped sets Dropped field to given value.


### GetCycleLast

`func (o *V0042StatsMsgRpcType) GetCycleLast() int32`

GetCycleLast returns the CycleLast field if non-nil, zero value otherwise.

### GetCycleLastOk

`func (o *V0042StatsMsgRpcType) GetCycleLastOk() (*int32, bool)`

GetCycleLastOk returns a tuple with the CycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLast

`func (o *V0042StatsMsgRpcType) SetCycleLast(v int32)`

SetCycleLast sets CycleLast field to given value.


### GetCycleMax

`func (o *V0042StatsMsgRpcType) GetCycleMax() int32`

GetCycleMax returns the CycleMax field if non-nil, zero value otherwise.

### GetCycleMaxOk

`func (o *V0042StatsMsgRpcType) GetCycleMaxOk() (*int32, bool)`

GetCycleMaxOk returns a tuple with the CycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleMax

`func (o *V0042StatsMsgRpcType) SetCycleMax(v int32)`

SetCycleMax sets CycleMax field to given value.


### GetTotalTime

`func (o *V0042StatsMsgRpcType) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0042StatsMsgRpcType) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0042StatsMsgRpcType) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *V0042StatsMsgRpcType) GetAverageTime() V0042Uint64NoValStruct`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0042StatsMsgRpcType) GetAverageTimeOk() (*V0042Uint64NoValStruct, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0042StatsMsgRpcType) SetAverageTime(v V0042Uint64NoValStruct)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


