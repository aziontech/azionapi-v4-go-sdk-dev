# ConnectorPolymorphicConnectorLiveIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | 
**IsVersioned** | **bool** |  | 
**Version** | **NullableInt64** |  | 
**VersionState** | **NullableString** |  | 
**Attributes** | [**ConnectorLiveIngestAttributes**](ConnectorLiveIngestAttributes.md) |  | 

## Methods

### NewConnectorPolymorphicConnectorLiveIngest

`func NewConnectorPolymorphicConnectorLiveIngest(type_ string, id int64, name string, lastEditor string, lastModified time.Time, createdAt time.Time, productVersion string, isVersioned bool, version NullableInt64, versionState NullableString, attributes ConnectorLiveIngestAttributes, ) *ConnectorPolymorphicConnectorLiveIngest`

NewConnectorPolymorphicConnectorLiveIngest instantiates a new ConnectorPolymorphicConnectorLiveIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorPolymorphicConnectorLiveIngestWithDefaults

`func NewConnectorPolymorphicConnectorLiveIngestWithDefaults() *ConnectorPolymorphicConnectorLiveIngest`

NewConnectorPolymorphicConnectorLiveIngestWithDefaults instantiates a new ConnectorPolymorphicConnectorLiveIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetCreatedAt

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActive

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorPolymorphicConnectorLiveIngest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetIsVersioned

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetIsVersioned() bool`

GetIsVersioned returns the IsVersioned field if non-nil, zero value otherwise.

### GetIsVersionedOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetIsVersionedOk() (*bool, bool)`

GetIsVersionedOk returns a tuple with the IsVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioned

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetIsVersioned(v bool)`

SetIsVersioned sets IsVersioned field to given value.


### GetVersion

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetVersion(v int64)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConnectorPolymorphicConnectorLiveIngest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionState

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetVersionState() string`

GetVersionState returns the VersionState field if non-nil, zero value otherwise.

### GetVersionStateOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetVersionStateOk() (*string, bool)`

GetVersionStateOk returns a tuple with the VersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionState

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetVersionState(v string)`

SetVersionState sets VersionState field to given value.


### SetVersionStateNil

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetVersionStateNil(b bool)`

 SetVersionStateNil sets the value for VersionState to be an explicit nil

### UnsetVersionState
`func (o *ConnectorPolymorphicConnectorLiveIngest) UnsetVersionState()`

UnsetVersionState ensures that no value is present for VersionState, not even an explicit nil
### GetAttributes

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetAttributes() ConnectorLiveIngestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorPolymorphicConnectorLiveIngest) GetAttributesOk() (*ConnectorLiveIngestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorPolymorphicConnectorLiveIngest) SetAttributes(v ConnectorLiveIngestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


