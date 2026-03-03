# SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of monthly rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time monthly rollup ran (UNIX timestamp) | [optional] 
**Duration** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyDuration**](SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyDuration.md) |  | [optional] 

## Methods

### NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly

`func NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly() *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly`

NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyWithDefaults

`func NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyWithDefaults() *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly`

NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetDuration() SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) GetDurationOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) SetDuration(v SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthlyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


