/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CloneEdgeFirewallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneEdgeFirewallRequest{}

// CloneEdgeFirewallRequest struct for CloneEdgeFirewallRequest
type CloneEdgeFirewallRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
}

type _CloneEdgeFirewallRequest CloneEdgeFirewallRequest

// NewCloneEdgeFirewallRequest instantiates a new CloneEdgeFirewallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneEdgeFirewallRequest(name string) *CloneEdgeFirewallRequest {
	this := CloneEdgeFirewallRequest{}
	this.Name = name
	return &this
}

// NewCloneEdgeFirewallRequestWithDefaults instantiates a new CloneEdgeFirewallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneEdgeFirewallRequestWithDefaults() *CloneEdgeFirewallRequest {
	this := CloneEdgeFirewallRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneEdgeFirewallRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneEdgeFirewallRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneEdgeFirewallRequest) SetName(v string) {
	o.Name = v
}

func (o CloneEdgeFirewallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneEdgeFirewallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CloneEdgeFirewallRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCloneEdgeFirewallRequest := _CloneEdgeFirewallRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloneEdgeFirewallRequest)

	if err != nil {
		return err
	}

	*o = CloneEdgeFirewallRequest(varCloneEdgeFirewallRequest)

	return err
}

type NullableCloneEdgeFirewallRequest struct {
	value *CloneEdgeFirewallRequest
	isSet bool
}

func (v NullableCloneEdgeFirewallRequest) Get() *CloneEdgeFirewallRequest {
	return v.value
}

func (v *NullableCloneEdgeFirewallRequest) Set(val *CloneEdgeFirewallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneEdgeFirewallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneEdgeFirewallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneEdgeFirewallRequest(val *CloneEdgeFirewallRequest) *NullableCloneEdgeFirewallRequest {
	return &NullableCloneEdgeFirewallRequest{value: val, isSet: true}
}

func (v NullableCloneEdgeFirewallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneEdgeFirewallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


