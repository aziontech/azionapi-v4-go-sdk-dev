# BehaviorFilterRespHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviorsRequest | 
**Attributes** | [**AppRuleFilterRespHeaderAttrsReq**](AppRuleFilterRespHeaderAttrsReq.md) |  | 

## Methods

### NewBehaviorFilterRespHeaderRequest

`func NewBehaviorFilterRespHeaderRequest(type_ string, attributes AppRuleFilterRespHeaderAttrsReq, ) *BehaviorFilterRespHeaderRequest`

NewBehaviorFilterRespHeaderRequest instantiates a new BehaviorFilterRespHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFilterRespHeaderRequestWithDefaults

`func NewBehaviorFilterRespHeaderRequestWithDefaults() *BehaviorFilterRespHeaderRequest`

NewBehaviorFilterRespHeaderRequestWithDefaults instantiates a new BehaviorFilterRespHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFilterRespHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFilterRespHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFilterRespHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFilterRespHeaderRequest) GetAttributes() AppRuleFilterRespHeaderAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFilterRespHeaderRequest) GetAttributesOk() (*AppRuleFilterRespHeaderAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFilterRespHeaderRequest) SetAttributes(v AppRuleFilterRespHeaderAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


