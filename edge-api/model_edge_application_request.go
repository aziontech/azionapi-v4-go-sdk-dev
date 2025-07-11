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

// checks if the EdgeApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRequest{}

// EdgeApplicationRequest struct for EdgeApplicationRequest
type EdgeApplicationRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	Modules *EdgeApplicationModulesRequest `json:"modules,omitempty"`
	Active *bool `json:"active,omitempty"`
	Debug *bool `json:"debug,omitempty"`
}

type _EdgeApplicationRequest EdgeApplicationRequest

// NewEdgeApplicationRequest instantiates a new EdgeApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRequest(name string) *EdgeApplicationRequest {
	this := EdgeApplicationRequest{}
	this.Name = name
	return &this
}

// NewEdgeApplicationRequestWithDefaults instantiates a new EdgeApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRequestWithDefaults() *EdgeApplicationRequest {
	this := EdgeApplicationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EdgeApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *EdgeApplicationRequest) GetModules() EdgeApplicationModulesRequest {
	if o == nil || IsNil(o.Modules) {
		var ret EdgeApplicationModulesRequest
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRequest) GetModulesOk() (*EdgeApplicationModulesRequest, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *EdgeApplicationRequest) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given EdgeApplicationModulesRequest and assigns it to the Modules field.
func (o *EdgeApplicationRequest) SetModules(v EdgeApplicationModulesRequest) {
	o.Modules = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeApplicationRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeApplicationRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeApplicationRequest) SetActive(v bool) {
	o.Active = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *EdgeApplicationRequest) GetDebug() bool {
	if o == nil || IsNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRequest) GetDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *EdgeApplicationRequest) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *EdgeApplicationRequest) SetDebug(v bool) {
	o.Debug = &v
}

func (o EdgeApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	return toSerialize, nil
}

func (o *EdgeApplicationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEdgeApplicationRequest := _EdgeApplicationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeApplicationRequest)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRequest(varEdgeApplicationRequest)

	return err
}

type NullableEdgeApplicationRequest struct {
	value *EdgeApplicationRequest
	isSet bool
}

func (v NullableEdgeApplicationRequest) Get() *EdgeApplicationRequest {
	return v.value
}

func (v *NullableEdgeApplicationRequest) Set(val *EdgeApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRequest(val *EdgeApplicationRequest) *NullableEdgeApplicationRequest {
	return &NullableEdgeApplicationRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


