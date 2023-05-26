# Go API client for slurmrestapi

API to access and control Slurm.

## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project. By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.37
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

| Class        | Method                                                                        | HTTP request                                          | Description                        |
| ------------ | ----------------------------------------------------------------------------- | ----------------------------------------------------- | ---------------------------------- |
| _OpenapiApi_ | [**OpenapiGet**](docs/OpenapiApi.md#openapiget)                               | **Get** /openapi                                      | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiJsonGet**](docs/OpenapiApi.md#openapijsonget)                       | **Get** /openapi.json                                 | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiV3Get**](docs/OpenapiApi.md#openapiv3get)                           | **Get** /openapi/v3                                   | Retrieve OpenAPI Specification     |
| _OpenapiApi_ | [**OpenapiYamlGet**](docs/OpenapiApi.md#openapiyamlget)                       | **Get** /openapi.yaml                                 | Retrieve OpenAPI Specification     |
| _SlurmApi_   | [**SlurmctldCancelJob**](docs/SlurmApi.md#slurmctldcanceljob)                 | **Delete** /slurm/v0.0.37/job/{job_id}                | cancel or signal job               |
| _SlurmApi_   | [**SlurmctldDiag**](docs/SlurmApi.md#slurmctlddiag)                           | **Get** /slurm/v0.0.37/diag                           | get diagnostics                    |
| _SlurmApi_   | [**SlurmctldGetJob**](docs/SlurmApi.md#slurmctldgetjob)                       | **Get** /slurm/v0.0.37/job/{job_id}                   | get job info                       |
| _SlurmApi_   | [**SlurmctldGetJobs**](docs/SlurmApi.md#slurmctldgetjobs)                     | **Get** /slurm/v0.0.37/jobs                           | get list of jobs                   |
| _SlurmApi_   | [**SlurmctldGetNode**](docs/SlurmApi.md#slurmctldgetnode)                     | **Get** /slurm/v0.0.37/node/{node_name}               | get node info                      |
| _SlurmApi_   | [**SlurmctldGetNodes**](docs/SlurmApi.md#slurmctldgetnodes)                   | **Get** /slurm/v0.0.37/nodes                          | get all node info                  |
| _SlurmApi_   | [**SlurmctldGetPartition**](docs/SlurmApi.md#slurmctldgetpartition)           | **Get** /slurm/v0.0.37/partition/{partition_name}     | get partition info                 |
| _SlurmApi_   | [**SlurmctldGetPartitions**](docs/SlurmApi.md#slurmctldgetpartitions)         | **Get** /slurm/v0.0.37/partitions                     | get all partition info             |
| _SlurmApi_   | [**SlurmctldGetReservation**](docs/SlurmApi.md#slurmctldgetreservation)       | **Get** /slurm/v0.0.37/reservation/{reservation_name} | get reservation info               |
| _SlurmApi_   | [**SlurmctldGetReservations**](docs/SlurmApi.md#slurmctldgetreservations)     | **Get** /slurm/v0.0.37/reservations                   | get all reservation info           |
| _SlurmApi_   | [**SlurmctldPing**](docs/SlurmApi.md#slurmctldping)                           | **Get** /slurm/v0.0.37/ping                           | ping test                          |
| _SlurmApi_   | [**SlurmctldSubmitJob**](docs/SlurmApi.md#slurmctldsubmitjob)                 | **Post** /slurm/v0.0.37/job/submit                    | submit new job                     |
| _SlurmApi_   | [**SlurmctldUpdateJob**](docs/SlurmApi.md#slurmctldupdatejob)                 | **Post** /slurm/v0.0.37/job/{job_id}                  | update job                         |
| _SlurmApi_   | [**SlurmdbdAddClusters**](docs/SlurmApi.md#slurmdbdaddclusters)               | **Post** /slurmdb/v0.0.37/clusters                    | Add clusters                       |
| _SlurmApi_   | [**SlurmdbdAddWckeys**](docs/SlurmApi.md#slurmdbdaddwckeys)                   | **Post** /slurmdb/v0.0.37/wckeys                      | Add wckeys                         |
| _SlurmApi_   | [**SlurmdbdDeleteAccount**](docs/SlurmApi.md#slurmdbddeleteaccount)           | **Delete** /slurmdb/v0.0.37/account/{account_name}    | Delete account                     |
| _SlurmApi_   | [**SlurmdbdDeleteAssociation**](docs/SlurmApi.md#slurmdbddeleteassociation)   | **Delete** /slurmdb/v0.0.37/association               | Delete association                 |
| _SlurmApi_   | [**SlurmdbdDeleteCluster**](docs/SlurmApi.md#slurmdbddeletecluster)           | **Delete** /slurmdb/v0.0.37/cluster/{cluster_name}    | Delete cluster                     |
| _SlurmApi_   | [**SlurmdbdDeleteQos**](docs/SlurmApi.md#slurmdbddeleteqos)                   | **Delete** /slurmdb/v0.0.37/qos/{qos_name}            | Delete QOS                         |
| _SlurmApi_   | [**SlurmdbdDeleteUser**](docs/SlurmApi.md#slurmdbddeleteuser)                 | **Delete** /slurmdb/v0.0.37/user/{user_name}          | Delete user                        |
| _SlurmApi_   | [**SlurmdbdDeleteWckey**](docs/SlurmApi.md#slurmdbddeletewckey)               | **Delete** /slurmdb/v0.0.37/wckey/{wckey}             | Delete wckey                       |
| _SlurmApi_   | [**SlurmdbdDiag**](docs/SlurmApi.md#slurmdbddiag)                             | **Get** /slurmdb/v0.0.37/diag                         | Get slurmdb diagnostics            |
| _SlurmApi_   | [**SlurmdbdGetAccount**](docs/SlurmApi.md#slurmdbdgetaccount)                 | **Get** /slurmdb/v0.0.37/account/{account_name}       | Get account info                   |
| _SlurmApi_   | [**SlurmdbdGetAccounts**](docs/SlurmApi.md#slurmdbdgetaccounts)               | **Get** /slurmdb/v0.0.37/accounts                     | Get account list                   |
| _SlurmApi_   | [**SlurmdbdGetAssociation**](docs/SlurmApi.md#slurmdbdgetassociation)         | **Get** /slurmdb/v0.0.37/association                  | Get association info               |
| _SlurmApi_   | [**SlurmdbdGetAssociations**](docs/SlurmApi.md#slurmdbdgetassociations)       | **Get** /slurmdb/v0.0.37/associations                 | Get association list               |
| _SlurmApi_   | [**SlurmdbdGetCluster**](docs/SlurmApi.md#slurmdbdgetcluster)                 | **Get** /slurmdb/v0.0.37/cluster/{cluster_name}       | Get cluster info                   |
| _SlurmApi_   | [**SlurmdbdGetClusters**](docs/SlurmApi.md#slurmdbdgetclusters)               | **Get** /slurmdb/v0.0.37/clusters                     | Get cluster list                   |
| _SlurmApi_   | [**SlurmdbdGetDbConfig**](docs/SlurmApi.md#slurmdbdgetdbconfig)               | **Get** /slurmdb/v0.0.37/config                       | Dump all configuration information |
| _SlurmApi_   | [**SlurmdbdGetJob**](docs/SlurmApi.md#slurmdbdgetjob)                         | **Get** /slurmdb/v0.0.37/job/{job_id}                 | Get job info                       |
| _SlurmApi_   | [**SlurmdbdGetJobs**](docs/SlurmApi.md#slurmdbdgetjobs)                       | **Get** /slurmdb/v0.0.37/jobs                         | Get job list                       |
| _SlurmApi_   | [**SlurmdbdGetQos**](docs/SlurmApi.md#slurmdbdgetqos)                         | **Get** /slurmdb/v0.0.37/qos                          | Get QOS list                       |
| _SlurmApi_   | [**SlurmdbdGetSingleQos**](docs/SlurmApi.md#slurmdbdgetsingleqos)             | **Get** /slurmdb/v0.0.37/qos/{qos_name}               | Get QOS info                       |
| _SlurmApi_   | [**SlurmdbdGetTres**](docs/SlurmApi.md#slurmdbdgettres)                       | **Get** /slurmdb/v0.0.37/tres                         | Get TRES info                      |
| _SlurmApi_   | [**SlurmdbdGetUser**](docs/SlurmApi.md#slurmdbdgetuser)                       | **Get** /slurmdb/v0.0.37/user/{user_name}             | Get user info                      |
| _SlurmApi_   | [**SlurmdbdGetUsers**](docs/SlurmApi.md#slurmdbdgetusers)                     | **Get** /slurmdb/v0.0.37/users                        | Get user list                      |
| _SlurmApi_   | [**SlurmdbdGetWckey**](docs/SlurmApi.md#slurmdbdgetwckey)                     | **Get** /slurmdb/v0.0.37/wckey/{wckey}                | Get wckey info                     |
| _SlurmApi_   | [**SlurmdbdGetWckeys**](docs/SlurmApi.md#slurmdbdgetwckeys)                   | **Get** /slurmdb/v0.0.37/wckeys                       | Get wckey list                     |
| _SlurmApi_   | [**SlurmdbdSetDbConfig**](docs/SlurmApi.md#slurmdbdsetdbconfig)               | **Post** /slurmdb/v0.0.37/config                      | Load all configuration information |
| _SlurmApi_   | [**SlurmdbdUpdateAccount**](docs/SlurmApi.md#slurmdbdupdateaccount)           | **Post** /slurmdb/v0.0.37/accounts                    | Update accounts                    |
| _SlurmApi_   | [**SlurmdbdUpdateAssociations**](docs/SlurmApi.md#slurmdbdupdateassociations) | **Post** /slurmdb/v0.0.37/associations                | Set associations info              |
| _SlurmApi_   | [**SlurmdbdUpdateTres**](docs/SlurmApi.md#slurmdbdupdatetres)                 | **Post** /slurmdb/v0.0.37/tres                        | Set TRES info                      |
| _SlurmApi_   | [**SlurmdbdUpdateUsers**](docs/SlurmApi.md#slurmdbdupdateusers)               | **Post** /slurmdb/v0.0.37/users                       | Update user                        |

## Documentation For Models

- [Dbv0037Account](docs/Dbv0037Account.md)
- [Dbv0037AccountInfo](docs/Dbv0037AccountInfo.md)
- [Dbv0037AccountResponse](docs/Dbv0037AccountResponse.md)
- [Dbv0037Association](docs/Dbv0037Association.md)
- [Dbv0037AssociationDefault](docs/Dbv0037AssociationDefault.md)
- [Dbv0037AssociationMax](docs/Dbv0037AssociationMax.md)
- [Dbv0037AssociationMaxJobs](docs/Dbv0037AssociationMaxJobs.md)
- [Dbv0037AssociationMaxJobsPer](docs/Dbv0037AssociationMaxJobsPer.md)
- [Dbv0037AssociationMaxPer](docs/Dbv0037AssociationMaxPer.md)
- [Dbv0037AssociationMaxPerAccount](docs/Dbv0037AssociationMaxPerAccount.md)
- [Dbv0037AssociationMaxTres](docs/Dbv0037AssociationMaxTres.md)
- [Dbv0037AssociationMaxTresGroup](docs/Dbv0037AssociationMaxTresGroup.md)
- [Dbv0037AssociationMaxTresMinutes](docs/Dbv0037AssociationMaxTresMinutes.md)
- [Dbv0037AssociationMaxTresMinutesPer](docs/Dbv0037AssociationMaxTresMinutesPer.md)
- [Dbv0037AssociationMaxTresPer](docs/Dbv0037AssociationMaxTresPer.md)
- [Dbv0037AssociationMin](docs/Dbv0037AssociationMin.md)
- [Dbv0037AssociationShortInfo](docs/Dbv0037AssociationShortInfo.md)
- [Dbv0037AssociationUsage](docs/Dbv0037AssociationUsage.md)
- [Dbv0037AssociationsInfo](docs/Dbv0037AssociationsInfo.md)
- [Dbv0037ClusterInfo](docs/Dbv0037ClusterInfo.md)
- [Dbv0037ClusterInfoAssociations](docs/Dbv0037ClusterInfoAssociations.md)
- [Dbv0037ClusterInfoController](docs/Dbv0037ClusterInfoController.md)
- [Dbv0037ConfigInfo](docs/Dbv0037ConfigInfo.md)
- [Dbv0037ConfigResponse](docs/Dbv0037ConfigResponse.md)
- [Dbv0037CoordinatorInfo](docs/Dbv0037CoordinatorInfo.md)
- [Dbv0037Diag](docs/Dbv0037Diag.md)
- [Dbv0037DiagStatistics](docs/Dbv0037DiagStatistics.md)
- [Dbv0037DiagStatisticsRPCsInner](docs/Dbv0037DiagStatisticsRPCsInner.md)
- [Dbv0037DiagStatisticsRPCsInnerTime](docs/Dbv0037DiagStatisticsRPCsInnerTime.md)
- [Dbv0037DiagStatisticsRollupsInner](docs/Dbv0037DiagStatisticsRollupsInner.md)
- [Dbv0037DiagStatisticsUsersInner](docs/Dbv0037DiagStatisticsUsersInner.md)
- [Dbv0037DiagStatisticsUsersInnerTime](docs/Dbv0037DiagStatisticsUsersInnerTime.md)
- [Dbv0037Error](docs/Dbv0037Error.md)
- [Dbv0037Job](docs/Dbv0037Job.md)
- [Dbv0037JobArray](docs/Dbv0037JobArray.md)
- [Dbv0037JobArrayLimits](docs/Dbv0037JobArrayLimits.md)
- [Dbv0037JobArrayLimitsMax](docs/Dbv0037JobArrayLimitsMax.md)
- [Dbv0037JobArrayLimitsMaxRunning](docs/Dbv0037JobArrayLimitsMaxRunning.md)
- [Dbv0037JobComment](docs/Dbv0037JobComment.md)
- [Dbv0037JobExitCode](docs/Dbv0037JobExitCode.md)
- [Dbv0037JobExitCodeSignal](docs/Dbv0037JobExitCodeSignal.md)
- [Dbv0037JobHet](docs/Dbv0037JobHet.md)
- [Dbv0037JobInfo](docs/Dbv0037JobInfo.md)
- [Dbv0037JobMcs](docs/Dbv0037JobMcs.md)
- [Dbv0037JobRequired](docs/Dbv0037JobRequired.md)
- [Dbv0037JobReservation](docs/Dbv0037JobReservation.md)
- [Dbv0037JobState](docs/Dbv0037JobState.md)
- [Dbv0037JobStep](docs/Dbv0037JobStep.md)
- [Dbv0037JobStepCPU](docs/Dbv0037JobStepCPU.md)
- [Dbv0037JobStepCPURequestedFrequency](docs/Dbv0037JobStepCPURequestedFrequency.md)
- [Dbv0037JobStepNodes](docs/Dbv0037JobStepNodes.md)
- [Dbv0037JobStepStatistics](docs/Dbv0037JobStepStatistics.md)
- [Dbv0037JobStepStatisticsCPU](docs/Dbv0037JobStepStatisticsCPU.md)
- [Dbv0037JobStepStatisticsEnergy](docs/Dbv0037JobStepStatisticsEnergy.md)
- [Dbv0037JobStepStep](docs/Dbv0037JobStepStep.md)
- [Dbv0037JobStepStepHet](docs/Dbv0037JobStepStepHet.md)
- [Dbv0037JobStepTask](docs/Dbv0037JobStepTask.md)
- [Dbv0037JobStepTasks](docs/Dbv0037JobStepTasks.md)
- [Dbv0037JobStepTime](docs/Dbv0037JobStepTime.md)
- [Dbv0037JobStepTres](docs/Dbv0037JobStepTres.md)
- [Dbv0037JobStepTresRequested](docs/Dbv0037JobStepTresRequested.md)
- [Dbv0037JobTime](docs/Dbv0037JobTime.md)
- [Dbv0037JobTimeSystem](docs/Dbv0037JobTimeSystem.md)
- [Dbv0037JobTimeTotal](docs/Dbv0037JobTimeTotal.md)
- [Dbv0037JobTimeUser](docs/Dbv0037JobTimeUser.md)
- [Dbv0037JobTres](docs/Dbv0037JobTres.md)
- [Dbv0037JobWckey](docs/Dbv0037JobWckey.md)
- [Dbv0037Qos](docs/Dbv0037Qos.md)
- [Dbv0037QosInfo](docs/Dbv0037QosInfo.md)
- [Dbv0037QosLimits](docs/Dbv0037QosLimits.md)
- [Dbv0037QosLimitsMax](docs/Dbv0037QosLimitsMax.md)
- [Dbv0037QosLimitsMaxAccruing](docs/Dbv0037QosLimitsMaxAccruing.md)
- [Dbv0037QosLimitsMaxAccruingPer](docs/Dbv0037QosLimitsMaxAccruingPer.md)
- [Dbv0037QosLimitsMaxJobs](docs/Dbv0037QosLimitsMaxJobs.md)
- [Dbv0037QosLimitsMaxJobsActiveJobs](docs/Dbv0037QosLimitsMaxJobsActiveJobs.md)
- [Dbv0037QosLimitsMaxJobsActiveJobsPer](docs/Dbv0037QosLimitsMaxJobsActiveJobsPer.md)
- [Dbv0037QosLimitsMaxTres](docs/Dbv0037QosLimitsMaxTres.md)
- [Dbv0037QosLimitsMaxTresMinutes](docs/Dbv0037QosLimitsMaxTresMinutes.md)
- [Dbv0037QosLimitsMaxTresMinutesPer](docs/Dbv0037QosLimitsMaxTresMinutesPer.md)
- [Dbv0037QosLimitsMaxTresPer](docs/Dbv0037QosLimitsMaxTresPer.md)
- [Dbv0037QosLimitsMaxWallClock](docs/Dbv0037QosLimitsMaxWallClock.md)
- [Dbv0037QosLimitsMaxWallClockPer](docs/Dbv0037QosLimitsMaxWallClockPer.md)
- [Dbv0037QosLimitsMin](docs/Dbv0037QosLimitsMin.md)
- [Dbv0037QosLimitsMinTres](docs/Dbv0037QosLimitsMinTres.md)
- [Dbv0037QosLimitsMinTresPer](docs/Dbv0037QosLimitsMinTresPer.md)
- [Dbv0037QosPreempt](docs/Dbv0037QosPreempt.md)
- [Dbv0037ResponseAccountDelete](docs/Dbv0037ResponseAccountDelete.md)
- [Dbv0037ResponseAssociationDelete](docs/Dbv0037ResponseAssociationDelete.md)
- [Dbv0037ResponseAssociations](docs/Dbv0037ResponseAssociations.md)
- [Dbv0037ResponseClusterAdd](docs/Dbv0037ResponseClusterAdd.md)
- [Dbv0037ResponseClusterDelete](docs/Dbv0037ResponseClusterDelete.md)
- [Dbv0037ResponseQosDelete](docs/Dbv0037ResponseQosDelete.md)
- [Dbv0037ResponseTres](docs/Dbv0037ResponseTres.md)
- [Dbv0037ResponseUserDelete](docs/Dbv0037ResponseUserDelete.md)
- [Dbv0037ResponseUserUpdate](docs/Dbv0037ResponseUserUpdate.md)
- [Dbv0037ResponseWckeyAdd](docs/Dbv0037ResponseWckeyAdd.md)
- [Dbv0037ResponseWckeyDelete](docs/Dbv0037ResponseWckeyDelete.md)
- [Dbv0037TresInfo](docs/Dbv0037TresInfo.md)
- [Dbv0037TresListInner](docs/Dbv0037TresListInner.md)
- [Dbv0037User](docs/Dbv0037User.md)
- [Dbv0037UserAssociations](docs/Dbv0037UserAssociations.md)
- [Dbv0037UserDefault](docs/Dbv0037UserDefault.md)
- [Dbv0037UserInfo](docs/Dbv0037UserInfo.md)
- [Dbv0037Wckey](docs/Dbv0037Wckey.md)
- [Dbv0037WckeyInfo](docs/Dbv0037WckeyInfo.md)
- [V0037Diag](docs/V0037Diag.md)
- [V0037DiagStatistics](docs/V0037DiagStatistics.md)
- [V0037Error](docs/V0037Error.md)
- [V0037JobProperties](docs/V0037JobProperties.md)
- [V0037JobResources](docs/V0037JobResources.md)
- [V0037JobResponseProperties](docs/V0037JobResponseProperties.md)
- [V0037JobSubmission](docs/V0037JobSubmission.md)
- [V0037JobSubmissionResponse](docs/V0037JobSubmissionResponse.md)
- [V0037JobsResponse](docs/V0037JobsResponse.md)
- [V0037Node](docs/V0037Node.md)
- [V0037NodeAllocation](docs/V0037NodeAllocation.md)
- [V0037NodesResponse](docs/V0037NodesResponse.md)
- [V0037Partition](docs/V0037Partition.md)
- [V0037PartitionsResponse](docs/V0037PartitionsResponse.md)
- [V0037Ping](docs/V0037Ping.md)
- [V0037Pings](docs/V0037Pings.md)
- [V0037Reservation](docs/V0037Reservation.md)
- [V0037ReservationPurgeCompleted](docs/V0037ReservationPurgeCompleted.md)
- [V0037ReservationsResponse](docs/V0037ReservationsResponse.md)
- [V0037Signal](docs/V0037Signal.md)

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