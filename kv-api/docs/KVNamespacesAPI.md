# \KVNamespacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageKvNamespacesCreate**](KVNamespacesAPI.md#StorageKvNamespacesCreate) | **Post** /storage/kv/namespaces | Create namespace
[**StorageKvNamespacesList**](KVNamespacesAPI.md#StorageKvNamespacesList) | **Get** /storage/kv/namespaces | List namespaces
[**StorageKvNamespacesRetrieve**](KVNamespacesAPI.md#StorageKvNamespacesRetrieve) | **Get** /storage/kv/namespaces/{namespace} | Retrieve namespace



## StorageKvNamespacesCreate

> NamespaceResponse StorageKvNamespacesCreate(ctx).NamespaceCreateRequest(namespaceCreateRequest).Execute()

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
	resp, r, err := apiClient.KVNamespacesAPI.StorageKvNamespacesCreate(context.Background()).NamespaceCreateRequest(namespaceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.StorageKvNamespacesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageKvNamespacesCreate`: NamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.StorageKvNamespacesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageKvNamespacesCreateRequest struct via the builder pattern


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


## StorageKvNamespacesList

> []NamespaceListResponse StorageKvNamespacesList(ctx).Fields(fields).Page(page).PageSize(pageSize).Execute()

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
	resp, r, err := apiClient.KVNamespacesAPI.StorageKvNamespacesList(context.Background()).Fields(fields).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.StorageKvNamespacesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageKvNamespacesList`: []NamespaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.StorageKvNamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageKvNamespacesListRequest struct via the builder pattern


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


## StorageKvNamespacesRetrieve

> NamespaceResponse StorageKvNamespacesRetrieve(ctx, namespace).Fields(fields).Execute()

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
	resp, r, err := apiClient.KVNamespacesAPI.StorageKvNamespacesRetrieve(context.Background(), namespace).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVNamespacesAPI.StorageKvNamespacesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageKvNamespacesRetrieve`: NamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `KVNamespacesAPI.StorageKvNamespacesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The unique identifier (name) of the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageKvNamespacesRetrieveRequest struct via the builder pattern


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

