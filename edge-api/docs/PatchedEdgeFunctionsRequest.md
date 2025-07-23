# PatchedEdgeFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | * &#x60;javascript&#x60; - JavaScript | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**DefaultArgs** | Pointer to [**EdgeFunctionsDefaultArgs**](EdgeFunctionsDefaultArgs.md) |  | [optional] 
**InitiatorType** | Pointer to **string** | * &#x60;edge_application&#x60; - Edge Application * &#x60;edge_firewall&#x60; - Edge Firewall | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEdgeFunctionsRequest

`func NewPatchedEdgeFunctionsRequest() *PatchedEdgeFunctionsRequest`

NewPatchedEdgeFunctionsRequest instantiates a new PatchedEdgeFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeFunctionsRequestWithDefaults

`func NewPatchedEdgeFunctionsRequestWithDefaults() *PatchedEdgeFunctionsRequest`

NewPatchedEdgeFunctionsRequestWithDefaults instantiates a new PatchedEdgeFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeFunctionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *PatchedEdgeFunctionsRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PatchedEdgeFunctionsRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PatchedEdgeFunctionsRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PatchedEdgeFunctionsRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *PatchedEdgeFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchedEdgeFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchedEdgeFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PatchedEdgeFunctionsRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefaultArgs

`func (o *PatchedEdgeFunctionsRequest) GetDefaultArgs() EdgeFunctionsDefaultArgs`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *PatchedEdgeFunctionsRequest) GetDefaultArgsOk() (*EdgeFunctionsDefaultArgs, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *PatchedEdgeFunctionsRequest) SetDefaultArgs(v EdgeFunctionsDefaultArgs)`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *PatchedEdgeFunctionsRequest) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### GetInitiatorType

`func (o *PatchedEdgeFunctionsRequest) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *PatchedEdgeFunctionsRequest) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *PatchedEdgeFunctionsRequest) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *PatchedEdgeFunctionsRequest) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


