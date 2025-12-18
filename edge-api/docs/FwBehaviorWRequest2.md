# FwBehaviorWRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviorsRequest | 
**Attributes** | [**FwBehaviorWAttrsReq**](FwBehaviorWAttrsReq.md) |  | 

## Methods

### NewFwBehaviorWRequest2

`func NewFwBehaviorWRequest2(type_ string, attributes FwBehaviorWAttrsReq, ) *FwBehaviorWRequest2`

NewFwBehaviorWRequest2 instantiates a new FwBehaviorWRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorWRequest2WithDefaults

`func NewFwBehaviorWRequest2WithDefaults() *FwBehaviorWRequest2`

NewFwBehaviorWRequest2WithDefaults instantiates a new FwBehaviorWRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorWRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorWRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorWRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorWRequest2) GetAttributes() FwBehaviorWAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorWRequest2) GetAttributesOk() (*FwBehaviorWAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorWRequest2) SetAttributes(v FwBehaviorWAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


