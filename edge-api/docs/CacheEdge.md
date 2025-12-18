# CacheEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | * &#x60;honor&#x60; - Honor Origin Cache Settings * &#x60;override&#x60; - Override Cache Settings | [optional] 
**MaxAge** | Pointer to **int64** | To use a value lower than 60s, the Application Acceleration module must be enabled on the Edge Application. | [optional] 
**StaleCache** | Pointer to [**StateCacheModule**](StateCacheModule.md) |  | [optional] 
**LargeFileCache** | Pointer to [**LargeFileCacheModule**](LargeFileCacheModule.md) |  | [optional] 
**TieredCache** | Pointer to [**NullableCacheSettingsTierCacheModule**](CacheSettingsTierCacheModule.md) |  | [optional] 

## Methods

### NewCacheEdge

`func NewCacheEdge() *CacheEdge`

NewCacheEdge instantiates a new CacheEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheEdgeWithDefaults

`func NewCacheEdgeWithDefaults() *CacheEdge`

NewCacheEdgeWithDefaults instantiates a new CacheEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheEdge) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheEdge) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheEdge) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *CacheEdge) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetMaxAge

`func (o *CacheEdge) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CacheEdge) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CacheEdge) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CacheEdge) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetStaleCache

`func (o *CacheEdge) GetStaleCache() StateCacheModule`

GetStaleCache returns the StaleCache field if non-nil, zero value otherwise.

### GetStaleCacheOk

`func (o *CacheEdge) GetStaleCacheOk() (*StateCacheModule, bool)`

GetStaleCacheOk returns a tuple with the StaleCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCache

`func (o *CacheEdge) SetStaleCache(v StateCacheModule)`

SetStaleCache sets StaleCache field to given value.

### HasStaleCache

`func (o *CacheEdge) HasStaleCache() bool`

HasStaleCache returns a boolean if a field has been set.

### GetLargeFileCache

`func (o *CacheEdge) GetLargeFileCache() LargeFileCacheModule`

GetLargeFileCache returns the LargeFileCache field if non-nil, zero value otherwise.

### GetLargeFileCacheOk

`func (o *CacheEdge) GetLargeFileCacheOk() (*LargeFileCacheModule, bool)`

GetLargeFileCacheOk returns a tuple with the LargeFileCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFileCache

`func (o *CacheEdge) SetLargeFileCache(v LargeFileCacheModule)`

SetLargeFileCache sets LargeFileCache field to given value.

### HasLargeFileCache

`func (o *CacheEdge) HasLargeFileCache() bool`

HasLargeFileCache returns a boolean if a field has been set.

### GetTieredCache

`func (o *CacheEdge) GetTieredCache() CacheSettingsTierCacheModule`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *CacheEdge) GetTieredCacheOk() (*CacheSettingsTierCacheModule, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *CacheEdge) SetTieredCache(v CacheSettingsTierCacheModule)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *CacheEdge) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.

### SetTieredCacheNil

`func (o *CacheEdge) SetTieredCacheNil(b bool)`

 SetTieredCacheNil sets the value for TieredCache to be an explicit nil

### UnsetTieredCache
`func (o *CacheEdge) UnsetTieredCache()`

UnsetTieredCache ensures that no value is present for TieredCache, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


