# V0043QosLimitsMaxTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0043Tres**](V0043Tres.md) |  | [optional] 
**Minutes** | Pointer to [**V0043QosLimitsMaxTresMinutes**](V0043QosLimitsMaxTresMinutes.md) |  | [optional] 
**Per** | Pointer to [**V0043QosLimitsMaxTresPer**](V0043QosLimitsMaxTresPer.md) |  | [optional] 

## Methods

### NewV0043QosLimitsMaxTres

`func NewV0043QosLimitsMaxTres() *V0043QosLimitsMaxTres`

NewV0043QosLimitsMaxTres instantiates a new V0043QosLimitsMaxTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043QosLimitsMaxTresWithDefaults

`func NewV0043QosLimitsMaxTresWithDefaults() *V0043QosLimitsMaxTres`

NewV0043QosLimitsMaxTresWithDefaults instantiates a new V0043QosLimitsMaxTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0043QosLimitsMaxTres) GetTotal() []V0043Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0043QosLimitsMaxTres) GetTotalOk() (*[]V0043Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0043QosLimitsMaxTres) SetTotal(v []V0043Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0043QosLimitsMaxTres) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMinutes

`func (o *V0043QosLimitsMaxTres) GetMinutes() V0043QosLimitsMaxTresMinutes`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *V0043QosLimitsMaxTres) GetMinutesOk() (*V0043QosLimitsMaxTresMinutes, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *V0043QosLimitsMaxTres) SetMinutes(v V0043QosLimitsMaxTresMinutes)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *V0043QosLimitsMaxTres) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetPer

`func (o *V0043QosLimitsMaxTres) GetPer() V0043QosLimitsMaxTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0043QosLimitsMaxTres) GetPerOk() (*V0043QosLimitsMaxTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0043QosLimitsMaxTres) SetPer(v V0043QosLimitsMaxTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0043QosLimitsMaxTres) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


