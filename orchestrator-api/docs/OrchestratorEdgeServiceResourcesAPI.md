# \OrchestratorEdgeServiceResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResource**](OrchestratorEdgeServiceResourcesAPI.md#CreateResource) | **Post** /edge_orchestrator/edge_services/{serviceId}/resources | Create Service Resource
[**DeleteResource**](OrchestratorEdgeServiceResourcesAPI.md#DeleteResource) | **Delete** /edge_orchestrator/edge_services/{serviceId}/resources/{resourceId} | Delete Resource
[**ListResourcesOfAService**](OrchestratorEdgeServiceResourcesAPI.md#ListResourcesOfAService) | **Get** /edge_orchestrator/edge_services/{serviceId}/resources | List Service Resources
[**RetrieveResource**](OrchestratorEdgeServiceResourcesAPI.md#RetrieveResource) | **Get** /edge_orchestrator/edge_services/{serviceId}/resources/{resourceId} | Retrieve details of a Resource
[**RetrieveResourceContent**](OrchestratorEdgeServiceResourcesAPI.md#RetrieveResourceContent) | **Get** /edge_orchestrator/edge_services/{serviceId}/resources/{resourceId}/content | Retrieve content of a Resource
[**UpdateResource**](OrchestratorEdgeServiceResourcesAPI.md#UpdateResource) | **Put** /edge_orchestrator/edge_services/{serviceId}/resources/{resourceId} | Update Resource
[**UploadResourceContent**](OrchestratorEdgeServiceResourcesAPI.md#UploadResourceContent) | **Put** /edge_orchestrator/edge_services/{serviceId}/resources/{resourceId}/content | Upload content of a Resource



## CreateResource

> ServiceResource CreateResource(ctx, serviceId).ServiceResourceRequest(serviceResourceRequest).Execute()

Create Service Resource



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
	serviceId := int64(789) // int64 | 
	serviceResourceRequest := *openapiclient.NewServiceResourceRequest("Name_example", "ContentType_example", "FileGroup_example", "FileMode_example", "FileOwner_example") // ServiceResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.CreateResource(context.Background(), serviceId).ServiceResourceRequest(serviceResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.CreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResource`: ServiceResource
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.CreateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceResourceRequest** | [**ServiceResourceRequest**](ServiceResourceRequest.md) |  | 

### Return type

[**ServiceResource**](ServiceResource.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResource

> DeleteResource(ctx, resourceId, serviceId).Execute()

Delete Resource



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
	resourceId := int64(789) // int64 | 
	serviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.DeleteResource(context.Background(), resourceId, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.DeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int64** |  | 
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourcesOfAService

> PaginatedServiceResourceList ListResourcesOfAService(ctx, serviceId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Service Resources



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
	serviceId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.ListResourcesOfAService(context.Background(), serviceId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.ListResourcesOfAService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourcesOfAService`: PaginatedServiceResourceList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.ListResourcesOfAService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesOfAServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedServiceResourceList**](PaginatedServiceResourceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveResource

> ServiceResourceId RetrieveResource(ctx, resourceId, serviceId).Fields(fields).Execute()

Retrieve details of a Resource



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
	resourceId := int64(789) // int64 | 
	serviceId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.RetrieveResource(context.Background(), resourceId, serviceId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.RetrieveResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveResource`: ServiceResourceId
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.RetrieveResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int64** |  | 
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ServiceResourceId**](ServiceResourceId.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveResourceContent

> Content RetrieveResourceContent(ctx, resourceId, serviceId).Fields(fields).Execute()

Retrieve content of a Resource



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
	resourceId := int64(789) // int64 | 
	serviceId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.RetrieveResourceContent(context.Background(), resourceId, serviceId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.RetrieveResourceContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveResourceContent`: Content
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.RetrieveResourceContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int64** |  | 
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveResourceContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**Content**](Content.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResource

> ServiceResourceId UpdateResource(ctx, resourceId, serviceId).ServiceResourceIdRequest(serviceResourceIdRequest).Execute()

Update Resource



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
	resourceId := int64(789) // int64 | 
	serviceId := int64(789) // int64 | 
	serviceResourceIdRequest := *openapiclient.NewServiceResourceIdRequest("Name_example", "FileGroup_example", "FileMode_example", "FileOwner_example") // ServiceResourceIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.UpdateResource(context.Background(), resourceId, serviceId).ServiceResourceIdRequest(serviceResourceIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.UpdateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResource`: ServiceResourceId
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.UpdateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int64** |  | 
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceResourceIdRequest** | [**ServiceResourceIdRequest**](ServiceResourceIdRequest.md) |  | 

### Return type

[**ServiceResourceId**](ServiceResourceId.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadResourceContent

> Content UploadResourceContent(ctx, resourceId, serviceId).ContentRequest(contentRequest).Execute()

Upload content of a Resource



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
	resourceId := int64(789) // int64 | 
	serviceId := int64(789) // int64 | 
	contentRequest := *openapiclient.NewContentRequest("Name_example", "ContentType_example", "ContentHash_example") // ContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServiceResourcesAPI.UploadResourceContent(context.Background(), resourceId, serviceId).ContentRequest(contentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServiceResourcesAPI.UploadResourceContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadResourceContent`: Content
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServiceResourcesAPI.UploadResourceContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int64** |  | 
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadResourceContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentRequest** | [**ContentRequest**](ContentRequest.md) |  | 

### Return type

[**Content**](Content.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

