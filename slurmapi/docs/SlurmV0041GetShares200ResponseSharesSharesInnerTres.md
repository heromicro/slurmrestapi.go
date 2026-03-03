# SlurmV0041GetShares200ResponseSharesSharesInnerTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunSeconds** | Pointer to [**[]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner**](SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner.md) | Currently running tres-secs &#x3D; grp_used_tres_run_secs | [optional] 
**GroupMinutes** | Pointer to [**[]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner**](SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner.md) | TRES-minute limit | [optional] 
**Usage** | Pointer to [**[]SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner**](SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner.md) | Measure of each TRES usage | [optional] 

## Methods

### NewSlurmV0041GetShares200ResponseSharesSharesInnerTres

`func NewSlurmV0041GetShares200ResponseSharesSharesInnerTres() *SlurmV0041GetShares200ResponseSharesSharesInnerTres`

NewSlurmV0041GetShares200ResponseSharesSharesInnerTres instantiates a new SlurmV0041GetShares200ResponseSharesSharesInnerTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseSharesSharesInnerTresWithDefaults

`func NewSlurmV0041GetShares200ResponseSharesSharesInnerTresWithDefaults() *SlurmV0041GetShares200ResponseSharesSharesInnerTres`

NewSlurmV0041GetShares200ResponseSharesSharesInnerTresWithDefaults instantiates a new SlurmV0041GetShares200ResponseSharesSharesInnerTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunSeconds

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetRunSeconds() []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner`

GetRunSeconds returns the RunSeconds field if non-nil, zero value otherwise.

### GetRunSecondsOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetRunSecondsOk() (*[]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner, bool)`

GetRunSecondsOk returns a tuple with the RunSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSeconds

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetRunSeconds(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner)`

SetRunSeconds sets RunSeconds field to given value.

### HasRunSeconds

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasRunSeconds() bool`

HasRunSeconds returns a boolean if a field has been set.

### GetGroupMinutes

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetGroupMinutes() []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner`

GetGroupMinutes returns the GroupMinutes field if non-nil, zero value otherwise.

### GetGroupMinutesOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetGroupMinutesOk() (*[]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner, bool)`

GetGroupMinutesOk returns a tuple with the GroupMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMinutes

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetGroupMinutes(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner)`

SetGroupMinutes sets GroupMinutes field to given value.

### HasGroupMinutes

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasGroupMinutes() bool`

HasGroupMinutes returns a boolean if a field has been set.

### GetUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetUsage() []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetUsageOk() (*[]SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetUsage(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


