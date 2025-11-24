# \SQLAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](SQLAPI.md#CreateDatabase) | **Post** /edge_sql/databases | Create a database
[**DeleteDatabase**](SQLAPI.md#DeleteDatabase) | **Delete** /edge_sql/databases/{database_id} | Delete a database
[**ExecuteQuery**](SQLAPI.md#ExecuteQuery) | **Post** /edge_sql/databases/{database_id}/query | Execute a query into a database
[**ListDatabases**](SQLAPI.md#ListDatabases) | **Get** /edge_sql/databases | List databases
[**RetrieveDatabase**](SQLAPI.md#RetrieveDatabase) | **Get** /edge_sql/databases/{database_id} | Retrieve details from a database



## CreateDatabase

> ResponseOpenAPISchema CreateDatabase(ctx).DatabaseRequest(databaseRequest).Execute()

Create a database



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
	databaseRequest := *openapiclient.NewDatabaseRequest("Name_example") // DatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SQLAPI.CreateDatabase(context.Background()).DatabaseRequest(databaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SQLAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: ResponseOpenAPISchema
	fmt.Fprintf(os.Stdout, "Response from `SQLAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseRequest** | [**DatabaseRequest**](DatabaseRequest.md) |  | 

### Return type

[**ResponseOpenAPISchema**](ResponseOpenAPISchema.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> ResponseDeleteOpenAPISchema DeleteDatabase(ctx, databaseId).Execute()

Delete a database



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
	databaseId := int64(789) // int64 | A unique integer value identifying this database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SQLAPI.DeleteDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SQLAPI.DeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabase`: ResponseDeleteOpenAPISchema
	fmt.Fprintf(os.Stdout, "Response from `SQLAPI.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int64** | A unique integer value identifying this database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteOpenAPISchema**](ResponseDeleteOpenAPISchema.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteQuery

> ResponseSQLResult ExecuteQuery(ctx, databaseId).SQLStatementsRequest(sQLStatementsRequest).Execute()

Execute a query into a database



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
	databaseId := int64(789) // int64 | A unique integer value identifying this database.
	sQLStatementsRequest := *openapiclient.NewSQLStatementsRequest([]string{"Statements_example"}) // SQLStatementsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SQLAPI.ExecuteQuery(context.Background(), databaseId).SQLStatementsRequest(sQLStatementsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SQLAPI.ExecuteQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteQuery`: ResponseSQLResult
	fmt.Fprintf(os.Stdout, "Response from `SQLAPI.ExecuteQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int64** | A unique integer value identifying this database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sQLStatementsRequest** | [**SQLStatementsRequest**](SQLStatementsRequest.md) |  | 

### Return type

[**ResponseSQLResult**](ResponseSQLResult.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> PaginatedOpenAPISchemaList ListDatabases(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List databases



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, status, active, last_modified) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SQLAPI.ListDatabases(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SQLAPI.ListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabases`: PaginatedOpenAPISchemaList
	fmt.Fprintf(os.Stdout, "Response from `SQLAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, status, active, last_modified) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedOpenAPISchemaList**](PaginatedOpenAPISchemaList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDatabase

> ResponseRetrieveOpenAPISchema RetrieveDatabase(ctx, databaseId).Fields(fields).Execute()

Retrieve details from a database



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
	databaseId := int64(789) // int64 | A unique integer value identifying this database.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SQLAPI.RetrieveDatabase(context.Background(), databaseId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SQLAPI.RetrieveDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDatabase`: ResponseRetrieveOpenAPISchema
	fmt.Fprintf(os.Stdout, "Response from `SQLAPI.RetrieveDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int64** | A unique integer value identifying this database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveOpenAPISchema**](ResponseRetrieveOpenAPISchema.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

