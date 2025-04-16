# V0042JobDescMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account associated with the job | [optional] 
**AccountGatherFrequency** | Pointer to **string** | Job accounting and profiling sampling intervals in seconds | [optional] 
**AdminComment** | Pointer to **string** | Arbitrary comment made by administrator | [optional] 
**AllocationNodeList** | Pointer to **string** | Local node making the resource allocation | [optional] 
**AllocationNodePort** | Pointer to **int32** | Port to send allocation confirmation to | [optional] 
**Argv** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to **string** | Job array index value specification | [optional] 
**BatchFeatures** | Pointer to **string** | Features required for batch script&#39;s node | [optional] 
**BeginTime** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
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
**CpuBindingFlags** | Pointer to **[]string** |  | [optional] 
**CpuFrequency** | Pointer to **string** | Requested CPU frequency range &lt;p1&gt;[-p2][:p3] | [optional] 
**CpusPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values values indicating how many CPUs should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Crontab** | Pointer to [**V0042CronEntry**](V0042CronEntry.md) |  | [optional] 
**Deadline** | Pointer to **int64** | Latest time that the job may start (UNIX timestamp) | [optional] 
**DelayBoot** | Pointer to **int32** | Number of seconds after job eligible start that nodes will be rebooted to satisfy feature specification | [optional] 
**Dependency** | Pointer to **string** | Other jobs that must meet certain criteria before this job can start | [optional] 
**EndTime** | Pointer to **int64** | Expected end time (UNIX timestamp) | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**Rlimits** | Pointer to [**V0042JobDescMsgRlimits**](V0042JobDescMsgRlimits.md) |  | [optional] 
**ExcludedNodes** | Pointer to **[]string** |  | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Constraints** | Pointer to **string** | Comma separated list of features that are required | [optional] 
**GroupId** | Pointer to **string** | Group ID of the user that owns the job | [optional] 
**HetjobGroup** | Pointer to **int32** | Unique sequence number applied to this component of the heterogeneous job | [optional] 
**Immediate** | Pointer to **bool** | If true, exit if resources are not available within the time period specified | [optional] 
**JobId** | Pointer to **int32** | Job ID | [optional] 
**KillOnNodeFail** | Pointer to **bool** | If true, kill job on node failure | [optional] 
**Licenses** | Pointer to **string** | License(s) required by the job | [optional] 
**MailType** | Pointer to **[]string** |  | [optional] 
**MailUser** | Pointer to **string** | User to receive email notifications | [optional] 
**McsLabel** | Pointer to **string** | Multi-Category Security label on the job | [optional] 
**MemoryBinding** | Pointer to **string** | Binding map for map/mask_cpu | [optional] 
**MemoryBindingType** | Pointer to **[]string** |  | [optional] 
**MemoryPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values indicating how much memory in megabytes should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Name** | Pointer to **string** | Job name | [optional] 
**Network** | Pointer to **string** | Network specs for job step | [optional] 
**Nice** | Pointer to **int32** | Requested job priority change | [optional] 
**Tasks** | Pointer to **int32** | Number of tasks | [optional] 
**OomKillStep** | Pointer to **int32** | Kill whole step in case of OOM in one of the tasks | [optional] 
**OpenMode** | Pointer to **[]string** |  | [optional] 
**ReservePorts** | Pointer to **int32** | Port to send various notification msg to | [optional] 
**Overcommit** | Pointer to **bool** | Overcommit resources | [optional] 
**Partition** | Pointer to **string** | Partition assigned to the job | [optional] 
**DistributionPlaneSize** | Pointer to [**V0042Uint16NoValStruct**](V0042Uint16NoValStruct.md) |  | [optional] 
**PowerFlags** | Pointer to **[]interface{}** |  | [optional] 
**Prefer** | Pointer to **string** | Comma separated list of features that are preferred but not required | [optional] 
**Hold** | Pointer to **bool** | Hold (true) or release (false) job | [optional] 
**Priority** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 
**Qos** | Pointer to **string** | Quality of Service assigned to the job | [optional] 
**Reboot** | Pointer to **bool** | Node reboot requested before start | [optional] 
**RequiredNodes** | Pointer to **[]string** |  | [optional] 
**Requeue** | Pointer to **bool** | Determines whether the job may be requeued | [optional] 
**Reservation** | Pointer to **string** | Name of reservation to use | [optional] 
**Script** | Pointer to **string** | Job batch script; only the first component in a HetJob is populated or honored | [optional] 
**Shared** | Pointer to **[]string** |  | [optional] 
**SiteFactor** | Pointer to **int32** | Site-specific priority factor | [optional] 
**SpankEnvironment** | Pointer to **[]string** |  | [optional] 
**Distribution** | Pointer to **string** | Layout | [optional] 
**TimeLimit** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**TresBind** | Pointer to **string** | Task to TRES binding directives | [optional] 
**TresFreq** | Pointer to **string** | TRES frequency directives | [optional] 
**TresPerJob** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every job | [optional] 
**TresPerNode** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every node | [optional] 
**TresPerSocket** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every socket | [optional] 
**TresPerTask** | Pointer to **string** | Comma separated list of TRES&#x3D;# values to be allocated for every task | [optional] 
**UserId** | Pointer to **string** | User ID that owns the job | [optional] 
**WaitAllNodes** | Pointer to **bool** | If true, wait to start until after all nodes have booted | [optional] 
**KillWarningFlags** | Pointer to **[]string** |  | [optional] 
**KillWarningSignal** | Pointer to **string** | Signal to send when approaching end time (e.g. \&quot;10\&quot; or \&quot;USR1\&quot;) | [optional] 
**KillWarningDelay** | Pointer to [**V0042Uint16NoValStruct**](V0042Uint16NoValStruct.md) |  | [optional] 
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
**MemoryPerCpu** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**TemporaryDiskPerNode** | Pointer to **int32** | Minimum tmp disk space required per node | [optional] 
**SelinuxContext** | Pointer to **string** | SELinux context | [optional] 
**RequiredSwitches** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**SegmentSize** | Pointer to [**V0042Uint16NoValStruct**](V0042Uint16NoValStruct.md) |  | [optional] 
**StandardError** | Pointer to **string** | Path to stderr file | [optional] 
**StandardInput** | Pointer to **string** | Path to stdin file | [optional] 
**StandardOutput** | Pointer to **string** | Path to stdout file | [optional] 
**WaitForSwitch** | Pointer to **int32** | Maximum time to wait for switches in seconds | [optional] 
**Wckey** | Pointer to **string** | Workload characterization key | [optional] 
**X11** | Pointer to **[]string** |  | [optional] 
**X11MagicCookie** | Pointer to **string** | Magic cookie for X11 forwarding | [optional] 
**X11TargetHost** | Pointer to **string** | Hostname or UNIX socket if x11_target_port&#x3D;0 | [optional] 
**X11TargetPort** | Pointer to **int32** | TCP port | [optional] 

## Methods

### NewV0042JobDescMsg

`func NewV0042JobDescMsg() *V0042JobDescMsg`

NewV0042JobDescMsg instantiates a new V0042JobDescMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobDescMsgWithDefaults

`func NewV0042JobDescMsgWithDefaults() *V0042JobDescMsg`

NewV0042JobDescMsgWithDefaults instantiates a new V0042JobDescMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0042JobDescMsg) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0042JobDescMsg) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0042JobDescMsg) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0042JobDescMsg) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGatherFrequency

`func (o *V0042JobDescMsg) GetAccountGatherFrequency() string`

GetAccountGatherFrequency returns the AccountGatherFrequency field if non-nil, zero value otherwise.

### GetAccountGatherFrequencyOk

`func (o *V0042JobDescMsg) GetAccountGatherFrequencyOk() (*string, bool)`

GetAccountGatherFrequencyOk returns a tuple with the AccountGatherFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGatherFrequency

`func (o *V0042JobDescMsg) SetAccountGatherFrequency(v string)`

SetAccountGatherFrequency sets AccountGatherFrequency field to given value.

### HasAccountGatherFrequency

`func (o *V0042JobDescMsg) HasAccountGatherFrequency() bool`

HasAccountGatherFrequency returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0042JobDescMsg) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0042JobDescMsg) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0042JobDescMsg) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0042JobDescMsg) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocationNodeList

`func (o *V0042JobDescMsg) GetAllocationNodeList() string`

GetAllocationNodeList returns the AllocationNodeList field if non-nil, zero value otherwise.

### GetAllocationNodeListOk

`func (o *V0042JobDescMsg) GetAllocationNodeListOk() (*string, bool)`

GetAllocationNodeListOk returns a tuple with the AllocationNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodeList

`func (o *V0042JobDescMsg) SetAllocationNodeList(v string)`

SetAllocationNodeList sets AllocationNodeList field to given value.

### HasAllocationNodeList

`func (o *V0042JobDescMsg) HasAllocationNodeList() bool`

HasAllocationNodeList returns a boolean if a field has been set.

### GetAllocationNodePort

`func (o *V0042JobDescMsg) GetAllocationNodePort() int32`

GetAllocationNodePort returns the AllocationNodePort field if non-nil, zero value otherwise.

### GetAllocationNodePortOk

`func (o *V0042JobDescMsg) GetAllocationNodePortOk() (*int32, bool)`

GetAllocationNodePortOk returns a tuple with the AllocationNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodePort

`func (o *V0042JobDescMsg) SetAllocationNodePort(v int32)`

SetAllocationNodePort sets AllocationNodePort field to given value.

### HasAllocationNodePort

`func (o *V0042JobDescMsg) HasAllocationNodePort() bool`

HasAllocationNodePort returns a boolean if a field has been set.

### GetArgv

`func (o *V0042JobDescMsg) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *V0042JobDescMsg) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *V0042JobDescMsg) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *V0042JobDescMsg) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetArray

`func (o *V0042JobDescMsg) GetArray() string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *V0042JobDescMsg) GetArrayOk() (*string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *V0042JobDescMsg) SetArray(v string)`

SetArray sets Array field to given value.

### HasArray

`func (o *V0042JobDescMsg) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0042JobDescMsg) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0042JobDescMsg) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0042JobDescMsg) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0042JobDescMsg) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBeginTime

`func (o *V0042JobDescMsg) GetBeginTime() V0042Uint64NoValStruct`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *V0042JobDescMsg) GetBeginTimeOk() (*V0042Uint64NoValStruct, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *V0042JobDescMsg) SetBeginTime(v V0042Uint64NoValStruct)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *V0042JobDescMsg) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetFlags

`func (o *V0042JobDescMsg) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0042JobDescMsg) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0042JobDescMsg) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0042JobDescMsg) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0042JobDescMsg) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0042JobDescMsg) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0042JobDescMsg) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0042JobDescMsg) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetClusters

`func (o *V0042JobDescMsg) GetClusters() string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0042JobDescMsg) GetClustersOk() (*string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0042JobDescMsg) SetClusters(v string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0042JobDescMsg) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetClusterConstraint

`func (o *V0042JobDescMsg) GetClusterConstraint() string`

GetClusterConstraint returns the ClusterConstraint field if non-nil, zero value otherwise.

### GetClusterConstraintOk

`func (o *V0042JobDescMsg) GetClusterConstraintOk() (*string, bool)`

GetClusterConstraintOk returns a tuple with the ClusterConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConstraint

`func (o *V0042JobDescMsg) SetClusterConstraint(v string)`

SetClusterConstraint sets ClusterConstraint field to given value.

### HasClusterConstraint

`func (o *V0042JobDescMsg) HasClusterConstraint() bool`

HasClusterConstraint returns a boolean if a field has been set.

### GetComment

`func (o *V0042JobDescMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0042JobDescMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0042JobDescMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0042JobDescMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContiguous

`func (o *V0042JobDescMsg) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0042JobDescMsg) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0042JobDescMsg) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0042JobDescMsg) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetContainer

`func (o *V0042JobDescMsg) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0042JobDescMsg) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0042JobDescMsg) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0042JobDescMsg) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0042JobDescMsg) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0042JobDescMsg) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0042JobDescMsg) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0042JobDescMsg) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetCoreSpecification

`func (o *V0042JobDescMsg) GetCoreSpecification() int32`

GetCoreSpecification returns the CoreSpecification field if non-nil, zero value otherwise.

### GetCoreSpecificationOk

`func (o *V0042JobDescMsg) GetCoreSpecificationOk() (*int32, bool)`

GetCoreSpecificationOk returns a tuple with the CoreSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecification

`func (o *V0042JobDescMsg) SetCoreSpecification(v int32)`

SetCoreSpecification sets CoreSpecification field to given value.

### HasCoreSpecification

`func (o *V0042JobDescMsg) HasCoreSpecification() bool`

HasCoreSpecification returns a boolean if a field has been set.

### GetThreadSpecification

`func (o *V0042JobDescMsg) GetThreadSpecification() int32`

GetThreadSpecification returns the ThreadSpecification field if non-nil, zero value otherwise.

### GetThreadSpecificationOk

`func (o *V0042JobDescMsg) GetThreadSpecificationOk() (*int32, bool)`

GetThreadSpecificationOk returns a tuple with the ThreadSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpecification

`func (o *V0042JobDescMsg) SetThreadSpecification(v int32)`

SetThreadSpecification sets ThreadSpecification field to given value.

### HasThreadSpecification

`func (o *V0042JobDescMsg) HasThreadSpecification() bool`

HasThreadSpecification returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0042JobDescMsg) GetCpuBinding() string`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0042JobDescMsg) GetCpuBindingOk() (*string, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0042JobDescMsg) SetCpuBinding(v string)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0042JobDescMsg) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuBindingFlags

`func (o *V0042JobDescMsg) GetCpuBindingFlags() []string`

GetCpuBindingFlags returns the CpuBindingFlags field if non-nil, zero value otherwise.

### GetCpuBindingFlagsOk

`func (o *V0042JobDescMsg) GetCpuBindingFlagsOk() (*[]string, bool)`

GetCpuBindingFlagsOk returns a tuple with the CpuBindingFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBindingFlags

`func (o *V0042JobDescMsg) SetCpuBindingFlags(v []string)`

SetCpuBindingFlags sets CpuBindingFlags field to given value.

### HasCpuBindingFlags

`func (o *V0042JobDescMsg) HasCpuBindingFlags() bool`

HasCpuBindingFlags returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *V0042JobDescMsg) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *V0042JobDescMsg) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *V0042JobDescMsg) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *V0042JobDescMsg) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0042JobDescMsg) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0042JobDescMsg) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0042JobDescMsg) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0042JobDescMsg) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCrontab

`func (o *V0042JobDescMsg) GetCrontab() V0042CronEntry`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *V0042JobDescMsg) GetCrontabOk() (*V0042CronEntry, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *V0042JobDescMsg) SetCrontab(v V0042CronEntry)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *V0042JobDescMsg) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### GetDeadline

`func (o *V0042JobDescMsg) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0042JobDescMsg) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0042JobDescMsg) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0042JobDescMsg) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0042JobDescMsg) GetDelayBoot() int32`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0042JobDescMsg) GetDelayBootOk() (*int32, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0042JobDescMsg) SetDelayBoot(v int32)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0042JobDescMsg) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0042JobDescMsg) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0042JobDescMsg) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0042JobDescMsg) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0042JobDescMsg) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetEndTime

`func (o *V0042JobDescMsg) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0042JobDescMsg) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0042JobDescMsg) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0042JobDescMsg) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *V0042JobDescMsg) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V0042JobDescMsg) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V0042JobDescMsg) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *V0042JobDescMsg) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRlimits

`func (o *V0042JobDescMsg) GetRlimits() V0042JobDescMsgRlimits`

GetRlimits returns the Rlimits field if non-nil, zero value otherwise.

### GetRlimitsOk

`func (o *V0042JobDescMsg) GetRlimitsOk() (*V0042JobDescMsgRlimits, bool)`

GetRlimitsOk returns a tuple with the Rlimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlimits

`func (o *V0042JobDescMsg) SetRlimits(v V0042JobDescMsgRlimits)`

SetRlimits sets Rlimits field to given value.

### HasRlimits

`func (o *V0042JobDescMsg) HasRlimits() bool`

HasRlimits returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0042JobDescMsg) GetExcludedNodes() []string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0042JobDescMsg) GetExcludedNodesOk() (*[]string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0042JobDescMsg) SetExcludedNodes(v []string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0042JobDescMsg) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExtra

`func (o *V0042JobDescMsg) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0042JobDescMsg) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0042JobDescMsg) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0042JobDescMsg) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetConstraints

`func (o *V0042JobDescMsg) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *V0042JobDescMsg) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *V0042JobDescMsg) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *V0042JobDescMsg) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetGroupId

`func (o *V0042JobDescMsg) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0042JobDescMsg) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0042JobDescMsg) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0042JobDescMsg) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHetjobGroup

`func (o *V0042JobDescMsg) GetHetjobGroup() int32`

GetHetjobGroup returns the HetjobGroup field if non-nil, zero value otherwise.

### GetHetjobGroupOk

`func (o *V0042JobDescMsg) GetHetjobGroupOk() (*int32, bool)`

GetHetjobGroupOk returns a tuple with the HetjobGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjobGroup

`func (o *V0042JobDescMsg) SetHetjobGroup(v int32)`

SetHetjobGroup sets HetjobGroup field to given value.

### HasHetjobGroup

`func (o *V0042JobDescMsg) HasHetjobGroup() bool`

HasHetjobGroup returns a boolean if a field has been set.

### GetImmediate

`func (o *V0042JobDescMsg) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *V0042JobDescMsg) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *V0042JobDescMsg) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *V0042JobDescMsg) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetJobId

`func (o *V0042JobDescMsg) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0042JobDescMsg) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0042JobDescMsg) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0042JobDescMsg) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetKillOnNodeFail

`func (o *V0042JobDescMsg) GetKillOnNodeFail() bool`

GetKillOnNodeFail returns the KillOnNodeFail field if non-nil, zero value otherwise.

### GetKillOnNodeFailOk

`func (o *V0042JobDescMsg) GetKillOnNodeFailOk() (*bool, bool)`

GetKillOnNodeFailOk returns a tuple with the KillOnNodeFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillOnNodeFail

`func (o *V0042JobDescMsg) SetKillOnNodeFail(v bool)`

SetKillOnNodeFail sets KillOnNodeFail field to given value.

### HasKillOnNodeFail

`func (o *V0042JobDescMsg) HasKillOnNodeFail() bool`

HasKillOnNodeFail returns a boolean if a field has been set.

### GetLicenses

`func (o *V0042JobDescMsg) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0042JobDescMsg) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0042JobDescMsg) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0042JobDescMsg) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0042JobDescMsg) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0042JobDescMsg) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0042JobDescMsg) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0042JobDescMsg) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0042JobDescMsg) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0042JobDescMsg) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0042JobDescMsg) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0042JobDescMsg) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0042JobDescMsg) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0042JobDescMsg) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0042JobDescMsg) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0042JobDescMsg) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryBinding

`func (o *V0042JobDescMsg) GetMemoryBinding() string`

GetMemoryBinding returns the MemoryBinding field if non-nil, zero value otherwise.

### GetMemoryBindingOk

`func (o *V0042JobDescMsg) GetMemoryBindingOk() (*string, bool)`

GetMemoryBindingOk returns a tuple with the MemoryBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBinding

`func (o *V0042JobDescMsg) SetMemoryBinding(v string)`

SetMemoryBinding sets MemoryBinding field to given value.

### HasMemoryBinding

`func (o *V0042JobDescMsg) HasMemoryBinding() bool`

HasMemoryBinding returns a boolean if a field has been set.

### GetMemoryBindingType

`func (o *V0042JobDescMsg) GetMemoryBindingType() []string`

GetMemoryBindingType returns the MemoryBindingType field if non-nil, zero value otherwise.

### GetMemoryBindingTypeOk

`func (o *V0042JobDescMsg) GetMemoryBindingTypeOk() (*[]string, bool)`

GetMemoryBindingTypeOk returns a tuple with the MemoryBindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBindingType

`func (o *V0042JobDescMsg) SetMemoryBindingType(v []string)`

SetMemoryBindingType sets MemoryBindingType field to given value.

### HasMemoryBindingType

`func (o *V0042JobDescMsg) HasMemoryBindingType() bool`

HasMemoryBindingType returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0042JobDescMsg) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0042JobDescMsg) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0042JobDescMsg) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0042JobDescMsg) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0042JobDescMsg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042JobDescMsg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042JobDescMsg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0042JobDescMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0042JobDescMsg) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0042JobDescMsg) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0042JobDescMsg) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0042JobDescMsg) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNice

`func (o *V0042JobDescMsg) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0042JobDescMsg) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0042JobDescMsg) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0042JobDescMsg) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasks

`func (o *V0042JobDescMsg) GetTasks() int32`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0042JobDescMsg) GetTasksOk() (*int32, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0042JobDescMsg) SetTasks(v int32)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0042JobDescMsg) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOomKillStep

`func (o *V0042JobDescMsg) GetOomKillStep() int32`

GetOomKillStep returns the OomKillStep field if non-nil, zero value otherwise.

### GetOomKillStepOk

`func (o *V0042JobDescMsg) GetOomKillStepOk() (*int32, bool)`

GetOomKillStepOk returns a tuple with the OomKillStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOomKillStep

`func (o *V0042JobDescMsg) SetOomKillStep(v int32)`

SetOomKillStep sets OomKillStep field to given value.

### HasOomKillStep

`func (o *V0042JobDescMsg) HasOomKillStep() bool`

HasOomKillStep returns a boolean if a field has been set.

### GetOpenMode

`func (o *V0042JobDescMsg) GetOpenMode() []string`

GetOpenMode returns the OpenMode field if non-nil, zero value otherwise.

### GetOpenModeOk

`func (o *V0042JobDescMsg) GetOpenModeOk() (*[]string, bool)`

GetOpenModeOk returns a tuple with the OpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMode

`func (o *V0042JobDescMsg) SetOpenMode(v []string)`

SetOpenMode sets OpenMode field to given value.

### HasOpenMode

`func (o *V0042JobDescMsg) HasOpenMode() bool`

HasOpenMode returns a boolean if a field has been set.

### GetReservePorts

`func (o *V0042JobDescMsg) GetReservePorts() int32`

GetReservePorts returns the ReservePorts field if non-nil, zero value otherwise.

### GetReservePortsOk

`func (o *V0042JobDescMsg) GetReservePortsOk() (*int32, bool)`

GetReservePortsOk returns a tuple with the ReservePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePorts

`func (o *V0042JobDescMsg) SetReservePorts(v int32)`

SetReservePorts sets ReservePorts field to given value.

### HasReservePorts

`func (o *V0042JobDescMsg) HasReservePorts() bool`

HasReservePorts returns a boolean if a field has been set.

### GetOvercommit

`func (o *V0042JobDescMsg) GetOvercommit() bool`

GetOvercommit returns the Overcommit field if non-nil, zero value otherwise.

### GetOvercommitOk

`func (o *V0042JobDescMsg) GetOvercommitOk() (*bool, bool)`

GetOvercommitOk returns a tuple with the Overcommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommit

`func (o *V0042JobDescMsg) SetOvercommit(v bool)`

SetOvercommit sets Overcommit field to given value.

### HasOvercommit

`func (o *V0042JobDescMsg) HasOvercommit() bool`

HasOvercommit returns a boolean if a field has been set.

### GetPartition

`func (o *V0042JobDescMsg) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0042JobDescMsg) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0042JobDescMsg) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0042JobDescMsg) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetDistributionPlaneSize

`func (o *V0042JobDescMsg) GetDistributionPlaneSize() V0042Uint16NoValStruct`

GetDistributionPlaneSize returns the DistributionPlaneSize field if non-nil, zero value otherwise.

### GetDistributionPlaneSizeOk

`func (o *V0042JobDescMsg) GetDistributionPlaneSizeOk() (*V0042Uint16NoValStruct, bool)`

GetDistributionPlaneSizeOk returns a tuple with the DistributionPlaneSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPlaneSize

`func (o *V0042JobDescMsg) SetDistributionPlaneSize(v V0042Uint16NoValStruct)`

SetDistributionPlaneSize sets DistributionPlaneSize field to given value.

### HasDistributionPlaneSize

`func (o *V0042JobDescMsg) HasDistributionPlaneSize() bool`

HasDistributionPlaneSize returns a boolean if a field has been set.

### GetPowerFlags

`func (o *V0042JobDescMsg) GetPowerFlags() []interface{}`

GetPowerFlags returns the PowerFlags field if non-nil, zero value otherwise.

### GetPowerFlagsOk

`func (o *V0042JobDescMsg) GetPowerFlagsOk() (*[]interface{}, bool)`

GetPowerFlagsOk returns a tuple with the PowerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFlags

`func (o *V0042JobDescMsg) SetPowerFlags(v []interface{})`

SetPowerFlags sets PowerFlags field to given value.

### HasPowerFlags

`func (o *V0042JobDescMsg) HasPowerFlags() bool`

HasPowerFlags returns a boolean if a field has been set.

### GetPrefer

`func (o *V0042JobDescMsg) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0042JobDescMsg) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0042JobDescMsg) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0042JobDescMsg) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetHold

`func (o *V0042JobDescMsg) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *V0042JobDescMsg) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *V0042JobDescMsg) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *V0042JobDescMsg) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPriority

`func (o *V0042JobDescMsg) GetPriority() V0042Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0042JobDescMsg) GetPriorityOk() (*V0042Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0042JobDescMsg) SetPriority(v V0042Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0042JobDescMsg) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0042JobDescMsg) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0042JobDescMsg) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0042JobDescMsg) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0042JobDescMsg) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0042JobDescMsg) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0042JobDescMsg) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0042JobDescMsg) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0042JobDescMsg) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0042JobDescMsg) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0042JobDescMsg) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0042JobDescMsg) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0042JobDescMsg) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0042JobDescMsg) GetRequiredNodes() []string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0042JobDescMsg) GetRequiredNodesOk() (*[]string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0042JobDescMsg) SetRequiredNodes(v []string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0042JobDescMsg) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetRequeue

`func (o *V0042JobDescMsg) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0042JobDescMsg) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0042JobDescMsg) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0042JobDescMsg) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetReservation

`func (o *V0042JobDescMsg) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0042JobDescMsg) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0042JobDescMsg) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0042JobDescMsg) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetScript

`func (o *V0042JobDescMsg) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0042JobDescMsg) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0042JobDescMsg) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0042JobDescMsg) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetShared

`func (o *V0042JobDescMsg) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0042JobDescMsg) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0042JobDescMsg) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0042JobDescMsg) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSiteFactor

`func (o *V0042JobDescMsg) GetSiteFactor() int32`

GetSiteFactor returns the SiteFactor field if non-nil, zero value otherwise.

### GetSiteFactorOk

`func (o *V0042JobDescMsg) GetSiteFactorOk() (*int32, bool)`

GetSiteFactorOk returns a tuple with the SiteFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFactor

`func (o *V0042JobDescMsg) SetSiteFactor(v int32)`

SetSiteFactor sets SiteFactor field to given value.

### HasSiteFactor

`func (o *V0042JobDescMsg) HasSiteFactor() bool`

HasSiteFactor returns a boolean if a field has been set.

### GetSpankEnvironment

`func (o *V0042JobDescMsg) GetSpankEnvironment() []string`

GetSpankEnvironment returns the SpankEnvironment field if non-nil, zero value otherwise.

### GetSpankEnvironmentOk

`func (o *V0042JobDescMsg) GetSpankEnvironmentOk() (*[]string, bool)`

GetSpankEnvironmentOk returns a tuple with the SpankEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpankEnvironment

`func (o *V0042JobDescMsg) SetSpankEnvironment(v []string)`

SetSpankEnvironment sets SpankEnvironment field to given value.

### HasSpankEnvironment

`func (o *V0042JobDescMsg) HasSpankEnvironment() bool`

HasSpankEnvironment returns a boolean if a field has been set.

### GetDistribution

`func (o *V0042JobDescMsg) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *V0042JobDescMsg) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *V0042JobDescMsg) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *V0042JobDescMsg) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0042JobDescMsg) GetTimeLimit() V0042Uint32NoValStruct`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0042JobDescMsg) GetTimeLimitOk() (*V0042Uint32NoValStruct, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0042JobDescMsg) SetTimeLimit(v V0042Uint32NoValStruct)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0042JobDescMsg) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0042JobDescMsg) GetTimeMinimum() V0042Uint32NoValStruct`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0042JobDescMsg) GetTimeMinimumOk() (*V0042Uint32NoValStruct, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0042JobDescMsg) SetTimeMinimum(v V0042Uint32NoValStruct)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0042JobDescMsg) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetTresBind

`func (o *V0042JobDescMsg) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0042JobDescMsg) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0042JobDescMsg) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0042JobDescMsg) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0042JobDescMsg) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0042JobDescMsg) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0042JobDescMsg) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0042JobDescMsg) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0042JobDescMsg) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0042JobDescMsg) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0042JobDescMsg) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0042JobDescMsg) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0042JobDescMsg) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0042JobDescMsg) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0042JobDescMsg) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0042JobDescMsg) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0042JobDescMsg) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0042JobDescMsg) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0042JobDescMsg) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0042JobDescMsg) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0042JobDescMsg) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0042JobDescMsg) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0042JobDescMsg) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0042JobDescMsg) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetUserId

`func (o *V0042JobDescMsg) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0042JobDescMsg) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0042JobDescMsg) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0042JobDescMsg) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitAllNodes

`func (o *V0042JobDescMsg) GetWaitAllNodes() bool`

GetWaitAllNodes returns the WaitAllNodes field if non-nil, zero value otherwise.

### GetWaitAllNodesOk

`func (o *V0042JobDescMsg) GetWaitAllNodesOk() (*bool, bool)`

GetWaitAllNodesOk returns a tuple with the WaitAllNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitAllNodes

`func (o *V0042JobDescMsg) SetWaitAllNodes(v bool)`

SetWaitAllNodes sets WaitAllNodes field to given value.

### HasWaitAllNodes

`func (o *V0042JobDescMsg) HasWaitAllNodes() bool`

HasWaitAllNodes returns a boolean if a field has been set.

### GetKillWarningFlags

`func (o *V0042JobDescMsg) GetKillWarningFlags() []string`

GetKillWarningFlags returns the KillWarningFlags field if non-nil, zero value otherwise.

### GetKillWarningFlagsOk

`func (o *V0042JobDescMsg) GetKillWarningFlagsOk() (*[]string, bool)`

GetKillWarningFlagsOk returns a tuple with the KillWarningFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningFlags

`func (o *V0042JobDescMsg) SetKillWarningFlags(v []string)`

SetKillWarningFlags sets KillWarningFlags field to given value.

### HasKillWarningFlags

`func (o *V0042JobDescMsg) HasKillWarningFlags() bool`

HasKillWarningFlags returns a boolean if a field has been set.

### GetKillWarningSignal

`func (o *V0042JobDescMsg) GetKillWarningSignal() string`

GetKillWarningSignal returns the KillWarningSignal field if non-nil, zero value otherwise.

### GetKillWarningSignalOk

`func (o *V0042JobDescMsg) GetKillWarningSignalOk() (*string, bool)`

GetKillWarningSignalOk returns a tuple with the KillWarningSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningSignal

`func (o *V0042JobDescMsg) SetKillWarningSignal(v string)`

SetKillWarningSignal sets KillWarningSignal field to given value.

### HasKillWarningSignal

`func (o *V0042JobDescMsg) HasKillWarningSignal() bool`

HasKillWarningSignal returns a boolean if a field has been set.

### GetKillWarningDelay

`func (o *V0042JobDescMsg) GetKillWarningDelay() V0042Uint16NoValStruct`

GetKillWarningDelay returns the KillWarningDelay field if non-nil, zero value otherwise.

### GetKillWarningDelayOk

`func (o *V0042JobDescMsg) GetKillWarningDelayOk() (*V0042Uint16NoValStruct, bool)`

GetKillWarningDelayOk returns a tuple with the KillWarningDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningDelay

`func (o *V0042JobDescMsg) SetKillWarningDelay(v V0042Uint16NoValStruct)`

SetKillWarningDelay sets KillWarningDelay field to given value.

### HasKillWarningDelay

`func (o *V0042JobDescMsg) HasKillWarningDelay() bool`

HasKillWarningDelay returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0042JobDescMsg) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0042JobDescMsg) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0042JobDescMsg) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0042JobDescMsg) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0042JobDescMsg) GetCpusPerTask() int32`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0042JobDescMsg) GetCpusPerTaskOk() (*int32, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0042JobDescMsg) SetCpusPerTask(v int32)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0042JobDescMsg) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetMinimumCpus

`func (o *V0042JobDescMsg) GetMinimumCpus() int32`

GetMinimumCpus returns the MinimumCpus field if non-nil, zero value otherwise.

### GetMinimumCpusOk

`func (o *V0042JobDescMsg) GetMinimumCpusOk() (*int32, bool)`

GetMinimumCpusOk returns a tuple with the MinimumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpus

`func (o *V0042JobDescMsg) SetMinimumCpus(v int32)`

SetMinimumCpus sets MinimumCpus field to given value.

### HasMinimumCpus

`func (o *V0042JobDescMsg) HasMinimumCpus() bool`

HasMinimumCpus returns a boolean if a field has been set.

### GetMaximumCpus

`func (o *V0042JobDescMsg) GetMaximumCpus() int32`

GetMaximumCpus returns the MaximumCpus field if non-nil, zero value otherwise.

### GetMaximumCpusOk

`func (o *V0042JobDescMsg) GetMaximumCpusOk() (*int32, bool)`

GetMaximumCpusOk returns a tuple with the MaximumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpus

`func (o *V0042JobDescMsg) SetMaximumCpus(v int32)`

SetMaximumCpus sets MaximumCpus field to given value.

### HasMaximumCpus

`func (o *V0042JobDescMsg) HasMaximumCpus() bool`

HasMaximumCpus returns a boolean if a field has been set.

### GetNodes

`func (o *V0042JobDescMsg) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042JobDescMsg) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042JobDescMsg) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042JobDescMsg) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *V0042JobDescMsg) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *V0042JobDescMsg) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *V0042JobDescMsg) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *V0042JobDescMsg) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *V0042JobDescMsg) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *V0042JobDescMsg) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *V0042JobDescMsg) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *V0042JobDescMsg) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetMinimumBoardsPerNode

`func (o *V0042JobDescMsg) GetMinimumBoardsPerNode() int32`

GetMinimumBoardsPerNode returns the MinimumBoardsPerNode field if non-nil, zero value otherwise.

### GetMinimumBoardsPerNodeOk

`func (o *V0042JobDescMsg) GetMinimumBoardsPerNodeOk() (*int32, bool)`

GetMinimumBoardsPerNodeOk returns a tuple with the MinimumBoardsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBoardsPerNode

`func (o *V0042JobDescMsg) SetMinimumBoardsPerNode(v int32)`

SetMinimumBoardsPerNode sets MinimumBoardsPerNode field to given value.

### HasMinimumBoardsPerNode

`func (o *V0042JobDescMsg) HasMinimumBoardsPerNode() bool`

HasMinimumBoardsPerNode returns a boolean if a field has been set.

### GetMinimumSocketsPerBoard

`func (o *V0042JobDescMsg) GetMinimumSocketsPerBoard() int32`

GetMinimumSocketsPerBoard returns the MinimumSocketsPerBoard field if non-nil, zero value otherwise.

### GetMinimumSocketsPerBoardOk

`func (o *V0042JobDescMsg) GetMinimumSocketsPerBoardOk() (*int32, bool)`

GetMinimumSocketsPerBoardOk returns a tuple with the MinimumSocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSocketsPerBoard

`func (o *V0042JobDescMsg) SetMinimumSocketsPerBoard(v int32)`

SetMinimumSocketsPerBoard sets MinimumSocketsPerBoard field to given value.

### HasMinimumSocketsPerBoard

`func (o *V0042JobDescMsg) HasMinimumSocketsPerBoard() bool`

HasMinimumSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0042JobDescMsg) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0042JobDescMsg) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0042JobDescMsg) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0042JobDescMsg) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0042JobDescMsg) GetThreadsPerCore() int32`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0042JobDescMsg) GetThreadsPerCoreOk() (*int32, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0042JobDescMsg) SetThreadsPerCore(v int32)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0042JobDescMsg) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0042JobDescMsg) GetTasksPerNode() int32`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0042JobDescMsg) GetTasksPerNodeOk() (*int32, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0042JobDescMsg) SetTasksPerNode(v int32)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0042JobDescMsg) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0042JobDescMsg) GetTasksPerSocket() int32`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0042JobDescMsg) GetTasksPerSocketOk() (*int32, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0042JobDescMsg) SetTasksPerSocket(v int32)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0042JobDescMsg) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0042JobDescMsg) GetTasksPerCore() int32`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0042JobDescMsg) GetTasksPerCoreOk() (*int32, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0042JobDescMsg) SetTasksPerCore(v int32)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0042JobDescMsg) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0042JobDescMsg) GetTasksPerBoard() int32`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0042JobDescMsg) GetTasksPerBoardOk() (*int32, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0042JobDescMsg) SetTasksPerBoard(v int32)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0042JobDescMsg) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetNtasksPerTres

`func (o *V0042JobDescMsg) GetNtasksPerTres() int32`

GetNtasksPerTres returns the NtasksPerTres field if non-nil, zero value otherwise.

### GetNtasksPerTresOk

`func (o *V0042JobDescMsg) GetNtasksPerTresOk() (*int32, bool)`

GetNtasksPerTresOk returns a tuple with the NtasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtasksPerTres

`func (o *V0042JobDescMsg) SetNtasksPerTres(v int32)`

SetNtasksPerTres sets NtasksPerTres field to given value.

### HasNtasksPerTres

`func (o *V0042JobDescMsg) HasNtasksPerTres() bool`

HasNtasksPerTres returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0042JobDescMsg) GetMinimumCpusPerNode() int32`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0042JobDescMsg) GetMinimumCpusPerNodeOk() (*int32, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0042JobDescMsg) SetMinimumCpusPerNode(v int32)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0042JobDescMsg) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0042JobDescMsg) GetMemoryPerCpu() V0042Uint64NoValStruct`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0042JobDescMsg) GetMemoryPerCpuOk() (*V0042Uint64NoValStruct, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0042JobDescMsg) SetMemoryPerCpu(v V0042Uint64NoValStruct)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0042JobDescMsg) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0042JobDescMsg) GetMemoryPerNode() V0042Uint64NoValStruct`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0042JobDescMsg) GetMemoryPerNodeOk() (*V0042Uint64NoValStruct, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0042JobDescMsg) SetMemoryPerNode(v V0042Uint64NoValStruct)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0042JobDescMsg) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetTemporaryDiskPerNode

`func (o *V0042JobDescMsg) GetTemporaryDiskPerNode() int32`

GetTemporaryDiskPerNode returns the TemporaryDiskPerNode field if non-nil, zero value otherwise.

### GetTemporaryDiskPerNodeOk

`func (o *V0042JobDescMsg) GetTemporaryDiskPerNodeOk() (*int32, bool)`

GetTemporaryDiskPerNodeOk returns a tuple with the TemporaryDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDiskPerNode

`func (o *V0042JobDescMsg) SetTemporaryDiskPerNode(v int32)`

SetTemporaryDiskPerNode sets TemporaryDiskPerNode field to given value.

### HasTemporaryDiskPerNode

`func (o *V0042JobDescMsg) HasTemporaryDiskPerNode() bool`

HasTemporaryDiskPerNode returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0042JobDescMsg) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0042JobDescMsg) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0042JobDescMsg) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0042JobDescMsg) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetRequiredSwitches

`func (o *V0042JobDescMsg) GetRequiredSwitches() V0042Uint32NoValStruct`

GetRequiredSwitches returns the RequiredSwitches field if non-nil, zero value otherwise.

### GetRequiredSwitchesOk

`func (o *V0042JobDescMsg) GetRequiredSwitchesOk() (*V0042Uint32NoValStruct, bool)`

GetRequiredSwitchesOk returns a tuple with the RequiredSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSwitches

`func (o *V0042JobDescMsg) SetRequiredSwitches(v V0042Uint32NoValStruct)`

SetRequiredSwitches sets RequiredSwitches field to given value.

### HasRequiredSwitches

`func (o *V0042JobDescMsg) HasRequiredSwitches() bool`

HasRequiredSwitches returns a boolean if a field has been set.

### GetSegmentSize

`func (o *V0042JobDescMsg) GetSegmentSize() V0042Uint16NoValStruct`

GetSegmentSize returns the SegmentSize field if non-nil, zero value otherwise.

### GetSegmentSizeOk

`func (o *V0042JobDescMsg) GetSegmentSizeOk() (*V0042Uint16NoValStruct, bool)`

GetSegmentSizeOk returns a tuple with the SegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentSize

`func (o *V0042JobDescMsg) SetSegmentSize(v V0042Uint16NoValStruct)`

SetSegmentSize sets SegmentSize field to given value.

### HasSegmentSize

`func (o *V0042JobDescMsg) HasSegmentSize() bool`

HasSegmentSize returns a boolean if a field has been set.

### GetStandardError

`func (o *V0042JobDescMsg) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0042JobDescMsg) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0042JobDescMsg) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0042JobDescMsg) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0042JobDescMsg) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0042JobDescMsg) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0042JobDescMsg) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0042JobDescMsg) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0042JobDescMsg) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0042JobDescMsg) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0042JobDescMsg) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0042JobDescMsg) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetWaitForSwitch

`func (o *V0042JobDescMsg) GetWaitForSwitch() int32`

GetWaitForSwitch returns the WaitForSwitch field if non-nil, zero value otherwise.

### GetWaitForSwitchOk

`func (o *V0042JobDescMsg) GetWaitForSwitchOk() (*int32, bool)`

GetWaitForSwitchOk returns a tuple with the WaitForSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForSwitch

`func (o *V0042JobDescMsg) SetWaitForSwitch(v int32)`

SetWaitForSwitch sets WaitForSwitch field to given value.

### HasWaitForSwitch

`func (o *V0042JobDescMsg) HasWaitForSwitch() bool`

HasWaitForSwitch returns a boolean if a field has been set.

### GetWckey

`func (o *V0042JobDescMsg) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0042JobDescMsg) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0042JobDescMsg) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0042JobDescMsg) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetX11

`func (o *V0042JobDescMsg) GetX11() []string`

GetX11 returns the X11 field if non-nil, zero value otherwise.

### GetX11Ok

`func (o *V0042JobDescMsg) GetX11Ok() (*[]string, bool)`

GetX11Ok returns a tuple with the X11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11

`func (o *V0042JobDescMsg) SetX11(v []string)`

SetX11 sets X11 field to given value.

### HasX11

`func (o *V0042JobDescMsg) HasX11() bool`

HasX11 returns a boolean if a field has been set.

### GetX11MagicCookie

`func (o *V0042JobDescMsg) GetX11MagicCookie() string`

GetX11MagicCookie returns the X11MagicCookie field if non-nil, zero value otherwise.

### GetX11MagicCookieOk

`func (o *V0042JobDescMsg) GetX11MagicCookieOk() (*string, bool)`

GetX11MagicCookieOk returns a tuple with the X11MagicCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11MagicCookie

`func (o *V0042JobDescMsg) SetX11MagicCookie(v string)`

SetX11MagicCookie sets X11MagicCookie field to given value.

### HasX11MagicCookie

`func (o *V0042JobDescMsg) HasX11MagicCookie() bool`

HasX11MagicCookie returns a boolean if a field has been set.

### GetX11TargetHost

`func (o *V0042JobDescMsg) GetX11TargetHost() string`

GetX11TargetHost returns the X11TargetHost field if non-nil, zero value otherwise.

### GetX11TargetHostOk

`func (o *V0042JobDescMsg) GetX11TargetHostOk() (*string, bool)`

GetX11TargetHostOk returns a tuple with the X11TargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetHost

`func (o *V0042JobDescMsg) SetX11TargetHost(v string)`

SetX11TargetHost sets X11TargetHost field to given value.

### HasX11TargetHost

`func (o *V0042JobDescMsg) HasX11TargetHost() bool`

HasX11TargetHost returns a boolean if a field has been set.

### GetX11TargetPort

`func (o *V0042JobDescMsg) GetX11TargetPort() int32`

GetX11TargetPort returns the X11TargetPort field if non-nil, zero value otherwise.

### GetX11TargetPortOk

`func (o *V0042JobDescMsg) GetX11TargetPortOk() (*int32, bool)`

GetX11TargetPortOk returns a tuple with the X11TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetPort

`func (o *V0042JobDescMsg) SetX11TargetPort(v int32)`

SetX11TargetPort sets X11TargetPort field to given value.

### HasX11TargetPort

`func (o *V0042JobDescMsg) HasX11TargetPort() bool`

HasX11TargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


