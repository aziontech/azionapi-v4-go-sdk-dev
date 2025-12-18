# BehaviorCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviors | 
**Attributes** | [**AppRuleSetCachePolicyAttrs**](AppRuleSetCachePolicyAttrs.md) |  | 

## Methods

### NewBehaviorCache

`func NewBehaviorCache(type_ string, attributes AppRuleSetCachePolicyAttrs, ) *BehaviorCache`

NewBehaviorCache instantiates a new BehaviorCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCacheWithDefaults

`func NewBehaviorCacheWithDefaults() *BehaviorCache`

NewBehaviorCacheWithDefaults instantiates a new BehaviorCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCache) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCache) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCache) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCache) GetAttributes() AppRuleSetCachePolicyAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCache) GetAttributesOk() (*AppRuleSetCachePolicyAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCache) SetAttributes(v AppRuleSetCachePolicyAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


