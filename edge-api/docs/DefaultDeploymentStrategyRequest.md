# DefaultDeploymentStrategyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeApplication** | **int64** |  | 
**EdgeFirewall** | Pointer to **NullableInt64** |  | [optional] 
**CustomPage** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewDefaultDeploymentStrategyRequest

`func NewDefaultDeploymentStrategyRequest(edgeApplication int64, ) *DefaultDeploymentStrategyRequest`

NewDefaultDeploymentStrategyRequest instantiates a new DefaultDeploymentStrategyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultDeploymentStrategyRequestWithDefaults

`func NewDefaultDeploymentStrategyRequestWithDefaults() *DefaultDeploymentStrategyRequest`

NewDefaultDeploymentStrategyRequestWithDefaults instantiates a new DefaultDeploymentStrategyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeApplication

`func (o *DefaultDeploymentStrategyRequest) GetEdgeApplication() int64`

GetEdgeApplication returns the EdgeApplication field if non-nil, zero value otherwise.

### GetEdgeApplicationOk

`func (o *DefaultDeploymentStrategyRequest) GetEdgeApplicationOk() (*int64, bool)`

GetEdgeApplicationOk returns a tuple with the EdgeApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplication

`func (o *DefaultDeploymentStrategyRequest) SetEdgeApplication(v int64)`

SetEdgeApplication sets EdgeApplication field to given value.


### GetEdgeFirewall

`func (o *DefaultDeploymentStrategyRequest) GetEdgeFirewall() int64`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *DefaultDeploymentStrategyRequest) GetEdgeFirewallOk() (*int64, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *DefaultDeploymentStrategyRequest) SetEdgeFirewall(v int64)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *DefaultDeploymentStrategyRequest) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### SetEdgeFirewallNil

`func (o *DefaultDeploymentStrategyRequest) SetEdgeFirewallNil(b bool)`

 SetEdgeFirewallNil sets the value for EdgeFirewall to be an explicit nil

### UnsetEdgeFirewall
`func (o *DefaultDeploymentStrategyRequest) UnsetEdgeFirewall()`

UnsetEdgeFirewall ensures that no value is present for EdgeFirewall, not even an explicit nil
### GetCustomPage

`func (o *DefaultDeploymentStrategyRequest) GetCustomPage() int64`

GetCustomPage returns the CustomPage field if non-nil, zero value otherwise.

### GetCustomPageOk

`func (o *DefaultDeploymentStrategyRequest) GetCustomPageOk() (*int64, bool)`

GetCustomPageOk returns a tuple with the CustomPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPage

`func (o *DefaultDeploymentStrategyRequest) SetCustomPage(v int64)`

SetCustomPage sets CustomPage field to given value.

### HasCustomPage

`func (o *DefaultDeploymentStrategyRequest) HasCustomPage() bool`

HasCustomPage returns a boolean if a field has been set.

### SetCustomPageNil

`func (o *DefaultDeploymentStrategyRequest) SetCustomPageNil(b bool)`

 SetCustomPageNil sets the value for CustomPage to be an explicit nil

### UnsetCustomPage
`func (o *DefaultDeploymentStrategyRequest) UnsetCustomPage()`

UnsetCustomPage ensures that no value is present for CustomPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


