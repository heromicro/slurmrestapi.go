# V0042Tres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | TRES type (CPU, MEM, etc) | 
**Name** | Pointer to **string** | TRES name (if applicable) | [optional] 
**Id** | Pointer to **int32** | ID used in the database | [optional] 
**Count** | Pointer to **int64** | TRES count (0 if listed generically) | [optional] 

## Methods

### NewV0042Tres

`func NewV0042Tres(type_ string, ) *V0042Tres`

NewV0042Tres instantiates a new V0042Tres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042TresWithDefaults

`func NewV0042TresWithDefaults() *V0042Tres`

NewV0042TresWithDefaults instantiates a new V0042Tres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V0042Tres) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0042Tres) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0042Tres) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *V0042Tres) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042Tres) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042Tres) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0042Tres) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *V0042Tres) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0042Tres) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0042Tres) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0042Tres) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *V0042Tres) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042Tres) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042Tres) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0042Tres) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


