# ConnectorLiveIngestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVersioned** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **NullableInt64** |  | [optional] 
**VersionId** | Pointer to **NullableString** | ID of the version metadata (use in /versions/{id} URLs) | [optional] 
**VersionState** | Pointer to **NullableString** | Build state of this version (queued, building, ready, error, ...) | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**LastEditor** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Attributes** | [**ConnectorLiveIngestAttributes**](ConnectorLiveIngestAttributes.md) |  | 

## Methods

### NewConnectorLiveIngestBase

`func NewConnectorLiveIngestBase(name string, type_ string, attributes ConnectorLiveIngestAttributes, ) *ConnectorLiveIngestBase`

NewConnectorLiveIngestBase instantiates a new ConnectorLiveIngestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorLiveIngestBaseWithDefaults

`func NewConnectorLiveIngestBaseWithDefaults() *ConnectorLiveIngestBase`

NewConnectorLiveIngestBaseWithDefaults instantiates a new ConnectorLiveIngestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVersioned

`func (o *ConnectorLiveIngestBase) GetIsVersioned() bool`

GetIsVersioned returns the IsVersioned field if non-nil, zero value otherwise.

### GetIsVersionedOk

`func (o *ConnectorLiveIngestBase) GetIsVersionedOk() (*bool, bool)`

GetIsVersionedOk returns a tuple with the IsVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioned

`func (o *ConnectorLiveIngestBase) SetIsVersioned(v bool)`

SetIsVersioned sets IsVersioned field to given value.

### HasIsVersioned

`func (o *ConnectorLiveIngestBase) HasIsVersioned() bool`

HasIsVersioned returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectorLiveIngestBase) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorLiveIngestBase) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorLiveIngestBase) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorLiveIngestBase) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConnectorLiveIngestBase) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConnectorLiveIngestBase) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionId

`func (o *ConnectorLiveIngestBase) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ConnectorLiveIngestBase) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ConnectorLiveIngestBase) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ConnectorLiveIngestBase) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### SetVersionIdNil

`func (o *ConnectorLiveIngestBase) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *ConnectorLiveIngestBase) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil
### GetVersionState

`func (o *ConnectorLiveIngestBase) GetVersionState() string`

GetVersionState returns the VersionState field if non-nil, zero value otherwise.

### GetVersionStateOk

`func (o *ConnectorLiveIngestBase) GetVersionStateOk() (*string, bool)`

GetVersionStateOk returns a tuple with the VersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionState

`func (o *ConnectorLiveIngestBase) SetVersionState(v string)`

SetVersionState sets VersionState field to given value.

### HasVersionState

`func (o *ConnectorLiveIngestBase) HasVersionState() bool`

HasVersionState returns a boolean if a field has been set.

### SetVersionStateNil

`func (o *ConnectorLiveIngestBase) SetVersionStateNil(b bool)`

 SetVersionStateNil sets the value for VersionState to be an explicit nil

### UnsetVersionState
`func (o *ConnectorLiveIngestBase) UnsetVersionState()`

UnsetVersionState ensures that no value is present for VersionState, not even an explicit nil
### GetId

`func (o *ConnectorLiveIngestBase) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorLiveIngestBase) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorLiveIngestBase) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorLiveIngestBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectorLiveIngestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorLiveIngestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorLiveIngestBase) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ConnectorLiveIngestBase) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ConnectorLiveIngestBase) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ConnectorLiveIngestBase) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ConnectorLiveIngestBase) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ConnectorLiveIngestBase) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ConnectorLiveIngestBase) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ConnectorLiveIngestBase) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ConnectorLiveIngestBase) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorLiveIngestBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorLiveIngestBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorLiveIngestBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorLiveIngestBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetActive

`func (o *ConnectorLiveIngestBase) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorLiveIngestBase) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorLiveIngestBase) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorLiveIngestBase) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ConnectorLiveIngestBase) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ConnectorLiveIngestBase) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ConnectorLiveIngestBase) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ConnectorLiveIngestBase) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetType

`func (o *ConnectorLiveIngestBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorLiveIngestBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorLiveIngestBase) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ConnectorLiveIngestBase) GetAttributes() ConnectorLiveIngestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorLiveIngestBase) GetAttributesOk() (*ConnectorLiveIngestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorLiveIngestBase) SetAttributes(v ConnectorLiveIngestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


