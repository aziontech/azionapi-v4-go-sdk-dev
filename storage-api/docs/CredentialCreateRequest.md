# CredentialCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Capabilities** | **[]string** |  | 
**Bucket** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCredentialCreateRequest

`func NewCredentialCreateRequest(name string, capabilities []string, ) *CredentialCreateRequest`

NewCredentialCreateRequest instantiates a new CredentialCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialCreateRequestWithDefaults

`func NewCredentialCreateRequestWithDefaults() *CredentialCreateRequest`

NewCredentialCreateRequestWithDefaults instantiates a new CredentialCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCapabilities

`func (o *CredentialCreateRequest) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CredentialCreateRequest) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CredentialCreateRequest) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetBucket

`func (o *CredentialCreateRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CredentialCreateRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CredentialCreateRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CredentialCreateRequest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CredentialCreateRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CredentialCreateRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CredentialCreateRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CredentialCreateRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


