# WorkloadBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** |  | 
**DeploymentId** | **string** |  | 
**Domains** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkloadBinding

`func NewWorkloadBinding(environmentId string, deploymentId string, ) *WorkloadBinding`

NewWorkloadBinding instantiates a new WorkloadBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBindingWithDefaults

`func NewWorkloadBindingWithDefaults() *WorkloadBinding`

NewWorkloadBindingWithDefaults instantiates a new WorkloadBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *WorkloadBinding) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WorkloadBinding) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WorkloadBinding) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDeploymentId

`func (o *WorkloadBinding) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *WorkloadBinding) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *WorkloadBinding) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDomains

`func (o *WorkloadBinding) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *WorkloadBinding) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *WorkloadBinding) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *WorkloadBinding) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


