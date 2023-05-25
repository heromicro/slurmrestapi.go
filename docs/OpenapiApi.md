# \OpenapiApi

All URIs are relative to _http://localhost_

| Method                                             | HTTP request          | Description                    |
| -------------------------------------------------- | --------------------- | ------------------------------ |
| [**OpenapiGet**](OpenapiApi.md#OpenapiGet)         | **Get** /openapi      | Retrieve OpenAPI Specification |
| [**OpenapiJsonGet**](OpenapiApi.md#OpenapiJsonGet) | **Get** /openapi.json | Retrieve OpenAPI Specification |
| [**OpenapiV3Get**](OpenapiApi.md#OpenapiV3Get)     | **Get** /openapi/v3   | Retrieve OpenAPI Specification |
| [**OpenapiYamlGet**](OpenapiApi.md#OpenapiYamlGet) | **Get** /openapi.yaml | Retrieve OpenAPI Specification |

## OpenapiGet

> OpenapiGet(ctx).Execute()

Retrieve OpenAPI Specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HEROMICRO/SLURMRESTAPI.GO"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenapiApi.OpenapiGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiGetRequest struct via the builder pattern

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OpenapiJsonGet

> OpenapiJsonGet(ctx).Execute()

Retrieve OpenAPI Specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HEROMICRO/SLURMRESTAPI.GO"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenapiApi.OpenapiJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiJsonGetRequest struct via the builder pattern

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OpenapiV3Get

> OpenapiV3Get(ctx).Execute()

Retrieve OpenAPI Specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HEROMICRO/SLURMRESTAPI.GO"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenapiApi.OpenapiV3Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiV3Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiV3GetRequest struct via the builder pattern

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OpenapiYamlGet

> OpenapiYamlGet(ctx).Execute()

Retrieve OpenAPI Specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/HEROMICRO/SLURMRESTAPI.GO"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenapiApi.OpenapiYamlGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiYamlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiYamlGetRequest struct via the builder pattern

### Return type

(empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
