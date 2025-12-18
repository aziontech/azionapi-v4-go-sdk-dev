# WEngineSettingsFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineVersion** | Pointer to **string** | * &#x60;2021-Q3&#x60; - 2021-Q3 | [optional] 
**Type** | Pointer to **string** | * &#x60;score&#x60; - score | [optional] 
**Attributes** | Pointer to [**WEngineSettingsAttributesFieldRequest**](WEngineSettingsAttributesFieldRequest.md) |  | [optional] 

## Methods

### NewWEngineSettingsFieldRequest

`func NewWEngineSettingsFieldRequest() *WEngineSettingsFieldRequest`

NewWEngineSettingsFieldRequest instantiates a new WEngineSettingsFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWEngineSettingsFieldRequestWithDefaults

`func NewWEngineSettingsFieldRequestWithDefaults() *WEngineSettingsFieldRequest`

NewWEngineSettingsFieldRequestWithDefaults instantiates a new WEngineSettingsFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineVersion

`func (o *WEngineSettingsFieldRequest) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *WEngineSettingsFieldRequest) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *WEngineSettingsFieldRequest) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *WEngineSettingsFieldRequest) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### GetType

`func (o *WEngineSettingsFieldRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WEngineSettingsFieldRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WEngineSettingsFieldRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WEngineSettingsFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WEngineSettingsFieldRequest) GetAttributes() WEngineSettingsAttributesFieldRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WEngineSettingsFieldRequest) GetAttributesOk() (*WEngineSettingsAttributesFieldRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WEngineSettingsFieldRequest) SetAttributes(v WEngineSettingsAttributesFieldRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WEngineSettingsFieldRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


