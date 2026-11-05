# ChatCompletionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | The number of tokens in the generated completion | [optional] 
**PromptTokens** | Pointer to **int64** | The number of tokens in the prompt | [optional] 
**TotalTokens** | Pointer to **int64** | The total number of tokens used in the request and response | [optional] 

## Methods

### NewChatCompletionUsage

`func NewChatCompletionUsage() *ChatCompletionUsage`

NewChatCompletionUsage instantiates a new ChatCompletionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionUsageWithDefaults

`func NewChatCompletionUsageWithDefaults() *ChatCompletionUsage`

NewChatCompletionUsageWithDefaults instantiates a new ChatCompletionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *ChatCompletionUsage) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ChatCompletionUsage) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ChatCompletionUsage) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *ChatCompletionUsage) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetPromptTokens

`func (o *ChatCompletionUsage) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ChatCompletionUsage) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ChatCompletionUsage) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *ChatCompletionUsage) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ChatCompletionUsage) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ChatCompletionUsage) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ChatCompletionUsage) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ChatCompletionUsage) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


