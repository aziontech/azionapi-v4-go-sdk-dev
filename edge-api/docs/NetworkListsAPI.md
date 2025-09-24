# \NetworkListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkList**](NetworkListsAPI.md#CreateNetworkList) | **Post** /workspace/network_lists | Create a Network List
[**DestroyNetworkList**](NetworkListsAPI.md#DestroyNetworkList) | **Delete** /workspace/network_lists/{id} | Destroy a Network List
[**ListNetworkLists**](NetworkListsAPI.md#ListNetworkLists) | **Get** /workspace/network_lists | List Network Lists
[**PartialUpdateNetworkList**](NetworkListsAPI.md#PartialUpdateNetworkList) | **Patch** /workspace/network_lists/{id} | Partially update a Network List
[**RetrieveNetworkList**](NetworkListsAPI.md#RetrieveNetworkList) | **Get** /workspace/network_lists/{id} | Retrieve details of a Network List
[**UpdateNetworkList**](NetworkListsAPI.md#UpdateNetworkList) | **Put** /workspace/network_lists/{id} | Update a Network List



## CreateNetworkList

> ResponseNetworkListDetail CreateNetworkList(ctx).NetworkListDetailRequest(networkListDetailRequest).Execute()

Create a Network List



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
	networkListDetailRequest := *openapiclient.NewNetworkListDetailRequest("Name_example", "Type_example", []string{"Items_example"}) // NetworkListDetailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.CreateNetworkList(context.Background()).NetworkListDetailRequest(networkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.CreateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.CreateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkListDetailRequest** | [**NetworkListDetailRequest**](NetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyNetworkList

> ResponseDeleteNetworkListDetail DestroyNetworkList(ctx, id).Execute()

Destroy a Network List



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.DestroyNetworkList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.DestroyNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyNetworkList`: ResponseDeleteNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.DestroyNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteNetworkListDetail**](ResponseDeleteNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkLists

> PaginatedNetworkListList ListNetworkLists(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Network Lists



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, type, last_editor, last_modified, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.ListNetworkLists(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.ListNetworkLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkLists`: PaginatedNetworkListList
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.ListNetworkLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, type, last_editor, last_modified, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedNetworkListList**](PaginatedNetworkListList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateNetworkList

> ResponseNetworkListDetail PartialUpdateNetworkList(ctx, id).PatchedNetworkListDetailRequest(patchedNetworkListDetailRequest).Execute()

Partially update a Network List



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
	id := "id_example" // string | 
	patchedNetworkListDetailRequest := *openapiclient.NewPatchedNetworkListDetailRequest() // PatchedNetworkListDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.PartialUpdateNetworkList(context.Background(), id).PatchedNetworkListDetailRequest(patchedNetworkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.PartialUpdateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.PartialUpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNetworkListDetailRequest** | [**PatchedNetworkListDetailRequest**](PatchedNetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNetworkList

> ResponseRetrieveNetworkListDetail RetrieveNetworkList(ctx, id).Fields(fields).Ipv4(ipv4).Ipv6(ipv6).Execute()

Retrieve details of a Network List



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
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ipv4 := true // bool | Filter by IPv4. Only applicable for network lists of type 'ip_cidr'. (optional)
	ipv6 := true // bool | Filter by IPv6. Only applicable for network lists of type 'ip_cidr'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.RetrieveNetworkList(context.Background(), id).Fields(fields).Ipv4(ipv4).Ipv6(ipv6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.RetrieveNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNetworkList`: ResponseRetrieveNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.RetrieveNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ipv4** | **bool** | Filter by IPv4. Only applicable for network lists of type &#39;ip_cidr&#39;. | 
 **ipv6** | **bool** | Filter by IPv6. Only applicable for network lists of type &#39;ip_cidr&#39;. | 

### Return type

[**ResponseRetrieveNetworkListDetail**](ResponseRetrieveNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkList

> ResponseNetworkListDetail UpdateNetworkList(ctx, id).NetworkListDetailRequest(networkListDetailRequest).Execute()

Update a Network List



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
	id := "id_example" // string | 
	networkListDetailRequest := *openapiclient.NewNetworkListDetailRequest("Name_example", "Type_example", []string{"Items_example"}) // NetworkListDetailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.UpdateNetworkList(context.Background(), id).NetworkListDetailRequest(networkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.UpdateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.UpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkListDetailRequest** | [**NetworkListDetailRequest**](NetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

