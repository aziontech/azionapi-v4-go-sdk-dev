# \WorkloadVersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkloadVersion**](WorkloadVersionsAPI.md#ArchiveWorkloadVersion) | **Post** /workspace/workloads/{resource_pk}/versions/{id}/archive | Archive a Workload version
[**BuildWorkloadVersion**](WorkloadVersionsAPI.md#BuildWorkloadVersion) | **Post** /workspace/workloads/{resource_pk}/versions/{id}/build | Build a Workload version
[**CancelWorkloadVersionBuild**](WorkloadVersionsAPI.md#CancelWorkloadVersionBuild) | **Post** /workspace/workloads/{resource_pk}/versions/{id}/cancel | Cancel a Workload version build
[**CreateWorkloadVersion**](WorkloadVersionsAPI.md#CreateWorkloadVersion) | **Post** /workspace/workloads/{resource_pk}/versions | Create a new Workload version
[**DeleteWorkloadVersion**](WorkloadVersionsAPI.md#DeleteWorkloadVersion) | **Delete** /workspace/workloads/{resource_pk}/versions/{id} | Delete a Workload version
[**ListWorkloadVersions**](WorkloadVersionsAPI.md#ListWorkloadVersions) | **Get** /workspace/workloads/{resource_pk}/versions | List Workload versions
[**PartialUpdateWorkloadVersion**](WorkloadVersionsAPI.md#PartialUpdateWorkloadVersion) | **Patch** /workspace/workloads/{resource_pk}/versions/{id} | Partially update a Workload version
[**RetrieveWorkloadVersion**](WorkloadVersionsAPI.md#RetrieveWorkloadVersion) | **Get** /workspace/workloads/{resource_pk}/versions/{id} | Retrieve a Workload version
[**RollbackWorkloadVersion**](WorkloadVersionsAPI.md#RollbackWorkloadVersion) | **Post** /workspace/workloads/{resource_pk}/versions/{id}/rollback | Rollback to a Workload version
[**UpdateWorkloadVersion**](WorkloadVersionsAPI.md#UpdateWorkloadVersion) | **Put** /workspace/workloads/{resource_pk}/versions/{id} | Update a Workload version



## ArchiveWorkloadVersion

> ArchiveWorkloadVersion(ctx, id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()

Archive a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	versionArchiveRequest := *openapiclient.NewVersionArchiveRequest() // VersionArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.ArchiveWorkloadVersion(context.Background(), id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.ArchiveWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveWorkloadVersionRequest struct via the builder pattern


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


## BuildWorkloadVersion

> BuildWorkloadVersion(ctx, id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()

Build a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	versionBuildRequest := *openapiclient.NewVersionBuildRequest() // VersionBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.BuildWorkloadVersion(context.Background(), id, resourcePk).VersionBuildRequest(versionBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.BuildWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildWorkloadVersionRequest struct via the builder pattern


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


## CancelWorkloadVersionBuild

> CancelWorkloadVersionBuild(ctx, id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()

Cancel a Workload version build



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	versionCancelRequest := *openapiclient.NewVersionCancelRequest() // VersionCancelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.CancelWorkloadVersionBuild(context.Background(), id, resourcePk).VersionCancelRequest(versionCancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.CancelWorkloadVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkloadVersionBuildRequest struct via the builder pattern


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


## CreateWorkloadVersion

> CreateWorkloadVersion(ctx, resourcePk).VersionCreateRequest(versionCreateRequest).Execute()

Create a new Workload version



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
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	versionCreateRequest := *openapiclient.NewVersionCreateRequest() // VersionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.CreateWorkloadVersion(context.Background(), resourcePk).VersionCreateRequest(versionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.CreateWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadVersionRequest struct via the builder pattern


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


## DeleteWorkloadVersion

> DeleteWorkloadVersion(ctx, id, resourcePk).Execute()

Delete a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.DeleteWorkloadVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.DeleteWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadVersionRequest struct via the builder pattern


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


## ListWorkloadVersions

> ListWorkloadVersions(ctx, resourcePk).Fields(fields).Execute()

List Workload versions



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
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.ListWorkloadVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.ListWorkloadVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadVersionsRequest struct via the builder pattern


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


## PartialUpdateWorkloadVersion

> PartialUpdateWorkloadVersion(ctx, id, resourcePk).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()

Partially update a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	patchedWorkloadRequest := *openapiclient.NewPatchedWorkloadRequest() // PatchedWorkloadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.PartialUpdateWorkloadVersion(context.Background(), id, resourcePk).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.PartialUpdateWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWorkloadVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedWorkloadRequest** | [**PatchedWorkloadRequest**](PatchedWorkloadRequest.md) |  | 

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


## RetrieveWorkloadVersion

> RetrieveWorkloadVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.RetrieveWorkloadVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.RetrieveWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWorkloadVersionRequest struct via the builder pattern


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


## RollbackWorkloadVersion

> RollbackWorkloadVersion(ctx, id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()

Rollback to a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	versionArchiveRequest := *openapiclient.NewVersionArchiveRequest() // VersionArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.RollbackWorkloadVersion(context.Background(), id, resourcePk).VersionArchiveRequest(versionArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.RollbackWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackWorkloadVersionRequest struct via the builder pattern


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


## UpdateWorkloadVersion

> UpdateWorkloadVersion(ctx, id, resourcePk).WorkloadRequest(workloadRequest).Execute()

Update a Workload version



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
	id := "id_example" // string | The short ID identifier of the version.
	resourcePk := int64(789) // int64 | The global_id of the Workload resource.
	workloadRequest := *openapiclient.NewWorkloadRequest("Name_example") // WorkloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadVersionsAPI.UpdateWorkloadVersion(context.Background(), id, resourcePk).WorkloadRequest(workloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadVersionsAPI.UpdateWorkloadVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The short ID identifier of the version. | 
**resourcePk** | **int64** | The global_id of the Workload resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workloadRequest** | [**WorkloadRequest**](WorkloadRequest.md) |  | 

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

