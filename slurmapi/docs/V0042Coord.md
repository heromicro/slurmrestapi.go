# V0042Coord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User name | 
**Direct** | Pointer to **bool** | Indicates whether the coordinator was directly assigned to this account | [optional] 

## Methods

### NewV0042Coord

`func NewV0042Coord(name string, ) *V0042Coord`

NewV0042Coord instantiates a new V0042Coord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042CoordWithDefaults

`func NewV0042CoordWithDefaults() *V0042Coord`

NewV0042CoordWithDefaults instantiates a new V0042Coord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V0042Coord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042Coord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042Coord) SetName(v string)`

SetName sets Name field to given value.


### GetDirect

`func (o *V0042Coord) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *V0042Coord) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *V0042Coord) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *V0042Coord) HasDirect() bool`

HasDirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


