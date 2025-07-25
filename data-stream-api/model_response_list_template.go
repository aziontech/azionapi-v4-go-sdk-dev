/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datastreamapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ResponseListTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListTemplate{}

// ResponseListTemplate struct for ResponseListTemplate
type ResponseListTemplate struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	Custom bool `json:"custom"`
	Active *bool `json:"active,omitempty"`
	DataSet string `json:"data_set" validate:"regexp=.*"`
}

type _ResponseListTemplate ResponseListTemplate

// NewResponseListTemplate instantiates a new ResponseListTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListTemplate(id int64, name string, lastEditor string, lastModified time.Time, custom bool, dataSet string) *ResponseListTemplate {
	this := ResponseListTemplate{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.Custom = custom
	this.DataSet = dataSet
	return &this
}

// NewResponseListTemplateWithDefaults instantiates a new ResponseListTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListTemplateWithDefaults() *ResponseListTemplate {
	this := ResponseListTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListTemplate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListTemplate) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListTemplate) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListTemplate) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListTemplate) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListTemplate) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListTemplate) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetCustom returns the Custom field value
func (o *ResponseListTemplate) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *ResponseListTemplate) SetCustom(v bool) {
	o.Custom = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListTemplate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListTemplate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListTemplate) SetActive(v bool) {
	o.Active = &v
}

// GetDataSet returns the DataSet field value
func (o *ResponseListTemplate) GetDataSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSet
}

// GetDataSetOk returns a tuple with the DataSet field value
// and a boolean to check if the value has been set.
func (o *ResponseListTemplate) GetDataSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSet, true
}

// SetDataSet sets field value
func (o *ResponseListTemplate) SetDataSet(v string) {
	o.DataSet = v
}

func (o ResponseListTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["custom"] = o.Custom
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["data_set"] = o.DataSet
	return toSerialize, nil
}

func (o *ResponseListTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"custom",
		"data_set",
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

	varResponseListTemplate := _ResponseListTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListTemplate)

	if err != nil {
		return err
	}

	*o = ResponseListTemplate(varResponseListTemplate)

	return err
}

type NullableResponseListTemplate struct {
	value *ResponseListTemplate
	isSet bool
}

func (v NullableResponseListTemplate) Get() *ResponseListTemplate {
	return v.value
}

func (v *NullableResponseListTemplate) Set(val *ResponseListTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListTemplate(val *ResponseListTemplate) *NullableResponseListTemplate {
	return &NullableResponseListTemplate{value: val, isSet: true}
}

func (v NullableResponseListTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


