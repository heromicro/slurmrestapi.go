# V0042OpenapiDiagResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | [**V0042StatsMsg**](V0042StatsMsg.md) |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiDiagResp

`func NewV0042OpenapiDiagResp(statistics V0042StatsMsg, ) *V0042OpenapiDiagResp`

NewV0042OpenapiDiagResp instantiates a new V0042OpenapiDiagResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiDiagRespWithDefaults

`func NewV0042OpenapiDiagRespWithDefaults() *V0042OpenapiDiagResp`

NewV0042OpenapiDiagRespWithDefaults instantiates a new V0042OpenapiDiagResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *V0042OpenapiDiagResp) GetStatistics() V0042StatsMsg`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0042OpenapiDiagResp) GetStatisticsOk() (*V0042StatsMsg, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0042OpenapiDiagResp) SetStatistics(v V0042StatsMsg)`

SetStatistics sets Statistics field to given value.


### GetMeta

`func (o *V0042OpenapiDiagResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiDiagResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiDiagResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiDiagResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiDiagResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiDiagResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiDiagResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiDiagResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiDiagResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiDiagResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiDiagResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiDiagResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


