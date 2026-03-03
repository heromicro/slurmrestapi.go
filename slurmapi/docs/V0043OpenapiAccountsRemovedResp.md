# V0043OpenapiAccountsRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedAccounts** | **[]string** |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiAccountsRemovedResp

`func NewV0043OpenapiAccountsRemovedResp(removedAccounts []string, ) *V0043OpenapiAccountsRemovedResp`

NewV0043OpenapiAccountsRemovedResp instantiates a new V0043OpenapiAccountsRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiAccountsRemovedRespWithDefaults

`func NewV0043OpenapiAccountsRemovedRespWithDefaults() *V0043OpenapiAccountsRemovedResp`

NewV0043OpenapiAccountsRemovedRespWithDefaults instantiates a new V0043OpenapiAccountsRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedAccounts

`func (o *V0043OpenapiAccountsRemovedResp) GetRemovedAccounts() []string`

GetRemovedAccounts returns the RemovedAccounts field if non-nil, zero value otherwise.

### GetRemovedAccountsOk

`func (o *V0043OpenapiAccountsRemovedResp) GetRemovedAccountsOk() (*[]string, bool)`

GetRemovedAccountsOk returns a tuple with the RemovedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAccounts

`func (o *V0043OpenapiAccountsRemovedResp) SetRemovedAccounts(v []string)`

SetRemovedAccounts sets RemovedAccounts field to given value.


### GetMeta

`func (o *V0043OpenapiAccountsRemovedResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiAccountsRemovedResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiAccountsRemovedResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiAccountsRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiAccountsRemovedResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiAccountsRemovedResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiAccountsRemovedResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiAccountsRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiAccountsRemovedResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiAccountsRemovedResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiAccountsRemovedResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiAccountsRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


