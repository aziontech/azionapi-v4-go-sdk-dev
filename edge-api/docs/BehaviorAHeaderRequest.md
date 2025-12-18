# BehaviorAHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleAHeaderAttrsReq**](AppRuleAHeaderAttrsReq.md) |  | 

## Methods

### NewBehaviorAHeaderRequest

`func NewBehaviorAHeaderRequest(type_ string, attributes AppRuleAHeaderAttrsReq, ) *BehaviorAHeaderRequest`

NewBehaviorAHeaderRequest instantiates a new BehaviorAHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorAHeaderRequestWithDefaults

`func NewBehaviorAHeaderRequestWithDefaults() *BehaviorAHeaderRequest`

NewBehaviorAHeaderRequestWithDefaults instantiates a new BehaviorAHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorAHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorAHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorAHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorAHeaderRequest) GetAttributes() AppRuleAHeaderAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorAHeaderRequest) GetAttributesOk() (*AppRuleAHeaderAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorAHeaderRequest) SetAttributes(v AppRuleAHeaderAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


