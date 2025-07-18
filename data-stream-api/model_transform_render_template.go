/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransformRenderTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformRenderTemplate{}

// TransformRenderTemplate Serializer for render_template transform attributes.
type TransformRenderTemplate struct {
	Template int64 `json:"template"`
}

type _TransformRenderTemplate TransformRenderTemplate

// NewTransformRenderTemplate instantiates a new TransformRenderTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformRenderTemplate(template int64) *TransformRenderTemplate {
	this := TransformRenderTemplate{}
	this.Template = template
	return &this
}

// NewTransformRenderTemplateWithDefaults instantiates a new TransformRenderTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformRenderTemplateWithDefaults() *TransformRenderTemplate {
	this := TransformRenderTemplate{}
	return &this
}

// GetTemplate returns the Template field value
func (o *TransformRenderTemplate) GetTemplate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *TransformRenderTemplate) GetTemplateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *TransformRenderTemplate) SetTemplate(v int64) {
	o.Template = v
}

func (o TransformRenderTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformRenderTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *TransformRenderTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template",
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

	varTransformRenderTemplate := _TransformRenderTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransformRenderTemplate)

	if err != nil {
		return err
	}

	*o = TransformRenderTemplate(varTransformRenderTemplate)

	return err
}

type NullableTransformRenderTemplate struct {
	value *TransformRenderTemplate
	isSet bool
}

func (v NullableTransformRenderTemplate) Get() *TransformRenderTemplate {
	return v.value
}

func (v *NullableTransformRenderTemplate) Set(val *TransformRenderTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformRenderTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformRenderTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformRenderTemplate(val *TransformRenderTemplate) *NullableTransformRenderTemplate {
	return &NullableTransformRenderTemplate{value: val, isSet: true}
}

func (v NullableTransformRenderTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformRenderTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


