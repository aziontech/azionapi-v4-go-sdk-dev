# CacheSettingModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserCache** | [**BrowserCacheModuleRequest**](BrowserCacheModuleRequest.md) |  | 
**EdgeCache** | [**EdgeCacheModuleRequest**](EdgeCacheModuleRequest.md) |  | 
**ApplicationControls** | [**ApplicationControlsModuleRequest**](ApplicationControlsModuleRequest.md) |  | 
**SliceControls** | [**SliceControlsModuleRequest**](SliceControlsModuleRequest.md) |  | 

## Methods

### NewCacheSettingModulesRequest

`func NewCacheSettingModulesRequest(browserCache BrowserCacheModuleRequest, edgeCache EdgeCacheModuleRequest, applicationControls ApplicationControlsModuleRequest, sliceControls SliceControlsModuleRequest, ) *CacheSettingModulesRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


