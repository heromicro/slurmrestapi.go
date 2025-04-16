# V0042OpenapiJobPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]V0042JobArrayResponseMsgEntry**](V0042JobArrayResponseMsgEntry.md) |  | [optional] 
**Meta** | Pointer to [**V0042OpenapiMeta**](V0042OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0042OpenapiError**](V0042OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0042OpenapiWarning**](V0042OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0042OpenapiJobPostResponse

`func NewV0042OpenapiJobPostResponse() *V0042OpenapiJobPostResponse`

NewV0042OpenapiJobPostResponse instantiates a new V0042OpenapiJobPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiJobPostResponseWithDefaults

`func NewV0042OpenapiJobPostResponseWithDefaults() *V0042OpenapiJobPostResponse`

NewV0042OpenapiJobPostResponseWithDefaults instantiates a new V0042OpenapiJobPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V0042OpenapiJobPostResponse) GetResults() []V0042JobArrayResponseMsgEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V0042OpenapiJobPostResponse) GetResultsOk() (*[]V0042JobArrayResponseMsgEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V0042OpenapiJobPostResponse) SetResults(v []V0042JobArrayResponseMsgEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *V0042OpenapiJobPostResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMeta

`func (o *V0042OpenapiJobPostResponse) GetMeta() V0042OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0042OpenapiJobPostResponse) GetMetaOk() (*V0042OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0042OpenapiJobPostResponse) SetMeta(v V0042OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0042OpenapiJobPostResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0042OpenapiJobPostResponse) GetErrors() []V0042OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0042OpenapiJobPostResponse) GetErrorsOk() (*[]V0042OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0042OpenapiJobPostResponse) SetErrors(v []V0042OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0042OpenapiJobPostResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0042OpenapiJobPostResponse) GetWarnings() []V0042OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0042OpenapiJobPostResponse) GetWarningsOk() (*[]V0042OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0042OpenapiJobPostResponse) SetWarnings(v []V0042OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0042OpenapiJobPostResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


