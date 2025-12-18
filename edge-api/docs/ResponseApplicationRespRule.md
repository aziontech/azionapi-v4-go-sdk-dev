# ResponseApplicationRespRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**AppRespRule**](AppRespRule.md) |  | 

## Methods

### NewResponseApplicationRespRule

`func NewResponseApplicationRespRule(data AppRespRule, ) *ResponseApplicationRespRule`

NewResponseApplicationRespRule instantiates a new ResponseApplicationRespRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseApplicationRespRuleWithDefaults

`func NewResponseApplicationRespRuleWithDefaults() *ResponseApplicationRespRule`

NewResponseApplicationRespRuleWithDefaults instantiates a new ResponseApplicationRespRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseApplicationRespRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseApplicationRespRule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseApplicationRespRule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseApplicationRespRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseApplicationRespRule) GetData() AppRespRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseApplicationRespRule) GetDataOk() (*AppRespRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseApplicationRespRule) SetData(v AppRespRule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


