# ApplicationRuleEngineSetOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ApplicationRuleEngineSetOriginTypeEnum**](ApplicationRuleEngineSetOriginTypeEnum.md) |  | 
**Attributes** | [**ApplicationRuleEngineSetOriginAttributesRequest**](ApplicationRuleEngineSetOriginAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineSetOriginRequest

`func NewApplicationRuleEngineSetOriginRequest(type_ ApplicationRuleEngineSetOriginTypeEnum, attributes ApplicationRuleEngineSetOriginAttributesRequest, ) *ApplicationRuleEngineSetOriginRequest`

NewApplicationRuleEngineSetOriginRequest instantiates a new ApplicationRuleEngineSetOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineSetOriginRequestWithDefaults

`func NewApplicationRuleEngineSetOriginRequestWithDefaults() *ApplicationRuleEngineSetOriginRequest`

NewApplicationRuleEngineSetOriginRequestWithDefaults instantiates a new ApplicationRuleEngineSetOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineSetOriginRequest) GetType() ApplicationRuleEngineSetOriginTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineSetOriginRequest) GetTypeOk() (*ApplicationRuleEngineSetOriginTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineSetOriginRequest) SetType(v ApplicationRuleEngineSetOriginTypeEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineSetOriginRequest) GetAttributes() ApplicationRuleEngineSetOriginAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineSetOriginRequest) GetAttributesOk() (*ApplicationRuleEngineSetOriginAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineSetOriginRequest) SetAttributes(v ApplicationRuleEngineSetOriginAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


