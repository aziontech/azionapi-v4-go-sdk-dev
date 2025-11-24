# \VCSContinuousDeploymentExecutorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveExecutor**](VCSContinuousDeploymentExecutorAPI.md#RetrieveExecutor) | **Get** /vcs/continuous_deployments/{continuous_deployment_id}/executor | Retrieve details from a executor
[**UpdateExecutor**](VCSContinuousDeploymentExecutorAPI.md#UpdateExecutor) | **Put** /vcs/continuous_deployments/{continuous_deployment_id}/executor | Update a executor



## RetrieveExecutor

> ResponseRetrieveExecutor RetrieveExecutor(ctx, continuousDeploymentId).Fields(fields).Execute()

Retrieve details from a executor



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
	continuousDeploymentId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VCSContinuousDeploymentExecutorAPI.RetrieveExecutor(context.Background(), continuousDeploymentId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VCSContinuousDeploymentExecutorAPI.RetrieveExecutor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveExecutor`: ResponseRetrieveExecutor
	fmt.Fprintf(os.Stdout, "Response from `VCSContinuousDeploymentExecutorAPI.RetrieveExecutor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**continuousDeploymentId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveExecutorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveExecutor**](ResponseRetrieveExecutor.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecutor

> ResponseExecutor UpdateExecutor(ctx, continuousDeploymentId).Execute()

Update a executor



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
	continuousDeploymentId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VCSContinuousDeploymentExecutorAPI.UpdateExecutor(context.Background(), continuousDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VCSContinuousDeploymentExecutorAPI.UpdateExecutor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExecutor`: ResponseExecutor
	fmt.Fprintf(os.Stdout, "Response from `VCSContinuousDeploymentExecutorAPI.UpdateExecutor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**continuousDeploymentId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExecutorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseExecutor**](ResponseExecutor.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

