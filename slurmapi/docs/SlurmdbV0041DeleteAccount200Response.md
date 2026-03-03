# SlurmdbV0041DeleteAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedAccounts** | **[]string** | removed_accounts | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmdbV0041DeleteAccount200Response

`func NewSlurmdbV0041DeleteAccount200Response(removedAccounts []string, ) *SlurmdbV0041DeleteAccount200Response`

NewSlurmdbV0041DeleteAccount200Response instantiates a new SlurmdbV0041DeleteAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041DeleteAccount200ResponseWithDefaults

`func NewSlurmdbV0041DeleteAccount200ResponseWithDefaults() *SlurmdbV0041DeleteAccount200Response`

NewSlurmdbV0041DeleteAccount200ResponseWithDefaults instantiates a new SlurmdbV0041DeleteAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedAccounts

`func (o *SlurmdbV0041DeleteAccount200Response) GetRemovedAccounts() []string`

GetRemovedAccounts returns the RemovedAccounts field if non-nil, zero value otherwise.

### GetRemovedAccountsOk

`func (o *SlurmdbV0041DeleteAccount200Response) GetRemovedAccountsOk() (*[]string, bool)`

GetRemovedAccountsOk returns a tuple with the RemovedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAccounts

`func (o *SlurmdbV0041DeleteAccount200Response) SetRemovedAccounts(v []string)`

SetRemovedAccounts sets RemovedAccounts field to given value.


### GetMeta

`func (o *SlurmdbV0041DeleteAccount200Response) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmdbV0041DeleteAccount200Response) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmdbV0041DeleteAccount200Response) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmdbV0041DeleteAccount200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmdbV0041DeleteAccount200Response) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmdbV0041DeleteAccount200Response) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmdbV0041DeleteAccount200Response) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmdbV0041DeleteAccount200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmdbV0041DeleteAccount200Response) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmdbV0041DeleteAccount200Response) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmdbV0041DeleteAccount200Response) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmdbV0041DeleteAccount200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


