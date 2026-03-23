# V0043Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]V0043AssocShort**](V0043AssocShort.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0043Coord**](V0043Coord.md) |  | [optional] 
**Description** | **string** | Arbitrary string describing the account | 
**Name** | **string** | Account name | 
**Organization** | **string** | Organization to which the account belongs | 
**Flags** | Pointer to **[]string** | Flags associated with this account | [optional] 

## Methods

### NewV0043Account

`func NewV0043Account(description string, name string, organization string, ) *V0043Account`

NewV0043Account instantiates a new V0043Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043AccountWithDefaults

`func NewV0043AccountWithDefaults() *V0043Account`

NewV0043AccountWithDefaults instantiates a new V0043Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *V0043Account) GetAssociations() []V0043AssocShort`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0043Account) GetAssociationsOk() (*[]V0043AssocShort, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0043Account) SetAssociations(v []V0043AssocShort)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0043Account) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0043Account) GetCoordinators() []V0043Coord`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0043Account) GetCoordinatorsOk() (*[]V0043Coord, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0043Account) SetCoordinators(v []V0043Coord)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0043Account) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDescription

`func (o *V0043Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0043Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0043Account) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *V0043Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0043Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0043Account) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *V0043Account) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V0043Account) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V0043Account) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetFlags

`func (o *V0043Account) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0043Account) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0043Account) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0043Account) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


