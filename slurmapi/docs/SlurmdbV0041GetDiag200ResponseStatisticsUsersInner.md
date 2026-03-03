# SlurmdbV0041GetDiag200ResponseStatisticsUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | User ID | [optional] 
**Count** | Pointer to **int32** | Number of RPCs processed | [optional] 
**Time** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime**](SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime.md) |  | [optional] 

## Methods

### NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInner

`func NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInner() *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner`

NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInner instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInnerWithDefaults

`func NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInnerWithDefaults() *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner`

NewSlurmdbV0041GetDiag200ResponseStatisticsUsersInnerWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetTime() SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) GetTimeOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) SetTime(v SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


