# V0043StatsRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** | When data collection started (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Rollups** | Pointer to [**V0043RollupStats**](V0043RollupStats.md) |  | [optional] 
**RPCs** | Pointer to [**[]V0043StatsRpc**](V0043StatsRpc.md) |  | [optional] 
**Users** | Pointer to [**[]V0043StatsUser**](V0043StatsUser.md) |  | [optional] 

## Methods

### NewV0043StatsRec

`func NewV0043StatsRec() *V0043StatsRec`

NewV0043StatsRec instantiates a new V0043StatsRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StatsRecWithDefaults

`func NewV0043StatsRecWithDefaults() *V0043StatsRec`

NewV0043StatsRecWithDefaults instantiates a new V0043StatsRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *V0043StatsRec) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0043StatsRec) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0043StatsRec) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0043StatsRec) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *V0043StatsRec) GetRollups() V0043RollupStats`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *V0043StatsRec) GetRollupsOk() (*V0043RollupStats, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *V0043StatsRec) SetRollups(v V0043RollupStats)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *V0043StatsRec) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *V0043StatsRec) GetRPCs() []V0043StatsRpc`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *V0043StatsRec) GetRPCsOk() (*[]V0043StatsRpc, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *V0043StatsRec) SetRPCs(v []V0043StatsRpc)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *V0043StatsRec) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *V0043StatsRec) GetUsers() []V0043StatsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0043StatsRec) GetUsersOk() (*[]V0043StatsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0043StatsRec) SetUsers(v []V0043StatsUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0043StatsRec) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


