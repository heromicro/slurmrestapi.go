# V0043OpenapiWckeyRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedWckeys** | **[]string** |  | 
**Meta** | Pointer to [**V0043OpenapiMeta**](V0043OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0043OpenapiError**](V0043OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0043OpenapiWarning**](V0043OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0043OpenapiWckeyRemovedResp

`func NewV0043OpenapiWckeyRemovedResp(deletedWckeys []string, ) *V0043OpenapiWckeyRemovedResp`

NewV0043OpenapiWckeyRemovedResp instantiates a new V0043OpenapiWckeyRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043OpenapiWckeyRemovedRespWithDefaults

`func NewV0043OpenapiWckeyRemovedRespWithDefaults() *V0043OpenapiWckeyRemovedResp`

NewV0043OpenapiWckeyRemovedRespWithDefaults instantiates a new V0043OpenapiWckeyRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedWckeys

`func (o *V0043OpenapiWckeyRemovedResp) GetDeletedWckeys() []string`

GetDeletedWckeys returns the DeletedWckeys field if non-nil, zero value otherwise.

### GetDeletedWckeysOk

`func (o *V0043OpenapiWckeyRemovedResp) GetDeletedWckeysOk() (*[]string, bool)`

GetDeletedWckeysOk returns a tuple with the DeletedWckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedWckeys

`func (o *V0043OpenapiWckeyRemovedResp) SetDeletedWckeys(v []string)`

SetDeletedWckeys sets DeletedWckeys field to given value.


### GetMeta

`func (o *V0043OpenapiWckeyRemovedResp) GetMeta() V0043OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0043OpenapiWckeyRemovedResp) GetMetaOk() (*V0043OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0043OpenapiWckeyRemovedResp) SetMeta(v V0043OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0043OpenapiWckeyRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0043OpenapiWckeyRemovedResp) GetErrors() []V0043OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0043OpenapiWckeyRemovedResp) GetErrorsOk() (*[]V0043OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0043OpenapiWckeyRemovedResp) SetErrors(v []V0043OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0043OpenapiWckeyRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0043OpenapiWckeyRemovedResp) GetWarnings() []V0043OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0043OpenapiWckeyRemovedResp) GetWarningsOk() (*[]V0043OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0043OpenapiWckeyRemovedResp) SetWarnings(v []V0043OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0043OpenapiWckeyRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


