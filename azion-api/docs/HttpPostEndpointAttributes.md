# HttpPostEndpointAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;standard&#x60; - Standard HTTP/HTTPS POST | 
**Attributes** | [**HttpPostEndpoint**](HttpPostEndpoint.md) |  | 

## Methods

### NewHttpPostEndpointAttributes

`func NewHttpPostEndpointAttributes(type_ string, attributes HttpPostEndpoint, ) *HttpPostEndpointAttributes`

NewHttpPostEndpointAttributes instantiates a new HttpPostEndpointAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPostEndpointAttributesWithDefaults

`func NewHttpPostEndpointAttributesWithDefaults() *HttpPostEndpointAttributes`

NewHttpPostEndpointAttributesWithDefaults instantiates a new HttpPostEndpointAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HttpPostEndpointAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpPostEndpointAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpPostEndpointAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *HttpPostEndpointAttributes) GetAttributes() HttpPostEndpoint`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpPostEndpointAttributes) GetAttributesOk() (*HttpPostEndpoint, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpPostEndpointAttributes) SetAttributes(v HttpPostEndpoint)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


