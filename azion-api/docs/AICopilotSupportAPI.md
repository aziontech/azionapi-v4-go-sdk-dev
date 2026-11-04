# \AICopilotSupportAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleSupportWebhook**](AICopilotSupportAPI.md#HandleSupportWebhook) | **Post** /copilot/support | Freshdesk support webhook



## HandleSupportWebhook

> string HandleSupportWebhook(ctx).SupportWebhookRequest(supportWebhookRequest).Execute()

Freshdesk support webhook



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
	supportWebhookRequest := *openapiclient.NewSupportWebhookRequest("TicketID_example", "Subject_example", "Description_example", "ContactFirstName_example") // SupportWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICopilotSupportAPI.HandleSupportWebhook(context.Background()).SupportWebhookRequest(supportWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICopilotSupportAPI.HandleSupportWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleSupportWebhook`: string
	fmt.Fprintf(os.Stdout, "Response from `AICopilotSupportAPI.HandleSupportWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSupportWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportWebhookRequest** | [**SupportWebhookRequest**](SupportWebhookRequest.md) |  | 

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

