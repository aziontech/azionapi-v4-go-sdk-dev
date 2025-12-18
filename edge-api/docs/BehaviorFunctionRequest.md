# BehaviorFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleRunFunctionAttrsReq**](AppRuleRunFunctionAttrsReq.md) |  | 

## Methods

### NewBehaviorFunctionRequest

`func NewBehaviorFunctionRequest(type_ string, attributes AppRuleRunFunctionAttrsReq, ) *BehaviorFunctionRequest`

NewBehaviorFunctionRequest instantiates a new BehaviorFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFunctionRequestWithDefaults

`func NewBehaviorFunctionRequestWithDefaults() *BehaviorFunctionRequest`

NewBehaviorFunctionRequestWithDefaults instantiates a new BehaviorFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFunctionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFunctionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFunctionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFunctionRequest) GetAttributes() AppRuleRunFunctionAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFunctionRequest) GetAttributesOk() (*AppRuleRunFunctionAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFunctionRequest) SetAttributes(v AppRuleRunFunctionAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


