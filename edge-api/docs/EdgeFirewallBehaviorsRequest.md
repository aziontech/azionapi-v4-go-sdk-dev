# EdgeFirewallBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**EdgeFirewallBehaviorRunFunctionAttributesRequest**](EdgeFirewallBehaviorRunFunctionAttributesRequest.md) |  | 

## Methods

### NewEdgeFirewallBehaviorsRequest

`func NewEdgeFirewallBehaviorsRequest(type_ string, attributes EdgeFirewallBehaviorRunFunctionAttributesRequest, ) *EdgeFirewallBehaviorsRequest`

NewEdgeFirewallBehaviorsRequest instantiates a new EdgeFirewallBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorsRequestWithDefaults

`func NewEdgeFirewallBehaviorsRequestWithDefaults() *EdgeFirewallBehaviorsRequest`

NewEdgeFirewallBehaviorsRequestWithDefaults instantiates a new EdgeFirewallBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeFirewallBehaviorsRequest) GetAttributes() EdgeFirewallBehaviorRunFunctionAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeFirewallBehaviorsRequest) GetAttributesOk() (*EdgeFirewallBehaviorRunFunctionAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeFirewallBehaviorsRequest) SetAttributes(v EdgeFirewallBehaviorRunFunctionAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


