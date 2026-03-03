# SlurmV0041GetShares200ResponseMetaSlurm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**SlurmV0041GetShares200ResponseMetaSlurmVersion**](SlurmV0041GetShares200ResponseMetaSlurmVersion.md) |  | [optional] 
**Release** | Pointer to **string** | Slurm release string | [optional] 
**Cluster** | Pointer to **string** | Slurm cluster name | [optional] 

## Methods

### NewSlurmV0041GetShares200ResponseMetaSlurm

`func NewSlurmV0041GetShares200ResponseMetaSlurm() *SlurmV0041GetShares200ResponseMetaSlurm`

NewSlurmV0041GetShares200ResponseMetaSlurm instantiates a new SlurmV0041GetShares200ResponseMetaSlurm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseMetaSlurmWithDefaults

`func NewSlurmV0041GetShares200ResponseMetaSlurmWithDefaults() *SlurmV0041GetShares200ResponseMetaSlurm`

NewSlurmV0041GetShares200ResponseMetaSlurmWithDefaults instantiates a new SlurmV0041GetShares200ResponseMetaSlurm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetVersion() SlurmV0041GetShares200ResponseMetaSlurmVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetVersionOk() (*SlurmV0041GetShares200ResponseMetaSlurmVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) SetVersion(v SlurmV0041GetShares200ResponseMetaSlurmVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCluster

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SlurmV0041GetShares200ResponseMetaSlurm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


