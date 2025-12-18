# PatchConnectorLiveIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | Pointer to [**ConnectorLiveIngestAttrsReq**](ConnectorLiveIngestAttrsReq.md) |  | [optional] 

## Methods

### NewPatchConnectorLiveIngestRequest

`func NewPatchConnectorLiveIngestRequest(type_ string, ) *PatchConnectorLiveIngestRequest`

NewPatchConnectorLiveIngestRequest instantiates a new PatchConnectorLiveIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchConnectorLiveIngestRequestWithDefaults

`func NewPatchConnectorLiveIngestRequestWithDefaults() *PatchConnectorLiveIngestRequest`

NewPatchConnectorLiveIngestRequestWithDefaults instantiates a new PatchConnectorLiveIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchConnectorLiveIngestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchConnectorLiveIngestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchConnectorLiveIngestRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchConnectorLiveIngestRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchConnectorLiveIngestRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchConnectorLiveIngestRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchConnectorLiveIngestRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchConnectorLiveIngestRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchConnectorLiveIngestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchConnectorLiveIngestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchConnectorLiveIngestRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PatchConnectorLiveIngestRequest) GetAttributes() ConnectorLiveIngestAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchConnectorLiveIngestRequest) GetAttributesOk() (*ConnectorLiveIngestAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchConnectorLiveIngestRequest) SetAttributes(v ConnectorLiveIngestAttrsReq)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchConnectorLiveIngestRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


