# V0041OpenapiSharesRespMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**V0040OpenapiMetaPlugin**](V0040OpenapiMetaPlugin.md) |  | [optional] 
**Client** | Pointer to [**V0040OpenapiMetaClient**](V0040OpenapiMetaClient.md) |  | [optional] 
**Command** | Pointer to **[]string** | CLI command (if applicable) | [optional] 
**Slurm** | Pointer to [**V0040OpenapiMetaSlurm**](V0040OpenapiMetaSlurm.md) |  | [optional] 

## Methods

### NewV0041OpenapiSharesRespMeta

`func NewV0041OpenapiSharesRespMeta() *V0041OpenapiSharesRespMeta`

NewV0041OpenapiSharesRespMeta instantiates a new V0041OpenapiSharesRespMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSharesRespMetaWithDefaults

`func NewV0041OpenapiSharesRespMetaWithDefaults() *V0041OpenapiSharesRespMeta`

NewV0041OpenapiSharesRespMetaWithDefaults instantiates a new V0041OpenapiSharesRespMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *V0041OpenapiSharesRespMeta) GetPlugin() V0040OpenapiMetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V0041OpenapiSharesRespMeta) GetPluginOk() (*V0040OpenapiMetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V0041OpenapiSharesRespMeta) SetPlugin(v V0040OpenapiMetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V0041OpenapiSharesRespMeta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetClient

`func (o *V0041OpenapiSharesRespMeta) GetClient() V0040OpenapiMetaClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *V0041OpenapiSharesRespMeta) GetClientOk() (*V0040OpenapiMetaClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *V0041OpenapiSharesRespMeta) SetClient(v V0040OpenapiMetaClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *V0041OpenapiSharesRespMeta) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommand

`func (o *V0041OpenapiSharesRespMeta) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041OpenapiSharesRespMeta) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041OpenapiSharesRespMeta) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041OpenapiSharesRespMeta) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSlurm

`func (o *V0041OpenapiSharesRespMeta) GetSlurm() V0040OpenapiMetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *V0041OpenapiSharesRespMeta) GetSlurmOk() (*V0040OpenapiMetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *V0041OpenapiSharesRespMeta) SetSlurm(v V0040OpenapiMetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *V0041OpenapiSharesRespMeta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


