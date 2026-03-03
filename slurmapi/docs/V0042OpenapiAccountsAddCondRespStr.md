# V0042OpenapiAccountsAddCondRespStr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAccounts** | **string** | added_accounts | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiAccountsAddCondRespStr

`func NewV0042OpenapiAccountsAddCondRespStr(addedAccounts string, ) *V0042OpenapiAccountsAddCondRespStr`

NewV0042OpenapiAccountsAddCondRespStr instantiates a new V0042OpenapiAccountsAddCondRespStr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiAccountsAddCondRespStrWithDefaults

`func NewV0042OpenapiAccountsAddCondRespStrWithDefaults() *V0042OpenapiAccountsAddCondRespStr`

NewV0042OpenapiAccountsAddCondRespStrWithDefaults instantiates a new V0042OpenapiAccountsAddCondRespStr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAccounts

`func (o *V0042OpenapiAccountsAddCondRespStr) GetAddedAccounts() string`

GetAddedAccounts returns the AddedAccounts field if non-nil, zero value otherwise.

### GetAddedAccountsOk

`func (o *V0042OpenapiAccountsAddCondRespStr) GetAddedAccountsOk() (*string, bool)`

GetAddedAccountsOk returns a tuple with the AddedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAccounts

`func (o *V0042OpenapiAccountsAddCondRespStr) SetAddedAccounts(v string)`

SetAddedAccounts sets AddedAccounts field to given value.


### GetMeta

`func (o *V0042OpenapiAccountsAddCondRespStr) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiAccountsAddCondRespStr) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiAccountsAddCondRespStr) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiAccountsAddCondRespStr) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiAccountsAddCondRespStr) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiAccountsAddCondRespStr) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiAccountsAddCondRespStr) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiAccountsAddCondRespStr) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiAccountsAddCondRespStr) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiAccountsAddCondRespStr) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiAccountsAddCondRespStr) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiAccountsAddCondRespStr) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


