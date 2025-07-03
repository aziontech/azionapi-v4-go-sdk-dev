# PageDefaultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**PageDefaultAttributesRequest**](PageDefaultAttributesRequest.md) |  | [optional] 

## Methods

### NewPageDefaultRequest

`func NewPageDefaultRequest() *PageDefaultRequest`

NewPageDefaultRequest instantiates a new PageDefaultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageDefaultRequestWithDefaults

`func NewPageDefaultRequestWithDefaults() *PageDefaultRequest`

NewPageDefaultRequestWithDefaults instantiates a new PageDefaultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PageDefaultRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageDefaultRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageDefaultRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PageDefaultRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PageDefaultRequest) GetAttributes() PageDefaultAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PageDefaultRequest) GetAttributesOk() (*PageDefaultAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PageDefaultRequest) SetAttributes(v PageDefaultAttributesRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PageDefaultRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


