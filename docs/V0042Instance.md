# V0042Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**InstanceId** | Pointer to **string** | Cloud instance ID | [optional] 
**InstanceType** | Pointer to **string** | Cloud instance type | [optional] 
**NodeName** | Pointer to **string** | NodeName | [optional] 
**Time** | Pointer to [**V0040InstanceTime**](V0040InstanceTime.md) |  | [optional] 

## Methods

### NewV0042Instance

`func NewV0042Instance() *V0042Instance`

NewV0042Instance instantiates a new V0042Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042InstanceWithDefaults

`func NewV0042InstanceWithDefaults() *V0042Instance`

NewV0042InstanceWithDefaults instantiates a new V0042Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *V0042Instance) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0042Instance) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0042Instance) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0042Instance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetExtra

`func (o *V0042Instance) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0042Instance) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0042Instance) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0042Instance) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetInstanceId

`func (o *V0042Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *V0042Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *V0042Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *V0042Instance) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceType

`func (o *V0042Instance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *V0042Instance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *V0042Instance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *V0042Instance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetNodeName

`func (o *V0042Instance) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *V0042Instance) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *V0042Instance) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *V0042Instance) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetTime

`func (o *V0042Instance) GetTime() V0040InstanceTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042Instance) GetTimeOk() (*V0040InstanceTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042Instance) SetTime(v V0040InstanceTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042Instance) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


