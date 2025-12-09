# \StorageBucketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](StorageBucketsAPI.md#CreateBucket) | **Post** /workspace/storage/buckets | Create a new bucket
[**DeleteBucket**](StorageBucketsAPI.md#DeleteBucket) | **Delete** /workspace/storage/buckets/{bucket_name} | Delete a bucket
[**ListBuckets**](StorageBucketsAPI.md#ListBuckets) | **Get** /workspace/storage/buckets | List buckets
[**RetrieveBucket**](StorageBucketsAPI.md#RetrieveBucket) | **Get** /workspace/storage/buckets/{bucket_name} | Retrieve a bucket
[**UpdateBucket**](StorageBucketsAPI.md#UpdateBucket) | **Patch** /workspace/storage/buckets/{bucket_name} | Update bucket info



## CreateBucket

> ResponseBucketCreate CreateBucket(ctx).BucketCreateRequest(bucketCreateRequest).Execute()

Create a new bucket



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
	bucketCreateRequest := *openapiclient.NewBucketCreateRequest("Name_example", "EdgeAccess_example") // BucketCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.CreateBucket(context.Background()).BucketCreateRequest(bucketCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.CreateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBucket`: ResponseBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketCreateRequest** | [**BucketCreateRequest**](BucketCreateRequest.md) |  | 

### Return type

[**ResponseBucketCreate**](ResponseBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> ResponseAsyncDeleteBucketCreate DeleteBucket(ctx, bucketName).Execute()

Delete a bucket



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
	bucketName := "bucketName_example" // string | The name of the bucket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.DeleteBucket(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.DeleteBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucket`: ResponseAsyncDeleteBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.DeleteBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteBucketCreate**](ResponseAsyncDeleteBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> PaginatedBucketList ListBuckets(ctx).EdgeAccess(edgeAccess).EdgeAccessIn(edgeAccessIn).Fields(fields).LastEditorIcontains(lastEditorIcontains).LastModified(lastModified).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List buckets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	edgeAccess := "edgeAccess_example" // string | Filter by edge access (exact match). (optional)
	edgeAccessIn := "edgeAccessIn_example" // string | Filter by multiple edge access values (comma-separated). (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	lastEditorIcontains := "lastEditorIcontains_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	lastModified := time.Now() // time.Time | Filter by last modified date (exact match). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified date (greater than or equal to). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified date (less than or equal to). (optional)
	nameIcontains := "nameIcontains_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.ListBuckets(context.Background()).EdgeAccess(edgeAccess).EdgeAccessIn(edgeAccessIn).Fields(fields).LastEditorIcontains(lastEditorIcontains).LastModified(lastModified).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuckets`: PaginatedBucketList
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeAccess** | **string** | Filter by edge access (exact match). | 
 **edgeAccessIn** | **string** | Filter by multiple edge access values (comma-separated). | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **lastEditorIcontains** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **lastModified** | **time.Time** | Filter by last modified date (exact match). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified date (greater than or equal to). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified date (less than or equal to). | 
 **nameIcontains** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedBucketList**](PaginatedBucketList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBucket

> ResponseBucketCreate RetrieveBucket(ctx, bucketName).Fields(fields).Execute()

Retrieve a bucket



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
	bucketName := "bucketName_example" // string | The name of the bucket
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.RetrieveBucket(context.Background(), bucketName).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.RetrieveBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBucket`: ResponseBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.RetrieveBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseBucketCreate**](ResponseBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> ResponseBucketCreate UpdateBucket(ctx, bucketName).PatchedBucketRequest(patchedBucketRequest).Execute()

Update bucket info



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
	bucketName := "bucketName_example" // string | The name of the bucket
	patchedBucketRequest := *openapiclient.NewPatchedBucketRequest() // PatchedBucketRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.UpdateBucket(context.Background(), bucketName).PatchedBucketRequest(patchedBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.UpdateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBucket`: ResponseBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBucketRequest** | [**PatchedBucketRequest**](PatchedBucketRequest.md) |  | 

### Return type

[**ResponseBucketCreate**](ResponseBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

