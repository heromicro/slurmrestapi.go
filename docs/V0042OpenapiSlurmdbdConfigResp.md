# V0042OpenapiSlurmdbdConfigResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V0042ClusterRec**](V0042ClusterRec.md) |  | [optional] 
**Tres** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Accounts** | Pointer to [**[]V0042Account**](V0042Account.md) |  | [optional] 
**Users** | Pointer to [**[]V0042User**](V0042User.md) |  | [optional] 
**Qos** | Pointer to [**[]V0042Qos**](V0042Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]V0042Wckey**](V0042Wckey.md) |  | [optional] 
**Associations** | Pointer to [**[]V0042Assoc**](V0042Assoc.md) |  | [optional] 
**Instances** | Pointer to [**[]V0042Instance**](V0042Instance.md) |  | [optional] 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiSlurmdbdConfigResp

`func NewV0042OpenapiSlurmdbdConfigResp() *V0042OpenapiSlurmdbdConfigResp`

NewV0042OpenapiSlurmdbdConfigResp instantiates a new V0042OpenapiSlurmdbdConfigResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiSlurmdbdConfigRespWithDefaults

`func NewV0042OpenapiSlurmdbdConfigRespWithDefaults() *V0042OpenapiSlurmdbdConfigResp`

NewV0042OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0042OpenapiSlurmdbdConfigResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V0042OpenapiSlurmdbdConfigResp) GetClusters() []V0042ClusterRec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetClustersOk() (*[]V0042ClusterRec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0042OpenapiSlurmdbdConfigResp) SetClusters(v []V0042ClusterRec)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0042OpenapiSlurmdbdConfigResp) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTres

`func (o *V0042OpenapiSlurmdbdConfigResp) GetTres() []V0042Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetTresOk() (*[]V0042Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0042OpenapiSlurmdbdConfigResp) SetTres(v []V0042Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0042OpenapiSlurmdbdConfigResp) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *V0042OpenapiSlurmdbdConfigResp) GetAccounts() []V0042Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetAccountsOk() (*[]V0042Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0042OpenapiSlurmdbdConfigResp) SetAccounts(v []V0042Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0042OpenapiSlurmdbdConfigResp) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *V0042OpenapiSlurmdbdConfigResp) GetUsers() []V0042User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetUsersOk() (*[]V0042User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0042OpenapiSlurmdbdConfigResp) SetUsers(v []V0042User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0042OpenapiSlurmdbdConfigResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *V0042OpenapiSlurmdbdConfigResp) GetQos() []V0042Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetQosOk() (*[]V0042Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0042OpenapiSlurmdbdConfigResp) SetQos(v []V0042Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0042OpenapiSlurmdbdConfigResp) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *V0042OpenapiSlurmdbdConfigResp) GetWckeys() []V0042Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetWckeysOk() (*[]V0042Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0042OpenapiSlurmdbdConfigResp) SetWckeys(v []V0042Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0042OpenapiSlurmdbdConfigResp) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *V0042OpenapiSlurmdbdConfigResp) GetAssociations() []V0042Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetAssociationsOk() (*[]V0042Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0042OpenapiSlurmdbdConfigResp) SetAssociations(v []V0042Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0042OpenapiSlurmdbdConfigResp) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetInstances

`func (o *V0042OpenapiSlurmdbdConfigResp) GetInstances() []V0042Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetInstancesOk() (*[]V0042Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0042OpenapiSlurmdbdConfigResp) SetInstances(v []V0042Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V0042OpenapiSlurmdbdConfigResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *V0042OpenapiSlurmdbdConfigResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiSlurmdbdConfigResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiSlurmdbdConfigResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiSlurmdbdConfigResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiSlurmdbdConfigResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiSlurmdbdConfigResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiSlurmdbdConfigResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiSlurmdbdConfigResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiSlurmdbdConfigResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiSlurmdbdConfigResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


