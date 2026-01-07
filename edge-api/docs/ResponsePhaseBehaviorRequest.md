# ResponsePhaseBehaviorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;capture_match_groups&#x60; - capture_match_groups | 
**Attributes** | [**BehaviorCaptureMatchGroupsAttributesRequest**](BehaviorCaptureMatchGroupsAttributesRequest.md) |  | 

## Methods

### NewResponsePhaseBehaviorRequest

`func NewResponsePhaseBehaviorRequest(type_ string, attributes BehaviorCaptureMatchGroupsAttributesRequest, ) *ResponsePhaseBehaviorRequest`

NewResponsePhaseBehaviorRequest instantiates a new ResponsePhaseBehaviorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePhaseBehaviorRequestWithDefaults

`func NewResponsePhaseBehaviorRequestWithDefaults() *ResponsePhaseBehaviorRequest`

NewResponsePhaseBehaviorRequestWithDefaults instantiates a new ResponsePhaseBehaviorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponsePhaseBehaviorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePhaseBehaviorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePhaseBehaviorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ResponsePhaseBehaviorRequest) GetAttributes() BehaviorCaptureMatchGroupsAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResponsePhaseBehaviorRequest) GetAttributesOk() (*BehaviorCaptureMatchGroupsAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResponsePhaseBehaviorRequest) SetAttributes(v BehaviorCaptureMatchGroupsAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


