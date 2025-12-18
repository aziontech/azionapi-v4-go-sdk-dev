# BehaviorFunctionResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviorsRequest | 
**Attributes** | [**AppRuleRunFunctionResponseAttrsReq**](AppRuleRunFunctionResponseAttrsReq.md) |  | 

## Methods

### NewBehaviorFunctionResponseRequest

`func NewBehaviorFunctionResponseRequest(type_ string, attributes AppRuleRunFunctionResponseAttrsReq, ) *BehaviorFunctionResponseRequest`

NewBehaviorFunctionResponseRequest instantiates a new BehaviorFunctionResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFunctionResponseRequestWithDefaults

`func NewBehaviorFunctionResponseRequestWithDefaults() *BehaviorFunctionResponseRequest`

NewBehaviorFunctionResponseRequestWithDefaults instantiates a new BehaviorFunctionResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFunctionResponseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFunctionResponseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFunctionResponseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFunctionResponseRequest) GetAttributes() AppRuleRunFunctionResponseAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFunctionResponseRequest) GetAttributesOk() (*AppRuleRunFunctionResponseAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFunctionResponseRequest) SetAttributes(v AppRuleRunFunctionResponseAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


