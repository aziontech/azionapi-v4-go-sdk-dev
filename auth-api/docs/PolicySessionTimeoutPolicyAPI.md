# \PolicySessionTimeoutPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveSessionTimeoutPolicy**](PolicySessionTimeoutPolicyAPI.md#RetrieveSessionTimeoutPolicy) | **Get** /auth/policies/session | Get Session Timeout Policy
[**UpdateSessionTimeoutPolicy**](PolicySessionTimeoutPolicyAPI.md#UpdateSessionTimeoutPolicy) | **Put** /auth/policies/session | Put Session Timeout Policy



## RetrieveSessionTimeoutPolicy

> ResponseRetrieveSessionTimeoutPolicy RetrieveSessionTimeoutPolicy(ctx).Fields(fields).Execute()

Get Session Timeout Policy



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySessionTimeoutPolicyAPI.RetrieveSessionTimeoutPolicy(context.Background()).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySessionTimeoutPolicyAPI.RetrieveSessionTimeoutPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSessionTimeoutPolicy`: ResponseRetrieveSessionTimeoutPolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicySessionTimeoutPolicyAPI.RetrieveSessionTimeoutPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSessionTimeoutPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveSessionTimeoutPolicy**](ResponseRetrieveSessionTimeoutPolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionTimeoutPolicy

> ResponseSessionTimeoutPolicy UpdateSessionTimeoutPolicy(ctx).SessionTimeoutPolicyRequest(sessionTimeoutPolicyRequest).Execute()

Put Session Timeout Policy



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
	sessionTimeoutPolicyRequest := *openapiclient.NewSessionTimeoutPolicyRequest(int64(123), int64(123)) // SessionTimeoutPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySessionTimeoutPolicyAPI.UpdateSessionTimeoutPolicy(context.Background()).SessionTimeoutPolicyRequest(sessionTimeoutPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySessionTimeoutPolicyAPI.UpdateSessionTimeoutPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSessionTimeoutPolicy`: ResponseSessionTimeoutPolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicySessionTimeoutPolicyAPI.UpdateSessionTimeoutPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionTimeoutPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionTimeoutPolicyRequest** | [**SessionTimeoutPolicyRequest**](SessionTimeoutPolicyRequest.md) |  | 

### Return type

[**ResponseSessionTimeoutPolicy**](ResponseSessionTimeoutPolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

