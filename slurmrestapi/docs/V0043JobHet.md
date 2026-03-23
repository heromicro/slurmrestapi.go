# V0043JobHet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** | Heterogeneous job ID, if applicable | [optional] 
**JobOffset** | Pointer to [**V0043Uint32NoValStruct**](V0043Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0043JobHet

`func NewV0043JobHet() *V0043JobHet`

NewV0043JobHet instantiates a new V0043JobHet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobHetWithDefaults

`func NewV0043JobHetWithDefaults() *V0043JobHet`

NewV0043JobHetWithDefaults instantiates a new V0043JobHet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0043JobHet) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0043JobHet) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0043JobHet) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0043JobHet) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobOffset

`func (o *V0043JobHet) GetJobOffset() V0043Uint32NoValStruct`

GetJobOffset returns the JobOffset field if non-nil, zero value otherwise.

### GetJobOffsetOk

`func (o *V0043JobHet) GetJobOffsetOk() (*V0043Uint32NoValStruct, bool)`

GetJobOffsetOk returns a tuple with the JobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOffset

`func (o *V0043JobHet) SetJobOffset(v V0043Uint32NoValStruct)`

SetJobOffset sets JobOffset field to given value.

### HasJobOffset

`func (o *V0043JobHet) HasJobOffset() bool`

HasJobOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


