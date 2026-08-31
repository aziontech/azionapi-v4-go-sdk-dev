# S3EndpointAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;s3&#x60; - Simple Storage Service (S3) | 
**Attributes** | [**S3EndpointRequest**](S3EndpointRequest.md) |  | 

## Methods

### NewS3EndpointAttributesRequest

`func NewS3EndpointAttributesRequest(type_ string, attributes S3EndpointRequest, ) *S3EndpointAttributesRequest`

NewS3EndpointAttributesRequest instantiates a new S3EndpointAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EndpointAttributesRequestWithDefaults

`func NewS3EndpointAttributesRequestWithDefaults() *S3EndpointAttributesRequest`

NewS3EndpointAttributesRequestWithDefaults instantiates a new S3EndpointAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *S3EndpointAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3EndpointAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3EndpointAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *S3EndpointAttributesRequest) GetAttributes() S3EndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *S3EndpointAttributesRequest) GetAttributesOk() (*S3EndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *S3EndpointAttributesRequest) SetAttributes(v S3EndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


