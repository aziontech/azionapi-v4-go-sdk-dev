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

// checks if the EdgeConnectorS3TypeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeConnectorS3TypeProperties{}

// EdgeConnectorS3TypeProperties struct for EdgeConnectorS3TypeProperties
type EdgeConnectorS3TypeProperties struct {
	Host NullableString `json:"host" validate:"regexp=.*"`
	Bucket string `json:"bucket" validate:"regexp=.*"`
	Path string `json:"path" validate:"regexp=.*"`
	Region string `json:"region" validate:"regexp=.*"`
	AccessKey string `json:"access_key" validate:"regexp=.*"`
	SecretKey string `json:"secret_key" validate:"regexp=.*"`
}

type _EdgeConnectorS3TypeProperties EdgeConnectorS3TypeProperties

// NewEdgeConnectorS3TypeProperties instantiates a new EdgeConnectorS3TypeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeConnectorS3TypeProperties(host NullableString, bucket string, path string, region string, accessKey string, secretKey string) *EdgeConnectorS3TypeProperties {
	this := EdgeConnectorS3TypeProperties{}
	this.Host = host
	this.Bucket = bucket
	this.Path = path
	this.Region = region
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	return &this
}

// NewEdgeConnectorS3TypePropertiesWithDefaults instantiates a new EdgeConnectorS3TypeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeConnectorS3TypePropertiesWithDefaults() *EdgeConnectorS3TypeProperties {
	this := EdgeConnectorS3TypeProperties{}
	return &this
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EdgeConnectorS3TypeProperties) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EdgeConnectorS3TypeProperties) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *EdgeConnectorS3TypeProperties) SetHost(v string) {
	o.Host.Set(&v)
}

// GetBucket returns the Bucket field value
func (o *EdgeConnectorS3TypeProperties) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorS3TypeProperties) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *EdgeConnectorS3TypeProperties) SetBucket(v string) {
	o.Bucket = v
}

// GetPath returns the Path field value
func (o *EdgeConnectorS3TypeProperties) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorS3TypeProperties) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *EdgeConnectorS3TypeProperties) SetPath(v string) {
	o.Path = v
}

// GetRegion returns the Region field value
func (o *EdgeConnectorS3TypeProperties) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorS3TypeProperties) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *EdgeConnectorS3TypeProperties) SetRegion(v string) {
	o.Region = v
}

// GetAccessKey returns the AccessKey field value
func (o *EdgeConnectorS3TypeProperties) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorS3TypeProperties) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *EdgeConnectorS3TypeProperties) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *EdgeConnectorS3TypeProperties) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorS3TypeProperties) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *EdgeConnectorS3TypeProperties) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o EdgeConnectorS3TypeProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeConnectorS3TypeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host.Get()
	toSerialize["bucket"] = o.Bucket
	toSerialize["path"] = o.Path
	toSerialize["region"] = o.Region
	toSerialize["access_key"] = o.AccessKey
	toSerialize["secret_key"] = o.SecretKey
	return toSerialize, nil
}

func (o *EdgeConnectorS3TypeProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"bucket",
		"path",
		"region",
		"access_key",
		"secret_key",
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

	varEdgeConnectorS3TypeProperties := _EdgeConnectorS3TypeProperties{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeConnectorS3TypeProperties)

	if err != nil {
		return err
	}

	*o = EdgeConnectorS3TypeProperties(varEdgeConnectorS3TypeProperties)

	return err
}

type NullableEdgeConnectorS3TypeProperties struct {
	value *EdgeConnectorS3TypeProperties
	isSet bool
}

func (v NullableEdgeConnectorS3TypeProperties) Get() *EdgeConnectorS3TypeProperties {
	return v.value
}

func (v *NullableEdgeConnectorS3TypeProperties) Set(val *EdgeConnectorS3TypeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeConnectorS3TypeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeConnectorS3TypeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeConnectorS3TypeProperties(val *EdgeConnectorS3TypeProperties) *NullableEdgeConnectorS3TypeProperties {
	return &NullableEdgeConnectorS3TypeProperties{value: val, isSet: true}
}

func (v NullableEdgeConnectorS3TypeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeConnectorS3TypeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


