# SlurmV0041GetDiag200ResponseStatisticsScheduleExit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndJobQueue** | Pointer to **int32** | Reached end of queue | [optional] 
**DefaultQueueDepth** | Pointer to **int32** | Reached number of jobs allowed to be tested | [optional] 
**MaxJobStart** | Pointer to **int32** | Reached number of jobs allowed to start | [optional] 
**MaxRpcCnt** | Pointer to **int32** | Reached RPC limit | [optional] 
**MaxSchedTime** | Pointer to **int32** | Reached maximum allowed scheduler time | [optional] 
**Licenses** | Pointer to **int32** | Blocked on licenses | [optional] 

## Methods

### NewSlurmV0041GetDiag200ResponseStatisticsScheduleExit

`func NewSlurmV0041GetDiag200ResponseStatisticsScheduleExit() *SlurmV0041GetDiag200ResponseStatisticsScheduleExit`

NewSlurmV0041GetDiag200ResponseStatisticsScheduleExit instantiates a new SlurmV0041GetDiag200ResponseStatisticsScheduleExit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetDiag200ResponseStatisticsScheduleExitWithDefaults

`func NewSlurmV0041GetDiag200ResponseStatisticsScheduleExitWithDefaults() *SlurmV0041GetDiag200ResponseStatisticsScheduleExit`

NewSlurmV0041GetDiag200ResponseStatisticsScheduleExitWithDefaults instantiates a new SlurmV0041GetDiag200ResponseStatisticsScheduleExit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetEndJobQueue() int32`

GetEndJobQueue returns the EndJobQueue field if non-nil, zero value otherwise.

### GetEndJobQueueOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetEndJobQueueOk() (*int32, bool)`

GetEndJobQueueOk returns a tuple with the EndJobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetEndJobQueue(v int32)`

SetEndJobQueue sets EndJobQueue field to given value.

### HasEndJobQueue

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasEndJobQueue() bool`

HasEndJobQueue returns a boolean if a field has been set.

### GetDefaultQueueDepth

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetDefaultQueueDepth() int32`

GetDefaultQueueDepth returns the DefaultQueueDepth field if non-nil, zero value otherwise.

### GetDefaultQueueDepthOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetDefaultQueueDepthOk() (*int32, bool)`

GetDefaultQueueDepthOk returns a tuple with the DefaultQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQueueDepth

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetDefaultQueueDepth(v int32)`

SetDefaultQueueDepth sets DefaultQueueDepth field to given value.

### HasDefaultQueueDepth

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasDefaultQueueDepth() bool`

HasDefaultQueueDepth returns a boolean if a field has been set.

### GetMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxJobStart() int32`

GetMaxJobStart returns the MaxJobStart field if non-nil, zero value otherwise.

### GetMaxJobStartOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxJobStartOk() (*int32, bool)`

GetMaxJobStartOk returns a tuple with the MaxJobStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetMaxJobStart(v int32)`

SetMaxJobStart sets MaxJobStart field to given value.

### HasMaxJobStart

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasMaxJobStart() bool`

HasMaxJobStart returns a boolean if a field has been set.

### GetMaxRpcCnt

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxRpcCnt() int32`

GetMaxRpcCnt returns the MaxRpcCnt field if non-nil, zero value otherwise.

### GetMaxRpcCntOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxRpcCntOk() (*int32, bool)`

GetMaxRpcCntOk returns a tuple with the MaxRpcCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRpcCnt

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetMaxRpcCnt(v int32)`

SetMaxRpcCnt sets MaxRpcCnt field to given value.

### HasMaxRpcCnt

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasMaxRpcCnt() bool`

HasMaxRpcCnt returns a boolean if a field has been set.

### GetMaxSchedTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxSchedTime() int32`

GetMaxSchedTime returns the MaxSchedTime field if non-nil, zero value otherwise.

### GetMaxSchedTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetMaxSchedTimeOk() (*int32, bool)`

GetMaxSchedTimeOk returns a tuple with the MaxSchedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSchedTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetMaxSchedTime(v int32)`

SetMaxSchedTime sets MaxSchedTime field to given value.

### HasMaxSchedTime

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasMaxSchedTime() bool`

HasMaxSchedTime returns a boolean if a field has been set.

### GetLicenses

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetLicenses() int32`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) GetLicensesOk() (*int32, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) SetLicenses(v int32)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *SlurmV0041GetDiag200ResponseStatisticsScheduleExit) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


