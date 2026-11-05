# PageConnectorBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | [**PageConnectorAttributesRequest**](PageConnectorAttributesRequest.md) |  | 

## Methods

### NewPageConnectorBaseRequest

`func NewPageConnectorBaseRequest(attributes PageConnectorAttributesRequest, ) *PageConnectorBaseRequest`

NewPageConnectorBaseRequest instantiates a new PageConnectorBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageConnectorBaseRequestWithDefaults

`func NewPageConnectorBaseRequestWithDefaults() *PageConnectorBaseRequest`

NewPageConnectorBaseRequestWithDefaults instantiates a new PageConnectorBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PageConnectorBaseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageConnectorBaseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageConnectorBaseRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PageConnectorBaseRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PageConnectorBaseRequest) GetAttributes() PageConnectorAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PageConnectorBaseRequest) GetAttributesOk() (*PageConnectorAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PageConnectorBaseRequest) SetAttributes(v PageConnectorAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


