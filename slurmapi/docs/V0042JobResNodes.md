# V0042JobResNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of allocated nodes | [optional] 
**SelectType** | Pointer to **[]string** |  | [optional] 
**List** | Pointer to **string** | Node(s) allocated to the job | [optional] 
**Whole** | Pointer to **bool** | Whether whole nodes were allocated | [optional] 
**Allocation** | Pointer to [**[]V0042JobResNode**](V0042JobResNode.md) | Job resources for a node | [optional] 

## Methods

### NewV0042JobResNodes

`func NewV0042JobResNodes() *V0042JobResNodes`

NewV0042JobResNodes instantiates a new V0042JobResNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobResNodesWithDefaults

`func NewV0042JobResNodesWithDefaults() *V0042JobResNodes`

NewV0042JobResNodesWithDefaults instantiates a new V0042JobResNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0042JobResNodes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042JobResNodes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042JobResNodes) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0042JobResNodes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSelectType

`func (o *V0042JobResNodes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0042JobResNodes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0042JobResNodes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0042JobResNodes) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetList

`func (o *V0042JobResNodes) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0042JobResNodes) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0042JobResNodes) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *V0042JobResNodes) HasList() bool`

HasList returns a boolean if a field has been set.

### GetWhole

`func (o *V0042JobResNodes) GetWhole() bool`

GetWhole returns the Whole field if non-nil, zero value otherwise.

### GetWholeOk

`func (o *V0042JobResNodes) GetWholeOk() (*bool, bool)`

GetWholeOk returns a tuple with the Whole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhole

`func (o *V0042JobResNodes) SetWhole(v bool)`

SetWhole sets Whole field to given value.

### HasWhole

`func (o *V0042JobResNodes) HasWhole() bool`

HasWhole returns a boolean if a field has been set.

### GetAllocation

`func (o *V0042JobResNodes) GetAllocation() []V0042JobResNode`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *V0042JobResNodes) GetAllocationOk() (*[]V0042JobResNode, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *V0042JobResNodes) SetAllocation(v []V0042JobResNode)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *V0042JobResNodes) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


