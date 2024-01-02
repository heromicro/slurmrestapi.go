# V0040StatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to [**V0040StatsRpcTime**](V0040StatsRpcTime.md) |  | [optional] 

## Methods

### NewV0040StatsUser

`func NewV0040StatsUser() *V0040StatsUser`

NewV0040StatsUser instantiates a new V0040StatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StatsUserWithDefaults

`func NewV0040StatsUserWithDefaults() *V0040StatsUser`

NewV0040StatsUserWithDefaults instantiates a new V0040StatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *V0040StatsUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040StatsUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040StatsUser) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0040StatsUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCount

`func (o *V0040StatsUser) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0040StatsUser) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0040StatsUser) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0040StatsUser) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *V0040StatsUser) GetTime() V0040StatsRpcTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0040StatsUser) GetTimeOk() (*V0040StatsRpcTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0040StatsUser) SetTime(v V0040StatsRpcTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0040StatsUser) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


