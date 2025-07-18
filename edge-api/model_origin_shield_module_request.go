/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapi

import (
	"encoding/json"
)

// checks if the OriginShieldModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginShieldModuleRequest{}

// OriginShieldModuleRequest struct for OriginShieldModuleRequest
type OriginShieldModuleRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	Config NullableOriginShieldConfigRequest `json:"config,omitempty"`
}

// NewOriginShieldModuleRequest instantiates a new OriginShieldModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginShieldModuleRequest() *OriginShieldModuleRequest {
	this := OriginShieldModuleRequest{}
	return &this
}

// NewOriginShieldModuleRequestWithDefaults instantiates a new OriginShieldModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginShieldModuleRequestWithDefaults() *OriginShieldModuleRequest {
	this := OriginShieldModuleRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OriginShieldModuleRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldModuleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OriginShieldModuleRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OriginShieldModuleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginShieldModuleRequest) GetConfig() OriginShieldConfigRequest {
	if o == nil || IsNil(o.Config.Get()) {
		var ret OriginShieldConfigRequest
		return ret
	}
	return *o.Config.Get()
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginShieldModuleRequest) GetConfigOk() (*OriginShieldConfigRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config.Get(), o.Config.IsSet()
}

// HasConfig returns a boolean if a field has been set.
func (o *OriginShieldModuleRequest) HasConfig() bool {
	if o != nil && o.Config.IsSet() {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NullableOriginShieldConfigRequest and assigns it to the Config field.
func (o *OriginShieldModuleRequest) SetConfig(v OriginShieldConfigRequest) {
	o.Config.Set(&v)
}
// SetConfigNil sets the value for Config to be an explicit nil
func (o *OriginShieldModuleRequest) SetConfigNil() {
	o.Config.Set(nil)
}

// UnsetConfig ensures that no value is present for Config, not even an explicit nil
func (o *OriginShieldModuleRequest) UnsetConfig() {
	o.Config.Unset()
}

func (o OriginShieldModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginShieldModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Config.IsSet() {
		toSerialize["config"] = o.Config.Get()
	}
	return toSerialize, nil
}

type NullableOriginShieldModuleRequest struct {
	value *OriginShieldModuleRequest
	isSet bool
}

func (v NullableOriginShieldModuleRequest) Get() *OriginShieldModuleRequest {
	return v.value
}

func (v *NullableOriginShieldModuleRequest) Set(val *OriginShieldModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginShieldModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginShieldModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginShieldModuleRequest(val *OriginShieldModuleRequest) *NullableOriginShieldModuleRequest {
	return &NullableOriginShieldModuleRequest{value: val, isSet: true}
}

func (v NullableOriginShieldModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginShieldModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


