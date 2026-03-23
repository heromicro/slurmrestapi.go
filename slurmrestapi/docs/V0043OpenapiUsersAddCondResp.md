# V0043OpenapiUsersAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | [**V0043UsersAddCond**](V0043UsersAddCond.md) |  | 
**User** | [**V0043UserShort**](V0043UserShort.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiUsersAddCondResp

`func NewV0043OpenapiUsersAddCondResp(associationCondition V0043UsersAddCond, user V0043UserShort, ) *V0043OpenapiUsersAddCondResp`

NewV0043OpenapiUsersAddCondResp instantiates a new V0043OpenapiUsersAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiUsersAddCondRespWithDefaults

`func NewV0043OpenapiUsersAddCondRespWithDefaults() *V0043OpenapiUsersAddCondResp`

NewV0043OpenapiUsersAddCondRespWithDefaults instantiates a new V0043OpenapiUsersAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0043OpenapiUsersAddCondResp) GetAssociationCondition() V0043UsersAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0043OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0043UsersAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0043OpenapiUsersAddCondResp) SetAssociationCondition(v V0043UsersAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.


### GetUser

`func (o *V0043OpenapiUsersAddCondResp) GetUser() V0043UserShort`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0043OpenapiUsersAddCondResp) GetUserOk() (*V0043UserShort, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0043OpenapiUsersAddCondResp) SetUser(v V0043UserShort)`

SetUser sets User field to given value.


### GetMeta

`func (o *V0043OpenapiUsersAddCondResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiUsersAddCondResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiUsersAddCondResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiUsersAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiUsersAddCondResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiUsersAddCondResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiUsersAddCondResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiUsersAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiUsersAddCondResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiUsersAddCondResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiUsersAddCondResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiUsersAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


