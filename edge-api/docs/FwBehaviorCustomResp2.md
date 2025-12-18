# FwBehaviorCustomResp2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviors | 
**Attributes** | [**FwBehaviorCustomRespAttrs**](FwBehaviorCustomRespAttrs.md) |  | 

## Methods

### NewFwBehaviorCustomResp2

`func NewFwBehaviorCustomResp2(type_ string, attributes FwBehaviorCustomRespAttrs, ) *FwBehaviorCustomResp2`

NewFwBehaviorCustomResp2 instantiates a new FwBehaviorCustomResp2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorCustomResp2WithDefaults

`func NewFwBehaviorCustomResp2WithDefaults() *FwBehaviorCustomResp2`

NewFwBehaviorCustomResp2WithDefaults instantiates a new FwBehaviorCustomResp2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorCustomResp2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorCustomResp2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorCustomResp2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorCustomResp2) GetAttributes() FwBehaviorCustomRespAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorCustomResp2) GetAttributesOk() (*FwBehaviorCustomRespAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorCustomResp2) SetAttributes(v FwBehaviorCustomRespAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


