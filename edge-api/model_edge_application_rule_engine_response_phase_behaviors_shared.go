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

// checks if the EdgeApplicationRuleEngineResponsePhaseBehaviorsShared type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineResponsePhaseBehaviorsShared{}

// EdgeApplicationRuleEngineResponsePhaseBehaviorsShared Polymorphic serializer base class.  Note that the discriminator field must exist at the same depth as the mapped serializer fields for the OpenAPI introspection. See https://swagger.io/docs/specification/data-models/inheritance-and-polymorphism/ for more information. As such, it's not possible to define something like:  {     \"object_type\": \"foo\",     \"polymorphic_context\": {         <foo-specific fields>     } }  without explicitly wrapping this in a parent serializer, i.e. - ``polymorphic_context`` can not be a PolymorphicSerializer itself, as it requires access to the ``object_type`` in the parent scope.
type EdgeApplicationRuleEngineResponsePhaseBehaviorsShared struct {
	Type string `json:"type" validate:"regexp=.*"`
	AdditionalProperties map[string]interface{}
}

type _EdgeApplicationRuleEngineResponsePhaseBehaviorsShared EdgeApplicationRuleEngineResponsePhaseBehaviorsShared

// NewEdgeApplicationRuleEngineResponsePhaseBehaviorsShared instantiates a new EdgeApplicationRuleEngineResponsePhaseBehaviorsShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineResponsePhaseBehaviorsShared(type_ string) *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared {
	this := EdgeApplicationRuleEngineResponsePhaseBehaviorsShared{}
	this.Type = type_
	return &this
}

// NewEdgeApplicationRuleEngineResponsePhaseBehaviorsSharedWithDefaults instantiates a new EdgeApplicationRuleEngineResponsePhaseBehaviorsShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineResponsePhaseBehaviorsSharedWithDefaults() *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared {
	this := EdgeApplicationRuleEngineResponsePhaseBehaviorsShared{}
	return &this
}

// GetType returns the Type field value
func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) SetType(v string) {
	o.Type = v
}

func (o EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeApplicationRuleEngineResponsePhaseBehaviorsShared := _EdgeApplicationRuleEngineResponsePhaseBehaviorsShared{}

	err = json.Unmarshal(data, &varEdgeApplicationRuleEngineResponsePhaseBehaviorsShared)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineResponsePhaseBehaviorsShared(varEdgeApplicationRuleEngineResponsePhaseBehaviorsShared)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared struct {
	value *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) Get() *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) Set(val *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared(val *EdgeApplicationRuleEngineResponsePhaseBehaviorsShared) *NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared {
	return &NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineResponsePhaseBehaviorsShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


