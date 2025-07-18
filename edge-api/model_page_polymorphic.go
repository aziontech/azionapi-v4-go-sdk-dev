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

// PagePolymorphic - struct for PagePolymorphic
type PagePolymorphic struct {
	PageConnector *PageConnector
	PageDefault *PageDefault
}

// PageConnectorAsPagePolymorphic is a convenience function that returns PageConnector wrapped in PagePolymorphic
func PageConnectorAsPagePolymorphic(v *PageConnector) PagePolymorphic {
	return PagePolymorphic{
		PageConnector: v,
	}
}

// PageDefaultAsPagePolymorphic is a convenience function that returns PageDefault wrapped in PagePolymorphic
func PageDefaultAsPagePolymorphic(v *PageDefault) PagePolymorphic {
	return PagePolymorphic{
		PageDefault: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PagePolymorphic) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PageConnector
	err = newStrictDecoder(data).Decode(&dst.PageConnector)
	if err == nil {
		jsonPageConnector, _ := json.Marshal(dst.PageConnector)
		if string(jsonPageConnector) == "{}" { // empty struct
			dst.PageConnector = nil
		} else {
			if err = validator.Validate(dst.PageConnector); err != nil {
				dst.PageConnector = nil
			} else {
				match++
			}
		}
	} else {
		dst.PageConnector = nil
	}

	// try to unmarshal data into PageDefault
	err = newStrictDecoder(data).Decode(&dst.PageDefault)
	if err == nil {
		jsonPageDefault, _ := json.Marshal(dst.PageDefault)
		if string(jsonPageDefault) == "{}" { // empty struct
			dst.PageDefault = nil
		} else {
			if err = validator.Validate(dst.PageDefault); err != nil {
				dst.PageDefault = nil
			} else {
				match++
			}
		}
	} else {
		dst.PageDefault = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PageConnector = nil
		dst.PageDefault = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PagePolymorphic)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PagePolymorphic)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PagePolymorphic) MarshalJSON() ([]byte, error) {
	if src.PageConnector != nil {
		return json.Marshal(&src.PageConnector)
	}

	if src.PageDefault != nil {
		return json.Marshal(&src.PageDefault)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PagePolymorphic) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PageConnector != nil {
		return obj.PageConnector
	}

	if obj.PageDefault != nil {
		return obj.PageDefault
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj PagePolymorphic) GetActualInstanceValue() (interface{}) {
	if obj.PageConnector != nil {
		return *obj.PageConnector
	}

	if obj.PageDefault != nil {
		return *obj.PageDefault
	}

	// all schemas are nil
	return nil
}

type NullablePagePolymorphic struct {
	value *PagePolymorphic
	isSet bool
}

func (v NullablePagePolymorphic) Get() *PagePolymorphic {
	return v.value
}

func (v *NullablePagePolymorphic) Set(val *PagePolymorphic) {
	v.value = val
	v.isSet = true
}

func (v NullablePagePolymorphic) IsSet() bool {
	return v.isSet
}

func (v *NullablePagePolymorphic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagePolymorphic(val *PagePolymorphic) *NullablePagePolymorphic {
	return &NullablePagePolymorphic{value: val, isSet: true}
}

func (v NullablePagePolymorphic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagePolymorphic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


