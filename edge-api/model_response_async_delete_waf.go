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

// checks if the ResponseAsyncDeleteWAF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAsyncDeleteWAF{}

// ResponseAsyncDeleteWAF struct for ResponseAsyncDeleteWAF
type ResponseAsyncDeleteWAF struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
}

// NewResponseAsyncDeleteWAF instantiates a new ResponseAsyncDeleteWAF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAsyncDeleteWAF() *ResponseAsyncDeleteWAF {
	this := ResponseAsyncDeleteWAF{}
	return &this
}

// NewResponseAsyncDeleteWAFWithDefaults instantiates a new ResponseAsyncDeleteWAF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAsyncDeleteWAFWithDefaults() *ResponseAsyncDeleteWAF {
	this := ResponseAsyncDeleteWAF{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseAsyncDeleteWAF) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAsyncDeleteWAF) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseAsyncDeleteWAF) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseAsyncDeleteWAF) SetState(v string) {
	o.State = &v
}

func (o ResponseAsyncDeleteWAF) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAsyncDeleteWAF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableResponseAsyncDeleteWAF struct {
	value *ResponseAsyncDeleteWAF
	isSet bool
}

func (v NullableResponseAsyncDeleteWAF) Get() *ResponseAsyncDeleteWAF {
	return v.value
}

func (v *NullableResponseAsyncDeleteWAF) Set(val *ResponseAsyncDeleteWAF) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAsyncDeleteWAF) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAsyncDeleteWAF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAsyncDeleteWAF(val *ResponseAsyncDeleteWAF) *NullableResponseAsyncDeleteWAF {
	return &NullableResponseAsyncDeleteWAF{value: val, isSet: true}
}

func (v NullableResponseAsyncDeleteWAF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAsyncDeleteWAF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


