# V0043AssocMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0043Tres**](V0043Tres.md) |  | [optional] 
**Per** | Pointer to [**V0043QosLimitsMinTresPer**](V0043QosLimitsMinTresPer.md) |  | [optional] 

## Methods

### NewV0043AssocMaxTresMinutes

`func NewV0043AssocMaxTresMinutes() *V0043AssocMaxTresMinutes`

NewV0043AssocMaxTresMinutes instantiates a new V0043AssocMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043AssocMaxTresMinutesWithDefaults

`func NewV0043AssocMaxTresMinutesWithDefaults() *V0043AssocMaxTresMinutes`

NewV0043AssocMaxTresMinutesWithDefaults instantiates a new V0043AssocMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0043AssocMaxTresMinutes) GetTotal() []V0043Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0043AssocMaxTresMinutes) GetTotalOk() (*[]V0043Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0043AssocMaxTresMinutes) SetTotal(v []V0043Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0043AssocMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0043AssocMaxTresMinutes) GetPer() V0043QosLimitsMinTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0043AssocMaxTresMinutes) GetPerOk() (*V0043QosLimitsMinTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0043AssocMaxTresMinutes) SetPer(v V0043QosLimitsMinTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0043AssocMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


