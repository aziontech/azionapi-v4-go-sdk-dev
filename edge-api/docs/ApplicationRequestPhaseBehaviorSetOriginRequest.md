# ApplicationRequestPhaseBehaviorSetOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_origin&#x60; - set_origin | 
**Attributes** | [**ApplicationRuleEngineSetOriginAttributesRequest**](ApplicationRuleEngineSetOriginAttributesRequest.md) |  | 

## Methods

### NewApplicationRequestPhaseBehaviorSetOriginRequest

`func NewApplicationRequestPhaseBehaviorSetOriginRequest(type_ string, attributes ApplicationRuleEngineSetOriginAttributesRequest, ) *ApplicationRequestPhaseBehaviorSetOriginRequest`

NewApplicationRequestPhaseBehaviorSetOriginRequest instantiates a new ApplicationRequestPhaseBehaviorSetOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseBehaviorSetOriginRequestWithDefaults

`func NewApplicationRequestPhaseBehaviorSetOriginRequestWithDefaults() *ApplicationRequestPhaseBehaviorSetOriginRequest`

NewApplicationRequestPhaseBehaviorSetOriginRequestWithDefaults instantiates a new ApplicationRequestPhaseBehaviorSetOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) GetAttributes() ApplicationRuleEngineSetOriginAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) GetAttributesOk() (*ApplicationRuleEngineSetOriginAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRequestPhaseBehaviorSetOriginRequest) SetAttributes(v ApplicationRuleEngineSetOriginAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


