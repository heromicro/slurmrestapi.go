# V0042KillJobsRespJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**V0040KillJobsRespJobError**](V0040KillJobsRespJobError.md) |  | [optional] 
**StepId** | **string** | Job or Step ID that signaling failed | 
**JobId** | [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | 
**Federation** | Pointer to [**V0040KillJobsRespJobFederation**](V0040KillJobsRespJobFederation.md) |  | [optional] 

## Methods

### NewV0042KillJobsRespJob

`func NewV0042KillJobsRespJob(stepId string, jobId V0042Uint32NoValStruct, ) *V0042KillJobsRespJob`

NewV0042KillJobsRespJob instantiates a new V0042KillJobsRespJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042KillJobsRespJobWithDefaults

`func NewV0042KillJobsRespJobWithDefaults() *V0042KillJobsRespJob`

NewV0042KillJobsRespJobWithDefaults instantiates a new V0042KillJobsRespJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V0042KillJobsRespJob) GetError() V0040KillJobsRespJobError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0042KillJobsRespJob) GetErrorOk() (*V0040KillJobsRespJobError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0042KillJobsRespJob) SetError(v V0040KillJobsRespJobError)`

SetError sets Error field to given value.

### HasError

`func (o *V0042KillJobsRespJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStepId

`func (o *V0042KillJobsRespJob) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0042KillJobsRespJob) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0042KillJobsRespJob) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetJobId

`func (o *V0042KillJobsRespJob) GetJobId() V0042Uint32NoValStruct`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0042KillJobsRespJob) GetJobIdOk() (*V0042Uint32NoValStruct, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0042KillJobsRespJob) SetJobId(v V0042Uint32NoValStruct)`

SetJobId sets JobId field to given value.


### GetFederation

`func (o *V0042KillJobsRespJob) GetFederation() V0040KillJobsRespJobFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *V0042KillJobsRespJob) GetFederationOk() (*V0040KillJobsRespJobFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *V0042KillJobsRespJob) SetFederation(v V0040KillJobsRespJobFederation)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *V0042KillJobsRespJob) HasFederation() bool`

HasFederation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


