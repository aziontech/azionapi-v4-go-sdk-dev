# EdgeApplicationRuleEngineResponseStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 * &#x60;filter_response_cookie&#x60; - filter_response_cookie | 
**Attributes** | [**EdgeApplicationRuleEngineResponseStringAttributesRequest**](EdgeApplicationRuleEngineResponseStringAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineResponseStringRequest

`func NewEdgeApplicationRuleEngineResponseStringRequest(type_ string, attributes EdgeApplicationRuleEngineResponseStringAttributesRequest, ) *EdgeApplicationRuleEngineResponseStringRequest`

NewEdgeApplicationRuleEngineResponseStringRequest instantiates a new EdgeApplicationRuleEngineResponseStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineResponseStringRequestWithDefaults

`func NewEdgeApplicationRuleEngineResponseStringRequestWithDefaults() *EdgeApplicationRuleEngineResponseStringRequest`

NewEdgeApplicationRuleEngineResponseStringRequestWithDefaults instantiates a new EdgeApplicationRuleEngineResponseStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineResponseStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineResponseStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineResponseStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineResponseStringRequest) GetAttributes() EdgeApplicationRuleEngineResponseStringAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineResponseStringRequest) GetAttributesOk() (*EdgeApplicationRuleEngineResponseStringAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineResponseStringRequest) SetAttributes(v EdgeApplicationRuleEngineResponseStringAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


