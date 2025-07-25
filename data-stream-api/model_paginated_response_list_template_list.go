/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
)

// checks if the PaginatedResponseListTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseListTemplateList{}

// PaginatedResponseListTemplateList struct for PaginatedResponseListTemplateList
type PaginatedResponseListTemplateList struct {
	Count *int64 `json:"count,omitempty"`
	Results []ResponseListTemplate `json:"results,omitempty"`
}

// NewPaginatedResponseListTemplateList instantiates a new PaginatedResponseListTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseListTemplateList() *PaginatedResponseListTemplateList {
	this := PaginatedResponseListTemplateList{}
	return &this
}

// NewPaginatedResponseListTemplateListWithDefaults instantiates a new PaginatedResponseListTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseListTemplateListWithDefaults() *PaginatedResponseListTemplateList {
	this := PaginatedResponseListTemplateList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedResponseListTemplateList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListTemplateList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedResponseListTemplateList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedResponseListTemplateList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedResponseListTemplateList) GetResults() []ResponseListTemplate {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseListTemplate
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListTemplateList) GetResultsOk() ([]ResponseListTemplate, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedResponseListTemplateList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseListTemplate and assigns it to the Results field.
func (o *PaginatedResponseListTemplateList) SetResults(v []ResponseListTemplate) {
	o.Results = v
}

func (o PaginatedResponseListTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseListTemplateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedResponseListTemplateList struct {
	value *PaginatedResponseListTemplateList
	isSet bool
}

func (v NullablePaginatedResponseListTemplateList) Get() *PaginatedResponseListTemplateList {
	return v.value
}

func (v *NullablePaginatedResponseListTemplateList) Set(val *PaginatedResponseListTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseListTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseListTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseListTemplateList(val *PaginatedResponseListTemplateList) *NullablePaginatedResponseListTemplateList {
	return &NullablePaginatedResponseListTemplateList{value: val, isSet: true}
}

func (v NullablePaginatedResponseListTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseListTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


