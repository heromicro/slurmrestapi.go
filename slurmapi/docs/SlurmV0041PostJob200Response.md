# SlurmV0041PostJob200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SlurmV0041PostJob200ResponseResultsInner**](SlurmV0041PostJob200ResponseResultsInner.md) | Job update results | [optional] 
**JobId** | Pointer to **string** | First updated Job ID - Use results instead | [optional] 
**StepId** | Pointer to **string** | First updated Step ID - Use results instead | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | First updated Job submission user message - Use results instead | [optional] 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmV0041PostJob200Response

`func NewSlurmV0041PostJob200Response() *SlurmV0041PostJob200Response`

NewSlurmV0041PostJob200Response instantiates a new SlurmV0041PostJob200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJob200ResponseWithDefaults

`func NewSlurmV0041PostJob200ResponseWithDefaults() *SlurmV0041PostJob200Response`

NewSlurmV0041PostJob200ResponseWithDefaults instantiates a new SlurmV0041PostJob200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SlurmV0041PostJob200Response) GetResults() []SlurmV0041PostJob200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SlurmV0041PostJob200Response) GetResultsOk() (*[]SlurmV0041PostJob200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SlurmV0041PostJob200Response) SetResults(v []SlurmV0041PostJob200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *SlurmV0041PostJob200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetJobId

`func (o *SlurmV0041PostJob200Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041PostJob200Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041PostJob200Response) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SlurmV0041PostJob200Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *SlurmV0041PostJob200Response) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SlurmV0041PostJob200Response) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SlurmV0041PostJob200Response) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *SlurmV0041PostJob200Response) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *SlurmV0041PostJob200Response) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *SlurmV0041PostJob200Response) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *SlurmV0041PostJob200Response) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *SlurmV0041PostJob200Response) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.

### GetMeta

`func (o *SlurmV0041PostJob200Response) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmV0041PostJob200Response) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmV0041PostJob200Response) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmV0041PostJob200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmV0041PostJob200Response) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmV0041PostJob200Response) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmV0041PostJob200Response) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmV0041PostJob200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmV0041PostJob200Response) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmV0041PostJob200Response) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmV0041PostJob200Response) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmV0041PostJob200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


