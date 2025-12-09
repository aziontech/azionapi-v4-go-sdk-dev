# ApplicationRequestPhaseBehaviorRewriteRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;rewrite_request&#x60; - rewrite_request | 
**Attributes** | [**ApplicationRuleEngineRewriteRequestAttributesRequest**](ApplicationRuleEngineRewriteRequestAttributesRequest.md) |  | 

## Methods

### NewApplicationRequestPhaseBehaviorRewriteRequestRequest

`func NewApplicationRequestPhaseBehaviorRewriteRequestRequest(type_ string, attributes ApplicationRuleEngineRewriteRequestAttributesRequest, ) *ApplicationRequestPhaseBehaviorRewriteRequestRequest`

NewApplicationRequestPhaseBehaviorRewriteRequestRequest instantiates a new ApplicationRequestPhaseBehaviorRewriteRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseBehaviorRewriteRequestRequestWithDefaults

`func NewApplicationRequestPhaseBehaviorRewriteRequestRequestWithDefaults() *ApplicationRequestPhaseBehaviorRewriteRequestRequest`

NewApplicationRequestPhaseBehaviorRewriteRequestRequestWithDefaults instantiates a new ApplicationRequestPhaseBehaviorRewriteRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) GetAttributes() ApplicationRuleEngineRewriteRequestAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) GetAttributesOk() (*ApplicationRuleEngineRewriteRequestAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRequestPhaseBehaviorRewriteRequestRequest) SetAttributes(v ApplicationRuleEngineRewriteRequestAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


