# V0042SlurmdbdPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Target for ping | 
**Responding** | **bool** | If ping RPC responded with pong from slurmdbd | 
**Latency** | **int64** | Number of microseconds it took to successfully ping or timeout | 
**Primary** | **bool** | Is responding slurmdbd the primary controller | 

## Methods

### NewV0042SlurmdbdPing

`func NewV0042SlurmdbdPing(hostname string, responding bool, latency int64, primary bool, ) *V0042SlurmdbdPing`

NewV0042SlurmdbdPing instantiates a new V0042SlurmdbdPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042SlurmdbdPingWithDefaults

`func NewV0042SlurmdbdPingWithDefaults() *V0042SlurmdbdPing`

NewV0042SlurmdbdPingWithDefaults instantiates a new V0042SlurmdbdPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V0042SlurmdbdPing) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0042SlurmdbdPing) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0042SlurmdbdPing) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetResponding

`func (o *V0042SlurmdbdPing) GetResponding() bool`

GetResponding returns the Responding field if non-nil, zero value otherwise.

### GetRespondingOk

`func (o *V0042SlurmdbdPing) GetRespondingOk() (*bool, bool)`

GetRespondingOk returns a tuple with the Responding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponding

`func (o *V0042SlurmdbdPing) SetResponding(v bool)`

SetResponding sets Responding field to given value.


### GetLatency

`func (o *V0042SlurmdbdPing) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V0042SlurmdbdPing) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V0042SlurmdbdPing) SetLatency(v int64)`

SetLatency sets Latency field to given value.


### GetPrimary

`func (o *V0042SlurmdbdPing) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *V0042SlurmdbdPing) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *V0042SlurmdbdPing) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


