# ResponseCustomPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**CustomPages**](CustomPages.md) |  | 

## Methods

### NewResponseCustomPages

`func NewResponseCustomPages(data CustomPages, ) *ResponseCustomPages`

NewResponseCustomPages instantiates a new ResponseCustomPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomPagesWithDefaults

`func NewResponseCustomPagesWithDefaults() *ResponseCustomPages`

NewResponseCustomPagesWithDefaults instantiates a new ResponseCustomPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseCustomPages) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseCustomPages) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseCustomPages) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseCustomPages) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseCustomPages) GetData() CustomPages`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCustomPages) GetDataOk() (*CustomPages, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCustomPages) SetData(v CustomPages)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


