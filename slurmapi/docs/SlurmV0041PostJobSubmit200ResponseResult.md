# SlurmV0041PostJobSubmit200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** | New job ID | [optional] 
**StepId** | Pointer to **string** | New job step ID | [optional] 
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**Error** | Pointer to **string** | Error message | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | Message to user from job_submit plugin | [optional] 

## Methods

### NewSlurmV0041PostJobSubmit200ResponseResult

`func NewSlurmV0041PostJobSubmit200ResponseResult() *SlurmV0041PostJobSubmit200ResponseResult`

NewSlurmV0041PostJobSubmit200ResponseResult instantiates a new SlurmV0041PostJobSubmit200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmit200ResponseResultWithDefaults

`func NewSlurmV0041PostJobSubmit200ResponseResultWithDefaults() *SlurmV0041PostJobSubmit200ResponseResult`

NewSlurmV0041PostJobSubmit200ResponseResultWithDefaults instantiates a new SlurmV0041PostJobSubmit200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *SlurmV0041PostJobSubmit200ResponseResult) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetErrorCode

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SlurmV0041PostJobSubmit200ResponseResult) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SlurmV0041PostJobSubmit200ResponseResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetError

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SlurmV0041PostJobSubmit200ResponseResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SlurmV0041PostJobSubmit200ResponseResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *SlurmV0041PostJobSubmit200ResponseResult) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200ResponseResult) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200ResponseResult) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


