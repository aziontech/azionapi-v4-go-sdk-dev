# BehaviorConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleSetConnectorAttrsReq**](AppRuleSetConnectorAttrsReq.md) |  | 

## Methods

### NewBehaviorConnectorRequest

`func NewBehaviorConnectorRequest(type_ string, attributes AppRuleSetConnectorAttrsReq, ) *BehaviorConnectorRequest`

NewBehaviorConnectorRequest instantiates a new BehaviorConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorConnectorRequestWithDefaults

`func NewBehaviorConnectorRequestWithDefaults() *BehaviorConnectorRequest`

NewBehaviorConnectorRequestWithDefaults instantiates a new BehaviorConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorConnectorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorConnectorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorConnectorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorConnectorRequest) GetAttributes() AppRuleSetConnectorAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorConnectorRequest) GetAttributesOk() (*AppRuleSetConnectorAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorConnectorRequest) SetAttributes(v AppRuleSetConnectorAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


