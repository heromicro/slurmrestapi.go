# SlurmdbV0041PostUsersAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | [**SlurmdbV0041PostUsersAssociationRequestAssociationCondition**](SlurmdbV0041PostUsersAssociationRequestAssociationCondition.md) |  | 
**User** | [**SlurmdbV0041PostUsersAssociationRequestUser**](SlurmdbV0041PostUsersAssociationRequestUser.md) |  | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmdbV0041PostUsersAssociationRequest

`func NewSlurmdbV0041PostUsersAssociationRequest(associationCondition SlurmdbV0041PostUsersAssociationRequestAssociationCondition, user SlurmdbV0041PostUsersAssociationRequestUser, ) *SlurmdbV0041PostUsersAssociationRequest`

NewSlurmdbV0041PostUsersAssociationRequest instantiates a new SlurmdbV0041PostUsersAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostUsersAssociationRequestWithDefaults

`func NewSlurmdbV0041PostUsersAssociationRequestWithDefaults() *SlurmdbV0041PostUsersAssociationRequest`

NewSlurmdbV0041PostUsersAssociationRequestWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetAssociationCondition() SlurmdbV0041PostUsersAssociationRequestAssociationCondition`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetAssociationConditionOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationCondition, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *SlurmdbV0041PostUsersAssociationRequest) SetAssociationCondition(v SlurmdbV0041PostUsersAssociationRequestAssociationCondition)`

SetAssociationCondition sets AssociationCondition field to given value.


### GetUser

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetUser() SlurmdbV0041PostUsersAssociationRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetUserOk() (*SlurmdbV0041PostUsersAssociationRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SlurmdbV0041PostUsersAssociationRequest) SetUser(v SlurmdbV0041PostUsersAssociationRequestUser)`

SetUser sets User field to given value.


### GetMeta

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmdbV0041PostUsersAssociationRequest) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmdbV0041PostUsersAssociationRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmdbV0041PostUsersAssociationRequest) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmdbV0041PostUsersAssociationRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmdbV0041PostUsersAssociationRequest) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmdbV0041PostUsersAssociationRequest) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmdbV0041PostUsersAssociationRequest) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


