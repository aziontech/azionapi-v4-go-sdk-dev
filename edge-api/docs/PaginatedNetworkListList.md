# PaginatedNetworkListList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]NetworkList**](NetworkList.md) |  | [optional] 

## Methods

### NewPaginatedNetworkListList

`func NewPaginatedNetworkListList() *PaginatedNetworkListList`

NewPaginatedNetworkListList instantiates a new PaginatedNetworkListList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedNetworkListListWithDefaults

`func NewPaginatedNetworkListListWithDefaults() *PaginatedNetworkListList`

NewPaginatedNetworkListListWithDefaults instantiates a new PaginatedNetworkListList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedNetworkListList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedNetworkListList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedNetworkListList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedNetworkListList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedNetworkListList) GetResults() []NetworkList`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedNetworkListList) GetResultsOk() (*[]NetworkList, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedNetworkListList) SetResults(v []NetworkList)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedNetworkListList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


