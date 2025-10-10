# KnowledgeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KbId** | **string** |  | 
**AccountId** | **NullableInt64** |  | 
**CreatedBy** | **NullableString** |  | 
**UpdatedBy** | **NullableString** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EmbeddingModel** | Pointer to **string** | * &#x60;text-embedding-3-small&#x60; - text-embedding-3-small | [optional] 
**EdgesqlDbId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewKnowledgeBase

`func NewKnowledgeBase(kbId string, accountId NullableInt64, createdBy NullableString, updatedBy NullableString, name string, edgesqlDbId string, createdAt time.Time, updatedAt time.Time, ) *KnowledgeBase`

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
### GetCreatedBy

`func (o *KnowledgeBase) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KnowledgeBase) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KnowledgeBase) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *KnowledgeBase) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *KnowledgeBase) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *KnowledgeBase) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *KnowledgeBase) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *KnowledgeBase) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### SetUpdatedByNil

`func (o *KnowledgeBase) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *KnowledgeBase) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
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


### GetCreatedAt

`func (o *KnowledgeBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KnowledgeBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KnowledgeBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KnowledgeBase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KnowledgeBase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KnowledgeBase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


