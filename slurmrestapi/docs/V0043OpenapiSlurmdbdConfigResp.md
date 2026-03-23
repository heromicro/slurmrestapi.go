# V0043OpenapiSlurmdbdConfigResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V0043ClusterRec**](V0043ClusterRec.md) |  | [optional] 
**Tres** | Pointer to [**[]V0043Tres**](V0043Tres.md) |  | [optional] 
**Accounts** | Pointer to [**[]V0043Account**](V0043Account.md) |  | [optional] 
**Users** | Pointer to [**[]V0043User**](V0043User.md) |  | [optional] 
**Qos** | Pointer to [**[]V0043Qos**](V0043Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]V0043Wckey**](V0043Wckey.md) |  | [optional] 
**Associations** | Pointer to [**[]V0043Assoc**](V0043Assoc.md) |  | [optional] 
**Instances** | Pointer to [**[]V0043Instance**](V0043Instance.md) |  | [optional] 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiSlurmdbdConfigResp

`func NewV0043OpenapiSlurmdbdConfigResp() *V0043OpenapiSlurmdbdConfigResp`

NewV0043OpenapiSlurmdbdConfigResp instantiates a new V0043OpenapiSlurmdbdConfigResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiSlurmdbdConfigRespWithDefaults

`func NewV0043OpenapiSlurmdbdConfigRespWithDefaults() *V0043OpenapiSlurmdbdConfigResp`

NewV0043OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0043OpenapiSlurmdbdConfigResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V0043OpenapiSlurmdbdConfigResp) GetClusters() []V0043ClusterRec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetClustersOk() (*[]V0043ClusterRec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0043OpenapiSlurmdbdConfigResp) SetClusters(v []V0043ClusterRec)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0043OpenapiSlurmdbdConfigResp) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTres

`func (o *V0043OpenapiSlurmdbdConfigResp) GetTres() []V0043Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetTresOk() (*[]V0043Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0043OpenapiSlurmdbdConfigResp) SetTres(v []V0043Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0043OpenapiSlurmdbdConfigResp) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *V0043OpenapiSlurmdbdConfigResp) GetAccounts() []V0043Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetAccountsOk() (*[]V0043Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0043OpenapiSlurmdbdConfigResp) SetAccounts(v []V0043Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0043OpenapiSlurmdbdConfigResp) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *V0043OpenapiSlurmdbdConfigResp) GetUsers() []V0043User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetUsersOk() (*[]V0043User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0043OpenapiSlurmdbdConfigResp) SetUsers(v []V0043User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0043OpenapiSlurmdbdConfigResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *V0043OpenapiSlurmdbdConfigResp) GetQos() []V0043Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetQosOk() (*[]V0043Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0043OpenapiSlurmdbdConfigResp) SetQos(v []V0043Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0043OpenapiSlurmdbdConfigResp) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *V0043OpenapiSlurmdbdConfigResp) GetWckeys() []V0043Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetWckeysOk() (*[]V0043Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0043OpenapiSlurmdbdConfigResp) SetWckeys(v []V0043Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0043OpenapiSlurmdbdConfigResp) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *V0043OpenapiSlurmdbdConfigResp) GetAssociations() []V0043Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetAssociationsOk() (*[]V0043Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0043OpenapiSlurmdbdConfigResp) SetAssociations(v []V0043Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0043OpenapiSlurmdbdConfigResp) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetInstances

`func (o *V0043OpenapiSlurmdbdConfigResp) GetInstances() []V0043Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetInstancesOk() (*[]V0043Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0043OpenapiSlurmdbdConfigResp) SetInstances(v []V0043Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V0043OpenapiSlurmdbdConfigResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *V0043OpenapiSlurmdbdConfigResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiSlurmdbdConfigResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiSlurmdbdConfigResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiSlurmdbdConfigResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiSlurmdbdConfigResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiSlurmdbdConfigResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiSlurmdbdConfigResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiSlurmdbdConfigResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiSlurmdbdConfigResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiSlurmdbdConfigResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


