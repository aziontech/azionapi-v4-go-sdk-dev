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

// checks if the InputDataSourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputDataSourceAttributes{}

// InputDataSourceAttributes struct for InputDataSourceAttributes
type InputDataSourceAttributes struct {
	// * `raw_logs` - Raw Logs
	Type string `json:"type"`
	Attributes InputDataSource `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _InputDataSourceAttributes InputDataSourceAttributes

// NewInputDataSourceAttributes instantiates a new InputDataSourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDataSourceAttributes(type_ string, attributes InputDataSource) *InputDataSourceAttributes {
	this := InputDataSourceAttributes{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInputDataSourceAttributesWithDefaults instantiates a new InputDataSourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDataSourceAttributesWithDefaults() *InputDataSourceAttributes {
	this := InputDataSourceAttributes{}
	return &this
}

// GetType returns the Type field value
func (o *InputDataSourceAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputDataSourceAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputDataSourceAttributes) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InputDataSourceAttributes) GetAttributes() InputDataSource {
	if o == nil {
		var ret InputDataSource
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InputDataSourceAttributes) GetAttributesOk() (*InputDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InputDataSourceAttributes) SetAttributes(v InputDataSource) {
	o.Attributes = v
}

func (o InputDataSourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputDataSourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InputDataSourceAttributes) UnmarshalJSON(data []byte) (err error) {
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

	varInputDataSourceAttributes := _InputDataSourceAttributes{}

	err = json.Unmarshal(data, &varInputDataSourceAttributes)

	if err != nil {
		return err
	}

	*o = InputDataSourceAttributes(varInputDataSourceAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInputDataSourceAttributes struct {
	value *InputDataSourceAttributes
	isSet bool
}

func (v NullableInputDataSourceAttributes) Get() *InputDataSourceAttributes {
	return v.value
}

func (v *NullableInputDataSourceAttributes) Set(val *InputDataSourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDataSourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDataSourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDataSourceAttributes(val *InputDataSourceAttributes) *NullableInputDataSourceAttributes {
	return &NullableInputDataSourceAttributes{value: val, isSet: true}
}

func (v NullableInputDataSourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDataSourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


