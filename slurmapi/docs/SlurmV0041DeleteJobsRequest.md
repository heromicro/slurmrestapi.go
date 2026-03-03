# SlurmV0041DeleteJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Filter jobs to a specific account | [optional] 
**Flags** | Pointer to **[]string** | Filter jobs according to flags | [optional] 
**JobName** | Pointer to **string** | Filter jobs to a specific name | [optional] 
**Jobs** | Pointer to **[]string** | List of jobs to signal | [optional] 
**Partition** | Pointer to **string** | Filter jobs to a specific partition | [optional] 
**Qos** | Pointer to **string** | Filter jobs to a specific QOS | [optional] 
**Reservation** | Pointer to **string** | Filter jobs to a specific reservation | [optional] 
**Signal** | Pointer to **string** | Signal to send to jobs | [optional] 
**JobState** | Pointer to **[]string** | Filter jobs to a specific state | [optional] 
**UserId** | Pointer to **string** | Filter jobs to a specific numeric user id | [optional] 
**UserName** | Pointer to **string** | Filter jobs to a specific user name | [optional] 
**Wckey** | Pointer to **string** | Filter jobs to a specific wckey | [optional] 
**Nodes** | Pointer to **[]string** | Filter jobs to a set of nodes | [optional] 

## Methods

### NewSlurmV0041DeleteJobsRequest

`func NewSlurmV0041DeleteJobsRequest() *SlurmV0041DeleteJobsRequest`

NewSlurmV0041DeleteJobsRequest instantiates a new SlurmV0041DeleteJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041DeleteJobsRequestWithDefaults

`func NewSlurmV0041DeleteJobsRequestWithDefaults() *SlurmV0041DeleteJobsRequest`

NewSlurmV0041DeleteJobsRequestWithDefaults instantiates a new SlurmV0041DeleteJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SlurmV0041DeleteJobsRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SlurmV0041DeleteJobsRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SlurmV0041DeleteJobsRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SlurmV0041DeleteJobsRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFlags

`func (o *SlurmV0041DeleteJobsRequest) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SlurmV0041DeleteJobsRequest) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SlurmV0041DeleteJobsRequest) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SlurmV0041DeleteJobsRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetJobName

`func (o *SlurmV0041DeleteJobsRequest) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *SlurmV0041DeleteJobsRequest) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *SlurmV0041DeleteJobsRequest) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *SlurmV0041DeleteJobsRequest) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobs

`func (o *SlurmV0041DeleteJobsRequest) GetJobs() []string`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *SlurmV0041DeleteJobsRequest) GetJobsOk() (*[]string, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *SlurmV0041DeleteJobsRequest) SetJobs(v []string)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *SlurmV0041DeleteJobsRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPartition

`func (o *SlurmV0041DeleteJobsRequest) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *SlurmV0041DeleteJobsRequest) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *SlurmV0041DeleteJobsRequest) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *SlurmV0041DeleteJobsRequest) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetQos

`func (o *SlurmV0041DeleteJobsRequest) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *SlurmV0041DeleteJobsRequest) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *SlurmV0041DeleteJobsRequest) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *SlurmV0041DeleteJobsRequest) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReservation

`func (o *SlurmV0041DeleteJobsRequest) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *SlurmV0041DeleteJobsRequest) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *SlurmV0041DeleteJobsRequest) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *SlurmV0041DeleteJobsRequest) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetSignal

`func (o *SlurmV0041DeleteJobsRequest) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *SlurmV0041DeleteJobsRequest) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *SlurmV0041DeleteJobsRequest) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *SlurmV0041DeleteJobsRequest) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetJobState

`func (o *SlurmV0041DeleteJobsRequest) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *SlurmV0041DeleteJobsRequest) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *SlurmV0041DeleteJobsRequest) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *SlurmV0041DeleteJobsRequest) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetUserId

`func (o *SlurmV0041DeleteJobsRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SlurmV0041DeleteJobsRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SlurmV0041DeleteJobsRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SlurmV0041DeleteJobsRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SlurmV0041DeleteJobsRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SlurmV0041DeleteJobsRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SlurmV0041DeleteJobsRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SlurmV0041DeleteJobsRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWckey

`func (o *SlurmV0041DeleteJobsRequest) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *SlurmV0041DeleteJobsRequest) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *SlurmV0041DeleteJobsRequest) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *SlurmV0041DeleteJobsRequest) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetNodes

`func (o *SlurmV0041DeleteJobsRequest) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SlurmV0041DeleteJobsRequest) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SlurmV0041DeleteJobsRequest) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *SlurmV0041DeleteJobsRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


