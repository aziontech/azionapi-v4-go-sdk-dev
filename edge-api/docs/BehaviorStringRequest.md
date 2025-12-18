# BehaviorStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ReqBehaviorsRequest | 
**Attributes** | [**AppRuleStringAttrsReq**](AppRuleStringAttrsReq.md) |  | 

## Methods

### NewBehaviorStringRequest

`func NewBehaviorStringRequest(type_ string, attributes AppRuleStringAttrsReq, ) *BehaviorStringRequest`

NewBehaviorStringRequest instantiates a new BehaviorStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorStringRequestWithDefaults

`func NewBehaviorStringRequestWithDefaults() *BehaviorStringRequest`

NewBehaviorStringRequestWithDefaults instantiates a new BehaviorStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorStringRequest) GetAttributes() AppRuleStringAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorStringRequest) GetAttributesOk() (*AppRuleStringAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorStringRequest) SetAttributes(v AppRuleStringAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


