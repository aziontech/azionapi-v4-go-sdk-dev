# PagePolymorphicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | [**PageConnectorAttributesRequest**](PageConnectorAttributesRequest.md) |  | 

## Methods

### NewPagePolymorphicRequest

`func NewPagePolymorphicRequest(attributes PageConnectorAttributesRequest, ) *PagePolymorphicRequest`

NewPagePolymorphicRequest instantiates a new PagePolymorphicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagePolymorphicRequestWithDefaults

`func NewPagePolymorphicRequestWithDefaults() *PagePolymorphicRequest`

NewPagePolymorphicRequestWithDefaults instantiates a new PagePolymorphicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PagePolymorphicRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PagePolymorphicRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PagePolymorphicRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PagePolymorphicRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PagePolymorphicRequest) GetAttributes() PageConnectorAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PagePolymorphicRequest) GetAttributesOk() (*PageConnectorAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PagePolymorphicRequest) SetAttributes(v PageConnectorAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


