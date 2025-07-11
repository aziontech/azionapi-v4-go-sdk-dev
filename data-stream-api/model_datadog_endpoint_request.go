/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DatadogEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadogEndpointRequest{}

// DatadogEndpointRequest struct for DatadogEndpointRequest
type DatadogEndpointRequest struct {
	Url string `json:"url"`
	ApiKey string `json:"api_key"`
}

type _DatadogEndpointRequest DatadogEndpointRequest

// NewDatadogEndpointRequest instantiates a new DatadogEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadogEndpointRequest(url string, apiKey string) *DatadogEndpointRequest {
	this := DatadogEndpointRequest{}
	this.Url = url
	this.ApiKey = apiKey
	return &this
}

// NewDatadogEndpointRequestWithDefaults instantiates a new DatadogEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadogEndpointRequestWithDefaults() *DatadogEndpointRequest {
	this := DatadogEndpointRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *DatadogEndpointRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DatadogEndpointRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DatadogEndpointRequest) SetUrl(v string) {
	o.Url = v
}

// GetApiKey returns the ApiKey field value
func (o *DatadogEndpointRequest) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *DatadogEndpointRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *DatadogEndpointRequest) SetApiKey(v string) {
	o.ApiKey = v
}

func (o DatadogEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadogEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["api_key"] = o.ApiKey
	return toSerialize, nil
}

func (o *DatadogEndpointRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"api_key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDatadogEndpointRequest := _DatadogEndpointRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDatadogEndpointRequest)

	if err != nil {
		return err
	}

	*o = DatadogEndpointRequest(varDatadogEndpointRequest)

	return err
}

type NullableDatadogEndpointRequest struct {
	value *DatadogEndpointRequest
	isSet bool
}

func (v NullableDatadogEndpointRequest) Get() *DatadogEndpointRequest {
	return v.value
}

func (v *NullableDatadogEndpointRequest) Set(val *DatadogEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadogEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadogEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadogEndpointRequest(val *DatadogEndpointRequest) *NullableDatadogEndpointRequest {
	return &NullableDatadogEndpointRequest{value: val, isSet: true}
}

func (v NullableDatadogEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadogEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


