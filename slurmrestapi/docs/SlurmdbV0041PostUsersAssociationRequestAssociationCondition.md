# SlurmdbV0041PostUsersAssociationRequestAssociationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | CSV accounts list | [optional] 
**Association** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation.md) |  | [optional] 
**Clusters** | Pointer to **[]string** | CSV clusters list | [optional] 
**Partitions** | Pointer to **[]string** | CSV partitions list | [optional] 
**Users** | **[]string** | CSV users list | 
**Wckeys** | Pointer to **[]string** | CSV WCKeys list | [optional] 

## Methods

### NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition(users []string, ) *SlurmdbV0041PostUsersAssociationRequestAssociationCondition`

NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionWithDefaults

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionWithDefaults() *SlurmdbV0041PostUsersAssociationRequestAssociationCondition`

NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociation

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAssociation() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAssociationOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetAssociation(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetPartitions

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetUsers

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetWckeys

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetWckeys() []string`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetWckeysOk() (*[]string, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetWckeys(v []string)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


