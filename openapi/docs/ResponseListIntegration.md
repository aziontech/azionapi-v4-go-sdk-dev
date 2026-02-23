# ResponseListIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Scope** | **string** |  | 
**ScopeType** | **string** |  | 
**Provider** | [**Platform**](Platform.md) |  | 

## Methods

### NewResponseListIntegration

`func NewResponseListIntegration(id int64, scope string, scopeType string, provider Platform, ) *ResponseListIntegration`

NewResponseListIntegration instantiates a new ResponseListIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListIntegrationWithDefaults

`func NewResponseListIntegrationWithDefaults() *ResponseListIntegration`

NewResponseListIntegrationWithDefaults instantiates a new ResponseListIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListIntegration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListIntegration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListIntegration) SetId(v int64)`

SetId sets Id field to given value.


### GetScope

`func (o *ResponseListIntegration) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ResponseListIntegration) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ResponseListIntegration) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeType

`func (o *ResponseListIntegration) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *ResponseListIntegration) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *ResponseListIntegration) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetProvider

`func (o *ResponseListIntegration) GetProvider() Platform`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResponseListIntegration) GetProviderOk() (*Platform, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResponseListIntegration) SetProvider(v Platform)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


