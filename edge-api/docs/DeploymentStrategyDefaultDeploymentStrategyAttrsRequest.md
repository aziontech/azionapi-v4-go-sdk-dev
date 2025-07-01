# DeploymentStrategyDefaultDeploymentStrategyAttrsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyType** | **string** | * &#x60;default&#x60; - Default | 
**Attributes** | [**DefaultDeploymentStrategyRequest**](DefaultDeploymentStrategyRequest.md) |  | 

## Methods

### NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequest

`func NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequest(strategyType string, attributes DefaultDeploymentStrategyRequest, ) *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest`

NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequest instantiates a new DeploymentStrategyDefaultDeploymentStrategyAttrsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequestWithDefaults

`func NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequestWithDefaults() *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest`

NewDeploymentStrategyDefaultDeploymentStrategyAttrsRequestWithDefaults instantiates a new DeploymentStrategyDefaultDeploymentStrategyAttrsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategyType

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.


### GetAttributes

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) GetAttributes() DefaultDeploymentStrategyRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) GetAttributesOk() (*DefaultDeploymentStrategyRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeploymentStrategyDefaultDeploymentStrategyAttrsRequest) SetAttributes(v DefaultDeploymentStrategyRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


