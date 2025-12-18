# BehaviorCaptureGroupsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviorsRequest | 
**Attributes** | [**AppRuleCaptureGroupsAttrsReq**](AppRuleCaptureGroupsAttrsReq.md) |  | 

## Methods

### NewBehaviorCaptureGroupsRequest2

`func NewBehaviorCaptureGroupsRequest2(type_ string, attributes AppRuleCaptureGroupsAttrsReq, ) *BehaviorCaptureGroupsRequest2`

NewBehaviorCaptureGroupsRequest2 instantiates a new BehaviorCaptureGroupsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCaptureGroupsRequest2WithDefaults

`func NewBehaviorCaptureGroupsRequest2WithDefaults() *BehaviorCaptureGroupsRequest2`

NewBehaviorCaptureGroupsRequest2WithDefaults instantiates a new BehaviorCaptureGroupsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCaptureGroupsRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCaptureGroupsRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCaptureGroupsRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCaptureGroupsRequest2) GetAttributes() AppRuleCaptureGroupsAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCaptureGroupsRequest2) GetAttributesOk() (*AppRuleCaptureGroupsAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCaptureGroupsRequest2) SetAttributes(v AppRuleCaptureGroupsAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


