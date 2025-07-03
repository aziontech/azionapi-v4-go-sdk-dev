# ResponseAsyncEdgeFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**EdgeFunctions**](EdgeFunctions.md) |  | 

## Methods

### NewResponseAsyncEdgeFunctions

`func NewResponseAsyncEdgeFunctions(data EdgeFunctions, ) *ResponseAsyncEdgeFunctions`

NewResponseAsyncEdgeFunctions instantiates a new ResponseAsyncEdgeFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncEdgeFunctionsWithDefaults

`func NewResponseAsyncEdgeFunctionsWithDefaults() *ResponseAsyncEdgeFunctions`

NewResponseAsyncEdgeFunctionsWithDefaults instantiates a new ResponseAsyncEdgeFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncEdgeFunctions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncEdgeFunctions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncEdgeFunctions) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncEdgeFunctions) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncEdgeFunctions) GetData() EdgeFunctions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncEdgeFunctions) GetDataOk() (*EdgeFunctions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncEdgeFunctions) SetData(v EdgeFunctions)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


