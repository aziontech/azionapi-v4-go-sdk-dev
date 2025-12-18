# MToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access_token is a short-lived token (e.g., 10 minutes) used for API authentication in the Authorization: Bearer &lt;access_token&gt; header. | 
**TwoFactorRequired** | **bool** | Specifies the type of two-factor authentication configured. Currently supports &#39;TOTP&#39; (Time-based One-Time Password). | 
**TwoFactorType** | **string** | Indicates whether the user has an active TOTP device configured. If false, the user must register a device before using MFA authentication. | 
**HasActiveDevice** | **bool** | Indicates whether the user already has an active TOTP device configured.If false, the user needs to set up a new device before using MFA authentication. | 
**ResponseType** | **string** | Discriminator field for LoginResponse | 

## Methods

### NewMToken

`func NewMToken(accessToken string, twoFactorRequired bool, twoFactorType string, hasActiveDevice bool, responseType string, ) *MToken`

NewMToken instantiates a new MToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTokenWithDefaults

`func NewMTokenWithDefaults() *MToken`

NewMTokenWithDefaults instantiates a new MToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTwoFactorRequired

`func (o *MToken) GetTwoFactorRequired() bool`

GetTwoFactorRequired returns the TwoFactorRequired field if non-nil, zero value otherwise.

### GetTwoFactorRequiredOk

`func (o *MToken) GetTwoFactorRequiredOk() (*bool, bool)`

GetTwoFactorRequiredOk returns a tuple with the TwoFactorRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRequired

`func (o *MToken) SetTwoFactorRequired(v bool)`

SetTwoFactorRequired sets TwoFactorRequired field to given value.


### GetTwoFactorType

`func (o *MToken) GetTwoFactorType() string`

GetTwoFactorType returns the TwoFactorType field if non-nil, zero value otherwise.

### GetTwoFactorTypeOk

`func (o *MToken) GetTwoFactorTypeOk() (*string, bool)`

GetTwoFactorTypeOk returns a tuple with the TwoFactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorType

`func (o *MToken) SetTwoFactorType(v string)`

SetTwoFactorType sets TwoFactorType field to given value.


### GetHasActiveDevice

`func (o *MToken) GetHasActiveDevice() bool`

GetHasActiveDevice returns the HasActiveDevice field if non-nil, zero value otherwise.

### GetHasActiveDeviceOk

`func (o *MToken) GetHasActiveDeviceOk() (*bool, bool)`

GetHasActiveDeviceOk returns a tuple with the HasActiveDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveDevice

`func (o *MToken) SetHasActiveDevice(v bool)`

SetHasActiveDevice sets HasActiveDevice field to given value.


### GetResponseType

`func (o *MToken) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *MToken) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *MToken) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


