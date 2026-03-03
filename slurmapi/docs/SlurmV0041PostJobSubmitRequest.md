# SlurmV0041PostJobSubmitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | Job batch script contents; Same as the script field in jobs[0] or job. | [optional] 
**Jobs** | Pointer to [**[]SlurmV0041PostJobSubmitRequestJobsInner**](SlurmV0041PostJobSubmitRequestJobsInner.md) | HetJob description | [optional] 
**Job** | Pointer to [**SlurmV0041PostJobSubmitRequestJob**](SlurmV0041PostJobSubmitRequestJob.md) |  | [optional] 

## Methods

### NewSlurmV0041PostJobSubmitRequest

`func NewSlurmV0041PostJobSubmitRequest() *SlurmV0041PostJobSubmitRequest`

NewSlurmV0041PostJobSubmitRequest instantiates a new SlurmV0041PostJobSubmitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmitRequestWithDefaults

`func NewSlurmV0041PostJobSubmitRequestWithDefaults() *SlurmV0041PostJobSubmitRequest`

NewSlurmV0041PostJobSubmitRequestWithDefaults instantiates a new SlurmV0041PostJobSubmitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *SlurmV0041PostJobSubmitRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *SlurmV0041PostJobSubmitRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *SlurmV0041PostJobSubmitRequest) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *SlurmV0041PostJobSubmitRequest) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJobs

`func (o *SlurmV0041PostJobSubmitRequest) GetJobs() []SlurmV0041PostJobSubmitRequestJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *SlurmV0041PostJobSubmitRequest) GetJobsOk() (*[]SlurmV0041PostJobSubmitRequestJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *SlurmV0041PostJobSubmitRequest) SetJobs(v []SlurmV0041PostJobSubmitRequestJobsInner)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *SlurmV0041PostJobSubmitRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetJob

`func (o *SlurmV0041PostJobSubmitRequest) GetJob() SlurmV0041PostJobSubmitRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SlurmV0041PostJobSubmitRequest) GetJobOk() (*SlurmV0041PostJobSubmitRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SlurmV0041PostJobSubmitRequest) SetJob(v SlurmV0041PostJobSubmitRequestJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *SlurmV0041PostJobSubmitRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


