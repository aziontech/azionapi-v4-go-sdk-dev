# EdgeFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Language** | Pointer to **string** | * &#x60;javascript&#x60; - JavaScript | [optional] 
**Code** | **string** |  | 
**DefaultArgs** | Pointer to [**EdgeFunctionsDefaultArgs**](EdgeFunctionsDefaultArgs.md) |  | [optional] 
**InitiatorType** | Pointer to **string** | * &#x60;edge_application&#x60; - Edge Application * &#x60;edge_firewall&#x60; - Edge Firewall | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeFunctionsRequest

`func NewEdgeFunctionsRequest(name string, code string, ) *EdgeFunctionsRequest`

NewEdgeFunctionsRequest instantiates a new EdgeFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionsRequestWithDefaults

`func NewEdgeFunctionsRequestWithDefaults() *EdgeFunctionsRequest`

NewEdgeFunctionsRequestWithDefaults instantiates a new EdgeFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *EdgeFunctionsRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EdgeFunctionsRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EdgeFunctionsRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EdgeFunctionsRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *EdgeFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EdgeFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EdgeFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetDefaultArgs

`func (o *EdgeFunctionsRequest) GetDefaultArgs() EdgeFunctionsDefaultArgs`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *EdgeFunctionsRequest) GetDefaultArgsOk() (*EdgeFunctionsDefaultArgs, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *EdgeFunctionsRequest) SetDefaultArgs(v EdgeFunctionsDefaultArgs)`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *EdgeFunctionsRequest) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### GetInitiatorType

`func (o *EdgeFunctionsRequest) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *EdgeFunctionsRequest) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *EdgeFunctionsRequest) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *EdgeFunctionsRequest) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetActive

`func (o *EdgeFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


