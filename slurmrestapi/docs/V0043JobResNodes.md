# V0043JobResNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of allocated nodes | [optional] 
**SelectType** | Pointer to **[]string** | Node scheduling selection method | [optional] 
**List** | Pointer to **string** | Node(s) allocated to the job | [optional] 
**Whole** | Pointer to **bool** | Whether whole nodes were allocated | [optional] 
**Allocation** | Pointer to [**[]V0043JobResNode**](V0043JobResNode.md) | Job resources for a node | [optional] 

## Methods

### NewV0043JobResNodes

`func NewV0043JobResNodes() *V0043JobResNodes`

NewV0043JobResNodes instantiates a new V0043JobResNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobResNodesWithDefaults

`func NewV0043JobResNodesWithDefaults() *V0043JobResNodes`

NewV0043JobResNodesWithDefaults instantiates a new V0043JobResNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0043JobResNodes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0043JobResNodes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0043JobResNodes) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0043JobResNodes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSelectType

`func (o *V0043JobResNodes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0043JobResNodes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0043JobResNodes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0043JobResNodes) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetList

`func (o *V0043JobResNodes) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0043JobResNodes) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0043JobResNodes) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *V0043JobResNodes) HasList() bool`

HasList returns a boolean if a field has been set.

### GetWhole

`func (o *V0043JobResNodes) GetWhole() bool`

GetWhole returns the Whole field if non-nil, zero value otherwise.

### GetWholeOk

`func (o *V0043JobResNodes) GetWholeOk() (*bool, bool)`

GetWholeOk returns a tuple with the Whole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhole

`func (o *V0043JobResNodes) SetWhole(v bool)`

SetWhole sets Whole field to given value.

### HasWhole

`func (o *V0043JobResNodes) HasWhole() bool`

HasWhole returns a boolean if a field has been set.

### GetAllocation

`func (o *V0043JobResNodes) GetAllocation() []V0043JobResNode`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *V0043JobResNodes) GetAllocationOk() (*[]V0043JobResNode, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *V0043JobResNodes) SetAllocation(v []V0043JobResNode)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *V0043JobResNodes) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


