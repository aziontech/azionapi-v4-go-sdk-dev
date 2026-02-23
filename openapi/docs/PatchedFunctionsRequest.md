# PatchedFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Runtime** | Pointer to **string** | * &#x60;azion_js&#x60; - Azion JavaScript | [optional] 
**ExecutionEnvironment** | Pointer to **string** | * &#x60;firewall&#x60; - Firewall * &#x60;application&#x60; - Application | [optional] 
**Code** | Pointer to **string** | String containing the function code. Maximum size: 20MB. | [optional] 
**DefaultArgs** | Pointer to **interface{}** |  | [optional] 
**AzionForm** | Pointer to [**FunctionsAzionForm**](FunctionsAzionForm.md) |  | [optional] 

## Methods

### NewPatchedFunctionsRequest

`func NewPatchedFunctionsRequest() *PatchedFunctionsRequest`

NewPatchedFunctionsRequest instantiates a new PatchedFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFunctionsRequestWithDefaults

`func NewPatchedFunctionsRequestWithDefaults() *PatchedFunctionsRequest`

NewPatchedFunctionsRequestWithDefaults instantiates a new PatchedFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFunctionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRuntime

`func (o *PatchedFunctionsRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *PatchedFunctionsRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *PatchedFunctionsRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *PatchedFunctionsRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *PatchedFunctionsRequest) GetExecutionEnvironment() string`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *PatchedFunctionsRequest) GetExecutionEnvironmentOk() (*string, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *PatchedFunctionsRequest) SetExecutionEnvironment(v string)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *PatchedFunctionsRequest) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetCode

`func (o *PatchedFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchedFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchedFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PatchedFunctionsRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefaultArgs

`func (o *PatchedFunctionsRequest) GetDefaultArgs() interface{}`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *PatchedFunctionsRequest) GetDefaultArgsOk() (*interface{}, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *PatchedFunctionsRequest) SetDefaultArgs(v interface{})`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *PatchedFunctionsRequest) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### SetDefaultArgsNil

`func (o *PatchedFunctionsRequest) SetDefaultArgsNil(b bool)`

 SetDefaultArgsNil sets the value for DefaultArgs to be an explicit nil

### UnsetDefaultArgs
`func (o *PatchedFunctionsRequest) UnsetDefaultArgs()`

UnsetDefaultArgs ensures that no value is present for DefaultArgs, not even an explicit nil
### GetAzionForm

`func (o *PatchedFunctionsRequest) GetAzionForm() FunctionsAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *PatchedFunctionsRequest) GetAzionFormOk() (*FunctionsAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *PatchedFunctionsRequest) SetAzionForm(v FunctionsAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *PatchedFunctionsRequest) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


