# V0043JobArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** | Job ID of job array, or 0 if N/A | [optional] 
**Limits** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayLimits**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayLimits.md) |  | [optional] 
**TaskId** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**Task** | Pointer to **string** | String expression of task IDs in this record | [optional] 

## Methods

### NewV0043JobArray

`func NewV0043JobArray() *V0043JobArray`

NewV0043JobArray instantiates a new V0043JobArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobArrayWithDefaults

`func NewV0043JobArrayWithDefaults() *V0043JobArray`

NewV0043JobArrayWithDefaults instantiates a new V0043JobArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0043JobArray) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0043JobArray) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0043JobArray) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0043JobArray) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLimits

`func (o *V0043JobArray) GetLimits() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *V0043JobArray) GetLimitsOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *V0043JobArray) SetLimits(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *V0043JobArray) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTaskId

`func (o *V0043JobArray) GetTaskId() V0043Uint32NoValStruct`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V0043JobArray) GetTaskIdOk() (*V0043Uint32NoValStruct, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V0043JobArray) SetTaskId(v V0043Uint32NoValStruct)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *V0043JobArray) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTask

`func (o *V0043JobArray) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0043JobArray) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0043JobArray) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0043JobArray) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


