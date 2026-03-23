# V0043QosLimitsMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0043Tres**](V0043Tres.md) |  | [optional] 
**Per** | Pointer to [**V0043QosLimitsMaxTresMinutesPer**](V0043QosLimitsMaxTresMinutesPer.md) |  | [optional] 

## Methods

### NewV0043QosLimitsMaxTresMinutes

`func NewV0043QosLimitsMaxTresMinutes() *V0043QosLimitsMaxTresMinutes`

NewV0043QosLimitsMaxTresMinutes instantiates a new V0043QosLimitsMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043QosLimitsMaxTresMinutesWithDefaults

`func NewV0043QosLimitsMaxTresMinutesWithDefaults() *V0043QosLimitsMaxTresMinutes`

NewV0043QosLimitsMaxTresMinutesWithDefaults instantiates a new V0043QosLimitsMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0043QosLimitsMaxTresMinutes) GetTotal() []V0043Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0043QosLimitsMaxTresMinutes) GetTotalOk() (*[]V0043Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0043QosLimitsMaxTresMinutes) SetTotal(v []V0043Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0043QosLimitsMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0043QosLimitsMaxTresMinutes) GetPer() V0043QosLimitsMaxTresMinutesPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0043QosLimitsMaxTresMinutes) GetPerOk() (*V0043QosLimitsMaxTresMinutesPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0043QosLimitsMaxTresMinutes) SetPer(v V0043QosLimitsMaxTresMinutesPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0043QosLimitsMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


