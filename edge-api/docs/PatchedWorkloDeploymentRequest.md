# PatchedWorkloDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Strategy** | Pointer to [**DeployStrategyDefaultDeployStrategyRequest**](DeployStrategyDefaultDeployStrategyRequest.md) |  | [optional] 

## Methods

### NewPatchedWorkloDeploymentRequest

`func NewPatchedWorkloDeploymentRequest() *PatchedWorkloDeploymentRequest`

NewPatchedWorkloDeploymentRequest instantiates a new PatchedWorkloDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWorkloDeploymentRequestWithDefaults

`func NewPatchedWorkloDeploymentRequestWithDefaults() *PatchedWorkloDeploymentRequest`

NewPatchedWorkloDeploymentRequestWithDefaults instantiates a new PatchedWorkloDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWorkloDeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWorkloDeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWorkloDeploymentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWorkloDeploymentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrent

`func (o *PatchedWorkloDeploymentRequest) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PatchedWorkloDeploymentRequest) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PatchedWorkloDeploymentRequest) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *PatchedWorkloDeploymentRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetActive

`func (o *PatchedWorkloDeploymentRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedWorkloDeploymentRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedWorkloDeploymentRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedWorkloDeploymentRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrategy

`func (o *PatchedWorkloDeploymentRequest) GetStrategy() DeployStrategyDefaultDeployStrategyRequest`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *PatchedWorkloDeploymentRequest) GetStrategyOk() (*DeployStrategyDefaultDeployStrategyRequest, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *PatchedWorkloDeploymentRequest) SetStrategy(v DeployStrategyDefaultDeployStrategyRequest)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *PatchedWorkloDeploymentRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


