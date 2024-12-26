# V0041OpenapiWckeyRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedWckeys** | **[]string** | deleted wckeys | 
**Meta** | Pointer to [**V0041OpenapiSharesRespMeta**](V0041OpenapiSharesRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSharesRespErrorsInner**](V0041OpenapiSharesRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSharesRespWarningsInner**](V0041OpenapiSharesRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiWckeyRemovedResp

`func NewV0041OpenapiWckeyRemovedResp(deletedWckeys []string, ) *V0041OpenapiWckeyRemovedResp`

NewV0041OpenapiWckeyRemovedResp instantiates a new V0041OpenapiWckeyRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiWckeyRemovedRespWithDefaults

`func NewV0041OpenapiWckeyRemovedRespWithDefaults() *V0041OpenapiWckeyRemovedResp`

NewV0041OpenapiWckeyRemovedRespWithDefaults instantiates a new V0041OpenapiWckeyRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedWckeys

`func (o *V0041OpenapiWckeyRemovedResp) GetDeletedWckeys() []string`

GetDeletedWckeys returns the DeletedWckeys field if non-nil, zero value otherwise.

### GetDeletedWckeysOk

`func (o *V0041OpenapiWckeyRemovedResp) GetDeletedWckeysOk() (*[]string, bool)`

GetDeletedWckeysOk returns a tuple with the DeletedWckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedWckeys

`func (o *V0041OpenapiWckeyRemovedResp) SetDeletedWckeys(v []string)`

SetDeletedWckeys sets DeletedWckeys field to given value.


### GetMeta

`func (o *V0041OpenapiWckeyRemovedResp) GetMeta() V0041OpenapiSharesRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiWckeyRemovedResp) GetMetaOk() (*V0041OpenapiSharesRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiWckeyRemovedResp) SetMeta(v V0041OpenapiSharesRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiWckeyRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiWckeyRemovedResp) GetErrors() []V0041OpenapiSharesRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiWckeyRemovedResp) GetErrorsOk() (*[]V0041OpenapiSharesRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiWckeyRemovedResp) SetErrors(v []V0041OpenapiSharesRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiWckeyRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiWckeyRemovedResp) GetWarnings() []V0041OpenapiSharesRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiWckeyRemovedResp) GetWarningsOk() (*[]V0041OpenapiSharesRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiWckeyRemovedResp) SetWarnings(v []V0041OpenapiSharesRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiWckeyRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


