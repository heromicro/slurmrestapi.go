# V0043JobRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectType** | **[]string** | Scheduler consumable resource selection type | 
**Nodes** | Pointer to [**V0043JobResNodes**](V0043JobResNodes.md) |  | [optional] 
**Cpus** | **int32** | Number of allocated CPUs | 
**ThreadsPerCore** | [**V0043Uint16NoValStruct**](V0043Uint16NoValStruct.md) |  | 

## Methods

### NewV0043JobRes

`func NewV0043JobRes(selectType []string, cpus int32, threadsPerCore V0043Uint16NoValStruct, ) *V0043JobRes`

NewV0043JobRes instantiates a new V0043JobRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobResWithDefaults

`func NewV0043JobResWithDefaults() *V0043JobRes`

NewV0043JobResWithDefaults instantiates a new V0043JobRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectType

`func (o *V0043JobRes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0043JobRes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0043JobRes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.


### GetNodes

`func (o *V0043JobRes) GetNodes() V0043JobResNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0043JobRes) GetNodesOk() (*V0043JobResNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0043JobRes) SetNodes(v V0043JobResNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0043JobRes) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCpus

`func (o *V0043JobRes) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0043JobRes) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0043JobRes) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetThreadsPerCore

`func (o *V0043JobRes) GetThreadsPerCore() V0043Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0043JobRes) GetThreadsPerCoreOk() (*V0043Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0043JobRes) SetThreadsPerCore(v V0043Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


