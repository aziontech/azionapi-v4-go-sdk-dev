# BehaviorRewriteRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleRewriteRequestAttrsReq**](AppRuleRewriteRequestAttrsReq.md) |  | 

## Methods

### NewBehaviorRewriteRequest2

`func NewBehaviorRewriteRequest2(type_ string, attributes AppRuleRewriteRequestAttrsReq, ) *BehaviorRewriteRequest2`

NewBehaviorRewriteRequest2 instantiates a new BehaviorRewriteRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRewriteRequest2WithDefaults

`func NewBehaviorRewriteRequest2WithDefaults() *BehaviorRewriteRequest2`

NewBehaviorRewriteRequest2WithDefaults instantiates a new BehaviorRewriteRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorRewriteRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRewriteRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRewriteRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorRewriteRequest2) GetAttributes() AppRuleRewriteRequestAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorRewriteRequest2) GetAttributesOk() (*AppRuleRewriteRequestAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorRewriteRequest2) SetAttributes(v AppRuleRewriteRequestAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


