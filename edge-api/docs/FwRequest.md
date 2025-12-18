# FwRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Modules** | Pointer to [**FwModulesRequest**](FwModulesRequest.md) |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewFwRequest

`func NewFwRequest(name string, ) *FwRequest`

NewFwRequest instantiates a new FwRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwRequestWithDefaults

`func NewFwRequestWithDefaults() *FwRequest`

NewFwRequestWithDefaults instantiates a new FwRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FwRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FwRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FwRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *FwRequest) GetModules() FwModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *FwRequest) GetModulesOk() (*FwModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *FwRequest) SetModules(v FwModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *FwRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebug

`func (o *FwRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *FwRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *FwRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *FwRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetActive

`func (o *FwRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FwRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FwRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FwRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


