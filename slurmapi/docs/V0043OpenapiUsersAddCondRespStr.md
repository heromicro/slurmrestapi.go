# V0043OpenapiUsersAddCondRespStr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedUsers** | **string** | added_users | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiUsersAddCondRespStr

`func NewV0043OpenapiUsersAddCondRespStr(addedUsers string, ) *V0043OpenapiUsersAddCondRespStr`

NewV0043OpenapiUsersAddCondRespStr instantiates a new V0043OpenapiUsersAddCondRespStr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiUsersAddCondRespStrWithDefaults

`func NewV0043OpenapiUsersAddCondRespStrWithDefaults() *V0043OpenapiUsersAddCondRespStr`

NewV0043OpenapiUsersAddCondRespStrWithDefaults instantiates a new V0043OpenapiUsersAddCondRespStr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedUsers

`func (o *V0043OpenapiUsersAddCondRespStr) GetAddedUsers() string`

GetAddedUsers returns the AddedUsers field if non-nil, zero value otherwise.

### GetAddedUsersOk

`func (o *V0043OpenapiUsersAddCondRespStr) GetAddedUsersOk() (*string, bool)`

GetAddedUsersOk returns a tuple with the AddedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedUsers

`func (o *V0043OpenapiUsersAddCondRespStr) SetAddedUsers(v string)`

SetAddedUsers sets AddedUsers field to given value.


### GetMeta

`func (o *V0043OpenapiUsersAddCondRespStr) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiUsersAddCondRespStr) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiUsersAddCondRespStr) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiUsersAddCondRespStr) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiUsersAddCondRespStr) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiUsersAddCondRespStr) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiUsersAddCondRespStr) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiUsersAddCondRespStr) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiUsersAddCondRespStr) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiUsersAddCondRespStr) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiUsersAddCondRespStr) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiUsersAddCondRespStr) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


