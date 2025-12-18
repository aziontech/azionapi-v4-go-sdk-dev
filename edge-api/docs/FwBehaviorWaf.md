# FwBehaviorWaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_waf&#x60; - set_waf | 
**Attributes** | [**FwBehaviorWAttrs**](FwBehaviorWAttrs.md) |  | 

## Methods

### NewFwBehaviorWaf

`func NewFwBehaviorWaf(type_ string, attributes FwBehaviorWAttrs, ) *FwBehaviorWaf`

NewFwBehaviorWaf instantiates a new FwBehaviorWaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorWafWithDefaults

`func NewFwBehaviorWafWithDefaults() *FwBehaviorWaf`

NewFwBehaviorWafWithDefaults instantiates a new FwBehaviorWaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorWaf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorWaf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorWaf) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorWaf) GetAttributes() FwBehaviorWAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorWaf) GetAttributesOk() (*FwBehaviorWAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorWaf) SetAttributes(v FwBehaviorWAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


