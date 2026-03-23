# V0042JobResNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Node index | 
**Name** | **string** | Node name | 
**Cpus** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus.md) |  | [optional] 
**Memory** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory.md) |  | [optional] 
**Sockets** | [**[]V0042JobResSocket**](V0042JobResSocket.md) |  | 

## Methods

### NewV0042JobResNode

`func NewV0042JobResNode(index int32, name string, sockets []V0042JobResSocket, ) *V0042JobResNode`

NewV0042JobResNode instantiates a new V0042JobResNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobResNodeWithDefaults

`func NewV0042JobResNodeWithDefaults() *V0042JobResNode`

NewV0042JobResNodeWithDefaults instantiates a new V0042JobResNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *V0042JobResNode) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V0042JobResNode) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V0042JobResNode) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *V0042JobResNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042JobResNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042JobResNode) SetName(v string)`

SetName sets Name field to given value.


### GetCpus

`func (o *V0042JobResNode) GetCpus() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0042JobResNode) GetCpusOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0042JobResNode) SetCpus(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0042JobResNode) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *V0042JobResNode) GetMemory() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0042JobResNode) GetMemoryOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0042JobResNode) SetMemory(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0042JobResNode) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSockets

`func (o *V0042JobResNode) GetSockets() []V0042JobResSocket`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0042JobResNode) GetSocketsOk() (*[]V0042JobResSocket, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0042JobResNode) SetSockets(v []V0042JobResSocket)`

SetSockets sets Sockets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


