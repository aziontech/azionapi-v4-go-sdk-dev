/*
edge-storage-api

REST API OpenAPI documentation for the Edge Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storageapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseBucketCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBucketCreate{}

// ResponseBucketCreate struct for ResponseBucketCreate
type ResponseBucketCreate struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
	Data BucketCreate `json:"data"`
}

type _ResponseBucketCreate ResponseBucketCreate

// NewResponseBucketCreate instantiates a new ResponseBucketCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBucketCreate(data BucketCreate) *ResponseBucketCreate {
	this := ResponseBucketCreate{}
	this.Data = data
	return &this
}

// NewResponseBucketCreateWithDefaults instantiates a new ResponseBucketCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBucketCreateWithDefaults() *ResponseBucketCreate {
	this := ResponseBucketCreate{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseBucketCreate) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBucketCreate) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseBucketCreate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseBucketCreate) SetState(v string) {
	o.State = &v
}

// GetData returns the Data field value
func (o *ResponseBucketCreate) GetData() BucketCreate {
	if o == nil {
		var ret BucketCreate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseBucketCreate) GetDataOk() (*BucketCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseBucketCreate) SetData(v BucketCreate) {
	o.Data = v
}

func (o ResponseBucketCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBucketCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseBucketCreate) UnmarshalJSON(data []byte) (err error) {
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

	varResponseBucketCreate := _ResponseBucketCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseBucketCreate)

	if err != nil {
		return err
	}

	*o = ResponseBucketCreate(varResponseBucketCreate)

	return err
}

type NullableResponseBucketCreate struct {
	value *ResponseBucketCreate
	isSet bool
}

func (v NullableResponseBucketCreate) Get() *ResponseBucketCreate {
	return v.value
}

func (v *NullableResponseBucketCreate) Set(val *ResponseBucketCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBucketCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBucketCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBucketCreate(val *ResponseBucketCreate) *NullableResponseBucketCreate {
	return &NullableResponseBucketCreate{value: val, isSet: true}
}

func (v NullableResponseBucketCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBucketCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


