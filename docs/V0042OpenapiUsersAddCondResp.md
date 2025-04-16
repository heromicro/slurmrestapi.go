# V0042OpenapiUsersAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | [**V0042UsersAddCond**](V0042UsersAddCond.md) |  | 
**User** | [**V0042UserShort**](V0042UserShort.md) |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiUsersAddCondResp

`func NewV0042OpenapiUsersAddCondResp(associationCondition V0042UsersAddCond, user V0042UserShort, ) *V0042OpenapiUsersAddCondResp`

NewV0042OpenapiUsersAddCondResp instantiates a new V0042OpenapiUsersAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiUsersAddCondRespWithDefaults

`func NewV0042OpenapiUsersAddCondRespWithDefaults() *V0042OpenapiUsersAddCondResp`

NewV0042OpenapiUsersAddCondRespWithDefaults instantiates a new V0042OpenapiUsersAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0042OpenapiUsersAddCondResp) GetAssociationCondition() V0042UsersAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0042OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0042UsersAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0042OpenapiUsersAddCondResp) SetAssociationCondition(v V0042UsersAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.


### GetUser

`func (o *V0042OpenapiUsersAddCondResp) GetUser() V0042UserShort`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0042OpenapiUsersAddCondResp) GetUserOk() (*V0042UserShort, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0042OpenapiUsersAddCondResp) SetUser(v V0042UserShort)`

SetUser sets User field to given value.


### GetMeta

`func (o *V0042OpenapiUsersAddCondResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiUsersAddCondResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiUsersAddCondResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiUsersAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiUsersAddCondResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiUsersAddCondResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiUsersAddCondResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiUsersAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiUsersAddCondResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiUsersAddCondResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiUsersAddCondResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiUsersAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


