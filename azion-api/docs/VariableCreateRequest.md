# VariableCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewVariableCreateRequest

`func NewVariableCreateRequest(key string, ) *VariableCreateRequest`

NewVariableCreateRequest instantiates a new VariableCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCreateRequestWithDefaults

`func NewVariableCreateRequestWithDefaults() *VariableCreateRequest`

NewVariableCreateRequestWithDefaults instantiates a new VariableCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableCreateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableCreateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableCreateRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableCreateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableCreateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableCreateRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableCreateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSecret

`func (o *VariableCreateRequest) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *VariableCreateRequest) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *VariableCreateRequest) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *VariableCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *VariableCreateRequest) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *VariableCreateRequest) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


