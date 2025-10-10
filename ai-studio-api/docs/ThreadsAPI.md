# \ThreadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreadsCreate**](ThreadsAPI.md#ThreadsCreate) | **Post** /v4/workspace/ai/threads | 
[**ThreadsDestroy**](ThreadsAPI.md#ThreadsDestroy) | **Delete** /v4/workspace/ai/threads/{threadId} | 
[**ThreadsList**](ThreadsAPI.md#ThreadsList) | **Get** /v4/workspace/ai/threads | 
[**ThreadsMessagesCreate**](ThreadsAPI.md#ThreadsMessagesCreate) | **Post** /v4/workspace/ai/threads/{threadId}/messages | 
[**ThreadsMessagesDestroy**](ThreadsAPI.md#ThreadsMessagesDestroy) | **Delete** /v4/workspace/ai/threads/{threadId}/messages/{messageId} | 
[**ThreadsMessagesList**](ThreadsAPI.md#ThreadsMessagesList) | **Get** /v4/workspace/ai/threads/{threadId}/messages | 
[**ThreadsMessagesRetrieve**](ThreadsAPI.md#ThreadsMessagesRetrieve) | **Get** /v4/workspace/ai/threads/{threadId}/messages/{messageId} | 
[**ThreadsMessagesUpdate**](ThreadsAPI.md#ThreadsMessagesUpdate) | **Put** /v4/workspace/ai/threads/{threadId}/messages/{messageId} | 
[**ThreadsPartialUpdate**](ThreadsAPI.md#ThreadsPartialUpdate) | **Patch** /v4/workspace/ai/threads/{threadId} | 
[**ThreadsRetrieve**](ThreadsAPI.md#ThreadsRetrieve) | **Get** /v4/workspace/ai/threads/{threadId} | 
[**ThreadsUpdate**](ThreadsAPI.md#ThreadsUpdate) | **Put** /v4/workspace/ai/threads/{threadId} | 



## ThreadsCreate

> ChatThread ThreadsCreate(ctx).ChatThreadRequest(chatThreadRequest).Execute()



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
	chatThreadRequest := *openapiclient.NewChatThreadRequest() // ChatThreadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsCreate(context.Background()).ChatThreadRequest(chatThreadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsCreate`: ChatThread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThreadsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatThreadRequest** | [**ChatThreadRequest**](ChatThreadRequest.md) |  | 

### Return type

[**ChatThread**](ChatThread.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsDestroy

> ThreadsDestroy(ctx, threadId).Execute()



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
	threadId := "threadId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreadsAPI.ThreadsDestroy(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsDestroyRequest struct via the builder pattern


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


## ThreadsList

> PaginatedChatThreadList ThreadsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.ThreadsAPI.ThreadsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsList`: PaginatedChatThreadList
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThreadsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedChatThreadList**](PaginatedChatThreadList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsMessagesCreate

> Message ThreadsMessagesCreate(ctx, threadId).MessageRequest(messageRequest).Execute()



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	messageRequest := *openapiclient.NewMessageRequest(map[string]interface{}{"key": interface{}(123)}) // MessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsMessagesCreate(context.Background(), threadId).MessageRequest(messageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsMessagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsMessagesCreate`: Message
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsMessagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsMessagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageRequest** | [**MessageRequest**](MessageRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsMessagesDestroy

> ThreadsMessagesDestroy(ctx, messageId, threadId).Execute()



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
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreadsAPI.ThreadsMessagesDestroy(context.Background(), messageId, threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsMessagesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsMessagesDestroyRequest struct via the builder pattern


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


## ThreadsMessagesList

> PaginatedMessageList ThreadsMessagesList(ctx, threadId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsMessagesList(context.Background(), threadId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsMessagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsMessagesList`: PaginatedMessageList
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsMessagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsMessagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedMessageList**](PaginatedMessageList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsMessagesRetrieve

> Message ThreadsMessagesRetrieve(ctx, messageId, threadId).Execute()



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
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsMessagesRetrieve(context.Background(), messageId, threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsMessagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsMessagesRetrieve`: Message
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsMessagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsMessagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Message**](Message.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsMessagesUpdate

> Message ThreadsMessagesUpdate(ctx, messageId, threadId).MessageRequest(messageRequest).Execute()



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
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	threadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	messageRequest := *openapiclient.NewMessageRequest(map[string]interface{}{"key": interface{}(123)}) // MessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsMessagesUpdate(context.Background(), messageId, threadId).MessageRequest(messageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsMessagesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsMessagesUpdate`: Message
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsMessagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsMessagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messageRequest** | [**MessageRequest**](MessageRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsPartialUpdate

> ChatThread ThreadsPartialUpdate(ctx, threadId).PatchedChatThreadRequest(patchedChatThreadRequest).Execute()



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
	threadId := "threadId_example" // string | 
	patchedChatThreadRequest := *openapiclient.NewPatchedChatThreadRequest() // PatchedChatThreadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsPartialUpdate(context.Background(), threadId).PatchedChatThreadRequest(patchedChatThreadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsPartialUpdate`: ChatThread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedChatThreadRequest** | [**PatchedChatThreadRequest**](PatchedChatThreadRequest.md) |  | 

### Return type

[**ChatThread**](ChatThread.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsRetrieve

> ChatThread ThreadsRetrieve(ctx, threadId).Execute()



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
	threadId := "threadId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsRetrieve(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsRetrieve`: ChatThread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatThread**](ChatThread.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadsUpdate

> ChatThread ThreadsUpdate(ctx, threadId).ChatThreadRequest(chatThreadRequest).Execute()



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
	threadId := "threadId_example" // string | 
	chatThreadRequest := *openapiclient.NewChatThreadRequest() // ChatThreadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.ThreadsUpdate(context.Background(), threadId).ChatThreadRequest(chatThreadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.ThreadsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadsUpdate`: ChatThread
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.ThreadsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatThreadRequest** | [**ChatThreadRequest**](ChatThreadRequest.md) |  | 

### Return type

[**ChatThread**](ChatThread.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

