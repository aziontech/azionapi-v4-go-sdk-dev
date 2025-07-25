/*
edge-storage-api

REST API OpenAPI documentation for the Edge Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storageapi

import (
	"encoding/json"
)

// checks if the ResponseBadRequestCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestCredential{}

// ResponseBadRequestCredential struct for ResponseBadRequestCredential
type ResponseBadRequestCredential struct {
	Name []string `json:"name,omitempty"`
	ExpirationDate []string `json:"expiration_date,omitempty"`
	LastModified []string `json:"last_modified,omitempty"`
	Capabilities *ResponseBadRequestCredentialCapabilities `json:"capabilities,omitempty"`
	Bucket []string `json:"bucket,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestCredential instantiates a new ResponseBadRequestCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestCredential() *ResponseBadRequestCredential {
	this := ResponseBadRequestCredential{}
	return &this
}

// NewResponseBadRequestCredentialWithDefaults instantiates a new ResponseBadRequestCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestCredentialWithDefaults() *ResponseBadRequestCredential {
	this := ResponseBadRequestCredential{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestCredential) SetName(v []string) {
	o.Name = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetExpirationDate() []string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret []string
		return ret
	}
	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetExpirationDateOk() ([]string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given []string and assigns it to the ExpirationDate field.
func (o *ResponseBadRequestCredential) SetExpirationDate(v []string) {
	o.ExpirationDate = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetLastModified() []string {
	if o == nil || IsNil(o.LastModified) {
		var ret []string
		return ret
	}
	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetLastModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given []string and assigns it to the LastModified field.
func (o *ResponseBadRequestCredential) SetLastModified(v []string) {
	o.LastModified = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetCapabilities() ResponseBadRequestCredentialCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret ResponseBadRequestCredentialCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetCapabilitiesOk() (*ResponseBadRequestCredentialCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given ResponseBadRequestCredentialCapabilities and assigns it to the Capabilities field.
func (o *ResponseBadRequestCredential) SetCapabilities(v ResponseBadRequestCredentialCapabilities) {
	o.Capabilities = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetBucket() []string {
	if o == nil || IsNil(o.Bucket) {
		var ret []string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetBucketOk() ([]string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given []string and assigns it to the Bucket field.
func (o *ResponseBadRequestCredential) SetBucket(v []string) {
	o.Bucket = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestCredential) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCredential) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestCredential) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestCredential) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestCredential struct {
	value *ResponseBadRequestCredential
	isSet bool
}

func (v NullableResponseBadRequestCredential) Get() *ResponseBadRequestCredential {
	return v.value
}

func (v *NullableResponseBadRequestCredential) Set(val *ResponseBadRequestCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestCredential(val *ResponseBadRequestCredential) *NullableResponseBadRequestCredential {
	return &NullableResponseBadRequestCredential{value: val, isSet: true}
}

func (v NullableResponseBadRequestCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


