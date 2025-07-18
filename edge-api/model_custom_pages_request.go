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

// checks if the CustomPagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPagesRequest{}

// CustomPagesRequest struct for CustomPagesRequest
type CustomPagesRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	Active *bool `json:"active,omitempty"`
	Pages []ItemPageRequest `json:"pages"`
}

type _CustomPagesRequest CustomPagesRequest

// NewCustomPagesRequest instantiates a new CustomPagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPagesRequest(name string, pages []ItemPageRequest) *CustomPagesRequest {
	this := CustomPagesRequest{}
	this.Name = name
	this.Pages = pages
	return &this
}

// NewCustomPagesRequestWithDefaults instantiates a new CustomPagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPagesRequestWithDefaults() *CustomPagesRequest {
	this := CustomPagesRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CustomPagesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomPagesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomPagesRequest) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CustomPagesRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPagesRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CustomPagesRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CustomPagesRequest) SetActive(v bool) {
	o.Active = &v
}

// GetPages returns the Pages field value
func (o *CustomPagesRequest) GetPages() []ItemPageRequest {
	if o == nil {
		var ret []ItemPageRequest
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *CustomPagesRequest) GetPagesOk() ([]ItemPageRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *CustomPagesRequest) SetPages(v []ItemPageRequest) {
	o.Pages = v
}

func (o CustomPagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["pages"] = o.Pages
	return toSerialize, nil
}

func (o *CustomPagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pages",
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

	varCustomPagesRequest := _CustomPagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomPagesRequest)

	if err != nil {
		return err
	}

	*o = CustomPagesRequest(varCustomPagesRequest)

	return err
}

type NullableCustomPagesRequest struct {
	value *CustomPagesRequest
	isSet bool
}

func (v NullableCustomPagesRequest) Get() *CustomPagesRequest {
	return v.value
}

func (v *NullableCustomPagesRequest) Set(val *CustomPagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPagesRequest(val *CustomPagesRequest) *NullableCustomPagesRequest {
	return &NullableCustomPagesRequest{value: val, isSet: true}
}

func (v NullableCustomPagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


