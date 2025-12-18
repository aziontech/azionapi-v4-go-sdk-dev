# AppRuleRunFunctionResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**AppRuleRunFunctionResponseAttrsReq**](AppRuleRunFunctionResponseAttrsReq.md) |  | 

## Methods

### NewAppRuleRunFunctionResponseRequest

`func NewAppRuleRunFunctionResponseRequest(type_ string, attributes AppRuleRunFunctionResponseAttrsReq, ) *AppRuleRunFunctionResponseRequest`

NewAppRuleRunFunctionResponseRequest instantiates a new AppRuleRunFunctionResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleRunFunctionResponseRequestWithDefaults

`func NewAppRuleRunFunctionResponseRequestWithDefaults() *AppRuleRunFunctionResponseRequest`

NewAppRuleRunFunctionResponseRequestWithDefaults instantiates a new AppRuleRunFunctionResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleRunFunctionResponseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleRunFunctionResponseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleRunFunctionResponseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleRunFunctionResponseRequest) GetAttributes() AppRuleRunFunctionResponseAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleRunFunctionResponseRequest) GetAttributesOk() (*AppRuleRunFunctionResponseAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleRunFunctionResponseRequest) SetAttributes(v AppRuleRunFunctionResponseAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


