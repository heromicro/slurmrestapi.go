# SlurmV0041PostJobAllocateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hetjob** | Pointer to [**[]SlurmV0041PostJobSubmitRequestJobsInner**](SlurmV0041PostJobSubmitRequestJobsInner.md) | HetJob description | [optional] 
**Job** | Pointer to [**SlurmV0041PostJobSubmitRequestJob**](SlurmV0041PostJobSubmitRequestJob.md) |  | [optional] 

## Methods

### NewSlurmV0041PostJobAllocateRequest

`func NewSlurmV0041PostJobAllocateRequest() *SlurmV0041PostJobAllocateRequest`

NewSlurmV0041PostJobAllocateRequest instantiates a new SlurmV0041PostJobAllocateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobAllocateRequestWithDefaults

`func NewSlurmV0041PostJobAllocateRequestWithDefaults() *SlurmV0041PostJobAllocateRequest`

NewSlurmV0041PostJobAllocateRequestWithDefaults instantiates a new SlurmV0041PostJobAllocateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHetjob

`func (o *SlurmV0041PostJobAllocateRequest) GetHetjob() []SlurmV0041PostJobSubmitRequestJobsInner`

GetHetjob returns the Hetjob field if non-nil, zero value otherwise.

### GetHetjobOk

`func (o *SlurmV0041PostJobAllocateRequest) GetHetjobOk() (*[]SlurmV0041PostJobSubmitRequestJobsInner, bool)`

GetHetjobOk returns a tuple with the Hetjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjob

`func (o *SlurmV0041PostJobAllocateRequest) SetHetjob(v []SlurmV0041PostJobSubmitRequestJobsInner)`

SetHetjob sets Hetjob field to given value.

### HasHetjob

`func (o *SlurmV0041PostJobAllocateRequest) HasHetjob() bool`

HasHetjob returns a boolean if a field has been set.

### GetJob

`func (o *SlurmV0041PostJobAllocateRequest) GetJob() SlurmV0041PostJobSubmitRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SlurmV0041PostJobAllocateRequest) GetJobOk() (*SlurmV0041PostJobSubmitRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SlurmV0041PostJobAllocateRequest) SetJob(v SlurmV0041PostJobSubmitRequestJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *SlurmV0041PostJobAllocateRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


