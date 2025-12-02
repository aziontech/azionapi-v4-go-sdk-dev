# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;error&#x60; - error | 
**Error** | [**ErrorDetail**](ErrorDetail.md) |  | 

## Methods

### NewErrorResponse

`func NewErrorResponse(state string, error_ ErrorDetail, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ErrorResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ErrorResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ErrorResponse) SetState(v string)`

SetState sets State field to given value.


### GetError

`func (o *ErrorResponse) GetError() ErrorDetail`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponse) GetErrorOk() (*ErrorDetail, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponse) SetError(v ErrorDetail)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


