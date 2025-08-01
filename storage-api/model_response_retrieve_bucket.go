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

// checks if the ResponseRetrieveBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveBucket{}

// ResponseRetrieveBucket struct for ResponseRetrieveBucket
type ResponseRetrieveBucket struct {
	Data Bucket `json:"data"`
}

type _ResponseRetrieveBucket ResponseRetrieveBucket

// NewResponseRetrieveBucket instantiates a new ResponseRetrieveBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveBucket(data Bucket) *ResponseRetrieveBucket {
	this := ResponseRetrieveBucket{}
	this.Data = data
	return &this
}

// NewResponseRetrieveBucketWithDefaults instantiates a new ResponseRetrieveBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveBucketWithDefaults() *ResponseRetrieveBucket {
	this := ResponseRetrieveBucket{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveBucket) GetData() Bucket {
	if o == nil {
		var ret Bucket
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveBucket) GetDataOk() (*Bucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveBucket) SetData(v Bucket) {
	o.Data = v
}

func (o ResponseRetrieveBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveBucket) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveBucket := _ResponseRetrieveBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveBucket)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveBucket(varResponseRetrieveBucket)

	return err
}

type NullableResponseRetrieveBucket struct {
	value *ResponseRetrieveBucket
	isSet bool
}

func (v NullableResponseRetrieveBucket) Get() *ResponseRetrieveBucket {
	return v.value
}

func (v *NullableResponseRetrieveBucket) Set(val *ResponseRetrieveBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveBucket(val *ResponseRetrieveBucket) *NullableResponseRetrieveBucket {
	return &NullableResponseRetrieveBucket{value: val, isSet: true}
}

func (v NullableResponseRetrieveBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


