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

// checks if the PaginatedResponseListDataSourceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseListDataSourceList{}

// PaginatedResponseListDataSourceList struct for PaginatedResponseListDataSourceList
type PaginatedResponseListDataSourceList struct {
	Count *int64 `json:"count,omitempty"`
	Results []ResponseListDataSource `json:"results,omitempty"`
}

// NewPaginatedResponseListDataSourceList instantiates a new PaginatedResponseListDataSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseListDataSourceList() *PaginatedResponseListDataSourceList {
	this := PaginatedResponseListDataSourceList{}
	return &this
}

// NewPaginatedResponseListDataSourceListWithDefaults instantiates a new PaginatedResponseListDataSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseListDataSourceListWithDefaults() *PaginatedResponseListDataSourceList {
	this := PaginatedResponseListDataSourceList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedResponseListDataSourceList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListDataSourceList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedResponseListDataSourceList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedResponseListDataSourceList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedResponseListDataSourceList) GetResults() []ResponseListDataSource {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseListDataSource
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListDataSourceList) GetResultsOk() ([]ResponseListDataSource, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedResponseListDataSourceList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseListDataSource and assigns it to the Results field.
func (o *PaginatedResponseListDataSourceList) SetResults(v []ResponseListDataSource) {
	o.Results = v
}

func (o PaginatedResponseListDataSourceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseListDataSourceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedResponseListDataSourceList struct {
	value *PaginatedResponseListDataSourceList
	isSet bool
}

func (v NullablePaginatedResponseListDataSourceList) Get() *PaginatedResponseListDataSourceList {
	return v.value
}

func (v *NullablePaginatedResponseListDataSourceList) Set(val *PaginatedResponseListDataSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseListDataSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseListDataSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseListDataSourceList(val *PaginatedResponseListDataSourceList) *NullablePaginatedResponseListDataSourceList {
	return &NullablePaginatedResponseListDataSourceList{value: val, isSet: true}
}

func (v NullablePaginatedResponseListDataSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseListDataSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


