# ApplicationRuleEngineRequestPhaseBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_origin&#x60; - set_origin | 
**Attributes** | [**ApplicationRuleEngineSetOriginAttributesRequest**](ApplicationRuleEngineSetOriginAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineRequestPhaseBehaviorsRequest

`func NewApplicationRuleEngineRequestPhaseBehaviorsRequest(type_ string, attributes ApplicationRuleEngineSetOriginAttributesRequest, ) *ApplicationRuleEngineRequestPhaseBehaviorsRequest`

NewApplicationRuleEngineRequestPhaseBehaviorsRequest instantiates a new ApplicationRuleEngineRequestPhaseBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults

`func NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults() *ApplicationRuleEngineRequestPhaseBehaviorsRequest`

NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults instantiates a new ApplicationRuleEngineRequestPhaseBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetAttributes() ApplicationRuleEngineSetOriginAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetAttributesOk() (*ApplicationRuleEngineSetOriginAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) SetAttributes(v ApplicationRuleEngineSetOriginAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


