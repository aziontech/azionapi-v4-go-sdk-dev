# WorkloDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Current** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Strategy** | [**DeployStrategyDefaultDeployStrategyRequest**](DeployStrategyDefaultDeployStrategyRequest.md) |  | 

## Methods

### NewWorkloDeploymentRequest

`func NewWorkloDeploymentRequest(name string, strategy DeployStrategyDefaultDeployStrategyRequest, ) *WorkloDeploymentRequest`

NewWorkloDeploymentRequest instantiates a new WorkloDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloDeploymentRequestWithDefaults

`func NewWorkloDeploymentRequestWithDefaults() *WorkloDeploymentRequest`

NewWorkloDeploymentRequestWithDefaults instantiates a new WorkloDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkloDeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloDeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloDeploymentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCurrent

`func (o *WorkloDeploymentRequest) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkloDeploymentRequest) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkloDeploymentRequest) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkloDeploymentRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetActive

`func (o *WorkloDeploymentRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkloDeploymentRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkloDeploymentRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WorkloDeploymentRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrategy

`func (o *WorkloDeploymentRequest) GetStrategy() DeployStrategyDefaultDeployStrategyRequest`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *WorkloDeploymentRequest) GetStrategyOk() (*DeployStrategyDefaultDeployStrategyRequest, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *WorkloDeploymentRequest) SetStrategy(v DeployStrategyDefaultDeployStrategyRequest)`

SetStrategy sets Strategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


