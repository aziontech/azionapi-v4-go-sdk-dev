# \KbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KbCreate**](KbAPI.md#KbCreate) | **Post** /workspace/ai/kb | 
[**KbDestroy**](KbAPI.md#KbDestroy) | **Delete** /workspace/ai/kb/{kbId} | 
[**KbDocumentsChunksList**](KbAPI.md#KbDocumentsChunksList) | **Get** /workspace/ai/kb/{kbId}/documents/{documentId}/chunks | 
[**KbDocumentsCreate**](KbAPI.md#KbDocumentsCreate) | **Post** /workspace/ai/kb/{kbId}/documents | 
[**KbDocumentsDestroy**](KbAPI.md#KbDocumentsDestroy) | **Delete** /workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbDocumentsList**](KbAPI.md#KbDocumentsList) | **Get** /workspace/ai/kb/{kbId}/documents | 
[**KbDocumentsRetrieve**](KbAPI.md#KbDocumentsRetrieve) | **Get** /workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbDocumentsUpdate**](KbAPI.md#KbDocumentsUpdate) | **Put** /workspace/ai/kb/{kbId}/documents/{documentId} | 
[**KbList**](KbAPI.md#KbList) | **Get** /workspace/ai/kb | 
[**KbPartialUpdate**](KbAPI.md#KbPartialUpdate) | **Patch** /workspace/ai/kb/{kbId} | 
[**KbRetrieve**](KbAPI.md#KbRetrieve) | **Get** /workspace/ai/kb/{kbId} | 
[**KbUpdate**](KbAPI.md#KbUpdate) | **Put** /workspace/ai/kb/{kbId} | 



## KbCreate

> ResponseKnowledgeBase KbCreate(ctx).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()



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
	// response from `KbCreate`: ResponseKnowledgeBase
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

[**ResponseKnowledgeBase**](ResponseKnowledgeBase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDestroy

> ResponseDeleteKnowledgeBase KbDestroy(ctx, kbId).Execute()



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
	resp, r, err := apiClient.KbAPI.KbDestroy(context.Background(), kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDestroy`: ResponseDeleteKnowledgeBase
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDestroy`: %v\n", resp)
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

[**ResponseDeleteKnowledgeBase**](ResponseDeleteKnowledgeBase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsChunksList

> PaginatedChunkList KbDocumentsChunksList(ctx, documentId, kbId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.KbDocumentsChunksList(context.Background(), documentId, kbId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsChunksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsChunksList`: PaginatedChunkList
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsChunksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDocumentsChunksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedChunkList**](PaginatedChunkList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsCreate

> ResponseDocument KbDocumentsCreate(ctx, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()



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
	// response from `KbDocumentsCreate`: ResponseDocument
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

[**ResponseDocument**](ResponseDocument.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsDestroy

> ResponseDeleteDocument KbDocumentsDestroy(ctx, documentId, kbId).Execute()



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
	resp, r, err := apiClient.KbAPI.KbDocumentsDestroy(context.Background(), documentId, kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.KbDocumentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDocumentsDestroy`: ResponseDeleteDocument
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.KbDocumentsDestroy`: %v\n", resp)
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

[**ResponseDeleteDocument**](ResponseDeleteDocument.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
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
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDocumentList**](PaginatedDocumentList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsRetrieve

> ResponseRetrieveDocument KbDocumentsRetrieve(ctx, documentId, kbId).Execute()



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
	// response from `KbDocumentsRetrieve`: ResponseRetrieveDocument
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

[**ResponseRetrieveDocument**](ResponseRetrieveDocument.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDocumentsUpdate

> ResponseDocument KbDocumentsUpdate(ctx, documentId, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()



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
	// response from `KbDocumentsUpdate`: ResponseDocument
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

[**ResponseDocument**](ResponseDocument.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

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
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
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
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedKnowledgeBaseList**](PaginatedKnowledgeBaseList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbPartialUpdate

> ResponseKnowledgeBase KbPartialUpdate(ctx, kbId).PatchedKnowledgeBaseRequest(patchedKnowledgeBaseRequest).Execute()



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
	// response from `KbPartialUpdate`: ResponseKnowledgeBase
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

[**ResponseKnowledgeBase**](ResponseKnowledgeBase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbRetrieve

> ResponseRetrieveKnowledgeBase KbRetrieve(ctx, kbId).Execute()



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
	// response from `KbRetrieve`: ResponseRetrieveKnowledgeBase
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

[**ResponseRetrieveKnowledgeBase**](ResponseRetrieveKnowledgeBase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbUpdate

> ResponseKnowledgeBase KbUpdate(ctx, kbId).KnowledgeBaseRequest(knowledgeBaseRequest).Execute()



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
	// response from `KbUpdate`: ResponseKnowledgeBase
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

[**ResponseKnowledgeBase**](ResponseKnowledgeBase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

