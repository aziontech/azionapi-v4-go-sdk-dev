# PaginatedResponseListIntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ResponseListIntegration**](ResponseListIntegration.md) |  | [optional] 

## Methods

### NewPaginatedResponseListIntegrationList

`func NewPaginatedResponseListIntegrationList() *PaginatedResponseListIntegrationList`

NewPaginatedResponseListIntegrationList instantiates a new PaginatedResponseListIntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListIntegrationListWithDefaults

`func NewPaginatedResponseListIntegrationListWithDefaults() *PaginatedResponseListIntegrationList`

NewPaginatedResponseListIntegrationListWithDefaults instantiates a new PaginatedResponseListIntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListIntegrationList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListIntegrationList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListIntegrationList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListIntegrationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListIntegrationList) GetResults() []ResponseListIntegration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListIntegrationList) GetResultsOk() (*[]ResponseListIntegration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListIntegrationList) SetResults(v []ResponseListIntegration)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListIntegrationList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


