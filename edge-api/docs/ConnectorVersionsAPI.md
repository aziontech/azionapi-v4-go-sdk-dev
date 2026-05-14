# \ConnectorVersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveConnectorVersion**](ConnectorVersionsAPI.md#ArchiveConnectorVersion) | **Post** /workspace/connectors/{resource_pk}/versions/{id}/archive | Archive a Connector version
[**BuildConnectorVersion**](ConnectorVersionsAPI.md#BuildConnectorVersion) | **Post** /workspace/connectors/{resource_pk}/versions/{id}/build | Build a Connector version
[**CancelConnectorVersionBuild**](ConnectorVersionsAPI.md#CancelConnectorVersionBuild) | **Post** /workspace/connectors/{resource_pk}/versions/{id}/cancel | Cancel a Connector version build
[**CreateConnectorVersion**](ConnectorVersionsAPI.md#CreateConnectorVersion) | **Post** /workspace/connectors/{resource_pk}/versions | Create a new Connector version
[**DeleteConnectorVersion**](ConnectorVersionsAPI.md#DeleteConnectorVersion) | **Delete** /workspace/connectors/{resource_pk}/versions/{id} | Delete a Connector version
[**ListConnectorVersions**](ConnectorVersionsAPI.md#ListConnectorVersions) | **Get** /workspace/connectors/{resource_pk}/versions | List Connector versions
[**PartialUpdateConnectorVersion**](ConnectorVersionsAPI.md#PartialUpdateConnectorVersion) | **Patch** /workspace/connectors/{resource_pk}/versions/{id} | Partially update a Connector version
[**RetrieveConnectorVersion**](ConnectorVersionsAPI.md#RetrieveConnectorVersion) | **Get** /workspace/connectors/{resource_pk}/versions/{id} | Retrieve a Connector version
[**UpdateConnectorVersion**](ConnectorVersionsAPI.md#UpdateConnectorVersion) | **Put** /workspace/connectors/{resource_pk}/versions/{id} | Update a Connector version



## ArchiveConnectorVersion

> ArchiveConnectorVersion(ctx, id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()

Archive a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	versionArchiveRequest := *openapiclient.NewVersionArchiveRequest() // VersionArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.ArchiveConnectorVersion(context.Background(), id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.ArchiveConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveConnectorVersionRequest struct via the builder pattern


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


## BuildConnectorVersion

> BuildConnectorVersion(ctx, id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()

Build a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	versionBuildRequest := *openapiclient.NewVersionBuildRequest() // VersionBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.BuildConnectorVersion(context.Background(), id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.BuildConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildConnectorVersionRequest struct via the builder pattern


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


## CancelConnectorVersionBuild

> CancelConnectorVersionBuild(ctx, id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()

Cancel a Connector version build



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	versionCancelRequest := *openapiclient.NewVersionCancelRequest() // VersionCancelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.CancelConnectorVersionBuild(context.Background(), id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.CancelConnectorVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelConnectorVersionBuildRequest struct via the builder pattern


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


## CreateConnectorVersion

> CreateConnectorVersion(ctx, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Create a new Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.CreateConnectorVersion(context.Background(), resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.CreateConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorVersionRequest struct via the builder pattern


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


## DeleteConnectorVersion

> DeleteConnectorVersion(ctx, id, resourcePk).Execute()

Delete a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.DeleteConnectorVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.DeleteConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorVersionRequest struct via the builder pattern


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


## ListConnectorVersions

> ListConnectorVersions(ctx, resourcePk).Fields(fields).Execute()

List Connector versions



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.ListConnectorVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.ListConnectorVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorVersionsRequest struct via the builder pattern


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


## PartialUpdateConnectorVersion

> PartialUpdateConnectorVersion(ctx, id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()

Partially update a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	patchedVersionCreateRequest := *openapiclient.NewPatchedVersionCreateRequest() // PatchedVersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.PartialUpdateConnectorVersion(context.Background(), id, resourcePk).PatchedVersionCreateRequest(patchedVersionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.PartialUpdateConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateConnectorVersionRequest struct via the builder pattern


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


## RetrieveConnectorVersion

> RetrieveConnectorVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.RetrieveConnectorVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.RetrieveConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConnectorVersionRequest struct via the builder pattern


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


## UpdateConnectorVersion

> UpdateConnectorVersion(ctx, id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Update a Connector version



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
	resourcePk := int64(789) // int64 | The ID of the Connector resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorVersionsAPI.UpdateConnectorVersion(context.Background(), id, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorVersionsAPI.UpdateConnectorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the Connector resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorVersionRequest struct via the builder pattern


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

