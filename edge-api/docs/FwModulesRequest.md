# FwModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | Pointer to [**FwModuleRequest**](FwModuleRequest.md) |  | [optional] 
**NetworkProtection** | Pointer to [**FwModuleRequest**](FwModuleRequest.md) |  | [optional] 
**Waf** | Pointer to [**FwModuleRequest**](FwModuleRequest.md) |  | [optional] 

## Methods

### NewFwModulesRequest

`func NewFwModulesRequest() *FwModulesRequest`

NewFwModulesRequest instantiates a new FwModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwModulesRequestWithDefaults

`func NewFwModulesRequestWithDefaults() *FwModulesRequest`

NewFwModulesRequestWithDefaults instantiates a new FwModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *FwModulesRequest) GetFunctions() FwModuleRequest`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *FwModulesRequest) GetFunctionsOk() (*FwModuleRequest, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *FwModulesRequest) SetFunctions(v FwModuleRequest)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *FwModulesRequest) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetNetworkProtection

`func (o *FwModulesRequest) GetNetworkProtection() FwModuleRequest`

GetNetworkProtection returns the NetworkProtection field if non-nil, zero value otherwise.

### GetNetworkProtectionOk

`func (o *FwModulesRequest) GetNetworkProtectionOk() (*FwModuleRequest, bool)`

GetNetworkProtectionOk returns a tuple with the NetworkProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtection

`func (o *FwModulesRequest) SetNetworkProtection(v FwModuleRequest)`

SetNetworkProtection sets NetworkProtection field to given value.

### HasNetworkProtection

`func (o *FwModulesRequest) HasNetworkProtection() bool`

HasNetworkProtection returns a boolean if a field has been set.

### GetWaf

`func (o *FwModulesRequest) GetWaf() FwModuleRequest`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *FwModulesRequest) GetWafOk() (*FwModuleRequest, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *FwModulesRequest) SetWaf(v FwModuleRequest)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *FwModulesRequest) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


