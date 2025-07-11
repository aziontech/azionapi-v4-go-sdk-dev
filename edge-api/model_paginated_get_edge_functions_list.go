/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapi

import (
	"encoding/json"
)

// checks if the PaginatedGetEdgeFunctionsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedGetEdgeFunctionsList{}

// PaginatedGetEdgeFunctionsList struct for PaginatedGetEdgeFunctionsList
type PaginatedGetEdgeFunctionsList struct {
	Count *int64 `json:"count,omitempty"`
	Results []GetEdgeFunctions `json:"results,omitempty"`
}

// NewPaginatedGetEdgeFunctionsList instantiates a new PaginatedGetEdgeFunctionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGetEdgeFunctionsList() *PaginatedGetEdgeFunctionsList {
	this := PaginatedGetEdgeFunctionsList{}
	return &this
}

// NewPaginatedGetEdgeFunctionsListWithDefaults instantiates a new PaginatedGetEdgeFunctionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGetEdgeFunctionsListWithDefaults() *PaginatedGetEdgeFunctionsList {
	this := PaginatedGetEdgeFunctionsList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedGetEdgeFunctionsList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGetEdgeFunctionsList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedGetEdgeFunctionsList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedGetEdgeFunctionsList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedGetEdgeFunctionsList) GetResults() []GetEdgeFunctions {
	if o == nil || IsNil(o.Results) {
		var ret []GetEdgeFunctions
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGetEdgeFunctionsList) GetResultsOk() ([]GetEdgeFunctions, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedGetEdgeFunctionsList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GetEdgeFunctions and assigns it to the Results field.
func (o *PaginatedGetEdgeFunctionsList) SetResults(v []GetEdgeFunctions) {
	o.Results = v
}

func (o PaginatedGetEdgeFunctionsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedGetEdgeFunctionsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedGetEdgeFunctionsList struct {
	value *PaginatedGetEdgeFunctionsList
	isSet bool
}

func (v NullablePaginatedGetEdgeFunctionsList) Get() *PaginatedGetEdgeFunctionsList {
	return v.value
}

func (v *NullablePaginatedGetEdgeFunctionsList) Set(val *PaginatedGetEdgeFunctionsList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGetEdgeFunctionsList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGetEdgeFunctionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGetEdgeFunctionsList(val *PaginatedGetEdgeFunctionsList) *NullablePaginatedGetEdgeFunctionsList {
	return &NullablePaginatedGetEdgeFunctionsList{value: val, isSet: true}
}

func (v NullablePaginatedGetEdgeFunctionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGetEdgeFunctionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


