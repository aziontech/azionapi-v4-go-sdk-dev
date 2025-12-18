# FwBehaviorCustomRespRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviorsRequest | 
**Attributes** | [**FwBehaviorCustomRespAttrsReq**](FwBehaviorCustomRespAttrsReq.md) |  | 

## Methods

### NewFwBehaviorCustomRespRequest2

`func NewFwBehaviorCustomRespRequest2(type_ string, attributes FwBehaviorCustomRespAttrsReq, ) *FwBehaviorCustomRespRequest2`

NewFwBehaviorCustomRespRequest2 instantiates a new FwBehaviorCustomRespRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorCustomRespRequest2WithDefaults

`func NewFwBehaviorCustomRespRequest2WithDefaults() *FwBehaviorCustomRespRequest2`

NewFwBehaviorCustomRespRequest2WithDefaults instantiates a new FwBehaviorCustomRespRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorCustomRespRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorCustomRespRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorCustomRespRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorCustomRespRequest2) GetAttributes() FwBehaviorCustomRespAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorCustomRespRequest2) GetAttributesOk() (*FwBehaviorCustomRespAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorCustomRespRequest2) SetAttributes(v FwBehaviorCustomRespAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


