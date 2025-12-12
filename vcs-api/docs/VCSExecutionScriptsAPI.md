# \VCSExecutionScriptsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExecutionScripts**](VCSExecutionScriptsAPI.md#ListExecutionScripts) | **Get** /vcs/execution_scripts | List execution scripts
[**RetriveExecutionScript**](VCSExecutionScriptsAPI.md#RetriveExecutionScript) | **Get** /vcs/execution_scripts/{execution_script_id} | Retrieve details from a execution script



## ListExecutionScripts

> PaginatedResponseListExecutionScriptList ListExecutionScripts(ctx).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List execution scripts



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
	id := "id_example" // string | Filter by id (accepts comma-separated values). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VCSExecutionScriptsAPI.ListExecutionScripts(context.Background()).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VCSExecutionScriptsAPI.ListExecutionScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExecutionScripts`: PaginatedResponseListExecutionScriptList
	fmt.Fprintf(os.Stdout, "Response from `VCSExecutionScriptsAPI.ListExecutionScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExecutionScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **string** | Filter by id (accepts comma-separated values). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListExecutionScriptList**](PaginatedResponseListExecutionScriptList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetriveExecutionScript

> ResponseRetrieveExecutionScript RetriveExecutionScript(ctx, executionScriptId).Fields(fields).Execute()

Retrieve details from a execution script



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
	executionScriptId := int64(789) // int64 | A unique integer value identifying this execution script.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VCSExecutionScriptsAPI.RetriveExecutionScript(context.Background(), executionScriptId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VCSExecutionScriptsAPI.RetriveExecutionScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveExecutionScript`: ResponseRetrieveExecutionScript
	fmt.Fprintf(os.Stdout, "Response from `VCSExecutionScriptsAPI.RetriveExecutionScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionScriptId** | **int64** | A unique integer value identifying this execution script. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveExecutionScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveExecutionScript**](ResponseRetrieveExecutionScript.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

