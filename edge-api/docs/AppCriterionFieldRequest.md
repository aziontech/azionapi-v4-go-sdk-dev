# AppCriterionFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditional** | **string** | * &#x60;if&#x60; - if * &#x60;or&#x60; - or * &#x60;and&#x60; - and | 
**Variable** | **string** | * &#x60;${arg_&lt;name&gt;}&#x60; - ${arg_&lt;name&gt;} * &#x60;${args}&#x60; - ${args} * &#x60;${cookie_&lt;name&gt;}&#x60; - ${cookie_&lt;name&gt;} * &#x60;${device_group}&#x60; - ${device_group} * &#x60;${geoip_city_continent_code}&#x60; - ${geoip_city_continent_code} * &#x60;${geoip_city_country_code}&#x60; - ${geoip_city_country_code} * &#x60;${geoip_city_country_name}&#x60; - ${geoip_city_country_name} * &#x60;${geoip_city}&#x60; - ${geoip_city} * &#x60;${geoip_continent_code}&#x60; - ${geoip_continent_code} * &#x60;${geoip_country_code}&#x60; - ${geoip_country_code} * &#x60;${geoip_country_name}&#x60; - ${geoip_country_name} * &#x60;${geoip_region_name}&#x60; - ${geoip_region_name} * &#x60;${geoip_region}&#x60; - ${geoip_region} * &#x60;${host}&#x60; - ${host} * &#x60;${domain}&#x60; - ${domain} * &#x60;${http_&lt;header_name&gt;}&#x60; - ${http_&lt;header_name&gt;} * &#x60;${remote_addr}&#x60; - ${remote_addr} * &#x60;${remote_user}&#x60; - ${remote_user} * &#x60;${request_method}&#x60; - ${request_method} * &#x60;${request_uri}&#x60; - ${request_uri} * &#x60;${request}&#x60; - ${request} * &#x60;${scheme}&#x60; - ${scheme} * &#x60;${sent_http_&lt;header_name&gt;}&#x60; - ${sent_http_&lt;header_name&gt;} * &#x60;${status}&#x60; - ${status} * &#x60;${upstream_addr}&#x60; - ${upstream_addr} * &#x60;${upstream_cookie_&lt;name&gt;}&#x60; - ${upstream_cookie_&lt;name&gt;} * &#x60;${upstream_http_&lt;header_name&gt;}&#x60; - ${upstream_http_&lt;header_name&gt;} * &#x60;${upstream_status}&#x60; - ${upstream_status} * &#x60;${uri}&#x60; - ${uri} * &#x60;${server_addr}&#x60; - ${server_addr} * &#x60;${server_port}&#x60; - ${server_port} * &#x60;${ssl_client_cert}&#x60; - ${ssl_client_cert} * &#x60;${ssl_client_escaped_cert}&#x60; - ${ssl_client_escaped_cert} * &#x60;${ssl_client_fingerprint}&#x60; - ${ssl_client_fingerprint} * &#x60;${ssl_client_i_dn}&#x60; - ${ssl_client_i_dn} * &#x60;${ssl_client_s_dn_parsed}&#x60; - ${ssl_client_s_dn_parsed} * &#x60;${ssl_client_s_dn}&#x60; - ${ssl_client_s_dn} * &#x60;${ssl_client_serial}&#x60; - ${ssl_client_serial} * &#x60;${ssl_client_v_end}&#x60; - ${ssl_client_v_end} * &#x60;${ssl_client_v_remain}&#x60; - ${ssl_client_v_remain} * &#x60;${ssl_client_v_start}&#x60; - ${ssl_client_v_start} * &#x60;${ssl_client_verify}&#x60; - ${ssl_client_verify} * &#x60;${tcpinfo_rtt}&#x60; - ${tcpinfo_rtt} * &#x60;${remote_port}&#x60; - ${remote_port} * &#x60;${request_body}&#x60; - ${request_body} | 
**Operator** | **string** | * &#x60;does_not_exist&#x60; - does_not_exist * &#x60;does_not_match&#x60; - does_not_match * &#x60;does_not_start_with&#x60; - does_not_start_with * &#x60;exists&#x60; - exists * &#x60;is_equal&#x60; - is_equal * &#x60;is_in_list&#x60; - is_in_list * &#x60;is_not_equal&#x60; - is_not_equal * &#x60;is_not_in_list&#x60; - is_not_in_list * &#x60;matches&#x60; - matches * &#x60;starts_with&#x60; - starts_with | 
**Argument** | Pointer to [**NullableAppCriterionArgumentRequest**](AppCriterionArgumentRequest.md) |  | [optional] 

## Methods

### NewAppCriterionFieldRequest

`func NewAppCriterionFieldRequest(conditional string, variable string, operator string, ) *AppCriterionFieldRequest`

NewAppCriterionFieldRequest instantiates a new AppCriterionFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCriterionFieldRequestWithDefaults

`func NewAppCriterionFieldRequestWithDefaults() *AppCriterionFieldRequest`

NewAppCriterionFieldRequestWithDefaults instantiates a new AppCriterionFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditional

`func (o *AppCriterionFieldRequest) GetConditional() string`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *AppCriterionFieldRequest) GetConditionalOk() (*string, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *AppCriterionFieldRequest) SetConditional(v string)`

SetConditional sets Conditional field to given value.


### GetVariable

`func (o *AppCriterionFieldRequest) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *AppCriterionFieldRequest) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *AppCriterionFieldRequest) SetVariable(v string)`

SetVariable sets Variable field to given value.


### GetOperator

`func (o *AppCriterionFieldRequest) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppCriterionFieldRequest) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppCriterionFieldRequest) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetArgument

`func (o *AppCriterionFieldRequest) GetArgument() AppCriterionArgumentRequest`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *AppCriterionFieldRequest) GetArgumentOk() (*AppCriterionArgumentRequest, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *AppCriterionFieldRequest) SetArgument(v AppCriterionArgumentRequest)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *AppCriterionFieldRequest) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *AppCriterionFieldRequest) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *AppCriterionFieldRequest) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


