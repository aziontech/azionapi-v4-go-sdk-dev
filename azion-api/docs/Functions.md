# Functions

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
**IsVersioned** | **bool** |  | 
**ResourceVersion** | **NullableInt64** |  | 
**VersionState** | **NullableString** |  | 
**VersionId** | **NullableString** |  | 
**Code** | **string** | String containing the function code. Maximum size: 50.0MB | 

## Methods

### NewFunctions

`func NewFunctions(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, referenceCount int64, version string, vendor string, isVersioned bool, resourceVersion NullableInt64, versionState NullableString, versionId NullableString, code string, ) *Functions`

NewFunctions instantiates a new Functions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsWithDefaults

`func NewFunctionsWithDefaults() *Functions`

NewFunctionsWithDefaults instantiates a new Functions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Functions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Functions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Functions) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Functions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Functions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Functions) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *Functions) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *Functions) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *Functions) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *Functions) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Functions) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Functions) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *Functions) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *Functions) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *Functions) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetActive

`func (o *Functions) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Functions) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Functions) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Functions) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRuntime

`func (o *Functions) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Functions) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Functions) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Functions) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *Functions) GetExecutionEnvironment() string`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *Functions) GetExecutionEnvironmentOk() (*string, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *Functions) SetExecutionEnvironment(v string)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *Functions) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetDefaultArgs

`func (o *Functions) GetDefaultArgs() interface{}`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *Functions) GetDefaultArgsOk() (*interface{}, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *Functions) SetDefaultArgs(v interface{})`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *Functions) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### SetDefaultArgsNil

`func (o *Functions) SetDefaultArgsNil(b bool)`

 SetDefaultArgsNil sets the value for DefaultArgs to be an explicit nil

### UnsetDefaultArgs
`func (o *Functions) UnsetDefaultArgs()`

UnsetDefaultArgs ensures that no value is present for DefaultArgs, not even an explicit nil
### GetAzionForm

`func (o *Functions) GetAzionForm() FunctionAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *Functions) GetAzionFormOk() (*FunctionAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *Functions) SetAzionForm(v FunctionAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *Functions) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetReferenceCount

`func (o *Functions) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *Functions) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *Functions) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.


### GetVersion

`func (o *Functions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Functions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Functions) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVendor

`func (o *Functions) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Functions) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Functions) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetIsVersioned

`func (o *Functions) GetIsVersioned() bool`

GetIsVersioned returns the IsVersioned field if non-nil, zero value otherwise.

### GetIsVersionedOk

`func (o *Functions) GetIsVersionedOk() (*bool, bool)`

GetIsVersionedOk returns a tuple with the IsVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioned

`func (o *Functions) SetIsVersioned(v bool)`

SetIsVersioned sets IsVersioned field to given value.


### GetResourceVersion

`func (o *Functions) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *Functions) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *Functions) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.


### SetResourceVersionNil

`func (o *Functions) SetResourceVersionNil(b bool)`

 SetResourceVersionNil sets the value for ResourceVersion to be an explicit nil

### UnsetResourceVersion
`func (o *Functions) UnsetResourceVersion()`

UnsetResourceVersion ensures that no value is present for ResourceVersion, not even an explicit nil
### GetVersionState

`func (o *Functions) GetVersionState() string`

GetVersionState returns the VersionState field if non-nil, zero value otherwise.

### GetVersionStateOk

`func (o *Functions) GetVersionStateOk() (*string, bool)`

GetVersionStateOk returns a tuple with the VersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionState

`func (o *Functions) SetVersionState(v string)`

SetVersionState sets VersionState field to given value.


### SetVersionStateNil

`func (o *Functions) SetVersionStateNil(b bool)`

 SetVersionStateNil sets the value for VersionState to be an explicit nil

### UnsetVersionState
`func (o *Functions) UnsetVersionState()`

UnsetVersionState ensures that no value is present for VersionState, not even an explicit nil
### GetVersionId

`func (o *Functions) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Functions) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Functions) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### SetVersionIdNil

`func (o *Functions) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *Functions) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil
### GetCode

`func (o *Functions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Functions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Functions) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


