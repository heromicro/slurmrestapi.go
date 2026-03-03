# V0043OpenapiWckeyResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wckeys** | [**[]V0043Wckey**](V0043Wckey.md) |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiWckeyResp

`func NewV0043OpenapiWckeyResp(wckeys []V0043Wckey, ) *V0043OpenapiWckeyResp`

NewV0043OpenapiWckeyResp instantiates a new V0043OpenapiWckeyResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiWckeyRespWithDefaults

`func NewV0043OpenapiWckeyRespWithDefaults() *V0043OpenapiWckeyResp`

NewV0043OpenapiWckeyRespWithDefaults instantiates a new V0043OpenapiWckeyResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWckeys

`func (o *V0043OpenapiWckeyResp) GetWckeys() []V0043Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0043OpenapiWckeyResp) GetWckeysOk() (*[]V0043Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0043OpenapiWckeyResp) SetWckeys(v []V0043Wckey)`

SetWckeys sets Wckeys field to given value.


### GetMeta

`func (o *V0043OpenapiWckeyResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiWckeyResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiWckeyResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiWckeyResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiWckeyResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiWckeyResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiWckeyResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiWckeyResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiWckeyResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiWckeyResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiWckeyResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiWckeyResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


