# V0043OpenapiSlurmdbdPingResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]V0043SlurmdbdPing**](V0043SlurmdbdPing.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiSlurmdbdPingResp

`func NewV0043OpenapiSlurmdbdPingResp(pings []V0043SlurmdbdPing, ) *V0043OpenapiSlurmdbdPingResp`

NewV0043OpenapiSlurmdbdPingResp instantiates a new V0043OpenapiSlurmdbdPingResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiSlurmdbdPingRespWithDefaults

`func NewV0043OpenapiSlurmdbdPingRespWithDefaults() *V0043OpenapiSlurmdbdPingResp`

NewV0043OpenapiSlurmdbdPingRespWithDefaults instantiates a new V0043OpenapiSlurmdbdPingResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *V0043OpenapiSlurmdbdPingResp) GetPings() []V0043SlurmdbdPing`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0043OpenapiSlurmdbdPingResp) GetPingsOk() (*[]V0043SlurmdbdPing, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0043OpenapiSlurmdbdPingResp) SetPings(v []V0043SlurmdbdPing)`

SetPings sets Pings field to given value.


### GetMeta

`func (o *V0043OpenapiSlurmdbdPingResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiSlurmdbdPingResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiSlurmdbdPingResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiSlurmdbdPingResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiSlurmdbdPingResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiSlurmdbdPingResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiSlurmdbdPingResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiSlurmdbdPingResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiSlurmdbdPingResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiSlurmdbdPingResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiSlurmdbdPingResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiSlurmdbdPingResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


