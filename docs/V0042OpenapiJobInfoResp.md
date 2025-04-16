# V0042OpenapiJobInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]V0042JobInfo**](V0042JobInfo.md) |  | 
**LastBackfill** | [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | 
**LastUpdate** | [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiJobInfoResp

`func NewV0042OpenapiJobInfoResp(jobs []V0042JobInfo, lastBackfill V0042Uint64NoValStruct, lastUpdate V0042Uint64NoValStruct, ) *V0042OpenapiJobInfoResp`

NewV0042OpenapiJobInfoResp instantiates a new V0042OpenapiJobInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiJobInfoRespWithDefaults

`func NewV0042OpenapiJobInfoRespWithDefaults() *V0042OpenapiJobInfoResp`

NewV0042OpenapiJobInfoRespWithDefaults instantiates a new V0042OpenapiJobInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V0042OpenapiJobInfoResp) GetJobs() []V0042JobInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0042OpenapiJobInfoResp) GetJobsOk() (*[]V0042JobInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0042OpenapiJobInfoResp) SetJobs(v []V0042JobInfo)`

SetJobs sets Jobs field to given value.


### GetLastBackfill

`func (o *V0042OpenapiJobInfoResp) GetLastBackfill() V0042Uint64NoValStruct`

GetLastBackfill returns the LastBackfill field if non-nil, zero value otherwise.

### GetLastBackfillOk

`func (o *V0042OpenapiJobInfoResp) GetLastBackfillOk() (*V0042Uint64NoValStruct, bool)`

GetLastBackfillOk returns a tuple with the LastBackfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackfill

`func (o *V0042OpenapiJobInfoResp) SetLastBackfill(v V0042Uint64NoValStruct)`

SetLastBackfill sets LastBackfill field to given value.


### GetLastUpdate

`func (o *V0042OpenapiJobInfoResp) GetLastUpdate() V0042Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0042OpenapiJobInfoResp) GetLastUpdateOk() (*V0042Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0042OpenapiJobInfoResp) SetLastUpdate(v V0042Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0042OpenapiJobInfoResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiJobInfoResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiJobInfoResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiJobInfoResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiJobInfoResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiJobInfoResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiJobInfoResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiJobInfoResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiJobInfoResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiJobInfoResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiJobInfoResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiJobInfoResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


