# SlurmV0041GetDiag200ResponseStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartsPacked** | Pointer to **int32** | Zero if only RPC statistic included | [optional] 
**ReqTime** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsReqTime**](SlurmV0041GetDiag200ResponseStatisticsReqTime.md) |  | [optional] 
**ReqTimeStart** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsReqTimeStart**](SlurmV0041GetDiag200ResponseStatisticsReqTimeStart.md) |  | [optional] 
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
**ScheduleExit** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsScheduleExit**](SlurmV0041GetDiag200ResponseStatisticsScheduleExit.md) |  | [optional] 
**ScheduleQueueLength** | Pointer to **int32** | Number of jobs pending in queue | [optional] 
**JobsSubmitted** | Pointer to **int32** | Number of jobs submitted since last reset | [optional] 
**JobsStarted** | Pointer to **int32** | Number of jobs started since last reset | [optional] 
**JobsCompleted** | Pointer to **int32** | Number of jobs completed since last reset | [optional] 
**JobsCanceled** | Pointer to **int32** | Number of jobs canceled since the last reset | [optional] 
**JobsFailed** | Pointer to **int32** | Number of jobs failed due to slurmd or other internal issues since last reset | [optional] 
**JobsPending** | Pointer to **int32** | Number of jobs pending at the time of listed in job_state_ts | [optional] 
**JobsRunning** | Pointer to **int32** | Number of jobs running at the time of listed in job_state_ts | [optional] 
**JobStatesTs** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsJobStatesTs**](SlurmV0041GetDiag200ResponseStatisticsJobStatesTs.md) |  | [optional] 
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
**BfExit** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsBfExit**](SlurmV0041GetDiag200ResponseStatisticsBfExit.md) |  | [optional] 
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
**BfWhenLastCycle** | Pointer to [**SlurmV0041GetDiag200ResponseStatisticsBfWhenLastCycle**](SlurmV0041GetDiag200ResponseStatisticsBfWhenLastCycle.md) |  | [optional] 
**BfActive** | Pointer to **bool** | Backfill scheduler currently running | [optional] 
**RpcsByMessageType** | Pointer to [**[]SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner**](SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner.md) | Most frequently issued remote procedure calls (RPCs) | [optional] 
**RpcsByUser** | Pointer to [**[]SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner**](SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner.md) | RPCs issued by user ID | [optional] 
**PendingRpcs** | Pointer to [**[]SlurmV0041GetDiag200ResponseStatisticsPendingRpcsInner**](SlurmV0041GetDiag200ResponseStatisticsPendingRpcsInner.md) | Pending RPC statistics | [optional] 
**PendingRpcsByHostlist** | Pointer to [**[]SlurmV0041GetDiag200ResponseStatisticsPendingRpcsByHostlistInner**](SlurmV0041GetDiag200ResponseStatisticsPendingRpcsByHostlistInner.md) | Pending RPCs hostlists | [optional] 

## Methods

### NewSlurmV0041GetDiag200ResponseStatistics

`func NewSlurmV0041GetDiag200ResponseStatistics() *SlurmV0041GetDiag200ResponseStatistics`

NewSlurmV0041GetDiag200ResponseStatistics instantiates a new SlurmV0041GetDiag200ResponseStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetDiag200ResponseStatisticsWithDefaults

`func NewSlurmV0041GetDiag200ResponseStatisticsWithDefaults() *SlurmV0041GetDiag200ResponseStatistics`

NewSlurmV0041GetDiag200ResponseStatisticsWithDefaults instantiates a new SlurmV0041GetDiag200ResponseStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartsPacked

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPartsPacked() int32`

GetPartsPacked returns the PartsPacked field if non-nil, zero value otherwise.

### GetPartsPackedOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPartsPackedOk() (*int32, bool)`

GetPartsPackedOk returns a tuple with the PartsPacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPacked

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetPartsPacked(v int32)`

SetPartsPacked sets PartsPacked field to given value.

### HasPartsPacked

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasPartsPacked() bool`

HasPartsPacked returns a boolean if a field has been set.

### GetReqTime

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetReqTime() SlurmV0041GetDiag200ResponseStatisticsReqTime`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetReqTimeOk() (*SlurmV0041GetDiag200ResponseStatisticsReqTime, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetReqTime(v SlurmV0041GetDiag200ResponseStatisticsReqTime)`

SetReqTime sets ReqTime field to given value.

### HasReqTime

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasReqTime() bool`

HasReqTime returns a boolean if a field has been set.

### GetReqTimeStart

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetReqTimeStart() SlurmV0041GetDiag200ResponseStatisticsReqTimeStart`

GetReqTimeStart returns the ReqTimeStart field if non-nil, zero value otherwise.

### GetReqTimeStartOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetReqTimeStartOk() (*SlurmV0041GetDiag200ResponseStatisticsReqTimeStart, bool)`

GetReqTimeStartOk returns a tuple with the ReqTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeStart

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetReqTimeStart(v SlurmV0041GetDiag200ResponseStatisticsReqTimeStart)`

SetReqTimeStart sets ReqTimeStart field to given value.

### HasReqTimeStart

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasReqTimeStart() bool`

HasReqTimeStart returns a boolean if a field has been set.

### GetServerThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetServerThreadCount() int32`

GetServerThreadCount returns the ServerThreadCount field if non-nil, zero value otherwise.

### GetServerThreadCountOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetServerThreadCountOk() (*int32, bool)`

GetServerThreadCountOk returns a tuple with the ServerThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetServerThreadCount(v int32)`

SetServerThreadCount sets ServerThreadCount field to given value.

### HasServerThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasServerThreadCount() bool`

HasServerThreadCount returns a boolean if a field has been set.

### GetAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentQueueSize() int32`

GetAgentQueueSize returns the AgentQueueSize field if non-nil, zero value otherwise.

### GetAgentQueueSizeOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentQueueSizeOk() (*int32, bool)`

GetAgentQueueSizeOk returns a tuple with the AgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetAgentQueueSize(v int32)`

SetAgentQueueSize sets AgentQueueSize field to given value.

### HasAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasAgentQueueSize() bool`

HasAgentQueueSize returns a boolean if a field has been set.

### GetAgentCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentCount() int32`

GetAgentCount returns the AgentCount field if non-nil, zero value otherwise.

### GetAgentCountOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentCountOk() (*int32, bool)`

GetAgentCountOk returns a tuple with the AgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetAgentCount(v int32)`

SetAgentCount sets AgentCount field to given value.

### HasAgentCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasAgentCount() bool`

HasAgentCount returns a boolean if a field has been set.

### GetAgentThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentThreadCount() int32`

GetAgentThreadCount returns the AgentThreadCount field if non-nil, zero value otherwise.

### GetAgentThreadCountOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetAgentThreadCountOk() (*int32, bool)`

GetAgentThreadCountOk returns a tuple with the AgentThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetAgentThreadCount(v int32)`

SetAgentThreadCount sets AgentThreadCount field to given value.

### HasAgentThreadCount

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasAgentThreadCount() bool`

HasAgentThreadCount returns a boolean if a field has been set.

### GetDbdAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetDbdAgentQueueSize() int32`

GetDbdAgentQueueSize returns the DbdAgentQueueSize field if non-nil, zero value otherwise.

### GetDbdAgentQueueSizeOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetDbdAgentQueueSizeOk() (*int32, bool)`

GetDbdAgentQueueSizeOk returns a tuple with the DbdAgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbdAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetDbdAgentQueueSize(v int32)`

SetDbdAgentQueueSize sets DbdAgentQueueSize field to given value.

### HasDbdAgentQueueSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasDbdAgentQueueSize() bool`

HasDbdAgentQueueSize returns a boolean if a field has been set.

### GetGettimeofdayLatency

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetGettimeofdayLatency() int32`

GetGettimeofdayLatency returns the GettimeofdayLatency field if non-nil, zero value otherwise.

### GetGettimeofdayLatencyOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetGettimeofdayLatencyOk() (*int32, bool)`

GetGettimeofdayLatencyOk returns a tuple with the GettimeofdayLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettimeofdayLatency

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetGettimeofdayLatency(v int32)`

SetGettimeofdayLatency sets GettimeofdayLatency field to given value.

### HasGettimeofdayLatency

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasGettimeofdayLatency() bool`

HasGettimeofdayLatency returns a boolean if a field has been set.

### GetScheduleCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMax() int32`

GetScheduleCycleMax returns the ScheduleCycleMax field if non-nil, zero value otherwise.

### GetScheduleCycleMaxOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMaxOk() (*int32, bool)`

GetScheduleCycleMaxOk returns a tuple with the ScheduleCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleMax(v int32)`

SetScheduleCycleMax sets ScheduleCycleMax field to given value.

### HasScheduleCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleMax() bool`

HasScheduleCycleMax returns a boolean if a field has been set.

### GetScheduleCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleLast() int32`

GetScheduleCycleLast returns the ScheduleCycleLast field if non-nil, zero value otherwise.

### GetScheduleCycleLastOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleLastOk() (*int32, bool)`

GetScheduleCycleLastOk returns a tuple with the ScheduleCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleLast(v int32)`

SetScheduleCycleLast sets ScheduleCycleLast field to given value.

### HasScheduleCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleLast() bool`

HasScheduleCycleLast returns a boolean if a field has been set.

### GetScheduleCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleSum() int64`

GetScheduleCycleSum returns the ScheduleCycleSum field if non-nil, zero value otherwise.

### GetScheduleCycleSumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleSumOk() (*int64, bool)`

GetScheduleCycleSumOk returns a tuple with the ScheduleCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleSum(v int64)`

SetScheduleCycleSum sets ScheduleCycleSum field to given value.

### HasScheduleCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleSum() bool`

HasScheduleCycleSum returns a boolean if a field has been set.

### GetScheduleCycleTotal

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleTotal() int32`

GetScheduleCycleTotal returns the ScheduleCycleTotal field if non-nil, zero value otherwise.

### GetScheduleCycleTotalOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleTotalOk() (*int32, bool)`

GetScheduleCycleTotalOk returns a tuple with the ScheduleCycleTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleTotal

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleTotal(v int32)`

SetScheduleCycleTotal sets ScheduleCycleTotal field to given value.

### HasScheduleCycleTotal

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleTotal() bool`

HasScheduleCycleTotal returns a boolean if a field has been set.

### GetScheduleCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMean() int64`

GetScheduleCycleMean returns the ScheduleCycleMean field if non-nil, zero value otherwise.

### GetScheduleCycleMeanOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMeanOk() (*int64, bool)`

GetScheduleCycleMeanOk returns a tuple with the ScheduleCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleMean(v int64)`

SetScheduleCycleMean sets ScheduleCycleMean field to given value.

### HasScheduleCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleMean() bool`

HasScheduleCycleMean returns a boolean if a field has been set.

### GetScheduleCycleMeanDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMeanDepth() int64`

GetScheduleCycleMeanDepth returns the ScheduleCycleMeanDepth field if non-nil, zero value otherwise.

### GetScheduleCycleMeanDepthOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleMeanDepthOk() (*int64, bool)`

GetScheduleCycleMeanDepthOk returns a tuple with the ScheduleCycleMeanDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMeanDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleMeanDepth(v int64)`

SetScheduleCycleMeanDepth sets ScheduleCycleMeanDepth field to given value.

### HasScheduleCycleMeanDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleMeanDepth() bool`

HasScheduleCycleMeanDepth returns a boolean if a field has been set.

### GetScheduleCyclePerMinute

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCyclePerMinute() int64`

GetScheduleCyclePerMinute returns the ScheduleCyclePerMinute field if non-nil, zero value otherwise.

### GetScheduleCyclePerMinuteOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCyclePerMinuteOk() (*int64, bool)`

GetScheduleCyclePerMinuteOk returns a tuple with the ScheduleCyclePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCyclePerMinute

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCyclePerMinute(v int64)`

SetScheduleCyclePerMinute sets ScheduleCyclePerMinute field to given value.

### HasScheduleCyclePerMinute

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCyclePerMinute() bool`

HasScheduleCyclePerMinute returns a boolean if a field has been set.

### GetScheduleCycleDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleDepth() int32`

GetScheduleCycleDepth returns the ScheduleCycleDepth field if non-nil, zero value otherwise.

### GetScheduleCycleDepthOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleCycleDepthOk() (*int32, bool)`

GetScheduleCycleDepthOk returns a tuple with the ScheduleCycleDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleCycleDepth(v int32)`

SetScheduleCycleDepth sets ScheduleCycleDepth field to given value.

### HasScheduleCycleDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleCycleDepth() bool`

HasScheduleCycleDepth returns a boolean if a field has been set.

### GetScheduleExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleExit() SlurmV0041GetDiag200ResponseStatisticsScheduleExit`

GetScheduleExit returns the ScheduleExit field if non-nil, zero value otherwise.

### GetScheduleExitOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleExitOk() (*SlurmV0041GetDiag200ResponseStatisticsScheduleExit, bool)`

GetScheduleExitOk returns a tuple with the ScheduleExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleExit(v SlurmV0041GetDiag200ResponseStatisticsScheduleExit)`

SetScheduleExit sets ScheduleExit field to given value.

### HasScheduleExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleExit() bool`

HasScheduleExit returns a boolean if a field has been set.

### GetScheduleQueueLength

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleQueueLength() int32`

GetScheduleQueueLength returns the ScheduleQueueLength field if non-nil, zero value otherwise.

### GetScheduleQueueLengthOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetScheduleQueueLengthOk() (*int32, bool)`

GetScheduleQueueLengthOk returns a tuple with the ScheduleQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleQueueLength

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetScheduleQueueLength(v int32)`

SetScheduleQueueLength sets ScheduleQueueLength field to given value.

### HasScheduleQueueLength

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasScheduleQueueLength() bool`

HasScheduleQueueLength returns a boolean if a field has been set.

### GetJobsSubmitted

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsSubmitted() int32`

GetJobsSubmitted returns the JobsSubmitted field if non-nil, zero value otherwise.

### GetJobsSubmittedOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsSubmittedOk() (*int32, bool)`

GetJobsSubmittedOk returns a tuple with the JobsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSubmitted

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsSubmitted(v int32)`

SetJobsSubmitted sets JobsSubmitted field to given value.

### HasJobsSubmitted

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsSubmitted() bool`

HasJobsSubmitted returns a boolean if a field has been set.

### GetJobsStarted

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsStarted() int32`

GetJobsStarted returns the JobsStarted field if non-nil, zero value otherwise.

### GetJobsStartedOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsStartedOk() (*int32, bool)`

GetJobsStartedOk returns a tuple with the JobsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsStarted

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsStarted(v int32)`

SetJobsStarted sets JobsStarted field to given value.

### HasJobsStarted

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsStarted() bool`

HasJobsStarted returns a boolean if a field has been set.

### GetJobsCompleted

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.

### HasJobsCompleted

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsCompleted() bool`

HasJobsCompleted returns a boolean if a field has been set.

### GetJobsCanceled

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsCanceled() int32`

GetJobsCanceled returns the JobsCanceled field if non-nil, zero value otherwise.

### GetJobsCanceledOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsCanceledOk() (*int32, bool)`

GetJobsCanceledOk returns a tuple with the JobsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCanceled

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsCanceled(v int32)`

SetJobsCanceled sets JobsCanceled field to given value.

### HasJobsCanceled

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsCanceled() bool`

HasJobsCanceled returns a boolean if a field has been set.

### GetJobsFailed

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsFailed() int32`

GetJobsFailed returns the JobsFailed field if non-nil, zero value otherwise.

### GetJobsFailedOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsFailedOk() (*int32, bool)`

GetJobsFailedOk returns a tuple with the JobsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFailed

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsFailed(v int32)`

SetJobsFailed sets JobsFailed field to given value.

### HasJobsFailed

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsFailed() bool`

HasJobsFailed returns a boolean if a field has been set.

### GetJobsPending

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsPending() int32`

GetJobsPending returns the JobsPending field if non-nil, zero value otherwise.

### GetJobsPendingOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsPendingOk() (*int32, bool)`

GetJobsPendingOk returns a tuple with the JobsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsPending

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsPending(v int32)`

SetJobsPending sets JobsPending field to given value.

### HasJobsPending

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsPending() bool`

HasJobsPending returns a boolean if a field has been set.

### GetJobsRunning

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsRunning() int32`

GetJobsRunning returns the JobsRunning field if non-nil, zero value otherwise.

### GetJobsRunningOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobsRunningOk() (*int32, bool)`

GetJobsRunningOk returns a tuple with the JobsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsRunning

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobsRunning(v int32)`

SetJobsRunning sets JobsRunning field to given value.

### HasJobsRunning

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobsRunning() bool`

HasJobsRunning returns a boolean if a field has been set.

### GetJobStatesTs

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobStatesTs() SlurmV0041GetDiag200ResponseStatisticsJobStatesTs`

GetJobStatesTs returns the JobStatesTs field if non-nil, zero value otherwise.

### GetJobStatesTsOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetJobStatesTsOk() (*SlurmV0041GetDiag200ResponseStatisticsJobStatesTs, bool)`

GetJobStatesTsOk returns a tuple with the JobStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatesTs

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetJobStatesTs(v SlurmV0041GetDiag200ResponseStatisticsJobStatesTs)`

SetJobStatesTs sets JobStatesTs field to given value.

### HasJobStatesTs

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasJobStatesTs() bool`

HasJobStatesTs returns a boolean if a field has been set.

### GetBfBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfBackfilledJobs() int32`

GetBfBackfilledJobs returns the BfBackfilledJobs field if non-nil, zero value otherwise.

### GetBfBackfilledJobsOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfBackfilledJobsOk() (*int32, bool)`

GetBfBackfilledJobsOk returns a tuple with the BfBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfBackfilledJobs(v int32)`

SetBfBackfilledJobs sets BfBackfilledJobs field to given value.

### HasBfBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfBackfilledJobs() bool`

HasBfBackfilledJobs returns a boolean if a field has been set.

### GetBfLastBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastBackfilledJobs() int32`

GetBfLastBackfilledJobs returns the BfLastBackfilledJobs field if non-nil, zero value otherwise.

### GetBfLastBackfilledJobsOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastBackfilledJobsOk() (*int32, bool)`

GetBfLastBackfilledJobsOk returns a tuple with the BfLastBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfLastBackfilledJobs(v int32)`

SetBfLastBackfilledJobs sets BfLastBackfilledJobs field to given value.

### HasBfLastBackfilledJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfLastBackfilledJobs() bool`

HasBfLastBackfilledJobs returns a boolean if a field has been set.

### GetBfBackfilledHetJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfBackfilledHetJobs() int32`

GetBfBackfilledHetJobs returns the BfBackfilledHetJobs field if non-nil, zero value otherwise.

### GetBfBackfilledHetJobsOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfBackfilledHetJobsOk() (*int32, bool)`

GetBfBackfilledHetJobsOk returns a tuple with the BfBackfilledHetJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledHetJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfBackfilledHetJobs(v int32)`

SetBfBackfilledHetJobs sets BfBackfilledHetJobs field to given value.

### HasBfBackfilledHetJobs

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfBackfilledHetJobs() bool`

HasBfBackfilledHetJobs returns a boolean if a field has been set.

### GetBfCycleCounter

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleCounter() int32`

GetBfCycleCounter returns the BfCycleCounter field if non-nil, zero value otherwise.

### GetBfCycleCounterOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleCounterOk() (*int32, bool)`

GetBfCycleCounterOk returns a tuple with the BfCycleCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleCounter

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfCycleCounter(v int32)`

SetBfCycleCounter sets BfCycleCounter field to given value.

### HasBfCycleCounter

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfCycleCounter() bool`

HasBfCycleCounter returns a boolean if a field has been set.

### GetBfCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleMean() int64`

GetBfCycleMean returns the BfCycleMean field if non-nil, zero value otherwise.

### GetBfCycleMeanOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleMeanOk() (*int64, bool)`

GetBfCycleMeanOk returns a tuple with the BfCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfCycleMean(v int64)`

SetBfCycleMean sets BfCycleMean field to given value.

### HasBfCycleMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfCycleMean() bool`

HasBfCycleMean returns a boolean if a field has been set.

### GetBfDepthMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthMean() int64`

GetBfDepthMean returns the BfDepthMean field if non-nil, zero value otherwise.

### GetBfDepthMeanOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthMeanOk() (*int64, bool)`

GetBfDepthMeanOk returns a tuple with the BfDepthMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfDepthMean(v int64)`

SetBfDepthMean sets BfDepthMean field to given value.

### HasBfDepthMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfDepthMean() bool`

HasBfDepthMean returns a boolean if a field has been set.

### GetBfDepthMeanTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthMeanTry() int64`

GetBfDepthMeanTry returns the BfDepthMeanTry field if non-nil, zero value otherwise.

### GetBfDepthMeanTryOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthMeanTryOk() (*int64, bool)`

GetBfDepthMeanTryOk returns a tuple with the BfDepthMeanTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMeanTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfDepthMeanTry(v int64)`

SetBfDepthMeanTry sets BfDepthMeanTry field to given value.

### HasBfDepthMeanTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfDepthMeanTry() bool`

HasBfDepthMeanTry returns a boolean if a field has been set.

### GetBfCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleSum() int64`

GetBfCycleSum returns the BfCycleSum field if non-nil, zero value otherwise.

### GetBfCycleSumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleSumOk() (*int64, bool)`

GetBfCycleSumOk returns a tuple with the BfCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfCycleSum(v int64)`

SetBfCycleSum sets BfCycleSum field to given value.

### HasBfCycleSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfCycleSum() bool`

HasBfCycleSum returns a boolean if a field has been set.

### GetBfCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleLast() int32`

GetBfCycleLast returns the BfCycleLast field if non-nil, zero value otherwise.

### GetBfCycleLastOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleLastOk() (*int32, bool)`

GetBfCycleLastOk returns a tuple with the BfCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfCycleLast(v int32)`

SetBfCycleLast sets BfCycleLast field to given value.

### HasBfCycleLast

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfCycleLast() bool`

HasBfCycleLast returns a boolean if a field has been set.

### GetBfCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleMax() int32`

GetBfCycleMax returns the BfCycleMax field if non-nil, zero value otherwise.

### GetBfCycleMaxOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfCycleMaxOk() (*int32, bool)`

GetBfCycleMaxOk returns a tuple with the BfCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfCycleMax(v int32)`

SetBfCycleMax sets BfCycleMax field to given value.

### HasBfCycleMax

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfCycleMax() bool`

HasBfCycleMax returns a boolean if a field has been set.

### GetBfExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfExit() SlurmV0041GetDiag200ResponseStatisticsBfExit`

GetBfExit returns the BfExit field if non-nil, zero value otherwise.

### GetBfExitOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfExitOk() (*SlurmV0041GetDiag200ResponseStatisticsBfExit, bool)`

GetBfExitOk returns a tuple with the BfExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfExit(v SlurmV0041GetDiag200ResponseStatisticsBfExit)`

SetBfExit sets BfExit field to given value.

### HasBfExit

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfExit() bool`

HasBfExit returns a boolean if a field has been set.

### GetBfLastDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastDepth() int32`

GetBfLastDepth returns the BfLastDepth field if non-nil, zero value otherwise.

### GetBfLastDepthOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastDepthOk() (*int32, bool)`

GetBfLastDepthOk returns a tuple with the BfLastDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfLastDepth(v int32)`

SetBfLastDepth sets BfLastDepth field to given value.

### HasBfLastDepth

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfLastDepth() bool`

HasBfLastDepth returns a boolean if a field has been set.

### GetBfLastDepthTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastDepthTry() int32`

GetBfLastDepthTry returns the BfLastDepthTry field if non-nil, zero value otherwise.

### GetBfLastDepthTryOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfLastDepthTryOk() (*int32, bool)`

GetBfLastDepthTryOk returns a tuple with the BfLastDepthTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepthTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfLastDepthTry(v int32)`

SetBfLastDepthTry sets BfLastDepthTry field to given value.

### HasBfLastDepthTry

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfLastDepthTry() bool`

HasBfLastDepthTry returns a boolean if a field has been set.

### GetBfDepthSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthSum() int32`

GetBfDepthSum returns the BfDepthSum field if non-nil, zero value otherwise.

### GetBfDepthSumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthSumOk() (*int32, bool)`

GetBfDepthSumOk returns a tuple with the BfDepthSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfDepthSum(v int32)`

SetBfDepthSum sets BfDepthSum field to given value.

### HasBfDepthSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfDepthSum() bool`

HasBfDepthSum returns a boolean if a field has been set.

### GetBfDepthTrySum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthTrySum() int32`

GetBfDepthTrySum returns the BfDepthTrySum field if non-nil, zero value otherwise.

### GetBfDepthTrySumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfDepthTrySumOk() (*int32, bool)`

GetBfDepthTrySumOk returns a tuple with the BfDepthTrySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthTrySum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfDepthTrySum(v int32)`

SetBfDepthTrySum sets BfDepthTrySum field to given value.

### HasBfDepthTrySum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfDepthTrySum() bool`

HasBfDepthTrySum returns a boolean if a field has been set.

### GetBfQueueLen

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLen() int32`

GetBfQueueLen returns the BfQueueLen field if non-nil, zero value otherwise.

### GetBfQueueLenOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLenOk() (*int32, bool)`

GetBfQueueLenOk returns a tuple with the BfQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLen

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfQueueLen(v int32)`

SetBfQueueLen sets BfQueueLen field to given value.

### HasBfQueueLen

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfQueueLen() bool`

HasBfQueueLen returns a boolean if a field has been set.

### GetBfQueueLenMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLenMean() int64`

GetBfQueueLenMean returns the BfQueueLenMean field if non-nil, zero value otherwise.

### GetBfQueueLenMeanOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLenMeanOk() (*int64, bool)`

GetBfQueueLenMeanOk returns a tuple with the BfQueueLenMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfQueueLenMean(v int64)`

SetBfQueueLenMean sets BfQueueLenMean field to given value.

### HasBfQueueLenMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfQueueLenMean() bool`

HasBfQueueLenMean returns a boolean if a field has been set.

### GetBfQueueLenSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLenSum() int32`

GetBfQueueLenSum returns the BfQueueLenSum field if non-nil, zero value otherwise.

### GetBfQueueLenSumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfQueueLenSumOk() (*int32, bool)`

GetBfQueueLenSumOk returns a tuple with the BfQueueLenSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfQueueLenSum(v int32)`

SetBfQueueLenSum sets BfQueueLenSum field to given value.

### HasBfQueueLenSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfQueueLenSum() bool`

HasBfQueueLenSum returns a boolean if a field has been set.

### GetBfTableSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSize() int32`

GetBfTableSize returns the BfTableSize field if non-nil, zero value otherwise.

### GetBfTableSizeOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSizeOk() (*int32, bool)`

GetBfTableSizeOk returns a tuple with the BfTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfTableSize(v int32)`

SetBfTableSize sets BfTableSize field to given value.

### HasBfTableSize

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfTableSize() bool`

HasBfTableSize returns a boolean if a field has been set.

### GetBfTableSizeSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSizeSum() int32`

GetBfTableSizeSum returns the BfTableSizeSum field if non-nil, zero value otherwise.

### GetBfTableSizeSumOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSizeSumOk() (*int32, bool)`

GetBfTableSizeSumOk returns a tuple with the BfTableSizeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfTableSizeSum(v int32)`

SetBfTableSizeSum sets BfTableSizeSum field to given value.

### HasBfTableSizeSum

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfTableSizeSum() bool`

HasBfTableSizeSum returns a boolean if a field has been set.

### GetBfTableSizeMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSizeMean() int64`

GetBfTableSizeMean returns the BfTableSizeMean field if non-nil, zero value otherwise.

### GetBfTableSizeMeanOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfTableSizeMeanOk() (*int64, bool)`

GetBfTableSizeMeanOk returns a tuple with the BfTableSizeMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfTableSizeMean(v int64)`

SetBfTableSizeMean sets BfTableSizeMean field to given value.

### HasBfTableSizeMean

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfTableSizeMean() bool`

HasBfTableSizeMean returns a boolean if a field has been set.

### GetBfWhenLastCycle

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfWhenLastCycle() SlurmV0041GetDiag200ResponseStatisticsBfWhenLastCycle`

GetBfWhenLastCycle returns the BfWhenLastCycle field if non-nil, zero value otherwise.

### GetBfWhenLastCycleOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfWhenLastCycleOk() (*SlurmV0041GetDiag200ResponseStatisticsBfWhenLastCycle, bool)`

GetBfWhenLastCycleOk returns a tuple with the BfWhenLastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfWhenLastCycle

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfWhenLastCycle(v SlurmV0041GetDiag200ResponseStatisticsBfWhenLastCycle)`

SetBfWhenLastCycle sets BfWhenLastCycle field to given value.

### HasBfWhenLastCycle

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfWhenLastCycle() bool`

HasBfWhenLastCycle returns a boolean if a field has been set.

### GetBfActive

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfActive() bool`

GetBfActive returns the BfActive field if non-nil, zero value otherwise.

### GetBfActiveOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetBfActiveOk() (*bool, bool)`

GetBfActiveOk returns a tuple with the BfActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfActive

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetBfActive(v bool)`

SetBfActive sets BfActive field to given value.

### HasBfActive

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasBfActive() bool`

HasBfActive returns a boolean if a field has been set.

### GetRpcsByMessageType

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetRpcsByMessageType() []SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner`

GetRpcsByMessageType returns the RpcsByMessageType field if non-nil, zero value otherwise.

### GetRpcsByMessageTypeOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetRpcsByMessageTypeOk() (*[]SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner, bool)`

GetRpcsByMessageTypeOk returns a tuple with the RpcsByMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByMessageType

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetRpcsByMessageType(v []SlurmV0041GetDiag200ResponseStatisticsRpcsByMessageTypeInner)`

SetRpcsByMessageType sets RpcsByMessageType field to given value.

### HasRpcsByMessageType

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasRpcsByMessageType() bool`

HasRpcsByMessageType returns a boolean if a field has been set.

### GetRpcsByUser

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetRpcsByUser() []SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner`

GetRpcsByUser returns the RpcsByUser field if non-nil, zero value otherwise.

### GetRpcsByUserOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetRpcsByUserOk() (*[]SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner, bool)`

GetRpcsByUserOk returns a tuple with the RpcsByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByUser

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetRpcsByUser(v []SlurmV0041GetDiag200ResponseStatisticsRpcsByUserInner)`

SetRpcsByUser sets RpcsByUser field to given value.

### HasRpcsByUser

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasRpcsByUser() bool`

HasRpcsByUser returns a boolean if a field has been set.

### GetPendingRpcs

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPendingRpcs() []SlurmV0041GetDiag200ResponseStatisticsPendingRpcsInner`

GetPendingRpcs returns the PendingRpcs field if non-nil, zero value otherwise.

### GetPendingRpcsOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPendingRpcsOk() (*[]SlurmV0041GetDiag200ResponseStatisticsPendingRpcsInner, bool)`

GetPendingRpcsOk returns a tuple with the PendingRpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcs

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetPendingRpcs(v []SlurmV0041GetDiag200ResponseStatisticsPendingRpcsInner)`

SetPendingRpcs sets PendingRpcs field to given value.

### HasPendingRpcs

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasPendingRpcs() bool`

HasPendingRpcs returns a boolean if a field has been set.

### GetPendingRpcsByHostlist

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPendingRpcsByHostlist() []SlurmV0041GetDiag200ResponseStatisticsPendingRpcsByHostlistInner`

GetPendingRpcsByHostlist returns the PendingRpcsByHostlist field if non-nil, zero value otherwise.

### GetPendingRpcsByHostlistOk

`func (o *SlurmV0041GetDiag200ResponseStatistics) GetPendingRpcsByHostlistOk() (*[]SlurmV0041GetDiag200ResponseStatisticsPendingRpcsByHostlistInner, bool)`

GetPendingRpcsByHostlistOk returns a tuple with the PendingRpcsByHostlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcsByHostlist

`func (o *SlurmV0041GetDiag200ResponseStatistics) SetPendingRpcsByHostlist(v []SlurmV0041GetDiag200ResponseStatisticsPendingRpcsByHostlistInner)`

SetPendingRpcsByHostlist sets PendingRpcsByHostlist field to given value.

### HasPendingRpcsByHostlist

`func (o *SlurmV0041GetDiag200ResponseStatistics) HasPendingRpcsByHostlist() bool`

HasPendingRpcsByHostlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


