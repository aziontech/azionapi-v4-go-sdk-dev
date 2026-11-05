# \EdgeDNSRecordsAPI

All URIs are relative to *https://stage-api.azion.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsRecord**](EdgeDNSRecordsAPI.md#CreateDnsRecord) | **Post** /workspace/dns/zones/{zone_id}/records | Create a DNS Record
[**DeleteDnsRecord**](EdgeDNSRecordsAPI.md#DeleteDnsRecord) | **Delete** /workspace/dns/zones/{zone_id}/records/{record_id} | Delete a DNS Record
[**ListDnsRecords**](EdgeDNSRecordsAPI.md#ListDnsRecords) | **Get** /workspace/dns/zones/{zone_id}/records | List DNS Records
[**PartialUpdateDnsRecord**](EdgeDNSRecordsAPI.md#PartialUpdateDnsRecord) | **Patch** /workspace/dns/zones/{zone_id}/records/{record_id} | Partially update a DNS Record
[**RetrieveDnsRecord**](EdgeDNSRecordsAPI.md#RetrieveDnsRecord) | **Get** /workspace/dns/zones/{zone_id}/records/{record_id} | Retrieve details of a DNS Record
[**UpdateDnsRecord**](EdgeDNSRecordsAPI.md#UpdateDnsRecord) | **Put** /workspace/dns/zones/{zone_id}/records/{record_id} | Update a DNS Record



## CreateDnsRecord

> RecordResponse CreateDnsRecord(ctx, zoneId).RecordRequest(recordRequest).Execute()

Create a DNS Record



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
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.
	recordRequest := *openapiclient.NewRecordRequest("Name_example", "Type_example", []string{"Rdata_example"}) // RecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.CreateDnsRecord(context.Background(), zoneId).RecordRequest(recordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.CreateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsRecord`: RecordResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.CreateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordRequest** | [**RecordRequest**](RecordRequest.md) |  | 

### Return type

[**RecordResponse**](RecordResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnsRecord

> DeleteResponse DeleteDnsRecord(ctx, recordId, zoneId).Execute()

Delete a DNS Record



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
	recordId := int64(789) // int64 | A unique integer value identifying the DNS Record.
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.DeleteDnsRecord(context.Background(), recordId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.DeleteDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDnsRecord`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.DeleteDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int64** | A unique integer value identifying the DNS Record. | 
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnsRecords

> PaginatedRecordList ListDnsRecords(ctx, zoneId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List DNS Records



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
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.ListDnsRecords(context.Background(), zoneId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.ListDnsRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDnsRecords`: PaginatedRecordList
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.ListDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRecordList**](PaginatedRecordList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDnsRecord

> RecordResponse PartialUpdateDnsRecord(ctx, recordId, zoneId).PatchedRecordRequest(patchedRecordRequest).Execute()

Partially update a DNS Record



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
	recordId := int64(789) // int64 | A unique integer value identifying the DNS Record.
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.
	patchedRecordRequest := *openapiclient.NewPatchedRecordRequest() // PatchedRecordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.PartialUpdateDnsRecord(context.Background(), recordId, zoneId).PatchedRecordRequest(patchedRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.PartialUpdateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDnsRecord`: RecordResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.PartialUpdateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int64** | A unique integer value identifying the DNS Record. | 
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedRecordRequest** | [**PatchedRecordRequest**](PatchedRecordRequest.md) |  | 

### Return type

[**RecordResponse**](RecordResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDnsRecord

> RecordResponse RetrieveDnsRecord(ctx, recordId, zoneId).Fields(fields).Execute()

Retrieve details of a DNS Record



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
	recordId := int64(789) // int64 | A unique integer value identifying the DNS Record.
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.RetrieveDnsRecord(context.Background(), recordId, zoneId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.RetrieveDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDnsRecord`: RecordResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.RetrieveDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int64** | A unique integer value identifying the DNS Record. | 
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. Nested fields can be accessed using dot notation. | 

### Return type

[**RecordResponse**](RecordResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDnsRecord

> RecordResponse UpdateDnsRecord(ctx, recordId, zoneId).RecordRequest(recordRequest).Execute()

Update a DNS Record



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
	recordId := int64(789) // int64 | A unique integer value identifying the DNS Record.
	zoneId := int64(789) // int64 | A unique integer value identifying the DNS Zone.
	recordRequest := *openapiclient.NewRecordRequest("Name_example", "Type_example", []string{"Rdata_example"}) // RecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeDNSRecordsAPI.UpdateDnsRecord(context.Background(), recordId, zoneId).RecordRequest(recordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeDNSRecordsAPI.UpdateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsRecord`: RecordResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeDNSRecordsAPI.UpdateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int64** | A unique integer value identifying the DNS Record. | 
**zoneId** | **int64** | A unique integer value identifying the DNS Zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordRequest** | [**RecordRequest**](RecordRequest.md) |  | 

### Return type

[**RecordResponse**](RecordResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

