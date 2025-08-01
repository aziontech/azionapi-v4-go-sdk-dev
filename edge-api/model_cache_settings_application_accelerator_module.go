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

// checks if the CacheSettingsApplicationAcceleratorModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CacheSettingsApplicationAcceleratorModule{}

// CacheSettingsApplicationAcceleratorModule struct for CacheSettingsApplicationAcceleratorModule
type CacheSettingsApplicationAcceleratorModule struct {
	CacheVaryByMethod []string `json:"cache_vary_by_method,omitempty"`
	CacheVaryByQuerystring *CacheVaryByQuerystringModule `json:"cache_vary_by_querystring,omitempty"`
	CacheVaryByCookies *CacheVaryByCookiesModule `json:"cache_vary_by_cookies,omitempty"`
	CacheVaryByDevices *CacheVaryByDevicesModule `json:"cache_vary_by_devices,omitempty"`
}

// NewCacheSettingsApplicationAcceleratorModule instantiates a new CacheSettingsApplicationAcceleratorModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheSettingsApplicationAcceleratorModule() *CacheSettingsApplicationAcceleratorModule {
	this := CacheSettingsApplicationAcceleratorModule{}
	return &this
}

// NewCacheSettingsApplicationAcceleratorModuleWithDefaults instantiates a new CacheSettingsApplicationAcceleratorModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheSettingsApplicationAcceleratorModuleWithDefaults() *CacheSettingsApplicationAcceleratorModule {
	this := CacheSettingsApplicationAcceleratorModule{}
	return &this
}

// GetCacheVaryByMethod returns the CacheVaryByMethod field value if set, zero value otherwise.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByMethod() []string {
	if o == nil || IsNil(o.CacheVaryByMethod) {
		var ret []string
		return ret
	}
	return o.CacheVaryByMethod
}

// GetCacheVaryByMethodOk returns a tuple with the CacheVaryByMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByMethodOk() ([]string, bool) {
	if o == nil || IsNil(o.CacheVaryByMethod) {
		return nil, false
	}
	return o.CacheVaryByMethod, true
}

// HasCacheVaryByMethod returns a boolean if a field has been set.
func (o *CacheSettingsApplicationAcceleratorModule) HasCacheVaryByMethod() bool {
	if o != nil && !IsNil(o.CacheVaryByMethod) {
		return true
	}

	return false
}

// SetCacheVaryByMethod gets a reference to the given []string and assigns it to the CacheVaryByMethod field.
func (o *CacheSettingsApplicationAcceleratorModule) SetCacheVaryByMethod(v []string) {
	o.CacheVaryByMethod = v
}

// GetCacheVaryByQuerystring returns the CacheVaryByQuerystring field value if set, zero value otherwise.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByQuerystring() CacheVaryByQuerystringModule {
	if o == nil || IsNil(o.CacheVaryByQuerystring) {
		var ret CacheVaryByQuerystringModule
		return ret
	}
	return *o.CacheVaryByQuerystring
}

// GetCacheVaryByQuerystringOk returns a tuple with the CacheVaryByQuerystring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByQuerystringOk() (*CacheVaryByQuerystringModule, bool) {
	if o == nil || IsNil(o.CacheVaryByQuerystring) {
		return nil, false
	}
	return o.CacheVaryByQuerystring, true
}

// HasCacheVaryByQuerystring returns a boolean if a field has been set.
func (o *CacheSettingsApplicationAcceleratorModule) HasCacheVaryByQuerystring() bool {
	if o != nil && !IsNil(o.CacheVaryByQuerystring) {
		return true
	}

	return false
}

// SetCacheVaryByQuerystring gets a reference to the given CacheVaryByQuerystringModule and assigns it to the CacheVaryByQuerystring field.
func (o *CacheSettingsApplicationAcceleratorModule) SetCacheVaryByQuerystring(v CacheVaryByQuerystringModule) {
	o.CacheVaryByQuerystring = &v
}

// GetCacheVaryByCookies returns the CacheVaryByCookies field value if set, zero value otherwise.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByCookies() CacheVaryByCookiesModule {
	if o == nil || IsNil(o.CacheVaryByCookies) {
		var ret CacheVaryByCookiesModule
		return ret
	}
	return *o.CacheVaryByCookies
}

// GetCacheVaryByCookiesOk returns a tuple with the CacheVaryByCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByCookiesOk() (*CacheVaryByCookiesModule, bool) {
	if o == nil || IsNil(o.CacheVaryByCookies) {
		return nil, false
	}
	return o.CacheVaryByCookies, true
}

// HasCacheVaryByCookies returns a boolean if a field has been set.
func (o *CacheSettingsApplicationAcceleratorModule) HasCacheVaryByCookies() bool {
	if o != nil && !IsNil(o.CacheVaryByCookies) {
		return true
	}

	return false
}

// SetCacheVaryByCookies gets a reference to the given CacheVaryByCookiesModule and assigns it to the CacheVaryByCookies field.
func (o *CacheSettingsApplicationAcceleratorModule) SetCacheVaryByCookies(v CacheVaryByCookiesModule) {
	o.CacheVaryByCookies = &v
}

// GetCacheVaryByDevices returns the CacheVaryByDevices field value if set, zero value otherwise.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByDevices() CacheVaryByDevicesModule {
	if o == nil || IsNil(o.CacheVaryByDevices) {
		var ret CacheVaryByDevicesModule
		return ret
	}
	return *o.CacheVaryByDevices
}

// GetCacheVaryByDevicesOk returns a tuple with the CacheVaryByDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsApplicationAcceleratorModule) GetCacheVaryByDevicesOk() (*CacheVaryByDevicesModule, bool) {
	if o == nil || IsNil(o.CacheVaryByDevices) {
		return nil, false
	}
	return o.CacheVaryByDevices, true
}

// HasCacheVaryByDevices returns a boolean if a field has been set.
func (o *CacheSettingsApplicationAcceleratorModule) HasCacheVaryByDevices() bool {
	if o != nil && !IsNil(o.CacheVaryByDevices) {
		return true
	}

	return false
}

// SetCacheVaryByDevices gets a reference to the given CacheVaryByDevicesModule and assigns it to the CacheVaryByDevices field.
func (o *CacheSettingsApplicationAcceleratorModule) SetCacheVaryByDevices(v CacheVaryByDevicesModule) {
	o.CacheVaryByDevices = &v
}

func (o CacheSettingsApplicationAcceleratorModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CacheSettingsApplicationAcceleratorModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheVaryByMethod) {
		toSerialize["cache_vary_by_method"] = o.CacheVaryByMethod
	}
	if !IsNil(o.CacheVaryByQuerystring) {
		toSerialize["cache_vary_by_querystring"] = o.CacheVaryByQuerystring
	}
	if !IsNil(o.CacheVaryByCookies) {
		toSerialize["cache_vary_by_cookies"] = o.CacheVaryByCookies
	}
	if !IsNil(o.CacheVaryByDevices) {
		toSerialize["cache_vary_by_devices"] = o.CacheVaryByDevices
	}
	return toSerialize, nil
}

type NullableCacheSettingsApplicationAcceleratorModule struct {
	value *CacheSettingsApplicationAcceleratorModule
	isSet bool
}

func (v NullableCacheSettingsApplicationAcceleratorModule) Get() *CacheSettingsApplicationAcceleratorModule {
	return v.value
}

func (v *NullableCacheSettingsApplicationAcceleratorModule) Set(val *CacheSettingsApplicationAcceleratorModule) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheSettingsApplicationAcceleratorModule) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheSettingsApplicationAcceleratorModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheSettingsApplicationAcceleratorModule(val *CacheSettingsApplicationAcceleratorModule) *NullableCacheSettingsApplicationAcceleratorModule {
	return &NullableCacheSettingsApplicationAcceleratorModule{value: val, isSet: true}
}

func (v NullableCacheSettingsApplicationAcceleratorModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheSettingsApplicationAcceleratorModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


