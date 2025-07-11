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

// checks if the PaginatedWorkloadDeploymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedWorkloadDeploymentList{}

// PaginatedWorkloadDeploymentList struct for PaginatedWorkloadDeploymentList
type PaginatedWorkloadDeploymentList struct {
	Count *int64 `json:"count,omitempty"`
	Results []WorkloadDeployment `json:"results,omitempty"`
}

// NewPaginatedWorkloadDeploymentList instantiates a new PaginatedWorkloadDeploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedWorkloadDeploymentList() *PaginatedWorkloadDeploymentList {
	this := PaginatedWorkloadDeploymentList{}
	return &this
}

// NewPaginatedWorkloadDeploymentListWithDefaults instantiates a new PaginatedWorkloadDeploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedWorkloadDeploymentListWithDefaults() *PaginatedWorkloadDeploymentList {
	this := PaginatedWorkloadDeploymentList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedWorkloadDeploymentList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedWorkloadDeploymentList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedWorkloadDeploymentList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedWorkloadDeploymentList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedWorkloadDeploymentList) GetResults() []WorkloadDeployment {
	if o == nil || IsNil(o.Results) {
		var ret []WorkloadDeployment
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedWorkloadDeploymentList) GetResultsOk() ([]WorkloadDeployment, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedWorkloadDeploymentList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WorkloadDeployment and assigns it to the Results field.
func (o *PaginatedWorkloadDeploymentList) SetResults(v []WorkloadDeployment) {
	o.Results = v
}

func (o PaginatedWorkloadDeploymentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedWorkloadDeploymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedWorkloadDeploymentList struct {
	value *PaginatedWorkloadDeploymentList
	isSet bool
}

func (v NullablePaginatedWorkloadDeploymentList) Get() *PaginatedWorkloadDeploymentList {
	return v.value
}

func (v *NullablePaginatedWorkloadDeploymentList) Set(val *PaginatedWorkloadDeploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedWorkloadDeploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedWorkloadDeploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedWorkloadDeploymentList(val *PaginatedWorkloadDeploymentList) *NullablePaginatedWorkloadDeploymentList {
	return &NullablePaginatedWorkloadDeploymentList{value: val, isSet: true}
}

func (v NullablePaginatedWorkloadDeploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedWorkloadDeploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


