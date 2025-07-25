/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Output{}

// Output struct for Output
type Output struct {
	// * `standard` - Standard HTTP/HTTPS POST * `kafka` - Apache Kafka * `s3` - Simple Storage Service (S3) * `big_query` - Google BigQuery * `elasticsearch` - Elasticsearch * `splunk` - Splunk * `aws_kinesis_firehose` - AWS Kinesis Data Firehose * `datadog` - Datadog * `qradar` - IBM QRadar * `azure_monitor` - Azure Monitor * `azure_blob_storage` - Azure Blob Storage
	Type string `json:"type"`
	Attributes OutputPolymorphic `json:"attributes"`
}

type _Output Output

// NewOutput instantiates a new Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutput(type_ string, attributes OutputPolymorphic) *Output {
	this := Output{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOutputWithDefaults instantiates a new Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputWithDefaults() *Output {
	this := Output{}
	return &this
}

// GetType returns the Type field value
func (o *Output) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Output) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Output) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *Output) GetAttributes() OutputPolymorphic {
	if o == nil {
		var ret OutputPolymorphic
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Output) GetAttributesOk() (*OutputPolymorphic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *Output) SetAttributes(v OutputPolymorphic) {
	o.Attributes = v
}

func (o Output) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *Output) UnmarshalJSON(data []byte) (err error) {
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

	varOutput := _Output{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutput)

	if err != nil {
		return err
	}

	*o = Output(varOutput)

	return err
}

type NullableOutput struct {
	value *Output
	isSet bool
}

func (v NullableOutput) Get() *Output {
	return v.value
}

func (v *NullableOutput) Set(val *Output) {
	v.value = val
	v.isSet = true
}

func (v NullableOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutput(val *Output) *NullableOutput {
	return &NullableOutput{value: val, isSet: true}
}

func (v NullableOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


