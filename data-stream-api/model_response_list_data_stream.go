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

// checks if the ResponseListDataStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListDataStream{}

// ResponseListDataStream struct for ResponseListDataStream
type ResponseListDataStream struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	ProductVersion string `json:"product_version" validate:"regexp=.*"`
	Active *bool `json:"active,omitempty"`
	Inputs []InputPolymorphicInputDataSourceAttributes `json:"inputs"`
	Transform []TransformPolymorphic `json:"transform"`
	Outputs []Output `json:"outputs"`
}

type _ResponseListDataStream ResponseListDataStream

// NewResponseListDataStream instantiates a new ResponseListDataStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListDataStream(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, inputs []InputPolymorphicInputDataSourceAttributes, transform []TransformPolymorphic, outputs []Output) *ResponseListDataStream {
	this := ResponseListDataStream{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	this.Inputs = inputs
	this.Transform = transform
	this.Outputs = outputs
	return &this
}

// NewResponseListDataStreamWithDefaults instantiates a new ResponseListDataStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListDataStreamWithDefaults() *ResponseListDataStream {
	this := ResponseListDataStream{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListDataStream) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListDataStream) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListDataStream) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListDataStream) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListDataStream) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListDataStream) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListDataStream) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListDataStream) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetProductVersion returns the ProductVersion field value
func (o *ResponseListDataStream) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *ResponseListDataStream) SetProductVersion(v string) {
	o.ProductVersion = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListDataStream) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListDataStream) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListDataStream) SetActive(v bool) {
	o.Active = &v
}

// GetInputs returns the Inputs field value
func (o *ResponseListDataStream) GetInputs() []InputPolymorphicInputDataSourceAttributes {
	if o == nil {
		var ret []InputPolymorphicInputDataSourceAttributes
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetInputsOk() ([]InputPolymorphicInputDataSourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *ResponseListDataStream) SetInputs(v []InputPolymorphicInputDataSourceAttributes) {
	o.Inputs = v
}

// GetTransform returns the Transform field value
func (o *ResponseListDataStream) GetTransform() []TransformPolymorphic {
	if o == nil {
		var ret []TransformPolymorphic
		return ret
	}

	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetTransformOk() ([]TransformPolymorphic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transform, true
}

// SetTransform sets field value
func (o *ResponseListDataStream) SetTransform(v []TransformPolymorphic) {
	o.Transform = v
}

// GetOutputs returns the Outputs field value
func (o *ResponseListDataStream) GetOutputs() []Output {
	if o == nil {
		var ret []Output
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *ResponseListDataStream) GetOutputsOk() ([]Output, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *ResponseListDataStream) SetOutputs(v []Output) {
	o.Outputs = v
}

func (o ResponseListDataStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListDataStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["product_version"] = o.ProductVersion
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["inputs"] = o.Inputs
	toSerialize["transform"] = o.Transform
	toSerialize["outputs"] = o.Outputs
	return toSerialize, nil
}

func (o *ResponseListDataStream) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"product_version",
		"inputs",
		"transform",
		"outputs",
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

	varResponseListDataStream := _ResponseListDataStream{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListDataStream)

	if err != nil {
		return err
	}

	*o = ResponseListDataStream(varResponseListDataStream)

	return err
}

type NullableResponseListDataStream struct {
	value *ResponseListDataStream
	isSet bool
}

func (v NullableResponseListDataStream) Get() *ResponseListDataStream {
	return v.value
}

func (v *NullableResponseListDataStream) Set(val *ResponseListDataStream) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListDataStream) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListDataStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListDataStream(val *ResponseListDataStream) *NullableResponseListDataStream {
	return &NullableResponseListDataStream{value: val, isSet: true}
}

func (v NullableResponseListDataStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListDataStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


