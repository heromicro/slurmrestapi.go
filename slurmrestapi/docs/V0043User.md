# V0043User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorLevel** | Pointer to **[]string** | AdminLevel granted to the user | [optional] 
**Associations** | Pointer to [**[]V0043AssocShort**](V0043AssocShort.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0043Coord**](V0043Coord.md) |  | [optional] 
**Default** | Pointer to [**V0042UserDefault**](V0042UserDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Flags associated with this user | [optional] 
**Name** | **string** | User name | 
**OldName** | Pointer to **string** | Previous user name | [optional] 
**Wckeys** | Pointer to [**[]V0043Wckey**](V0043Wckey.md) |  | [optional] 

## Methods

### NewV0043User

`func NewV0043User(name string, ) *V0043User`

NewV0043User instantiates a new V0043User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043UserWithDefaults

`func NewV0043UserWithDefaults() *V0043User`

NewV0043UserWithDefaults instantiates a new V0043User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorLevel

`func (o *V0043User) GetAdministratorLevel() []string`

GetAdministratorLevel returns the AdministratorLevel field if non-nil, zero value otherwise.

### GetAdministratorLevelOk

`func (o *V0043User) GetAdministratorLevelOk() (*[]string, bool)`

GetAdministratorLevelOk returns a tuple with the AdministratorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorLevel

`func (o *V0043User) SetAdministratorLevel(v []string)`

SetAdministratorLevel sets AdministratorLevel field to given value.

### HasAdministratorLevel

`func (o *V0043User) HasAdministratorLevel() bool`

HasAdministratorLevel returns a boolean if a field has been set.

### GetAssociations

`func (o *V0043User) GetAssociations() []V0043AssocShort`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0043User) GetAssociationsOk() (*[]V0043AssocShort, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0043User) SetAssociations(v []V0043AssocShort)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0043User) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0043User) GetCoordinators() []V0043Coord`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0043User) GetCoordinatorsOk() (*[]V0043Coord, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0043User) SetCoordinators(v []V0043Coord)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0043User) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDefault

`func (o *V0043User) GetDefault() V0042UserDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0043User) GetDefaultOk() (*V0042UserDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0043User) SetDefault(v V0042UserDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0043User) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0043User) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0043User) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0043User) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0043User) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0043User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0043User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0043User) SetName(v string)`

SetName sets Name field to given value.


### GetOldName

`func (o *V0043User) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *V0043User) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *V0043User) SetOldName(v string)`

SetOldName sets OldName field to given value.

### HasOldName

`func (o *V0043User) HasOldName() bool`

HasOldName returns a boolean if a field has been set.

### GetWckeys

`func (o *V0043User) GetWckeys() []V0043Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0043User) GetWckeysOk() (*[]V0043Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0043User) SetWckeys(v []V0043Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0043User) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


