# BehaviorCaptureGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviors | 
**Attributes** | [**AppRuleCaptureGroupsAttrs**](AppRuleCaptureGroupsAttrs.md) |  | 

## Methods

### NewBehaviorCaptureGroups

`func NewBehaviorCaptureGroups(type_ string, attributes AppRuleCaptureGroupsAttrs, ) *BehaviorCaptureGroups`

NewBehaviorCaptureGroups instantiates a new BehaviorCaptureGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCaptureGroupsWithDefaults

`func NewBehaviorCaptureGroupsWithDefaults() *BehaviorCaptureGroups`

NewBehaviorCaptureGroupsWithDefaults instantiates a new BehaviorCaptureGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCaptureGroups) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCaptureGroups) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCaptureGroups) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCaptureGroups) GetAttributes() AppRuleCaptureGroupsAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCaptureGroups) GetAttributesOk() (*AppRuleCaptureGroupsAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCaptureGroups) SetAttributes(v AppRuleCaptureGroupsAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


