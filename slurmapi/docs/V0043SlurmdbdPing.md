# V0043SlurmdbdPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Target for ping | 
**Responding** | **bool** | If ping RPC responded with pong from slurmdbd | 
**Latency** | **int64** | Number of microseconds it took to successfully ping or timeout | 
**Primary** | **bool** | Is responding slurmdbd the primary controller (Is responding slurmctld the primary controller) | 

## Methods

### NewV0043SlurmdbdPing

`func NewV0043SlurmdbdPing(hostname string, responding bool, latency int64, primary bool, ) *V0043SlurmdbdPing`

NewV0043SlurmdbdPing instantiates a new V0043SlurmdbdPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043SlurmdbdPingWithDefaults

`func NewV0043SlurmdbdPingWithDefaults() *V0043SlurmdbdPing`

NewV0043SlurmdbdPingWithDefaults instantiates a new V0043SlurmdbdPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V0043SlurmdbdPing) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0043SlurmdbdPing) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0043SlurmdbdPing) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetResponding

`func (o *V0043SlurmdbdPing) GetResponding() bool`

GetResponding returns the Responding field if non-nil, zero value otherwise.

### GetRespondingOk

`func (o *V0043SlurmdbdPing) GetRespondingOk() (*bool, bool)`

GetRespondingOk returns a tuple with the Responding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponding

`func (o *V0043SlurmdbdPing) SetResponding(v bool)`

SetResponding sets Responding field to given value.


### GetLatency

`func (o *V0043SlurmdbdPing) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V0043SlurmdbdPing) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V0043SlurmdbdPing) SetLatency(v int64)`

SetLatency sets Latency field to given value.


### GetPrimary

`func (o *V0043SlurmdbdPing) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *V0043SlurmdbdPing) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *V0043SlurmdbdPing) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


