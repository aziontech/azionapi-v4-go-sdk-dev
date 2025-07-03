# PagePolymorphic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | [**PageConnectorAttributes**](PageConnectorAttributes.md) |  | 

## Methods

### NewPagePolymorphic

`func NewPagePolymorphic(attributes PageConnectorAttributes, ) *PagePolymorphic`

NewPagePolymorphic instantiates a new PagePolymorphic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagePolymorphicWithDefaults

`func NewPagePolymorphicWithDefaults() *PagePolymorphic`

NewPagePolymorphicWithDefaults instantiates a new PagePolymorphic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PagePolymorphic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PagePolymorphic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PagePolymorphic) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PagePolymorphic) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PagePolymorphic) GetAttributes() PageConnectorAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PagePolymorphic) GetAttributesOk() (*PageConnectorAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PagePolymorphic) SetAttributes(v PageConnectorAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


