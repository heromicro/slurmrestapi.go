# SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User ID (numeric) | 
**User** | **string** | User name | 
**Count** | **int32** | Number of RPCs received | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime**](SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime.md) |  | 

## Methods

### NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner

`func NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner(userId int32, user string, count int32, totalTime int64, averageTime SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime, ) *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner`

NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner instantiates a new SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInnerWithDefaults

`func NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInnerWithDefaults() *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner`

NewSlurmV0041GetDiag200ResponseStatisticsRpcsByUserInnerWithDefaults instantiates a new SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) SetUser(v string)`

SetUser sets User field to given value.


### GetCount

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetAverageTime() SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) GetAverageTimeOk() (*SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner) SetAverageTime(v SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInnerAverageTime)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


