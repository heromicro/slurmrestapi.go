# V0040JobReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of requested reservation | [optional] 
**Name** | Pointer to **string** | Name of reservation to use | [optional] 

## Methods

### NewV0040JobReservation

`func NewV0040JobReservation() *V0040JobReservation`

NewV0040JobReservation instantiates a new V0040JobReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobReservationWithDefaults

`func NewV0040JobReservationWithDefaults() *V0040JobReservation`

NewV0040JobReservationWithDefaults instantiates a new V0040JobReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0040JobReservation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0040JobReservation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0040JobReservation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0040JobReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0040JobReservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040JobReservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040JobReservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040JobReservation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


