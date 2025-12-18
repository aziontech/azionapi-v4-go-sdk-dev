# BehaviorCaptureGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleCaptureGroupsAttrsReq**](AppRuleCaptureGroupsAttrsReq.md) |  | 

## Methods

### NewBehaviorCaptureGroupsRequest

`func NewBehaviorCaptureGroupsRequest(type_ string, attributes AppRuleCaptureGroupsAttrsReq, ) *BehaviorCaptureGroupsRequest`

NewBehaviorCaptureGroupsRequest instantiates a new BehaviorCaptureGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorCaptureGroupsRequestWithDefaults

`func NewBehaviorCaptureGroupsRequestWithDefaults() *BehaviorCaptureGroupsRequest`

NewBehaviorCaptureGroupsRequestWithDefaults instantiates a new BehaviorCaptureGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorCaptureGroupsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorCaptureGroupsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorCaptureGroupsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorCaptureGroupsRequest) GetAttributes() AppRuleCaptureGroupsAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorCaptureGroupsRequest) GetAttributesOk() (*AppRuleCaptureGroupsAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorCaptureGroupsRequest) SetAttributes(v AppRuleCaptureGroupsAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


