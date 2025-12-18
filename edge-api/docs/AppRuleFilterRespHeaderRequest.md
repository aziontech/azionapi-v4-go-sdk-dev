# AppRuleFilterRespHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_response_header&#x60; - filter_response_header | 
**Attributes** | [**AppRuleFilterRespHeaderAttrsReq**](AppRuleFilterRespHeaderAttrsReq.md) |  | 

## Methods

### NewAppRuleFilterRespHeaderRequest

`func NewAppRuleFilterRespHeaderRequest(type_ string, attributes AppRuleFilterRespHeaderAttrsReq, ) *AppRuleFilterRespHeaderRequest`

NewAppRuleFilterRespHeaderRequest instantiates a new AppRuleFilterRespHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleFilterRespHeaderRequestWithDefaults

`func NewAppRuleFilterRespHeaderRequestWithDefaults() *AppRuleFilterRespHeaderRequest`

NewAppRuleFilterRespHeaderRequestWithDefaults instantiates a new AppRuleFilterRespHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppRuleFilterRespHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRuleFilterRespHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRuleFilterRespHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppRuleFilterRespHeaderRequest) GetAttributes() AppRuleFilterRespHeaderAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppRuleFilterRespHeaderRequest) GetAttributesOk() (*AppRuleFilterRespHeaderAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppRuleFilterRespHeaderRequest) SetAttributes(v AppRuleFilterRespHeaderAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


