# \KbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KbCreate**](KbAPI.md#KbCreate) | **Post** /v4/workspace/ai/kb | 
[**KbDestroy**](KbAPI.md#KbDestroy) | **Delete** /v4/workspace/ai/kb/{kbId} | 
[**KbDocumentsChunksRetrieve**](KbAPI.md#KbDocumentsChunksRetrieve) | **Get** /v4/workspace/ai/kb/{kbId}/documents/{documentId}/chunks | 
[**KbDocumentsCreate**](KbAPI.md#KbDocumentsCreate) | **Post** /v4/workspace/ai/kb/{kbId}/documents | 
[**KbDocumentsDestroy**](KbAPI.md#KbDocumentsDestroy) | **Delete** /v4/workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbDocumentsList**](KbAPI.md#KbDocumentsList) | **Get** /v4/workspace/ai/kb/{kbId}/documents | 
[**KbDocumentsRetrieve**](KbAPI.md#KbDocumentsRetrieve) | **Get** /v4/workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbDocumentsUpdate**](KbAPI.md#KbDocumentsUpdate) | **Put** /v4/workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbList**](KbAPI.md#KbList) | **Get** /v4/workspace/ai/kb | 
[**KbPartialUpdate**](KbAPI.md#KbPartialUpdate) | **Patch** /v4/workspace/ai/kb/{kbId} | 
[**KbRetrieve**](KbAPI.md#KbRetrieve) | **Get** /v4/workspace/ai/kb/{kbId} | 
[**KbUpdate**](KbAPI.md#KbUpdate) | **Put** /v4/workspace/ai/kb/{kbId} | 



## KbCreate

> KnowledgeBase KbCreate(ctx).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()



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
	knowledgeBaseRequest := *openapiclient.NewKnowledgeBaseRequest("Name_example") // KnowledgeBaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbCreate(context.Background()).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbCreate`: KnowledgeBase
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKbCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeBaseRequest** | [**KnowledgeBaseRequest**](KnowledgeBaseRequest.md) |  | 

### Return type

[**KnowledgeBase**](KnowledgeBase.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDestroy

> KbDestroy(ctx, kbId).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this knowledge base.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KbAPI.KbDestroy(context.Background(), kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** | A UUID string identifying this knowledge base. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDestroyRequest struct via the builder pattern


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


## KbDocumentsChunksRetrieve

> KbDocumentsChunksRetrieve(ctx, documentId, kbId).Execute()



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KbAPI.KbDocumentsChunksRetrieve(context.Background(), documentId, kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsChunksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsChunksRetrieveRequest struct via the builder pattern


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


## KbDocumentsCreate

> Document KbDocumentsCreate(ctx, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string |  (optional)
	description := "description_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	sourceUri := "sourceUri_example" // string |  (optional)
	chunkStrategy := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbDocumentsCreate(context.Background(), kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsCreate`: Document
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **description** | **string** |  | 
 **type_** | **string** |  | 
 **sourceUri** | **string** |  | 
 **chunkStrategy** | **map[string]interface{}** |  | 

### Return type

[**Document**](Document.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsDestroy

> KbDocumentsDestroy(ctx, documentId, kbId).Execute()



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KbAPI.KbDocumentsDestroy(context.Background(), documentId, kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsDestroyRequest struct via the builder pattern


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


## KbDocumentsList

> PaginatedDocumentList KbDocumentsList(ctx, kbId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbDocumentsList(context.Background(), kbId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsList`: PaginatedDocumentList
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDocumentList**](PaginatedDocumentList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsRetrieve

> Document KbDocumentsRetrieve(ctx, documentId, kbId).Execute()



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbDocumentsRetrieve(context.Background(), documentId, kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsRetrieve`: Document
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Document**](Document.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsUpdate

> Document KbDocumentsUpdate(ctx, documentId, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string |  (optional)
	description := "description_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	sourceUri := "sourceUri_example" // string |  (optional)
	chunkStrategy := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbDocumentsUpdate(context.Background(), documentId, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsUpdate`: Document
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** |  | 
 **description** | **string** |  | 
 **type_** | **string** |  | 
 **sourceUri** | **string** |  | 
 **chunkStrategy** | **map[string]interface{}** |  | 

### Return type

[**Document**](Document.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbList

> PaginatedKnowledgeBaseList KbList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.KbAPI.KbList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbList`: PaginatedKnowledgeBaseList
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedKnowledgeBaseList**](PaginatedKnowledgeBaseList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbPartialUpdate

> KnowledgeBase KbPartialUpdate(ctx, kbId).PatchedKnowledgeBaseRequest(patchedKnowledgeBaseRequest).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this knowledge base.
	patchedKnowledgeBaseRequest := *openapiclient.NewPatchedKnowledgeBaseRequest() // PatchedKnowledgeBaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbPartialUpdate(context.Background(), kbId).PatchedKnowledgeBaseRequest(patchedKnowledgeBaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbPartialUpdate`: KnowledgeBase
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** | A UUID string identifying this knowledge base. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedKnowledgeBaseRequest** | [**PatchedKnowledgeBaseRequest**](PatchedKnowledgeBaseRequest.md) |  | 

### Return type

[**KnowledgeBase**](KnowledgeBase.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbRetrieve

> KnowledgeBase KbRetrieve(ctx, kbId).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this knowledge base.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbRetrieve(context.Background(), kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbRetrieve`: KnowledgeBase
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** | A UUID string identifying this knowledge base. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeBase**](KnowledgeBase.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbUpdate

> KnowledgeBase KbUpdate(ctx, kbId).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()



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
	kbId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this knowledge base.
	knowledgeBaseRequest := *openapiclient.NewKnowledgeBaseRequest("Name_example") // KnowledgeBaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbUpdate(context.Background(), kbId).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbUpdate`: KnowledgeBase
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** | A UUID string identifying this knowledge base. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeBaseRequest** | [**KnowledgeBaseRequest**](KnowledgeBaseRequest.md) |  | 

### Return type

[**KnowledgeBase**](KnowledgeBase.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

