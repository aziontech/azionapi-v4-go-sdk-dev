# \MetricsReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReport**](MetricsReportsAPI.md#CreateReport) | **Post** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports | Create a new report
[**DeleteReport**](MetricsReportsAPI.md#DeleteReport) | **Delete** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports/{reportId} | Delete a report
[**ListReports**](MetricsReportsAPI.md#ListReports) | **Get** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports | List of reports
[**OrderingReport**](MetricsReportsAPI.md#OrderingReport) | **Put** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports/order | Ordering reports in row
[**RetrieveReport**](MetricsReportsAPI.md#RetrieveReport) | **Get** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports/{reportId} | Retrieve details from a report
[**UpdateReport**](MetricsReportsAPI.md#UpdateReport) | **Put** /metrics/folders/{folderId}/dashboards/{dashboardId}/rows/{rowId}/reports/{reportId} | Update a report



## CreateReport

> ResponseReport CreateReport(ctx, dashboardId, folderId, rowId).ReportRequest(reportRequest).Execute()

Create a new report



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	rowId := "rowId_example" // string | 
	reportRequest := *openapiclient.NewReportRequest("Description_example", "Type_example", "AggregationType_example", "DataUnit_example", []openapiclient.BaseQueryRequest{*openapiclient.NewBaseQueryRequest("Dataset_example", int64(123), "OrderDirection_example")}, int64(123), "Name_example") // ReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.CreateReport(context.Background(), dashboardId, folderId, rowId).ReportRequest(reportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.CreateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReport`: ResponseReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.CreateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reportRequest** | [**ReportRequest**](ReportRequest.md) |  | 

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


## DeleteReport

> ResponseDeleteReport DeleteReport(ctx, dashboardId, folderId, reportId, rowId).Execute()

Delete a report



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	reportId := "reportId_example" // string | 
	rowId := "rowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.DeleteReport(context.Background(), dashboardId, folderId, reportId, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.DeleteReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReport`: ResponseDeleteReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.DeleteReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**reportId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportRequest struct via the builder pattern


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


## ListReports

> PaginatedResponseListReportList ListReports(ctx, dashboardId, folderId, rowId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of reports



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	rowId := "rowId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.ListReports(context.Background(), dashboardId, folderId, rowId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.ListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReports`: PaginatedResponseListReportList
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.ListReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


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


## OrderingReport

> ResponseOrder OrderingReport(ctx, dashboardId, folderId, rowId).OrderRequest(orderRequest).Execute()

Ordering reports in row



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	rowId := "rowId_example" // string | 
	orderRequest := *openapiclient.NewOrderRequest([]int64{int64(123)}) // OrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.OrderingReport(context.Background(), dashboardId, folderId, rowId).OrderRequest(orderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.OrderingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderingReport`: ResponseOrder
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.OrderingReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderingReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orderRequest** | [**OrderRequest**](OrderRequest.md) |  | 

### Return type

[**ResponseOrder**](ResponseOrder.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveReport

> ResponseRetrieveReport RetrieveReport(ctx, dashboardId, folderId, reportId, rowId).Fields(fields).Execute()

Retrieve details from a report



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	reportId := "reportId_example" // string | 
	rowId := "rowId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.RetrieveReport(context.Background(), dashboardId, folderId, reportId, rowId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.RetrieveReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveReport`: ResponseRetrieveReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.RetrieveReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**reportId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveReportRequest struct via the builder pattern


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


## UpdateReport

> ResponseReport UpdateReport(ctx, dashboardId, folderId, reportId, rowId).ReportRequest(reportRequest).Execute()

Update a report



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
	dashboardId := "dashboardId_example" // string | 
	folderId := "folderId_example" // string | 
	reportId := "reportId_example" // string | 
	rowId := "rowId_example" // string | 
	reportRequest := *openapiclient.NewReportRequest("Description_example", "Type_example", "AggregationType_example", "DataUnit_example", []openapiclient.BaseQueryRequest{*openapiclient.NewBaseQueryRequest("Dataset_example", int64(123), "OrderDirection_example")}, int64(123), "Name_example") // ReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsReportsAPI.UpdateReport(context.Background(), dashboardId, folderId, reportId, rowId).ReportRequest(reportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsReportsAPI.UpdateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReport`: ResponseReport
	fmt.Fprintf(os.Stdout, "Response from `MetricsReportsAPI.UpdateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** |  | 
**folderId** | **string** |  | 
**reportId** | **string** |  | 
**rowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **reportRequest** | [**ReportRequest**](ReportRequest.md) |  | 

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

