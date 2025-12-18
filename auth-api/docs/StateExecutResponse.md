# StateExecutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State of the operation (lowercase with underscores) | [optional] 

## Methods

### NewStateExecutResponse

`func NewStateExecutResponse() *StateExecutResponse`

NewStateExecutResponse instantiates a new StateExecutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateExecutResponseWithDefaults

`func NewStateExecutResponseWithDefaults() *StateExecutResponse`

NewStateExecutResponseWithDefaults instantiates a new StateExecutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StateExecutResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StateExecutResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StateExecutResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StateExecutResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


