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

// checks if the Page type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Page{}

// Page struct for Page
type Page struct {
	// * `default` - default * `400` - Bad Request * `401` - Unauthorized * `403` - Forbidden * `404` - Not Found * `405` - Method Not Allowed * `406` - Not Acceptable * `408` - Request Timeout * `409` - Conflict * `410` - Gone * `411` - Length Required * `414` - URI Too Long * `415` - Unsupported Media Type * `416` - Range Not Satisfiable * `426` - Upgrade Required * `429` - Too Many Requests * `431` - Request Header Fields Too Large * `500` - Internal Server Error * `501` - Not Implemented * `502` - Bad Gateway * `503` - Service Unavailable * `504` - Gateway Timeout * `505` - HTTP Version Not Supported
	Code string `json:"code"`
	Page PageConnector `json:"page"`
}

type _Page Page

// NewPage instantiates a new Page object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPage(code string, page PageConnector) *Page {
	this := Page{}
	this.Code = code
	this.Page = page
	return &this
}

// NewPageWithDefaults instantiates a new Page object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDefaults() *Page {
	this := Page{}
	return &this
}

// GetCode returns the Code field value
func (o *Page) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Page) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Page) SetCode(v string) {
	o.Code = v
}

// GetPage returns the Page field value
func (o *Page) GetPage() PageConnector {
	if o == nil {
		var ret PageConnector
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *Page) GetPageOk() (*PageConnector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *Page) SetPage(v PageConnector) {
	o.Page = v
}

func (o Page) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Page) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["page"] = o.Page
	return toSerialize, nil
}

func (o *Page) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"page",
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

	varPage := _Page{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPage)

	if err != nil {
		return err
	}

	*o = Page(varPage)

	return err
}

type NullablePage struct {
	value *Page
	isSet bool
}

func (v NullablePage) Get() *Page {
	return v.value
}

func (v *NullablePage) Set(val *Page) {
	v.value = val
	v.isSet = true
}

func (v NullablePage) IsSet() bool {
	return v.isSet
}

func (v *NullablePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePage(val *Page) *NullablePage {
	return &NullablePage{value: val, isSet: true}
}

func (v NullablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


