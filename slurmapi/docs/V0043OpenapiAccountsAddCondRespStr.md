# V0043OpenapiAccountsAddCondRespStr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAccounts** | **string** | added_accounts | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiAccountsAddCondRespStr

`func NewV0043OpenapiAccountsAddCondRespStr(addedAccounts string, ) *V0043OpenapiAccountsAddCondRespStr`

NewV0043OpenapiAccountsAddCondRespStr instantiates a new V0043OpenapiAccountsAddCondRespStr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiAccountsAddCondRespStrWithDefaults

`func NewV0043OpenapiAccountsAddCondRespStrWithDefaults() *V0043OpenapiAccountsAddCondRespStr`

NewV0043OpenapiAccountsAddCondRespStrWithDefaults instantiates a new V0043OpenapiAccountsAddCondRespStr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAccounts

`func (o *V0043OpenapiAccountsAddCondRespStr) GetAddedAccounts() string`

GetAddedAccounts returns the AddedAccounts field if non-nil, zero value otherwise.

### GetAddedAccountsOk

`func (o *V0043OpenapiAccountsAddCondRespStr) GetAddedAccountsOk() (*string, bool)`

GetAddedAccountsOk returns a tuple with the AddedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAccounts

`func (o *V0043OpenapiAccountsAddCondRespStr) SetAddedAccounts(v string)`

SetAddedAccounts sets AddedAccounts field to given value.


### GetMeta

`func (o *V0043OpenapiAccountsAddCondRespStr) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiAccountsAddCondRespStr) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiAccountsAddCondRespStr) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiAccountsAddCondRespStr) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiAccountsAddCondRespStr) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiAccountsAddCondRespStr) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiAccountsAddCondRespStr) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiAccountsAddCondRespStr) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiAccountsAddCondRespStr) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiAccountsAddCondRespStr) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiAccountsAddCondRespStr) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiAccountsAddCondRespStr) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


