# SlurmV0041PostJobSubmitRequestJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account associated with the job | [optional] 
**AccountGatherFrequency** | Pointer to **string** | Job accounting and profiling sampling intervals in seconds | [optional] 
**AdminComment** | Pointer to **string** | Arbitrary comment made by administrator | [optional] 
**AllocationNodeList** | Pointer to **string** | Local node making the resource allocation | [optional] 
**AllocationNodePort** | Pointer to **int32** | Port to send allocation confirmation to | [optional] 
**Argv** | Pointer to **[]string** | Arguments to the script | [optional] 
**Array** | Pointer to **string** | Job array index value specification | [optional] 
**BatchFeatures** | Pointer to **string** | Features required for batch script&#39;s node | [optional] 
**BeginTime** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerBeginTime**](SlurmV0041PostJobSubmitRequestJobsInnerBeginTime.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Job flags | [optional] 
**BurstBuffer** | Pointer to **string** | Burst buffer specifications | [optional] 
**Clusters** | Pointer to **string** | Clusters that a federated job can run on | [optional] 
**ClusterConstraint** | Pointer to **string** | Required features that a federated cluster must have to have a sibling job submitted to it | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment made by user | [optional] 
**Contiguous** | Pointer to **bool** | True if job requires contiguous nodes | [optional] 
**Container** | Pointer to **string** | Absolute path to OCI container bundle | [optional] 
**ContainerId** | Pointer to **string** | OCI container ID | [optional] 
**CoreSpecification** | Pointer to **int32** | Specialized core count | [optional] 
**ThreadSpecification** | Pointer to **int32** | Specialized thread count | [optional] 
**CpuBinding** | Pointer to **string** | Method for binding tasks to allocated CPUs | [optional] 
**CpuBindingFlags** | Pointer to **[]string** | Flags for CPU binding | [optional] 
**CpuFrequency** | Pointer to **string** | Requested CPU frequency range &lt;p1&gt;[-p2][:p3] | [optional] 
**CpusPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values values indicating how many CPUs should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Crontab** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerCrontab**](SlurmV0041PostJobSubmitRequestJobsInnerCrontab.md) |  | [optional] 
**Deadline** | Pointer to **int64** | Latest time that the job may start (UNIX timestamp) | [optional] 
**DelayBoot** | Pointer to **int32** | Number of seconds after job eligible start that nodes will be rebooted to satisfy feature specification | [optional] 
**Dependency** | Pointer to **string** | Other jobs that must meet certain criteria before this job can start | [optional] 
**EndTime** | Pointer to **int64** | Expected end time (UNIX timestamp) | [optional] 
**Environment** | Pointer to **[]string** | Environment variables to be set for the job | [optional] 
**Rlimits** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRlimits**](SlurmV0041PostJobSubmitRequestJobsInnerRlimits.md) |  | [optional] 
**ExcludedNodes** | Pointer to **[]string** | Comma separated list of nodes that may not be used | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Constraints** | Pointer to **string** | Comma separated list of features that are required | [optional] 
**GroupId** | Pointer to **string** | Group ID of the user that owns the job | [optional] 
**HetjobGroup** | Pointer to **int32** | Unique sequence number applied to this component of the heterogeneous job | [optional] 
**Immediate** | Pointer to **bool** | If true, exit if resources are not available within the time period specified | [optional] 
**JobId** | Pointer to **int32** | Job ID | [optional] 
**KillOnNodeFail** | Pointer to **bool** | If true, kill job on node failure | [optional] 
**Licenses** | Pointer to **string** | License(s) required by the job | [optional] 
**MailType** | Pointer to **[]string** | Mail event type(s) | [optional] 
**MailUser** | Pointer to **string** | User to receive email notifications | [optional] 
**McsLabel** | Pointer to **string** | Multi-Category Security label on the job | [optional] 
**MemoryBinding** | Pointer to **string** | Binding map for map/mask_cpu | [optional] 
**MemoryBindingType** | Pointer to **[]string** | Method for binding tasks to memory | [optional] 
**MemoryPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values indicating how much memory in megabytes should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Name** | Pointer to **string** | Job name | [optional] 
**Network** | Pointer to **string** | Network specs for job step | [optional] 
**Nice** | Pointer to **int32** | Requested job priority change | [optional] 
**Tasks** | Pointer to **int32** | Number of tasks | [optional] 
**OpenMode** | Pointer to **[]string** | Open mode used for stdout and stderr files | [optional] 
**ReservePorts** | Pointer to **int32** | Port to send various notification msg to | [optional] 
**Overcommit** | Pointer to **bool** | Overcommit resources | [optional] 
**Partition** | Pointer to **string** | Partition assigned to the job | [optional] 
**DistributionPlaneSize** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerDistributionPlaneSize**](SlurmV0041PostJobSubmitRequestJobsInnerDistributionPlaneSize.md) |  | [optional] 
**PowerFlags** | Pointer to **[]interface{}** |  | [optional] 
**Prefer** | Pointer to **string** | Comma separated list of features that are preferred but not required | [optional] 
**Hold** | Pointer to **bool** | Hold (true) or release (false) job | [optional] 
**Priority** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerPriority**](SlurmV0041PostJobSubmitRequestJobsInnerPriority.md) |  | [optional] 
**Profile** | Pointer to **[]string** | Profile used by the acct_gather_profile plugin | [optional] 
**Qos** | Pointer to **string** | Quality of Service assigned to the job | [optional] 
**Reboot** | Pointer to **bool** | Node reboot requested before start | [optional] 
**RequiredNodes** | Pointer to **[]string** | Comma separated list of required nodes | [optional] 
**Requeue** | Pointer to **bool** | Determines whether the job may be requeued | [optional] 
**Reservation** | Pointer to **string** | Name of reservation to use | [optional] 
**ResvMpiPorts** | Pointer to **int32** | Number of reserved communication ports; can only be used if slurmstepd step manager is enabled | [optional] 
**Script** | Pointer to **string** | Job batch script; only the first component in a HetJob is populated or honored | [optional] 
**Shared** | Pointer to **[]string** | How the job can share resources with other jobs, if at all | [optional] 
**Exclusive** | Pointer to **[]string** |  | [optional] 
**Oversubscribe** | Pointer to **bool** |  | [optional] 
**SiteFactor** | Pointer to **int32** | Site-specific priority factor | [optional] 
**SpankEnvironment** | Pointer to **[]string** | Environment variables for job prolog/epilog scripts as set by SPANK plugins | [optional] 
**Distribution** | Pointer to **string** | Layout | [optional] 
**TimeLimit** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerTimeLimit**](SlurmV0041PostJobSubmitRequestJobsInnerTimeLimit.md) |  | [optional] 
**TimeMinimum** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerTimeMinimum**](SlurmV0041PostJobSubmitRequestJobsInnerTimeMinimum.md) |  | [optional] 
**TresBind** | Pointer to **string** | Task to TRES binding directives | [optional] 
**TresFreq** | Pointer to **string** | TRES frequency directives | [optional] 
**TresPerJob** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every job | [optional] 
**TresPerNode** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every node | [optional] 
**TresPerSocket** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every socket | [optional] 
**TresPerTask** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every task | [optional] 
**UserId** | Pointer to **string** | User ID that owns the job | [optional] 
**WaitAllNodes** | Pointer to **bool** | If true, wait to start until after all nodes have booted | [optional] 
**KillWarningFlags** | Pointer to **[]string** | Flags related to job signals | [optional] 
**KillWarningSignal** | Pointer to **string** | Signal to send when approaching end time (e.g. \&quot;10\&quot; or \&quot;USR1\&quot;) | [optional] 
**KillWarningDelay** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerKillWarningDelay**](SlurmV0041PostJobSubmitRequestJobsInnerKillWarningDelay.md) |  | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** | Working directory to use for the job | [optional] 
**CpusPerTask** | Pointer to **int32** | Number of CPUs required by each task | [optional] 
**MinimumCpus** | Pointer to **int32** | Minimum number of CPUs required | [optional] 
**MaximumCpus** | Pointer to **int32** | Maximum number of CPUs required | [optional] 
**Nodes** | Pointer to **string** | Node count range specification (e.g. 1-15:4) | [optional] 
**MinimumNodes** | Pointer to **int32** | Minimum node count | [optional] 
**MaximumNodes** | Pointer to **int32** | Maximum node count | [optional] 
**MinimumBoardsPerNode** | Pointer to **int32** | Boards per node required | [optional] 
**MinimumSocketsPerBoard** | Pointer to **int32** | Sockets per board required | [optional] 
**SocketsPerNode** | Pointer to **int32** | Sockets per node required | [optional] 
**ThreadsPerCore** | Pointer to **int32** | Threads per core required | [optional] 
**TasksPerNode** | Pointer to **int32** | Number of tasks to invoke on each node | [optional] 
**TasksPerSocket** | Pointer to **int32** | Number of tasks to invoke on each socket | [optional] 
**TasksPerCore** | Pointer to **int32** | Number of tasks to invoke on each core | [optional] 
**TasksPerBoard** | Pointer to **int32** | Number of tasks to invoke on each board | [optional] 
**NtasksPerTres** | Pointer to **int32** | Number of tasks that can access each GPU | [optional] 
**MinimumCpusPerNode** | Pointer to **int32** | Minimum number of CPUs per node | [optional] 
**MemoryPerCpu** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu**](SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu**](SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu.md) |  | [optional] 
**TemporaryDiskPerNode** | Pointer to **int32** | Minimum tmp disk space required per node | [optional] 
**SelinuxContext** | Pointer to **string** | SELinux context | [optional] 
**RequiredSwitches** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerRequiredSwitches**](SlurmV0041PostJobSubmitRequestJobsInnerRequiredSwitches.md) |  | [optional] 
**SegmentSize** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerSegmentSize**](SlurmV0041PostJobSubmitRequestJobsInnerSegmentSize.md) |  | [optional] 
**StandardError** | Pointer to **string** | Path to stderr file | [optional] 
**StandardInput** | Pointer to **string** | Path to stdin file | [optional] 
**StandardOutput** | Pointer to **string** | Path to stdout file | [optional] 
**WaitForSwitch** | Pointer to **int32** | Maximum time to wait for switches in seconds | [optional] 
**Wckey** | Pointer to **string** | Workload characterization key | [optional] 
**X11** | Pointer to **[]string** | X11 forwarding options | [optional] 
**X11MagicCookie** | Pointer to **string** | Magic cookie for X11 forwarding | [optional] 
**X11TargetHost** | Pointer to **string** | Hostname or UNIX socket if x11_target_port&#x3D;0 | [optional] 
**X11TargetPort** | Pointer to **int32** | TCP port | [optional] 

## Methods

### NewSlurmV0041PostJobSubmitRequestJobsInner

`func NewSlurmV0041PostJobSubmitRequestJobsInner() *SlurmV0041PostJobSubmitRequestJobsInner`

NewSlurmV0041PostJobSubmitRequestJobsInner instantiates a new SlurmV0041PostJobSubmitRequestJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostJobSubmitRequestJobsInnerWithDefaults

`func NewSlurmV0041PostJobSubmitRequestJobsInnerWithDefaults() *SlurmV0041PostJobSubmitRequestJobsInner`

NewSlurmV0041PostJobSubmitRequestJobsInnerWithDefaults instantiates a new SlurmV0041PostJobSubmitRequestJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGatherFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAccountGatherFrequency() string`

GetAccountGatherFrequency returns the AccountGatherFrequency field if non-nil, zero value otherwise.

### GetAccountGatherFrequencyOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAccountGatherFrequencyOk() (*string, bool)`

GetAccountGatherFrequencyOk returns a tuple with the AccountGatherFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGatherFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetAccountGatherFrequency(v string)`

SetAccountGatherFrequency sets AccountGatherFrequency field to given value.

### HasAccountGatherFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasAccountGatherFrequency() bool`

HasAccountGatherFrequency returns a boolean if a field has been set.

### GetAdminComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocationNodeList

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAllocationNodeList() string`

GetAllocationNodeList returns the AllocationNodeList field if non-nil, zero value otherwise.

### GetAllocationNodeListOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAllocationNodeListOk() (*string, bool)`

GetAllocationNodeListOk returns a tuple with the AllocationNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodeList

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetAllocationNodeList(v string)`

SetAllocationNodeList sets AllocationNodeList field to given value.

### HasAllocationNodeList

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasAllocationNodeList() bool`

HasAllocationNodeList returns a boolean if a field has been set.

### GetAllocationNodePort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAllocationNodePort() int32`

GetAllocationNodePort returns the AllocationNodePort field if non-nil, zero value otherwise.

### GetAllocationNodePortOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetAllocationNodePortOk() (*int32, bool)`

GetAllocationNodePortOk returns a tuple with the AllocationNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodePort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetAllocationNodePort(v int32)`

SetAllocationNodePort sets AllocationNodePort field to given value.

### HasAllocationNodePort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasAllocationNodePort() bool`

HasAllocationNodePort returns a boolean if a field has been set.

### GetArgv

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetArray

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetArray() string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetArrayOk() (*string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetArray(v string)`

SetArray sets Array field to given value.

### HasArray

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBeginTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBeginTime() SlurmV0041PostJobSubmitRequestJobsInnerBeginTime`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBeginTimeOk() (*SlurmV0041PostJobSubmitRequestJobsInnerBeginTime, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetBeginTime(v SlurmV0041PostJobSubmitRequestJobsInnerBeginTime)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetClusters

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetClusters() string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetClustersOk() (*string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetClusters(v string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetClusterConstraint

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetClusterConstraint() string`

GetClusterConstraint returns the ClusterConstraint field if non-nil, zero value otherwise.

### GetClusterConstraintOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetClusterConstraintOk() (*string, bool)`

GetClusterConstraintOk returns a tuple with the ClusterConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConstraint

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetClusterConstraint(v string)`

SetClusterConstraint sets ClusterConstraint field to given value.

### HasClusterConstraint

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasClusterConstraint() bool`

HasClusterConstraint returns a boolean if a field has been set.

### GetComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContiguous

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetContainer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetCoreSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCoreSpecification() int32`

GetCoreSpecification returns the CoreSpecification field if non-nil, zero value otherwise.

### GetCoreSpecificationOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCoreSpecificationOk() (*int32, bool)`

GetCoreSpecificationOk returns a tuple with the CoreSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCoreSpecification(v int32)`

SetCoreSpecification sets CoreSpecification field to given value.

### HasCoreSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCoreSpecification() bool`

HasCoreSpecification returns a boolean if a field has been set.

### GetThreadSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetThreadSpecification() int32`

GetThreadSpecification returns the ThreadSpecification field if non-nil, zero value otherwise.

### GetThreadSpecificationOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetThreadSpecificationOk() (*int32, bool)`

GetThreadSpecificationOk returns a tuple with the ThreadSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetThreadSpecification(v int32)`

SetThreadSpecification sets ThreadSpecification field to given value.

### HasThreadSpecification

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasThreadSpecification() bool`

HasThreadSpecification returns a boolean if a field has been set.

### GetCpuBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuBinding() string`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuBindingOk() (*string, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCpuBinding(v string)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuBindingFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuBindingFlags() []string`

GetCpuBindingFlags returns the CpuBindingFlags field if non-nil, zero value otherwise.

### GetCpuBindingFlagsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuBindingFlagsOk() (*[]string, bool)`

GetCpuBindingFlagsOk returns a tuple with the CpuBindingFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBindingFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCpuBindingFlags(v []string)`

SetCpuBindingFlags sets CpuBindingFlags field to given value.

### HasCpuBindingFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCpuBindingFlags() bool`

HasCpuBindingFlags returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCrontab

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCrontab() SlurmV0041PostJobSubmitRequestJobsInnerCrontab`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCrontabOk() (*SlurmV0041PostJobSubmitRequestJobsInnerCrontab, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCrontab(v SlurmV0041PostJobSubmitRequestJobsInnerCrontab)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### GetDeadline

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDelayBoot() int32`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDelayBootOk() (*int32, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetDelayBoot(v int32)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetEndTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRlimits

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRlimits() SlurmV0041PostJobSubmitRequestJobsInnerRlimits`

GetRlimits returns the Rlimits field if non-nil, zero value otherwise.

### GetRlimitsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRlimitsOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRlimits, bool)`

GetRlimitsOk returns a tuple with the Rlimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlimits

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetRlimits(v SlurmV0041PostJobSubmitRequestJobsInnerRlimits)`

SetRlimits sets Rlimits field to given value.

### HasRlimits

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasRlimits() bool`

HasRlimits returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExcludedNodes() []string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExcludedNodesOk() (*[]string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetExcludedNodes(v []string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExtra

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetConstraints

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetGroupId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHetjobGroup

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetHetjobGroup() int32`

GetHetjobGroup returns the HetjobGroup field if non-nil, zero value otherwise.

### GetHetjobGroupOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetHetjobGroupOk() (*int32, bool)`

GetHetjobGroupOk returns a tuple with the HetjobGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjobGroup

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetHetjobGroup(v int32)`

SetHetjobGroup sets HetjobGroup field to given value.

### HasHetjobGroup

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasHetjobGroup() bool`

HasHetjobGroup returns a boolean if a field has been set.

### GetImmediate

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetJobId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetKillOnNodeFail

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillOnNodeFail() bool`

GetKillOnNodeFail returns the KillOnNodeFail field if non-nil, zero value otherwise.

### GetKillOnNodeFailOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillOnNodeFailOk() (*bool, bool)`

GetKillOnNodeFailOk returns a tuple with the KillOnNodeFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillOnNodeFail

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetKillOnNodeFail(v bool)`

SetKillOnNodeFail sets KillOnNodeFail field to given value.

### HasKillOnNodeFail

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasKillOnNodeFail() bool`

HasKillOnNodeFail returns a boolean if a field has been set.

### GetLicenses

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMcsLabel

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryBinding() string`

GetMemoryBinding returns the MemoryBinding field if non-nil, zero value otherwise.

### GetMemoryBindingOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryBindingOk() (*string, bool)`

GetMemoryBindingOk returns a tuple with the MemoryBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMemoryBinding(v string)`

SetMemoryBinding sets MemoryBinding field to given value.

### HasMemoryBinding

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMemoryBinding() bool`

HasMemoryBinding returns a boolean if a field has been set.

### GetMemoryBindingType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryBindingType() []string`

GetMemoryBindingType returns the MemoryBindingType field if non-nil, zero value otherwise.

### GetMemoryBindingTypeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryBindingTypeOk() (*[]string, bool)`

GetMemoryBindingTypeOk returns a tuple with the MemoryBindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBindingType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMemoryBindingType(v []string)`

SetMemoryBindingType sets MemoryBindingType field to given value.

### HasMemoryBindingType

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMemoryBindingType() bool`

HasMemoryBindingType returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNice

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasks

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasks() int32`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksOk() (*int32, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTasks(v int32)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOpenMode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOpenMode() []string`

GetOpenMode returns the OpenMode field if non-nil, zero value otherwise.

### GetOpenModeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOpenModeOk() (*[]string, bool)`

GetOpenModeOk returns a tuple with the OpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetOpenMode(v []string)`

SetOpenMode sets OpenMode field to given value.

### HasOpenMode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasOpenMode() bool`

HasOpenMode returns a boolean if a field has been set.

### GetReservePorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetReservePorts() int32`

GetReservePorts returns the ReservePorts field if non-nil, zero value otherwise.

### GetReservePortsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetReservePortsOk() (*int32, bool)`

GetReservePortsOk returns a tuple with the ReservePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetReservePorts(v int32)`

SetReservePorts sets ReservePorts field to given value.

### HasReservePorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasReservePorts() bool`

HasReservePorts returns a boolean if a field has been set.

### GetOvercommit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOvercommit() bool`

GetOvercommit returns the Overcommit field if non-nil, zero value otherwise.

### GetOvercommitOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOvercommitOk() (*bool, bool)`

GetOvercommitOk returns a tuple with the Overcommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetOvercommit(v bool)`

SetOvercommit sets Overcommit field to given value.

### HasOvercommit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasOvercommit() bool`

HasOvercommit returns a boolean if a field has been set.

### GetPartition

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetDistributionPlaneSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDistributionPlaneSize() SlurmV0041PostJobSubmitRequestJobsInnerDistributionPlaneSize`

GetDistributionPlaneSize returns the DistributionPlaneSize field if non-nil, zero value otherwise.

### GetDistributionPlaneSizeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDistributionPlaneSizeOk() (*SlurmV0041PostJobSubmitRequestJobsInnerDistributionPlaneSize, bool)`

GetDistributionPlaneSizeOk returns a tuple with the DistributionPlaneSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPlaneSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetDistributionPlaneSize(v SlurmV0041PostJobSubmitRequestJobsInnerDistributionPlaneSize)`

SetDistributionPlaneSize sets DistributionPlaneSize field to given value.

### HasDistributionPlaneSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasDistributionPlaneSize() bool`

HasDistributionPlaneSize returns a boolean if a field has been set.

### GetPowerFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPowerFlags() []interface{}`

GetPowerFlags returns the PowerFlags field if non-nil, zero value otherwise.

### GetPowerFlagsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPowerFlagsOk() (*[]interface{}, bool)`

GetPowerFlagsOk returns a tuple with the PowerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetPowerFlags(v []interface{})`

SetPowerFlags sets PowerFlags field to given value.

### HasPowerFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasPowerFlags() bool`

HasPowerFlags returns a boolean if a field has been set.

### GetPrefer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetHold

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPriority

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPriority() SlurmV0041PostJobSubmitRequestJobsInnerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetPriorityOk() (*SlurmV0041PostJobSubmitRequestJobsInnerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetPriority(v SlurmV0041PostJobSubmitRequestJobsInnerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequiredNodes() []string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequiredNodesOk() (*[]string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetRequiredNodes(v []string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetRequeue

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetReservation

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetResvMpiPorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetResvMpiPorts() int32`

GetResvMpiPorts returns the ResvMpiPorts field if non-nil, zero value otherwise.

### GetResvMpiPortsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetResvMpiPortsOk() (*int32, bool)`

GetResvMpiPortsOk returns a tuple with the ResvMpiPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvMpiPorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetResvMpiPorts(v int32)`

SetResvMpiPorts sets ResvMpiPorts field to given value.

### HasResvMpiPorts

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasResvMpiPorts() bool`

HasResvMpiPorts returns a boolean if a field has been set.

### GetScript

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetShared

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExclusive

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExclusive() []string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetExclusiveOk() (*[]string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetExclusive(v []string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetOversubscribe

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetSiteFactor

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSiteFactor() int32`

GetSiteFactor returns the SiteFactor field if non-nil, zero value otherwise.

### GetSiteFactorOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSiteFactorOk() (*int32, bool)`

GetSiteFactorOk returns a tuple with the SiteFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFactor

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetSiteFactor(v int32)`

SetSiteFactor sets SiteFactor field to given value.

### HasSiteFactor

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasSiteFactor() bool`

HasSiteFactor returns a boolean if a field has been set.

### GetSpankEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSpankEnvironment() []string`

GetSpankEnvironment returns the SpankEnvironment field if non-nil, zero value otherwise.

### GetSpankEnvironmentOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSpankEnvironmentOk() (*[]string, bool)`

GetSpankEnvironmentOk returns a tuple with the SpankEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpankEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetSpankEnvironment(v []string)`

SetSpankEnvironment sets SpankEnvironment field to given value.

### HasSpankEnvironment

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasSpankEnvironment() bool`

HasSpankEnvironment returns a boolean if a field has been set.

### GetDistribution

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetTimeLimit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTimeLimit() SlurmV0041PostJobSubmitRequestJobsInnerTimeLimit`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTimeLimitOk() (*SlurmV0041PostJobSubmitRequestJobsInnerTimeLimit, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTimeLimit(v SlurmV0041PostJobSubmitRequestJobsInnerTimeLimit)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTimeMinimum() SlurmV0041PostJobSubmitRequestJobsInnerTimeMinimum`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTimeMinimumOk() (*SlurmV0041PostJobSubmitRequestJobsInnerTimeMinimum, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTimeMinimum(v SlurmV0041PostJobSubmitRequestJobsInnerTimeMinimum)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetTresBind

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetUserId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitAllNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWaitAllNodes() bool`

GetWaitAllNodes returns the WaitAllNodes field if non-nil, zero value otherwise.

### GetWaitAllNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWaitAllNodesOk() (*bool, bool)`

GetWaitAllNodesOk returns a tuple with the WaitAllNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitAllNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetWaitAllNodes(v bool)`

SetWaitAllNodes sets WaitAllNodes field to given value.

### HasWaitAllNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasWaitAllNodes() bool`

HasWaitAllNodes returns a boolean if a field has been set.

### GetKillWarningFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningFlags() []string`

GetKillWarningFlags returns the KillWarningFlags field if non-nil, zero value otherwise.

### GetKillWarningFlagsOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningFlagsOk() (*[]string, bool)`

GetKillWarningFlagsOk returns a tuple with the KillWarningFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetKillWarningFlags(v []string)`

SetKillWarningFlags sets KillWarningFlags field to given value.

### HasKillWarningFlags

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasKillWarningFlags() bool`

HasKillWarningFlags returns a boolean if a field has been set.

### GetKillWarningSignal

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningSignal() string`

GetKillWarningSignal returns the KillWarningSignal field if non-nil, zero value otherwise.

### GetKillWarningSignalOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningSignalOk() (*string, bool)`

GetKillWarningSignalOk returns a tuple with the KillWarningSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningSignal

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetKillWarningSignal(v string)`

SetKillWarningSignal sets KillWarningSignal field to given value.

### HasKillWarningSignal

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasKillWarningSignal() bool`

HasKillWarningSignal returns a boolean if a field has been set.

### GetKillWarningDelay

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningDelay() SlurmV0041PostJobSubmitRequestJobsInnerKillWarningDelay`

GetKillWarningDelay returns the KillWarningDelay field if non-nil, zero value otherwise.

### GetKillWarningDelayOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetKillWarningDelayOk() (*SlurmV0041PostJobSubmitRequestJobsInnerKillWarningDelay, bool)`

GetKillWarningDelayOk returns a tuple with the KillWarningDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningDelay

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetKillWarningDelay(v SlurmV0041PostJobSubmitRequestJobsInnerKillWarningDelay)`

SetKillWarningDelay sets KillWarningDelay field to given value.

### HasKillWarningDelay

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasKillWarningDelay() bool`

HasKillWarningDelay returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpusPerTask() int32`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetCpusPerTaskOk() (*int32, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetCpusPerTask(v int32)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetMinimumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumCpus() int32`

GetMinimumCpus returns the MinimumCpus field if non-nil, zero value otherwise.

### GetMinimumCpusOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumCpusOk() (*int32, bool)`

GetMinimumCpusOk returns a tuple with the MinimumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMinimumCpus(v int32)`

SetMinimumCpus sets MinimumCpus field to given value.

### HasMinimumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMinimumCpus() bool`

HasMinimumCpus returns a boolean if a field has been set.

### GetMaximumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMaximumCpus() int32`

GetMaximumCpus returns the MaximumCpus field if non-nil, zero value otherwise.

### GetMaximumCpusOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMaximumCpusOk() (*int32, bool)`

GetMaximumCpusOk returns a tuple with the MaximumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMaximumCpus(v int32)`

SetMaximumCpus sets MaximumCpus field to given value.

### HasMaximumCpus

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMaximumCpus() bool`

HasMaximumCpus returns a boolean if a field has been set.

### GetNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetMinimumBoardsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumBoardsPerNode() int32`

GetMinimumBoardsPerNode returns the MinimumBoardsPerNode field if non-nil, zero value otherwise.

### GetMinimumBoardsPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumBoardsPerNodeOk() (*int32, bool)`

GetMinimumBoardsPerNodeOk returns a tuple with the MinimumBoardsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBoardsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMinimumBoardsPerNode(v int32)`

SetMinimumBoardsPerNode sets MinimumBoardsPerNode field to given value.

### HasMinimumBoardsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMinimumBoardsPerNode() bool`

HasMinimumBoardsPerNode returns a boolean if a field has been set.

### GetMinimumSocketsPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumSocketsPerBoard() int32`

GetMinimumSocketsPerBoard returns the MinimumSocketsPerBoard field if non-nil, zero value otherwise.

### GetMinimumSocketsPerBoardOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumSocketsPerBoardOk() (*int32, bool)`

GetMinimumSocketsPerBoardOk returns a tuple with the MinimumSocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSocketsPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMinimumSocketsPerBoard(v int32)`

SetMinimumSocketsPerBoard sets MinimumSocketsPerBoard field to given value.

### HasMinimumSocketsPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMinimumSocketsPerBoard() bool`

HasMinimumSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetThreadsPerCore() int32`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetThreadsPerCoreOk() (*int32, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetThreadsPerCore(v int32)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerNode() int32`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerNodeOk() (*int32, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTasksPerNode(v int32)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerSocket() int32`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerSocketOk() (*int32, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTasksPerSocket(v int32)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerCore() int32`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerCoreOk() (*int32, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTasksPerCore(v int32)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerBoard() int32`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTasksPerBoardOk() (*int32, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTasksPerBoard(v int32)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetNtasksPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNtasksPerTres() int32`

GetNtasksPerTres returns the NtasksPerTres field if non-nil, zero value otherwise.

### GetNtasksPerTresOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetNtasksPerTresOk() (*int32, bool)`

GetNtasksPerTresOk returns a tuple with the NtasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtasksPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetNtasksPerTres(v int32)`

SetNtasksPerTres sets NtasksPerTres field to given value.

### HasNtasksPerTres

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasNtasksPerTres() bool`

HasNtasksPerTres returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumCpusPerNode() int32`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMinimumCpusPerNodeOk() (*int32, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMinimumCpusPerNode(v int32)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerCpu() SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerCpuOk() (*SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMemoryPerCpu(v SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerNode() SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetMemoryPerNodeOk() (*SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetMemoryPerNode(v SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetTemporaryDiskPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTemporaryDiskPerNode() int32`

GetTemporaryDiskPerNode returns the TemporaryDiskPerNode field if non-nil, zero value otherwise.

### GetTemporaryDiskPerNodeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetTemporaryDiskPerNodeOk() (*int32, bool)`

GetTemporaryDiskPerNodeOk returns a tuple with the TemporaryDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDiskPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetTemporaryDiskPerNode(v int32)`

SetTemporaryDiskPerNode sets TemporaryDiskPerNode field to given value.

### HasTemporaryDiskPerNode

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasTemporaryDiskPerNode() bool`

HasTemporaryDiskPerNode returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetRequiredSwitches

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequiredSwitches() SlurmV0041PostJobSubmitRequestJobsInnerRequiredSwitches`

GetRequiredSwitches returns the RequiredSwitches field if non-nil, zero value otherwise.

### GetRequiredSwitchesOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetRequiredSwitchesOk() (*SlurmV0041PostJobSubmitRequestJobsInnerRequiredSwitches, bool)`

GetRequiredSwitchesOk returns a tuple with the RequiredSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSwitches

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetRequiredSwitches(v SlurmV0041PostJobSubmitRequestJobsInnerRequiredSwitches)`

SetRequiredSwitches sets RequiredSwitches field to given value.

### HasRequiredSwitches

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasRequiredSwitches() bool`

HasRequiredSwitches returns a boolean if a field has been set.

### GetSegmentSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSegmentSize() SlurmV0041PostJobSubmitRequestJobsInnerSegmentSize`

GetSegmentSize returns the SegmentSize field if non-nil, zero value otherwise.

### GetSegmentSizeOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetSegmentSizeOk() (*SlurmV0041PostJobSubmitRequestJobsInnerSegmentSize, bool)`

GetSegmentSizeOk returns a tuple with the SegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetSegmentSize(v SlurmV0041PostJobSubmitRequestJobsInnerSegmentSize)`

SetSegmentSize sets SegmentSize field to given value.

### HasSegmentSize

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasSegmentSize() bool`

HasSegmentSize returns a boolean if a field has been set.

### GetStandardError

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetWaitForSwitch

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWaitForSwitch() int32`

GetWaitForSwitch returns the WaitForSwitch field if non-nil, zero value otherwise.

### GetWaitForSwitchOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWaitForSwitchOk() (*int32, bool)`

GetWaitForSwitchOk returns a tuple with the WaitForSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForSwitch

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetWaitForSwitch(v int32)`

SetWaitForSwitch sets WaitForSwitch field to given value.

### HasWaitForSwitch

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasWaitForSwitch() bool`

HasWaitForSwitch returns a boolean if a field has been set.

### GetWckey

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetX11

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11() []string`

GetX11 returns the X11 field if non-nil, zero value otherwise.

### GetX11Ok

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11Ok() (*[]string, bool)`

GetX11Ok returns a tuple with the X11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetX11(v []string)`

SetX11 sets X11 field to given value.

### HasX11

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasX11() bool`

HasX11 returns a boolean if a field has been set.

### GetX11MagicCookie

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11MagicCookie() string`

GetX11MagicCookie returns the X11MagicCookie field if non-nil, zero value otherwise.

### GetX11MagicCookieOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11MagicCookieOk() (*string, bool)`

GetX11MagicCookieOk returns a tuple with the X11MagicCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11MagicCookie

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetX11MagicCookie(v string)`

SetX11MagicCookie sets X11MagicCookie field to given value.

### HasX11MagicCookie

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasX11MagicCookie() bool`

HasX11MagicCookie returns a boolean if a field has been set.

### GetX11TargetHost

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11TargetHost() string`

GetX11TargetHost returns the X11TargetHost field if non-nil, zero value otherwise.

### GetX11TargetHostOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11TargetHostOk() (*string, bool)`

GetX11TargetHostOk returns a tuple with the X11TargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetHost

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetX11TargetHost(v string)`

SetX11TargetHost sets X11TargetHost field to given value.

### HasX11TargetHost

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasX11TargetHost() bool`

HasX11TargetHost returns a boolean if a field has been set.

### GetX11TargetPort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11TargetPort() int32`

GetX11TargetPort returns the X11TargetPort field if non-nil, zero value otherwise.

### GetX11TargetPortOk

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) GetX11TargetPortOk() (*int32, bool)`

GetX11TargetPortOk returns a tuple with the X11TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetPort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) SetX11TargetPort(v int32)`

SetX11TargetPort sets X11TargetPort field to given value.

### HasX11TargetPort

`func (o *SlurmV0041PostJobSubmitRequestJobsInner) HasX11TargetPort() bool`

HasX11TargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


