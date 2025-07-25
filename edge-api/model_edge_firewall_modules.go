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

// checks if the EdgeFirewallModules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallModules{}

// EdgeFirewallModules struct for EdgeFirewallModules
type EdgeFirewallModules struct {
	DdosProtection EdgeFirewallModule `json:"ddos_protection"`
	EdgeFunctions *EdgeFirewallModule `json:"edge_functions,omitempty"`
	NetworkProtection *EdgeFirewallModule `json:"network_protection,omitempty"`
	Waf *EdgeFirewallModule `json:"waf,omitempty"`
}

type _EdgeFirewallModules EdgeFirewallModules

// NewEdgeFirewallModules instantiates a new EdgeFirewallModules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallModules(ddosProtection EdgeFirewallModule) *EdgeFirewallModules {
	this := EdgeFirewallModules{}
	this.DdosProtection = ddosProtection
	return &this
}

// NewEdgeFirewallModulesWithDefaults instantiates a new EdgeFirewallModules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallModulesWithDefaults() *EdgeFirewallModules {
	this := EdgeFirewallModules{}
	return &this
}

// GetDdosProtection returns the DdosProtection field value
func (o *EdgeFirewallModules) GetDdosProtection() EdgeFirewallModule {
	if o == nil {
		var ret EdgeFirewallModule
		return ret
	}

	return o.DdosProtection
}

// GetDdosProtectionOk returns a tuple with the DdosProtection field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModules) GetDdosProtectionOk() (*EdgeFirewallModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdosProtection, true
}

// SetDdosProtection sets field value
func (o *EdgeFirewallModules) SetDdosProtection(v EdgeFirewallModule) {
	o.DdosProtection = v
}

// GetEdgeFunctions returns the EdgeFunctions field value if set, zero value otherwise.
func (o *EdgeFirewallModules) GetEdgeFunctions() EdgeFirewallModule {
	if o == nil || IsNil(o.EdgeFunctions) {
		var ret EdgeFirewallModule
		return ret
	}
	return *o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModules) GetEdgeFunctionsOk() (*EdgeFirewallModule, bool) {
	if o == nil || IsNil(o.EdgeFunctions) {
		return nil, false
	}
	return o.EdgeFunctions, true
}

// HasEdgeFunctions returns a boolean if a field has been set.
func (o *EdgeFirewallModules) HasEdgeFunctions() bool {
	if o != nil && !IsNil(o.EdgeFunctions) {
		return true
	}

	return false
}

// SetEdgeFunctions gets a reference to the given EdgeFirewallModule and assigns it to the EdgeFunctions field.
func (o *EdgeFirewallModules) SetEdgeFunctions(v EdgeFirewallModule) {
	o.EdgeFunctions = &v
}

// GetNetworkProtection returns the NetworkProtection field value if set, zero value otherwise.
func (o *EdgeFirewallModules) GetNetworkProtection() EdgeFirewallModule {
	if o == nil || IsNil(o.NetworkProtection) {
		var ret EdgeFirewallModule
		return ret
	}
	return *o.NetworkProtection
}

// GetNetworkProtectionOk returns a tuple with the NetworkProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModules) GetNetworkProtectionOk() (*EdgeFirewallModule, bool) {
	if o == nil || IsNil(o.NetworkProtection) {
		return nil, false
	}
	return o.NetworkProtection, true
}

// HasNetworkProtection returns a boolean if a field has been set.
func (o *EdgeFirewallModules) HasNetworkProtection() bool {
	if o != nil && !IsNil(o.NetworkProtection) {
		return true
	}

	return false
}

// SetNetworkProtection gets a reference to the given EdgeFirewallModule and assigns it to the NetworkProtection field.
func (o *EdgeFirewallModules) SetNetworkProtection(v EdgeFirewallModule) {
	o.NetworkProtection = &v
}

// GetWaf returns the Waf field value if set, zero value otherwise.
func (o *EdgeFirewallModules) GetWaf() EdgeFirewallModule {
	if o == nil || IsNil(o.Waf) {
		var ret EdgeFirewallModule
		return ret
	}
	return *o.Waf
}

// GetWafOk returns a tuple with the Waf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModules) GetWafOk() (*EdgeFirewallModule, bool) {
	if o == nil || IsNil(o.Waf) {
		return nil, false
	}
	return o.Waf, true
}

// HasWaf returns a boolean if a field has been set.
func (o *EdgeFirewallModules) HasWaf() bool {
	if o != nil && !IsNil(o.Waf) {
		return true
	}

	return false
}

// SetWaf gets a reference to the given EdgeFirewallModule and assigns it to the Waf field.
func (o *EdgeFirewallModules) SetWaf(v EdgeFirewallModule) {
	o.Waf = &v
}

func (o EdgeFirewallModules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallModules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ddos_protection"] = o.DdosProtection
	if !IsNil(o.EdgeFunctions) {
		toSerialize["edge_functions"] = o.EdgeFunctions
	}
	if !IsNil(o.NetworkProtection) {
		toSerialize["network_protection"] = o.NetworkProtection
	}
	if !IsNil(o.Waf) {
		toSerialize["waf"] = o.Waf
	}
	return toSerialize, nil
}

func (o *EdgeFirewallModules) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ddos_protection",
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

	varEdgeFirewallModules := _EdgeFirewallModules{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeFirewallModules)

	if err != nil {
		return err
	}

	*o = EdgeFirewallModules(varEdgeFirewallModules)

	return err
}

type NullableEdgeFirewallModules struct {
	value *EdgeFirewallModules
	isSet bool
}

func (v NullableEdgeFirewallModules) Get() *EdgeFirewallModules {
	return v.value
}

func (v *NullableEdgeFirewallModules) Set(val *EdgeFirewallModules) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallModules) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallModules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallModules(val *EdgeFirewallModules) *NullableEdgeFirewallModules {
	return &NullableEdgeFirewallModules{value: val, isSet: true}
}

func (v NullableEdgeFirewallModules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallModules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


