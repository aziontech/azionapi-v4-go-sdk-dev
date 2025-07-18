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

// checks if the PatchedWAFRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWAFRuleRequest{}

// PatchedWAFRuleRequest struct for PatchedWAFRuleRequest
type PatchedWAFRuleRequest struct {
	// 0 - All Rules<br> 1 - Validation of protocol compliance: weird request, unable to parse<br> 2 - Request too big, stored on disk and not parsed<br> 10 - Validation of protocol compliance: invalid HEX encoding (null bytes)<br> 11 - Validation of protocol compliance: missing or unknown Content-Type header in a POST (this rule applies only to Request Body match zone)<br> 12 - Validation of protocol compliance: invalid formatted URL<br> 13 - Validation of protocol compliance: invalid POST format<br> 14 - Validation of protocol compliance: invalid POST boundary<br> 15 - Validation of protocol compliance: invalid JSON<br> 16 - Validation of protocol compliance: POST with no body<br> 17 - Possible SQL Injection attack: validation with libinjection_sql<br> 18 - Possible XSS attack: validation with libinjection_xss<br> 1000 - Possible SQL Injection attack: SQL keywords found in Body, Path, Query String or Cookies<br> 1001 - Possible SQL Injection or XSS attack: double quote (\") found in Body, Path, Query String or Cookies<br> 1002 - Possible SQL Injection attack: possible hex encoding (0x) found in Body, Path, Query String or Cookies<br> 1003 - Possible SQL Injection attack: MySQL comment (/_*) found in Body, Path, Query String or Cookies<br> 1004 - Possible SQL Injection attack: MySQL comment (*_/) found in Body, Path, Query String or Cookies<br> 1005 - Possible SQL Injection attack: MySQL keyword (|) found in Body, Path, Query String or Cookies<br> 1006 - Possible SQL Injection attack: MySQL keyword (&&) found in Body, Path, Query String or Cookies<br> 1007 - Possible SQL Injection attack: MySQL comment (--) found in Body, Path, Query String or Cookies<br> 1008 - Possible SQL Injection or XSS attack: semicolon (;) found in Body, Path or Query String<br> 1009 - Possible SQL Injection attack: equal sign (=) found in Body or Query String<br> 1010 - Possible SQL Injection or XSS attack: open parenthesis [(] found in Body, Path, Query String or Cookies<br> 1011 - Possible SQL Injection or XSS attack: close parenthesis [)] found in Body, Path, Query String or Cookies<br> 1013 - Possible SQL Injection or XSS attack: apostrophe (') found in Body, Path, Query String or Cookies<br> 1015 - Possible SQL Injection attack: comma (,) found in Body, Path, Query String or Cookies<br> 1016 - Possible SQL Injection attack: MySQL comment (#) found in Body, Path, Query String or Cookies<br> 1017 - Possible SQL Injection attack: double at sign (@@) found in Body, Path, Query String or Cookies<br> 1100 - Possible RFI attack: scheme \"http://\" found in Body, Query String or Cookies<br> 1101 - Possible RFI attack: scheme \"https://\" found in Body, Query String or Cookies<br> 1102 - Possible RFI attack: scheme \"ftp://\" found in Body, Query String or Cookies<br> 1103 - Possible RFI attack: scheme \"php://\" found in Body, Query String or Cookies<br> 1104 - Possible RFI attack: scheme \"sftp://\" found in Body, Query String or Cookies<br> 1105 - Possible RFI attack: scheme \"zlib://\" found in Body, Query String or Cookies<br> 1106 - Possible RFI attack: scheme \"data://\" found in Body, Query String or Cookies<br> 1107 - Possible RFI attack: scheme \"glob://\" found in Body, Query String or Cookies<br> 1108 - Possible RFI attack: scheme \"phar://\" found in Body, Query String or Cookies<br> 1109 - Possible RFI attack: scheme \"file://\" found in Body, Query String or Cookies<br> 1110 - Possible RFI attack: scheme \"gopher://\" found in Body, Query String or Cookies<br> 1198 - Possible RCE attack: validation with log4j (Log4Shell) in HEADERS_VAR<br> 1199 - Possible RCE attack: validation with log4j (Log4Shell) in Body, Path, Query String, Headers or Cookies<br> 1200 - Possible Directory Traversal attack: double dot (..) found in Body, Path, Query String or Cookies<br> 1202 - Possible Directory Traversal attack: obvious probe (/etc/passwd) found in Body, Path, Query String or Cookies<br> 1203 - Possible Directory Traversal attack: obvious windows path (c:\\) found in Body, Path, Query String or Cookies<br> 1204 - Possible Directory Traversal attack: obvious probe (cmd.exe) found in Body, Path, Query String or Cookies<br> 1205 - Possible Directory Traversal attack: backslash (\\) found in Body, Path, Query String or Cookies<br> 1206 - Possible Directory Traversal attack: slash (/) found in Body, Query String or Cookies<br> 1207 - Possible Directory Traversal attack: obvious path probe (/..;/) found in Body, Query String or Cookies<br> 1208 - Possible Directory Traversal attack: obvious path probe (/.;/) found in Body, Query String or Cookies<br> 1209 - Possible Directory Traversal attack: obvious path probe (/.%2e/) found in Body, Query String or Cookies<br> 1210 - Possible Directory Traversal attack: obvious path probe (/%2e./) found in Body, Query String or Cookies<br> 1302 - Possible XSS attack: html open tag (<) found in Body, Path, Query String or Cookies<br> 1303 - Possible XSS attack: html close tag (>) found in Body, Path, Query String or Cookies<br> 1310 - Possible XSS attack: open square bracket ([) found in Body, Path, Query String or Cookies<br> 1311 - Possible XSS attack: close square bracket (]) found in Body, Path, Query String or Cookies<br> 1312 - Possible XSS attack: tilde character (~) found in Body, Path, Query String or Cookies<br> 1314 - Possible XSS attack: back quote ( `) found in Body, Path, Query String or Cookies<br> 1315 - Possible XSS attack: double encoding (%[2|3]) found in Body, Path, Query String or Cookies<br> 1400 - Possible trick to evade protection: UTF7/8 encoding (&#) found in Body, Path, Query String or Cookies<br> 1401 - Possible trick to evade protection: MS encoding (%U) found in Body, Path, Query String or Cookies<br> 1402 - Possible trick to evade protection: encoded chars (%20-%3F) found in Body, Query String or Cookies<br> 1500 - Possible File Upload attempt: asp/php (.ph, .asp or .ht) found in filename in a multipart POST containing a file<br> 2001 - Possible CVE-2022-22965 attack: Tomcat Pipeline Context tampering  * `0` - All Rules * `1` - Validation of protocol compliance: weird request, unable to parse * `2` - Request too big, stored on disk and not parsed * `10` - Validation of protocol compliance: invalid HEX encoding (null bytes) * `11` - Validation of protocol compliance: missing or unknown Content-Type header in a POST (this rule applies only to Request Body match zone) * `12` - Validation of protocol compliance: invalid formatted URL * `13` - Validation of protocol compliance: invalid POST format * `14` - Validation of protocol compliance: invalid POST boundary * `15` - Validation of protocol compliance: invalid JSON * `16` - Validation of protocol compliance: POST with no body * `17` - Possible SQL Injection attack: validation with libinjection_sql * `18` - Possible XSS attack: validation with libinjection_xss * `1000` - Possible SQL Injection attack: SQL keywords found in Body, Path, Query String or Cookies * `1001` - Possible SQL Injection or XSS attack: double quote (\") found in Body, Path, Query String or Cookies * `1002` - Possible SQL Injection attack: possible hex encoding (0x) found in Body, Path, Query String or Cookies * `1003` - Possible SQL Injection attack: MySQL comment (/_*) found in Body, Path, Query String or Cookies * `1004` - Possible SQL Injection attack: MySQL comment (*_/) found in Body, Path, Query String or Cookies * `1005` - Possible SQL Injection attack: MySQL keyword (|) found in Body, Path, Query String or Cookies * `1006` - Possible SQL Injection attack: MySQL keyword (&&) found in Body, Path, Query String or Cookies * `1007` - Possible SQL Injection attack: MySQL comment (--) found in Body, Path, Query String or Cookies * `1008` - Possible SQL Injection or XSS attack: semicolon (;) found in Body, Path or Query String * `1009` - Possible SQL Injection attack: equal sign (=) found in Body or Query String * `1010` - Possible SQL Injection or XSS attack: open parenthesis [(] found in Body, Path, Query String or Cookies * `1011` - Possible SQL Injection or XSS attack: close parenthesis [)] found in Body, Path, Query String or Cookies * `1013` - Possible SQL Injection or XSS attack: apostrophe (') found in Body, Path, Query String or Cookies * `1015` - Possible SQL Injection attack: comma (,) found in Body, Path, Query String or Cookies * `1016` - Possible SQL Injection attack: MySQL comment (#) found in Body, Path, Query String or Cookies * `1017` - Possible SQL Injection attack: double at sign (@@) found in Body, Path, Query String or Cookies * `1100` - Possible RFI attack: scheme \"http://\" found in Body, Query String or Cookies * `1101` - Possible RFI attack: scheme \"https://\" found in Body, Query String or Cookies * `1102` - Possible RFI attack: scheme \"ftp://\" found in Body, Query String or Cookies * `1103` - Possible RFI attack: scheme \"php://\" found in Body, Query String or Cookies * `1104` - Possible RFI attack: scheme \"sftp://\" found in Body, Query String or Cookies * `1105` - Possible RFI attack: scheme \"zlib://\" found in Body, Query String or Cookies * `1106` - Possible RFI attack: scheme \"data://\" found in Body, Query String or Cookies * `1107` - Possible RFI attack: scheme \"glob://\" found in Body, Query String or Cookies * `1108` - Possible RFI attack: scheme \"phar://\" found in Body, Query String or Cookies * `1109` - Possible RFI attack: scheme \"file://\" found in Body, Query String or Cookies * `1110` - Possible RFI attack: scheme \"gopher://\" found in Body, Query String or Cookies * `1198` - Possible RCE attack: validation with log4j (Log4Shell) in HEADERS_VAR * `1199` - Possible RCE attack: validation with log4j (Log4Shell) in Body, Path, Query String, Headers or Cookies * `1200` - Possible Directory Traversal attack: double dot (..) found in Body, Path, Query String or Cookies * `1202` - Possible Directory Traversal attack: obvious probe (/etc/passwd) found in Body, Path, Query String or Cookies * `1203` - Possible Directory Traversal attack: obvious windows path (c:\\) found in Body, Path, Query String or Cookies * `1204` - Possible Directory Traversal attack: obvious probe (cmd.exe) found in Body, Path, Query String or Cookies * `1205` - Possible Directory Traversal attack: backslash (\\) found in Body, Path, Query String or Cookies * `1206` - Possible Directory Traversal attack: slash (/) found in Body, Query String or Cookies * `1207` - Possible Directory Traversal attack: obvious path probe (/..;/) found in Body, Query String or Cookies * `1208` - Possible Directory Traversal attack: obvious path probe (/.;/) found in Body, Query String or Cookies * `1209` - Possible Directory Traversal attack: obvious path probe (/.%2e/) found in Body, Query String or Cookies * `1210` - Possible Directory Traversal attack: obvious path probe (/%2e./) found in Body, Query String or Cookies * `1302` - Possible XSS attack: html open tag (<) found in Body, Path, Query String or Cookies * `1303` - Possible XSS attack: html close tag (>) found in Body, Path, Query String or Cookies * `1310` - Possible XSS attack: open square bracket ([) found in Body, Path, Query String or Cookies * `1311` - Possible XSS attack: close square bracket (]) found in Body, Path, Query String or Cookies * `1312` - Possible XSS attack: tilde character (~) found in Body, Path, Query String or Cookies * `1314` - Possible XSS attack: back quote ( `) found in Body, Path, Query String or Cookies * `1315` - Possible XSS attack: double encoding (%[2|3]) found in Body, Path, Query String or Cookies * `1400` - Possible trick to evade protection: UTF7/8 encoding (&#) found in Body, Path, Query String or Cookies * `1401` - Possible trick to evade protection: MS encoding (%U) found in Body, Path, Query String or Cookies * `1402` - Possible trick to evade protection: encoded chars (%20-%3F) found in Body, Query String or Cookies * `1500` - Possible File Upload attempt: asp/php (.ph, .asp or .ht) found in filename in a multipart POST containing a file * `2001` - Possible CVE-2022-22965 attack: Tomcat Pipeline Context tampering
	RuleId *int64 `json:"rule_id,omitempty"`
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	Path NullableString `json:"path,omitempty" validate:"regexp=^[^|]*$"`
	Conditions []WAFExceptionPolymorphicConditionRequest `json:"conditions,omitempty"`
	// * `regex` - regex * `contains` - contains
	Operator *string `json:"operator,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewPatchedWAFRuleRequest instantiates a new PatchedWAFRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWAFRuleRequest() *PatchedWAFRuleRequest {
	this := PatchedWAFRuleRequest{}
	return &this
}

// NewPatchedWAFRuleRequestWithDefaults instantiates a new PatchedWAFRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWAFRuleRequestWithDefaults() *PatchedWAFRuleRequest {
	this := PatchedWAFRuleRequest{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *PatchedWAFRuleRequest) GetRuleId() int64 {
	if o == nil || IsNil(o.RuleId) {
		var ret int64
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRuleRequest) GetRuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given int64 and assigns it to the RuleId field.
func (o *PatchedWAFRuleRequest) SetRuleId(v int64) {
	o.RuleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWAFRuleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWAFRuleRequest) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWAFRuleRequest) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWAFRuleRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *PatchedWAFRuleRequest) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *PatchedWAFRuleRequest) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *PatchedWAFRuleRequest) UnsetPath() {
	o.Path.Unset()
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PatchedWAFRuleRequest) GetConditions() []WAFExceptionPolymorphicConditionRequest {
	if o == nil || IsNil(o.Conditions) {
		var ret []WAFExceptionPolymorphicConditionRequest
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRuleRequest) GetConditionsOk() ([]WAFExceptionPolymorphicConditionRequest, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []WAFExceptionPolymorphicConditionRequest and assigns it to the Conditions field.
func (o *PatchedWAFRuleRequest) SetConditions(v []WAFExceptionPolymorphicConditionRequest) {
	o.Conditions = v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *PatchedWAFRuleRequest) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRuleRequest) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *PatchedWAFRuleRequest) SetOperator(v string) {
	o.Operator = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedWAFRuleRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRuleRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedWAFRuleRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedWAFRuleRequest) SetActive(v bool) {
	o.Active = &v
}

func (o PatchedWAFRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWAFRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullablePatchedWAFRuleRequest struct {
	value *PatchedWAFRuleRequest
	isSet bool
}

func (v NullablePatchedWAFRuleRequest) Get() *PatchedWAFRuleRequest {
	return v.value
}

func (v *NullablePatchedWAFRuleRequest) Set(val *PatchedWAFRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWAFRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWAFRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWAFRuleRequest(val *PatchedWAFRuleRequest) *NullablePatchedWAFRuleRequest {
	return &NullablePatchedWAFRuleRequest{value: val, isSet: true}
}

func (v NullablePatchedWAFRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWAFRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


