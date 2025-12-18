# AppRuleSetCookieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_cookie&#x60; - set_cookie | 
**Attributes** | [**AppRuleSetCookieAttrsReq**](AppRuleSetCookieAttrsReq.md) |  | 

## Methods

### NewAppRuleSetCookieRequest

`func NewAppRuleSetCookieRequest(type_ string, attributes AppRuleSetCookieAttrsReq, ) *AppRuleSetCookieRequest`

NewAppRuleSetCookieRequest instantiates a new AppRuleSetCookieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleSetCookieRequestWithDefaults

`func NewAppRuleSetCookieRequestWithDefaults() *AppRuleSetCookieRequest`

NewAppRuleSetCookieRequestWithDefaults instantiates a new AppRuleSetCookieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleSetCookieRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleSetCookieRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleSetCookieRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleSetCookieRequest) GetAttributes() AppRuleSetCookieAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleSetCookieRequest) GetAttributesOk() (*AppRuleSetCookieAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleSetCookieRequest) SetAttributes(v AppRuleSetCookieAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


