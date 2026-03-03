# SlurmV0041PostJobSubmit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**SlurmV0041PostJobSubmit200ResponseResult**](SlurmV0041PostJobSubmit200ResponseResult.md) |  | [optional] 
**JobId** | Pointer to **int32** | Submitted Job ID | [optional] 
**StepId** | Pointer to **string** | Submitted Step ID | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | Job submission user message | [optional] 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmV0041PostJobSubmit200Response

`func NewSlurmV0041PostJobSubmit200Response() *SlurmV0041PostJobSubmit200Response`

NewSlurmV0041PostJobSubmit200Response instantiates a new SlurmV0041PostJobSubmit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmit200ResponseWithDefaults

`func NewSlurmV0041PostJobSubmit200ResponseWithDefaults() *SlurmV0041PostJobSubmit200Response`

NewSlurmV0041PostJobSubmit200ResponseWithDefaults instantiates a new SlurmV0041PostJobSubmit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SlurmV0041PostJobSubmit200Response) GetResult() SlurmV0041PostJobSubmit200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SlurmV0041PostJobSubmit200Response) GetResultOk() (*SlurmV0041PostJobSubmit200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SlurmV0041PostJobSubmit200Response) SetResult(v SlurmV0041PostJobSubmit200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SlurmV0041PostJobSubmit200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetJobId

`func (o *SlurmV0041PostJobSubmit200Response) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041PostJobSubmit200Response) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041PostJobSubmit200Response) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SlurmV0041PostJobSubmit200Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *SlurmV0041PostJobSubmit200Response) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SlurmV0041PostJobSubmit200Response) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SlurmV0041PostJobSubmit200Response) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *SlurmV0041PostJobSubmit200Response) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200Response) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *SlurmV0041PostJobSubmit200Response) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200Response) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *SlurmV0041PostJobSubmit200Response) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.

### GetMeta

`func (o *SlurmV0041PostJobSubmit200Response) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmV0041PostJobSubmit200Response) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmV0041PostJobSubmit200Response) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmV0041PostJobSubmit200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmV0041PostJobSubmit200Response) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmV0041PostJobSubmit200Response) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmV0041PostJobSubmit200Response) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmV0041PostJobSubmit200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmV0041PostJobSubmit200Response) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmV0041PostJobSubmit200Response) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmV0041PostJobSubmit200Response) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmV0041PostJobSubmit200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


