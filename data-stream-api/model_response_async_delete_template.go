/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
)

// checks if the ResponseAsyncDeleteTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAsyncDeleteTemplate{}

// ResponseAsyncDeleteTemplate struct for ResponseAsyncDeleteTemplate
type ResponseAsyncDeleteTemplate struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
}

// NewResponseAsyncDeleteTemplate instantiates a new ResponseAsyncDeleteTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAsyncDeleteTemplate() *ResponseAsyncDeleteTemplate {
	this := ResponseAsyncDeleteTemplate{}
	return &this
}

// NewResponseAsyncDeleteTemplateWithDefaults instantiates a new ResponseAsyncDeleteTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAsyncDeleteTemplateWithDefaults() *ResponseAsyncDeleteTemplate {
	this := ResponseAsyncDeleteTemplate{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseAsyncDeleteTemplate) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAsyncDeleteTemplate) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseAsyncDeleteTemplate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseAsyncDeleteTemplate) SetState(v string) {
	o.State = &v
}

func (o ResponseAsyncDeleteTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAsyncDeleteTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableResponseAsyncDeleteTemplate struct {
	value *ResponseAsyncDeleteTemplate
	isSet bool
}

func (v NullableResponseAsyncDeleteTemplate) Get() *ResponseAsyncDeleteTemplate {
	return v.value
}

func (v *NullableResponseAsyncDeleteTemplate) Set(val *ResponseAsyncDeleteTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAsyncDeleteTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAsyncDeleteTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAsyncDeleteTemplate(val *ResponseAsyncDeleteTemplate) *NullableResponseAsyncDeleteTemplate {
	return &NullableResponseAsyncDeleteTemplate{value: val, isSet: true}
}

func (v NullableResponseAsyncDeleteTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAsyncDeleteTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


