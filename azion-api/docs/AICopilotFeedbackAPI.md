# \AICopilotFeedbackAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitFeedback**](AICopilotFeedbackAPI.md#SubmitFeedback) | **Post** /copilot/feedback | Submit feedback



## SubmitFeedback

> string SubmitFeedback(ctx).FeedbackRequest(feedbackRequest).Execute()

Submit feedback



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
	feedbackRequest := *openapiclient.NewFeedbackRequest(false, "RunId_example") // FeedbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICopilotFeedbackAPI.SubmitFeedback(context.Background()).FeedbackRequest(feedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICopilotFeedbackAPI.SubmitFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFeedback`: string
	fmt.Fprintf(os.Stdout, "Response from `AICopilotFeedbackAPI.SubmitFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackRequest** | [**FeedbackRequest**](FeedbackRequest.md) |  | 

### Return type

**string**

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

