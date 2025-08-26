# \ConnectorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorsAPI.md#CreateConnector) | **Post** /edge_connector/connectors | Create an Connector
[**DestroyConnector**](ConnectorsAPI.md#DestroyConnector) | **Delete** /edge_connector/connectors/{id} | Destroy an Connector
[**ListConnectors**](ConnectorsAPI.md#ListConnectors) | **Get** /edge_connector/connectors | List Connectors
[**PartialUpdateConnector**](ConnectorsAPI.md#PartialUpdateConnector) | **Patch** /edge_connector/connectors/{id} | Partially update an Connector
[**RetrieveConnector**](ConnectorsAPI.md#RetrieveConnector) | **Get** /edge_connector/connectors/{id} | Retrieve details of an Connector
[**UpdateConnector**](ConnectorsAPI.md#UpdateConnector) | **Put** /edge_connector/connectors/{id} | Update an Connector



## CreateConnector

> ResponseConnectorPolymorphic CreateConnector(ctx).ConnectorPolymorphicRequest(connectorPolymorphicRequest).Execute()

Create an Connector



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
	connectorPolymorphicRequest := openapiclient.ConnectorPolymorphicRequest{ConnectorHTTPRequest: openapiclient.NewConnectorHTTPRequest("Name_example", "Type_example", *openapiclient.NewConnectorHTTPAttributesRequest([]openapiclient.AddressRequest{*openapiclient.NewAddressRequest("Address_example")}))} // ConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CreateConnector(context.Background()).ConnectorPolymorphicRequest(connectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CreateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnector`: ResponseConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CreateConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorPolymorphicRequest** | [**ConnectorPolymorphicRequest**](ConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseConnectorPolymorphic**](ResponseConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyConnector

> ResponseDeleteConnectorPolymorphic DestroyConnector(ctx, id).Execute()

Destroy an Connector



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.DestroyConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.DestroyConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyConnector`: ResponseDeleteConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.DestroyConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteConnectorPolymorphic**](ResponseDeleteConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> PaginatedConnectorPolymorphicList ListConnectors(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Connectors



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: ) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.ListConnectors(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.ListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectors`: PaginatedConnectorPolymorphicList
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: ) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedConnectorPolymorphicList**](PaginatedConnectorPolymorphicList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateConnector

> ResponseConnectorPolymorphic PartialUpdateConnector(ctx, id).PatchedConnectorPolymorphicRequest(patchedConnectorPolymorphicRequest).Execute()

Partially update an Connector



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
	id := "id_example" // string | 
	patchedConnectorPolymorphicRequest := openapiclient.PatchedConnectorPolymorphicRequest{PatchedConnectorHTTPRequest: openapiclient.NewPatchedConnectorHTTPRequest()} // PatchedConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PartialUpdateConnector(context.Background(), id).PatchedConnectorPolymorphicRequest(patchedConnectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PartialUpdateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateConnector`: ResponseConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PartialUpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedConnectorPolymorphicRequest** | [**PatchedConnectorPolymorphicRequest**](PatchedConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseConnectorPolymorphic**](ResponseConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConnector

> ResponseRetrieveConnectorPolymorphic RetrieveConnector(ctx, id).Fields(fields).Execute()

Retrieve details of an Connector



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
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.RetrieveConnector(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.RetrieveConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveConnector`: ResponseRetrieveConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.RetrieveConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveConnectorPolymorphic**](ResponseRetrieveConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnector

> ResponseConnectorPolymorphic UpdateConnector(ctx, id).ConnectorPolymorphicRequest(connectorPolymorphicRequest).Execute()

Update an Connector



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
	id := "id_example" // string | 
	connectorPolymorphicRequest := openapiclient.ConnectorPolymorphicRequest{ConnectorHTTPRequest: openapiclient.NewConnectorHTTPRequest("Name_example", "Type_example", *openapiclient.NewConnectorHTTPAttributesRequest([]openapiclient.AddressRequest{*openapiclient.NewAddressRequest("Address_example")}))} // ConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.UpdateConnector(context.Background(), id).ConnectorPolymorphicRequest(connectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.UpdateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnector`: ResponseConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.UpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorPolymorphicRequest** | [**ConnectorPolymorphicRequest**](ConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseConnectorPolymorphic**](ResponseConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

