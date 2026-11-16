# VariableGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Secret** | Pointer to **NullableBool** |  | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVariableGet

`func NewVariableGet(createdAt time.Time, updatedAt time.Time, ) *VariableGet`

NewVariableGet instantiates a new VariableGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableGetWithDefaults

`func NewVariableGetWithDefaults() *VariableGet`

NewVariableGetWithDefaults instantiates a new VariableGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *VariableGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VariableGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VariableGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VariableGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *VariableGet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableGet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableGet) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VariableGet) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *VariableGet) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *VariableGet) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *VariableGet) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableGet) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableGet) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableGet) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableGet) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableGet) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *VariableGet) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *VariableGet) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *VariableGet) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *VariableGet) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *VariableGet) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *VariableGet) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetLastEditor

`func (o *VariableGet) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *VariableGet) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *VariableGet) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *VariableGet) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariableGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariableGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariableGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VariableGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariableGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariableGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


