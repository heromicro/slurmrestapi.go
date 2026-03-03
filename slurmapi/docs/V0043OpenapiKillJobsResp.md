# V0043OpenapiKillJobsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**[]V0043KillJobsRespJob**](V0043KillJobsRespJob.md) | List of jobs signal responses | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiKillJobsResp

`func NewV0043OpenapiKillJobsResp(status []V0043KillJobsRespJob, ) *V0043OpenapiKillJobsResp`

NewV0043OpenapiKillJobsResp instantiates a new V0043OpenapiKillJobsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiKillJobsRespWithDefaults

`func NewV0043OpenapiKillJobsRespWithDefaults() *V0043OpenapiKillJobsResp`

NewV0043OpenapiKillJobsRespWithDefaults instantiates a new V0043OpenapiKillJobsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0043OpenapiKillJobsResp) GetStatus() []V0043KillJobsRespJob`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0043OpenapiKillJobsResp) GetStatusOk() (*[]V0043KillJobsRespJob, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0043OpenapiKillJobsResp) SetStatus(v []V0043KillJobsRespJob)`

SetStatus sets Status field to given value.


### GetMeta

`func (o *V0043OpenapiKillJobsResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiKillJobsResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiKillJobsResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiKillJobsResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiKillJobsResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiKillJobsResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiKillJobsResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiKillJobsResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiKillJobsResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiKillJobsResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiKillJobsResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiKillJobsResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


