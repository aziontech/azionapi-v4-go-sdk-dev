# ResponseWRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**WRule**](WRule.md) |  | 

## Methods

### NewResponseWRule

`func NewResponseWRule(data WRule, ) *ResponseWRule`

NewResponseWRule instantiates a new ResponseWRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWRuleWithDefaults

`func NewResponseWRuleWithDefaults() *ResponseWRule`

NewResponseWRuleWithDefaults instantiates a new ResponseWRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseWRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseWRule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseWRule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseWRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseWRule) GetData() WRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseWRule) GetDataOk() (*WRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseWRule) SetData(v WRule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


