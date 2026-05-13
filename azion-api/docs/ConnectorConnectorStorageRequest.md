# ConnectorConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for ConnectorRequest | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Attributes** | [**ConnectorStorageAttributesRequest**](ConnectorStorageAttributesRequest.md) |  | 

## Methods

### NewConnectorConnectorStorageRequest

`func NewConnectorConnectorStorageRequest(type_ string, name string, attributes ConnectorStorageAttributesRequest, ) *ConnectorConnectorStorageRequest`

NewConnectorConnectorStorageRequest instantiates a new ConnectorConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConnectorStorageRequestWithDefaults

`func NewConnectorConnectorStorageRequestWithDefaults() *ConnectorConnectorStorageRequest`

NewConnectorConnectorStorageRequestWithDefaults instantiates a new ConnectorConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectorConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectorConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ConnectorConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttributes

`func (o *ConnectorConnectorStorageRequest) GetAttributes() ConnectorStorageAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorConnectorStorageRequest) GetAttributesOk() (*ConnectorStorageAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorConnectorStorageRequest) SetAttributes(v ConnectorStorageAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


