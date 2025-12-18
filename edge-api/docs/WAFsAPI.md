# \WAFsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneWaf**](WAFsAPI.md#CloneWaf) | **Post** /workspace/wafs/{waf_id}/clone | Clone a Web Application Firewall (WAF)
[**CreateWaf**](WAFsAPI.md#CreateWaf) | **Post** /workspace/wafs | Create a Web Application Firewall (WAF)
[**DeleteWaf**](WAFsAPI.md#DeleteWaf) | **Delete** /workspace/wafs/{waf_id} | Delete a Web Application Firewall (WAF)
[**ListWafs**](WAFsAPI.md#ListWafs) | **Get** /workspace/wafs | List Web Application Firewalls (WAFs)
[**PartialUpdateWaf**](WAFsAPI.md#PartialUpdateWaf) | **Patch** /workspace/wafs/{waf_id} | Partially update a Web Application Firewall (WAF)
[**RetrieveWaf**](WAFsAPI.md#RetrieveWaf) | **Get** /workspace/wafs/{waf_id} | Retrieve details from a Web Application Firewall (WAF)
[**UpdateWaf**](WAFsAPI.md#UpdateWaf) | **Put** /workspace/wafs/{waf_id} | Update a Web Application Firewall (WAF)



## CloneWaf

> ResponseWAF CloneWaf(ctx, wafId).CloneWRequest(cloneWRequest).Execute()

Clone a Web Application Firewall (WAF)



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
	cloneWRequest := *openapiclient.NewCloneWRequest("Name_example") // CloneWRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.CloneWaf(context.Background(), wafId).CloneWRequest(cloneWRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.CloneWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneWaf`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.CloneWaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneWRequest** | [**CloneWRequest**](CloneWRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWaf

> ResponseWAF CreateWaf(ctx).WRequest(wRequest).Execute()

Create a Web Application Firewall (WAF)



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
	wRequest := *openapiclient.NewWRequest("Name_example") // WRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.CreateWaf(context.Background()).WRequest(wRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.CreateWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWaf`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.CreateWaf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wRequest** | [**WRequest**](WRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWaf

> ResponseDeleteWAF DeleteWaf(ctx, wafId).Execute()

Delete a Web Application Firewall (WAF)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.DeleteWaf(context.Background(), wafId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.DeleteWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWaf`: ResponseDeleteWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.DeleteWaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteWAF**](ResponseDeleteWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWafs

> PaginatedWList ListWafs(ctx).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Web Application Firewalls (WAFs)



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
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, active, last_editor, last_modified, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.ListWafs(context.Background()).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.ListWafs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWafs`: PaginatedWList
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.ListWafs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWafsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, active, last_editor, last_modified, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWList**](PaginatedWList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWaf

> ResponseWAF PartialUpdateWaf(ctx, wafId).PatchedWRequest(patchedWRequest).Execute()

Partially update a Web Application Firewall (WAF)



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
	patchedWRequest := *openapiclient.NewPatchedWRequest() // PatchedWRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.PartialUpdateWaf(context.Background(), wafId).PatchedWRequest(patchedWRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.PartialUpdateWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWaf`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.PartialUpdateWaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWRequest** | [**PatchedWRequest**](PatchedWRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWaf

> ResponseRetrieveWAF RetrieveWaf(ctx, wafId).Fields(fields).Execute()

Retrieve details from a Web Application Firewall (WAF)



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
	resp, r, err := apiClient.WAFsAPI.RetrieveWaf(context.Background(), wafId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.RetrieveWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWaf`: ResponseRetrieveWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.RetrieveWaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveWAF**](ResponseRetrieveWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWaf

> ResponseWAF UpdateWaf(ctx, wafId).WRequest(wRequest).Execute()

Update a Web Application Firewall (WAF)



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
	wRequest := *openapiclient.NewWRequest("Name_example") // WRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.UpdateWaf(context.Background(), wafId).WRequest(wRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.UpdateWaf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWaf`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.UpdateWaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wRequest** | [**WRequest**](WRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

