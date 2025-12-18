# FwBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviorsRequest | 
**Attributes** | [**FwBehaviorFunctionAttrsReq**](FwBehaviorFunctionAttrsReq.md) |  | 

## Methods

### NewFwBehaviorsRequest

`func NewFwBehaviorsRequest(type_ string, attributes FwBehaviorFunctionAttrsReq, ) *FwBehaviorsRequest`

NewFwBehaviorsRequest instantiates a new FwBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorsRequestWithDefaults

`func NewFwBehaviorsRequestWithDefaults() *FwBehaviorsRequest`

NewFwBehaviorsRequestWithDefaults instantiates a new FwBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorsRequest) GetAttributes() FwBehaviorFunctionAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorsRequest) GetAttributesOk() (*FwBehaviorFunctionAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorsRequest) SetAttributes(v FwBehaviorFunctionAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


