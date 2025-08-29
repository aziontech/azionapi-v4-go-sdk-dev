# ApplicationRequestPhaseBehaviorSetCachePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_cache_policy&#x60; - set_cache_policy | 
**Attributes** | [**ApplicationRuleEngineSetCachePolicyAttributesRequest**](ApplicationRuleEngineSetCachePolicyAttributesRequest.md) |  | 

## Methods

### NewApplicationRequestPhaseBehaviorSetCachePolicyRequest

`func NewApplicationRequestPhaseBehaviorSetCachePolicyRequest(type_ string, attributes ApplicationRuleEngineSetCachePolicyAttributesRequest, ) *ApplicationRequestPhaseBehaviorSetCachePolicyRequest`

NewApplicationRequestPhaseBehaviorSetCachePolicyRequest instantiates a new ApplicationRequestPhaseBehaviorSetCachePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseBehaviorSetCachePolicyRequestWithDefaults

`func NewApplicationRequestPhaseBehaviorSetCachePolicyRequestWithDefaults() *ApplicationRequestPhaseBehaviorSetCachePolicyRequest`

NewApplicationRequestPhaseBehaviorSetCachePolicyRequestWithDefaults instantiates a new ApplicationRequestPhaseBehaviorSetCachePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) GetAttributes() ApplicationRuleEngineSetCachePolicyAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) GetAttributesOk() (*ApplicationRuleEngineSetCachePolicyAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRequestPhaseBehaviorSetCachePolicyRequest) SetAttributes(v ApplicationRuleEngineSetCachePolicyAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


