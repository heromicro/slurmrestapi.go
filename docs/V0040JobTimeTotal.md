# V0040JobTimeTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int64** | Sum of System and User CPU time used by the job in seconds | [optional] 
**Microseconds** | Pointer to **int64** | Sum of System and User CPU time used by the job in microseconds | [optional] 

## Methods

### NewV0040JobTimeTotal

`func NewV0040JobTimeTotal() *V0040JobTimeTotal`

NewV0040JobTimeTotal instantiates a new V0040JobTimeTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobTimeTotalWithDefaults

`func NewV0040JobTimeTotalWithDefaults() *V0040JobTimeTotal`

NewV0040JobTimeTotalWithDefaults instantiates a new V0040JobTimeTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *V0040JobTimeTotal) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *V0040JobTimeTotal) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *V0040JobTimeTotal) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *V0040JobTimeTotal) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetMicroseconds

`func (o *V0040JobTimeTotal) GetMicroseconds() int64`

GetMicroseconds returns the Microseconds field if non-nil, zero value otherwise.

### GetMicrosecondsOk

`func (o *V0040JobTimeTotal) GetMicrosecondsOk() (*int64, bool)`

GetMicrosecondsOk returns a tuple with the Microseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroseconds

`func (o *V0040JobTimeTotal) SetMicroseconds(v int64)`

SetMicroseconds sets Microseconds field to given value.

### HasMicroseconds

`func (o *V0040JobTimeTotal) HasMicroseconds() bool`

HasMicroseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


