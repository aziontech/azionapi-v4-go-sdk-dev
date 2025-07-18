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

// checks if the ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField{}

// ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField struct for ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField
type ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField struct {
	Enabled []string `json:"enabled,omitempty"`
	Offset []string `json:"offset,omitempty"`
}

// NewResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField instantiates a new ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField() *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField {
	this := ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField{}
	return &this
}

// NewResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheFieldWithDefaults instantiates a new ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheFieldWithDefaults() *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField {
	this := ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) GetEnabled() []string {
	if o == nil || IsNil(o.Enabled) {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) GetEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) SetEnabled(v []string) {
	o.Enabled = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) GetOffset() []string {
	if o == nil || IsNil(o.Offset) {
		var ret []string
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) GetOffsetOk() ([]string, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given []string and assigns it to the Offset field.
func (o *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) SetOffset(v []string) {
	o.Offset = v
}

func (o ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField struct {
	value *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField
	isSet bool
}

func (v NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) Get() *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField {
	return v.value
}

func (v *NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) Set(val *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField(val *ResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) *NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField {
	return &NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField{value: val, isSet: true}
}

func (v NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestCacheSettingsEdgeCacheModuleLargeFileCacheField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


