# SlurmV0041PostJobSubmitRequestJobsInnerCrontab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** | Flags | [optional] 
**Minute** | Pointer to **string** | Ranged string specifying eligible minute values (e.g. 0-10,50) | [optional] 
**Hour** | Pointer to **string** | Ranged string specifying eligible hour values (e.g. 0-5,23) | [optional] 
**DayOfMonth** | Pointer to **string** | Ranged string specifying eligible day of month values (e.g. 0-10,29) | [optional] 
**Month** | Pointer to **string** | Ranged string specifying eligible month values (e.g. 0-5,12) | [optional] 
**DayOfWeek** | Pointer to **string** | Ranged string specifying eligible day of week values (e.g.0-3,7) | [optional] 
**Specification** | Pointer to **string** | Time specification (* means valid for all allowed values) - minute hour day_of_month month day_of_week | [optional] 
**Command** | Pointer to **string** | Command to run | [optional] 
**Line** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerCrontabLine**](SlurmV0041PostJobSubmitRequestJobsInnerCrontabLine.md) |  | [optional] 

## Methods

### NewSlurmV0041PostJobSubmitRequestJobsInnerCrontab

`func NewSlurmV0041PostJobSubmitRequestJobsInnerCrontab() *SlurmV0041PostJobSubmitRequestJobsInnerCrontab`

NewSlurmV0041PostJobSubmitRequestJobsInnerCrontab instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerCrontab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmitRequestJobsInnerCrontabWithDefaults

`func NewSlurmV0041PostJobSubmitRequestJobsInnerCrontabWithDefaults() *SlurmV0041PostJobSubmitRequestJobsInnerCrontab`

NewSlurmV0041PostJobSubmitRequestJobsInnerCrontabWithDefaults instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerCrontab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMinute

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetMinute() string`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetMinuteOk() (*string, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetMinute(v string)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetSpecification() string`

GetSpecification returns the Specification field if non-nil, zero value otherwise.

### GetSpecificationOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetSpecificationOk() (*string, bool)`

GetSpecificationOk returns a tuple with the Specification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetSpecification(v string)`

SetSpecification sets Specification field to given value.

### HasSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasSpecification() bool`

HasSpecification returns a boolean if a field has been set.

### GetCommand

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetLine

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetLine() SlurmV0041PostJobSubmitRequestJobsInnerCrontabLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) GetLineOk() (*SlurmV0041PostJobSubmitRequestJobsInnerCrontabLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) SetLine(v SlurmV0041PostJobSubmitRequestJobsInnerCrontabLine)`

SetLine sets Line field to given value.

### HasLine

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerCrontab) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


