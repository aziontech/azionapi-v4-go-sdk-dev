# PatchApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to [**AppModulesRequest**](AppModulesRequest.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchApplicationRequest

`func NewPatchApplicationRequest() *PatchApplicationRequest`

NewPatchApplicationRequest instantiates a new PatchApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationRequestWithDefaults

`func NewPatchApplicationRequestWithDefaults() *PatchApplicationRequest`

NewPatchApplicationRequestWithDefaults instantiates a new PatchApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchApplicationRequest) GetModules() AppModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchApplicationRequest) GetModulesOk() (*AppModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchApplicationRequest) SetModules(v AppModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchApplicationRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *PatchApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebug

`func (o *PatchApplicationRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *PatchApplicationRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *PatchApplicationRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *PatchApplicationRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


