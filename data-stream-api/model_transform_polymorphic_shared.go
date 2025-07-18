/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TransformPolymorphicShared type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformPolymorphicShared{}

// TransformPolymorphicShared struct for TransformPolymorphicShared
type TransformPolymorphicShared struct {
	// 
	Type map[string]interface{} `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TransformPolymorphicShared TransformPolymorphicShared

// NewTransformPolymorphicShared instantiates a new TransformPolymorphicShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformPolymorphicShared(type_ map[string]interface{}) *TransformPolymorphicShared {
	this := TransformPolymorphicShared{}
	this.Type = type_
	return &this
}

// NewTransformPolymorphicSharedWithDefaults instantiates a new TransformPolymorphicShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformPolymorphicSharedWithDefaults() *TransformPolymorphicShared {
	this := TransformPolymorphicShared{}
	return &this
}

// GetType returns the Type field value
func (o *TransformPolymorphicShared) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransformPolymorphicShared) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TransformPolymorphicShared) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o TransformPolymorphicShared) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformPolymorphicShared) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransformPolymorphicShared) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varTransformPolymorphicShared := _TransformPolymorphicShared{}

	err = json.Unmarshal(data, &varTransformPolymorphicShared)

	if err != nil {
		return err
	}

	*o = TransformPolymorphicShared(varTransformPolymorphicShared)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransformPolymorphicShared struct {
	value *TransformPolymorphicShared
	isSet bool
}

func (v NullableTransformPolymorphicShared) Get() *TransformPolymorphicShared {
	return v.value
}

func (v *NullableTransformPolymorphicShared) Set(val *TransformPolymorphicShared) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformPolymorphicShared) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformPolymorphicShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformPolymorphicShared(val *TransformPolymorphicShared) *NullableTransformPolymorphicShared {
	return &NullableTransformPolymorphicShared{value: val, isSet: true}
}

func (v NullableTransformPolymorphicShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformPolymorphicShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


