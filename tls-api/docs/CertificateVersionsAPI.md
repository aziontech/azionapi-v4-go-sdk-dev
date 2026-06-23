# \CertificateVersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveCertificateVersion**](CertificateVersionsAPI.md#ArchiveCertificateVersion) | **Post** /workspace/tls/certificates/{resource_pk}/versions/{id}/archive | Archive a certificate version
[**CancelCertificateVersionBuild**](CertificateVersionsAPI.md#CancelCertificateVersionBuild) | **Post** /workspace/tls/certificates/{resource_pk}/versions/{id}/cancel | Cancel a certificate version build
[**CreateCertificateVersion**](CertificateVersionsAPI.md#CreateCertificateVersion) | **Post** /workspace/tls/certificates/{resource_pk}/versions | Create a new certificate version
[**DeleteCertificateVersion**](CertificateVersionsAPI.md#DeleteCertificateVersion) | **Delete** /workspace/tls/certificates/{resource_pk}/versions/{id} | Delete a certificate version
[**ListCertificateVersions**](CertificateVersionsAPI.md#ListCertificateVersions) | **Get** /workspace/tls/certificates/{resource_pk}/versions | List certificate versions
[**PartialUpdateCertificateVersion**](CertificateVersionsAPI.md#PartialUpdateCertificateVersion) | **Patch** /workspace/tls/certificates/{resource_pk}/versions/{id} | Partially update a certificate version
[**RetrieveCertificateVersion**](CertificateVersionsAPI.md#RetrieveCertificateVersion) | **Get** /workspace/tls/certificates/{resource_pk}/versions/{id} | Retrieve a certificate version
[**UpdateCertificateVersion**](CertificateVersionsAPI.md#UpdateCertificateVersion) | **Put** /workspace/tls/certificates/{resource_pk}/versions/{id} | Update a certificate version



## ArchiveCertificateVersion

> ArchiveCertificateVersion(ctx, id, resourcePk).VersionArchive(versionArchive).Execute()

Archive a certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	versionArchive := *openapiclient.NewVersionArchive() // VersionArchive |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.ArchiveCertificateVersion(context.Background(), id, resourcePk).VersionArchive(versionArchive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.ArchiveCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveCertificateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionArchive** | [**VersionArchive**](VersionArchive.md) |  | 

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


## CancelCertificateVersionBuild

> CancelCertificateVersionBuild(ctx, id, resourcePk).VersionCancel(versionCancel).Execute()

Cancel a certificate version build



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	versionCancel := *openapiclient.NewVersionCancel() // VersionCancel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.CancelCertificateVersionBuild(context.Background(), id, resourcePk).VersionCancel(versionCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.CancelCertificateVersionBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCertificateVersionBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionCancel** | [**VersionCancel**](VersionCancel.md) |  | 

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


## CreateCertificateVersion

> CreateCertificateVersion(ctx, resourcePk).CertificateVersionCreate(certificateVersionCreate).Execute()

Create a new certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	certificateVersionCreate := *openapiclient.NewCertificateVersionCreate() // CertificateVersionCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.CreateCertificateVersion(context.Background(), resourcePk).CertificateVersionCreate(certificateVersionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.CreateCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateVersionCreate** | [**CertificateVersionCreate**](CertificateVersionCreate.md) |  | 

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


## DeleteCertificateVersion

> DeleteCertificateVersion(ctx, id, resourcePk).Execute()

Delete a certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.DeleteCertificateVersion(context.Background(), id, resourcePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.DeleteCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateVersionRequest struct via the builder pattern


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


## ListCertificateVersions

> ListCertificateVersions(ctx, resourcePk).Fields(fields).Execute()

List certificate versions



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.ListCertificateVersions(context.Background(), resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.ListCertificateVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCertificateVersionsRequest struct via the builder pattern


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


## PartialUpdateCertificateVersion

> PartialUpdateCertificateVersion(ctx, id, resourcePk).PatchedCertificate(patchedCertificate).Execute()

Partially update a certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	patchedCertificate := *openapiclient.NewPatchedCertificate() // PatchedCertificate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.PartialUpdateCertificateVersion(context.Background(), id, resourcePk).PatchedCertificate(patchedCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.PartialUpdateCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCertificateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedCertificate** | [**PatchedCertificate**](PatchedCertificate.md) |  | 

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


## RetrieveCertificateVersion

> RetrieveCertificateVersion(ctx, id, resourcePk).Fields(fields).Execute()

Retrieve a certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.RetrieveCertificateVersion(context.Background(), id, resourcePk).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.RetrieveCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCertificateVersionRequest struct via the builder pattern


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


## UpdateCertificateVersion

> UpdateCertificateVersion(ctx, id, resourcePk).Certificate(certificate).Execute()

Update a certificate version



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
	resourcePk := int64(789) // int64 | The ID of the certificate resource.
	certificate := *openapiclient.NewCertificate("Name_example") // Certificate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateVersionsAPI.UpdateCertificateVersion(context.Background(), id, resourcePk).Certificate(certificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateVersionsAPI.UpdateCertificateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ULID identifier of the version. | 
**resourcePk** | **int64** | The ID of the certificate resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certificate** | [**Certificate**](Certificate.md) |  | 

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

