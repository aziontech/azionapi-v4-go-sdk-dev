# ScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewScopeRequest

`func NewScopeRequest(slug string, name string, ) *ScopeRequest`

NewScopeRequest instantiates a new ScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeRequestWithDefaults

`func NewScopeRequestWithDefaults() *ScopeRequest`

NewScopeRequestWithDefaults instantiates a new ScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *ScopeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ScopeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ScopeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *ScopeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ScopeRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScopeRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScopeRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScopeRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


