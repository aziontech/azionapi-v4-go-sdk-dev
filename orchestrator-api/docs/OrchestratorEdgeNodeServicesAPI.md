# \OrchestratorEdgeNodeServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindEdgeNodeServices**](OrchestratorEdgeNodeServicesAPI.md#BindEdgeNodeServices) | **Post** /edge_orchestrator/edge_nodes/{nodeId}/services | Bind Node Service
[**ListEdgeNodeServices**](OrchestratorEdgeNodeServicesAPI.md#ListEdgeNodeServices) | **Get** /edge_orchestrator/edge_nodes/{nodeId}/services | List Node Services
[**RetrieveEdgeNodeServiceBind**](OrchestratorEdgeNodeServicesAPI.md#RetrieveEdgeNodeServiceBind) | **Get** /edge_orchestrator/edge_nodes/{nodeId}/services/{bindId} | Retrieve details of an Edge Node Service Bind
[**UnbindEdgeNodeService**](OrchestratorEdgeNodeServicesAPI.md#UnbindEdgeNodeService) | **Delete** /edge_orchestrator/edge_nodes/{nodeId}/services/{bindId} | Unbind Node Service



## BindEdgeNodeServices

> ResponseAsyncNodeServices BindEdgeNodeServices(ctx, nodeId).NodeServicesRequest(nodeServicesRequest).Execute()

Bind Node Service



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
	nodeId := int64(789) // int64 | 
	nodeServicesRequest := *openapiclient.NewNodeServicesRequest(int64(123)) // NodeServicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeServicesAPI.BindEdgeNodeServices(context.Background(), nodeId).NodeServicesRequest(nodeServicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeServicesAPI.BindEdgeNodeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindEdgeNodeServices`: ResponseAsyncNodeServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeServicesAPI.BindEdgeNodeServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindEdgeNodeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeServicesRequest** | [**NodeServicesRequest**](NodeServicesRequest.md) |  | 

### Return type

[**ResponseAsyncNodeServices**](ResponseAsyncNodeServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeNodeServices

> PaginatedResponseListNodeServicesList ListEdgeNodeServices(ctx, nodeId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Node Services



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
	nodeId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: service_name, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeServicesAPI.ListEdgeNodeServices(context.Background(), nodeId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeServicesAPI.ListEdgeNodeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeNodeServices`: PaginatedResponseListNodeServicesList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeServicesAPI.ListEdgeNodeServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeNodeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: service_name, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListNodeServicesList**](PaginatedResponseListNodeServicesList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeNodeServiceBind

> ResponseRetrieveNodeServiceBind RetrieveEdgeNodeServiceBind(ctx, bindId, nodeId).Fields(fields).Execute()

Retrieve details of an Edge Node Service Bind



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
	bindId := int64(789) // int64 | 
	nodeId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeServicesAPI.RetrieveEdgeNodeServiceBind(context.Background(), bindId, nodeId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeServicesAPI.RetrieveEdgeNodeServiceBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeNodeServiceBind`: ResponseRetrieveNodeServiceBind
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeServicesAPI.RetrieveEdgeNodeServiceBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindId** | **int64** |  | 
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeNodeServiceBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveNodeServiceBind**](ResponseRetrieveNodeServiceBind.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindEdgeNodeService

> ResponseAsyncDeleteNodeServiceBind UnbindEdgeNodeService(ctx, bindId, nodeId).Execute()

Unbind Node Service



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
	bindId := int64(789) // int64 | 
	nodeId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeServicesAPI.UnbindEdgeNodeService(context.Background(), bindId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeServicesAPI.UnbindEdgeNodeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindEdgeNodeService`: ResponseAsyncDeleteNodeServiceBind
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeServicesAPI.UnbindEdgeNodeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindId** | **int64** |  | 
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindEdgeNodeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAsyncDeleteNodeServiceBind**](ResponseAsyncDeleteNodeServiceBind.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

