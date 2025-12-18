# BehaviorAResponseHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for RespBehaviorsRequest | 
**Attributes** | [**AppRuleAResponseHeaderAttrsReq**](AppRuleAResponseHeaderAttrsReq.md) |  | 

## Methods

### NewBehaviorAResponseHeaderRequest

`func NewBehaviorAResponseHeaderRequest(type_ string, attributes AppRuleAResponseHeaderAttrsReq, ) *BehaviorAResponseHeaderRequest`

NewBehaviorAResponseHeaderRequest instantiates a new BehaviorAResponseHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorAResponseHeaderRequestWithDefaults

`func NewBehaviorAResponseHeaderRequestWithDefaults() *BehaviorAResponseHeaderRequest`

NewBehaviorAResponseHeaderRequestWithDefaults instantiates a new BehaviorAResponseHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorAResponseHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorAResponseHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorAResponseHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorAResponseHeaderRequest) GetAttributes() AppRuleAResponseHeaderAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorAResponseHeaderRequest) GetAttributesOk() (*AppRuleAResponseHeaderAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorAResponseHeaderRequest) SetAttributes(v AppRuleAResponseHeaderAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


