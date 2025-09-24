# WAFExceptionSpecificConditionOnNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | * &#x60;specific_body_form_field_name&#x60; - specific_body_form_field_name * &#x60;specific_http_header_name&#x60; - specific_http_header_name * &#x60;specific_query_string_name&#x60; - specific_query_string_name | 
**Name** | **string** |  | 

## Methods

### NewWAFExceptionSpecificConditionOnNameRequest

`func NewWAFExceptionSpecificConditionOnNameRequest(match string, name string, ) *WAFExceptionSpecificConditionOnNameRequest`

NewWAFExceptionSpecificConditionOnNameRequest instantiates a new WAFExceptionSpecificConditionOnNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFExceptionSpecificConditionOnNameRequestWithDefaults

`func NewWAFExceptionSpecificConditionOnNameRequestWithDefaults() *WAFExceptionSpecificConditionOnNameRequest`

NewWAFExceptionSpecificConditionOnNameRequestWithDefaults instantiates a new WAFExceptionSpecificConditionOnNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *WAFExceptionSpecificConditionOnNameRequest) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WAFExceptionSpecificConditionOnNameRequest) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WAFExceptionSpecificConditionOnNameRequest) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetName

`func (o *WAFExceptionSpecificConditionOnNameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WAFExceptionSpecificConditionOnNameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WAFExceptionSpecificConditionOnNameRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


