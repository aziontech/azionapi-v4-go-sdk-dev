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
	"gopkg.in/validator.v2"
)

// EdgeFirewallCriterionPolymorphicArgument - struct for EdgeFirewallCriterionPolymorphicArgument
type EdgeFirewallCriterionPolymorphicArgument struct {
	Int64 *int64
	String *string
}

// int64AsEdgeFirewallCriterionPolymorphicArgument is a convenience function that returns int64 wrapped in EdgeFirewallCriterionPolymorphicArgument
func Int64AsEdgeFirewallCriterionPolymorphicArgument(v *int64) EdgeFirewallCriterionPolymorphicArgument {
	return EdgeFirewallCriterionPolymorphicArgument{
		Int64: v,
	}
}

// stringAsEdgeFirewallCriterionPolymorphicArgument is a convenience function that returns string wrapped in EdgeFirewallCriterionPolymorphicArgument
func StringAsEdgeFirewallCriterionPolymorphicArgument(v *string) EdgeFirewallCriterionPolymorphicArgument {
	return EdgeFirewallCriterionPolymorphicArgument{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EdgeFirewallCriterionPolymorphicArgument) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			if err = validator.Validate(dst.Int64); err != nil {
				dst.Int64 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EdgeFirewallCriterionPolymorphicArgument)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EdgeFirewallCriterionPolymorphicArgument)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EdgeFirewallCriterionPolymorphicArgument) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EdgeFirewallCriterionPolymorphicArgument) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EdgeFirewallCriterionPolymorphicArgument) GetActualInstanceValue() (interface{}) {
	if obj.Int64 != nil {
		return *obj.Int64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableEdgeFirewallCriterionPolymorphicArgument struct {
	value *EdgeFirewallCriterionPolymorphicArgument
	isSet bool
}

func (v NullableEdgeFirewallCriterionPolymorphicArgument) Get() *EdgeFirewallCriterionPolymorphicArgument {
	return v.value
}

func (v *NullableEdgeFirewallCriterionPolymorphicArgument) Set(val *EdgeFirewallCriterionPolymorphicArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallCriterionPolymorphicArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallCriterionPolymorphicArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallCriterionPolymorphicArgument(val *EdgeFirewallCriterionPolymorphicArgument) *NullableEdgeFirewallCriterionPolymorphicArgument {
	return &NullableEdgeFirewallCriterionPolymorphicArgument{value: val, isSet: true}
}

func (v NullableEdgeFirewallCriterionPolymorphicArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallCriterionPolymorphicArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


