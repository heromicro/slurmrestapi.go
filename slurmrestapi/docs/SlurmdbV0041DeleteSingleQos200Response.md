# SlurmdbV0041DeleteSingleQos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedQos** | **[]string** | removed QOS | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmdbV0041DeleteSingleQos200Response

`func NewSlurmdbV0041DeleteSingleQos200Response(removedQos []string, ) *SlurmdbV0041DeleteSingleQos200Response`

NewSlurmdbV0041DeleteSingleQos200Response instantiates a new SlurmdbV0041DeleteSingleQos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041DeleteSingleQos200ResponseWithDefaults

`func NewSlurmdbV0041DeleteSingleQos200ResponseWithDefaults() *SlurmdbV0041DeleteSingleQos200Response`

NewSlurmdbV0041DeleteSingleQos200ResponseWithDefaults instantiates a new SlurmdbV0041DeleteSingleQos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedQos

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetRemovedQos() []string`

GetRemovedQos returns the RemovedQos field if non-nil, zero value otherwise.

### GetRemovedQosOk

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetRemovedQosOk() (*[]string, bool)`

GetRemovedQosOk returns a tuple with the RemovedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedQos

`func (o *SlurmdbV0041DeleteSingleQos200Response) SetRemovedQos(v []string)`

SetRemovedQos sets RemovedQos field to given value.


### GetMeta

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmdbV0041DeleteSingleQos200Response) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmdbV0041DeleteSingleQos200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmdbV0041DeleteSingleQos200Response) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmdbV0041DeleteSingleQos200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmdbV0041DeleteSingleQos200Response) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmdbV0041DeleteSingleQos200Response) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmdbV0041DeleteSingleQos200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


