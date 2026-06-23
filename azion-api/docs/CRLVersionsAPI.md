# \CRLVersionsAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveCrlVersion**](CRLVersionsAPI.md#ArchiveCrlVersion) | **Post** /workspace/tls/crls/{resource_pk}/versions/{id}/archive | Archive a CRL version
[**CancelCrlVersionBuild**](CRLVersionsAPI.md#CancelCrlVersionBuild) | **Post** /workspace/tls/crls/{resource_pk}/versions/{id}/cancel | Cancel a CRL version build
[**CreateCrlVersion**](CRLVersionsAPI.md#CreateCrlVersion) | **Post** /workspace/tls/crls/{resource_pk}/versions | Create a new CRL version
[**DeleteCrlVersion**](CRLVersionsAPI.md#DeleteCrlVersion) | **Delete** /workspace/tls/crls/{resource_pk}/versions/{id} | Delete a CRL version
[**ListCrlVersions**](CRLVersionsAPI.md#ListCrlVersions) | **Get** /workspace/tls/crls/{resource_pk}/versions | List CRL versions
[**PartialUpdateCrlVersion**](CRLVersionsAPI.md#PartialUpdateCrlVersion) | **Patch** /workspace/tls/crls/{resource_pk}/versions/{id} | Partially update a CRL version
[**RetrieveCrlVersion**](CRLVersionsAPI.md#RetrieveCrlVersion) | **Get** /workspace/tls/crls/{resource_pk}/versions/{id} | Retrieve a CRL version
[**UpdateCrlVersion**](CRLVersionsAPI.md#UpdateCrlVersion) | **Put** /workspace/tls/crls/{resource_pk}/versions/{id} | Update a CRL version



## ArchiveCrlVersion

> ArchiveCrlVersion(ctx, id, resourcePk).VersionArchive(versionArchive).Execute()

Archive a CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	versionArchive := *openapiclient.NewVersionArchive() // VersionArchive |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.ArchiveCrlVersion(context.Background(), id, resourcePk).VersionArchive(versionArchive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.ArchiveCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveCrlVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionArchive** | [**VersionArchive**](VersionArchive.md) |  | 

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


## CancelCrlVersionBuild

> CancelCrlVersionBuild(ctx, id, resourcePk).VersionCancel(versionCancel).Execute()

Cancel a CRL version build



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	versionCancel := *openapiclient.NewVersionCancel() // VersionCancel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.CancelCrlVersionBuild(context.Background(), id, resourcePk).VersionCancel(versionCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.CancelCrlVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCrlVersionBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionCancel** | [**VersionCancel**](VersionCancel.md) |  | 

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


## CreateCrlVersion

> CreateCrlVersion(ctx, resourcePk).CRLVersionCreate(cRLVersionCreate).Execute()

Create a new CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	cRLVersionCreate := *openapiclient.NewCRLVersionCreate() // CRLVersionCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.CreateCrlVersion(context.Background(), resourcePk).CRLVersionCreate(cRLVersionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.CreateCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrlVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cRLVersionCreate** | [**CRLVersionCreate**](CRLVersionCreate.md) |  | 

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


## DeleteCrlVersion

> DeleteCrlVersion(ctx, id, resourcePk).Execute()

Delete a CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.DeleteCrlVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.DeleteCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrlVersionRequest struct via the builder pattern


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


## ListCrlVersions

> ListCrlVersions(ctx, resourcePk).Fields(fields).Execute()

List CRL versions



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.ListCrlVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.ListCrlVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCrlVersionsRequest struct via the builder pattern


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


## PartialUpdateCrlVersion

> PartialUpdateCrlVersion(ctx, id, resourcePk).PatchedCertificateRevocationList(patchedCertificateRevocationList).Execute()

Partially update a CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	patchedCertificateRevocationList := *openapiclient.NewPatchedCertificateRevocationList() // PatchedCertificateRevocationList |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.PartialUpdateCrlVersion(context.Background(), id, resourcePk).PatchedCertificateRevocationList(patchedCertificateRevocationList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.PartialUpdateCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCrlVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedCertificateRevocationList** | [**PatchedCertificateRevocationList**](PatchedCertificateRevocationList.md) |  | 

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


## RetrieveCrlVersion

> RetrieveCrlVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve a CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.RetrieveCrlVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.RetrieveCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCrlVersionRequest struct via the builder pattern


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


## UpdateCrlVersion

> UpdateCrlVersion(ctx, id, resourcePk).CertificateRevocationList(certificateRevocationList).Execute()

Update a CRL version



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
	resourcePk := int64(789) // int64 | The ID of the CRL resource.
	certificateRevocationList := *openapiclient.NewCertificateRevocationList("Name_example", "Crl_example") // CertificateRevocationList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CRLVersionsAPI.UpdateCrlVersion(context.Background(), id, resourcePk).CertificateRevocationList(certificateRevocationList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRLVersionsAPI.UpdateCrlVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the CRL resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCrlVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certificateRevocationList** | [**CertificateRevocationList**](CertificateRevocationList.md) |  | 

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

