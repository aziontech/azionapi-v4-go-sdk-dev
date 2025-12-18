# ResponseFirewallFuncInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**FwFuncInstance**](FwFuncInstance.md) |  | 

## Methods

### NewResponseFirewallFuncInstance

`func NewResponseFirewallFuncInstance(data FwFuncInstance, ) *ResponseFirewallFuncInstance`

NewResponseFirewallFuncInstance instantiates a new ResponseFirewallFuncInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFirewallFuncInstanceWithDefaults

`func NewResponseFirewallFuncInstanceWithDefaults() *ResponseFirewallFuncInstance`

NewResponseFirewallFuncInstanceWithDefaults instantiates a new ResponseFirewallFuncInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseFirewallFuncInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseFirewallFuncInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseFirewallFuncInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseFirewallFuncInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseFirewallFuncInstance) GetData() FwFuncInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseFirewallFuncInstance) GetDataOk() (*FwFuncInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseFirewallFuncInstance) SetData(v FwFuncInstance)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


