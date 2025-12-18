# CacheSettingsTierCacheModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topology** | Pointer to **string** | * &#x60;nearest-region&#x60; - nearest-region * &#x60;br-east-1&#x60; - br-east-1 * &#x60;us-east-1&#x60; - us-east-1 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCacheSettingsTierCacheModule

`func NewCacheSettingsTierCacheModule() *CacheSettingsTierCacheModule`

NewCacheSettingsTierCacheModule instantiates a new CacheSettingsTierCacheModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsTierCacheModuleWithDefaults

`func NewCacheSettingsTierCacheModuleWithDefaults() *CacheSettingsTierCacheModule`

NewCacheSettingsTierCacheModuleWithDefaults instantiates a new CacheSettingsTierCacheModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopology

`func (o *CacheSettingsTierCacheModule) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *CacheSettingsTierCacheModule) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *CacheSettingsTierCacheModule) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *CacheSettingsTierCacheModule) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetEnabled

`func (o *CacheSettingsTierCacheModule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CacheSettingsTierCacheModule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CacheSettingsTierCacheModule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CacheSettingsTierCacheModule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


