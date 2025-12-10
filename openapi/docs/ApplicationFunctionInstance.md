# ApplicationFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Args** | Pointer to [**ApplicationFunctionInstanceArgs**](ApplicationFunctionInstanceArgs.md) |  | [optional] 
**AzionForm** | Pointer to [**ApplicationFunctionInstanceAzionForm**](ApplicationFunctionInstanceAzionForm.md) |  | [optional] [default to {}]
**Function** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 

## Methods

### NewApplicationFunctionInstance

`func NewApplicationFunctionInstance(id int32, name string, function int64, lastEditor string, lastModified time.Time, ) *ApplicationFunctionInstance`

NewApplicationFunctionInstance instantiates a new ApplicationFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFunctionInstanceWithDefaults

`func NewApplicationFunctionInstanceWithDefaults() *ApplicationFunctionInstance`

NewApplicationFunctionInstanceWithDefaults instantiates a new ApplicationFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationFunctionInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationFunctionInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationFunctionInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationFunctionInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationFunctionInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationFunctionInstance) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *ApplicationFunctionInstance) GetArgs() ApplicationFunctionInstanceArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationFunctionInstance) GetArgsOk() (*ApplicationFunctionInstanceArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationFunctionInstance) SetArgs(v ApplicationFunctionInstanceArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ApplicationFunctionInstance) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAzionForm

`func (o *ApplicationFunctionInstance) GetAzionForm() ApplicationFunctionInstanceAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *ApplicationFunctionInstance) GetAzionFormOk() (*ApplicationFunctionInstanceAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *ApplicationFunctionInstance) SetAzionForm(v ApplicationFunctionInstanceAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *ApplicationFunctionInstance) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetFunction

`func (o *ApplicationFunctionInstance) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ApplicationFunctionInstance) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ApplicationFunctionInstance) SetFunction(v int64)`

SetFunction sets Function field to given value.


### GetActive

`func (o *ApplicationFunctionInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationFunctionInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationFunctionInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationFunctionInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *ApplicationFunctionInstance) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ApplicationFunctionInstance) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ApplicationFunctionInstance) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ApplicationFunctionInstance) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApplicationFunctionInstance) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApplicationFunctionInstance) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


