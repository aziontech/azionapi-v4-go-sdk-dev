# PatchConnectorHTTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | Pointer to [**ConnectorHTTPAttrsReq**](ConnectorHTTPAttrsReq.md) |  | [optional] 

## Methods

### NewPatchConnectorHTTPRequest

`func NewPatchConnectorHTTPRequest(type_ string, ) *PatchConnectorHTTPRequest`

NewPatchConnectorHTTPRequest instantiates a new PatchConnectorHTTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchConnectorHTTPRequestWithDefaults

`func NewPatchConnectorHTTPRequestWithDefaults() *PatchConnectorHTTPRequest`

NewPatchConnectorHTTPRequestWithDefaults instantiates a new PatchConnectorHTTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchConnectorHTTPRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchConnectorHTTPRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchConnectorHTTPRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchConnectorHTTPRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchConnectorHTTPRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchConnectorHTTPRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchConnectorHTTPRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchConnectorHTTPRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchConnectorHTTPRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchConnectorHTTPRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchConnectorHTTPRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PatchConnectorHTTPRequest) GetAttributes() ConnectorHTTPAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchConnectorHTTPRequest) GetAttributesOk() (*ConnectorHTTPAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchConnectorHTTPRequest) SetAttributes(v ConnectorHTTPAttrsReq)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchConnectorHTTPRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


