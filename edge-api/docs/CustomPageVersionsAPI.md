# \CustomPageVersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveCustomPageVersion**](CustomPageVersionsAPI.md#ArchiveCustomPageVersion) | **Post** /workspace/custom_pages/{resource_pk}/versions/{id}/archive | Archive a Custom Page version
[**BuildCustomPageVersion**](CustomPageVersionsAPI.md#BuildCustomPageVersion) | **Post** /workspace/custom_pages/{resource_pk}/versions/{id}/build | Build a Custom Page version
[**CancelCustomPageVersionBuild**](CustomPageVersionsAPI.md#CancelCustomPageVersionBuild) | **Post** /workspace/custom_pages/{resource_pk}/versions/{id}/cancel | Cancel a Custom Page version build
[**CreateCustomPageVersion**](CustomPageVersionsAPI.md#CreateCustomPageVersion) | **Post** /workspace/custom_pages/{resource_pk}/versions | Create a new Custom Page version
[**DeleteCustomPageVersion**](CustomPageVersionsAPI.md#DeleteCustomPageVersion) | **Delete** /workspace/custom_pages/{resource_pk}/versions/{id} | Delete a Custom Page version
[**ListCustomPageVersions**](CustomPageVersionsAPI.md#ListCustomPageVersions) | **Get** /workspace/custom_pages/{resource_pk}/versions | List Custom Page versions
[**PartialUpdateCustomPageVersion**](CustomPageVersionsAPI.md#PartialUpdateCustomPageVersion) | **Patch** /workspace/custom_pages/{resource_pk}/versions/{id} | Partially update a Custom Page version
[**RetrieveCustomPageVersion**](CustomPageVersionsAPI.md#RetrieveCustomPageVersion) | **Get** /workspace/custom_pages/{resource_pk}/versions/{id} | Retrieve a Custom Page version
[**UpdateCustomPageVersion**](CustomPageVersionsAPI.md#UpdateCustomPageVersion) | **Put** /workspace/custom_pages/{resource_pk}/versions/{id} | Update a Custom Page version



## ArchiveCustomPageVersion

> ArchiveCustomPageVersion(ctx, id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()

Archive a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	versionArchiveRequest := *openapiclient.NewVersionArchiveRequest() // VersionArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.ArchiveCustomPageVersion(context.Background(), id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.ArchiveCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionArchiveRequest** | [**VersionArchiveRequest**](VersionArchiveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCustomPageVersion

> BuildCustomPageVersion(ctx, id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()

Build a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	versionBuildRequest := *openapiclient.NewVersionBuildRequest() // VersionBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.BuildCustomPageVersion(context.Background(), id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.BuildCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionBuildRequest** | [**VersionBuildRequest**](VersionBuildRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelCustomPageVersionBuild

> CancelCustomPageVersionBuild(ctx, id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()

Cancel a Custom Page version build



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	versionCancelRequest := *openapiclient.NewVersionCancelRequest() // VersionCancelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.CancelCustomPageVersionBuild(context.Background(), id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.CancelCustomPageVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCustomPageVersionBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionCancelRequest** | [**VersionCancelRequest**](VersionCancelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomPageVersion

> CreateCustomPageVersion(ctx, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Create a new Custom Page version



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
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.CreateCustomPageVersion(context.Background(), resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.CreateCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionCreateRequest** | [**VersionCreateRequest**](VersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomPageVersion

> DeleteCustomPageVersion(ctx, id, resourcePk).Execute()

Delete a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.DeleteCustomPageVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.DeleteCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomPageVersions

> ListCustomPageVersions(ctx, resourcePk).Fields(fields).Execute()

List Custom Page versions



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
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.ListCustomPageVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.ListCustomPageVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomPageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateCustomPageVersion

> PartialUpdateCustomPageVersion(ctx, id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()

Partially update a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	patchedVersionCreateRequest := *openapiclient.NewPatchedVersionCreateRequest() // PatchedVersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.PartialUpdateCustomPageVersion(context.Background(), id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.PartialUpdateCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedVersionCreateRequest** | [**PatchedVersionCreateRequest**](PatchedVersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCustomPageVersion

> RetrieveCustomPageVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.RetrieveCustomPageVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.RetrieveCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomPageVersion

> UpdateCustomPageVersion(ctx, id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Update a Custom Page version



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
	id := "id_example" // string | The ULID identifier of the version.
	resourcePk := int64(789) // int64 | The ID of the Custom Page resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPageVersionsAPI.UpdateCustomPageVersion(context.Background(), id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPageVersionsAPI.UpdateCustomPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Custom Page resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionCreateRequest** | [**VersionCreateRequest**](VersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

