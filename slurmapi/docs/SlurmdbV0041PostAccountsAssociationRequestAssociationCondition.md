# SlurmdbV0041PostAccountsAssociationRequestAssociationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **[]string** | CSV accounts list | 
**Association** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation.md) |  | [optional] 
**Clusters** | Pointer to **[]string** | CSV clusters list | [optional] 

## Methods

### NewSlurmdbV0041PostAccountsAssociationRequestAssociationCondition

`func NewSlurmdbV0041PostAccountsAssociationRequestAssociationCondition(accounts []string, ) *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition`

NewSlurmdbV0041PostAccountsAssociationRequestAssociationCondition instantiates a new SlurmdbV0041PostAccountsAssociationRequestAssociationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostAccountsAssociationRequestAssociationConditionWithDefaults

`func NewSlurmdbV0041PostAccountsAssociationRequestAssociationConditionWithDefaults() *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition`

NewSlurmdbV0041PostAccountsAssociationRequestAssociationConditionWithDefaults instantiates a new SlurmdbV0041PostAccountsAssociationRequestAssociationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.


### GetAssociation

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetAssociation() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetAssociationOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) SetAssociation(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


