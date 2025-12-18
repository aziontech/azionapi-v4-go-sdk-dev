# PatchApplicationRespRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]AppCriterionFieldRequest**]([]AppCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]RespBehaviorsRequest**](RespBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchApplicationRespRuleRequest

`func NewPatchApplicationRespRuleRequest() *PatchApplicationRespRuleRequest`

NewPatchApplicationRespRuleRequest instantiates a new PatchApplicationRespRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationRespRuleRequestWithDefaults

`func NewPatchApplicationRespRuleRequestWithDefaults() *PatchApplicationRespRuleRequest`

NewPatchApplicationRespRuleRequestWithDefaults instantiates a new PatchApplicationRespRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApplicationRespRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApplicationRespRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApplicationRespRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApplicationRespRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchApplicationRespRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchApplicationRespRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchApplicationRespRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchApplicationRespRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchApplicationRespRuleRequest) GetCriteria() [][]AppCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchApplicationRespRuleRequest) GetCriteriaOk() (*[][]AppCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchApplicationRespRuleRequest) SetCriteria(v [][]AppCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchApplicationRespRuleRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchApplicationRespRuleRequest) GetBehaviors() []RespBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchApplicationRespRuleRequest) GetBehaviorsOk() (*[]RespBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchApplicationRespRuleRequest) SetBehaviors(v []RespBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchApplicationRespRuleRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchApplicationRespRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchApplicationRespRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchApplicationRespRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchApplicationRespRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


