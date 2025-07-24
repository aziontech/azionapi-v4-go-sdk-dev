# GetEdgeFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Runtime** | Pointer to **string** | * &#x60;azion_js&#x60; - Azion JavaScript * &#x60;azion_lua&#x60; - Azion Lua | [optional] 
**DefaultArgs** | Pointer to [**EdgeFunctionsDefaultArgs**](EdgeFunctionsDefaultArgs.md) |  | [optional] 
**ExecutionEnvironment** | Pointer to **string** | * &#x60;application&#x60; - application * &#x60;firewall&#x60; - firewall | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ReferenceCount** | **int64** |  | [readonly] 
**Version** | **string** | Installed version, which may not be the latest if the vendor has released updates since installation. | [readonly] 
**Vendor** | **string** |  | [readonly] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewGetEdgeFunctions

`func NewGetEdgeFunctions(id int64, name string, referenceCount int64, version string, vendor string, lastEditor string, lastModified time.Time, productVersion string, ) *GetEdgeFunctions`

NewGetEdgeFunctions instantiates a new GetEdgeFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEdgeFunctionsWithDefaults

`func NewGetEdgeFunctionsWithDefaults() *GetEdgeFunctions`

NewGetEdgeFunctionsWithDefaults instantiates a new GetEdgeFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEdgeFunctions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEdgeFunctions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEdgeFunctions) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *GetEdgeFunctions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEdgeFunctions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEdgeFunctions) SetName(v string)`

SetName sets Name field to given value.


### GetRuntime

`func (o *GetEdgeFunctions) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *GetEdgeFunctions) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *GetEdgeFunctions) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *GetEdgeFunctions) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetDefaultArgs

`func (o *GetEdgeFunctions) GetDefaultArgs() EdgeFunctionsDefaultArgs`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *GetEdgeFunctions) GetDefaultArgsOk() (*EdgeFunctionsDefaultArgs, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *GetEdgeFunctions) SetDefaultArgs(v EdgeFunctionsDefaultArgs)`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *GetEdgeFunctions) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *GetEdgeFunctions) GetExecutionEnvironment() string`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *GetEdgeFunctions) GetExecutionEnvironmentOk() (*string, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *GetEdgeFunctions) SetExecutionEnvironment(v string)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *GetEdgeFunctions) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetActive

`func (o *GetEdgeFunctions) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetEdgeFunctions) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetEdgeFunctions) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetEdgeFunctions) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReferenceCount

`func (o *GetEdgeFunctions) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *GetEdgeFunctions) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *GetEdgeFunctions) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.


### GetVersion

`func (o *GetEdgeFunctions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetEdgeFunctions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetEdgeFunctions) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVendor

`func (o *GetEdgeFunctions) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetEdgeFunctions) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetEdgeFunctions) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetLastEditor

`func (o *GetEdgeFunctions) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *GetEdgeFunctions) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *GetEdgeFunctions) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *GetEdgeFunctions) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *GetEdgeFunctions) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *GetEdgeFunctions) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *GetEdgeFunctions) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *GetEdgeFunctions) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *GetEdgeFunctions) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


