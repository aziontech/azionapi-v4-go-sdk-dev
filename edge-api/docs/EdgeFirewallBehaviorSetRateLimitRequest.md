# EdgeFirewallBehaviorSetRateLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_rate_limit&#x60; - set_rate_limit | 
**Attributes** | [**EdgeFirewallBehaviorSetRateLimitAttributesRequest**](EdgeFirewallBehaviorSetRateLimitAttributesRequest.md) |  | 

## Methods

### NewEdgeFirewallBehaviorSetRateLimitRequest

`func NewEdgeFirewallBehaviorSetRateLimitRequest(type_ string, attributes EdgeFirewallBehaviorSetRateLimitAttributesRequest, ) *EdgeFirewallBehaviorSetRateLimitRequest`

NewEdgeFirewallBehaviorSetRateLimitRequest instantiates a new EdgeFirewallBehaviorSetRateLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorSetRateLimitRequestWithDefaults

`func NewEdgeFirewallBehaviorSetRateLimitRequestWithDefaults() *EdgeFirewallBehaviorSetRateLimitRequest`

NewEdgeFirewallBehaviorSetRateLimitRequestWithDefaults instantiates a new EdgeFirewallBehaviorSetRateLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) GetAttributes() EdgeFirewallBehaviorSetRateLimitAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) GetAttributesOk() (*EdgeFirewallBehaviorSetRateLimitAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeFirewallBehaviorSetRateLimitRequest) SetAttributes(v EdgeFirewallBehaviorSetRateLimitAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


