# V0043JobTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Elapsed time in seconds | [optional] 
**Eligible** | Pointer to **int64** | Time when the job became eligible to run (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**End** | Pointer to **int64** | End time (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Planned** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**Start** | Pointer to **int64** | Time execution began (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Submission** | Pointer to **int64** | Time when the job was submitted (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Suspended** | Pointer to **int32** | Total time in suspended state in seconds | [optional] 
**System** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem.md) |  | [optional] 
**Limit** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 
**Total** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal.md) |  | [optional] 
**User** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser.md) |  | [optional] 

## Methods

### NewV0043JobTime

`func NewV0043JobTime() *V0043JobTime`

NewV0043JobTime instantiates a new V0043JobTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobTimeWithDefaults

`func NewV0043JobTimeWithDefaults() *V0043JobTime`

NewV0043JobTimeWithDefaults instantiates a new V0043JobTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0043JobTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0043JobTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0043JobTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0043JobTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *V0043JobTime) GetEligible() int64`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *V0043JobTime) GetEligibleOk() (*int64, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *V0043JobTime) SetEligible(v int64)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *V0043JobTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *V0043JobTime) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0043JobTime) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0043JobTime) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0043JobTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPlanned

`func (o *V0043JobTime) GetPlanned() V0043Uint64NoValStruct`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *V0043JobTime) GetPlannedOk() (*V0043Uint64NoValStruct, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *V0043JobTime) SetPlanned(v V0043Uint64NoValStruct)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *V0043JobTime) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetStart

`func (o *V0043JobTime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0043JobTime) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0043JobTime) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0043JobTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *V0043JobTime) GetSubmission() int64`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *V0043JobTime) GetSubmissionOk() (*int64, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *V0043JobTime) SetSubmission(v int64)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *V0043JobTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *V0043JobTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0043JobTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0043JobTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0043JobTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0043JobTime) GetSystem() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0043JobTime) GetSystemOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0043JobTime) SetSystem(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0043JobTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetLimit

`func (o *V0043JobTime) GetLimit() V0043Uint32NoValStruct`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V0043JobTime) GetLimitOk() (*V0043Uint32NoValStruct, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V0043JobTime) SetLimit(v V0043Uint32NoValStruct)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V0043JobTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *V0043JobTime) GetTotal() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0043JobTime) GetTotalOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0043JobTime) SetTotal(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0043JobTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0043JobTime) GetUser() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0043JobTime) GetUserOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0043JobTime) SetUser(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *V0043JobTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


