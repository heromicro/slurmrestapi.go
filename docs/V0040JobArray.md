# V0040JobArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** |  | [optional] 
**Limits** | Pointer to [**V0040JobArrayLimits**](V0040JobArrayLimits.md) |  | [optional] 
**TaskId** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 

## Methods

### NewV0040JobArray

`func NewV0040JobArray() *V0040JobArray`

NewV0040JobArray instantiates a new V0040JobArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobArrayWithDefaults

`func NewV0040JobArrayWithDefaults() *V0040JobArray`

NewV0040JobArrayWithDefaults instantiates a new V0040JobArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0040JobArray) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040JobArray) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040JobArray) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040JobArray) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLimits

`func (o *V0040JobArray) GetLimits() V0040JobArrayLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *V0040JobArray) GetLimitsOk() (*V0040JobArrayLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *V0040JobArray) SetLimits(v V0040JobArrayLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *V0040JobArray) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTaskId

`func (o *V0040JobArray) GetTaskId() V0040Uint32NoVal`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V0040JobArray) GetTaskIdOk() (*V0040Uint32NoVal, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V0040JobArray) SetTaskId(v V0040Uint32NoVal)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *V0040JobArray) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTask

`func (o *V0040JobArray) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0040JobArray) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0040JobArray) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0040JobArray) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


