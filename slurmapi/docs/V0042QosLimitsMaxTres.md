# V0042QosLimitsMaxTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Minutes** | Pointer to [**V0042QosLimitsMaxTresMinutes**](V0042QosLimitsMaxTresMinutes.md) |  | [optional] 
**Per** | Pointer to [**V0042QosLimitsMaxTresPer**](V0042QosLimitsMaxTresPer.md) |  | [optional] 

## Methods

### NewV0042QosLimitsMaxTres

`func NewV0042QosLimitsMaxTres() *V0042QosLimitsMaxTres`

NewV0042QosLimitsMaxTres instantiates a new V0042QosLimitsMaxTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042QosLimitsMaxTresWithDefaults

`func NewV0042QosLimitsMaxTresWithDefaults() *V0042QosLimitsMaxTres`

NewV0042QosLimitsMaxTresWithDefaults instantiates a new V0042QosLimitsMaxTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0042QosLimitsMaxTres) GetTotal() []V0042Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0042QosLimitsMaxTres) GetTotalOk() (*[]V0042Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0042QosLimitsMaxTres) SetTotal(v []V0042Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0042QosLimitsMaxTres) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMinutes

`func (o *V0042QosLimitsMaxTres) GetMinutes() V0042QosLimitsMaxTresMinutes`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *V0042QosLimitsMaxTres) GetMinutesOk() (*V0042QosLimitsMaxTresMinutes, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *V0042QosLimitsMaxTres) SetMinutes(v V0042QosLimitsMaxTresMinutes)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *V0042QosLimitsMaxTres) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetPer

`func (o *V0042QosLimitsMaxTres) GetPer() V0042QosLimitsMaxTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0042QosLimitsMaxTres) GetPerOk() (*V0042QosLimitsMaxTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0042QosLimitsMaxTres) SetPer(v V0042QosLimitsMaxTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0042QosLimitsMaxTres) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


