# V0042OpenapiAccountsAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**V0042AccountsAddCond**](V0042AccountsAddCond.md) |  | [optional] 
**Account** | Pointer to [**V0042AccountShort**](V0042AccountShort.md) |  | [optional] 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiAccountsAddCondResp

`func NewV0042OpenapiAccountsAddCondResp() *V0042OpenapiAccountsAddCondResp`

NewV0042OpenapiAccountsAddCondResp instantiates a new V0042OpenapiAccountsAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiAccountsAddCondRespWithDefaults

`func NewV0042OpenapiAccountsAddCondRespWithDefaults() *V0042OpenapiAccountsAddCondResp`

NewV0042OpenapiAccountsAddCondRespWithDefaults instantiates a new V0042OpenapiAccountsAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0042OpenapiAccountsAddCondResp) GetAssociationCondition() V0042AccountsAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0042OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0042AccountsAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0042OpenapiAccountsAddCondResp) SetAssociationCondition(v V0042AccountsAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *V0042OpenapiAccountsAddCondResp) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetAccount

`func (o *V0042OpenapiAccountsAddCondResp) GetAccount() V0042AccountShort`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0042OpenapiAccountsAddCondResp) GetAccountOk() (*V0042AccountShort, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0042OpenapiAccountsAddCondResp) SetAccount(v V0042AccountShort)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0042OpenapiAccountsAddCondResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *V0042OpenapiAccountsAddCondResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiAccountsAddCondResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiAccountsAddCondResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiAccountsAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiAccountsAddCondResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiAccountsAddCondResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiAccountsAddCondResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiAccountsAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiAccountsAddCondResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiAccountsAddCondResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiAccountsAddCondResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiAccountsAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


