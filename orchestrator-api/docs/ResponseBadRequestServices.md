# ResponseBadRequestServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**MinVersion** | Pointer to **[]string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**RefCount** | Pointer to **[]string** |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestServices

`func NewResponseBadRequestServices() *ResponseBadRequestServices`

NewResponseBadRequestServices instantiates a new ResponseBadRequestServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestServicesWithDefaults

`func NewResponseBadRequestServicesWithDefaults() *ResponseBadRequestServices`

NewResponseBadRequestServicesWithDefaults instantiates a new ResponseBadRequestServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseBadRequestServices) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestServices) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestServices) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestServices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestServices) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestServices) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestServices) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestServices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestServices) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestServices) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestServices) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestServices) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMinVersion

`func (o *ResponseBadRequestServices) GetMinVersion() []string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *ResponseBadRequestServices) GetMinVersionOk() (*[]string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *ResponseBadRequestServices) SetMinVersion(v []string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *ResponseBadRequestServices) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetPermissions

`func (o *ResponseBadRequestServices) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ResponseBadRequestServices) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ResponseBadRequestServices) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ResponseBadRequestServices) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestServices) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestServices) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestServices) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestServices) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestServices) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestServices) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestServices) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestServices) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetRefCount

`func (o *ResponseBadRequestServices) GetRefCount() []string`

GetRefCount returns the RefCount field if non-nil, zero value otherwise.

### GetRefCountOk

`func (o *ResponseBadRequestServices) GetRefCountOk() (*[]string, bool)`

GetRefCountOk returns a tuple with the RefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCount

`func (o *ResponseBadRequestServices) SetRefCount(v []string)`

SetRefCount sets RefCount field to given value.

### HasRefCount

`func (o *ResponseBadRequestServices) HasRefCount() bool`

HasRefCount returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestServices) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestServices) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestServices) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestServices) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestServices) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestServices) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestServices) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestServices) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


