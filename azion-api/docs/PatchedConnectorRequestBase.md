# PatchedConnectorRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | Pointer to [**ConnectorStorageAttributesRequest**](ConnectorStorageAttributesRequest.md) |  | [optional] 

## Methods

### NewPatchedConnectorRequestBase

`func NewPatchedConnectorRequestBase(type_ string, ) *PatchedConnectorRequestBase`

NewPatchedConnectorRequestBase instantiates a new PatchedConnectorRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConnectorRequestBaseWithDefaults

`func NewPatchedConnectorRequestBaseWithDefaults() *PatchedConnectorRequestBase`

NewPatchedConnectorRequestBaseWithDefaults instantiates a new PatchedConnectorRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedConnectorRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConnectorRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConnectorRequestBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConnectorRequestBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedConnectorRequestBase) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedConnectorRequestBase) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedConnectorRequestBase) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedConnectorRequestBase) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedConnectorRequestBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedConnectorRequestBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedConnectorRequestBase) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PatchedConnectorRequestBase) GetAttributes() ConnectorStorageAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedConnectorRequestBase) GetAttributesOk() (*ConnectorStorageAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedConnectorRequestBase) SetAttributes(v ConnectorStorageAttributesRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedConnectorRequestBase) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


