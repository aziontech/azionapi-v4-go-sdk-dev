# EdgeFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Runtime** | Pointer to [**RuntimeEnum**](RuntimeEnum.md) |  | [optional] [default to azion_js]
**ExecutionEnvironment** | Pointer to [**ExecutionEnvironmentEnum**](ExecutionEnvironmentEnum.md) |  | [optional] [default to application]
**Code** | **string** | String containing the function code. Maximum size: 20MB. | 
**DefaultArgs** | Pointer to [**EdgeFunctionsDefaultArgs**](EdgeFunctionsDefaultArgs.md) |  | [optional] [default to {}]
**AzionForm** | Pointer to [**ApplicationFunctionInstanceAzionForm**](ApplicationFunctionInstanceAzionForm.md) |  | [optional] [default to {}]

## Methods

### NewEdgeFunctionsRequest

`func NewEdgeFunctionsRequest(name string, code string, ) *EdgeFunctionsRequest`

NewEdgeFunctionsRequest instantiates a new EdgeFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionsRequestWithDefaults

`func NewEdgeFunctionsRequestWithDefaults() *EdgeFunctionsRequest`

NewEdgeFunctionsRequestWithDefaults instantiates a new EdgeFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *EdgeFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRuntime

`func (o *EdgeFunctionsRequest) GetRuntime() RuntimeEnum`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *EdgeFunctionsRequest) GetRuntimeOk() (*RuntimeEnum, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *EdgeFunctionsRequest) SetRuntime(v RuntimeEnum)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *EdgeFunctionsRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *EdgeFunctionsRequest) GetExecutionEnvironment() ExecutionEnvironmentEnum`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *EdgeFunctionsRequest) GetExecutionEnvironmentOk() (*ExecutionEnvironmentEnum, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *EdgeFunctionsRequest) SetExecutionEnvironment(v ExecutionEnvironmentEnum)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *EdgeFunctionsRequest) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetCode

`func (o *EdgeFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EdgeFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EdgeFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetDefaultArgs

`func (o *EdgeFunctionsRequest) GetDefaultArgs() EdgeFunctionsDefaultArgs`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *EdgeFunctionsRequest) GetDefaultArgsOk() (*EdgeFunctionsDefaultArgs, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *EdgeFunctionsRequest) SetDefaultArgs(v EdgeFunctionsDefaultArgs)`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *EdgeFunctionsRequest) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### GetAzionForm

`func (o *EdgeFunctionsRequest) GetAzionForm() ApplicationFunctionInstanceAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *EdgeFunctionsRequest) GetAzionFormOk() (*ApplicationFunctionInstanceAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *EdgeFunctionsRequest) SetAzionForm(v ApplicationFunctionInstanceAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *EdgeFunctionsRequest) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


