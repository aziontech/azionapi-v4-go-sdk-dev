# CacheSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Modules** | Pointer to [**CacheSettingModulesRequest**](CacheSettingModulesRequest.md) |  | [optional] 

## Methods

### NewCacheSettingRequest

`func NewCacheSettingRequest(name string, ) *CacheSettingRequest`

NewCacheSettingRequest instantiates a new CacheSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingRequestWithDefaults

`func NewCacheSettingRequestWithDefaults() *CacheSettingRequest`

NewCacheSettingRequestWithDefaults instantiates a new CacheSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CacheSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSettingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *CacheSettingRequest) GetModules() CacheSettingModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *CacheSettingRequest) GetModulesOk() (*CacheSettingModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *CacheSettingRequest) SetModules(v CacheSettingModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *CacheSettingRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


