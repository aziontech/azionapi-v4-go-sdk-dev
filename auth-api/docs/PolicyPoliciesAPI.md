# \PolicyPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](PolicyPoliciesAPI.md#CreatePolicy) | **Post** /auth/policies | Create a new policy
[**DeletePolicy**](PolicyPoliciesAPI.md#DeletePolicy) | **Delete** /auth/policies/{policy_id} | Delete a policy
[**ListPolicy**](PolicyPoliciesAPI.md#ListPolicy) | **Get** /auth/policies | List of account policies
[**PartialUpdatePolicy**](PolicyPoliciesAPI.md#PartialUpdatePolicy) | **Patch** /auth/policies/{policy_id} | Partially update a policy
[**RetrievePolicy**](PolicyPoliciesAPI.md#RetrievePolicy) | **Get** /auth/policies/{policy_id} | Retrieve details from a policy
[**UpdatePolicy**](PolicyPoliciesAPI.md#UpdatePolicy) | **Put** /auth/policies/{policy_id} | Update a policy



## CreatePolicy

> ResponsePolicy CreatePolicy(ctx).PolicyRequest(policyRequest).Execute()

Create a new policy



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
	policyRequest := *openapiclient.NewPolicyRequest("Name_example", false, []openapiclient.PolicyRuleRequest{*openapiclient.NewPolicyRuleRequest("Name_example", "Effect_example", "Resource_example", []string{"Actions_example"}, *openapiclient.NewPolicyRuleConditionRequest([]string{"IpAddress_example"}))}) // PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.CreatePolicy(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: ResponsePolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyRequest** | [**PolicyRequest**](PolicyRequest.md) |  | 

### Return type

[**ResponsePolicy**](ResponsePolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> ResponseDeletePolicy DeletePolicy(ctx, policyId).Execute()

Delete a policy



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
	policyId := int64(789) // int64 | A unique integer value identifying the policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.DeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: ResponseDeletePolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | A unique integer value identifying the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeletePolicy**](ResponseDeletePolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicy

> PaginatedPolicyList ListPolicy(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of account policies



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.ListPolicy(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.ListPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicy`: PaginatedPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.ListPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPolicyList**](PaginatedPolicyList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdatePolicy

> ResponsePolicy PartialUpdatePolicy(ctx, policyId).PatchPolicyRequest(patchPolicyRequest).Execute()

Partially update a policy



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
	policyId := int64(789) // int64 | A unique integer value identifying the policy.
	patchPolicyRequest := *openapiclient.NewPatchPolicyRequest() // PatchPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.PartialUpdatePolicy(context.Background(), policyId).PatchPolicyRequest(patchPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.PartialUpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdatePolicy`: ResponsePolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.PartialUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | A unique integer value identifying the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPolicyRequest** | [**PatchPolicyRequest**](PatchPolicyRequest.md) |  | 

### Return type

[**ResponsePolicy**](ResponsePolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePolicy

> ResponseRetrievePolicy RetrievePolicy(ctx, policyId).Fields(fields).Execute()

Retrieve details from a policy



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
	policyId := int64(789) // int64 | A unique integer value identifying the policy.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.RetrievePolicy(context.Background(), policyId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.RetrievePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievePolicy`: ResponseRetrievePolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.RetrievePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | A unique integer value identifying the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrievePolicy**](ResponseRetrievePolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> ResponsePolicy UpdatePolicy(ctx, policyId).PolicyRequest(policyRequest).Execute()

Update a policy



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
	policyId := int64(789) // int64 | A unique integer value identifying the policy.
	policyRequest := *openapiclient.NewPolicyRequest("Name_example", false, []openapiclient.PolicyRuleRequest{*openapiclient.NewPolicyRuleRequest("Name_example", "Effect_example", "Resource_example", []string{"Actions_example"}, *openapiclient.NewPolicyRuleConditionRequest([]string{"IpAddress_example"}))}) // PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPoliciesAPI.UpdatePolicy(context.Background(), policyId).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPoliciesAPI.UpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicy`: ResponsePolicy
	fmt.Fprintf(os.Stdout, "Response from `PolicyPoliciesAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | A unique integer value identifying the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRequest** | [**PolicyRequest**](PolicyRequest.md) |  | 

### Return type

[**ResponsePolicy**](ResponsePolicy.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

