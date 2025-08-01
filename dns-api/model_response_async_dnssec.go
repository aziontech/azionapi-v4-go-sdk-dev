/*
dns-api

REST API OpenAPI documentation for the DNS API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseAsyncDNSSEC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAsyncDNSSEC{}

// ResponseAsyncDNSSEC struct for ResponseAsyncDNSSEC
type ResponseAsyncDNSSEC struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
	Data DNSSEC `json:"data"`
}

type _ResponseAsyncDNSSEC ResponseAsyncDNSSEC

// NewResponseAsyncDNSSEC instantiates a new ResponseAsyncDNSSEC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAsyncDNSSEC(data DNSSEC) *ResponseAsyncDNSSEC {
	this := ResponseAsyncDNSSEC{}
	this.Data = data
	return &this
}

// NewResponseAsyncDNSSECWithDefaults instantiates a new ResponseAsyncDNSSEC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAsyncDNSSECWithDefaults() *ResponseAsyncDNSSEC {
	this := ResponseAsyncDNSSEC{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseAsyncDNSSEC) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAsyncDNSSEC) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseAsyncDNSSEC) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseAsyncDNSSEC) SetState(v string) {
	o.State = &v
}

// GetData returns the Data field value
func (o *ResponseAsyncDNSSEC) GetData() DNSSEC {
	if o == nil {
		var ret DNSSEC
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseAsyncDNSSEC) GetDataOk() (*DNSSEC, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseAsyncDNSSEC) SetData(v DNSSEC) {
	o.Data = v
}

func (o ResponseAsyncDNSSEC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAsyncDNSSEC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseAsyncDNSSEC) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varResponseAsyncDNSSEC := _ResponseAsyncDNSSEC{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseAsyncDNSSEC)

	if err != nil {
		return err
	}

	*o = ResponseAsyncDNSSEC(varResponseAsyncDNSSEC)

	return err
}

type NullableResponseAsyncDNSSEC struct {
	value *ResponseAsyncDNSSEC
	isSet bool
}

func (v NullableResponseAsyncDNSSEC) Get() *ResponseAsyncDNSSEC {
	return v.value
}

func (v *NullableResponseAsyncDNSSEC) Set(val *ResponseAsyncDNSSEC) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAsyncDNSSEC) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAsyncDNSSEC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAsyncDNSSEC(val *ResponseAsyncDNSSEC) *NullableResponseAsyncDNSSEC {
	return &NullableResponseAsyncDNSSEC{value: val, isSet: true}
}

func (v NullableResponseAsyncDNSSEC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAsyncDNSSEC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


