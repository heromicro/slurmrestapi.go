# V0041OpenapiSlurmdbdQosResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qos** | [**[]V0041OpenapiSlurmdbdConfigRespQosInner**](V0041OpenapiSlurmdbdConfigRespQosInner.md) | List of QOS | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdQosResp

`func NewV0041OpenapiSlurmdbdQosResp(qos []V0041OpenapiSlurmdbdConfigRespQosInner, ) *V0041OpenapiSlurmdbdQosResp`

NewV0041OpenapiSlurmdbdQosResp instantiates a new V0041OpenapiSlurmdbdQosResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdQosRespWithDefaults

`func NewV0041OpenapiSlurmdbdQosRespWithDefaults() *V0041OpenapiSlurmdbdQosResp`

NewV0041OpenapiSlurmdbdQosRespWithDefaults instantiates a new V0041OpenapiSlurmdbdQosResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQos

`func (o *V0041OpenapiSlurmdbdQosResp) GetQos() []V0041OpenapiSlurmdbdConfigRespQosInner`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiSlurmdbdQosResp) GetQosOk() (*[]V0041OpenapiSlurmdbdConfigRespQosInner, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiSlurmdbdQosResp) SetQos(v []V0041OpenapiSlurmdbdConfigRespQosInner)`

SetQos sets Qos field to given value.


### GetMeta

`func (o *V0041OpenapiSlurmdbdQosResp) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiSlurmdbdQosResp) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiSlurmdbdQosResp) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiSlurmdbdQosResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiSlurmdbdQosResp) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiSlurmdbdQosResp) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiSlurmdbdQosResp) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiSlurmdbdQosResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiSlurmdbdQosResp) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiSlurmdbdQosResp) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiSlurmdbdQosResp) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiSlurmdbdQosResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


