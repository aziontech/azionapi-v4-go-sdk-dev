# ConnectorConnectorHTTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ConnectorRequest | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Attributes** | [**ConnectorHTTPAttributesRequest**](ConnectorHTTPAttributesRequest.md) |  | 

## Methods

### NewConnectorConnectorHTTPRequest

`func NewConnectorConnectorHTTPRequest(type_ string, name string, attributes ConnectorHTTPAttributesRequest, ) *ConnectorConnectorHTTPRequest`

NewConnectorConnectorHTTPRequest instantiates a new ConnectorConnectorHTTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConnectorHTTPRequestWithDefaults

`func NewConnectorConnectorHTTPRequestWithDefaults() *ConnectorConnectorHTTPRequest`

NewConnectorConnectorHTTPRequestWithDefaults instantiates a new ConnectorConnectorHTTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectorConnectorHTTPRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorConnectorHTTPRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorConnectorHTTPRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectorConnectorHTTPRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorConnectorHTTPRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorConnectorHTTPRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ConnectorConnectorHTTPRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorConnectorHTTPRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorConnectorHTTPRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorConnectorHTTPRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttributes

`func (o *ConnectorConnectorHTTPRequest) GetAttributes() ConnectorHTTPAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorConnectorHTTPRequest) GetAttributesOk() (*ConnectorHTTPAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorConnectorHTTPRequest) SetAttributes(v ConnectorHTTPAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


