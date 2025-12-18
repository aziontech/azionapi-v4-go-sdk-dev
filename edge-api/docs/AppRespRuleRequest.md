# AppRespRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]AppCriterionFieldRequest**]([]AppCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]RespBehaviorsRequest**](RespBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppRespRuleRequest

`func NewAppRespRuleRequest(name string, criteria [][]AppCriterionFieldRequest, behaviors []RespBehaviorsRequest, ) *AppRespRuleRequest`

NewAppRespRuleRequest instantiates a new AppRespRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRespRuleRequestWithDefaults

`func NewAppRespRuleRequestWithDefaults() *AppRespRuleRequest`

NewAppRespRuleRequestWithDefaults instantiates a new AppRespRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppRespRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRespRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRespRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *AppRespRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AppRespRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AppRespRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AppRespRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *AppRespRuleRequest) GetCriteria() [][]AppCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AppRespRuleRequest) GetCriteriaOk() (*[][]AppCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AppRespRuleRequest) SetCriteria(v [][]AppCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *AppRespRuleRequest) GetBehaviors() []RespBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *AppRespRuleRequest) GetBehaviorsOk() (*[]RespBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *AppRespRuleRequest) SetBehaviors(v []RespBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *AppRespRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRespRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRespRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRespRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


