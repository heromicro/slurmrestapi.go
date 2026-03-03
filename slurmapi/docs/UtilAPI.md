# \UtilAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilV0044PostHostlist**](UtilAPI.md#UtilV0044PostHostlist) | **Post** /util/v0.0.44/hostlist | Convert an array of host names into hostlist expression
[**UtilV0044PostHostnames**](UtilAPI.md#UtilV0044PostHostnames) | **Post** /util/v0.0.44/hostnames | Convert a hostlist expression into array of host names



## UtilV0044PostHostlist

> V0044OpenapiHostlistReqResp UtilV0044PostHostlist(ctx).V0044OpenapiHostnamesReqResp(v0044OpenapiHostnamesReqResp).Execute()

Convert an array of host names into hostlist expression

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
	v0044OpenapiHostnamesReqResp := *openapiclient.NewV0044OpenapiHostnamesReqResp([]string{"Hostnames_example"}) // V0044OpenapiHostnamesReqResp | Array of host names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilAPI.UtilV0044PostHostlist(context.Background()).V0044OpenapiHostnamesReqResp(v0044OpenapiHostnamesReqResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilAPI.UtilV0044PostHostlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UtilV0044PostHostlist`: V0044OpenapiHostlistReqResp
	fmt.Fprintf(os.Stdout, "Response from `UtilAPI.UtilV0044PostHostlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUtilV0044PostHostlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiHostnamesReqResp** | [**V0044OpenapiHostnamesReqResp**](V0044OpenapiHostnamesReqResp.md) | Array of host names | 

### Return type

[**V0044OpenapiHostlistReqResp**](V0044OpenapiHostlistReqResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UtilV0044PostHostnames

> V0044OpenapiHostnamesReqResp UtilV0044PostHostnames(ctx).V0044OpenapiHostlistReqResp(v0044OpenapiHostlistReqResp).Execute()

Convert a hostlist expression into array of host names

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
	v0044OpenapiHostlistReqResp := *openapiclient.NewV0044OpenapiHostlistReqResp("Hostlist_example") // V0044OpenapiHostlistReqResp | Hostlist expression (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilAPI.UtilV0044PostHostnames(context.Background()).V0044OpenapiHostlistReqResp(v0044OpenapiHostlistReqResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilAPI.UtilV0044PostHostnames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UtilV0044PostHostnames`: V0044OpenapiHostnamesReqResp
	fmt.Fprintf(os.Stdout, "Response from `UtilAPI.UtilV0044PostHostnames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUtilV0044PostHostnamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiHostlistReqResp** | [**V0044OpenapiHostlistReqResp**](V0044OpenapiHostlistReqResp.md) | Hostlist expression | 

### Return type

[**V0044OpenapiHostnamesReqResp**](V0044OpenapiHostnamesReqResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

