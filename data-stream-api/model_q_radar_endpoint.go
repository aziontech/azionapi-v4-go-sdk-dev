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

// checks if the QRadarEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QRadarEndpoint{}

// QRadarEndpoint struct for QRadarEndpoint
type QRadarEndpoint struct {
	Url string `json:"url"`
}

type _QRadarEndpoint QRadarEndpoint

// NewQRadarEndpoint instantiates a new QRadarEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQRadarEndpoint(url string) *QRadarEndpoint {
	this := QRadarEndpoint{}
	this.Url = url
	return &this
}

// NewQRadarEndpointWithDefaults instantiates a new QRadarEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQRadarEndpointWithDefaults() *QRadarEndpoint {
	this := QRadarEndpoint{}
	return &this
}

// GetUrl returns the Url field value
func (o *QRadarEndpoint) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QRadarEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *QRadarEndpoint) SetUrl(v string) {
	o.Url = v
}

func (o QRadarEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QRadarEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *QRadarEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varQRadarEndpoint := _QRadarEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQRadarEndpoint)

	if err != nil {
		return err
	}

	*o = QRadarEndpoint(varQRadarEndpoint)

	return err
}

type NullableQRadarEndpoint struct {
	value *QRadarEndpoint
	isSet bool
}

func (v NullableQRadarEndpoint) Get() *QRadarEndpoint {
	return v.value
}

func (v *NullableQRadarEndpoint) Set(val *QRadarEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableQRadarEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableQRadarEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRadarEndpoint(val *QRadarEndpoint) *NullableQRadarEndpoint {
	return &NullableQRadarEndpoint{value: val, isSet: true}
}

func (v NullableQRadarEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRadarEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


