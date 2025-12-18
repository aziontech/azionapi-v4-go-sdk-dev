# PageConnectorAttrs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **int64** |  | 
**Ttl** | Pointer to **int64** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**CustomStatusCode** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewPageConnectorAttrs

`func NewPageConnectorAttrs(connector int64, ) *PageConnectorAttrs`

NewPageConnectorAttrs instantiates a new PageConnectorAttrs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageConnectorAttrsWithDefaults

`func NewPageConnectorAttrsWithDefaults() *PageConnectorAttrs`

NewPageConnectorAttrsWithDefaults instantiates a new PageConnectorAttrs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *PageConnectorAttrs) GetConnector() int64`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PageConnectorAttrs) GetConnectorOk() (*int64, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PageConnectorAttrs) SetConnector(v int64)`

SetConnector sets Connector field to given value.


### GetTtl

`func (o *PageConnectorAttrs) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PageConnectorAttrs) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PageConnectorAttrs) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PageConnectorAttrs) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUri

`func (o *PageConnectorAttrs) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PageConnectorAttrs) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PageConnectorAttrs) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PageConnectorAttrs) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *PageConnectorAttrs) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *PageConnectorAttrs) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetCustomStatusCode

`func (o *PageConnectorAttrs) GetCustomStatusCode() int64`

GetCustomStatusCode returns the CustomStatusCode field if non-nil, zero value otherwise.

### GetCustomStatusCodeOk

`func (o *PageConnectorAttrs) GetCustomStatusCodeOk() (*int64, bool)`

GetCustomStatusCodeOk returns a tuple with the CustomStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusCode

`func (o *PageConnectorAttrs) SetCustomStatusCode(v int64)`

SetCustomStatusCode sets CustomStatusCode field to given value.

### HasCustomStatusCode

`func (o *PageConnectorAttrs) HasCustomStatusCode() bool`

HasCustomStatusCode returns a boolean if a field has been set.

### SetCustomStatusCodeNil

`func (o *PageConnectorAttrs) SetCustomStatusCodeNil(b bool)`

 SetCustomStatusCodeNil sets the value for CustomStatusCode to be an explicit nil

### UnsetCustomStatusCode
`func (o *PageConnectorAttrs) UnsetCustomStatusCode()`

UnsetCustomStatusCode ensures that no value is present for CustomStatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


