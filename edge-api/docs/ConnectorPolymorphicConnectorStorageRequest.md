# ConnectorPolymorphicConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Attributes** | [**ConnectorStorageAttributesRequest**](ConnectorStorageAttributesRequest.md) |  | 

## Methods

### NewConnectorPolymorphicConnectorStorageRequest

`func NewConnectorPolymorphicConnectorStorageRequest(type_ string, name string, attributes ConnectorStorageAttributesRequest, ) *ConnectorPolymorphicConnectorStorageRequest`

NewConnectorPolymorphicConnectorStorageRequest instantiates a new ConnectorPolymorphicConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorPolymorphicConnectorStorageRequestWithDefaults

`func NewConnectorPolymorphicConnectorStorageRequestWithDefaults() *ConnectorPolymorphicConnectorStorageRequest`

NewConnectorPolymorphicConnectorStorageRequestWithDefaults instantiates a new ConnectorPolymorphicConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorPolymorphicConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorPolymorphicConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorPolymorphicConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorPolymorphicConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttributes

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetAttributes() ConnectorStorageAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorPolymorphicConnectorStorageRequest) GetAttributesOk() (*ConnectorStorageAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorPolymorphicConnectorStorageRequest) SetAttributes(v ConnectorStorageAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


