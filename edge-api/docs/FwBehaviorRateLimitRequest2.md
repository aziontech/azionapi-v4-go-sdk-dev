# FwBehaviorRateLimitRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviorsRequest | 
**Attributes** | [**FwBehaviorRateLimitAttrsReq**](FwBehaviorRateLimitAttrsReq.md) |  | 

## Methods

### NewFwBehaviorRateLimitRequest2

`func NewFwBehaviorRateLimitRequest2(type_ string, attributes FwBehaviorRateLimitAttrsReq, ) *FwBehaviorRateLimitRequest2`

NewFwBehaviorRateLimitRequest2 instantiates a new FwBehaviorRateLimitRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorRateLimitRequest2WithDefaults

`func NewFwBehaviorRateLimitRequest2WithDefaults() *FwBehaviorRateLimitRequest2`

NewFwBehaviorRateLimitRequest2WithDefaults instantiates a new FwBehaviorRateLimitRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorRateLimitRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorRateLimitRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorRateLimitRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorRateLimitRequest2) GetAttributes() FwBehaviorRateLimitAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorRateLimitRequest2) GetAttributesOk() (*FwBehaviorRateLimitAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorRateLimitRequest2) SetAttributes(v FwBehaviorRateLimitAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


