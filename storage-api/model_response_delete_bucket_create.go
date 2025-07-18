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

// checks if the ResponseDeleteBucketCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeleteBucketCreate{}

// ResponseDeleteBucketCreate struct for ResponseDeleteBucketCreate
type ResponseDeleteBucketCreate struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
}

// NewResponseDeleteBucketCreate instantiates a new ResponseDeleteBucketCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeleteBucketCreate() *ResponseDeleteBucketCreate {
	this := ResponseDeleteBucketCreate{}
	return &this
}

// NewResponseDeleteBucketCreateWithDefaults instantiates a new ResponseDeleteBucketCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeleteBucketCreateWithDefaults() *ResponseDeleteBucketCreate {
	this := ResponseDeleteBucketCreate{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseDeleteBucketCreate) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeleteBucketCreate) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseDeleteBucketCreate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseDeleteBucketCreate) SetState(v string) {
	o.State = &v
}

func (o ResponseDeleteBucketCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeleteBucketCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableResponseDeleteBucketCreate struct {
	value *ResponseDeleteBucketCreate
	isSet bool
}

func (v NullableResponseDeleteBucketCreate) Get() *ResponseDeleteBucketCreate {
	return v.value
}

func (v *NullableResponseDeleteBucketCreate) Set(val *ResponseDeleteBucketCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeleteBucketCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeleteBucketCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeleteBucketCreate(val *ResponseDeleteBucketCreate) *NullableResponseDeleteBucketCreate {
	return &NullableResponseDeleteBucketCreate{value: val, isSet: true}
}

func (v NullableResponseDeleteBucketCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeleteBucketCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


