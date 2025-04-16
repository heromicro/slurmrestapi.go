# V0042OpenapiWckeyRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedWckeys** | **[]string** |  | 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiWckeyRemovedResp

`func NewV0042OpenapiWckeyRemovedResp(deletedWckeys []string, ) *V0042OpenapiWckeyRemovedResp`

NewV0042OpenapiWckeyRemovedResp instantiates a new V0042OpenapiWckeyRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiWckeyRemovedRespWithDefaults

`func NewV0042OpenapiWckeyRemovedRespWithDefaults() *V0042OpenapiWckeyRemovedResp`

NewV0042OpenapiWckeyRemovedRespWithDefaults instantiates a new V0042OpenapiWckeyRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedWckeys

`func (o *V0042OpenapiWckeyRemovedResp) GetDeletedWckeys() []string`

GetDeletedWckeys returns the DeletedWckeys field if non-nil, zero value otherwise.

### GetDeletedWckeysOk

`func (o *V0042OpenapiWckeyRemovedResp) GetDeletedWckeysOk() (*[]string, bool)`

GetDeletedWckeysOk returns a tuple with the DeletedWckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedWckeys

`func (o *V0042OpenapiWckeyRemovedResp) SetDeletedWckeys(v []string)`

SetDeletedWckeys sets DeletedWckeys field to given value.


### GetMeta

`func (o *V0042OpenapiWckeyRemovedResp) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiWckeyRemovedResp) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiWckeyRemovedResp) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiWckeyRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiWckeyRemovedResp) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiWckeyRemovedResp) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiWckeyRemovedResp) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiWckeyRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiWckeyRemovedResp) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiWckeyRemovedResp) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiWckeyRemovedResp) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiWckeyRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


