# BehaviorFilterReqCookieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleFilterReqCookieAttrsReq**](AppRuleFilterReqCookieAttrsReq.md) |  | 

## Methods

### NewBehaviorFilterReqCookieRequest

`func NewBehaviorFilterReqCookieRequest(type_ string, attributes AppRuleFilterReqCookieAttrsReq, ) *BehaviorFilterReqCookieRequest`

NewBehaviorFilterReqCookieRequest instantiates a new BehaviorFilterReqCookieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFilterReqCookieRequestWithDefaults

`func NewBehaviorFilterReqCookieRequestWithDefaults() *BehaviorFilterReqCookieRequest`

NewBehaviorFilterReqCookieRequestWithDefaults instantiates a new BehaviorFilterReqCookieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFilterReqCookieRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFilterReqCookieRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFilterReqCookieRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFilterReqCookieRequest) GetAttributes() AppRuleFilterReqCookieAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFilterReqCookieRequest) GetAttributesOk() (*AppRuleFilterReqCookieAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFilterReqCookieRequest) SetAttributes(v AppRuleFilterReqCookieAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


