# PatchedCustomPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Pages** | Pointer to [**[]ItemPageRequest**](ItemPageRequest.md) |  | [optional] 

## Methods

### NewPatchedCustomPagesRequest

`func NewPatchedCustomPagesRequest() *PatchedCustomPagesRequest`

NewPatchedCustomPagesRequest instantiates a new PatchedCustomPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCustomPagesRequestWithDefaults

`func NewPatchedCustomPagesRequestWithDefaults() *PatchedCustomPagesRequest`

NewPatchedCustomPagesRequestWithDefaults instantiates a new PatchedCustomPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCustomPagesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCustomPagesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCustomPagesRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCustomPagesRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedCustomPagesRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedCustomPagesRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedCustomPagesRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedCustomPagesRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPages

`func (o *PatchedCustomPagesRequest) GetPages() []ItemPageRequest`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PatchedCustomPagesRequest) GetPagesOk() (*[]ItemPageRequest, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PatchedCustomPagesRequest) SetPages(v []ItemPageRequest)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PatchedCustomPagesRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


