# CacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Modules** | Pointer to [**CacheSettingModules**](CacheSettingModules.md) |  | [optional] 

## Methods

### NewCacheSetting

`func NewCacheSetting(id int64, name string, ) *CacheSetting`

NewCacheSetting instantiates a new CacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingWithDefaults

`func NewCacheSettingWithDefaults() *CacheSetting`

NewCacheSettingWithDefaults instantiates a new CacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CacheSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CacheSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CacheSetting) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CacheSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSetting) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *CacheSetting) GetModules() CacheSettingModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *CacheSetting) GetModulesOk() (*CacheSettingModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *CacheSetting) SetModules(v CacheSettingModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *CacheSetting) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


