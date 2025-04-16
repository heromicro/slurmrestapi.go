# V0042UsersAddCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** |  | [optional] 
**Association** | Pointer to [**V0042AssocRecSet**](V0042AssocRecSet.md) |  | [optional] 
**Clusters** | Pointer to **[]string** |  | [optional] 
**Partitions** | Pointer to **[]string** |  | [optional] 
**Users** | **[]string** |  | 
**Wckeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0042UsersAddCond

`func NewV0042UsersAddCond(users []string, ) *V0042UsersAddCond`

NewV0042UsersAddCond instantiates a new V0042UsersAddCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042UsersAddCondWithDefaults

`func NewV0042UsersAddCondWithDefaults() *V0042UsersAddCond`

NewV0042UsersAddCondWithDefaults instantiates a new V0042UsersAddCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0042UsersAddCond) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0042UsersAddCond) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0042UsersAddCond) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0042UsersAddCond) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociation

`func (o *V0042UsersAddCond) GetAssociation() V0042AssocRecSet`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0042UsersAddCond) GetAssociationOk() (*V0042AssocRecSet, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0042UsersAddCond) SetAssociation(v V0042AssocRecSet)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0042UsersAddCond) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *V0042UsersAddCond) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0042UsersAddCond) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0042UsersAddCond) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0042UsersAddCond) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetPartitions

`func (o *V0042UsersAddCond) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0042UsersAddCond) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0042UsersAddCond) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0042UsersAddCond) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetUsers

`func (o *V0042UsersAddCond) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0042UsersAddCond) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0042UsersAddCond) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetWckeys

`func (o *V0042UsersAddCond) GetWckeys() []string`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0042UsersAddCond) GetWckeysOk() (*[]string, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0042UsersAddCond) SetWckeys(v []string)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0042UsersAddCond) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


