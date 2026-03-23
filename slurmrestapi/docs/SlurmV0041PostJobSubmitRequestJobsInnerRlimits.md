# SlurmV0041PostJobSubmitRequestJobsInnerRlimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCpu**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCpu.md) |  | [optional] 
**Fsize** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsFsize**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsFsize.md) |  | [optional] 
**Data** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsData**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsData.md) |  | [optional] 
**Stack** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsStack**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsStack.md) |  | [optional] 
**Core** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCore**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCore.md) |  | [optional] 
**Rss** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsRss**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsRss.md) |  | [optional] 
**Nproc** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNproc**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNproc.md) |  | [optional] 
**Nofile** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNofile**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNofile.md) |  | [optional] 
**Memlock** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsMemlock**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsMemlock.md) |  | [optional] 
**As** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimitsAs**](SlurmV0041PostJobSubmitRequestJobsInnerRlimitsAs.md) |  | [optional] 

## Methods

### NewSlurmV0041PostJobSubmitRequestJobsInnerRlimits

`func NewSlurmV0041PostJobSubmitRequestJobsInnerRlimits() *SlurmV0041PostJobSubmitRequestJobsInnerRlimits`

NewSlurmV0041PostJobSubmitRequestJobsInnerRlimits instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerRlimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmitRequestJobsInnerRlimitsWithDefaults

`func NewSlurmV0041PostJobSubmitRequestJobsInnerRlimitsWithDefaults() *SlurmV0041PostJobSubmitRequestJobsInnerRlimits`

NewSlurmV0041PostJobSubmitRequestJobsInnerRlimitsWithDefaults instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerRlimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetCpu() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetCpuOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetCpu(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetFsize

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetFsize() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsFsize`

GetFsize returns the Fsize field if non-nil, zero value otherwise.

### GetFsizeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetFsizeOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsFsize, bool)`

GetFsizeOk returns a tuple with the Fsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsize

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetFsize(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsFsize)`

SetFsize sets Fsize field to given value.

### HasFsize

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasFsize() bool`

HasFsize returns a boolean if a field has been set.

### GetData

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetData() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetDataOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetData(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsData)`

SetData sets Data field to given value.

### HasData

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStack

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetStack() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetStackOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetStack(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetCore() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCore`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetCoreOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCore, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetCore(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsCore)`

SetCore sets Core field to given value.

### HasCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetRss

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetRss() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsRss`

GetRss returns the Rss field if non-nil, zero value otherwise.

### GetRssOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetRssOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsRss, bool)`

GetRssOk returns a tuple with the Rss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRss

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetRss(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsRss)`

SetRss sets Rss field to given value.

### HasRss

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasRss() bool`

HasRss returns a boolean if a field has been set.

### GetNproc

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetNproc() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNproc`

GetNproc returns the Nproc field if non-nil, zero value otherwise.

### GetNprocOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetNprocOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNproc, bool)`

GetNprocOk returns a tuple with the Nproc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNproc

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetNproc(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNproc)`

SetNproc sets Nproc field to given value.

### HasNproc

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasNproc() bool`

HasNproc returns a boolean if a field has been set.

### GetNofile

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetNofile() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNofile`

GetNofile returns the Nofile field if non-nil, zero value otherwise.

### GetNofileOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetNofileOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNofile, bool)`

GetNofileOk returns a tuple with the Nofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNofile

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetNofile(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsNofile)`

SetNofile sets Nofile field to given value.

### HasNofile

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasNofile() bool`

HasNofile returns a boolean if a field has been set.

### GetMemlock

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetMemlock() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsMemlock`

GetMemlock returns the Memlock field if non-nil, zero value otherwise.

### GetMemlockOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetMemlockOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsMemlock, bool)`

GetMemlockOk returns a tuple with the Memlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemlock

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetMemlock(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsMemlock)`

SetMemlock sets Memlock field to given value.

### HasMemlock

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasMemlock() bool`

HasMemlock returns a boolean if a field has been set.

### GetAs

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetAs() SlurmV0041PostJobSubmitRequestJobsInnerRlimitsAs`

GetAs returns the As field if non-nil, zero value otherwise.

### GetAsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) GetAsOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimitsAs, bool)`

GetAsOk returns a tuple with the As field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAs

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) SetAs(v SlurmV0041PostJobSubmitRequestJobsInnerRlimitsAs)`

SetAs sets As field to given value.

### HasAs

`func (o *SlurmV0041PostJobSubmitRequestJobsInnerRlimits) HasAs() bool`

HasAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


