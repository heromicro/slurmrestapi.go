# SlurmV0041GetDiag200ResponseStatisticsBfExit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndJobQueue** | Pointer to **int32** | Reached end of queue | [optional] 
**BfMaxJobStart** | Pointer to **int32** | Reached number of jobs allowed to start | [optional] 
**BfMaxJobTest** | Pointer to **int32** | Reached number of jobs allowed to be tested | [optional] 
**BfMaxTime** | Pointer to **int32** | Reached maximum allowed scheduler time | [optional] 
**BfNodeSpaceSize** | Pointer to **int32** | Reached table size limit | [optional] 
**StateChanged** | Pointer to **int32** | System state changed | [optional] 

## Methods

### NewSlurmV0041GetDiag200ResponseStatisticsBfExit

`func NewSlurmV0041GetDiag200ResponseStatisticsBfExit() *SlurmV0041GetDiag200ResponseStatisticsBfExit`

NewSlurmV0041GetDiag200ResponseStatisticsBfExit instantiates a new SlurmV0041GetDiag200ResponseStatisticsBfExit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetDiag200ResponseStatisticsBfExitWithDefaults

`func NewSlurmV0041GetDiag200ResponseStatisticsBfExitWithDefaults() *SlurmV0041GetDiag200ResponseStatisticsBfExit`

NewSlurmV0041GetDiag200ResponseStatisticsBfExitWithDefaults instantiates a new SlurmV0041GetDiag200ResponseStatisticsBfExit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetEndJobQueue() int32`

GetEndJobQueue returns the EndJobQueue field if non-nil, zero value otherwise.

### GetEndJobQueueOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetEndJobQueueOk() (*int32, bool)`

GetEndJobQueueOk returns a tuple with the EndJobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetEndJobQueue(v int32)`

SetEndJobQueue sets EndJobQueue field to given value.

### HasEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasEndJobQueue() bool`

HasEndJobQueue returns a boolean if a field has been set.

### GetBfMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxJobStart() int32`

GetBfMaxJobStart returns the BfMaxJobStart field if non-nil, zero value otherwise.

### GetBfMaxJobStartOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxJobStartOk() (*int32, bool)`

GetBfMaxJobStartOk returns a tuple with the BfMaxJobStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetBfMaxJobStart(v int32)`

SetBfMaxJobStart sets BfMaxJobStart field to given value.

### HasBfMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasBfMaxJobStart() bool`

HasBfMaxJobStart returns a boolean if a field has been set.

### GetBfMaxJobTest

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxJobTest() int32`

GetBfMaxJobTest returns the BfMaxJobTest field if non-nil, zero value otherwise.

### GetBfMaxJobTestOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxJobTestOk() (*int32, bool)`

GetBfMaxJobTestOk returns a tuple with the BfMaxJobTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfMaxJobTest

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetBfMaxJobTest(v int32)`

SetBfMaxJobTest sets BfMaxJobTest field to given value.

### HasBfMaxJobTest

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasBfMaxJobTest() bool`

HasBfMaxJobTest returns a boolean if a field has been set.

### GetBfMaxTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxTime() int32`

GetBfMaxTime returns the BfMaxTime field if non-nil, zero value otherwise.

### GetBfMaxTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfMaxTimeOk() (*int32, bool)`

GetBfMaxTimeOk returns a tuple with the BfMaxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfMaxTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetBfMaxTime(v int32)`

SetBfMaxTime sets BfMaxTime field to given value.

### HasBfMaxTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasBfMaxTime() bool`

HasBfMaxTime returns a boolean if a field has been set.

### GetBfNodeSpaceSize

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfNodeSpaceSize() int32`

GetBfNodeSpaceSize returns the BfNodeSpaceSize field if non-nil, zero value otherwise.

### GetBfNodeSpaceSizeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetBfNodeSpaceSizeOk() (*int32, bool)`

GetBfNodeSpaceSizeOk returns a tuple with the BfNodeSpaceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfNodeSpaceSize

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetBfNodeSpaceSize(v int32)`

SetBfNodeSpaceSize sets BfNodeSpaceSize field to given value.

### HasBfNodeSpaceSize

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasBfNodeSpaceSize() bool`

HasBfNodeSpaceSize returns a boolean if a field has been set.

### GetStateChanged

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetStateChanged() int32`

GetStateChanged returns the StateChanged field if non-nil, zero value otherwise.

### GetStateChangedOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) GetStateChangedOk() (*int32, bool)`

GetStateChangedOk returns a tuple with the StateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanged

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) SetStateChanged(v int32)`

SetStateChanged sets StateChanged field to given value.

### HasStateChanged

`func (o *SlurmV0041GetDiag200ResponseStatisticsBfExit) HasStateChanged() bool`

HasStateChanged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


