# V0041OpenapiSlurmdbdConfigRespUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorLevel** | Pointer to **[]string** | AdminLevel granted to the user | [optional] 
**Associations** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner**](V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner.md) | Associations created for this user | [optional] 
**Coordinators** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner**](V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner.md) | Accounts this user is a coordinator for | [optional] 
**Default** | Pointer to [**V0040UserDefault**](V0040UserDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Flags associated with user | [optional] 
**Name** | **string** | User name | 
**OldName** | Pointer to **string** | Previous user name | [optional] 
**Wckeys** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner**](V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner.md) | List of available WCKeys | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdConfigRespUsersInner

`func NewV0041OpenapiSlurmdbdConfigRespUsersInner(name string, ) *V0041OpenapiSlurmdbdConfigRespUsersInner`

NewV0041OpenapiSlurmdbdConfigRespUsersInner instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespUsersInnerWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespUsersInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespUsersInner`

NewV0041OpenapiSlurmdbdConfigRespUsersInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorLevel

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAdministratorLevel() []string`

GetAdministratorLevel returns the AdministratorLevel field if non-nil, zero value otherwise.

### GetAdministratorLevelOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAdministratorLevelOk() (*[]string, bool)`

GetAdministratorLevelOk returns a tuple with the AdministratorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorLevel

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetAdministratorLevel(v []string)`

SetAdministratorLevel sets AdministratorLevel field to given value.

### HasAdministratorLevel

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasAdministratorLevel() bool`

HasAdministratorLevel returns a boolean if a field has been set.

### GetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAssociations() []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAssociationsOk() (*[]V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetAssociations(v []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetCoordinators() []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetCoordinatorsOk() (*[]V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetCoordinators(v []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDefault

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetDefault() V0040UserDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetDefaultOk() (*V0040UserDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetDefault(v V0040UserDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetName(v string)`

SetName sets Name field to given value.


### GetOldName

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetOldName(v string)`

SetOldName sets OldName field to given value.

### HasOldName

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasOldName() bool`

HasOldName returns a boolean if a field has been set.

### GetWckeys

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetWckeys() []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetWckeysOk() (*[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetWckeys(v []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


