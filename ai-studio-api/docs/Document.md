# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**KbId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceUri** | Pointer to **string** |  | [optional] 
**ChunkStrategy** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | **string** | * &#x60;creating&#x60; - creating * &#x60;processing&#x60; - processing * &#x60;created&#x60; - created * &#x60;error&#x60; - error | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDocument

`func NewDocument(documentId string, kbId string, status string, createdAt time.Time, updatedAt time.Time, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *Document) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Document) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Document) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetKbId

`func (o *Document) GetKbId() string`

GetKbId returns the KbId field if non-nil, zero value otherwise.

### GetKbIdOk

`func (o *Document) GetKbIdOk() (*string, bool)`

GetKbIdOk returns a tuple with the KbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbId

`func (o *Document) SetKbId(v string)`

SetKbId sets KbId field to given value.


### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Document) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Document) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Document) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Document) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Document) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Document) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Document) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUri

`func (o *Document) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *Document) GetSourceUriOk() (*string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUri

`func (o *Document) SetSourceUri(v string)`

SetSourceUri sets SourceUri field to given value.

### HasSourceUri

`func (o *Document) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### GetChunkStrategy

`func (o *Document) GetChunkStrategy() map[string]interface{}`

GetChunkStrategy returns the ChunkStrategy field if non-nil, zero value otherwise.

### GetChunkStrategyOk

`func (o *Document) GetChunkStrategyOk() (*map[string]interface{}, bool)`

GetChunkStrategyOk returns a tuple with the ChunkStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkStrategy

`func (o *Document) SetChunkStrategy(v map[string]interface{})`

SetChunkStrategy sets ChunkStrategy field to given value.

### HasChunkStrategy

`func (o *Document) HasChunkStrategy() bool`

HasChunkStrategy returns a boolean if a field has been set.

### GetStatus

`func (o *Document) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Document) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Document) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Document) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Document) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Document) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Document) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Document) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Document) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


