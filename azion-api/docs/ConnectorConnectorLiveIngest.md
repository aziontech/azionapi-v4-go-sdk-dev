# ConnectorConnectorLiveIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for Connector | 
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
**VersionId** | **NullableString** |  | 
**Attributes** | [**ConnectorLiveIngestAttributes**](ConnectorLiveIngestAttributes.md) |  | 

## Methods

### NewConnectorConnectorLiveIngest

`func NewConnectorConnectorLiveIngest(type_ string, id int64, name string, lastEditor string, lastModified time.Time, createdAt time.Time, productVersion string, isVersioned bool, version NullableInt64, versionState NullableString, versionId NullableString, attributes ConnectorLiveIngestAttributes, ) *ConnectorConnectorLiveIngest`

NewConnectorConnectorLiveIngest instantiates a new ConnectorConnectorLiveIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConnectorLiveIngestWithDefaults

`func NewConnectorConnectorLiveIngestWithDefaults() *ConnectorConnectorLiveIngest`

NewConnectorConnectorLiveIngestWithDefaults instantiates a new ConnectorConnectorLiveIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectorConnectorLiveIngest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorConnectorLiveIngest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorConnectorLiveIngest) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ConnectorConnectorLiveIngest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorConnectorLiveIngest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorConnectorLiveIngest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectorConnectorLiveIngest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorConnectorLiveIngest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorConnectorLiveIngest) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ConnectorConnectorLiveIngest) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ConnectorConnectorLiveIngest) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ConnectorConnectorLiveIngest) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ConnectorConnectorLiveIngest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ConnectorConnectorLiveIngest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ConnectorConnectorLiveIngest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetCreatedAt

`func (o *ConnectorConnectorLiveIngest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorConnectorLiveIngest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorConnectorLiveIngest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActive

`func (o *ConnectorConnectorLiveIngest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorConnectorLiveIngest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorConnectorLiveIngest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorConnectorLiveIngest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ConnectorConnectorLiveIngest) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ConnectorConnectorLiveIngest) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ConnectorConnectorLiveIngest) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetIsVersioned

`func (o *ConnectorConnectorLiveIngest) GetIsVersioned() bool`

GetIsVersioned returns the IsVersioned field if non-nil, zero value otherwise.

### GetIsVersionedOk

`func (o *ConnectorConnectorLiveIngest) GetIsVersionedOk() (*bool, bool)`

GetIsVersionedOk returns a tuple with the IsVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioned

`func (o *ConnectorConnectorLiveIngest) SetIsVersioned(v bool)`

SetIsVersioned sets IsVersioned field to given value.


### GetVersion

`func (o *ConnectorConnectorLiveIngest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorConnectorLiveIngest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorConnectorLiveIngest) SetVersion(v int64)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *ConnectorConnectorLiveIngest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConnectorConnectorLiveIngest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionState

`func (o *ConnectorConnectorLiveIngest) GetVersionState() string`

GetVersionState returns the VersionState field if non-nil, zero value otherwise.

### GetVersionStateOk

`func (o *ConnectorConnectorLiveIngest) GetVersionStateOk() (*string, bool)`

GetVersionStateOk returns a tuple with the VersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionState

`func (o *ConnectorConnectorLiveIngest) SetVersionState(v string)`

SetVersionState sets VersionState field to given value.


### SetVersionStateNil

`func (o *ConnectorConnectorLiveIngest) SetVersionStateNil(b bool)`

 SetVersionStateNil sets the value for VersionState to be an explicit nil

### UnsetVersionState
`func (o *ConnectorConnectorLiveIngest) UnsetVersionState()`

UnsetVersionState ensures that no value is present for VersionState, not even an explicit nil
### GetVersionId

`func (o *ConnectorConnectorLiveIngest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ConnectorConnectorLiveIngest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ConnectorConnectorLiveIngest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### SetVersionIdNil

`func (o *ConnectorConnectorLiveIngest) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *ConnectorConnectorLiveIngest) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil
### GetAttributes

`func (o *ConnectorConnectorLiveIngest) GetAttributes() ConnectorLiveIngestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorConnectorLiveIngest) GetAttributesOk() (*ConnectorLiveIngestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorConnectorLiveIngest) SetAttributes(v ConnectorLiveIngestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


