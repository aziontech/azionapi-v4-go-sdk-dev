# ApplicationRequestPhaseBehaviorWithArgsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**ApplicationRuleEngineStringAttributes**](ApplicationRuleEngineStringAttributes.md) |  | 

## Methods

### NewApplicationRequestPhaseBehaviorWithArgsRequest

`func NewApplicationRequestPhaseBehaviorWithArgsRequest(type_ string, attributes ApplicationRuleEngineStringAttributes, ) *ApplicationRequestPhaseBehaviorWithArgsRequest`

NewApplicationRequestPhaseBehaviorWithArgsRequest instantiates a new ApplicationRequestPhaseBehaviorWithArgsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseBehaviorWithArgsRequestWithDefaults

`func NewApplicationRequestPhaseBehaviorWithArgsRequestWithDefaults() *ApplicationRequestPhaseBehaviorWithArgsRequest`

NewApplicationRequestPhaseBehaviorWithArgsRequestWithDefaults instantiates a new ApplicationRequestPhaseBehaviorWithArgsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) GetAttributes() ApplicationRuleEngineStringAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) GetAttributesOk() (*ApplicationRuleEngineStringAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRequestPhaseBehaviorWithArgsRequest) SetAttributes(v ApplicationRuleEngineStringAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


