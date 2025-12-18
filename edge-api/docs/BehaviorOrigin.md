# BehaviorOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviors | 
**Attributes** | [**AppRuleSetOriginAttrs**](AppRuleSetOriginAttrs.md) |  | 

## Methods

### NewBehaviorOrigin

`func NewBehaviorOrigin(type_ string, attributes AppRuleSetOriginAttrs, ) *BehaviorOrigin`

NewBehaviorOrigin instantiates a new BehaviorOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorOriginWithDefaults

`func NewBehaviorOriginWithDefaults() *BehaviorOrigin`

NewBehaviorOriginWithDefaults instantiates a new BehaviorOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorOrigin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorOrigin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorOrigin) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorOrigin) GetAttributes() AppRuleSetOriginAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorOrigin) GetAttributesOk() (*AppRuleSetOriginAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorOrigin) SetAttributes(v AppRuleSetOriginAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


