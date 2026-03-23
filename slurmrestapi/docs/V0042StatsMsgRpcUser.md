# V0042StatsMsgRpcUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User ID (numeric) | 
**User** | **string** | User name | 
**Count** | **int32** | Number of RPCs received | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | 

## Methods

### NewV0042StatsMsgRpcUser

`func NewV0042StatsMsgRpcUser(userId int32, user string, count int32, totalTime int64, averageTime V0042Uint64NoValStruct, ) *V0042StatsMsgRpcUser`

NewV0042StatsMsgRpcUser instantiates a new V0042StatsMsgRpcUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsMsgRpcUserWithDefaults

`func NewV0042StatsMsgRpcUserWithDefaults() *V0042StatsMsgRpcUser`

NewV0042StatsMsgRpcUserWithDefaults instantiates a new V0042StatsMsgRpcUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *V0042StatsMsgRpcUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0042StatsMsgRpcUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0042StatsMsgRpcUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *V0042StatsMsgRpcUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0042StatsMsgRpcUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0042StatsMsgRpcUser) SetUser(v string)`

SetUser sets User field to given value.


### GetCount

`func (o *V0042StatsMsgRpcUser) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042StatsMsgRpcUser) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042StatsMsgRpcUser) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalTime

`func (o *V0042StatsMsgRpcUser) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0042StatsMsgRpcUser) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0042StatsMsgRpcUser) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *V0042StatsMsgRpcUser) GetAverageTime() V0042Uint64NoValStruct`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0042StatsMsgRpcUser) GetAverageTimeOk() (*V0042Uint64NoValStruct, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0042StatsMsgRpcUser) SetAverageTime(v V0042Uint64NoValStruct)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


