# \MetricsDashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](MetricsDashboardAPI.md#CreateDashboard) | **Post** /metrics/folders/{folderId}/dashboards | Create a new dashboard
[**DeleteDashboard**](MetricsDashboardAPI.md#DeleteDashboard) | **Delete** /metrics/folders/{folderId}/dashboards/{dashboardId} | Delete a dashboard
[**ListDashboards**](MetricsDashboardAPI.md#ListDashboards) | **Get** /metrics/folders/{folderId}/dashboards | List of the dashboards
[**PartialUpdateDashboard**](MetricsDashboardAPI.md#PartialUpdateDashboard) | **Patch** /metrics/folders/{folderId}/dashboards/{dashboardId} | Partially update a dashboard
[**RetrieveDashboard**](MetricsDashboardAPI.md#RetrieveDashboard) | **Get** /metrics/folders/{folderId}/dashboards/{dashboardId} | Retrieve details from a dashboard
[**UpdateDashboard**](MetricsDashboardAPI.md#UpdateDashboard) | **Put** /metrics/folders/{folderId}/dashboards/{dashboardId} | Update a dashboard



## CreateDashboard

> ResponseDashboard CreateDashboard(ctx, folderId).DashboardRequest(dashboardRequest).Execute()

Create a new dashboard



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
	folderId := int64(789) // int64 | The unique identifier of the folder
	dashboardRequest := *openapiclient.NewDashboardRequest("Name_example", "Scope_example") // DashboardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsDashboardAPI.CreateDashboard(context.Background(), folderId).DashboardRequest(dashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: ResponseDashboard
	fmt.Fprintf(os.Stdout, "Response from `MetricsDashboardAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) |  | 

### Return type

[**ResponseDashboard**](ResponseDashboard.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx, dashboardId, folderId).Execute()

Delete a dashboard



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
	dashboardId := int64(789) // int64 | The unique identifier of the dashboard
	folderId := int64(789) // int64 | The unique identifier of the folder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsDashboardAPI.DeleteDashboard(context.Background(), dashboardId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The unique identifier of the dashboard | 
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## ListDashboards

> PaginatedResponseListDashboardList ListDashboards(ctx, folderId).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of the dashboards



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
	folderId := int64(789) // int64 | The unique identifier of the folder
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := "id_example" // string | Filter by id (accepts comma-separated values). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsDashboardAPI.ListDashboards(context.Background(), folderId).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.ListDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboards`: PaginatedResponseListDashboardList
	fmt.Fprintf(os.Stdout, "Response from `MetricsDashboardAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **string** | Filter by id (accepts comma-separated values). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListDashboardList**](PaginatedResponseListDashboardList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDashboard

> ResponseDashboard PartialUpdateDashboard(ctx, dashboardId, folderId).PatchedDashboardRequest(patchedDashboardRequest).Execute()

Partially update a dashboard



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
	dashboardId := int64(789) // int64 | The unique identifier of the dashboard
	folderId := int64(789) // int64 | The unique identifier of the folder
	patchedDashboardRequest := *openapiclient.NewPatchedDashboardRequest() // PatchedDashboardRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsDashboardAPI.PartialUpdateDashboard(context.Background(), dashboardId, folderId).PatchedDashboardRequest(patchedDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.PartialUpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDashboard`: ResponseDashboard
	fmt.Fprintf(os.Stdout, "Response from `MetricsDashboardAPI.PartialUpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The unique identifier of the dashboard | 
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedDashboardRequest** | [**PatchedDashboardRequest**](PatchedDashboardRequest.md) |  | 

### Return type

[**ResponseDashboard**](ResponseDashboard.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDashboard

> ResponseRetrieveDashboard RetrieveDashboard(ctx, dashboardId, folderId).Fields(fields).Execute()

Retrieve details from a dashboard



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
	dashboardId := int64(789) // int64 | The unique identifier of the dashboard
	folderId := int64(789) // int64 | The unique identifier of the folder
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsDashboardAPI.RetrieveDashboard(context.Background(), dashboardId, folderId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.RetrieveDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDashboard`: ResponseRetrieveDashboard
	fmt.Fprintf(os.Stdout, "Response from `MetricsDashboardAPI.RetrieveDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The unique identifier of the dashboard | 
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveDashboard**](ResponseRetrieveDashboard.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> ResponseDashboard UpdateDashboard(ctx, dashboardId, folderId).DashboardRequest(dashboardRequest).Execute()

Update a dashboard



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
	dashboardId := int64(789) // int64 | The unique identifier of the dashboard
	folderId := int64(789) // int64 | The unique identifier of the folder
	dashboardRequest := *openapiclient.NewDashboardRequest("Name_example", "Scope_example") // DashboardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsDashboardAPI.UpdateDashboard(context.Background(), dashboardId, folderId).DashboardRequest(dashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsDashboardAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: ResponseDashboard
	fmt.Fprintf(os.Stdout, "Response from `MetricsDashboardAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The unique identifier of the dashboard | 
**folderId** | **int64** | The unique identifier of the folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) |  | 

### Return type

[**ResponseDashboard**](ResponseDashboard.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

