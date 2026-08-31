# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**ProductVersion** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Runtime** | Pointer to **string** | * &#x60;azion_js&#x60; - Azion JavaScript | [optional] 
**ExecutionEnvironment** | Pointer to **string** | * &#x60;firewall&#x60; - Firewall * &#x60;application&#x60; - Application | [optional] 
**DefaultArgs** | Pointer to **interface{}** |  | [optional] 
**AzionForm** | Pointer to [**FunctionAzionForm**](FunctionAzionForm.md) |  | [optional] 
**ReferenceCount** | **int64** |  | 
**Version** | **string** | Installed version, which may not be the latest if the vendor has released updates since installation. | 
**Vendor** | **string** |  | 

## Methods

### NewFunction

`func NewFunction(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, referenceCount int64, version string, vendor string, ) *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Function) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Function) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Function) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Function) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Function) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Function) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *Function) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *Function) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *Function) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *Function) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Function) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Function) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *Function) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *Function) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *Function) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetActive

`func (o *Function) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Function) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Function) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Function) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRuntime

`func (o *Function) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Function) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Function) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Function) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *Function) GetExecutionEnvironment() string`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *Function) GetExecutionEnvironmentOk() (*string, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *Function) SetExecutionEnvironment(v string)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *Function) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetDefaultArgs

`func (o *Function) GetDefaultArgs() interface{}`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *Function) GetDefaultArgsOk() (*interface{}, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *Function) SetDefaultArgs(v interface{})`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *Function) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### SetDefaultArgsNil

`func (o *Function) SetDefaultArgsNil(b bool)`

 SetDefaultArgsNil sets the value for DefaultArgs to be an explicit nil

### UnsetDefaultArgs
`func (o *Function) UnsetDefaultArgs()`

UnsetDefaultArgs ensures that no value is present for DefaultArgs, not even an explicit nil
### GetAzionForm

`func (o *Function) GetAzionForm() FunctionAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *Function) GetAzionFormOk() (*FunctionAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *Function) SetAzionForm(v FunctionAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *Function) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetReferenceCount

`func (o *Function) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *Function) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *Function) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.


### GetVersion

`func (o *Function) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Function) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Function) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVendor

`func (o *Function) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Function) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Function) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


