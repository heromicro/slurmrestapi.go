# SlurmdbV0041PostAccountsAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**SlurmdbV0041PostAccountsAssociationRequestAssociationCondition**](SlurmdbV0041PostAccountsAssociationRequestAssociationCondition.md) |  | [optional] 
**Account** | Pointer to [**SlurmdbV0041PostAccountsAssociationRequestAccount**](SlurmdbV0041PostAccountsAssociationRequestAccount.md) |  | [optional] 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmdbV0041PostAccountsAssociationRequest

`func NewSlurmdbV0041PostAccountsAssociationRequest() *SlurmdbV0041PostAccountsAssociationRequest`

NewSlurmdbV0041PostAccountsAssociationRequest instantiates a new SlurmdbV0041PostAccountsAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostAccountsAssociationRequestWithDefaults

`func NewSlurmdbV0041PostAccountsAssociationRequestWithDefaults() *SlurmdbV0041PostAccountsAssociationRequest`

NewSlurmdbV0041PostAccountsAssociationRequestWithDefaults instantiates a new SlurmdbV0041PostAccountsAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAssociationCondition() SlurmdbV0041PostAccountsAssociationRequestAssociationCondition`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAssociationConditionOk() (*SlurmdbV0041PostAccountsAssociationRequestAssociationCondition, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *SlurmdbV0041PostAccountsAssociationRequest) SetAssociationCondition(v SlurmdbV0041PostAccountsAssociationRequestAssociationCondition)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *SlurmdbV0041PostAccountsAssociationRequest) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetAccount

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAccount() SlurmdbV0041PostAccountsAssociationRequestAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAccountOk() (*SlurmdbV0041PostAccountsAssociationRequestAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SlurmdbV0041PostAccountsAssociationRequest) SetAccount(v SlurmdbV0041PostAccountsAssociationRequestAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SlurmdbV0041PostAccountsAssociationRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmdbV0041PostAccountsAssociationRequest) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmdbV0041PostAccountsAssociationRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmdbV0041PostAccountsAssociationRequest) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmdbV0041PostAccountsAssociationRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmdbV0041PostAccountsAssociationRequest) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmdbV0041PostAccountsAssociationRequest) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmdbV0041PostAccountsAssociationRequest) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


