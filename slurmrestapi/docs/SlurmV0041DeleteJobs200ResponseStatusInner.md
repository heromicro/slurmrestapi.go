# SlurmV0041DeleteJobs200ResponseStatusInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**SlurmV0041DeleteJobs200ResponseStatusInnerError**](SlurmV0041DeleteJobs200ResponseStatusInnerError.md) |  | [optional] 
**StepId** | **string** | Job or Step ID that signaling failed | 
**JobId** | [**SlurmV0041DeleteJobs200ResponseStatusInnerJobId**](SlurmV0041DeleteJobs200ResponseStatusInnerJobId.md) |  | 
**Federation** | Pointer to [**SlurmV0041DeleteJobs200ResponseStatusInnerFederation**](SlurmV0041DeleteJobs200ResponseStatusInnerFederation.md) |  | [optional] 

## Methods

### NewSlurmV0041DeleteJobs200ResponseStatusInner

`func NewSlurmV0041DeleteJobs200ResponseStatusInner(stepId string, jobId SlurmV0041DeleteJobs200ResponseStatusInnerJobId, ) *SlurmV0041DeleteJobs200ResponseStatusInner`

NewSlurmV0041DeleteJobs200ResponseStatusInner instantiates a new SlurmV0041DeleteJobs200ResponseStatusInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041DeleteJobs200ResponseStatusInnerWithDefaults

`func NewSlurmV0041DeleteJobs200ResponseStatusInnerWithDefaults() *SlurmV0041DeleteJobs200ResponseStatusInner`

NewSlurmV0041DeleteJobs200ResponseStatusInnerWithDefaults instantiates a new SlurmV0041DeleteJobs200ResponseStatusInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetError() SlurmV0041DeleteJobs200ResponseStatusInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetErrorOk() (*SlurmV0041DeleteJobs200ResponseStatusInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) SetError(v SlurmV0041DeleteJobs200ResponseStatusInnerError)`

SetError sets Error field to given value.

### HasError

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStepId

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetJobId

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetJobId() SlurmV0041DeleteJobs200ResponseStatusInnerJobId`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetJobIdOk() (*SlurmV0041DeleteJobs200ResponseStatusInnerJobId, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) SetJobId(v SlurmV0041DeleteJobs200ResponseStatusInnerJobId)`

SetJobId sets JobId field to given value.


### GetFederation

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetFederation() SlurmV0041DeleteJobs200ResponseStatusInnerFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) GetFederationOk() (*SlurmV0041DeleteJobs200ResponseStatusInnerFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) SetFederation(v SlurmV0041DeleteJobs200ResponseStatusInnerFederation)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *SlurmV0041DeleteJobs200ResponseStatusInner) HasFederation() bool`

HasFederation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


