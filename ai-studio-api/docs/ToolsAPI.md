# \ToolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToolsCreate**](ToolsAPI.md#ToolsCreate) | **Post** /v4/workspace/ai/tools | 
[**ToolsDestroy**](ToolsAPI.md#ToolsDestroy) | **Delete** /v4/workspace/ai/tools/{toolId} | 
[**ToolsKbCreate**](ToolsAPI.md#ToolsKbCreate) | **Post** /v4/workspace/ai/tools/{toolId}/kb | 
[**ToolsKbDestroy**](ToolsAPI.md#ToolsKbDestroy) | **Delete** /v4/workspace/ai/tools/{toolId}/kb/{kbId} | 
[**ToolsKbRetrieve**](ToolsAPI.md#ToolsKbRetrieve) | **Get** /v4/workspace/ai/tools/{toolId}/kb | 
[**ToolsList**](ToolsAPI.md#ToolsList) | **Get** /v4/workspace/ai/tools | 
[**ToolsPartialUpdate**](ToolsAPI.md#ToolsPartialUpdate) | **Patch** /v4/workspace/ai/tools/{toolId} | 
[**ToolsRetrieve**](ToolsAPI.md#ToolsRetrieve) | **Get** /v4/workspace/ai/tools/{toolId} | 
[**ToolsUpdate**](ToolsAPI.md#ToolsUpdate) | **Put** /v4/workspace/ai/tools/{toolId} | 



## ToolsCreate

> Tool ToolsCreate(ctx).ToolRequest(toolRequest).Execute()



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
	toolRequest := *openapiclient.NewToolRequest("Name_example", "Type_example") // ToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsCreate(context.Background()).ToolRequest(toolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsCreate`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolRequest** | [**ToolRequest**](ToolRequest.md) |  | 

### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsDestroy

> ToolsDestroy(ctx, toolId).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.ToolsDestroy(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsKbCreate

> Tool ToolsKbCreate(ctx, toolId).ToolRequest(toolRequest).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.
	toolRequest := *openapiclient.NewToolRequest("Name_example", "Type_example") // ToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsKbCreate(context.Background(), toolId).ToolRequest(toolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsKbCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsKbCreate`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsKbCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsKbCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolRequest** | [**ToolRequest**](ToolRequest.md) |  | 

### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsKbDestroy

> ToolsKbDestroy(ctx, kbId, toolId).Execute()



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
	kbId := "kbId_example" // string | 
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.ToolsKbDestroy(context.Background(), kbId, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsKbDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** |  | 
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsKbDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsKbRetrieve

> Tool ToolsKbRetrieve(ctx, toolId).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsKbRetrieve(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsKbRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsKbRetrieve`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsKbRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsKbRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsList

> PaginatedToolList ToolsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsList`: PaginatedToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedToolList**](PaginatedToolList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsPartialUpdate

> Tool ToolsPartialUpdate(ctx, toolId).PatchedToolRequest(patchedToolRequest).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.
	patchedToolRequest := *openapiclient.NewPatchedToolRequest() // PatchedToolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsPartialUpdate(context.Background(), toolId).PatchedToolRequest(patchedToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsPartialUpdate`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedToolRequest** | [**PatchedToolRequest**](PatchedToolRequest.md) |  | 

### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsRetrieve

> Tool ToolsRetrieve(ctx, toolId).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsRetrieve(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsRetrieve`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsUpdate

> Tool ToolsUpdate(ctx, toolId).ToolRequest(toolRequest).Execute()



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
	toolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tool.
	toolRequest := *openapiclient.NewToolRequest("Name_example", "Type_example") // ToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsUpdate(context.Background(), toolId).ToolRequest(toolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsUpdate`: Tool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | A UUID string identifying this tool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolRequest** | [**ToolRequest**](ToolRequest.md) |  | 

### Return type

[**Tool**](Tool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

