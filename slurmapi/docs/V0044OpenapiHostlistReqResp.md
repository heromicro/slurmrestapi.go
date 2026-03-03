# V0044OpenapiHostlistReqResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostlist** | **string** | Hostlist expression string (Hostlist expression string) | 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiHostlistReqResp

`func NewV0044OpenapiHostlistReqResp(hostlist string, ) *V0044OpenapiHostlistReqResp`

NewV0044OpenapiHostlistReqResp instantiates a new V0044OpenapiHostlistReqResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiHostlistReqRespWithDefaults

`func NewV0044OpenapiHostlistReqRespWithDefaults() *V0044OpenapiHostlistReqResp`

NewV0044OpenapiHostlistReqRespWithDefaults instantiates a new V0044OpenapiHostlistReqResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostlist

`func (o *V0044OpenapiHostlistReqResp) GetHostlist() string`

GetHostlist returns the Hostlist field if non-nil, zero value otherwise.

### GetHostlistOk

`func (o *V0044OpenapiHostlistReqResp) GetHostlistOk() (*string, bool)`

GetHostlistOk returns a tuple with the Hostlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostlist

`func (o *V0044OpenapiHostlistReqResp) SetHostlist(v string)`

SetHostlist sets Hostlist field to given value.


### GetMeta

`func (o *V0044OpenapiHostlistReqResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiHostlistReqResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiHostlistReqResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiHostlistReqResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiHostlistReqResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiHostlistReqResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiHostlistReqResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiHostlistReqResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiHostlistReqResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiHostlistReqResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiHostlistReqResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiHostlistReqResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


