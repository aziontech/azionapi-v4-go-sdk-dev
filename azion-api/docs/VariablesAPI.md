# \VariablesAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](VariablesAPI.md#CreateVariable) | **Post** /variables | Create a Variable
[**DeleteVariable**](VariablesAPI.md#DeleteVariable) | **Delete** /variables/{uuid} | Delete a Variable
[**ListVariables**](VariablesAPI.md#ListVariables) | **Get** /variables | List Variables
[**PartialUpdateVariable**](VariablesAPI.md#PartialUpdateVariable) | **Patch** /variables/{uuid} | Partially update a Variable
[**RetrieveVariable**](VariablesAPI.md#RetrieveVariable) | **Get** /variables/{uuid} | Retrieve details of a Variable
[**UpdateVariable**](VariablesAPI.md#UpdateVariable) | **Put** /variables/{uuid} | Update a Variable



## CreateVariable

> VariableGet CreateVariable(ctx).VariableCreateRequest(variableCreateRequest).Execute()

Create a Variable



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
	variableCreateRequest := *openapiclient.NewVariableCreateRequest("Key_example") // VariableCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.CreateVariable(context.Background()).VariableCreateRequest(variableCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.CreateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVariable`: VariableGet
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.CreateVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableCreateRequest** | [**VariableCreateRequest**](VariableCreateRequest.md) |  | 

### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariable

> DeleteVariable(ctx, uuid).Execute()

Delete a Variable



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
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.DeleteVariable(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.DeleteVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariables

> PaginatedVariableGetList ListVariables(ctx).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Key(key).LastEditor(lastEditor).Ordering(ordering).Page(page).PageSize(pageSize).ResourceId(resourceId).ResourceType(resourceType).Scope(scope).ScopeId(scopeId).ScopeType(scopeType).Search(search).Secret(secret).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Uuid(uuid).Execute()

List Variables



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
	createdAtGte := time.Now() // time.Time | Filter by creation date (greater than or equal). (optional)
	createdAtLte := time.Now() // time.Time | Filter by creation date (less than or equal). (optional)
	key := "key_example" // string | Filter by key (case-insensitive, partial match). (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	resourceId := "resourceId_example" // string | Filter by scope resource id (the external id of the resource_type). Accepts comma-separated values, unless `resource_type` also holds several values. (optional)
	resourceType := "resourceType_example" // string | Filter by scope resource type (global, environment, deployment, application, firewall). Accepts comma-separated values. Combined with `resource_id`, both must match the same scope, and only one of the two may hold several values. Use `scope` to ask for several type and id pairs. (optional)
	scope := "scope_example" // string | Filter by scope. Comma-separated selectors, each `global` or `<resource_type>:<resource_id>` (e.g. `global,environment:AJQ6CZJE`). Returns variables bound to any selector. Up to 50 selectors. Sending `scope` without any selector is rejected. (optional)
	scopeId := "scopeId_example" // string | Deprecated alias for `resource_id`. (optional)
	scopeType := "scopeType_example" // string | Deprecated alias for `resource_type`. (optional)
	search := "search_example" // string | A search term. (optional)
	secret := true // bool | Filter by secret status (true for secret variables, false for non-secret variables). (optional)
	updatedAtGte := time.Now() // time.Time | Filter by updated date (greater than or equal). (optional)
	updatedAtLte := time.Now() // time.Time | Filter by updated date (less than or equal). (optional)
	uuid := "uuid_example" // string | Filter by uuid (accepts comma-separated values). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.ListVariables(context.Background()).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Key(key).LastEditor(lastEditor).Ordering(ordering).Page(page).PageSize(pageSize).ResourceId(resourceId).ResourceType(resourceType).Scope(scope).ScopeId(scopeId).ScopeType(scopeType).Search(search).Secret(secret).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ListVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVariables`: PaginatedVariableGetList
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.ListVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAtGte** | **time.Time** | Filter by creation date (greater than or equal). | 
 **createdAtLte** | **time.Time** | Filter by creation date (less than or equal). | 
 **key** | **string** | Filter by key (case-insensitive, partial match). | 
 **lastEditor** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **resourceId** | **string** | Filter by scope resource id (the external id of the resource_type). Accepts comma-separated values, unless &#x60;resource_type&#x60; also holds several values. | 
 **resourceType** | **string** | Filter by scope resource type (global, environment, deployment, application, firewall). Accepts comma-separated values. Combined with &#x60;resource_id&#x60;, both must match the same scope, and only one of the two may hold several values. Use &#x60;scope&#x60; to ask for several type and id pairs. | 
 **scope** | **string** | Filter by scope. Comma-separated selectors, each &#x60;global&#x60; or &#x60;&lt;resource_type&gt;:&lt;resource_id&gt;&#x60; (e.g. &#x60;global,environment:AJQ6CZJE&#x60;). Returns variables bound to any selector. Up to 50 selectors. Sending &#x60;scope&#x60; without any selector is rejected. | 
 **scopeId** | **string** | Deprecated alias for &#x60;resource_id&#x60;. | 
 **scopeType** | **string** | Deprecated alias for &#x60;resource_type&#x60;. | 
 **search** | **string** | A search term. | 
 **secret** | **bool** | Filter by secret status (true for secret variables, false for non-secret variables). | 
 **updatedAtGte** | **time.Time** | Filter by updated date (greater than or equal). | 
 **updatedAtLte** | **time.Time** | Filter by updated date (less than or equal). | 
 **uuid** | **string** | Filter by uuid (accepts comma-separated values). | 

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


## PartialUpdateVariable

> VariableGet PartialUpdateVariable(ctx, uuid).PatchedVariableCreateRequest(patchedVariableCreateRequest).Execute()

Partially update a Variable



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
	uuid := "uuid_example" // string | 
	patchedVariableCreateRequest := *openapiclient.NewPatchedVariableCreateRequest() // PatchedVariableCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.PartialUpdateVariable(context.Background(), uuid).PatchedVariableCreateRequest(patchedVariableCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.PartialUpdateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateVariable`: VariableGet
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.PartialUpdateVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedVariableCreateRequest** | [**PatchedVariableCreateRequest**](PatchedVariableCreateRequest.md) |  | 

### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVariable

> VariableGet RetrieveVariable(ctx, uuid).Execute()

Retrieve details of a Variable



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
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.RetrieveVariable(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.RetrieveVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVariable`: VariableGet
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.RetrieveVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariable

> VariableGet UpdateVariable(ctx, uuid).VariableCreateRequest(variableCreateRequest).Execute()

Update a Variable



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
	uuid := "uuid_example" // string | 
	variableCreateRequest := *openapiclient.NewVariableCreateRequest("Key_example") // VariableCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.UpdateVariable(context.Background(), uuid).VariableCreateRequest(variableCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.UpdateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVariable`: VariableGet
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.UpdateVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableCreateRequest** | [**VariableCreateRequest**](VariableCreateRequest.md) |  | 

### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

