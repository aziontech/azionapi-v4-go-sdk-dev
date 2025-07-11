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

// checks if the EdgeConnectorLiveIngestAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeConnectorLiveIngestAttributes{}

// EdgeConnectorLiveIngestAttributes struct for EdgeConnectorLiveIngestAttributes
type EdgeConnectorLiveIngestAttributes struct {
	// * `us-east-1` - us-east-1 * `us-east-2` - us-east-2 * `br-east-1` - br-east-1 * `br-east-2` - br-east-2 * `br-east-3` - br-east-3
	Region string `json:"region"`
}

type _EdgeConnectorLiveIngestAttributes EdgeConnectorLiveIngestAttributes

// NewEdgeConnectorLiveIngestAttributes instantiates a new EdgeConnectorLiveIngestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeConnectorLiveIngestAttributes(region string) *EdgeConnectorLiveIngestAttributes {
	this := EdgeConnectorLiveIngestAttributes{}
	this.Region = region
	return &this
}

// NewEdgeConnectorLiveIngestAttributesWithDefaults instantiates a new EdgeConnectorLiveIngestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeConnectorLiveIngestAttributesWithDefaults() *EdgeConnectorLiveIngestAttributes {
	this := EdgeConnectorLiveIngestAttributes{}
	return &this
}

// GetRegion returns the Region field value
func (o *EdgeConnectorLiveIngestAttributes) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorLiveIngestAttributes) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *EdgeConnectorLiveIngestAttributes) SetRegion(v string) {
	o.Region = v
}

func (o EdgeConnectorLiveIngestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeConnectorLiveIngestAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

func (o *EdgeConnectorLiveIngestAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
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

	varEdgeConnectorLiveIngestAttributes := _EdgeConnectorLiveIngestAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeConnectorLiveIngestAttributes)

	if err != nil {
		return err
	}

	*o = EdgeConnectorLiveIngestAttributes(varEdgeConnectorLiveIngestAttributes)

	return err
}

type NullableEdgeConnectorLiveIngestAttributes struct {
	value *EdgeConnectorLiveIngestAttributes
	isSet bool
}

func (v NullableEdgeConnectorLiveIngestAttributes) Get() *EdgeConnectorLiveIngestAttributes {
	return v.value
}

func (v *NullableEdgeConnectorLiveIngestAttributes) Set(val *EdgeConnectorLiveIngestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeConnectorLiveIngestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeConnectorLiveIngestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeConnectorLiveIngestAttributes(val *EdgeConnectorLiveIngestAttributes) *NullableEdgeConnectorLiveIngestAttributes {
	return &NullableEdgeConnectorLiveIngestAttributes{value: val, isSet: true}
}

func (v NullableEdgeConnectorLiveIngestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeConnectorLiveIngestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


