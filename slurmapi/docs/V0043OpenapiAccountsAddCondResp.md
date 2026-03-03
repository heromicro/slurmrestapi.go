# V0043OpenapiAccountsAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**V0043AccountsAddCond**](V0043AccountsAddCond.md) |  | [optional] 
**Account** | Pointer to [**V0043AccountShort**](V0043AccountShort.md) |  | [optional] 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiAccountsAddCondResp

`func NewV0043OpenapiAccountsAddCondResp() *V0043OpenapiAccountsAddCondResp`

NewV0043OpenapiAccountsAddCondResp instantiates a new V0043OpenapiAccountsAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiAccountsAddCondRespWithDefaults

`func NewV0043OpenapiAccountsAddCondRespWithDefaults() *V0043OpenapiAccountsAddCondResp`

NewV0043OpenapiAccountsAddCondRespWithDefaults instantiates a new V0043OpenapiAccountsAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0043OpenapiAccountsAddCondResp) GetAssociationCondition() V0043AccountsAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0043OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0043AccountsAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0043OpenapiAccountsAddCondResp) SetAssociationCondition(v V0043AccountsAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *V0043OpenapiAccountsAddCondResp) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetAccount

`func (o *V0043OpenapiAccountsAddCondResp) GetAccount() V0043AccountShort`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0043OpenapiAccountsAddCondResp) GetAccountOk() (*V0043AccountShort, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0043OpenapiAccountsAddCondResp) SetAccount(v V0043AccountShort)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0043OpenapiAccountsAddCondResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *V0043OpenapiAccountsAddCondResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiAccountsAddCondResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiAccountsAddCondResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiAccountsAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiAccountsAddCondResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiAccountsAddCondResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiAccountsAddCondResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiAccountsAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiAccountsAddCondResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiAccountsAddCondResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiAccountsAddCondResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiAccountsAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


