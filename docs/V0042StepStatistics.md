# V0042StepStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPU** | Pointer to [**V0040StepStatisticsCPU**](V0040StepStatisticsCPU.md) |  | [optional] 
**Energy** | Pointer to [**V0042StepStatisticsEnergy**](V0042StepStatisticsEnergy.md) |  | [optional] 

## Methods

### NewV0042StepStatistics

`func NewV0042StepStatistics() *V0042StepStatistics`

NewV0042StepStatistics instantiates a new V0042StepStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StepStatisticsWithDefaults

`func NewV0042StepStatisticsWithDefaults() *V0042StepStatistics`

NewV0042StepStatisticsWithDefaults instantiates a new V0042StepStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPU

`func (o *V0042StepStatistics) GetCPU() V0040StepStatisticsCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0042StepStatistics) GetCPUOk() (*V0040StepStatisticsCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0042StepStatistics) SetCPU(v V0040StepStatisticsCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0042StepStatistics) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetEnergy

`func (o *V0042StepStatistics) GetEnergy() V0042StepStatisticsEnergy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *V0042StepStatistics) GetEnergyOk() (*V0042StepStatisticsEnergy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *V0042StepStatistics) SetEnergy(v V0042StepStatisticsEnergy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *V0042StepStatistics) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


