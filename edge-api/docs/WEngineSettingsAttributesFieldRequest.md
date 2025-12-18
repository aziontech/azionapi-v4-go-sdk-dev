# WEngineSettingsAttributesFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rulesets** | Pointer to **[]int64** |  | [optional] 
**Thresholds** | Pointer to [**[]ThresholdsConfigFieldRequest**](ThresholdsConfigFieldRequest.md) |  | [optional] 

## Methods

### NewWEngineSettingsAttributesFieldRequest

`func NewWEngineSettingsAttributesFieldRequest() *WEngineSettingsAttributesFieldRequest`

NewWEngineSettingsAttributesFieldRequest instantiates a new WEngineSettingsAttributesFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWEngineSettingsAttributesFieldRequestWithDefaults

`func NewWEngineSettingsAttributesFieldRequestWithDefaults() *WEngineSettingsAttributesFieldRequest`

NewWEngineSettingsAttributesFieldRequestWithDefaults instantiates a new WEngineSettingsAttributesFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesets

`func (o *WEngineSettingsAttributesFieldRequest) GetRulesets() []int64`

GetRulesets returns the Rulesets field if non-nil, zero value otherwise.

### GetRulesetsOk

`func (o *WEngineSettingsAttributesFieldRequest) GetRulesetsOk() (*[]int64, bool)`

GetRulesetsOk returns a tuple with the Rulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesets

`func (o *WEngineSettingsAttributesFieldRequest) SetRulesets(v []int64)`

SetRulesets sets Rulesets field to given value.

### HasRulesets

`func (o *WEngineSettingsAttributesFieldRequest) HasRulesets() bool`

HasRulesets returns a boolean if a field has been set.

### GetThresholds

`func (o *WEngineSettingsAttributesFieldRequest) GetThresholds() []ThresholdsConfigFieldRequest`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *WEngineSettingsAttributesFieldRequest) GetThresholdsOk() (*[]ThresholdsConfigFieldRequest, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *WEngineSettingsAttributesFieldRequest) SetThresholds(v []ThresholdsConfigFieldRequest)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *WEngineSettingsAttributesFieldRequest) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


