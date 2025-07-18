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

// checks if the EdgeFirewallBehaviorRunFunctionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallBehaviorRunFunctionAttributes{}

// EdgeFirewallBehaviorRunFunctionAttributes struct for EdgeFirewallBehaviorRunFunctionAttributes
type EdgeFirewallBehaviorRunFunctionAttributes struct {
	Value int64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _EdgeFirewallBehaviorRunFunctionAttributes EdgeFirewallBehaviorRunFunctionAttributes

// NewEdgeFirewallBehaviorRunFunctionAttributes instantiates a new EdgeFirewallBehaviorRunFunctionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallBehaviorRunFunctionAttributes(value int64) *EdgeFirewallBehaviorRunFunctionAttributes {
	this := EdgeFirewallBehaviorRunFunctionAttributes{}
	this.Value = value
	return &this
}

// NewEdgeFirewallBehaviorRunFunctionAttributesWithDefaults instantiates a new EdgeFirewallBehaviorRunFunctionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallBehaviorRunFunctionAttributesWithDefaults() *EdgeFirewallBehaviorRunFunctionAttributes {
	this := EdgeFirewallBehaviorRunFunctionAttributes{}
	return &this
}

// GetValue returns the Value field value
func (o *EdgeFirewallBehaviorRunFunctionAttributes) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallBehaviorRunFunctionAttributes) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EdgeFirewallBehaviorRunFunctionAttributes) SetValue(v int64) {
	o.Value = v
}

func (o EdgeFirewallBehaviorRunFunctionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallBehaviorRunFunctionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeFirewallBehaviorRunFunctionAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varEdgeFirewallBehaviorRunFunctionAttributes := _EdgeFirewallBehaviorRunFunctionAttributes{}

	err = json.Unmarshal(data, &varEdgeFirewallBehaviorRunFunctionAttributes)

	if err != nil {
		return err
	}

	*o = EdgeFirewallBehaviorRunFunctionAttributes(varEdgeFirewallBehaviorRunFunctionAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeFirewallBehaviorRunFunctionAttributes struct {
	value *EdgeFirewallBehaviorRunFunctionAttributes
	isSet bool
}

func (v NullableEdgeFirewallBehaviorRunFunctionAttributes) Get() *EdgeFirewallBehaviorRunFunctionAttributes {
	return v.value
}

func (v *NullableEdgeFirewallBehaviorRunFunctionAttributes) Set(val *EdgeFirewallBehaviorRunFunctionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallBehaviorRunFunctionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallBehaviorRunFunctionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallBehaviorRunFunctionAttributes(val *EdgeFirewallBehaviorRunFunctionAttributes) *NullableEdgeFirewallBehaviorRunFunctionAttributes {
	return &NullableEdgeFirewallBehaviorRunFunctionAttributes{value: val, isSet: true}
}

func (v NullableEdgeFirewallBehaviorRunFunctionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallBehaviorRunFunctionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


