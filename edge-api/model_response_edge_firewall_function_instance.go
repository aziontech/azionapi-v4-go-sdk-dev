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

// checks if the ResponseEdgeFirewallFunctionInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEdgeFirewallFunctionInstance{}

// ResponseEdgeFirewallFunctionInstance struct for ResponseEdgeFirewallFunctionInstance
type ResponseEdgeFirewallFunctionInstance struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
	Data EdgeFirewallFunctionInstance `json:"data"`
}

type _ResponseEdgeFirewallFunctionInstance ResponseEdgeFirewallFunctionInstance

// NewResponseEdgeFirewallFunctionInstance instantiates a new ResponseEdgeFirewallFunctionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEdgeFirewallFunctionInstance(data EdgeFirewallFunctionInstance) *ResponseEdgeFirewallFunctionInstance {
	this := ResponseEdgeFirewallFunctionInstance{}
	this.Data = data
	return &this
}

// NewResponseEdgeFirewallFunctionInstanceWithDefaults instantiates a new ResponseEdgeFirewallFunctionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEdgeFirewallFunctionInstanceWithDefaults() *ResponseEdgeFirewallFunctionInstance {
	this := ResponseEdgeFirewallFunctionInstance{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseEdgeFirewallFunctionInstance) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFirewallFunctionInstance) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseEdgeFirewallFunctionInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseEdgeFirewallFunctionInstance) SetState(v string) {
	o.State = &v
}

// GetData returns the Data field value
func (o *ResponseEdgeFirewallFunctionInstance) GetData() EdgeFirewallFunctionInstance {
	if o == nil {
		var ret EdgeFirewallFunctionInstance
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFirewallFunctionInstance) GetDataOk() (*EdgeFirewallFunctionInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseEdgeFirewallFunctionInstance) SetData(v EdgeFirewallFunctionInstance) {
	o.Data = v
}

func (o ResponseEdgeFirewallFunctionInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEdgeFirewallFunctionInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseEdgeFirewallFunctionInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varResponseEdgeFirewallFunctionInstance := _ResponseEdgeFirewallFunctionInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseEdgeFirewallFunctionInstance)

	if err != nil {
		return err
	}

	*o = ResponseEdgeFirewallFunctionInstance(varResponseEdgeFirewallFunctionInstance)

	return err
}

type NullableResponseEdgeFirewallFunctionInstance struct {
	value *ResponseEdgeFirewallFunctionInstance
	isSet bool
}

func (v NullableResponseEdgeFirewallFunctionInstance) Get() *ResponseEdgeFirewallFunctionInstance {
	return v.value
}

func (v *NullableResponseEdgeFirewallFunctionInstance) Set(val *ResponseEdgeFirewallFunctionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEdgeFirewallFunctionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEdgeFirewallFunctionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEdgeFirewallFunctionInstance(val *ResponseEdgeFirewallFunctionInstance) *NullableResponseEdgeFirewallFunctionInstance {
	return &NullableResponseEdgeFirewallFunctionInstance{value: val, isSet: true}
}

func (v NullableResponseEdgeFirewallFunctionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEdgeFirewallFunctionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


