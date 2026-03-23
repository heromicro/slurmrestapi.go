# SlurmdbV0041GetDiag200ResponseStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** | When data collection started (UNIX timestamp) | [optional] 
**Rollups** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRollups**](SlurmdbV0041GetDiag200ResponseStatisticsRollups.md) |  | [optional] 
**RPCs** | Pointer to [**[]SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner**](SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner.md) | List of RPCs sent to the slurmdbd | [optional] 
**Users** | Pointer to [**[]SlurmdbV0041GetDiag200ResponseStatisticsUsersInner**](SlurmdbV0041GetDiag200ResponseStatisticsUsersInner.md) | List of users that issued RPCs | [optional] 

## Methods

### NewSlurmdbV0041GetDiag200ResponseStatistics

`func NewSlurmdbV0041GetDiag200ResponseStatistics() *SlurmdbV0041GetDiag200ResponseStatistics`

NewSlurmdbV0041GetDiag200ResponseStatistics instantiates a new SlurmdbV0041GetDiag200ResponseStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041GetDiag200ResponseStatisticsWithDefaults

`func NewSlurmdbV0041GetDiag200ResponseStatisticsWithDefaults() *SlurmdbV0041GetDiag200ResponseStatistics`

NewSlurmdbV0041GetDiag200ResponseStatisticsWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRollups() SlurmdbV0041GetDiag200ResponseStatisticsRollups`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRollupsOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollups, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetRollups(v SlurmdbV0041GetDiag200ResponseStatisticsRollups)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRPCs() []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRPCsOk() (*[]SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetRPCs(v []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetUsers() []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetUsersOk() (*[]SlurmdbV0041GetDiag200ResponseStatisticsUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetUsers(v []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


