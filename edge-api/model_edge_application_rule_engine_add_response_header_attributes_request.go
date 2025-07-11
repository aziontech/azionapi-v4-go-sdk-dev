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

// checks if the EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest{}

// EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest struct for EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest
type EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest struct {
	Value string `json:"value" validate:"regexp=.*"`
	AdditionalProperties map[string]interface{}
}

type _EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest

// NewEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest instantiates a new EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest(value string) *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest {
	this := EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest{}
	this.Value = value
	return &this
}

// NewEdgeApplicationRuleEngineAddResponseHeaderAttributesRequestWithDefaults instantiates a new EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineAddResponseHeaderAttributesRequestWithDefaults() *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest {
	this := EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) SetValue(v string) {
	o.Value = v
}

func (o EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest := _EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest{}

	err = json.Unmarshal(data, &varEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest(varEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest struct {
	value *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) Get() *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) Set(val *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest(val *EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) *NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest {
	return &NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineAddResponseHeaderAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


