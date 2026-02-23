# ResponseAsyncPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**PaymentMethod**](PaymentMethod.md) |  | 

## Methods

### NewResponseAsyncPaymentMethod

`func NewResponseAsyncPaymentMethod(data PaymentMethod, ) *ResponseAsyncPaymentMethod`

NewResponseAsyncPaymentMethod instantiates a new ResponseAsyncPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncPaymentMethodWithDefaults

`func NewResponseAsyncPaymentMethodWithDefaults() *ResponseAsyncPaymentMethod`

NewResponseAsyncPaymentMethodWithDefaults instantiates a new ResponseAsyncPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncPaymentMethod) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncPaymentMethod) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncPaymentMethod) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncPaymentMethod) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncPaymentMethod) GetData() PaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncPaymentMethod) GetDataOk() (*PaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncPaymentMethod) SetData(v PaymentMethod)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


