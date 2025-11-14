# ExecutionScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TemplateUuid** | **string** |  | 

## Methods

### NewExecutionScriptRequest

`func NewExecutionScriptRequest(name string, templateUuid string, ) *ExecutionScriptRequest`

NewExecutionScriptRequest instantiates a new ExecutionScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionScriptRequestWithDefaults

`func NewExecutionScriptRequestWithDefaults() *ExecutionScriptRequest`

NewExecutionScriptRequestWithDefaults instantiates a new ExecutionScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExecutionScriptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionScriptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionScriptRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTemplateUuid

`func (o *ExecutionScriptRequest) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *ExecutionScriptRequest) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *ExecutionScriptRequest) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


