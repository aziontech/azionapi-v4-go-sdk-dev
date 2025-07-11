/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EdgeFirewall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewall{}

// EdgeFirewall struct for EdgeFirewall
type EdgeFirewall struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	Modules *EdgeFirewallModules `json:"modules,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	Active *bool `json:"active,omitempty"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	ProductVersion string `json:"product_version" validate:"regexp=\\\\d+\\\\.\\\\d+"`
}

type _EdgeFirewall EdgeFirewall

// NewEdgeFirewall instantiates a new EdgeFirewall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewall(id int64, name string, lastEditor string, lastModified time.Time, productVersion string) *EdgeFirewall {
	this := EdgeFirewall{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	return &this
}

// NewEdgeFirewallWithDefaults instantiates a new EdgeFirewall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallWithDefaults() *EdgeFirewall {
	this := EdgeFirewall{}
	return &this
}

// GetId returns the Id field value
func (o *EdgeFirewall) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeFirewall) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EdgeFirewall) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeFirewall) SetName(v string) {
	o.Name = v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *EdgeFirewall) GetModules() EdgeFirewallModules {
	if o == nil || IsNil(o.Modules) {
		var ret EdgeFirewallModules
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetModulesOk() (*EdgeFirewallModules, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *EdgeFirewall) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given EdgeFirewallModules and assigns it to the Modules field.
func (o *EdgeFirewall) SetModules(v EdgeFirewallModules) {
	o.Modules = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *EdgeFirewall) GetDebug() bool {
	if o == nil || IsNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *EdgeFirewall) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *EdgeFirewall) SetDebug(v bool) {
	o.Debug = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeFirewall) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeFirewall) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeFirewall) SetActive(v bool) {
	o.Active = &v
}

// GetLastEditor returns the LastEditor field value
func (o *EdgeFirewall) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *EdgeFirewall) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *EdgeFirewall) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *EdgeFirewall) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetProductVersion returns the ProductVersion field value
func (o *EdgeFirewall) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewall) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *EdgeFirewall) SetProductVersion(v string) {
	o.ProductVersion = v
}

func (o EdgeFirewall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["product_version"] = o.ProductVersion
	return toSerialize, nil
}

func (o *EdgeFirewall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"product_version",
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

	varEdgeFirewall := _EdgeFirewall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeFirewall)

	if err != nil {
		return err
	}

	*o = EdgeFirewall(varEdgeFirewall)

	return err
}

type NullableEdgeFirewall struct {
	value *EdgeFirewall
	isSet bool
}

func (v NullableEdgeFirewall) Get() *EdgeFirewall {
	return v.value
}

func (v *NullableEdgeFirewall) Set(val *EdgeFirewall) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewall) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewall(val *EdgeFirewall) *NullableEdgeFirewall {
	return &NullableEdgeFirewall{value: val, isSet: true}
}

func (v NullableEdgeFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


