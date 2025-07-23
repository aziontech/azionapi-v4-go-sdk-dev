# EdgeApplicationRuleEngineRunFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**EdgeApplicationRuleEngineRunFunctionAttributesRequest**](EdgeApplicationRuleEngineRunFunctionAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineRunFunctionRequest

`func NewEdgeApplicationRuleEngineRunFunctionRequest(type_ string, attributes EdgeApplicationRuleEngineRunFunctionAttributesRequest, ) *EdgeApplicationRuleEngineRunFunctionRequest`

NewEdgeApplicationRuleEngineRunFunctionRequest instantiates a new EdgeApplicationRuleEngineRunFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineRunFunctionRequestWithDefaults

`func NewEdgeApplicationRuleEngineRunFunctionRequestWithDefaults() *EdgeApplicationRuleEngineRunFunctionRequest`

NewEdgeApplicationRuleEngineRunFunctionRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRunFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) GetAttributes() EdgeApplicationRuleEngineRunFunctionAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) GetAttributesOk() (*EdgeApplicationRuleEngineRunFunctionAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineRunFunctionRequest) SetAttributes(v EdgeApplicationRuleEngineRunFunctionAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


