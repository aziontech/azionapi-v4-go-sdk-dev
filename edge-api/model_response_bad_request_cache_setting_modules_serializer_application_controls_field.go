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

// checks if the ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField{}

// ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField struct for ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField
type ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField struct {
	CacheByQueryString []string `json:"cache_by_query_string,omitempty"`
	QueryStringFields *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields `json:"query_string_fields,omitempty"`
	QueryStringSortEnabled []string `json:"query_string_sort_enabled,omitempty"`
	CacheByCookies []string `json:"cache_by_cookies,omitempty"`
	CookieNames *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields `json:"cookie_names,omitempty"`
	AdaptiveDeliveryAction []string `json:"adaptive_delivery_action,omitempty"`
	DeviceGroup []string `json:"device_group,omitempty"`
}

// NewResponseBadRequestCacheSettingModulesSerializerApplicationControlsField instantiates a new ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestCacheSettingModulesSerializerApplicationControlsField() *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField {
	this := ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField{}
	return &this
}

// NewResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldWithDefaults instantiates a new ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldWithDefaults() *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField {
	this := ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField{}
	return &this
}

// GetCacheByQueryString returns the CacheByQueryString field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCacheByQueryString() []string {
	if o == nil || IsNil(o.CacheByQueryString) {
		var ret []string
		return ret
	}
	return o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCacheByQueryStringOk() ([]string, bool) {
	if o == nil || IsNil(o.CacheByQueryString) {
		return nil, false
	}
	return o.CacheByQueryString, true
}

// HasCacheByQueryString returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasCacheByQueryString() bool {
	if o != nil && !IsNil(o.CacheByQueryString) {
		return true
	}

	return false
}

// SetCacheByQueryString gets a reference to the given []string and assigns it to the CacheByQueryString field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetCacheByQueryString(v []string) {
	o.CacheByQueryString = v
}

// GetQueryStringFields returns the QueryStringFields field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetQueryStringFields() ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields {
	if o == nil || IsNil(o.QueryStringFields) {
		var ret ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields
		return ret
	}
	return *o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetQueryStringFieldsOk() (*ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields, bool) {
	if o == nil || IsNil(o.QueryStringFields) {
		return nil, false
	}
	return o.QueryStringFields, true
}

// HasQueryStringFields returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasQueryStringFields() bool {
	if o != nil && !IsNil(o.QueryStringFields) {
		return true
	}

	return false
}

// SetQueryStringFields gets a reference to the given ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields and assigns it to the QueryStringFields field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetQueryStringFields(v ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields) {
	o.QueryStringFields = &v
}

// GetQueryStringSortEnabled returns the QueryStringSortEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetQueryStringSortEnabled() []string {
	if o == nil || IsNil(o.QueryStringSortEnabled) {
		var ret []string
		return ret
	}
	return o.QueryStringSortEnabled
}

// GetQueryStringSortEnabledOk returns a tuple with the QueryStringSortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetQueryStringSortEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.QueryStringSortEnabled) {
		return nil, false
	}
	return o.QueryStringSortEnabled, true
}

// HasQueryStringSortEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasQueryStringSortEnabled() bool {
	if o != nil && !IsNil(o.QueryStringSortEnabled) {
		return true
	}

	return false
}

// SetQueryStringSortEnabled gets a reference to the given []string and assigns it to the QueryStringSortEnabled field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetQueryStringSortEnabled(v []string) {
	o.QueryStringSortEnabled = v
}

// GetCacheByCookies returns the CacheByCookies field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCacheByCookies() []string {
	if o == nil || IsNil(o.CacheByCookies) {
		var ret []string
		return ret
	}
	return o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCacheByCookiesOk() ([]string, bool) {
	if o == nil || IsNil(o.CacheByCookies) {
		return nil, false
	}
	return o.CacheByCookies, true
}

// HasCacheByCookies returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasCacheByCookies() bool {
	if o != nil && !IsNil(o.CacheByCookies) {
		return true
	}

	return false
}

// SetCacheByCookies gets a reference to the given []string and assigns it to the CacheByCookies field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetCacheByCookies(v []string) {
	o.CacheByCookies = v
}

// GetCookieNames returns the CookieNames field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCookieNames() ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields {
	if o == nil || IsNil(o.CookieNames) {
		var ret ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields
		return ret
	}
	return *o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetCookieNamesOk() (*ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields, bool) {
	if o == nil || IsNil(o.CookieNames) {
		return nil, false
	}
	return o.CookieNames, true
}

// HasCookieNames returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasCookieNames() bool {
	if o != nil && !IsNil(o.CookieNames) {
		return true
	}

	return false
}

// SetCookieNames gets a reference to the given ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields and assigns it to the CookieNames field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetCookieNames(v ResponseBadRequestCacheSettingModulesSerializerApplicationControlsFieldQueryStringFields) {
	o.CookieNames = &v
}

// GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetAdaptiveDeliveryAction() []string {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		var ret []string
		return ret
	}
	return o.AdaptiveDeliveryAction
}

// GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetAdaptiveDeliveryActionOk() ([]string, bool) {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		return nil, false
	}
	return o.AdaptiveDeliveryAction, true
}

// HasAdaptiveDeliveryAction returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasAdaptiveDeliveryAction() bool {
	if o != nil && !IsNil(o.AdaptiveDeliveryAction) {
		return true
	}

	return false
}

// SetAdaptiveDeliveryAction gets a reference to the given []string and assigns it to the AdaptiveDeliveryAction field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetAdaptiveDeliveryAction(v []string) {
	o.AdaptiveDeliveryAction = v
}

// GetDeviceGroup returns the DeviceGroup field value if set, zero value otherwise.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetDeviceGroup() []string {
	if o == nil || IsNil(o.DeviceGroup) {
		var ret []string
		return ret
	}
	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) GetDeviceGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceGroup) {
		return nil, false
	}
	return o.DeviceGroup, true
}

// HasDeviceGroup returns a boolean if a field has been set.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) HasDeviceGroup() bool {
	if o != nil && !IsNil(o.DeviceGroup) {
		return true
	}

	return false
}

// SetDeviceGroup gets a reference to the given []string and assigns it to the DeviceGroup field.
func (o *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) SetDeviceGroup(v []string) {
	o.DeviceGroup = v
}

func (o ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheByQueryString) {
		toSerialize["cache_by_query_string"] = o.CacheByQueryString
	}
	if !IsNil(o.QueryStringFields) {
		toSerialize["query_string_fields"] = o.QueryStringFields
	}
	if !IsNil(o.QueryStringSortEnabled) {
		toSerialize["query_string_sort_enabled"] = o.QueryStringSortEnabled
	}
	if !IsNil(o.CacheByCookies) {
		toSerialize["cache_by_cookies"] = o.CacheByCookies
	}
	if !IsNil(o.CookieNames) {
		toSerialize["cookie_names"] = o.CookieNames
	}
	if !IsNil(o.AdaptiveDeliveryAction) {
		toSerialize["adaptive_delivery_action"] = o.AdaptiveDeliveryAction
	}
	if !IsNil(o.DeviceGroup) {
		toSerialize["device_group"] = o.DeviceGroup
	}
	return toSerialize, nil
}

type NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField struct {
	value *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField
	isSet bool
}

func (v NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) Get() *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField {
	return v.value
}

func (v *NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) Set(val *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField(val *ResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) *NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField {
	return &NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField{value: val, isSet: true}
}

func (v NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestCacheSettingModulesSerializerApplicationControlsField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


