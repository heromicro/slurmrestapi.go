# V0042StatsRpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rpc** | Pointer to **string** | RPC type | [optional] 
**Count** | Pointer to **int32** | Number of RPCs processed | [optional] 
**Time** | Pointer to [**V0040StatsRpcTime**](V0040StatsRpcTime.md) |  | [optional] 

## Methods

### NewV0042StatsRpc

`func NewV0042StatsRpc() *V0042StatsRpc`

NewV0042StatsRpc instantiates a new V0042StatsRpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsRpcWithDefaults

`func NewV0042StatsRpcWithDefaults() *V0042StatsRpc`

NewV0042StatsRpcWithDefaults instantiates a new V0042StatsRpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpc

`func (o *V0042StatsRpc) GetRpc() string`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *V0042StatsRpc) GetRpcOk() (*string, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *V0042StatsRpc) SetRpc(v string)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *V0042StatsRpc) HasRpc() bool`

HasRpc returns a boolean if a field has been set.

### GetCount

`func (o *V0042StatsRpc) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042StatsRpc) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042StatsRpc) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0042StatsRpc) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *V0042StatsRpc) GetTime() V0040StatsRpcTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042StatsRpc) GetTimeOk() (*V0040StatsRpcTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042StatsRpc) SetTime(v V0040StatsRpcTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042StatsRpc) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


