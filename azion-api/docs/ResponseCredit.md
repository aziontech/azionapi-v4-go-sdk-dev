# ResponseCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**Credit**](Credit.md) |  | 

## Methods

### NewResponseCredit

`func NewResponseCredit(data Credit, ) *ResponseCredit`

NewResponseCredit instantiates a new ResponseCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCreditWithDefaults

`func NewResponseCreditWithDefaults() *ResponseCredit`

NewResponseCreditWithDefaults instantiates a new ResponseCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseCredit) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseCredit) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseCredit) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseCredit) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseCredit) GetData() Credit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCredit) GetDataOk() (*Credit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCredit) SetData(v Credit)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


