# NamespaceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]NamespaceResponse**](NamespaceResponse.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewNamespaceListResponse

`func NewNamespaceListResponse(results []NamespaceResponse, pagination Pagination, ) *NamespaceListResponse`

NewNamespaceListResponse instantiates a new NamespaceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceListResponseWithDefaults

`func NewNamespaceListResponseWithDefaults() *NamespaceListResponse`

NewNamespaceListResponseWithDefaults instantiates a new NamespaceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *NamespaceListResponse) GetResults() []NamespaceResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NamespaceListResponse) GetResultsOk() (*[]NamespaceResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NamespaceListResponse) SetResults(v []NamespaceResponse)`

SetResults sets Results field to given value.


### GetPagination

`func (o *NamespaceListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NamespaceListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NamespaceListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


