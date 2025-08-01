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

// checks if the ImageProcessorModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageProcessorModule{}

// ImageProcessorModule struct for ImageProcessorModule
type ImageProcessorModule struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewImageProcessorModule instantiates a new ImageProcessorModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageProcessorModule() *ImageProcessorModule {
	this := ImageProcessorModule{}
	return &this
}

// NewImageProcessorModuleWithDefaults instantiates a new ImageProcessorModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageProcessorModuleWithDefaults() *ImageProcessorModule {
	this := ImageProcessorModule{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ImageProcessorModule) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageProcessorModule) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ImageProcessorModule) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ImageProcessorModule) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ImageProcessorModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageProcessorModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableImageProcessorModule struct {
	value *ImageProcessorModule
	isSet bool
}

func (v NullableImageProcessorModule) Get() *ImageProcessorModule {
	return v.value
}

func (v *NullableImageProcessorModule) Set(val *ImageProcessorModule) {
	v.value = val
	v.isSet = true
}

func (v NullableImageProcessorModule) IsSet() bool {
	return v.isSet
}

func (v *NullableImageProcessorModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageProcessorModule(val *ImageProcessorModule) *NullableImageProcessorModule {
	return &NullableImageProcessorModule{value: val, isSet: true}
}

func (v NullableImageProcessorModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageProcessorModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


