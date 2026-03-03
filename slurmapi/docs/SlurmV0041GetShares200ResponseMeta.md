# SlurmV0041GetShares200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**SlurmV0041GetShares200ResponseMetaPlugin**](SlurmV0041GetShares200ResponseMetaPlugin.md) |  | [optional] 
**Client** | Pointer to [**SlurmV0041GetShares200ResponseMetaClient**](SlurmV0041GetShares200ResponseMetaClient.md) |  | [optional] 
**Command** | Pointer to **[]string** | CLI command (if applicable) | [optional] 
**Slurm** | Pointer to [**SlurmV0041GetShares200ResponseMetaSlurm**](SlurmV0041GetShares200ResponseMetaSlurm.md) |  | [optional] 

## Methods

### NewSlurmV0041GetShares200ResponseMeta

`func NewSlurmV0041GetShares200ResponseMeta() *SlurmV0041GetShares200ResponseMeta`

NewSlurmV0041GetShares200ResponseMeta instantiates a new SlurmV0041GetShares200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseMetaWithDefaults

`func NewSlurmV0041GetShares200ResponseMetaWithDefaults() *SlurmV0041GetShares200ResponseMeta`

NewSlurmV0041GetShares200ResponseMetaWithDefaults instantiates a new SlurmV0041GetShares200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *SlurmV0041GetShares200ResponseMeta) GetPlugin() SlurmV0041GetShares200ResponseMetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *SlurmV0041GetShares200ResponseMeta) GetPluginOk() (*SlurmV0041GetShares200ResponseMetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *SlurmV0041GetShares200ResponseMeta) SetPlugin(v SlurmV0041GetShares200ResponseMetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *SlurmV0041GetShares200ResponseMeta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetClient

`func (o *SlurmV0041GetShares200ResponseMeta) GetClient() SlurmV0041GetShares200ResponseMetaClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SlurmV0041GetShares200ResponseMeta) GetClientOk() (*SlurmV0041GetShares200ResponseMetaClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SlurmV0041GetShares200ResponseMeta) SetClient(v SlurmV0041GetShares200ResponseMetaClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *SlurmV0041GetShares200ResponseMeta) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommand

`func (o *SlurmV0041GetShares200ResponseMeta) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SlurmV0041GetShares200ResponseMeta) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SlurmV0041GetShares200ResponseMeta) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SlurmV0041GetShares200ResponseMeta) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSlurm

`func (o *SlurmV0041GetShares200ResponseMeta) GetSlurm() SlurmV0041GetShares200ResponseMetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *SlurmV0041GetShares200ResponseMeta) GetSlurmOk() (*SlurmV0041GetShares200ResponseMetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *SlurmV0041GetShares200ResponseMeta) SetSlurm(v SlurmV0041GetShares200ResponseMetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *SlurmV0041GetShares200ResponseMeta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


