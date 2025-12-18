# LoBalancerModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**NullableLoBalancerModuleConfig**](LoBalancerModuleConfig.md) |  | [optional] 

## Methods

### NewLoBalancerModule

`func NewLoBalancerModule() *LoBalancerModule`

NewLoBalancerModule instantiates a new LoBalancerModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoBalancerModuleWithDefaults

`func NewLoBalancerModuleWithDefaults() *LoBalancerModule`

NewLoBalancerModuleWithDefaults instantiates a new LoBalancerModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LoBalancerModule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoBalancerModule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoBalancerModule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LoBalancerModule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *LoBalancerModule) GetConfig() LoBalancerModuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoBalancerModule) GetConfigOk() (*LoBalancerModuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoBalancerModule) SetConfig(v LoBalancerModuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoBalancerModule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *LoBalancerModule) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *LoBalancerModule) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


