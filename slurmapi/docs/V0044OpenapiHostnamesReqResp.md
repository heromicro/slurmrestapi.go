# V0044OpenapiHostnamesReqResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | **[]string** |  | 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiHostnamesReqResp

`func NewV0044OpenapiHostnamesReqResp(hostnames []string, ) *V0044OpenapiHostnamesReqResp`

NewV0044OpenapiHostnamesReqResp instantiates a new V0044OpenapiHostnamesReqResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiHostnamesReqRespWithDefaults

`func NewV0044OpenapiHostnamesReqRespWithDefaults() *V0044OpenapiHostnamesReqResp`

NewV0044OpenapiHostnamesReqRespWithDefaults instantiates a new V0044OpenapiHostnamesReqResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnames

`func (o *V0044OpenapiHostnamesReqResp) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *V0044OpenapiHostnamesReqResp) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *V0044OpenapiHostnamesReqResp) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetMeta

`func (o *V0044OpenapiHostnamesReqResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiHostnamesReqResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiHostnamesReqResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiHostnamesReqResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiHostnamesReqResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiHostnamesReqResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiHostnamesReqResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiHostnamesReqResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiHostnamesReqResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiHostnamesReqResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiHostnamesReqResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiHostnamesReqResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


