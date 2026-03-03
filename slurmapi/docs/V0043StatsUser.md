# V0043StatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | User ID | [optional] 
**Count** | Pointer to **int32** | Number of RPCs processed | [optional] 
**Time** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime**](SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime.md) |  | [optional] 

## Methods

### NewV0043StatsUser

`func NewV0043StatsUser() *V0043StatsUser`

NewV0043StatsUser instantiates a new V0043StatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StatsUserWithDefaults

`func NewV0043StatsUserWithDefaults() *V0043StatsUser`

NewV0043StatsUserWithDefaults instantiates a new V0043StatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *V0043StatsUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0043StatsUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0043StatsUser) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0043StatsUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCount

`func (o *V0043StatsUser) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0043StatsUser) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0043StatsUser) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0043StatsUser) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *V0043StatsUser) GetTime() SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0043StatsUser) GetTimeOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0043StatsUser) SetTime(v SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0043StatsUser) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


