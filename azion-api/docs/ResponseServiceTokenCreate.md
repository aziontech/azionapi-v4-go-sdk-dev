# ResponseServiceTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**ServiceTokenCreate**](ServiceTokenCreate.md) |  | 

## Methods

### NewResponseServiceTokenCreate

`func NewResponseServiceTokenCreate(state string, data ServiceTokenCreate, ) *ResponseServiceTokenCreate`

NewResponseServiceTokenCreate instantiates a new ResponseServiceTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseServiceTokenCreateWithDefaults

`func NewResponseServiceTokenCreateWithDefaults() *ResponseServiceTokenCreate`

NewResponseServiceTokenCreateWithDefaults instantiates a new ResponseServiceTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseServiceTokenCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseServiceTokenCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseServiceTokenCreate) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseServiceTokenCreate) GetData() ServiceTokenCreate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseServiceTokenCreate) GetDataOk() (*ServiceTokenCreate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseServiceTokenCreate) SetData(v ServiceTokenCreate)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


