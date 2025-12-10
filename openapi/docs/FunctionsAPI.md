# \FunctionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsAPI.md#CreateFunction) | **Post** /workspace/functions | Create an Function
[**DeleteFunction**](FunctionsAPI.md#DeleteFunction) | **Delete** /workspace/functions/{function_id} | Delete an Function
[**ListFunctions**](FunctionsAPI.md#ListFunctions) | **Get** /workspace/functions | List Functions
[**PartialUpdateFunction**](FunctionsAPI.md#PartialUpdateFunction) | **Patch** /workspace/functions/{function_id} | Partially update an Function
[**RetrieveFunction**](FunctionsAPI.md#RetrieveFunction) | **Get** /workspace/functions/{function_id} | Retrieve details of an Function
[**UpdateFunction**](FunctionsAPI.md#UpdateFunction) | **Put** /workspace/functions/{function_id} | Update an Function



## CreateFunction

> ResponseFunctionsDoc CreateFunction(ctx).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()

Create an Function



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
	edgeFunctionsRequest := *openapiclient.NewEdgeFunctionsRequest("Name_example", "Code_example") // EdgeFunctionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.CreateFunction(context.Background()).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFunction`: ResponseFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeFunctionsRequest** | [**EdgeFunctionsRequest**](EdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseFunctionsDoc**](ResponseFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> ResponseAsyncDeleteFunctionsDoc DeleteFunction(ctx, functionId).Execute()

Delete an Function



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
	functionId := int32(56) // int32 | A unique integer value identifying the edge function.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.DeleteFunction(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFunction`: ResponseAsyncDeleteFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** | A unique integer value identifying the edge function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteFunctionsDoc**](ResponseAsyncDeleteFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctions

> PaginatedEdgeFunctionsList ListFunctions(ctx).Active(active).DefaultArgs(defaultArgs).ExecutionEnvironment(executionEnvironment).Fields(fields).Id(id).LanguageIn(languageIn).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Runtime(runtime).RuntimeEnvironmentIn(runtimeEnvironmentIn).Search(search).Vendor(vendor).Execute()

List Functions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	active := true // bool | Filter by active status. (optional)
	defaultArgs := "defaultArgs_example" // string | Filter by default args (partial search). (optional)
	executionEnvironment := "executionEnvironment_example" // string | Filter by execution environment. (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := "id_example" // string | Filter by ID. Can be multiple comma-separated values. (optional)
	languageIn := "languageIn_example" // string | Filter by language. Can be multiple comma-separated values. (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (partial search). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified (start). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified (end). (optional)
	name := "name_example" // string | Filter by name (partial search). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, language, json_args, runtime_environment, active, last_editor, last_modified, product_version) (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | A numeric value that indicates the number of items per page. (optional)
	runtime := "runtime_example" // string | Filter by runtime. (optional)
	runtimeEnvironmentIn := "runtimeEnvironmentIn_example" // string | Filter by runtime environment. Can be multiple comma-separated values. (optional)
	search := "search_example" // string | A search term. (optional)
	vendor := "vendor_example" // string | Filter by vendor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.ListFunctions(context.Background()).Active(active).DefaultArgs(defaultArgs).ExecutionEnvironment(executionEnvironment).Fields(fields).Id(id).LanguageIn(languageIn).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Runtime(runtime).RuntimeEnvironmentIn(runtimeEnvironmentIn).Search(search).Vendor(vendor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ListFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFunctions`: PaginatedEdgeFunctionsList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ListFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Filter by active status. | 
 **defaultArgs** | **string** | Filter by default args (partial search). | 
 **executionEnvironment** | **string** | Filter by execution environment. | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **string** | Filter by ID. Can be multiple comma-separated values. | 
 **languageIn** | **string** | Filter by language. Can be multiple comma-separated values. | 
 **lastEditor** | **string** | Filter by last editor (partial search). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified (start). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified (end). | 
 **name** | **string** | Filter by name (partial search). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, language, json_args, runtime_environment, active, last_editor, last_modified, product_version) | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | A numeric value that indicates the number of items per page. | 
 **runtime** | **string** | Filter by runtime. | 
 **runtimeEnvironmentIn** | **string** | Filter by runtime environment. Can be multiple comma-separated values. | 
 **search** | **string** | A search term. | 
 **vendor** | **string** | Filter by vendor. | 

### Return type

[**PaginatedEdgeFunctionsList**](PaginatedEdgeFunctionsList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateFunction

> ResponseFunctionsDoc PartialUpdateFunction(ctx, functionId).PatchedEdgeFunctionsRequest(patchedEdgeFunctionsRequest).Execute()

Partially update an Function



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
	functionId := int32(56) // int32 | A unique integer value identifying the edge function.
	patchedEdgeFunctionsRequest := *openapiclient.NewPatchedEdgeFunctionsRequest() // PatchedEdgeFunctionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.PartialUpdateFunction(context.Background(), functionId).PatchedEdgeFunctionsRequest(patchedEdgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.PartialUpdateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateFunction`: ResponseFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.PartialUpdateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** | A unique integer value identifying the edge function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEdgeFunctionsRequest** | [**PatchedEdgeFunctionsRequest**](PatchedEdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseFunctionsDoc**](ResponseFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFunction

> ResponseRetrieveFunctionsDoc RetrieveFunction(ctx, functionId).Fields(fields).Execute()

Retrieve details of an Function



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
	functionId := int32(56) // int32 | A unique integer value identifying the edge function.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.RetrieveFunction(context.Background(), functionId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.RetrieveFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFunction`: ResponseRetrieveFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.RetrieveFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** | A unique integer value identifying the edge function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveFunctionsDoc**](ResponseRetrieveFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> ResponseFunctionsDoc UpdateFunction(ctx, functionId).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()

Update an Function



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
	functionId := int32(56) // int32 | A unique integer value identifying the edge function.
	edgeFunctionsRequest := *openapiclient.NewEdgeFunctionsRequest("Name_example", "Code_example") // EdgeFunctionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunction(context.Background(), functionId).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunction`: ResponseFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** | A unique integer value identifying the edge function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFunctionsRequest** | [**EdgeFunctionsRequest**](EdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseFunctionsDoc**](ResponseFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

