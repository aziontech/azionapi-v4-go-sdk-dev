# AppFunctionInstanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Args** | Pointer to **interface{}** |  | [optional] 
**AzionForm** | Pointer to [**EdgeFunctionsAzionForm**](EdgeFunctionsAzionForm.md) |  | [optional] 
**Function** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppFunctionInstanRequest

`func NewAppFunctionInstanRequest(name string, function int64, ) *AppFunctionInstanRequest`

NewAppFunctionInstanRequest instantiates a new AppFunctionInstanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppFunctionInstanRequestWithDefaults

`func NewAppFunctionInstanRequestWithDefaults() *AppFunctionInstanRequest`

NewAppFunctionInstanRequestWithDefaults instantiates a new AppFunctionInstanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppFunctionInstanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppFunctionInstanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppFunctionInstanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *AppFunctionInstanRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppFunctionInstanRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppFunctionInstanRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppFunctionInstanRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *AppFunctionInstanRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *AppFunctionInstanRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetAzionForm

`func (o *AppFunctionInstanRequest) GetAzionForm() EdgeFunctionsAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *AppFunctionInstanRequest) GetAzionFormOk() (*EdgeFunctionsAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *AppFunctionInstanRequest) SetAzionForm(v EdgeFunctionsAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *AppFunctionInstanRequest) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetFunction

`func (o *AppFunctionInstanRequest) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AppFunctionInstanRequest) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AppFunctionInstanRequest) SetFunction(v int64)`

SetFunction sets Function field to given value.


### GetActive

`func (o *AppFunctionInstanRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AppFunctionInstanRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AppFunctionInstanRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AppFunctionInstanRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


