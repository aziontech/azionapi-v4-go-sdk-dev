# EdgeConnectorHTTPTypeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | Pointer to **[]string** |  | [optional] 
**Host** | **NullableString** |  | 
**Path** | **string** |  | 
**FollowingRedirect** | Pointer to **bool** |  | [optional] 
**RealIpHeader** | **string** |  | 
**RealPortHeader** | **string** |  | 

## Methods

### NewEdgeConnectorHTTPTypeProperties

`func NewEdgeConnectorHTTPTypeProperties(host NullableString, path string, realIpHeader string, realPortHeader string, ) *EdgeConnectorHTTPTypeProperties`

NewEdgeConnectorHTTPTypeProperties instantiates a new EdgeConnectorHTTPTypeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorHTTPTypePropertiesWithDefaults

`func NewEdgeConnectorHTTPTypePropertiesWithDefaults() *EdgeConnectorHTTPTypeProperties`

NewEdgeConnectorHTTPTypePropertiesWithDefaults instantiates a new EdgeConnectorHTTPTypeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *EdgeConnectorHTTPTypeProperties) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *EdgeConnectorHTTPTypeProperties) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *EdgeConnectorHTTPTypeProperties) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *EdgeConnectorHTTPTypeProperties) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetHost

`func (o *EdgeConnectorHTTPTypeProperties) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EdgeConnectorHTTPTypeProperties) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EdgeConnectorHTTPTypeProperties) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *EdgeConnectorHTTPTypeProperties) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *EdgeConnectorHTTPTypeProperties) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPath

`func (o *EdgeConnectorHTTPTypeProperties) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EdgeConnectorHTTPTypeProperties) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EdgeConnectorHTTPTypeProperties) SetPath(v string)`

SetPath sets Path field to given value.


### GetFollowingRedirect

`func (o *EdgeConnectorHTTPTypeProperties) GetFollowingRedirect() bool`

GetFollowingRedirect returns the FollowingRedirect field if non-nil, zero value otherwise.

### GetFollowingRedirectOk

`func (o *EdgeConnectorHTTPTypeProperties) GetFollowingRedirectOk() (*bool, bool)`

GetFollowingRedirectOk returns a tuple with the FollowingRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingRedirect

`func (o *EdgeConnectorHTTPTypeProperties) SetFollowingRedirect(v bool)`

SetFollowingRedirect sets FollowingRedirect field to given value.

### HasFollowingRedirect

`func (o *EdgeConnectorHTTPTypeProperties) HasFollowingRedirect() bool`

HasFollowingRedirect returns a boolean if a field has been set.

### GetRealIpHeader

`func (o *EdgeConnectorHTTPTypeProperties) GetRealIpHeader() string`

GetRealIpHeader returns the RealIpHeader field if non-nil, zero value otherwise.

### GetRealIpHeaderOk

`func (o *EdgeConnectorHTTPTypeProperties) GetRealIpHeaderOk() (*string, bool)`

GetRealIpHeaderOk returns a tuple with the RealIpHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealIpHeader

`func (o *EdgeConnectorHTTPTypeProperties) SetRealIpHeader(v string)`

SetRealIpHeader sets RealIpHeader field to given value.


### GetRealPortHeader

`func (o *EdgeConnectorHTTPTypeProperties) GetRealPortHeader() string`

GetRealPortHeader returns the RealPortHeader field if non-nil, zero value otherwise.

### GetRealPortHeaderOk

`func (o *EdgeConnectorHTTPTypeProperties) GetRealPortHeaderOk() (*string, bool)`

GetRealPortHeaderOk returns a tuple with the RealPortHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPortHeader

`func (o *EdgeConnectorHTTPTypeProperties) SetRealPortHeader(v string)`

SetRealPortHeader sets RealPortHeader field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


