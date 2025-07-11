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

// checks if the EdgeApplicationRuleEngineCaptureMatchGroupsAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineCaptureMatchGroupsAttributes{}

// EdgeApplicationRuleEngineCaptureMatchGroupsAttributes struct for EdgeApplicationRuleEngineCaptureMatchGroupsAttributes
type EdgeApplicationRuleEngineCaptureMatchGroupsAttributes struct {
	Subject string `json:"subject" validate:"regexp=.*"`
	Regex string `json:"regex" validate:"regexp=.*"`
	CapturedArray string `json:"captured_array" validate:"regexp=.*"`
	AdditionalProperties map[string]interface{}
}

type _EdgeApplicationRuleEngineCaptureMatchGroupsAttributes EdgeApplicationRuleEngineCaptureMatchGroupsAttributes

// NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributes instantiates a new EdgeApplicationRuleEngineCaptureMatchGroupsAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributes(subject string, regex string, capturedArray string) *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes {
	this := EdgeApplicationRuleEngineCaptureMatchGroupsAttributes{}
	this.Subject = subject
	this.Regex = regex
	this.CapturedArray = capturedArray
	return &this
}

// NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributesWithDefaults instantiates a new EdgeApplicationRuleEngineCaptureMatchGroupsAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributesWithDefaults() *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes {
	this := EdgeApplicationRuleEngineCaptureMatchGroupsAttributes{}
	return &this
}

// GetSubject returns the Subject field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) SetSubject(v string) {
	o.Subject = v
}

// GetRegex returns the Regex field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) SetRegex(v string) {
	o.Regex = v
}

// GetCapturedArray returns the CapturedArray field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetCapturedArray() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CapturedArray
}

// GetCapturedArrayOk returns a tuple with the CapturedArray field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) GetCapturedArrayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedArray, true
}

// SetCapturedArray sets field value
func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) SetCapturedArray(v string) {
	o.CapturedArray = v
}

func (o EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subject"] = o.Subject
	toSerialize["regex"] = o.Regex
	toSerialize["captured_array"] = o.CapturedArray

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subject",
		"regex",
		"captured_array",
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

	varEdgeApplicationRuleEngineCaptureMatchGroupsAttributes := _EdgeApplicationRuleEngineCaptureMatchGroupsAttributes{}

	err = json.Unmarshal(data, &varEdgeApplicationRuleEngineCaptureMatchGroupsAttributes)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineCaptureMatchGroupsAttributes(varEdgeApplicationRuleEngineCaptureMatchGroupsAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subject")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "captured_array")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes struct {
	value *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) Get() *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) Set(val *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes(val *EdgeApplicationRuleEngineCaptureMatchGroupsAttributes) *NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes {
	return &NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineCaptureMatchGroupsAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


