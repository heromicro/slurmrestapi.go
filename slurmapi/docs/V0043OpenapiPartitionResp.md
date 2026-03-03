# V0043OpenapiPartitionResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | [**[]V0043PartitionInfo**](V0043PartitionInfo.md) |  | 
**LastUpdate** | [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiPartitionResp

`func NewV0043OpenapiPartitionResp(partitions []V0043PartitionInfo, lastUpdate V0043Uint64NoValStruct, ) *V0043OpenapiPartitionResp`

NewV0043OpenapiPartitionResp instantiates a new V0043OpenapiPartitionResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiPartitionRespWithDefaults

`func NewV0043OpenapiPartitionRespWithDefaults() *V0043OpenapiPartitionResp`

NewV0043OpenapiPartitionRespWithDefaults instantiates a new V0043OpenapiPartitionResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *V0043OpenapiPartitionResp) GetPartitions() []V0043PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0043OpenapiPartitionResp) GetPartitionsOk() (*[]V0043PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0043OpenapiPartitionResp) SetPartitions(v []V0043PartitionInfo)`

SetPartitions sets Partitions field to given value.


### GetLastUpdate

`func (o *V0043OpenapiPartitionResp) GetLastUpdate() V0043Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0043OpenapiPartitionResp) GetLastUpdateOk() (*V0043Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0043OpenapiPartitionResp) SetLastUpdate(v V0043Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0043OpenapiPartitionResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiPartitionResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiPartitionResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiPartitionResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiPartitionResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiPartitionResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiPartitionResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiPartitionResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiPartitionResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiPartitionResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiPartitionResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiPartitionResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


