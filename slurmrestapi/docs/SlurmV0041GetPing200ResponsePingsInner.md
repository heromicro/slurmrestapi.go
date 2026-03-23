# SlurmV0041GetPing200ResponsePingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Target for ping | [optional] 
**Pinged** | Pointer to **string** | Ping result | [optional] 
**Latency** | Pointer to **int64** | Number of microseconds it took to successfully ping or timeout | [optional] 
**Mode** | Pointer to **string** | The operating mode of the responding slurmctld | [optional] 

## Methods

### NewSlurmV0041GetPing200ResponsePingsInner

`func NewSlurmV0041GetPing200ResponsePingsInner() *SlurmV0041GetPing200ResponsePingsInner`

NewSlurmV0041GetPing200ResponsePingsInner instantiates a new SlurmV0041GetPing200ResponsePingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetPing200ResponsePingsInnerWithDefaults

`func NewSlurmV0041GetPing200ResponsePingsInnerWithDefaults() *SlurmV0041GetPing200ResponsePingsInner`

NewSlurmV0041GetPing200ResponsePingsInnerWithDefaults instantiates a new SlurmV0041GetPing200ResponsePingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SlurmV0041GetPing200ResponsePingsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SlurmV0041GetPing200ResponsePingsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPinged

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetPinged() string`

GetPinged returns the Pinged field if non-nil, zero value otherwise.

### GetPingedOk

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetPingedOk() (*string, bool)`

GetPingedOk returns a tuple with the Pinged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinged

`func (o *SlurmV0041GetPing200ResponsePingsInner) SetPinged(v string)`

SetPinged sets Pinged field to given value.

### HasPinged

`func (o *SlurmV0041GetPing200ResponsePingsInner) HasPinged() bool`

HasPinged returns a boolean if a field has been set.

### GetLatency

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *SlurmV0041GetPing200ResponsePingsInner) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *SlurmV0041GetPing200ResponsePingsInner) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMode

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SlurmV0041GetPing200ResponsePingsInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SlurmV0041GetPing200ResponsePingsInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SlurmV0041GetPing200ResponsePingsInner) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


