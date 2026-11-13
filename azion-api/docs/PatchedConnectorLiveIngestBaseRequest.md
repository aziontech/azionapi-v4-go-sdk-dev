# PatchedConnectorLiveIngestBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to [**ConnectorLiveIngestAttributesRequest**](ConnectorLiveIngestAttributesRequest.md) |  | [optional] 

## Methods

### NewPatchedConnectorLiveIngestBaseRequest

`func NewPatchedConnectorLiveIngestBaseRequest() *PatchedConnectorLiveIngestBaseRequest`

NewPatchedConnectorLiveIngestBaseRequest instantiates a new PatchedConnectorLiveIngestBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConnectorLiveIngestBaseRequestWithDefaults

`func NewPatchedConnectorLiveIngestBaseRequestWithDefaults() *PatchedConnectorLiveIngestBaseRequest`

NewPatchedConnectorLiveIngestBaseRequestWithDefaults instantiates a new PatchedConnectorLiveIngestBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedConnectorLiveIngestBaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConnectorLiveIngestBaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConnectorLiveIngestBaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConnectorLiveIngestBaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedConnectorLiveIngestBaseRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedConnectorLiveIngestBaseRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedConnectorLiveIngestBaseRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedConnectorLiveIngestBaseRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedConnectorLiveIngestBaseRequest) GetAttributes() ConnectorLiveIngestAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedConnectorLiveIngestBaseRequest) GetAttributesOk() (*ConnectorLiveIngestAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedConnectorLiveIngestBaseRequest) SetAttributes(v ConnectorLiveIngestAttributesRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedConnectorLiveIngestBaseRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


