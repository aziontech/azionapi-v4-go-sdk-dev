# OutputS3Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**S3Endpoint**](S3Endpoint.md) |  | 

## Methods

### NewOutputS3Endpoint

`func NewOutputS3Endpoint(type_ string, attributes S3Endpoint, ) *OutputS3Endpoint`

NewOutputS3Endpoint instantiates a new OutputS3Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputS3EndpointWithDefaults

`func NewOutputS3EndpointWithDefaults() *OutputS3Endpoint`

NewOutputS3EndpointWithDefaults instantiates a new OutputS3Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputS3Endpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputS3Endpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputS3Endpoint) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OutputS3Endpoint) GetAttributes() S3Endpoint`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputS3Endpoint) GetAttributesOk() (*S3Endpoint, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputS3Endpoint) SetAttributes(v S3Endpoint)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


