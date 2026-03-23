# V0043OpenapiJobInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]V0043JobInfo**](V0043JobInfo.md) |  | 
**LastBackfill** | [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | 
**LastUpdate** | [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiJobInfoResp

`func NewV0043OpenapiJobInfoResp(jobs []V0043JobInfo, lastBackfill V0043Uint64NoValStruct, lastUpdate V0043Uint64NoValStruct, ) *V0043OpenapiJobInfoResp`

NewV0043OpenapiJobInfoResp instantiates a new V0043OpenapiJobInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiJobInfoRespWithDefaults

`func NewV0043OpenapiJobInfoRespWithDefaults() *V0043OpenapiJobInfoResp`

NewV0043OpenapiJobInfoRespWithDefaults instantiates a new V0043OpenapiJobInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V0043OpenapiJobInfoResp) GetJobs() []V0043JobInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0043OpenapiJobInfoResp) GetJobsOk() (*[]V0043JobInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0043OpenapiJobInfoResp) SetJobs(v []V0043JobInfo)`

SetJobs sets Jobs field to given value.


### GetLastBackfill

`func (o *V0043OpenapiJobInfoResp) GetLastBackfill() V0043Uint64NoValStruct`

GetLastBackfill returns the LastBackfill field if non-nil, zero value otherwise.

### GetLastBackfillOk

`func (o *V0043OpenapiJobInfoResp) GetLastBackfillOk() (*V0043Uint64NoValStruct, bool)`

GetLastBackfillOk returns a tuple with the LastBackfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackfill

`func (o *V0043OpenapiJobInfoResp) SetLastBackfill(v V0043Uint64NoValStruct)`

SetLastBackfill sets LastBackfill field to given value.


### GetLastUpdate

`func (o *V0043OpenapiJobInfoResp) GetLastUpdate() V0043Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0043OpenapiJobInfoResp) GetLastUpdateOk() (*V0043Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0043OpenapiJobInfoResp) SetLastUpdate(v V0043Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0043OpenapiJobInfoResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiJobInfoResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiJobInfoResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiJobInfoResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiJobInfoResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiJobInfoResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiJobInfoResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiJobInfoResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiJobInfoResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiJobInfoResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiJobInfoResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiJobInfoResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


