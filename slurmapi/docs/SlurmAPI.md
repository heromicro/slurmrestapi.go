# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0041DeleteJob**](SlurmAPI.md#SlurmV0041DeleteJob) | **Delete** /slurm/v0.0.41/job/{job_id} | cancel or signal job
[**SlurmV0041DeleteJobs**](SlurmAPI.md#SlurmV0041DeleteJobs) | **Delete** /slurm/v0.0.41/jobs/ | send signal to list of jobs
[**SlurmV0041DeleteNode**](SlurmAPI.md#SlurmV0041DeleteNode) | **Delete** /slurm/v0.0.41/node/{node_name} | delete node
[**SlurmV0041DeleteReservation**](SlurmAPI.md#SlurmV0041DeleteReservation) | **Delete** /slurm/v0.0.41/reservation/{reservation_name} | delete a reservation
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
[**SlurmV0042DeleteReservation**](SlurmAPI.md#SlurmV0042DeleteReservation) | **Delete** /slurm/v0.0.42/reservation/{reservation_name} | delete a reservation
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
[**SlurmV0043DeleteJob**](SlurmAPI.md#SlurmV0043DeleteJob) | **Delete** /slurm/v0.0.43/job/{job_id} | cancel or signal job
[**SlurmV0043DeleteJobs**](SlurmAPI.md#SlurmV0043DeleteJobs) | **Delete** /slurm/v0.0.43/jobs/ | send signal to list of jobs
[**SlurmV0043DeleteNode**](SlurmAPI.md#SlurmV0043DeleteNode) | **Delete** /slurm/v0.0.43/node/{node_name} | delete node
[**SlurmV0043DeleteReservation**](SlurmAPI.md#SlurmV0043DeleteReservation) | **Delete** /slurm/v0.0.43/reservation/{reservation_name} | delete a reservation
[**SlurmV0043GetDiag**](SlurmAPI.md#SlurmV0043GetDiag) | **Get** /slurm/v0.0.43/diag/ | get diagnostics
[**SlurmV0043GetJob**](SlurmAPI.md#SlurmV0043GetJob) | **Get** /slurm/v0.0.43/job/{job_id} | get job info
[**SlurmV0043GetJobs**](SlurmAPI.md#SlurmV0043GetJobs) | **Get** /slurm/v0.0.43/jobs/ | get list of jobs
[**SlurmV0043GetJobsState**](SlurmAPI.md#SlurmV0043GetJobsState) | **Get** /slurm/v0.0.43/jobs/state/ | get list of job states
[**SlurmV0043GetLicenses**](SlurmAPI.md#SlurmV0043GetLicenses) | **Get** /slurm/v0.0.43/licenses/ | get all Slurm tracked license info
[**SlurmV0043GetNode**](SlurmAPI.md#SlurmV0043GetNode) | **Get** /slurm/v0.0.43/node/{node_name} | get node info
[**SlurmV0043GetNodes**](SlurmAPI.md#SlurmV0043GetNodes) | **Get** /slurm/v0.0.43/nodes/ | get node(s) info
[**SlurmV0043GetPartition**](SlurmAPI.md#SlurmV0043GetPartition) | **Get** /slurm/v0.0.43/partition/{partition_name} | get partition info
[**SlurmV0043GetPartitions**](SlurmAPI.md#SlurmV0043GetPartitions) | **Get** /slurm/v0.0.43/partitions/ | get all partition info
[**SlurmV0043GetPing**](SlurmAPI.md#SlurmV0043GetPing) | **Get** /slurm/v0.0.43/ping/ | ping test
[**SlurmV0043GetReconfigure**](SlurmAPI.md#SlurmV0043GetReconfigure) | **Get** /slurm/v0.0.43/reconfigure/ | request slurmctld reconfigure
[**SlurmV0043GetReservation**](SlurmAPI.md#SlurmV0043GetReservation) | **Get** /slurm/v0.0.43/reservation/{reservation_name} | get reservation info
[**SlurmV0043GetReservations**](SlurmAPI.md#SlurmV0043GetReservations) | **Get** /slurm/v0.0.43/reservations/ | get all reservation info
[**SlurmV0043GetShares**](SlurmAPI.md#SlurmV0043GetShares) | **Get** /slurm/v0.0.43/shares | get fairshare info
[**SlurmV0043PostJob**](SlurmAPI.md#SlurmV0043PostJob) | **Post** /slurm/v0.0.43/job/{job_id} | update job
[**SlurmV0043PostJobAllocate**](SlurmAPI.md#SlurmV0043PostJobAllocate) | **Post** /slurm/v0.0.43/job/allocate | submit new job allocation without any steps that must be signaled to stop
[**SlurmV0043PostJobSubmit**](SlurmAPI.md#SlurmV0043PostJobSubmit) | **Post** /slurm/v0.0.43/job/submit | submit new job
[**SlurmV0043PostNode**](SlurmAPI.md#SlurmV0043PostNode) | **Post** /slurm/v0.0.43/node/{node_name} | update node properties
[**SlurmV0043PostNodes**](SlurmAPI.md#SlurmV0043PostNodes) | **Post** /slurm/v0.0.43/nodes/ | batch update node(s)
[**SlurmV0043PostReservation**](SlurmAPI.md#SlurmV0043PostReservation) | **Post** /slurm/v0.0.43/reservation | create or update a reservation
[**SlurmV0043PostReservations**](SlurmAPI.md#SlurmV0043PostReservations) | **Post** /slurm/v0.0.43/reservations/ | create or update reservations
[**SlurmV0044DeleteJob**](SlurmAPI.md#SlurmV0044DeleteJob) | **Delete** /slurm/v0.0.44/job/{job_id} | cancel or signal job
[**SlurmV0044DeleteJobs**](SlurmAPI.md#SlurmV0044DeleteJobs) | **Delete** /slurm/v0.0.44/jobs/ | send signal to list of jobs
[**SlurmV0044DeleteNode**](SlurmAPI.md#SlurmV0044DeleteNode) | **Delete** /slurm/v0.0.44/node/{node_name} | delete node
[**SlurmV0044DeleteReservation**](SlurmAPI.md#SlurmV0044DeleteReservation) | **Delete** /slurm/v0.0.44/reservation/{reservation_name} | delete a reservation
[**SlurmV0044GetDiag**](SlurmAPI.md#SlurmV0044GetDiag) | **Get** /slurm/v0.0.44/diag/ | get diagnostics
[**SlurmV0044GetJob**](SlurmAPI.md#SlurmV0044GetJob) | **Get** /slurm/v0.0.44/job/{job_id} | get job info
[**SlurmV0044GetJobs**](SlurmAPI.md#SlurmV0044GetJobs) | **Get** /slurm/v0.0.44/jobs/ | get list of jobs
[**SlurmV0044GetJobsState**](SlurmAPI.md#SlurmV0044GetJobsState) | **Get** /slurm/v0.0.44/jobs/state/ | get list of job states
[**SlurmV0044GetLicenses**](SlurmAPI.md#SlurmV0044GetLicenses) | **Get** /slurm/v0.0.44/licenses/ | get all Slurm tracked license info
[**SlurmV0044GetNode**](SlurmAPI.md#SlurmV0044GetNode) | **Get** /slurm/v0.0.44/node/{node_name} | get node info
[**SlurmV0044GetNodes**](SlurmAPI.md#SlurmV0044GetNodes) | **Get** /slurm/v0.0.44/nodes/ | get node(s) info
[**SlurmV0044GetPartition**](SlurmAPI.md#SlurmV0044GetPartition) | **Get** /slurm/v0.0.44/partition/{partition_name} | get partition info
[**SlurmV0044GetPartitions**](SlurmAPI.md#SlurmV0044GetPartitions) | **Get** /slurm/v0.0.44/partitions/ | get all partition info
[**SlurmV0044GetPing**](SlurmAPI.md#SlurmV0044GetPing) | **Get** /slurm/v0.0.44/ping/ | ping test
[**SlurmV0044GetReconfigure**](SlurmAPI.md#SlurmV0044GetReconfigure) | **Get** /slurm/v0.0.44/reconfigure/ | request slurmctld reconfigure
[**SlurmV0044GetReservation**](SlurmAPI.md#SlurmV0044GetReservation) | **Get** /slurm/v0.0.44/reservation/{reservation_name} | get reservation info
[**SlurmV0044GetReservations**](SlurmAPI.md#SlurmV0044GetReservations) | **Get** /slurm/v0.0.44/reservations/ | get all reservation info
[**SlurmV0044GetResources**](SlurmAPI.md#SlurmV0044GetResources) | **Get** /slurm/v0.0.44/resources/{job_id} | get resource layout info
[**SlurmV0044GetShares**](SlurmAPI.md#SlurmV0044GetShares) | **Get** /slurm/v0.0.44/shares | get fairshare info
[**SlurmV0044PostJob**](SlurmAPI.md#SlurmV0044PostJob) | **Post** /slurm/v0.0.44/job/{job_id} | update job
[**SlurmV0044PostJobAllocate**](SlurmAPI.md#SlurmV0044PostJobAllocate) | **Post** /slurm/v0.0.44/job/allocate | submit new job allocation without any steps that must be signaled to stop
[**SlurmV0044PostJobSubmit**](SlurmAPI.md#SlurmV0044PostJobSubmit) | **Post** /slurm/v0.0.44/job/submit | submit new job
[**SlurmV0044PostNewNode**](SlurmAPI.md#SlurmV0044PostNewNode) | **Post** /slurm/v0.0.44/new/node/ | create node
[**SlurmV0044PostNode**](SlurmAPI.md#SlurmV0044PostNode) | **Post** /slurm/v0.0.44/node/{node_name} | update node properties
[**SlurmV0044PostNodes**](SlurmAPI.md#SlurmV0044PostNodes) | **Post** /slurm/v0.0.44/nodes/ | batch update node(s)
[**SlurmV0044PostReservation**](SlurmAPI.md#SlurmV0044PostReservation) | **Post** /slurm/v0.0.44/reservation | create or update a reservation
[**SlurmV0044PostReservations**](SlurmAPI.md#SlurmV0044PostReservations) | **Post** /slurm/v0.0.44/reservations/ | create or update reservations



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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteJobs

> SlurmV0041DeleteJobs200Response SlurmV0041DeleteJobs(ctx).SlurmV0041DeleteJobsRequest(slurmV0041DeleteJobsRequest).Execute()

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
	slurmV0041DeleteJobsRequest := *openapiclient.NewSlurmV0041DeleteJobsRequest() // SlurmV0041DeleteJobsRequest | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJobs(context.Background()).SlurmV0041DeleteJobsRequest(slurmV0041DeleteJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJobs`: SlurmV0041DeleteJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041DeleteJobsRequest** | [**SlurmV0041DeleteJobsRequest**](SlurmV0041DeleteJobsRequest.md) | Signal or cancel jobs | 

### Return type

[**SlurmV0041DeleteJobs200Response**](SlurmV0041DeleteJobs200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteReservation

> V0041OpenapiResp SlurmV0041DeleteReservation(ctx, reservationName).Execute()

delete a reservation

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteReservation(context.Background(), reservationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteReservation`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetDiag

> SlurmV0041GetDiag200Response SlurmV0041GetDiag(ctx).Execute()

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
	// response from `SlurmV0041GetDiag`: SlurmV0041GetDiag200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetDiagRequest struct via the builder pattern


### Return type

[**SlurmV0041GetDiag200Response**](SlurmV0041GetDiag200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetLicenses

> SlurmV0041GetLicenses200Response SlurmV0041GetLicenses(ctx).Execute()

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
	// response from `SlurmV0041GetLicenses`: SlurmV0041GetLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetLicensesRequest struct via the builder pattern


### Return type

[**SlurmV0041GetLicenses200Response**](SlurmV0041GetLicenses200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPing

> SlurmV0041GetPing200Response SlurmV0041GetPing(ctx).Execute()

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
	// response from `SlurmV0041GetPing`: SlurmV0041GetPing200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPingRequest struct via the builder pattern


### Return type

[**SlurmV0041GetPing200Response**](SlurmV0041GetPing200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetShares

> SlurmV0041GetShares200Response SlurmV0041GetShares(ctx).Accounts(accounts).Users(users).Execute()

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
	// response from `SlurmV0041GetShares`: SlurmV0041GetShares200Response
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

[**SlurmV0041GetShares200Response**](SlurmV0041GetShares200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJob

> SlurmV0041PostJob200Response SlurmV0041PostJob(ctx, jobId).SlurmV0041PostJobSubmitRequestJobsInner(slurmV0041PostJobSubmitRequestJobsInner).Execute()

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
	slurmV0041PostJobSubmitRequestJobsInner := *openapiclient.NewSlurmV0041PostJobSubmitRequestJobsInner() // SlurmV0041PostJobSubmitRequestJobsInner | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJob(context.Background(), jobId).SlurmV0041PostJobSubmitRequestJobsInner(slurmV0041PostJobSubmitRequestJobsInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJob`: SlurmV0041PostJob200Response
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

 **slurmV0041PostJobSubmitRequestJobsInner** | [**SlurmV0041PostJobSubmitRequestJobsInner**](SlurmV0041PostJobSubmitRequestJobsInner.md) | Job update description | 

### Return type

[**SlurmV0041PostJob200Response**](SlurmV0041PostJob200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobAllocate

> SlurmV0041PostJobAllocate200Response SlurmV0041PostJobAllocate(ctx).SlurmV0041PostJobAllocateRequest(slurmV0041PostJobAllocateRequest).Execute()

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
	slurmV0041PostJobAllocateRequest := *openapiclient.NewSlurmV0041PostJobAllocateRequest() // SlurmV0041PostJobAllocateRequest | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobAllocate(context.Background()).SlurmV0041PostJobAllocateRequest(slurmV0041PostJobAllocateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobAllocate`: SlurmV0041PostJobAllocate200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041PostJobAllocateRequest** | [**SlurmV0041PostJobAllocateRequest**](SlurmV0041PostJobAllocateRequest.md) | Job allocation description | 

### Return type

[**SlurmV0041PostJobAllocate200Response**](SlurmV0041PostJobAllocate200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobSubmit

> SlurmV0041PostJobSubmit200Response SlurmV0041PostJobSubmit(ctx).SlurmV0041PostJobSubmitRequest(slurmV0041PostJobSubmitRequest).Execute()

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
	slurmV0041PostJobSubmitRequest := *openapiclient.NewSlurmV0041PostJobSubmitRequest() // SlurmV0041PostJobSubmitRequest | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobSubmit(context.Background()).SlurmV0041PostJobSubmitRequest(slurmV0041PostJobSubmitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobSubmit`: SlurmV0041PostJobSubmit200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041PostJobSubmitRequest** | [**SlurmV0041PostJobSubmitRequest**](SlurmV0041PostJobSubmitRequest.md) | Job description | 

### Return type

[**SlurmV0041PostJobSubmit200Response**](SlurmV0041PostJobSubmit200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0042DeleteReservation

> V0042OpenapiResp SlurmV0042DeleteReservation(ctx, reservationName).Execute()

delete a reservation

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0042DeleteReservation(context.Background(), reservationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0042DeleteReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0042DeleteReservation`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0042DeleteReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0042DeleteReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043DeleteJob

> V0043OpenapiKillJobResp SlurmV0043DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043DeleteJob`: V0043OpenapiKillJobResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0043OpenapiKillJobResp**](V0043OpenapiKillJobResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043DeleteJobs

> V0043OpenapiKillJobsResp SlurmV0043DeleteJobs(ctx).V0043KillJobsMsg(v0043KillJobsMsg).Execute()

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
	v0043KillJobsMsg := *openapiclient.NewV0043KillJobsMsg() // V0043KillJobsMsg | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043DeleteJobs(context.Background()).V0043KillJobsMsg(v0043KillJobsMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043DeleteJobs`: V0043OpenapiKillJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043KillJobsMsg** | [**V0043KillJobsMsg**](V0043KillJobsMsg.md) | Signal or cancel jobs | 

### Return type

[**V0043OpenapiKillJobsResp**](V0043OpenapiKillJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043DeleteNode

> V0043OpenapiResp SlurmV0043DeleteNode(ctx, nodeName).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043DeleteNode`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043DeleteReservation

> V0043OpenapiResp SlurmV0043DeleteReservation(ctx, reservationName).Execute()

delete a reservation

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043DeleteReservation(context.Background(), reservationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043DeleteReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043DeleteReservation`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043DeleteReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043DeleteReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetDiag

> V0043OpenapiDiagResp SlurmV0043GetDiag(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetDiag`: V0043OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetDiagRequest struct via the builder pattern


### Return type

[**V0043OpenapiDiagResp**](V0043OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetJob

> V0043OpenapiJobInfoResp SlurmV0043GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetJob`: V0043OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiJobInfoResp**](V0043OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetJobs

> V0043OpenapiJobInfoResp SlurmV0043GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetJobs`: V0043OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiJobInfoResp**](V0043OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetJobsState

> V0043OpenapiJobInfoResp SlurmV0043GetJobsState(ctx).JobId(jobId).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetJobsState(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetJobsState`: V0043OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | CSV list of Job IDs to search for | 

### Return type

[**V0043OpenapiJobInfoResp**](V0043OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetLicenses

> V0043OpenapiLicensesResp SlurmV0043GetLicenses(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetLicenses`: V0043OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetLicensesRequest struct via the builder pattern


### Return type

[**V0043OpenapiLicensesResp**](V0043OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetNode

> V0043OpenapiNodesResp SlurmV0043GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetNode`: V0043OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiNodesResp**](V0043OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetNodes

> V0043OpenapiNodesResp SlurmV0043GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetNodes`: V0043OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiNodesResp**](V0043OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetPartition

> V0043OpenapiPartitionResp SlurmV0043GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetPartition`: V0043OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiPartitionResp**](V0043OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetPartitions

> V0043OpenapiPartitionResp SlurmV0043GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetPartitions`: V0043OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0043OpenapiPartitionResp**](V0043OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetPing

> V0043OpenapiPingArrayResp SlurmV0043GetPing(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetPing`: V0043OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetPingRequest struct via the builder pattern


### Return type

[**V0043OpenapiPingArrayResp**](V0043OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetReconfigure

> V0043OpenapiResp SlurmV0043GetReconfigure(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetReconfigure`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetReconfigureRequest struct via the builder pattern


### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetReservation

> V0043OpenapiReservationResp SlurmV0043GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetReservation`: V0043OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0043OpenapiReservationResp**](V0043OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetReservations

> V0043OpenapiReservationResp SlurmV0043GetReservations(ctx).UpdateTime(updateTime).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetReservations`: V0043OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0043OpenapiReservationResp**](V0043OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043GetShares

> V0043OpenapiSharesResp SlurmV0043GetShares(ctx).Accounts(accounts).Users(users).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0043GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043GetShares`: V0043OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0043OpenapiSharesResp**](V0043OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostJob

> V0043OpenapiJobPostResponse SlurmV0043PostJob(ctx, jobId).V0043JobDescMsg(v0043JobDescMsg).Execute()

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
	v0043JobDescMsg := *openapiclient.NewV0043JobDescMsg() // V0043JobDescMsg | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostJob(context.Background(), jobId).V0043JobDescMsg(v0043JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostJob`: V0043OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0043JobDescMsg** | [**V0043JobDescMsg**](V0043JobDescMsg.md) | Job update description | 

### Return type

[**V0043OpenapiJobPostResponse**](V0043OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostJobAllocate

> V0043OpenapiJobAllocResp SlurmV0043PostJobAllocate(ctx).V0043JobAllocReq(v0043JobAllocReq).Execute()

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
	v0043JobAllocReq := *openapiclient.NewV0043JobAllocReq() // V0043JobAllocReq | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostJobAllocate(context.Background()).V0043JobAllocReq(v0043JobAllocReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostJobAllocate`: V0043OpenapiJobAllocResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043JobAllocReq** | [**V0043JobAllocReq**](V0043JobAllocReq.md) | Job allocation description | 

### Return type

[**V0043OpenapiJobAllocResp**](V0043OpenapiJobAllocResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostJobSubmit

> V0043OpenapiJobSubmitResponse SlurmV0043PostJobSubmit(ctx).V0043JobSubmitReq(v0043JobSubmitReq).Execute()

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
	v0043JobSubmitReq := *openapiclient.NewV0043JobSubmitReq() // V0043JobSubmitReq | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostJobSubmit(context.Background()).V0043JobSubmitReq(v0043JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostJobSubmit`: V0043OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043JobSubmitReq** | [**V0043JobSubmitReq**](V0043JobSubmitReq.md) | Job description | 

### Return type

[**V0043OpenapiJobSubmitResponse**](V0043OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostNode

> V0043OpenapiResp SlurmV0043PostNode(ctx, nodeName).V0043UpdateNodeMsg(v0043UpdateNodeMsg).Execute()

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
	v0043UpdateNodeMsg := *openapiclient.NewV0043UpdateNodeMsg() // V0043UpdateNodeMsg | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostNode(context.Background(), nodeName).V0043UpdateNodeMsg(v0043UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostNode`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0043UpdateNodeMsg** | [**V0043UpdateNodeMsg**](V0043UpdateNodeMsg.md) | Node update description | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostNodes

> V0043OpenapiResp SlurmV0043PostNodes(ctx).V0043UpdateNodeMsg(v0043UpdateNodeMsg).Execute()

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
	v0043UpdateNodeMsg := *openapiclient.NewV0043UpdateNodeMsg() // V0043UpdateNodeMsg | Nodelist update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostNodes(context.Background()).V0043UpdateNodeMsg(v0043UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostNodes`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043UpdateNodeMsg** | [**V0043UpdateNodeMsg**](V0043UpdateNodeMsg.md) | Nodelist update description | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostReservation

> V0043OpenapiReservationModResp SlurmV0043PostReservation(ctx).V0043ReservationDescMsg(v0043ReservationDescMsg).Execute()

create or update a reservation

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
	v0043ReservationDescMsg := *openapiclient.NewV0043ReservationDescMsg() // V0043ReservationDescMsg | reservation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostReservation(context.Background()).V0043ReservationDescMsg(v0043ReservationDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostReservation`: V0043OpenapiReservationModResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostReservation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043ReservationDescMsg** | [**V0043ReservationDescMsg**](V0043ReservationDescMsg.md) | reservation description | 

### Return type

[**V0043OpenapiReservationModResp**](V0043OpenapiReservationModResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0043PostReservations

> V0043OpenapiReservationModResp SlurmV0043PostReservations(ctx).V0043ReservationModReq(v0043ReservationModReq).Execute()

create or update reservations

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
	v0043ReservationModReq := *openapiclient.NewV0043ReservationModReq() // V0043ReservationModReq | reservation descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0043PostReservations(context.Background()).V0043ReservationModReq(v0043ReservationModReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0043PostReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0043PostReservations`: V0043OpenapiReservationModResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0043PostReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0043PostReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043ReservationModReq** | [**V0043ReservationModReq**](V0043ReservationModReq.md) | reservation descriptions | 

### Return type

[**V0043OpenapiReservationModResp**](V0043OpenapiReservationModResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044DeleteJob

> V0044OpenapiKillJobResp SlurmV0044DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044DeleteJob`: V0044OpenapiKillJobResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0044OpenapiKillJobResp**](V0044OpenapiKillJobResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044DeleteJobs

> V0044OpenapiKillJobsResp SlurmV0044DeleteJobs(ctx).V0044KillJobsMsg(v0044KillJobsMsg).Execute()

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
	v0044KillJobsMsg := *openapiclient.NewV0044KillJobsMsg() // V0044KillJobsMsg | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044DeleteJobs(context.Background()).V0044KillJobsMsg(v0044KillJobsMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044DeleteJobs`: V0044OpenapiKillJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044KillJobsMsg** | [**V0044KillJobsMsg**](V0044KillJobsMsg.md) | Signal or cancel jobs | 

### Return type

[**V0044OpenapiKillJobsResp**](V0044OpenapiKillJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044DeleteNode

> V0044OpenapiResp SlurmV0044DeleteNode(ctx, nodeName).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044DeleteNode`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044DeleteReservation

> V0044OpenapiResp SlurmV0044DeleteReservation(ctx, reservationName).Execute()

delete a reservation

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044DeleteReservation(context.Background(), reservationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044DeleteReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044DeleteReservation`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044DeleteReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044DeleteReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetDiag

> V0044OpenapiDiagResp SlurmV0044GetDiag(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetDiag`: V0044OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetDiagRequest struct via the builder pattern


### Return type

[**V0044OpenapiDiagResp**](V0044OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetJob

> V0044OpenapiJobInfoResp SlurmV0044GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetJob`: V0044OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiJobInfoResp**](V0044OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetJobs

> V0044OpenapiJobInfoResp SlurmV0044GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetJobs`: V0044OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiJobInfoResp**](V0044OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetJobsState

> V0044OpenapiJobInfoResp SlurmV0044GetJobsState(ctx).JobId(jobId).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetJobsState(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetJobsState`: V0044OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | CSV list of Job IDs to search for | 

### Return type

[**V0044OpenapiJobInfoResp**](V0044OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetLicenses

> V0044OpenapiLicensesResp SlurmV0044GetLicenses(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetLicenses`: V0044OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetLicensesRequest struct via the builder pattern


### Return type

[**V0044OpenapiLicensesResp**](V0044OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetNode

> V0044OpenapiNodesResp SlurmV0044GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetNode`: V0044OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiNodesResp**](V0044OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetNodes

> V0044OpenapiNodesResp SlurmV0044GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetNodes`: V0044OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query jobs updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiNodesResp**](V0044OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetPartition

> V0044OpenapiPartitionResp SlurmV0044GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetPartition`: V0044OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiPartitionResp**](V0044OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetPartitions

> V0044OpenapiPartitionResp SlurmV0044GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetPartitions`: V0044OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 

### Return type

[**V0044OpenapiPartitionResp**](V0044OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetPing

> V0044OpenapiPingArrayResp SlurmV0044GetPing(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetPing`: V0044OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetPingRequest struct via the builder pattern


### Return type

[**V0044OpenapiPingArrayResp**](V0044OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetReconfigure

> V0044OpenapiResp SlurmV0044GetReconfigure(ctx).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetReconfigure`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetReconfigureRequest struct via the builder pattern


### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetReservation

> V0044OpenapiReservationResp SlurmV0044GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetReservation`: V0044OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0044OpenapiReservationResp**](V0044OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetReservations

> V0044OpenapiReservationResp SlurmV0044GetReservations(ctx).UpdateTime(updateTime).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetReservations`: V0044OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0044OpenapiReservationResp**](V0044OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetResources

> V0044OpenapiResourceLayoutResp SlurmV0044GetResources(ctx, jobId).Execute()

get resource layout info

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetResources(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetResources`: V0044OpenapiResourceLayoutResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiResourceLayoutResp**](V0044OpenapiResourceLayoutResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044GetShares

> V0044OpenapiSharesResp SlurmV0044GetShares(ctx).Accounts(accounts).Users(users).Execute()

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0044GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044GetShares`: V0044OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0044OpenapiSharesResp**](V0044OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostJob

> V0044OpenapiJobPostResponse SlurmV0044PostJob(ctx, jobId).V0044JobDescMsg(v0044JobDescMsg).Execute()

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
	v0044JobDescMsg := *openapiclient.NewV0044JobDescMsg() // V0044JobDescMsg | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostJob(context.Background(), jobId).V0044JobDescMsg(v0044JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostJob`: V0044OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0044JobDescMsg** | [**V0044JobDescMsg**](V0044JobDescMsg.md) | Job update description | 

### Return type

[**V0044OpenapiJobPostResponse**](V0044OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostJobAllocate

> V0044OpenapiJobAllocResp SlurmV0044PostJobAllocate(ctx).V0044JobAllocReq(v0044JobAllocReq).Execute()

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
	v0044JobAllocReq := *openapiclient.NewV0044JobAllocReq() // V0044JobAllocReq | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostJobAllocate(context.Background()).V0044JobAllocReq(v0044JobAllocReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostJobAllocate`: V0044OpenapiJobAllocResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044JobAllocReq** | [**V0044JobAllocReq**](V0044JobAllocReq.md) | Job allocation description | 

### Return type

[**V0044OpenapiJobAllocResp**](V0044OpenapiJobAllocResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostJobSubmit

> V0044OpenapiJobSubmitResponse SlurmV0044PostJobSubmit(ctx).V0044JobSubmitReq(v0044JobSubmitReq).Execute()

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
	v0044JobSubmitReq := *openapiclient.NewV0044JobSubmitReq() // V0044JobSubmitReq | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostJobSubmit(context.Background()).V0044JobSubmitReq(v0044JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostJobSubmit`: V0044OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044JobSubmitReq** | [**V0044JobSubmitReq**](V0044JobSubmitReq.md) | Job description | 

### Return type

[**V0044OpenapiJobSubmitResponse**](V0044OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostNewNode

> V0044OpenapiResp SlurmV0044PostNewNode(ctx).V0044OpenapiCreateNodeReq(v0044OpenapiCreateNodeReq).Execute()

create node

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
	v0044OpenapiCreateNodeReq := *openapiclient.NewV0044OpenapiCreateNodeReq("NodeConf_example") // V0044OpenapiCreateNodeReq | node create request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostNewNode(context.Background()).V0044OpenapiCreateNodeReq(v0044OpenapiCreateNodeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostNewNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostNewNode`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostNewNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostNewNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiCreateNodeReq** | [**V0044OpenapiCreateNodeReq**](V0044OpenapiCreateNodeReq.md) | node create request | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostNode

> V0044OpenapiResp SlurmV0044PostNode(ctx, nodeName).V0044UpdateNodeMsg(v0044UpdateNodeMsg).Execute()

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
	v0044UpdateNodeMsg := *openapiclient.NewV0044UpdateNodeMsg() // V0044UpdateNodeMsg | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostNode(context.Background(), nodeName).V0044UpdateNodeMsg(v0044UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostNode`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0044UpdateNodeMsg** | [**V0044UpdateNodeMsg**](V0044UpdateNodeMsg.md) | Node update description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostNodes

> V0044OpenapiResp SlurmV0044PostNodes(ctx).V0044UpdateNodeMsg(v0044UpdateNodeMsg).Execute()

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
	v0044UpdateNodeMsg := *openapiclient.NewV0044UpdateNodeMsg() // V0044UpdateNodeMsg | Nodelist update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostNodes(context.Background()).V0044UpdateNodeMsg(v0044UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostNodes`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044UpdateNodeMsg** | [**V0044UpdateNodeMsg**](V0044UpdateNodeMsg.md) | Nodelist update description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostReservation

> V0044OpenapiReservationModResp SlurmV0044PostReservation(ctx).V0044ReservationDescMsg(v0044ReservationDescMsg).Execute()

create or update a reservation

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
	v0044ReservationDescMsg := *openapiclient.NewV0044ReservationDescMsg() // V0044ReservationDescMsg | reservation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostReservation(context.Background()).V0044ReservationDescMsg(v0044ReservationDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostReservation`: V0044OpenapiReservationModResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostReservation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044ReservationDescMsg** | [**V0044ReservationDescMsg**](V0044ReservationDescMsg.md) | reservation description | 

### Return type

[**V0044OpenapiReservationModResp**](V0044OpenapiReservationModResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0044PostReservations

> V0044OpenapiReservationModResp SlurmV0044PostReservations(ctx).V0044ReservationModReq(v0044ReservationModReq).Execute()

create or update reservations

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
	v0044ReservationModReq := *openapiclient.NewV0044ReservationModReq() // V0044ReservationModReq | reservation descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0044PostReservations(context.Background()).V0044ReservationModReq(v0044ReservationModReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0044PostReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0044PostReservations`: V0044OpenapiReservationModResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0044PostReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0044PostReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044ReservationModReq** | [**V0044ReservationModReq**](V0044ReservationModReq.md) | reservation descriptions | 

### Return type

[**V0044OpenapiReservationModResp**](V0044OpenapiReservationModResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

