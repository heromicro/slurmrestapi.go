# SlurmV0041PostJob200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** | Job ID for updated job | [optional] 
**StepId** | Pointer to **string** | Step ID for updated job | [optional] 
**Error** | Pointer to **string** | Verbose update status or error | [optional] 
**ErrorCode** | Pointer to **int32** | Verbose update status or error | [optional] 
**Why** | Pointer to **string** | Update response message | [optional] 

## Methods

### NewSlurmV0041PostJob200ResponseResultsInner

`func NewSlurmV0041PostJob200ResponseResultsInner() *SlurmV0041PostJob200ResponseResultsInner`

NewSlurmV0041PostJob200ResponseResultsInner instantiates a new SlurmV0041PostJob200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJob200ResponseResultsInnerWithDefaults

`func NewSlurmV0041PostJob200ResponseResultsInnerWithDefaults() *SlurmV0041PostJob200ResponseResultsInner`

NewSlurmV0041PostJob200ResponseResultsInnerWithDefaults instantiates a new SlurmV0041PostJob200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041PostJob200ResponseResultsInner) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SlurmV0041PostJob200ResponseResultsInner) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SlurmV0041PostJob200ResponseResultsInner) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *SlurmV0041PostJob200ResponseResultsInner) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetError

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SlurmV0041PostJob200ResponseResultsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SlurmV0041PostJob200ResponseResultsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SlurmV0041PostJob200ResponseResultsInner) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SlurmV0041PostJob200ResponseResultsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetWhy

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *SlurmV0041PostJob200ResponseResultsInner) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *SlurmV0041PostJob200ResponseResultsInner) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *SlurmV0041PostJob200ResponseResultsInner) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


