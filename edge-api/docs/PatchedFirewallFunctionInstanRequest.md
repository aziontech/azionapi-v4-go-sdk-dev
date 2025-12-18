# PatchedFirewallFunctionInstanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **interface{}** |  | [optional] 
**AzionForm** | Pointer to [**EdgeFunctionsAzionForm**](EdgeFunctionsAzionForm.md) |  | [optional] 
**Function** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedFirewallFunctionInstanRequest

`func NewPatchedFirewallFunctionInstanRequest() *PatchedFirewallFunctionInstanRequest`

NewPatchedFirewallFunctionInstanRequest instantiates a new PatchedFirewallFunctionInstanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFirewallFunctionInstanRequestWithDefaults

`func NewPatchedFirewallFunctionInstanRequestWithDefaults() *PatchedFirewallFunctionInstanRequest`

NewPatchedFirewallFunctionInstanRequestWithDefaults instantiates a new PatchedFirewallFunctionInstanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedFirewallFunctionInstanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFirewallFunctionInstanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFirewallFunctionInstanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFirewallFunctionInstanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArgs

`func (o *PatchedFirewallFunctionInstanRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PatchedFirewallFunctionInstanRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PatchedFirewallFunctionInstanRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *PatchedFirewallFunctionInstanRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *PatchedFirewallFunctionInstanRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *PatchedFirewallFunctionInstanRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetAzionForm

`func (o *PatchedFirewallFunctionInstanRequest) GetAzionForm() EdgeFunctionsAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *PatchedFirewallFunctionInstanRequest) GetAzionFormOk() (*EdgeFunctionsAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *PatchedFirewallFunctionInstanRequest) SetAzionForm(v EdgeFunctionsAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *PatchedFirewallFunctionInstanRequest) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetFunction

`func (o *PatchedFirewallFunctionInstanRequest) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *PatchedFirewallFunctionInstanRequest) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *PatchedFirewallFunctionInstanRequest) SetFunction(v int64)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *PatchedFirewallFunctionInstanRequest) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetActive

`func (o *PatchedFirewallFunctionInstanRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedFirewallFunctionInstanRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedFirewallFunctionInstanRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedFirewallFunctionInstanRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


