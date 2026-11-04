# \AICopilotDocumentationSearchAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchDocs**](AICopilotDocumentationSearchAPI.md#SearchDocs) | **Post** /copilot/docs | Documentation search



## SearchDocs

> []DocsResponseItem SearchDocs(ctx).DocsRequest(docsRequest).Execute()

Documentation search



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
	docsRequest := *openapiclient.NewDocsRequest([]openapiclient.ChatCompletionMessage{*openapiclient.NewChatCompletionMessage("Role_example", "Content_example")}) // DocsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICopilotDocumentationSearchAPI.SearchDocs(context.Background()).DocsRequest(docsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICopilotDocumentationSearchAPI.SearchDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDocs`: []DocsResponseItem
	fmt.Fprintf(os.Stdout, "Response from `AICopilotDocumentationSearchAPI.SearchDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **docsRequest** | [**DocsRequest**](DocsRequest.md) |  | 

### Return type

[**[]DocsResponseItem**](DocsResponseItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

