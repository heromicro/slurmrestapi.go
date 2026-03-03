# V0042ControllerPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Target for ping | [optional] 
**Pinged** | Pointer to **string** | Ping result | [optional] 
**Responding** | **bool** | If ping RPC responded with pong from controller | 
**Latency** | Pointer to **int64** | Number of microseconds it took to successfully ping or timeout | [optional] 
**Mode** | Pointer to **string** | The operating mode of the responding slurmctld | [optional] 
**Primary** | **bool** | Is responding slurmctld the primary controller | 

## Methods

### NewV0042ControllerPing

`func NewV0042ControllerPing(responding bool, primary bool, ) *V0042ControllerPing`

NewV0042ControllerPing instantiates a new V0042ControllerPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042ControllerPingWithDefaults

`func NewV0042ControllerPingWithDefaults() *V0042ControllerPing`

NewV0042ControllerPingWithDefaults instantiates a new V0042ControllerPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V0042ControllerPing) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0042ControllerPing) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0042ControllerPing) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0042ControllerPing) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPinged

`func (o *V0042ControllerPing) GetPinged() string`

GetPinged returns the Pinged field if non-nil, zero value otherwise.

### GetPingedOk

`func (o *V0042ControllerPing) GetPingedOk() (*string, bool)`

GetPingedOk returns a tuple with the Pinged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinged

`func (o *V0042ControllerPing) SetPinged(v string)`

SetPinged sets Pinged field to given value.

### HasPinged

`func (o *V0042ControllerPing) HasPinged() bool`

HasPinged returns a boolean if a field has been set.

### GetResponding

`func (o *V0042ControllerPing) GetResponding() bool`

GetResponding returns the Responding field if non-nil, zero value otherwise.

### GetRespondingOk

`func (o *V0042ControllerPing) GetRespondingOk() (*bool, bool)`

GetRespondingOk returns a tuple with the Responding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponding

`func (o *V0042ControllerPing) SetResponding(v bool)`

SetResponding sets Responding field to given value.


### GetLatency

`func (o *V0042ControllerPing) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V0042ControllerPing) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V0042ControllerPing) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *V0042ControllerPing) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMode

`func (o *V0042ControllerPing) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0042ControllerPing) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0042ControllerPing) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0042ControllerPing) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPrimary

`func (o *V0042ControllerPing) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *V0042ControllerPing) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *V0042ControllerPing) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


