# ChatCompletionChoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishReason** | Pointer to **string** | The reason the chat completion finished | [optional] 
**Index** | Pointer to **int64** | The index of the choice | [optional] 
**Message** | Pointer to [**ChatCompletionMessage**](ChatCompletionMessage.md) |  | [optional] 

## Methods

### NewChatCompletionChoicesInner

`func NewChatCompletionChoicesInner() *ChatCompletionChoicesInner`

NewChatCompletionChoicesInner instantiates a new ChatCompletionChoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionChoicesInnerWithDefaults

`func NewChatCompletionChoicesInnerWithDefaults() *ChatCompletionChoicesInner`

NewChatCompletionChoicesInnerWithDefaults instantiates a new ChatCompletionChoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishReason

`func (o *ChatCompletionChoicesInner) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *ChatCompletionChoicesInner) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *ChatCompletionChoicesInner) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *ChatCompletionChoicesInner) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetIndex

`func (o *ChatCompletionChoicesInner) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ChatCompletionChoicesInner) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ChatCompletionChoicesInner) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ChatCompletionChoicesInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMessage

`func (o *ChatCompletionChoicesInner) GetMessage() ChatCompletionMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatCompletionChoicesInner) GetMessageOk() (*ChatCompletionMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatCompletionChoicesInner) SetMessage(v ChatCompletionMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatCompletionChoicesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


