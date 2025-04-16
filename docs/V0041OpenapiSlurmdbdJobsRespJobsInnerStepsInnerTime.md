# V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Elapsed time in seconds | [optional] 
**End** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeEnd**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeEnd.md) |  | [optional] 
**Start** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeStart**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeStart.md) |  | [optional] 
**Suspended** | Pointer to **int32** | Time in suspended state in seconds | [optional] 
**System** | Pointer to [**V0040StepTimeSystem**](V0040StepTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**V0040StepTimeTotal**](V0040StepTimeTotal.md) |  | [optional] 
**User** | Pointer to [**V0040StepTimeUser**](V0040StepTimeUser.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetEnd() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetEndOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetEnd(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetStart() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetStartOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetStart(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSystem() V0040StepTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSystemOk() (*V0040StepTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetSystem(v V0040StepTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetTotal() V0040StepTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetTotalOk() (*V0040StepTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetTotal(v V0040StepTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetUser() V0040StepTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetUserOk() (*V0040StepTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetUser(v V0040StepTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


