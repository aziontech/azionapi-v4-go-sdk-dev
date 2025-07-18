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

// checks if the CacheSettingsModulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CacheSettingsModulesRequest{}

// CacheSettingsModulesRequest struct for CacheSettingsModulesRequest
type CacheSettingsModulesRequest struct {
	EdgeCache *CacheSettingsEdgeCacheModuleRequest `json:"edge_cache,omitempty"`
	TieredCache NullableCacheSettingsTieredCacheModuleRequest `json:"tiered_cache,omitempty"`
	ApplicationAccelerator *CacheSettingsApplicationAcceleratorModuleRequest `json:"application_accelerator,omitempty"`
}

// NewCacheSettingsModulesRequest instantiates a new CacheSettingsModulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheSettingsModulesRequest() *CacheSettingsModulesRequest {
	this := CacheSettingsModulesRequest{}
	return &this
}

// NewCacheSettingsModulesRequestWithDefaults instantiates a new CacheSettingsModulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheSettingsModulesRequestWithDefaults() *CacheSettingsModulesRequest {
	this := CacheSettingsModulesRequest{}
	return &this
}

// GetEdgeCache returns the EdgeCache field value if set, zero value otherwise.
func (o *CacheSettingsModulesRequest) GetEdgeCache() CacheSettingsEdgeCacheModuleRequest {
	if o == nil || IsNil(o.EdgeCache) {
		var ret CacheSettingsEdgeCacheModuleRequest
		return ret
	}
	return *o.EdgeCache
}

// GetEdgeCacheOk returns a tuple with the EdgeCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsModulesRequest) GetEdgeCacheOk() (*CacheSettingsEdgeCacheModuleRequest, bool) {
	if o == nil || IsNil(o.EdgeCache) {
		return nil, false
	}
	return o.EdgeCache, true
}

// HasEdgeCache returns a boolean if a field has been set.
func (o *CacheSettingsModulesRequest) HasEdgeCache() bool {
	if o != nil && !IsNil(o.EdgeCache) {
		return true
	}

	return false
}

// SetEdgeCache gets a reference to the given CacheSettingsEdgeCacheModuleRequest and assigns it to the EdgeCache field.
func (o *CacheSettingsModulesRequest) SetEdgeCache(v CacheSettingsEdgeCacheModuleRequest) {
	o.EdgeCache = &v
}

// GetTieredCache returns the TieredCache field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CacheSettingsModulesRequest) GetTieredCache() CacheSettingsTieredCacheModuleRequest {
	if o == nil || IsNil(o.TieredCache.Get()) {
		var ret CacheSettingsTieredCacheModuleRequest
		return ret
	}
	return *o.TieredCache.Get()
}

// GetTieredCacheOk returns a tuple with the TieredCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CacheSettingsModulesRequest) GetTieredCacheOk() (*CacheSettingsTieredCacheModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.TieredCache.Get(), o.TieredCache.IsSet()
}

// HasTieredCache returns a boolean if a field has been set.
func (o *CacheSettingsModulesRequest) HasTieredCache() bool {
	if o != nil && o.TieredCache.IsSet() {
		return true
	}

	return false
}

// SetTieredCache gets a reference to the given NullableCacheSettingsTieredCacheModuleRequest and assigns it to the TieredCache field.
func (o *CacheSettingsModulesRequest) SetTieredCache(v CacheSettingsTieredCacheModuleRequest) {
	o.TieredCache.Set(&v)
}
// SetTieredCacheNil sets the value for TieredCache to be an explicit nil
func (o *CacheSettingsModulesRequest) SetTieredCacheNil() {
	o.TieredCache.Set(nil)
}

// UnsetTieredCache ensures that no value is present for TieredCache, not even an explicit nil
func (o *CacheSettingsModulesRequest) UnsetTieredCache() {
	o.TieredCache.Unset()
}

// GetApplicationAccelerator returns the ApplicationAccelerator field value if set, zero value otherwise.
func (o *CacheSettingsModulesRequest) GetApplicationAccelerator() CacheSettingsApplicationAcceleratorModuleRequest {
	if o == nil || IsNil(o.ApplicationAccelerator) {
		var ret CacheSettingsApplicationAcceleratorModuleRequest
		return ret
	}
	return *o.ApplicationAccelerator
}

// GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsModulesRequest) GetApplicationAcceleratorOk() (*CacheSettingsApplicationAcceleratorModuleRequest, bool) {
	if o == nil || IsNil(o.ApplicationAccelerator) {
		return nil, false
	}
	return o.ApplicationAccelerator, true
}

// HasApplicationAccelerator returns a boolean if a field has been set.
func (o *CacheSettingsModulesRequest) HasApplicationAccelerator() bool {
	if o != nil && !IsNil(o.ApplicationAccelerator) {
		return true
	}

	return false
}

// SetApplicationAccelerator gets a reference to the given CacheSettingsApplicationAcceleratorModuleRequest and assigns it to the ApplicationAccelerator field.
func (o *CacheSettingsModulesRequest) SetApplicationAccelerator(v CacheSettingsApplicationAcceleratorModuleRequest) {
	o.ApplicationAccelerator = &v
}

func (o CacheSettingsModulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CacheSettingsModulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EdgeCache) {
		toSerialize["edge_cache"] = o.EdgeCache
	}
	if o.TieredCache.IsSet() {
		toSerialize["tiered_cache"] = o.TieredCache.Get()
	}
	if !IsNil(o.ApplicationAccelerator) {
		toSerialize["application_accelerator"] = o.ApplicationAccelerator
	}
	return toSerialize, nil
}

type NullableCacheSettingsModulesRequest struct {
	value *CacheSettingsModulesRequest
	isSet bool
}

func (v NullableCacheSettingsModulesRequest) Get() *CacheSettingsModulesRequest {
	return v.value
}

func (v *NullableCacheSettingsModulesRequest) Set(val *CacheSettingsModulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheSettingsModulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheSettingsModulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheSettingsModulesRequest(val *CacheSettingsModulesRequest) *NullableCacheSettingsModulesRequest {
	return &NullableCacheSettingsModulesRequest{value: val, isSet: true}
}

func (v NullableCacheSettingsModulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheSettingsModulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


