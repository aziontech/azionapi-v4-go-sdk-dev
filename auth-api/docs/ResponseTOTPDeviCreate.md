# ResponseTOTPDeviCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**TOTPDeviCreate**](TOTPDeviCreate.md) |  | 

## Methods

### NewResponseTOTPDeviCreate

`func NewResponseTOTPDeviCreate(data TOTPDeviCreate, ) *ResponseTOTPDeviCreate`

NewResponseTOTPDeviCreate instantiates a new ResponseTOTPDeviCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTOTPDeviCreateWithDefaults

`func NewResponseTOTPDeviCreateWithDefaults() *ResponseTOTPDeviCreate`

NewResponseTOTPDeviCreateWithDefaults instantiates a new ResponseTOTPDeviCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseTOTPDeviCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseTOTPDeviCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseTOTPDeviCreate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseTOTPDeviCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseTOTPDeviCreate) GetData() TOTPDeviCreate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseTOTPDeviCreate) GetDataOk() (*TOTPDeviCreate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseTOTPDeviCreate) SetData(v TOTPDeviCreate)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


