# DocsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountPerSearch** | Pointer to **int64** | Maximum number of results to return per search | [optional] 
**MetadataItems** | Pointer to **[]string** | Metadata fields to include in the search results | [optional] 
**Filters** | Pointer to [**[]DocsOptionsFilter**](DocsOptionsFilter.md) | Filters applied to narrow down the documentation search | [optional] 

## Methods

### NewDocsOptions

`func NewDocsOptions() *DocsOptions`

NewDocsOptions instantiates a new DocsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsOptionsWithDefaults

`func NewDocsOptionsWithDefaults() *DocsOptions`

NewDocsOptionsWithDefaults instantiates a new DocsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountPerSearch

`func (o *DocsOptions) GetAmountPerSearch() int64`

GetAmountPerSearch returns the AmountPerSearch field if non-nil, zero value otherwise.

### GetAmountPerSearchOk

`func (o *DocsOptions) GetAmountPerSearchOk() (*int64, bool)`

GetAmountPerSearchOk returns a tuple with the AmountPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPerSearch

`func (o *DocsOptions) SetAmountPerSearch(v int64)`

SetAmountPerSearch sets AmountPerSearch field to given value.

### HasAmountPerSearch

`func (o *DocsOptions) HasAmountPerSearch() bool`

HasAmountPerSearch returns a boolean if a field has been set.

### GetMetadataItems

`func (o *DocsOptions) GetMetadataItems() []string`

GetMetadataItems returns the MetadataItems field if non-nil, zero value otherwise.

### GetMetadataItemsOk

`func (o *DocsOptions) GetMetadataItemsOk() (*[]string, bool)`

GetMetadataItemsOk returns a tuple with the MetadataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataItems

`func (o *DocsOptions) SetMetadataItems(v []string)`

SetMetadataItems sets MetadataItems field to given value.

### HasMetadataItems

`func (o *DocsOptions) HasMetadataItems() bool`

HasMetadataItems returns a boolean if a field has been set.

### GetFilters

`func (o *DocsOptions) GetFilters() []DocsOptionsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DocsOptions) GetFiltersOk() (*[]DocsOptionsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DocsOptions) SetFilters(v []DocsOptionsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DocsOptions) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


