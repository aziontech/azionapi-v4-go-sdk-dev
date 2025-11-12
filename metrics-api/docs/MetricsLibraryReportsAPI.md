# \MetricsLibraryReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLibraryReport**](MetricsLibraryReportsAPI.md#CreateLibraryReport) | **Post** /metrics/library/reports | Create a new library report
[**DeleteLibraryReport**](MetricsLibraryReportsAPI.md#DeleteLibraryReport) | **Delete** /metrics/library/reports/{libraryReportId} | Delete a library report
[**ListLibraryReports**](MetricsLibraryReportsAPI.md#ListLibraryReports) | **Get** /metrics/library/reports | List of library reports
[**RetrieveLibraryReport**](MetricsLibraryReportsAPI.md#RetrieveLibraryReport) | **Get** /metrics/library/reports/{libraryReportId} | Retrieve details from a library report
[**UpdateLibraryReport**](MetricsLibraryReportsAPI.md#UpdateLibraryReport) | **Put** /metrics/library/reports/{libraryReportId} | Update a library report



## CreateLibraryReport

> ResponseReport CreateLibraryReport(ctx).LibraryReportRequest(libraryReportRequest).Execute()

Create a new library report



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
	libraryReportRequest := *openapiclient.NewLibraryReportRequest("Description_example", "Type_example", "AggregationType_example", "DataUnit_example", []openapiclient.BaseQueryRequest{*openapiclient.NewBaseQueryRequest("Dataset_example", int64(123), "OrderDirection_example")}, "Name_example", "Scope_example") // LibraryReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsLibraryReportsAPI.CreateLibraryReport(context.Background()).LibraryReportRequest(libraryReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsLibraryReportsAPI.CreateLibraryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLibraryReport`: ResponseReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsLibraryReportsAPI.CreateLibraryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLibraryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryReportRequest** | [**LibraryReportRequest**](LibraryReportRequest.md) |  | 

### Return type

[**ResponseReport**](ResponseReport.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLibraryReport

> ResponseDeleteReport DeleteLibraryReport(ctx, libraryReportId).Execute()

Delete a library report



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
	libraryReportId := "libraryReportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsLibraryReportsAPI.DeleteLibraryReport(context.Background(), libraryReportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsLibraryReportsAPI.DeleteLibraryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLibraryReport`: ResponseDeleteReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsLibraryReportsAPI.DeleteLibraryReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryReportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteReport**](ResponseDeleteReport.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLibraryReports

> PaginatedResponseListReportList ListLibraryReports(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of library reports



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsLibraryReportsAPI.ListLibraryReports(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsLibraryReportsAPI.ListLibraryReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLibraryReports`: PaginatedResponseListReportList
	fmt.Fprintf(os.Stdout, "Response from `MetricsLibraryReportsAPI.ListLibraryReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLibraryReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListReportList**](PaginatedResponseListReportList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveLibraryReport

> ResponseRetrieveReport RetrieveLibraryReport(ctx, libraryReportId).Fields(fields).Execute()

Retrieve details from a library report



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
	libraryReportId := "libraryReportId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsLibraryReportsAPI.RetrieveLibraryReport(context.Background(), libraryReportId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsLibraryReportsAPI.RetrieveLibraryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveLibraryReport`: ResponseRetrieveReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsLibraryReportsAPI.RetrieveLibraryReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryReportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveLibraryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveReport**](ResponseRetrieveReport.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLibraryReport

> ResponseReport UpdateLibraryReport(ctx, libraryReportId).LibraryReportRequest(libraryReportRequest).Execute()

Update a library report



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
	libraryReportId := "libraryReportId_example" // string | 
	libraryReportRequest := *openapiclient.NewLibraryReportRequest("Description_example", "Type_example", "AggregationType_example", "DataUnit_example", []openapiclient.BaseQueryRequest{*openapiclient.NewBaseQueryRequest("Dataset_example", int64(123), "OrderDirection_example")}, "Name_example", "Scope_example") // LibraryReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsLibraryReportsAPI.UpdateLibraryReport(context.Background(), libraryReportId).LibraryReportRequest(libraryReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsLibraryReportsAPI.UpdateLibraryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLibraryReport`: ResponseReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsLibraryReportsAPI.UpdateLibraryReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryReportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLibraryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryReportRequest** | [**LibraryReportRequest**](LibraryReportRequest.md) |  | 

### Return type

[**ResponseReport**](ResponseReport.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

