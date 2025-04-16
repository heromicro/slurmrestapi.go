# V0042StepTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Elapsed time in seconds | [optional] 
**End** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**Start** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**Suspended** | Pointer to **int32** | Total time in suspended state in seconds | [optional] 
**System** | Pointer to [**V0040StepTimeSystem**](V0040StepTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**V0040StepTimeTotal**](V0040StepTimeTotal.md) |  | [optional] 
**User** | Pointer to [**V0040StepTimeUser**](V0040StepTimeUser.md) |  | [optional] 

## Methods

### NewV0042StepTime

`func NewV0042StepTime() *V0042StepTime`

NewV0042StepTime instantiates a new V0042StepTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StepTimeWithDefaults

`func NewV0042StepTimeWithDefaults() *V0042StepTime`

NewV0042StepTimeWithDefaults instantiates a new V0042StepTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0042StepTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0042StepTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0042StepTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0042StepTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEnd

`func (o *V0042StepTime) GetEnd() V0042Uint64NoValStruct`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0042StepTime) GetEndOk() (*V0042Uint64NoValStruct, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0042StepTime) SetEnd(v V0042Uint64NoValStruct)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0042StepTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *V0042StepTime) GetStart() V0042Uint64NoValStruct`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0042StepTime) GetStartOk() (*V0042Uint64NoValStruct, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0042StepTime) SetStart(v V0042Uint64NoValStruct)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0042StepTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSuspended

`func (o *V0042StepTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0042StepTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0042StepTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0042StepTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0042StepTime) GetSystem() V0040StepTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0042StepTime) GetSystemOk() (*V0040StepTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0042StepTime) SetSystem(v V0040StepTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0042StepTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *V0042StepTime) GetTotal() V0040StepTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0042StepTime) GetTotalOk() (*V0040StepTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0042StepTime) SetTotal(v V0040StepTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0042StepTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0042StepTime) GetUser() V0040StepTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0042StepTime) GetUserOk() (*V0040StepTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0042StepTime) SetUser(v V0040StepTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *V0042StepTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


