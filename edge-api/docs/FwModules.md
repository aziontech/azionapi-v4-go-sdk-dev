# FwModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdosProtection** | [**FwModule**](FwModule.md) |  | 
**Functions** | Pointer to [**FwModule**](FwModule.md) |  | [optional] 
**NetworkProtection** | Pointer to [**FwModule**](FwModule.md) |  | [optional] 
**Waf** | Pointer to [**FwModule**](FwModule.md) |  | [optional] 

## Methods

### NewFwModules

`func NewFwModules(ddosProtection FwModule, ) *FwModules`

NewFwModules instantiates a new FwModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwModulesWithDefaults

`func NewFwModulesWithDefaults() *FwModules`

NewFwModulesWithDefaults instantiates a new FwModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdosProtection

`func (o *FwModules) GetDdosProtection() FwModule`

GetDdosProtection returns the DdosProtection field if non-nil, zero value otherwise.

### GetDdosProtectionOk

`func (o *FwModules) GetDdosProtectionOk() (*FwModule, bool)`

GetDdosProtectionOk returns a tuple with the DdosProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosProtection

`func (o *FwModules) SetDdosProtection(v FwModule)`

SetDdosProtection sets DdosProtection field to given value.


### GetFunctions

`func (o *FwModules) GetFunctions() FwModule`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *FwModules) GetFunctionsOk() (*FwModule, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *FwModules) SetFunctions(v FwModule)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *FwModules) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetNetworkProtection

`func (o *FwModules) GetNetworkProtection() FwModule`

GetNetworkProtection returns the NetworkProtection field if non-nil, zero value otherwise.

### GetNetworkProtectionOk

`func (o *FwModules) GetNetworkProtectionOk() (*FwModule, bool)`

GetNetworkProtectionOk returns a tuple with the NetworkProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtection

`func (o *FwModules) SetNetworkProtection(v FwModule)`

SetNetworkProtection sets NetworkProtection field to given value.

### HasNetworkProtection

`func (o *FwModules) HasNetworkProtection() bool`

HasNetworkProtection returns a boolean if a field has been set.

### GetWaf

`func (o *FwModules) GetWaf() FwModule`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *FwModules) GetWafOk() (*FwModule, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *FwModules) SetWaf(v FwModule)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *FwModules) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


