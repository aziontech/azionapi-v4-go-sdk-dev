# OutputS3EndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**S3EndpointRequest**](S3EndpointRequest.md) |  | 

## Methods

### NewOutputS3EndpointRequest

`func NewOutputS3EndpointRequest(type_ string, attributes S3EndpointRequest, ) *OutputS3EndpointRequest`

NewOutputS3EndpointRequest instantiates a new OutputS3EndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputS3EndpointRequestWithDefaults

`func NewOutputS3EndpointRequestWithDefaults() *OutputS3EndpointRequest`

NewOutputS3EndpointRequestWithDefaults instantiates a new OutputS3EndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputS3EndpointRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputS3EndpointRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputS3EndpointRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OutputS3EndpointRequest) GetAttributes() S3EndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputS3EndpointRequest) GetAttributesOk() (*S3EndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputS3EndpointRequest) SetAttributes(v S3EndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


