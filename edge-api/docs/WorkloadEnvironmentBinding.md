# WorkloadEnvironmentBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Environment ID from external service | 
**Domain** | Pointer to **NullableString** | Custom domain or wildcard (*.domain) | [optional] 
**Certificate** | Pointer to **NullableInt64** | Certificate ID | [optional] 
**CanonicalUrl** | **string** | Generated canonical URL | 

## Methods

### NewWorkloadEnvironmentBinding

`func NewWorkloadEnvironmentBinding(environmentId string, canonicalUrl string, ) *WorkloadEnvironmentBinding`

NewWorkloadEnvironmentBinding instantiates a new WorkloadEnvironmentBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEnvironmentBindingWithDefaults

`func NewWorkloadEnvironmentBindingWithDefaults() *WorkloadEnvironmentBinding`

NewWorkloadEnvironmentBindingWithDefaults instantiates a new WorkloadEnvironmentBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *WorkloadEnvironmentBinding) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WorkloadEnvironmentBinding) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WorkloadEnvironmentBinding) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDomain

`func (o *WorkloadEnvironmentBinding) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WorkloadEnvironmentBinding) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WorkloadEnvironmentBinding) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WorkloadEnvironmentBinding) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *WorkloadEnvironmentBinding) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *WorkloadEnvironmentBinding) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCertificate

`func (o *WorkloadEnvironmentBinding) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WorkloadEnvironmentBinding) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WorkloadEnvironmentBinding) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WorkloadEnvironmentBinding) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WorkloadEnvironmentBinding) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WorkloadEnvironmentBinding) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCanonicalUrl

`func (o *WorkloadEnvironmentBinding) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WorkloadEnvironmentBinding) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WorkloadEnvironmentBinding) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


