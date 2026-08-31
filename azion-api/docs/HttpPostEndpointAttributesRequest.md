# HttpPostEndpointAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;standard&#x60; - Standard HTTP/HTTPS POST | 
**Attributes** | [**HttpPostEndpointRequest**](HttpPostEndpointRequest.md) |  | 

## Methods

### NewHttpPostEndpointAttributesRequest

`func NewHttpPostEndpointAttributesRequest(type_ string, attributes HttpPostEndpointRequest, ) *HttpPostEndpointAttributesRequest`

NewHttpPostEndpointAttributesRequest instantiates a new HttpPostEndpointAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPostEndpointAttributesRequestWithDefaults

`func NewHttpPostEndpointAttributesRequestWithDefaults() *HttpPostEndpointAttributesRequest`

NewHttpPostEndpointAttributesRequestWithDefaults instantiates a new HttpPostEndpointAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HttpPostEndpointAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpPostEndpointAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpPostEndpointAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *HttpPostEndpointAttributesRequest) GetAttributes() HttpPostEndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpPostEndpointAttributesRequest) GetAttributesOk() (*HttpPostEndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpPostEndpointAttributesRequest) SetAttributes(v HttpPostEndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


