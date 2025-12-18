# FwBehaviorCustomRespRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_custom_response&#x60; - set_custom_response | 
**Attributes** | [**FwBehaviorCustomRespAttrsReq**](FwBehaviorCustomRespAttrsReq.md) |  | 

## Methods

### NewFwBehaviorCustomRespRequest

`func NewFwBehaviorCustomRespRequest(type_ string, attributes FwBehaviorCustomRespAttrsReq, ) *FwBehaviorCustomRespRequest`

NewFwBehaviorCustomRespRequest instantiates a new FwBehaviorCustomRespRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorCustomRespRequestWithDefaults

`func NewFwBehaviorCustomRespRequestWithDefaults() *FwBehaviorCustomRespRequest`

NewFwBehaviorCustomRespRequestWithDefaults instantiates a new FwBehaviorCustomRespRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorCustomRespRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorCustomRespRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorCustomRespRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorCustomRespRequest) GetAttributes() FwBehaviorCustomRespAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorCustomRespRequest) GetAttributesOk() (*FwBehaviorCustomRespAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorCustomRespRequest) SetAttributes(v FwBehaviorCustomRespAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


