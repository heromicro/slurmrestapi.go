# V0041OpenapiUsersResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V0041OpenapiSlurmdbdConfigRespUsersInner**](V0041OpenapiSlurmdbdConfigRespUsersInner.md) | users | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiUsersResp

`func NewV0041OpenapiUsersResp(users []V0041OpenapiSlurmdbdConfigRespUsersInner, ) *V0041OpenapiUsersResp`

NewV0041OpenapiUsersResp instantiates a new V0041OpenapiUsersResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiUsersRespWithDefaults

`func NewV0041OpenapiUsersRespWithDefaults() *V0041OpenapiUsersResp`

NewV0041OpenapiUsersRespWithDefaults instantiates a new V0041OpenapiUsersResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *V0041OpenapiUsersResp) GetUsers() []V0041OpenapiSlurmdbdConfigRespUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041OpenapiUsersResp) GetUsersOk() (*[]V0041OpenapiSlurmdbdConfigRespUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041OpenapiUsersResp) SetUsers(v []V0041OpenapiSlurmdbdConfigRespUsersInner)`

SetUsers sets Users field to given value.


### GetMeta

`func (o *V0041OpenapiUsersResp) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiUsersResp) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiUsersResp) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiUsersResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiUsersResp) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiUsersResp) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiUsersResp) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiUsersResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiUsersResp) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiUsersResp) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiUsersResp) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiUsersResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


