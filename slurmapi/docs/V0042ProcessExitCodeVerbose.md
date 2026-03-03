# V0042ProcessExitCodeVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** |  | [optional] 
**ReturnCode** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Signal** | Pointer to [**V0042ProcessExitCodeVerboseSignal**](V0042ProcessExitCodeVerboseSignal.md) |  | [optional] 

## Methods

### NewV0042ProcessExitCodeVerbose

`func NewV0042ProcessExitCodeVerbose() *V0042ProcessExitCodeVerbose`

NewV0042ProcessExitCodeVerbose instantiates a new V0042ProcessExitCodeVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042ProcessExitCodeVerboseWithDefaults

`func NewV0042ProcessExitCodeVerboseWithDefaults() *V0042ProcessExitCodeVerbose`

NewV0042ProcessExitCodeVerboseWithDefaults instantiates a new V0042ProcessExitCodeVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0042ProcessExitCodeVerbose) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0042ProcessExitCodeVerbose) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0042ProcessExitCodeVerbose) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0042ProcessExitCodeVerbose) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *V0042ProcessExitCodeVerbose) GetReturnCode() V0042Uint32NoValStruct`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *V0042ProcessExitCodeVerbose) GetReturnCodeOk() (*V0042Uint32NoValStruct, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *V0042ProcessExitCodeVerbose) SetReturnCode(v V0042Uint32NoValStruct)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *V0042ProcessExitCodeVerbose) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *V0042ProcessExitCodeVerbose) GetSignal() V0042ProcessExitCodeVerboseSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0042ProcessExitCodeVerbose) GetSignalOk() (*V0042ProcessExitCodeVerboseSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0042ProcessExitCodeVerbose) SetSignal(v V0042ProcessExitCodeVerboseSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0042ProcessExitCodeVerbose) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


