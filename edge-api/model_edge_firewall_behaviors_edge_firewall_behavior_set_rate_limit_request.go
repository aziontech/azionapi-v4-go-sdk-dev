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

// checks if the EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest{}

// EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest struct for EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest
type EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest struct {
	// * `set_rate_limit` - set_rate_limit
	Type string `json:"type"`
	Attributes EdgeFirewallBehaviorSetRateLimitAttributesRequest `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest

// NewEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest instantiates a new EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest(type_ string, attributes EdgeFirewallBehaviorSetRateLimitAttributesRequest) *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest {
	this := EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequestWithDefaults instantiates a new EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequestWithDefaults() *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest {
	this := EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest{}
	return &this
}

// GetType returns the Type field value
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) GetAttributes() EdgeFirewallBehaviorSetRateLimitAttributesRequest {
	if o == nil {
		var ret EdgeFirewallBehaviorSetRateLimitAttributesRequest
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) GetAttributesOk() (*EdgeFirewallBehaviorSetRateLimitAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) SetAttributes(v EdgeFirewallBehaviorSetRateLimitAttributesRequest) {
	o.Attributes = v
}

func (o EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest := _EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest{}

	err = json.Unmarshal(data, &varEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest)

	if err != nil {
		return err
	}

	*o = EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest(varEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest struct {
	value *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest
	isSet bool
}

func (v NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) Get() *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest {
	return v.value
}

func (v *NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) Set(val *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest(val *EdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) *NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest {
	return &NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest{value: val, isSet: true}
}

func (v NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallBehaviorsEdgeFirewallBehaviorSetRateLimitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


