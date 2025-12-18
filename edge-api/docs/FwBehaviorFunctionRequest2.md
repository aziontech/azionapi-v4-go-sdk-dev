# FwBehaviorFunctionRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviorsRequest | 
**Attributes** | [**FwBehaviorFunctionAttrsReq**](FwBehaviorFunctionAttrsReq.md) |  | 

## Methods

### NewFwBehaviorFunctionRequest2

`func NewFwBehaviorFunctionRequest2(type_ string, attributes FwBehaviorFunctionAttrsReq, ) *FwBehaviorFunctionRequest2`

NewFwBehaviorFunctionRequest2 instantiates a new FwBehaviorFunctionRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorFunctionRequest2WithDefaults

`func NewFwBehaviorFunctionRequest2WithDefaults() *FwBehaviorFunctionRequest2`

NewFwBehaviorFunctionRequest2WithDefaults instantiates a new FwBehaviorFunctionRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorFunctionRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorFunctionRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorFunctionRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorFunctionRequest2) GetAttributes() FwBehaviorFunctionAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorFunctionRequest2) GetAttributesOk() (*FwBehaviorFunctionAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorFunctionRequest2) SetAttributes(v FwBehaviorFunctionAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


