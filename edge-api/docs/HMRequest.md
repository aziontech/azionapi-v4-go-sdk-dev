# HMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Config** | Pointer to [**NullableAWS4HMRequest**](AWS4HMRequest.md) |  | [optional] 

## Methods

### NewHMRequest

`func NewHMRequest(enabled bool, ) *HMRequest`

NewHMRequest instantiates a new HMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHMRequestWithDefaults

`func NewHMRequestWithDefaults() *HMRequest`

NewHMRequestWithDefaults instantiates a new HMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HMRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HMRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HMRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConfig

`func (o *HMRequest) GetConfig() AWS4HMRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *HMRequest) GetConfigOk() (*AWS4HMRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *HMRequest) SetConfig(v AWS4HMRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *HMRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *HMRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *HMRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


