# \AIStudioDocumentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](AIStudioDocumentsAPI.md#CreateDocument) | **Post** /workspace/ai/kbs/{kbId}/documents | Create a document
[**DestroyADocument**](AIStudioDocumentsAPI.md#DestroyADocument) | **Delete** /workspace/ai/kbs/{kbId}/documents/{documentId} | Destroy a document
[**ListDocuments**](AIStudioDocumentsAPI.md#ListDocuments) | **Get** /workspace/ai/kbs/{kbId}/documents | List documents
[**RetriveDocument**](AIStudioDocumentsAPI.md#RetriveDocument) | **Get** /workspace/ai/kbs/{kbId}/documents/{documentId} | Retrieve details from a document
[**UpdateDocument**](AIStudioDocumentsAPI.md#UpdateDocument) | **Put** /workspace/ai/kbs/{kbId}/documents/{documentId} | Update a document



## CreateDocument

> ResponseDocument CreateDocument(ctx, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()

Create a document



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
	resp, r, err := apiClient.AIStudioDocumentsAPI.CreateDocument(context.Background(), kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIStudioDocumentsAPI.CreateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDocument`: ResponseDocument
	fmt.Fprintf(os.Stdout, "Response from `AIStudioDocumentsAPI.CreateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


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


## DestroyADocument

> ResponseDeleteDocument DestroyADocument(ctx, documentId, kbId).Execute()

Destroy a document



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
	resp, r, err := apiClient.AIStudioDocumentsAPI.DestroyADocument(context.Background(), documentId, kbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIStudioDocumentsAPI.DestroyADocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyADocument`: ResponseDeleteDocument
	fmt.Fprintf(os.Stdout, "Response from `AIStudioDocumentsAPI.DestroyADocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyADocumentRequest struct via the builder pattern


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


## ListDocuments

> PaginatedDocumentList ListDocuments(ctx, kbId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List documents



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIStudioDocumentsAPI.ListDocuments(context.Background(), kbId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIStudioDocumentsAPI.ListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocuments`: PaginatedDocumentList
	fmt.Fprintf(os.Stdout, "Response from `AIStudioDocumentsAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
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


## RetriveDocument

> ResponseRetrieveDocument RetriveDocument(ctx, documentId, kbId).Fields(fields).Execute()

Retrieve details from a document



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIStudioDocumentsAPI.RetriveDocument(context.Background(), documentId, kbId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIStudioDocumentsAPI.RetriveDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveDocument`: ResponseRetrieveDocument
	fmt.Fprintf(os.Stdout, "Response from `AIStudioDocumentsAPI.RetriveDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

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


## UpdateDocument

> ResponseDocument UpdateDocument(ctx, documentId, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()

Update a document



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
	resp, r, err := apiClient.AIStudioDocumentsAPI.UpdateDocument(context.Background(), documentId, kbId).Name(name).Description(description).Type_(type_).SourceUri(sourceUri).ChunkStrategy(chunkStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIStudioDocumentsAPI.UpdateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDocument`: ResponseDocument
	fmt.Fprintf(os.Stdout, "Response from `AIStudioDocumentsAPI.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 
**kbId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


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

