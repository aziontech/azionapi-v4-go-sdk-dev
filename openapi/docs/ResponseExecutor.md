# ResponseExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**Executor**](Executor.md) |  | 

## Methods

### NewResponseExecutor

`func NewResponseExecutor(data Executor, ) *ResponseExecutor`

NewResponseExecutor instantiates a new ResponseExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseExecutorWithDefaults

`func NewResponseExecutorWithDefaults() *ResponseExecutor`

NewResponseExecutorWithDefaults instantiates a new ResponseExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseExecutor) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseExecutor) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseExecutor) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseExecutor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseExecutor) GetData() Executor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseExecutor) GetDataOk() (*Executor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseExecutor) SetData(v Executor)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


