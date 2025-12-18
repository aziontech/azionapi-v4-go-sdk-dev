# AppRuleResponseString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 * &#x60;filter_response_cookie&#x60; - filter_response_cookie | 
**Attributes** | [**AppRuleResponseStringAttrs**](AppRuleResponseStringAttrs.md) |  | 

## Methods

### NewAppRuleResponseString

`func NewAppRuleResponseString(type_ string, attributes AppRuleResponseStringAttrs, ) *AppRuleResponseString`

NewAppRuleResponseString instantiates a new AppRuleResponseString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleResponseStringWithDefaults

`func NewAppRuleResponseStringWithDefaults() *AppRuleResponseString`

NewAppRuleResponseStringWithDefaults instantiates a new AppRuleResponseString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleResponseString) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleResponseString) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleResponseString) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleResponseString) GetAttributes() AppRuleResponseStringAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleResponseString) GetAttributesOk() (*AppRuleResponseStringAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleResponseString) SetAttributes(v AppRuleResponseStringAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


