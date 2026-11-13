# OutputKafkaEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**KafkaEndpointRequest**](KafkaEndpointRequest.md) |  | 

## Methods

### NewOutputKafkaEndpointRequest

`func NewOutputKafkaEndpointRequest(type_ string, attributes KafkaEndpointRequest, ) *OutputKafkaEndpointRequest`

NewOutputKafkaEndpointRequest instantiates a new OutputKafkaEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputKafkaEndpointRequestWithDefaults

`func NewOutputKafkaEndpointRequestWithDefaults() *OutputKafkaEndpointRequest`

NewOutputKafkaEndpointRequestWithDefaults instantiates a new OutputKafkaEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputKafkaEndpointRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputKafkaEndpointRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputKafkaEndpointRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OutputKafkaEndpointRequest) GetAttributes() KafkaEndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputKafkaEndpointRequest) GetAttributesOk() (*KafkaEndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputKafkaEndpointRequest) SetAttributes(v KafkaEndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


