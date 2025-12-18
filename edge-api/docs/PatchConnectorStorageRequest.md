# PatchConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | Pointer to [**ConnectorStorageAttrsReq**](ConnectorStorageAttrsReq.md) |  | [optional] 

## Methods

### NewPatchConnectorStorageRequest

`func NewPatchConnectorStorageRequest(type_ string, ) *PatchConnectorStorageRequest`

NewPatchConnectorStorageRequest instantiates a new PatchConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchConnectorStorageRequestWithDefaults

`func NewPatchConnectorStorageRequestWithDefaults() *PatchConnectorStorageRequest`

NewPatchConnectorStorageRequestWithDefaults instantiates a new PatchConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchConnectorStorageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PatchConnectorStorageRequest) GetAttributes() ConnectorStorageAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchConnectorStorageRequest) GetAttributesOk() (*ConnectorStorageAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchConnectorStorageRequest) SetAttributes(v ConnectorStorageAttrsReq)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchConnectorStorageRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


