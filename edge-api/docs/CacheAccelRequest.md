# CacheAccelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheVaryByMethod** | Pointer to **[]string** |  | [optional] 
**CacheVaryByQuerystring** | Pointer to [**CacheVaryQuerystringRequest**](CacheVaryQuerystringRequest.md) |  | [optional] 
**CacheVaryByCookies** | Pointer to [**CacheVaryCookiesRequest**](CacheVaryCookiesRequest.md) |  | [optional] 
**CacheVaryByDevices** | Pointer to [**CacheVaryDevicesRequest**](CacheVaryDevicesRequest.md) |  | [optional] 

## Methods

### NewCacheAccelRequest

`func NewCacheAccelRequest() *CacheAccelRequest`

NewCacheAccelRequest instantiates a new CacheAccelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheAccelRequestWithDefaults

`func NewCacheAccelRequestWithDefaults() *CacheAccelRequest`

NewCacheAccelRequestWithDefaults instantiates a new CacheAccelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheVaryByMethod

`func (o *CacheAccelRequest) GetCacheVaryByMethod() []string`

GetCacheVaryByMethod returns the CacheVaryByMethod field if non-nil, zero value otherwise.

### GetCacheVaryByMethodOk

`func (o *CacheAccelRequest) GetCacheVaryByMethodOk() (*[]string, bool)`

GetCacheVaryByMethodOk returns a tuple with the CacheVaryByMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheVaryByMethod

`func (o *CacheAccelRequest) SetCacheVaryByMethod(v []string)`

SetCacheVaryByMethod sets CacheVaryByMethod field to given value.

### HasCacheVaryByMethod

`func (o *CacheAccelRequest) HasCacheVaryByMethod() bool`

HasCacheVaryByMethod returns a boolean if a field has been set.

### GetCacheVaryByQuerystring

`func (o *CacheAccelRequest) GetCacheVaryByQuerystring() CacheVaryQuerystringRequest`

GetCacheVaryByQuerystring returns the CacheVaryByQuerystring field if non-nil, zero value otherwise.

### GetCacheVaryByQuerystringOk

`func (o *CacheAccelRequest) GetCacheVaryByQuerystringOk() (*CacheVaryQuerystringRequest, bool)`

GetCacheVaryByQuerystringOk returns a tuple with the CacheVaryByQuerystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheVaryByQuerystring

`func (o *CacheAccelRequest) SetCacheVaryByQuerystring(v CacheVaryQuerystringRequest)`

SetCacheVaryByQuerystring sets CacheVaryByQuerystring field to given value.

### HasCacheVaryByQuerystring

`func (o *CacheAccelRequest) HasCacheVaryByQuerystring() bool`

HasCacheVaryByQuerystring returns a boolean if a field has been set.

### GetCacheVaryByCookies

`func (o *CacheAccelRequest) GetCacheVaryByCookies() CacheVaryCookiesRequest`

GetCacheVaryByCookies returns the CacheVaryByCookies field if non-nil, zero value otherwise.

### GetCacheVaryByCookiesOk

`func (o *CacheAccelRequest) GetCacheVaryByCookiesOk() (*CacheVaryCookiesRequest, bool)`

GetCacheVaryByCookiesOk returns a tuple with the CacheVaryByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheVaryByCookies

`func (o *CacheAccelRequest) SetCacheVaryByCookies(v CacheVaryCookiesRequest)`

SetCacheVaryByCookies sets CacheVaryByCookies field to given value.

### HasCacheVaryByCookies

`func (o *CacheAccelRequest) HasCacheVaryByCookies() bool`

HasCacheVaryByCookies returns a boolean if a field has been set.

### GetCacheVaryByDevices

`func (o *CacheAccelRequest) GetCacheVaryByDevices() CacheVaryDevicesRequest`

GetCacheVaryByDevices returns the CacheVaryByDevices field if non-nil, zero value otherwise.

### GetCacheVaryByDevicesOk

`func (o *CacheAccelRequest) GetCacheVaryByDevicesOk() (*CacheVaryDevicesRequest, bool)`

GetCacheVaryByDevicesOk returns a tuple with the CacheVaryByDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheVaryByDevices

`func (o *CacheAccelRequest) SetCacheVaryByDevices(v CacheVaryDevicesRequest)`

SetCacheVaryByDevices sets CacheVaryByDevices field to given value.

### HasCacheVaryByDevices

`func (o *CacheAccelRequest) HasCacheVaryByDevices() bool`

HasCacheVaryByDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


