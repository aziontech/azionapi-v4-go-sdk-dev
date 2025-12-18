# BehaviorARequestCookieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleARequestCookieAttrsReq**](AppRuleARequestCookieAttrsReq.md) |  | 

## Methods

### NewBehaviorARequestCookieRequest

`func NewBehaviorARequestCookieRequest(type_ string, attributes AppRuleARequestCookieAttrsReq, ) *BehaviorARequestCookieRequest`

NewBehaviorARequestCookieRequest instantiates a new BehaviorARequestCookieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorARequestCookieRequestWithDefaults

`func NewBehaviorARequestCookieRequestWithDefaults() *BehaviorARequestCookieRequest`

NewBehaviorARequestCookieRequestWithDefaults instantiates a new BehaviorARequestCookieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorARequestCookieRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorARequestCookieRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorARequestCookieRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorARequestCookieRequest) GetAttributes() AppRuleARequestCookieAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorARequestCookieRequest) GetAttributesOk() (*AppRuleARequestCookieAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorARequestCookieRequest) SetAttributes(v AppRuleARequestCookieAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


