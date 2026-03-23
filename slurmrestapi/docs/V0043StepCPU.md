# V0043StepCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedFrequency** | Pointer to [**V0043StepCPURequestedFrequency**](V0043StepCPURequestedFrequency.md) |  | [optional] 
**Governor** | Pointer to **string** | Requested CPU frequency governor in kHz | [optional] 

## Methods

### NewV0043StepCPU

`func NewV0043StepCPU() *V0043StepCPU`

NewV0043StepCPU instantiates a new V0043StepCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StepCPUWithDefaults

`func NewV0043StepCPUWithDefaults() *V0043StepCPU`

NewV0043StepCPUWithDefaults instantiates a new V0043StepCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedFrequency

`func (o *V0043StepCPU) GetRequestedFrequency() V0043StepCPURequestedFrequency`

GetRequestedFrequency returns the RequestedFrequency field if non-nil, zero value otherwise.

### GetRequestedFrequencyOk

`func (o *V0043StepCPU) GetRequestedFrequencyOk() (*V0043StepCPURequestedFrequency, bool)`

GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFrequency

`func (o *V0043StepCPU) SetRequestedFrequency(v V0043StepCPURequestedFrequency)`

SetRequestedFrequency sets RequestedFrequency field to given value.

### HasRequestedFrequency

`func (o *V0043StepCPU) HasRequestedFrequency() bool`

HasRequestedFrequency returns a boolean if a field has been set.

### GetGovernor

`func (o *V0043StepCPU) GetGovernor() string`

GetGovernor returns the Governor field if non-nil, zero value otherwise.

### GetGovernorOk

`func (o *V0043StepCPU) GetGovernorOk() (*string, bool)`

GetGovernorOk returns a tuple with the Governor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernor

`func (o *V0043StepCPU) SetGovernor(v string)`

SetGovernor sets Governor field to given value.

### HasGovernor

`func (o *V0043StepCPU) HasGovernor() bool`

HasGovernor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


