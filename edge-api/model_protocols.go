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

// checks if the Protocols type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Protocols{}

// Protocols struct for Protocols
type Protocols struct {
	Http *HttpProtocol `json:"http,omitempty"`
}

// NewProtocols instantiates a new Protocols object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocols() *Protocols {
	this := Protocols{}
	return &this
}

// NewProtocolsWithDefaults instantiates a new Protocols object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolsWithDefaults() *Protocols {
	this := Protocols{}
	return &this
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *Protocols) GetHttp() HttpProtocol {
	if o == nil || IsNil(o.Http) {
		var ret HttpProtocol
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocols) GetHttpOk() (*HttpProtocol, bool) {
	if o == nil || IsNil(o.Http) {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *Protocols) HasHttp() bool {
	if o != nil && !IsNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given HttpProtocol and assigns it to the Http field.
func (o *Protocols) SetHttp(v HttpProtocol) {
	o.Http = &v
}

func (o Protocols) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Protocols) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	return toSerialize, nil
}

type NullableProtocols struct {
	value *Protocols
	isSet bool
}

func (v NullableProtocols) Get() *Protocols {
	return v.value
}

func (v *NullableProtocols) Set(val *Protocols) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocols(val *Protocols) *NullableProtocols {
	return &NullableProtocols{value: val, isSet: true}
}

func (v NullableProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


