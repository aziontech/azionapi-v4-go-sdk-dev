# DocsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatCompletionMessage**](ChatCompletionMessage.md) |  | 
**Azion** | Pointer to [**Azion**](Azion.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 
**DocsOptions** | Pointer to [**DocsOptions**](DocsOptions.md) |  | [optional] 

## Methods

### NewDocsRequest

`func NewDocsRequest(messages []ChatCompletionMessage, ) *DocsRequest`

NewDocsRequest instantiates a new DocsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsRequestWithDefaults

`func NewDocsRequestWithDefaults() *DocsRequest`

NewDocsRequestWithDefaults instantiates a new DocsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *DocsRequest) GetMessages() []ChatCompletionMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DocsRequest) GetMessagesOk() (*[]ChatCompletionMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DocsRequest) SetMessages(v []ChatCompletionMessage)`

SetMessages sets Messages field to given value.


### GetAzion

`func (o *DocsRequest) GetAzion() Azion`

GetAzion returns the Azion field if non-nil, zero value otherwise.

### GetAzionOk

`func (o *DocsRequest) GetAzionOk() (*Azion, bool)`

GetAzionOk returns a tuple with the Azion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzion

`func (o *DocsRequest) SetAzion(v Azion)`

SetAzion sets Azion field to given value.

### HasAzion

`func (o *DocsRequest) HasAzion() bool`

HasAzion returns a boolean if a field has been set.

### GetStream

`func (o *DocsRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *DocsRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *DocsRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *DocsRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetDocsOptions

`func (o *DocsRequest) GetDocsOptions() DocsOptions`

GetDocsOptions returns the DocsOptions field if non-nil, zero value otherwise.

### GetDocsOptionsOk

`func (o *DocsRequest) GetDocsOptionsOk() (*DocsOptions, bool)`

GetDocsOptionsOk returns a tuple with the DocsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsOptions

`func (o *DocsRequest) SetDocsOptions(v DocsOptions)`

SetDocsOptions sets DocsOptions field to given value.

### HasDocsOptions

`func (o *DocsRequest) HasDocsOptions() bool`

HasDocsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


