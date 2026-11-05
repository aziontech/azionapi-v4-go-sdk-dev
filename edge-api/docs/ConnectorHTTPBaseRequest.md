# ConnectorHTTPBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 
**Attributes** | [**ConnectorHTTPAttributesRequest**](ConnectorHTTPAttributesRequest.md) |  | 

## Methods

### NewConnectorHTTPBaseRequest

`func NewConnectorHTTPBaseRequest(name string, type_ string, attributes ConnectorHTTPAttributesRequest, ) *ConnectorHTTPBaseRequest`

NewConnectorHTTPBaseRequest instantiates a new ConnectorHTTPBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorHTTPBaseRequestWithDefaults

`func NewConnectorHTTPBaseRequestWithDefaults() *ConnectorHTTPBaseRequest`

NewConnectorHTTPBaseRequestWithDefaults instantiates a new ConnectorHTTPBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorHTTPBaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorHTTPBaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorHTTPBaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ConnectorHTTPBaseRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorHTTPBaseRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorHTTPBaseRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorHTTPBaseRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *ConnectorHTTPBaseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorHTTPBaseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorHTTPBaseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ConnectorHTTPBaseRequest) GetAttributes() ConnectorHTTPAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorHTTPBaseRequest) GetAttributesOk() (*ConnectorHTTPAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorHTTPBaseRequest) SetAttributes(v ConnectorHTTPAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


