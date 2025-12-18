# PatchCacheSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**BrowserCache** | Pointer to [**BrowserCacheModuleRequest**](BrowserCacheModuleRequest.md) |  | [optional] 
**Modules** | Pointer to [**CacheSettingsModulesRequest**](CacheSettingsModulesRequest.md) |  | [optional] 

## Methods

### NewPatchCacheSettingRequest

`func NewPatchCacheSettingRequest() *PatchCacheSettingRequest`

NewPatchCacheSettingRequest instantiates a new PatchCacheSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCacheSettingRequestWithDefaults

`func NewPatchCacheSettingRequestWithDefaults() *PatchCacheSettingRequest`

NewPatchCacheSettingRequestWithDefaults instantiates a new PatchCacheSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchCacheSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchCacheSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchCacheSettingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchCacheSettingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBrowserCache

`func (o *PatchCacheSettingRequest) GetBrowserCache() BrowserCacheModuleRequest`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *PatchCacheSettingRequest) GetBrowserCacheOk() (*BrowserCacheModuleRequest, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *PatchCacheSettingRequest) SetBrowserCache(v BrowserCacheModuleRequest)`

SetBrowserCache sets BrowserCache field to given value.

### HasBrowserCache

`func (o *PatchCacheSettingRequest) HasBrowserCache() bool`

HasBrowserCache returns a boolean if a field has been set.

### GetModules

`func (o *PatchCacheSettingRequest) GetModules() CacheSettingsModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchCacheSettingRequest) GetModulesOk() (*CacheSettingsModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchCacheSettingRequest) SetModules(v CacheSettingsModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchCacheSettingRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


