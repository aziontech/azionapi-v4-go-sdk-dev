# DeployStrategyDefaultDeployStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for DeployStrategy | 
**Attributes** | [**DefaultDeployStrategyAttrs**](DefaultDeployStrategyAttrs.md) |  | 

## Methods

### NewDeployStrategyDefaultDeployStrategy

`func NewDeployStrategyDefaultDeployStrategy(type_ string, attributes DefaultDeployStrategyAttrs, ) *DeployStrategyDefaultDeployStrategy`

NewDeployStrategyDefaultDeployStrategy instantiates a new DeployStrategyDefaultDeployStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployStrategyDefaultDeployStrategyWithDefaults

`func NewDeployStrategyDefaultDeployStrategyWithDefaults() *DeployStrategyDefaultDeployStrategy`

NewDeployStrategyDefaultDeployStrategyWithDefaults instantiates a new DeployStrategyDefaultDeployStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeployStrategyDefaultDeployStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeployStrategyDefaultDeployStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeployStrategyDefaultDeployStrategy) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DeployStrategyDefaultDeployStrategy) GetAttributes() DefaultDeployStrategyAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeployStrategyDefaultDeployStrategy) GetAttributesOk() (*DefaultDeployStrategyAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeployStrategyDefaultDeployStrategy) SetAttributes(v DefaultDeployStrategyAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


