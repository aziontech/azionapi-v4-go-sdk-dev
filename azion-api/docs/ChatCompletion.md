# ChatCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to [**[]ChatCompletionChoicesInner**](ChatCompletionChoicesInner.md) |  | [optional] 
**Created** | Pointer to **int64** | The Unix timestamp (in seconds) of when the chat completion was created | [optional] 
**Id** | Pointer to **string** | A unique identifier for the chat completion | [optional] 
**Model** | Pointer to **string** | The model used for the chat completion | [optional] 
**Object** | Pointer to **string** | The object type, which is always \&quot;chat.completion\&quot; for this endpoint | [optional] 
**Usage** | Pointer to [**ChatCompletionUsage**](ChatCompletionUsage.md) |  | [optional] 

## Methods

### NewChatCompletion

`func NewChatCompletion() *ChatCompletion`

NewChatCompletion instantiates a new ChatCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionWithDefaults

`func NewChatCompletionWithDefaults() *ChatCompletion`

NewChatCompletionWithDefaults instantiates a new ChatCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *ChatCompletion) GetChoices() []ChatCompletionChoicesInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ChatCompletion) GetChoicesOk() (*[]ChatCompletionChoicesInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ChatCompletion) SetChoices(v []ChatCompletionChoicesInner)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ChatCompletion) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetCreated

`func (o *ChatCompletion) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatCompletion) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatCompletion) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChatCompletion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ChatCompletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatCompletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatCompletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatCompletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *ChatCompletion) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletion) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletion) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatCompletion) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *ChatCompletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatCompletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatCompletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChatCompletion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUsage

`func (o *ChatCompletion) GetUsage() ChatCompletionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatCompletion) GetUsageOk() (*ChatCompletionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatCompletion) SetUsage(v ChatCompletionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatCompletion) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


