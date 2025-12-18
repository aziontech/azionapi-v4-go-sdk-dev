# BehaviorFilterRespHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviors | 
**Attributes** | [**AppRuleFilterRespHeaderAttrs**](AppRuleFilterRespHeaderAttrs.md) |  | 

## Methods

### NewBehaviorFilterRespHeader

`func NewBehaviorFilterRespHeader(type_ string, attributes AppRuleFilterRespHeaderAttrs, ) *BehaviorFilterRespHeader`

NewBehaviorFilterRespHeader instantiates a new BehaviorFilterRespHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFilterRespHeaderWithDefaults

`func NewBehaviorFilterRespHeaderWithDefaults() *BehaviorFilterRespHeader`

NewBehaviorFilterRespHeaderWithDefaults instantiates a new BehaviorFilterRespHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFilterRespHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFilterRespHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFilterRespHeader) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFilterRespHeader) GetAttributes() AppRuleFilterRespHeaderAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFilterRespHeader) GetAttributesOk() (*AppRuleFilterRespHeaderAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFilterRespHeader) SetAttributes(v AppRuleFilterRespHeaderAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


