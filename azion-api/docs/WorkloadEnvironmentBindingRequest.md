# WorkloadEnvironmentBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Environment ID from external service | 
**Domain** | Pointer to **NullableString** | Custom domain or wildcard (*.domain) | [optional] 
**Certificate** | Pointer to **NullableInt64** | Certificate ID | [optional] 

## Methods

### NewWorkloadEnvironmentBindingRequest

`func NewWorkloadEnvironmentBindingRequest(environmentId string, ) *WorkloadEnvironmentBindingRequest`

NewWorkloadEnvironmentBindingRequest instantiates a new WorkloadEnvironmentBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEnvironmentBindingRequestWithDefaults

`func NewWorkloadEnvironmentBindingRequestWithDefaults() *WorkloadEnvironmentBindingRequest`

NewWorkloadEnvironmentBindingRequestWithDefaults instantiates a new WorkloadEnvironmentBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *WorkloadEnvironmentBindingRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WorkloadEnvironmentBindingRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WorkloadEnvironmentBindingRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDomain

`func (o *WorkloadEnvironmentBindingRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WorkloadEnvironmentBindingRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WorkloadEnvironmentBindingRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WorkloadEnvironmentBindingRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *WorkloadEnvironmentBindingRequest) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *WorkloadEnvironmentBindingRequest) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCertificate

`func (o *WorkloadEnvironmentBindingRequest) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WorkloadEnvironmentBindingRequest) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WorkloadEnvironmentBindingRequest) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WorkloadEnvironmentBindingRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WorkloadEnvironmentBindingRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WorkloadEnvironmentBindingRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


