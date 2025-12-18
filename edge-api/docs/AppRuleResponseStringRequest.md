# AppRuleResponseStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 * &#x60;filter_response_cookie&#x60; - filter_response_cookie | 
**Attributes** | [**AppRuleResponseStringAttrsReq**](AppRuleResponseStringAttrsReq.md) |  | 

## Methods

### NewAppRuleResponseStringRequest

`func NewAppRuleResponseStringRequest(type_ string, attributes AppRuleResponseStringAttrsReq, ) *AppRuleResponseStringRequest`

NewAppRuleResponseStringRequest instantiates a new AppRuleResponseStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleResponseStringRequestWithDefaults

`func NewAppRuleResponseStringRequestWithDefaults() *AppRuleResponseStringRequest`

NewAppRuleResponseStringRequestWithDefaults instantiates a new AppRuleResponseStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleResponseStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleResponseStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleResponseStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleResponseStringRequest) GetAttributes() AppRuleResponseStringAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleResponseStringRequest) GetAttributesOk() (*AppRuleResponseStringAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleResponseStringRequest) SetAttributes(v AppRuleResponseStringAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


