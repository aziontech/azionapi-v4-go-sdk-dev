# AWS4HMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | * &#x60;aws4_hmac_sha256&#x60; - AWS for HMAC - SHA256 | [optional] 
**Attributes** | [**AWS4HMAttrsReq**](AWS4HMAttrsReq.md) |  | 

## Methods

### NewAWS4HMRequest

`func NewAWS4HMRequest(attributes AWS4HMAttrsReq, ) *AWS4HMRequest`

NewAWS4HMRequest instantiates a new AWS4HMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWS4HMRequestWithDefaults

`func NewAWS4HMRequestWithDefaults() *AWS4HMRequest`

NewAWS4HMRequestWithDefaults instantiates a new AWS4HMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWS4HMRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWS4HMRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWS4HMRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AWS4HMRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *AWS4HMRequest) GetAttributes() AWS4HMAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AWS4HMRequest) GetAttributesOk() (*AWS4HMAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AWS4HMRequest) SetAttributes(v AWS4HMAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


