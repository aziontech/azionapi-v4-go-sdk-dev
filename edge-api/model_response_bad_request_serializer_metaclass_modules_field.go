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

// checks if the ResponseBadRequestSerializerMetaclassModulesField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestSerializerMetaclassModulesField{}

// ResponseBadRequestSerializerMetaclassModulesField struct for ResponseBadRequestSerializerMetaclassModulesField
type ResponseBadRequestSerializerMetaclassModulesField struct {
	EdgeCache *ResponseBadRequestCacheSettingsModulesEdgeCacheField `json:"edge_cache,omitempty"`
	TieredCache *ResponseBadRequestCacheSettingsModulesTieredCacheField `json:"tiered_cache,omitempty"`
	ApplicationAccelerator *ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField `json:"application_accelerator,omitempty"`
}

// NewResponseBadRequestSerializerMetaclassModulesField instantiates a new ResponseBadRequestSerializerMetaclassModulesField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestSerializerMetaclassModulesField() *ResponseBadRequestSerializerMetaclassModulesField {
	this := ResponseBadRequestSerializerMetaclassModulesField{}
	return &this
}

// NewResponseBadRequestSerializerMetaclassModulesFieldWithDefaults instantiates a new ResponseBadRequestSerializerMetaclassModulesField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestSerializerMetaclassModulesFieldWithDefaults() *ResponseBadRequestSerializerMetaclassModulesField {
	this := ResponseBadRequestSerializerMetaclassModulesField{}
	return &this
}

// GetEdgeCache returns the EdgeCache field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetEdgeCache() ResponseBadRequestCacheSettingsModulesEdgeCacheField {
	if o == nil || IsNil(o.EdgeCache) {
		var ret ResponseBadRequestCacheSettingsModulesEdgeCacheField
		return ret
	}
	return *o.EdgeCache
}

// GetEdgeCacheOk returns a tuple with the EdgeCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetEdgeCacheOk() (*ResponseBadRequestCacheSettingsModulesEdgeCacheField, bool) {
	if o == nil || IsNil(o.EdgeCache) {
		return nil, false
	}
	return o.EdgeCache, true
}

// HasEdgeCache returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) HasEdgeCache() bool {
	if o != nil && !IsNil(o.EdgeCache) {
		return true
	}

	return false
}

// SetEdgeCache gets a reference to the given ResponseBadRequestCacheSettingsModulesEdgeCacheField and assigns it to the EdgeCache field.
func (o *ResponseBadRequestSerializerMetaclassModulesField) SetEdgeCache(v ResponseBadRequestCacheSettingsModulesEdgeCacheField) {
	o.EdgeCache = &v
}

// GetTieredCache returns the TieredCache field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetTieredCache() ResponseBadRequestCacheSettingsModulesTieredCacheField {
	if o == nil || IsNil(o.TieredCache) {
		var ret ResponseBadRequestCacheSettingsModulesTieredCacheField
		return ret
	}
	return *o.TieredCache
}

// GetTieredCacheOk returns a tuple with the TieredCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetTieredCacheOk() (*ResponseBadRequestCacheSettingsModulesTieredCacheField, bool) {
	if o == nil || IsNil(o.TieredCache) {
		return nil, false
	}
	return o.TieredCache, true
}

// HasTieredCache returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) HasTieredCache() bool {
	if o != nil && !IsNil(o.TieredCache) {
		return true
	}

	return false
}

// SetTieredCache gets a reference to the given ResponseBadRequestCacheSettingsModulesTieredCacheField and assigns it to the TieredCache field.
func (o *ResponseBadRequestSerializerMetaclassModulesField) SetTieredCache(v ResponseBadRequestCacheSettingsModulesTieredCacheField) {
	o.TieredCache = &v
}

// GetApplicationAccelerator returns the ApplicationAccelerator field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetApplicationAccelerator() ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField {
	if o == nil || IsNil(o.ApplicationAccelerator) {
		var ret ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField
		return ret
	}
	return *o.ApplicationAccelerator
}

// GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) GetApplicationAcceleratorOk() (*ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField, bool) {
	if o == nil || IsNil(o.ApplicationAccelerator) {
		return nil, false
	}
	return o.ApplicationAccelerator, true
}

// HasApplicationAccelerator returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassModulesField) HasApplicationAccelerator() bool {
	if o != nil && !IsNil(o.ApplicationAccelerator) {
		return true
	}

	return false
}

// SetApplicationAccelerator gets a reference to the given ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField and assigns it to the ApplicationAccelerator field.
func (o *ResponseBadRequestSerializerMetaclassModulesField) SetApplicationAccelerator(v ResponseBadRequestCacheSettingsModulesApplicationAcceleratorField) {
	o.ApplicationAccelerator = &v
}

func (o ResponseBadRequestSerializerMetaclassModulesField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestSerializerMetaclassModulesField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EdgeCache) {
		toSerialize["edge_cache"] = o.EdgeCache
	}
	if !IsNil(o.TieredCache) {
		toSerialize["tiered_cache"] = o.TieredCache
	}
	if !IsNil(o.ApplicationAccelerator) {
		toSerialize["application_accelerator"] = o.ApplicationAccelerator
	}
	return toSerialize, nil
}

type NullableResponseBadRequestSerializerMetaclassModulesField struct {
	value *ResponseBadRequestSerializerMetaclassModulesField
	isSet bool
}

func (v NullableResponseBadRequestSerializerMetaclassModulesField) Get() *ResponseBadRequestSerializerMetaclassModulesField {
	return v.value
}

func (v *NullableResponseBadRequestSerializerMetaclassModulesField) Set(val *ResponseBadRequestSerializerMetaclassModulesField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestSerializerMetaclassModulesField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestSerializerMetaclassModulesField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestSerializerMetaclassModulesField(val *ResponseBadRequestSerializerMetaclassModulesField) *NullableResponseBadRequestSerializerMetaclassModulesField {
	return &NullableResponseBadRequestSerializerMetaclassModulesField{value: val, isSet: true}
}

func (v NullableResponseBadRequestSerializerMetaclassModulesField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestSerializerMetaclassModulesField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


