/*
edge-storage-api

REST API OpenAPI documentation for the Edge Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storageapi

import (
	"encoding/json"
)

// checks if the ResponseAsyncDeleteBucketCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAsyncDeleteBucketCreate{}

// ResponseAsyncDeleteBucketCreate struct for ResponseAsyncDeleteBucketCreate
type ResponseAsyncDeleteBucketCreate struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
}

// NewResponseAsyncDeleteBucketCreate instantiates a new ResponseAsyncDeleteBucketCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAsyncDeleteBucketCreate() *ResponseAsyncDeleteBucketCreate {
	this := ResponseAsyncDeleteBucketCreate{}
	return &this
}

// NewResponseAsyncDeleteBucketCreateWithDefaults instantiates a new ResponseAsyncDeleteBucketCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAsyncDeleteBucketCreateWithDefaults() *ResponseAsyncDeleteBucketCreate {
	this := ResponseAsyncDeleteBucketCreate{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseAsyncDeleteBucketCreate) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAsyncDeleteBucketCreate) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseAsyncDeleteBucketCreate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseAsyncDeleteBucketCreate) SetState(v string) {
	o.State = &v
}

func (o ResponseAsyncDeleteBucketCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAsyncDeleteBucketCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableResponseAsyncDeleteBucketCreate struct {
	value *ResponseAsyncDeleteBucketCreate
	isSet bool
}

func (v NullableResponseAsyncDeleteBucketCreate) Get() *ResponseAsyncDeleteBucketCreate {
	return v.value
}

func (v *NullableResponseAsyncDeleteBucketCreate) Set(val *ResponseAsyncDeleteBucketCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAsyncDeleteBucketCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAsyncDeleteBucketCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAsyncDeleteBucketCreate(val *ResponseAsyncDeleteBucketCreate) *NullableResponseAsyncDeleteBucketCreate {
	return &NullableResponseAsyncDeleteBucketCreate{value: val, isSet: true}
}

func (v NullableResponseAsyncDeleteBucketCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAsyncDeleteBucketCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


