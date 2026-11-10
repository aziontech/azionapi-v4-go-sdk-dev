# WorkloadDomainAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **NullableString** | Why the domain cannot be used, or null when it is available.  * &#x60;already_in_use&#x60; - already_in_use * &#x60;invalid_format&#x60; - invalid_format | [optional] 

## Methods

### NewWorkloadDomainAvailability

`func NewWorkloadDomainAvailability() *WorkloadDomainAvailability`

NewWorkloadDomainAvailability instantiates a new WorkloadDomainAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDomainAvailabilityWithDefaults

`func NewWorkloadDomainAvailabilityWithDefaults() *WorkloadDomainAvailability`

NewWorkloadDomainAvailabilityWithDefaults instantiates a new WorkloadDomainAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *WorkloadDomainAvailability) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WorkloadDomainAvailability) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WorkloadDomainAvailability) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WorkloadDomainAvailability) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAvailable

`func (o *WorkloadDomainAvailability) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WorkloadDomainAvailability) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WorkloadDomainAvailability) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *WorkloadDomainAvailability) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetReason

`func (o *WorkloadDomainAvailability) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkloadDomainAvailability) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkloadDomainAvailability) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkloadDomainAvailability) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *WorkloadDomainAvailability) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *WorkloadDomainAvailability) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


