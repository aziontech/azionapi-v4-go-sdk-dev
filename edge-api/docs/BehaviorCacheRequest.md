# BehaviorCacheRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleSetCachePolicyAttrsReq**](AppRuleSetCachePolicyAttrsReq.md) |  | 

## Methods

### NewBehaviorCacheRequest

`func NewBehaviorCacheRequest(type_ string, attributes AppRuleSetCachePolicyAttrsReq, ) *BehaviorCacheRequest`

NewBehaviorCacheRequest instantiates a new BehaviorCacheRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCacheRequestWithDefaults

`func NewBehaviorCacheRequestWithDefaults() *BehaviorCacheRequest`

NewBehaviorCacheRequestWithDefaults instantiates a new BehaviorCacheRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCacheRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCacheRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCacheRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCacheRequest) GetAttributes() AppRuleSetCachePolicyAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCacheRequest) GetAttributesOk() (*AppRuleSetCachePolicyAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCacheRequest) SetAttributes(v AppRuleSetCachePolicyAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


