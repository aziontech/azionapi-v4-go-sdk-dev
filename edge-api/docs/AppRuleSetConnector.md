# AppRuleSetConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_connector&#x60; - set_connector | 
**Attributes** | [**AppRuleSetConnectorAttrs**](AppRuleSetConnectorAttrs.md) |  | 

## Methods

### NewAppRuleSetConnector

`func NewAppRuleSetConnector(type_ string, attributes AppRuleSetConnectorAttrs, ) *AppRuleSetConnector`

NewAppRuleSetConnector instantiates a new AppRuleSetConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleSetConnectorWithDefaults

`func NewAppRuleSetConnectorWithDefaults() *AppRuleSetConnector`

NewAppRuleSetConnectorWithDefaults instantiates a new AppRuleSetConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleSetConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleSetConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleSetConnector) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleSetConnector) GetAttributes() AppRuleSetConnectorAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleSetConnector) GetAttributesOk() (*AppRuleSetConnectorAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleSetConnector) SetAttributes(v AppRuleSetConnectorAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


