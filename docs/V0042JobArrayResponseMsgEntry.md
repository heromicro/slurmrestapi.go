# V0042JobArrayResponseMsgEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** | Job ID for updated job | [optional] 
**StepId** | Pointer to **string** | Step ID for updated job | [optional] 
**Error** | Pointer to **string** | Verbose update status or error | [optional] 
**ErrorCode** | Pointer to **int32** | Verbose update status or error | [optional] 
**Why** | Pointer to **string** | Update response message | [optional] 

## Methods

### NewV0042JobArrayResponseMsgEntry

`func NewV0042JobArrayResponseMsgEntry() *V0042JobArrayResponseMsgEntry`

NewV0042JobArrayResponseMsgEntry instantiates a new V0042JobArrayResponseMsgEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobArrayResponseMsgEntryWithDefaults

`func NewV0042JobArrayResponseMsgEntryWithDefaults() *V0042JobArrayResponseMsgEntry`

NewV0042JobArrayResponseMsgEntryWithDefaults instantiates a new V0042JobArrayResponseMsgEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0042JobArrayResponseMsgEntry) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0042JobArrayResponseMsgEntry) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0042JobArrayResponseMsgEntry) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0042JobArrayResponseMsgEntry) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0042JobArrayResponseMsgEntry) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0042JobArrayResponseMsgEntry) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0042JobArrayResponseMsgEntry) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0042JobArrayResponseMsgEntry) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetError

`func (o *V0042JobArrayResponseMsgEntry) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0042JobArrayResponseMsgEntry) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0042JobArrayResponseMsgEntry) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0042JobArrayResponseMsgEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *V0042JobArrayResponseMsgEntry) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *V0042JobArrayResponseMsgEntry) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *V0042JobArrayResponseMsgEntry) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *V0042JobArrayResponseMsgEntry) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetWhy

`func (o *V0042JobArrayResponseMsgEntry) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *V0042JobArrayResponseMsgEntry) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *V0042JobArrayResponseMsgEntry) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *V0042JobArrayResponseMsgEntry) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


