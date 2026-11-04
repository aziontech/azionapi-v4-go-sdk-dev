# PageConnectorBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | [**PageConnectorAttributes**](PageConnectorAttributes.md) |  | 

## Methods

### NewPageConnectorBase

`func NewPageConnectorBase(attributes PageConnectorAttributes, ) *PageConnectorBase`

NewPageConnectorBase instantiates a new PageConnectorBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageConnectorBaseWithDefaults

`func NewPageConnectorBaseWithDefaults() *PageConnectorBase`

NewPageConnectorBaseWithDefaults instantiates a new PageConnectorBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PageConnectorBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageConnectorBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageConnectorBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PageConnectorBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PageConnectorBase) GetAttributes() PageConnectorAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PageConnectorBase) GetAttributesOk() (*PageConnectorAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PageConnectorBase) SetAttributes(v PageConnectorAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


