# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0040DeleteJob**](SlurmAPI.md#SlurmV0040DeleteJob) | **Delete** /slurm/v0.0.40/job/{job_id} | cancel or signal job
[**SlurmV0040DeleteJobs**](SlurmAPI.md#SlurmV0040DeleteJobs) | **Delete** /slurm/v0.0.40/jobs/ | send signal to list of jobs
[**SlurmV0040DeleteNode**](SlurmAPI.md#SlurmV0040DeleteNode) | **Delete** /slurm/v0.0.40/node/{node_name} | delete node
[**SlurmV0040GetDiag**](SlurmAPI.md#SlurmV0040GetDiag) | **Get** /slurm/v0.0.40/diag/ | get diagnostics
[**SlurmV0040GetJob**](SlurmAPI.md#SlurmV0040GetJob) | **Get** /slurm/v0.0.40/job/{job_id} | get job info
[**SlurmV0040GetJobs**](SlurmAPI.md#SlurmV0040GetJobs) | **Get** /slurm/v0.0.40/jobs/ | get list of jobs
[**SlurmV0040GetJobsState**](SlurmAPI.md#SlurmV0040GetJobsState) | **Get** /slurm/v0.0.40/jobs/state/ | get list of job states
[**SlurmV0040GetLicenses**](SlurmAPI.md#SlurmV0040GetLicenses) | **Get** /slurm/v0.0.40/licenses/ | get all Slurm tracked license info
[**SlurmV0040GetNode**](SlurmAPI.md#SlurmV0040GetNode) | **Get** /slurm/v0.0.40/node/{node_name} | get node info
[**SlurmV0040GetNodes**](SlurmAPI.md#SlurmV0040GetNodes) | **Get** /slurm/v0.0.40/nodes/ | get node(s) info
[**SlurmV0040GetPartition**](SlurmAPI.md#SlurmV0040GetPartition) | **Get** /slurm/v0.0.40/partition/{partition_name} | get partition info
[**SlurmV0040GetPartitions**](SlurmAPI.md#SlurmV0040GetPartitions) | **Get** /slurm/v0.0.40/partitions/ | get all partition info
[**SlurmV0040GetPing**](SlurmAPI.md#SlurmV0040GetPing) | **Get** /slurm/v0.0.40/ping/ | ping test
[**SlurmV0040GetReconfigure**](SlurmAPI.md#SlurmV0040GetReconfigure) | **Get** /slurm/v0.0.40/reconfigure/ | request slurmctld reconfigure
[**SlurmV0040GetReservation**](SlurmAPI.md#SlurmV0040GetReservation) | **Get** /slurm/v0.0.40/reservation/{reservation_name} | get reservation info
[**SlurmV0040GetReservations**](SlurmAPI.md#SlurmV0040GetReservations) | **Get** /slurm/v0.0.40/reservations/ | get all reservation info
[**SlurmV0040GetShares**](SlurmAPI.md#SlurmV0040GetShares) | **Get** /slurm/v0.0.40/shares | get fairshare info
[**SlurmV0040PostJob**](SlurmAPI.md#SlurmV0040PostJob) | **Post** /slurm/v0.0.40/job/{job_id} | update job
[**SlurmV0040PostJobSubmit**](SlurmAPI.md#SlurmV0040PostJobSubmit) | **Post** /slurm/v0.0.40/job/submit | submit new job
[**SlurmV0040PostNode**](SlurmAPI.md#SlurmV0040PostNode) | **Post** /slurm/v0.0.40/node/{node_name} | update node properties
[**SlurmV0040PostNodes**](SlurmAPI.md#SlurmV0040PostNodes) | **Post** /slurm/v0.0.40/nodes/ | batch update node(s)
[**SlurmV0041DeleteJob**](SlurmAPI.md#SlurmV0041DeleteJob) | **Delete** /slurm/v0.0.41/job/{job_id} | cancel or signal job
[**SlurmV0041DeleteJobs**](SlurmAPI.md#SlurmV0041DeleteJobs) | **Delete** /slurm/v0.0.41/jobs/ | send signal to list of jobs
[**SlurmV0041DeleteNode**](SlurmAPI.md#SlurmV0041DeleteNode) | **Delete** /slurm/v0.0.41/node/{node_name} | delete node
[**SlurmV0041GetDiag**](SlurmAPI.md#SlurmV0041GetDiag) | **Get** /slurm/v0.0.41/diag/ | get diagnostics
[**SlurmV0041GetJob**](SlurmAPI.md#SlurmV0041GetJob) | **Get** /slurm/v0.0.41/job/{job_id} | get job info
[**SlurmV0041GetJobs**](SlurmAPI.md#SlurmV0041GetJobs) | **Get** /slurm/v0.0.41/jobs/ | get list of jobs
[**SlurmV0041GetJobsState**](SlurmAPI.md#SlurmV0041GetJobsState) | **Get** /slurm/v0.0.41/jobs/state/ | get list of job states
[**SlurmV0041GetLicenses**](SlurmAPI.md#SlurmV0041GetLicenses) | **Get** /slurm/v0.0.41/licenses/ | get all Slurm tracked license info
[**SlurmV0041GetNode**](SlurmAPI.md#SlurmV0041GetNode) | **Get** /slurm/v0.0.41/node/{node_name} | get node info
[**SlurmV0041GetNodes**](SlurmAPI.md#SlurmV0041GetNodes) | **Get** /slurm/v0.0.41/nodes/ | get node(s) info
[**SlurmV0041GetPartition**](SlurmAPI.md#SlurmV0041GetPartition) | **Get** /slurm/v0.0.41/partition/{partition_name} | get partition info
[**SlurmV0041GetPartitions**](SlurmAPI.md#SlurmV0041GetPartitions) | **Get** /slurm/v0.0.41/partitions/ | get all partition info
[**SlurmV0041GetPing**](SlurmAPI.md#SlurmV0041GetPing) | **Get** /slurm/v0.0.41/ping/ | ping test
[**SlurmV0041GetReconfigure**](SlurmAPI.md#SlurmV0041GetReconfigure) | **Get** /slurm/v0.0.41/reconfigure/ | request slurmctld reconfigure
[**SlurmV0041GetReservation**](SlurmAPI.md#SlurmV0041GetReservation) | **Get** /slurm/v0.0.41/reservation/{reservation_name} | get reservation info
[**SlurmV0041GetReservations**](SlurmAPI.md#SlurmV0041GetReservations) | **Get** /slurm/v0.0.41/reservations/ | get all reservation info
[**SlurmV0041GetShares**](SlurmAPI.md#SlurmV0041GetShares) | **Get** /slurm/v0.0.41/shares | get fairshare info
[**SlurmV0041PostJob**](SlurmAPI.md#SlurmV0041PostJob) | **Post** /slurm/v0.0.41/job/{job_id} | update job
[**SlurmV0041PostJobAllocate**](SlurmAPI.md#SlurmV0041PostJobAllocate) | **Post** /slurm/v0.0.41/job/allocate | submit new job allocation without any steps that must be signaled to stop
[**SlurmV0041PostJobSubmit**](SlurmAPI.md#SlurmV0041PostJobSubmit) | **Post** /slurm/v0.0.41/job/submit | submit new job
[**SlurmV0041PostNode**](SlurmAPI.md#SlurmV0041PostNode) | **Post** /slurm/v0.0.41/node/{node_name} | update node properties
[**SlurmV0041PostNodes**](SlurmAPI.md#SlurmV0041PostNodes) | **Post** /slurm/v0.0.41/nodes/ | batch update node(s)
[**SlurmV0042DeleteJob**](SlurmAPI.md#SlurmV0042DeleteJob) | **Delete** /slurm/v0.0.42/job/{job_id} | cancel or signal job
[**SlurmV0042DeleteJobs**](SlurmAPI.md#SlurmV0042DeleteJobs) | **Delete** /slurm/v0.0.42/jobs/ | send signal to list of jobs
[**SlurmV0042DeleteNode**](SlurmAPI.md#SlurmV0042DeleteNode) | **Delete** /slurm/v0.0.42/node/{node_name} | delete node
[**SlurmV0042GetDiag**](SlurmAPI.md#SlurmV0042GetDiag) | **Get** /slurm/v0.0.42/diag/ | get diagnostics
[**SlurmV0042GetJob**](SlurmAPI.md#SlurmV0042GetJob) | **Get** /slurm/v0.0.42/job/{job_id} | get job info
[**SlurmV0042GetJobs**](SlurmAPI.md#SlurmV0042GetJobs) | **Get** /slurm/v0.0.42/jobs/ | get list of jobs
[**SlurmV0042GetJobsState**](SlurmAPI.md#SlurmV0042GetJobsState) | **Get** /slurm/v0.0.42/jobs/state/ | get list of job states
[**SlurmV0042GetLicenses**](SlurmAPI.md#SlurmV0042GetLicenses) | **Get** /slurm/v0.0.42/licenses/ | get all Slurm tracked license info
[**SlurmV0042GetNode**](SlurmAPI.md#SlurmV0042GetNode) | **Get** /slurm/v0.0.42/node/{node_name} | get node info
[**SlurmV0042GetNodes**](SlurmAPI.md#SlurmV0042GetNodes) | **Get** /slurm/v0.0.42/nodes/ | get node(s) info
[**SlurmV0042GetPartition**](SlurmAPI.md#SlurmV0042GetPartition) | **Get** /slurm/v0.0.42/partition/{partition_name} | get partition info
[**SlurmV0042GetPartitions**](SlurmAPI.md#SlurmV0042GetPartitions) | **Get** /slurm/v0.0.42/partitions/ | get all partition info
[**SlurmV0042GetPing**](SlurmAPI.md#SlurmV0042GetPing) | **Get** /slurm/v0.0.42/ping/ | ping test
[**SlurmV0042GetReconfigure**](SlurmAPI.md#SlurmV0042GetReconfigure) | **Get** /slurm/v0.0.42/reconfigure/ | request slurmctld reconfigure
[**SlurmV0042GetReservation**](SlurmAPI.md#SlurmV0042GetReservation) | **Get** /slurm/v0.0.42/reservation/{reservation_name} | get reservation info
[**SlurmV0042GetReservations**](SlurmAPI.md#SlurmV0042GetReservations) | **Get** /slurm/v0.0.42/reservations/ | get all reservation info
[**SlurmV0042GetShares**](SlurmAPI.md#SlurmV0042GetShares) | **Get** /slurm/v0.0.42/shares | get fairshare info
[**SlurmV0042PostJob**](SlurmAPI.md#SlurmV0042PostJob) | **Post** /slurm/v0.0.42/job/{job_id} | update job
[**SlurmV0042PostJobAllocate**](SlurmAPI.md#SlurmV0042PostJobAllocate) | **Post** /slurm/v0.0.42/job/allocate | submit new job allocation without any steps that must be signaled to stop
[**SlurmV0042PostJobSubmit**](SlurmAPI.md#SlurmV0042PostJobSubmit) | **Post** /slurm/v0.0.42/job/submit | submit new job
[**SlurmV0042PostNode**](SlurmAPI.md#SlurmV0042PostNode) | **Post** /slurm/v0.0.42/node/{node_name} | update node properties
[**SlurmV0042PostNodes**](SlurmAPI.md#SlurmV0042PostNodes) | **Post** /slurm/v0.0.42/nodes/ | batch update node(s)



## SlurmV0040DeleteJob

> V0040OpenapiResp SlurmV0040DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040DeleteJob`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040DeleteJobs

> V0040OpenapiKillJobsResp SlurmV0040DeleteJobs(ctx).V0040KillJobsMsg(v0040KillJobsMsg).Execute()

send signal to list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040KillJobsMsg := *openapiclient.NewV0040KillJobsMsg() // V0040KillJobsMsg | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040DeleteJobs(context.Background()).V0040KillJobsMsg(v0040KillJobsMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040DeleteJobs`: V0040OpenapiKillJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040KillJobsMsg** | [**V0040KillJobsMsg**](V0040KillJobsMsg.md) | Signal or cancel jobs | 

### Return type

[**V0040OpenapiKillJobsResp**](V0040OpenapiKillJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040DeleteNode

> V0040OpenapiResp SlurmV0040DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040DeleteNode`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetDiag

> V0040OpenapiDiagResp SlurmV0040GetDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetDiag`: V0040OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetDiagRequest struct via the builder pattern


### Return type

[**V0040OpenapiDiagResp**](V0040OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetJob

> V0040OpenapiJobInfoResp SlurmV0040GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetJob`: V0040OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiJobInfoResp**](V0040OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetJobs

> V0040OpenapiJobInfoResp SlurmV0040GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetJobs`: V0040OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiJobInfoResp**](V0040OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetJobsState

> V0040OpenapiJobInfoResp SlurmV0040GetJobsState(ctx).JobId(jobId).Execute()

get list of job states

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Search for CSV list of Job IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetJobsState(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetJobsState`: V0040OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | Search for CSV list of Job IDs | 

### Return type

[**V0040OpenapiJobInfoResp**](V0040OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetLicenses

> V0040OpenapiLicensesResp SlurmV0040GetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetLicenses`: V0040OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetLicensesRequest struct via the builder pattern


### Return type

[**V0040OpenapiLicensesResp**](V0040OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetNode

> V0040OpenapiNodesResp SlurmV0040GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetNode`: V0040OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiNodesResp**](V0040OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetNodes

> V0040OpenapiNodesResp SlurmV0040GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get node(s) info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetNodes`: V0040OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiNodesResp**](V0040OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPartition

> V0040OpenapiPartitionResp SlurmV0040GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Partition name
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPartition`: V0040OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiPartitionResp**](V0040OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPartitions

> V0040OpenapiPartitionResp SlurmV0040GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPartitions`: V0040OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiPartitionResp**](V0040OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPing

> V0040OpenapiPingArrayResp SlurmV0040GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPing`: V0040OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPingRequest struct via the builder pattern


### Return type

[**V0040OpenapiPingArrayResp**](V0040OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReconfigure

> V0040OpenapiResp SlurmV0040GetReconfigure(ctx).Execute()

request slurmctld reconfigure

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReconfigure`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReconfigureRequest struct via the builder pattern


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReservation

> V0040OpenapiReservationResp SlurmV0040GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Reservation name
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReservation`: V0040OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0040OpenapiReservationResp**](V0040OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReservations

> V0040OpenapiReservationResp SlurmV0040GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReservations`: V0040OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0040OpenapiReservationResp**](V0040OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetShares

> V0040OpenapiSharesResp SlurmV0040GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetShares`: V0040OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0040OpenapiSharesResp**](V0040OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostJob

> V0040OpenapiJobPostResponse SlurmV0040PostJob(ctx, jobId).V0040JobDescMsg(v0040JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	v0040JobDescMsg := *openapiclient.NewV0040JobDescMsg() // V0040JobDescMsg | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostJob(context.Background(), jobId).V0040JobDescMsg(v0040JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostJob`: V0040OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0040JobDescMsg** | [**V0040JobDescMsg**](V0040JobDescMsg.md) | Job update description | 

### Return type

[**V0040OpenapiJobPostResponse**](V0040OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostJobSubmit

> V0040OpenapiJobSubmitResponse SlurmV0040PostJobSubmit(ctx).V0040JobSubmitReq(v0040JobSubmitReq).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040JobSubmitReq := *openapiclient.NewV0040JobSubmitReq() // V0040JobSubmitReq | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostJobSubmit(context.Background()).V0040JobSubmitReq(v0040JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostJobSubmit`: V0040OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040JobSubmitReq** | [**V0040JobSubmitReq**](V0040JobSubmitReq.md) | Job description | 

### Return type

[**V0040OpenapiJobSubmitResponse**](V0040OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostNode

> V0040OpenapiResp SlurmV0040PostNode(ctx, nodeName).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	v0040UpdateNodeMsg := *openapiclient.NewV0040UpdateNodeMsg() // V0040UpdateNodeMsg | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostNode(context.Background(), nodeName).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostNode`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0040UpdateNodeMsg** | [**V0040UpdateNodeMsg**](V0040UpdateNodeMsg.md) | Node update description | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostNodes

> V0040OpenapiResp SlurmV0040PostNodes(ctx).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()

batch update node(s)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040UpdateNodeMsg := *openapiclient.NewV0040UpdateNodeMsg() // V0040UpdateNodeMsg | Nodelist update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostNodes(context.Background()).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostNodes`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040UpdateNodeMsg** | [**V0040UpdateNodeMsg**](V0040UpdateNodeMsg.md) | Nodelist update description | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteJob

> V0041OpenapiResp SlurmV0041DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJob`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteJobs

> V0041OpenapiKillJobsResp SlurmV0041DeleteJobs(ctx).V0041KillJobsMsg(v0041KillJobsMsg).Execute()

send signal to list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041KillJobsMsg := *openapiclient.NewV0041KillJobsMsg() // V0041KillJobsMsg | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJobs(context.Background()).V0041KillJobsMsg(v0041KillJobsMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJobs`: V0041OpenapiKillJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041KillJobsMsg** | [**V0041KillJobsMsg**](V0041KillJobsMsg.md) | Signal or cancel jobs | 

### Return type

[**V0041OpenapiKillJobsResp**](V0041OpenapiKillJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteNode

> V0041OpenapiResp SlurmV0041DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetDiag

> V0041OpenapiDiagResp SlurmV0041GetDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetDiag`: V0041OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetDiagRequest struct via the builder pattern


### Return type

[**V0041OpenapiDiagResp**](V0041OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJob

> V0041OpenapiJobInfoResp SlurmV0041GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJob`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJobs

> V0041OpenapiJobInfoResp SlurmV0041GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJobs`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJobsState

> V0041OpenapiJobInfoResp SlurmV0041GetJobsState(ctx).JobId(jobId).Execute()

get list of job states

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Search for CSV list of Job IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJobsState(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJobsState`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | Search for CSV list of Job IDs | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetLicenses

> V0041OpenapiLicensesResp SlurmV0041GetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetLicenses`: V0041OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetLicensesRequest struct via the builder pattern


### Return type

[**V0041OpenapiLicensesResp**](V0041OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNode

> V0041OpenapiNodesResp SlurmV0041GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNode`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNodes

> V0041OpenapiNodesResp SlurmV0041GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get node(s) info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNodes`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartition

> V0041OpenapiPartitionResp SlurmV0041GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Partition name
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartition`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartitions

> V0041OpenapiPartitionResp SlurmV0041GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartitions`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPing

> V0041OpenapiPingArrayResp SlurmV0041GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPing`: V0041OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPingRequest struct via the builder pattern


### Return type

[**V0041OpenapiPingArrayResp**](V0041OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReconfigure

> V0041OpenapiResp SlurmV0041GetReconfigure(ctx).Execute()

request slurmctld reconfigure

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReconfigure`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReconfigureRequest struct via the builder pattern


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReservation

> V0041OpenapiReservationResp SlurmV0041GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Reservation name
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservation`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReservations

> V0041OpenapiReservationResp SlurmV0041GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservations`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetShares

> V0041OpenapiSharesResp SlurmV0041GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetShares`: V0041OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0041OpenapiSharesResp**](V0041OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJob

> V0041OpenapiJobPostResponse SlurmV0041PostJob(ctx, jobId).V0041JobDescMsg(v0041JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	v0041JobDescMsg := *openapiclient.NewV0041JobDescMsg() // V0041JobDescMsg | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJob(context.Background(), jobId).V0041JobDescMsg(v0041JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJob`: V0041OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0041JobDescMsg** | [**V0041JobDescMsg**](V0041JobDescMsg.md) | Job update description | 

### Return type

[**V0041OpenapiJobPostResponse**](V0041OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobAllocate

> V0041OpenapiJobAllocResp SlurmV0041PostJobAllocate(ctx).V0041JobAllocReq(v0041JobAllocReq).Execute()

submit new job allocation without any steps that must be signaled to stop

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041JobAllocReq := *openapiclient.NewV0041JobAllocReq() // V0041JobAllocReq | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobAllocate(context.Background()).V0041JobAllocReq(v0041JobAllocReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobAllocate`: V0041OpenapiJobAllocResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041JobAllocReq** | [**V0041JobAllocReq**](V0041JobAllocReq.md) | Job allocation description | 

### Return type

[**V0041OpenapiJobAllocResp**](V0041OpenapiJobAllocResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobSubmit

> V0041OpenapiJobSubmitResponse SlurmV0041PostJobSubmit(ctx).V0041JobSubmitReq(v0041JobSubmitReq).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041JobSubmitReq := *openapiclient.NewV0041JobSubmitReq() // V0041JobSubmitReq | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobSubmit(context.Background()).V0041JobSubmitReq(v0041JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobSubmit`: V0041OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041JobSubmitReq** | [**V0041JobSubmitReq**](V0041JobSubmitReq.md) | Job description | 

### Return type

[**V0041OpenapiJobSubmitResponse**](V0041OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostNode

> V0041OpenapiResp SlurmV0041PostNode(ctx, nodeName).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	v0041UpdateNodeMsg := *openapiclient.NewV0041UpdateNodeMsg() // V0041UpdateNodeMsg | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostNode(context.Background(), nodeName).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0041UpdateNodeMsg** | [**V0041UpdateNodeMsg**](V0041UpdateNodeMsg.md) | Node update description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostNodes

> V0041OpenapiResp SlurmV0041PostNodes(ctx).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()

batch update node(s)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041UpdateNodeMsg := *openapiclient.NewV0041UpdateNodeMsg() // V0041UpdateNodeMsg | Nodelist update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostNodes(context.Background()).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostNodes`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041UpdateNodeMsg** | [**V0041UpdateNodeMsg**](V0041UpdateNodeMsg.md) | Nodelist update description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042DeleteJob

> V0042OpenapiKillJobResp SlurmV0042DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042DeleteJob`: V0042OpenapiKillJobResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0042OpenapiKillJobResp**](V0042OpenapiKillJobResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042DeleteJobs

> V0042OpenapiKillJobsResp SlurmV0042DeleteJobs(ctx).V0042KillJobsMsg(v0042KillJobsMsg).Execute()

send signal to list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042KillJobsMsg := *openapiclient.NewV0042KillJobsMsg() // V0042KillJobsMsg | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042DeleteJobs(context.Background()).V0042KillJobsMsg(v0042KillJobsMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042DeleteJobs`: V0042OpenapiKillJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042KillJobsMsg** | [**V0042KillJobsMsg**](V0042KillJobsMsg.md) | Signal or cancel jobs | 

### Return type

[**V0042OpenapiKillJobsResp**](V0042OpenapiKillJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042DeleteNode

> V0042OpenapiResp SlurmV0042DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042DeleteNode`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetDiag

> V0042OpenapiDiagResp SlurmV0042GetDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetDiag`: V0042OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetDiagRequest struct via the builder pattern


### Return type

[**V0042OpenapiDiagResp**](V0042OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetJob

> V0042OpenapiJobInfoResp SlurmV0042GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	updateTime := "updateTime_example" // string | Query jobs updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetJob`: V0042OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiJobInfoResp**](V0042OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetJobs

> V0042OpenapiJobInfoResp SlurmV0042GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query jobs updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetJobs`: V0042OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiJobInfoResp**](V0042OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetJobsState

> V0042OpenapiJobInfoResp SlurmV0042GetJobsState(ctx).JobId(jobId).Execute()

get list of job states

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | CSV list of Job IDs to search for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetJobsState(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetJobsState`: V0042OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | CSV list of Job IDs to search for | 

### Return type

[**V0042OpenapiJobInfoResp**](V0042OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetLicenses

> V0042OpenapiLicensesResp SlurmV0042GetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetLicenses`: V0042OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetLicensesRequest struct via the builder pattern


### Return type

[**V0042OpenapiLicensesResp**](V0042OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetNode

> V0042OpenapiNodesResp SlurmV0042GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Query jobs updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetNode`: V0042OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiNodesResp**](V0042OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetNodes

> V0042OpenapiNodesResp SlurmV0042GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get node(s) info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query jobs updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetNodes`: V0042OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiNodesResp**](V0042OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetPartition

> V0042OpenapiPartitionResp SlurmV0042GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Partition name
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetPartition`: V0042OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiPartitionResp**](V0042OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetPartitions

> V0042OpenapiPartitionResp SlurmV0042GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetPartitions`: V0042OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0042OpenapiPartitionResp**](V0042OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetPing

> V0042OpenapiPingArrayResp SlurmV0042GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetPing`: V0042OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetPingRequest struct via the builder pattern


### Return type

[**V0042OpenapiPingArrayResp**](V0042OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetReconfigure

> V0042OpenapiResp SlurmV0042GetReconfigure(ctx).Execute()

request slurmctld reconfigure

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetReconfigure`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetReconfigureRequest struct via the builder pattern


### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetReservation

> V0042OpenapiReservationResp SlurmV0042GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Reservation name
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetReservation`: V0042OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0042OpenapiReservationResp**](V0042OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetReservations

> V0042OpenapiReservationResp SlurmV0042GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetReservations`: V0042OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0042OpenapiReservationResp**](V0042OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042GetShares

> V0042OpenapiSharesResp SlurmV0042GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042GetShares`: V0042OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0042OpenapiSharesResp**](V0042OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042PostJob

> V0042OpenapiJobPostResponse SlurmV0042PostJob(ctx, jobId).V0042JobDescMsg(v0042JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	v0042JobDescMsg := *openapiclient.NewV0042JobDescMsg() // V0042JobDescMsg | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042PostJob(context.Background(), jobId).V0042JobDescMsg(v0042JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042PostJob`: V0042OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0042JobDescMsg** | [**V0042JobDescMsg**](V0042JobDescMsg.md) | Job update description | 

### Return type

[**V0042OpenapiJobPostResponse**](V0042OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042PostJobAllocate

> V0042OpenapiJobAllocResp SlurmV0042PostJobAllocate(ctx).V0042JobAllocReq(v0042JobAllocReq).Execute()

submit new job allocation without any steps that must be signaled to stop

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042JobAllocReq := *openapiclient.NewV0042JobAllocReq() // V0042JobAllocReq | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042PostJobAllocate(context.Background()).V0042JobAllocReq(v0042JobAllocReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042PostJobAllocate`: V0042OpenapiJobAllocResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042JobAllocReq** | [**V0042JobAllocReq**](V0042JobAllocReq.md) | Job allocation description | 

### Return type

[**V0042OpenapiJobAllocResp**](V0042OpenapiJobAllocResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042PostJobSubmit

> V0042OpenapiJobSubmitResponse SlurmV0042PostJobSubmit(ctx).V0042JobSubmitReq(v0042JobSubmitReq).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042JobSubmitReq := *openapiclient.NewV0042JobSubmitReq() // V0042JobSubmitReq | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042PostJobSubmit(context.Background()).V0042JobSubmitReq(v0042JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042PostJobSubmit`: V0042OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042JobSubmitReq** | [**V0042JobSubmitReq**](V0042JobSubmitReq.md) | Job description | 

### Return type

[**V0042OpenapiJobSubmitResponse**](V0042OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042PostNode

> V0042OpenapiResp SlurmV0042PostNode(ctx, nodeName).V0042UpdateNodeMsg(v0042UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	v0042UpdateNodeMsg := *openapiclient.NewV0042UpdateNodeMsg() // V0042UpdateNodeMsg | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042PostNode(context.Background(), nodeName).V0042UpdateNodeMsg(v0042UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042PostNode`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0042UpdateNodeMsg** | [**V0042UpdateNodeMsg**](V0042UpdateNodeMsg.md) | Node update description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042PostNodes

> V0042OpenapiResp SlurmV0042PostNodes(ctx).V0042UpdateNodeMsg(v0042UpdateNodeMsg).Execute()

batch update node(s)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042UpdateNodeMsg := *openapiclient.NewV0042UpdateNodeMsg() // V0042UpdateNodeMsg | Nodelist update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042PostNodes(context.Background()).V0042UpdateNodeMsg(v0042UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042PostNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042PostNodes`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042PostNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042PostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042UpdateNodeMsg** | [**V0042UpdateNodeMsg**](V0042UpdateNodeMsg.md) | Nodelist update description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

