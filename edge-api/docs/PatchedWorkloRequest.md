# PatchedWorkloRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Infrastructure** | Pointer to **int64** | * &#x60;1&#x60; - Production Infrastructure (All Edge Locations) * &#x60;2&#x60; - Staging Infrastructure | [optional] 
**Tls** | Pointer to [**TLSWorkloRequest**](TLSWorkloRequest.md) |  | [optional] 
**Protocols** | Pointer to [**ProtocolsRequest**](ProtocolsRequest.md) |  | [optional] 
**Mtls** | Pointer to [**MTLSRequest**](MTLSRequest.md) |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**WorkloadDomainAllowAccess** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedWorkloRequest

`func NewPatchedWorkloRequest() *PatchedWorkloRequest`

NewPatchedWorkloRequest instantiates a new PatchedWorkloRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWorkloRequestWithDefaults

`func NewPatchedWorkloRequestWithDefaults() *PatchedWorkloRequest`

NewPatchedWorkloRequestWithDefaults instantiates a new PatchedWorkloRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWorkloRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWorkloRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWorkloRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWorkloRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedWorkloRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedWorkloRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedWorkloRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedWorkloRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInfrastructure

`func (o *PatchedWorkloRequest) GetInfrastructure() int64`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *PatchedWorkloRequest) GetInfrastructureOk() (*int64, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *PatchedWorkloRequest) SetInfrastructure(v int64)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *PatchedWorkloRequest) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetTls

`func (o *PatchedWorkloRequest) GetTls() TLSWorkloRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *PatchedWorkloRequest) GetTlsOk() (*TLSWorkloRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *PatchedWorkloRequest) SetTls(v TLSWorkloRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *PatchedWorkloRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *PatchedWorkloRequest) GetProtocols() ProtocolsRequest`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PatchedWorkloRequest) GetProtocolsOk() (*ProtocolsRequest, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PatchedWorkloRequest) SetProtocols(v ProtocolsRequest)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *PatchedWorkloRequest) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *PatchedWorkloRequest) GetMtls() MTLSRequest`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *PatchedWorkloRequest) GetMtlsOk() (*MTLSRequest, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *PatchedWorkloRequest) SetMtls(v MTLSRequest)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *PatchedWorkloRequest) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *PatchedWorkloRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *PatchedWorkloRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *PatchedWorkloRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *PatchedWorkloRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetWorkloadDomainAllowAccess

`func (o *PatchedWorkloRequest) GetWorkloadDomainAllowAccess() bool`

GetWorkloadDomainAllowAccess returns the WorkloadDomainAllowAccess field if non-nil, zero value otherwise.

### GetWorkloadDomainAllowAccessOk

`func (o *PatchedWorkloRequest) GetWorkloadDomainAllowAccessOk() (*bool, bool)`

GetWorkloadDomainAllowAccessOk returns a tuple with the WorkloadDomainAllowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDomainAllowAccess

`func (o *PatchedWorkloRequest) SetWorkloadDomainAllowAccess(v bool)`

SetWorkloadDomainAllowAccess sets WorkloadDomainAllowAccess field to given value.

### HasWorkloadDomainAllowAccess

`func (o *PatchedWorkloRequest) HasWorkloadDomainAllowAccess() bool`

HasWorkloadDomainAllowAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


