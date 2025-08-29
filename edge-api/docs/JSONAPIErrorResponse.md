# JSONAPIErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]JSONAPIErrorObject**](JSONAPIErrorObject.md) | Array of error objects | 

## Methods

### NewJSONAPIErrorResponse

`func NewJSONAPIErrorResponse(errors []JSONAPIErrorObject, ) *JSONAPIErrorResponse`

NewJSONAPIErrorResponse instantiates a new JSONAPIErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONAPIErrorResponseWithDefaults

`func NewJSONAPIErrorResponseWithDefaults() *JSONAPIErrorResponse`

NewJSONAPIErrorResponseWithDefaults instantiates a new JSONAPIErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *JSONAPIErrorResponse) GetErrors() []JSONAPIErrorObject`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *JSONAPIErrorResponse) GetErrorsOk() (*[]JSONAPIErrorObject, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *JSONAPIErrorResponse) SetErrors(v []JSONAPIErrorObject)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


