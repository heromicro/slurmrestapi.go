# V0039StatsMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartsPacked** | Pointer to **int32** |  | [optional] 
**ReqTime** | Pointer to **int64** |  | [optional] 
**ReqTimeStart** | Pointer to **int64** |  | [optional] 
**ServerThreadCount** | Pointer to **int32** |  | [optional] 
**AgentQueueSize** | Pointer to **int32** |  | [optional] 
**AgentCount** | Pointer to **int32** |  | [optional] 
**AgentThreadCount** | Pointer to **int32** |  | [optional] 
**DbdAgentQueueSize** | Pointer to **int32** |  | [optional] 
**GettimeofdayLatency** | Pointer to **int32** |  | [optional] 
**ScheduleCycleMax** | Pointer to **int32** |  | [optional] 
**ScheduleCycleLast** | Pointer to **int32** |  | [optional] 
**ScheduleCycleTotal** | Pointer to **int32** |  | [optional] 
**ScheduleCycleMean** | Pointer to **int64** |  | [optional] 
**ScheduleCycleMeanDepth** | Pointer to **int64** |  | [optional] 
**ScheduleCyclePerMinute** | Pointer to **int64** |  | [optional] 
**ScheduleQueueLength** | Pointer to **int32** |  | [optional] 
**JobsSubmitted** | Pointer to **int32** |  | [optional] 
**JobsStarted** | Pointer to **int32** |  | [optional] 
**JobsCompleted** | Pointer to **int32** |  | [optional] 
**JobsCanceled** | Pointer to **int32** |  | [optional] 
**JobsFailed** | Pointer to **int32** |  | [optional] 
**JobsPending** | Pointer to **int32** |  | [optional] 
**JobsRunning** | Pointer to **int32** |  | [optional] 
**JobStatesTs** | Pointer to **int64** |  | [optional] 
**BfBackfilledJobs** | Pointer to **int32** |  | [optional] 
**BfLastBackfilledJobs** | Pointer to **int32** |  | [optional] 
**BfBackfilledHetJobs** | Pointer to **int32** |  | [optional] 
**BfCycleCounter** | Pointer to **int32** |  | [optional] 
**BfCycleMean** | Pointer to **int64** |  | [optional] 
**BfDepthMean** | Pointer to **int64** |  | [optional] 
**BfDepthMeanTry** | Pointer to **int64** |  | [optional] 
**BfCycleSum** | Pointer to **int64** |  | [optional] 
**BfCycleLast** | Pointer to **int32** |  | [optional] 
**BfLastDepth** | Pointer to **int32** |  | [optional] 
**BfLastDepthTry** | Pointer to **int32** |  | [optional] 
**BfDepthSum** | Pointer to **int32** |  | [optional] 
**BfDepthTrySum** | Pointer to **int32** |  | [optional] 
**BfQueueLen** | Pointer to **int32** |  | [optional] 
**BfQueueLenMean** | Pointer to **int64** |  | [optional] 
**BfQueueLenSum** | Pointer to **int32** |  | [optional] 
**BfTableSize** | Pointer to **int32** |  | [optional] 
**BfTableSizeMean** | Pointer to **int64** |  | [optional] 
**BfWhenLastCycle** | Pointer to **int64** |  | [optional] 
**BfActive** | Pointer to **bool** |  | [optional] 
**RpcsByMessageType** | Pointer to [**[]V0039StatsMsgRpcsByTypeInner**](V0039StatsMsgRpcsByTypeInner.md) | RPCs by message type | [optional] 
**RpcsByUser** | Pointer to [**[]V0039StatsMsgRpcsByUserInner**](V0039StatsMsgRpcsByUserInner.md) | RPCs by user | [optional] 

## Methods

### NewV0039StatsMsg

`func NewV0039StatsMsg() *V0039StatsMsg`

NewV0039StatsMsg instantiates a new V0039StatsMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039StatsMsgWithDefaults

`func NewV0039StatsMsgWithDefaults() *V0039StatsMsg`

NewV0039StatsMsgWithDefaults instantiates a new V0039StatsMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartsPacked

`func (o *V0039StatsMsg) GetPartsPacked() int32`

GetPartsPacked returns the PartsPacked field if non-nil, zero value otherwise.

### GetPartsPackedOk

`func (o *V0039StatsMsg) GetPartsPackedOk() (*int32, bool)`

GetPartsPackedOk returns a tuple with the PartsPacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPacked

`func (o *V0039StatsMsg) SetPartsPacked(v int32)`

SetPartsPacked sets PartsPacked field to given value.

### HasPartsPacked

`func (o *V0039StatsMsg) HasPartsPacked() bool`

HasPartsPacked returns a boolean if a field has been set.

### GetReqTime

`func (o *V0039StatsMsg) GetReqTime() int64`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *V0039StatsMsg) GetReqTimeOk() (*int64, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *V0039StatsMsg) SetReqTime(v int64)`

SetReqTime sets ReqTime field to given value.

### HasReqTime

`func (o *V0039StatsMsg) HasReqTime() bool`

HasReqTime returns a boolean if a field has been set.

### GetReqTimeStart

`func (o *V0039StatsMsg) GetReqTimeStart() int64`

GetReqTimeStart returns the ReqTimeStart field if non-nil, zero value otherwise.

### GetReqTimeStartOk

`func (o *V0039StatsMsg) GetReqTimeStartOk() (*int64, bool)`

GetReqTimeStartOk returns a tuple with the ReqTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeStart

`func (o *V0039StatsMsg) SetReqTimeStart(v int64)`

SetReqTimeStart sets ReqTimeStart field to given value.

### HasReqTimeStart

`func (o *V0039StatsMsg) HasReqTimeStart() bool`

HasReqTimeStart returns a boolean if a field has been set.

### GetServerThreadCount

`func (o *V0039StatsMsg) GetServerThreadCount() int32`

GetServerThreadCount returns the ServerThreadCount field if non-nil, zero value otherwise.

### GetServerThreadCountOk

`func (o *V0039StatsMsg) GetServerThreadCountOk() (*int32, bool)`

GetServerThreadCountOk returns a tuple with the ServerThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThreadCount

`func (o *V0039StatsMsg) SetServerThreadCount(v int32)`

SetServerThreadCount sets ServerThreadCount field to given value.

### HasServerThreadCount

`func (o *V0039StatsMsg) HasServerThreadCount() bool`

HasServerThreadCount returns a boolean if a field has been set.

### GetAgentQueueSize

`func (o *V0039StatsMsg) GetAgentQueueSize() int32`

GetAgentQueueSize returns the AgentQueueSize field if non-nil, zero value otherwise.

### GetAgentQueueSizeOk

`func (o *V0039StatsMsg) GetAgentQueueSizeOk() (*int32, bool)`

GetAgentQueueSizeOk returns a tuple with the AgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentQueueSize

`func (o *V0039StatsMsg) SetAgentQueueSize(v int32)`

SetAgentQueueSize sets AgentQueueSize field to given value.

### HasAgentQueueSize

`func (o *V0039StatsMsg) HasAgentQueueSize() bool`

HasAgentQueueSize returns a boolean if a field has been set.

### GetAgentCount

`func (o *V0039StatsMsg) GetAgentCount() int32`

GetAgentCount returns the AgentCount field if non-nil, zero value otherwise.

### GetAgentCountOk

`func (o *V0039StatsMsg) GetAgentCountOk() (*int32, bool)`

GetAgentCountOk returns a tuple with the AgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCount

`func (o *V0039StatsMsg) SetAgentCount(v int32)`

SetAgentCount sets AgentCount field to given value.

### HasAgentCount

`func (o *V0039StatsMsg) HasAgentCount() bool`

HasAgentCount returns a boolean if a field has been set.

### GetAgentThreadCount

`func (o *V0039StatsMsg) GetAgentThreadCount() int32`

GetAgentThreadCount returns the AgentThreadCount field if non-nil, zero value otherwise.

### GetAgentThreadCountOk

`func (o *V0039StatsMsg) GetAgentThreadCountOk() (*int32, bool)`

GetAgentThreadCountOk returns a tuple with the AgentThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentThreadCount

`func (o *V0039StatsMsg) SetAgentThreadCount(v int32)`

SetAgentThreadCount sets AgentThreadCount field to given value.

### HasAgentThreadCount

`func (o *V0039StatsMsg) HasAgentThreadCount() bool`

HasAgentThreadCount returns a boolean if a field has been set.

### GetDbdAgentQueueSize

`func (o *V0039StatsMsg) GetDbdAgentQueueSize() int32`

GetDbdAgentQueueSize returns the DbdAgentQueueSize field if non-nil, zero value otherwise.

### GetDbdAgentQueueSizeOk

`func (o *V0039StatsMsg) GetDbdAgentQueueSizeOk() (*int32, bool)`

GetDbdAgentQueueSizeOk returns a tuple with the DbdAgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbdAgentQueueSize

`func (o *V0039StatsMsg) SetDbdAgentQueueSize(v int32)`

SetDbdAgentQueueSize sets DbdAgentQueueSize field to given value.

### HasDbdAgentQueueSize

`func (o *V0039StatsMsg) HasDbdAgentQueueSize() bool`

HasDbdAgentQueueSize returns a boolean if a field has been set.

### GetGettimeofdayLatency

`func (o *V0039StatsMsg) GetGettimeofdayLatency() int32`

GetGettimeofdayLatency returns the GettimeofdayLatency field if non-nil, zero value otherwise.

### GetGettimeofdayLatencyOk

`func (o *V0039StatsMsg) GetGettimeofdayLatencyOk() (*int32, bool)`

GetGettimeofdayLatencyOk returns a tuple with the GettimeofdayLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettimeofdayLatency

`func (o *V0039StatsMsg) SetGettimeofdayLatency(v int32)`

SetGettimeofdayLatency sets GettimeofdayLatency field to given value.

### HasGettimeofdayLatency

`func (o *V0039StatsMsg) HasGettimeofdayLatency() bool`

HasGettimeofdayLatency returns a boolean if a field has been set.

### GetScheduleCycleMax

`func (o *V0039StatsMsg) GetScheduleCycleMax() int32`

GetScheduleCycleMax returns the ScheduleCycleMax field if non-nil, zero value otherwise.

### GetScheduleCycleMaxOk

`func (o *V0039StatsMsg) GetScheduleCycleMaxOk() (*int32, bool)`

GetScheduleCycleMaxOk returns a tuple with the ScheduleCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMax

`func (o *V0039StatsMsg) SetScheduleCycleMax(v int32)`

SetScheduleCycleMax sets ScheduleCycleMax field to given value.

### HasScheduleCycleMax

`func (o *V0039StatsMsg) HasScheduleCycleMax() bool`

HasScheduleCycleMax returns a boolean if a field has been set.

### GetScheduleCycleLast

`func (o *V0039StatsMsg) GetScheduleCycleLast() int32`

GetScheduleCycleLast returns the ScheduleCycleLast field if non-nil, zero value otherwise.

### GetScheduleCycleLastOk

`func (o *V0039StatsMsg) GetScheduleCycleLastOk() (*int32, bool)`

GetScheduleCycleLastOk returns a tuple with the ScheduleCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleLast

`func (o *V0039StatsMsg) SetScheduleCycleLast(v int32)`

SetScheduleCycleLast sets ScheduleCycleLast field to given value.

### HasScheduleCycleLast

`func (o *V0039StatsMsg) HasScheduleCycleLast() bool`

HasScheduleCycleLast returns a boolean if a field has been set.

### GetScheduleCycleTotal

`func (o *V0039StatsMsg) GetScheduleCycleTotal() int32`

GetScheduleCycleTotal returns the ScheduleCycleTotal field if non-nil, zero value otherwise.

### GetScheduleCycleTotalOk

`func (o *V0039StatsMsg) GetScheduleCycleTotalOk() (*int32, bool)`

GetScheduleCycleTotalOk returns a tuple with the ScheduleCycleTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleTotal

`func (o *V0039StatsMsg) SetScheduleCycleTotal(v int32)`

SetScheduleCycleTotal sets ScheduleCycleTotal field to given value.

### HasScheduleCycleTotal

`func (o *V0039StatsMsg) HasScheduleCycleTotal() bool`

HasScheduleCycleTotal returns a boolean if a field has been set.

### GetScheduleCycleMean

`func (o *V0039StatsMsg) GetScheduleCycleMean() int64`

GetScheduleCycleMean returns the ScheduleCycleMean field if non-nil, zero value otherwise.

### GetScheduleCycleMeanOk

`func (o *V0039StatsMsg) GetScheduleCycleMeanOk() (*int64, bool)`

GetScheduleCycleMeanOk returns a tuple with the ScheduleCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMean

`func (o *V0039StatsMsg) SetScheduleCycleMean(v int64)`

SetScheduleCycleMean sets ScheduleCycleMean field to given value.

### HasScheduleCycleMean

`func (o *V0039StatsMsg) HasScheduleCycleMean() bool`

HasScheduleCycleMean returns a boolean if a field has been set.

### GetScheduleCycleMeanDepth

`func (o *V0039StatsMsg) GetScheduleCycleMeanDepth() int64`

GetScheduleCycleMeanDepth returns the ScheduleCycleMeanDepth field if non-nil, zero value otherwise.

### GetScheduleCycleMeanDepthOk

`func (o *V0039StatsMsg) GetScheduleCycleMeanDepthOk() (*int64, bool)`

GetScheduleCycleMeanDepthOk returns a tuple with the ScheduleCycleMeanDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMeanDepth

`func (o *V0039StatsMsg) SetScheduleCycleMeanDepth(v int64)`

SetScheduleCycleMeanDepth sets ScheduleCycleMeanDepth field to given value.

### HasScheduleCycleMeanDepth

`func (o *V0039StatsMsg) HasScheduleCycleMeanDepth() bool`

HasScheduleCycleMeanDepth returns a boolean if a field has been set.

### GetScheduleCyclePerMinute

`func (o *V0039StatsMsg) GetScheduleCyclePerMinute() int64`

GetScheduleCyclePerMinute returns the ScheduleCyclePerMinute field if non-nil, zero value otherwise.

### GetScheduleCyclePerMinuteOk

`func (o *V0039StatsMsg) GetScheduleCyclePerMinuteOk() (*int64, bool)`

GetScheduleCyclePerMinuteOk returns a tuple with the ScheduleCyclePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCyclePerMinute

`func (o *V0039StatsMsg) SetScheduleCyclePerMinute(v int64)`

SetScheduleCyclePerMinute sets ScheduleCyclePerMinute field to given value.

### HasScheduleCyclePerMinute

`func (o *V0039StatsMsg) HasScheduleCyclePerMinute() bool`

HasScheduleCyclePerMinute returns a boolean if a field has been set.

### GetScheduleQueueLength

`func (o *V0039StatsMsg) GetScheduleQueueLength() int32`

GetScheduleQueueLength returns the ScheduleQueueLength field if non-nil, zero value otherwise.

### GetScheduleQueueLengthOk

`func (o *V0039StatsMsg) GetScheduleQueueLengthOk() (*int32, bool)`

GetScheduleQueueLengthOk returns a tuple with the ScheduleQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleQueueLength

`func (o *V0039StatsMsg) SetScheduleQueueLength(v int32)`

SetScheduleQueueLength sets ScheduleQueueLength field to given value.

### HasScheduleQueueLength

`func (o *V0039StatsMsg) HasScheduleQueueLength() bool`

HasScheduleQueueLength returns a boolean if a field has been set.

### GetJobsSubmitted

`func (o *V0039StatsMsg) GetJobsSubmitted() int32`

GetJobsSubmitted returns the JobsSubmitted field if non-nil, zero value otherwise.

### GetJobsSubmittedOk

`func (o *V0039StatsMsg) GetJobsSubmittedOk() (*int32, bool)`

GetJobsSubmittedOk returns a tuple with the JobsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSubmitted

`func (o *V0039StatsMsg) SetJobsSubmitted(v int32)`

SetJobsSubmitted sets JobsSubmitted field to given value.

### HasJobsSubmitted

`func (o *V0039StatsMsg) HasJobsSubmitted() bool`

HasJobsSubmitted returns a boolean if a field has been set.

### GetJobsStarted

`func (o *V0039StatsMsg) GetJobsStarted() int32`

GetJobsStarted returns the JobsStarted field if non-nil, zero value otherwise.

### GetJobsStartedOk

`func (o *V0039StatsMsg) GetJobsStartedOk() (*int32, bool)`

GetJobsStartedOk returns a tuple with the JobsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsStarted

`func (o *V0039StatsMsg) SetJobsStarted(v int32)`

SetJobsStarted sets JobsStarted field to given value.

### HasJobsStarted

`func (o *V0039StatsMsg) HasJobsStarted() bool`

HasJobsStarted returns a boolean if a field has been set.

### GetJobsCompleted

`func (o *V0039StatsMsg) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *V0039StatsMsg) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *V0039StatsMsg) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.

### HasJobsCompleted

`func (o *V0039StatsMsg) HasJobsCompleted() bool`

HasJobsCompleted returns a boolean if a field has been set.

### GetJobsCanceled

`func (o *V0039StatsMsg) GetJobsCanceled() int32`

GetJobsCanceled returns the JobsCanceled field if non-nil, zero value otherwise.

### GetJobsCanceledOk

`func (o *V0039StatsMsg) GetJobsCanceledOk() (*int32, bool)`

GetJobsCanceledOk returns a tuple with the JobsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCanceled

`func (o *V0039StatsMsg) SetJobsCanceled(v int32)`

SetJobsCanceled sets JobsCanceled field to given value.

### HasJobsCanceled

`func (o *V0039StatsMsg) HasJobsCanceled() bool`

HasJobsCanceled returns a boolean if a field has been set.

### GetJobsFailed

`func (o *V0039StatsMsg) GetJobsFailed() int32`

GetJobsFailed returns the JobsFailed field if non-nil, zero value otherwise.

### GetJobsFailedOk

`func (o *V0039StatsMsg) GetJobsFailedOk() (*int32, bool)`

GetJobsFailedOk returns a tuple with the JobsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFailed

`func (o *V0039StatsMsg) SetJobsFailed(v int32)`

SetJobsFailed sets JobsFailed field to given value.

### HasJobsFailed

`func (o *V0039StatsMsg) HasJobsFailed() bool`

HasJobsFailed returns a boolean if a field has been set.

### GetJobsPending

`func (o *V0039StatsMsg) GetJobsPending() int32`

GetJobsPending returns the JobsPending field if non-nil, zero value otherwise.

### GetJobsPendingOk

`func (o *V0039StatsMsg) GetJobsPendingOk() (*int32, bool)`

GetJobsPendingOk returns a tuple with the JobsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsPending

`func (o *V0039StatsMsg) SetJobsPending(v int32)`

SetJobsPending sets JobsPending field to given value.

### HasJobsPending

`func (o *V0039StatsMsg) HasJobsPending() bool`

HasJobsPending returns a boolean if a field has been set.

### GetJobsRunning

`func (o *V0039StatsMsg) GetJobsRunning() int32`

GetJobsRunning returns the JobsRunning field if non-nil, zero value otherwise.

### GetJobsRunningOk

`func (o *V0039StatsMsg) GetJobsRunningOk() (*int32, bool)`

GetJobsRunningOk returns a tuple with the JobsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsRunning

`func (o *V0039StatsMsg) SetJobsRunning(v int32)`

SetJobsRunning sets JobsRunning field to given value.

### HasJobsRunning

`func (o *V0039StatsMsg) HasJobsRunning() bool`

HasJobsRunning returns a boolean if a field has been set.

### GetJobStatesTs

`func (o *V0039StatsMsg) GetJobStatesTs() int64`

GetJobStatesTs returns the JobStatesTs field if non-nil, zero value otherwise.

### GetJobStatesTsOk

`func (o *V0039StatsMsg) GetJobStatesTsOk() (*int64, bool)`

GetJobStatesTsOk returns a tuple with the JobStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatesTs

`func (o *V0039StatsMsg) SetJobStatesTs(v int64)`

SetJobStatesTs sets JobStatesTs field to given value.

### HasJobStatesTs

`func (o *V0039StatsMsg) HasJobStatesTs() bool`

HasJobStatesTs returns a boolean if a field has been set.

### GetBfBackfilledJobs

`func (o *V0039StatsMsg) GetBfBackfilledJobs() int32`

GetBfBackfilledJobs returns the BfBackfilledJobs field if non-nil, zero value otherwise.

### GetBfBackfilledJobsOk

`func (o *V0039StatsMsg) GetBfBackfilledJobsOk() (*int32, bool)`

GetBfBackfilledJobsOk returns a tuple with the BfBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledJobs

`func (o *V0039StatsMsg) SetBfBackfilledJobs(v int32)`

SetBfBackfilledJobs sets BfBackfilledJobs field to given value.

### HasBfBackfilledJobs

`func (o *V0039StatsMsg) HasBfBackfilledJobs() bool`

HasBfBackfilledJobs returns a boolean if a field has been set.

### GetBfLastBackfilledJobs

`func (o *V0039StatsMsg) GetBfLastBackfilledJobs() int32`

GetBfLastBackfilledJobs returns the BfLastBackfilledJobs field if non-nil, zero value otherwise.

### GetBfLastBackfilledJobsOk

`func (o *V0039StatsMsg) GetBfLastBackfilledJobsOk() (*int32, bool)`

GetBfLastBackfilledJobsOk returns a tuple with the BfLastBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastBackfilledJobs

`func (o *V0039StatsMsg) SetBfLastBackfilledJobs(v int32)`

SetBfLastBackfilledJobs sets BfLastBackfilledJobs field to given value.

### HasBfLastBackfilledJobs

`func (o *V0039StatsMsg) HasBfLastBackfilledJobs() bool`

HasBfLastBackfilledJobs returns a boolean if a field has been set.

### GetBfBackfilledHetJobs

`func (o *V0039StatsMsg) GetBfBackfilledHetJobs() int32`

GetBfBackfilledHetJobs returns the BfBackfilledHetJobs field if non-nil, zero value otherwise.

### GetBfBackfilledHetJobsOk

`func (o *V0039StatsMsg) GetBfBackfilledHetJobsOk() (*int32, bool)`

GetBfBackfilledHetJobsOk returns a tuple with the BfBackfilledHetJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledHetJobs

`func (o *V0039StatsMsg) SetBfBackfilledHetJobs(v int32)`

SetBfBackfilledHetJobs sets BfBackfilledHetJobs field to given value.

### HasBfBackfilledHetJobs

`func (o *V0039StatsMsg) HasBfBackfilledHetJobs() bool`

HasBfBackfilledHetJobs returns a boolean if a field has been set.

### GetBfCycleCounter

`func (o *V0039StatsMsg) GetBfCycleCounter() int32`

GetBfCycleCounter returns the BfCycleCounter field if non-nil, zero value otherwise.

### GetBfCycleCounterOk

`func (o *V0039StatsMsg) GetBfCycleCounterOk() (*int32, bool)`

GetBfCycleCounterOk returns a tuple with the BfCycleCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleCounter

`func (o *V0039StatsMsg) SetBfCycleCounter(v int32)`

SetBfCycleCounter sets BfCycleCounter field to given value.

### HasBfCycleCounter

`func (o *V0039StatsMsg) HasBfCycleCounter() bool`

HasBfCycleCounter returns a boolean if a field has been set.

### GetBfCycleMean

`func (o *V0039StatsMsg) GetBfCycleMean() int64`

GetBfCycleMean returns the BfCycleMean field if non-nil, zero value otherwise.

### GetBfCycleMeanOk

`func (o *V0039StatsMsg) GetBfCycleMeanOk() (*int64, bool)`

GetBfCycleMeanOk returns a tuple with the BfCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMean

`func (o *V0039StatsMsg) SetBfCycleMean(v int64)`

SetBfCycleMean sets BfCycleMean field to given value.

### HasBfCycleMean

`func (o *V0039StatsMsg) HasBfCycleMean() bool`

HasBfCycleMean returns a boolean if a field has been set.

### GetBfDepthMean

`func (o *V0039StatsMsg) GetBfDepthMean() int64`

GetBfDepthMean returns the BfDepthMean field if non-nil, zero value otherwise.

### GetBfDepthMeanOk

`func (o *V0039StatsMsg) GetBfDepthMeanOk() (*int64, bool)`

GetBfDepthMeanOk returns a tuple with the BfDepthMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMean

`func (o *V0039StatsMsg) SetBfDepthMean(v int64)`

SetBfDepthMean sets BfDepthMean field to given value.

### HasBfDepthMean

`func (o *V0039StatsMsg) HasBfDepthMean() bool`

HasBfDepthMean returns a boolean if a field has been set.

### GetBfDepthMeanTry

`func (o *V0039StatsMsg) GetBfDepthMeanTry() int64`

GetBfDepthMeanTry returns the BfDepthMeanTry field if non-nil, zero value otherwise.

### GetBfDepthMeanTryOk

`func (o *V0039StatsMsg) GetBfDepthMeanTryOk() (*int64, bool)`

GetBfDepthMeanTryOk returns a tuple with the BfDepthMeanTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMeanTry

`func (o *V0039StatsMsg) SetBfDepthMeanTry(v int64)`

SetBfDepthMeanTry sets BfDepthMeanTry field to given value.

### HasBfDepthMeanTry

`func (o *V0039StatsMsg) HasBfDepthMeanTry() bool`

HasBfDepthMeanTry returns a boolean if a field has been set.

### GetBfCycleSum

`func (o *V0039StatsMsg) GetBfCycleSum() int64`

GetBfCycleSum returns the BfCycleSum field if non-nil, zero value otherwise.

### GetBfCycleSumOk

`func (o *V0039StatsMsg) GetBfCycleSumOk() (*int64, bool)`

GetBfCycleSumOk returns a tuple with the BfCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleSum

`func (o *V0039StatsMsg) SetBfCycleSum(v int64)`

SetBfCycleSum sets BfCycleSum field to given value.

### HasBfCycleSum

`func (o *V0039StatsMsg) HasBfCycleSum() bool`

HasBfCycleSum returns a boolean if a field has been set.

### GetBfCycleLast

`func (o *V0039StatsMsg) GetBfCycleLast() int32`

GetBfCycleLast returns the BfCycleLast field if non-nil, zero value otherwise.

### GetBfCycleLastOk

`func (o *V0039StatsMsg) GetBfCycleLastOk() (*int32, bool)`

GetBfCycleLastOk returns a tuple with the BfCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleLast

`func (o *V0039StatsMsg) SetBfCycleLast(v int32)`

SetBfCycleLast sets BfCycleLast field to given value.

### HasBfCycleLast

`func (o *V0039StatsMsg) HasBfCycleLast() bool`

HasBfCycleLast returns a boolean if a field has been set.

### GetBfLastDepth

`func (o *V0039StatsMsg) GetBfLastDepth() int32`

GetBfLastDepth returns the BfLastDepth field if non-nil, zero value otherwise.

### GetBfLastDepthOk

`func (o *V0039StatsMsg) GetBfLastDepthOk() (*int32, bool)`

GetBfLastDepthOk returns a tuple with the BfLastDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepth

`func (o *V0039StatsMsg) SetBfLastDepth(v int32)`

SetBfLastDepth sets BfLastDepth field to given value.

### HasBfLastDepth

`func (o *V0039StatsMsg) HasBfLastDepth() bool`

HasBfLastDepth returns a boolean if a field has been set.

### GetBfLastDepthTry

`func (o *V0039StatsMsg) GetBfLastDepthTry() int32`

GetBfLastDepthTry returns the BfLastDepthTry field if non-nil, zero value otherwise.

### GetBfLastDepthTryOk

`func (o *V0039StatsMsg) GetBfLastDepthTryOk() (*int32, bool)`

GetBfLastDepthTryOk returns a tuple with the BfLastDepthTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepthTry

`func (o *V0039StatsMsg) SetBfLastDepthTry(v int32)`

SetBfLastDepthTry sets BfLastDepthTry field to given value.

### HasBfLastDepthTry

`func (o *V0039StatsMsg) HasBfLastDepthTry() bool`

HasBfLastDepthTry returns a boolean if a field has been set.

### GetBfDepthSum

`func (o *V0039StatsMsg) GetBfDepthSum() int32`

GetBfDepthSum returns the BfDepthSum field if non-nil, zero value otherwise.

### GetBfDepthSumOk

`func (o *V0039StatsMsg) GetBfDepthSumOk() (*int32, bool)`

GetBfDepthSumOk returns a tuple with the BfDepthSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthSum

`func (o *V0039StatsMsg) SetBfDepthSum(v int32)`

SetBfDepthSum sets BfDepthSum field to given value.

### HasBfDepthSum

`func (o *V0039StatsMsg) HasBfDepthSum() bool`

HasBfDepthSum returns a boolean if a field has been set.

### GetBfDepthTrySum

`func (o *V0039StatsMsg) GetBfDepthTrySum() int32`

GetBfDepthTrySum returns the BfDepthTrySum field if non-nil, zero value otherwise.

### GetBfDepthTrySumOk

`func (o *V0039StatsMsg) GetBfDepthTrySumOk() (*int32, bool)`

GetBfDepthTrySumOk returns a tuple with the BfDepthTrySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthTrySum

`func (o *V0039StatsMsg) SetBfDepthTrySum(v int32)`

SetBfDepthTrySum sets BfDepthTrySum field to given value.

### HasBfDepthTrySum

`func (o *V0039StatsMsg) HasBfDepthTrySum() bool`

HasBfDepthTrySum returns a boolean if a field has been set.

### GetBfQueueLen

`func (o *V0039StatsMsg) GetBfQueueLen() int32`

GetBfQueueLen returns the BfQueueLen field if non-nil, zero value otherwise.

### GetBfQueueLenOk

`func (o *V0039StatsMsg) GetBfQueueLenOk() (*int32, bool)`

GetBfQueueLenOk returns a tuple with the BfQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLen

`func (o *V0039StatsMsg) SetBfQueueLen(v int32)`

SetBfQueueLen sets BfQueueLen field to given value.

### HasBfQueueLen

`func (o *V0039StatsMsg) HasBfQueueLen() bool`

HasBfQueueLen returns a boolean if a field has been set.

### GetBfQueueLenMean

`func (o *V0039StatsMsg) GetBfQueueLenMean() int64`

GetBfQueueLenMean returns the BfQueueLenMean field if non-nil, zero value otherwise.

### GetBfQueueLenMeanOk

`func (o *V0039StatsMsg) GetBfQueueLenMeanOk() (*int64, bool)`

GetBfQueueLenMeanOk returns a tuple with the BfQueueLenMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenMean

`func (o *V0039StatsMsg) SetBfQueueLenMean(v int64)`

SetBfQueueLenMean sets BfQueueLenMean field to given value.

### HasBfQueueLenMean

`func (o *V0039StatsMsg) HasBfQueueLenMean() bool`

HasBfQueueLenMean returns a boolean if a field has been set.

### GetBfQueueLenSum

`func (o *V0039StatsMsg) GetBfQueueLenSum() int32`

GetBfQueueLenSum returns the BfQueueLenSum field if non-nil, zero value otherwise.

### GetBfQueueLenSumOk

`func (o *V0039StatsMsg) GetBfQueueLenSumOk() (*int32, bool)`

GetBfQueueLenSumOk returns a tuple with the BfQueueLenSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenSum

`func (o *V0039StatsMsg) SetBfQueueLenSum(v int32)`

SetBfQueueLenSum sets BfQueueLenSum field to given value.

### HasBfQueueLenSum

`func (o *V0039StatsMsg) HasBfQueueLenSum() bool`

HasBfQueueLenSum returns a boolean if a field has been set.

### GetBfTableSize

`func (o *V0039StatsMsg) GetBfTableSize() int32`

GetBfTableSize returns the BfTableSize field if non-nil, zero value otherwise.

### GetBfTableSizeOk

`func (o *V0039StatsMsg) GetBfTableSizeOk() (*int32, bool)`

GetBfTableSizeOk returns a tuple with the BfTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSize

`func (o *V0039StatsMsg) SetBfTableSize(v int32)`

SetBfTableSize sets BfTableSize field to given value.

### HasBfTableSize

`func (o *V0039StatsMsg) HasBfTableSize() bool`

HasBfTableSize returns a boolean if a field has been set.

### GetBfTableSizeMean

`func (o *V0039StatsMsg) GetBfTableSizeMean() int64`

GetBfTableSizeMean returns the BfTableSizeMean field if non-nil, zero value otherwise.

### GetBfTableSizeMeanOk

`func (o *V0039StatsMsg) GetBfTableSizeMeanOk() (*int64, bool)`

GetBfTableSizeMeanOk returns a tuple with the BfTableSizeMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeMean

`func (o *V0039StatsMsg) SetBfTableSizeMean(v int64)`

SetBfTableSizeMean sets BfTableSizeMean field to given value.

### HasBfTableSizeMean

`func (o *V0039StatsMsg) HasBfTableSizeMean() bool`

HasBfTableSizeMean returns a boolean if a field has been set.

### GetBfWhenLastCycle

`func (o *V0039StatsMsg) GetBfWhenLastCycle() int64`

GetBfWhenLastCycle returns the BfWhenLastCycle field if non-nil, zero value otherwise.

### GetBfWhenLastCycleOk

`func (o *V0039StatsMsg) GetBfWhenLastCycleOk() (*int64, bool)`

GetBfWhenLastCycleOk returns a tuple with the BfWhenLastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfWhenLastCycle

`func (o *V0039StatsMsg) SetBfWhenLastCycle(v int64)`

SetBfWhenLastCycle sets BfWhenLastCycle field to given value.

### HasBfWhenLastCycle

`func (o *V0039StatsMsg) HasBfWhenLastCycle() bool`

HasBfWhenLastCycle returns a boolean if a field has been set.

### GetBfActive

`func (o *V0039StatsMsg) GetBfActive() bool`

GetBfActive returns the BfActive field if non-nil, zero value otherwise.

### GetBfActiveOk

`func (o *V0039StatsMsg) GetBfActiveOk() (*bool, bool)`

GetBfActiveOk returns a tuple with the BfActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfActive

`func (o *V0039StatsMsg) SetBfActive(v bool)`

SetBfActive sets BfActive field to given value.

### HasBfActive

`func (o *V0039StatsMsg) HasBfActive() bool`

HasBfActive returns a boolean if a field has been set.

### GetRpcsByMessageType

`func (o *V0039StatsMsg) GetRpcsByMessageType() []V0039StatsMsgRpcsByTypeInner`

GetRpcsByMessageType returns the RpcsByMessageType field if non-nil, zero value otherwise.

### GetRpcsByMessageTypeOk

`func (o *V0039StatsMsg) GetRpcsByMessageTypeOk() (*[]V0039StatsMsgRpcsByTypeInner, bool)`

GetRpcsByMessageTypeOk returns a tuple with the RpcsByMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByMessageType

`func (o *V0039StatsMsg) SetRpcsByMessageType(v []V0039StatsMsgRpcsByTypeInner)`

SetRpcsByMessageType sets RpcsByMessageType field to given value.

### HasRpcsByMessageType

`func (o *V0039StatsMsg) HasRpcsByMessageType() bool`

HasRpcsByMessageType returns a boolean if a field has been set.

### GetRpcsByUser

`func (o *V0039StatsMsg) GetRpcsByUser() []V0039StatsMsgRpcsByUserInner`

GetRpcsByUser returns the RpcsByUser field if non-nil, zero value otherwise.

### GetRpcsByUserOk

`func (o *V0039StatsMsg) GetRpcsByUserOk() (*[]V0039StatsMsgRpcsByUserInner, bool)`

GetRpcsByUserOk returns a tuple with the RpcsByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByUser

`func (o *V0039StatsMsg) SetRpcsByUser(v []V0039StatsMsgRpcsByUserInner)`

SetRpcsByUser sets RpcsByUser field to given value.

### HasRpcsByUser

`func (o *V0039StatsMsg) HasRpcsByUser() bool`

HasRpcsByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


