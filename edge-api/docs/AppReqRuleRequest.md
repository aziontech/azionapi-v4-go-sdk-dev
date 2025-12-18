# AppReqRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]AppCriterionFieldRequest**]([]AppCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]ReqBehaviorsRequest**](ReqBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppReqRuleRequest

`func NewAppReqRuleRequest(name string, criteria [][]AppCriterionFieldRequest, behaviors []ReqBehaviorsRequest, ) *AppReqRuleRequest`

NewAppReqRuleRequest instantiates a new AppReqRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppReqRuleRequestWithDefaults

`func NewAppReqRuleRequestWithDefaults() *AppReqRuleRequest`

NewAppReqRuleRequestWithDefaults instantiates a new AppReqRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppReqRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppReqRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppReqRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *AppReqRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AppReqRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AppReqRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AppReqRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *AppReqRuleRequest) GetCriteria() [][]AppCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AppReqRuleRequest) GetCriteriaOk() (*[][]AppCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AppReqRuleRequest) SetCriteria(v [][]AppCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *AppReqRuleRequest) GetBehaviors() []ReqBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *AppReqRuleRequest) GetBehaviorsOk() (*[]ReqBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *AppReqRuleRequest) SetBehaviors(v []ReqBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *AppReqRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppReqRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppReqRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppReqRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


