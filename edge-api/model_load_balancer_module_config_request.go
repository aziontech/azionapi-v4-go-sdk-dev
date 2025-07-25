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

// checks if the LoadBalancerModuleConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerModuleConfigRequest{}

// LoadBalancerModuleConfigRequest struct for LoadBalancerModuleConfigRequest
type LoadBalancerModuleConfigRequest struct {
	// Load balancing method to use (e.g., round-robin, least_conn).  * `round_robin` - Round Robin * `least_conn` - Least Connections * `ip_hash` - IP Hash
	Method *string `json:"method,omitempty"`
	// Maximum number of retry attempts on connection failure.
	MaxRetries *int64 `json:"max_retries,omitempty"`
	// Maximum time (in seconds) to wait for a connection to be established.
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	// Maximum time (in seconds) to wait for data read/write after connection.
	ReadWriteTimeout *int64 `json:"read_write_timeout,omitempty"`
}

// NewLoadBalancerModuleConfigRequest instantiates a new LoadBalancerModuleConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerModuleConfigRequest() *LoadBalancerModuleConfigRequest {
	this := LoadBalancerModuleConfigRequest{}
	return &this
}

// NewLoadBalancerModuleConfigRequestWithDefaults instantiates a new LoadBalancerModuleConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerModuleConfigRequestWithDefaults() *LoadBalancerModuleConfigRequest {
	this := LoadBalancerModuleConfigRequest{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoadBalancerModuleConfigRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerModuleConfigRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoadBalancerModuleConfigRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoadBalancerModuleConfigRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *LoadBalancerModuleConfigRequest) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int64
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerModuleConfigRequest) GetMaxRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *LoadBalancerModuleConfigRequest) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int64 and assigns it to the MaxRetries field.
func (o *LoadBalancerModuleConfigRequest) SetMaxRetries(v int64) {
	o.MaxRetries = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *LoadBalancerModuleConfigRequest) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerModuleConfigRequest) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *LoadBalancerModuleConfigRequest) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *LoadBalancerModuleConfigRequest) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetReadWriteTimeout returns the ReadWriteTimeout field value if set, zero value otherwise.
func (o *LoadBalancerModuleConfigRequest) GetReadWriteTimeout() int64 {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		var ret int64
		return ret
	}
	return *o.ReadWriteTimeout
}

// GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerModuleConfigRequest) GetReadWriteTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		return nil, false
	}
	return o.ReadWriteTimeout, true
}

// HasReadWriteTimeout returns a boolean if a field has been set.
func (o *LoadBalancerModuleConfigRequest) HasReadWriteTimeout() bool {
	if o != nil && !IsNil(o.ReadWriteTimeout) {
		return true
	}

	return false
}

// SetReadWriteTimeout gets a reference to the given int64 and assigns it to the ReadWriteTimeout field.
func (o *LoadBalancerModuleConfigRequest) SetReadWriteTimeout(v int64) {
	o.ReadWriteTimeout = &v
}

func (o LoadBalancerModuleConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerModuleConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MaxRetries) {
		toSerialize["max_retries"] = o.MaxRetries
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connection_timeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.ReadWriteTimeout) {
		toSerialize["read_write_timeout"] = o.ReadWriteTimeout
	}
	return toSerialize, nil
}

type NullableLoadBalancerModuleConfigRequest struct {
	value *LoadBalancerModuleConfigRequest
	isSet bool
}

func (v NullableLoadBalancerModuleConfigRequest) Get() *LoadBalancerModuleConfigRequest {
	return v.value
}

func (v *NullableLoadBalancerModuleConfigRequest) Set(val *LoadBalancerModuleConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerModuleConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerModuleConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerModuleConfigRequest(val *LoadBalancerModuleConfigRequest) *NullableLoadBalancerModuleConfigRequest {
	return &NullableLoadBalancerModuleConfigRequest{value: val, isSet: true}
}

func (v NullableLoadBalancerModuleConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerModuleConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


