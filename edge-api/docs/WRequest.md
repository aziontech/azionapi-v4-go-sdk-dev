# WRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ProductVersion** | Pointer to **NullableString** |  | [optional] 
**EngineSettings** | Pointer to [**WEngineSettingsFieldRequest**](WEngineSettingsFieldRequest.md) |  | [optional] 

## Methods

### NewWRequest

`func NewWRequest(name string, ) *WRequest`

NewWRequest instantiates a new WRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWRequestWithDefaults

`func NewWRequestWithDefaults() *WRequest`

NewWRequestWithDefaults instantiates a new WRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *WRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProductVersion

`func (o *WRequest) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *WRequest) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *WRequest) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *WRequest) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### SetProductVersionNil

`func (o *WRequest) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *WRequest) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetEngineSettings

`func (o *WRequest) GetEngineSettings() WEngineSettingsFieldRequest`

GetEngineSettings returns the EngineSettings field if non-nil, zero value otherwise.

### GetEngineSettingsOk

`func (o *WRequest) GetEngineSettingsOk() (*WEngineSettingsFieldRequest, bool)`

GetEngineSettingsOk returns a tuple with the EngineSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineSettings

`func (o *WRequest) SetEngineSettings(v WEngineSettingsFieldRequest)`

SetEngineSettings sets EngineSettings field to given value.

### HasEngineSettings

`func (o *WRequest) HasEngineSettings() bool`

HasEngineSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


