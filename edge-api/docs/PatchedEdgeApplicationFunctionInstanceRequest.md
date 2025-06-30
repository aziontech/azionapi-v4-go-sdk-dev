# PatchedEdgeApplicationFunctionInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Args** | Pointer to [**EdgeApplicationFunctionInstanceArgs**](EdgeApplicationFunctionInstanceArgs.md) |  | [optional] 
**EdgeFunction** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEdgeApplicationFunctionInstanceRequest

`func NewPatchedEdgeApplicationFunctionInstanceRequest() *PatchedEdgeApplicationFunctionInstanceRequest`

NewPatchedEdgeApplicationFunctionInstanceRequest instantiates a new PatchedEdgeApplicationFunctionInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeApplicationFunctionInstanceRequestWithDefaults

`func NewPatchedEdgeApplicationFunctionInstanceRequestWithDefaults() *PatchedEdgeApplicationFunctionInstanceRequest`

NewPatchedEdgeApplicationFunctionInstanceRequestWithDefaults instantiates a new PatchedEdgeApplicationFunctionInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArgs

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetArgs() EdgeApplicationFunctionInstanceArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetArgsOk() (*EdgeApplicationFunctionInstanceArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) SetArgs(v EdgeApplicationFunctionInstanceArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEdgeFunction

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.

### HasEdgeFunction

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) HasEdgeFunction() bool`

HasEdgeFunction returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeApplicationFunctionInstanceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


