# CacheSettingModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserCache** | Pointer to [**BrowserCacheModuleRequest**](BrowserCacheModuleRequest.md) |  | [optional] 
**EdgeCache** | Pointer to [**EdgeCacheModuleRequest**](EdgeCacheModuleRequest.md) |  | [optional] 
**ApplicationControls** | Pointer to [**ApplicationControlsModuleRequest**](ApplicationControlsModuleRequest.md) |  | [optional] 
**SliceControls** | Pointer to [**SliceControlsModuleRequest**](SliceControlsModuleRequest.md) |  | [optional] 

## Methods

### NewCacheSettingModulesRequest

`func NewCacheSettingModulesRequest() *CacheSettingModulesRequest`

NewCacheSettingModulesRequest instantiates a new CacheSettingModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingModulesRequestWithDefaults

`func NewCacheSettingModulesRequestWithDefaults() *CacheSettingModulesRequest`

NewCacheSettingModulesRequestWithDefaults instantiates a new CacheSettingModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserCache

`func (o *CacheSettingModulesRequest) GetBrowserCache() BrowserCacheModuleRequest`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *CacheSettingModulesRequest) GetBrowserCacheOk() (*BrowserCacheModuleRequest, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *CacheSettingModulesRequest) SetBrowserCache(v BrowserCacheModuleRequest)`

SetBrowserCache sets BrowserCache field to given value.

### HasBrowserCache

`func (o *CacheSettingModulesRequest) HasBrowserCache() bool`

HasBrowserCache returns a boolean if a field has been set.

### GetEdgeCache

`func (o *CacheSettingModulesRequest) GetEdgeCache() EdgeCacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *CacheSettingModulesRequest) GetEdgeCacheOk() (*EdgeCacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *CacheSettingModulesRequest) SetEdgeCache(v EdgeCacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *CacheSettingModulesRequest) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetApplicationControls

`func (o *CacheSettingModulesRequest) GetApplicationControls() ApplicationControlsModuleRequest`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *CacheSettingModulesRequest) GetApplicationControlsOk() (*ApplicationControlsModuleRequest, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *CacheSettingModulesRequest) SetApplicationControls(v ApplicationControlsModuleRequest)`

SetApplicationControls sets ApplicationControls field to given value.

### HasApplicationControls

`func (o *CacheSettingModulesRequest) HasApplicationControls() bool`

HasApplicationControls returns a boolean if a field has been set.

### GetSliceControls

`func (o *CacheSettingModulesRequest) GetSliceControls() SliceControlsModuleRequest`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *CacheSettingModulesRequest) GetSliceControlsOk() (*SliceControlsModuleRequest, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *CacheSettingModulesRequest) SetSliceControls(v SliceControlsModuleRequest)`

SetSliceControls sets SliceControls field to given value.

### HasSliceControls

`func (o *CacheSettingModulesRequest) HasSliceControls() bool`

HasSliceControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


