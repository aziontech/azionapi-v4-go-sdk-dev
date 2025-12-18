# ResponseAsyncApplicationRespRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**AppRespRule**](AppRespRule.md) |  | 

## Methods

### NewResponseAsyncApplicationRespRule

`func NewResponseAsyncApplicationRespRule(data AppRespRule, ) *ResponseAsyncApplicationRespRule`

NewResponseAsyncApplicationRespRule instantiates a new ResponseAsyncApplicationRespRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncApplicationRespRuleWithDefaults

`func NewResponseAsyncApplicationRespRuleWithDefaults() *ResponseAsyncApplicationRespRule`

NewResponseAsyncApplicationRespRuleWithDefaults instantiates a new ResponseAsyncApplicationRespRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncApplicationRespRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncApplicationRespRule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncApplicationRespRule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncApplicationRespRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncApplicationRespRule) GetData() AppRespRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncApplicationRespRule) GetDataOk() (*AppRespRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncApplicationRespRule) SetData(v AppRespRule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


