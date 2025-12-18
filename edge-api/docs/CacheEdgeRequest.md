# CacheEdgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | * &#x60;honor&#x60; - Honor Origin Cache Settings * &#x60;override&#x60; - Override Cache Settings | [optional] 
**MaxAge** | Pointer to **int64** | To use a value lower than 60s, the Application Acceleration module must be enabled on the Edge Application. | [optional] 
**StaleCache** | Pointer to [**StateCacheModuleRequest**](StateCacheModuleRequest.md) |  | [optional] 
**LargeFileCache** | Pointer to [**LargeFileCacheModuleRequest**](LargeFileCacheModuleRequest.md) |  | [optional] 
**TieredCache** | Pointer to [**NullableCacheSettingsTierCacheModuleRequest**](CacheSettingsTierCacheModuleRequest.md) |  | [optional] 

## Methods

### NewCacheEdgeRequest

`func NewCacheEdgeRequest() *CacheEdgeRequest`

NewCacheEdgeRequest instantiates a new CacheEdgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheEdgeRequestWithDefaults

`func NewCacheEdgeRequestWithDefaults() *CacheEdgeRequest`

NewCacheEdgeRequestWithDefaults instantiates a new CacheEdgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheEdgeRequest) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheEdgeRequest) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheEdgeRequest) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *CacheEdgeRequest) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetMaxAge

`func (o *CacheEdgeRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CacheEdgeRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CacheEdgeRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CacheEdgeRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetStaleCache

`func (o *CacheEdgeRequest) GetStaleCache() StateCacheModuleRequest`

GetStaleCache returns the StaleCache field if non-nil, zero value otherwise.

### GetStaleCacheOk

`func (o *CacheEdgeRequest) GetStaleCacheOk() (*StateCacheModuleRequest, bool)`

GetStaleCacheOk returns a tuple with the StaleCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCache

`func (o *CacheEdgeRequest) SetStaleCache(v StateCacheModuleRequest)`

SetStaleCache sets StaleCache field to given value.

### HasStaleCache

`func (o *CacheEdgeRequest) HasStaleCache() bool`

HasStaleCache returns a boolean if a field has been set.

### GetLargeFileCache

`func (o *CacheEdgeRequest) GetLargeFileCache() LargeFileCacheModuleRequest`

GetLargeFileCache returns the LargeFileCache field if non-nil, zero value otherwise.

### GetLargeFileCacheOk

`func (o *CacheEdgeRequest) GetLargeFileCacheOk() (*LargeFileCacheModuleRequest, bool)`

GetLargeFileCacheOk returns a tuple with the LargeFileCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFileCache

`func (o *CacheEdgeRequest) SetLargeFileCache(v LargeFileCacheModuleRequest)`

SetLargeFileCache sets LargeFileCache field to given value.

### HasLargeFileCache

`func (o *CacheEdgeRequest) HasLargeFileCache() bool`

HasLargeFileCache returns a boolean if a field has been set.

### GetTieredCache

`func (o *CacheEdgeRequest) GetTieredCache() CacheSettingsTierCacheModuleRequest`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *CacheEdgeRequest) GetTieredCacheOk() (*CacheSettingsTierCacheModuleRequest, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *CacheEdgeRequest) SetTieredCache(v CacheSettingsTierCacheModuleRequest)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *CacheEdgeRequest) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.

### SetTieredCacheNil

`func (o *CacheEdgeRequest) SetTieredCacheNil(b bool)`

 SetTieredCacheNil sets the value for TieredCache to be an explicit nil

### UnsetTieredCache
`func (o *CacheEdgeRequest) UnsetTieredCache()`

UnsetTieredCache ensures that no value is present for TieredCache, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


