# WorkloadBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** |  | 
**DeploymentId** | **string** |  | 
**Domains** | Pointer to **[]string** |  | [optional] 
**AutoDomainAllowAccess** | Pointer to **bool** |  | [optional] 
**Certificate** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewWorkloadBindingRequest

`func NewWorkloadBindingRequest(environmentId string, deploymentId string, ) *WorkloadBindingRequest`

NewWorkloadBindingRequest instantiates a new WorkloadBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBindingRequestWithDefaults

`func NewWorkloadBindingRequestWithDefaults() *WorkloadBindingRequest`

NewWorkloadBindingRequestWithDefaults instantiates a new WorkloadBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *WorkloadBindingRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WorkloadBindingRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WorkloadBindingRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDeploymentId

`func (o *WorkloadBindingRequest) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *WorkloadBindingRequest) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *WorkloadBindingRequest) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDomains

`func (o *WorkloadBindingRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *WorkloadBindingRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *WorkloadBindingRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *WorkloadBindingRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetAutoDomainAllowAccess

`func (o *WorkloadBindingRequest) GetAutoDomainAllowAccess() bool`

GetAutoDomainAllowAccess returns the AutoDomainAllowAccess field if non-nil, zero value otherwise.

### GetAutoDomainAllowAccessOk

`func (o *WorkloadBindingRequest) GetAutoDomainAllowAccessOk() (*bool, bool)`

GetAutoDomainAllowAccessOk returns a tuple with the AutoDomainAllowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDomainAllowAccess

`func (o *WorkloadBindingRequest) SetAutoDomainAllowAccess(v bool)`

SetAutoDomainAllowAccess sets AutoDomainAllowAccess field to given value.

### HasAutoDomainAllowAccess

`func (o *WorkloadBindingRequest) HasAutoDomainAllowAccess() bool`

HasAutoDomainAllowAccess returns a boolean if a field has been set.

### GetCertificate

`func (o *WorkloadBindingRequest) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WorkloadBindingRequest) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WorkloadBindingRequest) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WorkloadBindingRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WorkloadBindingRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WorkloadBindingRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


