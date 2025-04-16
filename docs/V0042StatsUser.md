# V0042StatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | User ID | [optional] 
**Count** | Pointer to **int32** | Number of RPCs processed | [optional] 
**Time** | Pointer to [**V0040StatsRpcTime**](V0040StatsRpcTime.md) |  | [optional] 

## Methods

### NewV0042StatsUser

`func NewV0042StatsUser() *V0042StatsUser`

NewV0042StatsUser instantiates a new V0042StatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsUserWithDefaults

`func NewV0042StatsUserWithDefaults() *V0042StatsUser`

NewV0042StatsUserWithDefaults instantiates a new V0042StatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *V0042StatsUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0042StatsUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0042StatsUser) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0042StatsUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCount

`func (o *V0042StatsUser) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042StatsUser) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042StatsUser) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0042StatsUser) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *V0042StatsUser) GetTime() V0040StatsRpcTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042StatsUser) GetTimeOk() (*V0040StatsRpcTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042StatsUser) SetTime(v V0040StatsRpcTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042StatsUser) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


