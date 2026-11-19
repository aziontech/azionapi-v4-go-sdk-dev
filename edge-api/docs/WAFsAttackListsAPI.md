# \WAFsAttackListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveWafAttackLists**](WAFsAttackListsAPI.md#RetrieveWafAttackLists) | **Get** /workspace/wafs/{waf_id}/attacks | Retrieve the attack lists of a Web Application Firewall (WAF)
[**UpdateWafAttackLists**](WAFsAttackListsAPI.md#UpdateWafAttackLists) | **Put** /workspace/wafs/{waf_id}/attacks | Update the attack lists of a Web Application Firewall (WAF)



## RetrieveWafAttackLists

> WAFAttackListsResponse RetrieveWafAttackLists(ctx, wafId).Fields(fields).Execute()

Retrieve the attack lists of a Web Application Firewall (WAF)



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
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAttackListsAPI.RetrieveWafAttackLists(context.Background(), wafId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAttackListsAPI.RetrieveWafAttackLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWafAttackLists`: WAFAttackListsResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsAttackListsAPI.RetrieveWafAttackLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWafAttackListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**WAFAttackListsResponse**](WAFAttackListsResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWafAttackLists

> WAFAttackListsResponse UpdateWafAttackLists(ctx, wafId).WAFAttackListsRequest(wAFAttackListsRequest).Execute()

Update the attack lists of a Web Application Firewall (WAF)



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
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	wAFAttackListsRequest := *openapiclient.NewWAFAttackListsRequest([]string{"Active_example"}) // WAFAttackListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAttackListsAPI.UpdateWafAttackLists(context.Background(), wafId).WAFAttackListsRequest(wAFAttackListsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAttackListsAPI.UpdateWafAttackLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWafAttackLists`: WAFAttackListsResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsAttackListsAPI.UpdateWafAttackLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafAttackListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wAFAttackListsRequest** | [**WAFAttackListsRequest**](WAFAttackListsRequest.md) |  | 

### Return type

[**WAFAttackListsResponse**](WAFAttackListsResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

