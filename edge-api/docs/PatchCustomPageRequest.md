# PatchCustomPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Pages** | Pointer to [**[]PageRequest**](PageRequest.md) |  | [optional] 

## Methods

### NewPatchCustomPageRequest

`func NewPatchCustomPageRequest() *PatchCustomPageRequest`

NewPatchCustomPageRequest instantiates a new PatchCustomPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCustomPageRequestWithDefaults

`func NewPatchCustomPageRequestWithDefaults() *PatchCustomPageRequest`

NewPatchCustomPageRequestWithDefaults instantiates a new PatchCustomPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchCustomPageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchCustomPageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchCustomPageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchCustomPageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchCustomPageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchCustomPageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchCustomPageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchCustomPageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPages

`func (o *PatchCustomPageRequest) GetPages() []PageRequest`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PatchCustomPageRequest) GetPagesOk() (*[]PageRequest, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PatchCustomPageRequest) SetPages(v []PageRequest)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PatchCustomPageRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


