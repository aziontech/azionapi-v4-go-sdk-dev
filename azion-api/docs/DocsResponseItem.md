# DocsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tool** | Pointer to **string** | Name of the tool/node that produced this response | [optional] 
**Response** | Pointer to **string** | Content returned by the tool | [optional] 

## Methods

### NewDocsResponseItem

`func NewDocsResponseItem() *DocsResponseItem`

NewDocsResponseItem instantiates a new DocsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsResponseItemWithDefaults

`func NewDocsResponseItemWithDefaults() *DocsResponseItem`

NewDocsResponseItemWithDefaults instantiates a new DocsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTool

`func (o *DocsResponseItem) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *DocsResponseItem) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *DocsResponseItem) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *DocsResponseItem) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetResponse

`func (o *DocsResponseItem) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DocsResponseItem) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DocsResponseItem) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DocsResponseItem) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


