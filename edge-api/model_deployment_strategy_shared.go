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

// checks if the DeploymentStrategyShared type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentStrategyShared{}

// DeploymentStrategyShared Deployment strategy.
type DeploymentStrategyShared struct {
	Type string `json:"type" validate:"regexp=.*"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentStrategyShared DeploymentStrategyShared

// NewDeploymentStrategyShared instantiates a new DeploymentStrategyShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStrategyShared(type_ string) *DeploymentStrategyShared {
	this := DeploymentStrategyShared{}
	this.Type = type_
	return &this
}

// NewDeploymentStrategySharedWithDefaults instantiates a new DeploymentStrategyShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStrategySharedWithDefaults() *DeploymentStrategyShared {
	this := DeploymentStrategyShared{}
	return &this
}

// GetType returns the Type field value
func (o *DeploymentStrategyShared) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeploymentStrategyShared) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeploymentStrategyShared) SetType(v string) {
	o.Type = v
}

func (o DeploymentStrategyShared) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentStrategyShared) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentStrategyShared) UnmarshalJSON(data []byte) (err error) {
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

	varDeploymentStrategyShared := _DeploymentStrategyShared{}

	err = json.Unmarshal(data, &varDeploymentStrategyShared)

	if err != nil {
		return err
	}

	*o = DeploymentStrategyShared(varDeploymentStrategyShared)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStrategyShared struct {
	value *DeploymentStrategyShared
	isSet bool
}

func (v NullableDeploymentStrategyShared) Get() *DeploymentStrategyShared {
	return v.value
}

func (v *NullableDeploymentStrategyShared) Set(val *DeploymentStrategyShared) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStrategyShared) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStrategyShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStrategyShared(val *DeploymentStrategyShared) *NullableDeploymentStrategyShared {
	return &NullableDeploymentStrategyShared{value: val, isSet: true}
}

func (v NullableDeploymentStrategyShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStrategyShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


