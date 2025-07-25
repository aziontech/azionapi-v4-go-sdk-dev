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

// checks if the TransformPolymorphicTransformFilterWorkloadsAttributesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformPolymorphicTransformFilterWorkloadsAttributesRequest{}

// TransformPolymorphicTransformFilterWorkloadsAttributesRequest struct for TransformPolymorphicTransformFilterWorkloadsAttributesRequest
type TransformPolymorphicTransformFilterWorkloadsAttributesRequest struct {
	// * `filter_workloads` - Filter Workloads
	Type string `json:"type"`
	Attributes TransformFilterWorkloadsRequest `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _TransformPolymorphicTransformFilterWorkloadsAttributesRequest TransformPolymorphicTransformFilterWorkloadsAttributesRequest

// NewTransformPolymorphicTransformFilterWorkloadsAttributesRequest instantiates a new TransformPolymorphicTransformFilterWorkloadsAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformPolymorphicTransformFilterWorkloadsAttributesRequest(type_ string, attributes TransformFilterWorkloadsRequest) *TransformPolymorphicTransformFilterWorkloadsAttributesRequest {
	this := TransformPolymorphicTransformFilterWorkloadsAttributesRequest{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTransformPolymorphicTransformFilterWorkloadsAttributesRequestWithDefaults instantiates a new TransformPolymorphicTransformFilterWorkloadsAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformPolymorphicTransformFilterWorkloadsAttributesRequestWithDefaults() *TransformPolymorphicTransformFilterWorkloadsAttributesRequest {
	this := TransformPolymorphicTransformFilterWorkloadsAttributesRequest{}
	return &this
}

// GetType returns the Type field value
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) GetAttributes() TransformFilterWorkloadsRequest {
	if o == nil {
		var ret TransformFilterWorkloadsRequest
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) GetAttributesOk() (*TransformFilterWorkloadsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) SetAttributes(v TransformFilterWorkloadsRequest) {
	o.Attributes = v
}

func (o TransformPolymorphicTransformFilterWorkloadsAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformPolymorphicTransformFilterWorkloadsAttributesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varTransformPolymorphicTransformFilterWorkloadsAttributesRequest := _TransformPolymorphicTransformFilterWorkloadsAttributesRequest{}

	err = json.Unmarshal(data, &varTransformPolymorphicTransformFilterWorkloadsAttributesRequest)

	if err != nil {
		return err
	}

	*o = TransformPolymorphicTransformFilterWorkloadsAttributesRequest(varTransformPolymorphicTransformFilterWorkloadsAttributesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest struct {
	value *TransformPolymorphicTransformFilterWorkloadsAttributesRequest
	isSet bool
}

func (v NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) Get() *TransformPolymorphicTransformFilterWorkloadsAttributesRequest {
	return v.value
}

func (v *NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) Set(val *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest(val *TransformPolymorphicTransformFilterWorkloadsAttributesRequest) *NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest {
	return &NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest{value: val, isSet: true}
}

func (v NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformPolymorphicTransformFilterWorkloadsAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


