# PaginatedResponseListDashboardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ResponseListDashboard**](ResponseListDashboard.md) |  | [optional] 

## Methods

### NewPaginatedResponseListDashboardList

`func NewPaginatedResponseListDashboardList() *PaginatedResponseListDashboardList`

NewPaginatedResponseListDashboardList instantiates a new PaginatedResponseListDashboardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListDashboardListWithDefaults

`func NewPaginatedResponseListDashboardListWithDefaults() *PaginatedResponseListDashboardList`

NewPaginatedResponseListDashboardListWithDefaults instantiates a new PaginatedResponseListDashboardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListDashboardList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListDashboardList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListDashboardList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListDashboardList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListDashboardList) GetResults() []ResponseListDashboard`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListDashboardList) GetResultsOk() (*[]ResponseListDashboard, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListDashboardList) SetResults(v []ResponseListDashboard)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListDashboardList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


