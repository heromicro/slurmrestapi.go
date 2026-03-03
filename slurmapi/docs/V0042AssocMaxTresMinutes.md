# V0042AssocMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Per** | Pointer to [**V0042QosLimitsMinTresPer**](V0042QosLimitsMinTresPer.md) |  | [optional] 

## Methods

### NewV0042AssocMaxTresMinutes

`func NewV0042AssocMaxTresMinutes() *V0042AssocMaxTresMinutes`

NewV0042AssocMaxTresMinutes instantiates a new V0042AssocMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042AssocMaxTresMinutesWithDefaults

`func NewV0042AssocMaxTresMinutesWithDefaults() *V0042AssocMaxTresMinutes`

NewV0042AssocMaxTresMinutesWithDefaults instantiates a new V0042AssocMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0042AssocMaxTresMinutes) GetTotal() []V0042Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0042AssocMaxTresMinutes) GetTotalOk() (*[]V0042Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0042AssocMaxTresMinutes) SetTotal(v []V0042Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0042AssocMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0042AssocMaxTresMinutes) GetPer() V0042QosLimitsMinTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0042AssocMaxTresMinutes) GetPerOk() (*V0042QosLimitsMinTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0042AssocMaxTresMinutes) SetPer(v V0042QosLimitsMinTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0042AssocMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


