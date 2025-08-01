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

// checks if the EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest{}

// EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest struct for EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest
type EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest struct {
	// * `capture_match_groups` - capture_match_groups
	Type string `json:"type"`
	Attributes EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest

// NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest instantiates a new EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest(type_ string, attributes EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest) *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest {
	this := EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequestWithDefaults() *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest {
	this := EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest{}
	return &this
}

// GetType returns the Type field value
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) GetAttributes() EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest {
	if o == nil {
		var ret EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) GetAttributesOk() (*EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) SetAttributes(v EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest) {
	o.Attributes = v
}

func (o EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest := _EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest{}

	err = json.Unmarshal(data, &varEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest(varEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest struct {
	value *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) Get() *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) Set(val *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest(val *EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) *NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest {
	return &NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


