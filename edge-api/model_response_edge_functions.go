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

// checks if the ResponseEdgeFunctions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEdgeFunctions{}

// ResponseEdgeFunctions struct for ResponseEdgeFunctions
type ResponseEdgeFunctions struct {
	State *string `json:"state,omitempty" validate:"regexp=.*"`
	Data EdgeFunctions `json:"data"`
}

type _ResponseEdgeFunctions ResponseEdgeFunctions

// NewResponseEdgeFunctions instantiates a new ResponseEdgeFunctions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEdgeFunctions(data EdgeFunctions) *ResponseEdgeFunctions {
	this := ResponseEdgeFunctions{}
	this.Data = data
	return &this
}

// NewResponseEdgeFunctionsWithDefaults instantiates a new ResponseEdgeFunctions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEdgeFunctionsWithDefaults() *ResponseEdgeFunctions {
	this := ResponseEdgeFunctions{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResponseEdgeFunctions) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFunctions) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResponseEdgeFunctions) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ResponseEdgeFunctions) SetState(v string) {
	o.State = &v
}

// GetData returns the Data field value
func (o *ResponseEdgeFunctions) GetData() EdgeFunctions {
	if o == nil {
		var ret EdgeFunctions
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFunctions) GetDataOk() (*EdgeFunctions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseEdgeFunctions) SetData(v EdgeFunctions) {
	o.Data = v
}

func (o ResponseEdgeFunctions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEdgeFunctions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseEdgeFunctions) UnmarshalJSON(data []byte) (err error) {
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

	varResponseEdgeFunctions := _ResponseEdgeFunctions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseEdgeFunctions)

	if err != nil {
		return err
	}

	*o = ResponseEdgeFunctions(varResponseEdgeFunctions)

	return err
}

type NullableResponseEdgeFunctions struct {
	value *ResponseEdgeFunctions
	isSet bool
}

func (v NullableResponseEdgeFunctions) Get() *ResponseEdgeFunctions {
	return v.value
}

func (v *NullableResponseEdgeFunctions) Set(val *ResponseEdgeFunctions) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEdgeFunctions) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEdgeFunctions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEdgeFunctions(val *ResponseEdgeFunctions) *NullableResponseEdgeFunctions {
	return &NullableResponseEdgeFunctions{value: val, isSet: true}
}

func (v NullableResponseEdgeFunctions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEdgeFunctions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


