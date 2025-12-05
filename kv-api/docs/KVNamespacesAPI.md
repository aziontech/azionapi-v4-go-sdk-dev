# \KVNamespacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceKvNamespacesCreate**](KVNamespacesAPI.md#WorkspaceKvNamespacesCreate) | **Post** /workspace/kv/namespaces | Create namespace
[**WorkspaceKvNamespacesList**](KVNamespacesAPI.md#WorkspaceKvNamespacesList) | **Get** /workspace/kv/namespaces | List namespaces
[**WorkspaceKvNamespacesRetrieve**](KVNamespacesAPI.md#WorkspaceKvNamespacesRetrieve) | **Get** /workspace/kv/namespaces/{namespace} | Retrieve namespace



## WorkspaceKvNamespacesCreate

> NamespaceResponse WorkspaceKvNamespacesCreate(ctx).NamespaceCreateRequest(namespaceCreateRequest).Execute()

Create namespace



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
	namespaceCreateRequest := *openapiclient.NewNamespaceCreateRequest("Name_example") // NamespaceCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVNamespacesAPI.WorkspaceKvNamespacesCreate(context.Background()).NamespaceCreateRequest(namespaceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.WorkspaceKvNamespacesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceKvNamespacesCreate`: NamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.WorkspaceKvNamespacesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceKvNamespacesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceCreateRequest** | [**NamespaceCreateRequest**](NamespaceCreateRequest.md) |  | 

### Return type

[**NamespaceResponse**](NamespaceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceKvNamespacesList

> []NamespaceListResponse WorkspaceKvNamespacesList(ctx).Fields(fields).Page(page).PageSize(pageSize).Execute()

List namespaces



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
	fields := "fields_example" // string | Comma-separated list of fields to include in the response (optional)
	page := int64(789) // int64 | Page number for pagination (optional)
	pageSize := int64(789) // int64 | Number of items per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVNamespacesAPI.WorkspaceKvNamespacesList(context.Background()).Fields(fields).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.WorkspaceKvNamespacesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceKvNamespacesList`: []NamespaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.WorkspaceKvNamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceKvNamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of fields to include in the response | 
 **page** | **int64** | Page number for pagination | 
 **pageSize** | **int64** | Number of items per page | 

### Return type

[**[]NamespaceListResponse**](NamespaceListResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceKvNamespacesRetrieve

> NamespaceResponse WorkspaceKvNamespacesRetrieve(ctx, namespace).Fields(fields).Execute()

Retrieve namespace



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
	namespace := "namespace_example" // string | The unique identifier (name) of the namespace
	fields := "fields_example" // string | Comma-separated list of fields to include in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVNamespacesAPI.WorkspaceKvNamespacesRetrieve(context.Background(), namespace).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.WorkspaceKvNamespacesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceKvNamespacesRetrieve`: NamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.WorkspaceKvNamespacesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The unique identifier (name) of the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceKvNamespacesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of fields to include in the response | 

### Return type

[**NamespaceResponse**](NamespaceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

