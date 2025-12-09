# \MetricsRecommendationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecommendation**](MetricsRecommendationsAPI.md#CreateRecommendation) | **Post** /metrics/recommendations | Create a new recommendation
[**DeleteRecommendation**](MetricsRecommendationsAPI.md#DeleteRecommendation) | **Delete** /metrics/recommendations/{recommendationId} | Delete a recommendation
[**ListRecommendations**](MetricsRecommendationsAPI.md#ListRecommendations) | **Get** /metrics/recommendations | List of the recommendations



## CreateRecommendation

> ResponseFolder CreateRecommendation(ctx).RecommendationRequest(recommendationRequest).Execute()

Create a new recommendation



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
	recommendationRequest := *openapiclient.NewRecommendationRequest(int64(123)) // RecommendationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsRecommendationsAPI.CreateRecommendation(context.Background()).RecommendationRequest(recommendationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsRecommendationsAPI.CreateRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecommendation`: ResponseFolder
	fmt.Fprintf(os.Stdout, "Response from `MetricsRecommendationsAPI.CreateRecommendation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendationRequest** | [**RecommendationRequest**](RecommendationRequest.md) |  | 

### Return type

[**ResponseFolder**](ResponseFolder.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecommendation

> DeleteRecommendation(ctx, recommendationId).Execute()

Delete a recommendation



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
	recommendationId := int64(789) // int64 | The unique identifier of the recommendation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsRecommendationsAPI.DeleteRecommendation(context.Background(), recommendationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsRecommendationsAPI.DeleteRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recommendationId** | **int64** | The unique identifier of the recommendation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecommendationRequest struct via the builder pattern


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


## ListRecommendations

> PaginatedResponseListFolderList ListRecommendations(ctx).Dashboard(dashboard).Fields(fields).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of the recommendations



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
	dashboard := int64(789) // int64 | Filter by dashboard ID (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by recommendation ID (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, dashboard) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsRecommendationsAPI.ListRecommendations(context.Background()).Dashboard(dashboard).Fields(fields).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsRecommendationsAPI.ListRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRecommendations`: PaginatedResponseListFolderList
	fmt.Fprintf(os.Stdout, "Response from `MetricsRecommendationsAPI.ListRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | **int64** | Filter by dashboard ID | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by recommendation ID | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, dashboard) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListFolderList**](PaginatedResponseListFolderList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

