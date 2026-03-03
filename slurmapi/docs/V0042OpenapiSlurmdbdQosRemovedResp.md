# V0042OpenapiSlurmdbdQosRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedQos** | **[]string** |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiSlurmdbdQosRemovedResp

`func NewV0042OpenapiSlurmdbdQosRemovedResp(removedQos []string, ) *V0042OpenapiSlurmdbdQosRemovedResp`

NewV0042OpenapiSlurmdbdQosRemovedResp instantiates a new V0042OpenapiSlurmdbdQosRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiSlurmdbdQosRemovedRespWithDefaults

`func NewV0042OpenapiSlurmdbdQosRemovedRespWithDefaults() *V0042OpenapiSlurmdbdQosRemovedResp`

NewV0042OpenapiSlurmdbdQosRemovedRespWithDefaults instantiates a new V0042OpenapiSlurmdbdQosRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedQos

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetRemovedQos() []string`

GetRemovedQos returns the RemovedQos field if non-nil, zero value otherwise.

### GetRemovedQosOk

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetRemovedQosOk() (*[]string, bool)`

GetRemovedQosOk returns a tuple with the RemovedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedQos

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) SetRemovedQos(v []string)`

SetRemovedQos sets RemovedQos field to given value.


### GetMeta

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiSlurmdbdQosRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


