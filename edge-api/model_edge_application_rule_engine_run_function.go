/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapi

import (
	"encoding/json"
	"fmt"
)

// checks if the EdgeApplicationRuleEngineRunFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineRunFunction{}

// EdgeApplicationRuleEngineRunFunction Run a function
type EdgeApplicationRuleEngineRunFunction struct {
	// * `run_function` - run_function
	Type string `json:"type"`
	Attributes EdgeApplicationRuleEngineRunFunctionAttributes `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _EdgeApplicationRuleEngineRunFunction EdgeApplicationRuleEngineRunFunction

// NewEdgeApplicationRuleEngineRunFunction instantiates a new EdgeApplicationRuleEngineRunFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineRunFunction(type_ string, attributes EdgeApplicationRuleEngineRunFunctionAttributes) *EdgeApplicationRuleEngineRunFunction {
	this := EdgeApplicationRuleEngineRunFunction{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewEdgeApplicationRuleEngineRunFunctionWithDefaults instantiates a new EdgeApplicationRuleEngineRunFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineRunFunctionWithDefaults() *EdgeApplicationRuleEngineRunFunction {
	this := EdgeApplicationRuleEngineRunFunction{}
	return &this
}

// GetType returns the Type field value
func (o *EdgeApplicationRuleEngineRunFunction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRunFunction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeApplicationRuleEngineRunFunction) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EdgeApplicationRuleEngineRunFunction) GetAttributes() EdgeApplicationRuleEngineRunFunctionAttributes {
	if o == nil {
		var ret EdgeApplicationRuleEngineRunFunctionAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRunFunction) GetAttributesOk() (*EdgeApplicationRuleEngineRunFunctionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EdgeApplicationRuleEngineRunFunction) SetAttributes(v EdgeApplicationRuleEngineRunFunctionAttributes) {
	o.Attributes = v
}

func (o EdgeApplicationRuleEngineRunFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineRunFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineRunFunction) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeApplicationRuleEngineRunFunction := _EdgeApplicationRuleEngineRunFunction{}

	err = json.Unmarshal(data, &varEdgeApplicationRuleEngineRunFunction)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineRunFunction(varEdgeApplicationRuleEngineRunFunction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeApplicationRuleEngineRunFunction struct {
	value *EdgeApplicationRuleEngineRunFunction
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineRunFunction) Get() *EdgeApplicationRuleEngineRunFunction {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineRunFunction) Set(val *EdgeApplicationRuleEngineRunFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineRunFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineRunFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineRunFunction(val *EdgeApplicationRuleEngineRunFunction) *NullableEdgeApplicationRuleEngineRunFunction {
	return &NullableEdgeApplicationRuleEngineRunFunction{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineRunFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineRunFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


