# \ApplicationVersionsAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveApplicationVersion**](ApplicationVersionsAPI.md#ArchiveApplicationVersion) | **Post** /workspace/applications/{resource_pk}/versions/{id}/archive | Archive an Application version
[**BuildApplicationVersion**](ApplicationVersionsAPI.md#BuildApplicationVersion) | **Post** /workspace/applications/{resource_pk}/versions/{id}/build | Build an Application version
[**CancelApplicationVersionBuild**](ApplicationVersionsAPI.md#CancelApplicationVersionBuild) | **Post** /workspace/applications/{resource_pk}/versions/{id}/cancel | Cancel an Application version build
[**CreateApplicationVersion**](ApplicationVersionsAPI.md#CreateApplicationVersion) | **Post** /workspace/applications/{resource_pk}/versions | Create a new Application version
[**DeleteApplicationVersion**](ApplicationVersionsAPI.md#DeleteApplicationVersion) | **Delete** /workspace/applications/{resource_pk}/versions/{id} | Delete an Application version
[**ListApplicationVersions**](ApplicationVersionsAPI.md#ListApplicationVersions) | **Get** /workspace/applications/{resource_pk}/versions | List Application versions
[**PartialUpdateApplicationVersion**](ApplicationVersionsAPI.md#PartialUpdateApplicationVersion) | **Patch** /workspace/applications/{resource_pk}/versions/{id} | Partially update an Application version
[**RetrieveApplicationVersion**](ApplicationVersionsAPI.md#RetrieveApplicationVersion) | **Get** /workspace/applications/{resource_pk}/versions/{id} | Retrieve an Application version
[**UpdateApplicationVersion**](ApplicationVersionsAPI.md#UpdateApplicationVersion) | **Put** /workspace/applications/{resource_pk}/versions/{id} | Update an Application version



## ArchiveApplicationVersion

> ArchiveApplicationVersion(ctx, id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()

Archive an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	versionArchiveRequest := *openapiclient.NewVersionArchiveRequest() // VersionArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.ArchiveApplicationVersion(context.Background(), id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.ArchiveApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionArchiveRequest** | [**VersionArchiveRequest**](VersionArchiveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildApplicationVersion

> BuildApplicationVersion(ctx, id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()

Build an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	versionBuildRequest := *openapiclient.NewVersionBuildRequest() // VersionBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.BuildApplicationVersion(context.Background(), id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.BuildApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionBuildRequest** | [**VersionBuildRequest**](VersionBuildRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelApplicationVersionBuild

> CancelApplicationVersionBuild(ctx, id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()

Cancel an Application version build



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	versionBuildRequest := *openapiclient.NewVersionBuildRequest() // VersionBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.CancelApplicationVersionBuild(context.Background(), id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.CancelApplicationVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelApplicationVersionBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionBuildRequest** | [**VersionBuildRequest**](VersionBuildRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationVersion

> CreateApplicationVersion(ctx, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Create a new Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.CreateApplicationVersion(context.Background(), resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.CreateApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionCreateRequest** | [**VersionCreateRequest**](VersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationVersion

> DeleteApplicationVersion(ctx, id, resourcePk).Execute()

Delete an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.DeleteApplicationVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.DeleteApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationVersionRequest struct via the builder pattern


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


## ListApplicationVersions

> ListApplicationVersions(ctx, resourcePk).Fields(fields).Execute()

List Application versions



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.ListApplicationVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.ListApplicationVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. | 

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


## PartialUpdateApplicationVersion

> PartialUpdateApplicationVersion(ctx, id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()

Partially update an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	patchedVersionCreateRequest := *openapiclient.NewPatchedVersionCreateRequest() // PatchedVersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.PartialUpdateApplicationVersion(context.Background(), id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.PartialUpdateApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedVersionCreateRequest** | [**PatchedVersionCreateRequest**](PatchedVersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveApplicationVersion

> RetrieveApplicationVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.RetrieveApplicationVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.RetrieveApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. | 

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


## UpdateApplicationVersion

> UpdateApplicationVersion(ctx, id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Update an Application version



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
	resourcePk := int64(789) // int64 | The ID (global_id) of the Application resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationVersionsAPI.UpdateApplicationVersion(context.Background(), id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationVersionsAPI.UpdateApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID (global_id) of the Application resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionCreateRequest** | [**VersionCreateRequest**](VersionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

