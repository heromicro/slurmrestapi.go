# V0042QosLimitsMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Per** | Pointer to [**V0042QosLimitsMaxTresMinutesPer**](V0042QosLimitsMaxTresMinutesPer.md) |  | [optional] 

## Methods

### NewV0042QosLimitsMaxTresMinutes

`func NewV0042QosLimitsMaxTresMinutes() *V0042QosLimitsMaxTresMinutes`

NewV0042QosLimitsMaxTresMinutes instantiates a new V0042QosLimitsMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042QosLimitsMaxTresMinutesWithDefaults

`func NewV0042QosLimitsMaxTresMinutesWithDefaults() *V0042QosLimitsMaxTresMinutes`

NewV0042QosLimitsMaxTresMinutesWithDefaults instantiates a new V0042QosLimitsMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0042QosLimitsMaxTresMinutes) GetTotal() []V0042Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0042QosLimitsMaxTresMinutes) GetTotalOk() (*[]V0042Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0042QosLimitsMaxTresMinutes) SetTotal(v []V0042Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0042QosLimitsMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0042QosLimitsMaxTresMinutes) GetPer() V0042QosLimitsMaxTresMinutesPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0042QosLimitsMaxTresMinutes) GetPerOk() (*V0042QosLimitsMaxTresMinutesPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0042QosLimitsMaxTresMinutes) SetPer(v V0042QosLimitsMaxTresMinutesPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0042QosLimitsMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


