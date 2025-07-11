# ResponseListCacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Modules** | Pointer to [**CacheSettingModules**](CacheSettingModules.md) |  | [optional] 

## Methods

### NewResponseListCacheSetting

`func NewResponseListCacheSetting(id int64, name string, ) *ResponseListCacheSetting`

NewResponseListCacheSetting instantiates a new ResponseListCacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListCacheSettingWithDefaults

`func NewResponseListCacheSettingWithDefaults() *ResponseListCacheSetting`

NewResponseListCacheSettingWithDefaults instantiates a new ResponseListCacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListCacheSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListCacheSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListCacheSetting) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListCacheSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListCacheSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListCacheSetting) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *ResponseListCacheSetting) GetModules() CacheSettingModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseListCacheSetting) GetModulesOk() (*CacheSettingModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseListCacheSetting) SetModules(v CacheSettingModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseListCacheSetting) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


