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

// checks if the EdgeFirewallBehaviorTagEventAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallBehaviorTagEventAttributes{}

// EdgeFirewallBehaviorTagEventAttributes struct for EdgeFirewallBehaviorTagEventAttributes
type EdgeFirewallBehaviorTagEventAttributes struct {
	Value string `json:"value" validate:"regexp=.*"`
	AdditionalProperties map[string]interface{}
}

type _EdgeFirewallBehaviorTagEventAttributes EdgeFirewallBehaviorTagEventAttributes

// NewEdgeFirewallBehaviorTagEventAttributes instantiates a new EdgeFirewallBehaviorTagEventAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallBehaviorTagEventAttributes(value string) *EdgeFirewallBehaviorTagEventAttributes {
	this := EdgeFirewallBehaviorTagEventAttributes{}
	this.Value = value
	return &this
}

// NewEdgeFirewallBehaviorTagEventAttributesWithDefaults instantiates a new EdgeFirewallBehaviorTagEventAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallBehaviorTagEventAttributesWithDefaults() *EdgeFirewallBehaviorTagEventAttributes {
	this := EdgeFirewallBehaviorTagEventAttributes{}
	return &this
}

// GetValue returns the Value field value
func (o *EdgeFirewallBehaviorTagEventAttributes) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallBehaviorTagEventAttributes) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EdgeFirewallBehaviorTagEventAttributes) SetValue(v string) {
	o.Value = v
}

func (o EdgeFirewallBehaviorTagEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallBehaviorTagEventAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeFirewallBehaviorTagEventAttributes) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeFirewallBehaviorTagEventAttributes := _EdgeFirewallBehaviorTagEventAttributes{}

	err = json.Unmarshal(data, &varEdgeFirewallBehaviorTagEventAttributes)

	if err != nil {
		return err
	}

	*o = EdgeFirewallBehaviorTagEventAttributes(varEdgeFirewallBehaviorTagEventAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeFirewallBehaviorTagEventAttributes struct {
	value *EdgeFirewallBehaviorTagEventAttributes
	isSet bool
}

func (v NullableEdgeFirewallBehaviorTagEventAttributes) Get() *EdgeFirewallBehaviorTagEventAttributes {
	return v.value
}

func (v *NullableEdgeFirewallBehaviorTagEventAttributes) Set(val *EdgeFirewallBehaviorTagEventAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallBehaviorTagEventAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallBehaviorTagEventAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallBehaviorTagEventAttributes(val *EdgeFirewallBehaviorTagEventAttributes) *NullableEdgeFirewallBehaviorTagEventAttributes {
	return &NullableEdgeFirewallBehaviorTagEventAttributes{value: val, isSet: true}
}

func (v NullableEdgeFirewallBehaviorTagEventAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallBehaviorTagEventAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


