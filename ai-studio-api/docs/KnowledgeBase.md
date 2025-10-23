# KnowledgeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KbId** | **string** |  | 
**AccountId** | **NullableInt64** |  | 
**LastEditor** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EmbeddingModel** | Pointer to **string** | * &#x60;text-embedding-3-small&#x60; - text-embedding-3-small | [optional] 
**EdgesqlDbId** | **string** |  | 
**LastModified** | **time.Time** |  | 

## Methods

### NewKnowledgeBase

`func NewKnowledgeBase(kbId string, accountId NullableInt64, lastEditor string, name string, edgesqlDbId string, lastModified time.Time, ) *KnowledgeBase`

NewKnowledgeBase instantiates a new KnowledgeBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeBaseWithDefaults

`func NewKnowledgeBaseWithDefaults() *KnowledgeBase`

NewKnowledgeBaseWithDefaults instantiates a new KnowledgeBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKbId

`func (o *KnowledgeBase) GetKbId() string`

GetKbId returns the KbId field if non-nil, zero value otherwise.

### GetKbIdOk

`func (o *KnowledgeBase) GetKbIdOk() (*string, bool)`

GetKbIdOk returns a tuple with the KbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbId

`func (o *KnowledgeBase) SetKbId(v string)`

SetKbId sets KbId field to given value.


### GetAccountId

`func (o *KnowledgeBase) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *KnowledgeBase) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *KnowledgeBase) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *KnowledgeBase) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *KnowledgeBase) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetLastEditor

`func (o *KnowledgeBase) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *KnowledgeBase) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *KnowledgeBase) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetName

`func (o *KnowledgeBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KnowledgeBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KnowledgeBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KnowledgeBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KnowledgeBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *KnowledgeBase) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *KnowledgeBase) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *KnowledgeBase) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *KnowledgeBase) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetEdgesqlDbId

`func (o *KnowledgeBase) GetEdgesqlDbId() string`

GetEdgesqlDbId returns the EdgesqlDbId field if non-nil, zero value otherwise.

### GetEdgesqlDbIdOk

`func (o *KnowledgeBase) GetEdgesqlDbIdOk() (*string, bool)`

GetEdgesqlDbIdOk returns a tuple with the EdgesqlDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgesqlDbId

`func (o *KnowledgeBase) SetEdgesqlDbId(v string)`

SetEdgesqlDbId sets EdgesqlDbId field to given value.


### GetLastModified

`func (o *KnowledgeBase) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *KnowledgeBase) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *KnowledgeBase) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


