# V0043OpenapiSlurmdbdQosRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedQos** | **[]string** |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiSlurmdbdQosRemovedResp

`func NewV0043OpenapiSlurmdbdQosRemovedResp(removedQos []string, ) *V0043OpenapiSlurmdbdQosRemovedResp`

NewV0043OpenapiSlurmdbdQosRemovedResp instantiates a new V0043OpenapiSlurmdbdQosRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiSlurmdbdQosRemovedRespWithDefaults

`func NewV0043OpenapiSlurmdbdQosRemovedRespWithDefaults() *V0043OpenapiSlurmdbdQosRemovedResp`

NewV0043OpenapiSlurmdbdQosRemovedRespWithDefaults instantiates a new V0043OpenapiSlurmdbdQosRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedQos

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetRemovedQos() []string`

GetRemovedQos returns the RemovedQos field if non-nil, zero value otherwise.

### GetRemovedQosOk

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetRemovedQosOk() (*[]string, bool)`

GetRemovedQosOk returns a tuple with the RemovedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedQos

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) SetRemovedQos(v []string)`

SetRemovedQos sets RemovedQos field to given value.


### GetMeta

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiSlurmdbdQosRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


