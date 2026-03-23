# V0043StatsMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartsPacked** | Pointer to **int32** | Zero if only RPC statistic included | [optional] 
**ReqTime** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**ReqTimeStart** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**ServerThreadCount** | Pointer to **int32** | Number of current active slurmctld threads | [optional] 
**AgentQueueSize** | Pointer to **int32** | Number of enqueued outgoing RPC requests in an internal retry list | [optional] 
**AgentCount** | Pointer to **int32** | Number of agent threads | [optional] 
**AgentThreadCount** | Pointer to **int32** | Total number of active threads created by all agent threads | [optional] 
**DbdAgentQueueSize** | Pointer to **int32** | Number of messages for SlurmDBD that are queued | [optional] 
**GettimeofdayLatency** | Pointer to **int32** | Latency of 1000 calls to the gettimeofday() syscall in microseconds, as measured at controller startup | [optional] 
**ScheduleCycleMax** | Pointer to **int32** | Max time of any scheduling cycle in microseconds since last reset | [optional] 
**ScheduleCycleLast** | Pointer to **int32** | Time in microseconds for last scheduling cycle | [optional] 
**ScheduleCycleSum** | Pointer to **int64** | Total run time in microseconds for all scheduling cycles since last reset | [optional] 
**ScheduleCycleTotal** | Pointer to **int32** | Number of scheduling cycles since last reset | [optional] 
**ScheduleCycleMean** | Pointer to **int64** | Mean time in microseconds for all scheduling cycles since last reset | [optional] 
**ScheduleCycleMeanDepth** | Pointer to **int64** | Mean of the number of jobs processed in a scheduling cycle | [optional] 
**ScheduleCyclePerMinute** | Pointer to **int64** | Number of scheduling executions per minute | [optional] 
**ScheduleCycleDepth** | Pointer to **int32** | Total number of jobs processed in scheduling cycles | [optional] 
**ScheduleExit** | Pointer to [**V0043ScheduleExitFields**](V0043ScheduleExitFields.md) |  | [optional] 
**ScheduleQueueLength** | Pointer to **int32** | Number of jobs pending in queue | [optional] 
**JobsSubmitted** | Pointer to **int32** | Number of jobs submitted since last reset | [optional] 
**JobsStarted** | Pointer to **int32** | Number of jobs started since last reset | [optional] 
**JobsCompleted** | Pointer to **int32** | Number of jobs completed since last reset | [optional] 
**JobsCanceled** | Pointer to **int32** | Number of jobs canceled since the last reset | [optional] 
**JobsFailed** | Pointer to **int32** | Number of jobs failed due to slurmd or other internal issues since last reset | [optional] 
**JobsPending** | Pointer to **int32** | Number of jobs pending at the time of listed in job_state_ts | [optional] 
**JobsRunning** | Pointer to **int32** | Number of jobs running at the time of listed in job_state_ts | [optional] 
**JobStatesTs** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**BfBackfilledJobs** | Pointer to **int32** | Number of jobs started through backfilling since last slurm start | [optional] 
**BfLastBackfilledJobs** | Pointer to **int32** | Number of jobs started through backfilling since last reset | [optional] 
**BfBackfilledHetJobs** | Pointer to **int32** | Number of heterogeneous job components started through backfilling since last Slurm start | [optional] 
**BfCycleCounter** | Pointer to **int32** | Number of backfill scheduling cycles since last reset | [optional] 
**BfCycleMean** | Pointer to **int64** | Mean time in microseconds of backfilling scheduling cycles since last reset | [optional] 
**BfDepthMean** | Pointer to **int64** | Mean number of eligible to run jobs processed during all backfilling scheduling cycles since last reset | [optional] 
**BfDepthMeanTry** | Pointer to **int64** | The subset of Depth Mean that the backfill scheduler attempted to schedule | [optional] 
**BfCycleSum** | Pointer to **int64** | Total time in microseconds of backfilling scheduling cycles since last reset | [optional] 
**BfCycleLast** | Pointer to **int32** | Execution time in microseconds of last backfill scheduling cycle | [optional] 
**BfCycleMax** | Pointer to **int32** | Execution time in microseconds of longest backfill scheduling cycle | [optional] 
**BfExit** | Pointer to [**V0043BfExitFields**](V0043BfExitFields.md) |  | [optional] 
**BfLastDepth** | Pointer to **int32** | Number of processed jobs during last backfilling scheduling cycle | [optional] 
**BfLastDepthTry** | Pointer to **int32** | Number of processed jobs during last backfilling scheduling cycle that had a chance to start using available resources | [optional] 
**BfDepthSum** | Pointer to **int32** | Total number of jobs processed during all backfilling scheduling cycles since last reset | [optional] 
**BfDepthTrySum** | Pointer to **int32** | Subset of bf_depth_sum that the backfill scheduler attempted to schedule | [optional] 
**BfQueueLen** | Pointer to **int32** | Number of jobs pending to be processed by backfilling algorithm | [optional] 
**BfQueueLenMean** | Pointer to **int64** | Mean number of jobs pending to be processed by backfilling algorithm | [optional] 
**BfQueueLenSum** | Pointer to **int32** | Total number of jobs pending to be processed by backfilling algorithm since last reset | [optional] 
**BfTableSize** | Pointer to **int32** | Number of different time slots tested by the backfill scheduler in its last iteration | [optional] 
**BfTableSizeSum** | Pointer to **int32** | Total number of different time slots tested by the backfill scheduler | [optional] 
**BfTableSizeMean** | Pointer to **int64** | Mean number of different time slots tested by the backfill scheduler | [optional] 
**BfWhenLastCycle** | Pointer to [**V0043Uint64NoValStruct**](V0043Uint64NoValStruct.md) |  | [optional] 
**BfActive** | Pointer to **bool** | Backfill scheduler currently running | [optional] 
**RpcsByMessageType** | Pointer to [**[]V0043StatsMsgRpcType**](V0043StatsMsgRpcType.md) | RPCs by type | [optional] 
**RpcsByUser** | Pointer to [**[]V0043StatsMsgRpcUser**](V0043StatsMsgRpcUser.md) | RPCs by user | [optional] 
**PendingRpcs** | Pointer to [**[]V0043StatsMsgRpcQueue**](V0043StatsMsgRpcQueue.md) | Pending RPCs | [optional] 
**PendingRpcsByHostlist** | Pointer to [**[]V0043StatsMsgRpcDump**](V0043StatsMsgRpcDump.md) | Pending RPCs by hostlist | [optional] 

## Methods

### NewV0043StatsMsg

`func NewV0043StatsMsg() *V0043StatsMsg`

NewV0043StatsMsg instantiates a new V0043StatsMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StatsMsgWithDefaults

`func NewV0043StatsMsgWithDefaults() *V0043StatsMsg`

NewV0043StatsMsgWithDefaults instantiates a new V0043StatsMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartsPacked

`func (o *V0043StatsMsg) GetPartsPacked() int32`

GetPartsPacked returns the PartsPacked field if non-nil, zero value otherwise.

### GetPartsPackedOk

`func (o *V0043StatsMsg) GetPartsPackedOk() (*int32, bool)`

GetPartsPackedOk returns a tuple with the PartsPacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPacked

`func (o *V0043StatsMsg) SetPartsPacked(v int32)`

SetPartsPacked sets PartsPacked field to given value.

### HasPartsPacked

`func (o *V0043StatsMsg) HasPartsPacked() bool`

HasPartsPacked returns a boolean if a field has been set.

### GetReqTime

`func (o *V0043StatsMsg) GetReqTime() V0043Uint64NoValStruct`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *V0043StatsMsg) GetReqTimeOk() (*V0043Uint64NoValStruct, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *V0043StatsMsg) SetReqTime(v V0043Uint64NoValStruct)`

SetReqTime sets ReqTime field to given value.

### HasReqTime

`func (o *V0043StatsMsg) HasReqTime() bool`

HasReqTime returns a boolean if a field has been set.

### GetReqTimeStart

`func (o *V0043StatsMsg) GetReqTimeStart() V0043Uint64NoValStruct`

GetReqTimeStart returns the ReqTimeStart field if non-nil, zero value otherwise.

### GetReqTimeStartOk

`func (o *V0043StatsMsg) GetReqTimeStartOk() (*V0043Uint64NoValStruct, bool)`

GetReqTimeStartOk returns a tuple with the ReqTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeStart

`func (o *V0043StatsMsg) SetReqTimeStart(v V0043Uint64NoValStruct)`

SetReqTimeStart sets ReqTimeStart field to given value.

### HasReqTimeStart

`func (o *V0043StatsMsg) HasReqTimeStart() bool`

HasReqTimeStart returns a boolean if a field has been set.

### GetServerThreadCount

`func (o *V0043StatsMsg) GetServerThreadCount() int32`

GetServerThreadCount returns the ServerThreadCount field if non-nil, zero value otherwise.

### GetServerThreadCountOk

`func (o *V0043StatsMsg) GetServerThreadCountOk() (*int32, bool)`

GetServerThreadCountOk returns a tuple with the ServerThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThreadCount

`func (o *V0043StatsMsg) SetServerThreadCount(v int32)`

SetServerThreadCount sets ServerThreadCount field to given value.

### HasServerThreadCount

`func (o *V0043StatsMsg) HasServerThreadCount() bool`

HasServerThreadCount returns a boolean if a field has been set.

### GetAgentQueueSize

`func (o *V0043StatsMsg) GetAgentQueueSize() int32`

GetAgentQueueSize returns the AgentQueueSize field if non-nil, zero value otherwise.

### GetAgentQueueSizeOk

`func (o *V0043StatsMsg) GetAgentQueueSizeOk() (*int32, bool)`

GetAgentQueueSizeOk returns a tuple with the AgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentQueueSize

`func (o *V0043StatsMsg) SetAgentQueueSize(v int32)`

SetAgentQueueSize sets AgentQueueSize field to given value.

### HasAgentQueueSize

`func (o *V0043StatsMsg) HasAgentQueueSize() bool`

HasAgentQueueSize returns a boolean if a field has been set.

### GetAgentCount

`func (o *V0043StatsMsg) GetAgentCount() int32`

GetAgentCount returns the AgentCount field if non-nil, zero value otherwise.

### GetAgentCountOk

`func (o *V0043StatsMsg) GetAgentCountOk() (*int32, bool)`

GetAgentCountOk returns a tuple with the AgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCount

`func (o *V0043StatsMsg) SetAgentCount(v int32)`

SetAgentCount sets AgentCount field to given value.

### HasAgentCount

`func (o *V0043StatsMsg) HasAgentCount() bool`

HasAgentCount returns a boolean if a field has been set.

### GetAgentThreadCount

`func (o *V0043StatsMsg) GetAgentThreadCount() int32`

GetAgentThreadCount returns the AgentThreadCount field if non-nil, zero value otherwise.

### GetAgentThreadCountOk

`func (o *V0043StatsMsg) GetAgentThreadCountOk() (*int32, bool)`

GetAgentThreadCountOk returns a tuple with the AgentThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentThreadCount

`func (o *V0043StatsMsg) SetAgentThreadCount(v int32)`

SetAgentThreadCount sets AgentThreadCount field to given value.

### HasAgentThreadCount

`func (o *V0043StatsMsg) HasAgentThreadCount() bool`

HasAgentThreadCount returns a boolean if a field has been set.

### GetDbdAgentQueueSize

`func (o *V0043StatsMsg) GetDbdAgentQueueSize() int32`

GetDbdAgentQueueSize returns the DbdAgentQueueSize field if non-nil, zero value otherwise.

### GetDbdAgentQueueSizeOk

`func (o *V0043StatsMsg) GetDbdAgentQueueSizeOk() (*int32, bool)`

GetDbdAgentQueueSizeOk returns a tuple with the DbdAgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbdAgentQueueSize

`func (o *V0043StatsMsg) SetDbdAgentQueueSize(v int32)`

SetDbdAgentQueueSize sets DbdAgentQueueSize field to given value.

### HasDbdAgentQueueSize

`func (o *V0043StatsMsg) HasDbdAgentQueueSize() bool`

HasDbdAgentQueueSize returns a boolean if a field has been set.

### GetGettimeofdayLatency

`func (o *V0043StatsMsg) GetGettimeofdayLatency() int32`

GetGettimeofdayLatency returns the GettimeofdayLatency field if non-nil, zero value otherwise.

### GetGettimeofdayLatencyOk

`func (o *V0043StatsMsg) GetGettimeofdayLatencyOk() (*int32, bool)`

GetGettimeofdayLatencyOk returns a tuple with the GettimeofdayLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettimeofdayLatency

`func (o *V0043StatsMsg) SetGettimeofdayLatency(v int32)`

SetGettimeofdayLatency sets GettimeofdayLatency field to given value.

### HasGettimeofdayLatency

`func (o *V0043StatsMsg) HasGettimeofdayLatency() bool`

HasGettimeofdayLatency returns a boolean if a field has been set.

### GetScheduleCycleMax

`func (o *V0043StatsMsg) GetScheduleCycleMax() int32`

GetScheduleCycleMax returns the ScheduleCycleMax field if non-nil, zero value otherwise.

### GetScheduleCycleMaxOk

`func (o *V0043StatsMsg) GetScheduleCycleMaxOk() (*int32, bool)`

GetScheduleCycleMaxOk returns a tuple with the ScheduleCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMax

`func (o *V0043StatsMsg) SetScheduleCycleMax(v int32)`

SetScheduleCycleMax sets ScheduleCycleMax field to given value.

### HasScheduleCycleMax

`func (o *V0043StatsMsg) HasScheduleCycleMax() bool`

HasScheduleCycleMax returns a boolean if a field has been set.

### GetScheduleCycleLast

`func (o *V0043StatsMsg) GetScheduleCycleLast() int32`

GetScheduleCycleLast returns the ScheduleCycleLast field if non-nil, zero value otherwise.

### GetScheduleCycleLastOk

`func (o *V0043StatsMsg) GetScheduleCycleLastOk() (*int32, bool)`

GetScheduleCycleLastOk returns a tuple with the ScheduleCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleLast

`func (o *V0043StatsMsg) SetScheduleCycleLast(v int32)`

SetScheduleCycleLast sets ScheduleCycleLast field to given value.

### HasScheduleCycleLast

`func (o *V0043StatsMsg) HasScheduleCycleLast() bool`

HasScheduleCycleLast returns a boolean if a field has been set.

### GetScheduleCycleSum

`func (o *V0043StatsMsg) GetScheduleCycleSum() int64`

GetScheduleCycleSum returns the ScheduleCycleSum field if non-nil, zero value otherwise.

### GetScheduleCycleSumOk

`func (o *V0043StatsMsg) GetScheduleCycleSumOk() (*int64, bool)`

GetScheduleCycleSumOk returns a tuple with the ScheduleCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleSum

`func (o *V0043StatsMsg) SetScheduleCycleSum(v int64)`

SetScheduleCycleSum sets ScheduleCycleSum field to given value.

### HasScheduleCycleSum

`func (o *V0043StatsMsg) HasScheduleCycleSum() bool`

HasScheduleCycleSum returns a boolean if a field has been set.

### GetScheduleCycleTotal

`func (o *V0043StatsMsg) GetScheduleCycleTotal() int32`

GetScheduleCycleTotal returns the ScheduleCycleTotal field if non-nil, zero value otherwise.

### GetScheduleCycleTotalOk

`func (o *V0043StatsMsg) GetScheduleCycleTotalOk() (*int32, bool)`

GetScheduleCycleTotalOk returns a tuple with the ScheduleCycleTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleTotal

`func (o *V0043StatsMsg) SetScheduleCycleTotal(v int32)`

SetScheduleCycleTotal sets ScheduleCycleTotal field to given value.

### HasScheduleCycleTotal

`func (o *V0043StatsMsg) HasScheduleCycleTotal() bool`

HasScheduleCycleTotal returns a boolean if a field has been set.

### GetScheduleCycleMean

`func (o *V0043StatsMsg) GetScheduleCycleMean() int64`

GetScheduleCycleMean returns the ScheduleCycleMean field if non-nil, zero value otherwise.

### GetScheduleCycleMeanOk

`func (o *V0043StatsMsg) GetScheduleCycleMeanOk() (*int64, bool)`

GetScheduleCycleMeanOk returns a tuple with the ScheduleCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMean

`func (o *V0043StatsMsg) SetScheduleCycleMean(v int64)`

SetScheduleCycleMean sets ScheduleCycleMean field to given value.

### HasScheduleCycleMean

`func (o *V0043StatsMsg) HasScheduleCycleMean() bool`

HasScheduleCycleMean returns a boolean if a field has been set.

### GetScheduleCycleMeanDepth

`func (o *V0043StatsMsg) GetScheduleCycleMeanDepth() int64`

GetScheduleCycleMeanDepth returns the ScheduleCycleMeanDepth field if non-nil, zero value otherwise.

### GetScheduleCycleMeanDepthOk

`func (o *V0043StatsMsg) GetScheduleCycleMeanDepthOk() (*int64, bool)`

GetScheduleCycleMeanDepthOk returns a tuple with the ScheduleCycleMeanDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMeanDepth

`func (o *V0043StatsMsg) SetScheduleCycleMeanDepth(v int64)`

SetScheduleCycleMeanDepth sets ScheduleCycleMeanDepth field to given value.

### HasScheduleCycleMeanDepth

`func (o *V0043StatsMsg) HasScheduleCycleMeanDepth() bool`

HasScheduleCycleMeanDepth returns a boolean if a field has been set.

### GetScheduleCyclePerMinute

`func (o *V0043StatsMsg) GetScheduleCyclePerMinute() int64`

GetScheduleCyclePerMinute returns the ScheduleCyclePerMinute field if non-nil, zero value otherwise.

### GetScheduleCyclePerMinuteOk

`func (o *V0043StatsMsg) GetScheduleCyclePerMinuteOk() (*int64, bool)`

GetScheduleCyclePerMinuteOk returns a tuple with the ScheduleCyclePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCyclePerMinute

`func (o *V0043StatsMsg) SetScheduleCyclePerMinute(v int64)`

SetScheduleCyclePerMinute sets ScheduleCyclePerMinute field to given value.

### HasScheduleCyclePerMinute

`func (o *V0043StatsMsg) HasScheduleCyclePerMinute() bool`

HasScheduleCyclePerMinute returns a boolean if a field has been set.

### GetScheduleCycleDepth

`func (o *V0043StatsMsg) GetScheduleCycleDepth() int32`

GetScheduleCycleDepth returns the ScheduleCycleDepth field if non-nil, zero value otherwise.

### GetScheduleCycleDepthOk

`func (o *V0043StatsMsg) GetScheduleCycleDepthOk() (*int32, bool)`

GetScheduleCycleDepthOk returns a tuple with the ScheduleCycleDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleDepth

`func (o *V0043StatsMsg) SetScheduleCycleDepth(v int32)`

SetScheduleCycleDepth sets ScheduleCycleDepth field to given value.

### HasScheduleCycleDepth

`func (o *V0043StatsMsg) HasScheduleCycleDepth() bool`

HasScheduleCycleDepth returns a boolean if a field has been set.

### GetScheduleExit

`func (o *V0043StatsMsg) GetScheduleExit() V0043ScheduleExitFields`

GetScheduleExit returns the ScheduleExit field if non-nil, zero value otherwise.

### GetScheduleExitOk

`func (o *V0043StatsMsg) GetScheduleExitOk() (*V0043ScheduleExitFields, bool)`

GetScheduleExitOk returns a tuple with the ScheduleExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExit

`func (o *V0043StatsMsg) SetScheduleExit(v V0043ScheduleExitFields)`

SetScheduleExit sets ScheduleExit field to given value.

### HasScheduleExit

`func (o *V0043StatsMsg) HasScheduleExit() bool`

HasScheduleExit returns a boolean if a field has been set.

### GetScheduleQueueLength

`func (o *V0043StatsMsg) GetScheduleQueueLength() int32`

GetScheduleQueueLength returns the ScheduleQueueLength field if non-nil, zero value otherwise.

### GetScheduleQueueLengthOk

`func (o *V0043StatsMsg) GetScheduleQueueLengthOk() (*int32, bool)`

GetScheduleQueueLengthOk returns a tuple with the ScheduleQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleQueueLength

`func (o *V0043StatsMsg) SetScheduleQueueLength(v int32)`

SetScheduleQueueLength sets ScheduleQueueLength field to given value.

### HasScheduleQueueLength

`func (o *V0043StatsMsg) HasScheduleQueueLength() bool`

HasScheduleQueueLength returns a boolean if a field has been set.

### GetJobsSubmitted

`func (o *V0043StatsMsg) GetJobsSubmitted() int32`

GetJobsSubmitted returns the JobsSubmitted field if non-nil, zero value otherwise.

### GetJobsSubmittedOk

`func (o *V0043StatsMsg) GetJobsSubmittedOk() (*int32, bool)`

GetJobsSubmittedOk returns a tuple with the JobsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSubmitted

`func (o *V0043StatsMsg) SetJobsSubmitted(v int32)`

SetJobsSubmitted sets JobsSubmitted field to given value.

### HasJobsSubmitted

`func (o *V0043StatsMsg) HasJobsSubmitted() bool`

HasJobsSubmitted returns a boolean if a field has been set.

### GetJobsStarted

`func (o *V0043StatsMsg) GetJobsStarted() int32`

GetJobsStarted returns the JobsStarted field if non-nil, zero value otherwise.

### GetJobsStartedOk

`func (o *V0043StatsMsg) GetJobsStartedOk() (*int32, bool)`

GetJobsStartedOk returns a tuple with the JobsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsStarted

`func (o *V0043StatsMsg) SetJobsStarted(v int32)`

SetJobsStarted sets JobsStarted field to given value.

### HasJobsStarted

`func (o *V0043StatsMsg) HasJobsStarted() bool`

HasJobsStarted returns a boolean if a field has been set.

### GetJobsCompleted

`func (o *V0043StatsMsg) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *V0043StatsMsg) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *V0043StatsMsg) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.

### HasJobsCompleted

`func (o *V0043StatsMsg) HasJobsCompleted() bool`

HasJobsCompleted returns a boolean if a field has been set.

### GetJobsCanceled

`func (o *V0043StatsMsg) GetJobsCanceled() int32`

GetJobsCanceled returns the JobsCanceled field if non-nil, zero value otherwise.

### GetJobsCanceledOk

`func (o *V0043StatsMsg) GetJobsCanceledOk() (*int32, bool)`

GetJobsCanceledOk returns a tuple with the JobsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCanceled

`func (o *V0043StatsMsg) SetJobsCanceled(v int32)`

SetJobsCanceled sets JobsCanceled field to given value.

### HasJobsCanceled

`func (o *V0043StatsMsg) HasJobsCanceled() bool`

HasJobsCanceled returns a boolean if a field has been set.

### GetJobsFailed

`func (o *V0043StatsMsg) GetJobsFailed() int32`

GetJobsFailed returns the JobsFailed field if non-nil, zero value otherwise.

### GetJobsFailedOk

`func (o *V0043StatsMsg) GetJobsFailedOk() (*int32, bool)`

GetJobsFailedOk returns a tuple with the JobsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFailed

`func (o *V0043StatsMsg) SetJobsFailed(v int32)`

SetJobsFailed sets JobsFailed field to given value.

### HasJobsFailed

`func (o *V0043StatsMsg) HasJobsFailed() bool`

HasJobsFailed returns a boolean if a field has been set.

### GetJobsPending

`func (o *V0043StatsMsg) GetJobsPending() int32`

GetJobsPending returns the JobsPending field if non-nil, zero value otherwise.

### GetJobsPendingOk

`func (o *V0043StatsMsg) GetJobsPendingOk() (*int32, bool)`

GetJobsPendingOk returns a tuple with the JobsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsPending

`func (o *V0043StatsMsg) SetJobsPending(v int32)`

SetJobsPending sets JobsPending field to given value.

### HasJobsPending

`func (o *V0043StatsMsg) HasJobsPending() bool`

HasJobsPending returns a boolean if a field has been set.

### GetJobsRunning

`func (o *V0043StatsMsg) GetJobsRunning() int32`

GetJobsRunning returns the JobsRunning field if non-nil, zero value otherwise.

### GetJobsRunningOk

`func (o *V0043StatsMsg) GetJobsRunningOk() (*int32, bool)`

GetJobsRunningOk returns a tuple with the JobsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsRunning

`func (o *V0043StatsMsg) SetJobsRunning(v int32)`

SetJobsRunning sets JobsRunning field to given value.

### HasJobsRunning

`func (o *V0043StatsMsg) HasJobsRunning() bool`

HasJobsRunning returns a boolean if a field has been set.

### GetJobStatesTs

`func (o *V0043StatsMsg) GetJobStatesTs() V0043Uint64NoValStruct`

GetJobStatesTs returns the JobStatesTs field if non-nil, zero value otherwise.

### GetJobStatesTsOk

`func (o *V0043StatsMsg) GetJobStatesTsOk() (*V0043Uint64NoValStruct, bool)`

GetJobStatesTsOk returns a tuple with the JobStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatesTs

`func (o *V0043StatsMsg) SetJobStatesTs(v V0043Uint64NoValStruct)`

SetJobStatesTs sets JobStatesTs field to given value.

### HasJobStatesTs

`func (o *V0043StatsMsg) HasJobStatesTs() bool`

HasJobStatesTs returns a boolean if a field has been set.

### GetBfBackfilledJobs

`func (o *V0043StatsMsg) GetBfBackfilledJobs() int32`

GetBfBackfilledJobs returns the BfBackfilledJobs field if non-nil, zero value otherwise.

### GetBfBackfilledJobsOk

`func (o *V0043StatsMsg) GetBfBackfilledJobsOk() (*int32, bool)`

GetBfBackfilledJobsOk returns a tuple with the BfBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledJobs

`func (o *V0043StatsMsg) SetBfBackfilledJobs(v int32)`

SetBfBackfilledJobs sets BfBackfilledJobs field to given value.

### HasBfBackfilledJobs

`func (o *V0043StatsMsg) HasBfBackfilledJobs() bool`

HasBfBackfilledJobs returns a boolean if a field has been set.

### GetBfLastBackfilledJobs

`func (o *V0043StatsMsg) GetBfLastBackfilledJobs() int32`

GetBfLastBackfilledJobs returns the BfLastBackfilledJobs field if non-nil, zero value otherwise.

### GetBfLastBackfilledJobsOk

`func (o *V0043StatsMsg) GetBfLastBackfilledJobsOk() (*int32, bool)`

GetBfLastBackfilledJobsOk returns a tuple with the BfLastBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastBackfilledJobs

`func (o *V0043StatsMsg) SetBfLastBackfilledJobs(v int32)`

SetBfLastBackfilledJobs sets BfLastBackfilledJobs field to given value.

### HasBfLastBackfilledJobs

`func (o *V0043StatsMsg) HasBfLastBackfilledJobs() bool`

HasBfLastBackfilledJobs returns a boolean if a field has been set.

### GetBfBackfilledHetJobs

`func (o *V0043StatsMsg) GetBfBackfilledHetJobs() int32`

GetBfBackfilledHetJobs returns the BfBackfilledHetJobs field if non-nil, zero value otherwise.

### GetBfBackfilledHetJobsOk

`func (o *V0043StatsMsg) GetBfBackfilledHetJobsOk() (*int32, bool)`

GetBfBackfilledHetJobsOk returns a tuple with the BfBackfilledHetJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledHetJobs

`func (o *V0043StatsMsg) SetBfBackfilledHetJobs(v int32)`

SetBfBackfilledHetJobs sets BfBackfilledHetJobs field to given value.

### HasBfBackfilledHetJobs

`func (o *V0043StatsMsg) HasBfBackfilledHetJobs() bool`

HasBfBackfilledHetJobs returns a boolean if a field has been set.

### GetBfCycleCounter

`func (o *V0043StatsMsg) GetBfCycleCounter() int32`

GetBfCycleCounter returns the BfCycleCounter field if non-nil, zero value otherwise.

### GetBfCycleCounterOk

`func (o *V0043StatsMsg) GetBfCycleCounterOk() (*int32, bool)`

GetBfCycleCounterOk returns a tuple with the BfCycleCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleCounter

`func (o *V0043StatsMsg) SetBfCycleCounter(v int32)`

SetBfCycleCounter sets BfCycleCounter field to given value.

### HasBfCycleCounter

`func (o *V0043StatsMsg) HasBfCycleCounter() bool`

HasBfCycleCounter returns a boolean if a field has been set.

### GetBfCycleMean

`func (o *V0043StatsMsg) GetBfCycleMean() int64`

GetBfCycleMean returns the BfCycleMean field if non-nil, zero value otherwise.

### GetBfCycleMeanOk

`func (o *V0043StatsMsg) GetBfCycleMeanOk() (*int64, bool)`

GetBfCycleMeanOk returns a tuple with the BfCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMean

`func (o *V0043StatsMsg) SetBfCycleMean(v int64)`

SetBfCycleMean sets BfCycleMean field to given value.

### HasBfCycleMean

`func (o *V0043StatsMsg) HasBfCycleMean() bool`

HasBfCycleMean returns a boolean if a field has been set.

### GetBfDepthMean

`func (o *V0043StatsMsg) GetBfDepthMean() int64`

GetBfDepthMean returns the BfDepthMean field if non-nil, zero value otherwise.

### GetBfDepthMeanOk

`func (o *V0043StatsMsg) GetBfDepthMeanOk() (*int64, bool)`

GetBfDepthMeanOk returns a tuple with the BfDepthMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMean

`func (o *V0043StatsMsg) SetBfDepthMean(v int64)`

SetBfDepthMean sets BfDepthMean field to given value.

### HasBfDepthMean

`func (o *V0043StatsMsg) HasBfDepthMean() bool`

HasBfDepthMean returns a boolean if a field has been set.

### GetBfDepthMeanTry

`func (o *V0043StatsMsg) GetBfDepthMeanTry() int64`

GetBfDepthMeanTry returns the BfDepthMeanTry field if non-nil, zero value otherwise.

### GetBfDepthMeanTryOk

`func (o *V0043StatsMsg) GetBfDepthMeanTryOk() (*int64, bool)`

GetBfDepthMeanTryOk returns a tuple with the BfDepthMeanTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMeanTry

`func (o *V0043StatsMsg) SetBfDepthMeanTry(v int64)`

SetBfDepthMeanTry sets BfDepthMeanTry field to given value.

### HasBfDepthMeanTry

`func (o *V0043StatsMsg) HasBfDepthMeanTry() bool`

HasBfDepthMeanTry returns a boolean if a field has been set.

### GetBfCycleSum

`func (o *V0043StatsMsg) GetBfCycleSum() int64`

GetBfCycleSum returns the BfCycleSum field if non-nil, zero value otherwise.

### GetBfCycleSumOk

`func (o *V0043StatsMsg) GetBfCycleSumOk() (*int64, bool)`

GetBfCycleSumOk returns a tuple with the BfCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleSum

`func (o *V0043StatsMsg) SetBfCycleSum(v int64)`

SetBfCycleSum sets BfCycleSum field to given value.

### HasBfCycleSum

`func (o *V0043StatsMsg) HasBfCycleSum() bool`

HasBfCycleSum returns a boolean if a field has been set.

### GetBfCycleLast

`func (o *V0043StatsMsg) GetBfCycleLast() int32`

GetBfCycleLast returns the BfCycleLast field if non-nil, zero value otherwise.

### GetBfCycleLastOk

`func (o *V0043StatsMsg) GetBfCycleLastOk() (*int32, bool)`

GetBfCycleLastOk returns a tuple with the BfCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleLast

`func (o *V0043StatsMsg) SetBfCycleLast(v int32)`

SetBfCycleLast sets BfCycleLast field to given value.

### HasBfCycleLast

`func (o *V0043StatsMsg) HasBfCycleLast() bool`

HasBfCycleLast returns a boolean if a field has been set.

### GetBfCycleMax

`func (o *V0043StatsMsg) GetBfCycleMax() int32`

GetBfCycleMax returns the BfCycleMax field if non-nil, zero value otherwise.

### GetBfCycleMaxOk

`func (o *V0043StatsMsg) GetBfCycleMaxOk() (*int32, bool)`

GetBfCycleMaxOk returns a tuple with the BfCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMax

`func (o *V0043StatsMsg) SetBfCycleMax(v int32)`

SetBfCycleMax sets BfCycleMax field to given value.

### HasBfCycleMax

`func (o *V0043StatsMsg) HasBfCycleMax() bool`

HasBfCycleMax returns a boolean if a field has been set.

### GetBfExit

`func (o *V0043StatsMsg) GetBfExit() V0043BfExitFields`

GetBfExit returns the BfExit field if non-nil, zero value otherwise.

### GetBfExitOk

`func (o *V0043StatsMsg) GetBfExitOk() (*V0043BfExitFields, bool)`

GetBfExitOk returns a tuple with the BfExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfExit

`func (o *V0043StatsMsg) SetBfExit(v V0043BfExitFields)`

SetBfExit sets BfExit field to given value.

### HasBfExit

`func (o *V0043StatsMsg) HasBfExit() bool`

HasBfExit returns a boolean if a field has been set.

### GetBfLastDepth

`func (o *V0043StatsMsg) GetBfLastDepth() int32`

GetBfLastDepth returns the BfLastDepth field if non-nil, zero value otherwise.

### GetBfLastDepthOk

`func (o *V0043StatsMsg) GetBfLastDepthOk() (*int32, bool)`

GetBfLastDepthOk returns a tuple with the BfLastDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepth

`func (o *V0043StatsMsg) SetBfLastDepth(v int32)`

SetBfLastDepth sets BfLastDepth field to given value.

### HasBfLastDepth

`func (o *V0043StatsMsg) HasBfLastDepth() bool`

HasBfLastDepth returns a boolean if a field has been set.

### GetBfLastDepthTry

`func (o *V0043StatsMsg) GetBfLastDepthTry() int32`

GetBfLastDepthTry returns the BfLastDepthTry field if non-nil, zero value otherwise.

### GetBfLastDepthTryOk

`func (o *V0043StatsMsg) GetBfLastDepthTryOk() (*int32, bool)`

GetBfLastDepthTryOk returns a tuple with the BfLastDepthTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepthTry

`func (o *V0043StatsMsg) SetBfLastDepthTry(v int32)`

SetBfLastDepthTry sets BfLastDepthTry field to given value.

### HasBfLastDepthTry

`func (o *V0043StatsMsg) HasBfLastDepthTry() bool`

HasBfLastDepthTry returns a boolean if a field has been set.

### GetBfDepthSum

`func (o *V0043StatsMsg) GetBfDepthSum() int32`

GetBfDepthSum returns the BfDepthSum field if non-nil, zero value otherwise.

### GetBfDepthSumOk

`func (o *V0043StatsMsg) GetBfDepthSumOk() (*int32, bool)`

GetBfDepthSumOk returns a tuple with the BfDepthSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthSum

`func (o *V0043StatsMsg) SetBfDepthSum(v int32)`

SetBfDepthSum sets BfDepthSum field to given value.

### HasBfDepthSum

`func (o *V0043StatsMsg) HasBfDepthSum() bool`

HasBfDepthSum returns a boolean if a field has been set.

### GetBfDepthTrySum

`func (o *V0043StatsMsg) GetBfDepthTrySum() int32`

GetBfDepthTrySum returns the BfDepthTrySum field if non-nil, zero value otherwise.

### GetBfDepthTrySumOk

`func (o *V0043StatsMsg) GetBfDepthTrySumOk() (*int32, bool)`

GetBfDepthTrySumOk returns a tuple with the BfDepthTrySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthTrySum

`func (o *V0043StatsMsg) SetBfDepthTrySum(v int32)`

SetBfDepthTrySum sets BfDepthTrySum field to given value.

### HasBfDepthTrySum

`func (o *V0043StatsMsg) HasBfDepthTrySum() bool`

HasBfDepthTrySum returns a boolean if a field has been set.

### GetBfQueueLen

`func (o *V0043StatsMsg) GetBfQueueLen() int32`

GetBfQueueLen returns the BfQueueLen field if non-nil, zero value otherwise.

### GetBfQueueLenOk

`func (o *V0043StatsMsg) GetBfQueueLenOk() (*int32, bool)`

GetBfQueueLenOk returns a tuple with the BfQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLen

`func (o *V0043StatsMsg) SetBfQueueLen(v int32)`

SetBfQueueLen sets BfQueueLen field to given value.

### HasBfQueueLen

`func (o *V0043StatsMsg) HasBfQueueLen() bool`

HasBfQueueLen returns a boolean if a field has been set.

### GetBfQueueLenMean

`func (o *V0043StatsMsg) GetBfQueueLenMean() int64`

GetBfQueueLenMean returns the BfQueueLenMean field if non-nil, zero value otherwise.

### GetBfQueueLenMeanOk

`func (o *V0043StatsMsg) GetBfQueueLenMeanOk() (*int64, bool)`

GetBfQueueLenMeanOk returns a tuple with the BfQueueLenMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenMean

`func (o *V0043StatsMsg) SetBfQueueLenMean(v int64)`

SetBfQueueLenMean sets BfQueueLenMean field to given value.

### HasBfQueueLenMean

`func (o *V0043StatsMsg) HasBfQueueLenMean() bool`

HasBfQueueLenMean returns a boolean if a field has been set.

### GetBfQueueLenSum

`func (o *V0043StatsMsg) GetBfQueueLenSum() int32`

GetBfQueueLenSum returns the BfQueueLenSum field if non-nil, zero value otherwise.

### GetBfQueueLenSumOk

`func (o *V0043StatsMsg) GetBfQueueLenSumOk() (*int32, bool)`

GetBfQueueLenSumOk returns a tuple with the BfQueueLenSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenSum

`func (o *V0043StatsMsg) SetBfQueueLenSum(v int32)`

SetBfQueueLenSum sets BfQueueLenSum field to given value.

### HasBfQueueLenSum

`func (o *V0043StatsMsg) HasBfQueueLenSum() bool`

HasBfQueueLenSum returns a boolean if a field has been set.

### GetBfTableSize

`func (o *V0043StatsMsg) GetBfTableSize() int32`

GetBfTableSize returns the BfTableSize field if non-nil, zero value otherwise.

### GetBfTableSizeOk

`func (o *V0043StatsMsg) GetBfTableSizeOk() (*int32, bool)`

GetBfTableSizeOk returns a tuple with the BfTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSize

`func (o *V0043StatsMsg) SetBfTableSize(v int32)`

SetBfTableSize sets BfTableSize field to given value.

### HasBfTableSize

`func (o *V0043StatsMsg) HasBfTableSize() bool`

HasBfTableSize returns a boolean if a field has been set.

### GetBfTableSizeSum

`func (o *V0043StatsMsg) GetBfTableSizeSum() int32`

GetBfTableSizeSum returns the BfTableSizeSum field if non-nil, zero value otherwise.

### GetBfTableSizeSumOk

`func (o *V0043StatsMsg) GetBfTableSizeSumOk() (*int32, bool)`

GetBfTableSizeSumOk returns a tuple with the BfTableSizeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeSum

`func (o *V0043StatsMsg) SetBfTableSizeSum(v int32)`

SetBfTableSizeSum sets BfTableSizeSum field to given value.

### HasBfTableSizeSum

`func (o *V0043StatsMsg) HasBfTableSizeSum() bool`

HasBfTableSizeSum returns a boolean if a field has been set.

### GetBfTableSizeMean

`func (o *V0043StatsMsg) GetBfTableSizeMean() int64`

GetBfTableSizeMean returns the BfTableSizeMean field if non-nil, zero value otherwise.

### GetBfTableSizeMeanOk

`func (o *V0043StatsMsg) GetBfTableSizeMeanOk() (*int64, bool)`

GetBfTableSizeMeanOk returns a tuple with the BfTableSizeMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeMean

`func (o *V0043StatsMsg) SetBfTableSizeMean(v int64)`

SetBfTableSizeMean sets BfTableSizeMean field to given value.

### HasBfTableSizeMean

`func (o *V0043StatsMsg) HasBfTableSizeMean() bool`

HasBfTableSizeMean returns a boolean if a field has been set.

### GetBfWhenLastCycle

`func (o *V0043StatsMsg) GetBfWhenLastCycle() V0043Uint64NoValStruct`

GetBfWhenLastCycle returns the BfWhenLastCycle field if non-nil, zero value otherwise.

### GetBfWhenLastCycleOk

`func (o *V0043StatsMsg) GetBfWhenLastCycleOk() (*V0043Uint64NoValStruct, bool)`

GetBfWhenLastCycleOk returns a tuple with the BfWhenLastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfWhenLastCycle

`func (o *V0043StatsMsg) SetBfWhenLastCycle(v V0043Uint64NoValStruct)`

SetBfWhenLastCycle sets BfWhenLastCycle field to given value.

### HasBfWhenLastCycle

`func (o *V0043StatsMsg) HasBfWhenLastCycle() bool`

HasBfWhenLastCycle returns a boolean if a field has been set.

### GetBfActive

`func (o *V0043StatsMsg) GetBfActive() bool`

GetBfActive returns the BfActive field if non-nil, zero value otherwise.

### GetBfActiveOk

`func (o *V0043StatsMsg) GetBfActiveOk() (*bool, bool)`

GetBfActiveOk returns a tuple with the BfActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfActive

`func (o *V0043StatsMsg) SetBfActive(v bool)`

SetBfActive sets BfActive field to given value.

### HasBfActive

`func (o *V0043StatsMsg) HasBfActive() bool`

HasBfActive returns a boolean if a field has been set.

### GetRpcsByMessageType

`func (o *V0043StatsMsg) GetRpcsByMessageType() []V0043StatsMsgRpcType`

GetRpcsByMessageType returns the RpcsByMessageType field if non-nil, zero value otherwise.

### GetRpcsByMessageTypeOk

`func (o *V0043StatsMsg) GetRpcsByMessageTypeOk() (*[]V0043StatsMsgRpcType, bool)`

GetRpcsByMessageTypeOk returns a tuple with the RpcsByMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByMessageType

`func (o *V0043StatsMsg) SetRpcsByMessageType(v []V0043StatsMsgRpcType)`

SetRpcsByMessageType sets RpcsByMessageType field to given value.

### HasRpcsByMessageType

`func (o *V0043StatsMsg) HasRpcsByMessageType() bool`

HasRpcsByMessageType returns a boolean if a field has been set.

### GetRpcsByUser

`func (o *V0043StatsMsg) GetRpcsByUser() []V0043StatsMsgRpcUser`

GetRpcsByUser returns the RpcsByUser field if non-nil, zero value otherwise.

### GetRpcsByUserOk

`func (o *V0043StatsMsg) GetRpcsByUserOk() (*[]V0043StatsMsgRpcUser, bool)`

GetRpcsByUserOk returns a tuple with the RpcsByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByUser

`func (o *V0043StatsMsg) SetRpcsByUser(v []V0043StatsMsgRpcUser)`

SetRpcsByUser sets RpcsByUser field to given value.

### HasRpcsByUser

`func (o *V0043StatsMsg) HasRpcsByUser() bool`

HasRpcsByUser returns a boolean if a field has been set.

### GetPendingRpcs

`func (o *V0043StatsMsg) GetPendingRpcs() []V0043StatsMsgRpcQueue`

GetPendingRpcs returns the PendingRpcs field if non-nil, zero value otherwise.

### GetPendingRpcsOk

`func (o *V0043StatsMsg) GetPendingRpcsOk() (*[]V0043StatsMsgRpcQueue, bool)`

GetPendingRpcsOk returns a tuple with the PendingRpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcs

`func (o *V0043StatsMsg) SetPendingRpcs(v []V0043StatsMsgRpcQueue)`

SetPendingRpcs sets PendingRpcs field to given value.

### HasPendingRpcs

`func (o *V0043StatsMsg) HasPendingRpcs() bool`

HasPendingRpcs returns a boolean if a field has been set.

### GetPendingRpcsByHostlist

`func (o *V0043StatsMsg) GetPendingRpcsByHostlist() []V0043StatsMsgRpcDump`

GetPendingRpcsByHostlist returns the PendingRpcsByHostlist field if non-nil, zero value otherwise.

### GetPendingRpcsByHostlistOk

`func (o *V0043StatsMsg) GetPendingRpcsByHostlistOk() (*[]V0043StatsMsgRpcDump, bool)`

GetPendingRpcsByHostlistOk returns a tuple with the PendingRpcsByHostlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcsByHostlist

`func (o *V0043StatsMsg) SetPendingRpcsByHostlist(v []V0043StatsMsgRpcDump)`

SetPendingRpcsByHostlist sets PendingRpcsByHostlist field to given value.

### HasPendingRpcsByHostlist

`func (o *V0043StatsMsg) HasPendingRpcsByHostlist() bool`

HasPendingRpcsByHostlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


