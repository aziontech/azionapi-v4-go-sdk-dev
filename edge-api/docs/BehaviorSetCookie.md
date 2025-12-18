# BehaviorSetCookie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviors | 
**Attributes** | [**AppRuleSetCookieAttrs**](AppRuleSetCookieAttrs.md) |  | 

## Methods

### NewBehaviorSetCookie

`func NewBehaviorSetCookie(type_ string, attributes AppRuleSetCookieAttrs, ) *BehaviorSetCookie`

NewBehaviorSetCookie instantiates a new BehaviorSetCookie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorSetCookieWithDefaults

`func NewBehaviorSetCookieWithDefaults() *BehaviorSetCookie`

NewBehaviorSetCookieWithDefaults instantiates a new BehaviorSetCookie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorSetCookie) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorSetCookie) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorSetCookie) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorSetCookie) GetAttributes() AppRuleSetCookieAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorSetCookie) GetAttributesOk() (*AppRuleSetCookieAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorSetCookie) SetAttributes(v AppRuleSetCookieAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


