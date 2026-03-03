# V0042OpenapiSlurmdbdPingResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]V0042SlurmdbdPing**](V0042SlurmdbdPing.md) |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiSlurmdbdPingResp

`func NewV0042OpenapiSlurmdbdPingResp(pings []V0042SlurmdbdPing, ) *V0042OpenapiSlurmdbdPingResp`

NewV0042OpenapiSlurmdbdPingResp instantiates a new V0042OpenapiSlurmdbdPingResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiSlurmdbdPingRespWithDefaults

`func NewV0042OpenapiSlurmdbdPingRespWithDefaults() *V0042OpenapiSlurmdbdPingResp`

NewV0042OpenapiSlurmdbdPingRespWithDefaults instantiates a new V0042OpenapiSlurmdbdPingResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *V0042OpenapiSlurmdbdPingResp) GetPings() []V0042SlurmdbdPing`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0042OpenapiSlurmdbdPingResp) GetPingsOk() (*[]V0042SlurmdbdPing, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0042OpenapiSlurmdbdPingResp) SetPings(v []V0042SlurmdbdPing)`

SetPings sets Pings field to given value.


### GetMeta

`func (o *V0042OpenapiSlurmdbdPingResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiSlurmdbdPingResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiSlurmdbdPingResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiSlurmdbdPingResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiSlurmdbdPingResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiSlurmdbdPingResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiSlurmdbdPingResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiSlurmdbdPingResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiSlurmdbdPingResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiSlurmdbdPingResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiSlurmdbdPingResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiSlurmdbdPingResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


