# PatchFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]FwCriterionFieldRequest**]([]FwCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]FwBehaviorsRequest**](FwBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchFirewallRuleRequest

`func NewPatchFirewallRuleRequest() *PatchFirewallRuleRequest`

NewPatchFirewallRuleRequest instantiates a new PatchFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFirewallRuleRequestWithDefaults

`func NewPatchFirewallRuleRequestWithDefaults() *PatchFirewallRuleRequest`

NewPatchFirewallRuleRequestWithDefaults instantiates a new PatchFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchFirewallRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchFirewallRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchFirewallRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchFirewallRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchFirewallRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchFirewallRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchFirewallRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchFirewallRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchFirewallRuleRequest) GetCriteria() [][]FwCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchFirewallRuleRequest) GetCriteriaOk() (*[][]FwCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchFirewallRuleRequest) SetCriteria(v [][]FwCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchFirewallRuleRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchFirewallRuleRequest) GetBehaviors() []FwBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchFirewallRuleRequest) GetBehaviorsOk() (*[]FwBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchFirewallRuleRequest) SetBehaviors(v []FwBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchFirewallRuleRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchFirewallRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchFirewallRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchFirewallRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchFirewallRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


