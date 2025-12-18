# WorkloRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Infrastructure** | Pointer to **int64** | * &#x60;1&#x60; - Production Infrastructure (All Edge Locations) * &#x60;2&#x60; - Staging Infrastructure | [optional] 
**Tls** | Pointer to [**TLSWorkloRequest**](TLSWorkloRequest.md) |  | [optional] 
**Protocols** | Pointer to [**ProtocolsRequest**](ProtocolsRequest.md) |  | [optional] 
**Mtls** | Pointer to [**MTLSRequest**](MTLSRequest.md) |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**WorkloadDomainAllowAccess** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkloRequest

`func NewWorkloRequest(name string, ) *WorkloRequest`

NewWorkloRequest instantiates a new WorkloRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloRequestWithDefaults

`func NewWorkloRequestWithDefaults() *WorkloRequest`

NewWorkloRequestWithDefaults instantiates a new WorkloRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkloRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *WorkloRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkloRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkloRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WorkloRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInfrastructure

`func (o *WorkloRequest) GetInfrastructure() int64`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *WorkloRequest) GetInfrastructureOk() (*int64, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *WorkloRequest) SetInfrastructure(v int64)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *WorkloRequest) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetTls

`func (o *WorkloRequest) GetTls() TLSWorkloRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *WorkloRequest) GetTlsOk() (*TLSWorkloRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *WorkloRequest) SetTls(v TLSWorkloRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *WorkloRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *WorkloRequest) GetProtocols() ProtocolsRequest`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *WorkloRequest) GetProtocolsOk() (*ProtocolsRequest, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *WorkloRequest) SetProtocols(v ProtocolsRequest)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *WorkloRequest) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *WorkloRequest) GetMtls() MTLSRequest`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *WorkloRequest) GetMtlsOk() (*MTLSRequest, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *WorkloRequest) SetMtls(v MTLSRequest)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *WorkloRequest) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *WorkloRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *WorkloRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *WorkloRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *WorkloRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetWorkloadDomainAllowAccess

`func (o *WorkloRequest) GetWorkloadDomainAllowAccess() bool`

GetWorkloadDomainAllowAccess returns the WorkloadDomainAllowAccess field if non-nil, zero value otherwise.

### GetWorkloadDomainAllowAccessOk

`func (o *WorkloRequest) GetWorkloadDomainAllowAccessOk() (*bool, bool)`

GetWorkloadDomainAllowAccessOk returns a tuple with the WorkloadDomainAllowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDomainAllowAccess

`func (o *WorkloRequest) SetWorkloadDomainAllowAccess(v bool)`

SetWorkloadDomainAllowAccess sets WorkloadDomainAllowAccess field to given value.

### HasWorkloadDomainAllowAccess

`func (o *WorkloRequest) HasWorkloadDomainAllowAccess() bool`

HasWorkloadDomainAllowAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


