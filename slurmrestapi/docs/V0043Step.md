# V0043Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**V0043StepTime**](V0043StepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0043ProcessExitCodeVerbose**](V0043ProcessExitCodeVerbose.md) |  | [optional] 
**Nodes** | Pointer to [**V0043StepNodes**](V0043StepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | Deprecated; Process ID | [optional] 
**CPU** | Pointer to [**V0043StepCPU**](V0043StepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User ID that requested termination of the step | [optional] 
**State** | Pointer to **[]string** | Current state | [optional] 
**Statistics** | Pointer to [**V0043StepStatistics**](V0043StepStatistics.md) |  | [optional] 
**Step** | Pointer to [**V0044StepStep**](V0044StepStep.md) |  | [optional] 
**Task** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask.md) |  | [optional] 
**Tres** | Pointer to [**V0043StepTres**](V0043StepTres.md) |  | [optional] 

## Methods

### NewV0043Step

`func NewV0043Step() *V0043Step`

NewV0043Step instantiates a new V0043Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043StepWithDefaults

`func NewV0043StepWithDefaults() *V0043Step`

NewV0043StepWithDefaults instantiates a new V0043Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V0043Step) GetTime() V0043StepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0043Step) GetTimeOk() (*V0043StepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0043Step) SetTime(v V0043StepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0043Step) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0043Step) GetExitCode() V0043ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0043Step) GetExitCodeOk() (*V0043ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0043Step) SetExitCode(v V0043ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0043Step) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *V0043Step) GetNodes() V0043StepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0043Step) GetNodesOk() (*V0043StepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0043Step) SetNodes(v V0043StepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0043Step) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *V0043Step) GetTasks() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0043Step) GetTasksOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0043Step) SetTasks(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0043Step) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *V0043Step) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *V0043Step) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *V0043Step) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *V0043Step) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *V0043Step) GetCPU() V0043StepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0043Step) GetCPUOk() (*V0043StepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0043Step) SetCPU(v V0043StepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0043Step) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0043Step) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0043Step) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0043Step) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0043Step) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *V0043Step) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0043Step) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0043Step) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0043Step) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *V0043Step) GetStatistics() V0043StepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0043Step) GetStatisticsOk() (*V0043StepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0043Step) SetStatistics(v V0043StepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0043Step) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *V0043Step) GetStep() V0044StepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V0043Step) GetStepOk() (*V0044StepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V0043Step) SetStep(v V0044StepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *V0043Step) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *V0043Step) GetTask() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0043Step) GetTaskOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0043Step) SetTask(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0043Step) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *V0043Step) GetTres() V0043StepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0043Step) GetTresOk() (*V0043StepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0043Step) SetTres(v V0043StepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0043Step) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


