# \SlurmApi

All URIs are relative to _http://localhost_

| Method                                                                   | HTTP request                                          | Description                        |
| ------------------------------------------------------------------------ | ----------------------------------------------------- | ---------------------------------- |
| [**SlurmctldCancelJob**](SlurmApi.md#SlurmctldCancelJob)                 | **Delete** /slurm/v0.0.37/job/{job_id}                | cancel or signal job               |
| [**SlurmctldDiag**](SlurmApi.md#SlurmctldDiag)                           | **Get** /slurm/v0.0.37/diag                           | get diagnostics                    |
| [**SlurmctldGetJob**](SlurmApi.md#SlurmctldGetJob)                       | **Get** /slurm/v0.0.37/job/{job_id}                   | get job info                       |
| [**SlurmctldGetJobs**](SlurmApi.md#SlurmctldGetJobs)                     | **Get** /slurm/v0.0.37/jobs                           | get list of jobs                   |
| [**SlurmctldGetNode**](SlurmApi.md#SlurmctldGetNode)                     | **Get** /slurm/v0.0.37/node/{node_name}               | get node info                      |
| [**SlurmctldGetNodes**](SlurmApi.md#SlurmctldGetNodes)                   | **Get** /slurm/v0.0.37/nodes                          | get all node info                  |
| [**SlurmctldGetPartition**](SlurmApi.md#SlurmctldGetPartition)           | **Get** /slurm/v0.0.37/partition/{partition_name}     | get partition info                 |
| [**SlurmctldGetPartitions**](SlurmApi.md#SlurmctldGetPartitions)         | **Get** /slurm/v0.0.37/partitions                     | get all partition info             |
| [**SlurmctldGetReservation**](SlurmApi.md#SlurmctldGetReservation)       | **Get** /slurm/v0.0.37/reservation/{reservation_name} | get reservation info               |
| [**SlurmctldGetReservations**](SlurmApi.md#SlurmctldGetReservations)     | **Get** /slurm/v0.0.37/reservations                   | get all reservation info           |
| [**SlurmctldPing**](SlurmApi.md#SlurmctldPing)                           | **Get** /slurm/v0.0.37/ping                           | ping test                          |
| [**SlurmctldSubmitJob**](SlurmApi.md#SlurmctldSubmitJob)                 | **Post** /slurm/v0.0.37/job/submit                    | submit new job                     |
| [**SlurmctldUpdateJob**](SlurmApi.md#SlurmctldUpdateJob)                 | **Post** /slurm/v0.0.37/job/{job_id}                  | update job                         |
| [**SlurmdbdAddClusters**](SlurmApi.md#SlurmdbdAddClusters)               | **Post** /slurmdb/v0.0.37/clusters                    | Add clusters                       |
| [**SlurmdbdAddWckeys**](SlurmApi.md#SlurmdbdAddWckeys)                   | **Post** /slurmdb/v0.0.37/wckeys                      | Add wckeys                         |
| [**SlurmdbdDeleteAccount**](SlurmApi.md#SlurmdbdDeleteAccount)           | **Delete** /slurmdb/v0.0.37/account/{account_name}    | Delete account                     |
| [**SlurmdbdDeleteAssociation**](SlurmApi.md#SlurmdbdDeleteAssociation)   | **Delete** /slurmdb/v0.0.37/association               | Delete association                 |
| [**SlurmdbdDeleteCluster**](SlurmApi.md#SlurmdbdDeleteCluster)           | **Delete** /slurmdb/v0.0.37/cluster/{cluster_name}    | Delete cluster                     |
| [**SlurmdbdDeleteQos**](SlurmApi.md#SlurmdbdDeleteQos)                   | **Delete** /slurmdb/v0.0.37/qos/{qos_name}            | Delete QOS                         |
| [**SlurmdbdDeleteUser**](SlurmApi.md#SlurmdbdDeleteUser)                 | **Delete** /slurmdb/v0.0.37/user/{user_name}          | Delete user                        |
| [**SlurmdbdDeleteWckey**](SlurmApi.md#SlurmdbdDeleteWckey)               | **Delete** /slurmdb/v0.0.37/wckey/{wckey}             | Delete wckey                       |
| [**SlurmdbdDiag**](SlurmApi.md#SlurmdbdDiag)                             | **Get** /slurmdb/v0.0.37/diag                         | Get slurmdb diagnostics            |
| [**SlurmdbdGetAccount**](SlurmApi.md#SlurmdbdGetAccount)                 | **Get** /slurmdb/v0.0.37/account/{account_name}       | Get account info                   |
| [**SlurmdbdGetAccounts**](SlurmApi.md#SlurmdbdGetAccounts)               | **Get** /slurmdb/v0.0.37/accounts                     | Get account list                   |
| [**SlurmdbdGetAssociation**](SlurmApi.md#SlurmdbdGetAssociation)         | **Get** /slurmdb/v0.0.37/association                  | Get association info               |
| [**SlurmdbdGetAssociations**](SlurmApi.md#SlurmdbdGetAssociations)       | **Get** /slurmdb/v0.0.37/associations                 | Get association list               |
| [**SlurmdbdGetCluster**](SlurmApi.md#SlurmdbdGetCluster)                 | **Get** /slurmdb/v0.0.37/cluster/{cluster_name}       | Get cluster info                   |
| [**SlurmdbdGetClusters**](SlurmApi.md#SlurmdbdGetClusters)               | **Get** /slurmdb/v0.0.37/clusters                     | Get cluster list                   |
| [**SlurmdbdGetDbConfig**](SlurmApi.md#SlurmdbdGetDbConfig)               | **Get** /slurmdb/v0.0.37/config                       | Dump all configuration information |
| [**SlurmdbdGetJob**](SlurmApi.md#SlurmdbdGetJob)                         | **Get** /slurmdb/v0.0.37/job/{job_id}                 | Get job info                       |
| [**SlurmdbdGetJobs**](SlurmApi.md#SlurmdbdGetJobs)                       | **Get** /slurmdb/v0.0.37/jobs                         | Get job list                       |
| [**SlurmdbdGetQos**](SlurmApi.md#SlurmdbdGetQos)                         | **Get** /slurmdb/v0.0.37/qos                          | Get QOS list                       |
| [**SlurmdbdGetSingleQos**](SlurmApi.md#SlurmdbdGetSingleQos)             | **Get** /slurmdb/v0.0.37/qos/{qos_name}               | Get QOS info                       |
| [**SlurmdbdGetTres**](SlurmApi.md#SlurmdbdGetTres)                       | **Get** /slurmdb/v0.0.37/tres                         | Get TRES info                      |
| [**SlurmdbdGetUser**](SlurmApi.md#SlurmdbdGetUser)                       | **Get** /slurmdb/v0.0.37/user/{user_name}             | Get user info                      |
| [**SlurmdbdGetUsers**](SlurmApi.md#SlurmdbdGetUsers)                     | **Get** /slurmdb/v0.0.37/users                        | Get user list                      |
| [**SlurmdbdGetWckey**](SlurmApi.md#SlurmdbdGetWckey)                     | **Get** /slurmdb/v0.0.37/wckey/{wckey}                | Get wckey info                     |
| [**SlurmdbdGetWckeys**](SlurmApi.md#SlurmdbdGetWckeys)                   | **Get** /slurmdb/v0.0.37/wckeys                       | Get wckey list                     |
| [**SlurmdbdSetDbConfig**](SlurmApi.md#SlurmdbdSetDbConfig)               | **Post** /slurmdb/v0.0.37/config                      | Load all configuration information |
| [**SlurmdbdUpdateAccount**](SlurmApi.md#SlurmdbdUpdateAccount)           | **Post** /slurmdb/v0.0.37/accounts                    | Update accounts                    |
| [**SlurmdbdUpdateAssociations**](SlurmApi.md#SlurmdbdUpdateAssociations) | **Post** /slurmdb/v0.0.37/associations                | Set associations info              |
| [**SlurmdbdUpdateTres**](SlurmApi.md#SlurmdbdUpdateTres)                 | **Post** /slurmdb/v0.0.37/tres                        | Set TRES info                      |
| [**SlurmdbdUpdateUsers**](SlurmApi.md#SlurmdbdUpdateUsers)               | **Post** /slurmdb/v0.0.37/users                       | Update user                        |

## SlurmctldCancelJob

> SlurmctldCancelJob(ctx, jobId).Signal(signal).Execute()

cancel or signal job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    signal := openapiclient.v0.0.37_signal("HUP") // V0037Signal | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmApi.SlurmctldCancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldCancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          | Slurm Job ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldCancelJobRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**signal** | [**V0037Signal**](V0037Signal.md) | signal to send to job |

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldDiag

> V0037Diag SlurmctldDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldDiag`: V0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldDiagRequest struct via the builder pattern

### Return type

[**V0037Diag**](V0037Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetJob

> V0037JobsResponse SlurmctldGetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    jobId := "jobId_example" // string | Slurm JobID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJob`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetJob`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          | Slurm JobID                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetJobs

> V0037JobsResponse SlurmctldGetJobs(ctx).UpdateTime(updateTime).Execute()

get list of jobs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetJobs(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJobs`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetJobs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobsRequest struct via the builder pattern

| Name           | Type      | Description                                                                              | Notes |
| -------------- | --------- | ---------------------------------------------------------------------------------------- | ----- |
| **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. |

### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetNode

> V0037NodesResponse SlurmctldGetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNode`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetNode`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **nodeName** | **string**          | Slurm Node Name                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodeRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetNodes

> V0037NodesResponse SlurmctldGetNodes(ctx).UpdateTime(updateTime).Execute()

get all node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetNodes(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNodes`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetNodes`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodesRequest struct via the builder pattern

| Name           | Type      | Description                                                                              | Notes |
| -------------- | --------- | ---------------------------------------------------------------------------------------- | ----- |
| **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. |

### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetPartition

> V0037PartitionsResponse SlurmctldGetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    partitionName := "partitionName_example" // string | Slurm Partition Name
    updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartition`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetPartition`: %v\n", resp)
}
```

### Path Parameters

| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **partitionName** | **string**          | Slurm Partition Name                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. |

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetPartitions

> V0037PartitionsResponse SlurmctldGetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetPartitions(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartitions`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetPartitions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionsRequest struct via the builder pattern

| Name           | Type      | Description                                                                              | Notes |
| -------------- | --------- | ---------------------------------------------------------------------------------------- | ----- |
| **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. |

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetReservation

> V0037ReservationsResponse SlurmctldGetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    reservationName := "reservationName_example" // string | Slurm Reservation Name
    updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservation`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetReservation`: %v\n", resp)
}
```

### Path Parameters

| Name                | Type                | Description                                                                 | Notes |
| ------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**             | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **reservationName** | **string**          | Slurm Reservation Name                                                      |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. |

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldGetReservations

> V0037ReservationsResponse SlurmctldGetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldGetReservations(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservations`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetReservations`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationsRequest struct via the builder pattern

| Name           | Type      | Description                                                                              | Notes |
| -------------- | --------- | ---------------------------------------------------------------------------------------- | ----- |
| **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. |

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldPing

> V0037Pings SlurmctldPing(ctx).Execute()

ping test

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldPing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldPing`: V0037Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldPingRequest struct via the builder pattern

### Return type

[**V0037Pings**](V0037Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldSubmitJob

> V0037JobSubmissionResponse SlurmctldSubmitJob(ctx).V0037JobSubmission(v0037JobSubmission).Execute()

submit new job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    v0037JobSubmission := *openapiclient.NewV0037JobSubmission("Script_example") // V0037JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmctldSubmitJob(context.Background()).V0037JobSubmission(v0037JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldSubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldSubmitJob`: V0037JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldSubmitJob`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldSubmitJobRequest struct via the builder pattern

| Name                   | Type                                            | Description    | Notes |
| ---------------------- | ----------------------------------------------- | -------------- | ----- |
| **v0037JobSubmission** | [**V0037JobSubmission**](V0037JobSubmission.md) | submit new job |

### Return type

[**V0037JobSubmissionResponse**](V0037JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmctldUpdateJob

> SlurmctldUpdateJob(ctx, jobId).V0037JobProperties(v0037JobProperties).Execute()

update job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    v0037JobProperties := *openapiclient.NewV0037JobProperties(map[string]interface{}(123)) // V0037JobProperties | update job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmApi.SlurmctldUpdateJob(context.Background(), jobId).V0037JobProperties(v0037JobProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldUpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          | Slurm Job ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldUpdateJobRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**v0037JobProperties** | [**V0037JobProperties**](V0037JobProperties.md) | update job |

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdAddClusters

> Dbv0037ResponseClusterAdd SlurmdbdAddClusters(ctx).Execute()

Add clusters

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdAddClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdAddClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddClusters`: Dbv0037ResponseClusterAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdAddClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddClustersRequest struct via the builder pattern

### Return type

[**Dbv0037ResponseClusterAdd**](Dbv0037ResponseClusterAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdAddWckeys

> Dbv0037ResponseWckeyAdd SlurmdbdAddWckeys(ctx).Execute()

Add wckeys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdAddWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdAddWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddWckeys`: Dbv0037ResponseWckeyAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdAddWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddWckeysRequest struct via the builder pattern

### Return type

[**Dbv0037ResponseWckeyAdd**](Dbv0037ResponseWckeyAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteAccount

> Dbv0037ResponseAccountDelete SlurmdbdDeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAccount`: Dbv0037ResponseAccountDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteAccount`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **accountName** | **string**          | Slurm Account Name                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAccountRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ResponseAccountDelete**](Dbv0037ResponseAccountDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteAssociation

> Dbv0037ResponseAssociationDelete SlurmdbdDeleteAssociation(ctx).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()

Delete association

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    account := "account_example" // string | Account name
    user := "user_example" // string | User name
    cluster := "cluster_example" // string | Cluster name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteAssociation(context.Background()).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAssociation`: Dbv0037ResponseAssociationDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteAssociation`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAssociationRequest struct via the builder pattern

| Name          | Type       | Description    | Notes |
| ------------- | ---------- | -------------- | ----- |
| **account**   | **string** | Account name   |
| **user**      | **string** | User name      |
| **cluster**   | **string** | Cluster name   |
| **partition** | **string** | Partition Name |

### Return type

[**Dbv0037ResponseAssociationDelete**](Dbv0037ResponseAssociationDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteCluster

> Dbv0037ResponseClusterDelete SlurmdbdDeleteCluster(ctx, clusterName).Execute()

Delete cluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteCluster`: Dbv0037ResponseClusterDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteCluster`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **clusterName** | **string**          | Slurm cluster name                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteClusterRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ResponseClusterDelete**](Dbv0037ResponseClusterDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteQos

> Dbv0037ResponseQosDelete SlurmdbdDeleteQos(ctx, qosName).Execute()

Delete QOS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteQos`: Dbv0037ResponseQosDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteQos`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **qosName** | **string**          | Slurm QOS Name                                                              |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteQosRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ResponseQosDelete**](Dbv0037ResponseQosDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteUser

> Dbv0037ResponseUserDelete SlurmdbdDeleteUser(ctx, userName).Execute()

Delete user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    userName := "userName_example" // string | Slurm User Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteUser`: Dbv0037ResponseUserDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteUser`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **userName** | **string**          | Slurm User Name                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteUserRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ResponseUserDelete**](Dbv0037ResponseUserDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDeleteWckey

> Dbv0037ResponseWckeyDelete SlurmdbdDeleteWckey(ctx, wckey).Execute()

Delete wckey

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDeleteWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteWckey`: Dbv0037ResponseWckeyDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteWckey`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **wckey** | **string**          | Slurm wckey name                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteWckeyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ResponseWckeyDelete**](Dbv0037ResponseWckeyDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdDiag

> Dbv0037Diag SlurmdbdDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDiag`: Dbv0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDiagRequest struct via the builder pattern

### Return type

[**Dbv0037Diag**](Dbv0037Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetAccount

> Dbv0037AccountInfo SlurmdbdGetAccount(ctx, accountName).Execute()

Get account info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccount`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAccount`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **accountName** | **string**          | Slurm Account Name                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetAccounts

> Dbv0037AccountInfo SlurmdbdGetAccounts(ctx).Execute()

Get account list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccounts`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountsRequest struct via the builder pattern

### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetAssociation

> Dbv0037AssociationsInfo SlurmdbdGetAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Get association info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociation`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAssociation`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationRequest struct via the builder pattern

| Name          | Type       | Description    | Notes |
| ------------- | ---------- | -------------- | ----- |
| **cluster**   | **string** | Cluster name   |
| **account**   | **string** | Account name   |
| **user**      | **string** | User name      |
| **partition** | **string** | Partition Name |

### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetAssociations

> Dbv0037AssociationsInfo SlurmdbdGetAssociations(ctx).Execute()

Get association list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetAssociations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociations`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAssociations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationsRequest struct via the builder pattern

### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetCluster

> Dbv0037ClusterInfo SlurmdbdGetCluster(ctx, clusterName).Execute()

Get cluster info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetCluster`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetCluster`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **clusterName** | **string**          | Slurm cluster name                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClusterRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetClusters

> Dbv0037ClusterInfo SlurmdbdGetClusters(ctx).Execute()

Get cluster list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetClusters`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClustersRequest struct via the builder pattern

### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetDbConfig

> Dbv0037ConfigInfo SlurmdbdGetDbConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetDbConfig`: Dbv0037ConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetDbConfigRequest struct via the builder pattern

### Return type

[**Dbv0037ConfigInfo**](Dbv0037ConfigInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetJob

> Dbv0037JobInfo SlurmdbdGetJob(ctx, jobId).Execute()

Get job info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    jobId := int64(789) // int64 | Slurm Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJob`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetJob`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **int64**           | Slurm Job ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetJobs

> Dbv0037JobInfo SlurmdbdGetJobs(ctx).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    submitTime := "submitTime_example" // string | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    startTime := "startTime_example" // string | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    endTime := "endTime_example" // string | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    account := "account_example" // string | Comma delimited list of accounts to match (optional)
    association := "association_example" // string | Comma delimited list of associations to match (optional)
    cluster := "cluster_example" // string | Comma delimited list of cluster to match (optional)
    constraints := "constraints_example" // string | Comma delimited list of constraints to match (optional)
    cpusMax := "cpusMax_example" // string | Number of CPUs high range (optional)
    cpusMin := "cpusMin_example" // string | Number of CPUs low range (optional)
    skipSteps := true // bool | Report job step information (optional)
    disableWaitForResult := true // bool | Disable waiting for result from slurmdbd (optional)
    exitCode := "exitCode_example" // string | Exit code of job (optional)
    format := "format_example" // string | Comma delimited list of formats to match (optional)
    group := "group_example" // string | Comma delimited list of groups to match (optional)
    jobName := "jobName_example" // string | Comma delimited list of job names to match (optional)
    nodesMax := "nodesMax_example" // string | Number of nodes high range (optional)
    nodesMin := "nodesMin_example" // string | Number of nodes low range (optional)
    partition := "partition_example" // string | Comma delimited list of partitions to match (optional)
    qos := "qos_example" // string | Comma delimited list of QOS to match (optional)
    reason := "reason_example" // string | Comma delimited list of job reasons to match (optional)
    reservation := "reservation_example" // string | Comma delimited list of reservations to match (optional)
    state := "state_example" // string | Comma delimited list of states to match (optional)
    step := "step_example" // string | Comma delimited list of job steps to match (optional)
    node := "node_example" // string | Comma delimited list of used nodes to match (optional)
    wckey := "wckey_example" // string | Comma delimited list of wckeys to match (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetJobs(context.Background()).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJobs`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetJobs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobsRequest struct via the builder pattern

| Name                     | Type       | Description                                                | Notes                                                                                  |
| ------------------------ | ---------- | ---------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| **submitTime**           | **string** | Filter by submission time Accepted formats: HH:MM[:SS] [AM | PM] MMDD[YY] or MM/DD[/YY] or MM.DD[.YY] MM/DD[/YY]-HH:MM[:SS] YYYY-MM-DD[THH:MM[:SS]] |
| **startTime**            | **string** | Filter by start time Accepted formats: HH:MM[:SS] [AM      | PM] MMDD[YY] or MM/DD[/YY] or MM.DD[.YY] MM/DD[/YY]-HH:MM[:SS] YYYY-MM-DD[THH:MM[:SS]] |
| **endTime**              | **string** | Filter by end time Accepted formats: HH:MM[:SS] [AM        | PM] MMDD[YY] or MM/DD[/YY] or MM.DD[.YY] MM/DD[/YY]-HH:MM[:SS] YYYY-MM-DD[THH:MM[:SS]] |
| **account**              | **string** | Comma delimited list of accounts to match                  |
| **association**          | **string** | Comma delimited list of associations to match              |
| **cluster**              | **string** | Comma delimited list of cluster to match                   |
| **constraints**          | **string** | Comma delimited list of constraints to match               |
| **cpusMax**              | **string** | Number of CPUs high range                                  |
| **cpusMin**              | **string** | Number of CPUs low range                                   |
| **skipSteps**            | **bool**   | Report job step information                                |
| **disableWaitForResult** | **bool**   | Disable waiting for result from slurmdbd                   |
| **exitCode**             | **string** | Exit code of job                                           |
| **format**               | **string** | Comma delimited list of formats to match                   |
| **group**                | **string** | Comma delimited list of groups to match                    |
| **jobName**              | **string** | Comma delimited list of job names to match                 |
| **nodesMax**             | **string** | Number of nodes high range                                 |
| **nodesMin**             | **string** | Number of nodes low range                                  |
| **partition**            | **string** | Comma delimited list of partitions to match                |
| **qos**                  | **string** | Comma delimited list of QOS to match                       |
| **reason**               | **string** | Comma delimited list of job reasons to match               |
| **reservation**          | **string** | Comma delimited list of reservations to match              |
| **state**                | **string** | Comma delimited list of states to match                    |
| **step**                 | **string** | Comma delimited list of job steps to match                 |
| **node**                 | **string** | Comma delimited list of used nodes to match                |
| **wckey**                | **string** | Comma delimited list of wckeys to match                    |

### Return type

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetQos

> Dbv0037QosInfo SlurmdbdGetQos(ctx).Execute()

Get QOS list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetQos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetQos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetQosRequest struct via the builder pattern

### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetSingleQos

> Dbv0037QosInfo SlurmdbdGetSingleQos(ctx, qosName).Execute()

Get QOS info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetSingleQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetSingleQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetSingleQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetSingleQos`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **qosName** | **string**          | Slurm QOS Name                                                              |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetSingleQosRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetTres

> Dbv0037TresInfo SlurmdbdGetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetTres`: Dbv0037TresInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetTresRequest struct via the builder pattern

### Return type

[**Dbv0037TresInfo**](Dbv0037TresInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetUser

> Dbv0037UserInfo SlurmdbdGetUser(ctx, userName).Execute()

Get user info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    userName := "userName_example" // string | Slurm User Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUser`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetUser`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **userName** | **string**          | Slurm User Name                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUserRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetUsers

> Dbv0037UserInfo SlurmdbdGetUsers(ctx).Execute()

Get user list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUsers`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUsersRequest struct via the builder pattern

### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetWckey

> Dbv0037WckeyInfo SlurmdbdGetWckey(ctx, wckey).Execute()

Get wckey info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckey`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetWckey`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **wckey** | **string**          | Slurm wckey name                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdGetWckeys

> Dbv0037WckeyInfo SlurmdbdGetWckeys(ctx).Execute()

Get wckey list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdGetWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckeys`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeysRequest struct via the builder pattern

### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdSetDbConfig

> Dbv0037ConfigResponse SlurmdbdSetDbConfig(ctx).Execute()

Load all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdSetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdSetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdSetDbConfig`: Dbv0037ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdSetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdSetDbConfigRequest struct via the builder pattern

### Return type

[**Dbv0037ConfigResponse**](Dbv0037ConfigResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdUpdateAccount

> Dbv0037AccountResponse SlurmdbdUpdateAccount(ctx).Execute()

Update accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdUpdateAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateAccount`: Dbv0037AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateAccountRequest struct via the builder pattern

### Return type

[**Dbv0037AccountResponse**](Dbv0037AccountResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdUpdateAssociations

> Dbv0037ResponseAssociations SlurmdbdUpdateAssociations(ctx).Execute()

Set associations info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdUpdateAssociations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateAssociations`: Dbv0037ResponseAssociations
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateAssociations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateAssociationsRequest struct via the builder pattern

### Return type

[**Dbv0037ResponseAssociations**](Dbv0037ResponseAssociations.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdUpdateTres

> Dbv0037ResponseTres SlurmdbdUpdateTres(ctx).Execute()

Set TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdUpdateTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateTres`: Dbv0037ResponseTres
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateTresRequest struct via the builder pattern

### Return type

[**Dbv0037ResponseTres**](Dbv0037ResponseTres.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SlurmdbdUpdateUsers

> Dbv0037ResponseUserUpdate SlurmdbdUpdateUsers(ctx).Execute()

Update user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/heromicro/slurmrestapi.go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmApi.SlurmdbdUpdateUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateUsers`: Dbv0037ResponseUserUpdate
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateUsersRequest struct via the builder pattern

### Return type

[**Dbv0037ResponseUserUpdate**](Dbv0037ResponseUserUpdate.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
