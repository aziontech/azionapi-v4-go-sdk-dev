# \OrchestratorEdgeServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeService**](OrchestratorEdgeServicesAPI.md#CreateEdgeService) | **Post** /edge_orchestrator/edge_services | Create Edge Service
[**DestroyEdgeService**](OrchestratorEdgeServicesAPI.md#DestroyEdgeService) | **Delete** /edge_orchestrator/edge_services/{serviceId} | Destroy an Edge Service
[**ListEdgeServices**](OrchestratorEdgeServicesAPI.md#ListEdgeServices) | **Get** /edge_orchestrator/edge_services | List Edge Services
[**PartialUpdateEdgeService**](OrchestratorEdgeServicesAPI.md#PartialUpdateEdgeService) | **Patch** /edge_orchestrator/edge_services/{serviceId} | Partially update an Edge Service
[**RetrieveEdgeService**](OrchestratorEdgeServicesAPI.md#RetrieveEdgeService) | **Get** /edge_orchestrator/edge_services/{serviceId} | Retrieve details of an Edge Service
[**UpdateEdgeService**](OrchestratorEdgeServicesAPI.md#UpdateEdgeService) | **Put** /edge_orchestrator/edge_services/{serviceId} | Update an Edge Service



## CreateEdgeService

> ResponseAsyncServices CreateEdgeService(ctx).ServicesRequest(servicesRequest).Execute()

Create Edge Service



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
	servicesRequest := *openapiclient.NewServicesRequest() // ServicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.CreateEdgeService(context.Background()).ServicesRequest(servicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.CreateEdgeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeService`: ResponseAsyncServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.CreateEdgeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicesRequest** | [**ServicesRequest**](ServicesRequest.md) |  | 

### Return type

[**ResponseAsyncServices**](ResponseAsyncServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeService

> ResponseAsyncDeleteServices DestroyEdgeService(ctx, serviceId).Execute()

Destroy an Edge Service



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
	serviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.DestroyEdgeService(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.DestroyEdgeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeService`: ResponseAsyncDeleteServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.DestroyEdgeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteServices**](ResponseAsyncDeleteServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeServices

> PaginatedResponseListServicesList ListEdgeServices(ctx).Fields(fields).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Execute()

List Edge Services



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
	name := "name_example" // string | Search by name (optional)
	ordering := "ordering_example" // string | Field to order by, use '-' for descending (optional)
	page := int64(789) // int64 | Page number for pagination (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.ListEdgeServices(context.Background()).Fields(fields).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.ListEdgeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeServices`: PaginatedResponseListServicesList
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.ListEdgeServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **name** | **string** | Search by name | 
 **ordering** | **string** | Field to order by, use &#39;-&#39; for descending | 
 **page** | **int64** | Page number for pagination | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 

### Return type

[**PaginatedResponseListServicesList**](PaginatedResponseListServicesList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeService

> ResponseAsyncServices PartialUpdateEdgeService(ctx, serviceId).PatchedServicesRequest(patchedServicesRequest).Execute()

Partially update an Edge Service



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
	serviceId := int64(789) // int64 | 
	patchedServicesRequest := *openapiclient.NewPatchedServicesRequest() // PatchedServicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.PartialUpdateEdgeService(context.Background(), serviceId).PatchedServicesRequest(patchedServicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.PartialUpdateEdgeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeService`: ResponseAsyncServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.PartialUpdateEdgeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedServicesRequest** | [**PatchedServicesRequest**](PatchedServicesRequest.md) |  | 

### Return type

[**ResponseAsyncServices**](ResponseAsyncServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeService

> ResponseRetrieveServices RetrieveEdgeService(ctx, serviceId).Fields(fields).Execute()

Retrieve details of an Edge Service



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
	serviceId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.RetrieveEdgeService(context.Background(), serviceId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.RetrieveEdgeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeService`: ResponseRetrieveServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.RetrieveEdgeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveServices**](ResponseRetrieveServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeService

> ResponseAsyncServices UpdateEdgeService(ctx, serviceId).ServicesRequest(servicesRequest).Execute()

Update an Edge Service



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
	serviceId := "serviceId_example" // string | 
	servicesRequest := *openapiclient.NewServicesRequest() // ServicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrchestratorEdgeServicesAPI.UpdateEdgeService(context.Background(), serviceId).ServicesRequest(servicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorEdgeServicesAPI.UpdateEdgeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeService`: ResponseAsyncServices
	fmt.Fprintf(os.Stdout, "Response from `OrchestratorEdgeServicesAPI.UpdateEdgeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicesRequest** | [**ServicesRequest**](ServicesRequest.md) |  | 

### Return type

[**ResponseAsyncServices**](ResponseAsyncServices.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

