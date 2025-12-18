# BehaviorCaptureGroups2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviors | 
**Attributes** | [**AppRuleCaptureGroupsAttrs**](AppRuleCaptureGroupsAttrs.md) |  | 

## Methods

### NewBehaviorCaptureGroups2

`func NewBehaviorCaptureGroups2(type_ string, attributes AppRuleCaptureGroupsAttrs, ) *BehaviorCaptureGroups2`

NewBehaviorCaptureGroups2 instantiates a new BehaviorCaptureGroups2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCaptureGroups2WithDefaults

`func NewBehaviorCaptureGroups2WithDefaults() *BehaviorCaptureGroups2`

NewBehaviorCaptureGroups2WithDefaults instantiates a new BehaviorCaptureGroups2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCaptureGroups2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCaptureGroups2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCaptureGroups2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCaptureGroups2) GetAttributes() AppRuleCaptureGroupsAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCaptureGroups2) GetAttributesOk() (*AppRuleCaptureGroupsAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCaptureGroups2) SetAttributes(v AppRuleCaptureGroupsAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


