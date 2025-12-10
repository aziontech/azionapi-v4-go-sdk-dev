# FirewallBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Type689Enum**](Type689Enum.md) |  | 
**Attributes** | [**FirewallBehaviorRunFunctionAttributesRequest**](FirewallBehaviorRunFunctionAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorsRequest

`func NewFirewallBehaviorsRequest(type_ Type689Enum, attributes FirewallBehaviorRunFunctionAttributesRequest, ) *FirewallBehaviorsRequest`

NewFirewallBehaviorsRequest instantiates a new FirewallBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorsRequestWithDefaults

`func NewFirewallBehaviorsRequestWithDefaults() *FirewallBehaviorsRequest`

NewFirewallBehaviorsRequestWithDefaults instantiates a new FirewallBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorsRequest) GetType() Type689Enum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorsRequest) GetTypeOk() (*Type689Enum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorsRequest) SetType(v Type689Enum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorsRequest) GetAttributes() FirewallBehaviorRunFunctionAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorsRequest) GetAttributesOk() (*FirewallBehaviorRunFunctionAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorsRequest) SetAttributes(v FirewallBehaviorRunFunctionAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


