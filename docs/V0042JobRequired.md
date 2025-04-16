# V0042JobRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPUs** | Pointer to **int32** | Minimum number of CPUs required | [optional] 
**MemoryPerCpu** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 

## Methods

### NewV0042JobRequired

`func NewV0042JobRequired() *V0042JobRequired`

NewV0042JobRequired instantiates a new V0042JobRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobRequiredWithDefaults

`func NewV0042JobRequiredWithDefaults() *V0042JobRequired`

NewV0042JobRequiredWithDefaults instantiates a new V0042JobRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPUs

`func (o *V0042JobRequired) GetCPUs() int32`

GetCPUs returns the CPUs field if non-nil, zero value otherwise.

### GetCPUsOk

`func (o *V0042JobRequired) GetCPUsOk() (*int32, bool)`

GetCPUsOk returns a tuple with the CPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPUs

`func (o *V0042JobRequired) SetCPUs(v int32)`

SetCPUs sets CPUs field to given value.

### HasCPUs

`func (o *V0042JobRequired) HasCPUs() bool`

HasCPUs returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0042JobRequired) GetMemoryPerCpu() V0042Uint64NoValStruct`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0042JobRequired) GetMemoryPerCpuOk() (*V0042Uint64NoValStruct, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0042JobRequired) SetMemoryPerCpu(v V0042Uint64NoValStruct)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0042JobRequired) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0042JobRequired) GetMemoryPerNode() V0042Uint64NoValStruct`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0042JobRequired) GetMemoryPerNodeOk() (*V0042Uint64NoValStruct, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0042JobRequired) SetMemoryPerNode(v V0042Uint64NoValStruct)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0042JobRequired) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


