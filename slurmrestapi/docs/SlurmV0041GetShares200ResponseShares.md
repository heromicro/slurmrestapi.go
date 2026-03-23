# SlurmV0041GetShares200ResponseShares

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]SlurmV0041GetShares200ResponseSharesSharesInner**](SlurmV0041GetShares200ResponseSharesSharesInner.md) | Association shares | [optional] 
**TotalShares** | Pointer to **int64** | Total number of shares | [optional] 

## Methods

### NewSlurmV0041GetShares200ResponseShares

`func NewSlurmV0041GetShares200ResponseShares() *SlurmV0041GetShares200ResponseShares`

NewSlurmV0041GetShares200ResponseShares instantiates a new SlurmV0041GetShares200ResponseShares object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseSharesWithDefaults

`func NewSlurmV0041GetShares200ResponseSharesWithDefaults() *SlurmV0041GetShares200ResponseShares`

NewSlurmV0041GetShares200ResponseSharesWithDefaults instantiates a new SlurmV0041GetShares200ResponseShares object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *SlurmV0041GetShares200ResponseShares) GetShares() []SlurmV0041GetShares200ResponseSharesSharesInner`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *SlurmV0041GetShares200ResponseShares) GetSharesOk() (*[]SlurmV0041GetShares200ResponseSharesSharesInner, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *SlurmV0041GetShares200ResponseShares) SetShares(v []SlurmV0041GetShares200ResponseSharesSharesInner)`

SetShares sets Shares field to given value.

### HasShares

`func (o *SlurmV0041GetShares200ResponseShares) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTotalShares

`func (o *SlurmV0041GetShares200ResponseShares) GetTotalShares() int64`

GetTotalShares returns the TotalShares field if non-nil, zero value otherwise.

### GetTotalSharesOk

`func (o *SlurmV0041GetShares200ResponseShares) GetTotalSharesOk() (*int64, bool)`

GetTotalSharesOk returns a tuple with the TotalShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShares

`func (o *SlurmV0041GetShares200ResponseShares) SetTotalShares(v int64)`

SetTotalShares sets TotalShares field to given value.

### HasTotalShares

`func (o *SlurmV0041GetShares200ResponseShares) HasTotalShares() bool`

HasTotalShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


