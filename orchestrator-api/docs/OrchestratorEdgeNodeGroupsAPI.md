# \OrchestratorEdgeNodeGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindEdgeNodeGroup**](OrchestratorEdgeNodeGroupsAPI.md#BindEdgeNodeGroup) | **Post** /edge_orchestrator/edge_nodes/{nodeId}/groups | Bind Node Group
[**CreateEdgeNodeGroup**](OrchestratorEdgeNodeGroupsAPI.md#CreateEdgeNodeGroup) | **Post** /edge_orchestrator/edge_nodes/groups | Create Edge Node Group
[**ListEdgeNodeGroups**](OrchestratorEdgeNodeGroupsAPI.md#ListEdgeNodeGroups) | **Get** /edge_orchestrator/edge_nodes/groups | List Edge Node Groups
[**ListEdgeNodeGroupsByID**](OrchestratorEdgeNodeGroupsAPI.md#ListEdgeNodeGroupsByID) | **Get** /edge_orchestrator/edge_nodes/{nodeId}/groups | List Edge Node Groups by id
[**RemoveEdgeNodeGroup**](OrchestratorEdgeNodeGroupsAPI.md#RemoveEdgeNodeGroup) | **Delete** /edge_orchestrator/edge_nodes/groups/{groupId} | Remove Node Group



## BindEdgeNodeGroup

> ResponseAsyncNodeGroupsById BindEdgeNodeGroup(ctx, nodeId).NodeGroupsByIdRequest(nodeGroupsByIdRequest).Execute()

Bind Node Group



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
	nodeGroupsByIdRequest := *openapiclient.NewNodeGroupsByIdRequest() // NodeGroupsByIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeGroupsAPI.BindEdgeNodeGroup(context.Background(), nodeId).NodeGroupsByIdRequest(nodeGroupsByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeGroupsAPI.BindEdgeNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindEdgeNodeGroup`: ResponseAsyncNodeGroupsById
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeGroupsAPI.BindEdgeNodeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindEdgeNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeGroupsByIdRequest** | [**NodeGroupsByIdRequest**](NodeGroupsByIdRequest.md) |  | 

### Return type

[**ResponseAsyncNodeGroupsById**](ResponseAsyncNodeGroupsById.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEdgeNodeGroup

> NodeGroups CreateEdgeNodeGroup(ctx).NodeGroupsRequest(nodeGroupsRequest).Execute()

Create Edge Node Group



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
	nodeGroupsRequest := *openapiclient.NewNodeGroupsRequest("Name_example") // NodeGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeGroupsAPI.CreateEdgeNodeGroup(context.Background()).NodeGroupsRequest(nodeGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeGroupsAPI.CreateEdgeNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeNodeGroup`: NodeGroups
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeGroupsAPI.CreateEdgeNodeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeGroupsRequest** | [**NodeGroupsRequest**](NodeGroupsRequest.md) |  | 

### Return type

[**NodeGroups**](NodeGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeNodeGroups

> PaginatedNodeGroupsList ListEdgeNodeGroups(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Node Groups



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroups(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeNodeGroups`: PaginatedNodeGroupsList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeNodeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedNodeGroupsList**](PaginatedNodeGroupsList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeNodeGroupsByID

> PaginatedResponseListNodeGroupsByIdList ListEdgeNodeGroupsByID(ctx, nodeId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Node Groups by id



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroupsByID(context.Background(), nodeId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroupsByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeNodeGroupsByID`: PaginatedResponseListNodeGroupsByIdList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeNodeGroupsAPI.ListEdgeNodeGroupsByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeNodeGroupsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListNodeGroupsByIdList**](PaginatedResponseListNodeGroupsByIdList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEdgeNodeGroup

> RemoveEdgeNodeGroup(ctx, groupId).Execute()

Remove Node Group



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
	groupId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrchestratorEdgeNodeGroupsAPI.RemoveEdgeNodeGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeNodeGroupsAPI.RemoveEdgeNodeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEdgeNodeGroupRequest struct via the builder pattern


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

