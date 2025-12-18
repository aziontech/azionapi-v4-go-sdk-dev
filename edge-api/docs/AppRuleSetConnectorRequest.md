# AppRuleSetConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_connector&#x60; - set_connector | 
**Attributes** | [**AppRuleSetConnectorAttrsReq**](AppRuleSetConnectorAttrsReq.md) |  | 

## Methods

### NewAppRuleSetConnectorRequest

`func NewAppRuleSetConnectorRequest(type_ string, attributes AppRuleSetConnectorAttrsReq, ) *AppRuleSetConnectorRequest`

NewAppRuleSetConnectorRequest instantiates a new AppRuleSetConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleSetConnectorRequestWithDefaults

`func NewAppRuleSetConnectorRequestWithDefaults() *AppRuleSetConnectorRequest`

NewAppRuleSetConnectorRequestWithDefaults instantiates a new AppRuleSetConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleSetConnectorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleSetConnectorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleSetConnectorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleSetConnectorRequest) GetAttributes() AppRuleSetConnectorAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleSetConnectorRequest) GetAttributesOk() (*AppRuleSetConnectorAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleSetConnectorRequest) SetAttributes(v AppRuleSetConnectorAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


