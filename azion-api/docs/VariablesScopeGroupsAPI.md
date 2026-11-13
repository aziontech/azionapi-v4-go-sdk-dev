# \VariablesScopeGroupsAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScope**](VariablesScopeGroupsAPI.md#CreateScope) | **Post** /scope-groups/{group_slug}/scopes | Create a Scope
[**CreateScopeVariables**](VariablesScopeGroupsAPI.md#CreateScopeVariables) | **Post** /scope-groups/{group_slug}/scopes/{scope_slug}/variables | Create Scope Variables
[**ListScopeVariables**](VariablesScopeGroupsAPI.md#ListScopeVariables) | **Get** /scope-groups/{group_slug}/scopes/{scope_slug}/variables | List Scope Variables
[**UpdateScopeVariables**](VariablesScopeGroupsAPI.md#UpdateScopeVariables) | **Patch** /scope-groups/{group_slug}/scopes/{scope_slug}/variables | Update Scope Variables



## CreateScope

> Scope CreateScope(ctx, groupSlug).ScopeRequest(scopeRequest).Execute()

Create a Scope



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
	groupSlug := "global" // string | The slug of the Scope Group.
	scopeRequest := *openapiclient.NewScopeRequest("Slug_example", "Name_example") // ScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesScopeGroupsAPI.CreateScope(context.Background(), groupSlug).ScopeRequest(scopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesScopeGroupsAPI.CreateScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScope`: Scope
	fmt.Fprintf(os.Stdout, "Response from `VariablesScopeGroupsAPI.CreateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSlug** | **string** | The slug of the Scope Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeRequest** | [**ScopeRequest**](ScopeRequest.md) |  | 

### Return type

[**Scope**](Scope.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScopeVariables

> PaginatedVariableGetList CreateScopeVariables(ctx, groupSlug, scopeSlug).VariableCreateRequest(variableCreateRequest).Page(page).PageSize(pageSize).Execute()

Create Scope Variables



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
	groupSlug := "global" // string | The slug of the Scope Group.
	scopeSlug := "global" // string | The slug of the Scope.
	variableCreateRequest := *openapiclient.NewVariableCreateRequest("Key_example") // VariableCreateRequest | 
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesScopeGroupsAPI.CreateScopeVariables(context.Background(), groupSlug, scopeSlug).VariableCreateRequest(variableCreateRequest).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesScopeGroupsAPI.CreateScopeVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScopeVariables`: PaginatedVariableGetList
	fmt.Fprintf(os.Stdout, "Response from `VariablesScopeGroupsAPI.CreateScopeVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSlug** | **string** | The slug of the Scope Group. | 
**scopeSlug** | **string** | The slug of the Scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **variableCreateRequest** | [**VariableCreateRequest**](VariableCreateRequest.md) |  | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 

### Return type

[**PaginatedVariableGetList**](PaginatedVariableGetList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScopeVariables

> PaginatedVariableGetList ListScopeVariables(ctx, groupSlug, scopeSlug).ClientId(clientId).Page(page).PageSize(pageSize).Execute()

List Scope Variables



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
	clientId := "clientId_example" // string | ID from Client to get the variables
	groupSlug := "global" // string | The slug of the Scope Group.
	scopeSlug := "global" // string | The slug of the Scope.
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesScopeGroupsAPI.ListScopeVariables(context.Background(), groupSlug, scopeSlug).ClientId(clientId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesScopeGroupsAPI.ListScopeVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListScopeVariables`: PaginatedVariableGetList
	fmt.Fprintf(os.Stdout, "Response from `VariablesScopeGroupsAPI.ListScopeVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSlug** | **string** | The slug of the Scope Group. | 
**scopeSlug** | **string** | The slug of the Scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScopeVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ID from Client to get the variables | 


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 

### Return type

[**PaginatedVariableGetList**](PaginatedVariableGetList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScopeVariables

> PaginatedVariableGetList UpdateScopeVariables(ctx, groupSlug, scopeSlug).Page(page).PageSize(pageSize).PatchedVariableBatchUpdateRequest(patchedVariableBatchUpdateRequest).Execute()

Update Scope Variables



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
	groupSlug := "global" // string | The slug of the Scope Group.
	scopeSlug := "global" // string | The slug of the Scope.
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	patchedVariableBatchUpdateRequest := *openapiclient.NewPatchedVariableBatchUpdateRequest() // PatchedVariableBatchUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesScopeGroupsAPI.UpdateScopeVariables(context.Background(), groupSlug, scopeSlug).Page(page).PageSize(pageSize).PatchedVariableBatchUpdateRequest(patchedVariableBatchUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesScopeGroupsAPI.UpdateScopeVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScopeVariables`: PaginatedVariableGetList
	fmt.Fprintf(os.Stdout, "Response from `VariablesScopeGroupsAPI.UpdateScopeVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSlug** | **string** | The slug of the Scope Group. | 
**scopeSlug** | **string** | The slug of the Scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopeVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **patchedVariableBatchUpdateRequest** | [**PatchedVariableBatchUpdateRequest**](PatchedVariableBatchUpdateRequest.md) |  | 

### Return type

[**PaginatedVariableGetList**](PaginatedVariableGetList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

