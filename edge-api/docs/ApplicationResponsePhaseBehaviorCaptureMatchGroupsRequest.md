# ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;capture_match_groups&#x60; - capture_match_groups | 
**Attributes** | [**ApplicationRuleEngineCaptureMatchGroupsAttributes**](ApplicationRuleEngineCaptureMatchGroupsAttributes.md) |  | 

## Methods

### NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest

`func NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest(type_ string, attributes ApplicationRuleEngineCaptureMatchGroupsAttributes, ) *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest`

NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest instantiates a new ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequestWithDefaults

`func NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequestWithDefaults() *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest`

NewApplicationResponsePhaseBehaviorCaptureMatchGroupsRequestWithDefaults instantiates a new ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) GetAttributes() ApplicationRuleEngineCaptureMatchGroupsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) GetAttributesOk() (*ApplicationRuleEngineCaptureMatchGroupsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest) SetAttributes(v ApplicationRuleEngineCaptureMatchGroupsAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


