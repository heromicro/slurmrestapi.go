# V0043StepStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPU** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatisticsCPU**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatisticsCPU.md) |  | [optional] 
**Energy** | Pointer to [**V0043StepStatisticsEnergy**](V0043StepStatisticsEnergy.md) |  | [optional] 

## Methods

### NewV0043StepStatistics

`func NewV0043StepStatistics() *V0043StepStatistics`

NewV0043StepStatistics instantiates a new V0043StepStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StepStatisticsWithDefaults

`func NewV0043StepStatisticsWithDefaults() *V0043StepStatistics`

NewV0043StepStatisticsWithDefaults instantiates a new V0043StepStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPU

`func (o *V0043StepStatistics) GetCPU() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatisticsCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0043StepStatistics) GetCPUOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatisticsCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0043StepStatistics) SetCPU(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatisticsCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0043StepStatistics) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetEnergy

`func (o *V0043StepStatistics) GetEnergy() V0043StepStatisticsEnergy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *V0043StepStatistics) GetEnergyOk() (*V0043StepStatisticsEnergy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *V0043StepStatistics) SetEnergy(v V0043StepStatisticsEnergy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *V0043StepStatistics) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


