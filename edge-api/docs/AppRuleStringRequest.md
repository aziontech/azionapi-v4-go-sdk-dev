# AppRuleStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**AppRuleStringAttrsReq**](AppRuleStringAttrsReq.md) |  | 

## Methods

### NewAppRuleStringRequest

`func NewAppRuleStringRequest(type_ string, attributes AppRuleStringAttrsReq, ) *AppRuleStringRequest`

NewAppRuleStringRequest instantiates a new AppRuleStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleStringRequestWithDefaults

`func NewAppRuleStringRequestWithDefaults() *AppRuleStringRequest`

NewAppRuleStringRequestWithDefaults instantiates a new AppRuleStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleStringRequest) GetAttributes() AppRuleStringAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleStringRequest) GetAttributesOk() (*AppRuleStringAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleStringRequest) SetAttributes(v AppRuleStringAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


