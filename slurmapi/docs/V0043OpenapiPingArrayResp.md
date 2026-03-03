# V0043OpenapiPingArrayResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]V0043ControllerPing**](V0043ControllerPing.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiPingArrayResp

`func NewV0043OpenapiPingArrayResp(pings []V0043ControllerPing, ) *V0043OpenapiPingArrayResp`

NewV0043OpenapiPingArrayResp instantiates a new V0043OpenapiPingArrayResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiPingArrayRespWithDefaults

`func NewV0043OpenapiPingArrayRespWithDefaults() *V0043OpenapiPingArrayResp`

NewV0043OpenapiPingArrayRespWithDefaults instantiates a new V0043OpenapiPingArrayResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *V0043OpenapiPingArrayResp) GetPings() []V0043ControllerPing`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0043OpenapiPingArrayResp) GetPingsOk() (*[]V0043ControllerPing, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0043OpenapiPingArrayResp) SetPings(v []V0043ControllerPing)`

SetPings sets Pings field to given value.


### GetMeta

`func (o *V0043OpenapiPingArrayResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiPingArrayResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiPingArrayResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiPingArrayResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiPingArrayResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiPingArrayResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiPingArrayResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiPingArrayResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiPingArrayResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiPingArrayResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiPingArrayResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiPingArrayResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


