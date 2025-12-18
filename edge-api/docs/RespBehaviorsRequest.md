# RespBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviorsRequest | 
**Attributes** | [**AppRuleCaptureGroupsAttrsReq**](AppRuleCaptureGroupsAttrsReq.md) |  | 

## Methods

### NewRespBehaviorsRequest

`func NewRespBehaviorsRequest(type_ string, attributes AppRuleCaptureGroupsAttrsReq, ) *RespBehaviorsRequest`

NewRespBehaviorsRequest instantiates a new RespBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespBehaviorsRequestWithDefaults

`func NewRespBehaviorsRequestWithDefaults() *RespBehaviorsRequest`

NewRespBehaviorsRequestWithDefaults instantiates a new RespBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RespBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RespBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RespBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *RespBehaviorsRequest) GetAttributes() AppRuleCaptureGroupsAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RespBehaviorsRequest) GetAttributesOk() (*AppRuleCaptureGroupsAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RespBehaviorsRequest) SetAttributes(v AppRuleCaptureGroupsAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


