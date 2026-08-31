# ElasticsearchEndpointAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;elasticsearch&#x60; - Elasticsearch | 
**Attributes** | [**ElasticsearchEndpointRequest**](ElasticsearchEndpointRequest.md) |  | 

## Methods

### NewElasticsearchEndpointAttributesRequest

`func NewElasticsearchEndpointAttributesRequest(type_ string, attributes ElasticsearchEndpointRequest, ) *ElasticsearchEndpointAttributesRequest`

NewElasticsearchEndpointAttributesRequest instantiates a new ElasticsearchEndpointAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticsearchEndpointAttributesRequestWithDefaults

`func NewElasticsearchEndpointAttributesRequestWithDefaults() *ElasticsearchEndpointAttributesRequest`

NewElasticsearchEndpointAttributesRequestWithDefaults instantiates a new ElasticsearchEndpointAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ElasticsearchEndpointAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ElasticsearchEndpointAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ElasticsearchEndpointAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ElasticsearchEndpointAttributesRequest) GetAttributes() ElasticsearchEndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ElasticsearchEndpointAttributesRequest) GetAttributesOk() (*ElasticsearchEndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ElasticsearchEndpointAttributesRequest) SetAttributes(v ElasticsearchEndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


