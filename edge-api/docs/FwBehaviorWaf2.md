# FwBehaviorWaf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviors | 
**Attributes** | [**FwBehaviorWAttrs**](FwBehaviorWAttrs.md) |  | 

## Methods

### NewFwBehaviorWaf2

`func NewFwBehaviorWaf2(type_ string, attributes FwBehaviorWAttrs, ) *FwBehaviorWaf2`

NewFwBehaviorWaf2 instantiates a new FwBehaviorWaf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorWaf2WithDefaults

`func NewFwBehaviorWaf2WithDefaults() *FwBehaviorWaf2`

NewFwBehaviorWaf2WithDefaults instantiates a new FwBehaviorWaf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorWaf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorWaf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorWaf2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorWaf2) GetAttributes() FwBehaviorWAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorWaf2) GetAttributesOk() (*FwBehaviorWAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorWaf2) SetAttributes(v FwBehaviorWAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


