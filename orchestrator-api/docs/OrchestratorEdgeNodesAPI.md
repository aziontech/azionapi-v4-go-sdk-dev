# \OrchestratorEdgeNodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEdgeNode**](OrchestratorEdgeNodesAPI.md#DeleteEdgeNode) | **Delete** /edge_orchestrator/edge_nodes/{nodeId} | Delete an Edge Node
[**ListEdgeNodes**](OrchestratorEdgeNodesAPI.md#ListEdgeNodes) | **Get** /edge_orchestrator/edge_nodes | List Edge Nodes
[**PartialUpdateEdgeNode**](OrchestratorEdgeNodesAPI.md#PartialUpdateEdgeNode) | **Patch** /edge_orchestrator/edge_nodes/{nodeId} | Partially update an Edge Node
[**PartialUpdateEdgeNode2**](OrchestratorEdgeNodesAPI.md#PartialUpdateEdgeNode2) | **Patch** /edge_orchestrator/edge_nodes/{nodeId}/ | Partially update an Edge Node
[**RetrieveEdgeNode**](OrchestratorEdgeNodesAPI.md#RetrieveEdgeNode) | **Get** /edge_orchestrator/edge_nodes/{nodeId} | Retrieve details of an Edge Node
[**RetrieveEdgeNode2**](OrchestratorEdgeNodesAPI.md#RetrieveEdgeNode2) | **Get** /edge_orchestrator/edge_nodes/{nodeId}/ | Retrieve details of an Edge Node
[**UpdateEdgeNode**](OrchestratorEdgeNodesAPI.md#UpdateEdgeNode) | **Put** /edge_orchestrator/edge_nodes/{nodeId} | Update an Edge Node
[**UpdateEdgeNode2**](OrchestratorEdgeNodesAPI.md#UpdateEdgeNode2) | **Put** /edge_orchestrator/edge_nodes/{nodeId}/ | Update an Edge Node



## DeleteEdgeNode

> ResponseAsyncDeleteNodes DeleteEdgeNode(ctx, nodeId).Execute()

Delete an Edge Node



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
	nodeId := "nodeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.DeleteEdgeNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.DeleteEdgeNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEdgeNode`: ResponseAsyncDeleteNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.DeleteEdgeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEdgeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteNodes**](ResponseAsyncDeleteNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeNodes

> PaginatedResponseListNodesList ListEdgeNodes(ctx).Fields(fields).HashId(hashId).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Nodes



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	hashId := "hashId_example" // string | Search by hash_id (optional)
	name := "name_example" // string | Search by name (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.ListEdgeNodes(context.Background()).Fields(fields).HashId(hashId).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.ListEdgeNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeNodes`: PaginatedResponseListNodesList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.ListEdgeNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **hashId** | **string** | Search by hash_id | 
 **name** | **string** | Search by name | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListNodesList**](PaginatedResponseListNodesList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeNode

> ResponseAsyncNodes PartialUpdateEdgeNode(ctx, nodeId).PatchedNodesRequest(patchedNodesRequest).Execute()

Partially update an Edge Node



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
	nodeId := "nodeId_example" // string | 
	patchedNodesRequest := *openapiclient.NewPatchedNodesRequest() // PatchedNodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode(context.Background(), nodeId).PatchedNodesRequest(patchedNodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeNode`: ResponseAsyncNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNodesRequest** | [**PatchedNodesRequest**](PatchedNodesRequest.md) |  | 

### Return type

[**ResponseAsyncNodes**](ResponseAsyncNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeNode2

> ResponseAsyncNodes PartialUpdateEdgeNode2(ctx, nodeId).PatchedNodesRequest(patchedNodesRequest).Execute()

Partially update an Edge Node



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
	patchedNodesRequest := *openapiclient.NewPatchedNodesRequest() // PatchedNodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode2(context.Background(), nodeId).PatchedNodesRequest(patchedNodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeNode2`: ResponseAsyncNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.PartialUpdateEdgeNode2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeNode2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNodesRequest** | [**PatchedNodesRequest**](PatchedNodesRequest.md) |  | 

### Return type

[**ResponseAsyncNodes**](ResponseAsyncNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeNode

> ResponseRetrieveNodes RetrieveEdgeNode(ctx, nodeId).Fields(fields).Execute()

Retrieve details of an Edge Node



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
	nodeId := "nodeId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.RetrieveEdgeNode(context.Background(), nodeId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.RetrieveEdgeNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeNode`: ResponseRetrieveNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.RetrieveEdgeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveNodes**](ResponseRetrieveNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeNode2

> ResponseRetrieveNodes RetrieveEdgeNode2(ctx, nodeId).Fields(fields).Execute()

Retrieve details of an Edge Node



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.RetrieveEdgeNode2(context.Background(), nodeId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.RetrieveEdgeNode2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeNode2`: ResponseRetrieveNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.RetrieveEdgeNode2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeNode2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveNodes**](ResponseRetrieveNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeNode

> ResponseAsyncNodes UpdateEdgeNode(ctx, nodeId).NodesRequest(nodesRequest).Execute()

Update an Edge Node



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
	nodeId := "nodeId_example" // string | 
	nodesRequest := *openapiclient.NewNodesRequest("Name_example", "Status_example") // NodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.UpdateEdgeNode(context.Background(), nodeId).NodesRequest(nodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.UpdateEdgeNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeNode`: ResponseAsyncNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.UpdateEdgeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodesRequest** | [**NodesRequest**](NodesRequest.md) |  | 

### Return type

[**ResponseAsyncNodes**](ResponseAsyncNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeNode2

> ResponseAsyncNodes UpdateEdgeNode2(ctx, nodeId).NodesRequest(nodesRequest).Execute()

Update an Edge Node



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
	nodesRequest := *openapiclient.NewNodesRequest("Name_example", "Status_example") // NodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodesAPI.UpdateEdgeNode2(context.Background(), nodeId).NodesRequest(nodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodesAPI.UpdateEdgeNode2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeNode2`: ResponseAsyncNodes
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodesAPI.UpdateEdgeNode2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeNode2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodesRequest** | [**NodesRequest**](NodesRequest.md) |  | 

### Return type

[**ResponseAsyncNodes**](ResponseAsyncNodes.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

