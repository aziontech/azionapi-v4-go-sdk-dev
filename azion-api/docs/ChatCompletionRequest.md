# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatCompletionMessage**](ChatCompletionMessage.md) |  | 
**Azion** | Pointer to [**Azion**](Azion.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(messages []ChatCompletionMessage, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []ChatCompletionMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]ChatCompletionMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []ChatCompletionMessage)`

SetMessages sets Messages field to given value.


### GetAzion

`func (o *ChatCompletionRequest) GetAzion() Azion`

GetAzion returns the Azion field if non-nil, zero value otherwise.

### GetAzionOk

`func (o *ChatCompletionRequest) GetAzionOk() (*Azion, bool)`

GetAzionOk returns a tuple with the Azion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzion

`func (o *ChatCompletionRequest) SetAzion(v Azion)`

SetAzion sets Azion field to given value.

### HasAzion

`func (o *ChatCompletionRequest) HasAzion() bool`

HasAzion returns a boolean if a field has been set.

### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


