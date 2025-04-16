# V0042OpenapiWckeyResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wckeys** | [**[]V0042Wckey**](V0042Wckey.md) |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiWckeyResp

`func NewV0042OpenapiWckeyResp(wckeys []V0042Wckey, ) *V0042OpenapiWckeyResp`

NewV0042OpenapiWckeyResp instantiates a new V0042OpenapiWckeyResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiWckeyRespWithDefaults

`func NewV0042OpenapiWckeyRespWithDefaults() *V0042OpenapiWckeyResp`

NewV0042OpenapiWckeyRespWithDefaults instantiates a new V0042OpenapiWckeyResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWckeys

`func (o *V0042OpenapiWckeyResp) GetWckeys() []V0042Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0042OpenapiWckeyResp) GetWckeysOk() (*[]V0042Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0042OpenapiWckeyResp) SetWckeys(v []V0042Wckey)`

SetWckeys sets Wckeys field to given value.


### GetMeta

`func (o *V0042OpenapiWckeyResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiWckeyResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiWckeyResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiWckeyResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiWckeyResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiWckeyResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiWckeyResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiWckeyResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiWckeyResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiWckeyResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiWckeyResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiWckeyResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


