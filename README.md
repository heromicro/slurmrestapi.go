# Go API client for slurmrestapi

API to access and control Slurm.

## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project. By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.39
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
  For more information, please visit [https://www.schedmd.com/](https://www.schedmd.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import slurmrestapi "github.com/heromicro/slurmrestapi.go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), slurmrestapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), slurmrestapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), slurmrestapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), slurmrestapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to _http://localhost_

| Class        | Method                                                                                | HTTP request                                          | Description                        |
| ------------ | ------------------------------------------------------------------------------------- | ----------------------------------------------------- | ---------------------------------- |
| _OpenapiApi_ | [**OpenapiGet**](docs/OpenapiApi.md#openapiget)                                       | **Get** /openapi                                      | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiJsonGet**](docs/OpenapiApi.md#openapijsonget)                               | **Get** /openapi.json                                 | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiV3Get**](docs/OpenapiApi.md#openapiv3get)                                   | **Get** /openapi/v3                                   | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiYamlGet**](docs/OpenapiApi.md#openapiyamlget)                               | **Get** /openapi.yaml                                 | Retrieve OpenAPI Specification     |
| _SlurmApi_   | [**SlurmV0039CancelJob**](docs/SlurmApi.md#slurmv0039canceljob)                       | **Delete** /slurm/v0.0.39/job/{job_id}                | cancel or signal job               |
| _SlurmApi_   | [**SlurmV0039DeleteNode**](docs/SlurmApi.md#slurmv0039deletenode)                     | **Delete** /slurm/v0.0.39/node/{node_name}            | delete node                        |
| _SlurmApi_   | [**SlurmV0039Diag**](docs/SlurmApi.md#slurmv0039diag)                                 | **Get** /slurm/v0.0.39/diag                           | get diagnostics                    |
| _SlurmApi_   | [**SlurmV0039GetJob**](docs/SlurmApi.md#slurmv0039getjob)                             | **Get** /slurm/v0.0.39/job/{job_id}                   | get job info                       |
| _SlurmApi_   | [**SlurmV0039GetJobs**](docs/SlurmApi.md#slurmv0039getjobs)                           | **Get** /slurm/v0.0.39/jobs                           | get list of jobs                   |
| _SlurmApi_   | [**SlurmV0039GetNode**](docs/SlurmApi.md#slurmv0039getnode)                           | **Get** /slurm/v0.0.39/node/{node_name}               | get node info                      |
| _SlurmApi_   | [**SlurmV0039GetNodes**](docs/SlurmApi.md#slurmv0039getnodes)                         | **Get** /slurm/v0.0.39/nodes                          | get all node info                  |
| _SlurmApi_   | [**SlurmV0039GetPartition**](docs/SlurmApi.md#slurmv0039getpartition)                 | **Get** /slurm/v0.0.39/partition/{partition_name}     | get partition info                 |
| _SlurmApi_   | [**SlurmV0039GetPartitions**](docs/SlurmApi.md#slurmv0039getpartitions)               | **Get** /slurm/v0.0.39/partitions                     | get all partition info             |
| _SlurmApi_   | [**SlurmV0039GetReservation**](docs/SlurmApi.md#slurmv0039getreservation)             | **Get** /slurm/v0.0.39/reservation/{reservation_name} | get reservation info               |
| _SlurmApi_   | [**SlurmV0039GetReservations**](docs/SlurmApi.md#slurmv0039getreservations)           | **Get** /slurm/v0.0.39/reservations                   | get all reservation info           |
| _SlurmApi_   | [**SlurmV0039Ping**](docs/SlurmApi.md#slurmv0039ping)                                 | **Get** /slurm/v0.0.39/ping                           | ping test                          |
| _SlurmApi_   | [**SlurmV0039SlurmctldGetLicenses**](docs/SlurmApi.md#slurmv0039slurmctldgetlicenses) | **Get** /slurm/v0.0.39/licenses                       | get all Slurm tracked license info |
| _SlurmApi_   | [**SlurmV0039SubmitJob**](docs/SlurmApi.md#slurmv0039submitjob)                       | **Post** /slurm/v0.0.39/job/submit                    | submit new job                     |
| _SlurmApi_   | [**SlurmV0039UpdateJob**](docs/SlurmApi.md#slurmv0039updatejob)                       | **Post** /slurm/v0.0.39/job/{job_id}                  | update job                         |
| _SlurmApi_   | [**SlurmV0039UpdateNode**](docs/SlurmApi.md#slurmv0039updatenode)                     | **Post** /slurm/v0.0.39/node/{node_name}              | update node properties             |
| _SlurmApi_   | [**SlurmdbV0039AddClusters**](docs/SlurmApi.md#slurmdbv0039addclusters)               | **Post** /slurmdb/v0.0.39/clusters                    | Add clusters                       |
| _SlurmApi_   | [**SlurmdbV0039AddWckeys**](docs/SlurmApi.md#slurmdbv0039addwckeys)                   | **Post** /slurmdb/v0.0.39/wckeys                      | Add wckeys                         |
| _SlurmApi_   | [**SlurmdbV0039DeleteAccount**](docs/SlurmApi.md#slurmdbv0039deleteaccount)           | **Delete** /slurmdb/v0.0.39/account/{account_name}    | Delete account                     |
| _SlurmApi_   | [**SlurmdbV0039DeleteAssociation**](docs/SlurmApi.md#slurmdbv0039deleteassociation)   | **Delete** /slurmdb/v0.0.39/association               | Delete association                 |
| _SlurmApi_   | [**SlurmdbV0039DeleteAssociations**](docs/SlurmApi.md#slurmdbv0039deleteassociations) | **Delete** /slurmdb/v0.0.39/associations              | Delete associations                |
| _SlurmApi_   | [**SlurmdbV0039DeleteCluster**](docs/SlurmApi.md#slurmdbv0039deletecluster)           | **Delete** /slurmdb/v0.0.39/cluster/{cluster_name}    | Delete cluster                     |
| _SlurmApi_   | [**SlurmdbV0039DeleteQos**](docs/SlurmApi.md#slurmdbv0039deleteqos)                   | **Delete** /slurmdb/v0.0.39/qos/{qos_name}            | Delete QOS                         |
| _SlurmApi_   | [**SlurmdbV0039DeleteUser**](docs/SlurmApi.md#slurmdbv0039deleteuser)                 | **Delete** /slurmdb/v0.0.39/user/{user_name}          | Delete user                        |
| _SlurmApi_   | [**SlurmdbV0039DeleteWckey**](docs/SlurmApi.md#slurmdbv0039deletewckey)               | **Delete** /slurmdb/v0.0.39/wckey/{wckey}             | Delete wckey                       |
| _SlurmApi_   | [**SlurmdbV0039Diag**](docs/SlurmApi.md#slurmdbv0039diag)                             | **Get** /slurmdb/v0.0.39/diag                         | Get slurmdb diagnostics            |
| _SlurmApi_   | [**SlurmdbV0039GetAccount**](docs/SlurmApi.md#slurmdbv0039getaccount)                 | **Get** /slurmdb/v0.0.39/account/{account_name}       | Get account info                   |
| _SlurmApi_   | [**SlurmdbV0039GetAccounts**](docs/SlurmApi.md#slurmdbv0039getaccounts)               | **Get** /slurmdb/v0.0.39/accounts                     | Get account list                   |
| _SlurmApi_   | [**SlurmdbV0039GetAssociation**](docs/SlurmApi.md#slurmdbv0039getassociation)         | **Get** /slurmdb/v0.0.39/association                  | Get association info               |
| _SlurmApi_   | [**SlurmdbV0039GetAssociations**](docs/SlurmApi.md#slurmdbv0039getassociations)       | **Get** /slurmdb/v0.0.39/associations                 | Get association list               |
| _SlurmApi_   | [**SlurmdbV0039GetCluster**](docs/SlurmApi.md#slurmdbv0039getcluster)                 | **Get** /slurmdb/v0.0.39/cluster/{cluster_name}       | Get cluster info                   |
| _SlurmApi_   | [**SlurmdbV0039GetClusters**](docs/SlurmApi.md#slurmdbv0039getclusters)               | **Get** /slurmdb/v0.0.39/clusters                     | Get cluster list                   |
| _SlurmApi_   | [**SlurmdbV0039GetConfig**](docs/SlurmApi.md#slurmdbv0039getconfig)                   | **Get** /slurmdb/v0.0.39/config                       | Dump all configuration information |
| _SlurmApi_   | [**SlurmdbV0039GetJob**](docs/SlurmApi.md#slurmdbv0039getjob)                         | **Get** /slurmdb/v0.0.39/job/{job_id}                 | Get job info                       |
| _SlurmApi_   | [**SlurmdbV0039GetJobs**](docs/SlurmApi.md#slurmdbv0039getjobs)                       | **Get** /slurmdb/v0.0.39/jobs                         | Get job list                       |
| _SlurmApi_   | [**SlurmdbV0039GetQos**](docs/SlurmApi.md#slurmdbv0039getqos)                         | **Get** /slurmdb/v0.0.39/qos                          | Get QOS list                       |
| _SlurmApi_   | [**SlurmdbV0039GetSingleQos**](docs/SlurmApi.md#slurmdbv0039getsingleqos)             | **Get** /slurmdb/v0.0.39/qos/{qos_name}               | Get QOS info                       |
| _SlurmApi_   | [**SlurmdbV0039GetTres**](docs/SlurmApi.md#slurmdbv0039gettres)                       | **Get** /slurmdb/v0.0.39/tres                         | Get TRES info                      |
| _SlurmApi_   | [**SlurmdbV0039GetUser**](docs/SlurmApi.md#slurmdbv0039getuser)                       | **Get** /slurmdb/v0.0.39/user/{user_name}             | Get user info                      |
| _SlurmApi_   | [**SlurmdbV0039GetUsers**](docs/SlurmApi.md#slurmdbv0039getusers)                     | **Get** /slurmdb/v0.0.39/users                        | Get user list                      |
| _SlurmApi_   | [**SlurmdbV0039GetWckey**](docs/SlurmApi.md#slurmdbv0039getwckey)                     | **Get** /slurmdb/v0.0.39/wckey/{wckey}                | Get wckey info                     |
| _SlurmApi_   | [**SlurmdbV0039GetWckeys**](docs/SlurmApi.md#slurmdbv0039getwckeys)                   | **Get** /slurmdb/v0.0.39/wckeys                       | Get wckey list                     |
| _SlurmApi_   | [**SlurmdbV0039SetConfig**](docs/SlurmApi.md#slurmdbv0039setconfig)                   | **Post** /slurmdb/v0.0.39/config                      | Load all configuration information |
| _SlurmApi_   | [**SlurmdbV0039UpdateAccounts**](docs/SlurmApi.md#slurmdbv0039updateaccounts)         | **Post** /slurmdb/v0.0.39/accounts                    | Update accounts                    |
| _SlurmApi_   | [**SlurmdbV0039UpdateAssociations**](docs/SlurmApi.md#slurmdbv0039updateassociations) | **Post** /slurmdb/v0.0.39/associations                | Set associations info              |
| _SlurmApi_   | [**SlurmdbV0039UpdateQos**](docs/SlurmApi.md#slurmdbv0039updateqos)                   | **Post** /slurmdb/v0.0.39/qos                         | Set QOS info                       |
| _SlurmApi_   | [**SlurmdbV0039UpdateTres**](docs/SlurmApi.md#slurmdbv0039updatetres)                 | **Post** /slurmdb/v0.0.39/tres                        | Set TRES info                      |
| _SlurmApi_   | [**SlurmdbV0039UpdateUsers**](docs/SlurmApi.md#slurmdbv0039updateusers)               | **Post** /slurmdb/v0.0.39/users                       | Update user                        |

## Documentation For Models

- [Dbv0039AccountInfo](docs/Dbv0039AccountInfo.md)
- [Dbv0039AssociationsInfo](docs/Dbv0039AssociationsInfo.md)
- [Dbv0039ClustersInfo](docs/Dbv0039ClustersInfo.md)
- [Dbv0039ConfigInfo](docs/Dbv0039ConfigInfo.md)
- [Dbv0039Diag](docs/Dbv0039Diag.md)
- [Dbv0039Error](docs/Dbv0039Error.md)
- [Dbv0039JobInfo](docs/Dbv0039JobInfo.md)
- [Dbv0039Meta](docs/Dbv0039Meta.md)
- [Dbv0039QosInfo](docs/Dbv0039QosInfo.md)
- [Dbv0039ResponseAssociationsDelete](docs/Dbv0039ResponseAssociationsDelete.md)
- [Dbv0039SetConfig](docs/Dbv0039SetConfig.md)
- [Dbv0039TresInfo](docs/Dbv0039TresInfo.md)
- [Dbv0039TresUpdate](docs/Dbv0039TresUpdate.md)
- [Dbv0039UpdateQos](docs/Dbv0039UpdateQos.md)
- [Dbv0039UpdateUsers](docs/Dbv0039UpdateUsers.md)
- [Dbv0039UserInfo](docs/Dbv0039UserInfo.md)
- [Dbv0039Warning](docs/Dbv0039Warning.md)
- [Dbv0039WckeyInfo](docs/Dbv0039WckeyInfo.md)
- [Status](docs/Status.md)
- [V0039Account](docs/V0039Account.md)
- [V0039Accounting](docs/V0039Accounting.md)
- [V0039AccountingAllocated](docs/V0039AccountingAllocated.md)
- [V0039AcctGatherEnergy](docs/V0039AcctGatherEnergy.md)
- [V0039Assoc](docs/V0039Assoc.md)
- [V0039AssocDefault](docs/V0039AssocDefault.md)
- [V0039AssocMax](docs/V0039AssocMax.md)
- [V0039AssocMaxJobs](docs/V0039AssocMaxJobs.md)
- [V0039AssocMaxJobsPer](docs/V0039AssocMaxJobsPer.md)
- [V0039AssocMin](docs/V0039AssocMin.md)
- [V0039AssocShort](docs/V0039AssocShort.md)
- [V0039AssocUsage](docs/V0039AssocUsage.md)
- [V0039ClusterRec](docs/V0039ClusterRec.md)
- [V0039ClusterRecAssociations](docs/V0039ClusterRecAssociations.md)
- [V0039ClusterRecController](docs/V0039ClusterRecController.md)
- [V0039ControllerPing](docs/V0039ControllerPing.md)
- [V0039Coord](docs/V0039Coord.md)
- [V0039CronEntry](docs/V0039CronEntry.md)
- [V0039CronEntryLine](docs/V0039CronEntryLine.md)
- [V0039Diag](docs/V0039Diag.md)
- [V0039Error](docs/V0039Error.md)
- [V0039ExtSensorsData](docs/V0039ExtSensorsData.md)
- [V0039Float64NoVal](docs/V0039Float64NoVal.md)
- [V0039Job](docs/V0039Job.md)
- [V0039JobArray](docs/V0039JobArray.md)
- [V0039JobArrayResponseMsgInner](docs/V0039JobArrayResponseMsgInner.md)
- [V0039JobComment](docs/V0039JobComment.md)
- [V0039JobDescMsg](docs/V0039JobDescMsg.md)
- [V0039JobExitCode](docs/V0039JobExitCode.md)
- [V0039JobExitCodeSignal](docs/V0039JobExitCodeSignal.md)
- [V0039JobHet](docs/V0039JobHet.md)
- [V0039JobInfo](docs/V0039JobInfo.md)
- [V0039JobInfoPower](docs/V0039JobInfoPower.md)
- [V0039JobMcs](docs/V0039JobMcs.md)
- [V0039JobRequired](docs/V0039JobRequired.md)
- [V0039JobRes](docs/V0039JobRes.md)
- [V0039JobReservation](docs/V0039JobReservation.md)
- [V0039JobState](docs/V0039JobState.md)
- [V0039JobSubmission](docs/V0039JobSubmission.md)
- [V0039JobSubmissionResponse](docs/V0039JobSubmissionResponse.md)
- [V0039JobTime](docs/V0039JobTime.md)
- [V0039JobTimeUser](docs/V0039JobTimeUser.md)
- [V0039JobTres](docs/V0039JobTres.md)
- [V0039JobUpdateResponse](docs/V0039JobUpdateResponse.md)
- [V0039JobsResponse](docs/V0039JobsResponse.md)
- [V0039License](docs/V0039License.md)
- [V0039LicensesInfo](docs/V0039LicensesInfo.md)
- [V0039Meta](docs/V0039Meta.md)
- [V0039MetaPlugin](docs/V0039MetaPlugin.md)
- [V0039MetaSlurm](docs/V0039MetaSlurm.md)
- [V0039MetaSlurmVersion](docs/V0039MetaSlurmVersion.md)
- [V0039Node](docs/V0039Node.md)
- [V0039NodesResponse](docs/V0039NodesResponse.md)
- [V0039PartitionInfo](docs/V0039PartitionInfo.md)
- [V0039PartitionInfoAccounts](docs/V0039PartitionInfoAccounts.md)
- [V0039PartitionInfoDefaults](docs/V0039PartitionInfoDefaults.md)
- [V0039PartitionInfoGroups](docs/V0039PartitionInfoGroups.md)
- [V0039PartitionInfoMaximums](docs/V0039PartitionInfoMaximums.md)
- [V0039PartitionInfoMinimums](docs/V0039PartitionInfoMinimums.md)
- [V0039PartitionInfoNodes](docs/V0039PartitionInfoNodes.md)
- [V0039PartitionInfoPriority](docs/V0039PartitionInfoPriority.md)
- [V0039PartitionInfoQos](docs/V0039PartitionInfoQos.md)
- [V0039PartitionInfoTimeouts](docs/V0039PartitionInfoTimeouts.md)
- [V0039PartitionInfoTres](docs/V0039PartitionInfoTres.md)
- [V0039PartitionsResponse](docs/V0039PartitionsResponse.md)
- [V0039Pings](docs/V0039Pings.md)
- [V0039PowerMgmtData](docs/V0039PowerMgmtData.md)
- [V0039Qos](docs/V0039Qos.md)
- [V0039QosLimits](docs/V0039QosLimits.md)
- [V0039QosLimitsMin](docs/V0039QosLimitsMin.md)
- [V0039QosLimitsMinTres](docs/V0039QosLimitsMinTres.md)
- [V0039QosLimitsMinTresPer](docs/V0039QosLimitsMinTresPer.md)
- [V0039QosPreempt](docs/V0039QosPreempt.md)
- [V0039ReservationCoreSpec](docs/V0039ReservationCoreSpec.md)
- [V0039ReservationInfo](docs/V0039ReservationInfo.md)
- [V0039ReservationInfoPurgeCompleted](docs/V0039ReservationInfoPurgeCompleted.md)
- [V0039ReservationsResponse](docs/V0039ReservationsResponse.md)
- [V0039RollupStatsInner](docs/V0039RollupStatsInner.md)
- [V0039SlurmStepId](docs/V0039SlurmStepId.md)
- [V0039StatsMsg](docs/V0039StatsMsg.md)
- [V0039StatsMsgRpcsByTypeInner](docs/V0039StatsMsgRpcsByTypeInner.md)
- [V0039StatsMsgRpcsByUserInner](docs/V0039StatsMsgRpcsByUserInner.md)
- [V0039StatsRec](docs/V0039StatsRec.md)
- [V0039StatsRpc](docs/V0039StatsRpc.md)
- [V0039StatsRpcTime](docs/V0039StatsRpcTime.md)
- [V0039StatsUser](docs/V0039StatsUser.md)
- [V0039Step](docs/V0039Step.md)
- [V0039StepCPU](docs/V0039StepCPU.md)
- [V0039StepNodes](docs/V0039StepNodes.md)
- [V0039StepStatistics](docs/V0039StepStatistics.md)
- [V0039StepStatisticsEnergy](docs/V0039StepStatisticsEnergy.md)
- [V0039StepTask](docs/V0039StepTask.md)
- [V0039StepTasks](docs/V0039StepTasks.md)
- [V0039StepTime](docs/V0039StepTime.md)
- [V0039StepTimeUser](docs/V0039StepTimeUser.md)
- [V0039StepTres](docs/V0039StepTres.md)
- [V0039Tres](docs/V0039Tres.md)
- [V0039Uint16NoVal](docs/V0039Uint16NoVal.md)
- [V0039Uint32NoVal](docs/V0039Uint32NoVal.md)
- [V0039Uint64NoVal](docs/V0039Uint64NoVal.md)
- [V0039UpdateNodeMsg](docs/V0039UpdateNodeMsg.md)
- [V0039User](docs/V0039User.md)
- [V0039UserDefault](docs/V0039UserDefault.md)
- [V0039Warning](docs/V0039Warning.md)
- [V0039Wckey](docs/V0039Wckey.md)
- [V0039WckeyTag](docs/V0039WckeyTag.md)

## Documentation For Authorization

Authentication schemes defined for the API:

### user

- **Type**: API key
- **API key parameter name**: X-SLURM-USER-NAME
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-SLURM-USER-NAME and passed in as the auth context for each request.

### token

- **Type**: API key
- **API key parameter name**: X-SLURM-USER-TOKEN
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-SLURM-USER-TOKEN and passed in as the auth context for each request.

### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

- `PtrBool`
- `PtrInt`
- `PtrInt32`
- `PtrInt64`
- `PtrFloat`
- `PtrFloat32`
- `PtrFloat64`
- `PtrString`
- `PtrTime`

## Author

sales@schedmd.com
