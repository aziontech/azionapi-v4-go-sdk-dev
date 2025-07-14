# AddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the address is active for use | [optional] 
**Address** | **string** | IPv4/IPv6 address or CNAME to resolve | 
**HttpPort** | Pointer to **int64** | Port number for HTTP connections | [optional] 
**HttpsPort** | Pointer to **int64** | Port number for HTTPS connections | [optional] 
**Modules** | Pointer to [**NullableAddressModulesRequest**](AddressModulesRequest.md) |  | [optional] 

## Methods

### NewAddressRequest

`func NewAddressRequest(address string, ) *AddressRequest`

NewAddressRequest instantiates a new AddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRequestWithDefaults

`func NewAddressRequestWithDefaults() *AddressRequest`

NewAddressRequestWithDefaults instantiates a new AddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AddressRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddressRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddressRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddressRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress

`func (o *AddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetHttpPort

`func (o *AddressRequest) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *AddressRequest) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *AddressRequest) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *AddressRequest) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *AddressRequest) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *AddressRequest) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *AddressRequest) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *AddressRequest) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetModules

`func (o *AddressRequest) GetModules() AddressModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AddressRequest) GetModulesOk() (*AddressModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AddressRequest) SetModules(v AddressModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *AddressRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *AddressRequest) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *AddressRequest) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


