# V0042RollupStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hourly** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourly**](SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourly.md) |  | [optional] 
**Daily** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRollupsDaily**](SlurmdbV0041GetDiag200ResponseStatisticsRollupsDaily.md) |  | [optional] 
**Monthly** | Pointer to [**SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly**](SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly.md) |  | [optional] 

## Methods

### NewV0042RollupStats

`func NewV0042RollupStats() *V0042RollupStats`

NewV0042RollupStats instantiates a new V0042RollupStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042RollupStatsWithDefaults

`func NewV0042RollupStatsWithDefaults() *V0042RollupStats`

NewV0042RollupStatsWithDefaults instantiates a new V0042RollupStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourly

`func (o *V0042RollupStats) GetHourly() SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourly`

GetHourly returns the Hourly field if non-nil, zero value otherwise.

### GetHourlyOk

`func (o *V0042RollupStats) GetHourlyOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourly, bool)`

GetHourlyOk returns a tuple with the Hourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourly

`func (o *V0042RollupStats) SetHourly(v SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourly)`

SetHourly sets Hourly field to given value.

### HasHourly

`func (o *V0042RollupStats) HasHourly() bool`

HasHourly returns a boolean if a field has been set.

### GetDaily

`func (o *V0042RollupStats) GetDaily() SlurmdbV0041GetDiag200ResponseStatisticsRollupsDaily`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *V0042RollupStats) GetDailyOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollupsDaily, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *V0042RollupStats) SetDaily(v SlurmdbV0041GetDiag200ResponseStatisticsRollupsDaily)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *V0042RollupStats) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetMonthly

`func (o *V0042RollupStats) GetMonthly() SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *V0042RollupStats) GetMonthlyOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *V0042RollupStats) SetMonthly(v SlurmdbV0041GetDiag200ResponseStatisticsRollupsMonthly)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *V0042RollupStats) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


