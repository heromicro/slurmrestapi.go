# V0042JobRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectType** | **[]string** |  | 
**Nodes** | Pointer to [**V0042JobResNodes**](V0042JobResNodes.md) |  | [optional] 
**Cpus** | **int32** | Number of allocated CPUs | 
**ThreadsPerCore** | [**V0042Uint16NoValStruct**](V0042Uint16NoValStruct.md) |  | 

## Methods

### NewV0042JobRes

`func NewV0042JobRes(selectType []string, cpus int32, threadsPerCore V0042Uint16NoValStruct, ) *V0042JobRes`

NewV0042JobRes instantiates a new V0042JobRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobResWithDefaults

`func NewV0042JobResWithDefaults() *V0042JobRes`

NewV0042JobResWithDefaults instantiates a new V0042JobRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectType

`func (o *V0042JobRes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0042JobRes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0042JobRes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.


### GetNodes

`func (o *V0042JobRes) GetNodes() V0042JobResNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042JobRes) GetNodesOk() (*V0042JobResNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042JobRes) SetNodes(v V0042JobResNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042JobRes) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCpus

`func (o *V0042JobRes) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0042JobRes) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0042JobRes) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetThreadsPerCore

`func (o *V0042JobRes) GetThreadsPerCore() V0042Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0042JobRes) GetThreadsPerCoreOk() (*V0042Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0042JobRes) SetThreadsPerCore(v V0042Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


